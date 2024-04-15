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
// * projects/{{project}}/global/backendBuckets/{{name}}
//
// * {{project}}/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Compute Engine backendbucket IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendBucketIamMember:BackendBucketIamMember editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendBucketIamMember:BackendBucketIamMember editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:compute/backendBucketIamMember:BackendBucketIamMember editor projects/{{project}}/global/backendBuckets/{{backend_bucket}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackendBucketIamMember struct {
	pulumi.CustomResourceState

	Condition BackendBucketIamMemberConditionPtrOutput `pulumi:"condition"`
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
	// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBackendBucketIamMember registers a new resource with the given unique name, arguments, and options.
func NewBackendBucketIamMember(ctx *pulumi.Context,
	name string, args *BackendBucketIamMemberArgs, opts ...pulumi.ResourceOption) (*BackendBucketIamMember, error) {
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
	var resource BackendBucketIamMember
	err := ctx.RegisterResource("gcp:compute/backendBucketIamMember:BackendBucketIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucketIamMember gets an existing BackendBucketIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucketIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketIamMemberState, opts ...pulumi.ResourceOption) (*BackendBucketIamMember, error) {
	var resource BackendBucketIamMember
	err := ctx.ReadResource("gcp:compute/backendBucketIamMember:BackendBucketIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucketIamMember resources.
type backendBucketIamMemberState struct {
	Condition *BackendBucketIamMemberCondition `pulumi:"condition"`
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
	// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BackendBucketIamMemberState struct {
	Condition BackendBucketIamMemberConditionPtrInput
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
	// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BackendBucketIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamMemberState)(nil)).Elem()
}

type backendBucketIamMemberArgs struct {
	Condition *BackendBucketIamMemberCondition `pulumi:"condition"`
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
	// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BackendBucketIamMember resource.
type BackendBucketIamMemberArgs struct {
	Condition BackendBucketIamMemberConditionPtrInput
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
	// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BackendBucketIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamMemberArgs)(nil)).Elem()
}

type BackendBucketIamMemberInput interface {
	pulumi.Input

	ToBackendBucketIamMemberOutput() BackendBucketIamMemberOutput
	ToBackendBucketIamMemberOutputWithContext(ctx context.Context) BackendBucketIamMemberOutput
}

func (*BackendBucketIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamMember)(nil)).Elem()
}

func (i *BackendBucketIamMember) ToBackendBucketIamMemberOutput() BackendBucketIamMemberOutput {
	return i.ToBackendBucketIamMemberOutputWithContext(context.Background())
}

func (i *BackendBucketIamMember) ToBackendBucketIamMemberOutputWithContext(ctx context.Context) BackendBucketIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamMemberOutput)
}

// BackendBucketIamMemberArrayInput is an input type that accepts BackendBucketIamMemberArray and BackendBucketIamMemberArrayOutput values.
// You can construct a concrete instance of `BackendBucketIamMemberArrayInput` via:
//
//	BackendBucketIamMemberArray{ BackendBucketIamMemberArgs{...} }
type BackendBucketIamMemberArrayInput interface {
	pulumi.Input

	ToBackendBucketIamMemberArrayOutput() BackendBucketIamMemberArrayOutput
	ToBackendBucketIamMemberArrayOutputWithContext(context.Context) BackendBucketIamMemberArrayOutput
}

type BackendBucketIamMemberArray []BackendBucketIamMemberInput

func (BackendBucketIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamMember)(nil)).Elem()
}

func (i BackendBucketIamMemberArray) ToBackendBucketIamMemberArrayOutput() BackendBucketIamMemberArrayOutput {
	return i.ToBackendBucketIamMemberArrayOutputWithContext(context.Background())
}

func (i BackendBucketIamMemberArray) ToBackendBucketIamMemberArrayOutputWithContext(ctx context.Context) BackendBucketIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamMemberArrayOutput)
}

// BackendBucketIamMemberMapInput is an input type that accepts BackendBucketIamMemberMap and BackendBucketIamMemberMapOutput values.
// You can construct a concrete instance of `BackendBucketIamMemberMapInput` via:
//
//	BackendBucketIamMemberMap{ "key": BackendBucketIamMemberArgs{...} }
type BackendBucketIamMemberMapInput interface {
	pulumi.Input

	ToBackendBucketIamMemberMapOutput() BackendBucketIamMemberMapOutput
	ToBackendBucketIamMemberMapOutputWithContext(context.Context) BackendBucketIamMemberMapOutput
}

type BackendBucketIamMemberMap map[string]BackendBucketIamMemberInput

func (BackendBucketIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamMember)(nil)).Elem()
}

func (i BackendBucketIamMemberMap) ToBackendBucketIamMemberMapOutput() BackendBucketIamMemberMapOutput {
	return i.ToBackendBucketIamMemberMapOutputWithContext(context.Background())
}

func (i BackendBucketIamMemberMap) ToBackendBucketIamMemberMapOutputWithContext(ctx context.Context) BackendBucketIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamMemberMapOutput)
}

type BackendBucketIamMemberOutput struct{ *pulumi.OutputState }

func (BackendBucketIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamMember)(nil)).Elem()
}

func (o BackendBucketIamMemberOutput) ToBackendBucketIamMemberOutput() BackendBucketIamMemberOutput {
	return o
}

func (o BackendBucketIamMemberOutput) ToBackendBucketIamMemberOutputWithContext(ctx context.Context) BackendBucketIamMemberOutput {
	return o
}

func (o BackendBucketIamMemberOutput) Condition() BackendBucketIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) BackendBucketIamMemberConditionPtrOutput { return v.Condition }).(BackendBucketIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o BackendBucketIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
func (o BackendBucketIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackendBucketIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackendBucketIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o BackendBucketIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type BackendBucketIamMemberArrayOutput struct{ *pulumi.OutputState }

func (BackendBucketIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamMember)(nil)).Elem()
}

func (o BackendBucketIamMemberArrayOutput) ToBackendBucketIamMemberArrayOutput() BackendBucketIamMemberArrayOutput {
	return o
}

func (o BackendBucketIamMemberArrayOutput) ToBackendBucketIamMemberArrayOutputWithContext(ctx context.Context) BackendBucketIamMemberArrayOutput {
	return o
}

func (o BackendBucketIamMemberArrayOutput) Index(i pulumi.IntInput) BackendBucketIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendBucketIamMember {
		return vs[0].([]*BackendBucketIamMember)[vs[1].(int)]
	}).(BackendBucketIamMemberOutput)
}

type BackendBucketIamMemberMapOutput struct{ *pulumi.OutputState }

func (BackendBucketIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamMember)(nil)).Elem()
}

func (o BackendBucketIamMemberMapOutput) ToBackendBucketIamMemberMapOutput() BackendBucketIamMemberMapOutput {
	return o
}

func (o BackendBucketIamMemberMapOutput) ToBackendBucketIamMemberMapOutputWithContext(ctx context.Context) BackendBucketIamMemberMapOutput {
	return o
}

func (o BackendBucketIamMemberMapOutput) MapIndex(k pulumi.StringInput) BackendBucketIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendBucketIamMember {
		return vs[0].(map[string]*BackendBucketIamMember)[vs[1].(string)]
	}).(BackendBucketIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamMemberInput)(nil)).Elem(), &BackendBucketIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamMemberArrayInput)(nil)).Elem(), BackendBucketIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamMemberMapInput)(nil)).Elem(), BackendBucketIamMemberMap{})
	pulumi.RegisterOutputType(BackendBucketIamMemberOutput{})
	pulumi.RegisterOutputType(BackendBucketIamMemberArrayOutput{})
	pulumi.RegisterOutputType(BackendBucketIamMemberMapOutput{})
}
