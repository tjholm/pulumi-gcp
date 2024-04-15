// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Apigee Environment. Each of these resources serves a different use case:
//
// * `apigee.EnvironmentIamPolicy`: Authoritative. Sets the IAM policy for the environment and replaces any existing policy already attached.
// * `apigee.EnvironmentIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the environment are preserved.
// * `apigee.EnvironmentIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the environment are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `apigee.EnvironmentIamPolicy`: Retrieves the IAM policy for the environment
//
// > **Note:** `apigee.EnvironmentIamPolicy` **cannot** be used in conjunction with `apigee.EnvironmentIamBinding` and `apigee.EnvironmentIamMember` or they will fight over what your policy should be.
//
// > **Note:** `apigee.EnvironmentIamBinding` resources **can be** used in conjunction with `apigee.EnvironmentIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_apigee\_environment\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewEnvironmentIamPolicy(ctx, "policy", &apigee.EnvironmentIamPolicyArgs{
//				OrgId:      pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId:      pulumi.Any(apigeeEnvironment.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_apigee\_environment\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigee.NewEnvironmentIamBinding(ctx, "binding", &apigee.EnvironmentIamBindingArgs{
//				OrgId: pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId: pulumi.Any(apigeeEnvironment.Name),
//				Role:  pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_apigee\_environment\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigee.NewEnvironmentIamMember(ctx, "member", &apigee.EnvironmentIamMemberArgs{
//				OrgId:  pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId:  pulumi.Any(apigeeEnvironment.Name),
//				Role:   pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_apigee\_environment\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewEnvironmentIamPolicy(ctx, "policy", &apigee.EnvironmentIamPolicyArgs{
//				OrgId:      pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId:      pulumi.Any(apigeeEnvironment.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_apigee\_environment\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigee.NewEnvironmentIamBinding(ctx, "binding", &apigee.EnvironmentIamBindingArgs{
//				OrgId: pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId: pulumi.Any(apigeeEnvironment.Name),
//				Role:  pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_apigee\_environment\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigee.NewEnvironmentIamMember(ctx, "member", &apigee.EnvironmentIamMemberArgs{
//				OrgId:  pulumi.Any(apigeeEnvironment.OrgId),
//				EnvId:  pulumi.Any(apigeeEnvironment.Name),
//				Role:   pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * {{org_id}}/environments/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Apigee environment IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy editor "{{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy editor "{{org_id}}/environments/{{environment}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy editor {{org_id}}/environments/{{environment}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type EnvironmentIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringOutput `pulumi:"envId"`
	// (Computed) The etag of the IAM policy.
	Etag  pulumi.StringOutput `pulumi:"etag"`
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewEnvironmentIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentIamPolicy(ctx *pulumi.Context,
	name string, args *EnvironmentIamPolicyArgs, opts ...pulumi.ResourceOption) (*EnvironmentIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvironmentIamPolicy
	err := ctx.RegisterResource("gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentIamPolicy gets an existing EnvironmentIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentIamPolicyState, opts ...pulumi.ResourceOption) (*EnvironmentIamPolicy, error) {
	var resource EnvironmentIamPolicy
	err := ctx.ReadResource("gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentIamPolicy resources.
type environmentIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	EnvId *string `pulumi:"envId"`
	// (Computed) The etag of the IAM policy.
	Etag  *string `pulumi:"etag"`
	OrgId *string `pulumi:"orgId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type EnvironmentIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag  pulumi.StringPtrInput
	OrgId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (EnvironmentIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentIamPolicyState)(nil)).Elem()
}

type environmentIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	EnvId string `pulumi:"envId"`
	OrgId string `pulumi:"orgId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a EnvironmentIamPolicy resource.
type EnvironmentIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringInput
	OrgId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (EnvironmentIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentIamPolicyArgs)(nil)).Elem()
}

type EnvironmentIamPolicyInput interface {
	pulumi.Input

	ToEnvironmentIamPolicyOutput() EnvironmentIamPolicyOutput
	ToEnvironmentIamPolicyOutputWithContext(ctx context.Context) EnvironmentIamPolicyOutput
}

func (*EnvironmentIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentIamPolicy)(nil)).Elem()
}

func (i *EnvironmentIamPolicy) ToEnvironmentIamPolicyOutput() EnvironmentIamPolicyOutput {
	return i.ToEnvironmentIamPolicyOutputWithContext(context.Background())
}

func (i *EnvironmentIamPolicy) ToEnvironmentIamPolicyOutputWithContext(ctx context.Context) EnvironmentIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamPolicyOutput)
}

// EnvironmentIamPolicyArrayInput is an input type that accepts EnvironmentIamPolicyArray and EnvironmentIamPolicyArrayOutput values.
// You can construct a concrete instance of `EnvironmentIamPolicyArrayInput` via:
//
//	EnvironmentIamPolicyArray{ EnvironmentIamPolicyArgs{...} }
type EnvironmentIamPolicyArrayInput interface {
	pulumi.Input

	ToEnvironmentIamPolicyArrayOutput() EnvironmentIamPolicyArrayOutput
	ToEnvironmentIamPolicyArrayOutputWithContext(context.Context) EnvironmentIamPolicyArrayOutput
}

type EnvironmentIamPolicyArray []EnvironmentIamPolicyInput

func (EnvironmentIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentIamPolicy)(nil)).Elem()
}

func (i EnvironmentIamPolicyArray) ToEnvironmentIamPolicyArrayOutput() EnvironmentIamPolicyArrayOutput {
	return i.ToEnvironmentIamPolicyArrayOutputWithContext(context.Background())
}

func (i EnvironmentIamPolicyArray) ToEnvironmentIamPolicyArrayOutputWithContext(ctx context.Context) EnvironmentIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamPolicyArrayOutput)
}

// EnvironmentIamPolicyMapInput is an input type that accepts EnvironmentIamPolicyMap and EnvironmentIamPolicyMapOutput values.
// You can construct a concrete instance of `EnvironmentIamPolicyMapInput` via:
//
//	EnvironmentIamPolicyMap{ "key": EnvironmentIamPolicyArgs{...} }
type EnvironmentIamPolicyMapInput interface {
	pulumi.Input

	ToEnvironmentIamPolicyMapOutput() EnvironmentIamPolicyMapOutput
	ToEnvironmentIamPolicyMapOutputWithContext(context.Context) EnvironmentIamPolicyMapOutput
}

type EnvironmentIamPolicyMap map[string]EnvironmentIamPolicyInput

func (EnvironmentIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentIamPolicy)(nil)).Elem()
}

func (i EnvironmentIamPolicyMap) ToEnvironmentIamPolicyMapOutput() EnvironmentIamPolicyMapOutput {
	return i.ToEnvironmentIamPolicyMapOutputWithContext(context.Background())
}

func (i EnvironmentIamPolicyMap) ToEnvironmentIamPolicyMapOutputWithContext(ctx context.Context) EnvironmentIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamPolicyMapOutput)
}

type EnvironmentIamPolicyOutput struct{ *pulumi.OutputState }

func (EnvironmentIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentIamPolicy)(nil)).Elem()
}

func (o EnvironmentIamPolicyOutput) ToEnvironmentIamPolicyOutput() EnvironmentIamPolicyOutput {
	return o
}

func (o EnvironmentIamPolicyOutput) ToEnvironmentIamPolicyOutputWithContext(ctx context.Context) EnvironmentIamPolicyOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o EnvironmentIamPolicyOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamPolicy) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o EnvironmentIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o EnvironmentIamPolicyOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamPolicy) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o EnvironmentIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type EnvironmentIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentIamPolicy)(nil)).Elem()
}

func (o EnvironmentIamPolicyArrayOutput) ToEnvironmentIamPolicyArrayOutput() EnvironmentIamPolicyArrayOutput {
	return o
}

func (o EnvironmentIamPolicyArrayOutput) ToEnvironmentIamPolicyArrayOutputWithContext(ctx context.Context) EnvironmentIamPolicyArrayOutput {
	return o
}

func (o EnvironmentIamPolicyArrayOutput) Index(i pulumi.IntInput) EnvironmentIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvironmentIamPolicy {
		return vs[0].([]*EnvironmentIamPolicy)[vs[1].(int)]
	}).(EnvironmentIamPolicyOutput)
}

type EnvironmentIamPolicyMapOutput struct{ *pulumi.OutputState }

func (EnvironmentIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentIamPolicy)(nil)).Elem()
}

func (o EnvironmentIamPolicyMapOutput) ToEnvironmentIamPolicyMapOutput() EnvironmentIamPolicyMapOutput {
	return o
}

func (o EnvironmentIamPolicyMapOutput) ToEnvironmentIamPolicyMapOutputWithContext(ctx context.Context) EnvironmentIamPolicyMapOutput {
	return o
}

func (o EnvironmentIamPolicyMapOutput) MapIndex(k pulumi.StringInput) EnvironmentIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvironmentIamPolicy {
		return vs[0].(map[string]*EnvironmentIamPolicy)[vs[1].(string)]
	}).(EnvironmentIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamPolicyInput)(nil)).Elem(), &EnvironmentIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamPolicyArrayInput)(nil)).Elem(), EnvironmentIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamPolicyMapInput)(nil)).Elem(), EnvironmentIamPolicyMap{})
	pulumi.RegisterOutputType(EnvironmentIamPolicyOutput{})
	pulumi.RegisterOutputType(EnvironmentIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentIamPolicyMapOutput{})
}
