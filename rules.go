// Copyright © 2022 Moritz Poldrack
//
// SPDX-License-Identifier: MPL-2.0

package normalizeemail

import "sync"

type Rule struct {
	// TrimChars is a list of characters that are removed from the
	// localpart of the email
	//
	// john.doe@gmail.com → johndoe@gmail.com
	TrimChars []rune
	// SubAddressDelimiter specifies what character ends the localpart and
	// marks the beginning of the subaddress
	//
	// joe+service@example.com → joe@example.com
	SubAddressDelimiter string
	// ChangeDomain specifies the domain these email-addresses are to be
	// changed to. To avoid infinite recursions, the rules of the new
	// domain are not applied.
	//
	// joe@googlemail.com -> joe@gmail.com
	ChangeDomain string
	// CaseSensitiveLocalpart indicates that a localpart must be in the
	// case specified. Otherwise, all characters are made lowercase. This
	// does not affect the domain, which is always turned to lowercase.
	//
	// JOE@DOE.com -> JOE@doe.com
	CaseSensitiveLocalpart bool
	// DisableDefaultRules disabled the default rules to be applied, in
	// case a provider has different rules.
	DisableDefaultRules bool
}

// SetRule sets the rules applied to a domain. If the domain is an empty
// string, the rules overwrite the defaults.
func SetRule(domain string, rule Rule) {
	rulesMtx.Lock()
	defer rulesMtx.Unlock()

	rules[domain] = rule
}

var (
	rulesMtx sync.RWMutex
	rules    = map[string]Rule{
		"":                            DefaultRuleset,
		"gmail.com":                   GmailRuleset,
		"googlemail.com":              GooglemailRuleset,
		"yahoo.com":                   YahooRuleset,
		"cock.li":                     CockLiRuleset,
		"airmail.cc":                  CockLiRuleset,
		"420blaze.it":                 CockLiRuleset,
		"aaathats3as.com":             CockLiRuleset,
		"cumallover.me":               CockLiRuleset,
		"goat.si":                     CockLiRuleset,
		"horsefucker.org":             CockLiRuleset,
		"national.shitposting.agency": CockLiRuleset,
		"tfwno.gf":                    CockLiRuleset,
		"cock.lu":                     CockLiRuleset,
		"cock.email":                  CockLiRuleset,
		"firemail.cc":                 CockLiRuleset,
		"memeware.net":                CockLiRuleset,
		"cocaine.ninja":               CockLiRuleset,
		"waifu.club":                  CockLiRuleset,
		"dicksinhisan.us":             CockLiRuleset,
		"loves.dicksinhisan.us":       CockLiRuleset,
		"wants.dicksinhisan.us":       CockLiRuleset,
		"dicksinmyan.us":              CockLiRuleset,
		"loves.dicksinmyan.us":        CockLiRuleset,
		"wants.dicksinmyan.us":        CockLiRuleset,
	}
)

var (
	DefaultRuleset = Rule{
		SubAddressDelimiter: "+",
	}
	GmailRuleset = Rule{
		TrimChars: []rune{'.'},
	}
	GooglemailRuleset = Rule{
		TrimChars:    []rune{'.'},
		ChangeDomain: "gmail.com",
	}
	YahooRuleset = Rule{
		TrimChars:           []rune{'.'},
		SubAddressDelimiter: "-",
	}
	CockLiRuleset = Rule{
		SubAddressDelimiter: ".",
	}
)
