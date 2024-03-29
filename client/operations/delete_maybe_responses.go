// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteMaybeReader is a Reader for the DeleteMaybe structure.
type DeleteMaybeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMaybeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteMaybeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewDeleteMaybeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMaybeNoContent creates a DeleteMaybeNoContent with default headers values
func NewDeleteMaybeNoContent() *DeleteMaybeNoContent {
	return &DeleteMaybeNoContent{}
}

/*DeleteMaybeNoContent handles this case with default header values.

No Content
*/
type DeleteMaybeNoContent struct {
}

func (o *DeleteMaybeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /maybe][%d] deleteMaybeNoContent ", 204)
}

func (o *DeleteMaybeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMaybeInternalServerError creates a DeleteMaybeInternalServerError with default headers values
func NewDeleteMaybeInternalServerError() *DeleteMaybeInternalServerError {
	return &DeleteMaybeInternalServerError{}
}

/*DeleteMaybeInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteMaybeInternalServerError struct {
	Payload string
}

func (o *DeleteMaybeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /maybe][%d] deleteMaybeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMaybeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
