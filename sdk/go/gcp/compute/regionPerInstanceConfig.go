// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A config defined for a single managed instance that belongs to an instance group manager. It preserves the instance name
// across instance group manager operations and can define stateful disks or metadata that are unique to the instance.
// This resource works with regional instance group managers.
//
// To get more information about RegionPerInstanceConfig, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/stateful-migs#per-instance_configs)
type RegionPerInstanceConfig struct {
	pulumi.CustomResourceState

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
	// The preserved state for this instance.  Structure is documented below.
	PreservedState RegionPerInstanceConfigPreservedStatePtrOutput `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the containing instance group manager is located
	Region pulumi.StringOutput `pulumi:"region"`
	// The region instance group manager this instance config is part of.
	RegionInstanceGroupManager pulumi.StringOutput `pulumi:"regionInstanceGroupManager"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrOutput `pulumi:"removeInstanceStateOnDestroy"`
}

// NewRegionPerInstanceConfig registers a new resource with the given unique name, arguments, and options.
func NewRegionPerInstanceConfig(ctx *pulumi.Context,
	name string, args *RegionPerInstanceConfigArgs, opts ...pulumi.ResourceOption) (*RegionPerInstanceConfig, error) {
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.RegionInstanceGroupManager == nil {
		return nil, errors.New("missing required argument 'RegionInstanceGroupManager'")
	}
	if args == nil {
		args = &RegionPerInstanceConfigArgs{}
	}
	var resource RegionPerInstanceConfig
	err := ctx.RegisterResource("gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionPerInstanceConfig gets an existing RegionPerInstanceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionPerInstanceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionPerInstanceConfigState, opts ...pulumi.ResourceOption) (*RegionPerInstanceConfig, error) {
	var resource RegionPerInstanceConfig
	err := ctx.ReadResource("gcp:compute/regionPerInstanceConfig:RegionPerInstanceConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionPerInstanceConfig resources.
type regionPerInstanceConfigState struct {
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
	// The preserved state for this instance.  Structure is documented below.
	PreservedState *RegionPerInstanceConfigPreservedState `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the containing instance group manager is located
	Region *string `pulumi:"region"`
	// The region instance group manager this instance config is part of.
	RegionInstanceGroupManager *string `pulumi:"regionInstanceGroupManager"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `pulumi:"removeInstanceStateOnDestroy"`
}

type RegionPerInstanceConfigState struct {
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
	// The preserved state for this instance.  Structure is documented below.
	PreservedState RegionPerInstanceConfigPreservedStatePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the containing instance group manager is located
	Region pulumi.StringPtrInput
	// The region instance group manager this instance config is part of.
	RegionInstanceGroupManager pulumi.StringPtrInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
}

func (RegionPerInstanceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionPerInstanceConfigState)(nil)).Elem()
}

type regionPerInstanceConfigArgs struct {
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
	// The preserved state for this instance.  Structure is documented below.
	PreservedState *RegionPerInstanceConfigPreservedState `pulumi:"preservedState"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the containing instance group manager is located
	Region string `pulumi:"region"`
	// The region instance group manager this instance config is part of.
	RegionInstanceGroupManager string `pulumi:"regionInstanceGroupManager"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `pulumi:"removeInstanceStateOnDestroy"`
}

// The set of arguments for constructing a RegionPerInstanceConfig resource.
type RegionPerInstanceConfigArgs struct {
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
	// The preserved state for this instance.  Structure is documented below.
	PreservedState RegionPerInstanceConfigPreservedStatePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the containing instance group manager is located
	Region pulumi.StringInput
	// The region instance group manager this instance config is part of.
	RegionInstanceGroupManager pulumi.StringInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will *not* immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
}

func (RegionPerInstanceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionPerInstanceConfigArgs)(nil)).Elem()
}
