// Copyright © 2022 Moritz Poldrack
//
// SPDX-License-Identifier: MPL-2.0

package normalizeemail

import (
	"net/mail"
	"strings"
)

// Normalize takes in the email and returns it in normalized form. This
// function does not perform validation and the normalized form of an invalid
// email address is an empty string.
func Normalize(email string) string {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}

	parts := strings.SplitN(addr.Address, "@", 2)
	localpart := parts[0]
	domain := parts[1]

	rulesMtx.RLock()
	defer rulesMtx.RUnlock()

	rule := rules[""]

	if r, exists := rules[strings.ToLower(domain)]; exists {
		rule.TrimChars = r.TrimChars
		rule.ChangeDomain = r.ChangeDomain
		rule.CaseSensitiveLocalpart = r.CaseSensitiveLocalpart
		if r.SubAddressDelimiter != "" {
			rule.SubAddressDelimiter = r.SubAddressDelimiter
		}

		if r.DisableDefaultRules {
			rule = r
		}
	}

	// trim unwanted characters
	for i := 0; i < len(rule.TrimChars); i++ {
		localpart = strings.ReplaceAll(localpart, string(rule.TrimChars[i]), "")
	}

	// remove subaddresses
	if rule.SubAddressDelimiter != "" {
		parts = strings.Split(localpart, rule.SubAddressDelimiter)
		localpart = parts[0]
	}

	// rewrite domain
	if rule.ChangeDomain != "" {
		domain = rule.ChangeDomain
	}
	domain = strings.ToLower(domain)

	// make localpart lowercase
	if !rule.CaseSensitiveLocalpart {
		localpart = strings.ToLower(localpart)
	}

	return localpart + "@" + domain
}
