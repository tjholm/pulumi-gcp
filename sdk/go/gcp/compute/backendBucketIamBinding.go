// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/global/backendBuckets/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine backendbucket IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamBinding:BackendBucketIamBinding editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamBinding:BackendBucketIamBinding editor "projects/{{project}}/global/backendBuckets/{{backend_bucket}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/backendBucketIamBinding:BackendBucketIamBinding editor projects/{{project}}/global/backendBuckets/{{backend_bucket}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BackendBucketIamBinding struct {
	pulumi.CustomResourceState

	Condition BackendBucketIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
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

// NewBackendBucketIamBinding registers a new resource with the given unique name, arguments, and options.
func NewBackendBucketIamBinding(ctx *pulumi.Context,
	name string, args *BackendBucketIamBindingArgs, opts ...pulumi.ResourceOption) (*BackendBucketIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource BackendBucketIamBinding
	err := ctx.RegisterResource("gcp:compute/backendBucketIamBinding:BackendBucketIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendBucketIamBinding gets an existing BackendBucketIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendBucketIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendBucketIamBindingState, opts ...pulumi.ResourceOption) (*BackendBucketIamBinding, error) {
	var resource BackendBucketIamBinding
	err := ctx.ReadResource("gcp:compute/backendBucketIamBinding:BackendBucketIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendBucketIamBinding resources.
type backendBucketIamBindingState struct {
	Condition *BackendBucketIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
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

type BackendBucketIamBindingState struct {
	Condition BackendBucketIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
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

func (BackendBucketIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamBindingState)(nil)).Elem()
}

type backendBucketIamBindingArgs struct {
	Condition *BackendBucketIamBindingCondition `pulumi:"condition"`
	Members   []string                          `pulumi:"members"`
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

// The set of arguments for constructing a BackendBucketIamBinding resource.
type BackendBucketIamBindingArgs struct {
	Condition BackendBucketIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
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

func (BackendBucketIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendBucketIamBindingArgs)(nil)).Elem()
}

type BackendBucketIamBindingInput interface {
	pulumi.Input

	ToBackendBucketIamBindingOutput() BackendBucketIamBindingOutput
	ToBackendBucketIamBindingOutputWithContext(ctx context.Context) BackendBucketIamBindingOutput
}

func (*BackendBucketIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamBinding)(nil)).Elem()
}

func (i *BackendBucketIamBinding) ToBackendBucketIamBindingOutput() BackendBucketIamBindingOutput {
	return i.ToBackendBucketIamBindingOutputWithContext(context.Background())
}

func (i *BackendBucketIamBinding) ToBackendBucketIamBindingOutputWithContext(ctx context.Context) BackendBucketIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamBindingOutput)
}

// BackendBucketIamBindingArrayInput is an input type that accepts BackendBucketIamBindingArray and BackendBucketIamBindingArrayOutput values.
// You can construct a concrete instance of `BackendBucketIamBindingArrayInput` via:
//
//	BackendBucketIamBindingArray{ BackendBucketIamBindingArgs{...} }
type BackendBucketIamBindingArrayInput interface {
	pulumi.Input

	ToBackendBucketIamBindingArrayOutput() BackendBucketIamBindingArrayOutput
	ToBackendBucketIamBindingArrayOutputWithContext(context.Context) BackendBucketIamBindingArrayOutput
}

type BackendBucketIamBindingArray []BackendBucketIamBindingInput

func (BackendBucketIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamBinding)(nil)).Elem()
}

func (i BackendBucketIamBindingArray) ToBackendBucketIamBindingArrayOutput() BackendBucketIamBindingArrayOutput {
	return i.ToBackendBucketIamBindingArrayOutputWithContext(context.Background())
}

func (i BackendBucketIamBindingArray) ToBackendBucketIamBindingArrayOutputWithContext(ctx context.Context) BackendBucketIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamBindingArrayOutput)
}

// BackendBucketIamBindingMapInput is an input type that accepts BackendBucketIamBindingMap and BackendBucketIamBindingMapOutput values.
// You can construct a concrete instance of `BackendBucketIamBindingMapInput` via:
//
//	BackendBucketIamBindingMap{ "key": BackendBucketIamBindingArgs{...} }
type BackendBucketIamBindingMapInput interface {
	pulumi.Input

	ToBackendBucketIamBindingMapOutput() BackendBucketIamBindingMapOutput
	ToBackendBucketIamBindingMapOutputWithContext(context.Context) BackendBucketIamBindingMapOutput
}

type BackendBucketIamBindingMap map[string]BackendBucketIamBindingInput

func (BackendBucketIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamBinding)(nil)).Elem()
}

func (i BackendBucketIamBindingMap) ToBackendBucketIamBindingMapOutput() BackendBucketIamBindingMapOutput {
	return i.ToBackendBucketIamBindingMapOutputWithContext(context.Background())
}

func (i BackendBucketIamBindingMap) ToBackendBucketIamBindingMapOutputWithContext(ctx context.Context) BackendBucketIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendBucketIamBindingMapOutput)
}

type BackendBucketIamBindingOutput struct{ *pulumi.OutputState }

func (BackendBucketIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendBucketIamBinding)(nil)).Elem()
}

func (o BackendBucketIamBindingOutput) ToBackendBucketIamBindingOutput() BackendBucketIamBindingOutput {
	return o
}

func (o BackendBucketIamBindingOutput) ToBackendBucketIamBindingOutputWithContext(ctx context.Context) BackendBucketIamBindingOutput {
	return o
}

func (o BackendBucketIamBindingOutput) Condition() BackendBucketIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) BackendBucketIamBindingConditionPtrOutput { return v.Condition }).(BackendBucketIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o BackendBucketIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BackendBucketIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o BackendBucketIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o BackendBucketIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.BackendBucketIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o BackendBucketIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendBucketIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type BackendBucketIamBindingArrayOutput struct{ *pulumi.OutputState }

func (BackendBucketIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendBucketIamBinding)(nil)).Elem()
}

func (o BackendBucketIamBindingArrayOutput) ToBackendBucketIamBindingArrayOutput() BackendBucketIamBindingArrayOutput {
	return o
}

func (o BackendBucketIamBindingArrayOutput) ToBackendBucketIamBindingArrayOutputWithContext(ctx context.Context) BackendBucketIamBindingArrayOutput {
	return o
}

func (o BackendBucketIamBindingArrayOutput) Index(i pulumi.IntInput) BackendBucketIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendBucketIamBinding {
		return vs[0].([]*BackendBucketIamBinding)[vs[1].(int)]
	}).(BackendBucketIamBindingOutput)
}

type BackendBucketIamBindingMapOutput struct{ *pulumi.OutputState }

func (BackendBucketIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendBucketIamBinding)(nil)).Elem()
}

func (o BackendBucketIamBindingMapOutput) ToBackendBucketIamBindingMapOutput() BackendBucketIamBindingMapOutput {
	return o
}

func (o BackendBucketIamBindingMapOutput) ToBackendBucketIamBindingMapOutputWithContext(ctx context.Context) BackendBucketIamBindingMapOutput {
	return o
}

func (o BackendBucketIamBindingMapOutput) MapIndex(k pulumi.StringInput) BackendBucketIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendBucketIamBinding {
		return vs[0].(map[string]*BackendBucketIamBinding)[vs[1].(string)]
	}).(BackendBucketIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamBindingInput)(nil)).Elem(), &BackendBucketIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamBindingArrayInput)(nil)).Elem(), BackendBucketIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendBucketIamBindingMapInput)(nil)).Elem(), BackendBucketIamBindingMap{})
	pulumi.RegisterOutputType(BackendBucketIamBindingOutput{})
	pulumi.RegisterOutputType(BackendBucketIamBindingArrayOutput{})
	pulumi.RegisterOutputType(BackendBucketIamBindingMapOutput{})
}
