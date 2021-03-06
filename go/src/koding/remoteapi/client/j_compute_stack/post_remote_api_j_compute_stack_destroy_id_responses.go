package j_compute_stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJComputeStackDestroyIDReader is a Reader for the PostRemoteAPIJComputeStackDestroyID structure.
type PostRemoteAPIJComputeStackDestroyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJComputeStackDestroyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJComputeStackDestroyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJComputeStackDestroyIDOK creates a PostRemoteAPIJComputeStackDestroyIDOK with default headers values
func NewPostRemoteAPIJComputeStackDestroyIDOK() *PostRemoteAPIJComputeStackDestroyIDOK {
	return &PostRemoteAPIJComputeStackDestroyIDOK{}
}

/*PostRemoteAPIJComputeStackDestroyIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJComputeStackDestroyIDOK struct {
	Payload *models.JComputeStack
}

func (o *PostRemoteAPIJComputeStackDestroyIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JComputeStack.destroy/{id}][%d] postRemoteApiJComputeStackDestroyIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJComputeStackDestroyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JComputeStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
