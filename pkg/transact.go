package transact

import (
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

// Transact performs the transaction
// [Not thread safe]
func (t *Transaction) Transact() error {
	pErr := t.up()
	if pErr != nil {
		pErr = t.down(pErr)
	}
	return pErr
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
