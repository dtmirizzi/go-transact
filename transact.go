package transact

import (
	"sync"
)

// Process is a process in the transaction that can be undone
type Process struct {
	Name     string
	RollOut  func() error
	Rollback func() error
}

// Transaction Transaction is a set of processes
type Transaction struct {
	Processes []Process
}

// NewTransaction Returns a new transaction
func NewTransaction(p ...Process) *Transaction {
	return &Transaction{
		Processes: p,
	}
}

// AddProcess allows you to add processes to the transaction block
func (t *Transaction) AddProcess(p ...Process) {
	t.Processes = append(t.Processes, p...)
}

// Transact performs the transaction
func (t *Transaction) Transact() error {
	fp := make([]Process, 0)

	wg := sync.WaitGroup{}
	for _, p := range t.Processes {
		wg.Add(1)
		go func() {
			err := p.RollOut
			if err != nil {
				fp = append(fp, p)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	ex := except(t.Processes, fp)

	if len(ex) != len(t.Processes) {

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
