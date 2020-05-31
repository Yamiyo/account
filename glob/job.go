package glob

import (
	"reflect"
)

var Jobs []reflect.Value

func InitJob() error {
	jobs, err := ParseSuffix("Job")
	if err != nil {
		return err
	}

	Jobs = append(Jobs, jobs...)
	return nil
}
