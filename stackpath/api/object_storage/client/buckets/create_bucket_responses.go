// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateBucketReader is a Reader for the CreateBucket structure.
type CreateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBucketOK creates a CreateBucketOK with default headers values
func NewCreateBucketOK() *CreateBucketOK {
	return &CreateBucketOK{}
}

/*CreateBucketOK handles this case with default header values.

CreateBucketOK create bucket o k
*/
type CreateBucketOK struct {
	Payload *CreateBucketOKBody
}

func (o *CreateBucketOK) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] createBucketOK  %+v", 200, o.Payload)
}

func (o *CreateBucketOK) GetPayload() *CreateBucketOKBody {
	return o.Payload
}

func (o *CreateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateBucketOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketDefault creates a CreateBucketDefault with default headers values
func NewCreateBucketDefault(code int) *CreateBucketDefault {
	return &CreateBucketDefault{
		_statusCode: code,
	}
}

/*CreateBucketDefault handles this case with default header values.

Default error structure.
*/
type CreateBucketDefault struct {
	_statusCode int

	Payload *CreateBucketDefaultBody
}

// Code gets the status code for the create bucket default response
func (o *CreateBucketDefault) Code() int {
	return o._statusCode
}

func (o *CreateBucketDefault) Error() string {
	return fmt.Sprintf("[POST /storage/v1/stacks/{stack_id}/buckets][%d] CreateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBucketDefault) GetPayload() *CreateBucketDefaultBody {
	return o.Payload
}

func (o *CreateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateBucketDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateBucketBody create bucket body
swagger:model CreateBucketBody
*/
type CreateBucketBody struct {

	// The name of the bucket to be created
	Label string `json:"label,omitempty"`

	// The region where to create the bucket, defaults to us-east-1
	Region string `json:"region,omitempty"`
}

// Validate validates this create bucket body
func (o *CreateBucketBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBucketBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBucketBody) UnmarshalBinary(b []byte) error {
	var res CreateBucketBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateBucketDefaultBody create bucket default body
swagger:model CreateBucketDefaultBody
*/
type CreateBucketDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create bucket default body
func (o *CreateBucketDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBucketDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBucketDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateBucketDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateBucketOKBody The bucket created
swagger:model CreateBucketOKBody
*/
type CreateBucketOKBody struct {

	// bucket
	Bucket *CreateBucketOKBodyBucket `json:"bucket,omitempty"`
}

// Validate validates this create bucket o k body
func (o *CreateBucketOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateBucketOKBody) validateBucket(formats strfmt.Registry) error {

	if swag.IsZero(o.Bucket) { // not required
		return nil
	}

	if o.Bucket != nil {
		if err := o.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createBucketOK" + "." + "bucket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateBucketOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBucketOKBody) UnmarshalBinary(b []byte) error {
	var res CreateBucketOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateBucketOKBodyBucket create bucket o k body bucket
swagger:model CreateBucketOKBodyBucket
*/
type CreateBucketOKBodyBucket struct {

	// The date when the bucket was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// The URL used to access the bucket
	EndpointURL string `json:"endpointUrl,omitempty"`

	// The ID for the bucket
	ID string `json:"id,omitempty"`

	// The name of the bucket
	Label string `json:"label,omitempty"`

	// The region in which the bucket is created. Available regions are: us-east-1, us-west-1, eu-central-1
	Region string `json:"region,omitempty"`

	// The date when the bucket was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// - PRIVATE: The bucket is private and only accessibly with credentials
	//  - PUBLIC: The bucket is public and accessible over the internet
	// Enum: [PRIVATE PUBLIC]
	Visibility *string `json:"visibility,omitempty"`
}

// Validate validates this create bucket o k body bucket
func (o *CreateBucketOKBodyBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateBucketOKBodyBucket) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createBucketOK"+"."+"bucket"+"."+"createdAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateBucketOKBodyBucket) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createBucketOK"+"."+"bucket"+"."+"updatedAt", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var createBucketOKBodyBucketTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIVATE","PUBLIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBucketOKBodyBucketTypeVisibilityPropEnum = append(createBucketOKBodyBucketTypeVisibilityPropEnum, v)
	}
}

const (

	// CreateBucketOKBodyBucketVisibilityPRIVATE captures enum value "PRIVATE"
	CreateBucketOKBodyBucketVisibilityPRIVATE string = "PRIVATE"

	// CreateBucketOKBodyBucketVisibilityPUBLIC captures enum value "PUBLIC"
	CreateBucketOKBodyBucketVisibilityPUBLIC string = "PUBLIC"
)

// prop value enum
func (o *CreateBucketOKBodyBucket) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createBucketOKBodyBucketTypeVisibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *CreateBucketOKBodyBucket) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(o.Visibility) { // not required
		return nil
	}

	// value enum
	if err := o.validateVisibilityEnum("createBucketOK"+"."+"bucket"+"."+"visibility", "body", *o.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateBucketOKBodyBucket) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBucketOKBodyBucket) UnmarshalBinary(b []byte) error {
	var res CreateBucketOKBodyBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
