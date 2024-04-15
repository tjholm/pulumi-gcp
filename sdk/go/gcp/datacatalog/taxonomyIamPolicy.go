// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Data catalog Taxonomy. Each of these resources serves a different use case:
//
// * `datacatalog.TaxonomyIamPolicy`: Authoritative. Sets the IAM policy for the taxonomy and replaces any existing policy already attached.
// * `datacatalog.TaxonomyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the taxonomy are preserved.
// * `datacatalog.TaxonomyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the taxonomy are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `datacatalog.TaxonomyIamPolicy`: Retrieves the IAM policy for the taxonomy
//
// > **Note:** `datacatalog.TaxonomyIamPolicy` **cannot** be used in conjunction with `datacatalog.TaxonomyIamBinding` and `datacatalog.TaxonomyIamMember` or they will fight over what your policy should be.
//
// > **Note:** `datacatalog.TaxonomyIamBinding` resources **can be** used in conjunction with `datacatalog.TaxonomyIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_data\_catalog\_taxonomy\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
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
//			_, err = datacatalog.NewTaxonomyIamPolicy(ctx, "policy", &datacatalog.TaxonomyIamPolicyArgs{
//				Taxonomy:   pulumi.Any(basicTaxonomy.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_data\_catalog\_taxonomy\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTaxonomyIamBinding(ctx, "binding", &datacatalog.TaxonomyIamBindingArgs{
//				Taxonomy: pulumi.Any(basicTaxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_data\_catalog\_taxonomy\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTaxonomyIamMember(ctx, "member", &datacatalog.TaxonomyIamMemberArgs{
//				Taxonomy: pulumi.Any(basicTaxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## google\_data\_catalog\_taxonomy\_iam\_policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
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
//			_, err = datacatalog.NewTaxonomyIamPolicy(ctx, "policy", &datacatalog.TaxonomyIamPolicyArgs{
//				Taxonomy:   pulumi.Any(basicTaxonomy.Name),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_data\_catalog\_taxonomy\_iam\_binding
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTaxonomyIamBinding(ctx, "binding", &datacatalog.TaxonomyIamBindingArgs{
//				Taxonomy: pulumi.Any(basicTaxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
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
// <!--End PulumiCodeChooser -->
//
// ## google\_data\_catalog\_taxonomy\_iam\_member
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/datacatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datacatalog.NewTaxonomyIamMember(ctx, "member", &datacatalog.TaxonomyIamMemberArgs{
//				Taxonomy: pulumi.Any(basicTaxonomy.Name),
//				Role:     pulumi.String("roles/viewer"),
//				Member:   pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
//
// * {{project}}/{{region}}/{{taxonomy}}
//
// * {{region}}/{{taxonomy}}
//
// * {{taxonomy}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Data catalog taxonomy IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TaxonomyIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringOutput `pulumi:"taxonomy"`
}

// NewTaxonomyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyIamPolicy(ctx *pulumi.Context,
	name string, args *TaxonomyIamPolicyArgs, opts ...pulumi.ResourceOption) (*TaxonomyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TaxonomyIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyIamPolicy gets an existing TaxonomyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyIamPolicyState, opts ...pulumi.ResourceOption) (*TaxonomyIamPolicy, error) {
	var resource TaxonomyIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyIamPolicy resources.
type taxonomyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy *string `pulumi:"taxonomy"`
}

type TaxonomyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringPtrInput
}

func (TaxonomyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamPolicyState)(nil)).Elem()
}

type taxonomyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy string `pulumi:"taxonomy"`
}

// The set of arguments for constructing a TaxonomyIamPolicy resource.
type TaxonomyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringInput
}

func (TaxonomyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamPolicyArgs)(nil)).Elem()
}

type TaxonomyIamPolicyInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput
	ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput
}

func (*TaxonomyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamPolicy)(nil)).Elem()
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput {
	return i.ToTaxonomyIamPolicyOutputWithContext(context.Background())
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyOutput)
}

// TaxonomyIamPolicyArrayInput is an input type that accepts TaxonomyIamPolicyArray and TaxonomyIamPolicyArrayOutput values.
// You can construct a concrete instance of `TaxonomyIamPolicyArrayInput` via:
//
//	TaxonomyIamPolicyArray{ TaxonomyIamPolicyArgs{...} }
type TaxonomyIamPolicyArrayInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput
	ToTaxonomyIamPolicyArrayOutputWithContext(context.Context) TaxonomyIamPolicyArrayOutput
}

type TaxonomyIamPolicyArray []TaxonomyIamPolicyInput

func (TaxonomyIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxonomyIamPolicy)(nil)).Elem()
}

func (i TaxonomyIamPolicyArray) ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput {
	return i.ToTaxonomyIamPolicyArrayOutputWithContext(context.Background())
}

func (i TaxonomyIamPolicyArray) ToTaxonomyIamPolicyArrayOutputWithContext(ctx context.Context) TaxonomyIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyArrayOutput)
}

// TaxonomyIamPolicyMapInput is an input type that accepts TaxonomyIamPolicyMap and TaxonomyIamPolicyMapOutput values.
// You can construct a concrete instance of `TaxonomyIamPolicyMapInput` via:
//
//	TaxonomyIamPolicyMap{ "key": TaxonomyIamPolicyArgs{...} }
type TaxonomyIamPolicyMapInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput
	ToTaxonomyIamPolicyMapOutputWithContext(context.Context) TaxonomyIamPolicyMapOutput
}

type TaxonomyIamPolicyMap map[string]TaxonomyIamPolicyInput

func (TaxonomyIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxonomyIamPolicy)(nil)).Elem()
}

func (i TaxonomyIamPolicyMap) ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput {
	return i.ToTaxonomyIamPolicyMapOutputWithContext(context.Background())
}

func (i TaxonomyIamPolicyMap) ToTaxonomyIamPolicyMapOutputWithContext(ctx context.Context) TaxonomyIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyMapOutput)
}

type TaxonomyIamPolicyOutput struct{ *pulumi.OutputState }

func (TaxonomyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamPolicy)(nil)).Elem()
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput {
	return o
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o TaxonomyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o TaxonomyIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o TaxonomyIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o TaxonomyIamPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o TaxonomyIamPolicyOutput) Taxonomy() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyIamPolicy) pulumi.StringOutput { return v.Taxonomy }).(pulumi.StringOutput)
}

type TaxonomyIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TaxonomyIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxonomyIamPolicy)(nil)).Elem()
}

func (o TaxonomyIamPolicyArrayOutput) ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput {
	return o
}

func (o TaxonomyIamPolicyArrayOutput) ToTaxonomyIamPolicyArrayOutputWithContext(ctx context.Context) TaxonomyIamPolicyArrayOutput {
	return o
}

func (o TaxonomyIamPolicyArrayOutput) Index(i pulumi.IntInput) TaxonomyIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaxonomyIamPolicy {
		return vs[0].([]*TaxonomyIamPolicy)[vs[1].(int)]
	}).(TaxonomyIamPolicyOutput)
}

type TaxonomyIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TaxonomyIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxonomyIamPolicy)(nil)).Elem()
}

func (o TaxonomyIamPolicyMapOutput) ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput {
	return o
}

func (o TaxonomyIamPolicyMapOutput) ToTaxonomyIamPolicyMapOutputWithContext(ctx context.Context) TaxonomyIamPolicyMapOutput {
	return o
}

func (o TaxonomyIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TaxonomyIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaxonomyIamPolicy {
		return vs[0].(map[string]*TaxonomyIamPolicy)[vs[1].(string)]
	}).(TaxonomyIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamPolicyInput)(nil)).Elem(), &TaxonomyIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamPolicyArrayInput)(nil)).Elem(), TaxonomyIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyIamPolicyMapInput)(nil)).Elem(), TaxonomyIamPolicyMap{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyOutput{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyMapOutput{})
}
