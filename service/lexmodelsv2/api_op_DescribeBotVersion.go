// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	jmespath "github.com/jmespath/go-jmespath"
	"time"
)

// Provides metadata about a version of a bot.
func (c *Client) DescribeBotVersion(ctx context.Context, params *DescribeBotVersionInput, optFns ...func(*Options)) (*DescribeBotVersionOutput, error) {
	if params == nil {
		params = &DescribeBotVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotVersion", params, optFns, c.addOperationDescribeBotVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotVersionInput struct {

	// The identifier of the bot containing the version to return metadata for.
	//
	// This member is required.
	BotId *string

	// The version of the bot to return metadata for.
	//
	// This member is required.
	BotVersion *string

	noSmithyDocumentSerde
}

type DescribeBotVersionOutput struct {

	// The identifier of the bot that contains the version.
	BotId *string

	// The members of bot network in the version that was described.
	BotMembers []types.BotMember

	// The name of the bot that contains the version.
	BotName *string

	// The current status of the bot. When the status is Available , the bot version is
	// ready for use.
	BotStatus types.BotStatus

	// The type of the bot in the version that was described.
	BotType types.BotType

	// The version of the bot that was described.
	BotVersion *string

	// A timestamp of the date and time that the bot version was created.
	CreationDateTime *time.Time

	// Data privacy settings for the bot version.
	DataPrivacy *types.DataPrivacy

	// The description specified for the bot.
	Description *string

	// If the botStatus is Failed , this contains a list of reasons that the version
	// couldn't be built.
	FailureReasons []string

	// The number of seconds that a session with the bot remains active before it is
	// discarded by Amazon Lex.
	IdleSessionTTLInSeconds *int32

	// A list of the networks to which the bot version you described belongs.
	ParentBotNetworks []types.ParentBotNetwork

	// The Amazon Resource Name (ARN) of an IAM role that has permission to access the
	// bot version.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBotVersion"); err != nil {
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
	if err = addOpDescribeBotVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotVersion(options.Region), middleware.Before); err != nil {
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

// BotVersionAvailableWaiterOptions are waiter options for
// BotVersionAvailableWaiter
type BotVersionAvailableWaiterOptions struct {

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
	// BotVersionAvailableWaiter will use default minimum delay of 10 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, BotVersionAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeBotVersionInput, *DescribeBotVersionOutput, error) (bool, error)
}

// BotVersionAvailableWaiter defines the waiters for BotVersionAvailable
type BotVersionAvailableWaiter struct {
	client DescribeBotVersionAPIClient

	options BotVersionAvailableWaiterOptions
}

// NewBotVersionAvailableWaiter constructs a BotVersionAvailableWaiter.
func NewBotVersionAvailableWaiter(client DescribeBotVersionAPIClient, optFns ...func(*BotVersionAvailableWaiterOptions)) *BotVersionAvailableWaiter {
	options := BotVersionAvailableWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botVersionAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotVersionAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotVersionAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *BotVersionAvailableWaiter) Wait(ctx context.Context, params *DescribeBotVersionInput, maxWaitDur time.Duration, optFns ...func(*BotVersionAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotVersionAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *BotVersionAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeBotVersionInput, maxWaitDur time.Duration, optFns ...func(*BotVersionAvailableWaiterOptions)) (*DescribeBotVersionOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
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

		out, err := w.client.DescribeBotVersion(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for BotVersionAvailable waiter")
}

func botVersionAvailableStateRetryable(ctx context.Context, input *DescribeBotVersionInput, output *DescribeBotVersionOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("botStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Available"
		value, ok := pathValue.(types.BotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.BotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.BotStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// DescribeBotVersionAPIClient is a client that implements the DescribeBotVersion
// operation.
type DescribeBotVersionAPIClient interface {
	DescribeBotVersion(context.Context, *DescribeBotVersionInput, ...func(*Options)) (*DescribeBotVersionOutput, error)
}

var _ DescribeBotVersionAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeBotVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBotVersion",
	}
}
