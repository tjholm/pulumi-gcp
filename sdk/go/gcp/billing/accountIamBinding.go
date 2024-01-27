// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Three different resources help you manage IAM policies on billing accounts. Each of these resources serves a different use case:
//
// * `billing.AccountIamPolicy`: Authoritative. Sets the IAM policy for the billing accounts and replaces any existing policy already attached.
// * `billing.AccountIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
// * `billing.AccountIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the billing accounts are preserved.
//
// > **Note:** `billing.AccountIamPolicy` **cannot** be used in conjunction with `billing.AccountIamBinding` and `billing.AccountIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the billing account as `billing.AccountIamPolicy` replaces the entire policy.
//
// > **Note:** `billing.AccountIamBinding` resources **can be** used in conjunction with `billing.AccountIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_billing\_account\_iam\_policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/billing"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/billing.viewer",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = billing.NewAccountIamPolicy(ctx, "editor", &billing.AccountIamPolicyArgs{
//				BillingAccountId: pulumi.String("00AA00-000AAA-00AA0A"),
//				PolicyData:       *pulumi.String(admin.PolicyData),
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
// ## google\_billing\_account\_iam\_binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/billing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := billing.NewAccountIamBinding(ctx, "editor", &billing.AccountIamBindingArgs{
//				BillingAccountId: pulumi.String("00AA00-000AAA-00AA0A"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Role: pulumi.String("roles/billing.viewer"),
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
// ## google\_billing\_account\_iam\_member
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/billing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := billing.NewAccountIamMember(ctx, "editor", &billing.AccountIamMemberArgs{
//				BillingAccountId: pulumi.String("00AA00-000AAA-00AA0A"),
//				Member:           pulumi.String("user:jane@example.com"),
//				Role:             pulumi.String("roles/billing.viewer"),
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
// ### Importing IAM policies IAM policy imports use the `billing_account_id` identifier of the Billing Account resource only. For example* `{{billing_account_id}}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
//
//	id = {{billing_account_id}}
//
//	to = google_billing_account_iam_policy.default } The `pulumi import` command can also be used
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamBinding:AccountIamBinding default {{billing_account_id}}
//
// ```
type AccountIamBinding struct {
	pulumi.CustomResourceState

	// The billing account id.
	//
	// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	BillingAccountId pulumi.StringOutput                 `pulumi:"billingAccountId"`
	Condition        AccountIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAccountIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAccountIamBinding(ctx *pulumi.Context,
	name string, args *AccountIamBindingArgs, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountIamBinding
	err := ctx.RegisterResource("gcp:billing/accountIamBinding:AccountIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamBinding gets an existing AccountIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamBindingState, opts ...pulumi.ResourceOption) (*AccountIamBinding, error) {
	var resource AccountIamBinding
	err := ctx.ReadResource("gcp:billing/accountIamBinding:AccountIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamBinding resources.
type accountIamBindingState struct {
	// The billing account id.
	//
	// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	BillingAccountId *string                     `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role *string `pulumi:"role"`
}

type AccountIamBindingState struct {
	// The billing account id.
	//
	// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	BillingAccountId pulumi.StringPtrInput
	Condition        AccountIamBindingConditionPtrInput
	// (Computed) The etag of the billing account's IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringPtrInput
}

func (AccountIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingState)(nil)).Elem()
}

type accountIamBindingArgs struct {
	// The billing account id.
	//
	// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	BillingAccountId string                      `pulumi:"billingAccountId"`
	Condition        *AccountIamBindingCondition `pulumi:"condition"`
	Members          []string                    `pulumi:"members"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AccountIamBinding resource.
type AccountIamBindingArgs struct {
	// The billing account id.
	//
	// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
	//
	// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
	//   Each entry can have one of the following values:
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	BillingAccountId pulumi.StringInput
	Condition        AccountIamBindingConditionPtrInput
	Members          pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringInput
}

func (AccountIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamBindingArgs)(nil)).Elem()
}

type AccountIamBindingInput interface {
	pulumi.Input

	ToAccountIamBindingOutput() AccountIamBindingOutput
	ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput
}

func (*AccountIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamBinding)(nil)).Elem()
}

func (i *AccountIamBinding) ToAccountIamBindingOutput() AccountIamBindingOutput {
	return i.ToAccountIamBindingOutputWithContext(context.Background())
}

func (i *AccountIamBinding) ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamBindingOutput)
}

// AccountIamBindingArrayInput is an input type that accepts AccountIamBindingArray and AccountIamBindingArrayOutput values.
// You can construct a concrete instance of `AccountIamBindingArrayInput` via:
//
//	AccountIamBindingArray{ AccountIamBindingArgs{...} }
type AccountIamBindingArrayInput interface {
	pulumi.Input

	ToAccountIamBindingArrayOutput() AccountIamBindingArrayOutput
	ToAccountIamBindingArrayOutputWithContext(context.Context) AccountIamBindingArrayOutput
}

type AccountIamBindingArray []AccountIamBindingInput

func (AccountIamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamBinding)(nil)).Elem()
}

func (i AccountIamBindingArray) ToAccountIamBindingArrayOutput() AccountIamBindingArrayOutput {
	return i.ToAccountIamBindingArrayOutputWithContext(context.Background())
}

func (i AccountIamBindingArray) ToAccountIamBindingArrayOutputWithContext(ctx context.Context) AccountIamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamBindingArrayOutput)
}

// AccountIamBindingMapInput is an input type that accepts AccountIamBindingMap and AccountIamBindingMapOutput values.
// You can construct a concrete instance of `AccountIamBindingMapInput` via:
//
//	AccountIamBindingMap{ "key": AccountIamBindingArgs{...} }
type AccountIamBindingMapInput interface {
	pulumi.Input

	ToAccountIamBindingMapOutput() AccountIamBindingMapOutput
	ToAccountIamBindingMapOutputWithContext(context.Context) AccountIamBindingMapOutput
}

type AccountIamBindingMap map[string]AccountIamBindingInput

func (AccountIamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamBinding)(nil)).Elem()
}

func (i AccountIamBindingMap) ToAccountIamBindingMapOutput() AccountIamBindingMapOutput {
	return i.ToAccountIamBindingMapOutputWithContext(context.Background())
}

func (i AccountIamBindingMap) ToAccountIamBindingMapOutputWithContext(ctx context.Context) AccountIamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamBindingMapOutput)
}

type AccountIamBindingOutput struct{ *pulumi.OutputState }

func (AccountIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamBinding)(nil)).Elem()
}

func (o AccountIamBindingOutput) ToAccountIamBindingOutput() AccountIamBindingOutput {
	return o
}

func (o AccountIamBindingOutput) ToAccountIamBindingOutputWithContext(ctx context.Context) AccountIamBindingOutput {
	return o
}

// The billing account id.
//
// For `billing.AccountIamMember` or `billing.AccountIamBinding`:
//
//   - `member/members` - (Required) Identities that will be granted the privilege in `role`.
//     Each entry can have one of the following values:
//   - **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
//   - **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o AccountIamBindingOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamBinding) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

func (o AccountIamBindingOutput) Condition() AccountIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *AccountIamBinding) AccountIamBindingConditionPtrOutput { return v.Condition }).(AccountIamBindingConditionPtrOutput)
}

// (Computed) The etag of the billing account's IAM policy.
func (o AccountIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AccountIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AccountIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The role that should be applied. Only one
// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
//
// `billing.AccountIamPolicy` only:
func (o AccountIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AccountIamBindingArrayOutput struct{ *pulumi.OutputState }

func (AccountIamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamBinding)(nil)).Elem()
}

func (o AccountIamBindingArrayOutput) ToAccountIamBindingArrayOutput() AccountIamBindingArrayOutput {
	return o
}

func (o AccountIamBindingArrayOutput) ToAccountIamBindingArrayOutputWithContext(ctx context.Context) AccountIamBindingArrayOutput {
	return o
}

func (o AccountIamBindingArrayOutput) Index(i pulumi.IntInput) AccountIamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountIamBinding {
		return vs[0].([]*AccountIamBinding)[vs[1].(int)]
	}).(AccountIamBindingOutput)
}

type AccountIamBindingMapOutput struct{ *pulumi.OutputState }

func (AccountIamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamBinding)(nil)).Elem()
}

func (o AccountIamBindingMapOutput) ToAccountIamBindingMapOutput() AccountIamBindingMapOutput {
	return o
}

func (o AccountIamBindingMapOutput) ToAccountIamBindingMapOutputWithContext(ctx context.Context) AccountIamBindingMapOutput {
	return o
}

func (o AccountIamBindingMapOutput) MapIndex(k pulumi.StringInput) AccountIamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountIamBinding {
		return vs[0].(map[string]*AccountIamBinding)[vs[1].(string)]
	}).(AccountIamBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamBindingInput)(nil)).Elem(), &AccountIamBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamBindingArrayInput)(nil)).Elem(), AccountIamBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamBindingMapInput)(nil)).Elem(), AccountIamBindingMap{})
	pulumi.RegisterOutputType(AccountIamBindingOutput{})
	pulumi.RegisterOutputType(AccountIamBindingArrayOutput{})
	pulumi.RegisterOutputType(AccountIamBindingMapOutput{})
}
