package operator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOperator_Get(t *testing.T) {
	o := NewOperator().SetDefault("kubemq", "kubemq-cluster")
	dep, err := o.Get()
	require.NoError(t, err)
	require.NotNil(t, dep)
}
