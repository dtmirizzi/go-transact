package transact

// Process defines a subroutine in the transaction.
// Up describes changes, Down describes the changes to undo the up function.
// Name must be unique to insure proper error handling.
type Process interface {
	Name() string
	Up() error
	Down() error
}

// ProcConfig defines the most simple process possible
type ProcConfig struct {
	// Name should be a unique identifier for the process
	Name string
	// Up is a simple anon function that is the primary intent of the process
	Up func() error
	// Down is a simple anon function that is used to undo the the up function
	Down func() error
}

// Proc is the most simple process possible
// it wraps the fuction calls to meet its basic interface
type Proc struct {
	PName    string
	UpFunc   func() error
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

// Down wraps the Down anon funtion
func (p *Proc) Down() error {
	return p.DownFunc()
}

// NewProc builds a new proc from config
func NewProc(pc *ProcConfig) *Proc {
	return &Proc{
		PName:    pc.Name,
		UpFunc:   pc.Up,
		DownFunc: pc.Down,
	}
}
