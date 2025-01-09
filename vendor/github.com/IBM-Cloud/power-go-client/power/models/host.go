// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Host Host description
//
// swagger:model Host
type Host struct {

	// Capacities of the host
	Capacity *HostCapacity `json:"capacity,omitempty"`

	// crn
	Crn CRN `json:"crn,omitempty"`

	// Name of the host (chosen by the user)
	DisplayName string `json:"displayName,omitempty"`

	// Information about the owning host group
	HostGroup *HostGroupSummary `json:"hostGroup,omitempty"`

	// current physical ID of the host. Keep internal
	HostReference int64 `json:"hostReference,omitempty"`

	// ID of the host
	ID string `json:"id,omitempty"`

	// State of the host (up/down)
	State string `json:"state,omitempty"`

	// Status of the host (enabled/disabled)
	Status string `json:"status,omitempty"`

	// System type
	SysType string `json:"sysType,omitempty"`

	// user tags
	UserTags Tags `json:"userTags,omitempty"`
}

// Validate validates this host
func (m *Host) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateCrn(formats strfmt.Registry) error {
	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

func (m *Host) validateHostGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.HostGroup) { // not required
		return nil
	}

	if m.HostGroup != nil {
		if err := m.HostGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateUserTags(formats strfmt.Registry) error {
	if swag.IsZero(m.UserTags) { // not required
		return nil
	}

	if err := m.UserTags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

// ContextValidate validate this host based on the context it is used
func (m *Host) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {

		if swag.IsZero(m.Capacity) { // not required
			return nil
		}

		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateCrn(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

func (m *Host) contextValidateHostGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.HostGroup != nil {

		if swag.IsZero(m.HostGroup) { // not required
			return nil
		}

		if err := m.HostGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hostGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateUserTags(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserTags.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Host) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Host) UnmarshalBinary(b []byte) error {
	var res Host
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
