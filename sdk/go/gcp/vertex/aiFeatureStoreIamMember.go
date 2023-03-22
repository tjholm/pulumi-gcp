// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/featurestores/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Vertex AI featurestore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember editor "projects/{{project}}/locations/{{region}}/featurestores/{{featurestore}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember editor "projects/{{project}}/locations/{{region}}/featurestores/{{featurestore}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember editor projects/{{project}}/locations/{{region}}/featurestores/{{featurestore}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AiFeatureStoreIamMember struct {
	pulumi.CustomResourceState

	Condition AiFeatureStoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringOutput `pulumi:"featurestore"`
	Member       pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAiFeatureStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureStoreIamMember(ctx *pulumi.Context,
	name string, args *AiFeatureStoreIamMemberArgs, opts ...pulumi.ResourceOption) (*AiFeatureStoreIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Featurestore == nil {
		return nil, errors.New("invalid value for required argument 'Featurestore'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource AiFeatureStoreIamMember
	err := ctx.RegisterResource("gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureStoreIamMember gets an existing AiFeatureStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureStoreIamMemberState, opts ...pulumi.ResourceOption) (*AiFeatureStoreIamMember, error) {
	var resource AiFeatureStoreIamMember
	err := ctx.ReadResource("gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureStoreIamMember resources.
type aiFeatureStoreIamMemberState struct {
	Condition *AiFeatureStoreIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Featurestore *string `pulumi:"featurestore"`
	Member       *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type AiFeatureStoreIamMemberState struct {
	Condition AiFeatureStoreIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringPtrInput
	Member       pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (AiFeatureStoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreIamMemberState)(nil)).Elem()
}

type aiFeatureStoreIamMemberArgs struct {
	Condition *AiFeatureStoreIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Featurestore string `pulumi:"featurestore"`
	Member       string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AiFeatureStoreIamMember resource.
type AiFeatureStoreIamMemberArgs struct {
	Condition AiFeatureStoreIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringInput
	Member       pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (AiFeatureStoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreIamMemberArgs)(nil)).Elem()
}

type AiFeatureStoreIamMemberInput interface {
	pulumi.Input

	ToAiFeatureStoreIamMemberOutput() AiFeatureStoreIamMemberOutput
	ToAiFeatureStoreIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberOutput
}

func (*AiFeatureStoreIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreIamMember)(nil)).Elem()
}

func (i *AiFeatureStoreIamMember) ToAiFeatureStoreIamMemberOutput() AiFeatureStoreIamMemberOutput {
	return i.ToAiFeatureStoreIamMemberOutputWithContext(context.Background())
}

func (i *AiFeatureStoreIamMember) ToAiFeatureStoreIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreIamMemberOutput)
}

// AiFeatureStoreIamMemberArrayInput is an input type that accepts AiFeatureStoreIamMemberArray and AiFeatureStoreIamMemberArrayOutput values.
// You can construct a concrete instance of `AiFeatureStoreIamMemberArrayInput` via:
//
//	AiFeatureStoreIamMemberArray{ AiFeatureStoreIamMemberArgs{...} }
type AiFeatureStoreIamMemberArrayInput interface {
	pulumi.Input

	ToAiFeatureStoreIamMemberArrayOutput() AiFeatureStoreIamMemberArrayOutput
	ToAiFeatureStoreIamMemberArrayOutputWithContext(context.Context) AiFeatureStoreIamMemberArrayOutput
}

type AiFeatureStoreIamMemberArray []AiFeatureStoreIamMemberInput

func (AiFeatureStoreIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreIamMember)(nil)).Elem()
}

func (i AiFeatureStoreIamMemberArray) ToAiFeatureStoreIamMemberArrayOutput() AiFeatureStoreIamMemberArrayOutput {
	return i.ToAiFeatureStoreIamMemberArrayOutputWithContext(context.Background())
}

func (i AiFeatureStoreIamMemberArray) ToAiFeatureStoreIamMemberArrayOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreIamMemberArrayOutput)
}

// AiFeatureStoreIamMemberMapInput is an input type that accepts AiFeatureStoreIamMemberMap and AiFeatureStoreIamMemberMapOutput values.
// You can construct a concrete instance of `AiFeatureStoreIamMemberMapInput` via:
//
//	AiFeatureStoreIamMemberMap{ "key": AiFeatureStoreIamMemberArgs{...} }
type AiFeatureStoreIamMemberMapInput interface {
	pulumi.Input

	ToAiFeatureStoreIamMemberMapOutput() AiFeatureStoreIamMemberMapOutput
	ToAiFeatureStoreIamMemberMapOutputWithContext(context.Context) AiFeatureStoreIamMemberMapOutput
}

type AiFeatureStoreIamMemberMap map[string]AiFeatureStoreIamMemberInput

func (AiFeatureStoreIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreIamMember)(nil)).Elem()
}

func (i AiFeatureStoreIamMemberMap) ToAiFeatureStoreIamMemberMapOutput() AiFeatureStoreIamMemberMapOutput {
	return i.ToAiFeatureStoreIamMemberMapOutputWithContext(context.Background())
}

func (i AiFeatureStoreIamMemberMap) ToAiFeatureStoreIamMemberMapOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreIamMemberMapOutput)
}

type AiFeatureStoreIamMemberOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreIamMember)(nil)).Elem()
}

func (o AiFeatureStoreIamMemberOutput) ToAiFeatureStoreIamMemberOutput() AiFeatureStoreIamMemberOutput {
	return o
}

func (o AiFeatureStoreIamMemberOutput) ToAiFeatureStoreIamMemberOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberOutput {
	return o
}

func (o AiFeatureStoreIamMemberOutput) Condition() AiFeatureStoreIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) AiFeatureStoreIamMemberConditionPtrOutput { return v.Condition }).(AiFeatureStoreIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o AiFeatureStoreIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o AiFeatureStoreIamMemberOutput) Featurestore() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Featurestore }).(pulumi.StringOutput)
}

func (o AiFeatureStoreIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o AiFeatureStoreIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the dataset. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (o AiFeatureStoreIamMemberOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `vertex.AiFeatureStoreIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o AiFeatureStoreIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AiFeatureStoreIamMemberArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreIamMember)(nil)).Elem()
}

func (o AiFeatureStoreIamMemberArrayOutput) ToAiFeatureStoreIamMemberArrayOutput() AiFeatureStoreIamMemberArrayOutput {
	return o
}

func (o AiFeatureStoreIamMemberArrayOutput) ToAiFeatureStoreIamMemberArrayOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberArrayOutput {
	return o
}

func (o AiFeatureStoreIamMemberArrayOutput) Index(i pulumi.IntInput) AiFeatureStoreIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiFeatureStoreIamMember {
		return vs[0].([]*AiFeatureStoreIamMember)[vs[1].(int)]
	}).(AiFeatureStoreIamMemberOutput)
}

type AiFeatureStoreIamMemberMapOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreIamMember)(nil)).Elem()
}

func (o AiFeatureStoreIamMemberMapOutput) ToAiFeatureStoreIamMemberMapOutput() AiFeatureStoreIamMemberMapOutput {
	return o
}

func (o AiFeatureStoreIamMemberMapOutput) ToAiFeatureStoreIamMemberMapOutputWithContext(ctx context.Context) AiFeatureStoreIamMemberMapOutput {
	return o
}

func (o AiFeatureStoreIamMemberMapOutput) MapIndex(k pulumi.StringInput) AiFeatureStoreIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiFeatureStoreIamMember {
		return vs[0].(map[string]*AiFeatureStoreIamMember)[vs[1].(string)]
	}).(AiFeatureStoreIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreIamMemberInput)(nil)).Elem(), &AiFeatureStoreIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreIamMemberArrayInput)(nil)).Elem(), AiFeatureStoreIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreIamMemberMapInput)(nil)).Elem(), AiFeatureStoreIamMemberMap{})
	pulumi.RegisterOutputType(AiFeatureStoreIamMemberOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreIamMemberArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreIamMemberMapOutput{})
}
