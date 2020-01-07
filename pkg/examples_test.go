package transact

import "fmt"

func ExampleTransaction() {

	trans := NewTransaction(Process{
		Name: "p0",
		Up: func() error {
			// Do something
			return nil
		},
		Down: func() error {

			return nil
		},
	},
		Process{
			Name: "p1",
			Up: func() error {
				return nil
			},
			Down: func() error {
				return nil
			},
		})

	err := trans.Transact()
	if err != nil {
		tErr := err.(*TransactionError)
		fmt.Println(tErr)
		if !tErr.Safe() {
			panic("Failed to safely revert changes!")
		}
	}
}
