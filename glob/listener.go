package glob

import (
	"reflect"
)

// Listeners ...
var Listeners []reflect.Value

// InitListener ...
func InitListener() error {
	listeners, err := ParseSuffix("Listener")
	if err != nil {
		return err
	}

	Listeners = append(Listeners, listeners...)
	return nil
}
