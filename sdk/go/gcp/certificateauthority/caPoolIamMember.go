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
//  $ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type CaPoolIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringOutput               `pulumi:"caPool"`
	Condition CaPoolIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCaPoolIamMember registers a new resource with the given unique name, arguments, and options.
func NewCaPoolIamMember(ctx *pulumi.Context,
	name string, args *CaPoolIamMemberArgs, opts ...pulumi.ResourceOption) (*CaPoolIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource CaPoolIamMember
	err := ctx.RegisterResource("gcp:certificateauthority/caPoolIamMember:CaPoolIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPoolIamMember gets an existing CaPoolIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPoolIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolIamMemberState, opts ...pulumi.ResourceOption) (*CaPoolIamMember, error) {
	var resource CaPoolIamMember
	err := ctx.ReadResource("gcp:certificateauthority/caPoolIamMember:CaPoolIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPoolIamMember resources.
type caPoolIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    *string                   `pulumi:"caPool"`
	Condition *CaPoolIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CaPoolIamMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringPtrInput
	Condition CaPoolIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CaPoolIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamMemberState)(nil)).Elem()
}

type caPoolIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    string                    `pulumi:"caPool"`
	Condition *CaPoolIamMemberCondition `pulumi:"condition"`
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CaPoolIamMember resource.
type CaPoolIamMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CaPool    pulumi.StringInput
	Condition CaPoolIamMemberConditionPtrInput
	// Location of the CaPool. A full list of valid locations can be found by
	// running `gcloud privateca locations list`.
	// Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CaPoolIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolIamMemberArgs)(nil)).Elem()
}

type CaPoolIamMemberInput interface {
	pulumi.Input

	ToCaPoolIamMemberOutput() CaPoolIamMemberOutput
	ToCaPoolIamMemberOutputWithContext(ctx context.Context) CaPoolIamMemberOutput
}

func (*CaPoolIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamMember)(nil))
}

func (i *CaPoolIamMember) ToCaPoolIamMemberOutput() CaPoolIamMemberOutput {
	return i.ToCaPoolIamMemberOutputWithContext(context.Background())
}

func (i *CaPoolIamMember) ToCaPoolIamMemberOutputWithContext(ctx context.Context) CaPoolIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberOutput)
}

func (i *CaPoolIamMember) ToCaPoolIamMemberPtrOutput() CaPoolIamMemberPtrOutput {
	return i.ToCaPoolIamMemberPtrOutputWithContext(context.Background())
}

func (i *CaPoolIamMember) ToCaPoolIamMemberPtrOutputWithContext(ctx context.Context) CaPoolIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberPtrOutput)
}

type CaPoolIamMemberPtrInput interface {
	pulumi.Input

	ToCaPoolIamMemberPtrOutput() CaPoolIamMemberPtrOutput
	ToCaPoolIamMemberPtrOutputWithContext(ctx context.Context) CaPoolIamMemberPtrOutput
}

type caPoolIamMemberPtrType CaPoolIamMemberArgs

func (*caPoolIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamMember)(nil))
}

func (i *caPoolIamMemberPtrType) ToCaPoolIamMemberPtrOutput() CaPoolIamMemberPtrOutput {
	return i.ToCaPoolIamMemberPtrOutputWithContext(context.Background())
}

func (i *caPoolIamMemberPtrType) ToCaPoolIamMemberPtrOutputWithContext(ctx context.Context) CaPoolIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberPtrOutput)
}

// CaPoolIamMemberArrayInput is an input type that accepts CaPoolIamMemberArray and CaPoolIamMemberArrayOutput values.
// You can construct a concrete instance of `CaPoolIamMemberArrayInput` via:
//
//          CaPoolIamMemberArray{ CaPoolIamMemberArgs{...} }
type CaPoolIamMemberArrayInput interface {
	pulumi.Input

	ToCaPoolIamMemberArrayOutput() CaPoolIamMemberArrayOutput
	ToCaPoolIamMemberArrayOutputWithContext(context.Context) CaPoolIamMemberArrayOutput
}

type CaPoolIamMemberArray []CaPoolIamMemberInput

func (CaPoolIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamMember)(nil)).Elem()
}

func (i CaPoolIamMemberArray) ToCaPoolIamMemberArrayOutput() CaPoolIamMemberArrayOutput {
	return i.ToCaPoolIamMemberArrayOutputWithContext(context.Background())
}

func (i CaPoolIamMemberArray) ToCaPoolIamMemberArrayOutputWithContext(ctx context.Context) CaPoolIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberArrayOutput)
}

// CaPoolIamMemberMapInput is an input type that accepts CaPoolIamMemberMap and CaPoolIamMemberMapOutput values.
// You can construct a concrete instance of `CaPoolIamMemberMapInput` via:
//
//          CaPoolIamMemberMap{ "key": CaPoolIamMemberArgs{...} }
type CaPoolIamMemberMapInput interface {
	pulumi.Input

	ToCaPoolIamMemberMapOutput() CaPoolIamMemberMapOutput
	ToCaPoolIamMemberMapOutputWithContext(context.Context) CaPoolIamMemberMapOutput
}

type CaPoolIamMemberMap map[string]CaPoolIamMemberInput

func (CaPoolIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamMember)(nil)).Elem()
}

func (i CaPoolIamMemberMap) ToCaPoolIamMemberMapOutput() CaPoolIamMemberMapOutput {
	return i.ToCaPoolIamMemberMapOutputWithContext(context.Background())
}

func (i CaPoolIamMemberMap) ToCaPoolIamMemberMapOutputWithContext(ctx context.Context) CaPoolIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberMapOutput)
}

type CaPoolIamMemberOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CaPoolIamMember)(nil))
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberOutput() CaPoolIamMemberOutput {
	return o
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberOutputWithContext(ctx context.Context) CaPoolIamMemberOutput {
	return o
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberPtrOutput() CaPoolIamMemberPtrOutput {
	return o.ToCaPoolIamMemberPtrOutputWithContext(context.Background())
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberPtrOutputWithContext(ctx context.Context) CaPoolIamMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CaPoolIamMember) *CaPoolIamMember {
		return &v
	}).(CaPoolIamMemberPtrOutput)
}

type CaPoolIamMemberPtrOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPoolIamMember)(nil))
}

func (o CaPoolIamMemberPtrOutput) ToCaPoolIamMemberPtrOutput() CaPoolIamMemberPtrOutput {
	return o
}

func (o CaPoolIamMemberPtrOutput) ToCaPoolIamMemberPtrOutputWithContext(ctx context.Context) CaPoolIamMemberPtrOutput {
	return o
}

func (o CaPoolIamMemberPtrOutput) Elem() CaPoolIamMemberOutput {
	return o.ApplyT(func(v *CaPoolIamMember) CaPoolIamMember {
		if v != nil {
			return *v
		}
		var ret CaPoolIamMember
		return ret
	}).(CaPoolIamMemberOutput)
}

type CaPoolIamMemberArrayOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CaPoolIamMember)(nil))
}

func (o CaPoolIamMemberArrayOutput) ToCaPoolIamMemberArrayOutput() CaPoolIamMemberArrayOutput {
	return o
}

func (o CaPoolIamMemberArrayOutput) ToCaPoolIamMemberArrayOutputWithContext(ctx context.Context) CaPoolIamMemberArrayOutput {
	return o
}

func (o CaPoolIamMemberArrayOutput) Index(i pulumi.IntInput) CaPoolIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CaPoolIamMember {
		return vs[0].([]CaPoolIamMember)[vs[1].(int)]
	}).(CaPoolIamMemberOutput)
}

type CaPoolIamMemberMapOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CaPoolIamMember)(nil))
}

func (o CaPoolIamMemberMapOutput) ToCaPoolIamMemberMapOutput() CaPoolIamMemberMapOutput {
	return o
}

func (o CaPoolIamMemberMapOutput) ToCaPoolIamMemberMapOutputWithContext(ctx context.Context) CaPoolIamMemberMapOutput {
	return o
}

func (o CaPoolIamMemberMapOutput) MapIndex(k pulumi.StringInput) CaPoolIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CaPoolIamMember {
		return vs[0].(map[string]CaPoolIamMember)[vs[1].(string)]
	}).(CaPoolIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(CaPoolIamMemberOutput{})
	pulumi.RegisterOutputType(CaPoolIamMemberPtrOutput{})
	pulumi.RegisterOutputType(CaPoolIamMemberArrayOutput{})
	pulumi.RegisterOutputType(CaPoolIamMemberMapOutput{})
}
