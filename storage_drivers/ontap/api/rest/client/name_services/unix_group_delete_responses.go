// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// UnixGroupDeleteReader is a Reader for the UnixGroupDelete structure.
type UnixGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupDeleteOK creates a UnixGroupDeleteOK with default headers values
func NewUnixGroupDeleteOK() *UnixGroupDeleteOK {
	return &UnixGroupDeleteOK{}
}

/* UnixGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type UnixGroupDeleteOK struct {
}

func (o *UnixGroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /name-services/unix-groups/{svm.uuid}/{name}][%d] unixGroupDeleteOK ", 200)
}

func (o *UnixGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixGroupDeleteDefault creates a UnixGroupDeleteDefault with default headers values
func NewUnixGroupDeleteDefault(code int) *UnixGroupDeleteDefault {
	return &UnixGroupDeleteDefault{
		_statusCode: code,
	}
}

/* UnixGroupDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type UnixGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unix group delete default response
func (o *UnixGroupDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UnixGroupDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /name-services/unix-groups/{svm.uuid}/{name}][%d] unix_group_delete default  %+v", o._statusCode, o.Payload)
}
func (o *UnixGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}