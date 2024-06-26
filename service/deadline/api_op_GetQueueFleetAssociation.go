// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	jmespath "github.com/jmespath/go-jmespath"
	"time"
)

// Gets a queue-fleet association.
func (c *Client) GetQueueFleetAssociation(ctx context.Context, params *GetQueueFleetAssociationInput, optFns ...func(*Options)) (*GetQueueFleetAssociationOutput, error) {
	if params == nil {
		params = &GetQueueFleetAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueueFleetAssociation", params, optFns, c.addOperationGetQueueFleetAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueFleetAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueueFleetAssociationInput struct {

	// The farm ID of the farm that contains the queue-fleet association.
	//
	// This member is required.
	FarmId *string

	// The fleet ID for the queue-fleet association.
	//
	// This member is required.
	FleetId *string

	// The queue ID for the queue-fleet association.
	//
	// This member is required.
	QueueId *string

	noSmithyDocumentSerde
}

type GetQueueFleetAssociationOutput struct {

	// The date and time the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user or system that created this resource.
	//
	// This member is required.
	CreatedBy *string

	// The fleet ID for the queue-fleet association.
	//
	// This member is required.
	FleetId *string

	// The queue ID for the queue-fleet association.
	//
	// This member is required.
	QueueId *string

	// The status of the queue-fleet association.
	//
	// This member is required.
	Status types.QueueFleetAssociationStatus

	// The date and time the resource was updated.
	UpdatedAt *time.Time

	// The user or system that updated this resource.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueueFleetAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQueueFleetAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQueueFleetAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueueFleetAssociation"); err != nil {
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
	if err = addEndpointPrefix_opGetQueueFleetAssociationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetQueueFleetAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueFleetAssociation(options.Region), middleware.Before); err != nil {
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

// QueueFleetAssociationStoppedWaiterOptions are waiter options for
// QueueFleetAssociationStoppedWaiter
type QueueFleetAssociationStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// QueueFleetAssociationStoppedWaiter will use default minimum delay of 10 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, QueueFleetAssociationStoppedWaiter will use default max delay of
	// 600 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetQueueFleetAssociationInput, *GetQueueFleetAssociationOutput, error) (bool, error)
}

// QueueFleetAssociationStoppedWaiter defines the waiters for
// QueueFleetAssociationStopped
type QueueFleetAssociationStoppedWaiter struct {
	client GetQueueFleetAssociationAPIClient

	options QueueFleetAssociationStoppedWaiterOptions
}

// NewQueueFleetAssociationStoppedWaiter constructs a
// QueueFleetAssociationStoppedWaiter.
func NewQueueFleetAssociationStoppedWaiter(client GetQueueFleetAssociationAPIClient, optFns ...func(*QueueFleetAssociationStoppedWaiterOptions)) *QueueFleetAssociationStoppedWaiter {
	options := QueueFleetAssociationStoppedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = queueFleetAssociationStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &QueueFleetAssociationStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for QueueFleetAssociationStopped waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *QueueFleetAssociationStoppedWaiter) Wait(ctx context.Context, params *GetQueueFleetAssociationInput, maxWaitDur time.Duration, optFns ...func(*QueueFleetAssociationStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for QueueFleetAssociationStopped waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *QueueFleetAssociationStoppedWaiter) WaitForOutput(ctx context.Context, params *GetQueueFleetAssociationInput, maxWaitDur time.Duration, optFns ...func(*QueueFleetAssociationStoppedWaiterOptions)) (*GetQueueFleetAssociationOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetQueueFleetAssociation(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for QueueFleetAssociationStopped waiter")
}

func queueFleetAssociationStoppedStateRetryable(ctx context.Context, input *GetQueueFleetAssociationInput, output *GetQueueFleetAssociationOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STOPPED"
		value, ok := pathValue.(types.QueueFleetAssociationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.QueueFleetAssociationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

type endpointPrefix_opGetQueueFleetAssociationMiddleware struct {
}

func (*endpointPrefix_opGetQueueFleetAssociationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetQueueFleetAssociationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetQueueFleetAssociationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetQueueFleetAssociationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetQueueFleetAssociationAPIClient is a client that implements the
// GetQueueFleetAssociation operation.
type GetQueueFleetAssociationAPIClient interface {
	GetQueueFleetAssociation(context.Context, *GetQueueFleetAssociationInput, ...func(*Options)) (*GetQueueFleetAssociationOutput, error)
}

var _ GetQueueFleetAssociationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetQueueFleetAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQueueFleetAssociation",
	}
}
