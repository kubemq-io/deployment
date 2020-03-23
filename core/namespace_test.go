package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOperator_Get(t *testing.T) {
	o := NewNamespace().SetDefault("kubemq")
	ns, err := o.Get()
	require.NoError(t, err)
	require.NotNil(t, ns)
}
