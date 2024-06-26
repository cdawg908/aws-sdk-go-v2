// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/document"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_OperationWithDefaults_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *OperationWithDefaultsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Client populates default values in input.
		"RestJsonClientPopulatesDefaultValuesInInput": {
			Params: &OperationWithDefaultsInput{
				Defaults: &types.Defaults{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/OperationWithDefaults",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "defaults": {
			        "defaultString": "hi",
			        "defaultBoolean": true,
			        "defaultList": [],
			        "defaultDocumentMap": {},
			        "defaultDocumentString": "hi",
			        "defaultDocumentBoolean": true,
			        "defaultDocumentList": [],
			        "defaultTimestamp": 0,
			        "defaultBlob": "YWJj",
			        "defaultByte": 1,
			        "defaultShort": 1,
			        "defaultInteger": 10,
			        "defaultLong": 100,
			        "defaultFloat": 1.0,
			        "defaultDouble": 1.0,
			        "defaultMap": {},
			        "defaultEnum": "FOO",
			        "defaultIntEnum": 1,
			        "emptyString": "",
			        "falseBoolean": false,
			        "emptyBlob": "",
			        "zeroByte": 0,
			        "zeroShort": 0,
			        "zeroInteger": 0,
			        "zeroLong": 0,
			        "zeroFloat": 0.0,
			        "zeroDouble": 0.0
			    }
			}`))
			},
		},
		// Client skips top level default values in input.
		"RestJsonClientSkipsTopLevelDefaultValuesInInput": {
			Params:        &OperationWithDefaultsInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/OperationWithDefaults",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			}`))
			},
		},
		// Client uses explicitly provided member values over defaults
		"RestJsonClientUsesExplicitlyProvidedMemberValuesOverDefaults": {
			Params: &OperationWithDefaultsInput{
				Defaults: &types.Defaults{
					DefaultString:  ptr.String("bye"),
					DefaultBoolean: ptr.Bool(true),
					DefaultList: []string{
						"a",
					},
					DefaultDocumentMap: document.NewLazyDocument(map[string]interface{}{
						"name": "Jack",
					}),
					DefaultDocumentString:  document.NewLazyDocument("bye"),
					DefaultDocumentBoolean: document.NewLazyDocument(true),
					DefaultDocumentList: document.NewLazyDocument([]interface{}{
						"b",
					}),
					DefaultNullDocument: document.NewLazyDocument("notNull"),
					DefaultTimestamp:    ptr.Time(smithytime.ParseEpochSeconds(1)),
					DefaultBlob:         []byte("hi"),
					DefaultByte:         ptr.Int8(2),
					DefaultShort:        ptr.Int16(2),
					DefaultInteger:      ptr.Int32(20),
					DefaultLong:         ptr.Int64(200),
					DefaultFloat:        ptr.Float32(2.0),
					DefaultDouble:       ptr.Float64(2.0),
					DefaultMap: map[string]string{
						"name": "Jack",
					},
					DefaultEnum:    types.TestEnum("BAR"),
					DefaultIntEnum: 2,
					EmptyString:    ptr.String("foo"),
					FalseBoolean:   true,
					EmptyBlob:      []byte("hi"),
					ZeroByte:       1,
					ZeroShort:      1,
					ZeroInteger:    1,
					ZeroLong:       1,
					ZeroFloat:      1.0,
					ZeroDouble:     1.0,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/OperationWithDefaults",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "defaults": {
			        "defaultString": "bye",
			        "defaultBoolean": true,
			        "defaultList": ["a"],
			        "defaultDocumentMap": {"name": "Jack"},
			        "defaultDocumentString": "bye",
			        "defaultDocumentBoolean": true,
			        "defaultDocumentList": ["b"],
			        "defaultNullDocument": "notNull",
			        "defaultTimestamp": 1,
			        "defaultBlob": "aGk=",
			        "defaultByte": 2,
			        "defaultShort": 2,
			        "defaultInteger": 20,
			        "defaultLong": 200,
			        "defaultFloat": 2.0,
			        "defaultDouble": 2.0,
			        "defaultMap": {"name": "Jack"},
			        "defaultEnum": "BAR",
			        "defaultIntEnum": 2,
			        "emptyString": "foo",
			        "falseBoolean": true,
			        "emptyBlob": "aGk=",
			        "zeroByte": 1,
			        "zeroShort": 1,
			        "zeroInteger": 1,
			        "zeroLong": 1,
			        "zeroFloat": 1.0,
			        "zeroDouble": 1.0
			    }
			}`))
			},
		},
		// Any time a value is provided for a member in the top level of input, it is
		// used, regardless of if its the default.
		"RestJsonClientUsesExplicitlyProvidedValuesInTopLevel": {
			Params: &OperationWithDefaultsInput{
				TopLevelDefault:      ptr.String("hi"),
				OtherTopLevelDefault: 0,
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/OperationWithDefaults",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "topLevelDefault": "hi",
			    "otherTopLevelDefault": 0
			}`))
			},
		},
		// Typically, non top-level members would have defaults filled in, but if they
		// have the clientOptional trait, the defaults should be ignored.
		"RestJsonClientIgnoresNonTopLevelDefaultsOnMembersWithClientOptional": {
			Params: &OperationWithDefaultsInput{
				ClientOptionalDefaults: &types.ClientOptionalDefaults{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/OperationWithDefaults",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "clientOptionalDefaults": {}
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "RestJsonClientPopulatesDefaultValuesInInput" {
				t.Skip("disabled test aws.protocoltests.restjson#RestJson aws.protocoltests.restjson#OperationWithDefaults")
			}

			if name == "RestJsonClientUsesExplicitlyProvidedValuesInTopLevel" {
				t.Skip("disabled test aws.protocoltests.restjson#RestJson aws.protocoltests.restjson#OperationWithDefaults")
			}

			actualReq := &http.Request{}
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.OperationWithDefaults(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_OperationWithDefaults_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *OperationWithDefaultsOutput
	}{
		// Client populates default values when missing in response.
		"RestJsonClientPopulatesDefaultsValuesWhenMissingInResponse": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body:          []byte(`{}`),
			ExpectResult: &OperationWithDefaultsOutput{
				DefaultString:          ptr.String("hi"),
				DefaultBoolean:         ptr.Bool(true),
				DefaultList:            []string{},
				DefaultDocumentMap:     document.NewLazyDocument(map[string]interface{}{}),
				DefaultDocumentString:  document.NewLazyDocument("hi"),
				DefaultDocumentBoolean: document.NewLazyDocument(true),
				DefaultDocumentList:    document.NewLazyDocument([]interface{}{}),
				DefaultTimestamp:       ptr.Time(smithytime.ParseEpochSeconds(0)),
				DefaultBlob:            []byte("abc"),
				DefaultByte:            ptr.Int8(1),
				DefaultShort:           ptr.Int16(1),
				DefaultInteger:         ptr.Int32(10),
				DefaultLong:            ptr.Int64(100),
				DefaultFloat:           ptr.Float32(1.0),
				DefaultDouble:          ptr.Float64(1.0),
				DefaultMap:             map[string]string{},
				DefaultEnum:            types.TestEnum("FOO"),
				DefaultIntEnum:         1,
				EmptyString:            ptr.String(""),
				FalseBoolean:           false,
				EmptyBlob:              []byte(""),
				ZeroByte:               0,
				ZeroShort:              0,
				ZeroInteger:            0,
				ZeroLong:               0,
				ZeroFloat:              0.0,
				ZeroDouble:             0.0,
			},
		},
		// Client ignores default values if member values are present in the response.
		"RestJsonClientIgnoresDefaultValuesIfMemberValuesArePresentInResponse": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "defaultString": "bye",
			    "defaultBoolean": false,
			    "defaultList": ["a"],
			    "defaultDocumentMap": {"name": "Jack"},
			    "defaultDocumentString": "bye",
			    "defaultDocumentBoolean": false,
			    "defaultDocumentList": ["b"],
			    "defaultNullDocument": "notNull",
			    "defaultTimestamp": 2,
			    "defaultBlob": "aGk=",
			    "defaultByte": 2,
			    "defaultShort": 2,
			    "defaultInteger": 20,
			    "defaultLong": 200,
			    "defaultFloat": 2.0,
			    "defaultDouble": 2.0,
			    "defaultMap": {"name": "Jack"},
			    "defaultEnum": "BAR",
			    "defaultIntEnum": 2,
			    "emptyString": "foo",
			    "falseBoolean": true,
			    "emptyBlob": "aGk=",
			    "zeroByte": 1,
			    "zeroShort": 1,
			    "zeroInteger": 1,
			    "zeroLong": 1,
			    "zeroFloat": 1.0,
			    "zeroDouble": 1.0
			}`),
			ExpectResult: &OperationWithDefaultsOutput{
				DefaultString:  ptr.String("bye"),
				DefaultBoolean: ptr.Bool(false),
				DefaultList: []string{
					"a",
				},
				DefaultDocumentMap: document.NewLazyDocument(map[string]interface{}{
					"name": "Jack",
				}),
				DefaultDocumentString:  document.NewLazyDocument("bye"),
				DefaultDocumentBoolean: document.NewLazyDocument(false),
				DefaultDocumentList: document.NewLazyDocument([]interface{}{
					"b",
				}),
				DefaultNullDocument: document.NewLazyDocument("notNull"),
				DefaultTimestamp:    ptr.Time(smithytime.ParseEpochSeconds(1)),
				DefaultBlob:         []byte("hi"),
				DefaultByte:         ptr.Int8(2),
				DefaultShort:        ptr.Int16(2),
				DefaultInteger:      ptr.Int32(20),
				DefaultLong:         ptr.Int64(200),
				DefaultFloat:        ptr.Float32(2.0),
				DefaultDouble:       ptr.Float64(2.0),
				DefaultMap: map[string]string{
					"name": "Jack",
				},
				DefaultEnum:    types.TestEnum("BAR"),
				DefaultIntEnum: 2,
				EmptyString:    ptr.String("foo"),
				FalseBoolean:   true,
				EmptyBlob:      []byte("hi"),
				ZeroByte:       1,
				ZeroShort:      1,
				ZeroInteger:    1,
				ZeroLong:       1,
				ZeroFloat:      1.0,
				ZeroDouble:     1.0,
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "RestJsonClientPopulatesDefaultsValuesWhenMissingInResponse" {
				t.Skip("disabled test aws.protocoltests.restjson#RestJson aws.protocoltests.restjson#OperationWithDefaults")
			}

			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params OperationWithDefaultsInput
			result, err := client.OperationWithDefaults(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
