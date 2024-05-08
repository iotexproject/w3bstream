package persistence

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewPersistence(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToConnectPg", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(gorm.Open, nil, errors.New(t.Name()))
		_, err := NewPersistence("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToMigrate", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(gorm.Open, &gorm.DB{}, nil)
		p.ApplyMethodReturn(&gorm.DB{}, "AutoMigrate", errors.New(t.Name()))
		_, err := NewPersistence("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(gorm.Open, &gorm.DB{}, nil)
		p.ApplyMethodReturn(&gorm.DB{}, "AutoMigrate", nil)
		_, err := NewPersistence("")
		r.NoError(err)
	})
}

func TestPersistence_FetchMessage(t *testing.T) {
	r := require.New(t)

	ps := &Persistence{
		db: &gorm.DB{
			Error:     nil,
			Statement: &gorm.Statement{},
		},
	}

	t.Run("FailedToQuery", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Find", &gorm.DB{Error: errors.New(t.Name())})
		_, err := ps.FetchMessage("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{}, {}, {}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return ps.db
			},
		)
		ms, err := ps.FetchMessage("")
		r.NoError(err)
		r.Equal(3, len(ms))
	})
}

func TestPersistence_FetchTask(t *testing.T) {
	r := require.New(t)

	ps := &Persistence{
		db: &gorm.DB{
			Error:     nil,
			Statement: &gorm.Statement{},
		},
	}

	t.Run("FailedToQuery", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Find", &gorm.DB{Error: errors.New(t.Name())})
		_, err := ps.FetchTask("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Task{{}, {}, {}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return ps.db
			},
		)
		ms, err := ps.FetchTask("")
		r.NoError(err)
		r.Equal(3, len(ms))
	})
}

func TestPersistence_createMessageTx(t *testing.T) {
	r := require.New(t)

	ps := &Persistence{}

	t.Run("FailedToCreate", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{Error: errors.New(t.Name())})
		err := ps.createMessageTx(&gorm.DB{}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{})
		err := ps.createMessageTx(&gorm.DB{}, nil)
		r.NoError(err)
	})
}

func TestPersistence_aggregateTaskTx(t *testing.T) {
	r := require.New(t)

	ps := &Persistence{}

	t.Run("FailedToFetch", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Find", &gorm.DB{Error: errors.New(t.Name())})

		err := ps.aggregateTaskTx(&gorm.DB{}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FetchZero", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Where", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)
		err := ps.aggregateTaskTx(&gorm.DB{}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.NoError(err)
	})

	t.Run("FailedToUpdate", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1"}, {MessageID: "m2"}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodReturn(&gorm.DB{}, "Model", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Update", &gorm.DB{Error: errors.New(t.Name())})

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToMarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1"}, {MessageID: "m2"}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodReturn(&gorm.DB{}, "Model", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Update", &gorm.DB{Error: nil})
		p.ApplyFuncReturn(json.Marshal, nil, errors.New(t.Name()))

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToCreate", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1"}, {MessageID: "m2"}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodReturn(&gorm.DB{}, "Model", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Update", &gorm.DB{Error: nil})
		p.ApplyFuncReturn(json.Marshal, []byte(""), nil)
		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{Error: errors.New(t.Name())})

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToSign", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1", Data: []byte("")}, {MessageID: "m2", Data: []byte("")}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodReturn(&gorm.DB{}, "Model", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Update", &gorm.DB{Error: nil})
		p.ApplyFuncReturn(json.Marshal, []byte(""), nil)
		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{Error: nil})
		p.ApplyPrivateMethod(
			&Task{},
			"sign",
			func(sk *ecdsa.PrivateKey, projectID uint64, clientID string, messages ...[]byte) (string, error) {
				return "", errors.New(t.Name())
			},
		)

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToUpdateSign", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
			{Values: Params{&gorm.DB{Error: errors.New(t.Name())}}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1", Data: []byte("")}, {MessageID: "m2", Data: []byte("")}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodSeq(&gorm.DB{}, "Model", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodSeq(&gorm.DB{}, "Update", []OutputCell{
			{Values: Params{&gorm.DB{Error: nil}}},
			{Values: Params{&gorm.DB{Error: nil}}},
		})
		p.ApplyFuncReturn(json.Marshal, []byte(""), nil)
		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{Error: nil})
		p.ApplyPrivateMethod(
			&Task{},
			"sign",
			func(sk *ecdsa.PrivateKey, projectID uint64, clientID string, messages ...[]byte) (string, error) {
				return "", nil
			},
		)

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gorm.DB{}, "Clauses", ps.db)
		p.ApplyMethodReturn(&gorm.DB{}, "Order", ps.db)
		p.ApplyMethodSeq(&gorm.DB{}, "Where", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
			{Values: Params{&gorm.DB{Error: nil}}},
		})
		p.ApplyMethodReturn(&gorm.DB{}, "Limit", ps.db)
		p.ApplyMethod(
			&gorm.DB{},
			"Find",
			func(_ *gorm.DB, v any) *gorm.DB {
				vi := reflect.ValueOf(&([]*Message{{MessageID: "m1", Data: []byte("")}, {MessageID: "m2", Data: []byte("")}}))
				vo := reflect.ValueOf(v)
				if vi.IsValid() && vo.IsValid() {
					vo.Elem().Set(vi.Elem())
				}
				return &gorm.DB{Error: nil}
			},
		)

		p.ApplyMethodSeq(&gorm.DB{}, "Model", []OutputCell{
			{Values: Params{ps.db}},
			{Values: Params{ps.db}},
		})
		p.ApplyMethodSeq(&gorm.DB{}, "Update", []OutputCell{
			{Values: Params{&gorm.DB{Error: nil}}},
			{Values: Params{&gorm.DB{Error: nil}}},
		})
		p.ApplyFuncReturn(json.Marshal, []byte(""), nil)
		p.ApplyMethodReturn(&gorm.DB{}, "Create", &gorm.DB{Error: nil})
		p.ApplyPrivateMethod(
			&Task{},
			"sign",
			func(sk *ecdsa.PrivateKey, projectID uint64, clientID string, messages ...[]byte) (string, error) {
				return "", nil
			},
		)

		err := ps.aggregateTaskTx(&gorm.DB{Statement: &gorm.Statement{}}, 0, &Message{
			ClientID:       "clientID",
			ProjectID:      0,
			ProjectVersion: "0.1",
		}, nil)
		r.NoError(err)
	})
}

func TestTask_sign(t *testing.T) {
	r := require.New(t)

	task := &Task{}
	t.Run("FailedToWriteID", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(binary.Write, errors.New(t.Name()))

		_, err := task.sign(nil, uint64(0), "", nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToWriteProjectID", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		})

		_, err := task.sign(nil, uint64(0), "", nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToWriteClientID", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		})
		p.ApplyMethodReturn(&bytes.Buffer{}, "WriteString", 0, errors.New(t.Name()))

		_, err := task.sign(nil, uint64(0), "", nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToWriteMessages", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		})
		p.ApplyMethodReturn(&bytes.Buffer{}, "WriteString", 0, nil)
		p.ApplyMethodReturn(&bytes.Buffer{}, "Write", 0, errors.New(t.Name()))

		_, err := task.sign(nil, uint64(0), "", nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToSign", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		})
		p.ApplyMethodReturn(&bytes.Buffer{}, "WriteString", 0, nil)
		p.ApplyMethodReturn(&bytes.Buffer{}, "Write", 0, nil)
		p.ApplyFuncReturn(crypto.Sign, nil, errors.New(t.Name()))

		_, err := task.sign(nil, uint64(0), "", nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		})
		p.ApplyMethodReturn(&bytes.Buffer{}, "WriteString", 0, nil)
		p.ApplyMethodReturn(&bytes.Buffer{}, "Write", 0, nil)
		p.ApplyFuncReturn(crypto.Sign, nil, nil)

		_, err := task.sign(nil, uint64(0), "", nil)
		r.NoError(err)
	})
}
