// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Tags TagKey. Each of these resources serves a different use case:
//
// * `tags.TagKeyIamPolicy`: Authoritative. Sets the IAM policy for the tagkey and replaces any existing policy already attached.
// * `tags.TagKeyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagkey are preserved.
// * `tags.TagKeyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagkey are preserved.
//
// > **Note:** `tags.TagKeyIamPolicy` **cannot** be used in conjunction with `tags.TagKeyIamBinding` and `tags.TagKeyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `tags.TagKeyIamBinding` resources **can be** used in conjunction with `tags.TagKeyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_tags\_tag\_key\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = tags.NewTagKeyIamPolicy(ctx, "policy", &tags.TagKeyIamPolicyArgs{
//				TagKey:     pulumi.Any(google_tags_tag_key.Key.Name),
//				PolicyData: *pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_tags\_tag\_key\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tags.NewTagKeyIamBinding(ctx, "binding", &tags.TagKeyIamBindingArgs{
//				TagKey: pulumi.Any(google_tags_tag_key.Key.Name),
//				Role:   pulumi.String("roles/viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_tags\_tag\_key\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tags.NewTagKeyIamMember(ctx, "member", &tags.TagKeyIamMemberArgs{
//				TagKey: pulumi.Any(google_tags_tag_key.Key.Name),
//				Role:   pulumi.String("roles/viewer"),
//				Member: pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* tagKeys/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Tags tagkey IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:tags/tagKeyIamBinding:TagKeyIamBinding editor "tagKeys/{{tag_key}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:tags/tagKeyIamBinding:TagKeyIamBinding editor "tagKeys/{{tag_key}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:tags/tagKeyIamBinding:TagKeyIamBinding editor tagKeys/{{tag_key}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagKeyIamBinding struct {
	pulumi.CustomResourceState

	Condition TagKeyIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringOutput `pulumi:"tagKey"`
}

// NewTagKeyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTagKeyIamBinding(ctx *pulumi.Context,
	name string, args *TagKeyIamBindingArgs, opts ...pulumi.ResourceOption) (*TagKeyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagKey == nil {
		return nil, errors.New("invalid value for required argument 'TagKey'")
	}
	var resource TagKeyIamBinding
	err := ctx.RegisterResource("gcp:tags/tagKeyIamBinding:TagKeyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagKeyIamBinding gets an existing TagKeyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagKeyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagKeyIamBindingState, opts ...pulumi.ResourceOption) (*TagKeyIamBinding, error) {
	var resource TagKeyIamBinding
	err := ctx.ReadResource("gcp:tags/tagKeyIamBinding:TagKeyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagKeyIamBinding resources.
type tagKeyIamBindingState struct {
	Condition *TagKeyIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey *string `pulumi:"tagKey"`
}

type TagKeyIamBindingState struct {
	Condition TagKeyIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringPtrInput
}

func (TagKeyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamBindingState)(nil)).Elem()
}

type tagKeyIamBindingArgs struct {
	Condition *TagKeyIamBindingCondition `pulumi:"condition"`
	Members   []string                   `pulumi:"members"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey string `pulumi:"tagKey"`
}

// The set of arguments for constructing a TagKeyIamBinding resource.
type TagKeyIamBindingArgs struct {
	Condition TagKeyIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringInput
}

func (TagKeyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamBindingArgs)(nil)).Elem()
}

type TagKeyIamBindingInput interface {
	pulumi.Input

	ToTagKeyIamBindingOutput() TagKeyIamBindingOutput
	ToTagKeyIamBindingOutputWithContext(ctx context.Context) TagKeyIamBindingOutput
}

func (*TagKeyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamBinding)(nil)).Elem()
}

func (i *TagKeyIamBinding) ToTagKeyIamBindingOutput() TagKeyIamBindingOutput {
	return i.ToTagKeyIamBindingOutputWithContext(context.Background())
}

func (i *TagKeyIamBinding) ToTagKeyIamBindingOutputWithContext(ctx context.Context) TagKeyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingOutput)
}

// TagKeyIamBindingArrayInput is an input type that accepts TagKeyIamBindingArray and TagKeyIamBindingArrayOutput values.
// You can construct a concrete instance of `TagKeyIamBindingArrayInput` via:
//
//	TagKeyIamBindingArray{ TagKeyIamBindingArgs{...} }
type TagKeyIamBindingArrayInput interface {
	pulumi.Input

	ToTagKeyIamBindingArrayOutput() TagKeyIamBindingArrayOutput
	ToTagKeyIamBindingArrayOutputWithContext(context.Context) TagKeyIamBindingArrayOutput
}

type TagKeyIamBindingArray []TagKeyIamBindingInput

func (TagKeyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagKeyIamBinding)(nil)).Elem()
}

func (i TagKeyIamBindingArray) ToTagKeyIamBindingArrayOutput() TagKeyIamBindingArrayOutput {
	return i.ToTagKeyIamBindingArrayOutputWithContext(context.Background())
}

func (i TagKeyIamBindingArray) ToTagKeyIamBindingArrayOutputWithContext(ctx context.Context) TagKeyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingArrayOutput)
}

// TagKeyIamBindingMapInput is an input type that accepts TagKeyIamBindingMap and TagKeyIamBindingMapOutput values.
// You can construct a concrete instance of `TagKeyIamBindingMapInput` via:
//
//	TagKeyIamBindingMap{ "key": TagKeyIamBindingArgs{...} }
type TagKeyIamBindingMapInput interface {
	pulumi.Input

	ToTagKeyIamBindingMapOutput() TagKeyIamBindingMapOutput
	ToTagKeyIamBindingMapOutputWithContext(context.Context) TagKeyIamBindingMapOutput
}

type TagKeyIamBindingMap map[string]TagKeyIamBindingInput

func (TagKeyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagKeyIamBinding)(nil)).Elem()
}

func (i TagKeyIamBindingMap) ToTagKeyIamBindingMapOutput() TagKeyIamBindingMapOutput {
	return i.ToTagKeyIamBindingMapOutputWithContext(context.Background())
}

func (i TagKeyIamBindingMap) ToTagKeyIamBindingMapOutputWithContext(ctx context.Context) TagKeyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamBindingMapOutput)
}

type TagKeyIamBindingOutput struct{ *pulumi.OutputState }

func (TagKeyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamBinding)(nil)).Elem()
}

func (o TagKeyIamBindingOutput) ToTagKeyIamBindingOutput() TagKeyIamBindingOutput {
	return o
}

func (o TagKeyIamBindingOutput) ToTagKeyIamBindingOutputWithContext(ctx context.Context) TagKeyIamBindingOutput {
	return o
}

func (o TagKeyIamBindingOutput) Condition() TagKeyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *TagKeyIamBinding) TagKeyIamBindingConditionPtrOutput { return v.Condition }).(TagKeyIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o TagKeyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TagKeyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TagKeyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. Only one
// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o TagKeyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o TagKeyIamBindingOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamBinding) pulumi.StringOutput { return v.TagKey }).(pulumi.StringOutput)
}

type TagKeyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (TagKeyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagKeyIamBinding)(nil)).Elem()
}

func (o TagKeyIamBindingArrayOutput) ToTagKeyIamBindingArrayOutput() TagKeyIamBindingArrayOutput {
	return o
}

func (o TagKeyIamBindingArrayOutput) ToTagKeyIamBindingArrayOutputWithContext(ctx context.Context) TagKeyIamBindingArrayOutput {
	return o
}

func (o TagKeyIamBindingArrayOutput) Index(i pulumi.IntInput) TagKeyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagKeyIamBinding {
		return vs[0].([]*TagKeyIamBinding)[vs[1].(int)]
	}).(TagKeyIamBindingOutput)
}

type TagKeyIamBindingMapOutput struct{ *pulumi.OutputState }

func (TagKeyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagKeyIamBinding)(nil)).Elem()
}

func (o TagKeyIamBindingMapOutput) ToTagKeyIamBindingMapOutput() TagKeyIamBindingMapOutput {
	return o
}

func (o TagKeyIamBindingMapOutput) ToTagKeyIamBindingMapOutputWithContext(ctx context.Context) TagKeyIamBindingMapOutput {
	return o
}

func (o TagKeyIamBindingMapOutput) MapIndex(k pulumi.StringInput) TagKeyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagKeyIamBinding {
		return vs[0].(map[string]*TagKeyIamBinding)[vs[1].(string)]
	}).(TagKeyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamBindingInput)(nil)).Elem(), &TagKeyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamBindingArrayInput)(nil)).Elem(), TagKeyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamBindingMapInput)(nil)).Elem(), TagKeyIamBindingMap{})
	pulumi.RegisterOutputType(TagKeyIamBindingOutput{})
	pulumi.RegisterOutputType(TagKeyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(TagKeyIamBindingMapOutput{})
}
