// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagTemplateIamPolicy struct {
	pulumi.CustomResourceState

	Etag        pulumi.StringOutput `pulumi:"etag"`
	PolicyData  pulumi.StringOutput `pulumi:"policyData"`
	Project     pulumi.StringOutput `pulumi:"project"`
	Region      pulumi.StringOutput `pulumi:"region"`
	TagTemplate pulumi.StringOutput `pulumi:"tagTemplate"`
}

// NewTagTemplateIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTagTemplateIamPolicy(ctx *pulumi.Context,
	name string, args *TagTemplateIamPolicyArgs, opts ...pulumi.ResourceOption) (*TagTemplateIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.TagTemplate == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplate'")
	}
	var resource TagTemplateIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/tagTemplateIamPolicy:TagTemplateIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplateIamPolicy gets an existing TagTemplateIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplateIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateIamPolicyState, opts ...pulumi.ResourceOption) (*TagTemplateIamPolicy, error) {
	var resource TagTemplateIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/tagTemplateIamPolicy:TagTemplateIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplateIamPolicy resources.
type tagTemplateIamPolicyState struct {
	Etag        *string `pulumi:"etag"`
	PolicyData  *string `pulumi:"policyData"`
	Project     *string `pulumi:"project"`
	Region      *string `pulumi:"region"`
	TagTemplate *string `pulumi:"tagTemplate"`
}

type TagTemplateIamPolicyState struct {
	Etag        pulumi.StringPtrInput
	PolicyData  pulumi.StringPtrInput
	Project     pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	TagTemplate pulumi.StringPtrInput
}

func (TagTemplateIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamPolicyState)(nil)).Elem()
}

type tagTemplateIamPolicyArgs struct {
	PolicyData  string  `pulumi:"policyData"`
	Project     *string `pulumi:"project"`
	Region      *string `pulumi:"region"`
	TagTemplate string  `pulumi:"tagTemplate"`
}

// The set of arguments for constructing a TagTemplateIamPolicy resource.
type TagTemplateIamPolicyArgs struct {
	PolicyData  pulumi.StringInput
	Project     pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	TagTemplate pulumi.StringInput
}

func (TagTemplateIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamPolicyArgs)(nil)).Elem()
}

type TagTemplateIamPolicyInput interface {
	pulumi.Input

	ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput
	ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput
}

func (*TagTemplateIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplateIamPolicy)(nil))
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput {
	return i.ToTagTemplateIamPolicyOutputWithContext(context.Background())
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyOutput)
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyPtrOutput() TagTemplateIamPolicyPtrOutput {
	return i.ToTagTemplateIamPolicyPtrOutputWithContext(context.Background())
}

func (i *TagTemplateIamPolicy) ToTagTemplateIamPolicyPtrOutputWithContext(ctx context.Context) TagTemplateIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyPtrOutput)
}

type TagTemplateIamPolicyPtrInput interface {
	pulumi.Input

	ToTagTemplateIamPolicyPtrOutput() TagTemplateIamPolicyPtrOutput
	ToTagTemplateIamPolicyPtrOutputWithContext(ctx context.Context) TagTemplateIamPolicyPtrOutput
}

type tagTemplateIamPolicyPtrType TagTemplateIamPolicyArgs

func (*tagTemplateIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamPolicy)(nil))
}

func (i *tagTemplateIamPolicyPtrType) ToTagTemplateIamPolicyPtrOutput() TagTemplateIamPolicyPtrOutput {
	return i.ToTagTemplateIamPolicyPtrOutputWithContext(context.Background())
}

func (i *tagTemplateIamPolicyPtrType) ToTagTemplateIamPolicyPtrOutputWithContext(ctx context.Context) TagTemplateIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyPtrOutput)
}

// TagTemplateIamPolicyArrayInput is an input type that accepts TagTemplateIamPolicyArray and TagTemplateIamPolicyArrayOutput values.
// You can construct a concrete instance of `TagTemplateIamPolicyArrayInput` via:
//
//          TagTemplateIamPolicyArray{ TagTemplateIamPolicyArgs{...} }
type TagTemplateIamPolicyArrayInput interface {
	pulumi.Input

	ToTagTemplateIamPolicyArrayOutput() TagTemplateIamPolicyArrayOutput
	ToTagTemplateIamPolicyArrayOutputWithContext(context.Context) TagTemplateIamPolicyArrayOutput
}

type TagTemplateIamPolicyArray []TagTemplateIamPolicyInput

func (TagTemplateIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TagTemplateIamPolicy)(nil))
}

func (i TagTemplateIamPolicyArray) ToTagTemplateIamPolicyArrayOutput() TagTemplateIamPolicyArrayOutput {
	return i.ToTagTemplateIamPolicyArrayOutputWithContext(context.Background())
}

func (i TagTemplateIamPolicyArray) ToTagTemplateIamPolicyArrayOutputWithContext(ctx context.Context) TagTemplateIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyArrayOutput)
}

// TagTemplateIamPolicyMapInput is an input type that accepts TagTemplateIamPolicyMap and TagTemplateIamPolicyMapOutput values.
// You can construct a concrete instance of `TagTemplateIamPolicyMapInput` via:
//
//          TagTemplateIamPolicyMap{ "key": TagTemplateIamPolicyArgs{...} }
type TagTemplateIamPolicyMapInput interface {
	pulumi.Input

	ToTagTemplateIamPolicyMapOutput() TagTemplateIamPolicyMapOutput
	ToTagTemplateIamPolicyMapOutputWithContext(context.Context) TagTemplateIamPolicyMapOutput
}

type TagTemplateIamPolicyMap map[string]TagTemplateIamPolicyInput

func (TagTemplateIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TagTemplateIamPolicy)(nil))
}

func (i TagTemplateIamPolicyMap) ToTagTemplateIamPolicyMapOutput() TagTemplateIamPolicyMapOutput {
	return i.ToTagTemplateIamPolicyMapOutputWithContext(context.Background())
}

func (i TagTemplateIamPolicyMap) ToTagTemplateIamPolicyMapOutputWithContext(ctx context.Context) TagTemplateIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamPolicyMapOutput)
}

type TagTemplateIamPolicyOutput struct {
	*pulumi.OutputState
}

func (TagTemplateIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagTemplateIamPolicy)(nil))
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyOutput() TagTemplateIamPolicyOutput {
	return o
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyOutputWithContext(ctx context.Context) TagTemplateIamPolicyOutput {
	return o
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyPtrOutput() TagTemplateIamPolicyPtrOutput {
	return o.ToTagTemplateIamPolicyPtrOutputWithContext(context.Background())
}

func (o TagTemplateIamPolicyOutput) ToTagTemplateIamPolicyPtrOutputWithContext(ctx context.Context) TagTemplateIamPolicyPtrOutput {
	return o.ApplyT(func(v TagTemplateIamPolicy) *TagTemplateIamPolicy {
		return &v
	}).(TagTemplateIamPolicyPtrOutput)
}

type TagTemplateIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (TagTemplateIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamPolicy)(nil))
}

func (o TagTemplateIamPolicyPtrOutput) ToTagTemplateIamPolicyPtrOutput() TagTemplateIamPolicyPtrOutput {
	return o
}

func (o TagTemplateIamPolicyPtrOutput) ToTagTemplateIamPolicyPtrOutputWithContext(ctx context.Context) TagTemplateIamPolicyPtrOutput {
	return o
}

type TagTemplateIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TagTemplateIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagTemplateIamPolicy)(nil))
}

func (o TagTemplateIamPolicyArrayOutput) ToTagTemplateIamPolicyArrayOutput() TagTemplateIamPolicyArrayOutput {
	return o
}

func (o TagTemplateIamPolicyArrayOutput) ToTagTemplateIamPolicyArrayOutputWithContext(ctx context.Context) TagTemplateIamPolicyArrayOutput {
	return o
}

func (o TagTemplateIamPolicyArrayOutput) Index(i pulumi.IntInput) TagTemplateIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagTemplateIamPolicy {
		return vs[0].([]TagTemplateIamPolicy)[vs[1].(int)]
	}).(TagTemplateIamPolicyOutput)
}

type TagTemplateIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TagTemplateIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagTemplateIamPolicy)(nil))
}

func (o TagTemplateIamPolicyMapOutput) ToTagTemplateIamPolicyMapOutput() TagTemplateIamPolicyMapOutput {
	return o
}

func (o TagTemplateIamPolicyMapOutput) ToTagTemplateIamPolicyMapOutputWithContext(ctx context.Context) TagTemplateIamPolicyMapOutput {
	return o
}

func (o TagTemplateIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TagTemplateIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagTemplateIamPolicy {
		return vs[0].(map[string]TagTemplateIamPolicy)[vs[1].(string)]
	}).(TagTemplateIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(TagTemplateIamPolicyOutput{})
	pulumi.RegisterOutputType(TagTemplateIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(TagTemplateIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TagTemplateIamPolicyMapOutput{})
}
