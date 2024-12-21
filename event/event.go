package event

import (
	"errors"
	"reflect"
	"sync"
)

type event struct {
	sync.RWMutex

	events map[string]interface{}
}

type Dispatcher interface {
	On(name string, fn interface{}) error
	Cmd(name string, params ...interface{}) error
	Send(name string, params ...interface{}) error
	Recv(name string, params ...interface{}) ([]byte, error)
	Has(name string) bool
	Remove(names string)
}

func New() Dispatcher {
	return &event{
		events: make(map[string]interface{}),
	}
}

func (e *event) On(name string, fn interface{}) error {
	e.Lock()
	defer e.Unlock()

	if fn == nil {
		return errors.New("fn is nil")
	}

	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return errors.New("fn is not a function")
	}

	if t.NumOut() < 1 {
		return errors.New("fn must have one return value")
	}
	//if t.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
	//	return errors.New("fn must return an error message")
	//}

	if f, ok := e.events[name]; ok {
		tt := reflect.TypeOf(f)
		if tt.NumIn() != t.NumIn() {
			return errors.New("fn signature is not equal")
		}
		for i := 0; i < tt.NumIn(); i++ {
			if tt.In(i) != t.In(i) {
				return errors.New("fn signature is not equal")
			}
		}
	}
	e.events[name] = fn

	return nil
}

func (e *event) Cmd(name string, params ...interface{}) error {
	e.RLock()
	defer e.RUnlock()

	fn := e.events[name]

	err := e.call(fn, params...)
	if err != nil {
		return err
	}
	return nil
}

func (e *event) Send(name string, params ...interface{}) error {
	e.RLock()
	defer e.RUnlock()

	fn := e.events[name]
	if fn == nil {
		return errors.New("function " + name + " is nil")
	}

	retValue, err := e.callFun(fn, params...)
	if err != nil {
		return err
	}

	retValueLen := len(retValue)
	if retValueLen == 0 {
		return errors.New("no value return from function")
	}

	if err, ok := retValue[retValueLen-1].(error); ok {
		return err
	}

	return nil
}

func (e *event) Recv(name string, params ...interface{}) ([]byte, error) {
	e.RLock()
	defer e.RUnlock()

	fn := e.events[name]
	if fn == nil {
		return nil, errors.New("function " + name + " is nil")
	}

	retValue, err := e.callFun(fn, params...)
	if err != nil {
		return nil, err
	}

	retValueLen := len(retValue)
	if retValueLen != 2 {
		return nil, errors.New("function must return two values")
	}

	var ok bool
	var data []byte
	if data, ok = retValue[0].([]byte); !ok {
		return nil, errors.New("function must return []byte as the first value")
	}

	if err, ok := retValue[1].(error); ok {
		return data, err
	}

	return data, nil
}

func (e *event) callFun(fn interface{}, params ...interface{}) (retValue []interface{}, err error) {
	var (
		f     = reflect.ValueOf(fn)
		t     = f.Type()
		numIn = t.NumIn()
		in    = make([]reflect.Value, 0, numIn)
	)
	// 如果是可变参数，则最后一个参数是可变参数
	if t.IsVariadic() {
		n := numIn - 1
		if len(params) < n {
			return nil, errors.New("parameters mismatched")
		}

		for _, param := range params[:n] {
			in = append(in, reflect.ValueOf(param))
		}
		s := reflect.MakeSlice(t.In(n), 0, len(params[n:]))
		for _, param := range params[n:] {
			s = reflect.Append(s, reflect.ValueOf(param))
		}
		in = append(in, s)
		out := f.CallSlice(in)
		for _, v := range out {
			retValue = append(retValue, v.Interface())
		}
		return retValue, nil
	}

	// 如果是固定参数，则参数个数必须和函数的参数个数相同
	if len(params) != numIn {
		return nil, errors.New("parameters mismatched")
	}

	for _, param := range params {
		in = append(in, reflect.ValueOf(param))
	}
	out := f.Call(in)
	for _, v := range out {
		retValue = append(retValue, v.Interface())
	}

	return retValue, nil
}

func (e *event) call(fn interface{}, params ...interface{}) (err error) {
	var (
		f     = reflect.ValueOf(fn)
		t     = f.Type()
		numIn = t.NumIn()
		in    = make([]reflect.Value, 0, numIn)
	)

	if t.IsVariadic() {
		n := numIn - 1
		if len(params) < n {
			return errors.New("parameters mismatched")
		}
		for _, param := range params[:n] {
			in = append(in, reflect.ValueOf(param))
		}
		s := reflect.MakeSlice(t.In(n), 0, len(params[n:]))
		for _, param := range params[n:] {
			s = reflect.Append(s, reflect.ValueOf(param))
		}
		in = append(in, s)

		err, _ = f.CallSlice(in)[0].Interface().(error)
		return err
	}

	if len(params) != numIn {
		return errors.New("parameters mismatched")
	}
	for _, param := range params {
		in = append(in, reflect.ValueOf(param))
	}

	err, _ = f.Call(in)[0].Interface().(error)
	return err
}

func (e *event) Has(name string) bool {
	e.RLock()
	defer e.RUnlock()
	_, ok := e.events[name]
	return ok
}

func (e *event) Remove(name string) {
	e.Lock()
	defer e.Unlock()

	delete(e.events, name)
}
