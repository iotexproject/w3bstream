package db

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type scannedBlockNumber struct {
	gorm.Model
	Number uint64 `gorm:"not null"`
}

type project struct {
	gorm.Model
	ProjectID uint64      `gorm:"uniqueIndex:project_id,not null"`
	URI       string      `gorm:"not null"`
	Hash      common.Hash `gorm:"not null"`
}

type projectFile struct {
	gorm.Model
	ProjectID uint64      `gorm:"uniqueIndex:project_id,not null"`
	File      []byte      `gorm:"not null"`
	Hash      common.Hash `gorm:"not null"`
}

type task struct {
	gorm.Model
	TaskID    common.Hash `gorm:"uniqueIndex:task_uniq,not null"`
	ProjectID uint64      `gorm:"uniqueIndex:task_uniq,not null"`
	Processed bool        `gorm:"index:unprocessed_task,not null,default:false"`
}

type DB struct {
	db     *gorm.DB
	prover common.Address
}

func (p *DB) ScannedBlockNumber() (uint64, error) {
	t := scannedBlockNumber{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to query scanned block number")
	}
	return t.Number, nil
}

func (p *DB) UpsertScannedBlockNumber(number uint64) error {
	t := scannedBlockNumber{
		Model: gorm.Model{
			ID: 1,
		},
		Number: number,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"number"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert scanned block number")
}

func (p *DB) Project(projectID uint64) (string, common.Hash, error) {
	t := project{}
	if err := p.db.Where("project_id = ?", projectID).First(&t).Error; err != nil {
		return "", common.Hash{}, errors.Wrap(err, "failed to query project")
	}
	return t.URI, t.Hash, nil
}

func (p *DB) UpsertProject(projectID uint64, uri string, hash common.Hash) error {
	t := project{
		ProjectID: projectID,
		URI:       uri,
		Hash:      hash,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"uri", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project")
}

func (p *DB) ProjectFile(projectID uint64) ([]byte, common.Hash, error) {
	t := projectFile{}
	if err := p.db.Where("project_id = ?", projectID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.Hash{}, nil
		}
		return nil, common.Hash{}, errors.Wrap(err, "failed to query project file")
	}
	return t.File, t.Hash, nil
}

func (p *DB) UpsertProjectFile(projectID uint64, file []byte, hash common.Hash) error {
	t := projectFile{
		ProjectID: projectID,
		File:      file,
		Hash:      hash,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"file", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project file")
}

func (p *DB) CreateTask(projectID uint64, taskID common.Hash, prover common.Address) error {
	if !bytes.Equal(prover[:], p.prover[:]) {
		return nil
	}
	t := &task{
		TaskID:    taskID,
		ProjectID: projectID,
		Processed: false,
	}
	err := p.db.Create(t).Error
	return errors.Wrap(err, "failed to create task")
}

func (p *DB) ProcessTask(projectID uint64, taskID common.Hash) error {
	t := &task{
		Processed: true,
	}
	err := p.db.Model(t).Where("task_id = ?", taskID).Where("project_id = ?", projectID).Updates(t).Error
	return errors.Wrap(err, "failed to update task")
}

func (p *DB) DeleteTask(projectID uint64, taskID common.Hash) error {
	err := p.db.Where("task_id = ?", taskID).Where("project_id = ?", projectID).Delete(&task{}).Error
	return errors.Wrap(err, "failed to delete task")
}

func (p *DB) UnprocessedTask() (uint64, common.Hash, error) {
	t := task{}
	if err := p.db.Order("created_at ASC").Where("processed = false").First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, common.Hash{}, nil
		}
		return 0, common.Hash{}, errors.Wrap(err, "failed to query unprocessed task")
	}
	return t.ProjectID, t.TaskID, nil
}

func New(localDBDir string, prover common.Address) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(localDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect sqlite")
	}
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}, &project{}, &projectFile{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db: db, prover: prover}, nil
}
