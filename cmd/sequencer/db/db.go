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

type currentNBits struct {
	gorm.Model
	NBits uint32 `gorm:"not null"`
}

type blockHead struct {
	gorm.Model
	Hash   common.Hash `gorm:"not null"`
	Number uint64      `gorm:"not null"`
}

type task struct {
	gorm.Model
	TaskID    common.Hash `gorm:"index:task_fetch,not null"`
	ProjectID uint64      `gorm:"index:task_fetch,not null"`
	Assigned  bool        `gorm:"index:unassigned_task,not null,default:false"`
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

func (p *DB) NBits() (uint32, error) {
	t := currentNBits{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return 0, errors.Wrap(err, "failed to query nbits")
	}
	return t.NBits, nil
}

func (p *DB) UpsertNBits(nbits uint32) error {
	t := currentNBits{
		Model: gorm.Model{
			ID: 1,
		},
		NBits: nbits,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"n_bits"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert nbits")
	}
	return nil
}

func (p *DB) BlockHead() (uint64, common.Hash, error) {
	t := blockHead{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return 0, common.Hash{}, errors.Wrap(err, "failed to query block head")
	}
	return t.Number, t.Hash, nil
}

func (p *DB) UpsertBlockHead(number uint64, hash common.Hash) error {
	t := blockHead{
		Model: gorm.Model{
			ID: 1,
		},
		Hash:   hash,
		Number: number,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"hash", "number"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert block head")
	}
	return nil
}

func (p *DB) CreateTask(projectID uint64, taskID common.Hash) error {
	t := &task{
		TaskID:    taskID,
		ProjectID: projectID,
		Assigned:  false,
	}
	if err := p.db.Create(t).Error; err != nil {
		return errors.Wrap(err, "failed to create task")
	}
	return nil
}

func (p *DB) AssignTask(projectID uint64, taskID common.Hash) error {
	t := &task{
		Assigned: true,
	}
	if err := p.db.Model(t).Where("task_id = ?", taskID).Where("project_id = ?", projectID).Updates(t).Error; err != nil {
		return errors.Wrap(err, "failed to assign task")
	}
	return nil
}

func (p *DB) DeleteTask(projectID uint64, taskID common.Hash) error {
	if err := p.db.Where("task_id = ?", taskID).Where("project_id = ?", projectID).Delete(&task{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete task")
	}
	return nil
}

func (p *DB) UnassignedTask() (common.Hash, uint64, error) {
	ts := []*task{}
	if err := p.db.Order("created_at ASC").Where("assigned = false").Find(&ts).Limit(1).Error; err != nil {
		return common.Hash{}, 0, errors.Wrap(err, "failed to query unassigned task")
	}
	if len(ts) == 0 {
		return common.Hash{}, 0, nil
	}
	return ts[0].TaskID, ts[0].ProjectID, nil
}

func New(db *gorm.DB) (*DB, error) {
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}, &currentNBits{}, &blockHead{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db}, nil
}
