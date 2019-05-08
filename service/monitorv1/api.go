// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package monitorv1

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetMetricStatisticsBatch = "GetMetricStatisticsBatch"

// GetMetricStatisticsBatchRequest generates a "ksc/request.Request" representing the
// client's request for the GetMetricStatisticsBatch operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetMetricStatisticsBatch for more information on using the GetMetricStatisticsBatch
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetMetricStatisticsBatchRequest method.
//    req, resp := client.GetMetricStatisticsBatchRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2018-09-29/GetMetricStatisticsBatch
func (c *Monitorv1) GetMetricStatisticsBatchRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetMetricStatisticsBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetMetricStatisticsBatch API operation for monitorv1.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv1's
// API operation GetMetricStatisticsBatch for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2018-09-29/GetMetricStatisticsBatch
func (c *Monitorv1) GetMetricStatisticsBatch(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetMetricStatisticsBatchRequest(input)
	return out, req.Send()
}

// GetMetricStatisticsBatchWithContext is the same as GetMetricStatisticsBatch with the addition of
// the ability to pass a context and additional request options.
//
// See GetMetricStatisticsBatch for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv1) GetMetricStatisticsBatchWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetMetricStatisticsBatchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}