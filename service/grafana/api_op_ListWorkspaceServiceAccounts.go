// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of service accounts for a workspace.
//
// Service accounts are only available for workspaces that are compatible with
// Grafana version 9 and above.
func (c *Client) ListWorkspaceServiceAccounts(ctx context.Context, params *ListWorkspaceServiceAccountsInput, optFns ...func(*Options)) (*ListWorkspaceServiceAccountsOutput, error) {
	if params == nil {
		params = &ListWorkspaceServiceAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkspaceServiceAccounts", params, optFns, c.addOperationListWorkspaceServiceAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkspaceServiceAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkspaceServiceAccountsInput struct {

	// The workspace for which to list service accounts.
	//
	// This member is required.
	WorkspaceId *string

	// The maximum number of service accounts to include in the results.
	MaxResults *int32

	// The token for the next set of service accounts to return. (You receive this
	// token from a previous ListWorkspaceServiceAccounts operation.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkspaceServiceAccountsOutput struct {

	// An array of structures containing information about the service accounts.
	//
	// This member is required.
	ServiceAccounts []types.ServiceAccountSummary

	// The workspace to which the service accounts are associated.
	//
	// This member is required.
	WorkspaceId *string

	// The token to use when requesting the next set of service accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkspaceServiceAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkspaceServiceAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkspaceServiceAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkspaceServiceAccounts"); err != nil {
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
	if err = addOpListWorkspaceServiceAccountsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkspaceServiceAccounts(options.Region), middleware.Before); err != nil {
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

// ListWorkspaceServiceAccountsAPIClient is a client that implements the
// ListWorkspaceServiceAccounts operation.
type ListWorkspaceServiceAccountsAPIClient interface {
	ListWorkspaceServiceAccounts(context.Context, *ListWorkspaceServiceAccountsInput, ...func(*Options)) (*ListWorkspaceServiceAccountsOutput, error)
}

var _ ListWorkspaceServiceAccountsAPIClient = (*Client)(nil)

// ListWorkspaceServiceAccountsPaginatorOptions is the paginator options for
// ListWorkspaceServiceAccounts
type ListWorkspaceServiceAccountsPaginatorOptions struct {
	// The maximum number of service accounts to include in the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkspaceServiceAccountsPaginator is a paginator for
// ListWorkspaceServiceAccounts
type ListWorkspaceServiceAccountsPaginator struct {
	options   ListWorkspaceServiceAccountsPaginatorOptions
	client    ListWorkspaceServiceAccountsAPIClient
	params    *ListWorkspaceServiceAccountsInput
	nextToken *string
	firstPage bool
}

// NewListWorkspaceServiceAccountsPaginator returns a new
// ListWorkspaceServiceAccountsPaginator
func NewListWorkspaceServiceAccountsPaginator(client ListWorkspaceServiceAccountsAPIClient, params *ListWorkspaceServiceAccountsInput, optFns ...func(*ListWorkspaceServiceAccountsPaginatorOptions)) *ListWorkspaceServiceAccountsPaginator {
	if params == nil {
		params = &ListWorkspaceServiceAccountsInput{}
	}

	options := ListWorkspaceServiceAccountsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkspaceServiceAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkspaceServiceAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkspaceServiceAccounts page.
func (p *ListWorkspaceServiceAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkspaceServiceAccountsOutput, error) {
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

	result, err := p.client.ListWorkspaceServiceAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkspaceServiceAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkspaceServiceAccounts",
	}
}
