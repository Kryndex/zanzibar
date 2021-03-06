// Code generated by thriftrw v1.3.0
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Bar_ArgNotStruct_Args struct {
	Request string `json:"request,required"`
}

func (v *Bar_ArgNotStruct_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Request), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Bar_ArgNotStruct_Args) FromWire(w wire.Value) error {
	var err error
	requestIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Request, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}
	if !requestIsSet {
		return errors.New("field Request of Bar_ArgNotStruct_Args is required")
	}
	return nil
}

func (v *Bar_ArgNotStruct_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++
	return fmt.Sprintf("Bar_ArgNotStruct_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_ArgNotStruct_Args) Equals(rhs *Bar_ArgNotStruct_Args) bool {
	if !(v.Request == rhs.Request) {
		return false
	}
	return true
}

func (v *Bar_ArgNotStruct_Args) MethodName() string {
	return "argNotStruct"
}

func (v *Bar_ArgNotStruct_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Bar_ArgNotStruct_Helper = struct {
	Args           func(request string) *Bar_ArgNotStruct_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*Bar_ArgNotStruct_Result, error)
	UnwrapResponse func(*Bar_ArgNotStruct_Result) error
}{}

func init() {
	Bar_ArgNotStruct_Helper.Args = func(request string) *Bar_ArgNotStruct_Args {
		return &Bar_ArgNotStruct_Args{Request: request}
	}
	Bar_ArgNotStruct_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}
	Bar_ArgNotStruct_Helper.WrapResponse = func(err error) (*Bar_ArgNotStruct_Result, error) {
		if err == nil {
			return &Bar_ArgNotStruct_Result{}, nil
		}
		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_ArgNotStruct_Result.BarException")
			}
			return &Bar_ArgNotStruct_Result{BarException: e}, nil
		}
		return nil, err
	}
	Bar_ArgNotStruct_Helper.UnwrapResponse = func(result *Bar_ArgNotStruct_Result) (err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}
		return
	}
}

type Bar_ArgNotStruct_Result struct {
	BarException *BarException `json:"barException,omitempty"`
}

func (v *Bar_ArgNotStruct_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("Bar_ArgNotStruct_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarException_Read(w wire.Value) (*BarException, error) {
	var v BarException
	err := v.FromWire(w)
	return &v, err
}

func (v *Bar_ArgNotStruct_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.BarException != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Bar_ArgNotStruct_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *Bar_ArgNotStruct_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}
	return fmt.Sprintf("Bar_ArgNotStruct_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Bar_ArgNotStruct_Result) Equals(rhs *Bar_ArgNotStruct_Result) bool {
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}
	return true
}

func (v *Bar_ArgNotStruct_Result) MethodName() string {
	return "argNotStruct"
}

func (v *Bar_ArgNotStruct_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
