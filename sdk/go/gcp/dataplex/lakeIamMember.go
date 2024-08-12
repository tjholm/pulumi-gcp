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

// Three different resources help you manage your IAM policy for Dataplex Lake. Each of these resources serves a different use case:
//
// * `dataplex.LakeIamPolicy`: Authoritative. Sets the IAM policy for the lake and replaces any existing policy already attached.
// * `dataplex.LakeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the lake are preserved.
// * `dataplex.LakeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the lake are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.LakeIamPolicy`: Retrieves the IAM policy for the lake
//
// > **Note:** `dataplex.LakeIamPolicy` **cannot** be used in conjunction with `dataplex.LakeIamBinding` and `dataplex.LakeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.LakeIamBinding` resources **can be** used in conjunction with `dataplex.LakeIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.LakeIamPolicy
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
//			_, err = dataplex.NewLakeIamPolicy(ctx, "policy", &dataplex.LakeIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Location:   pulumi.Any(example.Location),
//				Lake:       pulumi.Any(example.Name),
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
// ## dataplex.LakeIamBinding
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
//			_, err := dataplex.NewLakeIamBinding(ctx, "binding", &dataplex.LakeIamBindingArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Name),
//				Role:     pulumi.String("roles/viewer"),
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
// ## dataplex.LakeIamMember
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
//			_, err := dataplex.NewLakeIamMember(ctx, "member", &dataplex.LakeIamMemberArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// # IAM policy for Dataplex Lake
// Three different resources help you manage your IAM policy for Dataplex Lake. Each of these resources serves a different use case:
//
// * `dataplex.LakeIamPolicy`: Authoritative. Sets the IAM policy for the lake and replaces any existing policy already attached.
// * `dataplex.LakeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the lake are preserved.
// * `dataplex.LakeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the lake are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `dataplex.LakeIamPolicy`: Retrieves the IAM policy for the lake
//
// > **Note:** `dataplex.LakeIamPolicy` **cannot** be used in conjunction with `dataplex.LakeIamBinding` and `dataplex.LakeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `dataplex.LakeIamBinding` resources **can be** used in conjunction with `dataplex.LakeIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## dataplex.LakeIamPolicy
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
//			_, err = dataplex.NewLakeIamPolicy(ctx, "policy", &dataplex.LakeIamPolicyArgs{
//				Project:    pulumi.Any(example.Project),
//				Location:   pulumi.Any(example.Location),
//				Lake:       pulumi.Any(example.Name),
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
// ## dataplex.LakeIamBinding
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
//			_, err := dataplex.NewLakeIamBinding(ctx, "binding", &dataplex.LakeIamBindingArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Name),
//				Role:     pulumi.String("roles/viewer"),
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
// ## dataplex.LakeIamMember
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
//			_, err := dataplex.NewLakeIamMember(ctx, "member", &dataplex.LakeIamMemberArgs{
//				Project:  pulumi.Any(example.Project),
//				Location: pulumi.Any(example.Location),
//				Lake:     pulumi.Any(example.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/locations/{{location}}/lakes/{{name}}
//
// * {{project}}/{{location}}/{{name}}
//
// * {{location}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Dataplex lake IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/lakeIamMember:LakeIamMember editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/lakeIamMember:LakeIamMember editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:dataplex/lakeIamMember:LakeIamMember editor projects/{{project}}/locations/{{location}}/lakes/{{lake}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type LakeIamMember struct {
	pulumi.CustomResourceState

	Condition LakeIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Lake     pulumi.StringOutput `pulumi:"lake"`
	Location pulumi.StringOutput `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewLakeIamMember registers a new resource with the given unique name, arguments, and options.
func NewLakeIamMember(ctx *pulumi.Context,
	name string, args *LakeIamMemberArgs, opts ...pulumi.ResourceOption) (*LakeIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LakeIamMember
	err := ctx.RegisterResource("gcp:dataplex/lakeIamMember:LakeIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLakeIamMember gets an existing LakeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLakeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LakeIamMemberState, opts ...pulumi.ResourceOption) (*LakeIamMember, error) {
	var resource LakeIamMember
	err := ctx.ReadResource("gcp:dataplex/lakeIamMember:LakeIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LakeIamMember resources.
type lakeIamMemberState struct {
	Condition *LakeIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Lake     *string `pulumi:"lake"`
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type LakeIamMemberState struct {
	Condition LakeIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Lake     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (LakeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*lakeIamMemberState)(nil)).Elem()
}

type lakeIamMemberArgs struct {
	Condition *LakeIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Lake     string  `pulumi:"lake"`
	Location *string `pulumi:"location"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a LakeIamMember resource.
type LakeIamMemberArgs struct {
	Condition LakeIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Lake     pulumi.StringInput
	Location pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Member pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (LakeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lakeIamMemberArgs)(nil)).Elem()
}

type LakeIamMemberInput interface {
	pulumi.Input

	ToLakeIamMemberOutput() LakeIamMemberOutput
	ToLakeIamMemberOutputWithContext(ctx context.Context) LakeIamMemberOutput
}

func (*LakeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**LakeIamMember)(nil)).Elem()
}

func (i *LakeIamMember) ToLakeIamMemberOutput() LakeIamMemberOutput {
	return i.ToLakeIamMemberOutputWithContext(context.Background())
}

func (i *LakeIamMember) ToLakeIamMemberOutputWithContext(ctx context.Context) LakeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LakeIamMemberOutput)
}

// LakeIamMemberArrayInput is an input type that accepts LakeIamMemberArray and LakeIamMemberArrayOutput values.
// You can construct a concrete instance of `LakeIamMemberArrayInput` via:
//
//	LakeIamMemberArray{ LakeIamMemberArgs{...} }
type LakeIamMemberArrayInput interface {
	pulumi.Input

	ToLakeIamMemberArrayOutput() LakeIamMemberArrayOutput
	ToLakeIamMemberArrayOutputWithContext(context.Context) LakeIamMemberArrayOutput
}

type LakeIamMemberArray []LakeIamMemberInput

func (LakeIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LakeIamMember)(nil)).Elem()
}

func (i LakeIamMemberArray) ToLakeIamMemberArrayOutput() LakeIamMemberArrayOutput {
	return i.ToLakeIamMemberArrayOutputWithContext(context.Background())
}

func (i LakeIamMemberArray) ToLakeIamMemberArrayOutputWithContext(ctx context.Context) LakeIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LakeIamMemberArrayOutput)
}

// LakeIamMemberMapInput is an input type that accepts LakeIamMemberMap and LakeIamMemberMapOutput values.
// You can construct a concrete instance of `LakeIamMemberMapInput` via:
//
//	LakeIamMemberMap{ "key": LakeIamMemberArgs{...} }
type LakeIamMemberMapInput interface {
	pulumi.Input

	ToLakeIamMemberMapOutput() LakeIamMemberMapOutput
	ToLakeIamMemberMapOutputWithContext(context.Context) LakeIamMemberMapOutput
}

type LakeIamMemberMap map[string]LakeIamMemberInput

func (LakeIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LakeIamMember)(nil)).Elem()
}

func (i LakeIamMemberMap) ToLakeIamMemberMapOutput() LakeIamMemberMapOutput {
	return i.ToLakeIamMemberMapOutputWithContext(context.Background())
}

func (i LakeIamMemberMap) ToLakeIamMemberMapOutputWithContext(ctx context.Context) LakeIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LakeIamMemberMapOutput)
}

type LakeIamMemberOutput struct{ *pulumi.OutputState }

func (LakeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LakeIamMember)(nil)).Elem()
}

func (o LakeIamMemberOutput) ToLakeIamMemberOutput() LakeIamMemberOutput {
	return o
}

func (o LakeIamMemberOutput) ToLakeIamMemberOutputWithContext(ctx context.Context) LakeIamMemberOutput {
	return o
}

func (o LakeIamMemberOutput) Condition() LakeIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *LakeIamMember) LakeIamMemberConditionPtrOutput { return v.Condition }).(LakeIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o LakeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o LakeIamMemberOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

func (o LakeIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`.
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o LakeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o LakeIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataplex.LakeIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o LakeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *LakeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type LakeIamMemberArrayOutput struct{ *pulumi.OutputState }

func (LakeIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LakeIamMember)(nil)).Elem()
}

func (o LakeIamMemberArrayOutput) ToLakeIamMemberArrayOutput() LakeIamMemberArrayOutput {
	return o
}

func (o LakeIamMemberArrayOutput) ToLakeIamMemberArrayOutputWithContext(ctx context.Context) LakeIamMemberArrayOutput {
	return o
}

func (o LakeIamMemberArrayOutput) Index(i pulumi.IntInput) LakeIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LakeIamMember {
		return vs[0].([]*LakeIamMember)[vs[1].(int)]
	}).(LakeIamMemberOutput)
}

type LakeIamMemberMapOutput struct{ *pulumi.OutputState }

func (LakeIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LakeIamMember)(nil)).Elem()
}

func (o LakeIamMemberMapOutput) ToLakeIamMemberMapOutput() LakeIamMemberMapOutput {
	return o
}

func (o LakeIamMemberMapOutput) ToLakeIamMemberMapOutputWithContext(ctx context.Context) LakeIamMemberMapOutput {
	return o
}

func (o LakeIamMemberMapOutput) MapIndex(k pulumi.StringInput) LakeIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LakeIamMember {
		return vs[0].(map[string]*LakeIamMember)[vs[1].(string)]
	}).(LakeIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LakeIamMemberInput)(nil)).Elem(), &LakeIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*LakeIamMemberArrayInput)(nil)).Elem(), LakeIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LakeIamMemberMapInput)(nil)).Elem(), LakeIamMemberMap{})
	pulumi.RegisterOutputType(LakeIamMemberOutput{})
	pulumi.RegisterOutputType(LakeIamMemberArrayOutput{})
	pulumi.RegisterOutputType(LakeIamMemberMapOutput{})
}
