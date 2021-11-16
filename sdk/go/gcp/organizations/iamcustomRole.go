// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of a customized Cloud IAM organization role. For more information see
// [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
// and
// [API](https://cloud.google.com/iam/reference/rest/v1/organizations.roles).
//
// > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
//  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
//  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
//  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
//  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
//  by the provider, and new roles cannot share that name.
//
// ## Example Usage
//
// This snippet creates a customized IAM organization role.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewIAMCustomRole(ctx, "my_custom_role", &organizations.IAMCustomRoleArgs{
// 			Description: pulumi.String("A description"),
// 			OrgId:       pulumi.String("123456789"),
// 			Permissions: pulumi.StringArray{
// 				pulumi.String("iam.roles.list"),
// 				pulumi.String("iam.roles.create"),
// 				pulumi.String("iam.roles.delete"),
// 			},
// 			RoleId: pulumi.String("myCustomRole"),
// 			Title:  pulumi.String("My Custom Role"),
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
// Customized IAM organization role can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:organizations/iAMCustomRole:IAMCustomRole my-custom-role organizations/123456789/roles/myCustomRole
// ```
type IAMCustomRole struct {
	pulumi.CustomResourceState

	// (Optional) The current deleted state of the role.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// A human-readable description for the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name pulumi.StringOutput `pulumi:"name"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// The role id to use for this role.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrOutput `pulumi:"stage"`
	// A human-readable title for the role.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewIAMCustomRole registers a new resource with the given unique name, arguments, and options.
func NewIAMCustomRole(ctx *pulumi.Context,
	name string, args *IAMCustomRoleArgs, opts ...pulumi.ResourceOption) (*IAMCustomRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource IAMCustomRole
	err := ctx.RegisterResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMCustomRole gets an existing IAMCustomRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMCustomRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMCustomRoleState, opts ...pulumi.ResourceOption) (*IAMCustomRole, error) {
	var resource IAMCustomRole
	err := ctx.ReadResource("gcp:organizations/iAMCustomRole:IAMCustomRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMCustomRole resources.
type iamcustomRoleState struct {
	// (Optional) The current deleted state of the role.
	Deleted *bool `pulumi:"deleted"`
	// A human-readable description for the role.
	Description *string `pulumi:"description"`
	// The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name *string `pulumi:"name"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId *string `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The role id to use for this role.
	RoleId *string `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage *string `pulumi:"stage"`
	// A human-readable title for the role.
	Title *string `pulumi:"title"`
}

type IAMCustomRoleState struct {
	// (Optional) The current deleted state of the role.
	Deleted pulumi.BoolPtrInput
	// A human-readable description for the role.
	Description pulumi.StringPtrInput
	// The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
	Name pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The role id to use for this role.
	RoleId pulumi.StringPtrInput
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrInput
	// A human-readable title for the role.
	Title pulumi.StringPtrInput
}

func (IAMCustomRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamcustomRoleState)(nil)).Elem()
}

type iamcustomRoleArgs struct {
	// A human-readable description for the role.
	Description *string `pulumi:"description"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `pulumi:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `pulumi:"permissions"`
	// The role id to use for this role.
	RoleId string `pulumi:"roleId"`
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage *string `pulumi:"stage"`
	// A human-readable title for the role.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a IAMCustomRole resource.
type IAMCustomRoleArgs struct {
	// A human-readable description for the role.
	Description pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringInput
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions pulumi.StringArrayInput
	// The role id to use for this role.
	RoleId pulumi.StringInput
	// The current launch stage of the role.
	// Defaults to `GA`.
	// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
	Stage pulumi.StringPtrInput
	// A human-readable title for the role.
	Title pulumi.StringInput
}

func (IAMCustomRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamcustomRoleArgs)(nil)).Elem()
}

type IAMCustomRoleInput interface {
	pulumi.Input

	ToIAMCustomRoleOutput() IAMCustomRoleOutput
	ToIAMCustomRoleOutputWithContext(ctx context.Context) IAMCustomRoleOutput
}

func (*IAMCustomRole) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMCustomRole)(nil))
}

func (i *IAMCustomRole) ToIAMCustomRoleOutput() IAMCustomRoleOutput {
	return i.ToIAMCustomRoleOutputWithContext(context.Background())
}

func (i *IAMCustomRole) ToIAMCustomRoleOutputWithContext(ctx context.Context) IAMCustomRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMCustomRoleOutput)
}

func (i *IAMCustomRole) ToIAMCustomRolePtrOutput() IAMCustomRolePtrOutput {
	return i.ToIAMCustomRolePtrOutputWithContext(context.Background())
}

func (i *IAMCustomRole) ToIAMCustomRolePtrOutputWithContext(ctx context.Context) IAMCustomRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMCustomRolePtrOutput)
}

type IAMCustomRolePtrInput interface {
	pulumi.Input

	ToIAMCustomRolePtrOutput() IAMCustomRolePtrOutput
	ToIAMCustomRolePtrOutputWithContext(ctx context.Context) IAMCustomRolePtrOutput
}

type iamcustomRolePtrType IAMCustomRoleArgs

func (*iamcustomRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMCustomRole)(nil))
}

func (i *iamcustomRolePtrType) ToIAMCustomRolePtrOutput() IAMCustomRolePtrOutput {
	return i.ToIAMCustomRolePtrOutputWithContext(context.Background())
}

func (i *iamcustomRolePtrType) ToIAMCustomRolePtrOutputWithContext(ctx context.Context) IAMCustomRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMCustomRolePtrOutput)
}

// IAMCustomRoleArrayInput is an input type that accepts IAMCustomRoleArray and IAMCustomRoleArrayOutput values.
// You can construct a concrete instance of `IAMCustomRoleArrayInput` via:
//
//          IAMCustomRoleArray{ IAMCustomRoleArgs{...} }
type IAMCustomRoleArrayInput interface {
	pulumi.Input

	ToIAMCustomRoleArrayOutput() IAMCustomRoleArrayOutput
	ToIAMCustomRoleArrayOutputWithContext(context.Context) IAMCustomRoleArrayOutput
}

type IAMCustomRoleArray []IAMCustomRoleInput

func (IAMCustomRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMCustomRole)(nil)).Elem()
}

func (i IAMCustomRoleArray) ToIAMCustomRoleArrayOutput() IAMCustomRoleArrayOutput {
	return i.ToIAMCustomRoleArrayOutputWithContext(context.Background())
}

func (i IAMCustomRoleArray) ToIAMCustomRoleArrayOutputWithContext(ctx context.Context) IAMCustomRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMCustomRoleArrayOutput)
}

// IAMCustomRoleMapInput is an input type that accepts IAMCustomRoleMap and IAMCustomRoleMapOutput values.
// You can construct a concrete instance of `IAMCustomRoleMapInput` via:
//
//          IAMCustomRoleMap{ "key": IAMCustomRoleArgs{...} }
type IAMCustomRoleMapInput interface {
	pulumi.Input

	ToIAMCustomRoleMapOutput() IAMCustomRoleMapOutput
	ToIAMCustomRoleMapOutputWithContext(context.Context) IAMCustomRoleMapOutput
}

type IAMCustomRoleMap map[string]IAMCustomRoleInput

func (IAMCustomRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMCustomRole)(nil)).Elem()
}

func (i IAMCustomRoleMap) ToIAMCustomRoleMapOutput() IAMCustomRoleMapOutput {
	return i.ToIAMCustomRoleMapOutputWithContext(context.Background())
}

func (i IAMCustomRoleMap) ToIAMCustomRoleMapOutputWithContext(ctx context.Context) IAMCustomRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMCustomRoleMapOutput)
}

type IAMCustomRoleOutput struct{ *pulumi.OutputState }

func (IAMCustomRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMCustomRole)(nil))
}

func (o IAMCustomRoleOutput) ToIAMCustomRoleOutput() IAMCustomRoleOutput {
	return o
}

func (o IAMCustomRoleOutput) ToIAMCustomRoleOutputWithContext(ctx context.Context) IAMCustomRoleOutput {
	return o
}

func (o IAMCustomRoleOutput) ToIAMCustomRolePtrOutput() IAMCustomRolePtrOutput {
	return o.ToIAMCustomRolePtrOutputWithContext(context.Background())
}

func (o IAMCustomRoleOutput) ToIAMCustomRolePtrOutputWithContext(ctx context.Context) IAMCustomRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IAMCustomRole) *IAMCustomRole {
		return &v
	}).(IAMCustomRolePtrOutput)
}

type IAMCustomRolePtrOutput struct{ *pulumi.OutputState }

func (IAMCustomRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMCustomRole)(nil))
}

func (o IAMCustomRolePtrOutput) ToIAMCustomRolePtrOutput() IAMCustomRolePtrOutput {
	return o
}

func (o IAMCustomRolePtrOutput) ToIAMCustomRolePtrOutputWithContext(ctx context.Context) IAMCustomRolePtrOutput {
	return o
}

func (o IAMCustomRolePtrOutput) Elem() IAMCustomRoleOutput {
	return o.ApplyT(func(v *IAMCustomRole) IAMCustomRole {
		if v != nil {
			return *v
		}
		var ret IAMCustomRole
		return ret
	}).(IAMCustomRoleOutput)
}

type IAMCustomRoleArrayOutput struct{ *pulumi.OutputState }

func (IAMCustomRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IAMCustomRole)(nil))
}

func (o IAMCustomRoleArrayOutput) ToIAMCustomRoleArrayOutput() IAMCustomRoleArrayOutput {
	return o
}

func (o IAMCustomRoleArrayOutput) ToIAMCustomRoleArrayOutputWithContext(ctx context.Context) IAMCustomRoleArrayOutput {
	return o
}

func (o IAMCustomRoleArrayOutput) Index(i pulumi.IntInput) IAMCustomRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IAMCustomRole {
		return vs[0].([]IAMCustomRole)[vs[1].(int)]
	}).(IAMCustomRoleOutput)
}

type IAMCustomRoleMapOutput struct{ *pulumi.OutputState }

func (IAMCustomRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IAMCustomRole)(nil))
}

func (o IAMCustomRoleMapOutput) ToIAMCustomRoleMapOutput() IAMCustomRoleMapOutput {
	return o
}

func (o IAMCustomRoleMapOutput) ToIAMCustomRoleMapOutputWithContext(ctx context.Context) IAMCustomRoleMapOutput {
	return o
}

func (o IAMCustomRoleMapOutput) MapIndex(k pulumi.StringInput) IAMCustomRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IAMCustomRole {
		return vs[0].(map[string]IAMCustomRole)[vs[1].(string)]
	}).(IAMCustomRoleOutput)
}

func init() {
	pulumi.RegisterOutputType(IAMCustomRoleOutput{})
	pulumi.RegisterOutputType(IAMCustomRolePtrOutput{})
	pulumi.RegisterOutputType(IAMCustomRoleArrayOutput{})
	pulumi.RegisterOutputType(IAMCustomRoleMapOutput{})
}
