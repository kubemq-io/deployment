package operator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoleBinding_Get(t *testing.T) {
	rb := NewRoleBinding().SetDefault("kubemq", "kubemq-cluster")
	roleBinding, err := rb.Get()
	require.NoError(t, err)
	require.NotNil(t, roleBinding)
}
