package db

import (
	"math/big"

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
	TaskID string `gorm:"uniqueIndex:assigned_task_uniq,not null"`
	Prover string
}

type SettledTask struct {
	gorm.Model
	TaskID string `gorm:"uniqueIndex:settled_task_uniq,not null"`
	Tx     string `gorm:"not null"`
}

type ProjectDevice struct {
	gorm.Model
	ProjectID     string `gorm:"uniqueIndex:project_device_uniq,not null"`
	DeviceAddress string `gorm:"uniqueIndex:project_device_uniq,not null"`
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

func (p *DB) CreateTask(t *Task) error {
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoNothing: true,
	}).Create(t).Error
	return errors.Wrap(err, "failed to create task")
}

func (p *DB) FetchAllTask() ([]*Task, error) {
	ts := []*Task{}
	if err := p.sqlite.Order("created_at ASC").Find(&ts).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query all tasks")
	}
	return ts, nil
}

func (p *DB) DeleteTasks(ts []*Task) error {
	taskIDs := make([]string, len(ts))
	for i, t := range ts {
		taskIDs[i] = t.TaskID
	}
	if len(taskIDs) == 0 {
		return nil
	}
	err := p.sqlite.Unscoped().Where("task_id IN ?", taskIDs).Delete(&Task{}).Error
	return errors.Wrap(err, "failed to delete tasks")
}

func (p *DB) UpsertAssignedTask(taskID common.Hash, prover common.Address) error {
	t := AssignedTask{
		TaskID: taskID.Hex(),
		Prover: prover.Hex(),
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"prover"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert assigned task")
}

func (p *DB) UpsertSettledTask(taskID, tx common.Hash) error {
	t := SettledTask{
		TaskID: taskID.Hex(),
		Tx:     tx.Hex(),
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"tx"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert settled task")
}

func (p *DB) FetchAssignedTask(taskID common.Hash) (*AssignedTask, error) {
	t := AssignedTask{}
	if err := p.sqlite.Where("task_id = ?", taskID.Hex()).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query assigned task")
	}
	return &t, nil
}

func (p *DB) FetchSettledTask(taskID common.Hash) (*SettledTask, error) {
	t := SettledTask{}
	if err := p.sqlite.Where("task_id = ?", taskID.Hex()).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to query settled task")
	}
	return &t, nil
}

func (p *DB) UpsertProjectDevice(projectID *big.Int, address common.Address) error {
	t := ProjectDevice{
		ProjectID:     projectID.String(),
		DeviceAddress: address.Hex(),
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}, {Name: "device_address"}},
		DoNothing: true,
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project device")
}

func (p *DB) IsDeviceApproved(projectID *big.Int, address common.Address) (bool, error) {
	t := ProjectDevice{}
	if err := p.sqlite.Where("project_id = ?", projectID.String()).Where("device_address = ?", address.Hex()).First(&t).Error; err != nil {
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

func (p *DB) Project(projectID *big.Int) (string, common.Hash, error) {
	t := project{}
	if err := p.sqlite.Where("project_id = ?", projectID.String()).First(&t).Error; err != nil {
		return "", common.Hash{}, errors.Wrap(err, "failed to query project")
	}
	return t.URI, common.HexToHash(t.Hash), nil
}

func (p *DB) UpsertProject(projectID *big.Int, uri string, hash common.Hash) error {
	t := project{
		ProjectID: projectID.String(),
		URI:       uri,
		Hash:      hash.Hex(),
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"uri", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project")
}

func (p *DB) ProjectFile(projectID *big.Int) ([]byte, common.Hash, error) {
	t := projectFile{}
	if err := p.sqlite.Where("project_id = ?", projectID.String()).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.Hash{}, nil
		}
		return nil, common.Hash{}, errors.Wrap(err, "failed to query project file")
	}
	return t.File, common.HexToHash(t.Hash), nil
}

func (p *DB) UpsertProjectFile(projectID *big.Int, file []byte, hash common.Hash) error {
	t := projectFile{
		ProjectID: projectID.String(),
		File:      file,
		Hash:      hash.Hex(),
	}
	err := p.sqlite.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"file", "hash"}),
	}).Create(&t).Error
	return errors.Wrap(err, "failed to upsert project file")
}

func newSqlite(localDBDir string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(localDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect sqlite")
	}
	if err := db.AutoMigrate(&scannedBlockNumber{}, &Task{}, &AssignedTask{}, &SettledTask{}, &ProjectDevice{}, &project{}, &projectFile{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return db, nil
}
