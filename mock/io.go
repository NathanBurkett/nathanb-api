package mock

type Reader struct {
	ReadErr error
}

func (r Reader) Read(p []byte) (n int, err error) {
	if r.ReadErr != nil {
		return 0, r.ReadErr
	}

	return 0, nil
}
