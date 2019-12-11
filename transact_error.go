package transact

// TransactionError Defines transaction errors
// Indexed by process name
type TransactionError struct {
	UpErrors   map[string]ProcessError
	DownErrors map[string]ProcessError
}

// ProcessError Processed error
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
