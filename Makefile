test:
	go test ./pkg -v  -coverprofile=coverage.txt -covermode=atomic -race

vendor:
	go mod vendor -v
