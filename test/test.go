// Copyright 2020 Jing Li. All rights reserved.

/*
A wrapper of the built-in "testing" package to facilitate the unit-testing
*/
package test

import (
	"testing"
)

/*
If $c is false, it fails the test with the caller's filename and line #, and
logs the error with $fs and $as, or else it does nothing
*/
func Ok(t *testing.T, c bool, fs string, as ...interface{}) {
	// Safe to skip sanity checks
	t.Helper()
	if !c {
		t.Errorf(fs, as...)
	}
}
