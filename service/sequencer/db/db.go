package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"github.com/iotexproject/w3bstream/metrics"
)

type scannedBlockNumber struct {
	gorm.Model
	Number uint64 `gorm:"not null"`
}

type prover struct {
	gorm.Model
	Prover string `gorm:"uniqueIndex:prover,not null"`
}

type task struct {
	gorm.Model
	TaskID   string `gorm:"uniqueIndex:task_uniq,not null"`
	Assigned bool   `gorm:"index:unassigned_task,not null,default:false"`
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

func (p *DB) Provers() ([]common.Address, error) {
	ts := []*prover{}
	if err := p.db.Find(&ts).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query provers")
	}
	res := make([]common.Address, 0, len(ts))
	for _, t := range ts {
		res = append(res, common.HexToAddress(t.Prover))
	}
	metrics.ProverMtc.Set(float64(len(ts)))
	return res, nil
}

func (p *DB) UpsertProver(addr common.Address) error {
	t := prover{
		Prover: addr.Hex(),
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "prover"}},
		DoNothing: true,
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert prover")
}

func (p *DB) CreateTask(taskID common.Hash) error {
	t := &task{
		TaskID:   taskID.Hex(),
		Assigned: false,
	}
	if err := p.db.Create(t).Error; err != nil {
		return errors.Wrap(err, "failed to create task")
	}
	return nil
}

func (p *DB) AssignTasks(ids []common.Hash) error {
	if len(ids) == 0 {
		return nil
	}
	sids := make([]string, 0, len(ids))
	for _, id := range ids {
		sids = append(sids, id.Hex())
	}
	t := &task{
		Assigned: true,
	}
	err := p.db.Model(t).Where("task_id IN ?", sids).Updates(t).Error
	return errors.Wrap(err, "failed to assign tasks")
}

func (p *DB) DeleteTask(taskID, tx common.Hash) error {
	err := p.db.Where("task_id = ?", taskID.Hex()).Delete(&task{}).Error
	return errors.Wrap(err, "failed to delete task")
}

func (p *DB) UnassignedTasks(limit int) ([]common.Hash, error) {
	ts := []*task{}
	if err := p.db.Order("created_at ASC").Where("assigned = false").Find(&ts).Limit(limit).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query unassigned tasks")
	}
	ids := []common.Hash{}
	for _, t := range ts {
		ids = append(ids, common.HexToHash(t.TaskID))
	}
	return ids, nil
}

func New(localDBDir string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(localDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect sqlite")
	}
	if err := db.AutoMigrate(&task{}, &scannedBlockNumber{}, &prover{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &DB{db}, nil
}
