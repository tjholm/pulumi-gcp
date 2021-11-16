// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudidentity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Cloud Identity resource representing a Group.
//
// To get more information about Group, see:
//
// * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/identity/docs/how-to/setup)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
// ### Cloud Identity Groups Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudidentity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudidentity.NewGroup(ctx, "cloudIdentityGroupBasic", &cloudidentity.GroupArgs{
// 			DisplayName: pulumi.String("my-identity-group"),
// 			GroupKey: &cloudidentity.GroupGroupKeyArgs{
// 				Id: pulumi.String("my-identity-group@example.com"),
// 			},
// 			InitialGroupConfig: pulumi.String("WITH_INITIAL_OWNER"),
// 			Labels: pulumi.StringMap{
// 				"cloudidentity.googleapis.com/groups.discussion_forum": pulumi.String(""),
// 			},
// 			Parent: pulumi.String("customers/A01b123xz"),
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
// Group can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudidentity/group:Group default {{name}}
// ```
type Group struct {
	pulumi.CustomResourceState

	// The time when the Group was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An extended description to help users determine the purpose of a Group.
	// Must not be longer than 4,096 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the Group.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// EntityKey of the Group.
	// Structure is documented below.
	GroupKey GroupGroupKeyOutput `pulumi:"groupKey"`
	// The initial configuration options for creating a Group.
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values.
	// Default value is `EMPTY`.
	// Possible values are `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, and `EMPTY`.
	InitialGroupConfig pulumi.StringPtrOutput `pulumi:"initialGroupConfig"`
	// The labels that apply to the Group.
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the entity under which this Group resides in the
	// Cloud Identity resource hierarchy.
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The time when the Group was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupKey == nil {
		return nil, errors.New("invalid value for required argument 'GroupKey'")
	}
	if args.Labels == nil {
		return nil, errors.New("invalid value for required argument 'Labels'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource Group
	err := ctx.RegisterResource("gcp:cloudidentity/group:Group", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:cloudidentity/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The time when the Group was created.
	CreateTime *string `pulumi:"createTime"`
	// An extended description to help users determine the purpose of a Group.
	// Must not be longer than 4,096 characters.
	Description *string `pulumi:"description"`
	// The display name of the Group.
	DisplayName *string `pulumi:"displayName"`
	// EntityKey of the Group.
	// Structure is documented below.
	GroupKey *GroupGroupKey `pulumi:"groupKey"`
	// The initial configuration options for creating a Group.
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values.
	// Default value is `EMPTY`.
	// Possible values are `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, and `EMPTY`.
	InitialGroupConfig *string `pulumi:"initialGroupConfig"`
	// The labels that apply to the Group.
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
	Name *string `pulumi:"name"`
	// The resource name of the entity under which this Group resides in the
	// Cloud Identity resource hierarchy.
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	Parent *string `pulumi:"parent"`
	// The time when the Group was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type GroupState struct {
	// The time when the Group was created.
	CreateTime pulumi.StringPtrInput
	// An extended description to help users determine the purpose of a Group.
	// Must not be longer than 4,096 characters.
	Description pulumi.StringPtrInput
	// The display name of the Group.
	DisplayName pulumi.StringPtrInput
	// EntityKey of the Group.
	// Structure is documented below.
	GroupKey GroupGroupKeyPtrInput
	// The initial configuration options for creating a Group.
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values.
	// Default value is `EMPTY`.
	// Possible values are `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, and `EMPTY`.
	InitialGroupConfig pulumi.StringPtrInput
	// The labels that apply to the Group.
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	Labels pulumi.StringMapInput
	// Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
	Name pulumi.StringPtrInput
	// The resource name of the entity under which this Group resides in the
	// Cloud Identity resource hierarchy.
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	Parent pulumi.StringPtrInput
	// The time when the Group was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// An extended description to help users determine the purpose of a Group.
	// Must not be longer than 4,096 characters.
	Description *string `pulumi:"description"`
	// The display name of the Group.
	DisplayName *string `pulumi:"displayName"`
	// EntityKey of the Group.
	// Structure is documented below.
	GroupKey GroupGroupKey `pulumi:"groupKey"`
	// The initial configuration options for creating a Group.
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values.
	// Default value is `EMPTY`.
	// Possible values are `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, and `EMPTY`.
	InitialGroupConfig *string `pulumi:"initialGroupConfig"`
	// The labels that apply to the Group.
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the entity under which this Group resides in the
	// Cloud Identity resource hierarchy.
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// An extended description to help users determine the purpose of a Group.
	// Must not be longer than 4,096 characters.
	Description pulumi.StringPtrInput
	// The display name of the Group.
	DisplayName pulumi.StringPtrInput
	// EntityKey of the Group.
	// Structure is documented below.
	GroupKey GroupGroupKeyInput
	// The initial configuration options for creating a Group.
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values.
	// Default value is `EMPTY`.
	// Possible values are `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, and `EMPTY`.
	InitialGroupConfig pulumi.StringPtrInput
	// The labels that apply to the Group.
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	Labels pulumi.StringMapInput
	// The resource name of the entity under which this Group resides in the
	// Cloud Identity resource hierarchy.
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	Parent pulumi.StringInput
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
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *Group) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

type GroupPtrInput interface {
	pulumi.Input

	ToGroupPtrOutput() GroupPtrOutput
	ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput
}

type groupPtrType GroupArgs

func (*groupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (i *groupPtrType) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *groupPtrType) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
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

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
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

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o.ToGroupPtrOutputWithContext(context.Background())
}

func (o GroupOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Group) *Group {
		return &v
	}).(GroupPtrOutput)
}

type GroupPtrOutput struct{ *pulumi.OutputState }

func (GroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (o GroupPtrOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) Elem() GroupOutput {
	return o.ApplyT(func(v *Group) Group {
		if v != nil {
			return *v
		}
		var ret Group
		return ret
	}).(GroupOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Group)(nil))
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Group {
		return vs[0].([]Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Group)(nil))
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Group {
		return vs[0].(map[string]Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupPtrOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
