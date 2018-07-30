// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ewilde/go-form3/models"
)

// PatchAccountconfigurationsIDReader is a Reader for the PatchAccountconfigurationsID structure.
type PatchAccountconfigurationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAccountconfigurationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAccountconfigurationsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAccountconfigurationsIDOK creates a PatchAccountconfigurationsIDOK with default headers values
func NewPatchAccountconfigurationsIDOK() *PatchAccountconfigurationsIDOK {
	return &PatchAccountconfigurationsIDOK{}
}

/*PatchAccountconfigurationsIDOK handles this case with default header values.

Configuration updated
*/
type PatchAccountconfigurationsIDOK struct {
	Payload *models.AccountConfigurationDetailsResponse
}

func (o *PatchAccountconfigurationsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /accountconfigurations/{id}][%d] patchAccountconfigurationsIdOK  %+v", 200, o.Payload)
}

func (o *PatchAccountconfigurationsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountConfigurationDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
