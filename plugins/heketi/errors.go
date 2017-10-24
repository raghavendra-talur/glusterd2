package heketi

// ErrHeketiSingleError only error for now
type ErrHeketiSingleError struct{}

func (e *ErrHeketiSingleError) Error() string {
	return "rtalur heketi error"
}
