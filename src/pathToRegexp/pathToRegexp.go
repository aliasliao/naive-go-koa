package pathToRegexp

import "regexp"

type Options struct {
	// When true the regexp will be case sensitive. (default: false)
	sensitive bool
	// When false the regexp allows an optional trailing delimiter to match. (default: false)
	strict bool
	// When true the regexp will match to the end of the string. (default: true)
	end bool
	// When true the regexp will match from the beginning of the string. (default: true)
	start bool
	// The default delimiter for segments. (default: '/')
	delimiter string
	// Optional character, or list of characters, to treat as "end" characters.
	endsWidth uint8
	// List of characters to consider delimiters when parsing. (default: undefined, any character)
	whitelist []uint8
}

func pathToRegexp(path string, options Options) *regexp.Regexp {

}
