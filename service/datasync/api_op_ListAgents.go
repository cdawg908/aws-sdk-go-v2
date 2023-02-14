// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DataSync agents that belong to an Amazon Web Services account
// in the Amazon Web Services Region specified in the request. With pagination, you
// can reduce the number of agents returned in a response. If you get a truncated
// list of agents in a response, the response contains a marker that you can
// specify in your next request to fetch the next page of agents. ListAgents is
// eventually consistent. This means the result of running the operation might not
// reflect that you just created or deleted an agent. For example, if you create an
// agent with CreateAgent
// (https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateAgent.html) and
// then immediately run ListAgents, that agent might not show up in the list right
// away. In situations like this, you can always confirm whether an agent has been
// created (or deleted) by using DescribeAgent
// (https://docs.aws.amazon.com/datasync/latest/userguide/API_DescribeAgent.html).
func (c *Client) ListAgents(ctx context.Context, params *ListAgentsInput, optFns ...func(*Options)) (*ListAgentsOutput, error) {
	if params == nil {
		params = &ListAgentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgents", params, optFns, c.addOperationListAgentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListAgentsRequest
type ListAgentsInput struct {

	// Specifies the maximum number of DataSync agents to list in a response. By
	// default, a response shows a maximum of 100 agents.
	MaxResults *int32

	// Specifies an opaque string that indicates the position to begin the next list of
	// results in the response.
	NextToken *string

	noSmithyDocumentSerde
}

// ListAgentsResponse
type ListAgentsOutput struct {

	// A list of DataSync agents in your Amazon Web Services account in the Amazon Web
	// Services Region specified in the request. The list is ordered by the agents'
	// Amazon Resource Names (ARNs).
	Agents []types.AgentListEntry

	// The opaque string that indicates the position to begin the next list of results
	// in the response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAgents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAgents{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListAgentsAPIClient is a client that implements the ListAgents operation.
type ListAgentsAPIClient interface {
	ListAgents(context.Context, *ListAgentsInput, ...func(*Options)) (*ListAgentsOutput, error)
}

var _ ListAgentsAPIClient = (*Client)(nil)

// ListAgentsPaginatorOptions is the paginator options for ListAgents
type ListAgentsPaginatorOptions struct {
	// Specifies the maximum number of DataSync agents to list in a response. By
	// default, a response shows a maximum of 100 agents.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgentsPaginator is a paginator for ListAgents
type ListAgentsPaginator struct {
	options   ListAgentsPaginatorOptions
	client    ListAgentsAPIClient
	params    *ListAgentsInput
	nextToken *string
	firstPage bool
}

// NewListAgentsPaginator returns a new ListAgentsPaginator
func NewListAgentsPaginator(client ListAgentsAPIClient, params *ListAgentsInput, optFns ...func(*ListAgentsPaginatorOptions)) *ListAgentsPaginator {
	if params == nil {
		params = &ListAgentsInput{}
	}

	options := ListAgentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgents page.
func (p *ListAgentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListAgents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAgents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "ListAgents",
	}
}
