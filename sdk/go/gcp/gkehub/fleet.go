// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fleet contains information about a group of clusters.
//
// To get more information about Fleet, see:
//
// * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.fleets)
// * How-to Guides
//   - [Registering a Cluster to a Fleet](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
//
// ## Example Usage
// ### Gkehub Fleet Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFleet(ctx, "default", &gkehub.FleetArgs{
//				DefaultClusterConfig: &gkehub.FleetDefaultClusterConfigArgs{
//					SecurityPostureConfig: &gkehub.FleetDefaultClusterConfigSecurityPostureConfigArgs{
//						Mode:              pulumi.String("DISABLED"),
//						VulnerabilityMode: pulumi.String("VULNERABILITY_DISABLED"),
//					},
//				},
//				DisplayName: pulumi.String("my production fleet"),
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
// Fleet can be imported using any of these accepted formats* `projects/{{project}}/locations/global/fleets/default` * `{{project}}` When using the `pulumi import` command, Fleet can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:gkehub/fleet:Fleet default projects/{{project}}/locations/global/fleets/default
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkehub/fleet:Fleet default {{project}}
//
// ```
type Fleet struct {
	pulumi.CustomResourceState

	// The time the fleet was created, in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The default cluster configurations to apply across the fleet.
	// Structure is documented below.
	DefaultClusterConfig FleetDefaultClusterConfigPtrOutput `pulumi:"defaultClusterConfig"`
	// The time the fleet was deleted, in RFC3339 text format.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The state of the fleet resource.
	// Structure is documented below.
	States FleetStateTypeArrayOutput `pulumi:"states"`
	// Google-generated UUID for this resource. This is unique across all
	// Fleet resources. If a Fleet resource is deleted and another
	// resource with the same name is created, it gets a different uid.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time the fleet was last updated, in RFC3339 text format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil {
		args = &FleetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Fleet
	err := ctx.RegisterResource("gcp:gkehub/fleet:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("gcp:gkehub/fleet:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
	// The time the fleet was created, in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// The default cluster configurations to apply across the fleet.
	// Structure is documented below.
	DefaultClusterConfig *FleetDefaultClusterConfig `pulumi:"defaultClusterConfig"`
	// The time the fleet was deleted, in RFC3339 text format.
	DeleteTime *string `pulumi:"deleteTime"`
	// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The state of the fleet resource.
	// Structure is documented below.
	States []FleetStateType `pulumi:"states"`
	// Google-generated UUID for this resource. This is unique across all
	// Fleet resources. If a Fleet resource is deleted and another
	// resource with the same name is created, it gets a different uid.
	Uid *string `pulumi:"uid"`
	// The time the fleet was last updated, in RFC3339 text format.
	UpdateTime *string `pulumi:"updateTime"`
}

type FleetState struct {
	// The time the fleet was created, in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// The default cluster configurations to apply across the fleet.
	// Structure is documented below.
	DefaultClusterConfig FleetDefaultClusterConfigPtrInput
	// The time the fleet was deleted, in RFC3339 text format.
	DeleteTime pulumi.StringPtrInput
	// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The state of the fleet resource.
	// Structure is documented below.
	States FleetStateTypeArrayInput
	// Google-generated UUID for this resource. This is unique across all
	// Fleet resources. If a Fleet resource is deleted and another
	// resource with the same name is created, it gets a different uid.
	Uid pulumi.StringPtrInput
	// The time the fleet was last updated, in RFC3339 text format.
	UpdateTime pulumi.StringPtrInput
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// The default cluster configurations to apply across the fleet.
	// Structure is documented below.
	DefaultClusterConfig *FleetDefaultClusterConfig `pulumi:"defaultClusterConfig"`
	// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// The default cluster configurations to apply across the fleet.
	// Structure is documented below.
	DefaultClusterConfig FleetDefaultClusterConfigPtrInput
	// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
	DisplayName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

type FleetInput interface {
	pulumi.Input

	ToFleetOutput() FleetOutput
	ToFleetOutputWithContext(ctx context.Context) FleetOutput
}

func (*Fleet) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (i *Fleet) ToFleetOutput() FleetOutput {
	return i.ToFleetOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetOutput)
}

// FleetArrayInput is an input type that accepts FleetArray and FleetArrayOutput values.
// You can construct a concrete instance of `FleetArrayInput` via:
//
//	FleetArray{ FleetArgs{...} }
type FleetArrayInput interface {
	pulumi.Input

	ToFleetArrayOutput() FleetArrayOutput
	ToFleetArrayOutputWithContext(context.Context) FleetArrayOutput
}

type FleetArray []FleetInput

func (FleetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fleet)(nil)).Elem()
}

func (i FleetArray) ToFleetArrayOutput() FleetArrayOutput {
	return i.ToFleetArrayOutputWithContext(context.Background())
}

func (i FleetArray) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetArrayOutput)
}

// FleetMapInput is an input type that accepts FleetMap and FleetMapOutput values.
// You can construct a concrete instance of `FleetMapInput` via:
//
//	FleetMap{ "key": FleetArgs{...} }
type FleetMapInput interface {
	pulumi.Input

	ToFleetMapOutput() FleetMapOutput
	ToFleetMapOutputWithContext(context.Context) FleetMapOutput
}

type FleetMap map[string]FleetInput

func (FleetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fleet)(nil)).Elem()
}

func (i FleetMap) ToFleetMapOutput() FleetMapOutput {
	return i.ToFleetMapOutputWithContext(context.Background())
}

func (i FleetMap) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMapOutput)
}

type FleetOutput struct{ *pulumi.OutputState }

func (FleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil)).Elem()
}

func (o FleetOutput) ToFleetOutput() FleetOutput {
	return o
}

func (o FleetOutput) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return o
}

// The time the fleet was created, in RFC3339 text format.
func (o FleetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The default cluster configurations to apply across the fleet.
// Structure is documented below.
func (o FleetOutput) DefaultClusterConfig() FleetDefaultClusterConfigPtrOutput {
	return o.ApplyT(func(v *Fleet) FleetDefaultClusterConfigPtrOutput { return v.DefaultClusterConfig }).(FleetDefaultClusterConfigPtrOutput)
}

// The time the fleet was deleted, in RFC3339 text format.
func (o FleetOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.
// Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.
func (o FleetOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o FleetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The state of the fleet resource.
// Structure is documented below.
func (o FleetOutput) States() FleetStateTypeArrayOutput {
	return o.ApplyT(func(v *Fleet) FleetStateTypeArrayOutput { return v.States }).(FleetStateTypeArrayOutput)
}

// Google-generated UUID for this resource. This is unique across all
// Fleet resources. If a Fleet resource is deleted and another
// resource with the same name is created, it gets a different uid.
func (o FleetOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time the fleet was last updated, in RFC3339 text format.
func (o FleetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Fleet) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type FleetArrayOutput struct{ *pulumi.OutputState }

func (FleetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fleet)(nil)).Elem()
}

func (o FleetArrayOutput) ToFleetArrayOutput() FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) Index(i pulumi.IntInput) FleetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fleet {
		return vs[0].([]*Fleet)[vs[1].(int)]
	}).(FleetOutput)
}

type FleetMapOutput struct{ *pulumi.OutputState }

func (FleetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fleet)(nil)).Elem()
}

func (o FleetMapOutput) ToFleetMapOutput() FleetMapOutput {
	return o
}

func (o FleetMapOutput) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return o
}

func (o FleetMapOutput) MapIndex(k pulumi.StringInput) FleetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fleet {
		return vs[0].(map[string]*Fleet)[vs[1].(string)]
	}).(FleetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FleetInput)(nil)).Elem(), &Fleet{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetArrayInput)(nil)).Elem(), FleetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetMapInput)(nil)).Elem(), FleetMap{})
	pulumi.RegisterOutputType(FleetOutput{})
	pulumi.RegisterOutputType(FleetArrayOutput{})
	pulumi.RegisterOutputType(FleetMapOutput{})
}
