package pathToRegexp

import (
	"regexp"
	"strings"
)

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
	segments := strings.Split(path, options.delimiter)
	newSegments := make([]string, 0)
	for _, segment := range segments {
		re := regexp.MustCompile(":(\\w+)")
		cuts := re.FindAllStringSubmatchIndex(segment, -1)
		if len(cuts) == 0 {
			cuts = append(cuts, []int{0, 0})
		}
		newSeg := regexp.QuoteMeta(segment[0:cuts[0][0]])
		for i, mts := range cuts {
			newSeg += "([^\\/]+?)"
			if i+1 == len(cuts) {
				newSeg += regexp.QuoteMeta(segment[mts[1]:])
			} else {
				newSeg += regexp.QuoteMeta(segment[mts[1]:cuts[i+1][0]])
			}
		}
		newSegments = append(newSegments, newSeg)
	}
	regPath := strings.Join(newSegments, regexp.QuoteMeta(options.delimiter))
	if !options.strict {
		regPath += "(:?\\/)?$"
	} else {
		regPath += "$"
	}

	if ret, err := regexp.Compile(regPath); err != nil {
		panic(err)
	} else {
		return ret
	}
}
