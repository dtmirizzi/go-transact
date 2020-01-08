package transact

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Transact(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p0 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return nil
			},
		})

		p1 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return nil
			},
		})

		trans := NewTransaction(p0, p1)

		err := trans.Transact()
		assert.Nil(t, err)
	})
}

func TestTransaction_Transact2(t *testing.T) {
	t.Run("Up Failure", func(t *testing.T) {
		p0 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return errors.New("process failed")
			},
			Down: func() error {
				return nil
			},
		})

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		assert.True(t, tErr.Safe())
		fmt.Println(tErr)
	})
}

func TestTransaction_Transact3(t *testing.T) {
	t.Run("Up and Down Fail", func(t *testing.T) {

		p0 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return errors.New("process 0 failed")
			},
			Down: func() error {
				// This should not run!
				return errors.New("process 0 down failed")
			},
		})

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		// TODO this should be safe the down func should not run!
		assert.False(t, tErr.Safe())
	})

}

func TestTransaction_Transact4(t *testing.T) {
	t.Run("Only Down Failure", func(t *testing.T) {
		p0 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return errors.New("process 0 down failed")
			},
		})

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.Nil(t, err)
	})
}

func TestTransaction_Transact5(t *testing.T) {
	t.Run("P0 Up Failure P1 Down Failure", func(t *testing.T) {
		p0 := NewProc(ProcConfig{
			Name: "p0",
			Up: func() error {
				return errors.New("process 0 up failed")
			},
			Down: func() error {
				return nil
			},
		})

		p1 := NewProc(ProcConfig{
			Name: "p1",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return errors.New("process 1 down failed")
			},
		})

		trans := NewTransaction(p0, p1)

		err := trans.Transact()
		assert.NotNil(t, err)

		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		// This is not safe return false
		assert.False(t, tErr.Safe())
	})
}

func TestExcept(t *testing.T) {

}
