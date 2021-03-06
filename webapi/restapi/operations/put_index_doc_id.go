// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutIndexDocIDHandlerFunc turns a function with the right signature into a put index doc id handler
type PutIndexDocIDHandlerFunc func(PutIndexDocIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutIndexDocIDHandlerFunc) Handle(params PutIndexDocIDParams) middleware.Responder {
	return fn(params)
}

// PutIndexDocIDHandler interface for that can handle valid put index doc id params
type PutIndexDocIDHandler interface {
	Handle(PutIndexDocIDParams) middleware.Responder
}

// NewPutIndexDocID creates a new http.Handler for the put index doc id operation
func NewPutIndexDocID(ctx *middleware.Context, handler PutIndexDocIDHandler) *PutIndexDocID {
	return &PutIndexDocID{Context: ctx, Handler: handler}
}

/*PutIndexDocID swagger:route PUT /{index}/_doc/{_id} putIndexDocId

Update an existing JSON document to the specified index and makes it searchable.

*/
type PutIndexDocID struct {
	Context *middleware.Context
	Handler PutIndexDocIDHandler
}

func (o *PutIndexDocID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutIndexDocIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutIndexDocIDBadRequestBody put index doc ID bad request body
//
// swagger:model PutIndexDocIDBadRequestBody
type PutIndexDocIDBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put index doc ID bad request body
func (o *PutIndexDocIDBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIndexDocIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIndexDocIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutIndexDocIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutIndexDocIDNotFoundBody put index doc ID not found body
//
// swagger:model PutIndexDocIDNotFoundBody
type PutIndexDocIDNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put index doc ID not found body
func (o *PutIndexDocIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIndexDocIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIndexDocIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutIndexDocIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
