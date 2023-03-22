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
//	$ pulumi import gcp:dataproc/metastoreFederationIamPolicy:MetastoreFederationIamPolicy editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamPolicy:MetastoreFederationIamPolicy editor "projects/{{project}}/locations/{{location}}/federations/{{federation_id}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:dataproc/metastoreFederationIamPolicy:MetastoreFederationIamPolicy editor projects/{{project}}/locations/{{location}}/federations/{{federation_id}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type MetastoreFederationIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringOutput `pulumi:"etag"`
	FederationId pulumi.StringOutput `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewMetastoreFederationIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewMetastoreFederationIamPolicy(ctx *pulumi.Context,
	name string, args *MetastoreFederationIamPolicyArgs, opts ...pulumi.ResourceOption) (*MetastoreFederationIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FederationId == nil {
		return nil, errors.New("invalid value for required argument 'FederationId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource MetastoreFederationIamPolicy
	err := ctx.RegisterResource("gcp:dataproc/metastoreFederationIamPolicy:MetastoreFederationIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetastoreFederationIamPolicy gets an existing MetastoreFederationIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetastoreFederationIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetastoreFederationIamPolicyState, opts ...pulumi.ResourceOption) (*MetastoreFederationIamPolicy, error) {
	var resource MetastoreFederationIamPolicy
	err := ctx.ReadResource("gcp:dataproc/metastoreFederationIamPolicy:MetastoreFederationIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetastoreFederationIamPolicy resources.
type metastoreFederationIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag         *string `pulumi:"etag"`
	FederationId *string `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type MetastoreFederationIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag         pulumi.StringPtrInput
	FederationId pulumi.StringPtrInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MetastoreFederationIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamPolicyState)(nil)).Elem()
}

type metastoreFederationIamPolicyArgs struct {
	FederationId string `pulumi:"federationId"`
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a MetastoreFederationIamPolicy resource.
type MetastoreFederationIamPolicyArgs struct {
	FederationId pulumi.StringInput
	// The location where the metastore federation should reside.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (MetastoreFederationIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metastoreFederationIamPolicyArgs)(nil)).Elem()
}

type MetastoreFederationIamPolicyInput interface {
	pulumi.Input

	ToMetastoreFederationIamPolicyOutput() MetastoreFederationIamPolicyOutput
	ToMetastoreFederationIamPolicyOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyOutput
}

func (*MetastoreFederationIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamPolicy)(nil)).Elem()
}

func (i *MetastoreFederationIamPolicy) ToMetastoreFederationIamPolicyOutput() MetastoreFederationIamPolicyOutput {
	return i.ToMetastoreFederationIamPolicyOutputWithContext(context.Background())
}

func (i *MetastoreFederationIamPolicy) ToMetastoreFederationIamPolicyOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamPolicyOutput)
}

// MetastoreFederationIamPolicyArrayInput is an input type that accepts MetastoreFederationIamPolicyArray and MetastoreFederationIamPolicyArrayOutput values.
// You can construct a concrete instance of `MetastoreFederationIamPolicyArrayInput` via:
//
//	MetastoreFederationIamPolicyArray{ MetastoreFederationIamPolicyArgs{...} }
type MetastoreFederationIamPolicyArrayInput interface {
	pulumi.Input

	ToMetastoreFederationIamPolicyArrayOutput() MetastoreFederationIamPolicyArrayOutput
	ToMetastoreFederationIamPolicyArrayOutputWithContext(context.Context) MetastoreFederationIamPolicyArrayOutput
}

type MetastoreFederationIamPolicyArray []MetastoreFederationIamPolicyInput

func (MetastoreFederationIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamPolicy)(nil)).Elem()
}

func (i MetastoreFederationIamPolicyArray) ToMetastoreFederationIamPolicyArrayOutput() MetastoreFederationIamPolicyArrayOutput {
	return i.ToMetastoreFederationIamPolicyArrayOutputWithContext(context.Background())
}

func (i MetastoreFederationIamPolicyArray) ToMetastoreFederationIamPolicyArrayOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamPolicyArrayOutput)
}

// MetastoreFederationIamPolicyMapInput is an input type that accepts MetastoreFederationIamPolicyMap and MetastoreFederationIamPolicyMapOutput values.
// You can construct a concrete instance of `MetastoreFederationIamPolicyMapInput` via:
//
//	MetastoreFederationIamPolicyMap{ "key": MetastoreFederationIamPolicyArgs{...} }
type MetastoreFederationIamPolicyMapInput interface {
	pulumi.Input

	ToMetastoreFederationIamPolicyMapOutput() MetastoreFederationIamPolicyMapOutput
	ToMetastoreFederationIamPolicyMapOutputWithContext(context.Context) MetastoreFederationIamPolicyMapOutput
}

type MetastoreFederationIamPolicyMap map[string]MetastoreFederationIamPolicyInput

func (MetastoreFederationIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamPolicy)(nil)).Elem()
}

func (i MetastoreFederationIamPolicyMap) ToMetastoreFederationIamPolicyMapOutput() MetastoreFederationIamPolicyMapOutput {
	return i.ToMetastoreFederationIamPolicyMapOutputWithContext(context.Background())
}

func (i MetastoreFederationIamPolicyMap) ToMetastoreFederationIamPolicyMapOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetastoreFederationIamPolicyMapOutput)
}

type MetastoreFederationIamPolicyOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetastoreFederationIamPolicy)(nil)).Elem()
}

func (o MetastoreFederationIamPolicyOutput) ToMetastoreFederationIamPolicyOutput() MetastoreFederationIamPolicyOutput {
	return o
}

func (o MetastoreFederationIamPolicyOutput) ToMetastoreFederationIamPolicyOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o MetastoreFederationIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MetastoreFederationIamPolicyOutput) FederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamPolicy) pulumi.StringOutput { return v.FederationId }).(pulumi.StringOutput)
}

// The location where the metastore federation should reside.
// Used to find the parent resource to bind the IAM policy to
func (o MetastoreFederationIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o MetastoreFederationIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o MetastoreFederationIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetastoreFederationIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type MetastoreFederationIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetastoreFederationIamPolicy)(nil)).Elem()
}

func (o MetastoreFederationIamPolicyArrayOutput) ToMetastoreFederationIamPolicyArrayOutput() MetastoreFederationIamPolicyArrayOutput {
	return o
}

func (o MetastoreFederationIamPolicyArrayOutput) ToMetastoreFederationIamPolicyArrayOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyArrayOutput {
	return o
}

func (o MetastoreFederationIamPolicyArrayOutput) Index(i pulumi.IntInput) MetastoreFederationIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetastoreFederationIamPolicy {
		return vs[0].([]*MetastoreFederationIamPolicy)[vs[1].(int)]
	}).(MetastoreFederationIamPolicyOutput)
}

type MetastoreFederationIamPolicyMapOutput struct{ *pulumi.OutputState }

func (MetastoreFederationIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetastoreFederationIamPolicy)(nil)).Elem()
}

func (o MetastoreFederationIamPolicyMapOutput) ToMetastoreFederationIamPolicyMapOutput() MetastoreFederationIamPolicyMapOutput {
	return o
}

func (o MetastoreFederationIamPolicyMapOutput) ToMetastoreFederationIamPolicyMapOutputWithContext(ctx context.Context) MetastoreFederationIamPolicyMapOutput {
	return o
}

func (o MetastoreFederationIamPolicyMapOutput) MapIndex(k pulumi.StringInput) MetastoreFederationIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetastoreFederationIamPolicy {
		return vs[0].(map[string]*MetastoreFederationIamPolicy)[vs[1].(string)]
	}).(MetastoreFederationIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamPolicyInput)(nil)).Elem(), &MetastoreFederationIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamPolicyArrayInput)(nil)).Elem(), MetastoreFederationIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetastoreFederationIamPolicyMapInput)(nil)).Elem(), MetastoreFederationIamPolicyMap{})
	pulumi.RegisterOutputType(MetastoreFederationIamPolicyOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(MetastoreFederationIamPolicyMapOutput{})
}
