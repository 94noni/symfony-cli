package humanlog

import (
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Wed Aug 12 16:39:56 2020 (310): [Debug] ...
var phpLogLineRegexp = regexp.MustCompile(`^(.+?) \((?:\d+)\)\: \[(.+?)\] (.+)\s*$`)

func convertPHPLog(in []byte) (*line, error) {
	allMatches := phpLogLineRegexp.FindAllSubmatch(in, -1)
	if allMatches == nil {
		return nil, nil
	}
	line := &line{
		source:  "PHP",
		level:   strings.ToLower(string(allMatches[0][2])),
		message: string(allMatches[0][3]),
		fields:  make(map[string]string),
	}
	// convert date (Wed Aug 12 16:39:56 2020)
	var err error
	m := string(allMatches[0][1])
	line.time, err = time.Parse(`Mon Jan 2 15:04:05 2006`, m)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return line, nil
}
