package nopio

// NopCloser represents Closer interface with no-op Close method.
type NopCloser struct{}

// Close will return nil.
func (NopCloser) Close() error {
	return nil
}

// NopWriter represents Writer interface with no-op Write method.
type NopWriter struct{}

// Write will return len(data) and nil error.
func (NopWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

// NopReader represents Reader interface with no-op Reade method.
type NopReader struct{}

// Read will return zero read bytes and nil error.
func (NopReader) Read(data []byte) (int, error) {
	return 0, nil
}

// NopWriteCloser represents no-op Writer and Closer.
type NopWriteCloser struct {
	NopWriter
	NopCloser
}

// NopReadCloser represents no-op Reader and Closer.
type NopReadCloser struct {
	NopReader
	NopCloser
}
