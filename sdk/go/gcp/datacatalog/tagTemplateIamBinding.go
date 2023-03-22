// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Data catalog TagTemplate. Each of these resources serves a different use case:
//
// * `datacatalog.TagTemplateIamPolicy`: Authoritative. Sets the IAM policy for the tagtemplate and replaces any existing policy already attached.
// * `datacatalog.TagTemplateIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagtemplate are preserved.
// * `datacatalog.TagTemplateIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagtemplate are preserved.
//
// > **Note:** `datacatalog.TagTemplateIamPolicy` **cannot** be used in conjunction with `datacatalog.TagTemplateIamBinding` and `datacatalog.TagTemplateIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.TagTemplateIamBinding` resources **can be** used in conjunction with `datacatalog.TagTemplateIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_data\_catalog\_tag\_template\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
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
//			_, err = datacatalog.NewTagTemplateIamPolicy(ctx, "policy", &datacatalog.TagTemplateIamPolicyArgs{
//				TagTemplate: pulumi.Any(google_data_catalog_tag_template.Basic_tag_template.Name),
//				PolicyData:  *pulumi.String(admin.PolicyData),
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
// ## google\_data\_catalog\_tag\_template\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTagTemplateIamBinding(ctx, "binding", &datacatalog.TagTemplateIamBindingArgs{
//				TagTemplate: pulumi.Any(google_data_catalog_tag_template.Basic_tag_template.Name),
//				Role:        pulumi.String("roles/viewer"),
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
// ## google\_data\_catalog\_tag\_template\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTagTemplateIamMember(ctx, "member", &datacatalog.TagTemplateIamMemberArgs{
//				TagTemplate: pulumi.Any(google_data_catalog_tag_template.Basic_tag_template.Name),
//				Role:        pulumi.String("roles/viewer"),
//				Member:      pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} * {{project}}/{{region}}/{{tag_template}} * {{region}}/{{tag_template}} * {{tag_template}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog tagtemplate IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/tagTemplateIamBinding:TagTemplateIamBinding editor "projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/tagTemplateIamBinding:TagTemplateIamBinding editor "projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/tagTemplateIamBinding:TagTemplateIamBinding editor projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TagTemplateIamBinding struct {
	pulumi.CustomResourceState

	Condition TagTemplateIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagTemplate pulumi.StringOutput `pulumi:"tagTemplate"`
}

// NewTagTemplateIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTagTemplateIamBinding(ctx *pulumi.Context,
	name string, args *TagTemplateIamBindingArgs, opts ...pulumi.ResourceOption) (*TagTemplateIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagTemplate == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplate'")
	}
	var resource TagTemplateIamBinding
	err := ctx.RegisterResource("gcp:datacatalog/tagTemplateIamBinding:TagTemplateIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagTemplateIamBinding gets an existing TagTemplateIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagTemplateIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagTemplateIamBindingState, opts ...pulumi.ResourceOption) (*TagTemplateIamBinding, error) {
	var resource TagTemplateIamBinding
	err := ctx.ReadResource("gcp:datacatalog/tagTemplateIamBinding:TagTemplateIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagTemplateIamBinding resources.
type tagTemplateIamBindingState struct {
	Condition *TagTemplateIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagTemplate *string `pulumi:"tagTemplate"`
}

type TagTemplateIamBindingState struct {
	Condition TagTemplateIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	TagTemplate pulumi.StringPtrInput
}

func (TagTemplateIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamBindingState)(nil)).Elem()
}

type tagTemplateIamBindingArgs struct {
	Condition *TagTemplateIamBindingCondition `pulumi:"condition"`
	Members   []string                        `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	TagTemplate string `pulumi:"tagTemplate"`
}

// The set of arguments for constructing a TagTemplateIamBinding resource.
type TagTemplateIamBindingArgs struct {
	Condition TagTemplateIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	TagTemplate pulumi.StringInput
}

func (TagTemplateIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagTemplateIamBindingArgs)(nil)).Elem()
}

type TagTemplateIamBindingInput interface {
	pulumi.Input

	ToTagTemplateIamBindingOutput() TagTemplateIamBindingOutput
	ToTagTemplateIamBindingOutputWithContext(ctx context.Context) TagTemplateIamBindingOutput
}

func (*TagTemplateIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamBinding)(nil)).Elem()
}

func (i *TagTemplateIamBinding) ToTagTemplateIamBindingOutput() TagTemplateIamBindingOutput {
	return i.ToTagTemplateIamBindingOutputWithContext(context.Background())
}

func (i *TagTemplateIamBinding) ToTagTemplateIamBindingOutputWithContext(ctx context.Context) TagTemplateIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamBindingOutput)
}

// TagTemplateIamBindingArrayInput is an input type that accepts TagTemplateIamBindingArray and TagTemplateIamBindingArrayOutput values.
// You can construct a concrete instance of `TagTemplateIamBindingArrayInput` via:
//
//	TagTemplateIamBindingArray{ TagTemplateIamBindingArgs{...} }
type TagTemplateIamBindingArrayInput interface {
	pulumi.Input

	ToTagTemplateIamBindingArrayOutput() TagTemplateIamBindingArrayOutput
	ToTagTemplateIamBindingArrayOutputWithContext(context.Context) TagTemplateIamBindingArrayOutput
}

type TagTemplateIamBindingArray []TagTemplateIamBindingInput

func (TagTemplateIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagTemplateIamBinding)(nil)).Elem()
}

func (i TagTemplateIamBindingArray) ToTagTemplateIamBindingArrayOutput() TagTemplateIamBindingArrayOutput {
	return i.ToTagTemplateIamBindingArrayOutputWithContext(context.Background())
}

func (i TagTemplateIamBindingArray) ToTagTemplateIamBindingArrayOutputWithContext(ctx context.Context) TagTemplateIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamBindingArrayOutput)
}

// TagTemplateIamBindingMapInput is an input type that accepts TagTemplateIamBindingMap and TagTemplateIamBindingMapOutput values.
// You can construct a concrete instance of `TagTemplateIamBindingMapInput` via:
//
//	TagTemplateIamBindingMap{ "key": TagTemplateIamBindingArgs{...} }
type TagTemplateIamBindingMapInput interface {
	pulumi.Input

	ToTagTemplateIamBindingMapOutput() TagTemplateIamBindingMapOutput
	ToTagTemplateIamBindingMapOutputWithContext(context.Context) TagTemplateIamBindingMapOutput
}

type TagTemplateIamBindingMap map[string]TagTemplateIamBindingInput

func (TagTemplateIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagTemplateIamBinding)(nil)).Elem()
}

func (i TagTemplateIamBindingMap) ToTagTemplateIamBindingMapOutput() TagTemplateIamBindingMapOutput {
	return i.ToTagTemplateIamBindingMapOutputWithContext(context.Background())
}

func (i TagTemplateIamBindingMap) ToTagTemplateIamBindingMapOutputWithContext(ctx context.Context) TagTemplateIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagTemplateIamBindingMapOutput)
}

type TagTemplateIamBindingOutput struct{ *pulumi.OutputState }

func (TagTemplateIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagTemplateIamBinding)(nil)).Elem()
}

func (o TagTemplateIamBindingOutput) ToTagTemplateIamBindingOutput() TagTemplateIamBindingOutput {
	return o
}

func (o TagTemplateIamBindingOutput) ToTagTemplateIamBindingOutputWithContext(ctx context.Context) TagTemplateIamBindingOutput {
	return o
}

func (o TagTemplateIamBindingOutput) Condition() TagTemplateIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) TagTemplateIamBindingConditionPtrOutput { return v.Condition }).(TagTemplateIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o TagTemplateIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TagTemplateIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TagTemplateIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o TagTemplateIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `datacatalog.TagTemplateIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o TagTemplateIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o TagTemplateIamBindingOutput) TagTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *TagTemplateIamBinding) pulumi.StringOutput { return v.TagTemplate }).(pulumi.StringOutput)
}

type TagTemplateIamBindingArrayOutput struct{ *pulumi.OutputState }

func (TagTemplateIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagTemplateIamBinding)(nil)).Elem()
}

func (o TagTemplateIamBindingArrayOutput) ToTagTemplateIamBindingArrayOutput() TagTemplateIamBindingArrayOutput {
	return o
}

func (o TagTemplateIamBindingArrayOutput) ToTagTemplateIamBindingArrayOutputWithContext(ctx context.Context) TagTemplateIamBindingArrayOutput {
	return o
}

func (o TagTemplateIamBindingArrayOutput) Index(i pulumi.IntInput) TagTemplateIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TagTemplateIamBinding {
		return vs[0].([]*TagTemplateIamBinding)[vs[1].(int)]
	}).(TagTemplateIamBindingOutput)
}

type TagTemplateIamBindingMapOutput struct{ *pulumi.OutputState }

func (TagTemplateIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagTemplateIamBinding)(nil)).Elem()
}

func (o TagTemplateIamBindingMapOutput) ToTagTemplateIamBindingMapOutput() TagTemplateIamBindingMapOutput {
	return o
}

func (o TagTemplateIamBindingMapOutput) ToTagTemplateIamBindingMapOutputWithContext(ctx context.Context) TagTemplateIamBindingMapOutput {
	return o
}

func (o TagTemplateIamBindingMapOutput) MapIndex(k pulumi.StringInput) TagTemplateIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TagTemplateIamBinding {
		return vs[0].(map[string]*TagTemplateIamBinding)[vs[1].(string)]
	}).(TagTemplateIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamBindingInput)(nil)).Elem(), &TagTemplateIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamBindingArrayInput)(nil)).Elem(), TagTemplateIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagTemplateIamBindingMapInput)(nil)).Elem(), TagTemplateIamBindingMap{})
	pulumi.RegisterOutputType(TagTemplateIamBindingOutput{})
	pulumi.RegisterOutputType(TagTemplateIamBindingArrayOutput{})
	pulumi.RegisterOutputType(TagTemplateIamBindingMapOutput{})
}
