// Copyright © 2022 Moritz Poldrack
//
// SPDX-License-Identifier: MPL-2.0

// This package has the purpose of normalizing emails in order to prevent
// multiple signups with the same email address. Please note, that mail should
// always be sent to the user provided address in order to allow automatic
// sorting and not restricting the user unnecessarily.
//
// By default adresses from hosts without special rules (see rules.go) have
// + subaddresses removed, and are transformed to lowercase. Special rules can
// be added per domain.
package normalizeemail
