package core

import (
	"github.com/ghodss/yaml"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var namespace = `
apiVersion: v1
kind: Namespace
metadata:
  name: {{.Name}}
`

type Namespace struct {
	Name string
	*apiv1.Namespace
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (ns *Namespace) SetDefault(namespace string) *Namespace {
	ns.Name = namespace
	return ns
}
func (ns *Namespace) Spec() ([]byte, error) {
	t := NewTemplate(namespace, ns)
	return t.Get()
}

func (ns *Namespace) Object() runtime.Object {
	return ns.Namespace
}
func (ns *Namespace) String() string {
	if ns.Namespace != nil {
		data, err := yaml.Marshal(ns.Namespace)
		if err != nil {
			return ""
		}
		return string(data)
	}
	data, err := ns.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (ns *Namespace) Get() (*apiv1.Namespace, error) {
	if ns.Namespace != nil {
		return ns.Namespace, nil
	}
	n := &apiv1.Namespace{}
	data, err := ns.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, n)
	if err != nil {
		return nil, err
	}
	ns.Namespace = n
	return n, nil
}
