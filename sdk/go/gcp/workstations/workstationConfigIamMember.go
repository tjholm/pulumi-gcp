// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workstations

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
// * projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}
//
// * {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
//
// * {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}
//
// * {{workstation_config_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Workstations workstationconfig IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationConfigIamMember:WorkstationConfigIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationConfigIamMember:WorkstationConfigIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationConfigIamMember:WorkstationConfigIamMember editor projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WorkstationConfigIamMember struct {
	pulumi.CustomResourceState

	Condition WorkstationConfigIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringOutput `pulumi:"role"`
	WorkstationClusterId pulumi.StringOutput `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringOutput `pulumi:"workstationConfigId"`
}

// NewWorkstationConfigIamMember registers a new resource with the given unique name, arguments, and options.
func NewWorkstationConfigIamMember(ctx *pulumi.Context,
	name string, args *WorkstationConfigIamMemberArgs, opts ...pulumi.ResourceOption) (*WorkstationConfigIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkstationClusterId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationClusterId'")
	}
	if args.WorkstationConfigId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationConfigId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkstationConfigIamMember
	err := ctx.RegisterResource("gcp:workstations/workstationConfigIamMember:WorkstationConfigIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationConfigIamMember gets an existing WorkstationConfigIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationConfigIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationConfigIamMemberState, opts ...pulumi.ResourceOption) (*WorkstationConfigIamMember, error) {
	var resource WorkstationConfigIamMember
	err := ctx.ReadResource("gcp:workstations/workstationConfigIamMember:WorkstationConfigIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationConfigIamMember resources.
type workstationConfigIamMemberState struct {
	Condition *WorkstationConfigIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 *string `pulumi:"role"`
	WorkstationClusterId *string `pulumi:"workstationClusterId"`
	WorkstationConfigId  *string `pulumi:"workstationConfigId"`
}

type WorkstationConfigIamMemberState struct {
	Condition WorkstationConfigIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringPtrInput
	WorkstationClusterId pulumi.StringPtrInput
	WorkstationConfigId  pulumi.StringPtrInput
}

func (WorkstationConfigIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationConfigIamMemberState)(nil)).Elem()
}

type workstationConfigIamMemberArgs struct {
	Condition *WorkstationConfigIamMemberCondition `pulumi:"condition"`
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 string `pulumi:"role"`
	WorkstationClusterId string `pulumi:"workstationClusterId"`
	WorkstationConfigId  string `pulumi:"workstationConfigId"`
}

// The set of arguments for constructing a WorkstationConfigIamMember resource.
type WorkstationConfigIamMemberArgs struct {
	Condition WorkstationConfigIamMemberConditionPtrInput
	// The location where the workstation cluster config should reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
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
	// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringInput
	WorkstationClusterId pulumi.StringInput
	WorkstationConfigId  pulumi.StringInput
}

func (WorkstationConfigIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationConfigIamMemberArgs)(nil)).Elem()
}

type WorkstationConfigIamMemberInput interface {
	pulumi.Input

	ToWorkstationConfigIamMemberOutput() WorkstationConfigIamMemberOutput
	ToWorkstationConfigIamMemberOutputWithContext(ctx context.Context) WorkstationConfigIamMemberOutput
}

func (*WorkstationConfigIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationConfigIamMember)(nil)).Elem()
}

func (i *WorkstationConfigIamMember) ToWorkstationConfigIamMemberOutput() WorkstationConfigIamMemberOutput {
	return i.ToWorkstationConfigIamMemberOutputWithContext(context.Background())
}

func (i *WorkstationConfigIamMember) ToWorkstationConfigIamMemberOutputWithContext(ctx context.Context) WorkstationConfigIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationConfigIamMemberOutput)
}

// WorkstationConfigIamMemberArrayInput is an input type that accepts WorkstationConfigIamMemberArray and WorkstationConfigIamMemberArrayOutput values.
// You can construct a concrete instance of `WorkstationConfigIamMemberArrayInput` via:
//
//	WorkstationConfigIamMemberArray{ WorkstationConfigIamMemberArgs{...} }
type WorkstationConfigIamMemberArrayInput interface {
	pulumi.Input

	ToWorkstationConfigIamMemberArrayOutput() WorkstationConfigIamMemberArrayOutput
	ToWorkstationConfigIamMemberArrayOutputWithContext(context.Context) WorkstationConfigIamMemberArrayOutput
}

type WorkstationConfigIamMemberArray []WorkstationConfigIamMemberInput

func (WorkstationConfigIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationConfigIamMember)(nil)).Elem()
}

func (i WorkstationConfigIamMemberArray) ToWorkstationConfigIamMemberArrayOutput() WorkstationConfigIamMemberArrayOutput {
	return i.ToWorkstationConfigIamMemberArrayOutputWithContext(context.Background())
}

func (i WorkstationConfigIamMemberArray) ToWorkstationConfigIamMemberArrayOutputWithContext(ctx context.Context) WorkstationConfigIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationConfigIamMemberArrayOutput)
}

// WorkstationConfigIamMemberMapInput is an input type that accepts WorkstationConfigIamMemberMap and WorkstationConfigIamMemberMapOutput values.
// You can construct a concrete instance of `WorkstationConfigIamMemberMapInput` via:
//
//	WorkstationConfigIamMemberMap{ "key": WorkstationConfigIamMemberArgs{...} }
type WorkstationConfigIamMemberMapInput interface {
	pulumi.Input

	ToWorkstationConfigIamMemberMapOutput() WorkstationConfigIamMemberMapOutput
	ToWorkstationConfigIamMemberMapOutputWithContext(context.Context) WorkstationConfigIamMemberMapOutput
}

type WorkstationConfigIamMemberMap map[string]WorkstationConfigIamMemberInput

func (WorkstationConfigIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationConfigIamMember)(nil)).Elem()
}

func (i WorkstationConfigIamMemberMap) ToWorkstationConfigIamMemberMapOutput() WorkstationConfigIamMemberMapOutput {
	return i.ToWorkstationConfigIamMemberMapOutputWithContext(context.Background())
}

func (i WorkstationConfigIamMemberMap) ToWorkstationConfigIamMemberMapOutputWithContext(ctx context.Context) WorkstationConfigIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationConfigIamMemberMapOutput)
}

type WorkstationConfigIamMemberOutput struct{ *pulumi.OutputState }

func (WorkstationConfigIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationConfigIamMember)(nil)).Elem()
}

func (o WorkstationConfigIamMemberOutput) ToWorkstationConfigIamMemberOutput() WorkstationConfigIamMemberOutput {
	return o
}

func (o WorkstationConfigIamMemberOutput) ToWorkstationConfigIamMemberOutputWithContext(ctx context.Context) WorkstationConfigIamMemberOutput {
	return o
}

func (o WorkstationConfigIamMemberOutput) Condition() WorkstationConfigIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) WorkstationConfigIamMemberConditionPtrOutput { return v.Condition }).(WorkstationConfigIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o WorkstationConfigIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where the workstation cluster config should reside.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o WorkstationConfigIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o WorkstationConfigIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WorkstationConfigIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `workstations.WorkstationConfigIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o WorkstationConfigIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o WorkstationConfigIamMemberOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationConfigIamMemberOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationConfigIamMember) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

type WorkstationConfigIamMemberArrayOutput struct{ *pulumi.OutputState }

func (WorkstationConfigIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationConfigIamMember)(nil)).Elem()
}

func (o WorkstationConfigIamMemberArrayOutput) ToWorkstationConfigIamMemberArrayOutput() WorkstationConfigIamMemberArrayOutput {
	return o
}

func (o WorkstationConfigIamMemberArrayOutput) ToWorkstationConfigIamMemberArrayOutputWithContext(ctx context.Context) WorkstationConfigIamMemberArrayOutput {
	return o
}

func (o WorkstationConfigIamMemberArrayOutput) Index(i pulumi.IntInput) WorkstationConfigIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkstationConfigIamMember {
		return vs[0].([]*WorkstationConfigIamMember)[vs[1].(int)]
	}).(WorkstationConfigIamMemberOutput)
}

type WorkstationConfigIamMemberMapOutput struct{ *pulumi.OutputState }

func (WorkstationConfigIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationConfigIamMember)(nil)).Elem()
}

func (o WorkstationConfigIamMemberMapOutput) ToWorkstationConfigIamMemberMapOutput() WorkstationConfigIamMemberMapOutput {
	return o
}

func (o WorkstationConfigIamMemberMapOutput) ToWorkstationConfigIamMemberMapOutputWithContext(ctx context.Context) WorkstationConfigIamMemberMapOutput {
	return o
}

func (o WorkstationConfigIamMemberMapOutput) MapIndex(k pulumi.StringInput) WorkstationConfigIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkstationConfigIamMember {
		return vs[0].(map[string]*WorkstationConfigIamMember)[vs[1].(string)]
	}).(WorkstationConfigIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationConfigIamMemberInput)(nil)).Elem(), &WorkstationConfigIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationConfigIamMemberArrayInput)(nil)).Elem(), WorkstationConfigIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationConfigIamMemberMapInput)(nil)).Elem(), WorkstationConfigIamMemberMap{})
	pulumi.RegisterOutputType(WorkstationConfigIamMemberOutput{})
	pulumi.RegisterOutputType(WorkstationConfigIamMemberArrayOutput{})
	pulumi.RegisterOutputType(WorkstationConfigIamMemberMapOutput{})
}
