// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEventTrackersInput struct {
	_ struct{} `type:"structure"`

	// The ARN of a dataset group used to filter the response.
	DatasetGroupArn *string `locationName:"datasetGroupArn" type:"string"`

	// The maximum number of event trackers to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListEventTrackers for getting
	// the next set of event trackers (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListEventTrackersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventTrackersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventTrackersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEventTrackersOutput struct {
	_ struct{} `type:"structure"`

	// A list of event trackers.
	EventTrackers []EventTrackerSummary `locationName:"eventTrackers" type:"list"`

	// A token for getting the next set of event trackers (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListEventTrackersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEventTrackers = "ListEventTrackers"

// ListEventTrackersRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Returns the list of event trackers associated with the account. The response
// provides the properties for each event tracker, including the Amazon Resource
// Name (ARN) and tracking ID. For more information on event trackers, see CreateEventTracker.
//
//    // Example sending a request using ListEventTrackersRequest.
//    req := client.ListEventTrackersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListEventTrackers
func (c *Client) ListEventTrackersRequest(input *ListEventTrackersInput) ListEventTrackersRequest {
	op := &aws.Operation{
		Name:       opListEventTrackers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListEventTrackersInput{}
	}

	req := c.newRequest(op, input, &ListEventTrackersOutput{})
	return ListEventTrackersRequest{Request: req, Input: input, Copy: c.ListEventTrackersRequest}
}

// ListEventTrackersRequest is the request type for the
// ListEventTrackers API operation.
type ListEventTrackersRequest struct {
	*aws.Request
	Input *ListEventTrackersInput
	Copy  func(*ListEventTrackersInput) ListEventTrackersRequest
}

// Send marshals and sends the ListEventTrackers API request.
func (r ListEventTrackersRequest) Send(ctx context.Context) (*ListEventTrackersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventTrackersResponse{
		ListEventTrackersOutput: r.Request.Data.(*ListEventTrackersOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEventTrackersRequestPaginator returns a paginator for ListEventTrackers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEventTrackersRequest(input)
//   p := personalize.NewListEventTrackersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEventTrackersPaginator(req ListEventTrackersRequest) ListEventTrackersPaginator {
	return ListEventTrackersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEventTrackersInput
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

// ListEventTrackersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEventTrackersPaginator struct {
	aws.Pager
}

func (p *ListEventTrackersPaginator) CurrentPage() *ListEventTrackersOutput {
	return p.Pager.CurrentPage().(*ListEventTrackersOutput)
}

// ListEventTrackersResponse is the response type for the
// ListEventTrackers API operation.
type ListEventTrackersResponse struct {
	*ListEventTrackersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventTrackers request.
func (r *ListEventTrackersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
