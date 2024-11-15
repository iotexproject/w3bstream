package persistence

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type scannedBlockNumber struct {
	gorm.Model
	Number uint64 `gorm:"not null"`
}

type Task struct {
	gorm.Model
	DeviceID       common.Address `gorm:"not null"`
	TaskID         common.Hash    `gorm:"uniqueIndex:task_id_uniq,not null"`
	Nonce          uint64         `gorm:"uniqueIndex:task_uniq,not null"`
	ProjectID      uint64         `gorm:"uniqueIndex:task_uniq,not null"`
	ProjectVersion string         `gorm:"not null"`
	Payloads       datatypes.JSON `gorm:"not null"`
	Signature      []byte         `gorm:"not null"`
}

type AssignedTask struct {
	gorm.Model
	TaskID    common.Hash `gorm:"uniqueIndex:assigned_task_uniq,not null"`
	ProjectID uint64      `gorm:"not null"`
	Prover    common.Address
}

type SettledTask struct {
	gorm.Model
	TaskID    common.Hash `gorm:"uniqueIndex:settled_task_uniq,not null"`
	ProjectID uint64      `gorm:"not null"`
	Tx        common.Hash `gorm:"not null"`
}

type Persistence struct {
	db *gorm.DB
}

func (p *Persistence) CreateTask(m *Task) error {
	err := p.db.Create(m).Error
	return errors.Wrap(err, "failed to create task")
}

func (p *Persistence) FetchTask(taskID common.Hash) (*Task, error) {
	t := Task{}
	if err := p.db.Where("task_id = ?", taskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query task")
	}
	return &t, nil
}

func (p *Persistence) UpsertAssignedTask(projectID uint64, taskID common.Hash, prover common.Address) error {
	t := AssignedTask{
		ProjectID: projectID,
		TaskID:    taskID,
		Prover:    prover,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"prover"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert assigned task")
}

func (p *Persistence) UpsertSettledTask(projectID uint64, taskID, tx common.Hash) error {
	t := SettledTask{
		ProjectID: projectID,
		TaskID:    taskID,
		Tx:        tx,
	}
	err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"tx"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert settled task")
}

func (p *Persistence) FetchAssignedTask(taskID common.Hash) (*AssignedTask, error) {
	t := AssignedTask{}
	if err := p.db.Where("task_id = ?", taskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query assigned task")
	}
	return &t, nil
}

func (p *Persistence) FetchSettledTask(taskID common.Hash) (*SettledTask, error) {
	t := SettledTask{}
	if err := p.db.Where("task_id = ?", taskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query settled task")
	}
	return &t, nil
}

func (p *Persistence) ScannedBlockNumber() (uint64, error) {
	t := scannedBlockNumber{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to query scanned block number")
	}
	return t.Number, nil
}

func (p *Persistence) UpsertScannedBlockNumber(number uint64) error {
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

func NewPersistence(pgEndpoint string) (*Persistence, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	if err := db.AutoMigrate(&scannedBlockNumber{}, &Task{}, &AssignedTask{}, &SettledTask{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &Persistence{db}, nil
}
