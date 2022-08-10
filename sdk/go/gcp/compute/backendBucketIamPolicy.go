// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/backendBuckets/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine backendbucket IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamPolicy:BackendBucketIamPolicy editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamPolicy:BackendBucketIamPolicy editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamPolicy:BackendBucketIamPolicy editor projects/{{project}}/global/backendBuckets/{{backend_bucket}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackendBucketIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackendBucketIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackendBucketIamPolicy(ctx *pulumi.Context,
	name string, args *BackendBucketIamPolicyArgs, opts ...pulumi.ResourceOption) (*BackendBucketIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource BackendBucketIamPolicy
	err := ctx.RegisterResource("gcp:compute/backendBucketIamPolicy:BackendBucketIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucketIamPolicy gets an existing BackendBucketIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucketIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketIamPolicyState, opts ...pulumi.ResourceOption) (*BackendBucketIamPolicy, error) {
	var resource BackendBucketIamPolicy
	err := ctx.ReadResource("gcp:compute/backendBucketIamPolicy:BackendBucketIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucketIamPolicy resources.
type backendBucketIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackendBucketIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamPolicyState)(nil)).Elem()
}

type backendBucketIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendBucketIamPolicy resource.
type BackendBucketIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendBucketIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamPolicyArgs)(nil)).Elem()
}

type BackendBucketIamPolicyInput interface {
	pulumi.Input

	ToBackendBucketIamPolicyOutput() BackendBucketIamPolicyOutput
	ToBackendBucketIamPolicyOutputWithContext(ctx context.Context) BackendBucketIamPolicyOutput
}

func (*BackendBucketIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamPolicy)(nil)).Elem()
}

func (i *BackendBucketIamPolicy) ToBackendBucketIamPolicyOutput() BackendBucketIamPolicyOutput {
	return i.ToBackendBucketIamPolicyOutputWithContext(context.Background())
}

func (i *BackendBucketIamPolicy) ToBackendBucketIamPolicyOutputWithContext(ctx context.Context) BackendBucketIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamPolicyOutput)
}

// BackendBucketIamPolicyArrayInput is an input type that accepts BackendBucketIamPolicyArray and BackendBucketIamPolicyArrayOutput values.
// You can construct a concrete instance of `BackendBucketIamPolicyArrayInput` via:
//
//	BackendBucketIamPolicyArray{ BackendBucketIamPolicyArgs{...} }
type BackendBucketIamPolicyArrayInput interface {
	pulumi.Input

	ToBackendBucketIamPolicyArrayOutput() BackendBucketIamPolicyArrayOutput
	ToBackendBucketIamPolicyArrayOutputWithContext(context.Context) BackendBucketIamPolicyArrayOutput
}

type BackendBucketIamPolicyArray []BackendBucketIamPolicyInput

func (BackendBucketIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamPolicy)(nil)).Elem()
}

func (i BackendBucketIamPolicyArray) ToBackendBucketIamPolicyArrayOutput() BackendBucketIamPolicyArrayOutput {
	return i.ToBackendBucketIamPolicyArrayOutputWithContext(context.Background())
}

func (i BackendBucketIamPolicyArray) ToBackendBucketIamPolicyArrayOutputWithContext(ctx context.Context) BackendBucketIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamPolicyArrayOutput)
}

// BackendBucketIamPolicyMapInput is an input type that accepts BackendBucketIamPolicyMap and BackendBucketIamPolicyMapOutput values.
// You can construct a concrete instance of `BackendBucketIamPolicyMapInput` via:
//
//	BackendBucketIamPolicyMap{ "key": BackendBucketIamPolicyArgs{...} }
type BackendBucketIamPolicyMapInput interface {
	pulumi.Input

	ToBackendBucketIamPolicyMapOutput() BackendBucketIamPolicyMapOutput
	ToBackendBucketIamPolicyMapOutputWithContext(context.Context) BackendBucketIamPolicyMapOutput
}

type BackendBucketIamPolicyMap map[string]BackendBucketIamPolicyInput

func (BackendBucketIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamPolicy)(nil)).Elem()
}

func (i BackendBucketIamPolicyMap) ToBackendBucketIamPolicyMapOutput() BackendBucketIamPolicyMapOutput {
	return i.ToBackendBucketIamPolicyMapOutputWithContext(context.Background())
}

func (i BackendBucketIamPolicyMap) ToBackendBucketIamPolicyMapOutputWithContext(ctx context.Context) BackendBucketIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamPolicyMapOutput)
}

type BackendBucketIamPolicyOutput struct{ *pulumi.OutputState }

func (BackendBucketIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamPolicy)(nil)).Elem()
}

func (o BackendBucketIamPolicyOutput) ToBackendBucketIamPolicyOutput() BackendBucketIamPolicyOutput {
	return o
}

func (o BackendBucketIamPolicyOutput) ToBackendBucketIamPolicyOutputWithContext(ctx context.Context) BackendBucketIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o BackendBucketIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackendBucketIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o BackendBucketIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackendBucketIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type BackendBucketIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackendBucketIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamPolicy)(nil)).Elem()
}

func (o BackendBucketIamPolicyArrayOutput) ToBackendBucketIamPolicyArrayOutput() BackendBucketIamPolicyArrayOutput {
	return o
}

func (o BackendBucketIamPolicyArrayOutput) ToBackendBucketIamPolicyArrayOutputWithContext(ctx context.Context) BackendBucketIamPolicyArrayOutput {
	return o
}

func (o BackendBucketIamPolicyArrayOutput) Index(i pulumi.IntInput) BackendBucketIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendBucketIamPolicy {
		return vs[0].([]*BackendBucketIamPolicy)[vs[1].(int)]
	}).(BackendBucketIamPolicyOutput)
}

type BackendBucketIamPolicyMapOutput struct{ *pulumi.OutputState }

func (BackendBucketIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamPolicy)(nil)).Elem()
}

func (o BackendBucketIamPolicyMapOutput) ToBackendBucketIamPolicyMapOutput() BackendBucketIamPolicyMapOutput {
	return o
}

func (o BackendBucketIamPolicyMapOutput) ToBackendBucketIamPolicyMapOutputWithContext(ctx context.Context) BackendBucketIamPolicyMapOutput {
	return o
}

func (o BackendBucketIamPolicyMapOutput) MapIndex(k pulumi.StringInput) BackendBucketIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendBucketIamPolicy {
		return vs[0].(map[string]*BackendBucketIamPolicy)[vs[1].(string)]
	}).(BackendBucketIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamPolicyInput)(nil)).Elem(), &BackendBucketIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamPolicyArrayInput)(nil)).Elem(), BackendBucketIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamPolicyMapInput)(nil)).Elem(), BackendBucketIamPolicyMap{})
	pulumi.RegisterOutputType(BackendBucketIamPolicyOutput{})
	pulumi.RegisterOutputType(BackendBucketIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackendBucketIamPolicyMapOutput{})
}
