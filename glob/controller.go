package glob

import (
	"reflect"
)

// Controllers ...
var Controllers []reflect.Value

// InitController ...
func InitController() error {
	controller, err := ParseSuffix("Controller")
	if err != nil {
		return err
	}

	for _, c := range controller {
		Controllers = append(Controllers, c.MethodByName("SetupRouter"))
	}

	return nil
}
