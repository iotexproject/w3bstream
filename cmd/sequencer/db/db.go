package db

import (
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

type currentNBits struct {
	gorm.Model
	NBits uint32 `gorm:"not null"`
}

type blockHead struct {
	gorm.Model
	Hash   common.Hash `gorm:"not null"`
	Number uint64      `gorm:"not null"`
}

type prover struct {
	gorm.Model
	Prover common.Address `gorm:"uniqueIndex:prover,not null"`
}

type task struct {
	gorm.Model
	TaskID    common.Hash `gorm:"uniqueIndex:task_uniq,not null"`
	ProjectID uint64      `gorm:"uniqueIndex:task_uniq,not null"`
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
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"number"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert scanned block number")
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
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"n_bits"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert nbits")
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
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"hash", "number"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert block head")
}

func (p *DB) Provers() ([]common.Address, error) {
	ts := []*prover{}
	if err := p.db.Find(&ts).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query provers")
	}
	res := make([]common.Address, 0, len(ts))
	for _, t := range ts {
		res = append(res, t.Prover)
	}
	return res, nil
}

func (p *DB) UpsertProver(addr common.Address) error {
	t := prover{
		Prover: addr,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "prover"}},
		DoNothing: true,
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert prover")
}

func (p *DB) CreateTask(projectID uint64, taskID common.Hash) error {
	t := &task{
		TaskID:    taskID,
		ProjectID: projectID,
		Assigned:  false,
	}
	err := p.db.Create(t).Error
	return errors.Wrap(err, "failed to create task")
}

func (p *DB) AssignTask(projectID uint64, taskID common.Hash, prover common.Address) error {
	t := &task{
		Assigned: true,
	}
	err := p.db.Model(t).Where("task_id = ?", taskID).Where("project_id = ?", projectID).Updates(t).Error
	return errors.Wrap(err, "failed to assign task")
}

func (p *DB) DeleteTask(projectID uint64, taskID, tx common.Hash) error {
	err := p.db.Where("task_id = ?", taskID).Where("project_id = ?", projectID).Delete(&task{}).Error
	return errors.Wrap(err, "failed to delete task")
}

func (p *DB) UnassignedTask() (uint64, common.Hash, error) {
	t := task{}
	if err := p.db.Order("created_at ASC").Where("assigned = false").First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, common.Hash{}, nil
		}
		return 0, common.Hash{}, errors.Wrap(err, "failed to query unassigned task")
	}
	return t.ProjectID, t.TaskID, nil
}

func New(localDBDir string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(localDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect sqlite")
	}
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}, &currentNBits{}, &blockHead{}, &prover{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db}, nil
}
