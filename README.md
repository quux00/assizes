# Overview

The assizes library is a helper library for the standard Go testing package.  It provides some basic `assert` functionality, allowing your Go tests to be less verbose.  The assert style is patterned after Java's JUnit style.

Only a few helpers have been added so far - the ones I've needed for my tests.  More will be added.  Feel free to add your own.

The library name refers to the old British [Courts of Assizes](https://en.wikipedia.org/wiki/Assizes), which, being an American, I used to read and wonder about in my college days while reading lots of Dickens, Thackeray and George Eliot.


# Install

    go get quux00/assizes


# Usage

Wrap the `testing.T` pointer in an Assize and then use the `Assert` methods on the Assize.  Any errors will be reported as failing on the line you called the assert from.

    import (
        az "github.com/quux00/assizes"
        "testing"
    )
    
    func TestQueue(t *testing.T) {
    	a := az.Assize{t}
    
    	q := NewQueue()
        a.AssertNotNil(q)
        a.AssertTrue(q.IsEmpty())
        a.AssertEqual(0, q.Size())

    	q.Enqueue("a")
    	q.Enqueue("b")
        
        a.AssertFalse(q.IsEmpty())
        var (
            v interface{}
            err error
        )
        v, err = q.Dequeue()
        a.AssertNil(err)
        a.AssertEqual("a", v)

        // &etc.
    }


# LICENSE

Released under the [MIT License](http://opensource.org/licenses/MIT).
