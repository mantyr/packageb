package packageb

type B struct {}

func (b B) Check() error {
	return errors.New("test")
}
