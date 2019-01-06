package configuration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfiguration(t *testing.T) {
	configuration, err := GetConfiguration()

	require.NoError(t, err)
	require.NotNil(t, configuration)
	require.NotNil(t, configuration.Server)
	require.NotNil(t, configuration.Psql)
}
