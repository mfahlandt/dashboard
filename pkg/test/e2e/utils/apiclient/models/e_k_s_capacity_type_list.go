// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// EKSCapacityTypeList EKSCapacityTypeList represents a list of EKS Capacity Types for node group.
//
// swagger:model EKSCapacityTypeList
type EKSCapacityTypeList []string

// Validate validates this e k s capacity type list
func (m EKSCapacityTypeList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e k s capacity type list based on context it is used
func (m EKSCapacityTypeList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}