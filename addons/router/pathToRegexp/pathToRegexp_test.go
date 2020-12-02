package pathToRegexp

import (
	"regexp"
	"testing"
)

func TestPathToRegexp(t *testing.T) {

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

	t.Run("replaceAllWith", func(t *testing.T) {
		converter := func(s string) string { return "[" + s + "]" }
		newStr := replaceAllWith("/foo/:bar", regexp.MustCompile(`:\w+`), converter)
		expectNewStr := "/foo/[bar]"
		if newStr != expectNewStr {
			t.Errorf("Expect: %s, got: %s", expectNewStr, newStr)
		}

		newStr = replaceAllWith("/foo/:bar/:baz", regexp.MustCompile(`:\w+`), converter)
		expectNewStr = "/foo/[bar]/[baz]"
		if newStr != expectNewStr {
			t.Errorf("Expect: %s, got: %s", expectNewStr, newStr)
		}
	})

	t.Run("pathToRegexpStr", func(t *testing.T) {
		reStr := pathToRegexpStr("/foo/:bar/:baz", nil)
		expectReStr := `(?i:^/foo/(?P<bar>[^/]+?)/(?P<baz>[^/]+?)/?$)`
		if reStr != expectReStr {
			t.Errorf("Expect: %s, got: %s", expectReStr, reStr)
		}
	})

	t.Run("PathToRegexp", func(t *testing.T) {
		re := PathToRegexp("/foo/:bar/:baz", nil)
		t.Run("should match specific paths", func(t *testing.T) {
			paths := map[string]bool{
				// match cases
				"/foo/123/666":  true,
				"/foo/123/666/": true,
				"/FOO/123/666":  true,
				// not match cases
				"":                  false,
				"/":                 false,
				"/foo":              false,
				"/foo/":             false,
				"/foo//":            false,
				"/foo/123":          false,
				"/foo/123/":         false,
				"/foo/123//":        false,
				"/foo/123/666/bar":  false,
				"/foo/123//666/bar": false,
			}
			for path, expect := range paths {
				if res := re.MatchString(path); res != expect {
					t.Errorf("Path: %v, Expect: %v", path, expect)
				}
			}
		})
		t.Run("should capture path variables", func(t *testing.T) {
			keys := re.SubexpNames()
			values := re.FindAllStringSubmatch("/foo/123/666", -1)[0]
			if len(keys) != len(values) {
				t.Errorf("keys.length != values.length")
			}
			pathVariables := map[string]string{
				"":    "/foo/123/666",
				"bar": "123",
				"baz": "666",
			}
			for i, key := range keys {
				if pathVariables[key] != values[i] {
					t.Errorf(
						"Expect pathVariables[%s] = %s, Actual value = %s",
						key, pathVariables[key], values[i],
					)
				}
			}
		})
	})
}
