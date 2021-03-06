package j_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJMachineDenyIDReader is a Reader for the PostRemoteAPIJMachineDenyID structure.
type PostRemoteAPIJMachineDenyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJMachineDenyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJMachineDenyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJMachineDenyIDOK creates a PostRemoteAPIJMachineDenyIDOK with default headers values
func NewPostRemoteAPIJMachineDenyIDOK() *PostRemoteAPIJMachineDenyIDOK {
	return &PostRemoteAPIJMachineDenyIDOK{}
}

/*PostRemoteAPIJMachineDenyIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJMachineDenyIDOK struct {
	Payload *models.JMachine
}

func (o *PostRemoteAPIJMachineDenyIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JMachine.deny/{id}][%d] postRemoteApiJMachineDenyIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJMachineDenyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
