// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Defines a ProfileObjectType.
//
// To add or remove tags on an existing ObjectType, see [TagResource]/[UntagResource] .
//
// [TagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html
func (c *Client) PutProfileObjectType(ctx context.Context, params *PutProfileObjectTypeInput, optFns ...func(*Options)) (*PutProfileObjectTypeOutput, error) {
	if params == nil {
		params = &PutProfileObjectTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutProfileObjectType", params, optFns, c.addOperationPutProfileObjectTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutProfileObjectTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProfileObjectTypeInput struct {

	// Description of the profile object type.
	//
	// This member is required.
	Description *string

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// Indicates whether a profile should be created when data is received if one
	// doesn’t exist for an object of this type. The default is FALSE . If the
	// AllowProfileCreation flag is set to FALSE , then the service tries to fetch a
	// standard profile and associate this object with the profile. If it is set to
	// TRUE , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation bool

	// The customer-provided key to encrypt the profile object that will be created in
	// this profile object type.
	EncryptionKey *string

	// The number of days until the data in the object expires.
	ExpirationDays *int32

	// A map of the name and ObjectType field.
	Fields map[string]types.ObjectTypeField

	// A list of unique keys that can be used to map data to the profile.
	Keys map[string][]types.ObjectTypeKey

	// The amount of profile object max count assigned to the object type
	MaxProfileObjectCount *int32

	// The format of your sourceLastUpdatedTimestamp that was previously set up.
	SourceLastUpdatedTimestampFormat *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// A unique identifier for the object template. For some attributes in the
	// request, the service will use the default value from the object template when
	// TemplateId is present. If these attributes are present in the request, the
	// service may return a BadRequestException . These attributes include:
	// AllowProfileCreation, SourceLastUpdatedTimestampFormat, Fields, and Keys. For
	// example, if AllowProfileCreation is set to true when TemplateId is set, the
	// service may return a BadRequestException .
	TemplateId *string

	noSmithyDocumentSerde
}

type PutProfileObjectTypeOutput struct {

	// Description of the profile object type.
	//
	// This member is required.
	Description *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// Indicates whether a profile should be created when data is received if one
	// doesn’t exist for an object of this type. The default is FALSE . If the
	// AllowProfileCreation flag is set to FALSE , then the service tries to fetch a
	// standard profile and associate this object with the profile. If it is set to
	// TRUE , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation bool

	// The timestamp of when the domain was created.
	CreatedAt *time.Time

	// The customer-provided key to encrypt the profile object that will be created in
	// this profile object type.
	EncryptionKey *string

	// The number of days until the data in the object expires.
	ExpirationDays *int32

	// A map of the name and ObjectType field.
	Fields map[string]types.ObjectTypeField

	// A list of unique keys that can be used to map data to the profile.
	Keys map[string][]types.ObjectTypeKey

	// The timestamp of when the domain was most recently edited.
	LastUpdatedAt *time.Time

	// The amount of provisioned profile object max count available.
	MaxAvailableProfileObjectCount *int32

	// The amount of profile object max count assigned to the object type.
	MaxProfileObjectCount *int32

	// The format of your sourceLastUpdatedTimestamp that was previously set up in
	// fields that were parsed using [SimpleDateFormat]. If you have sourceLastUpdatedTimestamp in your
	// field, you must set up sourceLastUpdatedTimestampFormat .
	//
	// [SimpleDateFormat]: https://docs.oracle.com/javase/10/docs/api/java/text/SimpleDateFormat.html
	SourceLastUpdatedTimestampFormat *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// A unique identifier for the object template.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutProfileObjectTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutProfileObjectType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutProfileObjectType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutProfileObjectType"); err != nil {
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
	if err = addOpPutProfileObjectTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutProfileObjectType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutProfileObjectType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutProfileObjectType",
	}
}
