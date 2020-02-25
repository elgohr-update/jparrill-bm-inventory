// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/filanov/bm-inventory/models"
)

// RegisterNodeCreatedCode is the HTTP code returned for type RegisterNodeCreated
const RegisterNodeCreatedCode int = 201

/*RegisterNodeCreated Created

swagger:response registerNodeCreated
*/
type RegisterNodeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RegisteredNode `json:"body,omitempty"`
}

// NewRegisterNodeCreated creates RegisterNodeCreated with default headers values
func NewRegisterNodeCreated() *RegisterNodeCreated {

	return &RegisterNodeCreated{}
}

// WithPayload adds the payload to the register node created response
func (o *RegisterNodeCreated) WithPayload(payload *models.RegisteredNode) *RegisterNodeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register node created response
func (o *RegisterNodeCreated) SetPayload(payload *models.RegisteredNode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterNodeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterNodeMethodNotAllowedCode is the HTTP code returned for type RegisterNodeMethodNotAllowed
const RegisterNodeMethodNotAllowedCode int = 405

/*RegisterNodeMethodNotAllowed Invalid input

swagger:response registerNodeMethodNotAllowed
*/
type RegisterNodeMethodNotAllowed struct {
}

// NewRegisterNodeMethodNotAllowed creates RegisterNodeMethodNotAllowed with default headers values
func NewRegisterNodeMethodNotAllowed() *RegisterNodeMethodNotAllowed {

	return &RegisterNodeMethodNotAllowed{}
}

// WriteResponse to the client
func (o *RegisterNodeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}