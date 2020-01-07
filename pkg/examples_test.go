package transact

func ExampleTransaction_Transact() {

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

	}
}
