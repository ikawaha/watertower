// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/future-architect/watertower/webapi/models"
)

// GetIndexSourceIDOKCode is the HTTP code returned for type GetIndexSourceIDOK
const GetIndexSourceIDOKCode int = 200

/*GetIndexSourceIDOK OK

swagger:response getIndexSourceIdOK
*/
type GetIndexSourceIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Document `json:"body,omitempty"`
}

// NewGetIndexSourceIDOK creates GetIndexSourceIDOK with default headers values
func NewGetIndexSourceIDOK() *GetIndexSourceIDOK {

	return &GetIndexSourceIDOK{}
}

// WithPayload adds the payload to the get index source Id o k response
func (o *GetIndexSourceIDOK) WithPayload(payload *models.Document) *GetIndexSourceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get index source Id o k response
func (o *GetIndexSourceIDOK) SetPayload(payload *models.Document) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIndexSourceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIndexSourceIDBadRequestCode is the HTTP code returned for type GetIndexSourceIDBadRequest
const GetIndexSourceIDBadRequestCode int = 400

/*GetIndexSourceIDBadRequest Bad Request

swagger:response getIndexSourceIdBadRequest
*/
type GetIndexSourceIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetIndexSourceIDBadRequestBody `json:"body,omitempty"`
}

// NewGetIndexSourceIDBadRequest creates GetIndexSourceIDBadRequest with default headers values
func NewGetIndexSourceIDBadRequest() *GetIndexSourceIDBadRequest {

	return &GetIndexSourceIDBadRequest{}
}

// WithPayload adds the payload to the get index source Id bad request response
func (o *GetIndexSourceIDBadRequest) WithPayload(payload *GetIndexSourceIDBadRequestBody) *GetIndexSourceIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get index source Id bad request response
func (o *GetIndexSourceIDBadRequest) SetPayload(payload *GetIndexSourceIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIndexSourceIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetIndexSourceIDNotFoundCode is the HTTP code returned for type GetIndexSourceIDNotFound
const GetIndexSourceIDNotFoundCode int = 404

/*GetIndexSourceIDNotFound Not Found

swagger:response getIndexSourceIdNotFound
*/
type GetIndexSourceIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetIndexSourceIDNotFoundBody `json:"body,omitempty"`
}

// NewGetIndexSourceIDNotFound creates GetIndexSourceIDNotFound with default headers values
func NewGetIndexSourceIDNotFound() *GetIndexSourceIDNotFound {

	return &GetIndexSourceIDNotFound{}
}

// WithPayload adds the payload to the get index source Id not found response
func (o *GetIndexSourceIDNotFound) WithPayload(payload *GetIndexSourceIDNotFoundBody) *GetIndexSourceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get index source Id not found response
func (o *GetIndexSourceIDNotFound) SetPayload(payload *GetIndexSourceIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIndexSourceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}