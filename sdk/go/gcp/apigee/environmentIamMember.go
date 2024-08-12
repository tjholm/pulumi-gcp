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
// ## apigee.EnvironmentIamPolicy
//
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
//
// ## apigee.EnvironmentIamBinding
//
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
//
// ## apigee.EnvironmentIamMember
//
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
//
// ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
// ---
//
// # IAM policy for Apigee Environment
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
// ## apigee.EnvironmentIamPolicy
//
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
//
// ## apigee.EnvironmentIamBinding
//
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
//
// ## apigee.EnvironmentIamMember
//
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
// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor "{{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor "{{org_id}}/environments/{{environment}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor {{org_id}}/environments/{{environment}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type EnvironmentIamMember struct {
	pulumi.CustomResourceState

	Condition EnvironmentIamMemberConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringOutput `pulumi:"envId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
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
	OrgId  pulumi.StringOutput `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewEnvironmentIamMember registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentIamMember(ctx *pulumi.Context,
	name string, args *EnvironmentIamMemberArgs, opts ...pulumi.ResourceOption) (*EnvironmentIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvironmentIamMember
	err := ctx.RegisterResource("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentIamMember gets an existing EnvironmentIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentIamMemberState, opts ...pulumi.ResourceOption) (*EnvironmentIamMember, error) {
	var resource EnvironmentIamMember
	err := ctx.ReadResource("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentIamMember resources.
type environmentIamMemberState struct {
	Condition *EnvironmentIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EnvId *string `pulumi:"envId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
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
	OrgId  *string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type EnvironmentIamMemberState struct {
	Condition EnvironmentIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
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
	OrgId  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (EnvironmentIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentIamMemberState)(nil)).Elem()
}

type environmentIamMemberArgs struct {
	Condition *EnvironmentIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EnvId string `pulumi:"envId"`
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
	OrgId  string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a EnvironmentIamMember resource.
type EnvironmentIamMemberArgs struct {
	Condition EnvironmentIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	EnvId pulumi.StringInput
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
	OrgId  pulumi.StringInput
	// The role that should be applied. Only one
	// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (EnvironmentIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentIamMemberArgs)(nil)).Elem()
}

type EnvironmentIamMemberInput interface {
	pulumi.Input

	ToEnvironmentIamMemberOutput() EnvironmentIamMemberOutput
	ToEnvironmentIamMemberOutputWithContext(ctx context.Context) EnvironmentIamMemberOutput
}

func (*EnvironmentIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentIamMember)(nil)).Elem()
}

func (i *EnvironmentIamMember) ToEnvironmentIamMemberOutput() EnvironmentIamMemberOutput {
	return i.ToEnvironmentIamMemberOutputWithContext(context.Background())
}

func (i *EnvironmentIamMember) ToEnvironmentIamMemberOutputWithContext(ctx context.Context) EnvironmentIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamMemberOutput)
}

// EnvironmentIamMemberArrayInput is an input type that accepts EnvironmentIamMemberArray and EnvironmentIamMemberArrayOutput values.
// You can construct a concrete instance of `EnvironmentIamMemberArrayInput` via:
//
//	EnvironmentIamMemberArray{ EnvironmentIamMemberArgs{...} }
type EnvironmentIamMemberArrayInput interface {
	pulumi.Input

	ToEnvironmentIamMemberArrayOutput() EnvironmentIamMemberArrayOutput
	ToEnvironmentIamMemberArrayOutputWithContext(context.Context) EnvironmentIamMemberArrayOutput
}

type EnvironmentIamMemberArray []EnvironmentIamMemberInput

func (EnvironmentIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentIamMember)(nil)).Elem()
}

func (i EnvironmentIamMemberArray) ToEnvironmentIamMemberArrayOutput() EnvironmentIamMemberArrayOutput {
	return i.ToEnvironmentIamMemberArrayOutputWithContext(context.Background())
}

func (i EnvironmentIamMemberArray) ToEnvironmentIamMemberArrayOutputWithContext(ctx context.Context) EnvironmentIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamMemberArrayOutput)
}

// EnvironmentIamMemberMapInput is an input type that accepts EnvironmentIamMemberMap and EnvironmentIamMemberMapOutput values.
// You can construct a concrete instance of `EnvironmentIamMemberMapInput` via:
//
//	EnvironmentIamMemberMap{ "key": EnvironmentIamMemberArgs{...} }
type EnvironmentIamMemberMapInput interface {
	pulumi.Input

	ToEnvironmentIamMemberMapOutput() EnvironmentIamMemberMapOutput
	ToEnvironmentIamMemberMapOutputWithContext(context.Context) EnvironmentIamMemberMapOutput
}

type EnvironmentIamMemberMap map[string]EnvironmentIamMemberInput

func (EnvironmentIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentIamMember)(nil)).Elem()
}

func (i EnvironmentIamMemberMap) ToEnvironmentIamMemberMapOutput() EnvironmentIamMemberMapOutput {
	return i.ToEnvironmentIamMemberMapOutputWithContext(context.Background())
}

func (i EnvironmentIamMemberMap) ToEnvironmentIamMemberMapOutputWithContext(ctx context.Context) EnvironmentIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentIamMemberMapOutput)
}

type EnvironmentIamMemberOutput struct{ *pulumi.OutputState }

func (EnvironmentIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentIamMember)(nil)).Elem()
}

func (o EnvironmentIamMemberOutput) ToEnvironmentIamMemberOutput() EnvironmentIamMemberOutput {
	return o
}

func (o EnvironmentIamMemberOutput) ToEnvironmentIamMemberOutputWithContext(ctx context.Context) EnvironmentIamMemberOutput {
	return o
}

func (o EnvironmentIamMemberOutput) Condition() EnvironmentIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) EnvironmentIamMemberConditionPtrOutput { return v.Condition }).(EnvironmentIamMemberConditionPtrOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o EnvironmentIamMemberOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o EnvironmentIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o EnvironmentIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o EnvironmentIamMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o EnvironmentIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type EnvironmentIamMemberArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentIamMember)(nil)).Elem()
}

func (o EnvironmentIamMemberArrayOutput) ToEnvironmentIamMemberArrayOutput() EnvironmentIamMemberArrayOutput {
	return o
}

func (o EnvironmentIamMemberArrayOutput) ToEnvironmentIamMemberArrayOutputWithContext(ctx context.Context) EnvironmentIamMemberArrayOutput {
	return o
}

func (o EnvironmentIamMemberArrayOutput) Index(i pulumi.IntInput) EnvironmentIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvironmentIamMember {
		return vs[0].([]*EnvironmentIamMember)[vs[1].(int)]
	}).(EnvironmentIamMemberOutput)
}

type EnvironmentIamMemberMapOutput struct{ *pulumi.OutputState }

func (EnvironmentIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentIamMember)(nil)).Elem()
}

func (o EnvironmentIamMemberMapOutput) ToEnvironmentIamMemberMapOutput() EnvironmentIamMemberMapOutput {
	return o
}

func (o EnvironmentIamMemberMapOutput) ToEnvironmentIamMemberMapOutputWithContext(ctx context.Context) EnvironmentIamMemberMapOutput {
	return o
}

func (o EnvironmentIamMemberMapOutput) MapIndex(k pulumi.StringInput) EnvironmentIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvironmentIamMember {
		return vs[0].(map[string]*EnvironmentIamMember)[vs[1].(string)]
	}).(EnvironmentIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamMemberInput)(nil)).Elem(), &EnvironmentIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamMemberArrayInput)(nil)).Elem(), EnvironmentIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentIamMemberMapInput)(nil)).Elem(), EnvironmentIamMemberMap{})
	pulumi.RegisterOutputType(EnvironmentIamMemberOutput{})
	pulumi.RegisterOutputType(EnvironmentIamMemberArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentIamMemberMapOutput{})
}
