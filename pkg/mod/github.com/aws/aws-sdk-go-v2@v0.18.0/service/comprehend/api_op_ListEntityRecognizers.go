// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEntityRecognizersInput struct {
	_ struct{} `type:"structure"`

	// Filters the list of entities returned. You can filter on Status, SubmitTimeBefore,
	// or SubmitTimeAfter. You can only set one filter at a time.
	Filter *EntityRecognizerFilter `type:"structure"`

	// The maximum number of results to return on each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntityRecognizersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntityRecognizersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntityRecognizersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEntityRecognizersOutput struct {
	_ struct{} `type:"structure"`

	// The list of properties of an entity recognizer.
	EntityRecognizerPropertiesList []EntityRecognizerProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntityRecognizersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEntityRecognizers = "ListEntityRecognizers"

// ListEntityRecognizersRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of the properties of all entity recognizers that you created,
// including recognizers currently in training. Allows you to filter the list
// of recognizers based on criteria such as status and submission time. This
// call returns up to 500 entity recognizers in the list, with a default number
// of 100 recognizers in the list.
//
// The results of this list are not in any particular order. Please get the
// list and sort locally if needed.
//
//    // Example sending a request using ListEntityRecognizersRequest.
//    req := client.ListEntityRecognizersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListEntityRecognizers
func (c *Client) ListEntityRecognizersRequest(input *ListEntityRecognizersInput) ListEntityRecognizersRequest {
	op := &aws.Operation{
		Name:       opListEntityRecognizers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEntityRecognizersInput{}
	}

	req := c.newRequest(op, input, &ListEntityRecognizersOutput{})
	return ListEntityRecognizersRequest{Request: req, Input: input, Copy: c.ListEntityRecognizersRequest}
}

// ListEntityRecognizersRequest is the request type for the
// ListEntityRecognizers API operation.
type ListEntityRecognizersRequest struct {
	*aws.Request
	Input *ListEntityRecognizersInput
	Copy  func(*ListEntityRecognizersInput) ListEntityRecognizersRequest
}

// Send marshals and sends the ListEntityRecognizers API request.
func (r ListEntityRecognizersRequest) Send(ctx context.Context) (*ListEntityRecognizersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEntityRecognizersResponse{
		ListEntityRecognizersOutput: r.Request.Data.(*ListEntityRecognizersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEntityRecognizersRequestPaginator returns a paginator for ListEntityRecognizers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEntityRecognizersRequest(input)
//   p := comprehend.NewListEntityRecognizersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEntityRecognizersPaginator(req ListEntityRecognizersRequest) ListEntityRecognizersPaginator {
	return ListEntityRecognizersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEntityRecognizersInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListEntityRecognizersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEntityRecognizersPaginator struct {
	aws.Pager
}

func (p *ListEntityRecognizersPaginator) CurrentPage() *ListEntityRecognizersOutput {
	return p.Pager.CurrentPage().(*ListEntityRecognizersOutput)
}

// ListEntityRecognizersResponse is the response type for the
// ListEntityRecognizers API operation.
type ListEntityRecognizersResponse struct {
	*ListEntityRecognizersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEntityRecognizers request.
func (r *ListEntityRecognizersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
