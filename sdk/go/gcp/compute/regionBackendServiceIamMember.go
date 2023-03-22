// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/backendServices/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine regionbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionBackendServiceIamMember:RegionBackendServiceIamMember editor "projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}} roles/compute.admin user:jane@example.com"
//
// ```
//
//	IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionBackendServiceIamMember:RegionBackendServiceIamMember editor "projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}} roles/compute.admin"
//
// ```
//
//	IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//
//	$ pulumi import gcp:compute/regionBackendServiceIamMember:RegionBackendServiceIamMember editor projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}}
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type RegionBackendServiceIamMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition RegionBackendServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRegionBackendServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewRegionBackendServiceIamMember(ctx *pulumi.Context,
	name string, args *RegionBackendServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*RegionBackendServiceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource RegionBackendServiceIamMember
	err := ctx.RegisterResource("gcp:compute/regionBackendServiceIamMember:RegionBackendServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionBackendServiceIamMember gets an existing RegionBackendServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionBackendServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionBackendServiceIamMemberState, opts ...pulumi.ResourceOption) (*RegionBackendServiceIamMember, error) {
	var resource RegionBackendServiceIamMember
	err := ctx.ReadResource("gcp:compute/regionBackendServiceIamMember:RegionBackendServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionBackendServiceIamMember resources.
type regionBackendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *RegionBackendServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RegionBackendServiceIamMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition RegionBackendServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RegionBackendServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceIamMemberState)(nil)).Elem()
}

type regionBackendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *RegionBackendServiceIamMemberCondition `pulumi:"condition"`
	Member    string                                  `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RegionBackendServiceIamMember resource.
type RegionBackendServiceIamMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition RegionBackendServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created backend service should reside.
	// If it is not provided, the provider region is used.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RegionBackendServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionBackendServiceIamMemberArgs)(nil)).Elem()
}

type RegionBackendServiceIamMemberInput interface {
	pulumi.Input

	ToRegionBackendServiceIamMemberOutput() RegionBackendServiceIamMemberOutput
	ToRegionBackendServiceIamMemberOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberOutput
}

func (*RegionBackendServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceIamMember)(nil)).Elem()
}

func (i *RegionBackendServiceIamMember) ToRegionBackendServiceIamMemberOutput() RegionBackendServiceIamMemberOutput {
	return i.ToRegionBackendServiceIamMemberOutputWithContext(context.Background())
}

func (i *RegionBackendServiceIamMember) ToRegionBackendServiceIamMemberOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceIamMemberOutput)
}

// RegionBackendServiceIamMemberArrayInput is an input type that accepts RegionBackendServiceIamMemberArray and RegionBackendServiceIamMemberArrayOutput values.
// You can construct a concrete instance of `RegionBackendServiceIamMemberArrayInput` via:
//
//	RegionBackendServiceIamMemberArray{ RegionBackendServiceIamMemberArgs{...} }
type RegionBackendServiceIamMemberArrayInput interface {
	pulumi.Input

	ToRegionBackendServiceIamMemberArrayOutput() RegionBackendServiceIamMemberArrayOutput
	ToRegionBackendServiceIamMemberArrayOutputWithContext(context.Context) RegionBackendServiceIamMemberArrayOutput
}

type RegionBackendServiceIamMemberArray []RegionBackendServiceIamMemberInput

func (RegionBackendServiceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionBackendServiceIamMember)(nil)).Elem()
}

func (i RegionBackendServiceIamMemberArray) ToRegionBackendServiceIamMemberArrayOutput() RegionBackendServiceIamMemberArrayOutput {
	return i.ToRegionBackendServiceIamMemberArrayOutputWithContext(context.Background())
}

func (i RegionBackendServiceIamMemberArray) ToRegionBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceIamMemberArrayOutput)
}

// RegionBackendServiceIamMemberMapInput is an input type that accepts RegionBackendServiceIamMemberMap and RegionBackendServiceIamMemberMapOutput values.
// You can construct a concrete instance of `RegionBackendServiceIamMemberMapInput` via:
//
//	RegionBackendServiceIamMemberMap{ "key": RegionBackendServiceIamMemberArgs{...} }
type RegionBackendServiceIamMemberMapInput interface {
	pulumi.Input

	ToRegionBackendServiceIamMemberMapOutput() RegionBackendServiceIamMemberMapOutput
	ToRegionBackendServiceIamMemberMapOutputWithContext(context.Context) RegionBackendServiceIamMemberMapOutput
}

type RegionBackendServiceIamMemberMap map[string]RegionBackendServiceIamMemberInput

func (RegionBackendServiceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionBackendServiceIamMember)(nil)).Elem()
}

func (i RegionBackendServiceIamMemberMap) ToRegionBackendServiceIamMemberMapOutput() RegionBackendServiceIamMemberMapOutput {
	return i.ToRegionBackendServiceIamMemberMapOutputWithContext(context.Background())
}

func (i RegionBackendServiceIamMemberMap) ToRegionBackendServiceIamMemberMapOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionBackendServiceIamMemberMapOutput)
}

type RegionBackendServiceIamMemberOutput struct{ *pulumi.OutputState }

func (RegionBackendServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionBackendServiceIamMember)(nil)).Elem()
}

func (o RegionBackendServiceIamMemberOutput) ToRegionBackendServiceIamMemberOutput() RegionBackendServiceIamMemberOutput {
	return o
}

func (o RegionBackendServiceIamMemberOutput) ToRegionBackendServiceIamMemberOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o RegionBackendServiceIamMemberOutput) Condition() RegionBackendServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) RegionBackendServiceIamMemberConditionPtrOutput {
		return v.Condition
	}).(RegionBackendServiceIamMemberConditionPtrOutput)
}

// (Computed) The etag of the IAM policy.
func (o RegionBackendServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RegionBackendServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o RegionBackendServiceIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (o RegionBackendServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The Region in which the created backend service should reside.
// If it is not provided, the provider region is used.
// Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (o RegionBackendServiceIamMemberOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `compute.RegionBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o RegionBackendServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionBackendServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type RegionBackendServiceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (RegionBackendServiceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionBackendServiceIamMember)(nil)).Elem()
}

func (o RegionBackendServiceIamMemberArrayOutput) ToRegionBackendServiceIamMemberArrayOutput() RegionBackendServiceIamMemberArrayOutput {
	return o
}

func (o RegionBackendServiceIamMemberArrayOutput) ToRegionBackendServiceIamMemberArrayOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberArrayOutput {
	return o
}

func (o RegionBackendServiceIamMemberArrayOutput) Index(i pulumi.IntInput) RegionBackendServiceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionBackendServiceIamMember {
		return vs[0].([]*RegionBackendServiceIamMember)[vs[1].(int)]
	}).(RegionBackendServiceIamMemberOutput)
}

type RegionBackendServiceIamMemberMapOutput struct{ *pulumi.OutputState }

func (RegionBackendServiceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionBackendServiceIamMember)(nil)).Elem()
}

func (o RegionBackendServiceIamMemberMapOutput) ToRegionBackendServiceIamMemberMapOutput() RegionBackendServiceIamMemberMapOutput {
	return o
}

func (o RegionBackendServiceIamMemberMapOutput) ToRegionBackendServiceIamMemberMapOutputWithContext(ctx context.Context) RegionBackendServiceIamMemberMapOutput {
	return o
}

func (o RegionBackendServiceIamMemberMapOutput) MapIndex(k pulumi.StringInput) RegionBackendServiceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionBackendServiceIamMember {
		return vs[0].(map[string]*RegionBackendServiceIamMember)[vs[1].(string)]
	}).(RegionBackendServiceIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionBackendServiceIamMemberInput)(nil)).Elem(), &RegionBackendServiceIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionBackendServiceIamMemberArrayInput)(nil)).Elem(), RegionBackendServiceIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionBackendServiceIamMemberMapInput)(nil)).Elem(), RegionBackendServiceIamMemberMap{})
	pulumi.RegisterOutputType(RegionBackendServiceIamMemberOutput{})
	pulumi.RegisterOutputType(RegionBackendServiceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(RegionBackendServiceIamMemberMapOutput{})
}
