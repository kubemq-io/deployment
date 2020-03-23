package deployment

import (
	"fmt"
	"github.com/kubemq-io/deployment/core"
	"github.com/kubemq-io/deployment/crd"
	"github.com/kubemq-io/deployment/operator"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"strings"
)

type Deployment struct {
	CRDs []*v1beta1.CustomResourceDefinition
	*apiv1.Namespace
	*appsv1.Deployment
	*rbac.Role
	*rbac.RoleBinding
	*apiv1.ServiceAccount
	yamls []string
}

func CreateDeployment(namespace string) (*Deployment, error) {
	dep := &Deployment{
		CRDs:           nil,
		Namespace:      nil,
		Deployment:     nil,
		Role:           nil,
		RoleBinding:    nil,
		ServiceAccount: nil,
		yamls:          []string{},
	}
	var err error
	ns := core.NewNamespace().SetDefault(namespace)
	dep.Namespace, err = ns.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, namespace error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, ns.String())

	role := operator.NewRole().SetDefault(namespace, "kubemq-role")
	dep.Role, err = role.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, role error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, role.String())

	roleBinding := operator.NewRoleBinding().SetDefault(namespace, "kubemq-role-binding")
	dep.RoleBinding, err = roleBinding.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, role binding error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, roleBinding.String())

	serviceAccount := operator.NewServiceAccount().SetDefault(namespace, "kubemq-service-account")
	dep.ServiceAccount, err = serviceAccount.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, service account error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, serviceAccount.String())

	kubemqCluster := crd.NewKubemqCluster().SetDefault()
	kubemqClusterCrd, err := kubemqCluster.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, kubemq cluster crd error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, kubemqCluster.String())
	dep.CRDs = append(dep.CRDs, kubemqClusterCrd)

	kubemqDashboard := crd.NewKubemqDashboard().SetDefault()
	kubemqDashboardCrd, err := kubemqDashboard.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, kubemq dashboard crd error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, kubemqDashboard.String())
	dep.CRDs = append(dep.CRDs, kubemqDashboardCrd)

	operator := operator.NewOperator().SetDefault(namespace, "kubemq-operator")
	dep.Deployment, err = operator.Get()
	if err != nil {
		return nil, fmt.Errorf("error create deployment, operator error: %s", err.Error())
	}
	dep.yamls = append(dep.yamls, operator.String())

	return dep, nil
}

func (b *Deployment) IsValid() error {
	if b.CRDs == nil {
		return fmt.Errorf("no crd exsits or defined")
	}
	if b.Namespace == nil {
		return fmt.Errorf("no name deployment exsits or defined")
	}
	if b.Deployment == nil {
		return fmt.Errorf("no operator deployment exsits or defined")
	}
	if b.Role == nil {
		return fmt.Errorf("no role exsits or defined")
	}

	if b.RoleBinding == nil {
		return fmt.Errorf("no role binding exsits or defined")
	}
	if b.ServiceAccount == nil {
		return fmt.Errorf("no service account exsits or defined")
	}
	return nil
}

func (b *Deployment) String() string {
	return strings.Join(b.yamls, "---\n")
}
