// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/federations/{{federation_id}} * {{project}}/{{location}}/{{federation_id}} * {{location}}/{{federation_id}} * {{federation_id}} Any variables not passed in the import command will be taken from the provider configuration. Dataproc metastore federation IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamBinding:MetastoreFederationIamBinding editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamBinding:MetastoreFederationIamBinding editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamBinding:MetastoreFederationIamBinding editor projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MetastoreFederationIamBinding struct {
	pulumi.CustomResourceState

	Condition MetastoreFederationIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringOutput `pulumi:"etag"`
	FederationId pulumi.StringOutput `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewMetastoreFederationIamBinding registers a new resource with the given unique name, arguments, and options.
func NewMetastoreFederationIamBinding(ctx *pulumi.Context,
	name string, args *MetastoreFederationIamBindingArgs, opts ...pulumi.ResourceOption) (*MetastoreFederationIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FederationId == nil {
		return nil, errors.New("invalid value for required argument 'FederationId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource MetastoreFederationIamBinding
	err := ctx.RegisterResource("gcp:dataproc/metastoreFederationIamBinding:MetastoreFederationIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreFederationIamBinding gets an existing MetastoreFederationIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreFederationIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreFederationIamBindingState, opts ...pulumi.ResourceOption) (*MetastoreFederationIamBinding, error) {
	var resource MetastoreFederationIamBinding
	err := ctx.ReadResource("gcp:dataproc/metastoreFederationIamBinding:MetastoreFederationIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreFederationIamBinding resources.
type metastoreFederationIamBindingState struct {
	Condition *MetastoreFederationIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag         *string `pulumi:"etag"`
	FederationId *string `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type MetastoreFederationIamBindingState struct {
	Condition MetastoreFederationIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringPtrInput
	FederationId pulumi.StringPtrInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (MetastoreFederationIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamBindingState)(nil)).Elem()
}

type metastoreFederationIamBindingArgs struct {
	Condition    *MetastoreFederationIamBindingCondition `pulumi:"condition"`
	FederationId string                                  `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a MetastoreFederationIamBinding resource.
type MetastoreFederationIamBindingArgs struct {
	Condition    MetastoreFederationIamBindingConditionPtrInput
	FederationId pulumi.StringInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (MetastoreFederationIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamBindingArgs)(nil)).Elem()
}

type MetastoreFederationIamBindingInput interface {
	pulumi.Input

	ToMetastoreFederationIamBindingOutput() MetastoreFederationIamBindingOutput
	ToMetastoreFederationIamBindingOutputWithContext(ctx context.Context) MetastoreFederationIamBindingOutput
}

func (*MetastoreFederationIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamBinding)(nil)).Elem()
}

func (i *MetastoreFederationIamBinding) ToMetastoreFederationIamBindingOutput() MetastoreFederationIamBindingOutput {
	return i.ToMetastoreFederationIamBindingOutputWithContext(context.Background())
}

func (i *MetastoreFederationIamBinding) ToMetastoreFederationIamBindingOutputWithContext(ctx context.Context) MetastoreFederationIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamBindingOutput)
}

// MetastoreFederationIamBindingArrayInput is an input type that accepts MetastoreFederationIamBindingArray and MetastoreFederationIamBindingArrayOutput values.
// You can construct a concrete instance of `MetastoreFederationIamBindingArrayInput` via:
//
//	MetastoreFederationIamBindingArray{ MetastoreFederationIamBindingArgs{...} }
type MetastoreFederationIamBindingArrayInput interface {
	pulumi.Input

	ToMetastoreFederationIamBindingArrayOutput() MetastoreFederationIamBindingArrayOutput
	ToMetastoreFederationIamBindingArrayOutputWithContext(context.Context) MetastoreFederationIamBindingArrayOutput
}

type MetastoreFederationIamBindingArray []MetastoreFederationIamBindingInput

func (MetastoreFederationIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamBinding)(nil)).Elem()
}

func (i MetastoreFederationIamBindingArray) ToMetastoreFederationIamBindingArrayOutput() MetastoreFederationIamBindingArrayOutput {
	return i.ToMetastoreFederationIamBindingArrayOutputWithContext(context.Background())
}

func (i MetastoreFederationIamBindingArray) ToMetastoreFederationIamBindingArrayOutputWithContext(ctx context.Context) MetastoreFederationIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamBindingArrayOutput)
}

// MetastoreFederationIamBindingMapInput is an input type that accepts MetastoreFederationIamBindingMap and MetastoreFederationIamBindingMapOutput values.
// You can construct a concrete instance of `MetastoreFederationIamBindingMapInput` via:
//
//	MetastoreFederationIamBindingMap{ "key": MetastoreFederationIamBindingArgs{...} }
type MetastoreFederationIamBindingMapInput interface {
	pulumi.Input

	ToMetastoreFederationIamBindingMapOutput() MetastoreFederationIamBindingMapOutput
	ToMetastoreFederationIamBindingMapOutputWithContext(context.Context) MetastoreFederationIamBindingMapOutput
}

type MetastoreFederationIamBindingMap map[string]MetastoreFederationIamBindingInput

func (MetastoreFederationIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamBinding)(nil)).Elem()
}

func (i MetastoreFederationIamBindingMap) ToMetastoreFederationIamBindingMapOutput() MetastoreFederationIamBindingMapOutput {
	return i.ToMetastoreFederationIamBindingMapOutputWithContext(context.Background())
}

func (i MetastoreFederationIamBindingMap) ToMetastoreFederationIamBindingMapOutputWithContext(ctx context.Context) MetastoreFederationIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamBindingMapOutput)
}

type MetastoreFederationIamBindingOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamBinding)(nil)).Elem()
}

func (o MetastoreFederationIamBindingOutput) ToMetastoreFederationIamBindingOutput() MetastoreFederationIamBindingOutput {
	return o
}

func (o MetastoreFederationIamBindingOutput) ToMetastoreFederationIamBindingOutputWithContext(ctx context.Context) MetastoreFederationIamBindingOutput {
	return o
}

func (o MetastoreFederationIamBindingOutput) Condition() MetastoreFederationIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) MetastoreFederationIamBindingConditionPtrOutput {
		return v.Condition
	}).(MetastoreFederationIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o MetastoreFederationIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MetastoreFederationIamBindingOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringOutput { return v.FederationId }).(pulumi.StringOutput)
}

// The location where the metastore federation should reside.
// Used to find the parent resource to bind the IAM policy to
func (o MetastoreFederationIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MetastoreFederationIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MetastoreFederationIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o MetastoreFederationIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type MetastoreFederationIamBindingArrayOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamBinding)(nil)).Elem()
}

func (o MetastoreFederationIamBindingArrayOutput) ToMetastoreFederationIamBindingArrayOutput() MetastoreFederationIamBindingArrayOutput {
	return o
}

func (o MetastoreFederationIamBindingArrayOutput) ToMetastoreFederationIamBindingArrayOutputWithContext(ctx context.Context) MetastoreFederationIamBindingArrayOutput {
	return o
}

func (o MetastoreFederationIamBindingArrayOutput) Index(i pulumi.IntInput) MetastoreFederationIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreFederationIamBinding {
		return vs[0].([]*MetastoreFederationIamBinding)[vs[1].(int)]
	}).(MetastoreFederationIamBindingOutput)
}

type MetastoreFederationIamBindingMapOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamBinding)(nil)).Elem()
}

func (o MetastoreFederationIamBindingMapOutput) ToMetastoreFederationIamBindingMapOutput() MetastoreFederationIamBindingMapOutput {
	return o
}

func (o MetastoreFederationIamBindingMapOutput) ToMetastoreFederationIamBindingMapOutputWithContext(ctx context.Context) MetastoreFederationIamBindingMapOutput {
	return o
}

func (o MetastoreFederationIamBindingMapOutput) MapIndex(k pulumi.StringInput) MetastoreFederationIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreFederationIamBinding {
		return vs[0].(map[string]*MetastoreFederationIamBinding)[vs[1].(string)]
	}).(MetastoreFederationIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamBindingInput)(nil)).Elem(), &MetastoreFederationIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamBindingArrayInput)(nil)).Elem(), MetastoreFederationIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamBindingMapInput)(nil)).Elem(), MetastoreFederationIamBindingMap{})
	pulumi.RegisterOutputType(MetastoreFederationIamBindingOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamBindingArrayOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamBindingMapOutput{})
}
