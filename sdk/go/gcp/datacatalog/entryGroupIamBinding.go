// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Data catalog EntryGroup. Each of these resources serves a different use case:
//
// * `datacatalog.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
// * `datacatalog.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
// * `datacatalog.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.
//
// > **Note:** `datacatalog.EntryGroupIamPolicy` **cannot** be used in conjunction with `datacatalog.EntryGroupIamBinding` and `datacatalog.EntryGroupIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.EntryGroupIamBinding` resources **can be** used in conjunction with `datacatalog.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_data\_catalog\_entry\_group\_iam\_policy
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
//					organizations.GetIAMPolicyBinding{
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
//			_, err = datacatalog.NewEntryGroupIamPolicy(ctx, "policy", &datacatalog.EntryGroupIamPolicyArgs{
//				EntryGroup: pulumi.Any(google_data_catalog_entry_group.Basic_entry_group.Name),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## google\_data\_catalog\_entry\_group\_iam\_binding
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
//			_, err := datacatalog.NewEntryGroupIamBinding(ctx, "binding", &datacatalog.EntryGroupIamBindingArgs{
//				EntryGroup: pulumi.Any(google_data_catalog_entry_group.Basic_entry_group.Name),
//				Role:       pulumi.String("roles/viewer"),
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
// ## google\_data\_catalog\_entry\_group\_iam\_member
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
//			_, err := datacatalog.NewEntryGroupIamMember(ctx, "member", &datacatalog.EntryGroupIamMemberArgs{
//				EntryGroup: pulumi.Any(google_data_catalog_entry_group.Basic_entry_group.Name),
//				Role:       pulumi.String("roles/viewer"),
//				Member:     pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} * {{project}}/{{region}}/{{entry_group}} * {{region}}/{{entry_group}} * {{entry_group}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog entrygroup IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding editor projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type EntryGroupIamBinding struct {
	pulumi.CustomResourceState

	Condition EntryGroupIamBindingConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringOutput `pulumi:"entryGroup"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewEntryGroupIamBinding registers a new resource with the given unique name, arguments, and options.
func NewEntryGroupIamBinding(ctx *pulumi.Context,
	name string, args *EntryGroupIamBindingArgs, opts ...pulumi.ResourceOption) (*EntryGroupIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroup == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroup'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource EntryGroupIamBinding
	err := ctx.RegisterResource("gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroupIamBinding gets an existing EntryGroupIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroupIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupIamBindingState, opts ...pulumi.ResourceOption) (*EntryGroupIamBinding, error) {
	var resource EntryGroupIamBinding
	err := ctx.ReadResource("gcp:datacatalog/entryGroupIamBinding:EntryGroupIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroupIamBinding resources.
type entryGroupIamBindingState struct {
	Condition *EntryGroupIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup *string `pulumi:"entryGroup"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type EntryGroupIamBindingState struct {
	Condition EntryGroupIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (EntryGroupIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamBindingState)(nil)).Elem()
}

type entryGroupIamBindingArgs struct {
	Condition *EntryGroupIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup string   `pulumi:"entryGroup"`
	Members    []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a EntryGroupIamBinding resource.
type EntryGroupIamBindingArgs struct {
	Condition EntryGroupIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	EntryGroup pulumi.StringInput
	Members    pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (EntryGroupIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamBindingArgs)(nil)).Elem()
}

type EntryGroupIamBindingInput interface {
	pulumi.Input

	ToEntryGroupIamBindingOutput() EntryGroupIamBindingOutput
	ToEntryGroupIamBindingOutputWithContext(ctx context.Context) EntryGroupIamBindingOutput
}

func (*EntryGroupIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamBinding)(nil)).Elem()
}

func (i *EntryGroupIamBinding) ToEntryGroupIamBindingOutput() EntryGroupIamBindingOutput {
	return i.ToEntryGroupIamBindingOutputWithContext(context.Background())
}

func (i *EntryGroupIamBinding) ToEntryGroupIamBindingOutputWithContext(ctx context.Context) EntryGroupIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamBindingOutput)
}

// EntryGroupIamBindingArrayInput is an input type that accepts EntryGroupIamBindingArray and EntryGroupIamBindingArrayOutput values.
// You can construct a concrete instance of `EntryGroupIamBindingArrayInput` via:
//
//	EntryGroupIamBindingArray{ EntryGroupIamBindingArgs{...} }
type EntryGroupIamBindingArrayInput interface {
	pulumi.Input

	ToEntryGroupIamBindingArrayOutput() EntryGroupIamBindingArrayOutput
	ToEntryGroupIamBindingArrayOutputWithContext(context.Context) EntryGroupIamBindingArrayOutput
}

type EntryGroupIamBindingArray []EntryGroupIamBindingInput

func (EntryGroupIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroupIamBinding)(nil)).Elem()
}

func (i EntryGroupIamBindingArray) ToEntryGroupIamBindingArrayOutput() EntryGroupIamBindingArrayOutput {
	return i.ToEntryGroupIamBindingArrayOutputWithContext(context.Background())
}

func (i EntryGroupIamBindingArray) ToEntryGroupIamBindingArrayOutputWithContext(ctx context.Context) EntryGroupIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamBindingArrayOutput)
}

// EntryGroupIamBindingMapInput is an input type that accepts EntryGroupIamBindingMap and EntryGroupIamBindingMapOutput values.
// You can construct a concrete instance of `EntryGroupIamBindingMapInput` via:
//
//	EntryGroupIamBindingMap{ "key": EntryGroupIamBindingArgs{...} }
type EntryGroupIamBindingMapInput interface {
	pulumi.Input

	ToEntryGroupIamBindingMapOutput() EntryGroupIamBindingMapOutput
	ToEntryGroupIamBindingMapOutputWithContext(context.Context) EntryGroupIamBindingMapOutput
}

type EntryGroupIamBindingMap map[string]EntryGroupIamBindingInput

func (EntryGroupIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroupIamBinding)(nil)).Elem()
}

func (i EntryGroupIamBindingMap) ToEntryGroupIamBindingMapOutput() EntryGroupIamBindingMapOutput {
	return i.ToEntryGroupIamBindingMapOutputWithContext(context.Background())
}

func (i EntryGroupIamBindingMap) ToEntryGroupIamBindingMapOutputWithContext(ctx context.Context) EntryGroupIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamBindingMapOutput)
}

type EntryGroupIamBindingOutput struct{ *pulumi.OutputState }

func (EntryGroupIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamBinding)(nil)).Elem()
}

func (o EntryGroupIamBindingOutput) ToEntryGroupIamBindingOutput() EntryGroupIamBindingOutput {
	return o
}

func (o EntryGroupIamBindingOutput) ToEntryGroupIamBindingOutputWithContext(ctx context.Context) EntryGroupIamBindingOutput {
	return o
}

func (o EntryGroupIamBindingOutput) Condition() EntryGroupIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) EntryGroupIamBindingConditionPtrOutput { return v.Condition }).(EntryGroupIamBindingConditionPtrOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o EntryGroupIamBindingOutput) EntryGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringOutput { return v.EntryGroup }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o EntryGroupIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o EntryGroupIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o EntryGroupIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o EntryGroupIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o EntryGroupIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type EntryGroupIamBindingArrayOutput struct{ *pulumi.OutputState }

func (EntryGroupIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntryGroupIamBinding)(nil)).Elem()
}

func (o EntryGroupIamBindingArrayOutput) ToEntryGroupIamBindingArrayOutput() EntryGroupIamBindingArrayOutput {
	return o
}

func (o EntryGroupIamBindingArrayOutput) ToEntryGroupIamBindingArrayOutputWithContext(ctx context.Context) EntryGroupIamBindingArrayOutput {
	return o
}

func (o EntryGroupIamBindingArrayOutput) Index(i pulumi.IntInput) EntryGroupIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EntryGroupIamBinding {
		return vs[0].([]*EntryGroupIamBinding)[vs[1].(int)]
	}).(EntryGroupIamBindingOutput)
}

type EntryGroupIamBindingMapOutput struct{ *pulumi.OutputState }

func (EntryGroupIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntryGroupIamBinding)(nil)).Elem()
}

func (o EntryGroupIamBindingMapOutput) ToEntryGroupIamBindingMapOutput() EntryGroupIamBindingMapOutput {
	return o
}

func (o EntryGroupIamBindingMapOutput) ToEntryGroupIamBindingMapOutputWithContext(ctx context.Context) EntryGroupIamBindingMapOutput {
	return o
}

func (o EntryGroupIamBindingMapOutput) MapIndex(k pulumi.StringInput) EntryGroupIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EntryGroupIamBinding {
		return vs[0].(map[string]*EntryGroupIamBinding)[vs[1].(string)]
	}).(EntryGroupIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamBindingInput)(nil)).Elem(), &EntryGroupIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamBindingArrayInput)(nil)).Elem(), EntryGroupIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamBindingMapInput)(nil)).Elem(), EntryGroupIamBindingMap{})
	pulumi.RegisterOutputType(EntryGroupIamBindingOutput{})
	pulumi.RegisterOutputType(EntryGroupIamBindingArrayOutput{})
	pulumi.RegisterOutputType(EntryGroupIamBindingMapOutput{})
}
