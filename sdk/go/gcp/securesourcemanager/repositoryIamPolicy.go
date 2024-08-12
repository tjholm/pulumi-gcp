// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securesourcemanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Secure Source Manager Repository. Each of these resources serves a different use case:
//
// * `securesourcemanager.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
// * `securesourcemanager.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
// * `securesourcemanager.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `securesourcemanager.RepositoryIamPolicy`: Retrieves the IAM policy for the repository
//
// > **Note:** `securesourcemanager.RepositoryIamPolicy` **cannot** be used in conjunction with `securesourcemanager.RepositoryIamBinding` and `securesourcemanager.RepositoryIamMember` or they will fight over what your policy should be.
//
// > **Note:** `securesourcemanager.RepositoryIamBinding` resources **can be** used in conjunction with `securesourcemanager.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## securesourcemanager.RepositoryIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/securesourcemanager.repoAdmin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = securesourcemanager.NewRepositoryIamPolicy(ctx, "policy", &securesourcemanager.RepositoryIamPolicyArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## securesourcemanager.RepositoryIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := securesourcemanager.NewRepositoryIamBinding(ctx, "binding", &securesourcemanager.RepositoryIamBindingArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				Role:         pulumi.String("roles/securesourcemanager.repoAdmin"),
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
// ## securesourcemanager.RepositoryIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := securesourcemanager.NewRepositoryIamMember(ctx, "member", &securesourcemanager.RepositoryIamMemberArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				Role:         pulumi.String("roles/securesourcemanager.repoAdmin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// ## This resource supports User Project Overrides.
//
// -
//
// # IAM policy for Secure Source Manager Repository
// Three different resources help you manage your IAM policy for Secure Source Manager Repository. Each of these resources serves a different use case:
//
// * `securesourcemanager.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
// * `securesourcemanager.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
// * `securesourcemanager.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `securesourcemanager.RepositoryIamPolicy`: Retrieves the IAM policy for the repository
//
// > **Note:** `securesourcemanager.RepositoryIamPolicy` **cannot** be used in conjunction with `securesourcemanager.RepositoryIamBinding` and `securesourcemanager.RepositoryIamMember` or they will fight over what your policy should be.
//
// > **Note:** `securesourcemanager.RepositoryIamBinding` resources **can be** used in conjunction with `securesourcemanager.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## securesourcemanager.RepositoryIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/securesourcemanager.repoAdmin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = securesourcemanager.NewRepositoryIamPolicy(ctx, "policy", &securesourcemanager.RepositoryIamPolicyArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				PolicyData:   pulumi.String(admin.PolicyData),
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
// ## securesourcemanager.RepositoryIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := securesourcemanager.NewRepositoryIamBinding(ctx, "binding", &securesourcemanager.RepositoryIamBindingArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				Role:         pulumi.String("roles/securesourcemanager.repoAdmin"),
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
// ## securesourcemanager.RepositoryIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/securesourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := securesourcemanager.NewRepositoryIamMember(ctx, "member", &securesourcemanager.RepositoryIamMemberArgs{
//				Project:      pulumi.Any(_default.Project),
//				Location:     pulumi.Any(_default.Location),
//				RepositoryId: pulumi.Any(_default.RepositoryId),
//				Role:         pulumi.String("roles/securesourcemanager.repoAdmin"),
//				Member:       pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
//
// * {{project}}/{{location}}/{{repository_id}}
//
// * {{location}}/{{repository_id}}
//
// * {{repository_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Secure Source Manager repository IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}} roles/securesourcemanager.repoAdmin user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor "projects/{{project}}/locations/{{location}}/repositories/{{repository_id}} roles/securesourcemanager.repoAdmin"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy editor projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RepositoryIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location for the Repository.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The ID for the Repository.
	// Used to find the parent resource to bind the IAM policy to
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
}

// NewRepositoryIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamPolicy(ctx *pulumi.Context,
	name string, args *RepositoryIamPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryIamPolicy
	err := ctx.RegisterResource("gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamPolicy gets an existing RepositoryIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamPolicyState, opts ...pulumi.ResourceOption) (*RepositoryIamPolicy, error) {
	var resource RepositoryIamPolicy
	err := ctx.ReadResource("gcp:securesourcemanager/repositoryIamPolicy:RepositoryIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamPolicy resources.
type repositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location for the Repository.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID for the Repository.
	// Used to find the parent resource to bind the IAM policy to
	RepositoryId *string `pulumi:"repositoryId"`
}

type RepositoryIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location for the Repository.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID for the Repository.
	// Used to find the parent resource to bind the IAM policy to
	RepositoryId pulumi.StringPtrInput
}

func (RepositoryIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyState)(nil)).Elem()
}

type repositoryIamPolicyArgs struct {
	// The location for the Repository.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID for the Repository.
	// Used to find the parent resource to bind the IAM policy to
	RepositoryId string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a RepositoryIamPolicy resource.
type RepositoryIamPolicyArgs struct {
	// The location for the Repository.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID for the Repository.
	// Used to find the parent resource to bind the IAM policy to
	RepositoryId pulumi.StringInput
}

func (RepositoryIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamPolicyArgs)(nil)).Elem()
}

type RepositoryIamPolicyInput interface {
	pulumi.Input

	ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput
	ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput
}

func (*RepositoryIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return i.ToRepositoryIamPolicyOutputWithContext(context.Background())
}

func (i *RepositoryIamPolicy) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyOutput)
}

// RepositoryIamPolicyArrayInput is an input type that accepts RepositoryIamPolicyArray and RepositoryIamPolicyArrayOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyArrayInput` via:
//
//	RepositoryIamPolicyArray{ RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyArrayInput interface {
	pulumi.Input

	ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput
	ToRepositoryIamPolicyArrayOutputWithContext(context.Context) RepositoryIamPolicyArrayOutput
}

type RepositoryIamPolicyArray []RepositoryIamPolicyInput

func (RepositoryIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamPolicy)(nil)).Elem()
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return i.ToRepositoryIamPolicyArrayOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyArray) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyArrayOutput)
}

// RepositoryIamPolicyMapInput is an input type that accepts RepositoryIamPolicyMap and RepositoryIamPolicyMapOutput values.
// You can construct a concrete instance of `RepositoryIamPolicyMapInput` via:
//
//	RepositoryIamPolicyMap{ "key": RepositoryIamPolicyArgs{...} }
type RepositoryIamPolicyMapInput interface {
	pulumi.Input

	ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput
	ToRepositoryIamPolicyMapOutputWithContext(context.Context) RepositoryIamPolicyMapOutput
}

type RepositoryIamPolicyMap map[string]RepositoryIamPolicyInput

func (RepositoryIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamPolicy)(nil)).Elem()
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return i.ToRepositoryIamPolicyMapOutputWithContext(context.Background())
}

func (i RepositoryIamPolicyMap) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryIamPolicyMapOutput)
}

type RepositoryIamPolicyOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutput() RepositoryIamPolicyOutput {
	return o
}

func (o RepositoryIamPolicyOutput) ToRepositoryIamPolicyOutputWithContext(ctx context.Context) RepositoryIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o RepositoryIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location for the Repository.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o RepositoryIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o RepositoryIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o RepositoryIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The ID for the Repository.
// Used to find the parent resource to bind the IAM policy to
func (o RepositoryIamPolicyOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryIamPolicy) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

type RepositoryIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutput() RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) ToRepositoryIamPolicyArrayOutputWithContext(ctx context.Context) RepositoryIamPolicyArrayOutput {
	return o
}

func (o RepositoryIamPolicyArrayOutput) Index(i pulumi.IntInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryIamPolicy {
		return vs[0].([]*RepositoryIamPolicy)[vs[1].(int)]
	}).(RepositoryIamPolicyOutput)
}

type RepositoryIamPolicyMapOutput struct{ *pulumi.OutputState }

func (RepositoryIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryIamPolicy)(nil)).Elem()
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutput() RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) ToRepositoryIamPolicyMapOutputWithContext(ctx context.Context) RepositoryIamPolicyMapOutput {
	return o
}

func (o RepositoryIamPolicyMapOutput) MapIndex(k pulumi.StringInput) RepositoryIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryIamPolicy {
		return vs[0].(map[string]*RepositoryIamPolicy)[vs[1].(string)]
	}).(RepositoryIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyInput)(nil)).Elem(), &RepositoryIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyArrayInput)(nil)).Elem(), RepositoryIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryIamPolicyMapInput)(nil)).Elem(), RepositoryIamPolicyMap{})
	pulumi.RegisterOutputType(RepositoryIamPolicyOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryIamPolicyMapOutput{})
}
