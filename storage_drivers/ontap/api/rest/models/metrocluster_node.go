// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetroclusterNode Data for a node in a MetroCluster. REST: /api/cluster/metrocluster/nodes
//
// swagger:model metrocluster_node
type MetroclusterNode struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Specifies if automatic unplanned switchover is enabled.
	// Read Only: true
	AutomaticUso *bool `json:"automatic_uso,omitempty"`

	// cluster
	Cluster *MetroclusterNodeCluster `json:"cluster,omitempty"`

	// Configuration state of the node.
	// Read Only: true
	// Enum: [unreachable configured]
	ConfigurationState string `json:"configuration_state,omitempty"`

	// DR Group ID.
	// Read Only: true
	DrGroupID int64 `json:"dr_group_id,omitempty"`

	// State of the DR mirroring configuration.
	// Read Only: true
	// Enum: [enabled disabled unreachable configured]
	DrMirroringState string `json:"dr_mirroring_state,omitempty"`

	// State of the DR operation.
	// Read Only: true
	// Enum: [normal switchover_bypassed switchover_in_progress switchover_completed switchover_failed heal_aggrs_in_progress heal_aggrs_completed heal_aggrs_failed heal_roots_in_progress heal_roots_completed heal_roots_failed switchback_vetoed switchback_vetocheck_locked switchback_pre_commit_completed switchback_in_progress switchback_completed switchback_failed negotiated_switchover_vetoed negotiated_switchover_vetocheck_locked negotiated_switchover_pre_commit_completed negotiated_switchover_in_progress negotiated_switchover_completed negotiated_switchover_in_progress_waiting_for_DR_partner negotiated_switchover_incomplete negotiated_switchover_failed negotiated_switchover_failed_on_DR_partner switchback_recovery_in_progress switchback_recovery_complete waiting_for_switchback_recovery unknown]
	DrOperationState string `json:"dr_operation_state,omitempty"`

	// node
	Node *MetroclusterNodeNode `json:"node,omitempty"`
}

// Validate validates this metrocluster node
func (m *MetroclusterNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrMirroringState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrOperationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterNode) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

var metroclusterNodeTypeConfigurationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unreachable","configured"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterNodeTypeConfigurationStatePropEnum = append(metroclusterNodeTypeConfigurationStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// configuration_state
	// ConfigurationState
	// unreachable
	// END DEBUGGING
	// MetroclusterNodeConfigurationStateUnreachable captures enum value "unreachable"
	MetroclusterNodeConfigurationStateUnreachable string = "unreachable"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// configuration_state
	// ConfigurationState
	// configured
	// END DEBUGGING
	// MetroclusterNodeConfigurationStateConfigured captures enum value "configured"
	MetroclusterNodeConfigurationStateConfigured string = "configured"
)

// prop value enum
func (m *MetroclusterNode) validateConfigurationStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterNodeTypeConfigurationStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterNode) validateConfigurationState(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigurationStateEnum("configuration_state", "body", m.ConfigurationState); err != nil {
		return err
	}

	return nil
}

var metroclusterNodeTypeDrMirroringStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled","unreachable","configured"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterNodeTypeDrMirroringStatePropEnum = append(metroclusterNodeTypeDrMirroringStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_mirroring_state
	// DrMirroringState
	// enabled
	// END DEBUGGING
	// MetroclusterNodeDrMirroringStateEnabled captures enum value "enabled"
	MetroclusterNodeDrMirroringStateEnabled string = "enabled"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_mirroring_state
	// DrMirroringState
	// disabled
	// END DEBUGGING
	// MetroclusterNodeDrMirroringStateDisabled captures enum value "disabled"
	MetroclusterNodeDrMirroringStateDisabled string = "disabled"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_mirroring_state
	// DrMirroringState
	// unreachable
	// END DEBUGGING
	// MetroclusterNodeDrMirroringStateUnreachable captures enum value "unreachable"
	MetroclusterNodeDrMirroringStateUnreachable string = "unreachable"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_mirroring_state
	// DrMirroringState
	// configured
	// END DEBUGGING
	// MetroclusterNodeDrMirroringStateConfigured captures enum value "configured"
	MetroclusterNodeDrMirroringStateConfigured string = "configured"
)

// prop value enum
func (m *MetroclusterNode) validateDrMirroringStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterNodeTypeDrMirroringStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterNode) validateDrMirroringState(formats strfmt.Registry) error {
	if swag.IsZero(m.DrMirroringState) { // not required
		return nil
	}

	// value enum
	if err := m.validateDrMirroringStateEnum("dr_mirroring_state", "body", m.DrMirroringState); err != nil {
		return err
	}

	return nil
}

var metroclusterNodeTypeDrOperationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","switchover_bypassed","switchover_in_progress","switchover_completed","switchover_failed","heal_aggrs_in_progress","heal_aggrs_completed","heal_aggrs_failed","heal_roots_in_progress","heal_roots_completed","heal_roots_failed","switchback_vetoed","switchback_vetocheck_locked","switchback_pre_commit_completed","switchback_in_progress","switchback_completed","switchback_failed","negotiated_switchover_vetoed","negotiated_switchover_vetocheck_locked","negotiated_switchover_pre_commit_completed","negotiated_switchover_in_progress","negotiated_switchover_completed","negotiated_switchover_in_progress_waiting_for_DR_partner","negotiated_switchover_incomplete","negotiated_switchover_failed","negotiated_switchover_failed_on_DR_partner","switchback_recovery_in_progress","switchback_recovery_complete","waiting_for_switchback_recovery","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metroclusterNodeTypeDrOperationStatePropEnum = append(metroclusterNodeTypeDrOperationStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// normal
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNormal captures enum value "normal"
	MetroclusterNodeDrOperationStateNormal string = "normal"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchover_bypassed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchoverBypassed captures enum value "switchover_bypassed"
	MetroclusterNodeDrOperationStateSwitchoverBypassed string = "switchover_bypassed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchover_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchoverInProgress captures enum value "switchover_in_progress"
	MetroclusterNodeDrOperationStateSwitchoverInProgress string = "switchover_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchover_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchoverCompleted captures enum value "switchover_completed"
	MetroclusterNodeDrOperationStateSwitchoverCompleted string = "switchover_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchover_failed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchoverFailed captures enum value "switchover_failed"
	MetroclusterNodeDrOperationStateSwitchoverFailed string = "switchover_failed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_aggrs_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealAggrsInProgress captures enum value "heal_aggrs_in_progress"
	MetroclusterNodeDrOperationStateHealAggrsInProgress string = "heal_aggrs_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_aggrs_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealAggrsCompleted captures enum value "heal_aggrs_completed"
	MetroclusterNodeDrOperationStateHealAggrsCompleted string = "heal_aggrs_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_aggrs_failed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealAggrsFailed captures enum value "heal_aggrs_failed"
	MetroclusterNodeDrOperationStateHealAggrsFailed string = "heal_aggrs_failed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_roots_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealRootsInProgress captures enum value "heal_roots_in_progress"
	MetroclusterNodeDrOperationStateHealRootsInProgress string = "heal_roots_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_roots_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealRootsCompleted captures enum value "heal_roots_completed"
	MetroclusterNodeDrOperationStateHealRootsCompleted string = "heal_roots_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// heal_roots_failed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateHealRootsFailed captures enum value "heal_roots_failed"
	MetroclusterNodeDrOperationStateHealRootsFailed string = "heal_roots_failed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_vetoed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackVetoed captures enum value "switchback_vetoed"
	MetroclusterNodeDrOperationStateSwitchbackVetoed string = "switchback_vetoed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_vetocheck_locked
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackVetocheckLocked captures enum value "switchback_vetocheck_locked"
	MetroclusterNodeDrOperationStateSwitchbackVetocheckLocked string = "switchback_vetocheck_locked"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_pre_commit_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackPreCommitCompleted captures enum value "switchback_pre_commit_completed"
	MetroclusterNodeDrOperationStateSwitchbackPreCommitCompleted string = "switchback_pre_commit_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackInProgress captures enum value "switchback_in_progress"
	MetroclusterNodeDrOperationStateSwitchbackInProgress string = "switchback_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackCompleted captures enum value "switchback_completed"
	MetroclusterNodeDrOperationStateSwitchbackCompleted string = "switchback_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_failed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackFailed captures enum value "switchback_failed"
	MetroclusterNodeDrOperationStateSwitchbackFailed string = "switchback_failed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_vetoed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverVetoed captures enum value "negotiated_switchover_vetoed"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverVetoed string = "negotiated_switchover_vetoed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_vetocheck_locked
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverVetocheckLocked captures enum value "negotiated_switchover_vetocheck_locked"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverVetocheckLocked string = "negotiated_switchover_vetocheck_locked"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_pre_commit_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverPreCommitCompleted captures enum value "negotiated_switchover_pre_commit_completed"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverPreCommitCompleted string = "negotiated_switchover_pre_commit_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverInProgress captures enum value "negotiated_switchover_in_progress"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverInProgress string = "negotiated_switchover_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_completed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverCompleted captures enum value "negotiated_switchover_completed"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverCompleted string = "negotiated_switchover_completed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_in_progress_waiting_for_DR_partner
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverInProgressWaitingForDRPartner captures enum value "negotiated_switchover_in_progress_waiting_for_DR_partner"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverInProgressWaitingForDRPartner string = "negotiated_switchover_in_progress_waiting_for_DR_partner"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_incomplete
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverIncomplete captures enum value "negotiated_switchover_incomplete"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverIncomplete string = "negotiated_switchover_incomplete"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_failed
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverFailed captures enum value "negotiated_switchover_failed"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverFailed string = "negotiated_switchover_failed"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// negotiated_switchover_failed_on_DR_partner
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateNegotiatedSwitchoverFailedOnDRPartner captures enum value "negotiated_switchover_failed_on_DR_partner"
	MetroclusterNodeDrOperationStateNegotiatedSwitchoverFailedOnDRPartner string = "negotiated_switchover_failed_on_DR_partner"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_recovery_in_progress
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackRecoveryInProgress captures enum value "switchback_recovery_in_progress"
	MetroclusterNodeDrOperationStateSwitchbackRecoveryInProgress string = "switchback_recovery_in_progress"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// switchback_recovery_complete
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateSwitchbackRecoveryComplete captures enum value "switchback_recovery_complete"
	MetroclusterNodeDrOperationStateSwitchbackRecoveryComplete string = "switchback_recovery_complete"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// waiting_for_switchback_recovery
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateWaitingForSwitchbackRecovery captures enum value "waiting_for_switchback_recovery"
	MetroclusterNodeDrOperationStateWaitingForSwitchbackRecovery string = "waiting_for_switchback_recovery"

	// BEGIN DEBUGGING
	// metrocluster_node
	// MetroclusterNode
	// dr_operation_state
	// DrOperationState
	// unknown
	// END DEBUGGING
	// MetroclusterNodeDrOperationStateUnknown captures enum value "unknown"
	MetroclusterNodeDrOperationStateUnknown string = "unknown"
)

// prop value enum
func (m *MetroclusterNode) validateDrOperationStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metroclusterNodeTypeDrOperationStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetroclusterNode) validateDrOperationState(formats strfmt.Registry) error {
	if swag.IsZero(m.DrOperationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateDrOperationStateEnum("dr_operation_state", "body", m.DrOperationState); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster node based on the context it is used
func (m *MetroclusterNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutomaticUso(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurationState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrMirroringState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrOperationState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterNode) contextValidateAutomaticUso(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "automatic_uso", "body", m.AutomaticUso); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterNode) contextValidateConfigurationState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "configuration_state", "body", string(m.ConfigurationState)); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) contextValidateDrGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dr_group_id", "body", int64(m.DrGroupID)); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) contextValidateDrMirroringState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dr_mirroring_state", "body", string(m.DrMirroringState)); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) contextValidateDrOperationState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dr_operation_state", "body", string(m.DrOperationState)); err != nil {
		return err
	}

	return nil
}

func (m *MetroclusterNode) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterNodeCluster metrocluster node cluster
//
// swagger:model MetroclusterNodeCluster
type MetroclusterNodeCluster struct {

	// links
	Links *MetroclusterNodeClusterLinks `json:"_links,omitempty"`

	// name
	// Example: cluster1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this metrocluster node cluster
func (m *MetroclusterNodeCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterNodeCluster) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster node cluster based on the context it is used
func (m *MetroclusterNodeCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNodeCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNodeCluster) UnmarshalBinary(b []byte) error {
	var res MetroclusterNodeCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterNodeClusterLinks metrocluster node cluster links
//
// swagger:model MetroclusterNodeClusterLinks
type MetroclusterNodeClusterLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster node cluster links
func (m *MetroclusterNodeClusterLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeClusterLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster node cluster links based on the context it is used
func (m *MetroclusterNodeClusterLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeClusterLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNodeClusterLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNodeClusterLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterNodeClusterLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterNodeNode metrocluster node node
//
// swagger:model MetroclusterNodeNode
type MetroclusterNodeNode struct {

	// links
	Links *MetroclusterNodeNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster node node
func (m *MetroclusterNodeNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster node node based on the context it is used
func (m *MetroclusterNodeNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNodeNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNodeNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterNodeNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterNodeNodeLinks metrocluster node node links
//
// swagger:model MetroclusterNodeNodeLinks
type MetroclusterNodeNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster node node links
func (m *MetroclusterNodeNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeNodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster node node links based on the context it is used
func (m *MetroclusterNodeNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterNodeNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterNodeNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterNodeNodeLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterNodeNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}