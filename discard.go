package watchman

type null struct{}

func Discard() Logger {
	return null{}
}

func (null) Close() error {
	return nil
}

func (null) Log(_ ...any) {}
