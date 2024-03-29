package transact

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// Transaction is a set of dependant sub-processes
type Transaction struct {
	Processes []Process
}

// NewTransaction returns a new transaction
func NewTransaction(p ...Process) *Transaction {
	return &Transaction{
		Processes: p,
	}
}

// AddProcess adds processes to the transaction block
func (t *Transaction) AddProcess(p ...Process) {
	t.Processes = append(t.Processes, p...)
}

// ValidateTransacton validates the tranaction
// it ensures that there are no repeated names
func (t *Transaction) ValidateTransacton() error {
	m := make(map[string]struct{})

	for _, p := range t.Processes {
		if _, exists := m[p.Name()]; exists {
			return fmt.Errorf("process %s has duplicates", p.Name())
		}
		m[p.Name()] = struct{}{}
	}

	return nil
}

// Transact performs the transaction
// [Not thread safe]
func (t *Transaction) Transact() error {
	err := t.ValidateTransacton()
	if err != nil {
		return err
	}

	pErr := t.up()
	if pErr != nil {
		pErr = t.down(pErr)
	}

	if pErr != nil {
		return pErr
	}

	return nil
}

// TransactCtx performs a transaction with a context deadline
func (t *Transaction) TransactCtx(ctx context.Context) error {

	var err error
	done := make(chan struct{})

	go func() {
		err = t.Transact()

		done <- struct{}{}
	}()

	select {
	case <-done:
		break
	case <-ctx.Done():
		return errors.New("context timeout exceeded")
	}
	if err != nil {
		return err
	}

	return nil

}

func (t *Transaction) up() *TransactionError {
	var wg sync.WaitGroup
	errs := make([]ProcessError, 0)

	for _, p := range t.Processes {
		wg.Add(1)
		go func(process Process) {
			err := process.Up()
			if err != nil {
				errs = append(errs, ProcessError{
					Process: process,
					Error:   err,
				})
			}
			wg.Done()
		}(p)
	}
	wg.Wait()

	if len(errs) > 0 {
		err := NewTransactionError()
		err.AppendUpError(errs...)
		return err
	}

	return nil
}

func (t *Transaction) down(pErr *TransactionError) *TransactionError {
	var wg sync.WaitGroup

	dp := Except(t.Processes, pErr.FailedProcesses())

	for _, p := range dp {
		wg.Add(1)
		go func(process Process) {
			err := process.Down()
			if err != nil {
				pErr.AppendDownError(ProcessError{
					Process: process,
					Error:   err,
				})
			}
			wg.Done()
		}(p)
	}
	wg.Wait()

	return pErr
}

// helpers

// Except returns all left unique processes
func Except(l []Process, r []Process) (pe []Process) {
	for _, l0 := range l {
		if !contains(l0, r) {
			pe = append(pe, l0)
		}
	}
	return pe
}

func contains(needle Process, haystack []Process) bool {
	for _, p := range haystack {
		if needle == p {
			return true
		}
	}
	return false
}
