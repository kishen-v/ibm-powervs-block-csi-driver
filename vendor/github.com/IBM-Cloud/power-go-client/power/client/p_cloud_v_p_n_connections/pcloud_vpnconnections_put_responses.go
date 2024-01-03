// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPutReader is a Reader for the PcloudVpnconnectionsPut structure.
type PcloudVpnconnectionsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}] pcloud.vpnconnections.put", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPutOK creates a PcloudVpnconnectionsPutOK with default headers values
func NewPcloudVpnconnectionsPutOK() *PcloudVpnconnectionsPutOK {
	return &PcloudVpnconnectionsPutOK{}
}

/*
PcloudVpnconnectionsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsPutOK struct {
	Payload *models.VPNConnection
}

// IsSuccess returns true when this pcloud vpnconnections put o k response has a 2xx status code
func (o *PcloudVpnconnectionsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections put o k response has a 3xx status code
func (o *PcloudVpnconnectionsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put o k response has a 4xx status code
func (o *PcloudVpnconnectionsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections put o k response has a 5xx status code
func (o *PcloudVpnconnectionsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put o k response a status code equal to that given
func (o *PcloudVpnconnectionsPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud vpnconnections put o k response
func (o *PcloudVpnconnectionsPutOK) Code() int {
	return 200
}

func (o *PcloudVpnconnectionsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPutOK) GetPayload() *models.VPNConnection {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutBadRequest creates a PcloudVpnconnectionsPutBadRequest with default headers values
func NewPcloudVpnconnectionsPutBadRequest() *PcloudVpnconnectionsPutBadRequest {
	return &PcloudVpnconnectionsPutBadRequest{}
}

/*
PcloudVpnconnectionsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put bad request response has a 2xx status code
func (o *PcloudVpnconnectionsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put bad request response has a 3xx status code
func (o *PcloudVpnconnectionsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put bad request response has a 4xx status code
func (o *PcloudVpnconnectionsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections put bad request response has a 5xx status code
func (o *PcloudVpnconnectionsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put bad request response a status code equal to that given
func (o *PcloudVpnconnectionsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections put bad request response
func (o *PcloudVpnconnectionsPutBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutUnauthorized creates a PcloudVpnconnectionsPutUnauthorized with default headers values
func NewPcloudVpnconnectionsPutUnauthorized() *PcloudVpnconnectionsPutUnauthorized {
	return &PcloudVpnconnectionsPutUnauthorized{}
}

/*
PcloudVpnconnectionsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections put unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections put unauthorized response
func (o *PcloudVpnconnectionsPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutForbidden creates a PcloudVpnconnectionsPutForbidden with default headers values
func NewPcloudVpnconnectionsPutForbidden() *PcloudVpnconnectionsPutForbidden {
	return &PcloudVpnconnectionsPutForbidden{}
}

/*
PcloudVpnconnectionsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections put forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections put forbidden response
func (o *PcloudVpnconnectionsPutForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutNotFound creates a PcloudVpnconnectionsPutNotFound with default headers values
func NewPcloudVpnconnectionsPutNotFound() *PcloudVpnconnectionsPutNotFound {
	return &PcloudVpnconnectionsPutNotFound{}
}

/*
PcloudVpnconnectionsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put not found response has a 2xx status code
func (o *PcloudVpnconnectionsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put not found response has a 3xx status code
func (o *PcloudVpnconnectionsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put not found response has a 4xx status code
func (o *PcloudVpnconnectionsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections put not found response has a 5xx status code
func (o *PcloudVpnconnectionsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put not found response a status code equal to that given
func (o *PcloudVpnconnectionsPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections put not found response
func (o *PcloudVpnconnectionsPutNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutUnprocessableEntity creates a PcloudVpnconnectionsPutUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPutUnprocessableEntity() *PcloudVpnconnectionsPutUnprocessableEntity {
	return &PcloudVpnconnectionsPutUnprocessableEntity{}
}

/*
PcloudVpnconnectionsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put unprocessable entity response has a 2xx status code
func (o *PcloudVpnconnectionsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put unprocessable entity response has a 3xx status code
func (o *PcloudVpnconnectionsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put unprocessable entity response has a 4xx status code
func (o *PcloudVpnconnectionsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections put unprocessable entity response has a 5xx status code
func (o *PcloudVpnconnectionsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections put unprocessable entity response a status code equal to that given
func (o *PcloudVpnconnectionsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud vpnconnections put unprocessable entity response
func (o *PcloudVpnconnectionsPutUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPutInternalServerError creates a PcloudVpnconnectionsPutInternalServerError with default headers values
func NewPcloudVpnconnectionsPutInternalServerError() *PcloudVpnconnectionsPutInternalServerError {
	return &PcloudVpnconnectionsPutInternalServerError{}
}

/*
PcloudVpnconnectionsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections put internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections put internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections put internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections put internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections put internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections put internal server error response
func (o *PcloudVpnconnectionsPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}