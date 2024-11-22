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

type AssignedTask struct {
	gorm.Model
	TaskID common.Hash `gorm:"uniqueIndex:assigned_task_uniq,not null"`
	Prover common.Address
}

type SettledTask struct {
	gorm.Model
	TaskID common.Hash `gorm:"uniqueIndex:settled_task_uniq,not null"`
	Tx     common.Hash `gorm:"not null"`
}

type ProjectDevice struct {
	gorm.Model
	ProjectID     uint64         `gorm:"uniqueIndex:project_device_uniq,not null"`
	DeviceAddress common.Address `gorm:"uniqueIndex:project_device_uniq,not null"`
}

func (p *DB) UpsertAssignedTask(taskID common.Hash, prover common.Address) error {
	t := AssignedTask{
		TaskID: taskID,
		Prover: prover,
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"prover"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert assigned task")
}

func (p *DB) UpsertSettledTask(taskID, tx common.Hash) error {
	t := SettledTask{
		TaskID: taskID,
		Tx:     tx,
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"tx"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert settled task")
}

func (p *DB) FetchAssignedTask(taskID common.Hash) (*AssignedTask, error) {
	t := AssignedTask{}
	if err := p.sqlite.Where("task_id = ?", taskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query assigned task")
	}
	return &t, nil
}

func (p *DB) FetchSettledTask(taskID common.Hash) (*SettledTask, error) {
	t := SettledTask{}
	if err := p.sqlite.Where("task_id = ?", taskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query settled task")
	}
	return &t, nil
}

func (p *DB) UpsertProjectDevice(projectID uint64, address common.Address) error {
	t := ProjectDevice{
		ProjectID:     projectID,
		DeviceAddress: address,
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}, {Name: "device_address"}},
		DoNothing: true,
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project device")
}

func (p *DB) DeleteProjectDevice(projectID uint64, address common.Address) error {
	err := p.sqlite.Where("project_id = ?", projectID).Where("device_address = ?", address).Delete(&ProjectDevice{}).Error
	return errors.Wrap(err, "failed to delete project device")
}

func (p *DB) IsDeviceApproved(projectID uint64, address common.Address) (bool, error) {
	t := ProjectDevice{}
	if err := p.sqlite.Where("project_id = ?", projectID).Where("device_address = ?", address).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, errors.Wrap(err, "failed to query project device")
	}
	return true, nil
}

func (p *DB) ScannedBlockNumber() (uint64, error) {
	t := scannedBlockNumber{}
	if err := p.sqlite.Where("id = ?", 1).First(&t).Error; err != nil {
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
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"number"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert scanned block number")
}

func newSqlite(localDBDir string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(localDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect sqlite")
	}
	if err := db.AutoMigrate(&scannedBlockNumber{}, &AssignedTask{}, &SettledTask{}, &ProjectDevice{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return db, nil
}
