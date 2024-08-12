// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:
//
// * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
// * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
// * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note
//
// > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.
//
// > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## containeranalysis.NoteIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
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
//						Role: "roles/containeranalysis.notes.occurrences.viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = containeranalysis.NewNoteIamPolicy(ctx, "policy", &containeranalysis.NoteIamPolicyArgs{
//				Project:    pulumi.Any(note.Project),
//				Note:       pulumi.Any(note.Name),
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
// ## containeranalysis.NoteIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamBinding(ctx, "binding", &containeranalysis.NoteIamBindingArgs{
//				Project: pulumi.Any(note.Project),
//				Note:    pulumi.Any(note.Name),
//				Role:    pulumi.String("roles/containeranalysis.notes.occurrences.viewer"),
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
// ## containeranalysis.NoteIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamMember(ctx, "member", &containeranalysis.NoteIamMemberArgs{
//				Project: pulumi.Any(note.Project),
//				Note:    pulumi.Any(note.Name),
//				Role:    pulumi.String("roles/containeranalysis.notes.occurrences.viewer"),
//				Member:  pulumi.String("user:jane@example.com"),
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
// # IAM policy for Container Registry Note
// Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:
//
// * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
// * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
// * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note
//
// > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.
//
// > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## containeranalysis.NoteIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
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
//						Role: "roles/containeranalysis.notes.occurrences.viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = containeranalysis.NewNoteIamPolicy(ctx, "policy", &containeranalysis.NoteIamPolicyArgs{
//				Project:    pulumi.Any(note.Project),
//				Note:       pulumi.Any(note.Name),
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
// ## containeranalysis.NoteIamBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamBinding(ctx, "binding", &containeranalysis.NoteIamBindingArgs{
//				Project: pulumi.Any(note.Project),
//				Note:    pulumi.Any(note.Name),
//				Role:    pulumi.String("roles/containeranalysis.notes.occurrences.viewer"),
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
// ## containeranalysis.NoteIamMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamMember(ctx, "member", &containeranalysis.NoteIamMemberArgs{
//				Project: pulumi.Any(note.Project),
//				Note:    pulumi.Any(note.Name),
//				Role:    pulumi.String("roles/containeranalysis.notes.occurrences.viewer"),
//				Member:  pulumi.String("user:jane@example.com"),
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
// * projects/{{project}}/notes/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Container Registry note IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor projects/{{project}}/notes/{{note}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NoteIamBinding struct {
	pulumi.CustomResourceState

	Condition NoteIamBindingConditionPtrOutput `pulumi:"condition"`
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
	Note pulumi.StringOutput `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewNoteIamBinding registers a new resource with the given unique name, arguments, and options.
func NewNoteIamBinding(ctx *pulumi.Context,
	name string, args *NoteIamBindingArgs, opts ...pulumi.ResourceOption) (*NoteIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Note == nil {
		return nil, errors.New("invalid value for required argument 'Note'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NoteIamBinding
	err := ctx.RegisterResource("gcp:containeranalysis/noteIamBinding:NoteIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNoteIamBinding gets an existing NoteIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNoteIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteIamBindingState, opts ...pulumi.ResourceOption) (*NoteIamBinding, error) {
	var resource NoteIamBinding
	err := ctx.ReadResource("gcp:containeranalysis/noteIamBinding:NoteIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NoteIamBinding resources.
type noteIamBindingState struct {
	Condition *NoteIamBindingCondition `pulumi:"condition"`
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
	Note *string `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type NoteIamBindingState struct {
	Condition NoteIamBindingConditionPtrInput
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
	Note pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (NoteIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamBindingState)(nil)).Elem()
}

type noteIamBindingArgs struct {
	Condition *NoteIamBindingCondition `pulumi:"condition"`
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
	Note string `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a NoteIamBinding resource.
type NoteIamBindingArgs struct {
	Condition NoteIamBindingConditionPtrInput
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
	Note pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (NoteIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamBindingArgs)(nil)).Elem()
}

type NoteIamBindingInput interface {
	pulumi.Input

	ToNoteIamBindingOutput() NoteIamBindingOutput
	ToNoteIamBindingOutputWithContext(ctx context.Context) NoteIamBindingOutput
}

func (*NoteIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteIamBinding)(nil)).Elem()
}

func (i *NoteIamBinding) ToNoteIamBindingOutput() NoteIamBindingOutput {
	return i.ToNoteIamBindingOutputWithContext(context.Background())
}

func (i *NoteIamBinding) ToNoteIamBindingOutputWithContext(ctx context.Context) NoteIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamBindingOutput)
}

// NoteIamBindingArrayInput is an input type that accepts NoteIamBindingArray and NoteIamBindingArrayOutput values.
// You can construct a concrete instance of `NoteIamBindingArrayInput` via:
//
//	NoteIamBindingArray{ NoteIamBindingArgs{...} }
type NoteIamBindingArrayInput interface {
	pulumi.Input

	ToNoteIamBindingArrayOutput() NoteIamBindingArrayOutput
	ToNoteIamBindingArrayOutputWithContext(context.Context) NoteIamBindingArrayOutput
}

type NoteIamBindingArray []NoteIamBindingInput

func (NoteIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoteIamBinding)(nil)).Elem()
}

func (i NoteIamBindingArray) ToNoteIamBindingArrayOutput() NoteIamBindingArrayOutput {
	return i.ToNoteIamBindingArrayOutputWithContext(context.Background())
}

func (i NoteIamBindingArray) ToNoteIamBindingArrayOutputWithContext(ctx context.Context) NoteIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamBindingArrayOutput)
}

// NoteIamBindingMapInput is an input type that accepts NoteIamBindingMap and NoteIamBindingMapOutput values.
// You can construct a concrete instance of `NoteIamBindingMapInput` via:
//
//	NoteIamBindingMap{ "key": NoteIamBindingArgs{...} }
type NoteIamBindingMapInput interface {
	pulumi.Input

	ToNoteIamBindingMapOutput() NoteIamBindingMapOutput
	ToNoteIamBindingMapOutputWithContext(context.Context) NoteIamBindingMapOutput
}

type NoteIamBindingMap map[string]NoteIamBindingInput

func (NoteIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoteIamBinding)(nil)).Elem()
}

func (i NoteIamBindingMap) ToNoteIamBindingMapOutput() NoteIamBindingMapOutput {
	return i.ToNoteIamBindingMapOutputWithContext(context.Background())
}

func (i NoteIamBindingMap) ToNoteIamBindingMapOutputWithContext(ctx context.Context) NoteIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamBindingMapOutput)
}

type NoteIamBindingOutput struct{ *pulumi.OutputState }

func (NoteIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteIamBinding)(nil)).Elem()
}

func (o NoteIamBindingOutput) ToNoteIamBindingOutput() NoteIamBindingOutput {
	return o
}

func (o NoteIamBindingOutput) ToNoteIamBindingOutputWithContext(ctx context.Context) NoteIamBindingOutput {
	return o
}

func (o NoteIamBindingOutput) Condition() NoteIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *NoteIamBinding) NoteIamBindingConditionPtrOutput { return v.Condition }).(NoteIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o NoteIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o NoteIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NoteIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o NoteIamBindingOutput) Note() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamBinding) pulumi.StringOutput { return v.Note }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o NoteIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o NoteIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type NoteIamBindingArrayOutput struct{ *pulumi.OutputState }

func (NoteIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoteIamBinding)(nil)).Elem()
}

func (o NoteIamBindingArrayOutput) ToNoteIamBindingArrayOutput() NoteIamBindingArrayOutput {
	return o
}

func (o NoteIamBindingArrayOutput) ToNoteIamBindingArrayOutputWithContext(ctx context.Context) NoteIamBindingArrayOutput {
	return o
}

func (o NoteIamBindingArrayOutput) Index(i pulumi.IntInput) NoteIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NoteIamBinding {
		return vs[0].([]*NoteIamBinding)[vs[1].(int)]
	}).(NoteIamBindingOutput)
}

type NoteIamBindingMapOutput struct{ *pulumi.OutputState }

func (NoteIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoteIamBinding)(nil)).Elem()
}

func (o NoteIamBindingMapOutput) ToNoteIamBindingMapOutput() NoteIamBindingMapOutput {
	return o
}

func (o NoteIamBindingMapOutput) ToNoteIamBindingMapOutputWithContext(ctx context.Context) NoteIamBindingMapOutput {
	return o
}

func (o NoteIamBindingMapOutput) MapIndex(k pulumi.StringInput) NoteIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NoteIamBinding {
		return vs[0].(map[string]*NoteIamBinding)[vs[1].(string)]
	}).(NoteIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamBindingInput)(nil)).Elem(), &NoteIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamBindingArrayInput)(nil)).Elem(), NoteIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamBindingMapInput)(nil)).Elem(), NoteIamBindingMap{})
	pulumi.RegisterOutputType(NoteIamBindingOutput{})
	pulumi.RegisterOutputType(NoteIamBindingArrayOutput{})
	pulumi.RegisterOutputType(NoteIamBindingMapOutput{})
}
