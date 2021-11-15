package app

const (
	// These constants define the application semantic
	// and follow the semantic versioning 2.0.0 spec (http://semver.org/).
	verMajor string = "0"
	verMinor string = "1"
	verPatch string = "0"

	// verPreRelease flag to append suffix contain "-dev"
	// per the semantic versioning spec.
	verPreRelease = true

	// verPrefix use this prefix on version printing.
	verPrefix = "v"

	// verSeparator use this separator on version printing.
	verSeparator = "."
)

// version returns the application version as a properly formed string
// per the semantic versioning 2.0.0 spec (http://semver.org/).
func version() string {
	// start with the major, minor, and patch versions.
	ver := verPrefix + verMajor + verSeparator + verMinor + verSeparator + verPatch
	// append pre-release if there is one.
	if verPreRelease {
		ver += "-dev"
	}

	return ver
}
