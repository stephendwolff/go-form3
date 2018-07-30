// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ewilde/go-form3/models"
)

// GetStarlingReader is a Reader for the GetStarling structure.
type GetStarlingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStarlingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStarlingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarlingOK creates a GetStarlingOK with default headers values
func NewGetStarlingOK() *GetStarlingOK {
	return &GetStarlingOK{}
}

/*GetStarlingOK handles this case with default header values.

List of associations
*/
type GetStarlingOK struct {
	Payload *models.AssociationDetailsListResponse
}

func (o *GetStarlingOK) Error() string {
	return fmt.Sprintf("[GET /starling][%d] getStarlingOK  %+v", 200, o.Payload)
}

func (o *GetStarlingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
