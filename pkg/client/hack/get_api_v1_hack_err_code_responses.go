// Code generated by go-swagger; DO NOT EDIT.

package hack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/avanibbles/flowflow/pkg/models"
)

// GetAPIV1HackErrCodeReader is a Reader for the GetAPIV1HackErrCode structure.
type GetAPIV1HackErrCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV1HackErrCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV1HackErrCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV1HackErrCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAPIV1HackErrCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIV1HackErrCodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV1HackErrCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAPIV1HackErrCodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIV1HackErrCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV1HackErrCodeOK creates a GetAPIV1HackErrCodeOK with default headers values
func NewGetAPIV1HackErrCodeOK() *GetAPIV1HackErrCodeOK {
	return &GetAPIV1HackErrCodeOK{}
}

/* GetAPIV1HackErrCodeOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIV1HackErrCodeOK struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeOK  %+v", 200, o.Payload)
}
func (o *GetAPIV1HackErrCodeOK) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeBadRequest creates a GetAPIV1HackErrCodeBadRequest with default headers values
func NewGetAPIV1HackErrCodeBadRequest() *GetAPIV1HackErrCodeBadRequest {
	return &GetAPIV1HackErrCodeBadRequest{}
}

/* GetAPIV1HackErrCodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAPIV1HackErrCodeBadRequest struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeBadRequest  %+v", 400, o.Payload)
}
func (o *GetAPIV1HackErrCodeBadRequest) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeUnauthorized creates a GetAPIV1HackErrCodeUnauthorized with default headers values
func NewGetAPIV1HackErrCodeUnauthorized() *GetAPIV1HackErrCodeUnauthorized {
	return &GetAPIV1HackErrCodeUnauthorized{}
}

/* GetAPIV1HackErrCodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAPIV1HackErrCodeUnauthorized struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAPIV1HackErrCodeUnauthorized) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeForbidden creates a GetAPIV1HackErrCodeForbidden with default headers values
func NewGetAPIV1HackErrCodeForbidden() *GetAPIV1HackErrCodeForbidden {
	return &GetAPIV1HackErrCodeForbidden{}
}

/* GetAPIV1HackErrCodeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAPIV1HackErrCodeForbidden struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIV1HackErrCodeForbidden) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeNotFound creates a GetAPIV1HackErrCodeNotFound with default headers values
func NewGetAPIV1HackErrCodeNotFound() *GetAPIV1HackErrCodeNotFound {
	return &GetAPIV1HackErrCodeNotFound{}
}

/* GetAPIV1HackErrCodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAPIV1HackErrCodeNotFound struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIV1HackErrCodeNotFound) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeConflict creates a GetAPIV1HackErrCodeConflict with default headers values
func NewGetAPIV1HackErrCodeConflict() *GetAPIV1HackErrCodeConflict {
	return &GetAPIV1HackErrCodeConflict{}
}

/* GetAPIV1HackErrCodeConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetAPIV1HackErrCodeConflict struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeConflict) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeConflict  %+v", 409, o.Payload)
}
func (o *GetAPIV1HackErrCodeConflict) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV1HackErrCodeInternalServerError creates a GetAPIV1HackErrCodeInternalServerError with default headers values
func NewGetAPIV1HackErrCodeInternalServerError() *GetAPIV1HackErrCodeInternalServerError {
	return &GetAPIV1HackErrCodeInternalServerError{}
}

/* GetAPIV1HackErrCodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAPIV1HackErrCodeInternalServerError struct {
	Payload *models.ApimodelsHTTPError
}

func (o *GetAPIV1HackErrCodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/hack/err/{code}][%d] getApiV1HackErrCodeInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIV1HackErrCodeInternalServerError) GetPayload() *models.ApimodelsHTTPError {
	return o.Payload
}

func (o *GetAPIV1HackErrCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApimodelsHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
