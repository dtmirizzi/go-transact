package transact

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
	return "placeholder"
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
