package crd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKubemqCluster_Get(t *testing.T) {
	o := NewKubemqCluster().SetDefault()
	crd, err := o.Get()
	require.NoError(t, err)
	require.NotNil(t, crd)
}
