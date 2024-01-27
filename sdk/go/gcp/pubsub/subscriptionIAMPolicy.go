// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
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
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = pubsub.NewSubscriptionIAMPolicy(ctx, "editor", &pubsub.SubscriptionIAMPolicyArgs{
//				Subscription: pulumi.String("your-subscription-name"),
//				PolicyData:   *pulumi.String(admin.PolicyData),
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
// ## google\_pubsub\_subscription\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSubscriptionIAMBinding(ctx, "editor", &pubsub.SubscriptionIAMBindingArgs{
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role:         pulumi.String("roles/editor"),
//				Subscription: pulumi.String("your-subscription-name"),
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
// ## google\_pubsub\_subscription\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pubsub.NewSubscriptionIAMMember(ctx, "editor", &pubsub.SubscriptionIAMMemberArgs{
//				Member:       pulumi.String("user:jane@example.com"),
//				Role:         pulumi.String("roles/editor"),
//				Subscription: pulumi.String("your-subscription-name"),
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
// ### Importing IAM policies IAM policy imports use the identifier of the Pubsub Subscription resource. For example* `"projects/{{project_id}}/subscriptions/{{subscription}}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
//
//	id = "projects/{{project_id}}/subscriptions/{{subscription}}"
//
//	to = google_pubsub_subscription_iam_policy.default } The `pulumi import` command can also be used
//
// ```sh
//
//	$ pulumi import gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy default projects/{{project_id}}/subscriptions/{{subscription}}
//
// ```
type SubscriptionIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the subscription's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Subscription pulumi.StringOutput `pulumi:"subscription"`
}

// NewSubscriptionIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionIAMPolicy(ctx *pulumi.Context,
	name string, args *SubscriptionIAMPolicyArgs, opts ...pulumi.ResourceOption) (*SubscriptionIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Subscription == nil {
		return nil, errors.New("invalid value for required argument 'Subscription'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SubscriptionIAMPolicy
	err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionIAMPolicy gets an existing SubscriptionIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionIAMPolicyState, opts ...pulumi.ResourceOption) (*SubscriptionIAMPolicy, error) {
	var resource SubscriptionIAMPolicy
	err := ctx.ReadResource("gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionIAMPolicy resources.
type subscriptionIAMPolicyState struct {
	// (Computed) The etag of the subscription's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData *string `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Subscription *string `pulumi:"subscription"`
}

type SubscriptionIAMPolicyState struct {
	// (Computed) The etag of the subscription's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Subscription pulumi.StringPtrInput
}

func (SubscriptionIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMPolicyState)(nil)).Elem()
}

type subscriptionIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData string `pulumi:"policyData"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The subscription name or id to bind to attach IAM policy to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Subscription string `pulumi:"subscription"`
}

// The set of arguments for constructing a SubscriptionIAMPolicy resource.
type SubscriptionIAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The subscription name or id to bind to attach IAM policy to.
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Subscription pulumi.StringInput
}

func (SubscriptionIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionIAMPolicyArgs)(nil)).Elem()
}

type SubscriptionIAMPolicyInput interface {
	pulumi.Input

	ToSubscriptionIAMPolicyOutput() SubscriptionIAMPolicyOutput
	ToSubscriptionIAMPolicyOutputWithContext(ctx context.Context) SubscriptionIAMPolicyOutput
}

func (*SubscriptionIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMPolicy)(nil)).Elem()
}

func (i *SubscriptionIAMPolicy) ToSubscriptionIAMPolicyOutput() SubscriptionIAMPolicyOutput {
	return i.ToSubscriptionIAMPolicyOutputWithContext(context.Background())
}

func (i *SubscriptionIAMPolicy) ToSubscriptionIAMPolicyOutputWithContext(ctx context.Context) SubscriptionIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMPolicyOutput)
}

// SubscriptionIAMPolicyArrayInput is an input type that accepts SubscriptionIAMPolicyArray and SubscriptionIAMPolicyArrayOutput values.
// You can construct a concrete instance of `SubscriptionIAMPolicyArrayInput` via:
//
//	SubscriptionIAMPolicyArray{ SubscriptionIAMPolicyArgs{...} }
type SubscriptionIAMPolicyArrayInput interface {
	pulumi.Input

	ToSubscriptionIAMPolicyArrayOutput() SubscriptionIAMPolicyArrayOutput
	ToSubscriptionIAMPolicyArrayOutputWithContext(context.Context) SubscriptionIAMPolicyArrayOutput
}

type SubscriptionIAMPolicyArray []SubscriptionIAMPolicyInput

func (SubscriptionIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionIAMPolicy)(nil)).Elem()
}

func (i SubscriptionIAMPolicyArray) ToSubscriptionIAMPolicyArrayOutput() SubscriptionIAMPolicyArrayOutput {
	return i.ToSubscriptionIAMPolicyArrayOutputWithContext(context.Background())
}

func (i SubscriptionIAMPolicyArray) ToSubscriptionIAMPolicyArrayOutputWithContext(ctx context.Context) SubscriptionIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMPolicyArrayOutput)
}

// SubscriptionIAMPolicyMapInput is an input type that accepts SubscriptionIAMPolicyMap and SubscriptionIAMPolicyMapOutput values.
// You can construct a concrete instance of `SubscriptionIAMPolicyMapInput` via:
//
//	SubscriptionIAMPolicyMap{ "key": SubscriptionIAMPolicyArgs{...} }
type SubscriptionIAMPolicyMapInput interface {
	pulumi.Input

	ToSubscriptionIAMPolicyMapOutput() SubscriptionIAMPolicyMapOutput
	ToSubscriptionIAMPolicyMapOutputWithContext(context.Context) SubscriptionIAMPolicyMapOutput
}

type SubscriptionIAMPolicyMap map[string]SubscriptionIAMPolicyInput

func (SubscriptionIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionIAMPolicy)(nil)).Elem()
}

func (i SubscriptionIAMPolicyMap) ToSubscriptionIAMPolicyMapOutput() SubscriptionIAMPolicyMapOutput {
	return i.ToSubscriptionIAMPolicyMapOutputWithContext(context.Background())
}

func (i SubscriptionIAMPolicyMap) ToSubscriptionIAMPolicyMapOutputWithContext(ctx context.Context) SubscriptionIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionIAMPolicyMapOutput)
}

type SubscriptionIAMPolicyOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionIAMPolicy)(nil)).Elem()
}

func (o SubscriptionIAMPolicyOutput) ToSubscriptionIAMPolicyOutput() SubscriptionIAMPolicyOutput {
	return o
}

func (o SubscriptionIAMPolicyOutput) ToSubscriptionIAMPolicyOutputWithContext(ctx context.Context) SubscriptionIAMPolicyOutput {
	return o
}

// (Computed) The etag of the subscription's IAM policy.
func (o SubscriptionIAMPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionIAMPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
//
// ***
func (o SubscriptionIAMPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionIAMPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o SubscriptionIAMPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionIAMPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The subscription name or id to bind to attach IAM policy to.
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
//   - **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o SubscriptionIAMPolicyOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscriptionIAMPolicy) pulumi.StringOutput { return v.Subscription }).(pulumi.StringOutput)
}

type SubscriptionIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscriptionIAMPolicy)(nil)).Elem()
}

func (o SubscriptionIAMPolicyArrayOutput) ToSubscriptionIAMPolicyArrayOutput() SubscriptionIAMPolicyArrayOutput {
	return o
}

func (o SubscriptionIAMPolicyArrayOutput) ToSubscriptionIAMPolicyArrayOutputWithContext(ctx context.Context) SubscriptionIAMPolicyArrayOutput {
	return o
}

func (o SubscriptionIAMPolicyArrayOutput) Index(i pulumi.IntInput) SubscriptionIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscriptionIAMPolicy {
		return vs[0].([]*SubscriptionIAMPolicy)[vs[1].(int)]
	}).(SubscriptionIAMPolicyOutput)
}

type SubscriptionIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (SubscriptionIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscriptionIAMPolicy)(nil)).Elem()
}

func (o SubscriptionIAMPolicyMapOutput) ToSubscriptionIAMPolicyMapOutput() SubscriptionIAMPolicyMapOutput {
	return o
}

func (o SubscriptionIAMPolicyMapOutput) ToSubscriptionIAMPolicyMapOutputWithContext(ctx context.Context) SubscriptionIAMPolicyMapOutput {
	return o
}

func (o SubscriptionIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) SubscriptionIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscriptionIAMPolicy {
		return vs[0].(map[string]*SubscriptionIAMPolicy)[vs[1].(string)]
	}).(SubscriptionIAMPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMPolicyInput)(nil)).Elem(), &SubscriptionIAMPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMPolicyArrayInput)(nil)).Elem(), SubscriptionIAMPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionIAMPolicyMapInput)(nil)).Elem(), SubscriptionIAMPolicyMap{})
	pulumi.RegisterOutputType(SubscriptionIAMPolicyOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionIAMPolicyMapOutput{})
}
