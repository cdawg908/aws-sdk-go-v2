// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all MLflow Tracking Servers.
func (c *Client) ListMlflowTrackingServers(ctx context.Context, params *ListMlflowTrackingServersInput, optFns ...func(*Options)) (*ListMlflowTrackingServersOutput, error) {
	if params == nil {
		params = &ListMlflowTrackingServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMlflowTrackingServers", params, optFns, c.addOperationListMlflowTrackingServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMlflowTrackingServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMlflowTrackingServersInput struct {

	// Use the CreatedAfter filter to only list tracking servers created after a
	// specific date and time. Listed tracking servers are shown with a date and time
	// such as "2024-03-16T01:46:56+00:00" . The CreatedAfter parameter takes in a
	// Unix timestamp. To convert a date and time into a Unix timestamp, see [EpochConverter].
	//
	// [EpochConverter]: https://www.epochconverter.com/
	CreatedAfter *time.Time

	// Use the CreatedBefore filter to only list tracking servers created before a
	// specific date and time. Listed tracking servers are shown with a date and time
	// such as "2024-03-16T01:46:56+00:00" . The CreatedBefore parameter takes in a
	// Unix timestamp. To convert a date and time into a Unix timestamp, see [EpochConverter].
	//
	// [EpochConverter]: https://www.epochconverter.com/
	CreatedBefore *time.Time

	// The maximum number of tracking servers to list.
	MaxResults *int32

	// Filter for tracking servers using the specified MLflow version.
	MlflowVersion *string

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// Filter for trackings servers sorting by name, creation time, or creation status.
	SortBy types.SortTrackingServerBy

	// Change the order of the listed tracking servers. By default, tracking servers
	// are listed in Descending order by creation time. To change the list order, you
	// can specify SortOrder to be Ascending .
	SortOrder types.SortOrder

	// Filter for tracking servers with a specified creation status.
	TrackingServerStatus types.TrackingServerStatus

	noSmithyDocumentSerde
}

type ListMlflowTrackingServersOutput struct {

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// A list of tracking servers according to chosen filters.
	TrackingServerSummaries []types.TrackingServerSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMlflowTrackingServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMlflowTrackingServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMlflowTrackingServers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMlflowTrackingServers"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMlflowTrackingServers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListMlflowTrackingServersPaginatorOptions is the paginator options for
// ListMlflowTrackingServers
type ListMlflowTrackingServersPaginatorOptions struct {
	// The maximum number of tracking servers to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMlflowTrackingServersPaginator is a paginator for ListMlflowTrackingServers
type ListMlflowTrackingServersPaginator struct {
	options   ListMlflowTrackingServersPaginatorOptions
	client    ListMlflowTrackingServersAPIClient
	params    *ListMlflowTrackingServersInput
	nextToken *string
	firstPage bool
}

// NewListMlflowTrackingServersPaginator returns a new
// ListMlflowTrackingServersPaginator
func NewListMlflowTrackingServersPaginator(client ListMlflowTrackingServersAPIClient, params *ListMlflowTrackingServersInput, optFns ...func(*ListMlflowTrackingServersPaginatorOptions)) *ListMlflowTrackingServersPaginator {
	if params == nil {
		params = &ListMlflowTrackingServersInput{}
	}

	options := ListMlflowTrackingServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMlflowTrackingServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMlflowTrackingServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMlflowTrackingServers page.
func (p *ListMlflowTrackingServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMlflowTrackingServersOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListMlflowTrackingServers(ctx, &params, optFns...)
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

// ListMlflowTrackingServersAPIClient is a client that implements the
// ListMlflowTrackingServers operation.
type ListMlflowTrackingServersAPIClient interface {
	ListMlflowTrackingServers(context.Context, *ListMlflowTrackingServersInput, ...func(*Options)) (*ListMlflowTrackingServersOutput, error)
}

var _ ListMlflowTrackingServersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMlflowTrackingServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMlflowTrackingServers",
	}
}
