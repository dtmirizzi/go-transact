package transact

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Transact(t *testing.T) {
	t.Run("clear", func(t *testing.T) {
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

	t.Run("roll out failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			Up: func() error {
				return errors.New("process failed")
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
		assert.NotNil(t, err)
	})

}
