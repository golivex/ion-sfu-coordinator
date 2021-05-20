// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

type DetailType string

// Enum values for DetailType
const (
	DetailTypeBasic DetailType = "BASIC"
	DetailTypeFull  DetailType = "FULL"
)

func (enum DetailType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DetailType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ListEventTypesFilterName string

// Enum values for ListEventTypesFilterName
const (
	ListEventTypesFilterNameResourceType ListEventTypesFilterName = "RESOURCE_TYPE"
	ListEventTypesFilterNameServiceName  ListEventTypesFilterName = "SERVICE_NAME"
)

func (enum ListEventTypesFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ListEventTypesFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ListNotificationRulesFilterName string

// Enum values for ListNotificationRulesFilterName
const (
	ListNotificationRulesFilterNameEventTypeId   ListNotificationRulesFilterName = "EVENT_TYPE_ID"
	ListNotificationRulesFilterNameCreatedBy     ListNotificationRulesFilterName = "CREATED_BY"
	ListNotificationRulesFilterNameResource      ListNotificationRulesFilterName = "RESOURCE"
	ListNotificationRulesFilterNameTargetAddress ListNotificationRulesFilterName = "TARGET_ADDRESS"
)

func (enum ListNotificationRulesFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ListNotificationRulesFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ListTargetsFilterName string

// Enum values for ListTargetsFilterName
const (
	ListTargetsFilterNameTargetType    ListTargetsFilterName = "TARGET_TYPE"
	ListTargetsFilterNameTargetAddress ListTargetsFilterName = "TARGET_ADDRESS"
	ListTargetsFilterNameTargetStatus  ListTargetsFilterName = "TARGET_STATUS"
)

func (enum ListTargetsFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ListTargetsFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NotificationRuleStatus string

// Enum values for NotificationRuleStatus
const (
	NotificationRuleStatusEnabled  NotificationRuleStatus = "ENABLED"
	NotificationRuleStatusDisabled NotificationRuleStatus = "DISABLED"
)

func (enum NotificationRuleStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NotificationRuleStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TargetStatus string

// Enum values for TargetStatus
const (
	TargetStatusPending     TargetStatus = "PENDING"
	TargetStatusActive      TargetStatus = "ACTIVE"
	TargetStatusUnreachable TargetStatus = "UNREACHABLE"
	TargetStatusInactive    TargetStatus = "INACTIVE"
	TargetStatusDeactivated TargetStatus = "DEACTIVATED"
)

func (enum TargetStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TargetStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
