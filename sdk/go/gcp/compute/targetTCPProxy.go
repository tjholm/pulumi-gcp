// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a TargetTcpProxy resource, which is used by one or more
// global forwarding rule to route incoming TCP requests to a Backend
// service.
//
// To get more information about TargetTcpProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
// * How-to Guides
//   - [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
//
// ## Example Usage
//
// ## Import
//
// # TargetTcpProxy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default projects/{{project}}/global/targetTcpProxies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{name}}
//
// ```
type TargetTCPProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
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
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewTargetTCPProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetTCPProxy(ctx *pulumi.Context,
	name string, args *TargetTCPProxyArgs, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	var resource TargetTCPProxy
	err := ctx.RegisterResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetTCPProxy gets an existing TargetTCPProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetTCPProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetTCPProxyState, opts ...pulumi.ResourceOption) (*TargetTCPProxy, error) {
	var resource TargetTCPProxy
	err := ctx.ReadResource("gcp:compute/targetTCPProxy:TargetTCPProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetTCPProxy resources.
type targetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService *string `pulumi:"backendService"`
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
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type TargetTCPProxyState struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringPtrInput
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
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (TargetTCPProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyState)(nil)).Elem()
}

type targetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService string `pulumi:"backendService"`
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
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader *string `pulumi:"proxyHeader"`
}

// The set of arguments for constructing a TargetTCPProxy resource.
type TargetTCPProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService pulumi.StringInput
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
	// Specifies the type of proxy header to append before sending data to
	// the backend.
	// Default value is `NONE`.
	// Possible values are `NONE` and `PROXY_V1`.
	ProxyHeader pulumi.StringPtrInput
}

func (TargetTCPProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetTCPProxyArgs)(nil)).Elem()
}

type TargetTCPProxyInput interface {
	pulumi.Input

	ToTargetTCPProxyOutput() TargetTCPProxyOutput
	ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput
}

func (*TargetTCPProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetTCPProxy)(nil)).Elem()
}

func (i *TargetTCPProxy) ToTargetTCPProxyOutput() TargetTCPProxyOutput {
	return i.ToTargetTCPProxyOutputWithContext(context.Background())
}

func (i *TargetTCPProxy) ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetTCPProxyOutput)
}

// TargetTCPProxyArrayInput is an input type that accepts TargetTCPProxyArray and TargetTCPProxyArrayOutput values.
// You can construct a concrete instance of `TargetTCPProxyArrayInput` via:
//
//	TargetTCPProxyArray{ TargetTCPProxyArgs{...} }
type TargetTCPProxyArrayInput interface {
	pulumi.Input

	ToTargetTCPProxyArrayOutput() TargetTCPProxyArrayOutput
	ToTargetTCPProxyArrayOutputWithContext(context.Context) TargetTCPProxyArrayOutput
}

type TargetTCPProxyArray []TargetTCPProxyInput

func (TargetTCPProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetTCPProxy)(nil)).Elem()
}

func (i TargetTCPProxyArray) ToTargetTCPProxyArrayOutput() TargetTCPProxyArrayOutput {
	return i.ToTargetTCPProxyArrayOutputWithContext(context.Background())
}

func (i TargetTCPProxyArray) ToTargetTCPProxyArrayOutputWithContext(ctx context.Context) TargetTCPProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetTCPProxyArrayOutput)
}

// TargetTCPProxyMapInput is an input type that accepts TargetTCPProxyMap and TargetTCPProxyMapOutput values.
// You can construct a concrete instance of `TargetTCPProxyMapInput` via:
//
//	TargetTCPProxyMap{ "key": TargetTCPProxyArgs{...} }
type TargetTCPProxyMapInput interface {
	pulumi.Input

	ToTargetTCPProxyMapOutput() TargetTCPProxyMapOutput
	ToTargetTCPProxyMapOutputWithContext(context.Context) TargetTCPProxyMapOutput
}

type TargetTCPProxyMap map[string]TargetTCPProxyInput

func (TargetTCPProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetTCPProxy)(nil)).Elem()
}

func (i TargetTCPProxyMap) ToTargetTCPProxyMapOutput() TargetTCPProxyMapOutput {
	return i.ToTargetTCPProxyMapOutputWithContext(context.Background())
}

func (i TargetTCPProxyMap) ToTargetTCPProxyMapOutputWithContext(ctx context.Context) TargetTCPProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetTCPProxyMapOutput)
}

type TargetTCPProxyOutput struct{ *pulumi.OutputState }

func (TargetTCPProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetTCPProxy)(nil)).Elem()
}

func (o TargetTCPProxyOutput) ToTargetTCPProxyOutput() TargetTCPProxyOutput {
	return o
}

func (o TargetTCPProxyOutput) ToTargetTCPProxyOutputWithContext(ctx context.Context) TargetTCPProxyOutput {
	return o
}

// A reference to the BackendService resource.
func (o TargetTCPProxyOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringOutput { return v.BackendService }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o TargetTCPProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o TargetTCPProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o TargetTCPProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TargetTCPProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references
// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o TargetTCPProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// Specifies the type of proxy header to append before sending data to
// the backend.
// Default value is `NONE`.
// Possible values are `NONE` and `PROXY_V1`.
func (o TargetTCPProxyOutput) ProxyHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringPtrOutput { return v.ProxyHeader }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource.
func (o TargetTCPProxyOutput) ProxyId() pulumi.IntOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.IntOutput { return v.ProxyId }).(pulumi.IntOutput)
}

// The URI of the created resource.
func (o TargetTCPProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetTCPProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

type TargetTCPProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetTCPProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetTCPProxy)(nil)).Elem()
}

func (o TargetTCPProxyArrayOutput) ToTargetTCPProxyArrayOutput() TargetTCPProxyArrayOutput {
	return o
}

func (o TargetTCPProxyArrayOutput) ToTargetTCPProxyArrayOutputWithContext(ctx context.Context) TargetTCPProxyArrayOutput {
	return o
}

func (o TargetTCPProxyArrayOutput) Index(i pulumi.IntInput) TargetTCPProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TargetTCPProxy {
		return vs[0].([]*TargetTCPProxy)[vs[1].(int)]
	}).(TargetTCPProxyOutput)
}

type TargetTCPProxyMapOutput struct{ *pulumi.OutputState }

func (TargetTCPProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetTCPProxy)(nil)).Elem()
}

func (o TargetTCPProxyMapOutput) ToTargetTCPProxyMapOutput() TargetTCPProxyMapOutput {
	return o
}

func (o TargetTCPProxyMapOutput) ToTargetTCPProxyMapOutputWithContext(ctx context.Context) TargetTCPProxyMapOutput {
	return o
}

func (o TargetTCPProxyMapOutput) MapIndex(k pulumi.StringInput) TargetTCPProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TargetTCPProxy {
		return vs[0].(map[string]*TargetTCPProxy)[vs[1].(string)]
	}).(TargetTCPProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetTCPProxyInput)(nil)).Elem(), &TargetTCPProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetTCPProxyArrayInput)(nil)).Elem(), TargetTCPProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetTCPProxyMapInput)(nil)).Elem(), TargetTCPProxyMap{})
	pulumi.RegisterOutputType(TargetTCPProxyOutput{})
	pulumi.RegisterOutputType(TargetTCPProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetTCPProxyMapOutput{})
}
