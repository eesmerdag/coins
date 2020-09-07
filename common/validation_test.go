package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLimitValidation(t *testing.T) {
	for _, tc := range []struct {
		name       string
		isDefined  bool
		RawQuery   string
		errMessage string
		limit string
	}{
		{
			name:       "undefined",
			errMessage: "limit is not defined",
			RawQuery:   "",
		},
		{
			name:       "not a number",
			isDefined:  true,
			RawQuery:   "limit=notanumber",
			errMessage: "limit must be integer",
		},
		{
			name:       "lower than limit",
			isDefined:  true,
			RawQuery:   "limit=9",
			errMessage: "limit must be between 10 and 100",
		},
		{
			name:       "above than limit",
			isDefined:  true,
			RawQuery:   "limit=101",
			errMessage: "limit must be between 10 and 100",
		},
		{
			name:      "successful case",
			isDefined: true,
			RawQuery:  "limit=50",
			limit: "50",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l, err := IsLimitValid(tc.RawQuery)
			if !tc.isDefined {
				assert.Equal(t, err.Error(), tc.errMessage)
			} else if tc.isDefined && err != nil {
				assert.Equal(t, err.Error(), tc.errMessage)
			} else{
				assert.Equal(t, l, tc.limit)
			}
		})
	}
}
