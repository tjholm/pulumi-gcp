// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EdgeCacheOrigin represents a HTTP-reachable backend for an EdgeCacheService.
//
// ## Example Usage
// ### Network Services Edge Cache Origin Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/networkservices"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkservices.NewEdgeCacheOrigin(ctx, "default", &networkservices.EdgeCacheOriginArgs{
// 			Description:   pulumi.String("The default bucket for media edge test"),
// 			OriginAddress: pulumi.String("gs://media-edge-default"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Network Services Edge Cache Origin Advanced
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/networkservices"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fallback, err := networkservices.NewEdgeCacheOrigin(ctx, "fallback", &networkservices.EdgeCacheOriginArgs{
// 			OriginAddress: pulumi.String("gs://media-edge-fallback"),
// 			Description:   pulumi.String("The default bucket for media edge test"),
// 			MaxAttempts:   pulumi.Int(3),
// 			Protocol:      pulumi.String("HTTP"),
// 			Port:          pulumi.Int(80),
// 			RetryConditions: pulumi.StringArray{
// 				pulumi.String("CONNECT_FAILURE"),
// 				pulumi.String("NOT_FOUND"),
// 				pulumi.String("HTTP_5XX"),
// 			},
// 			Timeout: &networkservices.EdgeCacheOriginTimeoutArgs{
// 				ConnectTimeout:     pulumi.String("10s"),
// 				MaxAttemptsTimeout: pulumi.String("10s"),
// 				ResponseTimeout:    pulumi.String("10s"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = networkservices.NewEdgeCacheOrigin(ctx, "default", &networkservices.EdgeCacheOriginArgs{
// 			OriginAddress:  pulumi.String("gs://media-edge-default"),
// 			FailoverOrigin: fallback.ID(),
// 			Description:    pulumi.String("The default bucket for media edge test"),
// 			MaxAttempts:    pulumi.Int(2),
// 			Labels: pulumi.StringMap{
// 				"a": pulumi.String("b"),
// 			},
// 			Timeout: &networkservices.EdgeCacheOriginTimeoutArgs{
// 				ConnectTimeout: pulumi.String("10s"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// EdgeCacheOrigin can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default {{name}}
// ```
type EdgeCacheOrigin struct {
	pulumi.CustomResourceState

	// A human-readable description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	FailoverOrigin pulumi.StringPtrOutput `pulumi:"failoverOrigin"`
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	MaxAttempts pulumi.IntPtrOutput `pulumi:"maxAttempts"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress pulumi.StringOutput `pulumi:"originAddress"`
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
	// Possible values are `HTTP2`, `HTTPS`, and `HTTP`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Specifies one or more retry conditions for the configured origin.
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	// The default retryCondition is "CONNECT_FAILURE".
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	// Valid values are:
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	//   Each value may be one of `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, and `NOT_FOUND`.
	RetryConditions pulumi.StringArrayOutput `pulumi:"retryConditions"`
	// The connection and HTTP timeout configuration for this origin.
	// Structure is documented below.
	Timeout EdgeCacheOriginTimeoutPtrOutput `pulumi:"timeout"`
}

// NewEdgeCacheOrigin registers a new resource with the given unique name, arguments, and options.
func NewEdgeCacheOrigin(ctx *pulumi.Context,
	name string, args *EdgeCacheOriginArgs, opts ...pulumi.ResourceOption) (*EdgeCacheOrigin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OriginAddress == nil {
		return nil, errors.New("invalid value for required argument 'OriginAddress'")
	}
	var resource EdgeCacheOrigin
	err := ctx.RegisterResource("gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeCacheOrigin gets an existing EdgeCacheOrigin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeCacheOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeCacheOriginState, opts ...pulumi.ResourceOption) (*EdgeCacheOrigin, error) {
	var resource EdgeCacheOrigin
	err := ctx.ReadResource("gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeCacheOrigin resources.
type edgeCacheOriginState struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	FailoverOrigin *string `pulumi:"failoverOrigin"`
	// Set of label tags associated with the EdgeCache resource.
	Labels map[string]string `pulumi:"labels"`
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	MaxAttempts *int `pulumi:"maxAttempts"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress *string `pulumi:"originAddress"`
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
	// Possible values are `HTTP2`, `HTTPS`, and `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// Specifies one or more retry conditions for the configured origin.
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	// The default retryCondition is "CONNECT_FAILURE".
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	// Valid values are:
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	//   Each value may be one of `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, and `NOT_FOUND`.
	RetryConditions []string `pulumi:"retryConditions"`
	// The connection and HTTP timeout configuration for this origin.
	// Structure is documented below.
	Timeout *EdgeCacheOriginTimeout `pulumi:"timeout"`
}

type EdgeCacheOriginState struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	FailoverOrigin pulumi.StringPtrInput
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapInput
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	MaxAttempts pulumi.IntPtrInput
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress pulumi.StringPtrInput
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
	// Possible values are `HTTP2`, `HTTPS`, and `HTTP`.
	Protocol pulumi.StringPtrInput
	// Specifies one or more retry conditions for the configured origin.
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	// The default retryCondition is "CONNECT_FAILURE".
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	// Valid values are:
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	//   Each value may be one of `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, and `NOT_FOUND`.
	RetryConditions pulumi.StringArrayInput
	// The connection and HTTP timeout configuration for this origin.
	// Structure is documented below.
	Timeout EdgeCacheOriginTimeoutPtrInput
}

func (EdgeCacheOriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeCacheOriginState)(nil)).Elem()
}

type edgeCacheOriginArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	FailoverOrigin *string `pulumi:"failoverOrigin"`
	// Set of label tags associated with the EdgeCache resource.
	Labels map[string]string `pulumi:"labels"`
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	MaxAttempts *int `pulumi:"maxAttempts"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `pulumi:"name"`
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress string `pulumi:"originAddress"`
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
	// Possible values are `HTTP2`, `HTTPS`, and `HTTP`.
	Protocol *string `pulumi:"protocol"`
	// Specifies one or more retry conditions for the configured origin.
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	// The default retryCondition is "CONNECT_FAILURE".
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	// Valid values are:
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	//   Each value may be one of `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, and `NOT_FOUND`.
	RetryConditions []string `pulumi:"retryConditions"`
	// The connection and HTTP timeout configuration for this origin.
	// Structure is documented below.
	Timeout *EdgeCacheOriginTimeout `pulumi:"timeout"`
}

// The set of arguments for constructing a EdgeCacheOrigin resource.
type EdgeCacheOriginArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	FailoverOrigin pulumi.StringPtrInput
	// Set of label tags associated with the EdgeCache resource.
	Labels pulumi.StringMapInput
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	MaxAttempts pulumi.IntPtrInput
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name pulumi.StringPtrInput
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress pulumi.StringInput
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
	// Possible values are `HTTP2`, `HTTPS`, and `HTTP`.
	Protocol pulumi.StringPtrInput
	// Specifies one or more retry conditions for the configured origin.
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	// The default retryCondition is "CONNECT_FAILURE".
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	// Valid values are:
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
	//   Each value may be one of `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, and `NOT_FOUND`.
	RetryConditions pulumi.StringArrayInput
	// The connection and HTTP timeout configuration for this origin.
	// Structure is documented below.
	Timeout EdgeCacheOriginTimeoutPtrInput
}

func (EdgeCacheOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeCacheOriginArgs)(nil)).Elem()
}

type EdgeCacheOriginInput interface {
	pulumi.Input

	ToEdgeCacheOriginOutput() EdgeCacheOriginOutput
	ToEdgeCacheOriginOutputWithContext(ctx context.Context) EdgeCacheOriginOutput
}

func (*EdgeCacheOrigin) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeCacheOrigin)(nil)).Elem()
}

func (i *EdgeCacheOrigin) ToEdgeCacheOriginOutput() EdgeCacheOriginOutput {
	return i.ToEdgeCacheOriginOutputWithContext(context.Background())
}

func (i *EdgeCacheOrigin) ToEdgeCacheOriginOutputWithContext(ctx context.Context) EdgeCacheOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeCacheOriginOutput)
}

// EdgeCacheOriginArrayInput is an input type that accepts EdgeCacheOriginArray and EdgeCacheOriginArrayOutput values.
// You can construct a concrete instance of `EdgeCacheOriginArrayInput` via:
//
//          EdgeCacheOriginArray{ EdgeCacheOriginArgs{...} }
type EdgeCacheOriginArrayInput interface {
	pulumi.Input

	ToEdgeCacheOriginArrayOutput() EdgeCacheOriginArrayOutput
	ToEdgeCacheOriginArrayOutputWithContext(context.Context) EdgeCacheOriginArrayOutput
}

type EdgeCacheOriginArray []EdgeCacheOriginInput

func (EdgeCacheOriginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeCacheOrigin)(nil)).Elem()
}

func (i EdgeCacheOriginArray) ToEdgeCacheOriginArrayOutput() EdgeCacheOriginArrayOutput {
	return i.ToEdgeCacheOriginArrayOutputWithContext(context.Background())
}

func (i EdgeCacheOriginArray) ToEdgeCacheOriginArrayOutputWithContext(ctx context.Context) EdgeCacheOriginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeCacheOriginArrayOutput)
}

// EdgeCacheOriginMapInput is an input type that accepts EdgeCacheOriginMap and EdgeCacheOriginMapOutput values.
// You can construct a concrete instance of `EdgeCacheOriginMapInput` via:
//
//          EdgeCacheOriginMap{ "key": EdgeCacheOriginArgs{...} }
type EdgeCacheOriginMapInput interface {
	pulumi.Input

	ToEdgeCacheOriginMapOutput() EdgeCacheOriginMapOutput
	ToEdgeCacheOriginMapOutputWithContext(context.Context) EdgeCacheOriginMapOutput
}

type EdgeCacheOriginMap map[string]EdgeCacheOriginInput

func (EdgeCacheOriginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeCacheOrigin)(nil)).Elem()
}

func (i EdgeCacheOriginMap) ToEdgeCacheOriginMapOutput() EdgeCacheOriginMapOutput {
	return i.ToEdgeCacheOriginMapOutputWithContext(context.Background())
}

func (i EdgeCacheOriginMap) ToEdgeCacheOriginMapOutputWithContext(ctx context.Context) EdgeCacheOriginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeCacheOriginMapOutput)
}

type EdgeCacheOriginOutput struct{ *pulumi.OutputState }

func (EdgeCacheOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeCacheOrigin)(nil)).Elem()
}

func (o EdgeCacheOriginOutput) ToEdgeCacheOriginOutput() EdgeCacheOriginOutput {
	return o
}

func (o EdgeCacheOriginOutput) ToEdgeCacheOriginOutputWithContext(ctx context.Context) EdgeCacheOriginOutput {
	return o
}

type EdgeCacheOriginArrayOutput struct{ *pulumi.OutputState }

func (EdgeCacheOriginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgeCacheOrigin)(nil)).Elem()
}

func (o EdgeCacheOriginArrayOutput) ToEdgeCacheOriginArrayOutput() EdgeCacheOriginArrayOutput {
	return o
}

func (o EdgeCacheOriginArrayOutput) ToEdgeCacheOriginArrayOutputWithContext(ctx context.Context) EdgeCacheOriginArrayOutput {
	return o
}

func (o EdgeCacheOriginArrayOutput) Index(i pulumi.IntInput) EdgeCacheOriginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgeCacheOrigin {
		return vs[0].([]*EdgeCacheOrigin)[vs[1].(int)]
	}).(EdgeCacheOriginOutput)
}

type EdgeCacheOriginMapOutput struct{ *pulumi.OutputState }

func (EdgeCacheOriginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgeCacheOrigin)(nil)).Elem()
}

func (o EdgeCacheOriginMapOutput) ToEdgeCacheOriginMapOutput() EdgeCacheOriginMapOutput {
	return o
}

func (o EdgeCacheOriginMapOutput) ToEdgeCacheOriginMapOutputWithContext(ctx context.Context) EdgeCacheOriginMapOutput {
	return o
}

func (o EdgeCacheOriginMapOutput) MapIndex(k pulumi.StringInput) EdgeCacheOriginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgeCacheOrigin {
		return vs[0].(map[string]*EdgeCacheOrigin)[vs[1].(string)]
	}).(EdgeCacheOriginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeCacheOriginInput)(nil)).Elem(), &EdgeCacheOrigin{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeCacheOriginArrayInput)(nil)).Elem(), EdgeCacheOriginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeCacheOriginMapInput)(nil)).Elem(), EdgeCacheOriginMap{})
	pulumi.RegisterOutputType(EdgeCacheOriginOutput{})
	pulumi.RegisterOutputType(EdgeCacheOriginArrayOutput{})
	pulumi.RegisterOutputType(EdgeCacheOriginMapOutput{})
}
