// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for Certificate Authority Service CaPool. Each of these resources serves a different use case:
//
// * `certificateauthority.CaPoolIamPolicy`: Authoritative. Sets the IAM policy for the capool and replaces any existing policy already attached.
// * `certificateauthority.CaPoolIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the capool are preserved.
// * `certificateauthority.CaPoolIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the capool are preserved.
//
// > **Note:** `certificateauthority.CaPoolIamPolicy` **cannot** be used in conjunction with `certificateauthority.CaPoolIamBinding` and `certificateauthority.CaPoolIamMember` or they will fight over what your policy should be.
//
// > **Note:** `certificateauthority.CaPoolIamBinding` resources **can be** used in conjunction with `certificateauthority.CaPoolIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## google\_privateca\_ca\_pool\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
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
//						Role: "roles/privateca.certificateManager",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = certificateauthority.NewCaPoolIamPolicy(ctx, "policy", &certificateauthority.CaPoolIamPolicyArgs{
//				CaPool:     pulumi.Any(google_privateca_ca_pool.Default.Id),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
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
//						Role: "roles/privateca.certificateManager",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = certificateauthority.NewCaPoolIamPolicy(ctx, "policy", &certificateauthority.CaPoolIamPolicyArgs{
//				CaPool:     pulumi.Any(google_privateca_ca_pool.Default.Id),
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
// ## google\_privateca\_ca\_pool\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificateauthority.NewCaPoolIamBinding(ctx, "binding", &certificateauthority.CaPoolIamBindingArgs{
//				CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
//				Role:   pulumi.String("roles/privateca.certificateManager"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificateauthority.NewCaPoolIamBinding(ctx, "binding", &certificateauthority.CaPoolIamBindingArgs{
//				CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
//				Role:   pulumi.String("roles/privateca.certificateManager"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &certificateauthority.CaPoolIamBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## google\_privateca\_ca\_pool\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificateauthority.NewCaPoolIamMember(ctx, "member", &certificateauthority.CaPoolIamMemberArgs{
//				CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
//				Role:   pulumi.String("roles/privateca.certificateManager"),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := certificateauthority.NewCaPoolIamMember(ctx, "member", &certificateauthority.CaPoolIamMemberArgs{
//				CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
//				Role:   pulumi.String("roles/privateca.certificateManager"),
//				Member: pulumi.String("user:jane@example.com"),
//				Condition: &certificateauthority.CaPoolIamMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/caPools/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} Any variables not passed in the import command will be taken from the provider configuration. Certificate Authority Service capool IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding editor projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type CaPoolIamBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CaPoolIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCaPoolIamBinding registers a new resource with the given unique name, arguments, and options.
func NewCaPoolIamBinding(ctx *pulumi.Context,
	name string, args *CaPoolIamBindingArgs, opts ...pulumi.ResourceOption) (*CaPoolIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource CaPoolIamBinding
	err := ctx.RegisterResource("gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPoolIamBinding gets an existing CaPoolIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPoolIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolIamBindingState, opts ...pulumi.ResourceOption) (*CaPoolIamBinding, error) {
	var resource CaPoolIamBinding
	err := ctx.ReadResource("gcp:certificateauthority/caPoolIamBinding:CaPoolIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPoolIamBinding resources.
type caPoolIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool *string `pulumi:"caPool"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CaPoolIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CaPoolIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringPtrInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CaPoolIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CaPoolIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamBindingState)(nil)).Elem()
}

type caPoolIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool string `pulumi:"caPool"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CaPoolIamBindingCondition `pulumi:"condition"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CaPoolIamBinding resource.
type CaPoolIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CaPoolIamBindingConditionPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CaPoolIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamBindingArgs)(nil)).Elem()
}

type CaPoolIamBindingInput interface {
	pulumi.Input

	ToCaPoolIamBindingOutput() CaPoolIamBindingOutput
	ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput
}

func (*CaPoolIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamBinding)(nil)).Elem()
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingOutput() CaPoolIamBindingOutput {
	return i.ToCaPoolIamBindingOutputWithContext(context.Background())
}

func (i *CaPoolIamBinding) ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingOutput)
}

// CaPoolIamBindingArrayInput is an input type that accepts CaPoolIamBindingArray and CaPoolIamBindingArrayOutput values.
// You can construct a concrete instance of `CaPoolIamBindingArrayInput` via:
//
//	CaPoolIamBindingArray{ CaPoolIamBindingArgs{...} }
type CaPoolIamBindingArrayInput interface {
	pulumi.Input

	ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput
	ToCaPoolIamBindingArrayOutputWithContext(context.Context) CaPoolIamBindingArrayOutput
}

type CaPoolIamBindingArray []CaPoolIamBindingInput

func (CaPoolIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamBinding)(nil)).Elem()
}

func (i CaPoolIamBindingArray) ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput {
	return i.ToCaPoolIamBindingArrayOutputWithContext(context.Background())
}

func (i CaPoolIamBindingArray) ToCaPoolIamBindingArrayOutputWithContext(ctx context.Context) CaPoolIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingArrayOutput)
}

// CaPoolIamBindingMapInput is an input type that accepts CaPoolIamBindingMap and CaPoolIamBindingMapOutput values.
// You can construct a concrete instance of `CaPoolIamBindingMapInput` via:
//
//	CaPoolIamBindingMap{ "key": CaPoolIamBindingArgs{...} }
type CaPoolIamBindingMapInput interface {
	pulumi.Input

	ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput
	ToCaPoolIamBindingMapOutputWithContext(context.Context) CaPoolIamBindingMapOutput
}

type CaPoolIamBindingMap map[string]CaPoolIamBindingInput

func (CaPoolIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamBinding)(nil)).Elem()
}

func (i CaPoolIamBindingMap) ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput {
	return i.ToCaPoolIamBindingMapOutputWithContext(context.Background())
}

func (i CaPoolIamBindingMap) ToCaPoolIamBindingMapOutputWithContext(ctx context.Context) CaPoolIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamBindingMapOutput)
}

type CaPoolIamBindingOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamBinding)(nil)).Elem()
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingOutput() CaPoolIamBindingOutput {
	return o
}

func (o CaPoolIamBindingOutput) ToCaPoolIamBindingOutputWithContext(ctx context.Context) CaPoolIamBindingOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o CaPoolIamBindingOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o CaPoolIamBindingOutput) Condition() CaPoolIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) CaPoolIamBindingConditionPtrOutput { return v.Condition }).(CaPoolIamBindingConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o CaPoolIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Location of the CaPool. A full list of valid locations can be found by
// running `gcloud privateca locations list`.
// Used to find the parent resource to bind the IAM policy to
func (o CaPoolIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CaPoolIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o CaPoolIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o CaPoolIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type CaPoolIamBindingArrayOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamBinding)(nil)).Elem()
}

func (o CaPoolIamBindingArrayOutput) ToCaPoolIamBindingArrayOutput() CaPoolIamBindingArrayOutput {
	return o
}

func (o CaPoolIamBindingArrayOutput) ToCaPoolIamBindingArrayOutputWithContext(ctx context.Context) CaPoolIamBindingArrayOutput {
	return o
}

func (o CaPoolIamBindingArrayOutput) Index(i pulumi.IntInput) CaPoolIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CaPoolIamBinding {
		return vs[0].([]*CaPoolIamBinding)[vs[1].(int)]
	}).(CaPoolIamBindingOutput)
}

type CaPoolIamBindingMapOutput struct{ *pulumi.OutputState }

func (CaPoolIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamBinding)(nil)).Elem()
}

func (o CaPoolIamBindingMapOutput) ToCaPoolIamBindingMapOutput() CaPoolIamBindingMapOutput {
	return o
}

func (o CaPoolIamBindingMapOutput) ToCaPoolIamBindingMapOutputWithContext(ctx context.Context) CaPoolIamBindingMapOutput {
	return o
}

func (o CaPoolIamBindingMapOutput) MapIndex(k pulumi.StringInput) CaPoolIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CaPoolIamBinding {
		return vs[0].(map[string]*CaPoolIamBinding)[vs[1].(string)]
	}).(CaPoolIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamBindingInput)(nil)).Elem(), &CaPoolIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamBindingArrayInput)(nil)).Elem(), CaPoolIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamBindingMapInput)(nil)).Elem(), CaPoolIamBindingMap{})
	pulumi.RegisterOutputType(CaPoolIamBindingOutput{})
	pulumi.RegisterOutputType(CaPoolIamBindingArrayOutput{})
	pulumi.RegisterOutputType(CaPoolIamBindingMapOutput{})
}
