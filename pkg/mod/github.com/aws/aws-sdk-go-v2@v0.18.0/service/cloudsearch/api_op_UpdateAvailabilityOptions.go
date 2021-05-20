// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the UpdateAvailabilityOptions operation.
// Specifies the name of the domain you want to update and the Multi-AZ availability
// option.
type UpdateAvailabilityOptionsInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// You expand an existing search domain to a second Availability Zone by setting
	// the Multi-AZ option to true. Similarly, you can turn off the Multi-AZ option
	// to downgrade the domain to a single Availability Zone by setting the Multi-AZ
	// option to false.
	//
	// MultiAZ is a required field
	MultiAZ *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s UpdateAvailabilityOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAvailabilityOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAvailabilityOptionsInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if s.MultiAZ == nil {
		invalidParams.Add(aws.NewErrParamRequired("MultiAZ"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a UpdateAvailabilityOptions request. Contains the status of
// the domain's availability options.
type UpdateAvailabilityOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The newly-configured availability options. Indicates whether Multi-AZ is
	// enabled for the domain.
	AvailabilityOptions *AvailabilityOptionsStatus `type:"structure"`
}

// String returns the string representation
func (s UpdateAvailabilityOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAvailabilityOptions = "UpdateAvailabilityOptions"

// UpdateAvailabilityOptionsRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Configures the availability options for a domain. Enabling the Multi-AZ option
// expands an Amazon CloudSearch domain to an additional Availability Zone in
// the same Region to increase fault tolerance in the event of a service disruption.
// Changes to the Multi-AZ option can take about half an hour to become active.
// For more information, see Configuring Availability Options (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-availability-options.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using UpdateAvailabilityOptionsRequest.
//    req := client.UpdateAvailabilityOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateAvailabilityOptionsRequest(input *UpdateAvailabilityOptionsInput) UpdateAvailabilityOptionsRequest {
	op := &aws.Operation{
		Name:       opUpdateAvailabilityOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAvailabilityOptionsInput{}
	}

	req := c.newRequest(op, input, &UpdateAvailabilityOptionsOutput{})
	return UpdateAvailabilityOptionsRequest{Request: req, Input: input, Copy: c.UpdateAvailabilityOptionsRequest}
}

// UpdateAvailabilityOptionsRequest is the request type for the
// UpdateAvailabilityOptions API operation.
type UpdateAvailabilityOptionsRequest struct {
	*aws.Request
	Input *UpdateAvailabilityOptionsInput
	Copy  func(*UpdateAvailabilityOptionsInput) UpdateAvailabilityOptionsRequest
}

// Send marshals and sends the UpdateAvailabilityOptions API request.
func (r UpdateAvailabilityOptionsRequest) Send(ctx context.Context) (*UpdateAvailabilityOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAvailabilityOptionsResponse{
		UpdateAvailabilityOptionsOutput: r.Request.Data.(*UpdateAvailabilityOptionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAvailabilityOptionsResponse is the response type for the
// UpdateAvailabilityOptions API operation.
type UpdateAvailabilityOptionsResponse struct {
	*UpdateAvailabilityOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAvailabilityOptions request.
func (r *UpdateAvailabilityOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
