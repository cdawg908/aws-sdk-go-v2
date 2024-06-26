// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns a list of S3 Access Grants instances. An S3 Access Grants instance
// serves as a logical grouping for your individual access grants. You can only
// have one S3 Access Grants instance per Region per account.
//
// Permissions You must have the s3:ListAccessGrantsInstances permission to use
// this operation.
func (c *Client) ListAccessGrantsInstances(ctx context.Context, params *ListAccessGrantsInstancesInput, optFns ...func(*Options)) (*ListAccessGrantsInstancesOutput, error) {
	if params == nil {
		params = &ListAccessGrantsInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessGrantsInstances", params, optFns, c.addOperationListAccessGrantsInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessGrantsInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessGrantsInstancesInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	MaxResults int32

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants Instances request in order to retrieve the next
	// page of results.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListAccessGrantsInstancesInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type ListAccessGrantsInstancesOutput struct {

	// A container for a list of S3 Access Grants instances.
	AccessGrantsInstancesList []types.ListAccessGrantsInstanceEntry

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants Instances request in order to retrieve the next
	// page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessGrantsInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListAccessGrantsInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListAccessGrantsInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessGrantsInstances"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListAccessGrantsInstancesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccessGrantsInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessGrantsInstances(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListAccessGrantsInstancesUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

// ListAccessGrantsInstancesPaginatorOptions is the paginator options for
// ListAccessGrantsInstances
type ListAccessGrantsInstancesPaginatorOptions struct {
	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessGrantsInstancesPaginator is a paginator for ListAccessGrantsInstances
type ListAccessGrantsInstancesPaginator struct {
	options   ListAccessGrantsInstancesPaginatorOptions
	client    ListAccessGrantsInstancesAPIClient
	params    *ListAccessGrantsInstancesInput
	nextToken *string
	firstPage bool
}

// NewListAccessGrantsInstancesPaginator returns a new
// ListAccessGrantsInstancesPaginator
func NewListAccessGrantsInstancesPaginator(client ListAccessGrantsInstancesAPIClient, params *ListAccessGrantsInstancesInput, optFns ...func(*ListAccessGrantsInstancesPaginatorOptions)) *ListAccessGrantsInstancesPaginator {
	if params == nil {
		params = &ListAccessGrantsInstancesInput{}
	}

	options := ListAccessGrantsInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessGrantsInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessGrantsInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessGrantsInstances page.
func (p *ListAccessGrantsInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessGrantsInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAccessGrantsInstances(ctx, &params, optFns...)
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

type endpointPrefix_opListAccessGrantsInstancesMiddleware struct {
}

func (*endpointPrefix_opListAccessGrantsInstancesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessGrantsInstancesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ListAccessGrantsInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListAccessGrantsInstancesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAccessGrantsInstancesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAccessGrantsInstancesAPIClient is a client that implements the
// ListAccessGrantsInstances operation.
type ListAccessGrantsInstancesAPIClient interface {
	ListAccessGrantsInstances(context.Context, *ListAccessGrantsInstancesInput, ...func(*Options)) (*ListAccessGrantsInstancesOutput, error)
}

var _ ListAccessGrantsInstancesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccessGrantsInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessGrantsInstances",
	}
}

func copyListAccessGrantsInstancesInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListAccessGrantsInstancesInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListAccessGrantsInstancesInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListAccessGrantsInstancesInput) copy() interface{} {
	v := *in
	return &v
}
func backFillListAccessGrantsInstancesAccountID(input interface{}, v string) error {
	in := input.(*ListAccessGrantsInstancesInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListAccessGrantsInstancesUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListAccessGrantsInstancesInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
