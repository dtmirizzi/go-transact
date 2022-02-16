package transact

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_NewTransaction(t *testing.T) {
	trans := NewTransaction()
	assert.NotNil(t, trans)
}

func TestTransaction_ValidateTransaction(t *testing.T) {
	cases := []struct {
		Name      string
		Processes []*Proc
		err       error
	}{
		{Name: "Valid",
			Processes: []*Proc{{
				PName: "p0",
				UpFunc: func() error {
					return nil
				},
				DownFunc: func() error {
					return nil
				},
			},
				{
					PName: "p1",
					UpFunc: func() error {
						return nil
					},
					DownFunc: func() error {
						return nil
					},
				}},
			err: nil},
		{
			Name: "Invalid",
			Processes: []*Proc{{
				PName: "p0",
				UpFunc: func() error {
					return nil
				},
				DownFunc: func() error {
					return nil
				},
			}, {
				PName: "p0",
				UpFunc: func() error {
					return nil
				},
				DownFunc: func() error {
					return nil
				},
			},
			},
			err: fmt.Errorf("process p0 has duplicates"),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			trans := NewTransaction()
			for _, p := range c.Processes {
				trans.AddProcess(p)
			}
			err := trans.Transact()
			assert.Equal(t, c.err, err)
		})
	}
}

func TestTransaction_Transact(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return nil
			},
			DownFunc: func() error {
				return nil
			},
		}

		p1 := &Proc{
			PName: "p1",
			UpFunc: func() error {
				return nil
			},
			DownFunc: func() error {
				return nil
			},
		}

		trans := NewTransaction(p0, p1)

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("Up Failure", func(t *testing.T) {
		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return errors.New("process failed")
			},
			DownFunc: func() error {
				return nil
			},
		}

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		assert.True(t, tErr.Safe())
	})

	t.Run("Up and Down Fail", func(t *testing.T) {

		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return errors.New("process 0 failed")
			},
			DownFunc: func() error {
				// This should not run!
				return errors.New("process 0 down failed")
			},
		}

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		assert.True(t, tErr.Safe())
	})

	t.Run("Up Failure", func(t *testing.T) {
		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return errors.New("process failed")
			},
			DownFunc: func() error {
				return nil
			},
		}

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.NotNil(t, err)
		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		assert.True(t, tErr.Safe())
	})

	t.Run("Only Down Failure", func(t *testing.T) {
		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return nil
			},
			DownFunc: func() error {
				return errors.New("process 0 down failed")
			},
		}

		trans := NewTransaction(p0)

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("P0 Up Failure P1 Down Failure", func(t *testing.T) {
		p0 := &Proc{
			PName: "p0",
			UpFunc: func() error {
				return errors.New("process 0 up failed")
			},
			DownFunc: func() error {
				return nil
			},
		}

		p1 := &Proc{
			PName: "p1",
			UpFunc: func() error {
				return nil
			},
			DownFunc: func() error {
				return errors.New("process 1 down failed")
			},
		}

		trans := NewTransaction(p0, p1)

		err := trans.Transact()
		assert.NotNil(t, err)

		tErr, ok := err.(*TransactionError)
		assert.True(t, ok)
		// This is not safe return false
		assert.False(t, tErr.Safe())
	})
}
