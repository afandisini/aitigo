package migrate

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type direction string

const (
	directionUp   direction = "up"
	directionDown direction = "down"
)

var migrationPattern = regexp.MustCompile(`^(\d+)_([a-z0-9_-]+)\.(up|down)\.sql$`)

type parsedFile struct {
	version   int
	name      string
	direction direction
}

func parseFilename(filename string) (parsedFile, bool) {
	base := filepath.Base(filename)
	matches := migrationPattern.FindStringSubmatch(base)
	if len(matches) != 4 {
		return parsedFile{}, false
	}

	version, err := strconv.Atoi(matches[1])
	if err != nil {
		return parsedFile{}, false
	}

	name := strings.ToLower(matches[2])
	dir := direction(matches[3])
	return parsedFile{version: version, name: name, direction: dir}, true
}
