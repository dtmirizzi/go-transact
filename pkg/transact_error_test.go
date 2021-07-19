package transact

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionError_New(t *testing.T) {
	err := NewTransactionError()
	assert.NotNil(t, err)

	err.AppendDownError(ProcessError{Error: errors.New(""), Process: &Proc{}})
	assert.NotNil(t, err.AppendDownError)
	err.AppendUpError(ProcessError{Error: errors.New(""), Process: &Proc{}})
	assert.NotNil(t, err.AppendUpError)

	assert.False(t, err.Safe())

}
