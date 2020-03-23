package operator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceAccount_Get(t *testing.T) {
	sa := NewServiceAccount().SetDefault("kubemq", "kubemq-cluster")
	serviceAccount, err := sa.Get()
	require.NoError(t, err)
	require.NotNil(t, serviceAccount)
}
