// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a single binding within IAM policy for
// an existing Google Cloud Platform Organization.
//
// > **Note:** This resource __must not__ be used in conjunction with
//    `organizations.IAMMember` for the __same role__ or they will fight over
//    what your policy should be.
//
// > **Note:** On create, this resource will overwrite members of any existing roles.
//     Use `pulumi import` and inspect the `output to ensure
//     your existing members are preserved.
//
// ## Example Usage
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
// 		_, err := organizations.NewIAMBinding(ctx, "binding", &organizations.IAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:alice@gmail.com"),
// 			},
// 			OrgId: pulumi.String("123456789"),
// 			Role:  pulumi.String("roles/browser"),
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
// IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
//
// These bindings can be imported using the `org_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:organizations/iAMBinding:IAMBinding my_org "your-org-id roles/viewer"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IAMBinding struct {
	pulumi.CustomResourceState

	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:organizations/iAMBinding:IAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMBindingState, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	var resource IAMBinding
	err := ctx.ReadResource("gcp:organizations/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag *string `pulumi:"etag"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId *string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMBindingState struct {
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the organization's IAM policy.
	Etag pulumi.StringPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	Condition *IAMBindingCondition `pulumi:"condition"`
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members []string `pulumi:"members"`
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	Condition IAMBindingConditionPtrInput
	// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
	Members pulumi.StringArrayInput
	// The numeric ID of the organization in which you want to create a custom role.
	OrgId pulumi.StringInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}

type IAMBindingInput interface {
	pulumi.Input

	ToIAMBindingOutput() IAMBindingOutput
	ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput
}

func (*IAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMBinding)(nil)).Elem()
}

func (i *IAMBinding) ToIAMBindingOutput() IAMBindingOutput {
	return i.ToIAMBindingOutputWithContext(context.Background())
}

func (i *IAMBinding) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingOutput)
}

// IAMBindingArrayInput is an input type that accepts IAMBindingArray and IAMBindingArrayOutput values.
// You can construct a concrete instance of `IAMBindingArrayInput` via:
//
//          IAMBindingArray{ IAMBindingArgs{...} }
type IAMBindingArrayInput interface {
	pulumi.Input

	ToIAMBindingArrayOutput() IAMBindingArrayOutput
	ToIAMBindingArrayOutputWithContext(context.Context) IAMBindingArrayOutput
}

type IAMBindingArray []IAMBindingInput

func (IAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMBinding)(nil)).Elem()
}

func (i IAMBindingArray) ToIAMBindingArrayOutput() IAMBindingArrayOutput {
	return i.ToIAMBindingArrayOutputWithContext(context.Background())
}

func (i IAMBindingArray) ToIAMBindingArrayOutputWithContext(ctx context.Context) IAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingArrayOutput)
}

// IAMBindingMapInput is an input type that accepts IAMBindingMap and IAMBindingMapOutput values.
// You can construct a concrete instance of `IAMBindingMapInput` via:
//
//          IAMBindingMap{ "key": IAMBindingArgs{...} }
type IAMBindingMapInput interface {
	pulumi.Input

	ToIAMBindingMapOutput() IAMBindingMapOutput
	ToIAMBindingMapOutputWithContext(context.Context) IAMBindingMapOutput
}

type IAMBindingMap map[string]IAMBindingInput

func (IAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMBinding)(nil)).Elem()
}

func (i IAMBindingMap) ToIAMBindingMapOutput() IAMBindingMapOutput {
	return i.ToIAMBindingMapOutputWithContext(context.Background())
}

func (i IAMBindingMap) ToIAMBindingMapOutputWithContext(ctx context.Context) IAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingMapOutput)
}

type IAMBindingOutput struct{ *pulumi.OutputState }

func (IAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMBinding)(nil)).Elem()
}

func (o IAMBindingOutput) ToIAMBindingOutput() IAMBindingOutput {
	return o
}

func (o IAMBindingOutput) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return o
}

type IAMBindingArrayOutput struct{ *pulumi.OutputState }

func (IAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMBinding)(nil)).Elem()
}

func (o IAMBindingArrayOutput) ToIAMBindingArrayOutput() IAMBindingArrayOutput {
	return o
}

func (o IAMBindingArrayOutput) ToIAMBindingArrayOutputWithContext(ctx context.Context) IAMBindingArrayOutput {
	return o
}

func (o IAMBindingArrayOutput) Index(i pulumi.IntInput) IAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMBinding {
		return vs[0].([]*IAMBinding)[vs[1].(int)]
	}).(IAMBindingOutput)
}

type IAMBindingMapOutput struct{ *pulumi.OutputState }

func (IAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMBinding)(nil)).Elem()
}

func (o IAMBindingMapOutput) ToIAMBindingMapOutput() IAMBindingMapOutput {
	return o
}

func (o IAMBindingMapOutput) ToIAMBindingMapOutputWithContext(ctx context.Context) IAMBindingMapOutput {
	return o
}

func (o IAMBindingMapOutput) MapIndex(k pulumi.StringInput) IAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMBinding {
		return vs[0].(map[string]*IAMBinding)[vs[1].(string)]
	}).(IAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingInput)(nil)).Elem(), &IAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingArrayInput)(nil)).Elem(), IAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingMapInput)(nil)).Elem(), IAMBindingMap{})
	pulumi.RegisterOutputType(IAMBindingOutput{})
	pulumi.RegisterOutputType(IAMBindingArrayOutput{})
	pulumi.RegisterOutputType(IAMBindingMapOutput{})
}
