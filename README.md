
![](assets/logo.png)

![](https://github.com/dtmirizzi/go-transact/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dtmirizzi/go-transact)](https://goreportcard.com/report/github.com/dtmirizzi/go-transact)

## Install 
```$xslt
 go get -u github.com/dtmirizzi/go-transact
```
## Docs 
[Documentation](https://pkg.go.dev/github.com/dtmirizzi/go-transact/pkg)
## Problem
Often times in `synchronous` distributed systems you will have a set of processes 
that need to be performed at roughly the same time. 
To declare success all sub-processes must complete successfully, 
If one fails it is often common to revert state and retry. 
If synchronicity is not required the problem
is often better solved in an `asynchronous` fashion. This library is ment to provide a simple interface for handling said transactions.  

## Use Case
You are tasked with creating a realtime user management engine by D. Corp.
To create an user you need to add a db table and insert the creds to AWS Incognito.
If either fail you would like to roll back your changes so that you can retry at another time.  

## Example 
```go
import "github.com/dtmirizzi/go-transact"

t := NewTransaction()
	
/// Create DB Table Process

createDBTable := func() error {
    // Do something 
	return nil
}
deleteDBTable := func() error {
    // Do the opposite 
	return nil
}

// This adds the process the transaction 
// Proc is the most basic process possible 
// You may add any struct that meets the Process interface...
t.AddProcess( &Proc{
	Name: "p0", // NAMES MUST BE UNIQUE!
	Up:   createDBTable,
	Down: deleteDBTable,
})


insertCredsToIncognito := func() error {
	return nil
}
removeCredsFromIncognito := func() error {
	return nil
}
t.AddProcess( &Proc{
	Name: "p1",
	Up:   insertCredsToIncognito,
	Down: removeCredsFromIncognito,
})

/// Make Transaction concurrently (NOT TREAD SAFE)
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
