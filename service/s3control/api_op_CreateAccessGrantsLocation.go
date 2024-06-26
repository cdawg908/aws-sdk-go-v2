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
	"time"
)

// The S3 data location that you would like to register in your S3 Access Grants
// instance. Your S3 data must be in the same Region as your S3 Access Grants
// instance. The location can be one of the following:
//
//   - The default S3 location s3://
//
//   - A bucket - S3://
//
//   - A bucket and prefix - S3:///
//
// When you register a location, you must include the IAM role that has permission
// to manage the S3 location that you are registering. Give S3 Access Grants
// permission to assume this role [using a policy]. S3 Access Grants assumes this role to manage
// access to the location and to vend temporary credentials to grantees or client
// applications.
//
// Permissions You must have the s3:CreateAccessGrantsLocation permission to use
// this operation.
//
// Additional Permissions You must also have the following permission for the
// specified IAM role: iam:PassRole
//
// [using a policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-location.html
func (c *Client) CreateAccessGrantsLocation(ctx context.Context, params *CreateAccessGrantsLocationInput, optFns ...func(*Options)) (*CreateAccessGrantsLocationOutput, error) {
	if params == nil {
		params = &CreateAccessGrantsLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessGrantsLocation", params, optFns, c.addOperationCreateAccessGrantsLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessGrantsLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessGrantsLocationInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of the IAM role for the registered location. S3
	// Access Grants assumes this role to manage access to the registered location.
	//
	// This member is required.
	IAMRoleArn *string

	// The S3 path to the location that you are registering. The location scope can be
	// the default S3 location s3:// , the S3 path to a bucket s3:// , or the S3 path
	// to a bucket and prefix s3:/// . A prefix in S3 is a string of characters at the
	// beginning of an object key name used to organize the objects that you store in
	// your S3 buckets. For example, object key names that start with the engineering/
	// prefix or object key names that start with the marketing/campaigns/ prefix.
	//
	// This member is required.
	LocationScope *string

	// The Amazon Web Services resource tags that you are adding to the S3 Access
	// Grants location. Each tag is a label consisting of a user-defined key and value.
	// Tags can help you manage, identify, organize, search for, and filter resources.
	Tags []types.Tag

	noSmithyDocumentSerde
}

func (in *CreateAccessGrantsLocationInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type CreateAccessGrantsLocationOutput struct {

	// The Amazon Resource Name (ARN) of the location you are registering.
	AccessGrantsLocationArn *string

	// The ID of the registered location to which you are granting access. S3 Access
	// Grants assigns this ID when you register the location. S3 Access Grants assigns
	// the ID default to the default location s3:// and assigns an auto-generated ID
	// to other locations that you register.
	AccessGrantsLocationId *string

	// The date and time when you registered the location.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the IAM role for the registered location. S3
	// Access Grants assumes this role to manage access to the registered location.
	IAMRoleArn *string

	// The S3 URI path to the location that you are registering. The location scope
	// can be the default S3 location s3:// , the S3 path to a bucket, or the S3 path
	// to a bucket and prefix. A prefix in S3 is a string of characters at the
	// beginning of an object key name used to organize the objects that you store in
	// your S3 buckets. For example, object key names that start with the engineering/
	// prefix or object key names that start with the marketing/campaigns/ prefix.
	LocationScope *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessGrantsLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateAccessGrantsLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateAccessGrantsLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessGrantsLocation"); err != nil {
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
	if err = addEndpointPrefix_opCreateAccessGrantsLocationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAccessGrantsLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessGrantsLocation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateAccessGrantsLocationUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opCreateAccessGrantsLocationMiddleware struct {
}

func (*endpointPrefix_opCreateAccessGrantsLocationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateAccessGrantsLocationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*CreateAccessGrantsLocationInput)
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
func addEndpointPrefix_opCreateAccessGrantsLocationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateAccessGrantsLocationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateAccessGrantsLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessGrantsLocation",
	}
}

func copyCreateAccessGrantsLocationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateAccessGrantsLocationInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateAccessGrantsLocationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *CreateAccessGrantsLocationInput) copy() interface{} {
	v := *in
	return &v
}
func backFillCreateAccessGrantsLocationAccountID(input interface{}, v string) error {
	in := input.(*CreateAccessGrantsLocationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addCreateAccessGrantsLocationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyCreateAccessGrantsLocationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
