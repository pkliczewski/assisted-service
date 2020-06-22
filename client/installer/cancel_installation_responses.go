// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// CancelInstallationReader is a Reader for the CancelInstallation structure.
type CancelInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCancelInstallationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCancelInstallationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCancelInstallationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelInstallationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelInstallationAccepted creates a CancelInstallationAccepted with default headers values
func NewCancelInstallationAccepted() *CancelInstallationAccepted {
	return &CancelInstallationAccepted{}
}

/*CancelInstallationAccepted handles this case with default header values.

Success.
*/
type CancelInstallationAccepted struct {
	Payload *models.Cluster
}

func (o *CancelInstallationAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/cancel][%d] cancelInstallationAccepted  %+v", 202, o.Payload)
}

func (o *CancelInstallationAccepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *CancelInstallationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInstallationNotFound creates a CancelInstallationNotFound with default headers values
func NewCancelInstallationNotFound() *CancelInstallationNotFound {
	return &CancelInstallationNotFound{}
}

/*CancelInstallationNotFound handles this case with default header values.

Error.
*/
type CancelInstallationNotFound struct {
	Payload *models.Error
}

func (o *CancelInstallationNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/cancel][%d] cancelInstallationNotFound  %+v", 404, o.Payload)
}

func (o *CancelInstallationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelInstallationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInstallationConflict creates a CancelInstallationConflict with default headers values
func NewCancelInstallationConflict() *CancelInstallationConflict {
	return &CancelInstallationConflict{}
}

/*CancelInstallationConflict handles this case with default header values.

Error.
*/
type CancelInstallationConflict struct {
	Payload *models.Error
}

func (o *CancelInstallationConflict) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/cancel][%d] cancelInstallationConflict  %+v", 409, o.Payload)
}

func (o *CancelInstallationConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelInstallationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelInstallationInternalServerError creates a CancelInstallationInternalServerError with default headers values
func NewCancelInstallationInternalServerError() *CancelInstallationInternalServerError {
	return &CancelInstallationInternalServerError{}
}

/*CancelInstallationInternalServerError handles this case with default header values.

Error.
*/
type CancelInstallationInternalServerError struct {
	Payload *models.Error
}

func (o *CancelInstallationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/cancel][%d] cancelInstallationInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelInstallationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelInstallationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
