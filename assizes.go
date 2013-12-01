//
// Assertions and Helpers for the go test framework
//
package assizes

import (
	"fmt"
	"runtime"
	"testing"
)

type Assize struct {
	T *testing.T
}

func (a Assize) AssertTrue(b bool) {
	if !b {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("assertTrue failed: line %v", line)
	}
}

func (a Assize) AssertFalse(b bool) {
	if b {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("assertFalse failed: line %v", line)
	}
}

func (a Assize) AssertEqual(exp, act interface{}) {
	sexp := fmt.Sprintf("%v", exp)
	sact := fmt.Sprintf("%v", act)
	if sexp != sact {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("assertEqual failed (line %d). exp: %s; act: %s",
			line, sexp, sact)
	}
}

func (a Assize) AssertNotNil(act interface{}) {
	if act == nil {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("assertNotNil failed (line %d).", line)
	}
}

func (a Assize) AssertNil(act interface{}) {
	if act != nil {
		_, _, line, _ := runtime.Caller(1)
		a.T.Errorf("assertNil failed (line %d); act = %v", line, act)
	}
}

func (a Assize) AssertIntSlicesEqual(exp, act []int) {
	_, _, line, _ := runtime.Caller(1)
	if len(exp) != len(act) {
		a.T.Errorf("assertIntSlicesEqual failed (line %d). Lengths differ: len(exp): %d; len(act): %d",
			line, len(exp), len(act))
		return
	}
	for i, v := range exp {
		if v != act[i] {
			a.T.Errorf("assertIntSlicesEqual failed (line %d). Differ at index %d: exp[%d]: %d; act[%d]: %d",
				line, i, i, v, i, act[i])
		}
	}
}
