// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMergeCommitInput struct {
	_ struct{} `type:"structure"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which returns a not-mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation
	// is successful.
	ConflictResolutionStrategy ConflictResolutionStrategyTypeEnum `locationName:"conflictResolutionStrategy" type:"string" enum:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit (for example, a branch name or a full commit ID).
	//
	// DestinationCommitSpecifier is a required field
	DestinationCommitSpecifier *string `locationName:"destinationCommitSpecifier" type:"string" required:"true"`

	// The name of the repository that contains the merge commit about which you
	// want to get information.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit (for example, a branch name or a full commit ID).
	//
	// SourceCommitSpecifier is a required field
	SourceCommitSpecifier *string `locationName:"sourceCommitSpecifier" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMergeCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMergeCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMergeCommitInput"}

	if s.DestinationCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCommitSpecifier"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if s.SourceCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceCommitSpecifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMergeCommitOutput struct {
	_ struct{} `type:"structure"`

	// The commit ID of the merge base.
	BaseCommitId *string `locationName:"baseCommitId" type:"string"`

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	DestinationCommitId *string `locationName:"destinationCommitId" type:"string"`

	// The commit ID for the merge commit created when the source branch was merged
	// into the destination branch. If the fast-forward merge strategy was used,
	// there is no merge commit.
	MergedCommitId *string `locationName:"mergedCommitId" type:"string"`

	// The commit ID of the source commit specifier that was used in the merge evaluation.
	SourceCommitId *string `locationName:"sourceCommitId" type:"string"`
}

// String returns the string representation
func (s GetMergeCommitOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMergeCommit = "GetMergeCommit"

// GetMergeCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a specified merge commit.
//
//    // Example sending a request using GetMergeCommitRequest.
//    req := client.GetMergeCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetMergeCommit
func (c *Client) GetMergeCommitRequest(input *GetMergeCommitInput) GetMergeCommitRequest {
	op := &aws.Operation{
		Name:       opGetMergeCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMergeCommitInput{}
	}

	req := c.newRequest(op, input, &GetMergeCommitOutput{})
	return GetMergeCommitRequest{Request: req, Input: input, Copy: c.GetMergeCommitRequest}
}

// GetMergeCommitRequest is the request type for the
// GetMergeCommit API operation.
type GetMergeCommitRequest struct {
	*aws.Request
	Input *GetMergeCommitInput
	Copy  func(*GetMergeCommitInput) GetMergeCommitRequest
}

// Send marshals and sends the GetMergeCommit API request.
func (r GetMergeCommitRequest) Send(ctx context.Context) (*GetMergeCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMergeCommitResponse{
		GetMergeCommitOutput: r.Request.Data.(*GetMergeCommitOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMergeCommitResponse is the response type for the
// GetMergeCommit API operation.
type GetMergeCommitResponse struct {
	*GetMergeCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMergeCommit request.
func (r *GetMergeCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
