package kstring

import (
	"strings"

	"github.com/alioygur/godash"
)

// ToLowerFirstCamelCase returns the given string in camelcase formatted string
// but with the first letter being lowercase.
func ToLowerFirstCamelCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(string(s[0]))
	}
	return strings.ToLower(string(s[0])) + godash.ToCamelCase(s)[1:]
}

// ToUpperFirst returns the given string with the first letter being uppercase.
func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(string(s[0]))
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToLowerSnakeCase the given string in snake-case format.
func ToLowerSnakeCase(s string) string {
	return strings.ToLower(godash.ToSnakeCase(s))
}

// ToCamelCase the given string in camelcase format.
func ToCamelCase(s string) string {
	return godash.ToCamelCase(s)
}

func ToCacheKey(prefix, s1 string, s ...string) string {
	return prefix + strings.Join(append([]string{s1}, s...), "::")
}
