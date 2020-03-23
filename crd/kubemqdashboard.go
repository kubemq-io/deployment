package crd

import (
	"github.com/ghodss/yaml"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

var crdKubemqDashboard = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: kubemqdashboards.core.k8s.kubemq.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.status
    name: Status
    type: string
  - JSONPath: .status.address
    name: Address
    type: string
  - JSONPath: .status.prometheus_version
    name: Prometheus-Version
    type: string
  - JSONPath: .status.grafana_version
    name: Grafana-Version
    type: string
  group: core.k8s.kubemq.io
  names:
    kind: KubemqDashboard
    listKind: KubemqDashboardList
    plural: kubemqdashboards
    singular: kubemqdashboard
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: KubemqDashboard is the Schema for the kubemqdashboards API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: KubemqDashboardSpec defines the desired state of KubemqDashboard
          properties:
            grafana:
              properties:
                dashboardUrl:
                  type: string
                image:
                  type: string
              type: object
            port:
              format: int32
              type: integer
            prometheus:
              properties:
                image:
                  type: string
                nodePort:
                  format: int32
                  type: integer
              type: object
          type: object
        status:
          description: KubemqDashboardStatus defines the observed state of KubemqDashboard
          properties:
            address:
              type: string
            grafana_version:
              type: string
            prometheus_version:
              type: string
            status:
              type: string
          required:
          - address
          - grafana_version
          - prometheus_version
          - status
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
`

type KubemqDashboard struct {
	*v1beta1.CustomResourceDefinition
}

func NewKubemqDashboard() *KubemqDashboard {
	return &KubemqDashboard{}
}
func (c *KubemqDashboard) SetDefault() *KubemqDashboard {
	return c
}

func (c *KubemqDashboard) Spec() ([]byte, error) {
	t := NewTemplate(crdKubemqDashboard, c)
	return t.Get()
}

func (c *KubemqDashboard) Object() runtime.Object {
	return c.CustomResourceDefinition
}
func (c *KubemqDashboard) String() string {
	if c.CustomResourceDefinition != nil {
		data, err := yaml.Marshal(c.CustomResourceDefinition)
		if err != nil {
			return ""
		}
		return string(data)
	}
	data, err := c.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (c *KubemqDashboard) Get() (*v1beta1.CustomResourceDefinition, error) {
	if c.CustomResourceDefinition != nil {
		return c.CustomResourceDefinition, nil
	}
	crd := &v1beta1.CustomResourceDefinition{}
	data, err := c.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, crd)
	if err != nil {
		return nil, err
	}
	c.CustomResourceDefinition = crd
	return crd, nil
}
