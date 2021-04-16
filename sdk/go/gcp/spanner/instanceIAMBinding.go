// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:
//
// * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
//
// > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
//
// * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_spanner\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = spanner.NewInstanceIAMPolicy(ctx, "instance", &spanner.InstanceIAMPolicyArgs{
// 			Instance:   pulumi.String("your-instance-name"),
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
// ## google\_spanner\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewInstanceIAMBinding(ctx, "instance", &spanner.InstanceIAMBindingArgs{
// 			Instance: pulumi.String("your-instance-name"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/compute.networkUser"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_spanner\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.NewInstanceIAMMember(ctx, "instance", &spanner.InstanceIAMMemberArgs{
// 			Instance: pulumi.String("your-instance-name"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/compute.networkUser"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{name}} * {{name}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the account, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance "project-name/instance-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/instanceIAMBinding:InstanceIAMBinding instance project-name/instance-name
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIAMBinding struct {
	pulumi.CustomResourceState

	Condition InstanceIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the instance.
	Instance pulumi.StringOutput      `pulumi:"instance"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMBinding(ctx *pulumi.Context,
	name string, args *InstanceIAMBindingArgs, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
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
	var resource InstanceIAMBinding
	err := ctx.RegisterResource("gcp:spanner/instanceIAMBinding:InstanceIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMBinding gets an existing InstanceIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMBindingState, opts ...pulumi.ResourceOption) (*InstanceIAMBinding, error) {
	var resource InstanceIAMBinding
	err := ctx.ReadResource("gcp:spanner/instanceIAMBinding:InstanceIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMBinding resources.
type instanceIAMBindingState struct {
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the instance.
	Instance *string  `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIAMBindingState struct {
	Condition InstanceIAMBindingConditionPtrInput
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the instance.
	Instance pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingState)(nil)).Elem()
}

type instanceIAMBindingArgs struct {
	Condition *InstanceIAMBindingCondition `pulumi:"condition"`
	// The name of the instance.
	Instance string   `pulumi:"instance"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIAMBinding resource.
type InstanceIAMBindingArgs struct {
	Condition InstanceIAMBindingConditionPtrInput
	// The name of the instance.
	Instance pulumi.StringInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMBindingArgs)(nil)).Elem()
}

type InstanceIAMBindingInput interface {
	pulumi.Input

	ToInstanceIAMBindingOutput() InstanceIAMBindingOutput
	ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput
}

func (*InstanceIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMBinding)(nil))
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return i.ToInstanceIAMBindingOutputWithContext(context.Background())
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingOutput)
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingPtrOutput() InstanceIAMBindingPtrOutput {
	return i.ToInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (i *InstanceIAMBinding) ToInstanceIAMBindingPtrOutputWithContext(ctx context.Context) InstanceIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingPtrOutput)
}

type InstanceIAMBindingPtrInput interface {
	pulumi.Input

	ToInstanceIAMBindingPtrOutput() InstanceIAMBindingPtrOutput
	ToInstanceIAMBindingPtrOutputWithContext(ctx context.Context) InstanceIAMBindingPtrOutput
}

type instanceIAMBindingPtrType InstanceIAMBindingArgs

func (*instanceIAMBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMBinding)(nil))
}

func (i *instanceIAMBindingPtrType) ToInstanceIAMBindingPtrOutput() InstanceIAMBindingPtrOutput {
	return i.ToInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (i *instanceIAMBindingPtrType) ToInstanceIAMBindingPtrOutputWithContext(ctx context.Context) InstanceIAMBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingPtrOutput)
}

// InstanceIAMBindingArrayInput is an input type that accepts InstanceIAMBindingArray and InstanceIAMBindingArrayOutput values.
// You can construct a concrete instance of `InstanceIAMBindingArrayInput` via:
//
//          InstanceIAMBindingArray{ InstanceIAMBindingArgs{...} }
type InstanceIAMBindingArrayInput interface {
	pulumi.Input

	ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput
	ToInstanceIAMBindingArrayOutputWithContext(context.Context) InstanceIAMBindingArrayOutput
}

type InstanceIAMBindingArray []InstanceIAMBindingInput

func (InstanceIAMBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InstanceIAMBinding)(nil))
}

func (i InstanceIAMBindingArray) ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput {
	return i.ToInstanceIAMBindingArrayOutputWithContext(context.Background())
}

func (i InstanceIAMBindingArray) ToInstanceIAMBindingArrayOutputWithContext(ctx context.Context) InstanceIAMBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingArrayOutput)
}

// InstanceIAMBindingMapInput is an input type that accepts InstanceIAMBindingMap and InstanceIAMBindingMapOutput values.
// You can construct a concrete instance of `InstanceIAMBindingMapInput` via:
//
//          InstanceIAMBindingMap{ "key": InstanceIAMBindingArgs{...} }
type InstanceIAMBindingMapInput interface {
	pulumi.Input

	ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput
	ToInstanceIAMBindingMapOutputWithContext(context.Context) InstanceIAMBindingMapOutput
}

type InstanceIAMBindingMap map[string]InstanceIAMBindingInput

func (InstanceIAMBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InstanceIAMBinding)(nil))
}

func (i InstanceIAMBindingMap) ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput {
	return i.ToInstanceIAMBindingMapOutputWithContext(context.Background())
}

func (i InstanceIAMBindingMap) ToInstanceIAMBindingMapOutputWithContext(ctx context.Context) InstanceIAMBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMBindingMapOutput)
}

type InstanceIAMBindingOutput struct {
	*pulumi.OutputState
}

func (InstanceIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMBinding)(nil))
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutput() InstanceIAMBindingOutput {
	return o
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingOutputWithContext(ctx context.Context) InstanceIAMBindingOutput {
	return o
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingPtrOutput() InstanceIAMBindingPtrOutput {
	return o.ToInstanceIAMBindingPtrOutputWithContext(context.Background())
}

func (o InstanceIAMBindingOutput) ToInstanceIAMBindingPtrOutputWithContext(ctx context.Context) InstanceIAMBindingPtrOutput {
	return o.ApplyT(func(v InstanceIAMBinding) *InstanceIAMBinding {
		return &v
	}).(InstanceIAMBindingPtrOutput)
}

type InstanceIAMBindingPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceIAMBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMBinding)(nil))
}

func (o InstanceIAMBindingPtrOutput) ToInstanceIAMBindingPtrOutput() InstanceIAMBindingPtrOutput {
	return o
}

func (o InstanceIAMBindingPtrOutput) ToInstanceIAMBindingPtrOutputWithContext(ctx context.Context) InstanceIAMBindingPtrOutput {
	return o
}

type InstanceIAMBindingArrayOutput struct{ *pulumi.OutputState }

func (InstanceIAMBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIAMBinding)(nil))
}

func (o InstanceIAMBindingArrayOutput) ToInstanceIAMBindingArrayOutput() InstanceIAMBindingArrayOutput {
	return o
}

func (o InstanceIAMBindingArrayOutput) ToInstanceIAMBindingArrayOutputWithContext(ctx context.Context) InstanceIAMBindingArrayOutput {
	return o
}

func (o InstanceIAMBindingArrayOutput) Index(i pulumi.IntInput) InstanceIAMBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIAMBinding {
		return vs[0].([]InstanceIAMBinding)[vs[1].(int)]
	}).(InstanceIAMBindingOutput)
}

type InstanceIAMBindingMapOutput struct{ *pulumi.OutputState }

func (InstanceIAMBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceIAMBinding)(nil))
}

func (o InstanceIAMBindingMapOutput) ToInstanceIAMBindingMapOutput() InstanceIAMBindingMapOutput {
	return o
}

func (o InstanceIAMBindingMapOutput) ToInstanceIAMBindingMapOutputWithContext(ctx context.Context) InstanceIAMBindingMapOutput {
	return o
}

func (o InstanceIAMBindingMapOutput) MapIndex(k pulumi.StringInput) InstanceIAMBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceIAMBinding {
		return vs[0].(map[string]InstanceIAMBinding)[vs[1].(string)]
	}).(InstanceIAMBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceIAMBindingOutput{})
	pulumi.RegisterOutputType(InstanceIAMBindingPtrOutput{})
	pulumi.RegisterOutputType(InstanceIAMBindingArrayOutput{})
	pulumi.RegisterOutputType(InstanceIAMBindingMapOutput{})
}
