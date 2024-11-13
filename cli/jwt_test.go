package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJwtEncode(t *testing.T) {
	tests := []struct {
		name      string
		secret    string
		algorithm string
		payload   string
		timeFn    func() time.Time
		res       string
		err       error
	}{
		{
			name: "not ok - secret missing",
			err:  errJwtSecretMissing,
		},
		{
			name:   "not ok - algorithm missing",
			secret: "x",
			err:    errJwtAlgorithmMissing,
		},
		{
			name:      "not ok - payload missing",
			secret:    "x",
			algorithm: "hs256",
			err:       errJwtPayloadMissing,
		},
		{
			name:      "not ok - algorithm unknown",
			secret:    "x",
			algorithm: "y",
			payload:   "z",
			err:       errJwtAlgorithmUnknown("y"),
		},
		{
			name:      "ok",
			secret:    "x",
			algorithm: "hs256",
			payload:   "z",
			timeFn: func() time.Time {
				return time.Unix(0, 0)
			}
		},
	}

	for _, tt := range tests {
		now := time.Now()
		timeFn := func() time.Time {
			return now
		}

		t.Run(tt.name, func(t *testing.T) {
			res, err := jwtEncode(tt.secret, tt.algorithm, tt.payload, tt.timeFn)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.res, res)
		})
	}
}
