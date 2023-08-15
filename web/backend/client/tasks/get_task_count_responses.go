// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ipfs-force-community/brightbird/models"
)

// GetTaskCountReader is a Reader for the GetTaskCount structure.
type GetTaskCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTaskCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /task-count] getTaskCount", response, response.Code())
	}
}

// NewGetTaskCountOK creates a GetTaskCountOK with default headers values
func NewGetTaskCountOK() *GetTaskCountOK {
	return &GetTaskCountOK{}
}

/*
GetTaskCountOK describes a response with status code 200, with default header values.

	//todo fix correctstruct
*/
type GetTaskCountOK struct {
	Payload models.MyString
}

// IsSuccess returns true when this get task count o k response has a 2xx status code
func (o *GetTaskCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task count o k response has a 3xx status code
func (o *GetTaskCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task count o k response has a 4xx status code
func (o *GetTaskCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task count o k response has a 5xx status code
func (o *GetTaskCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task count o k response a status code equal to that given
func (o *GetTaskCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get task count o k response
func (o *GetTaskCountOK) Code() int {
	return 200
}

func (o *GetTaskCountOK) Error() string {
	return fmt.Sprintf("[GET /task-count][%d] getTaskCountOK  %+v", 200, o.Payload)
}

func (o *GetTaskCountOK) String() string {
	return fmt.Sprintf("[GET /task-count][%d] getTaskCountOK  %+v", 200, o.Payload)
}

func (o *GetTaskCountOK) GetPayload() models.MyString {
	return o.Payload
}

func (o *GetTaskCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskCountInternalServerError creates a GetTaskCountInternalServerError with default headers values
func NewGetTaskCountInternalServerError() *GetTaskCountInternalServerError {
	return &GetTaskCountInternalServerError{}
}

/*
GetTaskCountInternalServerError describes a response with status code 500, with default header values.

apiError
*/
type GetTaskCountInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this get task count internal server error response has a 2xx status code
func (o *GetTaskCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task count internal server error response has a 3xx status code
func (o *GetTaskCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task count internal server error response has a 4xx status code
func (o *GetTaskCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task count internal server error response has a 5xx status code
func (o *GetTaskCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get task count internal server error response a status code equal to that given
func (o *GetTaskCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get task count internal server error response
func (o *GetTaskCountInternalServerError) Code() int {
	return 500
}

func (o *GetTaskCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /task-count][%d] getTaskCountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /task-count][%d] getTaskCountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskCountInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetTaskCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}