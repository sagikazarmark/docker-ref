// Package reference provides a general type to represent any way of referencing images within the registry.
// Its main purpose is to abstract tags and digests (content-addressable hash).
//
// Grammar
//
//	reference                       := name [ ":" tag ] [ "@" digest ]
//	name                            := [domain '/'] remote-name
//	domain                          := host [':' port-number]
//	host                            := domain-name | IPv4address | \[ IPv6address \]	; rfc3986 appendix-A
//	domain-name                     := domain-component ['.' domain-component]*
//	domain-component                := /([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])/
//	port-number                     := /[0-9]+/
//	path-component                  := alpha-numeric [separator alpha-numeric]*
//	path (or "remote-name")         := path-component ['/' path-component]*
//	alpha-numeric                   := /[a-z0-9]+/
//	separator                       := /[_.]|__|[-]*/
//
//	tag                             := /[\w][\w.-]{0,127}/
//
//	digest                          := digest-algorithm ":" digest-hex
//	digest-algorithm                := digest-algorithm-component [ digest-algorithm-separator digest-algorithm-component ]*
//	digest-algorithm-separator      := /[+.-_]/
//	digest-algorithm-component      := /[A-Za-z][A-Za-z0-9]*/
//	digest-hex                      := /[0-9a-fA-F]{32,}/ ; At least 128 bit digest value
//
//	identifier                      := /[a-f0-9]{64}/
package reference

import (
	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
)

const (
	// NameTotalLengthMax is the maximum total number of characters in a repository name.
	NameTotalLengthMax = reference.RepositoryNameTotalLengthMax
)

var (
	// ErrReferenceInvalidFormat represents an error while trying to parse a string as a reference.
	ErrReferenceInvalidFormat = reference.ErrReferenceInvalidFormat

	// ErrTagInvalidFormat represents an error while trying to parse a string as a tag.
	ErrTagInvalidFormat = reference.ErrTagInvalidFormat

	// ErrDigestInvalidFormat represents an error while trying to parse a string as a tag.
	ErrDigestInvalidFormat = reference.ErrDigestInvalidFormat

	// ErrNameContainsUppercase is returned for invalid repository names that contain uppercase characters.
	ErrNameContainsUppercase = reference.ErrNameContainsUppercase

	// ErrNameEmpty is returned for empty, invalid repository names.
	ErrNameEmpty = reference.ErrNameEmpty

	// ErrNameTooLong is returned when a repository name is longer than NameTotalLengthMax.
	ErrNameTooLong = reference.ErrNameTooLong

	// ErrNameNotCanonical is returned when a name is not canonical.
	ErrNameNotCanonical = reference.ErrNameNotCanonical
)

// Reference is an opaque object reference identifier that may include
// modifiers such as a hostname, name, tag, and digest.
type Reference = reference.Reference

// Field provides a wrapper type for resolving correct reference types when
// working with encoding.
type Field = reference.Field

// AsField wraps a reference in a Field for encoding.
func AsField(ref Reference) Field {
	return reference.AsField(ref)
}

// Named is an object with a full name
type Named = reference.Named

// Tagged is an object which has a tag
type Tagged = reference.Tagged

// NamedTagged is an object including a name and tag.
type NamedTagged = reference.NamedTagged

// Digested is an object which has a digest
// in which it can be referenced by
type Digested = reference.Digested

// Canonical reference is an object with a fully unique
// name including a name with domain and digest
type Canonical = reference.Canonical

// Domain returns the domain part of the [Named] reference.
func Domain(named Named) string {
	return reference.Domain(named)
}

// Path returns the name without the domain part of the [Named] reference.
func Path(named Named) (name string) {
	return reference.Path(named)
}

// SplitHostname splits a named reference into a
// hostname and name string. If no valid hostname is
// found, the hostname is empty and the full value
// is returned as name
//
// Deprecated: Use [Domain] or [Path].
func SplitHostname(named Named) (string, string) {
	return reference.Domain(named), reference.Path(named)
}

// Parse parses s and returns a syntactically valid Reference.
// If an error was encountered it is returned, along with a nil Reference.
func Parse(s string) (Reference, error) {
	return reference.Parse(s)
}

// ParseNamed parses s and returns a syntactically valid reference implementing
// the Named interface. The reference must have a name and be in the canonical
// form, otherwise an error is returned.
// If an error was encountered it is returned, along with a nil Reference.
func ParseNamed(s string) (Named, error) {
	return reference.ParseNamed(s)
}

// WithName returns a named object representing the given string. If the input
// is invalid ErrReferenceInvalidFormat will be returned.
func WithName(name string) (Named, error) {
	return reference.WithName(name)
}

// WithTag combines the name from "name" and the tag from "tag" to form a
// reference incorporating both the name and the tag.
func WithTag(name Named, tag string) (NamedTagged, error) {
	return reference.WithTag(name, tag)
}

// WithDigest combines the name from "name" and the digest from "digest" to form
// a reference incorporating both the name and the digest.
func WithDigest(name Named, digest digest.Digest) (Canonical, error) {
	return reference.WithDigest(name, digest)
}

// TrimNamed removes any tag or digest from the named reference.
func TrimNamed(ref Named) Named {
	return reference.TrimNamed(ref)
}
