package cerrors

type TransactionNotMatchError struct {
}

func (t TransactionNotMatchError) Error() string {
	return "transaction string not match with pattern"
}
