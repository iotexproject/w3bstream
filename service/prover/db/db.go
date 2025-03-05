package db

import (
	"bytes"
	"time"

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
	ProjectID string `gorm:"uniqueIndex:project_id_project,not null"`
	URI       string `gorm:"not null"`
	Hash      string `gorm:"not null"`
}

type projectFile struct {
	gorm.Model
	ProjectID string `gorm:"uniqueIndex:project_id_project_file,not null"`
	File      []byte `gorm:"not null"`
	Hash      string `gorm:"not null"`
}

type task struct {
	gorm.Model
	ProjectID string `gorm:"not null"`
	TaskID    string `gorm:"uniqueIndex:task_uniq,not null"`
	Processed bool   `gorm:"index:unprocessed_task,not null,default:false"`
	Error     string `gorm:"not null,default:''"`
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

func (p *DB) Project(projectID string) (string, common.Hash, error) {
	t := project{}
	if err := p.db.Where("project_id = ?", projectID).First(&t).Error; err != nil {
		return "", common.Hash{}, errors.Wrap(err, "failed to query project")
	}
	return t.URI, common.HexToHash(t.Hash), nil
}

func (p *DB) UpsertProject(projectID string, uri string, hash common.Hash) error {
	t := project{
		ProjectID: projectID,
		URI:       uri,
		Hash:      hash.Hex(),
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"uri", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project")
}

func (p *DB) ProjectFile(projectID string) ([]byte, common.Hash, error) {
	t := projectFile{}
	if err := p.db.Where("project_id = ?", projectID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.Hash{}, nil
		}
		return nil, common.Hash{}, errors.Wrap(err, "failed to query project file")
	}
	return t.File, common.HexToHash(t.Hash), nil
}

func (p *DB) UpsertProjectFile(projectID string, file []byte, hash common.Hash) error {
	t := projectFile{
		ProjectID: projectID,
		File:      file,
		Hash:      hash.Hex(),
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"file", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project file")
}

func (p *DB) CreateTask(taskID common.Hash, prover common.Address, projectID string) error {
	if !bytes.Equal(prover[:], p.prover[:]) {
		return nil
	}
	t := &task{
		ProjectID: projectID,
		TaskID:    taskID.Hex(),
		Processed: false,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoNothing: true,
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert task")
}

func (p *DB) ProcessTasks(taskIDs []common.Hash, err error) error {
	t := &task{
		Processed: true,
	}
	if err != nil {
		t.Error = err.Error()
	}
	idStrs := []string{}
	for _, t := range taskIDs {
		idStrs = append(idStrs, t.Hex())
	}
	err = p.db.Model(t).Where("task_id IN ?", idStrs).Updates(t).Error
	return errors.Wrap(err, "failed to update tasks")
}

func (p *DB) DeleteTask(taskID, tx common.Hash) error {
	err := p.db.Unscoped().Where("task_id = ?", taskID.Hex()).Delete(&task{}).Error
	return errors.Wrap(err, "failed to delete task")
}

func (p *DB) ProcessedTask(taskID common.Hash) (bool, string, time.Time, error) {
	t := task{}
	if err := p.db.Where("task_id = ?", taskID.Hex()).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, "", time.Now(), nil
		}
		return false, "", time.Time{}, errors.Wrap(err, "failed to query processed task")
	}
	return t.Processed, t.Error, t.CreatedAt, nil
}

func (p *DB) UnprocessedTasks(projectID string) ([]common.Hash, error) {
	ts := []*task{}
	if err := p.db.Order("created_at ASC").Where("processed = false").Where("project_id = ?", projectID).Find(&ts).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query unprocessed tasks")
	}
	hs := []common.Hash{}
	for _, t := range ts {
		hs = append(hs, common.HexToHash(t.TaskID))
	}
	return hs, nil
}

func (p *DB) UnprocessedProjects() (map[string]uint64, error) {
	res := []*struct {
		ProjectID string
		Total     uint64
	}{}
	if err := p.db.Model(&task{}).Order("created_at ASC").Select("project_id, count(*) as total").
		Where("processed = false").Group("project_id").Find(&res).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query unprocessed projects")
	}
	resM := map[string]uint64{}
	for _, r := range res {
		resM[r.ProjectID] = r.Total
	}
	return resM, nil
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
