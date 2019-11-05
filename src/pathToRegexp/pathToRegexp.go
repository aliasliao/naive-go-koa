package pathToRegexp

import (
	"regexp"
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
	//delimiter string
	// Optional character, or list of characters, to treat as "end" characters.
	//endsWidth uint8
	// List of characters to consider delimiters when parsing. (default: undefined, any character)
	//whitelist []uint8
}

func replaceAll(str string, re *regexp.Regexp, newSubStr string) string {
	posPairs := re.FindAllStringIndex(str, -1)
	start := 0
	ret := ""
	for _, posPair := range posPairs {
		ret += str[start:posPair[0]] + newSubStr
		start = posPair[1]
	}
	ret += str[start:]
	return ret
}

func replaceAllWith(str string, re *regexp.Regexp, genNewSubStr func(string) string) string {
	posPairs := re.FindAllStringIndex(str, -1)
	start := 0
	ret := ""
	for _, posPair := range posPairs {
		ret += str[start:posPair[0]] + genNewSubStr(str[posPair[0]+1:posPair[1]])
		start = posPair[1]
	}
	ret += str[start:]
	return ret
}

func pathToRegexpStr(path string, options *Options) string {
	if options == nil {
		options = &Options{
			sensitive: false,
			strict:    false,
			end:       true,
			start:     true,
		}
	}
	reStr := replaceAllWith(
		regexp.QuoteMeta(path),
		regexp.MustCompile(`:\w+`),
		func(pathVar string) string {
			return `(?P<` + pathVar + `>[^\/]+?)`
		},
	)
	if options.start {
		reStr = `^` + reStr
	}
	if !options.strict {
		reStr += `\/?`
	}
	if options.end {
		reStr += `$`
	}
	if !options.sensitive {
		reStr = `(?i:` + reStr + `)`
	}
	return reStr
}

func PathToRegexp(path string, options *Options) *regexp.Regexp {
	return regexp.MustCompile(pathToRegexpStr(path, options))
}
