// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hunjixin/brightbird/models"
)

// ListPluginParamsReader is a Reader for the ListPluginParams structure.
type ListPluginParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPluginParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPluginParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewListPluginParamsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPluginParamsOK creates a ListPluginParamsOK with default headers values
func NewListPluginParamsOK() *ListPluginParamsOK {
	return &ListPluginParamsOK{}
}

/*
ListPluginParamsOK describes a response with status code 200, with default header values.

pluginDetail
*/
type ListPluginParamsOK struct {
	Payload []*models.PluginDetail
}

// IsSuccess returns true when this list plugin params o k response has a 2xx status code
func (o *ListPluginParamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list plugin params o k response has a 3xx status code
func (o *ListPluginParamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list plugin params o k response has a 4xx status code
func (o *ListPluginParamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list plugin params o k response has a 5xx status code
func (o *ListPluginParamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list plugin params o k response a status code equal to that given
func (o *ListPluginParamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list plugin params o k response
func (o *ListPluginParamsOK) Code() int {
	return 200
}

func (o *ListPluginParamsOK) Error() string {
	return fmt.Sprintf("[GET /plugin][%d] listPluginParamsOK  %+v", 200, o.Payload)
}

func (o *ListPluginParamsOK) String() string {
	return fmt.Sprintf("[GET /plugin][%d] listPluginParamsOK  %+v", 200, o.Payload)
}

func (o *ListPluginParamsOK) GetPayload() []*models.PluginDetail {
	return o.Payload
}

func (o *ListPluginParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginParamsServiceUnavailable creates a ListPluginParamsServiceUnavailable with default headers values
func NewListPluginParamsServiceUnavailable() *ListPluginParamsServiceUnavailable {
	return &ListPluginParamsServiceUnavailable{}
}

/*
ListPluginParamsServiceUnavailable describes a response with status code 503, with default header values.

apiError
*/
type ListPluginParamsServiceUnavailable struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list plugin params service unavailable response has a 2xx status code
func (o *ListPluginParamsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list plugin params service unavailable response has a 3xx status code
func (o *ListPluginParamsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list plugin params service unavailable response has a 4xx status code
func (o *ListPluginParamsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list plugin params service unavailable response has a 5xx status code
func (o *ListPluginParamsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list plugin params service unavailable response a status code equal to that given
func (o *ListPluginParamsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the list plugin params service unavailable response
func (o *ListPluginParamsServiceUnavailable) Code() int {
	return 503
}

func (o *ListPluginParamsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /plugin][%d] listPluginParamsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListPluginParamsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /plugin][%d] listPluginParamsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListPluginParamsServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListPluginParamsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
