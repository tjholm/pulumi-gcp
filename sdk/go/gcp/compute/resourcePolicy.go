// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A policy that can be attached to a resource to specify or schedule actions on that resource.
//
// ## Example Usage
// ### Resource Policy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewResourcePolicy(ctx, "foo", &compute.ResourcePolicyArgs{
// 			Region: pulumi.String("us-central1"),
// 			SnapshotSchedulePolicy: &compute.ResourcePolicySnapshotSchedulePolicyArgs{
// 				Schedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs{
// 					DailySchedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs{
// 						DaysInCycle: pulumi.Int(1),
// 						StartTime:   pulumi.String("04:00"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Resource Policy Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewResourcePolicy(ctx, "bar", &compute.ResourcePolicyArgs{
// 			Region: pulumi.String("us-central1"),
// 			SnapshotSchedulePolicy: &compute.ResourcePolicySnapshotSchedulePolicyArgs{
// 				RetentionPolicy: &compute.ResourcePolicySnapshotSchedulePolicyRetentionPolicyArgs{
// 					MaxRetentionDays:   pulumi.Int(10),
// 					OnSourceDiskDelete: pulumi.String("KEEP_AUTO_SNAPSHOTS"),
// 				},
// 				Schedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleArgs{
// 					HourlySchedule: &compute.ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleArgs{
// 						HoursInCycle: pulumi.Int(20),
// 						StartTime:    pulumi.String("23:00"),
// 					},
// 				},
// 				SnapshotProperties: &compute.ResourcePolicySnapshotSchedulePolicySnapshotPropertiesArgs{
// 					GuestFlush: pulumi.Bool(true),
// 					Labels: pulumi.StringMap{
// 						"myLabel": pulumi.String("value"),
// 					},
// 					StorageLocations: pulumi.String("us"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Resource Policy Placement Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewResourcePolicy(ctx, "baz", &compute.ResourcePolicyArgs{
// 			GroupPlacementPolicy: &compute.ResourcePolicyGroupPlacementPolicyArgs{
// 				Collocation: pulumi.String("COLLOCATED"),
// 				VmCount:     pulumi.Int(2),
// 			},
// 			Region: pulumi.String("us-central1"),
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
// ResourcePolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/resourcePolicy:ResourcePolicy default {{name}}
// ```
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrOutput `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where resource policy resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrOutput `pulumi:"snapshotSchedulePolicy"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		args = &ResourcePolicyArgs{}
	}

	var resource ResourcePolicy
	err := ctx.RegisterResource("gcp:compute/resourcePolicy:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("gcp:compute/resourcePolicy:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where resource policy resides.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicy"`
}

type ResourcePolicyState struct {
	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrInput
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where resource policy resides.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicy"`
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where resource policy resides.
	Region *string `pulumi:"region"`
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicy"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrInput
	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z`? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where resource policy resides.
	Region pulumi.StringPtrInput
	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

func (i *ResourcePolicy) ToResourcePolicyPtrOutput() ResourcePolicyPtrOutput {
	return i.ToResourcePolicyPtrOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyPtrOutputWithContext(ctx context.Context) ResourcePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyPtrOutput)
}

type ResourcePolicyPtrInput interface {
	pulumi.Input

	ToResourcePolicyPtrOutput() ResourcePolicyPtrOutput
	ToResourcePolicyPtrOutputWithContext(ctx context.Context) ResourcePolicyPtrOutput
}

type resourcePolicyPtrType ResourcePolicyArgs

func (*resourcePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil))
}

func (i *resourcePolicyPtrType) ToResourcePolicyPtrOutput() ResourcePolicyPtrOutput {
	return i.ToResourcePolicyPtrOutputWithContext(context.Background())
}

func (i *resourcePolicyPtrType) ToResourcePolicyPtrOutputWithContext(ctx context.Context) ResourcePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyPtrOutput)
}

// ResourcePolicyArrayInput is an input type that accepts ResourcePolicyArray and ResourcePolicyArrayOutput values.
// You can construct a concrete instance of `ResourcePolicyArrayInput` via:
//
//          ResourcePolicyArray{ ResourcePolicyArgs{...} }
type ResourcePolicyArrayInput interface {
	pulumi.Input

	ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput
	ToResourcePolicyArrayOutputWithContext(context.Context) ResourcePolicyArrayOutput
}

type ResourcePolicyArray []ResourcePolicyInput

func (ResourcePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResourcePolicy)(nil))
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return i.ToResourcePolicyArrayOutputWithContext(context.Background())
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyArrayOutput)
}

// ResourcePolicyMapInput is an input type that accepts ResourcePolicyMap and ResourcePolicyMapOutput values.
// You can construct a concrete instance of `ResourcePolicyMapInput` via:
//
//          ResourcePolicyMap{ "key": ResourcePolicyArgs{...} }
type ResourcePolicyMapInput interface {
	pulumi.Input

	ToResourcePolicyMapOutput() ResourcePolicyMapOutput
	ToResourcePolicyMapOutputWithContext(context.Context) ResourcePolicyMapOutput
}

type ResourcePolicyMap map[string]ResourcePolicyInput

func (ResourcePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResourcePolicy)(nil))
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return i.ToResourcePolicyMapOutputWithContext(context.Background())
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyMapOutput)
}

type ResourcePolicyOutput struct {
	*pulumi.OutputState
}

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyPtrOutput() ResourcePolicyPtrOutput {
	return o.ToResourcePolicyPtrOutputWithContext(context.Background())
}

func (o ResourcePolicyOutput) ToResourcePolicyPtrOutputWithContext(ctx context.Context) ResourcePolicyPtrOutput {
	return o.ApplyT(func(v ResourcePolicy) *ResourcePolicy {
		return &v
	}).(ResourcePolicyPtrOutput)
}

type ResourcePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ResourcePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil))
}

func (o ResourcePolicyPtrOutput) ToResourcePolicyPtrOutput() ResourcePolicyPtrOutput {
	return o
}

func (o ResourcePolicyPtrOutput) ToResourcePolicyPtrOutputWithContext(ctx context.Context) ResourcePolicyPtrOutput {
	return o
}

type ResourcePolicyArrayOutput struct{ *pulumi.OutputState }

func (ResourcePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourcePolicy)(nil))
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) Index(i pulumi.IntInput) ResourcePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourcePolicy {
		return vs[0].([]ResourcePolicy)[vs[1].(int)]
	}).(ResourcePolicyOutput)
}

type ResourcePolicyMapOutput struct{ *pulumi.OutputState }

func (ResourcePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourcePolicy)(nil))
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) MapIndex(k pulumi.StringInput) ResourcePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourcePolicy {
		return vs[0].(map[string]ResourcePolicy)[vs[1].(string)]
	}).(ResourcePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
	pulumi.RegisterOutputType(ResourcePolicyPtrOutput{})
	pulumi.RegisterOutputType(ResourcePolicyArrayOutput{})
	pulumi.RegisterOutputType(ResourcePolicyMapOutput{})
}
