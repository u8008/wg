// Copyright 2020 Jing Li. All rights reserved.

package test

import (
	"testing"
)

func TestOk(t *testing.T) {
	// Call a func to return value(s) and/or an error
	var e error

	/*
		Ok() makes the following test simpler and more readable
		if e != nil {
			t.Errorf("It failed for <%v>", e)
		}
	*/
	Ok(t, e == nil, "It failed for <%v>", e)
}
