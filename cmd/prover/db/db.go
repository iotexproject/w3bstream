package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	TaskID    common.Hash `gorm:"index:task_fetch,not null"`
	ProjectID uint64      `gorm:"index:task_fetch,not null"`
}

type DB struct {
	db *gorm.DB
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

func (p *DB) CreateTask(projectID uint64, taskID common.Hash) error {
	t := &task{
		TaskID:    taskID,
		ProjectID: projectID,
	}
	err := p.db.Create(t).Error
	return errors.Wrap(err, "failed to create task")
}

func (p *DB) DeleteTask(projectID uint64, taskID common.Hash) error {
	err := p.db.Where("task_id = ?", taskID).Where("project_id = ?", projectID).Delete(&task{}).Error
	return errors.Wrap(err, "failed to delete task")
}

func New(db *gorm.DB) (*DB, error) {
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db}, nil
}
