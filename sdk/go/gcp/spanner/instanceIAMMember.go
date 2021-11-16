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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
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
//  $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance project-name/instance-name
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIAMMember struct {
	pulumi.CustomResourceState

	Condition InstanceIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the instance.
	Instance pulumi.StringOutput `pulumi:"instance"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIAMMember registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMMember(ctx *pulumi.Context,
	name string, args *InstanceIAMMemberArgs, opts ...pulumi.ResourceOption) (*InstanceIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource InstanceIAMMember
	err := ctx.RegisterResource("gcp:spanner/instanceIAMMember:InstanceIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIAMMember gets an existing InstanceIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIAMMemberState, opts ...pulumi.ResourceOption) (*InstanceIAMMember, error) {
	var resource InstanceIAMMember
	err := ctx.ReadResource("gcp:spanner/instanceIAMMember:InstanceIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIAMMember resources.
type instanceIAMMemberState struct {
	Condition *InstanceIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the instance's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name of the instance.
	Instance *string `pulumi:"instance"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIAMMemberState struct {
	Condition InstanceIAMMemberConditionPtrInput
	// (Computed) The etag of the instance's IAM policy.
	Etag pulumi.StringPtrInput
	// The name of the instance.
	Instance pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMMemberState)(nil)).Elem()
}

type instanceIAMMemberArgs struct {
	Condition *InstanceIAMMemberCondition `pulumi:"condition"`
	// The name of the instance.
	Instance string `pulumi:"instance"`
	Member   string `pulumi:"member"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIAMMember resource.
type InstanceIAMMemberArgs struct {
	Condition InstanceIAMMemberConditionPtrInput
	// The name of the instance.
	Instance pulumi.StringInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIAMMemberArgs)(nil)).Elem()
}

type InstanceIAMMemberInput interface {
	pulumi.Input

	ToInstanceIAMMemberOutput() InstanceIAMMemberOutput
	ToInstanceIAMMemberOutputWithContext(ctx context.Context) InstanceIAMMemberOutput
}

func (*InstanceIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMMember)(nil))
}

func (i *InstanceIAMMember) ToInstanceIAMMemberOutput() InstanceIAMMemberOutput {
	return i.ToInstanceIAMMemberOutputWithContext(context.Background())
}

func (i *InstanceIAMMember) ToInstanceIAMMemberOutputWithContext(ctx context.Context) InstanceIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMMemberOutput)
}

func (i *InstanceIAMMember) ToInstanceIAMMemberPtrOutput() InstanceIAMMemberPtrOutput {
	return i.ToInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (i *InstanceIAMMember) ToInstanceIAMMemberPtrOutputWithContext(ctx context.Context) InstanceIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMMemberPtrOutput)
}

type InstanceIAMMemberPtrInput interface {
	pulumi.Input

	ToInstanceIAMMemberPtrOutput() InstanceIAMMemberPtrOutput
	ToInstanceIAMMemberPtrOutputWithContext(ctx context.Context) InstanceIAMMemberPtrOutput
}

type instanceIAMMemberPtrType InstanceIAMMemberArgs

func (*instanceIAMMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMMember)(nil))
}

func (i *instanceIAMMemberPtrType) ToInstanceIAMMemberPtrOutput() InstanceIAMMemberPtrOutput {
	return i.ToInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (i *instanceIAMMemberPtrType) ToInstanceIAMMemberPtrOutputWithContext(ctx context.Context) InstanceIAMMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMMemberPtrOutput)
}

// InstanceIAMMemberArrayInput is an input type that accepts InstanceIAMMemberArray and InstanceIAMMemberArrayOutput values.
// You can construct a concrete instance of `InstanceIAMMemberArrayInput` via:
//
//          InstanceIAMMemberArray{ InstanceIAMMemberArgs{...} }
type InstanceIAMMemberArrayInput interface {
	pulumi.Input

	ToInstanceIAMMemberArrayOutput() InstanceIAMMemberArrayOutput
	ToInstanceIAMMemberArrayOutputWithContext(context.Context) InstanceIAMMemberArrayOutput
}

type InstanceIAMMemberArray []InstanceIAMMemberInput

func (InstanceIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIAMMember)(nil)).Elem()
}

func (i InstanceIAMMemberArray) ToInstanceIAMMemberArrayOutput() InstanceIAMMemberArrayOutput {
	return i.ToInstanceIAMMemberArrayOutputWithContext(context.Background())
}

func (i InstanceIAMMemberArray) ToInstanceIAMMemberArrayOutputWithContext(ctx context.Context) InstanceIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMMemberArrayOutput)
}

// InstanceIAMMemberMapInput is an input type that accepts InstanceIAMMemberMap and InstanceIAMMemberMapOutput values.
// You can construct a concrete instance of `InstanceIAMMemberMapInput` via:
//
//          InstanceIAMMemberMap{ "key": InstanceIAMMemberArgs{...} }
type InstanceIAMMemberMapInput interface {
	pulumi.Input

	ToInstanceIAMMemberMapOutput() InstanceIAMMemberMapOutput
	ToInstanceIAMMemberMapOutputWithContext(context.Context) InstanceIAMMemberMapOutput
}

type InstanceIAMMemberMap map[string]InstanceIAMMemberInput

func (InstanceIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIAMMember)(nil)).Elem()
}

func (i InstanceIAMMemberMap) ToInstanceIAMMemberMapOutput() InstanceIAMMemberMapOutput {
	return i.ToInstanceIAMMemberMapOutputWithContext(context.Background())
}

func (i InstanceIAMMemberMap) ToInstanceIAMMemberMapOutputWithContext(ctx context.Context) InstanceIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIAMMemberMapOutput)
}

type InstanceIAMMemberOutput struct{ *pulumi.OutputState }

func (InstanceIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIAMMember)(nil))
}

func (o InstanceIAMMemberOutput) ToInstanceIAMMemberOutput() InstanceIAMMemberOutput {
	return o
}

func (o InstanceIAMMemberOutput) ToInstanceIAMMemberOutputWithContext(ctx context.Context) InstanceIAMMemberOutput {
	return o
}

func (o InstanceIAMMemberOutput) ToInstanceIAMMemberPtrOutput() InstanceIAMMemberPtrOutput {
	return o.ToInstanceIAMMemberPtrOutputWithContext(context.Background())
}

func (o InstanceIAMMemberOutput) ToInstanceIAMMemberPtrOutputWithContext(ctx context.Context) InstanceIAMMemberPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceIAMMember) *InstanceIAMMember {
		return &v
	}).(InstanceIAMMemberPtrOutput)
}

type InstanceIAMMemberPtrOutput struct{ *pulumi.OutputState }

func (InstanceIAMMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIAMMember)(nil))
}

func (o InstanceIAMMemberPtrOutput) ToInstanceIAMMemberPtrOutput() InstanceIAMMemberPtrOutput {
	return o
}

func (o InstanceIAMMemberPtrOutput) ToInstanceIAMMemberPtrOutputWithContext(ctx context.Context) InstanceIAMMemberPtrOutput {
	return o
}

func (o InstanceIAMMemberPtrOutput) Elem() InstanceIAMMemberOutput {
	return o.ApplyT(func(v *InstanceIAMMember) InstanceIAMMember {
		if v != nil {
			return *v
		}
		var ret InstanceIAMMember
		return ret
	}).(InstanceIAMMemberOutput)
}

type InstanceIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (InstanceIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIAMMember)(nil))
}

func (o InstanceIAMMemberArrayOutput) ToInstanceIAMMemberArrayOutput() InstanceIAMMemberArrayOutput {
	return o
}

func (o InstanceIAMMemberArrayOutput) ToInstanceIAMMemberArrayOutputWithContext(ctx context.Context) InstanceIAMMemberArrayOutput {
	return o
}

func (o InstanceIAMMemberArrayOutput) Index(i pulumi.IntInput) InstanceIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIAMMember {
		return vs[0].([]InstanceIAMMember)[vs[1].(int)]
	}).(InstanceIAMMemberOutput)
}

type InstanceIAMMemberMapOutput struct{ *pulumi.OutputState }

func (InstanceIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceIAMMember)(nil))
}

func (o InstanceIAMMemberMapOutput) ToInstanceIAMMemberMapOutput() InstanceIAMMemberMapOutput {
	return o
}

func (o InstanceIAMMemberMapOutput) ToInstanceIAMMemberMapOutputWithContext(ctx context.Context) InstanceIAMMemberMapOutput {
	return o
}

func (o InstanceIAMMemberMapOutput) MapIndex(k pulumi.StringInput) InstanceIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceIAMMember {
		return vs[0].(map[string]InstanceIAMMember)[vs[1].(string)]
	}).(InstanceIAMMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceIAMMemberOutput{})
	pulumi.RegisterOutputType(InstanceIAMMemberPtrOutput{})
	pulumi.RegisterOutputType(InstanceIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(InstanceIAMMemberMapOutput{})
}
