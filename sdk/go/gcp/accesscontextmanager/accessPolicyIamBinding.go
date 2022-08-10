// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Access Context Manager (VPC Service Controls) AccessPolicy. Each of these resources serves a different use case:
//
// * `accesscontextmanager.AccessPolicyIamPolicy`: Authoritative. Sets the IAM policy for the accesspolicy and replaces any existing policy already attached.
// * `accesscontextmanager.AccessPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the accesspolicy are preserved.
// * `accesscontextmanager.AccessPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the accesspolicy are preserved.
//
// > **Note:** `accesscontextmanager.AccessPolicyIamPolicy` **cannot** be used in conjunction with `accesscontextmanager.AccessPolicyIamBinding` and `accesscontextmanager.AccessPolicyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `accesscontextmanager.AccessPolicyIamBinding` resources **can be** used in conjunction with `accesscontextmanager.AccessPolicyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_access\_context\_manager\_access\_policy\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					organizations.GetIAMPolicyBinding{
//						Role: "roles/accesscontextmanager.policyAdmin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = accesscontextmanager.NewAccessPolicyIamPolicy(ctx, "policy", &accesscontextmanager.AccessPolicyIamPolicyArgs{
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## google\_access\_context\_manager\_access\_policy\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.NewAccessPolicyIamBinding(ctx, "binding", &accesscontextmanager.AccessPolicyIamBindingArgs{
//				Role: pulumi.String("roles/accesscontextmanager.policyAdmin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
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
// ## google\_access\_context\_manager\_access\_policy\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accesscontextmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := accesscontextmanager.NewAccessPolicyIamMember(ctx, "member", &accesscontextmanager.AccessPolicyIamMemberArgs{
//				Role:   pulumi.String("roles/accesscontextmanager.policyAdmin"),
//				Member: pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* accessPolicies/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Access Context Manager (VPC Service Controls) accesspolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding editor "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding editor "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding editor accessPolicies/{{access_policy}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AccessPolicyIamBinding struct {
	pulumi.CustomResourceState

	Condition AccessPolicyIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAccessPolicyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicyIamBinding(ctx *pulumi.Context,
	name string, args *AccessPolicyIamBindingArgs, opts ...pulumi.ResourceOption) (*AccessPolicyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AccessPolicyIamBinding
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicyIamBinding gets an existing AccessPolicyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyIamBindingState, opts ...pulumi.ResourceOption) (*AccessPolicyIamBinding, error) {
	var resource AccessPolicyIamBinding
	err := ctx.ReadResource("gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicyIamBinding resources.
type accessPolicyIamBindingState struct {
	Condition *AccessPolicyIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AccessPolicyIamBindingState struct {
	Condition AccessPolicyIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AccessPolicyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyIamBindingState)(nil)).Elem()
}

type accessPolicyIamBindingArgs struct {
	Condition *AccessPolicyIamBindingCondition `pulumi:"condition"`
	Members   []string                         `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AccessPolicyIamBinding resource.
type AccessPolicyIamBindingArgs struct {
	Condition AccessPolicyIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AccessPolicyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyIamBindingArgs)(nil)).Elem()
}

type AccessPolicyIamBindingInput interface {
	pulumi.Input

	ToAccessPolicyIamBindingOutput() AccessPolicyIamBindingOutput
	ToAccessPolicyIamBindingOutputWithContext(ctx context.Context) AccessPolicyIamBindingOutput
}

func (*AccessPolicyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyIamBinding)(nil)).Elem()
}

func (i *AccessPolicyIamBinding) ToAccessPolicyIamBindingOutput() AccessPolicyIamBindingOutput {
	return i.ToAccessPolicyIamBindingOutputWithContext(context.Background())
}

func (i *AccessPolicyIamBinding) ToAccessPolicyIamBindingOutputWithContext(ctx context.Context) AccessPolicyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyIamBindingOutput)
}

// AccessPolicyIamBindingArrayInput is an input type that accepts AccessPolicyIamBindingArray and AccessPolicyIamBindingArrayOutput values.
// You can construct a concrete instance of `AccessPolicyIamBindingArrayInput` via:
//
//	AccessPolicyIamBindingArray{ AccessPolicyIamBindingArgs{...} }
type AccessPolicyIamBindingArrayInput interface {
	pulumi.Input

	ToAccessPolicyIamBindingArrayOutput() AccessPolicyIamBindingArrayOutput
	ToAccessPolicyIamBindingArrayOutputWithContext(context.Context) AccessPolicyIamBindingArrayOutput
}

type AccessPolicyIamBindingArray []AccessPolicyIamBindingInput

func (AccessPolicyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPolicyIamBinding)(nil)).Elem()
}

func (i AccessPolicyIamBindingArray) ToAccessPolicyIamBindingArrayOutput() AccessPolicyIamBindingArrayOutput {
	return i.ToAccessPolicyIamBindingArrayOutputWithContext(context.Background())
}

func (i AccessPolicyIamBindingArray) ToAccessPolicyIamBindingArrayOutputWithContext(ctx context.Context) AccessPolicyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyIamBindingArrayOutput)
}

// AccessPolicyIamBindingMapInput is an input type that accepts AccessPolicyIamBindingMap and AccessPolicyIamBindingMapOutput values.
// You can construct a concrete instance of `AccessPolicyIamBindingMapInput` via:
//
//	AccessPolicyIamBindingMap{ "key": AccessPolicyIamBindingArgs{...} }
type AccessPolicyIamBindingMapInput interface {
	pulumi.Input

	ToAccessPolicyIamBindingMapOutput() AccessPolicyIamBindingMapOutput
	ToAccessPolicyIamBindingMapOutputWithContext(context.Context) AccessPolicyIamBindingMapOutput
}

type AccessPolicyIamBindingMap map[string]AccessPolicyIamBindingInput

func (AccessPolicyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPolicyIamBinding)(nil)).Elem()
}

func (i AccessPolicyIamBindingMap) ToAccessPolicyIamBindingMapOutput() AccessPolicyIamBindingMapOutput {
	return i.ToAccessPolicyIamBindingMapOutputWithContext(context.Background())
}

func (i AccessPolicyIamBindingMap) ToAccessPolicyIamBindingMapOutputWithContext(ctx context.Context) AccessPolicyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyIamBindingMapOutput)
}

type AccessPolicyIamBindingOutput struct{ *pulumi.OutputState }

func (AccessPolicyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyIamBinding)(nil)).Elem()
}

func (o AccessPolicyIamBindingOutput) ToAccessPolicyIamBindingOutput() AccessPolicyIamBindingOutput {
	return o
}

func (o AccessPolicyIamBindingOutput) ToAccessPolicyIamBindingOutputWithContext(ctx context.Context) AccessPolicyIamBindingOutput {
	return o
}

func (o AccessPolicyIamBindingOutput) Condition() AccessPolicyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *AccessPolicyIamBinding) AccessPolicyIamBindingConditionPtrOutput { return v.Condition }).(AccessPolicyIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o AccessPolicyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AccessPolicyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccessPolicyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o AccessPolicyIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o AccessPolicyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPolicyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AccessPolicyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPolicyIamBinding)(nil)).Elem()
}

func (o AccessPolicyIamBindingArrayOutput) ToAccessPolicyIamBindingArrayOutput() AccessPolicyIamBindingArrayOutput {
	return o
}

func (o AccessPolicyIamBindingArrayOutput) ToAccessPolicyIamBindingArrayOutputWithContext(ctx context.Context) AccessPolicyIamBindingArrayOutput {
	return o
}

func (o AccessPolicyIamBindingArrayOutput) Index(i pulumi.IntInput) AccessPolicyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPolicyIamBinding {
		return vs[0].([]*AccessPolicyIamBinding)[vs[1].(int)]
	}).(AccessPolicyIamBindingOutput)
}

type AccessPolicyIamBindingMapOutput struct{ *pulumi.OutputState }

func (AccessPolicyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPolicyIamBinding)(nil)).Elem()
}

func (o AccessPolicyIamBindingMapOutput) ToAccessPolicyIamBindingMapOutput() AccessPolicyIamBindingMapOutput {
	return o
}

func (o AccessPolicyIamBindingMapOutput) ToAccessPolicyIamBindingMapOutputWithContext(ctx context.Context) AccessPolicyIamBindingMapOutput {
	return o
}

func (o AccessPolicyIamBindingMapOutput) MapIndex(k pulumi.StringInput) AccessPolicyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPolicyIamBinding {
		return vs[0].(map[string]*AccessPolicyIamBinding)[vs[1].(string)]
	}).(AccessPolicyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyIamBindingInput)(nil)).Elem(), &AccessPolicyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyIamBindingArrayInput)(nil)).Elem(), AccessPolicyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyIamBindingMapInput)(nil)).Elem(), AccessPolicyIamBindingMap{})
	pulumi.RegisterOutputType(AccessPolicyIamBindingOutput{})
	pulumi.RegisterOutputType(AccessPolicyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyIamBindingMapOutput{})
}
