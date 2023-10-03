// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an existing workload.
func (c *Client) UpdateWorkload(ctx context.Context, params *UpdateWorkloadInput, optFns ...func(*Options)) (*UpdateWorkloadOutput, error) {
	if params == nil {
		params = &UpdateWorkloadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkload", params, optFns, c.addOperationUpdateWorkloadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to update a workload.
type UpdateWorkloadInput struct {

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The list of Amazon Web Services account IDs associated with the workload.
	AccountIds []string

	// List of AppRegistry application ARNs to associate to the workload.
	Applications []string

	// The URL of the architectural design for the workload.
	ArchitecturalDesign *string

	// The list of Amazon Web Services Regions associated with the workload, for
	// example, us-east-2 , or ca-central-1 .
	AwsRegions []string

	// The description for the workload.
	Description *string

	// Well-Architected discovery configuration settings to associate to the workload.
	DiscoveryConfig *types.WorkloadDiscoveryConfig

	// The environment for the workload.
	Environment types.WorkloadEnvironment

	// The improvement status for a workload.
	ImprovementStatus types.WorkloadImprovementStatus

	// The industry for the workload.
	Industry *string

	// The industry type for the workload. If specified, must be one of the following:
	//   - Agriculture
	//   - Automobile
	//   - Defense
	//   - Design and Engineering
	//   - Digital Advertising
	//   - Education
	//   - Environmental Protection
	//   - Financial Services
	//   - Gaming
	//   - General Public Services
	//   - Healthcare
	//   - Hospitality
	//   - InfoTech
	//   - Justice and Public Safety
	//   - Life Sciences
	//   - Manufacturing
	//   - Media & Entertainment
	//   - Mining & Resources
	//   - Oil & Gas
	//   - Power & Utilities
	//   - Professional Services
	//   - Real Estate & Construction
	//   - Retail & Wholesale
	//   - Social Protection
	//   - Telecommunications
	//   - Travel, Transportation & Logistics
	//   - Other
	IndustryType *string

	// Flag indicating whether the workload owner has acknowledged that the Review
	// owner field is required. If a Review owner is not added to the workload within
	// 60 days of acknowledgement, access to the workload is restricted until an owner
	// is added.
	IsReviewOwnerUpdateAcknowledged bool

	// The list of non-Amazon Web Services Regions associated with the workload.
	NonAwsRegions []string

	// The notes associated with the workload. For a review template, these are the
	// notes that will be associated with the workload when the template is applied.
	Notes *string

	// The priorities of the pillars, which are used to order items in the improvement
	// plan. Each pillar is represented by its PillarReviewSummary$PillarId .
	PillarPriorities []string

	// The review owner of the workload. The name, email address, or identifier for
	// the primary group or individual that owns the workload review process.
	ReviewOwner *string

	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	WorkloadName *string

	noSmithyDocumentSerde
}

// Output of an update workload call.
type UpdateWorkloadOutput struct {

	// A workload return object.
	Workload *types.Workload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkload{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateWorkloadResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateWorkloadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkload(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateWorkload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpdateWorkload",
	}
}

type opUpdateWorkloadResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateWorkloadResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateWorkloadResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "wellarchitected"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "wellarchitected"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("wellarchitected")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateWorkloadResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateWorkloadResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
