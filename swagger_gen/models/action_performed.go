// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionPerformed when evaluating a mock, an object capturing what would happen if the mock's action is performed
//
// swagger:model ActionPerformed
type ActionPerformed struct {

	// publish kafka action performed
	PublishKafkaActionPerformed *PublishKafkaActionPerformed `json:"publish_kafka_action_performed,omitempty"`

	// reply http action performed
	ReplyHTTPActionPerformed *ReplyHTTPActionPerformed `json:"reply_http_action_performed,omitempty"`
}

// Validate validates this action performed
func (m *ActionPerformed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublishKafkaActionPerformed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyHTTPActionPerformed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionPerformed) validatePublishKafkaActionPerformed(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishKafkaActionPerformed) { // not required
		return nil
	}

	if m.PublishKafkaActionPerformed != nil {
		if err := m.PublishKafkaActionPerformed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publish_kafka_action_performed")
			}
			return err
		}
	}

	return nil
}

func (m *ActionPerformed) validateReplyHTTPActionPerformed(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyHTTPActionPerformed) { // not required
		return nil
	}

	if m.ReplyHTTPActionPerformed != nil {
		if err := m.ReplyHTTPActionPerformed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_http_action_performed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionPerformed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionPerformed) UnmarshalBinary(b []byte) error {
	var res ActionPerformed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
