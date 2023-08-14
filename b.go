package packageb

import (
	"errors"
)

type B struct {}

func (b B) Check() (error, error){
	return errors.New("test"), errors.New("test2")
}
