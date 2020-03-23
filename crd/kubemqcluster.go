package crd

import (
	"github.com/ghodss/yaml"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

var crdKubemqClusters = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: kubemqclusters.core.k8s.kubemq.io
spec:
  additionalPrinterColumns:
    - JSONPath: .status.version
      name: Version
      type: string
    - JSONPath: .status.status
      name: Status
      type: string
    - JSONPath: .status.replicas
      name: Replicas
      type: string
    - JSONPath: .status.ready
      name: Ready
      type: string
    - JSONPath: .status.grpc
      name: gRPC
      type: string
    - JSONPath: .status.rest
      name: Rest
      type: string
    - JSONPath: .status.api
      name: API
      type: string
    - JSONPath: .status.license_type
      name: License-type
      type: string
    - JSONPath: .status.license_to
      name: License-To
      type: string
    - JSONPath: .status.license_expire
      name: License-Expire
      type: string
  group: core.k8s.kubemq.io
  names:
    kind: KubemqCluster
    listKind: KubemqClusterList
    plural: kubemqclusters
    singular: kubemqcluster
  scope: Namespaced
  subresources:
    scale:
      labelSelectorPath: .status.selector
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
    status: {}
  validation:
    openAPIV3Schema:
      description: KubemqCluster is the Schema for the kubemqclusters API
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
          properties:
            api:
              properties:
                disabled:
                  type: boolean
                expose:
                  pattern: (ClusterIP|NodePort|LoadBalancer)
                  type: string
                nodePort:
                  format: int32
                  type: integer
                port:
                  format: int32
                  type: integer
              type: object
            authentication:
              properties:
                key:
                  type: string
                type:
                  type: string
              type: object
            authorization:
              properties:
                autoReload:
                  format: int32
                  type: integer
                policy:
                  type: string
                url:
                  type: string
              type: object
            configData:
              type: string
            gateways:
              properties:
                ca:
                  type: string
                cert:
                  type: string
                key:
                  type: string
                port:
                  format: int32
                  type: integer
                remotes:
                  items:
                    type: string
                  type: array
              type: object
            grpc:
              properties:
                bodyLimit:
                  format: int32
                  type: integer
                bufferSize:
                  format: int32
                  type: integer
                disabled:
                  type: boolean
                expose:
                  pattern: (ClusterIP|NodePort|LoadBalancer)
                  type: string
                nodePort:
                  format: int32
                  type: integer
                port:
                  format: int32
                  type: integer
              type: object
            health:
              properties:
                enabled:
                  type: boolean
                failureThreshold:
                  format: int32
                  type: integer
                initialDelaySeconds:
                  format: int32
                  type: integer
                periodSeconds:
                  format: int32
                  type: integer
                successThreshold:
                  format: int32
                  type: integer
                timeoutSeconds:
                  format: int32
                  type: integer
              type: object
            image:
              properties:
                image:
                  type: string
                pullPolicy:
                  pattern: (IfNotPresent|Always|Never)
                  type: string
              type: object
            license:
              type: string
            log:
              properties:
                file:
                  type: string
                level:
                  format: int32
                  type: integer
              type: object
            nodeSelectors:
              properties:
                keys:
                  additionalProperties:
                    type: string
                  type: object
              type: object
            notification:
              properties:
                enabled:
                  type: boolean
                log:
                  type: boolean
                prefix:
                  type: string
              type: object
            queue:
              properties:
                defaultVisibilitySeconds:
                  format: int32
                  minimum: 0
                  type: integer
                defaultWaitTimeoutSeconds:
                  format: int32
                  minimum: 0
                  type: integer
                maxDelaySeconds:
                  format: int32
                  minimum: 0
                  type: integer
                maxExpirationSeconds:
                  format: int32
                  minimum: 0
                  type: integer
                maxReQueues:
                  format: int32
                  minimum: 0
                  type: integer
                maxReceiveMessagesRequest:
                  format: int32
                  minimum: 0
                  type: integer
                maxVisibilitySeconds:
                  format: int32
                  minimum: 0
                  type: integer
                maxWaitTimeoutSeconds:
                  format: int32
                  minimum: 0
                  type: integer
              type: object
            replicas:
              format: int32
              minimum: 0
              type: integer
            resources:
              properties:
                limitsCpu:
                  type: string
                limitsMemory:
                  type: string
                requestsCpu:
                  type: string
                requestsMemory:
                  type: string
              type: object
            rest:
              properties:
                bodyLimit:
                  format: int32
                  type: integer
                bufferSize:
                  format: int32
                  type: integer
                disabled:
                  type: boolean
                expose:
                  pattern: (ClusterIP|NodePort|LoadBalancer)
                  type: string
                nodePort:
                  format: int32
                  type: integer
                port:
                  format: int32
                  type: integer
              type: object
            routing:
              properties:
                autoReload:
                  format: int32
                  type: integer
                data:
                  type: string
                url:
                  type: string
              type: object
            store:
              properties:
                clean:
                  type: boolean
                maxChannelSize:
                  format: int32
                  minimum: 0
                  type: integer
                maxChannels:
                  format: int32
                  minimum: 0
                  type: integer
                maxMessages:
                  format: int32
                  minimum: 0
                  type: integer
                maxSubscribers:
                  format: int32
                  minimum: 0
                  type: integer
                messagesRetentionMinutes:
                  format: int32
                  minimum: 0
                  type: integer
                path:
                  type: string
                purgeInactiveMinutes:
                  format: int32
                  minimum: 0
                  type: integer
              type: object
            tls:
              properties:
                ca:
                  type: string
                cert:
                  type: string
                key:
                  type: string
              type: object
            volume:
              properties:
                size:
                  type: string
              type: object
          type: object
        status:
          description: KubemqClusterStatus defines the observed state of KubemqCluster
          properties:
            api:
              type: string
            grpc:
              type: string
            license_expire:
              type: string
            license_to:
              type: string
            license_type:
              type: string
            ready:
              format: int32
              type: integer
            replicas:
              format: int32
              type: integer
            rest:
              type: string
            selector:
              type: string
            status:
              type: string
            version:
              type: string
          required:
            - api
            - grpc
            - license_expire
            - license_to
            - license_type
            - ready
            - replicas
            - rest
            - selector
            - status
            - version
          type: object
      type: object
  version: v1alpha1
  versions:
    - name: v1alpha1
      served: true
      storage: true
`

type KubemqCluster struct {
	*v1beta1.CustomResourceDefinition
}

func NewKubemqCluster() *KubemqCluster {
	return &KubemqCluster{}
}
func (c *KubemqCluster) SetDefault() *KubemqCluster {
	return c
}
func (c *KubemqCluster) Spec() ([]byte, error) {
	t := NewTemplate(crdKubemqClusters, c)
	return t.Get()
}

func (c *KubemqCluster) Object() runtime.Object {
	return c.CustomResourceDefinition
}
func (c *KubemqCluster) String() string {
	data, err := c.Spec()
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}
func (c *KubemqCluster) Get() (*v1beta1.CustomResourceDefinition, error) {
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