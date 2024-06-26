// Code generated by smithy-go-codegen DO NOT EDIT.

package smithyrpcv2cbor

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
)

func TestClient_Float16_smithyRpcv2cborDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *Float16Output
	}{
		// Ensures that clients can correctly parse float16 +Inf.
		"RpcV2CborFloat16Inf": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`oWV2YWx1Zfl8AA==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &Float16Output{
				Value: ptr.Float64(math.Inf(1)),
			},
		},
		// Ensures that clients can correctly parse float16 -Inf.
		"RpcV2CborFloat16NegInf": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`oWV2YWx1Zfn8AA==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &Float16Output{
				Value: ptr.Float64(math.Inf(-1)),
			},
		},
		// Ensures that clients can correctly parse float16 NaN with high LSB.
		"RpcV2CborFloat16LSBNaN": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`oWV2YWx1Zfl8AQ==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &Float16Output{
				Value: ptr.Float64(math.NaN()),
			},
		},
		// Ensures that clients can correctly parse float16 NaN with high MSB.
		"RpcV2CborFloat16MSBNaN": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`oWV2YWx1Zfl+AA==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &Float16Output{
				Value: ptr.Float64(math.NaN()),
			},
		},
		// Ensures that clients can correctly parse a subnormal float16.
		"RpcV2CborFloat16Subnormal": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type":    []string{"application/cbor"},
				"smithy-protocol": []string{"rpc-v2-cbor"},
			},
			BodyMediaType: "application/cbor",
			Body: func() []byte {
				p, err := base64.StdEncoding.DecodeString(`oWV2YWx1ZfkAUA==`)
				if err != nil {
					panic(err)
				}

				return p
			}(),
			ExpectResult: &Float16Output{
				Value: ptr.Float64(4.76837158203125e-6),
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
			var params Float16Input
			result, err := client.Float16(context.Background(), &params)
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
