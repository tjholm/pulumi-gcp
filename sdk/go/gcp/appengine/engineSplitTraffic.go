// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Traffic routing configuration for versions within a single service. Traffic splits define how traffic directed to the service is assigned to versions.
//
// To get more information about ServiceSplitTraffic, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)
//
// ## Example Usage
//
// ## Import
//
// ServiceSplitTraffic can be imported using any of these accepted formats* `apps/{{project}}/services/{{service}}` * `{{project}}/{{service}}` * `{{service}}` When using the `pulumi import` command, ServiceSplitTraffic can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default apps/{{project}}/services/{{service}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default {{project}}/{{service}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:appengine/engineSplitTraffic:EngineSplitTraffic default {{service}}
//
// ```
type EngineSplitTraffic struct {
	pulumi.CustomResourceState

	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrOutput `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the service these settings apply to.
	Service pulumi.StringOutput `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitOutput `pulumi:"split"`
}

// NewEngineSplitTraffic registers a new resource with the given unique name, arguments, and options.
func NewEngineSplitTraffic(ctx *pulumi.Context,
	name string, args *EngineSplitTrafficArgs, opts ...pulumi.ResourceOption) (*EngineSplitTraffic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.Split == nil {
		return nil, errors.New("invalid value for required argument 'Split'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EngineSplitTraffic
	err := ctx.RegisterResource("gcp:appengine/engineSplitTraffic:EngineSplitTraffic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEngineSplitTraffic gets an existing EngineSplitTraffic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEngineSplitTraffic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EngineSplitTrafficState, opts ...pulumi.ResourceOption) (*EngineSplitTraffic, error) {
	var resource EngineSplitTraffic
	err := ctx.ReadResource("gcp:appengine/engineSplitTraffic:EngineSplitTraffic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EngineSplitTraffic resources.
type engineSplitTrafficState struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic *bool `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the service these settings apply to.
	Service *string `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split *EngineSplitTrafficSplit `pulumi:"split"`
}

type EngineSplitTrafficState struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the service these settings apply to.
	Service pulumi.StringPtrInput
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitPtrInput
}

func (EngineSplitTrafficState) ElementType() reflect.Type {
	return reflect.TypeOf((*engineSplitTrafficState)(nil)).Elem()
}

type engineSplitTrafficArgs struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic *bool `pulumi:"migrateTraffic"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The name of the service these settings apply to.
	Service string `pulumi:"service"`
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplit `pulumi:"split"`
}

// The set of arguments for constructing a EngineSplitTraffic resource.
type EngineSplitTrafficArgs struct {
	// If set to true traffic will be migrated to this version.
	MigrateTraffic pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The name of the service these settings apply to.
	Service pulumi.StringInput
	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// Structure is documented below.
	Split EngineSplitTrafficSplitInput
}

func (EngineSplitTrafficArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*engineSplitTrafficArgs)(nil)).Elem()
}

type EngineSplitTrafficInput interface {
	pulumi.Input

	ToEngineSplitTrafficOutput() EngineSplitTrafficOutput
	ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput
}

func (*EngineSplitTraffic) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineSplitTraffic)(nil)).Elem()
}

func (i *EngineSplitTraffic) ToEngineSplitTrafficOutput() EngineSplitTrafficOutput {
	return i.ToEngineSplitTrafficOutputWithContext(context.Background())
}

func (i *EngineSplitTraffic) ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineSplitTrafficOutput)
}

// EngineSplitTrafficArrayInput is an input type that accepts EngineSplitTrafficArray and EngineSplitTrafficArrayOutput values.
// You can construct a concrete instance of `EngineSplitTrafficArrayInput` via:
//
//	EngineSplitTrafficArray{ EngineSplitTrafficArgs{...} }
type EngineSplitTrafficArrayInput interface {
	pulumi.Input

	ToEngineSplitTrafficArrayOutput() EngineSplitTrafficArrayOutput
	ToEngineSplitTrafficArrayOutputWithContext(context.Context) EngineSplitTrafficArrayOutput
}

type EngineSplitTrafficArray []EngineSplitTrafficInput

func (EngineSplitTrafficArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EngineSplitTraffic)(nil)).Elem()
}

func (i EngineSplitTrafficArray) ToEngineSplitTrafficArrayOutput() EngineSplitTrafficArrayOutput {
	return i.ToEngineSplitTrafficArrayOutputWithContext(context.Background())
}

func (i EngineSplitTrafficArray) ToEngineSplitTrafficArrayOutputWithContext(ctx context.Context) EngineSplitTrafficArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineSplitTrafficArrayOutput)
}

// EngineSplitTrafficMapInput is an input type that accepts EngineSplitTrafficMap and EngineSplitTrafficMapOutput values.
// You can construct a concrete instance of `EngineSplitTrafficMapInput` via:
//
//	EngineSplitTrafficMap{ "key": EngineSplitTrafficArgs{...} }
type EngineSplitTrafficMapInput interface {
	pulumi.Input

	ToEngineSplitTrafficMapOutput() EngineSplitTrafficMapOutput
	ToEngineSplitTrafficMapOutputWithContext(context.Context) EngineSplitTrafficMapOutput
}

type EngineSplitTrafficMap map[string]EngineSplitTrafficInput

func (EngineSplitTrafficMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EngineSplitTraffic)(nil)).Elem()
}

func (i EngineSplitTrafficMap) ToEngineSplitTrafficMapOutput() EngineSplitTrafficMapOutput {
	return i.ToEngineSplitTrafficMapOutputWithContext(context.Background())
}

func (i EngineSplitTrafficMap) ToEngineSplitTrafficMapOutputWithContext(ctx context.Context) EngineSplitTrafficMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EngineSplitTrafficMapOutput)
}

type EngineSplitTrafficOutput struct{ *pulumi.OutputState }

func (EngineSplitTrafficOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineSplitTraffic)(nil)).Elem()
}

func (o EngineSplitTrafficOutput) ToEngineSplitTrafficOutput() EngineSplitTrafficOutput {
	return o
}

func (o EngineSplitTrafficOutput) ToEngineSplitTrafficOutputWithContext(ctx context.Context) EngineSplitTrafficOutput {
	return o
}

// If set to true traffic will be migrated to this version.
func (o EngineSplitTrafficOutput) MigrateTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EngineSplitTraffic) pulumi.BoolPtrOutput { return v.MigrateTraffic }).(pulumi.BoolPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o EngineSplitTrafficOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EngineSplitTraffic) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the service these settings apply to.
func (o EngineSplitTrafficOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *EngineSplitTraffic) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
// Structure is documented below.
func (o EngineSplitTrafficOutput) Split() EngineSplitTrafficSplitOutput {
	return o.ApplyT(func(v *EngineSplitTraffic) EngineSplitTrafficSplitOutput { return v.Split }).(EngineSplitTrafficSplitOutput)
}

type EngineSplitTrafficArrayOutput struct{ *pulumi.OutputState }

func (EngineSplitTrafficArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EngineSplitTraffic)(nil)).Elem()
}

func (o EngineSplitTrafficArrayOutput) ToEngineSplitTrafficArrayOutput() EngineSplitTrafficArrayOutput {
	return o
}

func (o EngineSplitTrafficArrayOutput) ToEngineSplitTrafficArrayOutputWithContext(ctx context.Context) EngineSplitTrafficArrayOutput {
	return o
}

func (o EngineSplitTrafficArrayOutput) Index(i pulumi.IntInput) EngineSplitTrafficOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EngineSplitTraffic {
		return vs[0].([]*EngineSplitTraffic)[vs[1].(int)]
	}).(EngineSplitTrafficOutput)
}

type EngineSplitTrafficMapOutput struct{ *pulumi.OutputState }

func (EngineSplitTrafficMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EngineSplitTraffic)(nil)).Elem()
}

func (o EngineSplitTrafficMapOutput) ToEngineSplitTrafficMapOutput() EngineSplitTrafficMapOutput {
	return o
}

func (o EngineSplitTrafficMapOutput) ToEngineSplitTrafficMapOutputWithContext(ctx context.Context) EngineSplitTrafficMapOutput {
	return o
}

func (o EngineSplitTrafficMapOutput) MapIndex(k pulumi.StringInput) EngineSplitTrafficOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EngineSplitTraffic {
		return vs[0].(map[string]*EngineSplitTraffic)[vs[1].(string)]
	}).(EngineSplitTrafficOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EngineSplitTrafficInput)(nil)).Elem(), &EngineSplitTraffic{})
	pulumi.RegisterInputType(reflect.TypeOf((*EngineSplitTrafficArrayInput)(nil)).Elem(), EngineSplitTrafficArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EngineSplitTrafficMapInput)(nil)).Elem(), EngineSplitTrafficMap{})
	pulumi.RegisterOutputType(EngineSplitTrafficOutput{})
	pulumi.RegisterOutputType(EngineSplitTrafficArrayOutput{})
	pulumi.RegisterOutputType(EngineSplitTrafficMapOutput{})
}
