package crd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKubemqDashboard_Get(t *testing.T) {
	o := NewKubemqDashboard().SetDefault()
	crd, err := o.Get()
	require.NoError(t, err)
	require.NotNil(t, crd)
}
