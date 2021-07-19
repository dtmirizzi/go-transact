test:
	go test ./pkg -v  -coverprofile=coverage.txt -covermode=atomic

vendor:
	go mod vendor -v
