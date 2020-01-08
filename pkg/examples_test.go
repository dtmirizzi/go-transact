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
	t.AddProcess(&Proc{
		PName:    "p0",
		UpFunc:   CreateDBTable,
		DownFunc: DeleteDBTable,
	})

	PutMessageOnQueue := func() error {
		// Do something
		return nil
	}
	DeleteMessageFromQueue := func() error {
		// Do something
		return nil
	}
	t.AddProcess(&Proc{
		PName:    "p1",
		UpFunc:   PutMessageOnQueue,
		DownFunc: DeleteMessageFromQueue,
	})

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
