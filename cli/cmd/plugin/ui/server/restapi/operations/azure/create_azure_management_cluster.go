// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateAzureManagementClusterHandlerFunc turns a function with the right signature into a create azure management cluster handler
type CreateAzureManagementClusterHandlerFunc func(CreateAzureManagementClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAzureManagementClusterHandlerFunc) Handle(params CreateAzureManagementClusterParams) middleware.Responder {
	return fn(params)
}

// CreateAzureManagementClusterHandler interface for that can handle valid create azure management cluster params
type CreateAzureManagementClusterHandler interface {
	Handle(CreateAzureManagementClusterParams) middleware.Responder
}

// NewCreateAzureManagementCluster creates a new http.Handler for the create azure management cluster operation
func NewCreateAzureManagementCluster(ctx *middleware.Context, handler CreateAzureManagementClusterHandler) *CreateAzureManagementCluster {
	return &CreateAzureManagementCluster{Context: ctx, Handler: handler}
}

/*CreateAzureManagementCluster swagger:route POST /api/provider/azure/create azure createAzureManagementCluster

Create Azure management cluster

*/
type CreateAzureManagementCluster struct {
	Context *middleware.Context
	Handler CreateAzureManagementClusterHandler
}

func (o *CreateAzureManagementCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAzureManagementClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
