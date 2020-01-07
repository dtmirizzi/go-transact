# Go-Transact

![](https://github.com/dtmirizzi/go-transact/workflows/Test/badge.svg)


## Install 
```$xslt
 go get -u github.com/dtmirizzi/go-transact
```
## Usage
[Documentation](https://godoc.org/github.com/dtmirizzi/go-transact/pkg)
### Basic Example 
```
    t := NewTransaction()
	
	// Create DB Table 
	// This function describes how to create the table
	CreateDBTable := func() error {
		// Do something
		return nil
	}
	// This function describes how to remove the table 
	DeleteDBTable := func() error {
		// Do something
		return nil
	}
	// This adds the process the the queue 
	t.AddProcess(Process{
		Name: "p0",
		Up:   CreateDBTable,
		Down: DeleteDBTable,
	})

	err := t.Transact()
	if err != nil {
		tErr := err.(*TransactionError)
		fmt.Println(tErr)
		if !tErr.Safe() {
			panic("Failed to safely revert changes!")
		}
	}
```

## Development
- Install [Precommit](https://pre-commit.com/), [go-acc](https://github.com/ory/go-acc), and [golangci-lint](https://github.com/golangci/golangci-lint).
- Run ```pre-commit install```
- Ship it!! 
