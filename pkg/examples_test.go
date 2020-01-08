package transact

import "fmt"

func ExampleTransaction() {
	t := NewTransaction()

	CreateDBTable := func() error {
		// Do something
		return nil
	}
	DeleteDBTable := func() error {
		// Do something
		return nil
	}
	t.AddProcess(NewProc(ProcConfig{
		Name: "p0",
		Up:   CreateDBTable,
		Down: DeleteDBTable,
	}))

	PutMessageOnQueue := func() error {
		// Do something
		return nil
	}
	DeleteMessageFromQueue := func() error {
		// Do something
		return nil
	}
	t.AddProcess(NewProc(ProcConfig{
		Name: "p1",
		Up:   PutMessageOnQueue,
		Down: DeleteMessageFromQueue,
	}))

	err := t.Transact()
	if err != nil {
		if tErr, ok := err.(*TransactionError); ok {
			fmt.Println(tErr)
			if !tErr.Safe() {
				panic("Failed to safely revert changes!")
			}
		}
	}
}
