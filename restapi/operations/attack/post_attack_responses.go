// Code generated by go-swagger; DO NOT EDIT.

package attack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "vegeta-server/models"
)

// PostAttackOKCode is the HTTP code returned for type PostAttackOK
const PostAttackOKCode int = 200

/*PostAttackOK Success

swagger:response postAttackOK
*/
type PostAttackOK struct {

	/*
	  In: Body
	*/
	Payload *models.AttackResponse `json:"body,omitempty"`
}

// NewPostAttackOK creates PostAttackOK with default headers values
func NewPostAttackOK() *PostAttackOK {

	return &PostAttackOK{}
}

// WithPayload adds the payload to the post attack o k response
func (o *PostAttackOK) WithPayload(payload *models.AttackResponse) *PostAttackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post attack o k response
func (o *PostAttackOK) SetPayload(payload *models.AttackResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAttackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
