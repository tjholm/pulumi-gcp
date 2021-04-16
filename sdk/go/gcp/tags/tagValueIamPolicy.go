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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/tags"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/tags"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/tags"
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
//  $ pulumi import gcp:tags/tagValueIamPolicy:TagValueIamPolicy editor "tagValues/{{tag_value}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagValueIamPolicy:TagValueIamPolicy editor "tagValues/{{tag_value}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:tags/tagValueIamPolicy:TagValueIamPolicy editor tagValues/{{tag_value}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagValueIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringOutput `pulumi:"tagValue"`
}

// NewTagValueIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTagValueIamPolicy(ctx *pulumi.Context,
	name string, args *TagValueIamPolicyArgs, opts ...pulumi.ResourceOption) (*TagValueIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.TagValue == nil {
		return nil, errors.New("invalid value for required argument 'TagValue'")
	}
	var resource TagValueIamPolicy
	err := ctx.RegisterResource("gcp:tags/tagValueIamPolicy:TagValueIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagValueIamPolicy gets an existing TagValueIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagValueIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagValueIamPolicyState, opts ...pulumi.ResourceOption) (*TagValueIamPolicy, error) {
	var resource TagValueIamPolicy
	err := ctx.ReadResource("gcp:tags/tagValueIamPolicy:TagValueIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagValueIamPolicy resources.
type tagValueIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue *string `pulumi:"tagValue"`
}

type TagValueIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringPtrInput
}

func (TagValueIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueIamPolicyState)(nil)).Elem()
}

type tagValueIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	TagValue string `pulumi:"tagValue"`
}

// The set of arguments for constructing a TagValueIamPolicy resource.
type TagValueIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	TagValue pulumi.StringInput
}

func (TagValueIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagValueIamPolicyArgs)(nil)).Elem()
}

type TagValueIamPolicyInput interface {
	pulumi.Input

	ToTagValueIamPolicyOutput() TagValueIamPolicyOutput
	ToTagValueIamPolicyOutputWithContext(ctx context.Context) TagValueIamPolicyOutput
}

func (*TagValueIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamPolicy)(nil))
}

func (i *TagValueIamPolicy) ToTagValueIamPolicyOutput() TagValueIamPolicyOutput {
	return i.ToTagValueIamPolicyOutputWithContext(context.Background())
}

func (i *TagValueIamPolicy) ToTagValueIamPolicyOutputWithContext(ctx context.Context) TagValueIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamPolicyOutput)
}

func (i *TagValueIamPolicy) ToTagValueIamPolicyPtrOutput() TagValueIamPolicyPtrOutput {
	return i.ToTagValueIamPolicyPtrOutputWithContext(context.Background())
}

func (i *TagValueIamPolicy) ToTagValueIamPolicyPtrOutputWithContext(ctx context.Context) TagValueIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamPolicyPtrOutput)
}

type TagValueIamPolicyPtrInput interface {
	pulumi.Input

	ToTagValueIamPolicyPtrOutput() TagValueIamPolicyPtrOutput
	ToTagValueIamPolicyPtrOutputWithContext(ctx context.Context) TagValueIamPolicyPtrOutput
}

type tagValueIamPolicyPtrType TagValueIamPolicyArgs

func (*tagValueIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamPolicy)(nil))
}

func (i *tagValueIamPolicyPtrType) ToTagValueIamPolicyPtrOutput() TagValueIamPolicyPtrOutput {
	return i.ToTagValueIamPolicyPtrOutputWithContext(context.Background())
}

func (i *tagValueIamPolicyPtrType) ToTagValueIamPolicyPtrOutputWithContext(ctx context.Context) TagValueIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamPolicyPtrOutput)
}

// TagValueIamPolicyArrayInput is an input type that accepts TagValueIamPolicyArray and TagValueIamPolicyArrayOutput values.
// You can construct a concrete instance of `TagValueIamPolicyArrayInput` via:
//
//          TagValueIamPolicyArray{ TagValueIamPolicyArgs{...} }
type TagValueIamPolicyArrayInput interface {
	pulumi.Input

	ToTagValueIamPolicyArrayOutput() TagValueIamPolicyArrayOutput
	ToTagValueIamPolicyArrayOutputWithContext(context.Context) TagValueIamPolicyArrayOutput
}

type TagValueIamPolicyArray []TagValueIamPolicyInput

func (TagValueIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TagValueIamPolicy)(nil))
}

func (i TagValueIamPolicyArray) ToTagValueIamPolicyArrayOutput() TagValueIamPolicyArrayOutput {
	return i.ToTagValueIamPolicyArrayOutputWithContext(context.Background())
}

func (i TagValueIamPolicyArray) ToTagValueIamPolicyArrayOutputWithContext(ctx context.Context) TagValueIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamPolicyArrayOutput)
}

// TagValueIamPolicyMapInput is an input type that accepts TagValueIamPolicyMap and TagValueIamPolicyMapOutput values.
// You can construct a concrete instance of `TagValueIamPolicyMapInput` via:
//
//          TagValueIamPolicyMap{ "key": TagValueIamPolicyArgs{...} }
type TagValueIamPolicyMapInput interface {
	pulumi.Input

	ToTagValueIamPolicyMapOutput() TagValueIamPolicyMapOutput
	ToTagValueIamPolicyMapOutputWithContext(context.Context) TagValueIamPolicyMapOutput
}

type TagValueIamPolicyMap map[string]TagValueIamPolicyInput

func (TagValueIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TagValueIamPolicy)(nil))
}

func (i TagValueIamPolicyMap) ToTagValueIamPolicyMapOutput() TagValueIamPolicyMapOutput {
	return i.ToTagValueIamPolicyMapOutputWithContext(context.Background())
}

func (i TagValueIamPolicyMap) ToTagValueIamPolicyMapOutputWithContext(ctx context.Context) TagValueIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagValueIamPolicyMapOutput)
}

type TagValueIamPolicyOutput struct {
	*pulumi.OutputState
}

func (TagValueIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagValueIamPolicy)(nil))
}

func (o TagValueIamPolicyOutput) ToTagValueIamPolicyOutput() TagValueIamPolicyOutput {
	return o
}

func (o TagValueIamPolicyOutput) ToTagValueIamPolicyOutputWithContext(ctx context.Context) TagValueIamPolicyOutput {
	return o
}

func (o TagValueIamPolicyOutput) ToTagValueIamPolicyPtrOutput() TagValueIamPolicyPtrOutput {
	return o.ToTagValueIamPolicyPtrOutputWithContext(context.Background())
}

func (o TagValueIamPolicyOutput) ToTagValueIamPolicyPtrOutputWithContext(ctx context.Context) TagValueIamPolicyPtrOutput {
	return o.ApplyT(func(v TagValueIamPolicy) *TagValueIamPolicy {
		return &v
	}).(TagValueIamPolicyPtrOutput)
}

type TagValueIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (TagValueIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagValueIamPolicy)(nil))
}

func (o TagValueIamPolicyPtrOutput) ToTagValueIamPolicyPtrOutput() TagValueIamPolicyPtrOutput {
	return o
}

func (o TagValueIamPolicyPtrOutput) ToTagValueIamPolicyPtrOutputWithContext(ctx context.Context) TagValueIamPolicyPtrOutput {
	return o
}

type TagValueIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TagValueIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagValueIamPolicy)(nil))
}

func (o TagValueIamPolicyArrayOutput) ToTagValueIamPolicyArrayOutput() TagValueIamPolicyArrayOutput {
	return o
}

func (o TagValueIamPolicyArrayOutput) ToTagValueIamPolicyArrayOutputWithContext(ctx context.Context) TagValueIamPolicyArrayOutput {
	return o
}

func (o TagValueIamPolicyArrayOutput) Index(i pulumi.IntInput) TagValueIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagValueIamPolicy {
		return vs[0].([]TagValueIamPolicy)[vs[1].(int)]
	}).(TagValueIamPolicyOutput)
}

type TagValueIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TagValueIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagValueIamPolicy)(nil))
}

func (o TagValueIamPolicyMapOutput) ToTagValueIamPolicyMapOutput() TagValueIamPolicyMapOutput {
	return o
}

func (o TagValueIamPolicyMapOutput) ToTagValueIamPolicyMapOutputWithContext(ctx context.Context) TagValueIamPolicyMapOutput {
	return o
}

func (o TagValueIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TagValueIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagValueIamPolicy {
		return vs[0].(map[string]TagValueIamPolicy)[vs[1].(string)]
	}).(TagValueIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(TagValueIamPolicyOutput{})
	pulumi.RegisterOutputType(TagValueIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(TagValueIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TagValueIamPolicyMapOutput{})
}
