// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
// ## google\_privateca\_ca\_pool\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/privateca.certificateManager",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = certificateauthority.NewCaPoolIamPolicy(ctx, "policy", &certificateauthority.CaPoolIamPolicyArgs{
// 			CaPool:     pulumi.Any(google_privateca_ca_pool.Default.Id),
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
// ## google\_privateca\_ca\_pool\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPoolIamBinding(ctx, "binding", &certificateauthority.CaPoolIamBindingArgs{
// 			CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
// 			Role:   pulumi.String("roles/privateca.certificateManager"),
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
// ## google\_privateca\_ca\_pool\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/certificateauthority"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := certificateauthority.NewCaPoolIamMember(ctx, "member", &certificateauthority.CaPoolIamMemberArgs{
// 			CaPool: pulumi.Any(google_privateca_ca_pool.Default.Id),
// 			Role:   pulumi.String("roles/privateca.certificateManager"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/caPools/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} Any variables not passed in the import command will be taken from the provider configuration. Certificate Authority Service capool IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamPolicy:CaPoolIamPolicy editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamPolicy:CaPoolIamPolicy editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamPolicy:CaPoolIamPolicy editor projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type CaPoolIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewCaPoolIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewCaPoolIamPolicy(ctx *pulumi.Context,
	name string, args *CaPoolIamPolicyArgs, opts ...pulumi.ResourceOption) (*CaPoolIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource CaPoolIamPolicy
	err := ctx.RegisterResource("gcp:certificateauthority/caPoolIamPolicy:CaPoolIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPoolIamPolicy gets an existing CaPoolIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPoolIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolIamPolicyState, opts ...pulumi.ResourceOption) (*CaPoolIamPolicy, error) {
	var resource CaPoolIamPolicy
	err := ctx.ReadResource("gcp:certificateauthority/caPoolIamPolicy:CaPoolIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPoolIamPolicy resources.
type caPoolIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool *string `pulumi:"caPool"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type CaPoolIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (CaPoolIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamPolicyState)(nil)).Elem()
}

type caPoolIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool string `pulumi:"caPool"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a CaPoolIamPolicy resource.
type CaPoolIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (CaPoolIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamPolicyArgs)(nil)).Elem()
}

type CaPoolIamPolicyInput interface {
	pulumi.Input

	ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput
	ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput
}

func (*CaPoolIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamPolicy)(nil)).Elem()
}

func (i *CaPoolIamPolicy) ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput {
	return i.ToCaPoolIamPolicyOutputWithContext(context.Background())
}

func (i *CaPoolIamPolicy) ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamPolicyOutput)
}

// CaPoolIamPolicyArrayInput is an input type that accepts CaPoolIamPolicyArray and CaPoolIamPolicyArrayOutput values.
// You can construct a concrete instance of `CaPoolIamPolicyArrayInput` via:
//
//          CaPoolIamPolicyArray{ CaPoolIamPolicyArgs{...} }
type CaPoolIamPolicyArrayInput interface {
	pulumi.Input

	ToCaPoolIamPolicyArrayOutput() CaPoolIamPolicyArrayOutput
	ToCaPoolIamPolicyArrayOutputWithContext(context.Context) CaPoolIamPolicyArrayOutput
}

type CaPoolIamPolicyArray []CaPoolIamPolicyInput

func (CaPoolIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamPolicy)(nil)).Elem()
}

func (i CaPoolIamPolicyArray) ToCaPoolIamPolicyArrayOutput() CaPoolIamPolicyArrayOutput {
	return i.ToCaPoolIamPolicyArrayOutputWithContext(context.Background())
}

func (i CaPoolIamPolicyArray) ToCaPoolIamPolicyArrayOutputWithContext(ctx context.Context) CaPoolIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamPolicyArrayOutput)
}

// CaPoolIamPolicyMapInput is an input type that accepts CaPoolIamPolicyMap and CaPoolIamPolicyMapOutput values.
// You can construct a concrete instance of `CaPoolIamPolicyMapInput` via:
//
//          CaPoolIamPolicyMap{ "key": CaPoolIamPolicyArgs{...} }
type CaPoolIamPolicyMapInput interface {
	pulumi.Input

	ToCaPoolIamPolicyMapOutput() CaPoolIamPolicyMapOutput
	ToCaPoolIamPolicyMapOutputWithContext(context.Context) CaPoolIamPolicyMapOutput
}

type CaPoolIamPolicyMap map[string]CaPoolIamPolicyInput

func (CaPoolIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamPolicy)(nil)).Elem()
}

func (i CaPoolIamPolicyMap) ToCaPoolIamPolicyMapOutput() CaPoolIamPolicyMapOutput {
	return i.ToCaPoolIamPolicyMapOutputWithContext(context.Background())
}

func (i CaPoolIamPolicyMap) ToCaPoolIamPolicyMapOutputWithContext(ctx context.Context) CaPoolIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamPolicyMapOutput)
}

type CaPoolIamPolicyOutput struct{ *pulumi.OutputState }

func (CaPoolIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamPolicy)(nil)).Elem()
}

func (o CaPoolIamPolicyOutput) ToCaPoolIamPolicyOutput() CaPoolIamPolicyOutput {
	return o
}

func (o CaPoolIamPolicyOutput) ToCaPoolIamPolicyOutputWithContext(ctx context.Context) CaPoolIamPolicyOutput {
	return o
}

type CaPoolIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (CaPoolIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamPolicy)(nil)).Elem()
}

func (o CaPoolIamPolicyArrayOutput) ToCaPoolIamPolicyArrayOutput() CaPoolIamPolicyArrayOutput {
	return o
}

func (o CaPoolIamPolicyArrayOutput) ToCaPoolIamPolicyArrayOutputWithContext(ctx context.Context) CaPoolIamPolicyArrayOutput {
	return o
}

func (o CaPoolIamPolicyArrayOutput) Index(i pulumi.IntInput) CaPoolIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CaPoolIamPolicy {
		return vs[0].([]*CaPoolIamPolicy)[vs[1].(int)]
	}).(CaPoolIamPolicyOutput)
}

type CaPoolIamPolicyMapOutput struct{ *pulumi.OutputState }

func (CaPoolIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamPolicy)(nil)).Elem()
}

func (o CaPoolIamPolicyMapOutput) ToCaPoolIamPolicyMapOutput() CaPoolIamPolicyMapOutput {
	return o
}

func (o CaPoolIamPolicyMapOutput) ToCaPoolIamPolicyMapOutputWithContext(ctx context.Context) CaPoolIamPolicyMapOutput {
	return o
}

func (o CaPoolIamPolicyMapOutput) MapIndex(k pulumi.StringInput) CaPoolIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CaPoolIamPolicy {
		return vs[0].(map[string]*CaPoolIamPolicy)[vs[1].(string)]
	}).(CaPoolIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamPolicyInput)(nil)).Elem(), &CaPoolIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamPolicyArrayInput)(nil)).Elem(), CaPoolIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamPolicyMapInput)(nil)).Elem(), CaPoolIamPolicyMap{})
	pulumi.RegisterOutputType(CaPoolIamPolicyOutput{})
	pulumi.RegisterOutputType(CaPoolIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(CaPoolIamPolicyMapOutput{})
}
