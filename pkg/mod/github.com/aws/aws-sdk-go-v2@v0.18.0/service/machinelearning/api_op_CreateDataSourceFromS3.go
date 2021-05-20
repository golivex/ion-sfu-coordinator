// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDataSourceFromS3Input struct {
	_ struct{} `type:"structure"`

	// The compute statistics for a DataSource. The statistics are generated from
	// the observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if
	// the DataSource needs to be used for MLModel training.
	ComputeStatistics *bool `type:"boolean"`

	// A user-supplied identifier that uniquely identifies the DataSource.
	//
	// DataSourceId is a required field
	DataSourceId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the DataSource.
	DataSourceName *string `type:"string"`

	// The data specification of a DataSource:
	//
	//    * DataLocationS3 - The Amazon S3 location of the observation data.
	//
	//    * DataSchemaLocationS3 - The Amazon S3 location of the DataSchema.
	//
	//    * DataSchema - A JSON string representing the schema. This is not required
	//    if DataSchemaUri is specified.
	//
	//    * DataRearrangement - A JSON string that represents the splitting and
	//    rearrangement requirements for the Datasource. Sample - "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// DataSpec is a required field
	DataSpec *S3DataSpec `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateDataSourceFromS3Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceFromS3Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceFromS3Input"}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}
	if s.DataSourceId != nil && len(*s.DataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSourceId", 1))
	}

	if s.DataSpec == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSpec"))
	}
	if s.DataSpec != nil {
		if err := s.DataSpec.Validate(); err != nil {
			invalidParams.AddNested("DataSpec", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDataSourceFromS3 operation, and is an acknowledgement
// that Amazon ML received the request.
//
// The CreateDataSourceFromS3 operation is asynchronous. You can poll for updates
// by using the GetBatchPrediction operation and checking the Status parameter.
type CreateDataSourceFromS3Output struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the DataSource. This value should
	// be identical to the value of the DataSourceID in the request.
	DataSourceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDataSourceFromS3Output) String() string {
	return awsutil.Prettify(s)
}

const opCreateDataSourceFromS3 = "CreateDataSourceFromS3"

// CreateDataSourceFromS3Request returns a request value for making API operation for
// Amazon Machine Learning.
//
// Creates a DataSource object. A DataSource references data that can be used
// to perform CreateMLModel, CreateEvaluation, or CreateBatchPrediction operations.
//
// CreateDataSourceFromS3 is an asynchronous operation. In response to CreateDataSourceFromS3,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the DataSource
// status to PENDING. After the DataSource has been created and is ready for
// use, Amazon ML sets the Status parameter to COMPLETED. DataSource in the
// COMPLETED or PENDING state can be used to perform only CreateMLModel, CreateEvaluation
// or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter
// to FAILED and includes an error message in the Message attribute of the GetDataSource
// operation response.
//
// The observation data used in a DataSource should be ready to use; that is,
// it should have a consistent structure, and missing data values should be
// kept to a minimum. The observation data must reside in one or more .csv files
// in an Amazon Simple Storage Service (Amazon S3) location, along with a schema
// that describes the data items by name and type. The same schema must be used
// for all of the data files referenced by the DataSource.
//
// After the DataSource has been created, it's ready to use in evaluations and
// batch predictions. If you plan to use the DataSource to train an MLModel,
// the DataSource also needs a recipe. A recipe describes how each input variable
// will be used in training an MLModel. Will the variable be included or excluded
// from training? Will the variable be manipulated; for example, will it be
// combined with another variable or will it be split apart into word combinations?
// The recipe provides answers to these questions.
//
//    // Example sending a request using CreateDataSourceFromS3Request.
//    req := client.CreateDataSourceFromS3Request(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDataSourceFromS3Request(input *CreateDataSourceFromS3Input) CreateDataSourceFromS3Request {
	op := &aws.Operation{
		Name:       opCreateDataSourceFromS3,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDataSourceFromS3Input{}
	}

	req := c.newRequest(op, input, &CreateDataSourceFromS3Output{})
	return CreateDataSourceFromS3Request{Request: req, Input: input, Copy: c.CreateDataSourceFromS3Request}
}

// CreateDataSourceFromS3Request is the request type for the
// CreateDataSourceFromS3 API operation.
type CreateDataSourceFromS3Request struct {
	*aws.Request
	Input *CreateDataSourceFromS3Input
	Copy  func(*CreateDataSourceFromS3Input) CreateDataSourceFromS3Request
}

// Send marshals and sends the CreateDataSourceFromS3 API request.
func (r CreateDataSourceFromS3Request) Send(ctx context.Context) (*CreateDataSourceFromS3Response, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSourceFromS3Response{
		CreateDataSourceFromS3Output: r.Request.Data.(*CreateDataSourceFromS3Output),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSourceFromS3Response is the response type for the
// CreateDataSourceFromS3 API operation.
type CreateDataSourceFromS3Response struct {
	*CreateDataSourceFromS3Output

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSourceFromS3 request.
func (r *CreateDataSourceFromS3Response) SDKResponseMetdata() *aws.Response {
	return r.response
}
