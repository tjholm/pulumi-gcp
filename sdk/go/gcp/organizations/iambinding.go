// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `org_id`, role, and member e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMBinding:IAMBinding my_organization "your-orgid roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `org_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMBinding:IAMBinding my_organization "your-org-id roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `org_id`.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMBinding:IAMBinding my_organization your-org-id
//
// ```
//
//	IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMBinding:IAMBinding my_organization "your-organization-id foo.googleapis.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
//
// ```sh
//
//	$ pulumi import gcp:organizations/iAMBinding:IAMBinding to include the title of condition, e.g. `google_organization_iam_binding.my_organization "your-org-id roles/{{role_id}} condition-title"`
//
// ```
type IAMBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewIAMBinding(ctx *pulumi.Context,
	name string, args *IAMBindingArgs, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource IAMBinding
	err := ctx.RegisterResource("gcp:organizations/iAMBinding:IAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMBinding gets an existing IAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMBindingState, opts ...pulumi.ResourceOption) (*IAMBinding, error) {
	var resource IAMBinding
	err := ctx.ReadResource("gcp:organizations/iAMBinding:IAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMBinding resources.
type iambindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the organization's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId *string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role *string `pulumi:"role"`
}

type IAMBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	// (Computed) The etag of the organization's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringPtrInput
}

func (IAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingState)(nil)).Elem()
}

type iambindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMBindingCondition `pulumi:"condition"`
	Members   []string             `pulumi:"members"`
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId string `pulumi:"orgId"`
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMBinding resource.
type IAMBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
	// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
	// will not be inferred from the provider.
	OrgId pulumi.StringInput
	// The role that should be applied. Only one
	// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `organizations/{{org_id}}/roles/{{role_id}}`.
	Role pulumi.StringInput
}

func (IAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iambindingArgs)(nil)).Elem()
}

type IAMBindingInput interface {
	pulumi.Input

	ToIAMBindingOutput() IAMBindingOutput
	ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput
}

func (*IAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMBinding)(nil)).Elem()
}

func (i *IAMBinding) ToIAMBindingOutput() IAMBindingOutput {
	return i.ToIAMBindingOutputWithContext(context.Background())
}

func (i *IAMBinding) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingOutput)
}

// IAMBindingArrayInput is an input type that accepts IAMBindingArray and IAMBindingArrayOutput values.
// You can construct a concrete instance of `IAMBindingArrayInput` via:
//
//	IAMBindingArray{ IAMBindingArgs{...} }
type IAMBindingArrayInput interface {
	pulumi.Input

	ToIAMBindingArrayOutput() IAMBindingArrayOutput
	ToIAMBindingArrayOutputWithContext(context.Context) IAMBindingArrayOutput
}

type IAMBindingArray []IAMBindingInput

func (IAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMBinding)(nil)).Elem()
}

func (i IAMBindingArray) ToIAMBindingArrayOutput() IAMBindingArrayOutput {
	return i.ToIAMBindingArrayOutputWithContext(context.Background())
}

func (i IAMBindingArray) ToIAMBindingArrayOutputWithContext(ctx context.Context) IAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingArrayOutput)
}

// IAMBindingMapInput is an input type that accepts IAMBindingMap and IAMBindingMapOutput values.
// You can construct a concrete instance of `IAMBindingMapInput` via:
//
//	IAMBindingMap{ "key": IAMBindingArgs{...} }
type IAMBindingMapInput interface {
	pulumi.Input

	ToIAMBindingMapOutput() IAMBindingMapOutput
	ToIAMBindingMapOutputWithContext(context.Context) IAMBindingMapOutput
}

type IAMBindingMap map[string]IAMBindingInput

func (IAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMBinding)(nil)).Elem()
}

func (i IAMBindingMap) ToIAMBindingMapOutput() IAMBindingMapOutput {
	return i.ToIAMBindingMapOutputWithContext(context.Background())
}

func (i IAMBindingMap) ToIAMBindingMapOutputWithContext(ctx context.Context) IAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMBindingMapOutput)
}

type IAMBindingOutput struct{ *pulumi.OutputState }

func (IAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMBinding)(nil)).Elem()
}

func (o IAMBindingOutput) ToIAMBindingOutput() IAMBindingOutput {
	return o
}

func (o IAMBindingOutput) ToIAMBindingOutputWithContext(ctx context.Context) IAMBindingOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IAMBindingOutput) Condition() IAMBindingConditionPtrOutput {
	return o.ApplyT(func(v *IAMBinding) IAMBindingConditionPtrOutput { return v.Condition }).(IAMBindingConditionPtrOutput)
}

// (Computed) The etag of the organization's IAM policy.
func (o IAMBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IAMBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IAMBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The organization ID. If not specified for `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
// Required for `organizations.IAMPolicy` - you must explicitly set the organization, and it
// will not be inferred from the provider.
func (o IAMBindingOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMBinding) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
// `organizations/{{org_id}}/roles/{{role_id}}`.
func (o IAMBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type IAMBindingArrayOutput struct{ *pulumi.OutputState }

func (IAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMBinding)(nil)).Elem()
}

func (o IAMBindingArrayOutput) ToIAMBindingArrayOutput() IAMBindingArrayOutput {
	return o
}

func (o IAMBindingArrayOutput) ToIAMBindingArrayOutputWithContext(ctx context.Context) IAMBindingArrayOutput {
	return o
}

func (o IAMBindingArrayOutput) Index(i pulumi.IntInput) IAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMBinding {
		return vs[0].([]*IAMBinding)[vs[1].(int)]
	}).(IAMBindingOutput)
}

type IAMBindingMapOutput struct{ *pulumi.OutputState }

func (IAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMBinding)(nil)).Elem()
}

func (o IAMBindingMapOutput) ToIAMBindingMapOutput() IAMBindingMapOutput {
	return o
}

func (o IAMBindingMapOutput) ToIAMBindingMapOutputWithContext(ctx context.Context) IAMBindingMapOutput {
	return o
}

func (o IAMBindingMapOutput) MapIndex(k pulumi.StringInput) IAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMBinding {
		return vs[0].(map[string]*IAMBinding)[vs[1].(string)]
	}).(IAMBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingInput)(nil)).Elem(), &IAMBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingArrayInput)(nil)).Elem(), IAMBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMBindingMapInput)(nil)).Elem(), IAMBindingMap{})
	pulumi.RegisterOutputType(IAMBindingOutput{})
	pulumi.RegisterOutputType(IAMBindingArrayOutput{})
	pulumi.RegisterOutputType(IAMBindingMapOutput{})
}
