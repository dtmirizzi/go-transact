
![](assets/logo.png)

![](https://github.com/dtmirizzi/go-transact/workflows/Test/badge.svg)


## Install 
```$xslt
 go get -u github.com/dtmirizzi/go-transact
```
## Docs 
[Documentation](https://godoc.org/github.com/dtmirizzi/go-transact/pkg)
## Problem
Often times in `synchronous` distributed systems you will have a set of processes 
that need to be performed at roughly the same time. 
To declare success all sub-processes must complete successfully, 
If one fails it is often common to undo all processes, revert state and retry. 
If synchronicity is not required the problem
is often better solved in an `asynchronous` fashion.  

## Use Case
You are tasked with creating a realtime user management engine by D. Corp.
To create an user you need to add a db table and insert the creds to AWS Incognito.
If either fail you would like to roll back your changes so that you can retry at another time.  

## Example 
```go
t := NewTransaction()
	
//-- Create DB Table Process  --//
// This function describes how to create the table
CreateDBTable := func() error {
    // Do something 
	return nil
}
// This function describes how to remove the table 
DeleteDBTable := func() error {
    // Do the opposite 
	return nil
}

// This adds the process the the queue 
// NewProc is the most basic process defined process possible 
// You may add any struct that meets the Process interface...
t.AddProcess( &ProcC{
	Name: "p0", // PROCESS NAMES MUST BE UNIQUE!
	Up:   CreateDBTable,
	Down: DeleteDBTable,
})


//-- Insert Creds To Incognito process --//
InsertCredsToIncognito := func() error {
	return nil
}
RemoveCredsToIncognito := func() error {
	return nil
}
t.AddProcess( &Proc{
	Name: "p1",
	Up:   InsertCredsToIncognito,
	Down: RemoveCredsToIncognito,
})

//-- Make Transaction concurrently (NOT TREAD SAFE) --//
err := t.Transact()
if err != nil {
    // You may cast the error to gain helper methods 
	if tErr, ok := err.(*TransactionError); ok {
        // this ensures that all transactions were undone. 
        if !tErr.Safe() {
            panic("Failed to safely revert changes!")
        }
    }
    fmt.Println(err)	
}
```
## Development
- Install [Precommit](https://pre-commit.com/), [go-acc](https://github.com/ory/go-acc), 
[gocyclo](https://github.com/fzipp/gocyclo), 
and [golangci-lint](https://github.com/golangci/golangci-lint).
- Run ```pre-commit install```
- Ship it!! 
