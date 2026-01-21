package cli

import (
	"strings"
	"unicode"
)

func normalizeModule(name string) string {
	return toSnake(name)
}

func toSnake(input string) string {
	if input == "" {
		return ""
	}

	var b strings.Builder
	var prevLower bool
	var lastUnderscore bool
	for _, r := range input {
		switch {
		case r == '-' || r == ' ' || r == '.':
			if b.Len() > 0 && !lastUnderscore {
				b.WriteRune('_')
				lastUnderscore = true
			}
			prevLower = false
		case unicode.IsUpper(r):
			if prevLower {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
			prevLower = false
			lastUnderscore = false
		default:
			b.WriteRune(unicode.ToLower(r))
			prevLower = unicode.IsLower(r) || unicode.IsDigit(r)
			lastUnderscore = r == '_'
		}
	}

	return strings.Trim(b.String(), "_")
}

func toPascal(input string) string {
	if input == "" {
		return ""
	}

	normalized := strings.NewReplacer("-", "_", " ", "_", ".", "_").Replace(input)
	parts := strings.Split(normalized, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		hasUpper := false
		for i := 1; i < len(runes); i++ {
			if unicode.IsUpper(runes[i]) {
				hasUpper = true
				break
			}
		}
		if !hasUpper {
			for i := 1; i < len(runes); i++ {
				runes[i] = unicode.ToLower(runes[i])
			}
		}
		b.WriteString(string(runes))
	}
	return b.String()
}
