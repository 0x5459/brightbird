// Code generated by go-swagger; DO NOT EDIT.

package testflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ipfs-force-community/brightbird/models"
)

// ChangGroupRequestReader is a Reader for the ChangGroupRequest structure.
type ChangGroupRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangGroupRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangGroupRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewChangGroupRequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangGroupRequestOK creates a ChangGroupRequestOK with default headers values
func NewChangGroupRequestOK() *ChangGroupRequestOK {
	return &ChangGroupRequestOK{}
}

/*
ChangGroupRequestOK describes a response with status code 200, with default header values.

ChangGroupRequestOK chang group request o k
*/
type ChangGroupRequestOK struct {
}

// IsSuccess returns true when this chang group request o k response has a 2xx status code
func (o *ChangGroupRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chang group request o k response has a 3xx status code
func (o *ChangGroupRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chang group request o k response has a 4xx status code
func (o *ChangGroupRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chang group request o k response has a 5xx status code
func (o *ChangGroupRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chang group request o k response a status code equal to that given
func (o *ChangGroupRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the chang group request o k response
func (o *ChangGroupRequestOK) Code() int {
	return 200
}

func (o *ChangGroupRequestOK) Error() string {
	return fmt.Sprintf("[POST /changegroup][%d] changGroupRequestOK ", 200)
}

func (o *ChangGroupRequestOK) String() string {
	return fmt.Sprintf("[POST /changegroup][%d] changGroupRequestOK ", 200)
}

func (o *ChangGroupRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangGroupRequestServiceUnavailable creates a ChangGroupRequestServiceUnavailable with default headers values
func NewChangGroupRequestServiceUnavailable() *ChangGroupRequestServiceUnavailable {
	return &ChangGroupRequestServiceUnavailable{}
}

/*
ChangGroupRequestServiceUnavailable describes a response with status code 503, with default header values.

apiError
*/
type ChangGroupRequestServiceUnavailable struct {
	Payload *models.APIError
}

// IsSuccess returns true when this chang group request service unavailable response has a 2xx status code
func (o *ChangGroupRequestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this chang group request service unavailable response has a 3xx status code
func (o *ChangGroupRequestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chang group request service unavailable response has a 4xx status code
func (o *ChangGroupRequestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this chang group request service unavailable response has a 5xx status code
func (o *ChangGroupRequestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this chang group request service unavailable response a status code equal to that given
func (o *ChangGroupRequestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the chang group request service unavailable response
func (o *ChangGroupRequestServiceUnavailable) Code() int {
	return 503
}

func (o *ChangGroupRequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /changegroup][%d] changGroupRequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ChangGroupRequestServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /changegroup][%d] changGroupRequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ChangGroupRequestServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ChangGroupRequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
