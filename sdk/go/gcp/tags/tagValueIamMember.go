// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Tags TagValue. Each of these resources serves a different use case:
//
// * `tags.TagValueIamPolicy`: Authoritative. Sets the IAM policy for the tagvalue and replaces any existing policy already attached.
// * `tags.TagValueIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagvalue are preserved.
// * `tags.TagValueIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagvalue are preserved.
//
// > **Note:** `tags.TagValueIamPolicy` **cannot** be used in conjunction with `tags.TagValueIamBinding` and `tags.TagValueIamMember` or they will fight over what your policy should be.
//
// > **Note:** `tags.TagValueIamBinding` resources **can be** used in conjunction with `tags.TagValueIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_tags\_tag\_value\_iam\_policy
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
// 		_, err = tags.NewTagValueIamPolicy(ctx, "policy", &tags.TagValueIamPolicyArgs{
// 			TagValue:   pulumi.Any(google_tags_tag_value.Value.Name),
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
// ## google\_tags\_tag\_value\_iam\_binding
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
// 		_, err := tags.NewTagValueIamBinding(ctx, "binding", &tags.TagValueIamBindingArgs{
// 			TagValue: pulumi.Any(google_tags_tag_value.Value.Name),
// 			Role:     pulumi.String("roles/viewer"),
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
// ## google\_tags\_tag\_value\_iam\_member
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
// 		_, err := tags.NewTagValueIamMember(ctx, "member", &tags.TagValueIamMemberArgs{
// 			TagValue: pulumi.Any(google_tags_tag_value.Value.Name),
// 			Role:     pulumi.String("roles/viewer"),
// 			Member:   pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* tagValues/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Tags tagvalue IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagValueIamMember:TagValueIamMember editor "tagValues/{{tag_value}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagValueIamMember:TagValueIamMember editor "tagValues/{{tag_value}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagValueIamMember:TagValueIamMember editor tagValues/{{tag_value}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagValueIamMember struct {
	pulumi.CustomResourceState

	Condition TagValueIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringOutput `pulumi:"tagValue"`
}

// NewTagValueIamMember registers a new resource with the given unique name, arguments, and options.
func NewTagValueIamMember(ctx *pulumi.Context,
	name string, args *TagValueIamMemberArgs, opts ...pulumi.ResourceOption) (*TagValueIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagValue == nil {
		return nil, errors.New("invalid value for required argument 'TagValue'")
	}
	var resource TagValueIamMember
	err := ctx.RegisterResource("gcp:tags/tagValueIamMember:TagValueIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagValueIamMember gets an existing TagValueIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagValueIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagValueIamMemberState, opts ...pulumi.ResourceOption) (*TagValueIamMember, error) {
	var resource TagValueIamMember
	err := ctx.ReadResource("gcp:tags/tagValueIamMember:TagValueIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagValueIamMember resources.
type tagValueIamMemberState struct {
	Condition *TagValueIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue *string `pulumi:"tagValue"`
}

type TagValueIamMemberState struct {
	Condition TagValueIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringPtrInput
}

func (TagValueIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueIamMemberState)(nil)).Elem()
}

type tagValueIamMemberArgs struct {
	Condition *TagValueIamMemberCondition `pulumi:"condition"`
	Member    string                      `pulumi:"member"`
	// The role that should be applied. Only one
	// `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue string `pulumi:"tagValue"`
}

// The set of arguments for constructing a TagValueIamMember resource.
type TagValueIamMemberArgs struct {
	Condition TagValueIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringInput
}

func (TagValueIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueIamMemberArgs)(nil)).Elem()
}

type TagValueIamMemberInput interface {
	pulumi.Input

	ToTagValueIamMemberOutput() TagValueIamMemberOutput
	ToTagValueIamMemberOutputWithContext(ctx context.Context) TagValueIamMemberOutput
}

func (*TagValueIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamMember)(nil))
}

func (i *TagValueIamMember) ToTagValueIamMemberOutput() TagValueIamMemberOutput {
	return i.ToTagValueIamMemberOutputWithContext(context.Background())
}

func (i *TagValueIamMember) ToTagValueIamMemberOutputWithContext(ctx context.Context) TagValueIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberOutput)
}

func (i *TagValueIamMember) ToTagValueIamMemberPtrOutput() TagValueIamMemberPtrOutput {
	return i.ToTagValueIamMemberPtrOutputWithContext(context.Background())
}

func (i *TagValueIamMember) ToTagValueIamMemberPtrOutputWithContext(ctx context.Context) TagValueIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberPtrOutput)
}

type TagValueIamMemberPtrInput interface {
	pulumi.Input

	ToTagValueIamMemberPtrOutput() TagValueIamMemberPtrOutput
	ToTagValueIamMemberPtrOutputWithContext(ctx context.Context) TagValueIamMemberPtrOutput
}

type tagValueIamMemberPtrType TagValueIamMemberArgs

func (*tagValueIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamMember)(nil))
}

func (i *tagValueIamMemberPtrType) ToTagValueIamMemberPtrOutput() TagValueIamMemberPtrOutput {
	return i.ToTagValueIamMemberPtrOutputWithContext(context.Background())
}

func (i *tagValueIamMemberPtrType) ToTagValueIamMemberPtrOutputWithContext(ctx context.Context) TagValueIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberPtrOutput)
}

// TagValueIamMemberArrayInput is an input type that accepts TagValueIamMemberArray and TagValueIamMemberArrayOutput values.
// You can construct a concrete instance of `TagValueIamMemberArrayInput` via:
//
//          TagValueIamMemberArray{ TagValueIamMemberArgs{...} }
type TagValueIamMemberArrayInput interface {
	pulumi.Input

	ToTagValueIamMemberArrayOutput() TagValueIamMemberArrayOutput
	ToTagValueIamMemberArrayOutputWithContext(context.Context) TagValueIamMemberArrayOutput
}

type TagValueIamMemberArray []TagValueIamMemberInput

func (TagValueIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagValueIamMember)(nil)).Elem()
}

func (i TagValueIamMemberArray) ToTagValueIamMemberArrayOutput() TagValueIamMemberArrayOutput {
	return i.ToTagValueIamMemberArrayOutputWithContext(context.Background())
}

func (i TagValueIamMemberArray) ToTagValueIamMemberArrayOutputWithContext(ctx context.Context) TagValueIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberArrayOutput)
}

// TagValueIamMemberMapInput is an input type that accepts TagValueIamMemberMap and TagValueIamMemberMapOutput values.
// You can construct a concrete instance of `TagValueIamMemberMapInput` via:
//
//          TagValueIamMemberMap{ "key": TagValueIamMemberArgs{...} }
type TagValueIamMemberMapInput interface {
	pulumi.Input

	ToTagValueIamMemberMapOutput() TagValueIamMemberMapOutput
	ToTagValueIamMemberMapOutputWithContext(context.Context) TagValueIamMemberMapOutput
}

type TagValueIamMemberMap map[string]TagValueIamMemberInput

func (TagValueIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagValueIamMember)(nil)).Elem()
}

func (i TagValueIamMemberMap) ToTagValueIamMemberMapOutput() TagValueIamMemberMapOutput {
	return i.ToTagValueIamMemberMapOutputWithContext(context.Background())
}

func (i TagValueIamMemberMap) ToTagValueIamMemberMapOutputWithContext(ctx context.Context) TagValueIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamMemberMapOutput)
}

type TagValueIamMemberOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamMember)(nil))
}

func (o TagValueIamMemberOutput) ToTagValueIamMemberOutput() TagValueIamMemberOutput {
	return o
}

func (o TagValueIamMemberOutput) ToTagValueIamMemberOutputWithContext(ctx context.Context) TagValueIamMemberOutput {
	return o
}

func (o TagValueIamMemberOutput) ToTagValueIamMemberPtrOutput() TagValueIamMemberPtrOutput {
	return o.ToTagValueIamMemberPtrOutputWithContext(context.Background())
}

func (o TagValueIamMemberOutput) ToTagValueIamMemberPtrOutputWithContext(ctx context.Context) TagValueIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagValueIamMember) *TagValueIamMember {
		return &v
	}).(TagValueIamMemberPtrOutput)
}

type TagValueIamMemberPtrOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamMember)(nil))
}

func (o TagValueIamMemberPtrOutput) ToTagValueIamMemberPtrOutput() TagValueIamMemberPtrOutput {
	return o
}

func (o TagValueIamMemberPtrOutput) ToTagValueIamMemberPtrOutputWithContext(ctx context.Context) TagValueIamMemberPtrOutput {
	return o
}

func (o TagValueIamMemberPtrOutput) Elem() TagValueIamMemberOutput {
	return o.ApplyT(func(v *TagValueIamMember) TagValueIamMember {
		if v != nil {
			return *v
		}
		var ret TagValueIamMember
		return ret
	}).(TagValueIamMemberOutput)
}

type TagValueIamMemberArrayOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagValueIamMember)(nil))
}

func (o TagValueIamMemberArrayOutput) ToTagValueIamMemberArrayOutput() TagValueIamMemberArrayOutput {
	return o
}

func (o TagValueIamMemberArrayOutput) ToTagValueIamMemberArrayOutputWithContext(ctx context.Context) TagValueIamMemberArrayOutput {
	return o
}

func (o TagValueIamMemberArrayOutput) Index(i pulumi.IntInput) TagValueIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagValueIamMember {
		return vs[0].([]TagValueIamMember)[vs[1].(int)]
	}).(TagValueIamMemberOutput)
}

type TagValueIamMemberMapOutput struct{ *pulumi.OutputState }

func (TagValueIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagValueIamMember)(nil))
}

func (o TagValueIamMemberMapOutput) ToTagValueIamMemberMapOutput() TagValueIamMemberMapOutput {
	return o
}

func (o TagValueIamMemberMapOutput) ToTagValueIamMemberMapOutputWithContext(ctx context.Context) TagValueIamMemberMapOutput {
	return o
}

func (o TagValueIamMemberMapOutput) MapIndex(k pulumi.StringInput) TagValueIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagValueIamMember {
		return vs[0].(map[string]TagValueIamMember)[vs[1].(string)]
	}).(TagValueIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(TagValueIamMemberOutput{})
	pulumi.RegisterOutputType(TagValueIamMemberPtrOutput{})
	pulumi.RegisterOutputType(TagValueIamMemberArrayOutput{})
	pulumi.RegisterOutputType(TagValueIamMemberMapOutput{})
}
