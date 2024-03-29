// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMaybeCreatedCode is the HTTP code returned for type PostMaybeCreated
const PostMaybeCreatedCode int = 201

/*PostMaybeCreated Created

swagger:response postMaybeCreated
*/
type PostMaybeCreated struct {
}

// NewPostMaybeCreated creates PostMaybeCreated with default headers values
func NewPostMaybeCreated() *PostMaybeCreated {

	return &PostMaybeCreated{}
}

// WriteResponse to the client
func (o *PostMaybeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostMaybeForbiddenCode is the HTTP code returned for type PostMaybeForbidden
const PostMaybeForbiddenCode int = 403

/*PostMaybeForbidden Forbidden

swagger:response postMaybeForbidden
*/
type PostMaybeForbidden struct {
}

// NewPostMaybeForbidden creates PostMaybeForbidden with default headers values
func NewPostMaybeForbidden() *PostMaybeForbidden {

	return &PostMaybeForbidden{}
}

// WriteResponse to the client
func (o *PostMaybeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostMaybeInternalServerErrorCode is the HTTP code returned for type PostMaybeInternalServerError
const PostMaybeInternalServerErrorCode int = 500

/*PostMaybeInternalServerError Internal Server Error

swagger:response postMaybeInternalServerError
*/
type PostMaybeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostMaybeInternalServerError creates PostMaybeInternalServerError with default headers values
func NewPostMaybeInternalServerError() *PostMaybeInternalServerError {

	return &PostMaybeInternalServerError{}
}

// WithPayload adds the payload to the post maybe internal server error response
func (o *PostMaybeInternalServerError) WithPayload(payload string) *PostMaybeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post maybe internal server error response
func (o *PostMaybeInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMaybeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
