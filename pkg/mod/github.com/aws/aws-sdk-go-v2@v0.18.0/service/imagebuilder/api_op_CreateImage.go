// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateImageInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of the distribution configuration that defines
	// and configures the outputs of your pipeline.
	DistributionConfigurationArn *string `locationName:"distributionConfigurationArn" type:"string"`

	// The Amazon Resource Name (ARN) of the image recipe that defines how images
	// are configured, tested and assessed.
	//
	// ImageRecipeArn is a required field
	ImageRecipeArn *string `locationName:"imageRecipeArn" type:"string" required:"true"`

	// The image tests configuration of the image.
	ImageTestsConfiguration *ImageTestsConfiguration `locationName:"imageTestsConfiguration" type:"structure"`

	// The Amazon Resource Name (ARN) of the infrastructure configuration that defines
	// the environment in which your image will be built and tested.
	//
	// InfrastructureConfigurationArn is a required field
	InfrastructureConfigurationArn *string `locationName:"infrastructureConfigurationArn" type:"string" required:"true"`

	// The tags of the image.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateImageInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.ImageRecipeArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageRecipeArn"))
	}

	if s.InfrastructureConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InfrastructureConfigurationArn"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.ImageTestsConfiguration != nil {
		if err := s.ImageTestsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ImageTestsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateImageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DistributionConfigurationArn != nil {
		v := *s.DistributionConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageRecipeArn != nil {
		v := *s.ImageRecipeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageRecipeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageTestsConfiguration != nil {
		v := s.ImageTestsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "imageTestsConfiguration", v, metadata)
	}
	if s.InfrastructureConfigurationArn != nil {
		v := *s.InfrastructureConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "infrastructureConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateImageOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the image that was created by this request.
	ImageBuildVersionArn *string `locationName:"imageBuildVersionArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateImageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateImageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateImage = "CreateImage"

// CreateImageRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Creates a new image. This request will create a new image along with all
// of the configured output resources defined in the distribution configuration.
//
//    // Example sending a request using CreateImageRequest.
//    req := client.CreateImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/CreateImage
func (c *Client) CreateImageRequest(input *CreateImageInput) CreateImageRequest {
	op := &aws.Operation{
		Name:       opCreateImage,
		HTTPMethod: "PUT",
		HTTPPath:   "/CreateImage",
	}

	if input == nil {
		input = &CreateImageInput{}
	}

	req := c.newRequest(op, input, &CreateImageOutput{})
	return CreateImageRequest{Request: req, Input: input, Copy: c.CreateImageRequest}
}

// CreateImageRequest is the request type for the
// CreateImage API operation.
type CreateImageRequest struct {
	*aws.Request
	Input *CreateImageInput
	Copy  func(*CreateImageInput) CreateImageRequest
}

// Send marshals and sends the CreateImage API request.
func (r CreateImageRequest) Send(ctx context.Context) (*CreateImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateImageResponse{
		CreateImageOutput: r.Request.Data.(*CreateImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateImageResponse is the response type for the
// CreateImage API operation.
type CreateImageResponse struct {
	*CreateImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateImage request.
func (r *CreateImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
