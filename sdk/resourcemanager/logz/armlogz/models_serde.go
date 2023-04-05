//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogz

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "info", &e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
			err = unpopulate(val, "Info", &e.Info)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
			err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FilteringTag.
func (f FilteringTag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "action", f.Action)
	populate(objectMap, "name", f.Name)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FilteringTag.
func (f *FilteringTag) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "action":
			err = unpopulate(val, "Action", &f.Action)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &f.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &f.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IdentityProperties.
func (i IdentityProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IdentityProperties.
func (i *IdentityProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &i.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &i.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LogRules.
func (l LogRules) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "filteringTags", l.FilteringTags)
	populate(objectMap, "sendAadLogs", l.SendAADLogs)
	populate(objectMap, "sendActivityLogs", l.SendActivityLogs)
	populate(objectMap, "sendSubscriptionLogs", l.SendSubscriptionLogs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogRules.
func (l *LogRules) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "filteringTags":
			err = unpopulate(val, "FilteringTags", &l.FilteringTags)
			delete(rawMsg, key)
		case "sendAadLogs":
			err = unpopulate(val, "SendAADLogs", &l.SendAADLogs)
			delete(rawMsg, key)
		case "sendActivityLogs":
			err = unpopulate(val, "SendActivityLogs", &l.SendActivityLogs)
			delete(rawMsg, key)
		case "sendSubscriptionLogs":
			err = unpopulate(val, "SendSubscriptionLogs", &l.SendSubscriptionLogs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorProperties.
func (m MonitorProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "liftrResourceCategory", m.LiftrResourceCategory)
	populate(objectMap, "liftrResourcePreference", m.LiftrResourcePreference)
	populate(objectMap, "logzOrganizationProperties", m.LogzOrganizationProperties)
	populate(objectMap, "marketplaceSubscriptionStatus", m.MarketplaceSubscriptionStatus)
	populate(objectMap, "monitoringStatus", m.MonitoringStatus)
	populate(objectMap, "planData", m.PlanData)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "userInfo", m.UserInfo)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorProperties.
func (m *MonitorProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "liftrResourceCategory":
			err = unpopulate(val, "LiftrResourceCategory", &m.LiftrResourceCategory)
			delete(rawMsg, key)
		case "liftrResourcePreference":
			err = unpopulate(val, "LiftrResourcePreference", &m.LiftrResourcePreference)
			delete(rawMsg, key)
		case "logzOrganizationProperties":
			err = unpopulate(val, "LogzOrganizationProperties", &m.LogzOrganizationProperties)
			delete(rawMsg, key)
		case "marketplaceSubscriptionStatus":
			err = unpopulate(val, "MarketplaceSubscriptionStatus", &m.MarketplaceSubscriptionStatus)
			delete(rawMsg, key)
		case "monitoringStatus":
			err = unpopulate(val, "MonitoringStatus", &m.MonitoringStatus)
			delete(rawMsg, key)
		case "planData":
			err = unpopulate(val, "PlanData", &m.PlanData)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &m.ProvisioningState)
			delete(rawMsg, key)
		case "userInfo":
			err = unpopulate(val, "UserInfo", &m.UserInfo)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorResource.
func (m MonitorResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorResource.
func (m *MonitorResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &m.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &m.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &m.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &m.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorResourceListResponse.
func (m MonitorResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorResourceListResponse.
func (m *MonitorResourceListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &m.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorResourceUpdateParameters.
func (m MonitorResourceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorResourceUpdateParameters.
func (m *MonitorResourceUpdateParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &m.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitorUpdateProperties.
func (m MonitorUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "monitoringStatus", m.MonitoringStatus)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitorUpdateProperties.
func (m *MonitorUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "monitoringStatus":
			err = unpopulate(val, "MonitoringStatus", &m.MonitoringStatus)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitoredResource.
func (m MonitoredResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "reasonForLogsStatus", m.ReasonForLogsStatus)
	populate(objectMap, "reasonForMetricsStatus", m.ReasonForMetricsStatus)
	populate(objectMap, "sendingLogs", m.SendingLogs)
	populate(objectMap, "sendingMetrics", m.SendingMetrics)
	populate(objectMap, "systemData", m.SystemData)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitoredResource.
func (m *MonitoredResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "reasonForLogsStatus":
			err = unpopulate(val, "ReasonForLogsStatus", &m.ReasonForLogsStatus)
			delete(rawMsg, key)
		case "reasonForMetricsStatus":
			err = unpopulate(val, "ReasonForMetricsStatus", &m.ReasonForMetricsStatus)
			delete(rawMsg, key)
		case "sendingLogs":
			err = unpopulate(val, "SendingLogs", &m.SendingLogs)
			delete(rawMsg, key)
		case "sendingMetrics":
			err = unpopulate(val, "SendingMetrics", &m.SendingMetrics)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &m.SystemData)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitoredResourceListResponse.
func (m MonitoredResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitoredResourceListResponse.
func (m *MonitoredResourceListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &m.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitoringTagRules.
func (m MonitoringTagRules) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitoringTagRules.
func (m *MonitoringTagRules) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &m.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitoringTagRulesListResponse.
func (m MonitoringTagRulesListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitoringTagRulesListResponse.
func (m *MonitoringTagRulesListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &m.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MonitoringTagRulesProperties.
func (m MonitoringTagRulesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "logRules", m.LogRules)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "systemData", m.SystemData)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MonitoringTagRulesProperties.
func (m *MonitoringTagRulesProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "logRules":
			err = unpopulate(val, "LogRules", &m.LogRules)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &m.ProvisioningState)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &m.SystemData)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationResult.
func (o OperationResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResult.
func (o *OperationResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrganizationProperties.
func (o OrganizationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "companyName", o.CompanyName)
	populate(objectMap, "enterpriseAppId", o.EnterpriseAppID)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "singleSignOnUrl", o.SingleSignOnURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationProperties.
func (o *OrganizationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "companyName":
			err = unpopulate(val, "CompanyName", &o.CompanyName)
			delete(rawMsg, key)
		case "enterpriseAppId":
			err = unpopulate(val, "EnterpriseAppID", &o.EnterpriseAppID)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "singleSignOnUrl":
			err = unpopulate(val, "SingleSignOnURL", &o.SingleSignOnURL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanData.
func (p PlanData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "billingCycle", p.BillingCycle)
	populateTimeRFC3339(objectMap, "effectiveDate", p.EffectiveDate)
	populate(objectMap, "planDetails", p.PlanDetails)
	populate(objectMap, "usageType", p.UsageType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanData.
func (p *PlanData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "billingCycle":
			err = unpopulate(val, "BillingCycle", &p.BillingCycle)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &p.EffectiveDate)
			delete(rawMsg, key)
		case "planDetails":
			err = unpopulate(val, "PlanDetails", &p.PlanDetails)
			delete(rawMsg, key)
		case "usageType":
			err = unpopulate(val, "UsageType", &p.UsageType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SingleSignOnProperties.
func (s SingleSignOnProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enterpriseAppId", s.EnterpriseAppID)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "singleSignOnState", s.SingleSignOnState)
	populate(objectMap, "singleSignOnUrl", s.SingleSignOnURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleSignOnProperties.
func (s *SingleSignOnProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enterpriseAppId":
			err = unpopulate(val, "EnterpriseAppID", &s.EnterpriseAppID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "singleSignOnState":
			err = unpopulate(val, "SingleSignOnState", &s.SingleSignOnState)
			delete(rawMsg, key)
		case "singleSignOnUrl":
			err = unpopulate(val, "SingleSignOnURL", &s.SingleSignOnURL)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SingleSignOnResource.
func (s SingleSignOnResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleSignOnResource.
func (s *SingleSignOnResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &s.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SingleSignOnResourceListResponse.
func (s SingleSignOnResourceListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleSignOnResourceListResponse.
func (s *SingleSignOnResourceListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserInfo.
func (u UserInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "emailAddress", u.EmailAddress)
	populate(objectMap, "firstName", u.FirstName)
	populate(objectMap, "lastName", u.LastName)
	populate(objectMap, "phoneNumber", u.PhoneNumber)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserInfo.
func (u *UserInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "emailAddress":
			err = unpopulate(val, "EmailAddress", &u.EmailAddress)
			delete(rawMsg, key)
		case "firstName":
			err = unpopulate(val, "FirstName", &u.FirstName)
			delete(rawMsg, key)
		case "lastName":
			err = unpopulate(val, "LastName", &u.LastName)
			delete(rawMsg, key)
		case "phoneNumber":
			err = unpopulate(val, "PhoneNumber", &u.PhoneNumber)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserRoleListResponse.
func (u UserRoleListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserRoleListResponse.
func (u *UserRoleListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &u.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &u.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserRoleRequest.
func (u UserRoleRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "emailAddress", u.EmailAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserRoleRequest.
func (u *UserRoleRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "emailAddress":
			err = unpopulate(val, "EmailAddress", &u.EmailAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserRoleResponse.
func (u UserRoleResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "role", u.Role)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserRoleResponse.
func (u *UserRoleResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "role":
			err = unpopulate(val, "Role", &u.Role)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VMExtensionPayload.
func (v VMExtensionPayload) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "apiKey", v.APIKey)
	populate(objectMap, "region", v.Region)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMExtensionPayload.
func (v *VMExtensionPayload) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiKey":
			err = unpopulate(val, "APIKey", &v.APIKey)
			delete(rawMsg, key)
		case "region":
			err = unpopulate(val, "Region", &v.Region)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VMHostUpdateRequest.
func (v VMHostUpdateRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "state", v.State)
	populate(objectMap, "vmResourceIds", v.VMResourceIDs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMHostUpdateRequest.
func (v *VMHostUpdateRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "state":
			err = unpopulate(val, "State", &v.State)
			delete(rawMsg, key)
		case "vmResourceIds":
			err = unpopulate(val, "VMResourceIDs", &v.VMResourceIDs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VMResources.
func (v VMResources) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "agentVersion", v.AgentVersion)
	populate(objectMap, "id", v.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMResources.
func (v *VMResources) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "agentVersion":
			err = unpopulate(val, "AgentVersion", &v.AgentVersion)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &v.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VMResourcesListResponse.
func (v VMResourcesListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMResourcesListResponse.
func (v *VMResourcesListResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &v.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &v.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}