package operator

import (
	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var operator = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
spec:
  replicas: 1
  selector:
    matchLabels:
      name: {{.Name}}
  template:
    metadata:
      labels:
        name: {{.Name}}
    spec:
      serviceAccountName: {{.ServiceAccountName}}
      containers:
        - name: {{.Name}}
          image: {{.Image}}
          command:
          - kubemq-operator
          imagePullPolicy: Always
          env:
            - name: WATCH_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: "{{.Name}}"
            - name: RELATED_IMAGE_KUBEMQ_CLUSTER
              value: {{.KubemqImage}}
            - name: RELATED_IMAGE_PROMETHEUS
              value: {{.PrometheusImage}}
            - name: RELATED_IMAGE_GRAFANA
              value: {{.GrafanaImage}}
            - name: KUBEMQ_VIEW_DASHBOARD_SOURCE
              value: {{.KubemqDashboardDashboardSource}}
            - name: KUBEMQ_LICENSE_MODE
              value: {{.LicenseMode}}
`

type Operator struct {
	Name                           string
	Namespace                      string
	Image                          string
	KubemqImage                    string
	PrometheusImage                string
	GrafanaImage                   string
	KubemqDashboardDashboardSource string
	LicenseMode                    string
	ServiceAccountName             string
	*appsv1.Deployment
}

func NewOperator() *Operator {
	return &Operator{}
}

func (o *Operator) SetDefault(namespace, name, serviceAccountName string) *Operator {

	o.Name = name
	o.Namespace = namespace
	o.Image = "docker.io/kubemq/kubemq-operator:latest"
	o.KubemqImage = "docker.io/kubemq/kubemq:latest"
	o.PrometheusImage = "prom/prometheus"
	o.GrafanaImage = "grafana/grafana:latest"
	o.KubemqDashboardDashboardSource = "https://raw.githubusercontent.com/kubemq-io/kubemq-dashboard/master/dashboard.json"
	o.LicenseMode = "COMMUNITY"
	o.Deployment = nil
	o.ServiceAccountName = serviceAccountName
	return o
}

func (o *Operator) Spec() ([]byte, error) {
	t := NewTemplate(operator, o)
	return t.Get()
}

func (o *Operator) Object() runtime.Object {
	return o.Deployment
}
func (o *Operator) String() string {

	data, err := o.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (o *Operator) Get() (*appsv1.Deployment, error) {
	if o.Deployment != nil {
		return o.Deployment, nil
	}
	deployment := &appsv1.Deployment{}
	data, err := o.Spec()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, deployment)
	if err != nil {
		return nil, err
	}
	o.Deployment = deployment
	return deployment, nil
}
