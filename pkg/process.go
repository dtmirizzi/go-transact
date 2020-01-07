package transact

// Process defines a subroutine in the transaction.
// Up describes changes, Down describes the changes to undo the up function.
// Name must be unique to insure proper error handling.
type Process struct {
	Name string
	Up   func() error
	Down func() error
}
