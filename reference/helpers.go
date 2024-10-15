package reference

import (
	"github.com/distribution/reference"
)

// IsNameOnly returns true if reference only contains a repo name.
func IsNameOnly(ref Named) bool {
	return reference.IsNameOnly(ref)
}

// FamiliarName returns the familiar name string
// for the given named, familiarizing if needed.
func FamiliarName(ref Named) string {
	return reference.FamiliarName(ref)
}

// FamiliarString returns the familiar string representation
// for the given reference, familiarizing if needed.
func FamiliarString(ref Reference) string {
	return reference.FamiliarString(ref)
}

// FamiliarMatch reports whether ref matches the specified pattern.
// See [path.Match] for supported patterns.
func FamiliarMatch(pattern string, ref Reference) (bool, error) {
	return reference.FamiliarMatch(pattern, ref)
}
