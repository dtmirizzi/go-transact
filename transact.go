package transact

import (
	"sync"
)

// Process defines a subroutine in the transaction.
// Up describes changes, Down describes the changes to undo the up function.
// Name must be unique to insure proper error handling.
type Process struct {
	Name string
	Up   func() error
	Down func() error
}

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
func (t *Transaction) Transact() error {
	pErr := TransactionError{}

	// Up
	wg := sync.WaitGroup{}
	for _, p := range t.Processes {
		wg.Add(1)
		go func() {
			err := p.Up()
			if err != nil {
				pErr.UpErrors[p.Name] = ProcessError{
					Process: p,
					Error:   err,
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Validate success
	ex := except(t.Processes, pErr.FailedProcesses())
	if len(ex) != len(t.Processes) {

	}

	// Down

	// Todo Better error handling
	if len(pErr.FailedProcesses()) > 0 {
		return &pErr
	}
	return nil
}

// helpers
func except(l []Process, r []Process) (pe []Process) {
	for _, l0 := range l {
		if !contains(l0, r) {
			pe = append(pe, l0)
		}
	}
	return pe
}

func contains(needle Process, haystack []Process) bool {
	for _, p := range haystack {
		if &needle == &p {
			return true
		}
	}
	return false
}
