// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stores a geofence geometry in a given geofence collection, or updates the
// geometry of an existing geofence if a geofence ID is included in the request.
func (c *Client) PutGeofence(ctx context.Context, params *PutGeofenceInput, optFns ...func(*Options)) (*PutGeofenceOutput, error) {
	if params == nil {
		params = &PutGeofenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGeofence", params, optFns, c.addOperationPutGeofenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGeofenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutGeofenceInput struct {

	// The geofence collection to store the geofence in.
	//
	// This member is required.
	CollectionName *string

	// An identifier for the geofence. For example, ExampleGeofence-1 .
	//
	// This member is required.
	GeofenceId *string

	// Contains the details to specify the position of the geofence. Can be a polygon,
	// a circle or a polygon encoded in Geobuf format. Including multiple selections
	// will return a validation error.
	//
	// The [geofence polygon] format supports a maximum of 1,000 vertices. The [Geofence Geobuf] format supports a
	// maximum of 100,000 vertices.
	//
	// [Geofence Geobuf]: https://docs.aws.amazon.com/location-geofences/latest/APIReference/API_GeofenceGeometry.html
	// [geofence polygon]: https://docs.aws.amazon.com/location-geofences/latest/APIReference/API_GeofenceGeometry.html
	//
	// This member is required.
	Geometry *types.GeofenceGeometry

	// Associates one of more properties with the geofence. A property is a key-value
	// pair stored with the geofence and added to any geofence event triggered with
	// that geofence.
	//
	// Format: "key" : "value"
	GeofenceProperties map[string]string

	noSmithyDocumentSerde
}

type PutGeofenceOutput struct {

	// The timestamp for when the geofence was created in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	CreateTime *time.Time

	// The geofence identifier entered in the request.
	//
	// This member is required.
	GeofenceId *string

	// The timestamp for when the geofence was last updated in [ISO 8601] format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// [ISO 8601]: https://www.iso.org/iso-8601-date-and-time-format.html
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutGeofenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutGeofence"); err != nil {
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
	if err = addEndpointPrefix_opPutGeofenceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutGeofenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGeofence(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opPutGeofenceMiddleware struct {
}

func (*endpointPrefix_opPutGeofenceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutGeofenceMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opPutGeofenceMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutGeofenceMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutGeofence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutGeofence",
	}
}
