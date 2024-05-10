// Code generated by smithy-go-codegen DO NOT EDIT.

package smithyrpcv2cbor

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_SparseNullsOperation_smithyRpcv2cborSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *SparseNullsOperationInput
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
		// Serializes null values in maps
		"RpcV2CborSparseMapsSerializeNullValues": {
			Params: &SparseNullsOperationInput{
				SparseStringMap: map[string]*string{
					"foo": nil,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/service/RpcV2Protocol/operation/SparseNullsOperation",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/cbor",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareCBOR(actual, `v29zcGFyc2VTdHJpbmdNYXC/Y2Zvb/b//w==`)
			},
		},
		// Serializes null values in lists
		"RpcV2CborSparseListsSerializeNull": {
			Params: &SparseNullsOperationInput{
				SparseStringList: []*string{
					nil,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/service/RpcV2Protocol/operation/SparseNullsOperation",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/cbor",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareCBOR(actual, `v3BzcGFyc2VTdHJpbmdMaXN0n/b//w==`)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
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
				HTTPClient: protocoltesthttp.NewClient(),
				Region:     "us-west-2",
			})
			result, err := client.SparseNullsOperation(context.Background(), c.Params, func(options *Options) {
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

func TestClient_SparseNullsOperation_smithyRpcv2cborDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SparseNullsOperationOutput
	}{
		// Deserializes null values in maps
		"RpcV2CborSparseMapsDeserializeNullValues": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`v29zcGFyc2VTdHJpbmdNYXC/Y2Zvb/b//w==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &SparseNullsOperationOutput{
				SparseStringMap: map[string]*string{
					"foo": nil,
				},
			},
		},
		// Deserializes null values in lists
		"RpcV2CborSparseListsDeserializeNull": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`v3BzcGFyc2VTdHJpbmdMaXN0n/b//w==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &SparseNullsOperationOutput{
				SparseStringList: []*string{
					nil,
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
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
				Region: "us-west-2",
			})
			var params SparseNullsOperationInput
			result, err := client.SparseNullsOperation(context.Background(), &params)
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
