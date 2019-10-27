package pathToRegexp

import (
	"regexp"
	"testing"
)

func TestPathToRegexp(t *testing.T) {
	options := &Options{
		sensitive: false,
		strict:    false,
		end:       true,
		start:     true,
	}
	t.Run("replaceAll", func(t *testing.T) {
		newStr := replaceAll("/foo/:bar", regexp.MustCompile(`:\w+`), "placeholder")
		expectNewStr := "/foo/placeholder"
		if newStr != expectNewStr {
			t.Errorf("Expect: %s, got: %s", expectNewStr, newStr)
		}

		newStr = replaceAll("/foo/:bar/:baz", regexp.MustCompile(`:\w+`), "test")
		expectNewStr = "/foo/test/test"
		if newStr != expectNewStr {
			t.Errorf("Expect: %s, got: %s", expectNewStr, newStr)
		}
	})
	t.Run("pathToRegexpStr", func(t *testing.T) {
		reStr := pathToRegexpStr("/foo/:bar", options)
		expectReStr := `(?i:^/foo/([^\/]+?)\/?$)`
		if reStr != expectReStr {
			t.Errorf("Expect: %s, got: %s", expectReStr, reStr)
		}
	})
	t.Run("PathToRegexp", func(t *testing.T) {
		re := PathToRegexp("/foo/:bar", options)
		paths := map[string]bool{
			"/foo/123":     true,
			"/foo/123/":    true,
			"/FOO/123":     true,
			"":             false,
			"/foo":         false,
			"/foo/":        false,
			"/foo/123/bar": false,
			"123":          false,
		}
		for path, expect := range paths {
			if res := re.MatchString(path); res != expect {
				t.Errorf("Path: %v, Expect: %v", path, expect)
			}
		}
	})
}
