// Package tonicpow is the official golang implementation for the TonicPow API
//
// If you have any suggestions or comments, please feel free to open an issue on
// this GitHub repository!
//
// By TonicPow Inc (https://tonicpow.com)
package tonicpow

// Version will return the version of the library
func Version() string {
	return version
}

// UserAgent will return the default user agent string
func UserAgent() string {
	return defaultUserAgent
}
