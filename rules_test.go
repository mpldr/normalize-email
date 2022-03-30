// Copyright © 2022 Moritz Poldrack
//
// SPDX-License-Identifier: MPL-2.0

package normalizeemail

import (
	"fmt"
	"testing"
)

func ExampleSetRule() {
	SetRule("some-provider.com", Rule{
		TrimChars:           []rune{'.'},
		SubAddressDelimiter: "-",
		ChangeDomain:        "new-provider-address.com",
	})
	fmt.Println(Normalize("eric.hoffman-subaddress@some-proviDER.Com"))
	// Output: erichoffman@new-provider-address.com
}

func TestCustomRules(t *testing.T) {
	testValues := map[string]string{
		"j.o.e@gmail.com":              "j.o.e@gmail.com",
		"J.OE@GMAIL.COM":               "j.oe@gmail.com",
		"moritz+git@poldrack.dev":      "moritz+git@poldrack.dev",
		"SomeStrangeMail@provider.com": "SomeStrangeMail@provider.com",
		"SomeStrangeMail@PROVIDER.com": "SomeStrangeMail@provider.com",
	}

	SetRule("provider.com", Rule{
		CaseSensitiveLocalpart: true,
	})

	SetRule("poldrack.dev", Rule{
		DisableDefaultRules: true,
	})

	SetRule("gmail.com", Rule{})

	for in, out := range testValues {
		if res := Normalize(in); res != out {
			t.Logf("expected: %s; got: %s; input: %s", out, res, in)
			t.Fail()
		}
	}
}
