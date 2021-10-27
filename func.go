package funny

// DoAndEarlyReturn do method (func() error) in order, if got an error, stop and return the error.
func DoAndEarlyReturn(fs...func() error) error {
	for _, f := range fs {
		if e := f(); e != nil {
			return e
		}
	}

	return nil
}
