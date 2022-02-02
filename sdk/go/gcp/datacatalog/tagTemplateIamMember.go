// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagTemplateIamMember struct {
	pulumi.CustomResourceState

	Condition   TagTemplateIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag        pulumi.StringOutput                    `pulumi:"etag"`
	Member      pulumi.StringOutput                    `pulumi:"member"`
	Project     pulumi.StringOutput                    `pulumi:"project"`
	Region      pulumi.StringOutput                    `pulumi:"region"`
	Role        pulumi.StringOutput                    `pulumi:"role"`
	TagTemplate pulumi.StringOutput                    `pulumi:"tagTemplate"`
}

// NewTagTemplateIamMember registers a new resource with the given unique name, arguments, and options.
func NewTagTemplateIamMember(ctx *pulumi.Context,
	name string, args *TagTemplateIamMemberArgs, opts ...pulumi.ResourceOption) (*TagTemplateIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagTemplate == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplate'")
	}
	var resource TagTemplateIamMember
	err := ctx.RegisterResource("gcp:datacatalog/tagTemplateIamMember:TagTemplateIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplateIamMember gets an existing TagTemplateIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplateIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateIamMemberState, opts ...pulumi.ResourceOption) (*TagTemplateIamMember, error) {
	var resource TagTemplateIamMember
	err := ctx.ReadResource("gcp:datacatalog/tagTemplateIamMember:TagTemplateIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplateIamMember resources.
type tagTemplateIamMemberState struct {
	Condition   *TagTemplateIamMemberCondition `pulumi:"condition"`
	Etag        *string                        `pulumi:"etag"`
	Member      *string                        `pulumi:"member"`
	Project     *string                        `pulumi:"project"`
	Region      *string                        `pulumi:"region"`
	Role        *string                        `pulumi:"role"`
	TagTemplate *string                        `pulumi:"tagTemplate"`
}

type TagTemplateIamMemberState struct {
	Condition   TagTemplateIamMemberConditionPtrInput
	Etag        pulumi.StringPtrInput
	Member      pulumi.StringPtrInput
	Project     pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	Role        pulumi.StringPtrInput
	TagTemplate pulumi.StringPtrInput
}

func (TagTemplateIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamMemberState)(nil)).Elem()
}

type tagTemplateIamMemberArgs struct {
	Condition   *TagTemplateIamMemberCondition `pulumi:"condition"`
	Member      string                         `pulumi:"member"`
	Project     *string                        `pulumi:"project"`
	Region      *string                        `pulumi:"region"`
	Role        string                         `pulumi:"role"`
	TagTemplate string                         `pulumi:"tagTemplate"`
}

// The set of arguments for constructing a TagTemplateIamMember resource.
type TagTemplateIamMemberArgs struct {
	Condition   TagTemplateIamMemberConditionPtrInput
	Member      pulumi.StringInput
	Project     pulumi.StringPtrInput
	Region      pulumi.StringPtrInput
	Role        pulumi.StringInput
	TagTemplate pulumi.StringInput
}

func (TagTemplateIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamMemberArgs)(nil)).Elem()
}

type TagTemplateIamMemberInput interface {
	pulumi.Input

	ToTagTemplateIamMemberOutput() TagTemplateIamMemberOutput
	ToTagTemplateIamMemberOutputWithContext(ctx context.Context) TagTemplateIamMemberOutput
}

func (*TagTemplateIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamMember)(nil)).Elem()
}

func (i *TagTemplateIamMember) ToTagTemplateIamMemberOutput() TagTemplateIamMemberOutput {
	return i.ToTagTemplateIamMemberOutputWithContext(context.Background())
}

func (i *TagTemplateIamMember) ToTagTemplateIamMemberOutputWithContext(ctx context.Context) TagTemplateIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamMemberOutput)
}

// TagTemplateIamMemberArrayInput is an input type that accepts TagTemplateIamMemberArray and TagTemplateIamMemberArrayOutput values.
// You can construct a concrete instance of `TagTemplateIamMemberArrayInput` via:
//
//          TagTemplateIamMemberArray{ TagTemplateIamMemberArgs{...} }
type TagTemplateIamMemberArrayInput interface {
	pulumi.Input

	ToTagTemplateIamMemberArrayOutput() TagTemplateIamMemberArrayOutput
	ToTagTemplateIamMemberArrayOutputWithContext(context.Context) TagTemplateIamMemberArrayOutput
}

type TagTemplateIamMemberArray []TagTemplateIamMemberInput

func (TagTemplateIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagTemplateIamMember)(nil)).Elem()
}

func (i TagTemplateIamMemberArray) ToTagTemplateIamMemberArrayOutput() TagTemplateIamMemberArrayOutput {
	return i.ToTagTemplateIamMemberArrayOutputWithContext(context.Background())
}

func (i TagTemplateIamMemberArray) ToTagTemplateIamMemberArrayOutputWithContext(ctx context.Context) TagTemplateIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamMemberArrayOutput)
}

// TagTemplateIamMemberMapInput is an input type that accepts TagTemplateIamMemberMap and TagTemplateIamMemberMapOutput values.
// You can construct a concrete instance of `TagTemplateIamMemberMapInput` via:
//
//          TagTemplateIamMemberMap{ "key": TagTemplateIamMemberArgs{...} }
type TagTemplateIamMemberMapInput interface {
	pulumi.Input

	ToTagTemplateIamMemberMapOutput() TagTemplateIamMemberMapOutput
	ToTagTemplateIamMemberMapOutputWithContext(context.Context) TagTemplateIamMemberMapOutput
}

type TagTemplateIamMemberMap map[string]TagTemplateIamMemberInput

func (TagTemplateIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagTemplateIamMember)(nil)).Elem()
}

func (i TagTemplateIamMemberMap) ToTagTemplateIamMemberMapOutput() TagTemplateIamMemberMapOutput {
	return i.ToTagTemplateIamMemberMapOutputWithContext(context.Background())
}

func (i TagTemplateIamMemberMap) ToTagTemplateIamMemberMapOutputWithContext(ctx context.Context) TagTemplateIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamMemberMapOutput)
}

type TagTemplateIamMemberOutput struct{ *pulumi.OutputState }

func (TagTemplateIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamMember)(nil)).Elem()
}

func (o TagTemplateIamMemberOutput) ToTagTemplateIamMemberOutput() TagTemplateIamMemberOutput {
	return o
}

func (o TagTemplateIamMemberOutput) ToTagTemplateIamMemberOutputWithContext(ctx context.Context) TagTemplateIamMemberOutput {
	return o
}

type TagTemplateIamMemberArrayOutput struct{ *pulumi.OutputState }

func (TagTemplateIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagTemplateIamMember)(nil)).Elem()
}

func (o TagTemplateIamMemberArrayOutput) ToTagTemplateIamMemberArrayOutput() TagTemplateIamMemberArrayOutput {
	return o
}

func (o TagTemplateIamMemberArrayOutput) ToTagTemplateIamMemberArrayOutputWithContext(ctx context.Context) TagTemplateIamMemberArrayOutput {
	return o
}

func (o TagTemplateIamMemberArrayOutput) Index(i pulumi.IntInput) TagTemplateIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagTemplateIamMember {
		return vs[0].([]*TagTemplateIamMember)[vs[1].(int)]
	}).(TagTemplateIamMemberOutput)
}

type TagTemplateIamMemberMapOutput struct{ *pulumi.OutputState }

func (TagTemplateIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagTemplateIamMember)(nil)).Elem()
}

func (o TagTemplateIamMemberMapOutput) ToTagTemplateIamMemberMapOutput() TagTemplateIamMemberMapOutput {
	return o
}

func (o TagTemplateIamMemberMapOutput) ToTagTemplateIamMemberMapOutputWithContext(ctx context.Context) TagTemplateIamMemberMapOutput {
	return o
}

func (o TagTemplateIamMemberMapOutput) MapIndex(k pulumi.StringInput) TagTemplateIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagTemplateIamMember {
		return vs[0].(map[string]*TagTemplateIamMember)[vs[1].(string)]
	}).(TagTemplateIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamMemberInput)(nil)).Elem(), &TagTemplateIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamMemberArrayInput)(nil)).Elem(), TagTemplateIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamMemberMapInput)(nil)).Elem(), TagTemplateIamMemberMap{})
	pulumi.RegisterOutputType(TagTemplateIamMemberOutput{})
	pulumi.RegisterOutputType(TagTemplateIamMemberArrayOutput{})
	pulumi.RegisterOutputType(TagTemplateIamMemberMapOutput{})
}
