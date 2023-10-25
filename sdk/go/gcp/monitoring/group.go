// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The description of a dynamic collection of monitored resources. Each group
// has a filter that is matched against monitored resources and their
// associated metadata. If a group's filter matches an available monitored
// resource, then that resource is a member of that group.
//
// To get more information about Group, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.groups)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/monitoring/groups/)
//
// ## Example Usage
// ### Monitoring Group Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := monitoring.NewGroup(ctx, "basic", &monitoring.GroupArgs{
//				DisplayName: pulumi.String("tf-test MonitoringGroup"),
//				Filter:      pulumi.String("resource.metadata.region=\"europe-west2\""),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Monitoring Group Subgroup
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			parent, err := monitoring.NewGroup(ctx, "parent", &monitoring.GroupArgs{
//				DisplayName: pulumi.String("tf-test MonitoringParentGroup"),
//				Filter:      pulumi.String("resource.metadata.region=\"europe-west2\""),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = monitoring.NewGroup(ctx, "subgroup", &monitoring.GroupArgs{
//				DisplayName: pulumi.String("tf-test MonitoringSubGroup"),
//				Filter:      pulumi.String("resource.metadata.region=\"europe-west2\""),
//				ParentName:  parent.Name,
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
// Group can be imported using any of these accepted formats:
//
// ```sh
//
//	$ pulumi import gcp:monitoring/group:Group default {{name}}
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	//
	// ***
	Filter pulumi.StringOutput `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrOutput `pulumi:"isCluster"`
	// A unique identifier for this group. The format is
	// "projects/{project_id_or_number}/groups/{group_id}".
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrOutput `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("gcp:monitoring/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("gcp:monitoring/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName *string `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	//
	// ***
	Filter *string `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster *bool `pulumi:"isCluster"`
	// A unique identifier for this group. The format is
	// "projects/{project_id_or_number}/groups/{group_id}".
	Name *string `pulumi:"name"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName *string `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type GroupState struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringPtrInput
	// The filter used to determine which monitored resources
	// belong to this group.
	//
	// ***
	Filter pulumi.StringPtrInput
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrInput
	// A unique identifier for this group. The format is
	// "projects/{project_id_or_number}/groups/{group_id}".
	Name pulumi.StringPtrInput
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName string `pulumi:"displayName"`
	// The filter used to determine which monitored resources
	// belong to this group.
	//
	// ***
	Filter string `pulumi:"filter"`
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster *bool `pulumi:"isCluster"`
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName *string `pulumi:"parentName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A user-assigned name for this group, used only for display
	// purposes.
	DisplayName pulumi.StringInput
	// The filter used to determine which monitored resources
	// belong to this group.
	//
	// ***
	Filter pulumi.StringInput
	// If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on
	// groups that are clusters.
	IsCluster pulumi.BoolPtrInput
	// The name of the group's parent, if it has one. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	ParentName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: i.ToGroupOutputWithContext(ctx).OutputState,
	}
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

func (i GroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*Group] {
	return pulumix.Output[[]*Group]{
		OutputState: i.ToGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

func (i GroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Group] {
	return pulumix.Output[map[string]*Group]{
		OutputState: i.ToGroupMapOutputWithContext(ctx).OutputState,
	}
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToOutput(ctx context.Context) pulumix.Output[*Group] {
	return pulumix.Output[*Group]{
		OutputState: o.OutputState,
	}
}

// A user-assigned name for this group, used only for display
// purposes.
func (o GroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The filter used to determine which monitored resources
// belong to this group.
//
// ***
func (o GroupOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// If true, the members of this group are considered to be a
// cluster. The system can perform additional analysis on
// groups that are clusters.
func (o GroupOutput) IsCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.IsCluster }).(pulumi.BoolPtrOutput)
}

// A unique identifier for this group. The format is
// "projects/{project_id_or_number}/groups/{group_id}".
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the group's parent, if it has one. The format is
// "projects/{project_id_or_number}/groups/{group_id}". For
// groups with no parent, parentName is the empty string, "".
func (o GroupOutput) ParentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.ParentName }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o GroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Group] {
	return pulumix.Output[[]*Group]{
		OutputState: o.OutputState,
	}
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Group] {
	return pulumix.Output[map[string]*Group]{
		OutputState: o.OutputState,
	}
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
