// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tags

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Three different resources help you manage your IAM policy for Tags TagKey. Each of these resources serves a different use case:
//
// * `tags.TagKeyIamPolicy`: Authoritative. Sets the IAM policy for the tagkey and replaces any existing policy already attached.
// * `tags.TagKeyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagkey are preserved.
// * `tags.TagKeyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagkey are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `tags.TagKeyIamPolicy`: Retrieves the IAM policy for the tagkey
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
//	$ pulumi import gcp:tags/tagKeyIamPolicy:TagKeyIamPolicy editor "tagKeys/{{tag_key}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:tags/tagKeyIamPolicy:TagKeyIamPolicy editor "tagKeys/{{tag_key}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:tags/tagKeyIamPolicy:TagKeyIamPolicy editor tagKeys/{{tag_key}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagKeyIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	TagKey pulumi.StringOutput `pulumi:"tagKey"`
}

// NewTagKeyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTagKeyIamPolicy(ctx *pulumi.Context,
	name string, args *TagKeyIamPolicyArgs, opts ...pulumi.ResourceOption) (*TagKeyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.TagKey == nil {
		return nil, errors.New("invalid value for required argument 'TagKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TagKeyIamPolicy
	err := ctx.RegisterResource("gcp:tags/tagKeyIamPolicy:TagKeyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagKeyIamPolicy gets an existing TagKeyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagKeyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagKeyIamPolicyState, opts ...pulumi.ResourceOption) (*TagKeyIamPolicy, error) {
	var resource TagKeyIamPolicy
	err := ctx.ReadResource("gcp:tags/tagKeyIamPolicy:TagKeyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagKeyIamPolicy resources.
type tagKeyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	TagKey *string `pulumi:"tagKey"`
}

type TagKeyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	TagKey pulumi.StringPtrInput
}

func (TagKeyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamPolicyState)(nil)).Elem()
}

type tagKeyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// Used to find the parent resource to bind the IAM policy to
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	TagKey string `pulumi:"tagKey"`
}

// The set of arguments for constructing a TagKeyIamPolicy resource.
type TagKeyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
	// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
	// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
	TagKey pulumi.StringInput
}

func (TagKeyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagKeyIamPolicyArgs)(nil)).Elem()
}

type TagKeyIamPolicyInput interface {
	pulumi.Input

	ToTagKeyIamPolicyOutput() TagKeyIamPolicyOutput
	ToTagKeyIamPolicyOutputWithContext(ctx context.Context) TagKeyIamPolicyOutput
}

func (*TagKeyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamPolicy)(nil)).Elem()
}

func (i *TagKeyIamPolicy) ToTagKeyIamPolicyOutput() TagKeyIamPolicyOutput {
	return i.ToTagKeyIamPolicyOutputWithContext(context.Background())
}

func (i *TagKeyIamPolicy) ToTagKeyIamPolicyOutputWithContext(ctx context.Context) TagKeyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamPolicyOutput)
}

func (i *TagKeyIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*TagKeyIamPolicy] {
	return pulumix.Output[*TagKeyIamPolicy]{
		OutputState: i.ToTagKeyIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// TagKeyIamPolicyArrayInput is an input type that accepts TagKeyIamPolicyArray and TagKeyIamPolicyArrayOutput values.
// You can construct a concrete instance of `TagKeyIamPolicyArrayInput` via:
//
//	TagKeyIamPolicyArray{ TagKeyIamPolicyArgs{...} }
type TagKeyIamPolicyArrayInput interface {
	pulumi.Input

	ToTagKeyIamPolicyArrayOutput() TagKeyIamPolicyArrayOutput
	ToTagKeyIamPolicyArrayOutputWithContext(context.Context) TagKeyIamPolicyArrayOutput
}

type TagKeyIamPolicyArray []TagKeyIamPolicyInput

func (TagKeyIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagKeyIamPolicy)(nil)).Elem()
}

func (i TagKeyIamPolicyArray) ToTagKeyIamPolicyArrayOutput() TagKeyIamPolicyArrayOutput {
	return i.ToTagKeyIamPolicyArrayOutputWithContext(context.Background())
}

func (i TagKeyIamPolicyArray) ToTagKeyIamPolicyArrayOutputWithContext(ctx context.Context) TagKeyIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamPolicyArrayOutput)
}

func (i TagKeyIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*TagKeyIamPolicy] {
	return pulumix.Output[[]*TagKeyIamPolicy]{
		OutputState: i.ToTagKeyIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// TagKeyIamPolicyMapInput is an input type that accepts TagKeyIamPolicyMap and TagKeyIamPolicyMapOutput values.
// You can construct a concrete instance of `TagKeyIamPolicyMapInput` via:
//
//	TagKeyIamPolicyMap{ "key": TagKeyIamPolicyArgs{...} }
type TagKeyIamPolicyMapInput interface {
	pulumi.Input

	ToTagKeyIamPolicyMapOutput() TagKeyIamPolicyMapOutput
	ToTagKeyIamPolicyMapOutputWithContext(context.Context) TagKeyIamPolicyMapOutput
}

type TagKeyIamPolicyMap map[string]TagKeyIamPolicyInput

func (TagKeyIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagKeyIamPolicy)(nil)).Elem()
}

func (i TagKeyIamPolicyMap) ToTagKeyIamPolicyMapOutput() TagKeyIamPolicyMapOutput {
	return i.ToTagKeyIamPolicyMapOutputWithContext(context.Background())
}

func (i TagKeyIamPolicyMap) ToTagKeyIamPolicyMapOutputWithContext(ctx context.Context) TagKeyIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagKeyIamPolicyMapOutput)
}

func (i TagKeyIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TagKeyIamPolicy] {
	return pulumix.Output[map[string]*TagKeyIamPolicy]{
		OutputState: i.ToTagKeyIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type TagKeyIamPolicyOutput struct{ *pulumi.OutputState }

func (TagKeyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagKeyIamPolicy)(nil)).Elem()
}

func (o TagKeyIamPolicyOutput) ToTagKeyIamPolicyOutput() TagKeyIamPolicyOutput {
	return o
}

func (o TagKeyIamPolicyOutput) ToTagKeyIamPolicyOutputWithContext(ctx context.Context) TagKeyIamPolicyOutput {
	return o
}

func (o TagKeyIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*TagKeyIamPolicy] {
	return pulumix.Output[*TagKeyIamPolicy]{
		OutputState: o.OutputState,
	}
}

// (Computed) The etag of the IAM policy.
func (o TagKeyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o TagKeyIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
//   - **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
//   - **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
//   - **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
func (o TagKeyIamPolicyOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v *TagKeyIamPolicy) pulumi.StringOutput { return v.TagKey }).(pulumi.StringOutput)
}

type TagKeyIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TagKeyIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagKeyIamPolicy)(nil)).Elem()
}

func (o TagKeyIamPolicyArrayOutput) ToTagKeyIamPolicyArrayOutput() TagKeyIamPolicyArrayOutput {
	return o
}

func (o TagKeyIamPolicyArrayOutput) ToTagKeyIamPolicyArrayOutputWithContext(ctx context.Context) TagKeyIamPolicyArrayOutput {
	return o
}

func (o TagKeyIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TagKeyIamPolicy] {
	return pulumix.Output[[]*TagKeyIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o TagKeyIamPolicyArrayOutput) Index(i pulumi.IntInput) TagKeyIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagKeyIamPolicy {
		return vs[0].([]*TagKeyIamPolicy)[vs[1].(int)]
	}).(TagKeyIamPolicyOutput)
}

type TagKeyIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TagKeyIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagKeyIamPolicy)(nil)).Elem()
}

func (o TagKeyIamPolicyMapOutput) ToTagKeyIamPolicyMapOutput() TagKeyIamPolicyMapOutput {
	return o
}

func (o TagKeyIamPolicyMapOutput) ToTagKeyIamPolicyMapOutputWithContext(ctx context.Context) TagKeyIamPolicyMapOutput {
	return o
}

func (o TagKeyIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TagKeyIamPolicy] {
	return pulumix.Output[map[string]*TagKeyIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o TagKeyIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TagKeyIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagKeyIamPolicy {
		return vs[0].(map[string]*TagKeyIamPolicy)[vs[1].(string)]
	}).(TagKeyIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamPolicyInput)(nil)).Elem(), &TagKeyIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamPolicyArrayInput)(nil)).Elem(), TagKeyIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagKeyIamPolicyMapInput)(nil)).Elem(), TagKeyIamPolicyMap{})
	pulumi.RegisterOutputType(TagKeyIamPolicyOutput{})
	pulumi.RegisterOutputType(TagKeyIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TagKeyIamPolicyMapOutput{})
}
