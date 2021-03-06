// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Document Document
//
// swagger:model Document
type Document struct {

	// content
	Content string `json:"content,omitempty"`

	// lang
	Lang string `json:"lang,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// title
	Title string `json:"title,omitempty"`

	// unique key
	UniqueKey string `json:"unique_key,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
