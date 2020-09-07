package common

import (
	"errors"
	"net/url"
	"strconv"
)

const min = 10
const max = 100

func IsLimitValid(urlRawQuery string) (string, error) {
	params, _ := url.ParseQuery(urlRawQuery)
	if _, ok := params["limit"]; !ok {
		return "", errors.New("limit is not defined")
	}

	limit := params["limit"][0]
	i, err := strconv.Atoi(limit);
	if err != nil {
		return "", errors.New("limit must be integer")
	}

	if !(i >= min && i <= max) {
		return "", errors.New("limit must be between 10 and 100")
	}

	return limit, nil
}
