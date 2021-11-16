// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = tags.NewTagKeyIamPolicy(ctx, "policy", &tags.TagKeyIamPolicyArgs{
// 			TagKey:     pulumi.Any(google_tags_tag_key.Key.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_tags\_tag\_key\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tags.NewTagKeyIamBinding(ctx, "binding", &tags.TagKeyIamBindingArgs{
// 			TagKey: pulumi.Any(google_tags_tag_key.Key.Name),
// 			Role:   pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_tags\_tag\_key\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/tags"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tags.NewTagKeyIamMember(ctx, "member", &tags.TagKeyIamMemberArgs{
// 			TagKey: pulumi.Any(google_tags_tag_key.Key.Name),
// 			Role:   pulumi.String("roles/viewer"),
// 			Member: pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* tagKeys/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Tags tagkey IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagKeyIamMember:TagKeyIamMember editor "tagKeys/{{tag_key}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagKeyIamMember:TagKeyIamMember editor "tagKeys/{{tag_key}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagKeyIamMember:TagKeyIamMember editor tagKeys/{{tag_key}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagKeyIamMember struct {
	pulumi.CustomResourceState

	Condition TagKeyIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringOutput `pulumi:"tagKey"`
}

// NewTagKeyIamMember registers a new resource with the given unique name, arguments, and options.
func NewTagKeyIamMember(ctx *pulumi.Context,
	name string, args *TagKeyIamMemberArgs, opts ...pulumi.ResourceOption) (*TagKeyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagKey == nil {
		return nil, errors.New("invalid value for required argument 'TagKey'")
	}
	var resource TagKeyIamMember
	err := ctx.RegisterResource("gcp:tags/tagKeyIamMember:TagKeyIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagKeyIamMember gets an existing TagKeyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagKeyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagKeyIamMemberState, opts ...pulumi.ResourceOption) (*TagKeyIamMember, error) {
	var resource TagKeyIamMember
	err := ctx.ReadResource("gcp:tags/tagKeyIamMember:TagKeyIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagKeyIamMember resources.
type tagKeyIamMemberState struct {
	Condition *TagKeyIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey *string `pulumi:"tagKey"`
}

type TagKeyIamMemberState struct {
	Condition TagKeyIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringPtrInput
}

func (TagKeyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamMemberState)(nil)).Elem()
}

type tagKeyIamMemberArgs struct {
	Condition *TagKeyIamMemberCondition `pulumi:"condition"`
	Member    string                    `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagKey string `pulumi:"tagKey"`
}

// The set of arguments for constructing a TagKeyIamMember resource.
type TagKeyIamMemberArgs struct {
	Condition TagKeyIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `tags.TagKeyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	TagKey pulumi.StringInput
}

func (TagKeyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamMemberArgs)(nil)).Elem()
}

type TagKeyIamMemberInput interface {
	pulumi.Input

	ToTagKeyIamMemberOutput() TagKeyIamMemberOutput
	ToTagKeyIamMemberOutputWithContext(ctx context.Context) TagKeyIamMemberOutput
}

func (*TagKeyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamMember)(nil))
}

func (i *TagKeyIamMember) ToTagKeyIamMemberOutput() TagKeyIamMemberOutput {
	return i.ToTagKeyIamMemberOutputWithContext(context.Background())
}

func (i *TagKeyIamMember) ToTagKeyIamMemberOutputWithContext(ctx context.Context) TagKeyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberOutput)
}

func (i *TagKeyIamMember) ToTagKeyIamMemberPtrOutput() TagKeyIamMemberPtrOutput {
	return i.ToTagKeyIamMemberPtrOutputWithContext(context.Background())
}

func (i *TagKeyIamMember) ToTagKeyIamMemberPtrOutputWithContext(ctx context.Context) TagKeyIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberPtrOutput)
}

type TagKeyIamMemberPtrInput interface {
	pulumi.Input

	ToTagKeyIamMemberPtrOutput() TagKeyIamMemberPtrOutput
	ToTagKeyIamMemberPtrOutputWithContext(ctx context.Context) TagKeyIamMemberPtrOutput
}

type tagKeyIamMemberPtrType TagKeyIamMemberArgs

func (*tagKeyIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamMember)(nil))
}

func (i *tagKeyIamMemberPtrType) ToTagKeyIamMemberPtrOutput() TagKeyIamMemberPtrOutput {
	return i.ToTagKeyIamMemberPtrOutputWithContext(context.Background())
}

func (i *tagKeyIamMemberPtrType) ToTagKeyIamMemberPtrOutputWithContext(ctx context.Context) TagKeyIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberPtrOutput)
}

// TagKeyIamMemberArrayInput is an input type that accepts TagKeyIamMemberArray and TagKeyIamMemberArrayOutput values.
// You can construct a concrete instance of `TagKeyIamMemberArrayInput` via:
//
//          TagKeyIamMemberArray{ TagKeyIamMemberArgs{...} }
type TagKeyIamMemberArrayInput interface {
	pulumi.Input

	ToTagKeyIamMemberArrayOutput() TagKeyIamMemberArrayOutput
	ToTagKeyIamMemberArrayOutputWithContext(context.Context) TagKeyIamMemberArrayOutput
}

type TagKeyIamMemberArray []TagKeyIamMemberInput

func (TagKeyIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagKeyIamMember)(nil)).Elem()
}

func (i TagKeyIamMemberArray) ToTagKeyIamMemberArrayOutput() TagKeyIamMemberArrayOutput {
	return i.ToTagKeyIamMemberArrayOutputWithContext(context.Background())
}

func (i TagKeyIamMemberArray) ToTagKeyIamMemberArrayOutputWithContext(ctx context.Context) TagKeyIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberArrayOutput)
}

// TagKeyIamMemberMapInput is an input type that accepts TagKeyIamMemberMap and TagKeyIamMemberMapOutput values.
// You can construct a concrete instance of `TagKeyIamMemberMapInput` via:
//
//          TagKeyIamMemberMap{ "key": TagKeyIamMemberArgs{...} }
type TagKeyIamMemberMapInput interface {
	pulumi.Input

	ToTagKeyIamMemberMapOutput() TagKeyIamMemberMapOutput
	ToTagKeyIamMemberMapOutputWithContext(context.Context) TagKeyIamMemberMapOutput
}

type TagKeyIamMemberMap map[string]TagKeyIamMemberInput

func (TagKeyIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagKeyIamMember)(nil)).Elem()
}

func (i TagKeyIamMemberMap) ToTagKeyIamMemberMapOutput() TagKeyIamMemberMapOutput {
	return i.ToTagKeyIamMemberMapOutputWithContext(context.Background())
}

func (i TagKeyIamMemberMap) ToTagKeyIamMemberMapOutputWithContext(ctx context.Context) TagKeyIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamMemberMapOutput)
}

type TagKeyIamMemberOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagKeyIamMember)(nil))
}

func (o TagKeyIamMemberOutput) ToTagKeyIamMemberOutput() TagKeyIamMemberOutput {
	return o
}

func (o TagKeyIamMemberOutput) ToTagKeyIamMemberOutputWithContext(ctx context.Context) TagKeyIamMemberOutput {
	return o
}

func (o TagKeyIamMemberOutput) ToTagKeyIamMemberPtrOutput() TagKeyIamMemberPtrOutput {
	return o.ToTagKeyIamMemberPtrOutputWithContext(context.Background())
}

func (o TagKeyIamMemberOutput) ToTagKeyIamMemberPtrOutputWithContext(ctx context.Context) TagKeyIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagKeyIamMember) *TagKeyIamMember {
		return &v
	}).(TagKeyIamMemberPtrOutput)
}

type TagKeyIamMemberPtrOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamMember)(nil))
}

func (o TagKeyIamMemberPtrOutput) ToTagKeyIamMemberPtrOutput() TagKeyIamMemberPtrOutput {
	return o
}

func (o TagKeyIamMemberPtrOutput) ToTagKeyIamMemberPtrOutputWithContext(ctx context.Context) TagKeyIamMemberPtrOutput {
	return o
}

func (o TagKeyIamMemberPtrOutput) Elem() TagKeyIamMemberOutput {
	return o.ApplyT(func(v *TagKeyIamMember) TagKeyIamMember {
		if v != nil {
			return *v
		}
		var ret TagKeyIamMember
		return ret
	}).(TagKeyIamMemberOutput)
}

type TagKeyIamMemberArrayOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagKeyIamMember)(nil))
}

func (o TagKeyIamMemberArrayOutput) ToTagKeyIamMemberArrayOutput() TagKeyIamMemberArrayOutput {
	return o
}

func (o TagKeyIamMemberArrayOutput) ToTagKeyIamMemberArrayOutputWithContext(ctx context.Context) TagKeyIamMemberArrayOutput {
	return o
}

func (o TagKeyIamMemberArrayOutput) Index(i pulumi.IntInput) TagKeyIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagKeyIamMember {
		return vs[0].([]TagKeyIamMember)[vs[1].(int)]
	}).(TagKeyIamMemberOutput)
}

type TagKeyIamMemberMapOutput struct{ *pulumi.OutputState }

func (TagKeyIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagKeyIamMember)(nil))
}

func (o TagKeyIamMemberMapOutput) ToTagKeyIamMemberMapOutput() TagKeyIamMemberMapOutput {
	return o
}

func (o TagKeyIamMemberMapOutput) ToTagKeyIamMemberMapOutputWithContext(ctx context.Context) TagKeyIamMemberMapOutput {
	return o
}

func (o TagKeyIamMemberMapOutput) MapIndex(k pulumi.StringInput) TagKeyIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagKeyIamMember {
		return vs[0].(map[string]TagKeyIamMember)[vs[1].(string)]
	}).(TagKeyIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(TagKeyIamMemberOutput{})
	pulumi.RegisterOutputType(TagKeyIamMemberPtrOutput{})
	pulumi.RegisterOutputType(TagKeyIamMemberArrayOutput{})
	pulumi.RegisterOutputType(TagKeyIamMemberMapOutput{})
}
