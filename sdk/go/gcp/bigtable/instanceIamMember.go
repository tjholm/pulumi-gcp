// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/bigtable.user",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigtable.NewInstanceIamPolicy(ctx, "editor", &bigtable.InstanceIamPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Instance:   pulumi.String("your-bigtable-instance"),
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
// ## google\_bigtable\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamBinding(ctx, "editor", &bigtable.InstanceIamBindingArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/bigtable.user"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigtable\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamMember(ctx, "editor", &bigtable.InstanceIamMemberArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/bigtable.user"),
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
// Instance IAM resources can be imported using the project, instance name, role and/or member.
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance}"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamMember struct {
	pulumi.CustomResourceState

	Condition InstanceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringOutput `pulumi:"instance"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIamMember registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamMember(ctx *pulumi.Context,
	name string, args *InstanceIamMemberArgs, opts ...pulumi.ResourceOption) (*InstanceIamMember, error) {
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
	var resource InstanceIamMember
	err := ctx.RegisterResource("gcp:bigtable/instanceIamMember:InstanceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamMember gets an existing InstanceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamMemberState, opts ...pulumi.ResourceOption) (*InstanceIamMember, error) {
	var resource InstanceIamMember
	err := ctx.ReadResource("gcp:bigtable/instanceIamMember:InstanceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamMember resources.
type instanceIamMemberState struct {
	Condition *InstanceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the instances's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance *string `pulumi:"instance"`
	Member   *string `pulumi:"member"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role *string `pulumi:"role"`
}

type InstanceIamMemberState struct {
	Condition InstanceIamMemberConditionPtrInput
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringPtrInput
}

func (InstanceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamMemberState)(nil)).Elem()
}

type instanceIamMemberArgs struct {
	Condition *InstanceIamMemberCondition `pulumi:"condition"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance string `pulumi:"instance"`
	Member   string `pulumi:"member"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIamMember resource.
type InstanceIamMemberArgs struct {
	Condition InstanceIamMemberConditionPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringInput
	Member   pulumi.StringInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	Role pulumi.StringInput
}

func (InstanceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamMemberArgs)(nil)).Elem()
}

type InstanceIamMemberInput interface {
	pulumi.Input

	ToInstanceIamMemberOutput() InstanceIamMemberOutput
	ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput
}

func (*InstanceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamMember)(nil))
}

func (i *InstanceIamMember) ToInstanceIamMemberOutput() InstanceIamMemberOutput {
	return i.ToInstanceIamMemberOutputWithContext(context.Background())
}

func (i *InstanceIamMember) ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberOutput)
}

func (i *InstanceIamMember) ToInstanceIamMemberPtrOutput() InstanceIamMemberPtrOutput {
	return i.ToInstanceIamMemberPtrOutputWithContext(context.Background())
}

func (i *InstanceIamMember) ToInstanceIamMemberPtrOutputWithContext(ctx context.Context) InstanceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberPtrOutput)
}

type InstanceIamMemberPtrInput interface {
	pulumi.Input

	ToInstanceIamMemberPtrOutput() InstanceIamMemberPtrOutput
	ToInstanceIamMemberPtrOutputWithContext(ctx context.Context) InstanceIamMemberPtrOutput
}

type instanceIamMemberPtrType InstanceIamMemberArgs

func (*instanceIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamMember)(nil))
}

func (i *instanceIamMemberPtrType) ToInstanceIamMemberPtrOutput() InstanceIamMemberPtrOutput {
	return i.ToInstanceIamMemberPtrOutputWithContext(context.Background())
}

func (i *instanceIamMemberPtrType) ToInstanceIamMemberPtrOutputWithContext(ctx context.Context) InstanceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberPtrOutput)
}

// InstanceIamMemberArrayInput is an input type that accepts InstanceIamMemberArray and InstanceIamMemberArrayOutput values.
// You can construct a concrete instance of `InstanceIamMemberArrayInput` via:
//
//          InstanceIamMemberArray{ InstanceIamMemberArgs{...} }
type InstanceIamMemberArrayInput interface {
	pulumi.Input

	ToInstanceIamMemberArrayOutput() InstanceIamMemberArrayOutput
	ToInstanceIamMemberArrayOutputWithContext(context.Context) InstanceIamMemberArrayOutput
}

type InstanceIamMemberArray []InstanceIamMemberInput

func (InstanceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InstanceIamMember)(nil))
}

func (i InstanceIamMemberArray) ToInstanceIamMemberArrayOutput() InstanceIamMemberArrayOutput {
	return i.ToInstanceIamMemberArrayOutputWithContext(context.Background())
}

func (i InstanceIamMemberArray) ToInstanceIamMemberArrayOutputWithContext(ctx context.Context) InstanceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberArrayOutput)
}

// InstanceIamMemberMapInput is an input type that accepts InstanceIamMemberMap and InstanceIamMemberMapOutput values.
// You can construct a concrete instance of `InstanceIamMemberMapInput` via:
//
//          InstanceIamMemberMap{ "key": InstanceIamMemberArgs{...} }
type InstanceIamMemberMapInput interface {
	pulumi.Input

	ToInstanceIamMemberMapOutput() InstanceIamMemberMapOutput
	ToInstanceIamMemberMapOutputWithContext(context.Context) InstanceIamMemberMapOutput
}

type InstanceIamMemberMap map[string]InstanceIamMemberInput

func (InstanceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InstanceIamMember)(nil))
}

func (i InstanceIamMemberMap) ToInstanceIamMemberMapOutput() InstanceIamMemberMapOutput {
	return i.ToInstanceIamMemberMapOutputWithContext(context.Background())
}

func (i InstanceIamMemberMap) ToInstanceIamMemberMapOutputWithContext(ctx context.Context) InstanceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberMapOutput)
}

type InstanceIamMemberOutput struct {
	*pulumi.OutputState
}

func (InstanceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamMember)(nil))
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberOutput() InstanceIamMemberOutput {
	return o
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput {
	return o
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberPtrOutput() InstanceIamMemberPtrOutput {
	return o.ToInstanceIamMemberPtrOutputWithContext(context.Background())
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberPtrOutputWithContext(ctx context.Context) InstanceIamMemberPtrOutput {
	return o.ApplyT(func(v InstanceIamMember) *InstanceIamMember {
		return &v
	}).(InstanceIamMemberPtrOutput)
}

type InstanceIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamMember)(nil))
}

func (o InstanceIamMemberPtrOutput) ToInstanceIamMemberPtrOutput() InstanceIamMemberPtrOutput {
	return o
}

func (o InstanceIamMemberPtrOutput) ToInstanceIamMemberPtrOutputWithContext(ctx context.Context) InstanceIamMemberPtrOutput {
	return o
}

type InstanceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (InstanceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIamMember)(nil))
}

func (o InstanceIamMemberArrayOutput) ToInstanceIamMemberArrayOutput() InstanceIamMemberArrayOutput {
	return o
}

func (o InstanceIamMemberArrayOutput) ToInstanceIamMemberArrayOutputWithContext(ctx context.Context) InstanceIamMemberArrayOutput {
	return o
}

func (o InstanceIamMemberArrayOutput) Index(i pulumi.IntInput) InstanceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIamMember {
		return vs[0].([]InstanceIamMember)[vs[1].(int)]
	}).(InstanceIamMemberOutput)
}

type InstanceIamMemberMapOutput struct{ *pulumi.OutputState }

func (InstanceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceIamMember)(nil))
}

func (o InstanceIamMemberMapOutput) ToInstanceIamMemberMapOutput() InstanceIamMemberMapOutput {
	return o
}

func (o InstanceIamMemberMapOutput) ToInstanceIamMemberMapOutputWithContext(ctx context.Context) InstanceIamMemberMapOutput {
	return o
}

func (o InstanceIamMemberMapOutput) MapIndex(k pulumi.StringInput) InstanceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceIamMember {
		return vs[0].(map[string]InstanceIamMember)[vs[1].(string)]
	}).(InstanceIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceIamMemberOutput{})
	pulumi.RegisterOutputType(InstanceIamMemberPtrOutput{})
	pulumi.RegisterOutputType(InstanceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(InstanceIamMemberMapOutput{})
}
