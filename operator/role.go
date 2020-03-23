package operator

import (
	"github.com/ghodss/yaml"
	rbac "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var role = `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - services/finalizers
  - endpoints
  - persistentvolumeclaims
  - events
  - configmaps
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  - daemonsets
  - replicasets
  - statefulsets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
- apiGroups:
  - apps
  resourceNames:
  - kubemq-operator
  resources:
  - deployments/finalizers
  verbs:
  - update
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
- apiGroups:
  - apps
  resources:
  - replicasets
  - deployments
  verbs:
  - get
- apiGroups:
  - core.k8s.kubemq.io
  resources:
  - '*'
  - kubemqclusters
  - kubemqdashboards
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs:
  - get
  - list
  - watch
  - create
  - delete
  - update
- apiGroups:
  - extensions
  resources:
  - ingresses/status
  verbs:
  - update
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses
  verbs:
  - get
  - list
  - watch
  - create
  - delete
  - update
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses/status
  verbs:
  - update
`

type Role struct {
	Name      string
	Namespace string
	*rbac.Role
}

func NewRole() *Role {
	return &Role{}
}
func (r *Role) SetDefault(namespace, name string) *Role {
	r.Name = name
	r.Namespace = namespace

	return r
}
func (r *Role) Spec() ([]byte, error) {
	t := NewTemplate(role, r)
	return t.Get()
}

func (r *Role) Object() runtime.Object {
	return r.Role
}
func (r *Role) String() string {

	data, err := r.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}

func (r *Role) Get() (*rbac.Role, error) {
	if r.Role != nil {
		return r.Role, nil
	}
	role := &rbac.Role{}
	data, err := r.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, role)
	if err != nil {
		return nil, err
	}
	r.Role = role
	return role, nil
}
