# Go-Transact

## Install 
```$xslt
 go get -u github.com/dtmirizzi/go-transact
```
## Usage
[Documentation](https://godoc.org/github.com/dtmirizzi/go-transact/pkg)
### Basic Example 
```
t := NewTransaction(Process{
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

err := t.Transact()
if err != nil {

}
```

## Development
- Install [Precommit](https://pre-commit.com/), [go-acc](https://github.com/ory/go-acc), and [golangci-lint](https://github.com/golangci/golangci-lint).
- Run ```pre-commit install```
- Ship it!! 
