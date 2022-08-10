// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a TargetHttpProxy resource, which is used by one or more global
// forwarding rule to route incoming HTTP requests to a URL map.
//
// To get more information about TargetHttpProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
//
// ## Example Usage
// ### Target Http Proxy Https Redirect
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultURLMap, err := compute.NewURLMap(ctx, "defaultURLMap", &compute.URLMapArgs{
//				DefaultUrlRedirect: &compute.URLMapDefaultUrlRedirectArgs{
//					HttpsRedirect: pulumi.Bool(true),
//					StripQuery:    pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewTargetHttpProxy(ctx, "defaultTargetHttpProxy", &compute.TargetHttpProxyArgs{
//				UrlMap: defaultURLMap.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # TargetHttpProxy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default projects/{{project}}/global/targetHttpProxies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{name}}
//
// ```
type TargetHttpProxy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpProxy(ctx *pulumi.Context,
	name string, args *TargetHttpProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlMap == nil {
		return nil, errors.New("invalid value for required argument 'UrlMap'")
	}
	var resource TargetHttpProxy
	err := ctx.RegisterResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpProxy gets an existing TargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpProxyState, opts ...pulumi.ResourceOption) (*TargetHttpProxy, error) {
	var resource TargetHttpProxy
	err := ctx.ReadResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpProxy resources.
type targetHttpProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyState)(nil)).Elem()
}

type targetHttpProxyArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpProxy resource.
type TargetHttpProxyArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	UrlMap pulumi.StringInput
}

func (TargetHttpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpProxyArgs)(nil)).Elem()
}

type TargetHttpProxyInput interface {
	pulumi.Input

	ToTargetHttpProxyOutput() TargetHttpProxyOutput
	ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput
}

func (*TargetHttpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpProxy)(nil)).Elem()
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return i.ToTargetHttpProxyOutputWithContext(context.Background())
}

func (i *TargetHttpProxy) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpProxyOutput)
}

// TargetHttpProxyArrayInput is an input type that accepts TargetHttpProxyArray and TargetHttpProxyArrayOutput values.
// You can construct a concrete instance of `TargetHttpProxyArrayInput` via:
//
//	TargetHttpProxyArray{ TargetHttpProxyArgs{...} }
type TargetHttpProxyArrayInput interface {
	pulumi.Input

	ToTargetHttpProxyArrayOutput() TargetHttpProxyArrayOutput
	ToTargetHttpProxyArrayOutputWithContext(context.Context) TargetHttpProxyArrayOutput
}

type TargetHttpProxyArray []TargetHttpProxyInput

func (TargetHttpProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetHttpProxy)(nil)).Elem()
}

func (i TargetHttpProxyArray) ToTargetHttpProxyArrayOutput() TargetHttpProxyArrayOutput {
	return i.ToTargetHttpProxyArrayOutputWithContext(context.Background())
}

func (i TargetHttpProxyArray) ToTargetHttpProxyArrayOutputWithContext(ctx context.Context) TargetHttpProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpProxyArrayOutput)
}

// TargetHttpProxyMapInput is an input type that accepts TargetHttpProxyMap and TargetHttpProxyMapOutput values.
// You can construct a concrete instance of `TargetHttpProxyMapInput` via:
//
//	TargetHttpProxyMap{ "key": TargetHttpProxyArgs{...} }
type TargetHttpProxyMapInput interface {
	pulumi.Input

	ToTargetHttpProxyMapOutput() TargetHttpProxyMapOutput
	ToTargetHttpProxyMapOutputWithContext(context.Context) TargetHttpProxyMapOutput
}

type TargetHttpProxyMap map[string]TargetHttpProxyInput

func (TargetHttpProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetHttpProxy)(nil)).Elem()
}

func (i TargetHttpProxyMap) ToTargetHttpProxyMapOutput() TargetHttpProxyMapOutput {
	return i.ToTargetHttpProxyMapOutputWithContext(context.Background())
}

func (i TargetHttpProxyMap) ToTargetHttpProxyMapOutputWithContext(ctx context.Context) TargetHttpProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpProxyMapOutput)
}

type TargetHttpProxyOutput struct{ *pulumi.OutputState }

func (TargetHttpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpProxy)(nil)).Elem()
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutput() TargetHttpProxyOutput {
	return o
}

func (o TargetHttpProxyOutput) ToTargetHttpProxyOutputWithContext(ctx context.Context) TargetHttpProxyOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o TargetHttpProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o TargetHttpProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o TargetHttpProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TargetHttpProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references
// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o TargetHttpProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// The unique identifier for the resource.
func (o TargetHttpProxyOutput) ProxyId() pulumi.IntOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.IntOutput { return v.ProxyId }).(pulumi.IntOutput)
}

// The URI of the created resource.
func (o TargetHttpProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// A reference to the UrlMap resource that defines the mapping from URL
// to the BackendService.
func (o TargetHttpProxyOutput) UrlMap() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpProxy) pulumi.StringOutput { return v.UrlMap }).(pulumi.StringOutput)
}

type TargetHttpProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetHttpProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetHttpProxy)(nil)).Elem()
}

func (o TargetHttpProxyArrayOutput) ToTargetHttpProxyArrayOutput() TargetHttpProxyArrayOutput {
	return o
}

func (o TargetHttpProxyArrayOutput) ToTargetHttpProxyArrayOutputWithContext(ctx context.Context) TargetHttpProxyArrayOutput {
	return o
}

func (o TargetHttpProxyArrayOutput) Index(i pulumi.IntInput) TargetHttpProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TargetHttpProxy {
		return vs[0].([]*TargetHttpProxy)[vs[1].(int)]
	}).(TargetHttpProxyOutput)
}

type TargetHttpProxyMapOutput struct{ *pulumi.OutputState }

func (TargetHttpProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetHttpProxy)(nil)).Elem()
}

func (o TargetHttpProxyMapOutput) ToTargetHttpProxyMapOutput() TargetHttpProxyMapOutput {
	return o
}

func (o TargetHttpProxyMapOutput) ToTargetHttpProxyMapOutputWithContext(ctx context.Context) TargetHttpProxyMapOutput {
	return o
}

func (o TargetHttpProxyMapOutput) MapIndex(k pulumi.StringInput) TargetHttpProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TargetHttpProxy {
		return vs[0].(map[string]*TargetHttpProxy)[vs[1].(string)]
	}).(TargetHttpProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpProxyInput)(nil)).Elem(), &TargetHttpProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpProxyArrayInput)(nil)).Elem(), TargetHttpProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpProxyMapInput)(nil)).Elem(), TargetHttpProxyMap{})
	pulumi.RegisterOutputType(TargetHttpProxyOutput{})
	pulumi.RegisterOutputType(TargetHttpProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetHttpProxyMapOutput{})
}
