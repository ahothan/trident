// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ConsistencyGroupSnapshotCollectionGetReader is a Reader for the ConsistencyGroupSnapshotCollectionGet structure.
type ConsistencyGroupSnapshotCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsistencyGroupSnapshotCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConsistencyGroupSnapshotCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConsistencyGroupSnapshotCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConsistencyGroupSnapshotCollectionGetOK creates a ConsistencyGroupSnapshotCollectionGetOK with default headers values
func NewConsistencyGroupSnapshotCollectionGetOK() *ConsistencyGroupSnapshotCollectionGetOK {
	return &ConsistencyGroupSnapshotCollectionGetOK{}
}

/* ConsistencyGroupSnapshotCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ConsistencyGroupSnapshotCollectionGetOK struct {
	Payload *models.ConsistencyGroupSnapshotResponse
}

func (o *ConsistencyGroupSnapshotCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistencyGroupSnapshotCollectionGetOK  %+v", 200, o.Payload)
}
func (o *ConsistencyGroupSnapshotCollectionGetOK) GetPayload() *models.ConsistencyGroupSnapshotResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsistencyGroupSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsistencyGroupSnapshotCollectionGetDefault creates a ConsistencyGroupSnapshotCollectionGetDefault with default headers values
func NewConsistencyGroupSnapshotCollectionGetDefault(code int) *ConsistencyGroupSnapshotCollectionGetDefault {
	return &ConsistencyGroupSnapshotCollectionGetDefault{
		_statusCode: code,
	}
}

/* ConsistencyGroupSnapshotCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ConsistencyGroupSnapshotCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the consistency group snapshot collection get default response
func (o *ConsistencyGroupSnapshotCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ConsistencyGroupSnapshotCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /application/consistency-groups/{consistency_group.uuid}/snapshots][%d] consistency_group_snapshot_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *ConsistencyGroupSnapshotCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConsistencyGroupSnapshotCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}