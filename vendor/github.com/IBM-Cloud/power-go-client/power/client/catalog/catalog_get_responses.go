// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// CatalogGetReader is a Reader for the CatalogGet structure.
type CatalogGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCatalogGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCatalogGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCatalogGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/catalog] catalog.get", response, response.Code())
	}
}

// NewCatalogGetOK creates a CatalogGetOK with default headers values
func NewCatalogGetOK() *CatalogGetOK {
	return &CatalogGetOK{}
}

/*
CatalogGetOK describes a response with status code 200, with default header values.

catalog response
*/
type CatalogGetOK struct {
	Payload *models.Catalog
}

// IsSuccess returns true when this catalog get o k response has a 2xx status code
func (o *CatalogGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog get o k response has a 3xx status code
func (o *CatalogGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get o k response has a 4xx status code
func (o *CatalogGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog get o k response has a 5xx status code
func (o *CatalogGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get o k response a status code equal to that given
func (o *CatalogGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog get o k response
func (o *CatalogGetOK) Code() int {
	return 200
}

func (o *CatalogGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetOK  %+v", 200, o.Payload)
}

func (o *CatalogGetOK) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetOK  %+v", 200, o.Payload)
}

func (o *CatalogGetOK) GetPayload() *models.Catalog {
	return o.Payload
}

func (o *CatalogGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Catalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogGetBadRequest creates a CatalogGetBadRequest with default headers values
func NewCatalogGetBadRequest() *CatalogGetBadRequest {
	return &CatalogGetBadRequest{}
}

/*
CatalogGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CatalogGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this catalog get bad request response has a 2xx status code
func (o *CatalogGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog get bad request response has a 3xx status code
func (o *CatalogGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get bad request response has a 4xx status code
func (o *CatalogGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog get bad request response has a 5xx status code
func (o *CatalogGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get bad request response a status code equal to that given
func (o *CatalogGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the catalog get bad request response
func (o *CatalogGetBadRequest) Code() int {
	return 400
}

func (o *CatalogGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogGetBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetBadRequest  %+v", 400, o.Payload)
}

func (o *CatalogGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CatalogGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogGetUnauthorized creates a CatalogGetUnauthorized with default headers values
func NewCatalogGetUnauthorized() *CatalogGetUnauthorized {
	return &CatalogGetUnauthorized{}
}

/*
CatalogGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CatalogGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this catalog get unauthorized response has a 2xx status code
func (o *CatalogGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog get unauthorized response has a 3xx status code
func (o *CatalogGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get unauthorized response has a 4xx status code
func (o *CatalogGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog get unauthorized response has a 5xx status code
func (o *CatalogGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get unauthorized response a status code equal to that given
func (o *CatalogGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the catalog get unauthorized response
func (o *CatalogGetUnauthorized) Code() int {
	return 401
}

func (o *CatalogGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetUnauthorized  %+v", 401, o.Payload)
}

func (o *CatalogGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CatalogGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogGetForbidden creates a CatalogGetForbidden with default headers values
func NewCatalogGetForbidden() *CatalogGetForbidden {
	return &CatalogGetForbidden{}
}

/*
CatalogGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CatalogGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this catalog get forbidden response has a 2xx status code
func (o *CatalogGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog get forbidden response has a 3xx status code
func (o *CatalogGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get forbidden response has a 4xx status code
func (o *CatalogGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog get forbidden response has a 5xx status code
func (o *CatalogGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get forbidden response a status code equal to that given
func (o *CatalogGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the catalog get forbidden response
func (o *CatalogGetForbidden) Code() int {
	return 403
}

func (o *CatalogGetForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetForbidden  %+v", 403, o.Payload)
}

func (o *CatalogGetForbidden) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetForbidden  %+v", 403, o.Payload)
}

func (o *CatalogGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CatalogGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogGetNotFound creates a CatalogGetNotFound with default headers values
func NewCatalogGetNotFound() *CatalogGetNotFound {
	return &CatalogGetNotFound{}
}

/*
CatalogGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CatalogGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this catalog get not found response has a 2xx status code
func (o *CatalogGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this catalog get not found response has a 3xx status code
func (o *CatalogGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get not found response has a 4xx status code
func (o *CatalogGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this catalog get not found response has a 5xx status code
func (o *CatalogGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get not found response a status code equal to that given
func (o *CatalogGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the catalog get not found response
func (o *CatalogGetNotFound) Code() int {
	return 404
}

func (o *CatalogGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetNotFound  %+v", 404, o.Payload)
}

func (o *CatalogGetNotFound) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetNotFound  %+v", 404, o.Payload)
}

func (o *CatalogGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CatalogGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}