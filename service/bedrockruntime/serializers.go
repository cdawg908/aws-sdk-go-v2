// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/document"
	internaldocument "github.com/aws/aws-sdk-go-v2/service/bedrockruntime/internal/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"math"
)

type awsRestjson1_serializeOpConverse struct {
}

func (*awsRestjson1_serializeOpConverse) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpConverse) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ConverseInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/model/{modelId}/converse")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsConverseInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentConverseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsConverseInput(v *ConverseInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ModelId == nil || len(*v.ModelId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member modelId must not be empty")}
	}
	if v.ModelId != nil {
		if err := encoder.SetURI("modelId").String(*v.ModelId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentConverseInput(v *ConverseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AdditionalModelRequestFields != nil {
		ok := object.Key("additionalModelRequestFields")
		if err := awsRestjson1_serializeDocumentDocument(v.AdditionalModelRequestFields, ok); err != nil {
			return err
		}
	}

	if v.AdditionalModelResponseFieldPaths != nil {
		ok := object.Key("additionalModelResponseFieldPaths")
		if err := awsRestjson1_serializeDocumentAdditionalModelResponseFieldPaths(v.AdditionalModelResponseFieldPaths, ok); err != nil {
			return err
		}
	}

	if v.GuardrailConfig != nil {
		ok := object.Key("guardrailConfig")
		if err := awsRestjson1_serializeDocumentGuardrailConfiguration(v.GuardrailConfig, ok); err != nil {
			return err
		}
	}

	if v.InferenceConfig != nil {
		ok := object.Key("inferenceConfig")
		if err := awsRestjson1_serializeDocumentInferenceConfiguration(v.InferenceConfig, ok); err != nil {
			return err
		}
	}

	if v.Messages != nil {
		ok := object.Key("messages")
		if err := awsRestjson1_serializeDocumentMessages(v.Messages, ok); err != nil {
			return err
		}
	}

	if v.System != nil {
		ok := object.Key("system")
		if err := awsRestjson1_serializeDocumentSystemContentBlocks(v.System, ok); err != nil {
			return err
		}
	}

	if v.ToolConfig != nil {
		ok := object.Key("toolConfig")
		if err := awsRestjson1_serializeDocumentToolConfiguration(v.ToolConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpConverseStream struct {
}

func (*awsRestjson1_serializeOpConverseStream) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpConverseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ConverseStreamInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/model/{modelId}/converse-stream")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsConverseStreamInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentConverseStreamInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsConverseStreamInput(v *ConverseStreamInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ModelId == nil || len(*v.ModelId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member modelId must not be empty")}
	}
	if v.ModelId != nil {
		if err := encoder.SetURI("modelId").String(*v.ModelId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentConverseStreamInput(v *ConverseStreamInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AdditionalModelRequestFields != nil {
		ok := object.Key("additionalModelRequestFields")
		if err := awsRestjson1_serializeDocumentDocument(v.AdditionalModelRequestFields, ok); err != nil {
			return err
		}
	}

	if v.AdditionalModelResponseFieldPaths != nil {
		ok := object.Key("additionalModelResponseFieldPaths")
		if err := awsRestjson1_serializeDocumentAdditionalModelResponseFieldPaths(v.AdditionalModelResponseFieldPaths, ok); err != nil {
			return err
		}
	}

	if v.GuardrailConfig != nil {
		ok := object.Key("guardrailConfig")
		if err := awsRestjson1_serializeDocumentGuardrailStreamConfiguration(v.GuardrailConfig, ok); err != nil {
			return err
		}
	}

	if v.InferenceConfig != nil {
		ok := object.Key("inferenceConfig")
		if err := awsRestjson1_serializeDocumentInferenceConfiguration(v.InferenceConfig, ok); err != nil {
			return err
		}
	}

	if v.Messages != nil {
		ok := object.Key("messages")
		if err := awsRestjson1_serializeDocumentMessages(v.Messages, ok); err != nil {
			return err
		}
	}

	if v.System != nil {
		ok := object.Key("system")
		if err := awsRestjson1_serializeDocumentSystemContentBlocks(v.System, ok); err != nil {
			return err
		}
	}

	if v.ToolConfig != nil {
		ok := object.Key("toolConfig")
		if err := awsRestjson1_serializeDocumentToolConfiguration(v.ToolConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpInvokeModel struct {
}

func (*awsRestjson1_serializeOpInvokeModel) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpInvokeModel) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*InvokeModelInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/model/{modelId}/invoke")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsInvokeModelInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if !restEncoder.HasHeader("Content-Type") {
		ctx = smithyhttp.SetIsContentTypeDefaultValue(ctx, true)
		restEncoder.SetHeader("Content-Type").String("application/octet-stream")
	}

	if input.Body != nil {
		payload := bytes.NewReader(input.Body)
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsInvokeModelInput(v *InvokeModelInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Accept != nil && len(*v.Accept) > 0 {
		locationName := "Accept"
		encoder.SetHeader(locationName).String(*v.Accept)
	}

	if v.ContentType != nil && len(*v.ContentType) > 0 {
		locationName := "Content-Type"
		encoder.SetHeader(locationName).String(*v.ContentType)
	}

	if v.GuardrailIdentifier != nil && len(*v.GuardrailIdentifier) > 0 {
		locationName := "X-Amzn-Bedrock-Guardrailidentifier"
		encoder.SetHeader(locationName).String(*v.GuardrailIdentifier)
	}

	if v.GuardrailVersion != nil && len(*v.GuardrailVersion) > 0 {
		locationName := "X-Amzn-Bedrock-Guardrailversion"
		encoder.SetHeader(locationName).String(*v.GuardrailVersion)
	}

	if v.ModelId == nil || len(*v.ModelId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member modelId must not be empty")}
	}
	if v.ModelId != nil {
		if err := encoder.SetURI("modelId").String(*v.ModelId); err != nil {
			return err
		}
	}

	if len(v.Trace) > 0 {
		locationName := "X-Amzn-Bedrock-Trace"
		encoder.SetHeader(locationName).String(string(v.Trace))
	}

	return nil
}

type awsRestjson1_serializeOpInvokeModelWithResponseStream struct {
}

func (*awsRestjson1_serializeOpInvokeModelWithResponseStream) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpInvokeModelWithResponseStream) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*InvokeModelWithResponseStreamInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/model/{modelId}/invoke-with-response-stream")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsInvokeModelWithResponseStreamInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if !restEncoder.HasHeader("Content-Type") {
		ctx = smithyhttp.SetIsContentTypeDefaultValue(ctx, true)
		restEncoder.SetHeader("Content-Type").String("application/octet-stream")
	}

	if input.Body != nil {
		payload := bytes.NewReader(input.Body)
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsInvokeModelWithResponseStreamInput(v *InvokeModelWithResponseStreamInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Accept != nil && len(*v.Accept) > 0 {
		locationName := "X-Amzn-Bedrock-Accept"
		encoder.SetHeader(locationName).String(*v.Accept)
	}

	if v.ContentType != nil && len(*v.ContentType) > 0 {
		locationName := "Content-Type"
		encoder.SetHeader(locationName).String(*v.ContentType)
	}

	if v.GuardrailIdentifier != nil && len(*v.GuardrailIdentifier) > 0 {
		locationName := "X-Amzn-Bedrock-Guardrailidentifier"
		encoder.SetHeader(locationName).String(*v.GuardrailIdentifier)
	}

	if v.GuardrailVersion != nil && len(*v.GuardrailVersion) > 0 {
		locationName := "X-Amzn-Bedrock-Guardrailversion"
		encoder.SetHeader(locationName).String(*v.GuardrailVersion)
	}

	if v.ModelId == nil || len(*v.ModelId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member modelId must not be empty")}
	}
	if v.ModelId != nil {
		if err := encoder.SetURI("modelId").String(*v.ModelId); err != nil {
			return err
		}
	}

	if len(v.Trace) > 0 {
		locationName := "X-Amzn-Bedrock-Trace"
		encoder.SetHeader(locationName).String(string(v.Trace))
	}

	return nil
}

func awsRestjson1_serializeDocumentAdditionalModelResponseFieldPaths(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentAnyToolChoice(v *types.AnyToolChoice, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsRestjson1_serializeDocumentAutoToolChoice(v *types.AutoToolChoice, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsRestjson1_serializeDocumentContentBlock(v types.ContentBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ContentBlockMemberGuardContent:
		av := object.Key("guardContent")
		if err := awsRestjson1_serializeDocumentGuardrailConverseContentBlock(uv.Value, av); err != nil {
			return err
		}

	case *types.ContentBlockMemberImage:
		av := object.Key("image")
		if err := awsRestjson1_serializeDocumentImageBlock(&uv.Value, av); err != nil {
			return err
		}

	case *types.ContentBlockMemberText:
		av := object.Key("text")
		av.String(uv.Value)

	case *types.ContentBlockMemberToolResult:
		av := object.Key("toolResult")
		if err := awsRestjson1_serializeDocumentToolResultBlock(&uv.Value, av); err != nil {
			return err
		}

	case *types.ContentBlockMemberToolUse:
		av := object.Key("toolUse")
		if err := awsRestjson1_serializeDocumentToolUseBlock(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentContentBlocks(v []types.ContentBlock, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentContentBlock(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentGuardrailConfiguration(v *types.GuardrailConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GuardrailIdentifier != nil {
		ok := object.Key("guardrailIdentifier")
		ok.String(*v.GuardrailIdentifier)
	}

	if v.GuardrailVersion != nil {
		ok := object.Key("guardrailVersion")
		ok.String(*v.GuardrailVersion)
	}

	if len(v.Trace) > 0 {
		ok := object.Key("trace")
		ok.String(string(v.Trace))
	}

	return nil
}

func awsRestjson1_serializeDocumentGuardrailConverseContentBlock(v types.GuardrailConverseContentBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.GuardrailConverseContentBlockMemberText:
		av := object.Key("text")
		if err := awsRestjson1_serializeDocumentGuardrailConverseTextBlock(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentGuardrailConverseTextBlock(v *types.GuardrailConverseTextBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Text != nil {
		ok := object.Key("text")
		ok.String(*v.Text)
	}

	return nil
}

func awsRestjson1_serializeDocumentGuardrailStreamConfiguration(v *types.GuardrailStreamConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.GuardrailIdentifier != nil {
		ok := object.Key("guardrailIdentifier")
		ok.String(*v.GuardrailIdentifier)
	}

	if v.GuardrailVersion != nil {
		ok := object.Key("guardrailVersion")
		ok.String(*v.GuardrailVersion)
	}

	if len(v.StreamProcessingMode) > 0 {
		ok := object.Key("streamProcessingMode")
		ok.String(string(v.StreamProcessingMode))
	}

	if len(v.Trace) > 0 {
		ok := object.Key("trace")
		ok.String(string(v.Trace))
	}

	return nil
}

func awsRestjson1_serializeDocumentImageBlock(v *types.ImageBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Format) > 0 {
		ok := object.Key("format")
		ok.String(string(v.Format))
	}

	if v.Source != nil {
		ok := object.Key("source")
		if err := awsRestjson1_serializeDocumentImageSource(v.Source, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentImageSource(v types.ImageSource, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ImageSourceMemberBytes:
		av := object.Key("bytes")
		av.Base64EncodeBytes(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentInferenceConfiguration(v *types.InferenceConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxTokens != nil {
		ok := object.Key("maxTokens")
		ok.Integer(*v.MaxTokens)
	}

	if v.StopSequences != nil {
		ok := object.Key("stopSequences")
		if err := awsRestjson1_serializeDocumentNonEmptyStringList(v.StopSequences, ok); err != nil {
			return err
		}
	}

	if v.Temperature != nil {
		ok := object.Key("temperature")
		switch {
		case math.IsNaN(float64(*v.Temperature)):
			ok.String("NaN")

		case math.IsInf(float64(*v.Temperature), 1):
			ok.String("Infinity")

		case math.IsInf(float64(*v.Temperature), -1):
			ok.String("-Infinity")

		default:
			ok.Float(*v.Temperature)

		}
	}

	if v.TopP != nil {
		ok := object.Key("topP")
		switch {
		case math.IsNaN(float64(*v.TopP)):
			ok.String("NaN")

		case math.IsInf(float64(*v.TopP), 1):
			ok.String("Infinity")

		case math.IsInf(float64(*v.TopP), -1):
			ok.String("-Infinity")

		default:
			ok.Float(*v.TopP)

		}
	}

	return nil
}

func awsRestjson1_serializeDocumentMessage(v *types.Message, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Content != nil {
		ok := object.Key("content")
		if err := awsRestjson1_serializeDocumentContentBlocks(v.Content, ok); err != nil {
			return err
		}
	}

	if len(v.Role) > 0 {
		ok := object.Key("role")
		ok.String(string(v.Role))
	}

	return nil
}

func awsRestjson1_serializeDocumentMessages(v []types.Message, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentMessage(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentNonEmptyStringList(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSpecificToolChoice(v *types.SpecificToolChoice, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	return nil
}

func awsRestjson1_serializeDocumentSystemContentBlock(v types.SystemContentBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.SystemContentBlockMemberGuardContent:
		av := object.Key("guardContent")
		if err := awsRestjson1_serializeDocumentGuardrailConverseContentBlock(uv.Value, av); err != nil {
			return err
		}

	case *types.SystemContentBlockMemberText:
		av := object.Key("text")
		av.String(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentSystemContentBlocks(v []types.SystemContentBlock, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentSystemContentBlock(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentTool(v types.Tool, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ToolMemberToolSpec:
		av := object.Key("toolSpec")
		if err := awsRestjson1_serializeDocumentToolSpecification(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentToolChoice(v types.ToolChoice, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ToolChoiceMemberAny:
		av := object.Key("any")
		if err := awsRestjson1_serializeDocumentAnyToolChoice(&uv.Value, av); err != nil {
			return err
		}

	case *types.ToolChoiceMemberAuto:
		av := object.Key("auto")
		if err := awsRestjson1_serializeDocumentAutoToolChoice(&uv.Value, av); err != nil {
			return err
		}

	case *types.ToolChoiceMemberTool:
		av := object.Key("tool")
		if err := awsRestjson1_serializeDocumentSpecificToolChoice(&uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentToolConfiguration(v *types.ToolConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ToolChoice != nil {
		ok := object.Key("toolChoice")
		if err := awsRestjson1_serializeDocumentToolChoice(v.ToolChoice, ok); err != nil {
			return err
		}
	}

	if v.Tools != nil {
		ok := object.Key("tools")
		if err := awsRestjson1_serializeDocumentTools(v.Tools, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentToolInputSchema(v types.ToolInputSchema, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ToolInputSchemaMemberJson:
		av := object.Key("json")
		if err := awsRestjson1_serializeDocumentDocument(uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentToolResultBlock(v *types.ToolResultBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Content != nil {
		ok := object.Key("content")
		if err := awsRestjson1_serializeDocumentToolResultContentBlocks(v.Content, ok); err != nil {
			return err
		}
	}

	if len(v.Status) > 0 {
		ok := object.Key("status")
		ok.String(string(v.Status))
	}

	if v.ToolUseId != nil {
		ok := object.Key("toolUseId")
		ok.String(*v.ToolUseId)
	}

	return nil
}

func awsRestjson1_serializeDocumentToolResultContentBlock(v types.ToolResultContentBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ToolResultContentBlockMemberImage:
		av := object.Key("image")
		if err := awsRestjson1_serializeDocumentImageBlock(&uv.Value, av); err != nil {
			return err
		}

	case *types.ToolResultContentBlockMemberJson:
		av := object.Key("json")
		if err := awsRestjson1_serializeDocumentDocument(uv.Value, av); err != nil {
			return err
		}

	case *types.ToolResultContentBlockMemberText:
		av := object.Key("text")
		av.String(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentToolResultContentBlocks(v []types.ToolResultContentBlock, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentToolResultContentBlock(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentTools(v []types.Tool, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentTool(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentToolSpecification(v *types.ToolSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Description != nil {
		ok := object.Key("description")
		ok.String(*v.Description)
	}

	if v.InputSchema != nil {
		ok := object.Key("inputSchema")
		if err := awsRestjson1_serializeDocumentToolInputSchema(v.InputSchema, ok); err != nil {
			return err
		}
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	return nil
}

func awsRestjson1_serializeDocumentToolUseBlock(v *types.ToolUseBlock, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Input != nil {
		ok := object.Key("input")
		if err := awsRestjson1_serializeDocumentDocument(v.Input, ok); err != nil {
			return err
		}
	}

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.ToolUseId != nil {
		ok := object.Key("toolUseId")
		ok.String(*v.ToolUseId)
	}

	return nil
}

func awsRestjson1_serializeDocumentDocument(v document.Interface, value smithyjson.Value) error {
	if v == nil {
		return nil
	}
	if !internaldocument.IsInterface(v) {
		return fmt.Errorf("%T is not a compatible document type", v)
	}
	db, err := v.MarshalSmithyDocument()
	if err != nil {
		return err
	}
	value.Write(db)
	return nil
}
