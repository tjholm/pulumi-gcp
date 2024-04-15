// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Directory Namespace. Each of these resources serves a different use case:
//
// * `servicedirectory.NamespaceIamPolicy`: Authoritative. Sets the IAM policy for the namespace and replaces any existing policy already attached.
// * `servicedirectory.NamespaceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the namespace are preserved.
// * `servicedirectory.NamespaceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the namespace are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `servicedirectory.NamespaceIamPolicy`: Retrieves the IAM policy for the namespace
//
// > **Note:** `servicedirectory.NamespaceIamPolicy` **cannot** be used in conjunction with `servicedirectory.NamespaceIamBinding` and `servicedirectory.NamespaceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `servicedirectory.NamespaceIamBinding` resources **can be** used in conjunction with `servicedirectory.NamespaceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_directory\_namespace\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
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
//			_, err = servicedirectory.NewNamespaceIamPolicy(ctx, "policy", &servicedirectory.NamespaceIamPolicyArgs{
//				Name:       pulumi.Any(example.Name),
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
// ## google\_service\_directory\_namespace\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicedirectory.NewNamespaceIamBinding(ctx, "binding", &servicedirectory.NamespaceIamBindingArgs{
//				Name: pulumi.Any(example.Name),
//				Role: pulumi.String("roles/viewer"),
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
// ## google\_service\_directory\_namespace\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicedirectory.NewNamespaceIamMember(ctx, "member", &servicedirectory.NamespaceIamMemberArgs{
//				Name:   pulumi.Any(example.Name),
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
// ## google\_service\_directory\_namespace\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
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
//			_, err = servicedirectory.NewNamespaceIamPolicy(ctx, "policy", &servicedirectory.NamespaceIamPolicyArgs{
//				Name:       pulumi.Any(example.Name),
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
// ## google\_service\_directory\_namespace\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicedirectory.NewNamespaceIamBinding(ctx, "binding", &servicedirectory.NamespaceIamBindingArgs{
//				Name: pulumi.Any(example.Name),
//				Role: pulumi.String("roles/viewer"),
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
// ## google\_service\_directory\_namespace\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicedirectory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicedirectory.NewNamespaceIamMember(ctx, "member", &servicedirectory.NamespaceIamMemberArgs{
//				Name:   pulumi.Any(example.Name),
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
// * projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
//
// * {{project}}/{{location}}/{{namespace_id}}
//
// * {{location}}/{{namespace_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Service Directory namespace IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NamespaceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewNamespaceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewNamespaceIamPolicy(ctx *pulumi.Context,
	name string, args *NamespaceIamPolicyArgs, opts ...pulumi.ResourceOption) (*NamespaceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamespaceIamPolicy
	err := ctx.RegisterResource("gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceIamPolicy gets an existing NamespaceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIamPolicyState, opts ...pulumi.ResourceOption) (*NamespaceIamPolicy, error) {
	var resource NamespaceIamPolicy
	err := ctx.ReadResource("gcp:servicedirectory/namespaceIamPolicy:NamespaceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceIamPolicy resources.
type namespaceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type NamespaceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (NamespaceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamPolicyState)(nil)).Elem()
}

type namespaceIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a NamespaceIamPolicy resource.
type NamespaceIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (NamespaceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamPolicyArgs)(nil)).Elem()
}

type NamespaceIamPolicyInput interface {
	pulumi.Input

	ToNamespaceIamPolicyOutput() NamespaceIamPolicyOutput
	ToNamespaceIamPolicyOutputWithContext(ctx context.Context) NamespaceIamPolicyOutput
}

func (*NamespaceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamPolicy)(nil)).Elem()
}

func (i *NamespaceIamPolicy) ToNamespaceIamPolicyOutput() NamespaceIamPolicyOutput {
	return i.ToNamespaceIamPolicyOutputWithContext(context.Background())
}

func (i *NamespaceIamPolicy) ToNamespaceIamPolicyOutputWithContext(ctx context.Context) NamespaceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamPolicyOutput)
}

// NamespaceIamPolicyArrayInput is an input type that accepts NamespaceIamPolicyArray and NamespaceIamPolicyArrayOutput values.
// You can construct a concrete instance of `NamespaceIamPolicyArrayInput` via:
//
//	NamespaceIamPolicyArray{ NamespaceIamPolicyArgs{...} }
type NamespaceIamPolicyArrayInput interface {
	pulumi.Input

	ToNamespaceIamPolicyArrayOutput() NamespaceIamPolicyArrayOutput
	ToNamespaceIamPolicyArrayOutputWithContext(context.Context) NamespaceIamPolicyArrayOutput
}

type NamespaceIamPolicyArray []NamespaceIamPolicyInput

func (NamespaceIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceIamPolicy)(nil)).Elem()
}

func (i NamespaceIamPolicyArray) ToNamespaceIamPolicyArrayOutput() NamespaceIamPolicyArrayOutput {
	return i.ToNamespaceIamPolicyArrayOutputWithContext(context.Background())
}

func (i NamespaceIamPolicyArray) ToNamespaceIamPolicyArrayOutputWithContext(ctx context.Context) NamespaceIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamPolicyArrayOutput)
}

// NamespaceIamPolicyMapInput is an input type that accepts NamespaceIamPolicyMap and NamespaceIamPolicyMapOutput values.
// You can construct a concrete instance of `NamespaceIamPolicyMapInput` via:
//
//	NamespaceIamPolicyMap{ "key": NamespaceIamPolicyArgs{...} }
type NamespaceIamPolicyMapInput interface {
	pulumi.Input

	ToNamespaceIamPolicyMapOutput() NamespaceIamPolicyMapOutput
	ToNamespaceIamPolicyMapOutputWithContext(context.Context) NamespaceIamPolicyMapOutput
}

type NamespaceIamPolicyMap map[string]NamespaceIamPolicyInput

func (NamespaceIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceIamPolicy)(nil)).Elem()
}

func (i NamespaceIamPolicyMap) ToNamespaceIamPolicyMapOutput() NamespaceIamPolicyMapOutput {
	return i.ToNamespaceIamPolicyMapOutputWithContext(context.Background())
}

func (i NamespaceIamPolicyMap) ToNamespaceIamPolicyMapOutputWithContext(ctx context.Context) NamespaceIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamPolicyMapOutput)
}

type NamespaceIamPolicyOutput struct{ *pulumi.OutputState }

func (NamespaceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamPolicy)(nil)).Elem()
}

func (o NamespaceIamPolicyOutput) ToNamespaceIamPolicyOutput() NamespaceIamPolicyOutput {
	return o
}

func (o NamespaceIamPolicyOutput) ToNamespaceIamPolicyOutputWithContext(ctx context.Context) NamespaceIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o NamespaceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o NamespaceIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o NamespaceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type NamespaceIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (NamespaceIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceIamPolicy)(nil)).Elem()
}

func (o NamespaceIamPolicyArrayOutput) ToNamespaceIamPolicyArrayOutput() NamespaceIamPolicyArrayOutput {
	return o
}

func (o NamespaceIamPolicyArrayOutput) ToNamespaceIamPolicyArrayOutputWithContext(ctx context.Context) NamespaceIamPolicyArrayOutput {
	return o
}

func (o NamespaceIamPolicyArrayOutput) Index(i pulumi.IntInput) NamespaceIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamespaceIamPolicy {
		return vs[0].([]*NamespaceIamPolicy)[vs[1].(int)]
	}).(NamespaceIamPolicyOutput)
}

type NamespaceIamPolicyMapOutput struct{ *pulumi.OutputState }

func (NamespaceIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceIamPolicy)(nil)).Elem()
}

func (o NamespaceIamPolicyMapOutput) ToNamespaceIamPolicyMapOutput() NamespaceIamPolicyMapOutput {
	return o
}

func (o NamespaceIamPolicyMapOutput) ToNamespaceIamPolicyMapOutputWithContext(ctx context.Context) NamespaceIamPolicyMapOutput {
	return o
}

func (o NamespaceIamPolicyMapOutput) MapIndex(k pulumi.StringInput) NamespaceIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamespaceIamPolicy {
		return vs[0].(map[string]*NamespaceIamPolicy)[vs[1].(string)]
	}).(NamespaceIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamPolicyInput)(nil)).Elem(), &NamespaceIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamPolicyArrayInput)(nil)).Elem(), NamespaceIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamPolicyMapInput)(nil)).Elem(), NamespaceIamPolicyMap{})
	pulumi.RegisterOutputType(NamespaceIamPolicyOutput{})
	pulumi.RegisterOutputType(NamespaceIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(NamespaceIamPolicyMapOutput{})
}
