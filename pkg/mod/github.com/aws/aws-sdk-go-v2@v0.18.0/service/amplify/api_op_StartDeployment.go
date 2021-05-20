// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for start a deployment.
type StartDeploymentInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch, for the Job.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// The job id for this deployment, generated by create deployment request.
	JobId *string `locationName:"jobId" type:"string"`

	// The sourceUrl for this deployment, used when calling start deployment without
	// create deployment. SourceUrl can be any HTTP GET url that is public accessible
	// and downloads a single zip.
	SourceUrl *string `locationName:"sourceUrl" type:"string"`
}

// String returns the string representation
func (s StartDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDeploymentInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceUrl != nil {
		v := *s.SourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for start a deployment.
type StartDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// Summary for the Job.
	//
	// JobSummary is a required field
	JobSummary *JobSummary `locationName:"jobSummary" type:"structure" required:"true"`
}

// String returns the string representation
func (s StartDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobSummary != nil {
		v := s.JobSummary

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobSummary", v, metadata)
	}
	return nil
}

const opStartDeployment = "StartDeployment"

// StartDeploymentRequest returns a request value for making API operation for
// AWS Amplify.
//
// Start a deployment for manual deploy apps. (Apps are not connected to repository)
//
//    // Example sending a request using StartDeploymentRequest.
//    req := client.StartDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/StartDeployment
func (c *Client) StartDeploymentRequest(input *StartDeploymentInput) StartDeploymentRequest {
	op := &aws.Operation{
		Name:       opStartDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/branches/{branchName}/deployments/start",
	}

	if input == nil {
		input = &StartDeploymentInput{}
	}

	req := c.newRequest(op, input, &StartDeploymentOutput{})
	return StartDeploymentRequest{Request: req, Input: input, Copy: c.StartDeploymentRequest}
}

// StartDeploymentRequest is the request type for the
// StartDeployment API operation.
type StartDeploymentRequest struct {
	*aws.Request
	Input *StartDeploymentInput
	Copy  func(*StartDeploymentInput) StartDeploymentRequest
}

// Send marshals and sends the StartDeployment API request.
func (r StartDeploymentRequest) Send(ctx context.Context) (*StartDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDeploymentResponse{
		StartDeploymentOutput: r.Request.Data.(*StartDeploymentOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDeploymentResponse is the response type for the
// StartDeployment API operation.
type StartDeploymentResponse struct {
	*StartDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDeployment request.
func (r *StartDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
