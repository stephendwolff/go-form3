// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/models"
)

// PatchUsersUserIDReader is a Reader for the PatchUsersUserID structure.
type PatchUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchUsersUserIDOK creates a PatchUsersUserIDOK with default headers values
func NewPatchUsersUserIDOK() *PatchUsersUserIDOK {
	return &PatchUsersUserIDOK{}
}

/*PatchUsersUserIDOK handles this case with default header values.

User details
*/
type PatchUsersUserIDOK struct {
	Payload *models.UserDetailsResponse
}

func (o *PatchUsersUserIDOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_id}][%d] patchUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *PatchUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
