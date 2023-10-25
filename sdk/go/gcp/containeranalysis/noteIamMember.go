// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
// ## google\_container\_analysis\_note\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
//				Project:    pulumi.Any(google_container_analysis_note.Note.Project),
//				Note:       pulumi.Any(google_container_analysis_note.Note.Name),
//				PolicyData: *pulumi.String(admin.PolicyData),
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
// ## google\_container\_analysis\_note\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamBinding(ctx, "binding", &containeranalysis.NoteIamBindingArgs{
//				Project: pulumi.Any(google_container_analysis_note.Note.Project),
//				Note:    pulumi.Any(google_container_analysis_note.Note.Name),
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
// ## google\_container\_analysis\_note\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/containeranalysis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containeranalysis.NewNoteIamMember(ctx, "member", &containeranalysis.NoteIamMemberArgs{
//				Project: pulumi.Any(google_container_analysis_note.Note.Project),
//				Note:    pulumi.Any(google_container_analysis_note.Note.Name),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/notes/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Container Registry note IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:containeranalysis/noteIamMember:NoteIamMember editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:containeranalysis/noteIamMember:NoteIamMember editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:containeranalysis/noteIamMember:NoteIamMember editor projects/{{project}}/notes/{{note}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NoteIamMember struct {
	pulumi.CustomResourceState

	Condition NoteIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Note pulumi.StringOutput `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewNoteIamMember registers a new resource with the given unique name, arguments, and options.
func NewNoteIamMember(ctx *pulumi.Context,
	name string, args *NoteIamMemberArgs, opts ...pulumi.ResourceOption) (*NoteIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Note == nil {
		return nil, errors.New("invalid value for required argument 'Note'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NoteIamMember
	err := ctx.RegisterResource("gcp:containeranalysis/noteIamMember:NoteIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNoteIamMember gets an existing NoteIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNoteIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteIamMemberState, opts ...pulumi.ResourceOption) (*NoteIamMember, error) {
	var resource NoteIamMember
	err := ctx.ReadResource("gcp:containeranalysis/noteIamMember:NoteIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NoteIamMember resources.
type noteIamMemberState struct {
	Condition *NoteIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Note *string `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type NoteIamMemberState struct {
	Condition NoteIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Note pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (NoteIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamMemberState)(nil)).Elem()
}

type noteIamMemberArgs struct {
	Condition *NoteIamMemberCondition `pulumi:"condition"`
	Member    string                  `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Note string `pulumi:"note"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a NoteIamMember resource.
type NoteIamMemberArgs struct {
	Condition NoteIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Note pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (NoteIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteIamMemberArgs)(nil)).Elem()
}

type NoteIamMemberInput interface {
	pulumi.Input

	ToNoteIamMemberOutput() NoteIamMemberOutput
	ToNoteIamMemberOutputWithContext(ctx context.Context) NoteIamMemberOutput
}

func (*NoteIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteIamMember)(nil)).Elem()
}

func (i *NoteIamMember) ToNoteIamMemberOutput() NoteIamMemberOutput {
	return i.ToNoteIamMemberOutputWithContext(context.Background())
}

func (i *NoteIamMember) ToNoteIamMemberOutputWithContext(ctx context.Context) NoteIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamMemberOutput)
}

func (i *NoteIamMember) ToOutput(ctx context.Context) pulumix.Output[*NoteIamMember] {
	return pulumix.Output[*NoteIamMember]{
		OutputState: i.ToNoteIamMemberOutputWithContext(ctx).OutputState,
	}
}

// NoteIamMemberArrayInput is an input type that accepts NoteIamMemberArray and NoteIamMemberArrayOutput values.
// You can construct a concrete instance of `NoteIamMemberArrayInput` via:
//
//	NoteIamMemberArray{ NoteIamMemberArgs{...} }
type NoteIamMemberArrayInput interface {
	pulumi.Input

	ToNoteIamMemberArrayOutput() NoteIamMemberArrayOutput
	ToNoteIamMemberArrayOutputWithContext(context.Context) NoteIamMemberArrayOutput
}

type NoteIamMemberArray []NoteIamMemberInput

func (NoteIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoteIamMember)(nil)).Elem()
}

func (i NoteIamMemberArray) ToNoteIamMemberArrayOutput() NoteIamMemberArrayOutput {
	return i.ToNoteIamMemberArrayOutputWithContext(context.Background())
}

func (i NoteIamMemberArray) ToNoteIamMemberArrayOutputWithContext(ctx context.Context) NoteIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamMemberArrayOutput)
}

func (i NoteIamMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*NoteIamMember] {
	return pulumix.Output[[]*NoteIamMember]{
		OutputState: i.ToNoteIamMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// NoteIamMemberMapInput is an input type that accepts NoteIamMemberMap and NoteIamMemberMapOutput values.
// You can construct a concrete instance of `NoteIamMemberMapInput` via:
//
//	NoteIamMemberMap{ "key": NoteIamMemberArgs{...} }
type NoteIamMemberMapInput interface {
	pulumi.Input

	ToNoteIamMemberMapOutput() NoteIamMemberMapOutput
	ToNoteIamMemberMapOutputWithContext(context.Context) NoteIamMemberMapOutput
}

type NoteIamMemberMap map[string]NoteIamMemberInput

func (NoteIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoteIamMember)(nil)).Elem()
}

func (i NoteIamMemberMap) ToNoteIamMemberMapOutput() NoteIamMemberMapOutput {
	return i.ToNoteIamMemberMapOutputWithContext(context.Background())
}

func (i NoteIamMemberMap) ToNoteIamMemberMapOutputWithContext(ctx context.Context) NoteIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteIamMemberMapOutput)
}

func (i NoteIamMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NoteIamMember] {
	return pulumix.Output[map[string]*NoteIamMember]{
		OutputState: i.ToNoteIamMemberMapOutputWithContext(ctx).OutputState,
	}
}

type NoteIamMemberOutput struct{ *pulumi.OutputState }

func (NoteIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoteIamMember)(nil)).Elem()
}

func (o NoteIamMemberOutput) ToNoteIamMemberOutput() NoteIamMemberOutput {
	return o
}

func (o NoteIamMemberOutput) ToNoteIamMemberOutputWithContext(ctx context.Context) NoteIamMemberOutput {
	return o
}

func (o NoteIamMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*NoteIamMember] {
	return pulumix.Output[*NoteIamMember]{
		OutputState: o.OutputState,
	}
}

func (o NoteIamMemberOutput) Condition() NoteIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *NoteIamMember) NoteIamMemberConditionPtrOutput { return v.Condition }).(NoteIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o NoteIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NoteIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o NoteIamMemberOutput) Note() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamMember) pulumi.StringOutput { return v.Note }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o NoteIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o NoteIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NoteIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type NoteIamMemberArrayOutput struct{ *pulumi.OutputState }

func (NoteIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NoteIamMember)(nil)).Elem()
}

func (o NoteIamMemberArrayOutput) ToNoteIamMemberArrayOutput() NoteIamMemberArrayOutput {
	return o
}

func (o NoteIamMemberArrayOutput) ToNoteIamMemberArrayOutputWithContext(ctx context.Context) NoteIamMemberArrayOutput {
	return o
}

func (o NoteIamMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NoteIamMember] {
	return pulumix.Output[[]*NoteIamMember]{
		OutputState: o.OutputState,
	}
}

func (o NoteIamMemberArrayOutput) Index(i pulumi.IntInput) NoteIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NoteIamMember {
		return vs[0].([]*NoteIamMember)[vs[1].(int)]
	}).(NoteIamMemberOutput)
}

type NoteIamMemberMapOutput struct{ *pulumi.OutputState }

func (NoteIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NoteIamMember)(nil)).Elem()
}

func (o NoteIamMemberMapOutput) ToNoteIamMemberMapOutput() NoteIamMemberMapOutput {
	return o
}

func (o NoteIamMemberMapOutput) ToNoteIamMemberMapOutputWithContext(ctx context.Context) NoteIamMemberMapOutput {
	return o
}

func (o NoteIamMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NoteIamMember] {
	return pulumix.Output[map[string]*NoteIamMember]{
		OutputState: o.OutputState,
	}
}

func (o NoteIamMemberMapOutput) MapIndex(k pulumi.StringInput) NoteIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NoteIamMember {
		return vs[0].(map[string]*NoteIamMember)[vs[1].(string)]
	}).(NoteIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamMemberInput)(nil)).Elem(), &NoteIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamMemberArrayInput)(nil)).Elem(), NoteIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NoteIamMemberMapInput)(nil)).Elem(), NoteIamMemberMap{})
	pulumi.RegisterOutputType(NoteIamMemberOutput{})
	pulumi.RegisterOutputType(NoteIamMemberArrayOutput{})
	pulumi.RegisterOutputType(NoteIamMemberMapOutput{})
}
