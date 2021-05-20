// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeNetworkAclsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * association.association-id - The ID of an association ID for the ACL.
	//
	//    * association.network-acl-id - The ID of the network ACL involved in the
	//    association.
	//
	//    * association.subnet-id - The ID of the subnet involved in the association.
	//
	//    * default - Indicates whether the ACL is the default network ACL for the
	//    VPC.
	//
	//    * entry.cidr - The IPv4 CIDR range specified in the entry.
	//
	//    * entry.icmp.code - The ICMP code specified in the entry, if any.
	//
	//    * entry.icmp.type - The ICMP type specified in the entry, if any.
	//
	//    * entry.ipv6-cidr - The IPv6 CIDR range specified in the entry.
	//
	//    * entry.port-range.from - The start of the port range specified in the
	//    entry.
	//
	//    * entry.port-range.to - The end of the port range specified in the entry.
	//
	//    * entry.protocol - The protocol specified in the entry (tcp | udp | icmp
	//    or a protocol number).
	//
	//    * entry.rule-action - Allows or denies the matching traffic (allow | deny).
	//
	//    * entry.rule-number - The number of an entry (in other words, rule) in
	//    the set of ACL entries.
	//
	//    * network-acl-id - The ID of the network ACL.
	//
	//    * owner-id - The ID of the AWS account that owns the network ACL.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-id - The ID of the VPC for the network ACL.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// One or more network ACL IDs.
	//
	// Default: Describes all your network ACLs.
	NetworkAclIds []string `locationName:"NetworkAclId" locationNameList:"item" type:"list"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkAclsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkAclsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNetworkAclsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeNetworkAclsOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more network ACLs.
	NetworkAcls []NetworkAcl `locationName:"networkAclSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeNetworkAclsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNetworkAcls = "DescribeNetworkAcls"

// DescribeNetworkAclsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your network ACLs.
//
// For more information, see Network ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using DescribeNetworkAclsRequest.
//    req := client.DescribeNetworkAclsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkAcls
func (c *Client) DescribeNetworkAclsRequest(input *DescribeNetworkAclsInput) DescribeNetworkAclsRequest {
	op := &aws.Operation{
		Name:       opDescribeNetworkAcls,
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
		input = &DescribeNetworkAclsInput{}
	}

	req := c.newRequest(op, input, &DescribeNetworkAclsOutput{})
	return DescribeNetworkAclsRequest{Request: req, Input: input, Copy: c.DescribeNetworkAclsRequest}
}

// DescribeNetworkAclsRequest is the request type for the
// DescribeNetworkAcls API operation.
type DescribeNetworkAclsRequest struct {
	*aws.Request
	Input *DescribeNetworkAclsInput
	Copy  func(*DescribeNetworkAclsInput) DescribeNetworkAclsRequest
}

// Send marshals and sends the DescribeNetworkAcls API request.
func (r DescribeNetworkAclsRequest) Send(ctx context.Context) (*DescribeNetworkAclsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNetworkAclsResponse{
		DescribeNetworkAclsOutput: r.Request.Data.(*DescribeNetworkAclsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNetworkAclsRequestPaginator returns a paginator for DescribeNetworkAcls.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNetworkAclsRequest(input)
//   p := ec2.NewDescribeNetworkAclsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNetworkAclsPaginator(req DescribeNetworkAclsRequest) DescribeNetworkAclsPaginator {
	return DescribeNetworkAclsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeNetworkAclsInput
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

// DescribeNetworkAclsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNetworkAclsPaginator struct {
	aws.Pager
}

func (p *DescribeNetworkAclsPaginator) CurrentPage() *DescribeNetworkAclsOutput {
	return p.Pager.CurrentPage().(*DescribeNetworkAclsOutput)
}

// DescribeNetworkAclsResponse is the response type for the
// DescribeNetworkAcls API operation.
type DescribeNetworkAclsResponse struct {
	*DescribeNetworkAclsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNetworkAcls request.
func (r *DescribeNetworkAclsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
