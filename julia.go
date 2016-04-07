package julia

/*
#cgo CFLAGS: -fPIC -DJULIA_INIT_DIR="/Applications/Julia-0.4.5.app/Contents/Resources/julia/lib" -I/Applications/Julia-0.4.5.app/Contents/Resources/julia/include/julia
#cgo LDFLAGS: -L/Applications/Julia-0.4.5.app/Contents/Resources/julia/lib/julia -Wl,-rpath,/Applications/Julia-0.4.5.app/Contents/Resources/julia/lib/julia -ljulia
#include <julia.h>

int my_jl_is_nothing(jl_value_t* t) {
  return jl_is_nothing(t);
}

int my_jl_is_tuple(jl_value_t* t) {
  return jl_is_tuple(t);
}

int my_jl_is_array(jl_value_t* t) {
  return jl_is_array(t);
}

int my_jl_array_len(jl_value_t* t) {
  return jl_array_len(t);
}

int my_jl_nfields(jl_value_t* t) {
  return jl_nfields(t);
}

int my_jl_is_float64(jl_value_t* t) {
  return jl_is_float64(t);
}

int my_jl_is_int64(jl_value_t* t) {
  return jl_is_int64(t);
}

int my_jl_is_int32(jl_value_t *t) {
  return jl_is_int32(t);
}

int my_jl_is_int8(jl_value_t *t) {
  return jl_is_int8(t);
}

int my_jl_is_utf8_string(jl_value_t *t) {
  return jl_is_utf8_string(t);
}

int my_jl_is_ascii_string(jl_value_t *t) {
  return jl_is_ascii_string(t);
}

int my_jl_is_float32(jl_value_t *t) {
  return jl_is_float32(t);
}

jl_value_t * my_jl_typeof(jl_value_t *t) {
  return jl_typeof(t);
}

char * my_jl_string_data(jl_value_t *t) {
  return jl_string_data(t);
}

*/
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type Julia struct {
}

func New() *Julia {
	if C.jl_is_initialized() == C.int(0) {
		init_dir := C.CString("/Applications/Julia-0.4.5.app/Contents/Resources/julia/lib")
		C.jl_init(init_dir)
		C.free(unsafe.Pointer(init_dir))
	}
	t := &Julia{}
	return t
}

func (this *Julia) Eval(in string) (res interface{}, err error) {

	ptr := C.CString(in)
	defer C.free(unsafe.Pointer(ptr))

	ret := C.jl_eval_string(ptr)

	if C.jl_exception_occurred() != nil {
		ex := C.jl_typeof_str(C.jl_exception_occurred())
		s := C.GoString(ex)
		return nil, errors.New(s)
	}

	fmt.Println(reflect.TypeOf(ret))

	if C.my_jl_is_nothing(ret) > C.int(0) {
		return nil, nil
	}

	if C.my_jl_is_array(ret) > C.int(0) {
		fmt.Println("array")
		if C.my_jl_array_len(ret) == 0 {
			return nil, nil
		}
	}

	if C.my_jl_is_tuple(ret) > C.int(0) {
		fmt.Println("tuple")
	}

	if C.my_jl_is_ascii_string(ret) > C.int(0) {
		c := C.GoString(C.my_jl_string_data(ret))
		return c, nil
	}

	if C.my_jl_is_utf8_string(ret) > C.int(0) {
		fmt.Println("utf8_string")
	}

	if C.my_jl_is_float64(ret) > C.int(0) {
		fmt.Println("float64")
		return float64(C.jl_unbox_float64(ret)), nil
	} else if C.my_jl_is_int64(ret) > C.int(0) {
		fmt.Println("int64")
		return int64(C.jl_unbox_int64(ret)), nil
	} else if C.my_jl_is_int32(ret) > C.int(0) {
		fmt.Println("int32")
		return C.jl_unbox_int32(ret), nil
	} else if C.my_jl_is_int8(ret) > C.int(0) {
		fmt.Println("int8")
		return C.jl_unbox_int8(ret), nil
	} else if C.my_jl_is_float32(ret) > C.int(0) {
		fmt.Println("float32")
	}

	return nil, nil
}

func (this *Julia) Free() {
	C.jl_atexit_hook(0)
}
