// Code generated by go-swagger; DO NOT EDIT.

package management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// DeleteMgmtClusterOKCode is the HTTP code returned for type DeleteMgmtClusterOK
const DeleteMgmtClusterOKCode int = 200

/*DeleteMgmtClusterOK Cluster deletion submitted.

swagger:response deleteMgmtClusterOK
*/
type DeleteMgmtClusterOK struct {
}

// NewDeleteMgmtClusterOK creates DeleteMgmtClusterOK with default headers values
func NewDeleteMgmtClusterOK() *DeleteMgmtClusterOK {

	return &DeleteMgmtClusterOK{}
}

// WriteResponse to the client
func (o *DeleteMgmtClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteMgmtClusterBadRequestCode is the HTTP code returned for type DeleteMgmtClusterBadRequest
const DeleteMgmtClusterBadRequestCode int = 400

/*DeleteMgmtClusterBadRequest Bad request

swagger:response deleteMgmtClusterBadRequest
*/
type DeleteMgmtClusterBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMgmtClusterBadRequest creates DeleteMgmtClusterBadRequest with default headers values
func NewDeleteMgmtClusterBadRequest() *DeleteMgmtClusterBadRequest {

	return &DeleteMgmtClusterBadRequest{}
}

// WithPayload adds the payload to the delete mgmt cluster bad request response
func (o *DeleteMgmtClusterBadRequest) WithPayload(payload *models.Error) *DeleteMgmtClusterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete mgmt cluster bad request response
func (o *DeleteMgmtClusterBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMgmtClusterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMgmtClusterUnauthorizedCode is the HTTP code returned for type DeleteMgmtClusterUnauthorized
const DeleteMgmtClusterUnauthorizedCode int = 401

/*DeleteMgmtClusterUnauthorized Incorrect credentials

swagger:response deleteMgmtClusterUnauthorized
*/
type DeleteMgmtClusterUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMgmtClusterUnauthorized creates DeleteMgmtClusterUnauthorized with default headers values
func NewDeleteMgmtClusterUnauthorized() *DeleteMgmtClusterUnauthorized {

	return &DeleteMgmtClusterUnauthorized{}
}

// WithPayload adds the payload to the delete mgmt cluster unauthorized response
func (o *DeleteMgmtClusterUnauthorized) WithPayload(payload *models.Error) *DeleteMgmtClusterUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete mgmt cluster unauthorized response
func (o *DeleteMgmtClusterUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMgmtClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMgmtClusterInternalServerErrorCode is the HTTP code returned for type DeleteMgmtClusterInternalServerError
const DeleteMgmtClusterInternalServerErrorCode int = 500

/*DeleteMgmtClusterInternalServerError Internal server error

swagger:response deleteMgmtClusterInternalServerError
*/
type DeleteMgmtClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMgmtClusterInternalServerError creates DeleteMgmtClusterInternalServerError with default headers values
func NewDeleteMgmtClusterInternalServerError() *DeleteMgmtClusterInternalServerError {

	return &DeleteMgmtClusterInternalServerError{}
}

// WithPayload adds the payload to the delete mgmt cluster internal server error response
func (o *DeleteMgmtClusterInternalServerError) WithPayload(payload *models.Error) *DeleteMgmtClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete mgmt cluster internal server error response
func (o *DeleteMgmtClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMgmtClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
