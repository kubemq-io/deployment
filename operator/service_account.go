package operator

import (
	"github.com/ghodss/yaml"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var serviceAccount = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
`

type ServiceAccount struct {
	Name      string
	Namespace string
	*apiv1.ServiceAccount
}

func NewServiceAccount() *ServiceAccount {
	return &ServiceAccount{}
}

func (sa *ServiceAccount) SetDefault(namespace, name string) *ServiceAccount {
	sa.Name = name
	sa.Namespace = namespace

	return sa
}
func (sa *ServiceAccount) Spec() ([]byte, error) {
	t := NewTemplate(serviceAccount, sa)
	return t.Get()
}

func (sa *ServiceAccount) Object() runtime.Object {
	return sa.ServiceAccount
}
func (sa *ServiceAccount) String() string {

	data, err := sa.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (sa *ServiceAccount) Get() (*apiv1.ServiceAccount, error) {
	if sa.ServiceAccount != nil {
		return sa.ServiceAccount, nil
	}
	serviceAccount := &apiv1.ServiceAccount{}
	data, err := sa.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, serviceAccount)
	if err != nil {
		return nil, err
	}
	sa.ServiceAccount = serviceAccount
	return serviceAccount, nil
}
