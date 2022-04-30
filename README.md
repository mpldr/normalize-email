<!--
SPDX-FileCopyrightText: 2022 Moritz Poldrack
SPDX-License-Identifier: Unlicense
-->

# normalize-email

[![Go Documentation](https://godocs.io/mpldr.codes/normalize-email?status.svg)](https://godocs.io/mpldr.codes/normalize-email)
[![REUSE status](https://api.reuse.software/badge/git.sr.ht/~poldi1405/normalize-email)](https://api.reuse.software/info/git.sr.ht/~poldi1405/normalize-email)
[![License: MPL-2](https://img.shields.io/static/v1?color=green&label=License&message=MPL-2&logo=open-source-initiative&color=3DA639)](https://git.sr.ht/~poldi1405/normalize-email/tree/master/item/LICENSES/MPL-2.0.txt)

This package is intended to help reduce multiple sign-ups using the same
mail-address by applying general and provider specific rules to email
addresses, thereby transforming them into a standardised form.

With this `John.Doe+marketing@GoogleMail.com` becomes `johndoe@gmail.com`. As
these addresses are both for the same mail account, there is risk in this
method. The following transformations can be applied based on the address'
domain:

- Remove characters (for example the `.` in Gmail-Addresses)
- Remove subaddresses (usually the part after the +)
- Rewrite domains (`googlemail.com` → `gmail.com`)
- make localpart and domain lowercase (`JohnDoe@…` → `johndoe@…`)

## Missing rules

If you know about more rules, please [open an
issue](https://todo.sr.ht/~poldi1405/issues?title=[normalize-email]+rules+for+&description=%E2%80%A6+has+the+following+rules:%0A-+%E2%80%A6%0A-+%E2%80%A6%0A%0AThese+rules+are+also+documented+here:)
or [send it via
email](mailto:~poldi1405/issues@todo.sr.ht?cc=moritz%40poldrack.dev&subject=[normalize+email]+rules+for+%E2%80%A6&body=%E2%80%A6+has+the+following+rules:%0A-+%E2%80%A6%0A-+%E2%80%A6%0A%0AThese+rules+are+also+documented+here:)

## Important notes

***Do not*** send mails to the normalised addresses! Always use the
user-provided address as things like subaddresses and receiver-side sieves can
affect deliverability. Just because `joe+site@gmail.com` and `joe@gmail.com`
belong to the same account, does not mean that they are equal.

This library can not detect server side settings. If the subaddress delimiter
is set to `-`, this library will still assume it's `+`. There is simply no way
to do that. This may – though highly unlikely – lead to non-functional
addresses being produced. The low likelyhood of this happening to two users on
the same domain with the same odd address (`joe+doe@…` and `joe+deere@…`) makes
this something considered acceptably rare. Again: *do not send mails to
normalised addresses!*

## Stable version

This library is considered stable and considered feature-complete. The only
planned updates are updating or adding rules. To avoid creating a new tag at
every commit, no tags will be added to this repository and origin/HEAD should
be considered stable.

