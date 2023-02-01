// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

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
// This member resource can be imported using the `project_id`, role, and member e.g.
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id roles/viewer user:foo@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `project_id` and role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id roles/viewer"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `project_id`.
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember my_project your-project-id
//
// ```
//
//	IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id foo.googleapis.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
//
// ```sh
//
//	$ pulumi import gcp:projects/iAMMember:IAMMember to include the title of condition, e.g. `google_project_iam_binding.my_project "{{your-project-id}} roles/{{role_id}} condition-title"`
//
// ```
type IAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource IAMMember
	err := ctx.RegisterResource("gcp:projects/iAMMember:IAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMMemberState, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	var resource IAMMember
	err := ctx.ReadResource("gcp:projects/iAMMember:IAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMMember resources.
type iammemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the project's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type IAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (IAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberState)(nil)).Elem()
}

type iammemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	Member    string              `pulumi:"member"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project string `pulumi:"project"`
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringInput
	// The role that should be applied. Only one
	// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (IAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberArgs)(nil)).Elem()
}

type IAMMemberInput interface {
	pulumi.Input

	ToIAMMemberOutput() IAMMemberOutput
	ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput
}

func (*IAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (i *IAMMember) ToIAMMemberOutput() IAMMemberOutput {
	return i.ToIAMMemberOutputWithContext(context.Background())
}

func (i *IAMMember) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberOutput)
}

// IAMMemberArrayInput is an input type that accepts IAMMemberArray and IAMMemberArrayOutput values.
// You can construct a concrete instance of `IAMMemberArrayInput` via:
//
//	IAMMemberArray{ IAMMemberArgs{...} }
type IAMMemberArrayInput interface {
	pulumi.Input

	ToIAMMemberArrayOutput() IAMMemberArrayOutput
	ToIAMMemberArrayOutputWithContext(context.Context) IAMMemberArrayOutput
}

type IAMMemberArray []IAMMemberInput

func (IAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (i IAMMemberArray) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return i.ToIAMMemberArrayOutputWithContext(context.Background())
}

func (i IAMMemberArray) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberArrayOutput)
}

// IAMMemberMapInput is an input type that accepts IAMMemberMap and IAMMemberMapOutput values.
// You can construct a concrete instance of `IAMMemberMapInput` via:
//
//	IAMMemberMap{ "key": IAMMemberArgs{...} }
type IAMMemberMapInput interface {
	pulumi.Input

	ToIAMMemberMapOutput() IAMMemberMapOutput
	ToIAMMemberMapOutputWithContext(context.Context) IAMMemberMapOutput
}

type IAMMemberMap map[string]IAMMemberInput

func (IAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (i IAMMemberMap) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return i.ToIAMMemberMapOutputWithContext(context.Background())
}

func (i IAMMemberMap) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberMapOutput)
}

type IAMMemberOutput struct{ *pulumi.OutputState }

func (IAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (o IAMMemberOutput) ToIAMMemberOutput() IAMMemberOutput {
	return o
}

func (o IAMMemberOutput) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IAMMemberOutput) Condition() IAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *IAMMember) IAMMemberConditionPtrOutput { return v.Condition }).(IAMMemberConditionPtrOutput)
}

// (Computed) The etag of the project's IAM policy.
func (o IAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The project id of the target project. This is not
// inferred from the provider.
func (o IAMMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o IAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type IAMMemberArrayOutput struct{ *pulumi.OutputState }

func (IAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) Index(i pulumi.IntInput) IAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].([]*IAMMember)[vs[1].(int)]
	}).(IAMMemberOutput)
}

type IAMMemberMapOutput struct{ *pulumi.OutputState }

func (IAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) MapIndex(k pulumi.StringInput) IAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].(map[string]*IAMMember)[vs[1].(string)]
	}).(IAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberInput)(nil)).Elem(), &IAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberArrayInput)(nil)).Elem(), IAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberMapInput)(nil)).Elem(), IAMMemberMap{})
	pulumi.RegisterOutputType(IAMMemberOutput{})
	pulumi.RegisterOutputType(IAMMemberArrayOutput{})
	pulumi.RegisterOutputType(IAMMemberMapOutput{})
}
