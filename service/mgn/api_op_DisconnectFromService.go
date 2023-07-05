// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects specific Source Servers from Application Migration Service. Data
// replication is stopped immediately. All AWS resources created by Application
// Migration Service for enabling the replication of these source servers will be
// terminated / deleted within 90 minutes. Launched Test or Cutover instances will
// NOT be terminated. If the agent on the source server has not been prevented from
// communicating with the Application Migration Service service, then it will
// receive a command to uninstall itself (within approximately 10 minutes). The
// following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func (c *Client) DisconnectFromService(ctx context.Context, params *DisconnectFromServiceInput, optFns ...func(*Options)) (*DisconnectFromServiceOutput, error) {
	if params == nil {
		params = &DisconnectFromServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectFromService", params, optFns, c.addOperationDisconnectFromServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectFromServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectFromServiceInput struct {

	// Request to disconnect Source Server from service by Server ID.
	//
	// This member is required.
	SourceServerID *string

	// Request to disconnect Source Server from service by Account ID.
	AccountID *string

	noSmithyDocumentSerde
}

type DisconnectFromServiceOutput struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server fqdn for action framework.
	FqdnForActionFramework *string

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server replication type.
	ReplicationType types.ReplicationType

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server user provided ID.
	UserProvidedID *string

	// Source server vCenter client id.
	VcenterClientID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectFromServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisconnectFromService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisconnectFromService{}, middleware.After)
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
	if err = addOpDisconnectFromServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectFromService(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDisconnectFromService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "DisconnectFromService",
	}
}
