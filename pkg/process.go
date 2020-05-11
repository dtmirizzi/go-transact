package transact

// Process defines a subroutine in the transaction.
// Up describes changes, Down describes the changes to undo the up function.
// Name must be unique to insure proper error handling.
type Process interface {
	Name() string
	Up() error
	Down() error
}

// Proc is the most simple process possible
// it wraps the fuction calls to meet its basic interface
type Proc struct {
	// PName should be a unique identifier for the process
	PName string
	// UpFunc is a simple anon function that is the primary intent of the process
	UpFunc func() error
	// DownFunc is a simple anon function that is the primary intent of the process
	DownFunc func() error
}

// Name Wraps the ProcName
func (p *Proc) Name() string {
	return p.PName
}

// Up wraps the Up anon function
func (p *Proc) Up() error {
	return p.UpFunc()
}

// Down wraps the Down anon function
func (p *Proc) Down() error {
	return p.DownFunc()
}
