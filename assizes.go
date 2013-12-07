//
// Assertions and Helpers for the go test framework
//
package assizes

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

type Assize struct {
	T *testing.T
}

func (a Assize) AssertTrue(b bool) {
	if !b {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertTrue failed: line %v", line)
	}
}

func (a Assize) AssertFalse(b bool) {
	if b {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertFalse failed: line %v", line)
	}
}

func (a Assize) AssertEqual(exp, act interface{}) {
	sexp := fmt.Sprintf("%v", exp)
	sact := fmt.Sprintf("%v", act)
	if sexp != sact {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertEqual failed (line %d). exp: %s; act: %s",
			line, sexp, sact)
	}
}

func (a Assize) AssertNotNil(act interface{}) {
	if act == nil {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertNotNil failed (line %d).", line)
	}
}

func (a Assize) AssertNil(act interface{}) {
	if act != nil {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertNil failed (line %d); act = %v", line, act)
	}
}

func (a Assize) AssertIntSlicesEqual(exp, act []int) {
	_, _, line, _ := runtime.Caller(1)
	if len(exp) != len(act) {
		a.T.Errorf("AssertIntSlicesEqual failed (line %d). Lengths differ: len(exp): %d; len(act): %d",
			line, len(exp), len(act))
		return
	}
	for i, v := range exp {
		if v != act[i] {
			a.T.Errorf("AssertIntSlicesEqual failed (line %d). Differ at index %d: exp[%d]: %d; act[%d]: %d",
				line, i, i, v, i, act[i])
		}
	}
}

func (a Assize) AssertIntSlicesNotEqual(exp, act []int) {
	if len(exp) != len(act) {
		// if not same length, then not equal, so assertion passes; just return
		return
	}
	for i, v := range exp {
		if v != act[i] {
			return
		}
	}
	// if get here the slices have equal size and values
	_, _, line, _ := runtime.Caller(1)
	a.T.Errorf("AssertIntSlicesNotEqual failed (line %d). Len and values all match", line)
}

//
// AssertSlicesNotEqual assumes that the values in the interface{}
// can be successfully compared to each other with ==.
//
func (a Assize) AssertSlicesNotEqual(exp, act []interface{}) {
	if len(exp) != len(act) {
		// if not same length, then not equal, so assertion passes; just return
		return
	}
	for i, v := range exp {
		if v != act[i] {
			return
		}
	}
	// if get here the slices have equal size and values
	_, _, line, _ := runtime.Caller(1)
	a.T.Errorf("AssertSlicesNotEqual failed (line %d). Len and values all match", line)
}

//
// AssertSlicesEqual assumes that the values in the interface{}
// can be successfully compared to each other with == (or rather !=).
//
func (a Assize) AssertSlicesEqual(exp, act []interface{}) {
	_, _, line, _ := runtime.Caller(1)
	if len(exp) != len(act) {
		a.T.Errorf("AssertSlicesEqual failed (line %d). Lengths differ: len(exp): %d; len(act): %d",
			line, len(exp), len(act))
		return
	}
	for i, v := range exp {
		if v != act[i] {
			a.T.Errorf("AssertSlicesEqual failed (line %d). Differ at index %d: exp[%d]: %d; act[%d]: %d",
				line, i, i, v, i, act[i])
		}
	}
}


//
// AssertSlicesDeepEqual uses reflect.DeepEqual. If the slices are the same
// size but not equal, no information is printed out as to where they differ
// since reflect.DeepEqual does not provide any information about that.
//
func (a Assize) AssertSlicesDeepEqual(exp, act []interface{}) {
	_, _, line, _ := runtime.Caller(1)
	if len(exp) != len(act) {
		a.T.Errorf("AssertSlicesDeepEqual failed (line %d). Lengths differ: len(exp): %d; len(act): %d",
			line, len(exp), len(act))
		return
	}
	if !reflect.DeepEqual(exp, act) {
		a.T.Errorf("AssertSlicesDeepEqual failed (line %d).", line)
	}
}

func (a Assize) AssertSliceNotEmptyOrNil(act []interface{}) {
	if act == nil || len(act) == 0 {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertSliceNotEmptyOrNil failed (line %d); act = %v", line, act)
	}
}

func (a Assize) AssertSliceEmptyOrNil(act []interface{}) {
	if act != nil && len(act) != 0 {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("AssertSliceEmptyOrNil failed (line %d); act = %v", line, act)
	}
}
