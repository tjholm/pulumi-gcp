// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// UrlMaps are used to route requests to a backend service based on rules
// that you define for the host and path of an incoming URL.
//
// To get more information about UrlMap, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps)
//
// ## Example Usage
//
// ## Import
//
// UrlMap can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/uRLMap:URLMap default projects/{{project}}/global/urlMaps/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/uRLMap:URLMap default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/uRLMap:URLMap default {{name}}
// ```
type URLMap struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrOutput `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrOutput `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrOutput `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrOutput `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayOutput `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId pulumi.IntOutput `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayOutput `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayOutput `pulumi:"tests"`
}

// NewURLMap registers a new resource with the given unique name, arguments, and options.
func NewURLMap(ctx *pulumi.Context,
	name string, args *URLMapArgs, opts ...pulumi.ResourceOption) (*URLMap, error) {
	if args == nil {
		args = &URLMapArgs{}
	}

	var resource URLMap
	err := ctx.RegisterResource("gcp:compute/uRLMap:URLMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURLMap gets an existing URLMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetURLMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *URLMapState, opts ...pulumi.ResourceOption) (*URLMap, error) {
	var resource URLMap
	err := ctx.ReadResource("gcp:compute/uRLMap:URLMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering URLMap resources.
type urlmapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction *URLMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *URLMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint *string `pulumi:"fingerprint"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction *URLMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []URLMapHostRule `pulumi:"hostRules"`
	// The unique identifier for the resource.
	MapId *int `pulumi:"mapId"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers []URLMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests []URLMapTest `pulumi:"tests"`
}

type URLMapState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	Fingerprint pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayInput
	// The unique identifier for the resource.
	MapId pulumi.IntPtrInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayInput
}

func (URLMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlmapState)(nil)).Elem()
}

type urlmapArgs struct {
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction *URLMapDefaultRouteAction `pulumi:"defaultRouteAction"`
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService *string `pulumi:"defaultService"`
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect *URLMapDefaultUrlRedirect `pulumi:"defaultUrlRedirect"`
	// Description of this test case.
	Description *string `pulumi:"description"`
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction *URLMapHeaderAction `pulumi:"headerAction"`
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules []URLMapHostRule `pulumi:"hostRules"`
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name *string `pulumi:"name"`
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers []URLMapPathMatcher `pulumi:"pathMatchers"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests []URLMapTest `pulumi:"tests"`
}

// The set of arguments for constructing a URLMap resource.
type URLMapArgs struct {
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
	// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
	// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
	// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
	// Only one of defaultRouteAction or defaultUrlRedirect must be set.
	// Structure is documented below.
	DefaultRouteAction URLMapDefaultRouteActionPtrInput
	// The backend service or backend bucket to use when none of the given paths match.
	DefaultService pulumi.StringPtrInput
	// When none of the specified hostRules match, the request is redirected to a URL specified
	// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
	// defaultRouteAction must not be set.
	// Structure is documented below.
	DefaultUrlRedirect URLMapDefaultUrlRedirectPtrInput
	// Description of this test case.
	Description pulumi.StringPtrInput
	// Specifies changes to request and response headers that need to take effect for
	// the selected backendService.
	// headerAction specified here take effect before headerAction in the enclosing
	// HttpRouteRule, PathMatcher and UrlMap.
	// Structure is documented below.
	HeaderAction URLMapHeaderActionPtrInput
	// The list of HostRules to use against the URL.
	// Structure is documented below.
	HostRules URLMapHostRuleArrayInput
	// The name of the query parameter to match. The query parameter must exist in the
	// request, in the absence of which the request match fails.
	Name pulumi.StringPtrInput
	// The name of the PathMatcher to use to match the path portion of the URL if the
	// hostRule matches the URL's host portion.
	PathMatchers URLMapPathMatcherArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The list of expected URL mapping tests. Request to update this UrlMap will
	// succeed only if all of the test cases pass. You can specify a maximum of 100
	// tests per UrlMap.
	// Structure is documented below.
	Tests URLMapTestArrayInput
}

func (URLMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlmapArgs)(nil)).Elem()
}

type URLMapInput interface {
	pulumi.Input

	ToURLMapOutput() URLMapOutput
	ToURLMapOutputWithContext(ctx context.Context) URLMapOutput
}

func (*URLMap) ElementType() reflect.Type {
	return reflect.TypeOf((**URLMap)(nil)).Elem()
}

func (i *URLMap) ToURLMapOutput() URLMapOutput {
	return i.ToURLMapOutputWithContext(context.Background())
}

func (i *URLMap) ToURLMapOutputWithContext(ctx context.Context) URLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapOutput)
}

// URLMapArrayInput is an input type that accepts URLMapArray and URLMapArrayOutput values.
// You can construct a concrete instance of `URLMapArrayInput` via:
//
//          URLMapArray{ URLMapArgs{...} }
type URLMapArrayInput interface {
	pulumi.Input

	ToURLMapArrayOutput() URLMapArrayOutput
	ToURLMapArrayOutputWithContext(context.Context) URLMapArrayOutput
}

type URLMapArray []URLMapInput

func (URLMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*URLMap)(nil)).Elem()
}

func (i URLMapArray) ToURLMapArrayOutput() URLMapArrayOutput {
	return i.ToURLMapArrayOutputWithContext(context.Background())
}

func (i URLMapArray) ToURLMapArrayOutputWithContext(ctx context.Context) URLMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapArrayOutput)
}

// URLMapMapInput is an input type that accepts URLMapMap and URLMapMapOutput values.
// You can construct a concrete instance of `URLMapMapInput` via:
//
//          URLMapMap{ "key": URLMapArgs{...} }
type URLMapMapInput interface {
	pulumi.Input

	ToURLMapMapOutput() URLMapMapOutput
	ToURLMapMapOutputWithContext(context.Context) URLMapMapOutput
}

type URLMapMap map[string]URLMapInput

func (URLMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*URLMap)(nil)).Elem()
}

func (i URLMapMap) ToURLMapMapOutput() URLMapMapOutput {
	return i.ToURLMapMapOutputWithContext(context.Background())
}

func (i URLMapMap) ToURLMapMapOutputWithContext(ctx context.Context) URLMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(URLMapMapOutput)
}

type URLMapOutput struct{ *pulumi.OutputState }

func (URLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**URLMap)(nil)).Elem()
}

func (o URLMapOutput) ToURLMapOutput() URLMapOutput {
	return o
}

func (o URLMapOutput) ToURLMapOutputWithContext(ctx context.Context) URLMapOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o URLMapOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
// advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
// to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
// Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
// Only one of defaultRouteAction or defaultUrlRedirect must be set.
// Structure is documented below.
func (o URLMapOutput) DefaultRouteAction() URLMapDefaultRouteActionPtrOutput {
	return o.ApplyT(func(v *URLMap) URLMapDefaultRouteActionPtrOutput { return v.DefaultRouteAction }).(URLMapDefaultRouteActionPtrOutput)
}

// The backend service or backend bucket to use when none of the given paths match.
func (o URLMapOutput) DefaultService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringPtrOutput { return v.DefaultService }).(pulumi.StringPtrOutput)
}

// When none of the specified hostRules match, the request is redirected to a URL specified
// by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
// defaultRouteAction must not be set.
// Structure is documented below.
func (o URLMapOutput) DefaultUrlRedirect() URLMapDefaultUrlRedirectPtrOutput {
	return o.ApplyT(func(v *URLMap) URLMapDefaultUrlRedirectPtrOutput { return v.DefaultUrlRedirect }).(URLMapDefaultUrlRedirectPtrOutput)
}

// Description of this test case.
func (o URLMapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
func (o URLMapOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Specifies changes to request and response headers that need to take effect for
// the selected backendService.
// headerAction specified here take effect before headerAction in the enclosing
// HttpRouteRule, PathMatcher and UrlMap.
// Structure is documented below.
func (o URLMapOutput) HeaderAction() URLMapHeaderActionPtrOutput {
	return o.ApplyT(func(v *URLMap) URLMapHeaderActionPtrOutput { return v.HeaderAction }).(URLMapHeaderActionPtrOutput)
}

// The list of HostRules to use against the URL.
// Structure is documented below.
func (o URLMapOutput) HostRules() URLMapHostRuleArrayOutput {
	return o.ApplyT(func(v *URLMap) URLMapHostRuleArrayOutput { return v.HostRules }).(URLMapHostRuleArrayOutput)
}

// The unique identifier for the resource.
func (o URLMapOutput) MapId() pulumi.IntOutput {
	return o.ApplyT(func(v *URLMap) pulumi.IntOutput { return v.MapId }).(pulumi.IntOutput)
}

// The name of the query parameter to match. The query parameter must exist in the
// request, in the absence of which the request match fails.
func (o URLMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the PathMatcher to use to match the path portion of the URL if the
// hostRule matches the URL's host portion.
func (o URLMapOutput) PathMatchers() URLMapPathMatcherArrayOutput {
	return o.ApplyT(func(v *URLMap) URLMapPathMatcherArrayOutput { return v.PathMatchers }).(URLMapPathMatcherArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o URLMapOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o URLMapOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *URLMap) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The list of expected URL mapping tests. Request to update this UrlMap will
// succeed only if all of the test cases pass. You can specify a maximum of 100
// tests per UrlMap.
// Structure is documented below.
func (o URLMapOutput) Tests() URLMapTestArrayOutput {
	return o.ApplyT(func(v *URLMap) URLMapTestArrayOutput { return v.Tests }).(URLMapTestArrayOutput)
}

type URLMapArrayOutput struct{ *pulumi.OutputState }

func (URLMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*URLMap)(nil)).Elem()
}

func (o URLMapArrayOutput) ToURLMapArrayOutput() URLMapArrayOutput {
	return o
}

func (o URLMapArrayOutput) ToURLMapArrayOutputWithContext(ctx context.Context) URLMapArrayOutput {
	return o
}

func (o URLMapArrayOutput) Index(i pulumi.IntInput) URLMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *URLMap {
		return vs[0].([]*URLMap)[vs[1].(int)]
	}).(URLMapOutput)
}

type URLMapMapOutput struct{ *pulumi.OutputState }

func (URLMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*URLMap)(nil)).Elem()
}

func (o URLMapMapOutput) ToURLMapMapOutput() URLMapMapOutput {
	return o
}

func (o URLMapMapOutput) ToURLMapMapOutputWithContext(ctx context.Context) URLMapMapOutput {
	return o
}

func (o URLMapMapOutput) MapIndex(k pulumi.StringInput) URLMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *URLMap {
		return vs[0].(map[string]*URLMap)[vs[1].(string)]
	}).(URLMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*URLMapInput)(nil)).Elem(), &URLMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*URLMapArrayInput)(nil)).Elem(), URLMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*URLMapMapInput)(nil)).Elem(), URLMapMap{})
	pulumi.RegisterOutputType(URLMapOutput{})
	pulumi.RegisterOutputType(URLMapArrayOutput{})
	pulumi.RegisterOutputType(URLMapMapOutput{})
}
