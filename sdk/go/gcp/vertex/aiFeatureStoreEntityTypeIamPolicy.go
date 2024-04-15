// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

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
// * {{featurestore}}/entityTypes/{{name}}
//
// * {{name}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Vertex AI featurestoreentitytype IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor "{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor "{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor {{featurestore}}/entityTypes/{{featurestore_entitytype}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AiFeatureStoreEntityTypeIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringOutput `pulumi:"entitytype"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringOutput `pulumi:"featurestore"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewAiFeatureStoreEntityTypeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureStoreEntityTypeIamPolicy(ctx *pulumi.Context,
	name string, args *AiFeatureStoreEntityTypeIamPolicyArgs, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entitytype == nil {
		return nil, errors.New("invalid value for required argument 'Entitytype'")
	}
	if args.Featurestore == nil {
		return nil, errors.New("invalid value for required argument 'Featurestore'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiFeatureStoreEntityTypeIamPolicy
	err := ctx.RegisterResource("gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureStoreEntityTypeIamPolicy gets an existing AiFeatureStoreEntityTypeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureStoreEntityTypeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureStoreEntityTypeIamPolicyState, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeIamPolicy, error) {
	var resource AiFeatureStoreEntityTypeIamPolicy
	err := ctx.ReadResource("gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureStoreEntityTypeIamPolicy resources.
type aiFeatureStoreEntityTypeIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Entitytype *string `pulumi:"entitytype"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore *string `pulumi:"featurestore"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type AiFeatureStoreEntityTypeIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (AiFeatureStoreEntityTypeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeIamPolicyState)(nil)).Elem()
}

type aiFeatureStoreEntityTypeIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Entitytype string `pulumi:"entitytype"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore string `pulumi:"featurestore"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a AiFeatureStoreEntityTypeIamPolicy resource.
type AiFeatureStoreEntityTypeIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Entitytype pulumi.StringInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
	Featurestore pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (AiFeatureStoreEntityTypeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeIamPolicyArgs)(nil)).Elem()
}

type AiFeatureStoreEntityTypeIamPolicyInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamPolicyOutput() AiFeatureStoreEntityTypeIamPolicyOutput
	ToAiFeatureStoreEntityTypeIamPolicyOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyOutput
}

func (*AiFeatureStoreEntityTypeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (i *AiFeatureStoreEntityTypeIamPolicy) ToAiFeatureStoreEntityTypeIamPolicyOutput() AiFeatureStoreEntityTypeIamPolicyOutput {
	return i.ToAiFeatureStoreEntityTypeIamPolicyOutputWithContext(context.Background())
}

func (i *AiFeatureStoreEntityTypeIamPolicy) ToAiFeatureStoreEntityTypeIamPolicyOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamPolicyOutput)
}

// AiFeatureStoreEntityTypeIamPolicyArrayInput is an input type that accepts AiFeatureStoreEntityTypeIamPolicyArray and AiFeatureStoreEntityTypeIamPolicyArrayOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeIamPolicyArrayInput` via:
//
//	AiFeatureStoreEntityTypeIamPolicyArray{ AiFeatureStoreEntityTypeIamPolicyArgs{...} }
type AiFeatureStoreEntityTypeIamPolicyArrayInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamPolicyArrayOutput() AiFeatureStoreEntityTypeIamPolicyArrayOutput
	ToAiFeatureStoreEntityTypeIamPolicyArrayOutputWithContext(context.Context) AiFeatureStoreEntityTypeIamPolicyArrayOutput
}

type AiFeatureStoreEntityTypeIamPolicyArray []AiFeatureStoreEntityTypeIamPolicyInput

func (AiFeatureStoreEntityTypeIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeIamPolicyArray) ToAiFeatureStoreEntityTypeIamPolicyArrayOutput() AiFeatureStoreEntityTypeIamPolicyArrayOutput {
	return i.ToAiFeatureStoreEntityTypeIamPolicyArrayOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeIamPolicyArray) ToAiFeatureStoreEntityTypeIamPolicyArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamPolicyArrayOutput)
}

// AiFeatureStoreEntityTypeIamPolicyMapInput is an input type that accepts AiFeatureStoreEntityTypeIamPolicyMap and AiFeatureStoreEntityTypeIamPolicyMapOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeIamPolicyMapInput` via:
//
//	AiFeatureStoreEntityTypeIamPolicyMap{ "key": AiFeatureStoreEntityTypeIamPolicyArgs{...} }
type AiFeatureStoreEntityTypeIamPolicyMapInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeIamPolicyMapOutput() AiFeatureStoreEntityTypeIamPolicyMapOutput
	ToAiFeatureStoreEntityTypeIamPolicyMapOutputWithContext(context.Context) AiFeatureStoreEntityTypeIamPolicyMapOutput
}

type AiFeatureStoreEntityTypeIamPolicyMap map[string]AiFeatureStoreEntityTypeIamPolicyInput

func (AiFeatureStoreEntityTypeIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeIamPolicyMap) ToAiFeatureStoreEntityTypeIamPolicyMapOutput() AiFeatureStoreEntityTypeIamPolicyMapOutput {
	return i.ToAiFeatureStoreEntityTypeIamPolicyMapOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeIamPolicyMap) ToAiFeatureStoreEntityTypeIamPolicyMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeIamPolicyMapOutput)
}

type AiFeatureStoreEntityTypeIamPolicyOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamPolicyOutput) ToAiFeatureStoreEntityTypeIamPolicyOutput() AiFeatureStoreEntityTypeIamPolicyOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamPolicyOutput) ToAiFeatureStoreEntityTypeIamPolicyOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o AiFeatureStoreEntityTypeIamPolicyOutput) Entitytype() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamPolicy) pulumi.StringOutput { return v.Entitytype }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o AiFeatureStoreEntityTypeIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
func (o AiFeatureStoreEntityTypeIamPolicyOutput) Featurestore() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamPolicy) pulumi.StringOutput { return v.Featurestore }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o AiFeatureStoreEntityTypeIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type AiFeatureStoreEntityTypeIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamPolicyArrayOutput) ToAiFeatureStoreEntityTypeIamPolicyArrayOutput() AiFeatureStoreEntityTypeIamPolicyArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamPolicyArrayOutput) ToAiFeatureStoreEntityTypeIamPolicyArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamPolicyArrayOutput) Index(i pulumi.IntInput) AiFeatureStoreEntityTypeIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeIamPolicy {
		return vs[0].([]*AiFeatureStoreEntityTypeIamPolicy)[vs[1].(int)]
	}).(AiFeatureStoreEntityTypeIamPolicyOutput)
}

type AiFeatureStoreEntityTypeIamPolicyMapOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeIamPolicy)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeIamPolicyMapOutput) ToAiFeatureStoreEntityTypeIamPolicyMapOutput() AiFeatureStoreEntityTypeIamPolicyMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamPolicyMapOutput) ToAiFeatureStoreEntityTypeIamPolicyMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeIamPolicyMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeIamPolicyMapOutput) MapIndex(k pulumi.StringInput) AiFeatureStoreEntityTypeIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeIamPolicy {
		return vs[0].(map[string]*AiFeatureStoreEntityTypeIamPolicy)[vs[1].(string)]
	}).(AiFeatureStoreEntityTypeIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamPolicyInput)(nil)).Elem(), &AiFeatureStoreEntityTypeIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamPolicyArrayInput)(nil)).Elem(), AiFeatureStoreEntityTypeIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeIamPolicyMapInput)(nil)).Elem(), AiFeatureStoreEntityTypeIamPolicyMap{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamPolicyOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeIamPolicyMapOutput{})
}
