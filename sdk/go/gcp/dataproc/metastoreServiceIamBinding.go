// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/services/{{service_id}} * {{project}}/{{location}}/{{service_id}} * {{location}}/{{service_id}} * {{service_id}} Any variables not passed in the import command will be taken from the provider configuration. Dataproc metastore service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding editor projects/{{project}}/locations/{{location}}/services/{{service_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MetastoreServiceIamBinding struct {
	pulumi.CustomResourceState

	Condition MetastoreServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringOutput `pulumi:"role"`
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewMetastoreServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewMetastoreServiceIamBinding(ctx *pulumi.Context,
	name string, args *MetastoreServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*MetastoreServiceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource MetastoreServiceIamBinding
	err := ctx.RegisterResource("gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreServiceIamBinding gets an existing MetastoreServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreServiceIamBindingState, opts ...pulumi.ResourceOption) (*MetastoreServiceIamBinding, error) {
	var resource MetastoreServiceIamBinding
	err := ctx.ReadResource("gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreServiceIamBinding resources.
type metastoreServiceIamBindingState struct {
	Condition *MetastoreServiceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      *string `pulumi:"role"`
	ServiceId *string `pulumi:"serviceId"`
}

type MetastoreServiceIamBindingState struct {
	Condition MetastoreServiceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringPtrInput
	ServiceId pulumi.StringPtrInput
}

func (MetastoreServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreServiceIamBindingState)(nil)).Elem()
}

type metastoreServiceIamBindingArgs struct {
	Condition *MetastoreServiceIamBindingCondition `pulumi:"condition"`
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      string `pulumi:"role"`
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a MetastoreServiceIamBinding resource.
type MetastoreServiceIamBindingArgs struct {
	Condition MetastoreServiceIamBindingConditionPtrInput
	// The  location where the autoscaling policy should reside.
	// The default value is `global`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role      pulumi.StringInput
	ServiceId pulumi.StringInput
}

func (MetastoreServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreServiceIamBindingArgs)(nil)).Elem()
}

type MetastoreServiceIamBindingInput interface {
	pulumi.Input

	ToMetastoreServiceIamBindingOutput() MetastoreServiceIamBindingOutput
	ToMetastoreServiceIamBindingOutputWithContext(ctx context.Context) MetastoreServiceIamBindingOutput
}

func (*MetastoreServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreServiceIamBinding)(nil)).Elem()
}

func (i *MetastoreServiceIamBinding) ToMetastoreServiceIamBindingOutput() MetastoreServiceIamBindingOutput {
	return i.ToMetastoreServiceIamBindingOutputWithContext(context.Background())
}

func (i *MetastoreServiceIamBinding) ToMetastoreServiceIamBindingOutputWithContext(ctx context.Context) MetastoreServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamBindingOutput)
}

// MetastoreServiceIamBindingArrayInput is an input type that accepts MetastoreServiceIamBindingArray and MetastoreServiceIamBindingArrayOutput values.
// You can construct a concrete instance of `MetastoreServiceIamBindingArrayInput` via:
//
//          MetastoreServiceIamBindingArray{ MetastoreServiceIamBindingArgs{...} }
type MetastoreServiceIamBindingArrayInput interface {
	pulumi.Input

	ToMetastoreServiceIamBindingArrayOutput() MetastoreServiceIamBindingArrayOutput
	ToMetastoreServiceIamBindingArrayOutputWithContext(context.Context) MetastoreServiceIamBindingArrayOutput
}

type MetastoreServiceIamBindingArray []MetastoreServiceIamBindingInput

func (MetastoreServiceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreServiceIamBinding)(nil)).Elem()
}

func (i MetastoreServiceIamBindingArray) ToMetastoreServiceIamBindingArrayOutput() MetastoreServiceIamBindingArrayOutput {
	return i.ToMetastoreServiceIamBindingArrayOutputWithContext(context.Background())
}

func (i MetastoreServiceIamBindingArray) ToMetastoreServiceIamBindingArrayOutputWithContext(ctx context.Context) MetastoreServiceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamBindingArrayOutput)
}

// MetastoreServiceIamBindingMapInput is an input type that accepts MetastoreServiceIamBindingMap and MetastoreServiceIamBindingMapOutput values.
// You can construct a concrete instance of `MetastoreServiceIamBindingMapInput` via:
//
//          MetastoreServiceIamBindingMap{ "key": MetastoreServiceIamBindingArgs{...} }
type MetastoreServiceIamBindingMapInput interface {
	pulumi.Input

	ToMetastoreServiceIamBindingMapOutput() MetastoreServiceIamBindingMapOutput
	ToMetastoreServiceIamBindingMapOutputWithContext(context.Context) MetastoreServiceIamBindingMapOutput
}

type MetastoreServiceIamBindingMap map[string]MetastoreServiceIamBindingInput

func (MetastoreServiceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreServiceIamBinding)(nil)).Elem()
}

func (i MetastoreServiceIamBindingMap) ToMetastoreServiceIamBindingMapOutput() MetastoreServiceIamBindingMapOutput {
	return i.ToMetastoreServiceIamBindingMapOutputWithContext(context.Background())
}

func (i MetastoreServiceIamBindingMap) ToMetastoreServiceIamBindingMapOutputWithContext(ctx context.Context) MetastoreServiceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreServiceIamBindingMapOutput)
}

type MetastoreServiceIamBindingOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreServiceIamBinding)(nil)).Elem()
}

func (o MetastoreServiceIamBindingOutput) ToMetastoreServiceIamBindingOutput() MetastoreServiceIamBindingOutput {
	return o
}

func (o MetastoreServiceIamBindingOutput) ToMetastoreServiceIamBindingOutputWithContext(ctx context.Context) MetastoreServiceIamBindingOutput {
	return o
}

func (o MetastoreServiceIamBindingOutput) Condition() MetastoreServiceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) MetastoreServiceIamBindingConditionPtrOutput { return v.Condition }).(MetastoreServiceIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o MetastoreServiceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The  location where the autoscaling policy should reside.
// The default value is `global`.
// Used to find the parent resource to bind the IAM policy to
func (o MetastoreServiceIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MetastoreServiceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MetastoreServiceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.MetastoreServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o MetastoreServiceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o MetastoreServiceIamBindingOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreServiceIamBinding) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

type MetastoreServiceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreServiceIamBinding)(nil)).Elem()
}

func (o MetastoreServiceIamBindingArrayOutput) ToMetastoreServiceIamBindingArrayOutput() MetastoreServiceIamBindingArrayOutput {
	return o
}

func (o MetastoreServiceIamBindingArrayOutput) ToMetastoreServiceIamBindingArrayOutputWithContext(ctx context.Context) MetastoreServiceIamBindingArrayOutput {
	return o
}

func (o MetastoreServiceIamBindingArrayOutput) Index(i pulumi.IntInput) MetastoreServiceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreServiceIamBinding {
		return vs[0].([]*MetastoreServiceIamBinding)[vs[1].(int)]
	}).(MetastoreServiceIamBindingOutput)
}

type MetastoreServiceIamBindingMapOutput struct{ *pulumi.OutputState }

func (MetastoreServiceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreServiceIamBinding)(nil)).Elem()
}

func (o MetastoreServiceIamBindingMapOutput) ToMetastoreServiceIamBindingMapOutput() MetastoreServiceIamBindingMapOutput {
	return o
}

func (o MetastoreServiceIamBindingMapOutput) ToMetastoreServiceIamBindingMapOutputWithContext(ctx context.Context) MetastoreServiceIamBindingMapOutput {
	return o
}

func (o MetastoreServiceIamBindingMapOutput) MapIndex(k pulumi.StringInput) MetastoreServiceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreServiceIamBinding {
		return vs[0].(map[string]*MetastoreServiceIamBinding)[vs[1].(string)]
	}).(MetastoreServiceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamBindingInput)(nil)).Elem(), &MetastoreServiceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamBindingArrayInput)(nil)).Elem(), MetastoreServiceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreServiceIamBindingMapInput)(nil)).Elem(), MetastoreServiceIamBindingMap{})
	pulumi.RegisterOutputType(MetastoreServiceIamBindingOutput{})
	pulumi.RegisterOutputType(MetastoreServiceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(MetastoreServiceIamBindingMapOutput{})
}
