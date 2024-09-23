package postgres

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"github.com/iotexproject/w3bstream/task"
)

type blockNumber struct {
	gorm.Model
	Number uint64 `gorm:"not null"`
}

type currentNBits struct {
	gorm.Model
	NBits uint32 `gorm:"not null"`
}

type chainHead struct {
	gorm.Model
	Hash   common.Hash `gorm:"not null"`
	Number uint64      `gorm:"not null"`
}

type projectProcessedTask struct {
	gorm.Model
	TaskID    uint64 `gorm:"not null"`
	ProjectID uint64 `gorm:"uniqueIndex:project_id,not null"`
}

type taskStateLog struct {
	gorm.Model
	TaskID         uint64     `gorm:"index:state_fetch,not null"`
	ProjectID      uint64     `gorm:"index:state_fetch,not null"`
	ProjectVersion string     `gorm:"index:state_fetch,not null"`
	State          task.State `gorm:"not null"`
	Comment        string
	Result         []byte
}

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) BlockNumber() (uint64, error) {
	t := blockNumber{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to query block number")
	}
	return t.Number, nil
}

func (p *Postgres) UpsertBlockNumber(number uint64) error {
	t := blockNumber{
		Model: gorm.Model{
			ID: 1,
		},
		Number: number,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"number"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert block number")
	}
	return nil
}

func (p *Postgres) NBits() (uint32, error) {
	t := currentNBits{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return 0, errors.Wrap(err, "failed to query nbits")
	}
	return t.NBits, nil
}

func (p *Postgres) UpsertNBits(nbits uint32) error {
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

func (p *Postgres) ChainHead() (uint64, common.Hash, error) {
	t := chainHead{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return 0, common.Hash{}, errors.Wrap(err, "failed to query chain head")
	}
	return t.Number, t.Hash, nil
}

func (p *Postgres) UpsertPrevHash(number uint64, hash common.Hash) error {
	t := chainHead{
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
		return errors.Wrap(err, "failed to upsert chain head")
	}
	return nil
}

func (p *Postgres) ProcessedTaskID(projectID uint64) (uint64, error) {
	t := projectProcessedTask{}
	if err := p.db.Where("project_id = ?", projectID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, errors.Wrapf(err, "failed to query project processed task_id, project_id %v", projectID)
	}
	return t.TaskID, nil
}

func (p *Postgres) UpsertProcessedTask(projectID, taskID uint64) error {
	t := projectProcessedTask{
		ProjectID: projectID,
		TaskID:    taskID,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"task_id"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrapf(err, "failed to upsert project processed task, project_id %v, task_id %v", projectID, taskID)
	}
	return nil
}

func (p *Postgres) Create(tl *task.StateLog, t *task.Task) error {
	l := &taskStateLog{
		TaskID:         tl.TaskID,
		ProjectID:      t.ProjectID,
		ProjectVersion: t.ProjectVersion,
		State:          tl.State,
		Comment:        tl.Comment,
		Result:         tl.Result,
		Model: gorm.Model{
			CreatedAt: tl.CreatedAt,
		},
	}
	if err := p.db.Create(l).Error; err != nil {
		return errors.Wrap(err, "failed to create task state log")
	}
	return nil
}

func (p *Postgres) Fetch(taskID, projectID uint64) ([]*task.StateLog, error) {
	ls := []*taskStateLog{}
	if err := p.db.Order("created_at").Where("task_id = ? AND project_id = ?", taskID, projectID).Find(&ls).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task state log, task_id %v, project_id %v", taskID, projectID)
	}
	tls := []*task.StateLog{}
	for _, l := range ls {
		tls = append(tls, &task.StateLog{
			TaskID:    taskID,
			State:     l.State,
			Comment:   l.Comment,
			Result:    l.Result,
			CreatedAt: l.CreatedAt,
		})
	}
	return tls, nil
}

func New(pgEndpoint string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	if err := db.AutoMigrate(&taskStateLog{}, &projectProcessedTask{}, &blockNumber{}, &currentNBits{}, &chainHead{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &Postgres{db}, nil
}
