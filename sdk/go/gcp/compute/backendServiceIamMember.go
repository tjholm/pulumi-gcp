// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/global/backendServices/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Compute Engine backendservice IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendServiceIamMember:BackendServiceIamMember editor "projects/{{project}}/global/backendServices/{{backend_service}} roles/compute.admin user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendServiceIamMember:BackendServiceIamMember editor "projects/{{project}}/global/backendServices/{{backend_service}} roles/compute.admin"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendServiceIamMember:BackendServiceIamMember editor projects/{{project}}/global/backendServices/{{backend_service}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackendServiceIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *BackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*BackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackendServiceIamMember
	err := ctx.RegisterResource("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServiceIamMember gets an existing BackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*BackendServiceIamMember, error) {
	var resource BackendServiceIamMember
	err := ctx.ReadResource("gcp:compute/backendServiceIamMember:BackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServiceIamMember resources.
type backendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BackendServiceIamMemberCondition `pulumi:"condition"`
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
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BackendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamMemberConditionPtrInput
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamMemberState)(nil)).Elem()
}

type backendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BackendServiceIamMemberCondition `pulumi:"condition"`
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
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BackendServiceIamMember resource.
type BackendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BackendServiceIamMemberConditionPtrInput
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
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceIamMemberArgs)(nil)).Elem()
}

type BackendServiceIamMemberInput interface {
	pulumi.Input

	ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput
	ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput
}

func (*BackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamMember)(nil)).Elem()
}

func (i *BackendServiceIamMember) ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput {
	return i.ToBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i *BackendServiceIamMember) ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberOutput)
}

// BackendServiceIamMemberArrayInput is an input type that accepts BackendServiceIamMemberArray and BackendServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `BackendServiceIamMemberArrayInput` via:
//
//	BackendServiceIamMemberArray{ BackendServiceIamMemberArgs{...} }
type BackendServiceIamMemberArrayInput interface {
	pulumi.Input

	ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput
	ToBackendServiceIamMemberArrayOutputWithContext(context.Context) BackendServiceIamMemberArrayOutput
}

type BackendServiceIamMemberArray []BackendServiceIamMemberInput

func (BackendServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamMember)(nil)).Elem()
}

func (i BackendServiceIamMemberArray) ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput {
	return i.ToBackendServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i BackendServiceIamMemberArray) ToBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) BackendServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberArrayOutput)
}

// BackendServiceIamMemberMapInput is an input type that accepts BackendServiceIamMemberMap and BackendServiceIamMemberMapOutput values.
// You can construct a concrete instance of `BackendServiceIamMemberMapInput` via:
//
//	BackendServiceIamMemberMap{ "key": BackendServiceIamMemberArgs{...} }
type BackendServiceIamMemberMapInput interface {
	pulumi.Input

	ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput
	ToBackendServiceIamMemberMapOutputWithContext(context.Context) BackendServiceIamMemberMapOutput
}

type BackendServiceIamMemberMap map[string]BackendServiceIamMemberInput

func (BackendServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamMember)(nil)).Elem()
}

func (i BackendServiceIamMemberMap) ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput {
	return i.ToBackendServiceIamMemberMapOutputWithContext(context.Background())
}

func (i BackendServiceIamMemberMap) ToBackendServiceIamMemberMapOutputWithContext(ctx context.Context) BackendServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceIamMemberMapOutput)
}

type BackendServiceIamMemberOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberOutput) ToBackendServiceIamMemberOutput() BackendServiceIamMemberOutput {
	return o
}

func (o BackendServiceIamMemberOutput) ToBackendServiceIamMemberOutputWithContext(ctx context.Context) BackendServiceIamMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o BackendServiceIamMemberOutput) Condition() BackendServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) BackendServiceIamMemberConditionPtrOutput { return v.Condition }).(BackendServiceIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o BackendServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o BackendServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackendServiceIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackendServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.BackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o BackendServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type BackendServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberArrayOutput) ToBackendServiceIamMemberArrayOutput() BackendServiceIamMemberArrayOutput {
	return o
}

func (o BackendServiceIamMemberArrayOutput) ToBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) BackendServiceIamMemberArrayOutput {
	return o
}

func (o BackendServiceIamMemberArrayOutput) Index(i pulumi.IntInput) BackendServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendServiceIamMember {
		return vs[0].([]*BackendServiceIamMember)[vs[1].(int)]
	}).(BackendServiceIamMemberOutput)
}

type BackendServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (BackendServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceIamMember)(nil)).Elem()
}

func (o BackendServiceIamMemberMapOutput) ToBackendServiceIamMemberMapOutput() BackendServiceIamMemberMapOutput {
	return o
}

func (o BackendServiceIamMemberMapOutput) ToBackendServiceIamMemberMapOutputWithContext(ctx context.Context) BackendServiceIamMemberMapOutput {
	return o
}

func (o BackendServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) BackendServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendServiceIamMember {
		return vs[0].(map[string]*BackendServiceIamMember)[vs[1].(string)]
	}).(BackendServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberInput)(nil)).Elem(), &BackendServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberArrayInput)(nil)).Elem(), BackendServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceIamMemberMapInput)(nil)).Elem(), BackendServiceIamMemberMap{})
	pulumi.RegisterOutputType(BackendServiceIamMemberOutput{})
	pulumi.RegisterOutputType(BackendServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(BackendServiceIamMemberMapOutput{})
}
