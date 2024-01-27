// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ## Import
//
// TlsRoute can be imported using any of these accepted formats* `projects/{{project}}/locations/global/tlsRoutes/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, TlsRoute can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:networkservices/tlsRoute:TlsRoute default projects/{{project}}/locations/global/tlsRoutes/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/tlsRoute:TlsRoute default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/tlsRoute:TlsRoute default {{name}}
//
// ```
type TlsRoute struct {
	pulumi.CustomResourceState

	// Time the TlsRoute was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayOutput `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayOutput `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules TlsRouteRuleArrayOutput `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Time the TlsRoute was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTlsRoute registers a new resource with the given unique name, arguments, and options.
func NewTlsRoute(ctx *pulumi.Context,
	name string, args *TlsRouteArgs, opts ...pulumi.ResourceOption) (*TlsRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TlsRoute
	err := ctx.RegisterResource("gcp:networkservices/tlsRoute:TlsRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTlsRoute gets an existing TlsRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTlsRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TlsRouteState, opts ...pulumi.ResourceOption) (*TlsRoute, error) {
	var resource TlsRoute
	err := ctx.ReadResource("gcp:networkservices/tlsRoute:TlsRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TlsRoute resources.
type tlsRouteState struct {
	// Time the TlsRoute was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules []TlsRouteRule `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// Time the TlsRoute was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type TlsRouteState struct {
	// Time the TlsRoute was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayInput
	// Name of the TlsRoute resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules TlsRouteRuleArrayInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	// Time the TlsRoute was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (TlsRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsRouteState)(nil)).Elem()
}

type tlsRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules []TlsRouteRule `pulumi:"rules"`
}

// The set of arguments for constructing a TlsRoute resource.
type TlsRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayInput
	// Name of the TlsRoute resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules TlsRouteRuleArrayInput
}

func (TlsRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsRouteArgs)(nil)).Elem()
}

type TlsRouteInput interface {
	pulumi.Input

	ToTlsRouteOutput() TlsRouteOutput
	ToTlsRouteOutputWithContext(ctx context.Context) TlsRouteOutput
}

func (*TlsRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsRoute)(nil)).Elem()
}

func (i *TlsRoute) ToTlsRouteOutput() TlsRouteOutput {
	return i.ToTlsRouteOutputWithContext(context.Background())
}

func (i *TlsRoute) ToTlsRouteOutputWithContext(ctx context.Context) TlsRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsRouteOutput)
}

// TlsRouteArrayInput is an input type that accepts TlsRouteArray and TlsRouteArrayOutput values.
// You can construct a concrete instance of `TlsRouteArrayInput` via:
//
//	TlsRouteArray{ TlsRouteArgs{...} }
type TlsRouteArrayInput interface {
	pulumi.Input

	ToTlsRouteArrayOutput() TlsRouteArrayOutput
	ToTlsRouteArrayOutputWithContext(context.Context) TlsRouteArrayOutput
}

type TlsRouteArray []TlsRouteInput

func (TlsRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsRoute)(nil)).Elem()
}

func (i TlsRouteArray) ToTlsRouteArrayOutput() TlsRouteArrayOutput {
	return i.ToTlsRouteArrayOutputWithContext(context.Background())
}

func (i TlsRouteArray) ToTlsRouteArrayOutputWithContext(ctx context.Context) TlsRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsRouteArrayOutput)
}

// TlsRouteMapInput is an input type that accepts TlsRouteMap and TlsRouteMapOutput values.
// You can construct a concrete instance of `TlsRouteMapInput` via:
//
//	TlsRouteMap{ "key": TlsRouteArgs{...} }
type TlsRouteMapInput interface {
	pulumi.Input

	ToTlsRouteMapOutput() TlsRouteMapOutput
	ToTlsRouteMapOutputWithContext(context.Context) TlsRouteMapOutput
}

type TlsRouteMap map[string]TlsRouteInput

func (TlsRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsRoute)(nil)).Elem()
}

func (i TlsRouteMap) ToTlsRouteMapOutput() TlsRouteMapOutput {
	return i.ToTlsRouteMapOutputWithContext(context.Background())
}

func (i TlsRouteMap) ToTlsRouteMapOutputWithContext(ctx context.Context) TlsRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsRouteMapOutput)
}

type TlsRouteOutput struct{ *pulumi.OutputState }

func (TlsRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsRoute)(nil)).Elem()
}

func (o TlsRouteOutput) ToTlsRouteOutput() TlsRouteOutput {
	return o
}

func (o TlsRouteOutput) ToTlsRouteOutputWithContext(ctx context.Context) TlsRouteOutput {
	return o
}

// Time the TlsRoute was created in UTC.
func (o TlsRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o TlsRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
// Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>
func (o TlsRouteOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringArrayOutput { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
// Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
// The attached Mesh should be of a type SIDECAR
func (o TlsRouteOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringArrayOutput { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the TlsRoute resource.
func (o TlsRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TlsRouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Rules that define how traffic is routed and handled.
// Structure is documented below.
func (o TlsRouteOutput) Rules() TlsRouteRuleArrayOutput {
	return o.ApplyT(func(v *TlsRoute) TlsRouteRuleArrayOutput { return v.Rules }).(TlsRouteRuleArrayOutput)
}

// Server-defined URL of this resource.
func (o TlsRouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Time the TlsRoute was updated in UTC.
func (o TlsRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type TlsRouteArrayOutput struct{ *pulumi.OutputState }

func (TlsRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TlsRoute)(nil)).Elem()
}

func (o TlsRouteArrayOutput) ToTlsRouteArrayOutput() TlsRouteArrayOutput {
	return o
}

func (o TlsRouteArrayOutput) ToTlsRouteArrayOutputWithContext(ctx context.Context) TlsRouteArrayOutput {
	return o
}

func (o TlsRouteArrayOutput) Index(i pulumi.IntInput) TlsRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TlsRoute {
		return vs[0].([]*TlsRoute)[vs[1].(int)]
	}).(TlsRouteOutput)
}

type TlsRouteMapOutput struct{ *pulumi.OutputState }

func (TlsRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TlsRoute)(nil)).Elem()
}

func (o TlsRouteMapOutput) ToTlsRouteMapOutput() TlsRouteMapOutput {
	return o
}

func (o TlsRouteMapOutput) ToTlsRouteMapOutputWithContext(ctx context.Context) TlsRouteMapOutput {
	return o
}

func (o TlsRouteMapOutput) MapIndex(k pulumi.StringInput) TlsRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TlsRoute {
		return vs[0].(map[string]*TlsRoute)[vs[1].(string)]
	}).(TlsRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TlsRouteInput)(nil)).Elem(), &TlsRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsRouteArrayInput)(nil)).Elem(), TlsRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TlsRouteMapInput)(nil)).Elem(), TlsRouteMap{})
	pulumi.RegisterOutputType(TlsRouteOutput{})
	pulumi.RegisterOutputType(TlsRouteArrayOutput{})
	pulumi.RegisterOutputType(TlsRouteMapOutput{})
}
