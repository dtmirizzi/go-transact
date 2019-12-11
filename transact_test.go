package transact

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	t.Run("clear", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			RollOut: func() error {
				return nil
			},
			Rollback: func() error {
				return nil
			},
		},
			Process{
				Name: "p1",
				RollOut: func() error {
					return nil
				},
				Rollback: func() error {
					return nil
				},
			})

		err := trans.Transact()
		assert.Nil(t, err)
	})

	t.Run("roll out failure", func(t *testing.T) {
		trans := NewTransaction(Process{
			Name: "p0",
			RollOut: func() error {
				return errors.New("process failed")
			},
			Rollback: func() error {
				return nil
			},
		},
			Process{
				Name: "p1",
				RollOut: func() error {
					return nil
				},
				Rollback: func() error {
					return nil
				},
			})

		err := trans.Transact()
		assert.NotNil(t, err)
	})

}
