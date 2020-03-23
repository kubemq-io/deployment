package operator

import (
	"github.com/ghodss/yaml"
	rbac "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var roleBinding = `
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
subjects:
- kind: ServiceAccount
  name: {{.Name}}
roleRef:
  kind: Role
  name: {{.Name}}
  apiGroup: rbac.authorization.k8s.io
`

type RoleBinding struct {
	Name      string
	Namespace string
	*rbac.RoleBinding
}

func NewRoleBinding() *RoleBinding {
	return &RoleBinding{}
}

func (rb *RoleBinding) SetDefault(namespace, name string) *RoleBinding {
	rb.Name = name
	rb.Namespace = namespace
	return rb
}
func (rb *RoleBinding) Spec() ([]byte, error) {
	t := NewTemplate(roleBinding, rb)
	return t.Get()
}

func (rb *RoleBinding) Object() runtime.Object {
	return rb.RoleBinding
}
func (rb *RoleBinding) String() string {
	if rb.RoleBinding != nil {
		data, err := yaml.Marshal(rb.RoleBinding)
		if err != nil {
			return ""
		}
		return string(data)
	}
	data, err := rb.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}

func (rb *RoleBinding) Get() (*rbac.RoleBinding, error) {
	if rb.RoleBinding != nil {
		return rb.RoleBinding, nil
	}
	roleBinding := &rbac.RoleBinding{}
	data, err := rb.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, roleBinding)
	if err != nil {
		return nil, err
	}
	rb.RoleBinding = roleBinding
	return roleBinding, nil
}
