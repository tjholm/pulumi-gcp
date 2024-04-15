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
// * projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
//
// * {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}}
//
// * {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}}
//
// * {{workstation_id}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Workstations workstation IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationIamBinding:WorkstationIamBinding editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationIamBinding:WorkstationIamBinding editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationIamBinding:WorkstationIamBinding editor projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WorkstationIamBinding struct {
	pulumi.CustomResourceState

	Condition WorkstationIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to
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
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringOutput `pulumi:"role"`
	WorkstationClusterId pulumi.StringOutput `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringOutput `pulumi:"workstationConfigId"`
	WorkstationId        pulumi.StringOutput `pulumi:"workstationId"`
}

// NewWorkstationIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWorkstationIamBinding(ctx *pulumi.Context,
	name string, args *WorkstationIamBindingArgs, opts ...pulumi.ResourceOption) (*WorkstationIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
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
	if args.WorkstationId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkstationIamBinding
	err := ctx.RegisterResource("gcp:workstations/workstationIamBinding:WorkstationIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationIamBinding gets an existing WorkstationIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationIamBindingState, opts ...pulumi.ResourceOption) (*WorkstationIamBinding, error) {
	var resource WorkstationIamBinding
	err := ctx.ReadResource("gcp:workstations/workstationIamBinding:WorkstationIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationIamBinding resources.
type workstationIamBindingState struct {
	Condition *WorkstationIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 *string `pulumi:"role"`
	WorkstationClusterId *string `pulumi:"workstationClusterId"`
	WorkstationConfigId  *string `pulumi:"workstationConfigId"`
	WorkstationId        *string `pulumi:"workstationId"`
}

type WorkstationIamBindingState struct {
	Condition WorkstationIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringPtrInput
	WorkstationClusterId pulumi.StringPtrInput
	WorkstationConfigId  pulumi.StringPtrInput
	WorkstationId        pulumi.StringPtrInput
}

func (WorkstationIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationIamBindingState)(nil)).Elem()
}

type workstationIamBindingArgs struct {
	Condition *WorkstationIamBindingCondition `pulumi:"condition"`
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to
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
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 string `pulumi:"role"`
	WorkstationClusterId string `pulumi:"workstationClusterId"`
	WorkstationConfigId  string `pulumi:"workstationConfigId"`
	WorkstationId        string `pulumi:"workstationId"`
}

// The set of arguments for constructing a WorkstationIamBinding resource.
type WorkstationIamBindingArgs struct {
	Condition WorkstationIamBindingConditionPtrInput
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to
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
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role                 pulumi.StringInput
	WorkstationClusterId pulumi.StringInput
	WorkstationConfigId  pulumi.StringInput
	WorkstationId        pulumi.StringInput
}

func (WorkstationIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationIamBindingArgs)(nil)).Elem()
}

type WorkstationIamBindingInput interface {
	pulumi.Input

	ToWorkstationIamBindingOutput() WorkstationIamBindingOutput
	ToWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationIamBindingOutput
}

func (*WorkstationIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationIamBinding)(nil)).Elem()
}

func (i *WorkstationIamBinding) ToWorkstationIamBindingOutput() WorkstationIamBindingOutput {
	return i.ToWorkstationIamBindingOutputWithContext(context.Background())
}

func (i *WorkstationIamBinding) ToWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamBindingOutput)
}

// WorkstationIamBindingArrayInput is an input type that accepts WorkstationIamBindingArray and WorkstationIamBindingArrayOutput values.
// You can construct a concrete instance of `WorkstationIamBindingArrayInput` via:
//
//	WorkstationIamBindingArray{ WorkstationIamBindingArgs{...} }
type WorkstationIamBindingArrayInput interface {
	pulumi.Input

	ToWorkstationIamBindingArrayOutput() WorkstationIamBindingArrayOutput
	ToWorkstationIamBindingArrayOutputWithContext(context.Context) WorkstationIamBindingArrayOutput
}

type WorkstationIamBindingArray []WorkstationIamBindingInput

func (WorkstationIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationIamBinding)(nil)).Elem()
}

func (i WorkstationIamBindingArray) ToWorkstationIamBindingArrayOutput() WorkstationIamBindingArrayOutput {
	return i.ToWorkstationIamBindingArrayOutputWithContext(context.Background())
}

func (i WorkstationIamBindingArray) ToWorkstationIamBindingArrayOutputWithContext(ctx context.Context) WorkstationIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamBindingArrayOutput)
}

// WorkstationIamBindingMapInput is an input type that accepts WorkstationIamBindingMap and WorkstationIamBindingMapOutput values.
// You can construct a concrete instance of `WorkstationIamBindingMapInput` via:
//
//	WorkstationIamBindingMap{ "key": WorkstationIamBindingArgs{...} }
type WorkstationIamBindingMapInput interface {
	pulumi.Input

	ToWorkstationIamBindingMapOutput() WorkstationIamBindingMapOutput
	ToWorkstationIamBindingMapOutputWithContext(context.Context) WorkstationIamBindingMapOutput
}

type WorkstationIamBindingMap map[string]WorkstationIamBindingInput

func (WorkstationIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationIamBinding)(nil)).Elem()
}

func (i WorkstationIamBindingMap) ToWorkstationIamBindingMapOutput() WorkstationIamBindingMapOutput {
	return i.ToWorkstationIamBindingMapOutputWithContext(context.Background())
}

func (i WorkstationIamBindingMap) ToWorkstationIamBindingMapOutputWithContext(ctx context.Context) WorkstationIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamBindingMapOutput)
}

type WorkstationIamBindingOutput struct{ *pulumi.OutputState }

func (WorkstationIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationIamBinding)(nil)).Elem()
}

func (o WorkstationIamBindingOutput) ToWorkstationIamBindingOutput() WorkstationIamBindingOutput {
	return o
}

func (o WorkstationIamBindingOutput) ToWorkstationIamBindingOutputWithContext(ctx context.Context) WorkstationIamBindingOutput {
	return o
}

func (o WorkstationIamBindingOutput) Condition() WorkstationIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) WorkstationIamBindingConditionPtrOutput { return v.Condition }).(WorkstationIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o WorkstationIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where the workstation parent resources reside.
// Used to find the parent resource to bind the IAM policy to
func (o WorkstationIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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
func (o WorkstationIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WorkstationIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o WorkstationIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o WorkstationIamBindingOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationIamBindingOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func (o WorkstationIamBindingOutput) WorkstationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamBinding) pulumi.StringOutput { return v.WorkstationId }).(pulumi.StringOutput)
}

type WorkstationIamBindingArrayOutput struct{ *pulumi.OutputState }

func (WorkstationIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationIamBinding)(nil)).Elem()
}

func (o WorkstationIamBindingArrayOutput) ToWorkstationIamBindingArrayOutput() WorkstationIamBindingArrayOutput {
	return o
}

func (o WorkstationIamBindingArrayOutput) ToWorkstationIamBindingArrayOutputWithContext(ctx context.Context) WorkstationIamBindingArrayOutput {
	return o
}

func (o WorkstationIamBindingArrayOutput) Index(i pulumi.IntInput) WorkstationIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkstationIamBinding {
		return vs[0].([]*WorkstationIamBinding)[vs[1].(int)]
	}).(WorkstationIamBindingOutput)
}

type WorkstationIamBindingMapOutput struct{ *pulumi.OutputState }

func (WorkstationIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationIamBinding)(nil)).Elem()
}

func (o WorkstationIamBindingMapOutput) ToWorkstationIamBindingMapOutput() WorkstationIamBindingMapOutput {
	return o
}

func (o WorkstationIamBindingMapOutput) ToWorkstationIamBindingMapOutputWithContext(ctx context.Context) WorkstationIamBindingMapOutput {
	return o
}

func (o WorkstationIamBindingMapOutput) MapIndex(k pulumi.StringInput) WorkstationIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkstationIamBinding {
		return vs[0].(map[string]*WorkstationIamBinding)[vs[1].(string)]
	}).(WorkstationIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamBindingInput)(nil)).Elem(), &WorkstationIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamBindingArrayInput)(nil)).Elem(), WorkstationIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamBindingMapInput)(nil)).Elem(), WorkstationIamBindingMap{})
	pulumi.RegisterOutputType(WorkstationIamBindingOutput{})
	pulumi.RegisterOutputType(WorkstationIamBindingArrayOutput{})
	pulumi.RegisterOutputType(WorkstationIamBindingMapOutput{})
}
