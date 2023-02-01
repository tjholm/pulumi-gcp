// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
//
// * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigtable\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
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
//						Role: "roles/bigtable.user",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = bigtable.NewInstanceIamPolicy(ctx, "editor", &bigtable.InstanceIamPolicyArgs{
//				Project:    pulumi.String("your-project"),
//				Instance:   pulumi.String("your-bigtable-instance"),
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
// ## google\_bigtable\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewInstanceIamBinding(ctx, "editor", &bigtable.InstanceIamBindingArgs{
//				Instance: pulumi.String("your-bigtable-instance"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/bigtable.user"),
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
// ## google\_bigtable\_instance\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigtable"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigtable.NewInstanceIamMember(ctx, "editor", &bigtable.InstanceIamMemberArgs{
//				Instance: pulumi.String("your-bigtable-instance"),
//				Member:   pulumi.String("user:jane@example.com"),
//				Role:     pulumi.String("roles/bigtable.user"),
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
// Instance IAM resources can be imported using the project, instance name, role and/or member.
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance}"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance} roles/editor"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamBinding struct {
	pulumi.CustomResourceState

	Condition InstanceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	Project  pulumi.StringOutput      `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamBinding(ctx *pulumi.Context,
	name string, args *InstanceIamBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource InstanceIamBinding
	err := ctx.RegisterResource("gcp:bigtable/instanceIamBinding:InstanceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamBinding gets an existing InstanceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamBindingState, opts ...pulumi.ResourceOption) (*InstanceIamBinding, error) {
	var resource InstanceIamBinding
	err := ctx.ReadResource("gcp:bigtable/instanceIamBinding:InstanceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamBinding resources.
type instanceIamBindingState struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	Project  *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role *string `pulumi:"role"`
}

type InstanceIamBindingState struct {
	Condition InstanceIamBindingConditionPtrInput
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringPtrInput
}

func (InstanceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingState)(nil)).Elem()
}

type instanceIamBindingArgs struct {
	Condition *InstanceIamBindingCondition `pulumi:"condition"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	Project  *string  `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIamBinding resource.
type InstanceIamBindingArgs struct {
	Condition InstanceIamBindingConditionPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	Project  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringInput
}

func (InstanceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamBindingArgs)(nil)).Elem()
}

type InstanceIamBindingInput interface {
	pulumi.Input

	ToInstanceIamBindingOutput() InstanceIamBindingOutput
	ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput
}

func (*InstanceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamBinding)(nil)).Elem()
}

func (i *InstanceIamBinding) ToInstanceIamBindingOutput() InstanceIamBindingOutput {
	return i.ToInstanceIamBindingOutputWithContext(context.Background())
}

func (i *InstanceIamBinding) ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingOutput)
}

// InstanceIamBindingArrayInput is an input type that accepts InstanceIamBindingArray and InstanceIamBindingArrayOutput values.
// You can construct a concrete instance of `InstanceIamBindingArrayInput` via:
//
//	InstanceIamBindingArray{ InstanceIamBindingArgs{...} }
type InstanceIamBindingArrayInput interface {
	pulumi.Input

	ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput
	ToInstanceIamBindingArrayOutputWithContext(context.Context) InstanceIamBindingArrayOutput
}

type InstanceIamBindingArray []InstanceIamBindingInput

func (InstanceIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIamBinding)(nil)).Elem()
}

func (i InstanceIamBindingArray) ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput {
	return i.ToInstanceIamBindingArrayOutputWithContext(context.Background())
}

func (i InstanceIamBindingArray) ToInstanceIamBindingArrayOutputWithContext(ctx context.Context) InstanceIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingArrayOutput)
}

// InstanceIamBindingMapInput is an input type that accepts InstanceIamBindingMap and InstanceIamBindingMapOutput values.
// You can construct a concrete instance of `InstanceIamBindingMapInput` via:
//
//	InstanceIamBindingMap{ "key": InstanceIamBindingArgs{...} }
type InstanceIamBindingMapInput interface {
	pulumi.Input

	ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput
	ToInstanceIamBindingMapOutputWithContext(context.Context) InstanceIamBindingMapOutput
}

type InstanceIamBindingMap map[string]InstanceIamBindingInput

func (InstanceIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIamBinding)(nil)).Elem()
}

func (i InstanceIamBindingMap) ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput {
	return i.ToInstanceIamBindingMapOutputWithContext(context.Background())
}

func (i InstanceIamBindingMap) ToInstanceIamBindingMapOutputWithContext(ctx context.Context) InstanceIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamBindingMapOutput)
}

type InstanceIamBindingOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamBinding)(nil)).Elem()
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingOutput() InstanceIamBindingOutput {
	return o
}

func (o InstanceIamBindingOutput) ToInstanceIamBindingOutputWithContext(ctx context.Context) InstanceIamBindingOutput {
	return o
}

func (o InstanceIamBindingOutput) Condition() InstanceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *InstanceIamBinding) InstanceIamBindingConditionPtrOutput { return v.Condition }).(InstanceIamBindingConditionPtrOutput)
}

// (Computed) The etag of the instances's IAM policy.
func (o InstanceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name or relative resource id of the instance to manage IAM policies for.
func (o InstanceIamBindingOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIamBinding) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

func (o InstanceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o InstanceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
func (o InstanceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type InstanceIamBindingArrayOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIamBinding)(nil)).Elem()
}

func (o InstanceIamBindingArrayOutput) ToInstanceIamBindingArrayOutput() InstanceIamBindingArrayOutput {
	return o
}

func (o InstanceIamBindingArrayOutput) ToInstanceIamBindingArrayOutputWithContext(ctx context.Context) InstanceIamBindingArrayOutput {
	return o
}

func (o InstanceIamBindingArrayOutput) Index(i pulumi.IntInput) InstanceIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceIamBinding {
		return vs[0].([]*InstanceIamBinding)[vs[1].(int)]
	}).(InstanceIamBindingOutput)
}

type InstanceIamBindingMapOutput struct{ *pulumi.OutputState }

func (InstanceIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIamBinding)(nil)).Elem()
}

func (o InstanceIamBindingMapOutput) ToInstanceIamBindingMapOutput() InstanceIamBindingMapOutput {
	return o
}

func (o InstanceIamBindingMapOutput) ToInstanceIamBindingMapOutputWithContext(ctx context.Context) InstanceIamBindingMapOutput {
	return o
}

func (o InstanceIamBindingMapOutput) MapIndex(k pulumi.StringInput) InstanceIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceIamBinding {
		return vs[0].(map[string]*InstanceIamBinding)[vs[1].(string)]
	}).(InstanceIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIamBindingInput)(nil)).Elem(), &InstanceIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIamBindingArrayInput)(nil)).Elem(), InstanceIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIamBindingMapInput)(nil)).Elem(), InstanceIamBindingMap{})
	pulumi.RegisterOutputType(InstanceIamBindingOutput{})
	pulumi.RegisterOutputType(InstanceIamBindingArrayOutput{})
	pulumi.RegisterOutputType(InstanceIamBindingMapOutput{})
}
