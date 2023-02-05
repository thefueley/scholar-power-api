// build +integration
package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDatabase(t *testing.T) {
	db, err := NewDatabase("")
	require.NoError(t, err)
	require.NotEmpty(t, db)
}
