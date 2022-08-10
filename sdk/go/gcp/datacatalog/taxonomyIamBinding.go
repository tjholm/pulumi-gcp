// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Data catalog Taxonomy. Each of these resources serves a different use case:
//
// * `datacatalog.TaxonomyIamPolicy`: Authoritative. Sets the IAM policy for the taxonomy and replaces any existing policy already attached.
// * `datacatalog.TaxonomyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the taxonomy are preserved.
// * `datacatalog.TaxonomyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the taxonomy are preserved.
//
// > **Note:** `datacatalog.TaxonomyIamPolicy` **cannot** be used in conjunction with `datacatalog.TaxonomyIamBinding` and `datacatalog.TaxonomyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.TaxonomyIamBinding` resources **can be** used in conjunction with `datacatalog.TaxonomyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_data\_catalog\_taxonomy\_iam\_policy
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
//			_, err = datacatalog.NewTaxonomyIamPolicy(ctx, "policy", &datacatalog.TaxonomyIamPolicyArgs{
//				Taxonomy:   pulumi.Any(google_data_catalog_taxonomy.Basic_taxonomy.Name),
//				PolicyData: pulumi.String(admin.PolicyData),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_data\_catalog\_taxonomy\_iam\_binding
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
//			_, err := datacatalog.NewTaxonomyIamBinding(ctx, "binding", &datacatalog.TaxonomyIamBindingArgs{
//				Taxonomy: pulumi.Any(google_data_catalog_taxonomy.Basic_taxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## google\_data\_catalog\_taxonomy\_iam\_member
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
//			_, err := datacatalog.NewTaxonomyIamMember(ctx, "member", &datacatalog.TaxonomyIamMemberArgs{
//				Taxonomy: pulumi.Any(google_data_catalog_taxonomy.Basic_taxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
//			}, pulumi.Provider(google_beta))
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} * {{project}}/{{region}}/{{taxonomy}} * {{region}}/{{taxonomy}} * {{taxonomy}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog taxonomy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/taxonomyIamBinding:TaxonomyIamBinding editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/taxonomyIamBinding:TaxonomyIamBinding editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:datacatalog/taxonomyIamBinding:TaxonomyIamBinding editor projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TaxonomyIamBinding struct {
	pulumi.CustomResourceState

	Condition TaxonomyIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringOutput `pulumi:"taxonomy"`
}

// NewTaxonomyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyIamBinding(ctx *pulumi.Context,
	name string, args *TaxonomyIamBindingArgs, opts ...pulumi.ResourceOption) (*TaxonomyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	var resource TaxonomyIamBinding
	err := ctx.RegisterResource("gcp:datacatalog/taxonomyIamBinding:TaxonomyIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyIamBinding gets an existing TaxonomyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyIamBindingState, opts ...pulumi.ResourceOption) (*TaxonomyIamBinding, error) {
	var resource TaxonomyIamBinding
	err := ctx.ReadResource("gcp:datacatalog/taxonomyIamBinding:TaxonomyIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyIamBinding resources.
type taxonomyIamBindingState struct {
	Condition *TaxonomyIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy *string `pulumi:"taxonomy"`
}

type TaxonomyIamBindingState struct {
	Condition TaxonomyIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringPtrInput
}

func (TaxonomyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamBindingState)(nil)).Elem()
}

type taxonomyIamBindingArgs struct {
	Condition *TaxonomyIamBindingCondition `pulumi:"condition"`
	Members   []string                     `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy string `pulumi:"taxonomy"`
}

// The set of arguments for constructing a TaxonomyIamBinding resource.
type TaxonomyIamBindingArgs struct {
	Condition TaxonomyIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringInput
}

func (TaxonomyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamBindingArgs)(nil)).Elem()
}

type TaxonomyIamBindingInput interface {
	pulumi.Input

	ToTaxonomyIamBindingOutput() TaxonomyIamBindingOutput
	ToTaxonomyIamBindingOutputWithContext(ctx context.Context) TaxonomyIamBindingOutput
}

func (*TaxonomyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamBinding)(nil)).Elem()
}

func (i *TaxonomyIamBinding) ToTaxonomyIamBindingOutput() TaxonomyIamBindingOutput {
	return i.ToTaxonomyIamBindingOutputWithContext(context.Background())
}

func (i *TaxonomyIamBinding) ToTaxonomyIamBindingOutputWithContext(ctx context.Context) TaxonomyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamBindingOutput)
}

// TaxonomyIamBindingArrayInput is an input type that accepts TaxonomyIamBindingArray and TaxonomyIamBindingArrayOutput values.
// You can construct a concrete instance of `TaxonomyIamBindingArrayInput` via:
//
//	TaxonomyIamBindingArray{ TaxonomyIamBindingArgs{...} }
type TaxonomyIamBindingArrayInput interface {
	pulumi.Input

	ToTaxonomyIamBindingArrayOutput() TaxonomyIamBindingArrayOutput
	ToTaxonomyIamBindingArrayOutputWithContext(context.Context) TaxonomyIamBindingArrayOutput
}

type TaxonomyIamBindingArray []TaxonomyIamBindingInput

func (TaxonomyIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxonomyIamBinding)(nil)).Elem()
}

func (i TaxonomyIamBindingArray) ToTaxonomyIamBindingArrayOutput() TaxonomyIamBindingArrayOutput {
	return i.ToTaxonomyIamBindingArrayOutputWithContext(context.Background())
}

func (i TaxonomyIamBindingArray) ToTaxonomyIamBindingArrayOutputWithContext(ctx context.Context) TaxonomyIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamBindingArrayOutput)
}

// TaxonomyIamBindingMapInput is an input type that accepts TaxonomyIamBindingMap and TaxonomyIamBindingMapOutput values.
// You can construct a concrete instance of `TaxonomyIamBindingMapInput` via:
//
//	TaxonomyIamBindingMap{ "key": TaxonomyIamBindingArgs{...} }
type TaxonomyIamBindingMapInput interface {
	pulumi.Input

	ToTaxonomyIamBindingMapOutput() TaxonomyIamBindingMapOutput
	ToTaxonomyIamBindingMapOutputWithContext(context.Context) TaxonomyIamBindingMapOutput
}

type TaxonomyIamBindingMap map[string]TaxonomyIamBindingInput

func (TaxonomyIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxonomyIamBinding)(nil)).Elem()
}

func (i TaxonomyIamBindingMap) ToTaxonomyIamBindingMapOutput() TaxonomyIamBindingMapOutput {
	return i.ToTaxonomyIamBindingMapOutputWithContext(context.Background())
}

func (i TaxonomyIamBindingMap) ToTaxonomyIamBindingMapOutputWithContext(ctx context.Context) TaxonomyIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamBindingMapOutput)
}

type TaxonomyIamBindingOutput struct{ *pulumi.OutputState }

func (TaxonomyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamBinding)(nil)).Elem()
}

func (o TaxonomyIamBindingOutput) ToTaxonomyIamBindingOutput() TaxonomyIamBindingOutput {
	return o
}

func (o TaxonomyIamBindingOutput) ToTaxonomyIamBindingOutputWithContext(ctx context.Context) TaxonomyIamBindingOutput {
	return o
}

func (o TaxonomyIamBindingOutput) Condition() TaxonomyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) TaxonomyIamBindingConditionPtrOutput { return v.Condition }).(TaxonomyIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o TaxonomyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TaxonomyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TaxonomyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o TaxonomyIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o TaxonomyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o TaxonomyIamBindingOutput) Taxonomy() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamBinding) pulumi.StringOutput { return v.Taxonomy }).(pulumi.StringOutput)
}

type TaxonomyIamBindingArrayOutput struct{ *pulumi.OutputState }

func (TaxonomyIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxonomyIamBinding)(nil)).Elem()
}

func (o TaxonomyIamBindingArrayOutput) ToTaxonomyIamBindingArrayOutput() TaxonomyIamBindingArrayOutput {
	return o
}

func (o TaxonomyIamBindingArrayOutput) ToTaxonomyIamBindingArrayOutputWithContext(ctx context.Context) TaxonomyIamBindingArrayOutput {
	return o
}

func (o TaxonomyIamBindingArrayOutput) Index(i pulumi.IntInput) TaxonomyIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaxonomyIamBinding {
		return vs[0].([]*TaxonomyIamBinding)[vs[1].(int)]
	}).(TaxonomyIamBindingOutput)
}

type TaxonomyIamBindingMapOutput struct{ *pulumi.OutputState }

func (TaxonomyIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxonomyIamBinding)(nil)).Elem()
}

func (o TaxonomyIamBindingMapOutput) ToTaxonomyIamBindingMapOutput() TaxonomyIamBindingMapOutput {
	return o
}

func (o TaxonomyIamBindingMapOutput) ToTaxonomyIamBindingMapOutputWithContext(ctx context.Context) TaxonomyIamBindingMapOutput {
	return o
}

func (o TaxonomyIamBindingMapOutput) MapIndex(k pulumi.StringInput) TaxonomyIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaxonomyIamBinding {
		return vs[0].(map[string]*TaxonomyIamBinding)[vs[1].(string)]
	}).(TaxonomyIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamBindingInput)(nil)).Elem(), &TaxonomyIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamBindingArrayInput)(nil)).Elem(), TaxonomyIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamBindingMapInput)(nil)).Elem(), TaxonomyIamBindingMap{})
	pulumi.RegisterOutputType(TaxonomyIamBindingOutput{})
	pulumi.RegisterOutputType(TaxonomyIamBindingArrayOutput{})
	pulumi.RegisterOutputType(TaxonomyIamBindingMapOutput{})
}
