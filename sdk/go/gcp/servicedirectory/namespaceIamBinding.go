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
//
// ## servicedirectory.NamespaceIamBinding
//
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
//
// ## servicedirectory.NamespaceIamMember
//
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
//
// ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
// ---
//
// # IAM policy for Service Directory Namespace
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
//
// ## servicedirectory.NamespaceIamBinding
//
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
//
// ## servicedirectory.NamespaceIamMember
//
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
// $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NamespaceIamBinding struct {
	pulumi.CustomResourceState

	Condition NamespaceIamBindingConditionPtrOutput `pulumi:"condition"`
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
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewNamespaceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewNamespaceIamBinding(ctx *pulumi.Context,
	name string, args *NamespaceIamBindingArgs, opts ...pulumi.ResourceOption) (*NamespaceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamespaceIamBinding
	err := ctx.RegisterResource("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceIamBinding gets an existing NamespaceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIamBindingState, opts ...pulumi.ResourceOption) (*NamespaceIamBinding, error) {
	var resource NamespaceIamBinding
	err := ctx.ReadResource("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceIamBinding resources.
type namespaceIamBindingState struct {
	Condition *NamespaceIamBindingCondition `pulumi:"condition"`
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
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type NamespaceIamBindingState struct {
	Condition NamespaceIamBindingConditionPtrInput
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
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (NamespaceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamBindingState)(nil)).Elem()
}

type namespaceIamBindingArgs struct {
	Condition *NamespaceIamBindingCondition `pulumi:"condition"`
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
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a NamespaceIamBinding resource.
type NamespaceIamBindingArgs struct {
	Condition NamespaceIamBindingConditionPtrInput
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
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (NamespaceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamBindingArgs)(nil)).Elem()
}

type NamespaceIamBindingInput interface {
	pulumi.Input

	ToNamespaceIamBindingOutput() NamespaceIamBindingOutput
	ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput
}

func (*NamespaceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamBinding)(nil)).Elem()
}

func (i *NamespaceIamBinding) ToNamespaceIamBindingOutput() NamespaceIamBindingOutput {
	return i.ToNamespaceIamBindingOutputWithContext(context.Background())
}

func (i *NamespaceIamBinding) ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamBindingOutput)
}

// NamespaceIamBindingArrayInput is an input type that accepts NamespaceIamBindingArray and NamespaceIamBindingArrayOutput values.
// You can construct a concrete instance of `NamespaceIamBindingArrayInput` via:
//
//	NamespaceIamBindingArray{ NamespaceIamBindingArgs{...} }
type NamespaceIamBindingArrayInput interface {
	pulumi.Input

	ToNamespaceIamBindingArrayOutput() NamespaceIamBindingArrayOutput
	ToNamespaceIamBindingArrayOutputWithContext(context.Context) NamespaceIamBindingArrayOutput
}

type NamespaceIamBindingArray []NamespaceIamBindingInput

func (NamespaceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceIamBinding)(nil)).Elem()
}

func (i NamespaceIamBindingArray) ToNamespaceIamBindingArrayOutput() NamespaceIamBindingArrayOutput {
	return i.ToNamespaceIamBindingArrayOutputWithContext(context.Background())
}

func (i NamespaceIamBindingArray) ToNamespaceIamBindingArrayOutputWithContext(ctx context.Context) NamespaceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamBindingArrayOutput)
}

// NamespaceIamBindingMapInput is an input type that accepts NamespaceIamBindingMap and NamespaceIamBindingMapOutput values.
// You can construct a concrete instance of `NamespaceIamBindingMapInput` via:
//
//	NamespaceIamBindingMap{ "key": NamespaceIamBindingArgs{...} }
type NamespaceIamBindingMapInput interface {
	pulumi.Input

	ToNamespaceIamBindingMapOutput() NamespaceIamBindingMapOutput
	ToNamespaceIamBindingMapOutputWithContext(context.Context) NamespaceIamBindingMapOutput
}

type NamespaceIamBindingMap map[string]NamespaceIamBindingInput

func (NamespaceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceIamBinding)(nil)).Elem()
}

func (i NamespaceIamBindingMap) ToNamespaceIamBindingMapOutput() NamespaceIamBindingMapOutput {
	return i.ToNamespaceIamBindingMapOutputWithContext(context.Background())
}

func (i NamespaceIamBindingMap) ToNamespaceIamBindingMapOutputWithContext(ctx context.Context) NamespaceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamBindingMapOutput)
}

type NamespaceIamBindingOutput struct{ *pulumi.OutputState }

func (NamespaceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamBinding)(nil)).Elem()
}

func (o NamespaceIamBindingOutput) ToNamespaceIamBindingOutput() NamespaceIamBindingOutput {
	return o
}

func (o NamespaceIamBindingOutput) ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput {
	return o
}

func (o NamespaceIamBindingOutput) Condition() NamespaceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *NamespaceIamBinding) NamespaceIamBindingConditionPtrOutput { return v.Condition }).(NamespaceIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o NamespaceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o NamespaceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o NamespaceIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o NamespaceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type NamespaceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (NamespaceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamespaceIamBinding)(nil)).Elem()
}

func (o NamespaceIamBindingArrayOutput) ToNamespaceIamBindingArrayOutput() NamespaceIamBindingArrayOutput {
	return o
}

func (o NamespaceIamBindingArrayOutput) ToNamespaceIamBindingArrayOutputWithContext(ctx context.Context) NamespaceIamBindingArrayOutput {
	return o
}

func (o NamespaceIamBindingArrayOutput) Index(i pulumi.IntInput) NamespaceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamespaceIamBinding {
		return vs[0].([]*NamespaceIamBinding)[vs[1].(int)]
	}).(NamespaceIamBindingOutput)
}

type NamespaceIamBindingMapOutput struct{ *pulumi.OutputState }

func (NamespaceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamespaceIamBinding)(nil)).Elem()
}

func (o NamespaceIamBindingMapOutput) ToNamespaceIamBindingMapOutput() NamespaceIamBindingMapOutput {
	return o
}

func (o NamespaceIamBindingMapOutput) ToNamespaceIamBindingMapOutputWithContext(ctx context.Context) NamespaceIamBindingMapOutput {
	return o
}

func (o NamespaceIamBindingMapOutput) MapIndex(k pulumi.StringInput) NamespaceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamespaceIamBinding {
		return vs[0].(map[string]*NamespaceIamBinding)[vs[1].(string)]
	}).(NamespaceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamBindingInput)(nil)).Elem(), &NamespaceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamBindingArrayInput)(nil)).Elem(), NamespaceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceIamBindingMapInput)(nil)).Elem(), NamespaceIamBindingMap{})
	pulumi.RegisterOutputType(NamespaceIamBindingOutput{})
	pulumi.RegisterOutputType(NamespaceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(NamespaceIamBindingMapOutput{})
}
