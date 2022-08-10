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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/federations/{{federation_id}} * {{project}}/{{location}}/{{federation_id}} * {{location}}/{{federation_id}} * {{federation_id}} Any variables not passed in the import command will be taken from the provider configuration. Dataproc metastore federation IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamMember:MetastoreFederationIamMember editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamMember:MetastoreFederationIamMember editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamMember:MetastoreFederationIamMember editor projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MetastoreFederationIamMember struct {
	pulumi.CustomResourceState

	Condition MetastoreFederationIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringOutput `pulumi:"etag"`
	FederationId pulumi.StringOutput `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewMetastoreFederationIamMember registers a new resource with the given unique name, arguments, and options.
func NewMetastoreFederationIamMember(ctx *pulumi.Context,
	name string, args *MetastoreFederationIamMemberArgs, opts ...pulumi.ResourceOption) (*MetastoreFederationIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FederationId == nil {
		return nil, errors.New("invalid value for required argument 'FederationId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource MetastoreFederationIamMember
	err := ctx.RegisterResource("gcp:dataproc/metastoreFederationIamMember:MetastoreFederationIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreFederationIamMember gets an existing MetastoreFederationIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreFederationIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreFederationIamMemberState, opts ...pulumi.ResourceOption) (*MetastoreFederationIamMember, error) {
	var resource MetastoreFederationIamMember
	err := ctx.ReadResource("gcp:dataproc/metastoreFederationIamMember:MetastoreFederationIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreFederationIamMember resources.
type metastoreFederationIamMemberState struct {
	Condition *MetastoreFederationIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag         *string `pulumi:"etag"`
	FederationId *string `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type MetastoreFederationIamMemberState struct {
	Condition MetastoreFederationIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringPtrInput
	FederationId pulumi.StringPtrInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (MetastoreFederationIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamMemberState)(nil)).Elem()
}

type metastoreFederationIamMemberArgs struct {
	Condition    *MetastoreFederationIamMemberCondition `pulumi:"condition"`
	FederationId string                                 `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a MetastoreFederationIamMember resource.
type MetastoreFederationIamMemberArgs struct {
	Condition    MetastoreFederationIamMemberConditionPtrInput
	FederationId pulumi.StringInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (MetastoreFederationIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamMemberArgs)(nil)).Elem()
}

type MetastoreFederationIamMemberInput interface {
	pulumi.Input

	ToMetastoreFederationIamMemberOutput() MetastoreFederationIamMemberOutput
	ToMetastoreFederationIamMemberOutputWithContext(ctx context.Context) MetastoreFederationIamMemberOutput
}

func (*MetastoreFederationIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamMember)(nil)).Elem()
}

func (i *MetastoreFederationIamMember) ToMetastoreFederationIamMemberOutput() MetastoreFederationIamMemberOutput {
	return i.ToMetastoreFederationIamMemberOutputWithContext(context.Background())
}

func (i *MetastoreFederationIamMember) ToMetastoreFederationIamMemberOutputWithContext(ctx context.Context) MetastoreFederationIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamMemberOutput)
}

// MetastoreFederationIamMemberArrayInput is an input type that accepts MetastoreFederationIamMemberArray and MetastoreFederationIamMemberArrayOutput values.
// You can construct a concrete instance of `MetastoreFederationIamMemberArrayInput` via:
//
//	MetastoreFederationIamMemberArray{ MetastoreFederationIamMemberArgs{...} }
type MetastoreFederationIamMemberArrayInput interface {
	pulumi.Input

	ToMetastoreFederationIamMemberArrayOutput() MetastoreFederationIamMemberArrayOutput
	ToMetastoreFederationIamMemberArrayOutputWithContext(context.Context) MetastoreFederationIamMemberArrayOutput
}

type MetastoreFederationIamMemberArray []MetastoreFederationIamMemberInput

func (MetastoreFederationIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamMember)(nil)).Elem()
}

func (i MetastoreFederationIamMemberArray) ToMetastoreFederationIamMemberArrayOutput() MetastoreFederationIamMemberArrayOutput {
	return i.ToMetastoreFederationIamMemberArrayOutputWithContext(context.Background())
}

func (i MetastoreFederationIamMemberArray) ToMetastoreFederationIamMemberArrayOutputWithContext(ctx context.Context) MetastoreFederationIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamMemberArrayOutput)
}

// MetastoreFederationIamMemberMapInput is an input type that accepts MetastoreFederationIamMemberMap and MetastoreFederationIamMemberMapOutput values.
// You can construct a concrete instance of `MetastoreFederationIamMemberMapInput` via:
//
//	MetastoreFederationIamMemberMap{ "key": MetastoreFederationIamMemberArgs{...} }
type MetastoreFederationIamMemberMapInput interface {
	pulumi.Input

	ToMetastoreFederationIamMemberMapOutput() MetastoreFederationIamMemberMapOutput
	ToMetastoreFederationIamMemberMapOutputWithContext(context.Context) MetastoreFederationIamMemberMapOutput
}

type MetastoreFederationIamMemberMap map[string]MetastoreFederationIamMemberInput

func (MetastoreFederationIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamMember)(nil)).Elem()
}

func (i MetastoreFederationIamMemberMap) ToMetastoreFederationIamMemberMapOutput() MetastoreFederationIamMemberMapOutput {
	return i.ToMetastoreFederationIamMemberMapOutputWithContext(context.Background())
}

func (i MetastoreFederationIamMemberMap) ToMetastoreFederationIamMemberMapOutputWithContext(ctx context.Context) MetastoreFederationIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamMemberMapOutput)
}

type MetastoreFederationIamMemberOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamMember)(nil)).Elem()
}

func (o MetastoreFederationIamMemberOutput) ToMetastoreFederationIamMemberOutput() MetastoreFederationIamMemberOutput {
	return o
}

func (o MetastoreFederationIamMemberOutput) ToMetastoreFederationIamMemberOutputWithContext(ctx context.Context) MetastoreFederationIamMemberOutput {
	return o
}

func (o MetastoreFederationIamMemberOutput) Condition() MetastoreFederationIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) MetastoreFederationIamMemberConditionPtrOutput {
		return v.Condition
	}).(MetastoreFederationIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o MetastoreFederationIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MetastoreFederationIamMemberOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.FederationId }).(pulumi.StringOutput)
}

// The location where the metastore federation should reside.
// Used to find the parent resource to bind the IAM policy to
func (o MetastoreFederationIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MetastoreFederationIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MetastoreFederationIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `dataproc.MetastoreFederationIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o MetastoreFederationIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type MetastoreFederationIamMemberArrayOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamMember)(nil)).Elem()
}

func (o MetastoreFederationIamMemberArrayOutput) ToMetastoreFederationIamMemberArrayOutput() MetastoreFederationIamMemberArrayOutput {
	return o
}

func (o MetastoreFederationIamMemberArrayOutput) ToMetastoreFederationIamMemberArrayOutputWithContext(ctx context.Context) MetastoreFederationIamMemberArrayOutput {
	return o
}

func (o MetastoreFederationIamMemberArrayOutput) Index(i pulumi.IntInput) MetastoreFederationIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreFederationIamMember {
		return vs[0].([]*MetastoreFederationIamMember)[vs[1].(int)]
	}).(MetastoreFederationIamMemberOutput)
}

type MetastoreFederationIamMemberMapOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamMember)(nil)).Elem()
}

func (o MetastoreFederationIamMemberMapOutput) ToMetastoreFederationIamMemberMapOutput() MetastoreFederationIamMemberMapOutput {
	return o
}

func (o MetastoreFederationIamMemberMapOutput) ToMetastoreFederationIamMemberMapOutputWithContext(ctx context.Context) MetastoreFederationIamMemberMapOutput {
	return o
}

func (o MetastoreFederationIamMemberMapOutput) MapIndex(k pulumi.StringInput) MetastoreFederationIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreFederationIamMember {
		return vs[0].(map[string]*MetastoreFederationIamMember)[vs[1].(string)]
	}).(MetastoreFederationIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamMemberInput)(nil)).Elem(), &MetastoreFederationIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamMemberArrayInput)(nil)).Elem(), MetastoreFederationIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamMemberMapInput)(nil)).Elem(), MetastoreFederationIamMemberMap{})
	pulumi.RegisterOutputType(MetastoreFederationIamMemberOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamMemberArrayOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamMemberMapOutput{})
}
