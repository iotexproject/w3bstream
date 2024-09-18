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

type currentDifficulty struct {
	gorm.Model
	Difficulty [8]byte `gorm:"not null"`
}

type prevHash struct {
	gorm.Model
	PrevHash common.Hash `gorm:"not null"`
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

func (p *Postgres) Difficulty() ([8]byte, error) {
	t := currentDifficulty{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return [8]byte{}, errors.Wrap(err, "failed to query difficulty")
	}
	return t.Difficulty, nil
}

func (p *Postgres) UpsertDifficulty(difficulty [8]byte) error {
	t := currentDifficulty{
		Model: gorm.Model{
			ID: 1,
		},
		Difficulty: difficulty,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"difficulty"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert difficulty")
	}
	return nil
}

func (p *Postgres) PrevHash() (common.Hash, error) {
	t := prevHash{}
	if err := p.db.Where("id = ?", 1).First(&t).Error; err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to query prev hash")
	}
	return t.PrevHash, nil
}

func (p *Postgres) UpsertPrevHash(hash common.Hash) error {
	t := prevHash{
		Model: gorm.Model{
			ID: 1,
		},
		PrevHash: hash,
	}
	if err := p.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"prev_hash"}),
	}).Create(&t).Error; err != nil {
		return errors.Wrap(err, "failed to upsert prev hash")
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
	if err := db.AutoMigrate(&taskStateLog{}, &projectProcessedTask{}, &blockNumber{}, &currentDifficulty{}, &prevHash{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &Postgres{db}, nil
}
