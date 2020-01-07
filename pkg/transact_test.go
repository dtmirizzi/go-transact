package transact

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Transact(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return nil
			},
		},
			Process{
				Name: "p1",
				Up: func() error {
					return nil
				},
				Down: func() error {
					return nil
				},
			})

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("Up Failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return errors.New("process failed")
			},
			Down: func() error {
				return nil
			},
		})

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		assert.True(t, tErr.Safe())
		fmt.Println(tErr)
	})

	t.Run("Up and Down Fail", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return errors.New("process 0 failed")
			},
			Down: func() error {
				// This should not run!
				return errors.New("process 0 down failed")
			},
		})

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		// TODO this should be safe the down func should not run!
		assert.True(t, tErr.Safe())
	})

	t.Run("Only Down Failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return errors.New("process 0 down failed")
			},
		})

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("Only Down Failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return errors.New("process 0 down failed")
			},
		})

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("P0 Up Failure P1 Down Failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return errors.New("process 0 up failed")
			},
			Down: func() error {
				return nil
			},
		},
			Process{
				Name: "p1",
				Up: func() error {
					return nil
				},
				Down: func() error {
					return errors.New("process 1 down failed")
				},
			})

		err := trans.Transact()
		assert.NotNil(t, err)

		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		// This is not safe return false
		assert.False(t, tErr.Safe())
	})

}
