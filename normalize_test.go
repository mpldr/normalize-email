// Copyright © 2022 Moritz Poldrack
//
// SPDX-License-Identifier: MPL-2.0

package normalizeemail

import (
	"fmt"
	"testing"
)

func ExampleNormalize() {
	fmt.Println(Normalize("j.Oe+some-subaddress@googlemail.com"))
	// Output: joe@gmail.com
}

func TestNormalize(t *testing.T) {
	testValues := map[string]string{
		"j.o.e@gmail.com":     "joe@gmail.com",
		"j.Oe@googlemail.com": "joe@gmail.com",
		"j.o.e+some-subaddress+with+extra+plus@gmail.com": "joe@gmail.com",
		"j.o.e+subaddress@googlemail.com":                 "joe@gmail.com",
		"JOE@GMAIL.COM":                                   "joe@gmail.com",
		"j.o.e@yahoo.com":                                 "joe@yahoo.com",
		"joe-is+using-sub.addresses@yahoo.com":            "joe@yahoo.com",
		"moritz+git@poldrack.dev":                         "moritz@poldrack.dev",
	}

	for in, out := range testValues {
		if res := Normalize(in); res != out {
			t.Logf("expected: %s; got: %s; input: %s", out, res, in)
			t.Fail()
		}
	}
}
