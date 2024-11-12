// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetHelloHandlerFunc turns a function with the right signature into a get hello handler
type GetHelloHandlerFunc func(GetHelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHelloHandlerFunc) Handle(params GetHelloParams) middleware.Responder {
	return fn(params)
}

// GetHelloHandler interface for that can handle valid get hello params
type GetHelloHandler interface {
	Handle(GetHelloParams) middleware.Responder
}

// NewGetHello creates a new http.Handler for the get hello operation
func NewGetHello(ctx *middleware.Context, handler GetHelloHandler) *GetHello {
	return &GetHello{Context: ctx, Handler: handler}
}

/*
	GetHello swagger:route GET /hello getHello

# Say Hello

Returns a greeting message.
*/
type GetHello struct {
	Context *middleware.Context
	Handler GetHelloHandler
}

func (o *GetHello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHelloParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetHelloOKBody get hello o k body
//
// swagger:model GetHelloOKBody
type GetHelloOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get hello o k body
func (o *GetHelloOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get hello o k body based on context it is used
func (o *GetHelloOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHelloOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHelloOKBody) UnmarshalBinary(b []byte) error {
	var res GetHelloOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}