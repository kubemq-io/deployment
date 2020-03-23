package deployment

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDeployment(t *testing.T) {
	dep, err := CreateDeployment("kubemq", "kubemq-cluster")
	require.NoError(t, err)
	require.NotNil(t, dep)
	require.NoError(t, dep.IsValid())
	require.NotEmpty(t, dep.String())
	fmt.Println(dep.String())
}
