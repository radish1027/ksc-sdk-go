// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kcsv2iface provides an interface to enable mocking the kcsv2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kcsv2iface

import (
	"github.com/KscSDK/ksc-sdk-go/service/kcsv2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// Kcsv2API provides an interface to enable mocking the
// kcsv2.Kcsv2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // kcsv2.
//    func myFunc(svc kcsv2iface.Kcsv2API) bool {
//        // Make svc.AddCacheSlaveNode request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kcsv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKcsv2Client struct {
//        kcsv2iface.Kcsv2API
//    }
//    func (m *mockKcsv2Client) AddCacheSlaveNode(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKcsv2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Kcsv2API interface {
	AddCacheSlaveNode(*map[string]interface{}) (*map[string]interface{}, error)
	AddCacheSlaveNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddCacheSlaveNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheSlaveNode(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheSlaveNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheSlaveNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheReadonlyNode(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheReadonlyNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheReadonlyNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ Kcsv2API = (*kcsv2.Kcsv2)(nil)
