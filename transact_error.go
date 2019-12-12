package transact

import (
	"fmt"
)


// TransactionError defines errors that occur during a transaction, and provide helpers for operating and handling.
// errors indexed by process name.
type TransactionError struct {
	UpErrors   map[string]ProcessError
	DownErrors map[string]ProcessError
}

// ProcessError contains a process and error that occurred when running said process.
type ProcessError struct {
	Process Process
	Error   error
}

// Error Implements the error
func (t *TransactionError) Error() string {
	e := "UpErrors:[ "
	for _, p := range t.UpErrors {
		e += formatProcessError(p)
	}
	e +=  "] DownErrors:[ "
	for _, p := range t.DownErrors {
		e += formatProcessError(p)
	}
	e +=  "]"
	return e
}

// FailedProcesses return all failed processes
func (t *TransactionError) FailedProcesses() (ps []Process) {
	for _, p := range t.UpErrors {
		ps = append(ps, p.Process)
	}
	for _, p := range t.DownErrors {
		ps = append(ps, p.Process)
	}
	return ps
}

// FailedProcessErrors return all failed processes
func (t *TransactionError) FailedProcessErrors() (ps []ProcessError) {
	for _, p := range t.UpErrors {
		ps = append(ps, p)
	}
	for _, p := range t.DownErrors {
		ps = append(ps, p)
	}
	return ps
}

// Safe is a helper a returns true when an error occurred but all the down operations ran successfully.
func (t *TransactionError) Safe() bool {
	return len(t.DownErrors) > 0
}

func formatProcessError(p ProcessError) string {
	return fmt.Sprintf("\n Process: %s -> err: %s,", p.Process.Name, p.Error)
}