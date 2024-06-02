package text

import "strings"

func ToKebabCase(v string) string {
	return strings.ReplaceAll(strings.ToLower(v), " ", "-")
}
