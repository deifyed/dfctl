package storage

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	testCases := []struct {
		name          string
		withStorePath string
		expectErr     error
	}{
		{
			name:          "Should return store without problems even when file does not exist",
			withStorePath: "/store",
		},
		{
			name:          "Should return error when StorePath is empty",
			withStorePath: "",
			expectErr:     ErrStorePathEmpty,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fs := &afero.Afero{Fs: afero.NewMemMapFs()}

			db := Store{Fs: fs, StorePath: tc.withStorePath}

			err := db.open()

			if tc.expectErr != nil {
				assert.EqualError(t, err, tc.expectErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
