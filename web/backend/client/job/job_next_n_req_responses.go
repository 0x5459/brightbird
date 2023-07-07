// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hunjixin/brightbird/models"
)

// JobNextNReqReader is a Reader for the JobNextNReq structure.
type JobNextNReqReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobNextNReqReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobNextNReqOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewJobNextNReqServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobNextNReqOK creates a JobNextNReqOK with default headers values
func NewJobNextNReqOK() *JobNextNReqOK {
	return &JobNextNReqOK{}
}

/*
JobNextNReqOK describes a response with status code 200, with default header values.

int64Arr
*/
type JobNextNReqOK struct {
	Payload models.Int64Array
}

// IsSuccess returns true when this job next n req o k response has a 2xx status code
func (o *JobNextNReqOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this job next n req o k response has a 3xx status code
func (o *JobNextNReqOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job next n req o k response has a 4xx status code
func (o *JobNextNReqOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this job next n req o k response has a 5xx status code
func (o *JobNextNReqOK) IsServerError() bool {
	return false
}

// IsCode returns true when this job next n req o k response a status code equal to that given
func (o *JobNextNReqOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the job next n req o k response
func (o *JobNextNReqOK) Code() int {
	return 200
}

func (o *JobNextNReqOK) Error() string {
	return fmt.Sprintf("[GET /job/next][%d] jobNextNReqOK  %+v", 200, o.Payload)
}

func (o *JobNextNReqOK) String() string {
	return fmt.Sprintf("[GET /job/next][%d] jobNextNReqOK  %+v", 200, o.Payload)
}

func (o *JobNextNReqOK) GetPayload() models.Int64Array {
	return o.Payload
}

func (o *JobNextNReqOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJobNextNReqServiceUnavailable creates a JobNextNReqServiceUnavailable with default headers values
func NewJobNextNReqServiceUnavailable() *JobNextNReqServiceUnavailable {
	return &JobNextNReqServiceUnavailable{}
}

/*
JobNextNReqServiceUnavailable describes a response with status code 503, with default header values.

apiError
*/
type JobNextNReqServiceUnavailable struct {
	Payload *models.APIError
}

// IsSuccess returns true when this job next n req service unavailable response has a 2xx status code
func (o *JobNextNReqServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this job next n req service unavailable response has a 3xx status code
func (o *JobNextNReqServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this job next n req service unavailable response has a 4xx status code
func (o *JobNextNReqServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this job next n req service unavailable response has a 5xx status code
func (o *JobNextNReqServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this job next n req service unavailable response a status code equal to that given
func (o *JobNextNReqServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the job next n req service unavailable response
func (o *JobNextNReqServiceUnavailable) Code() int {
	return 503
}

func (o *JobNextNReqServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /job/next][%d] jobNextNReqServiceUnavailable  %+v", 503, o.Payload)
}

func (o *JobNextNReqServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /job/next][%d] jobNextNReqServiceUnavailable  %+v", 503, o.Payload)
}

func (o *JobNextNReqServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *JobNextNReqServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
