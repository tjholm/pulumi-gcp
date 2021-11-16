// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Target gRPC Proxy resource. A target gRPC proxy is a component
// of load balancers intended for load balancing gRPC traffic. Global forwarding
// rules reference a target gRPC proxy. The Target gRPC Proxy references
// a URL map which specifies how traffic routes to gRPC backend services.
//
// To get more information about TargetGrpcProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetGrpcProxies)
// * How-to Guides
//     * [Using Target gRPC Proxies](https://cloud.google.com/traffic-director/docs/proxyless-overview)
//
// ## Example Usage
// ### Target Grpc Proxy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultHealthCheck, err := compute.NewHealthCheck(ctx, "defaultHealthCheck", &compute.HealthCheckArgs{
// 			TimeoutSec:       pulumi.Int(1),
// 			CheckIntervalSec: pulumi.Int(1),
// 			GrpcHealthCheck: &compute.HealthCheckGrpcHealthCheckArgs{
// 				PortName:          pulumi.String("health-check-port"),
// 				PortSpecification: pulumi.String("USE_NAMED_PORT"),
// 				GrpcServiceName:   pulumi.String("testservice"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		home, err := compute.NewBackendService(ctx, "home", &compute.BackendServiceArgs{
// 			PortName:   pulumi.String("grpc"),
// 			Protocol:   pulumi.String("GRPC"),
// 			TimeoutSec: pulumi.Int(10),
// 			HealthChecks: pulumi.String{
// 				defaultHealthCheck.ID(),
// 			},
// 			LoadBalancingScheme: pulumi.String("INTERNAL_SELF_MANAGED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		urlmap, err := compute.NewURLMap(ctx, "urlmap", &compute.URLMapArgs{
// 			Description:    pulumi.String("a description"),
// 			DefaultService: home.ID(),
// 			HostRules: compute.URLMapHostRuleArray{
// 				&compute.URLMapHostRuleArgs{
// 					Hosts: pulumi.StringArray{
// 						pulumi.String("mysite.com"),
// 					},
// 					PathMatcher: pulumi.String("allpaths"),
// 				},
// 			},
// 			PathMatchers: compute.URLMapPathMatcherArray{
// 				&compute.URLMapPathMatcherArgs{
// 					Name:           pulumi.String("allpaths"),
// 					DefaultService: home.ID(),
// 					RouteRules: compute.URLMapPathMatcherRouteRuleArray{
// 						&compute.URLMapPathMatcherRouteRuleArgs{
// 							Priority: pulumi.Int(1),
// 							HeaderAction: &compute.URLMapPathMatcherRouteRuleHeaderActionArgs{
// 								RequestHeadersToRemoves: pulumi.StringArray{
// 									pulumi.String("RemoveMe2"),
// 								},
// 								RequestHeadersToAdds: compute.URLMapPathMatcherRouteRuleHeaderActionRequestHeadersToAddArray{
// 									&compute.URLMapPathMatcherRouteRuleHeaderActionRequestHeadersToAddArgs{
// 										HeaderName:  pulumi.String("AddSomethingElse"),
// 										HeaderValue: pulumi.String("MyOtherValue"),
// 										Replace:     pulumi.Bool(true),
// 									},
// 								},
// 								ResponseHeadersToRemoves: pulumi.StringArray{
// 									pulumi.String("RemoveMe3"),
// 								},
// 								ResponseHeadersToAdds: compute.URLMapPathMatcherRouteRuleHeaderActionResponseHeadersToAddArray{
// 									&compute.URLMapPathMatcherRouteRuleHeaderActionResponseHeadersToAddArgs{
// 										HeaderName:  pulumi.String("AddMe"),
// 										HeaderValue: pulumi.String("MyValue"),
// 										Replace:     pulumi.Bool(false),
// 									},
// 								},
// 							},
// 							MatchRules: compute.URLMapPathMatcherRouteRuleMatchRuleArray{
// 								&compute.URLMapPathMatcherRouteRuleMatchRuleArgs{
// 									FullPathMatch: pulumi.String("a full path"),
// 									HeaderMatches: compute.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArray{
// 										&compute.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs{
// 											HeaderName:  pulumi.String("someheader"),
// 											ExactMatch:  pulumi.String("match this exactly"),
// 											InvertMatch: pulumi.Bool(true),
// 										},
// 									},
// 									IgnoreCase: pulumi.Bool(true),
// 									MetadataFilters: compute.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArray{
// 										&compute.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs{
// 											FilterMatchCriteria: pulumi.String("MATCH_ANY"),
// 											FilterLabels: compute.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArray{
// 												&compute.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs{
// 													Name:  pulumi.String("PLANET"),
// 													Value: pulumi.String("MARS"),
// 												},
// 											},
// 										},
// 									},
// 									QueryParameterMatches: compute.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArray{
// 										&compute.URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs{
// 											Name:         pulumi.String("a query parameter"),
// 											PresentMatch: pulumi.Bool(true),
// 										},
// 									},
// 								},
// 							},
// 							UrlRedirect: &compute.URLMapPathMatcherRouteRuleUrlRedirectArgs{
// 								HostRedirect:         pulumi.String("A host"),
// 								HttpsRedirect:        pulumi.Bool(false),
// 								PathRedirect:         pulumi.String("some/path"),
// 								RedirectResponseCode: pulumi.String("TEMPORARY_REDIRECT"),
// 								StripQuery:           pulumi.Bool(true),
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Tests: compute.URLMapTestArray{
// 				&compute.URLMapTestArgs{
// 					Service: home.ID(),
// 					Host:    pulumi.String("hi.com"),
// 					Path:    pulumi.String("/home"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewTargetGrpcProxy(ctx, "defaultTargetGrpcProxy", &compute.TargetGrpcProxyArgs{
// 			UrlMap:               urlmap.ID(),
// 			ValidateForProxyless: pulumi.Bool(true),
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
// TargetGrpcProxy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default projects/{{project}}/global/targetGrpcProxies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{name}}
// ```
type TargetGrpcProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrOutput `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrOutput `pulumi:"validateForProxyless"`
}

// NewTargetGrpcProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetGrpcProxy(ctx *pulumi.Context,
	name string, args *TargetGrpcProxyArgs, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	if args == nil {
		args = &TargetGrpcProxyArgs{}
	}

	var resource TargetGrpcProxy
	err := ctx.RegisterResource("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGrpcProxy gets an existing TargetGrpcProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGrpcProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGrpcProxyState, opts ...pulumi.ResourceOption) (*TargetGrpcProxy, error) {
	var resource TargetGrpcProxy
	err := ctx.ReadResource("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGrpcProxy resources.
type targetGrpcProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint *string `pulumi:"fingerprint"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL with id for the resource.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap *string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless *bool `pulumi:"validateForProxyless"`
}

type TargetGrpcProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// This field will be ignored when inserting a TargetGrpcProxy. An up-to-date fingerprint must be provided in order to
	// patch/update the TargetGrpcProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest
	// fingerprint, make a get() request to retrieve the TargetGrpcProxy. A base64-encoded string.
	Fingerprint pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL with id for the resource.
	SelfLinkWithId pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrInput
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrInput
}

func (TargetGrpcProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyState)(nil)).Elem()
}

type targetGrpcProxyArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap *string `pulumi:"urlMap"`
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless *bool `pulumi:"validateForProxyless"`
}

// The set of arguments for constructing a TargetGrpcProxy resource.
type TargetGrpcProxyArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource
	// is created. The name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long
	// and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// URL to the UrlMap resource that defines the mapping from URL to
	// the BackendService. The protocol field in the BackendService
	// must be set to GRPC.
	UrlMap pulumi.StringPtrInput
	// If true, indicates that the BackendServices referenced by
	// the urlMap may be accessed by gRPC applications without using
	// a sidecar proxy. This will enable configuration checks on urlMap
	// and its referenced BackendServices to not allow unsupported features.
	// A gRPC application must use "xds:///" scheme in the target URI
	// of the service it is connecting to. If false, indicates that the
	// BackendServices referenced by the urlMap will be accessed by gRPC
	// applications via a sidecar proxy. In this case, a gRPC application
	// must not use "xds:///" scheme in the target URI of the service
	// it is connecting to
	ValidateForProxyless pulumi.BoolPtrInput
}

func (TargetGrpcProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGrpcProxyArgs)(nil)).Elem()
}

type TargetGrpcProxyInput interface {
	pulumi.Input

	ToTargetGrpcProxyOutput() TargetGrpcProxyOutput
	ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput
}

func (*TargetGrpcProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxy)(nil))
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return i.ToTargetGrpcProxyOutputWithContext(context.Background())
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyOutput)
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyPtrOutput() TargetGrpcProxyPtrOutput {
	return i.ToTargetGrpcProxyPtrOutputWithContext(context.Background())
}

func (i *TargetGrpcProxy) ToTargetGrpcProxyPtrOutputWithContext(ctx context.Context) TargetGrpcProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyPtrOutput)
}

type TargetGrpcProxyPtrInput interface {
	pulumi.Input

	ToTargetGrpcProxyPtrOutput() TargetGrpcProxyPtrOutput
	ToTargetGrpcProxyPtrOutputWithContext(ctx context.Context) TargetGrpcProxyPtrOutput
}

type targetGrpcProxyPtrType TargetGrpcProxyArgs

func (*targetGrpcProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGrpcProxy)(nil))
}

func (i *targetGrpcProxyPtrType) ToTargetGrpcProxyPtrOutput() TargetGrpcProxyPtrOutput {
	return i.ToTargetGrpcProxyPtrOutputWithContext(context.Background())
}

func (i *targetGrpcProxyPtrType) ToTargetGrpcProxyPtrOutputWithContext(ctx context.Context) TargetGrpcProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyPtrOutput)
}

// TargetGrpcProxyArrayInput is an input type that accepts TargetGrpcProxyArray and TargetGrpcProxyArrayOutput values.
// You can construct a concrete instance of `TargetGrpcProxyArrayInput` via:
//
//          TargetGrpcProxyArray{ TargetGrpcProxyArgs{...} }
type TargetGrpcProxyArrayInput interface {
	pulumi.Input

	ToTargetGrpcProxyArrayOutput() TargetGrpcProxyArrayOutput
	ToTargetGrpcProxyArrayOutputWithContext(context.Context) TargetGrpcProxyArrayOutput
}

type TargetGrpcProxyArray []TargetGrpcProxyInput

func (TargetGrpcProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetGrpcProxy)(nil)).Elem()
}

func (i TargetGrpcProxyArray) ToTargetGrpcProxyArrayOutput() TargetGrpcProxyArrayOutput {
	return i.ToTargetGrpcProxyArrayOutputWithContext(context.Background())
}

func (i TargetGrpcProxyArray) ToTargetGrpcProxyArrayOutputWithContext(ctx context.Context) TargetGrpcProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyArrayOutput)
}

// TargetGrpcProxyMapInput is an input type that accepts TargetGrpcProxyMap and TargetGrpcProxyMapOutput values.
// You can construct a concrete instance of `TargetGrpcProxyMapInput` via:
//
//          TargetGrpcProxyMap{ "key": TargetGrpcProxyArgs{...} }
type TargetGrpcProxyMapInput interface {
	pulumi.Input

	ToTargetGrpcProxyMapOutput() TargetGrpcProxyMapOutput
	ToTargetGrpcProxyMapOutputWithContext(context.Context) TargetGrpcProxyMapOutput
}

type TargetGrpcProxyMap map[string]TargetGrpcProxyInput

func (TargetGrpcProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetGrpcProxy)(nil)).Elem()
}

func (i TargetGrpcProxyMap) ToTargetGrpcProxyMapOutput() TargetGrpcProxyMapOutput {
	return i.ToTargetGrpcProxyMapOutputWithContext(context.Background())
}

func (i TargetGrpcProxyMap) ToTargetGrpcProxyMapOutputWithContext(ctx context.Context) TargetGrpcProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGrpcProxyMapOutput)
}

type TargetGrpcProxyOutput struct{ *pulumi.OutputState }

func (TargetGrpcProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetGrpcProxy)(nil))
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutput() TargetGrpcProxyOutput {
	return o
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyOutputWithContext(ctx context.Context) TargetGrpcProxyOutput {
	return o
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyPtrOutput() TargetGrpcProxyPtrOutput {
	return o.ToTargetGrpcProxyPtrOutputWithContext(context.Background())
}

func (o TargetGrpcProxyOutput) ToTargetGrpcProxyPtrOutputWithContext(ctx context.Context) TargetGrpcProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetGrpcProxy) *TargetGrpcProxy {
		return &v
	}).(TargetGrpcProxyPtrOutput)
}

type TargetGrpcProxyPtrOutput struct{ *pulumi.OutputState }

func (TargetGrpcProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGrpcProxy)(nil))
}

func (o TargetGrpcProxyPtrOutput) ToTargetGrpcProxyPtrOutput() TargetGrpcProxyPtrOutput {
	return o
}

func (o TargetGrpcProxyPtrOutput) ToTargetGrpcProxyPtrOutputWithContext(ctx context.Context) TargetGrpcProxyPtrOutput {
	return o
}

func (o TargetGrpcProxyPtrOutput) Elem() TargetGrpcProxyOutput {
	return o.ApplyT(func(v *TargetGrpcProxy) TargetGrpcProxy {
		if v != nil {
			return *v
		}
		var ret TargetGrpcProxy
		return ret
	}).(TargetGrpcProxyOutput)
}

type TargetGrpcProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetGrpcProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetGrpcProxy)(nil))
}

func (o TargetGrpcProxyArrayOutput) ToTargetGrpcProxyArrayOutput() TargetGrpcProxyArrayOutput {
	return o
}

func (o TargetGrpcProxyArrayOutput) ToTargetGrpcProxyArrayOutputWithContext(ctx context.Context) TargetGrpcProxyArrayOutput {
	return o
}

func (o TargetGrpcProxyArrayOutput) Index(i pulumi.IntInput) TargetGrpcProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetGrpcProxy {
		return vs[0].([]TargetGrpcProxy)[vs[1].(int)]
	}).(TargetGrpcProxyOutput)
}

type TargetGrpcProxyMapOutput struct{ *pulumi.OutputState }

func (TargetGrpcProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TargetGrpcProxy)(nil))
}

func (o TargetGrpcProxyMapOutput) ToTargetGrpcProxyMapOutput() TargetGrpcProxyMapOutput {
	return o
}

func (o TargetGrpcProxyMapOutput) ToTargetGrpcProxyMapOutputWithContext(ctx context.Context) TargetGrpcProxyMapOutput {
	return o
}

func (o TargetGrpcProxyMapOutput) MapIndex(k pulumi.StringInput) TargetGrpcProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TargetGrpcProxy {
		return vs[0].(map[string]TargetGrpcProxy)[vs[1].(string)]
	}).(TargetGrpcProxyOutput)
}

func init() {
	pulumi.RegisterOutputType(TargetGrpcProxyOutput{})
	pulumi.RegisterOutputType(TargetGrpcProxyPtrOutput{})
	pulumi.RegisterOutputType(TargetGrpcProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetGrpcProxyMapOutput{})
}
