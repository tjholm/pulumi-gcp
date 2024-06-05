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
// $ pulumi import gcp:workstations/workstationIamPolicy:WorkstationIamPolicy editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationIamPolicy:WorkstationIamPolicy editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:workstations/workstationIamPolicy:WorkstationIamPolicy editor projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WorkstationIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              pulumi.StringOutput `pulumi:"project"`
	WorkstationClusterId pulumi.StringOutput `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringOutput `pulumi:"workstationConfigId"`
	WorkstationId        pulumi.StringOutput `pulumi:"workstationId"`
}

// NewWorkstationIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkstationIamPolicy(ctx *pulumi.Context,
	name string, args *WorkstationIamPolicyArgs, opts ...pulumi.ResourceOption) (*WorkstationIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
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
	var resource WorkstationIamPolicy
	err := ctx.RegisterResource("gcp:workstations/workstationIamPolicy:WorkstationIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationIamPolicy gets an existing WorkstationIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationIamPolicyState, opts ...pulumi.ResourceOption) (*WorkstationIamPolicy, error) {
	var resource WorkstationIamPolicy
	err := ctx.ReadResource("gcp:workstations/workstationIamPolicy:WorkstationIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationIamPolicy resources.
type workstationIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              *string `pulumi:"project"`
	WorkstationClusterId *string `pulumi:"workstationClusterId"`
	WorkstationConfigId  *string `pulumi:"workstationConfigId"`
	WorkstationId        *string `pulumi:"workstationId"`
}

type WorkstationIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              pulumi.StringPtrInput
	WorkstationClusterId pulumi.StringPtrInput
	WorkstationConfigId  pulumi.StringPtrInput
	WorkstationId        pulumi.StringPtrInput
}

func (WorkstationIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationIamPolicyState)(nil)).Elem()
}

type workstationIamPolicyArgs struct {
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              *string `pulumi:"project"`
	WorkstationClusterId string  `pulumi:"workstationClusterId"`
	WorkstationConfigId  string  `pulumi:"workstationConfigId"`
	WorkstationId        string  `pulumi:"workstationId"`
}

// The set of arguments for constructing a WorkstationIamPolicy resource.
type WorkstationIamPolicyArgs struct {
	// The location where the workstation parent resources reside.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project              pulumi.StringPtrInput
	WorkstationClusterId pulumi.StringInput
	WorkstationConfigId  pulumi.StringInput
	WorkstationId        pulumi.StringInput
}

func (WorkstationIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationIamPolicyArgs)(nil)).Elem()
}

type WorkstationIamPolicyInput interface {
	pulumi.Input

	ToWorkstationIamPolicyOutput() WorkstationIamPolicyOutput
	ToWorkstationIamPolicyOutputWithContext(ctx context.Context) WorkstationIamPolicyOutput
}

func (*WorkstationIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationIamPolicy)(nil)).Elem()
}

func (i *WorkstationIamPolicy) ToWorkstationIamPolicyOutput() WorkstationIamPolicyOutput {
	return i.ToWorkstationIamPolicyOutputWithContext(context.Background())
}

func (i *WorkstationIamPolicy) ToWorkstationIamPolicyOutputWithContext(ctx context.Context) WorkstationIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamPolicyOutput)
}

// WorkstationIamPolicyArrayInput is an input type that accepts WorkstationIamPolicyArray and WorkstationIamPolicyArrayOutput values.
// You can construct a concrete instance of `WorkstationIamPolicyArrayInput` via:
//
//	WorkstationIamPolicyArray{ WorkstationIamPolicyArgs{...} }
type WorkstationIamPolicyArrayInput interface {
	pulumi.Input

	ToWorkstationIamPolicyArrayOutput() WorkstationIamPolicyArrayOutput
	ToWorkstationIamPolicyArrayOutputWithContext(context.Context) WorkstationIamPolicyArrayOutput
}

type WorkstationIamPolicyArray []WorkstationIamPolicyInput

func (WorkstationIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationIamPolicy)(nil)).Elem()
}

func (i WorkstationIamPolicyArray) ToWorkstationIamPolicyArrayOutput() WorkstationIamPolicyArrayOutput {
	return i.ToWorkstationIamPolicyArrayOutputWithContext(context.Background())
}

func (i WorkstationIamPolicyArray) ToWorkstationIamPolicyArrayOutputWithContext(ctx context.Context) WorkstationIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamPolicyArrayOutput)
}

// WorkstationIamPolicyMapInput is an input type that accepts WorkstationIamPolicyMap and WorkstationIamPolicyMapOutput values.
// You can construct a concrete instance of `WorkstationIamPolicyMapInput` via:
//
//	WorkstationIamPolicyMap{ "key": WorkstationIamPolicyArgs{...} }
type WorkstationIamPolicyMapInput interface {
	pulumi.Input

	ToWorkstationIamPolicyMapOutput() WorkstationIamPolicyMapOutput
	ToWorkstationIamPolicyMapOutputWithContext(context.Context) WorkstationIamPolicyMapOutput
}

type WorkstationIamPolicyMap map[string]WorkstationIamPolicyInput

func (WorkstationIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationIamPolicy)(nil)).Elem()
}

func (i WorkstationIamPolicyMap) ToWorkstationIamPolicyMapOutput() WorkstationIamPolicyMapOutput {
	return i.ToWorkstationIamPolicyMapOutputWithContext(context.Background())
}

func (i WorkstationIamPolicyMap) ToWorkstationIamPolicyMapOutputWithContext(ctx context.Context) WorkstationIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationIamPolicyMapOutput)
}

type WorkstationIamPolicyOutput struct{ *pulumi.OutputState }

func (WorkstationIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationIamPolicy)(nil)).Elem()
}

func (o WorkstationIamPolicyOutput) ToWorkstationIamPolicyOutput() WorkstationIamPolicyOutput {
	return o
}

func (o WorkstationIamPolicyOutput) ToWorkstationIamPolicyOutputWithContext(ctx context.Context) WorkstationIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o WorkstationIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location where the workstation parent resources reside.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
// location is specified, it is taken from the provider configuration.
func (o WorkstationIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o WorkstationIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o WorkstationIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o WorkstationIamPolicyOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationIamPolicyOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func (o WorkstationIamPolicyOutput) WorkstationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationIamPolicy) pulumi.StringOutput { return v.WorkstationId }).(pulumi.StringOutput)
}

type WorkstationIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (WorkstationIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationIamPolicy)(nil)).Elem()
}

func (o WorkstationIamPolicyArrayOutput) ToWorkstationIamPolicyArrayOutput() WorkstationIamPolicyArrayOutput {
	return o
}

func (o WorkstationIamPolicyArrayOutput) ToWorkstationIamPolicyArrayOutputWithContext(ctx context.Context) WorkstationIamPolicyArrayOutput {
	return o
}

func (o WorkstationIamPolicyArrayOutput) Index(i pulumi.IntInput) WorkstationIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkstationIamPolicy {
		return vs[0].([]*WorkstationIamPolicy)[vs[1].(int)]
	}).(WorkstationIamPolicyOutput)
}

type WorkstationIamPolicyMapOutput struct{ *pulumi.OutputState }

func (WorkstationIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationIamPolicy)(nil)).Elem()
}

func (o WorkstationIamPolicyMapOutput) ToWorkstationIamPolicyMapOutput() WorkstationIamPolicyMapOutput {
	return o
}

func (o WorkstationIamPolicyMapOutput) ToWorkstationIamPolicyMapOutputWithContext(ctx context.Context) WorkstationIamPolicyMapOutput {
	return o
}

func (o WorkstationIamPolicyMapOutput) MapIndex(k pulumi.StringInput) WorkstationIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkstationIamPolicy {
		return vs[0].(map[string]*WorkstationIamPolicy)[vs[1].(string)]
	}).(WorkstationIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamPolicyInput)(nil)).Elem(), &WorkstationIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamPolicyArrayInput)(nil)).Elem(), WorkstationIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationIamPolicyMapInput)(nil)).Elem(), WorkstationIamPolicyMap{})
	pulumi.RegisterOutputType(WorkstationIamPolicyOutput{})
	pulumi.RegisterOutputType(WorkstationIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(WorkstationIamPolicyMapOutput{})
}
