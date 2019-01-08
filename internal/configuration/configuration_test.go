package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfiguration(t *testing.T) {
	os.Setenv("SERVER_GIN_MODE", "debug")
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("PSQL_DRIVER_NAME", "postgres")

	configuration := GetConfiguration()

	require.NotNil(t, configuration)
	require.NotNil(t, configuration.Server)
	require.Equal(t, configuration.Server.GinMode, "debug")
	require.Equal(t, configuration.Server.Host, "localhost")
	require.NotNil(t, configuration.Psql)
	require.NotNil(t, configuration.Psql.DriverName, "postgres")
}
