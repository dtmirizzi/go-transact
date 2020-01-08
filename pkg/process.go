package transact

// Process defines a subroutine in the transaction.
// Up describes changes, Down describes the changes to undo the up function.
// Name must be unique to insure proper error handling.
type Process interface {
	Name() string
	Up() error
	Down() error
}

// ProcConfig defines a subroutine in the transaction.
type ProcConfig struct {
	Name string
	Up   func() error
	Down func() error
}

// Proc asd
type Proc struct {
	N string
	U func() error
	D func() error
}

// Name asd
func (p *Proc) Name() string {
	return p.N
}

// Up asd
func (p *Proc) Up() error {
	return p.U()
}

// Down asd
func (p *Proc) Down() error {
	return p.D()
}

// NewProc asd
func NewProc(pc ProcConfig) *Proc {
	return &Proc{
		N: pc.Name,
		U: pc.Up,
		D: pc.Down,
	}
}
