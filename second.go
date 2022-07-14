package second

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Helpers.
const (
	s = 1
	m = s * 60
	h = m * 60
	d = h * 24
	w = d * 7
	y = d * 365.25

	regex = `(-?(?:\d+)?\.?\d+) *(?i)(seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|weeks?|w|years?|yrs?|y)?`
)

func Parse(str string) (float64, error) {
	if len(str) > 100 {
		return 0, errors.New("Value exceeds the maximum length of 100 characters.")
	}

	// find the first substring that satisfies the regex.
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(str)

	if len(match) < 3 || match[2] == "" {
		return 0, fmt.Errorf("The string %s was not matched.", str)
	}

	n, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, err
	}

	t := strings.ToLower(match[2])

	switch t {
	case "years":
	case "year":
	case "yrs":
	case "yr":
	case "y":
		return n * y, nil

	case "weeks":
	case "week":
	case "w":
		return n * w, nil

	case "days":
	case "day":
	case "d":
		return n * d, nil

	case "hours":
	case "hour":
	case "hrs":
	case "hr":
	case "h":
		return n * h, nil

	case "minutes":
	case "minute":
	case "mins":
	case "min":
	case "m":
		return n * m, nil

	case "seconds":
	case "second":
	case "secs":
	case "sec":
	case "s":
		return n * s, nil
	}

	return 0, fmt.Errorf("The unit %s was matched, but no matching case exists.", t)
}
