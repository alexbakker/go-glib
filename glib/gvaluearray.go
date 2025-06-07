package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"runtime"
	"unsafe"
)

type ValueArray struct {
	GValueArray *C.GValueArray
}

// native returns a pointer to the underlying GValue.
func (a *ValueArray) native() *C.GValueArray {
	return a.GValueArray
}

// Unsafe returns an unsafe.Pointer to the underlying GValue.
func (a *ValueArray) Unsafe() unsafe.Pointer {
	return unsafe.Pointer(a.native())
}

// Native refturns a pointer to the underlying GValue.
func (a *ValueArray) Native() unsafe.Pointer {
	return unsafe.Pointer(a.native())
}

func (a *ValueArray) Free() {
	C.g_value_array_free(a.native())
}

func (a *ValueArray) Append(v *Value) {
	C.g_value_array_append(a.native(), v.native())
}

func NewValueArray(preAllocSize uint) (*ValueArray, error) {
	c := C.g_value_array_new(C.guint(preAllocSize))
	if c == nil {
		return nil, errNilPtr
	}

	v := &ValueArray{c}
	runtime.SetFinalizer(v, (*ValueArray).Free)
	return v, nil
}
