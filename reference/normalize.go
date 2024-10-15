package reference

import (
	"github.com/distribution/reference"
)

// ParseNormalizedNamed parses a string into a named reference
// transforming a familiar name from Docker UI to a fully
// qualified reference. If the value may be an identifier
// use ParseAnyReference.
func ParseNormalizedNamed(s string) (Named, error) {
	return reference.ParseNormalizedNamed(s)
}

// ParseDockerRef normalizes the image reference following the docker convention,
// which allows for references to contain both a tag and a digest. It returns a
// reference that is either tagged or digested. For references containing both
// a tag and a digest, it returns a digested reference. For example, the following
// reference:
//
//	docker.io/library/busybox:latest@sha256:7cc4b5aefd1d0cadf8d97d4350462ba51c694ebca145b08d7d41b41acc8db5aa
//
// Is returned as a digested reference (with the ":latest" tag removed):
//
//	docker.io/library/busybox@sha256:7cc4b5aefd1d0cadf8d97d4350462ba51c694ebca145b08d7d41b41acc8db5aa
//
// References that are already "tagged" or "digested" are returned unmodified:
//
//	// Already a digested reference
//	docker.io/library/busybox@sha256:7cc4b5aefd1d0cadf8d97d4350462ba51c694ebca145b08d7d41b41acc8db5aa
//
//	// Already a named reference
//	docker.io/library/busybox:latest
func ParseDockerRef(ref string) (Named, error) {
	return reference.ParseDockerRef(ref)
}

// TagNameOnly adds the default tag "latest" to a reference if it only has
// a repo name.
func TagNameOnly(ref Named) Named {
	return reference.TagNameOnly(ref)
}

// ParseAnyReference parses a reference string as a possible identifier,
// full digest, or familiar name.
func ParseAnyReference(ref string) (Reference, error) {
	return reference.ParseAnyReference(ref)
}
