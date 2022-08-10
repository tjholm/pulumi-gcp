// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificateauthority

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/caPools/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} Any variables not passed in the import command will be taken from the provider configuration. Certificate Authority Service capool IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor "projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}} roles/privateca.certificateManager"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:certificateauthority/caPoolIamMember:CaPoolIamMember editor projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type CaPoolIamMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
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
	CaPool *string `pulumi:"caPool"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
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
	CaPool pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
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
	CaPool string `pulumi:"caPool"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
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
	CaPool pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
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
	return reflect.TypeOf((**CaPoolIamMember)(nil)).Elem()
}

func (i *CaPoolIamMember) ToCaPoolIamMemberOutput() CaPoolIamMemberOutput {
	return i.ToCaPoolIamMemberOutputWithContext(context.Background())
}

func (i *CaPoolIamMember) ToCaPoolIamMemberOutputWithContext(ctx context.Context) CaPoolIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolIamMemberOutput)
}

// CaPoolIamMemberArrayInput is an input type that accepts CaPoolIamMemberArray and CaPoolIamMemberArrayOutput values.
// You can construct a concrete instance of `CaPoolIamMemberArrayInput` via:
//
//	CaPoolIamMemberArray{ CaPoolIamMemberArgs{...} }
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
//	CaPoolIamMemberMap{ "key": CaPoolIamMemberArgs{...} }
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
	return reflect.TypeOf((**CaPoolIamMember)(nil)).Elem()
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberOutput() CaPoolIamMemberOutput {
	return o
}

func (o CaPoolIamMemberOutput) ToCaPoolIamMemberOutputWithContext(ctx context.Context) CaPoolIamMemberOutput {
	return o
}

// Used to find the parent resource to bind the IAM policy to
func (o CaPoolIamMemberOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o CaPoolIamMemberOutput) Condition() CaPoolIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *CaPoolIamMember) CaPoolIamMemberConditionPtrOutput { return v.Condition }).(CaPoolIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o CaPoolIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Location of the CaPool. A full list of valid locations can be found by
// running `gcloud privateca locations list`.
// Used to find the parent resource to bind the IAM policy to
func (o CaPoolIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CaPoolIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o CaPoolIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `certificateauthority.CaPoolIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o CaPoolIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CaPoolIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type CaPoolIamMemberArrayOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CaPoolIamMember)(nil)).Elem()
}

func (o CaPoolIamMemberArrayOutput) ToCaPoolIamMemberArrayOutput() CaPoolIamMemberArrayOutput {
	return o
}

func (o CaPoolIamMemberArrayOutput) ToCaPoolIamMemberArrayOutputWithContext(ctx context.Context) CaPoolIamMemberArrayOutput {
	return o
}

func (o CaPoolIamMemberArrayOutput) Index(i pulumi.IntInput) CaPoolIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CaPoolIamMember {
		return vs[0].([]*CaPoolIamMember)[vs[1].(int)]
	}).(CaPoolIamMemberOutput)
}

type CaPoolIamMemberMapOutput struct{ *pulumi.OutputState }

func (CaPoolIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CaPoolIamMember)(nil)).Elem()
}

func (o CaPoolIamMemberMapOutput) ToCaPoolIamMemberMapOutput() CaPoolIamMemberMapOutput {
	return o
}

func (o CaPoolIamMemberMapOutput) ToCaPoolIamMemberMapOutputWithContext(ctx context.Context) CaPoolIamMemberMapOutput {
	return o
}

func (o CaPoolIamMemberMapOutput) MapIndex(k pulumi.StringInput) CaPoolIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CaPoolIamMember {
		return vs[0].(map[string]*CaPoolIamMember)[vs[1].(string)]
	}).(CaPoolIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamMemberInput)(nil)).Elem(), &CaPoolIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamMemberArrayInput)(nil)).Elem(), CaPoolIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolIamMemberMapInput)(nil)).Elem(), CaPoolIamMemberMap{})
	pulumi.RegisterOutputType(CaPoolIamMemberOutput{})
	pulumi.RegisterOutputType(CaPoolIamMemberArrayOutput{})
	pulumi.RegisterOutputType(CaPoolIamMemberMapOutput{})
}
