package glob

import (
	"errors"
	"fmt"
	"github.com/Yamiyo/account/glob/consts"
	"reflect"
	"strings"
	"sync"
	"unsafe"
)

// Injector ...
type Injector struct {
	rwlock sync.RWMutex
	objs   map[string]reflect.Value
}

var injector *Injector

func init() {
	injector = &Injector{
		objs: make(map[string]reflect.Value, 0),
	}
}

func (i *Injector) has(key string) (bool, error) {
	i.rwlock.Lock()
	defer i.rwlock.Unlock()

	if injector.objs == nil {
		return false, errors.New("has fail, the injector with nil objs")
	}

	_, ok := i.objs[key]

	return ok, nil
}

func (i *Injector) put(key string, value interface{}) error {
	i.rwlock.Lock()
	defer i.rwlock.Unlock()

	if injector.objs == nil {
		return errors.New("put fail, the injector with nil objs")
	}

	i.objs[key] = reflect.ValueOf(value)

	return nil
}

func (i *Injector) delete(key string) error {
	i.rwlock.Lock()
	defer i.rwlock.Unlock()

	if injector.objs == nil {
		return errors.New("delete fail, the injector with nil objs")
	}

	delete(i.objs, key)

	return nil
}

func (i *Injector) get(key string) (reflect.Value, error) {
	i.rwlock.Lock()
	defer i.rwlock.Unlock()

	if injector.objs == nil {
		return reflect.Value{}, errors.New("get fail, the injector with nil objs")
	}

	val, _ := i.objs[key]

	return val, nil
}

// Register ...
func Register(name string, v interface{}) error {
	if injector == nil {
		return errors.New("fail with nil injector")
	}

	ok, err := injector.has(name)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("AutoRegister fail, the entry key object already exist")
	}

	return injector.put(name, v)
}

// AutoRegister ...
func AutoRegister(value interface{}) error {
	if injector == nil {
		return errors.New("fail with nil injector")
	}

	typ := reflect.TypeOf(value).String()
	sub := "."

	if !strings.Contains(typ, sub) {
		return fmt.Errorf("AutoRegister fail, unexpected type %v", typ)
	}

	tSlice := strings.Split(typ, ".")
	k := tSlice[len(tSlice)-1:][0]

	ok, err := injector.has(k)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("AutoRegister fail, the entry key object already exist")
	}

	return injector.put(k, value)
}

// Remove ...
func Remove(key string) error {
	if injector == nil {
		return errors.New("fail with nil injector")
	}

	ok, err := injector.has(key)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("Remove fail, the entry key object not exist")
	}

	return injector.delete(key)
}

// ParseSuffix ...
func ParseSuffix(suffix string) ([]reflect.Value, error) {
	if injector == nil {
		return nil, errors.New("fail with nil injector")
	}

	injector.rwlock.Lock()
	defer injector.rwlock.Unlock()

	ret := make([]reflect.Value, 0)
	for k, v := range injector.objs {
		if strings.HasSuffix(k, suffix) {
			ret = append(ret, v)
		}
	}

	return ret, nil
}

// Inject ...
func Inject() error {
	for _, v := range injector.objs {
		value := v
		// 執行初始化 function
		initFunction := value.MethodByName("Init")
		if initFunction.IsValid() {
			if err := initFunction.Interface().(func() error)(); err != nil {
				return err
			}
		}
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		for i := 0; i < value.NumField(); i++ {
			name := value.Type().Field(i).Tag.Get(consts.INJECTION)
			temp, ok := injector.objs[name]
			if ok && reflect.TypeOf(temp).Kind() == reflect.Struct {
				field := value.Field(i)
				if field.CanSet() {
					field.Set(temp)
				} else {
					field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
					field.Set(temp)
				}
			} else if name != "" {
				return errors.New("injection error not find " + name)
			}
		}
	}
	return nil
}
