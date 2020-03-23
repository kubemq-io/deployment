package operator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRole_Get(t *testing.T) {
	r := NewRole().SetDefault("kubemq", "kubemq-cluster")
	role, err := r.Get()
	require.NoError(t, err)
	require.NotNil(t, role)
}
