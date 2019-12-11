package transact

// TransactionError Defines transaction errors
type TransactionError struct {
	OutErrorMapping  map[Process]error
	BackErrorMapping map[Process]error
}

// Error Implements the error
func (TransactionError) Error() string {
	return "placeholder"
}
