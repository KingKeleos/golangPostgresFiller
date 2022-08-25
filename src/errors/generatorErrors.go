package errors

type WrongColumnTypes struct{}

func (WCT *WrongColumnTypes) Error() string {
	return "The Column did not match either text, nor Integer"
}
