// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

type CustomHealthStatus string

// Enum values for CustomHealthStatus
const (
	CustomHealthStatusHealthy   CustomHealthStatus = "HEALTHY"
	CustomHealthStatusUnhealthy CustomHealthStatus = "UNHEALTHY"
)

func (enum CustomHealthStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CustomHealthStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FilterCondition string

// Enum values for FilterCondition
const (
	FilterConditionEq      FilterCondition = "EQ"
	FilterConditionIn      FilterCondition = "IN"
	FilterConditionBetween FilterCondition = "BETWEEN"
)

func (enum FilterCondition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FilterCondition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthCheckType string

// Enum values for HealthCheckType
const (
	HealthCheckTypeHttp  HealthCheckType = "HTTP"
	HealthCheckTypeHttps HealthCheckType = "HTTPS"
	HealthCheckTypeTcp   HealthCheckType = "TCP"
)

func (enum HealthCheckType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthCheckType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthStatus string

// Enum values for HealthStatus
const (
	HealthStatusHealthy   HealthStatus = "HEALTHY"
	HealthStatusUnhealthy HealthStatus = "UNHEALTHY"
	HealthStatusUnknown   HealthStatus = "UNKNOWN"
)

func (enum HealthStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthStatusFilter string

// Enum values for HealthStatusFilter
const (
	HealthStatusFilterHealthy   HealthStatusFilter = "HEALTHY"
	HealthStatusFilterUnhealthy HealthStatusFilter = "UNHEALTHY"
	HealthStatusFilterAll       HealthStatusFilter = "ALL"
)

func (enum HealthStatusFilter) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthStatusFilter) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NamespaceFilterName string

// Enum values for NamespaceFilterName
const (
	NamespaceFilterNameType NamespaceFilterName = "TYPE"
)

func (enum NamespaceFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NamespaceFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NamespaceType string

// Enum values for NamespaceType
const (
	NamespaceTypeDnsPublic  NamespaceType = "DNS_PUBLIC"
	NamespaceTypeDnsPrivate NamespaceType = "DNS_PRIVATE"
	NamespaceTypeHttp       NamespaceType = "HTTP"
)

func (enum NamespaceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NamespaceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperationFilterName string

// Enum values for OperationFilterName
const (
	OperationFilterNameNamespaceId OperationFilterName = "NAMESPACE_ID"
	OperationFilterNameServiceId   OperationFilterName = "SERVICE_ID"
	OperationFilterNameStatus      OperationFilterName = "STATUS"
	OperationFilterNameType        OperationFilterName = "TYPE"
	OperationFilterNameUpdateDate  OperationFilterName = "UPDATE_DATE"
)

func (enum OperationFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperationFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusSubmitted OperationStatus = "SUBMITTED"
	OperationStatusPending   OperationStatus = "PENDING"
	OperationStatusSuccess   OperationStatus = "SUCCESS"
	OperationStatusFail      OperationStatus = "FAIL"
)

func (enum OperationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperationTargetType string

// Enum values for OperationTargetType
const (
	OperationTargetTypeNamespace OperationTargetType = "NAMESPACE"
	OperationTargetTypeService   OperationTargetType = "SERVICE"
	OperationTargetTypeInstance  OperationTargetType = "INSTANCE"
)

func (enum OperationTargetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperationTargetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeCreateNamespace    OperationType = "CREATE_NAMESPACE"
	OperationTypeDeleteNamespace    OperationType = "DELETE_NAMESPACE"
	OperationTypeUpdateService      OperationType = "UPDATE_SERVICE"
	OperationTypeRegisterInstance   OperationType = "REGISTER_INSTANCE"
	OperationTypeDeregisterInstance OperationType = "DEREGISTER_INSTANCE"
)

func (enum OperationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OperationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecordType string

// Enum values for RecordType
const (
	RecordTypeSrv   RecordType = "SRV"
	RecordTypeA     RecordType = "A"
	RecordTypeAaaa  RecordType = "AAAA"
	RecordTypeCname RecordType = "CNAME"
)

func (enum RecordType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecordType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RoutingPolicy string

// Enum values for RoutingPolicy
const (
	RoutingPolicyMultivalue RoutingPolicy = "MULTIVALUE"
	RoutingPolicyWeighted   RoutingPolicy = "WEIGHTED"
)

func (enum RoutingPolicy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RoutingPolicy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceFilterName string

// Enum values for ServiceFilterName
const (
	ServiceFilterNameNamespaceId ServiceFilterName = "NAMESPACE_ID"
)

func (enum ServiceFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
