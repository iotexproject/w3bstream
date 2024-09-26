package postgres

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
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"number"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert scanned block number")
	}
	return nil
}

func (p *DB) CreateTask(taskID common.Hash, projectID uint64) error {
	t := &task{
		TaskID:    taskID,
		ProjectID: projectID,
	}
	if err := p.db.Create(t).Error; err != nil {
		return errors.Wrap(err, "failed to create task")
	}
	return nil
}

func (p *DB) DeleteTask(taskID common.Hash, projectID uint64) error {
	if err := p.db.Where("task_id = ?", taskID).Where("project_id = ?", projectID).Delete(&task{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete task")
	}
	return nil
}

func New(db *gorm.DB) (*DB, error) {
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db}, nil
}
