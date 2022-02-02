// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
//
// * `pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
// * `pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
// * `pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
//
// > **Note:** `pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `pubsub.SubscriptionIAMBinding` and `pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_pubsub\_subscription\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
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
// 		_, err = pubsub.NewSubscriptionIAMPolicy(ctx, "editor", &pubsub.SubscriptionIAMPolicyArgs{
// 			Subscription: pulumi.String("your-subscription-name"),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewSubscriptionIAMBinding(ctx, "editor", &pubsub.SubscriptionIAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role:         pulumi.String("roles/editor"),
// 			Subscription: pulumi.String("your-subscription-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_pubsub\_subscription\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewSubscriptionIAMMember(ctx, "editor", &pubsub.SubscriptionIAMMemberArgs{
// 			Member:       pulumi.String("user:jane@example.com"),
// 			Role:         pulumi.String("roles/editor"),
// 			Subscription: pulumi.String("your-subscription-name"),
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
// Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor projects/{your-project-id}/subscriptions/{your-subscription-name}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SubscriptionIAMMember struct {
	pulumi.CustomResourceState

	Condition SubscriptionIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringOutput `pulumi:"subscription"`
}

// NewSubscriptionIAMMember registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMMember(ctx *pulumi.Context,
	name string, args *SubscriptionIAMMemberArgs, opts ...pulumi.ResourceOption) (*SubscriptionIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Subscription == nil {
		return nil, errors.New("invalid value for required argument 'Subscription'")
	}
	var resource SubscriptionIAMMember
	err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionIAMMember gets an existing SubscriptionIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionIAMMemberState, opts ...pulumi.ResourceOption) (*SubscriptionIAMMember, error) {
	var resource SubscriptionIAMMember
	err := ctx.ReadResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionIAMMember resources.
type subscriptionIAMMemberState struct {
	Condition *SubscriptionIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the subscription's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription *string `pulumi:"subscription"`
}

type SubscriptionIAMMemberState struct {
	Condition SubscriptionIAMMemberConditionPtrInput
	// (Computed) The etag of the subscription's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringPtrInput
}

func (SubscriptionIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMMemberState)(nil)).Elem()
}

type subscriptionIAMMemberArgs struct {
	Condition *SubscriptionIAMMemberCondition `pulumi:"condition"`
	Member    string                          `pulumi:"member"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The subscription name or id to bind to attach IAM policy to.
	Subscription string `pulumi:"subscription"`
}

// The set of arguments for constructing a SubscriptionIAMMember resource.
type SubscriptionIAMMemberArgs struct {
	Condition SubscriptionIAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The subscription name or id to bind to attach IAM policy to.
	Subscription pulumi.StringInput
}

func (SubscriptionIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMMemberArgs)(nil)).Elem()
}

type SubscriptionIAMMemberInput interface {
	pulumi.Input

	ToSubscriptionIAMMemberOutput() SubscriptionIAMMemberOutput
	ToSubscriptionIAMMemberOutputWithContext(ctx context.Context) SubscriptionIAMMemberOutput
}

func (*SubscriptionIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMMember)(nil)).Elem()
}

func (i *SubscriptionIAMMember) ToSubscriptionIAMMemberOutput() SubscriptionIAMMemberOutput {
	return i.ToSubscriptionIAMMemberOutputWithContext(context.Background())
}

func (i *SubscriptionIAMMember) ToSubscriptionIAMMemberOutputWithContext(ctx context.Context) SubscriptionIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMMemberOutput)
}

// SubscriptionIAMMemberArrayInput is an input type that accepts SubscriptionIAMMemberArray and SubscriptionIAMMemberArrayOutput values.
// You can construct a concrete instance of `SubscriptionIAMMemberArrayInput` via:
//
//          SubscriptionIAMMemberArray{ SubscriptionIAMMemberArgs{...} }
type SubscriptionIAMMemberArrayInput interface {
	pulumi.Input

	ToSubscriptionIAMMemberArrayOutput() SubscriptionIAMMemberArrayOutput
	ToSubscriptionIAMMemberArrayOutputWithContext(context.Context) SubscriptionIAMMemberArrayOutput
}

type SubscriptionIAMMemberArray []SubscriptionIAMMemberInput

func (SubscriptionIAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionIAMMember)(nil)).Elem()
}

func (i SubscriptionIAMMemberArray) ToSubscriptionIAMMemberArrayOutput() SubscriptionIAMMemberArrayOutput {
	return i.ToSubscriptionIAMMemberArrayOutputWithContext(context.Background())
}

func (i SubscriptionIAMMemberArray) ToSubscriptionIAMMemberArrayOutputWithContext(ctx context.Context) SubscriptionIAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMMemberArrayOutput)
}

// SubscriptionIAMMemberMapInput is an input type that accepts SubscriptionIAMMemberMap and SubscriptionIAMMemberMapOutput values.
// You can construct a concrete instance of `SubscriptionIAMMemberMapInput` via:
//
//          SubscriptionIAMMemberMap{ "key": SubscriptionIAMMemberArgs{...} }
type SubscriptionIAMMemberMapInput interface {
	pulumi.Input

	ToSubscriptionIAMMemberMapOutput() SubscriptionIAMMemberMapOutput
	ToSubscriptionIAMMemberMapOutputWithContext(context.Context) SubscriptionIAMMemberMapOutput
}

type SubscriptionIAMMemberMap map[string]SubscriptionIAMMemberInput

func (SubscriptionIAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionIAMMember)(nil)).Elem()
}

func (i SubscriptionIAMMemberMap) ToSubscriptionIAMMemberMapOutput() SubscriptionIAMMemberMapOutput {
	return i.ToSubscriptionIAMMemberMapOutputWithContext(context.Background())
}

func (i SubscriptionIAMMemberMap) ToSubscriptionIAMMemberMapOutputWithContext(ctx context.Context) SubscriptionIAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMMemberMapOutput)
}

type SubscriptionIAMMemberOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMMember)(nil)).Elem()
}

func (o SubscriptionIAMMemberOutput) ToSubscriptionIAMMemberOutput() SubscriptionIAMMemberOutput {
	return o
}

func (o SubscriptionIAMMemberOutput) ToSubscriptionIAMMemberOutputWithContext(ctx context.Context) SubscriptionIAMMemberOutput {
	return o
}

type SubscriptionIAMMemberArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionIAMMember)(nil)).Elem()
}

func (o SubscriptionIAMMemberArrayOutput) ToSubscriptionIAMMemberArrayOutput() SubscriptionIAMMemberArrayOutput {
	return o
}

func (o SubscriptionIAMMemberArrayOutput) ToSubscriptionIAMMemberArrayOutputWithContext(ctx context.Context) SubscriptionIAMMemberArrayOutput {
	return o
}

func (o SubscriptionIAMMemberArrayOutput) Index(i pulumi.IntInput) SubscriptionIAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscriptionIAMMember {
		return vs[0].([]*SubscriptionIAMMember)[vs[1].(int)]
	}).(SubscriptionIAMMemberOutput)
}

type SubscriptionIAMMemberMapOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionIAMMember)(nil)).Elem()
}

func (o SubscriptionIAMMemberMapOutput) ToSubscriptionIAMMemberMapOutput() SubscriptionIAMMemberMapOutput {
	return o
}

func (o SubscriptionIAMMemberMapOutput) ToSubscriptionIAMMemberMapOutputWithContext(ctx context.Context) SubscriptionIAMMemberMapOutput {
	return o
}

func (o SubscriptionIAMMemberMapOutput) MapIndex(k pulumi.StringInput) SubscriptionIAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscriptionIAMMember {
		return vs[0].(map[string]*SubscriptionIAMMember)[vs[1].(string)]
	}).(SubscriptionIAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMMemberInput)(nil)).Elem(), &SubscriptionIAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMMemberArrayInput)(nil)).Elem(), SubscriptionIAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMMemberMapInput)(nil)).Elem(), SubscriptionIAMMemberMap{})
	pulumi.RegisterOutputType(SubscriptionIAMMemberOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMMemberArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMMemberMapOutput{})
}
