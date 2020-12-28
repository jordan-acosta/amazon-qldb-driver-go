package sqldriver

import (
	"errors"
	"regexp"
)

var connRegex = regexp.MustCompile(`^([a-z-\d]+):([a-z-]+)$`)

func parseConnString(connString string) (region string, ledger string, err error) {
	matches := connRegex.FindStringSubmatch(connString)
	if len(matches) < 3 {
		err = errors.New("malformed connection string")
		return
	}

	region = matches[1]
	ledger = matches[2]
	return
}
