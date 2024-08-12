// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataplex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Dataplex EntryGroup. Each of these resources serves a different use case:
//
// * `dataplex.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
// * `dataplex.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
// * `dataplex.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.EntryGroupIamPolicy`: Retrieves the IAM policy for the entrygroup
//
// > **Note:** `dataplex.EntryGroupIamPolicy` **cannot** be used in conjunction with `dataplex.EntryGroupIamBinding` and `dataplex.EntryGroupIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.EntryGroupIamBinding` resources **can be** used in conjunction with `dataplex.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.EntryGroupIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewEntryGroupIamPolicy(ctx, "policy", &dataplex.EntryGroupIamPolicyArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
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
// ## dataplex.EntryGroupIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewEntryGroupIamBinding(ctx, "binding", &dataplex.EntryGroupIamBindingArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## dataplex.EntryGroupIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewEntryGroupIamMember(ctx, "member", &dataplex.EntryGroupIamMemberArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
//				Role:         pulumi.String("roles/viewer"),
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
// # IAM policy for Dataplex EntryGroup
// Three different resources help you manage your IAM policy for Dataplex EntryGroup. Each of these resources serves a different use case:
//
// * `dataplex.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
// * `dataplex.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
// * `dataplex.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.EntryGroupIamPolicy`: Retrieves the IAM policy for the entrygroup
//
// > **Note:** `dataplex.EntryGroupIamPolicy` **cannot** be used in conjunction with `dataplex.EntryGroupIamBinding` and `dataplex.EntryGroupIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.EntryGroupIamBinding` resources **can be** used in conjunction with `dataplex.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.EntryGroupIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
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
//			_, err = dataplex.NewEntryGroupIamPolicy(ctx, "policy", &dataplex.EntryGroupIamPolicyArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
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
// ## dataplex.EntryGroupIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewEntryGroupIamBinding(ctx, "binding", &dataplex.EntryGroupIamBindingArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
//				Role:         pulumi.String("roles/viewer"),
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
// ## dataplex.EntryGroupIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataplex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataplex.NewEntryGroupIamMember(ctx, "member", &dataplex.EntryGroupIamMemberArgs{
//				Project:      pulumi.Any(testEntryGroupBasic.Project),
//				Location:     pulumi.Any(testEntryGroupBasic.Location),
//				EntryGroupId: pulumi.Any(testEntryGroupBasic.EntryGroupId),
//				Role:         pulumi.String("roles/viewer"),
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
// * projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}
//
// * {{project}}/{{location}}/{{entry_group_id}}
//
// * {{location}}/{{entry_group_id}}
//
// * {{entry_group_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataplex entrygroup IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/entryGroupIamPolicy:EntryGroupIamPolicy editor "projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/entryGroupIamPolicy:EntryGroupIamPolicy editor "projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/entryGroupIamPolicy:EntryGroupIamPolicy editor projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type EntryGroupIamPolicy struct {
	pulumi.CustomResourceState

	EntryGroupId pulumi.StringOutput `pulumi:"entryGroupId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where entry group will be created in.
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
}

// NewEntryGroupIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, args *EntryGroupIamPolicyArgs, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EntryGroupIamPolicy
	err := ctx.RegisterResource("gcp:dataplex/entryGroupIamPolicy:EntryGroupIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroupIamPolicy gets an existing EntryGroupIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroupIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupIamPolicyState, opts ...pulumi.ResourceOption) (*EntryGroupIamPolicy, error) {
	var resource EntryGroupIamPolicy
	err := ctx.ReadResource("gcp:dataplex/entryGroupIamPolicy:EntryGroupIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroupIamPolicy resources.
type entryGroupIamPolicyState struct {
	EntryGroupId *string `pulumi:"entryGroupId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where entry group will be created in.
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
}

type EntryGroupIamPolicyState struct {
	EntryGroupId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where entry group will be created in.
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
}

func (EntryGroupIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyState)(nil)).Elem()
}

type entryGroupIamPolicyArgs struct {
	EntryGroupId string `pulumi:"entryGroupId"`
	// The location where entry group will be created in.
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
}

// The set of arguments for constructing a EntryGroupIamPolicy resource.
type EntryGroupIamPolicyArgs struct {
	EntryGroupId pulumi.StringInput
	// The location where entry group will be created in.
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
}

func (EntryGroupIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamPolicyArgs)(nil)).Elem()
}

type EntryGroupIamPolicyInput interface {
	pulumi.Input

	ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput
	ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput
}

func (*EntryGroupIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamPolicy)(nil)).Elem()
}

func (i *EntryGroupIamPolicy) ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput {
	return i.ToEntryGroupIamPolicyOutputWithContext(context.Background())
}

func (i *EntryGroupIamPolicy) ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamPolicyOutput)
}

// EntryGroupIamPolicyArrayInput is an input type that accepts EntryGroupIamPolicyArray and EntryGroupIamPolicyArrayOutput values.
// You can construct a concrete instance of `EntryGroupIamPolicyArrayInput` via:
//
//	EntryGroupIamPolicyArray{ EntryGroupIamPolicyArgs{...} }
type EntryGroupIamPolicyArrayInput interface {
	pulumi.Input

	ToEntryGroupIamPolicyArrayOutput() EntryGroupIamPolicyArrayOutput
	ToEntryGroupIamPolicyArrayOutputWithContext(context.Context) EntryGroupIamPolicyArrayOutput
}

type EntryGroupIamPolicyArray []EntryGroupIamPolicyInput

func (EntryGroupIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroupIamPolicy)(nil)).Elem()
}

func (i EntryGroupIamPolicyArray) ToEntryGroupIamPolicyArrayOutput() EntryGroupIamPolicyArrayOutput {
	return i.ToEntryGroupIamPolicyArrayOutputWithContext(context.Background())
}

func (i EntryGroupIamPolicyArray) ToEntryGroupIamPolicyArrayOutputWithContext(ctx context.Context) EntryGroupIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamPolicyArrayOutput)
}

// EntryGroupIamPolicyMapInput is an input type that accepts EntryGroupIamPolicyMap and EntryGroupIamPolicyMapOutput values.
// You can construct a concrete instance of `EntryGroupIamPolicyMapInput` via:
//
//	EntryGroupIamPolicyMap{ "key": EntryGroupIamPolicyArgs{...} }
type EntryGroupIamPolicyMapInput interface {
	pulumi.Input

	ToEntryGroupIamPolicyMapOutput() EntryGroupIamPolicyMapOutput
	ToEntryGroupIamPolicyMapOutputWithContext(context.Context) EntryGroupIamPolicyMapOutput
}

type EntryGroupIamPolicyMap map[string]EntryGroupIamPolicyInput

func (EntryGroupIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroupIamPolicy)(nil)).Elem()
}

func (i EntryGroupIamPolicyMap) ToEntryGroupIamPolicyMapOutput() EntryGroupIamPolicyMapOutput {
	return i.ToEntryGroupIamPolicyMapOutputWithContext(context.Background())
}

func (i EntryGroupIamPolicyMap) ToEntryGroupIamPolicyMapOutputWithContext(ctx context.Context) EntryGroupIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamPolicyMapOutput)
}

type EntryGroupIamPolicyOutput struct{ *pulumi.OutputState }

func (EntryGroupIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamPolicy)(nil)).Elem()
}

func (o EntryGroupIamPolicyOutput) ToEntryGroupIamPolicyOutput() EntryGroupIamPolicyOutput {
	return o
}

func (o EntryGroupIamPolicyOutput) ToEntryGroupIamPolicyOutputWithContext(ctx context.Context) EntryGroupIamPolicyOutput {
	return o
}

func (o EntryGroupIamPolicyOutput) EntryGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamPolicy) pulumi.StringOutput { return v.EntryGroupId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o EntryGroupIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where entry group will be created in.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o EntryGroupIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o EntryGroupIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o EntryGroupIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type EntryGroupIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (EntryGroupIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroupIamPolicy)(nil)).Elem()
}

func (o EntryGroupIamPolicyArrayOutput) ToEntryGroupIamPolicyArrayOutput() EntryGroupIamPolicyArrayOutput {
	return o
}

func (o EntryGroupIamPolicyArrayOutput) ToEntryGroupIamPolicyArrayOutputWithContext(ctx context.Context) EntryGroupIamPolicyArrayOutput {
	return o
}

func (o EntryGroupIamPolicyArrayOutput) Index(i pulumi.IntInput) EntryGroupIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EntryGroupIamPolicy {
		return vs[0].([]*EntryGroupIamPolicy)[vs[1].(int)]
	}).(EntryGroupIamPolicyOutput)
}

type EntryGroupIamPolicyMapOutput struct{ *pulumi.OutputState }

func (EntryGroupIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroupIamPolicy)(nil)).Elem()
}

func (o EntryGroupIamPolicyMapOutput) ToEntryGroupIamPolicyMapOutput() EntryGroupIamPolicyMapOutput {
	return o
}

func (o EntryGroupIamPolicyMapOutput) ToEntryGroupIamPolicyMapOutputWithContext(ctx context.Context) EntryGroupIamPolicyMapOutput {
	return o
}

func (o EntryGroupIamPolicyMapOutput) MapIndex(k pulumi.StringInput) EntryGroupIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EntryGroupIamPolicy {
		return vs[0].(map[string]*EntryGroupIamPolicy)[vs[1].(string)]
	}).(EntryGroupIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamPolicyInput)(nil)).Elem(), &EntryGroupIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamPolicyArrayInput)(nil)).Elem(), EntryGroupIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamPolicyMapInput)(nil)).Elem(), EntryGroupIamPolicyMap{})
	pulumi.RegisterOutputType(EntryGroupIamPolicyOutput{})
	pulumi.RegisterOutputType(EntryGroupIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(EntryGroupIamPolicyMapOutput{})
}
