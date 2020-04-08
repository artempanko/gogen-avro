// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
package avro

import (
	"fmt"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

func writeMapLong(r map[string]int64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteLong(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapLongWrapper struct {
	Target *map[string]int64
	keys   []string
	values []int64
}

func (_ *MapLongWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapLongWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapLongWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapLongWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapLongWrapper) Finalize() {
	fmt.Printf("Finalizing!\n")
	for i := range r.keys {
		fmt.Printf("%v\n", r.keys[i])
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapLongWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int64
	r.values = append(r.values, v)
	return &types.Long{Target: &r.values[len(r.values)-1]}
}

func (_ *MapLongWrapper) AppendArray() types.Field { panic("Unsupported operation") }
