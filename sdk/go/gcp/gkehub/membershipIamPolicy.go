// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for GKEHub Membership. Each of these resources serves a different use case:
//
// * `gkehub.MembershipIamPolicy`: Authoritative. Sets the IAM policy for the membership and replaces any existing policy already attached.
// * `gkehub.MembershipIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the membership are preserved.
// * `gkehub.MembershipIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the membership are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `gkehub.MembershipIamPolicy`: Retrieves the IAM policy for the membership
//
// > **Note:** `gkehub.MembershipIamPolicy` **cannot** be used in conjunction with `gkehub.MembershipIamBinding` and `gkehub.MembershipIamMember` or they will fight over what your policy should be.
//
// > **Note:** `gkehub.MembershipIamBinding` resources **can be** used in conjunction with `gkehub.MembershipIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## gkehub.MembershipIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
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
//			_, err = gkehub.NewMembershipIamPolicy(ctx, "policy", &gkehub.MembershipIamPolicyArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// ## gkehub.MembershipIamBinding
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
//			_, err := gkehub.NewMembershipIamBinding(ctx, "binding", &gkehub.MembershipIamBindingArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// ## gkehub.MembershipIamMember
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
//			_, err := gkehub.NewMembershipIamMember(ctx, "member", &gkehub.MembershipIamMemberArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// # IAM policy for GKEHub Membership
// Three different resources help you manage your IAM policy for GKEHub Membership. Each of these resources serves a different use case:
//
// * `gkehub.MembershipIamPolicy`: Authoritative. Sets the IAM policy for the membership and replaces any existing policy already attached.
// * `gkehub.MembershipIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the membership are preserved.
// * `gkehub.MembershipIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the membership are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `gkehub.MembershipIamPolicy`: Retrieves the IAM policy for the membership
//
// > **Note:** `gkehub.MembershipIamPolicy` **cannot** be used in conjunction with `gkehub.MembershipIamBinding` and `gkehub.MembershipIamMember` or they will fight over what your policy should be.
//
// > **Note:** `gkehub.MembershipIamBinding` resources **can be** used in conjunction with `gkehub.MembershipIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## gkehub.MembershipIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
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
//			_, err = gkehub.NewMembershipIamPolicy(ctx, "policy", &gkehub.MembershipIamPolicyArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// ## gkehub.MembershipIamBinding
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
//			_, err := gkehub.NewMembershipIamBinding(ctx, "binding", &gkehub.MembershipIamBindingArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// ## gkehub.MembershipIamMember
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
//			_, err := gkehub.NewMembershipIamMember(ctx, "member", &gkehub.MembershipIamMemberArgs{
//				Project:      pulumi.Any(membership.Project),
//				Location:     pulumi.Any(membership.Location),
//				MembershipId: pulumi.Any(membership.MembershipId),
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
// * projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
//
// * {{project}}/{{location}}/{{membership_id}}
//
// * {{location}}/{{membership_id}}
//
// * {{membership_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// GKEHub membership IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor "projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor "projects/{{project}}/locations/{{location}}/memberships/{{membership_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:gkehub/membershipIamPolicy:MembershipIamPolicy editor projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MembershipIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Location of the membership.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location     pulumi.StringOutput `pulumi:"location"`
	MembershipId pulumi.StringOutput `pulumi:"membershipId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewMembershipIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewMembershipIamPolicy(ctx *pulumi.Context,
	name string, args *MembershipIamPolicyArgs, opts ...pulumi.ResourceOption) (*MembershipIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MembershipId == nil {
		return nil, errors.New("invalid value for required argument 'MembershipId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MembershipIamPolicy
	err := ctx.RegisterResource("gcp:gkehub/membershipIamPolicy:MembershipIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembershipIamPolicy gets an existing MembershipIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembershipIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipIamPolicyState, opts ...pulumi.ResourceOption) (*MembershipIamPolicy, error) {
	var resource MembershipIamPolicy
	err := ctx.ReadResource("gcp:gkehub/membershipIamPolicy:MembershipIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MembershipIamPolicy resources.
type membershipIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Location of the membership.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location     *string `pulumi:"location"`
	MembershipId *string `pulumi:"membershipId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type MembershipIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Location of the membership.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location     pulumi.StringPtrInput
	MembershipId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MembershipIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipIamPolicyState)(nil)).Elem()
}

type membershipIamPolicyArgs struct {
	// Location of the membership.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location     *string `pulumi:"location"`
	MembershipId string  `pulumi:"membershipId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a MembershipIamPolicy resource.
type MembershipIamPolicyArgs struct {
	// Location of the membership.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location     pulumi.StringPtrInput
	MembershipId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MembershipIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipIamPolicyArgs)(nil)).Elem()
}

type MembershipIamPolicyInput interface {
	pulumi.Input

	ToMembershipIamPolicyOutput() MembershipIamPolicyOutput
	ToMembershipIamPolicyOutputWithContext(ctx context.Context) MembershipIamPolicyOutput
}

func (*MembershipIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipIamPolicy)(nil)).Elem()
}

func (i *MembershipIamPolicy) ToMembershipIamPolicyOutput() MembershipIamPolicyOutput {
	return i.ToMembershipIamPolicyOutputWithContext(context.Background())
}

func (i *MembershipIamPolicy) ToMembershipIamPolicyOutputWithContext(ctx context.Context) MembershipIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipIamPolicyOutput)
}

// MembershipIamPolicyArrayInput is an input type that accepts MembershipIamPolicyArray and MembershipIamPolicyArrayOutput values.
// You can construct a concrete instance of `MembershipIamPolicyArrayInput` via:
//
//	MembershipIamPolicyArray{ MembershipIamPolicyArgs{...} }
type MembershipIamPolicyArrayInput interface {
	pulumi.Input

	ToMembershipIamPolicyArrayOutput() MembershipIamPolicyArrayOutput
	ToMembershipIamPolicyArrayOutputWithContext(context.Context) MembershipIamPolicyArrayOutput
}

type MembershipIamPolicyArray []MembershipIamPolicyInput

func (MembershipIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MembershipIamPolicy)(nil)).Elem()
}

func (i MembershipIamPolicyArray) ToMembershipIamPolicyArrayOutput() MembershipIamPolicyArrayOutput {
	return i.ToMembershipIamPolicyArrayOutputWithContext(context.Background())
}

func (i MembershipIamPolicyArray) ToMembershipIamPolicyArrayOutputWithContext(ctx context.Context) MembershipIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipIamPolicyArrayOutput)
}

// MembershipIamPolicyMapInput is an input type that accepts MembershipIamPolicyMap and MembershipIamPolicyMapOutput values.
// You can construct a concrete instance of `MembershipIamPolicyMapInput` via:
//
//	MembershipIamPolicyMap{ "key": MembershipIamPolicyArgs{...} }
type MembershipIamPolicyMapInput interface {
	pulumi.Input

	ToMembershipIamPolicyMapOutput() MembershipIamPolicyMapOutput
	ToMembershipIamPolicyMapOutputWithContext(context.Context) MembershipIamPolicyMapOutput
}

type MembershipIamPolicyMap map[string]MembershipIamPolicyInput

func (MembershipIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MembershipIamPolicy)(nil)).Elem()
}

func (i MembershipIamPolicyMap) ToMembershipIamPolicyMapOutput() MembershipIamPolicyMapOutput {
	return i.ToMembershipIamPolicyMapOutputWithContext(context.Background())
}

func (i MembershipIamPolicyMap) ToMembershipIamPolicyMapOutputWithContext(ctx context.Context) MembershipIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipIamPolicyMapOutput)
}

type MembershipIamPolicyOutput struct{ *pulumi.OutputState }

func (MembershipIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipIamPolicy)(nil)).Elem()
}

func (o MembershipIamPolicyOutput) ToMembershipIamPolicyOutput() MembershipIamPolicyOutput {
	return o
}

func (o MembershipIamPolicyOutput) ToMembershipIamPolicyOutputWithContext(ctx context.Context) MembershipIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o MembershipIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MembershipIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Location of the membership.
// The default value is `global`.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o MembershipIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MembershipIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MembershipIamPolicyOutput) MembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v *MembershipIamPolicy) pulumi.StringOutput { return v.MembershipId }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o MembershipIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *MembershipIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MembershipIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MembershipIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type MembershipIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (MembershipIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MembershipIamPolicy)(nil)).Elem()
}

func (o MembershipIamPolicyArrayOutput) ToMembershipIamPolicyArrayOutput() MembershipIamPolicyArrayOutput {
	return o
}

func (o MembershipIamPolicyArrayOutput) ToMembershipIamPolicyArrayOutputWithContext(ctx context.Context) MembershipIamPolicyArrayOutput {
	return o
}

func (o MembershipIamPolicyArrayOutput) Index(i pulumi.IntInput) MembershipIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MembershipIamPolicy {
		return vs[0].([]*MembershipIamPolicy)[vs[1].(int)]
	}).(MembershipIamPolicyOutput)
}

type MembershipIamPolicyMapOutput struct{ *pulumi.OutputState }

func (MembershipIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MembershipIamPolicy)(nil)).Elem()
}

func (o MembershipIamPolicyMapOutput) ToMembershipIamPolicyMapOutput() MembershipIamPolicyMapOutput {
	return o
}

func (o MembershipIamPolicyMapOutput) ToMembershipIamPolicyMapOutputWithContext(ctx context.Context) MembershipIamPolicyMapOutput {
	return o
}

func (o MembershipIamPolicyMapOutput) MapIndex(k pulumi.StringInput) MembershipIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MembershipIamPolicy {
		return vs[0].(map[string]*MembershipIamPolicy)[vs[1].(string)]
	}).(MembershipIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipIamPolicyInput)(nil)).Elem(), &MembershipIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipIamPolicyArrayInput)(nil)).Elem(), MembershipIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipIamPolicyMapInput)(nil)).Elem(), MembershipIamPolicyMap{})
	pulumi.RegisterOutputType(MembershipIamPolicyOutput{})
	pulumi.RegisterOutputType(MembershipIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(MembershipIamPolicyMapOutput{})
}
