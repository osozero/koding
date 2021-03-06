package j_proxy_restriction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProxyRestrictionClearReader is a Reader for the PostRemoteAPIJProxyRestrictionClear structure.
type PostRemoteAPIJProxyRestrictionClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProxyRestrictionClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProxyRestrictionClearOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJProxyRestrictionClearUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProxyRestrictionClearOK creates a PostRemoteAPIJProxyRestrictionClearOK with default headers values
func NewPostRemoteAPIJProxyRestrictionClearOK() *PostRemoteAPIJProxyRestrictionClearOK {
	return &PostRemoteAPIJProxyRestrictionClearOK{}
}

/*PostRemoteAPIJProxyRestrictionClearOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJProxyRestrictionClearOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJProxyRestrictionClearOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyRestriction.clear][%d] postRemoteApiJProxyRestrictionClearOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProxyRestrictionClearOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJProxyRestrictionClearUnauthorized creates a PostRemoteAPIJProxyRestrictionClearUnauthorized with default headers values
func NewPostRemoteAPIJProxyRestrictionClearUnauthorized() *PostRemoteAPIJProxyRestrictionClearUnauthorized {
	return &PostRemoteAPIJProxyRestrictionClearUnauthorized{}
}

/*PostRemoteAPIJProxyRestrictionClearUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJProxyRestrictionClearUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJProxyRestrictionClearUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProxyRestriction.clear][%d] postRemoteApiJProxyRestrictionClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJProxyRestrictionClearUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
