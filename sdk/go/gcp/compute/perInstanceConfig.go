// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A config defined for a single managed instance that belongs to an instance group manager. It preserves the instance name
// across instance group manager operations and can define stateful disks or metadata that are unique to the instance.
//
// To get more information about PerInstanceConfig, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/stateful-migs#per-instance_configs)
//
// ## Example Usage
//
// ## Import
//
// PerInstanceConfig can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{project}}/{{zone}}/{{instance_group_manager}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{zone}}/{{instance_group_manager}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{instance_group_manager}}/{{name}}
// ```
type PerInstanceConfig struct {
	pulumi.CustomResourceState

	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringOutput `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update.
	// Default is `NONE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MinimalAction pulumi.StringPtrOutput `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update.
	// Default is `REPLACE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrOutput `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState PerInstanceConfigPreservedStatePtrOutput `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrOutput `pulumi:"removeInstanceStateOnDestroy"`
	// Zone where the containing instance group manager is located
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewPerInstanceConfig registers a new resource with the given unique name, arguments, and options.
func NewPerInstanceConfig(ctx *pulumi.Context,
	name string, args *PerInstanceConfigArgs, opts ...pulumi.ResourceOption) (*PerInstanceConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceGroupManager == nil {
		return nil, errors.New("invalid value for required argument 'InstanceGroupManager'")
	}
	var resource PerInstanceConfig
	err := ctx.RegisterResource("gcp:compute/perInstanceConfig:PerInstanceConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPerInstanceConfig gets an existing PerInstanceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPerInstanceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PerInstanceConfigState, opts ...pulumi.ResourceOption) (*PerInstanceConfig, error) {
	var resource PerInstanceConfig
	err := ctx.ReadResource("gcp:compute/perInstanceConfig:PerInstanceConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PerInstanceConfig resources.
type perInstanceConfigState struct {
	// The instance group manager this instance config is part of.
	InstanceGroupManager *string `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update.
	// Default is `NONE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MinimalAction *string `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update.
	// Default is `REPLACE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MostDisruptiveAllowedAction *string `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name *string `pulumi:"name"`
	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState *PerInstanceConfigPreservedState `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `pulumi:"removeInstanceStateOnDestroy"`
	// Zone where the containing instance group manager is located
	Zone *string `pulumi:"zone"`
}

type PerInstanceConfigState struct {
	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringPtrInput
	// The minimal action to perform on the instance during an update.
	// Default is `NONE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MinimalAction pulumi.StringPtrInput
	// The most disruptive action to perform on the instance during an update.
	// Default is `REPLACE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrInput
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringPtrInput
	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState PerInstanceConfigPreservedStatePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
	// Zone where the containing instance group manager is located
	Zone pulumi.StringPtrInput
}

func (PerInstanceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*perInstanceConfigState)(nil)).Elem()
}

type perInstanceConfigArgs struct {
	// The instance group manager this instance config is part of.
	InstanceGroupManager string `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update.
	// Default is `NONE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MinimalAction *string `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update.
	// Default is `REPLACE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MostDisruptiveAllowedAction *string `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name *string `pulumi:"name"`
	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState *PerInstanceConfigPreservedState `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `pulumi:"removeInstanceStateOnDestroy"`
	// Zone where the containing instance group manager is located
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a PerInstanceConfig resource.
type PerInstanceConfigArgs struct {
	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringInput
	// The minimal action to perform on the instance during an update.
	// Default is `NONE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MinimalAction pulumi.StringPtrInput
	// The most disruptive action to perform on the instance during an update.
	// Default is `REPLACE`. Possible values are:
	// * REPLACE
	// * RESTART
	// * REFRESH
	// * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrInput
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringPtrInput
	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState PerInstanceConfigPreservedStatePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
	// Zone where the containing instance group manager is located
	Zone pulumi.StringPtrInput
}

func (PerInstanceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*perInstanceConfigArgs)(nil)).Elem()
}

type PerInstanceConfigInput interface {
	pulumi.Input

	ToPerInstanceConfigOutput() PerInstanceConfigOutput
	ToPerInstanceConfigOutputWithContext(ctx context.Context) PerInstanceConfigOutput
}

func (*PerInstanceConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**PerInstanceConfig)(nil)).Elem()
}

func (i *PerInstanceConfig) ToPerInstanceConfigOutput() PerInstanceConfigOutput {
	return i.ToPerInstanceConfigOutputWithContext(context.Background())
}

func (i *PerInstanceConfig) ToPerInstanceConfigOutputWithContext(ctx context.Context) PerInstanceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerInstanceConfigOutput)
}

// PerInstanceConfigArrayInput is an input type that accepts PerInstanceConfigArray and PerInstanceConfigArrayOutput values.
// You can construct a concrete instance of `PerInstanceConfigArrayInput` via:
//
//          PerInstanceConfigArray{ PerInstanceConfigArgs{...} }
type PerInstanceConfigArrayInput interface {
	pulumi.Input

	ToPerInstanceConfigArrayOutput() PerInstanceConfigArrayOutput
	ToPerInstanceConfigArrayOutputWithContext(context.Context) PerInstanceConfigArrayOutput
}

type PerInstanceConfigArray []PerInstanceConfigInput

func (PerInstanceConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PerInstanceConfig)(nil)).Elem()
}

func (i PerInstanceConfigArray) ToPerInstanceConfigArrayOutput() PerInstanceConfigArrayOutput {
	return i.ToPerInstanceConfigArrayOutputWithContext(context.Background())
}

func (i PerInstanceConfigArray) ToPerInstanceConfigArrayOutputWithContext(ctx context.Context) PerInstanceConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerInstanceConfigArrayOutput)
}

// PerInstanceConfigMapInput is an input type that accepts PerInstanceConfigMap and PerInstanceConfigMapOutput values.
// You can construct a concrete instance of `PerInstanceConfigMapInput` via:
//
//          PerInstanceConfigMap{ "key": PerInstanceConfigArgs{...} }
type PerInstanceConfigMapInput interface {
	pulumi.Input

	ToPerInstanceConfigMapOutput() PerInstanceConfigMapOutput
	ToPerInstanceConfigMapOutputWithContext(context.Context) PerInstanceConfigMapOutput
}

type PerInstanceConfigMap map[string]PerInstanceConfigInput

func (PerInstanceConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PerInstanceConfig)(nil)).Elem()
}

func (i PerInstanceConfigMap) ToPerInstanceConfigMapOutput() PerInstanceConfigMapOutput {
	return i.ToPerInstanceConfigMapOutputWithContext(context.Background())
}

func (i PerInstanceConfigMap) ToPerInstanceConfigMapOutputWithContext(ctx context.Context) PerInstanceConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerInstanceConfigMapOutput)
}

type PerInstanceConfigOutput struct{ *pulumi.OutputState }

func (PerInstanceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PerInstanceConfig)(nil)).Elem()
}

func (o PerInstanceConfigOutput) ToPerInstanceConfigOutput() PerInstanceConfigOutput {
	return o
}

func (o PerInstanceConfigOutput) ToPerInstanceConfigOutputWithContext(ctx context.Context) PerInstanceConfigOutput {
	return o
}

type PerInstanceConfigArrayOutput struct{ *pulumi.OutputState }

func (PerInstanceConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PerInstanceConfig)(nil)).Elem()
}

func (o PerInstanceConfigArrayOutput) ToPerInstanceConfigArrayOutput() PerInstanceConfigArrayOutput {
	return o
}

func (o PerInstanceConfigArrayOutput) ToPerInstanceConfigArrayOutputWithContext(ctx context.Context) PerInstanceConfigArrayOutput {
	return o
}

func (o PerInstanceConfigArrayOutput) Index(i pulumi.IntInput) PerInstanceConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PerInstanceConfig {
		return vs[0].([]*PerInstanceConfig)[vs[1].(int)]
	}).(PerInstanceConfigOutput)
}

type PerInstanceConfigMapOutput struct{ *pulumi.OutputState }

func (PerInstanceConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PerInstanceConfig)(nil)).Elem()
}

func (o PerInstanceConfigMapOutput) ToPerInstanceConfigMapOutput() PerInstanceConfigMapOutput {
	return o
}

func (o PerInstanceConfigMapOutput) ToPerInstanceConfigMapOutputWithContext(ctx context.Context) PerInstanceConfigMapOutput {
	return o
}

func (o PerInstanceConfigMapOutput) MapIndex(k pulumi.StringInput) PerInstanceConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PerInstanceConfig {
		return vs[0].(map[string]*PerInstanceConfig)[vs[1].(string)]
	}).(PerInstanceConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PerInstanceConfigInput)(nil)).Elem(), &PerInstanceConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerInstanceConfigArrayInput)(nil)).Elem(), PerInstanceConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PerInstanceConfigMapInput)(nil)).Elem(), PerInstanceConfigMap{})
	pulumi.RegisterOutputType(PerInstanceConfigOutput{})
	pulumi.RegisterOutputType(PerInstanceConfigArrayOutput{})
	pulumi.RegisterOutputType(PerInstanceConfigMapOutput{})
}
