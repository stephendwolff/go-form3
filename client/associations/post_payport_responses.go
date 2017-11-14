// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ewilde/go-form3/models"
)

// PostPayportReader is a Reader for the PostPayport structure.
type PostPayportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPayportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPayportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPayportCreated creates a PostPayportCreated with default headers values
func NewPostPayportCreated() *PostPayportCreated {
	return &PostPayportCreated{}
}

/*PostPayportCreated handles this case with default header values.

creation response
*/
type PostPayportCreated struct {
	Payload *models.PayportAssociationCreationResponse
}

func (o *PostPayportCreated) Error() string {
	return fmt.Sprintf("[POST /payport][%d] postPayportCreated  %+v", 201, o.Payload)
}

func (o *PostPayportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayportAssociationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
