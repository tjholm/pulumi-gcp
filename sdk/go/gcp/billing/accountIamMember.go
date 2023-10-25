// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/billing"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/billing"
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/billing"
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
// Instance IAM resources can be imported using the project, table name, role and/or member.
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamMember:AccountIamMember binding "your-billing-account-id"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamMember:AccountIamMember binding "your-billing-account-id roles/billing.user"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamMember:AccountIamMember binding "your-billing-account-id roles/billing.user user:jane@example.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `organizations/my-org-id/roles/my-custom-role`.
type AccountIamMember struct {
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
	BillingAccountId pulumi.StringOutput                `pulumi:"billingAccountId"`
	Condition        AccountIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAccountIamMember registers a new resource with the given unique name, arguments, and options.
func NewAccountIamMember(ctx *pulumi.Context,
	name string, args *AccountIamMemberArgs, opts ...pulumi.ResourceOption) (*AccountIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountIamMember
	err := ctx.RegisterResource("gcp:billing/accountIamMember:AccountIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamMember gets an existing AccountIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamMemberState, opts ...pulumi.ResourceOption) (*AccountIamMember, error) {
	var resource AccountIamMember
	err := ctx.ReadResource("gcp:billing/accountIamMember:AccountIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamMember resources.
type accountIamMemberState struct {
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
	BillingAccountId *string                    `pulumi:"billingAccountId"`
	Condition        *AccountIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role *string `pulumi:"role"`
}

type AccountIamMemberState struct {
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
	Condition        AccountIamMemberConditionPtrInput
	// (Computed) The etag of the billing account's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringPtrInput
}

func (AccountIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamMemberState)(nil)).Elem()
}

type accountIamMemberArgs struct {
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
	BillingAccountId string                     `pulumi:"billingAccountId"`
	Condition        *AccountIamMemberCondition `pulumi:"condition"`
	Member           string                     `pulumi:"member"`
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AccountIamMember resource.
type AccountIamMemberArgs struct {
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
	Condition        AccountIamMemberConditionPtrInput
	Member           pulumi.StringInput
	// The role that should be applied. Only one
	// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
	//
	// `billing.AccountIamPolicy` only:
	Role pulumi.StringInput
}

func (AccountIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamMemberArgs)(nil)).Elem()
}

type AccountIamMemberInput interface {
	pulumi.Input

	ToAccountIamMemberOutput() AccountIamMemberOutput
	ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput
}

func (*AccountIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamMember)(nil)).Elem()
}

func (i *AccountIamMember) ToAccountIamMemberOutput() AccountIamMemberOutput {
	return i.ToAccountIamMemberOutputWithContext(context.Background())
}

func (i *AccountIamMember) ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamMemberOutput)
}

func (i *AccountIamMember) ToOutput(ctx context.Context) pulumix.Output[*AccountIamMember] {
	return pulumix.Output[*AccountIamMember]{
		OutputState: i.ToAccountIamMemberOutputWithContext(ctx).OutputState,
	}
}

// AccountIamMemberArrayInput is an input type that accepts AccountIamMemberArray and AccountIamMemberArrayOutput values.
// You can construct a concrete instance of `AccountIamMemberArrayInput` via:
//
//	AccountIamMemberArray{ AccountIamMemberArgs{...} }
type AccountIamMemberArrayInput interface {
	pulumi.Input

	ToAccountIamMemberArrayOutput() AccountIamMemberArrayOutput
	ToAccountIamMemberArrayOutputWithContext(context.Context) AccountIamMemberArrayOutput
}

type AccountIamMemberArray []AccountIamMemberInput

func (AccountIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamMember)(nil)).Elem()
}

func (i AccountIamMemberArray) ToAccountIamMemberArrayOutput() AccountIamMemberArrayOutput {
	return i.ToAccountIamMemberArrayOutputWithContext(context.Background())
}

func (i AccountIamMemberArray) ToAccountIamMemberArrayOutputWithContext(ctx context.Context) AccountIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamMemberArrayOutput)
}

func (i AccountIamMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*AccountIamMember] {
	return pulumix.Output[[]*AccountIamMember]{
		OutputState: i.ToAccountIamMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// AccountIamMemberMapInput is an input type that accepts AccountIamMemberMap and AccountIamMemberMapOutput values.
// You can construct a concrete instance of `AccountIamMemberMapInput` via:
//
//	AccountIamMemberMap{ "key": AccountIamMemberArgs{...} }
type AccountIamMemberMapInput interface {
	pulumi.Input

	ToAccountIamMemberMapOutput() AccountIamMemberMapOutput
	ToAccountIamMemberMapOutputWithContext(context.Context) AccountIamMemberMapOutput
}

type AccountIamMemberMap map[string]AccountIamMemberInput

func (AccountIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamMember)(nil)).Elem()
}

func (i AccountIamMemberMap) ToAccountIamMemberMapOutput() AccountIamMemberMapOutput {
	return i.ToAccountIamMemberMapOutputWithContext(context.Background())
}

func (i AccountIamMemberMap) ToAccountIamMemberMapOutputWithContext(ctx context.Context) AccountIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamMemberMapOutput)
}

func (i AccountIamMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountIamMember] {
	return pulumix.Output[map[string]*AccountIamMember]{
		OutputState: i.ToAccountIamMemberMapOutputWithContext(ctx).OutputState,
	}
}

type AccountIamMemberOutput struct{ *pulumi.OutputState }

func (AccountIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamMember)(nil)).Elem()
}

func (o AccountIamMemberOutput) ToAccountIamMemberOutput() AccountIamMemberOutput {
	return o
}

func (o AccountIamMemberOutput) ToAccountIamMemberOutputWithContext(ctx context.Context) AccountIamMemberOutput {
	return o
}

func (o AccountIamMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*AccountIamMember] {
	return pulumix.Output[*AccountIamMember]{
		OutputState: o.OutputState,
	}
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
func (o AccountIamMemberOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamMember) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

func (o AccountIamMemberOutput) Condition() AccountIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *AccountIamMember) AccountIamMemberConditionPtrOutput { return v.Condition }).(AccountIamMemberConditionPtrOutput)
}

// (Computed) The etag of the billing account's IAM policy.
func (o AccountIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AccountIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
//
// `billing.AccountIamPolicy` only:
func (o AccountIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AccountIamMemberArrayOutput struct{ *pulumi.OutputState }

func (AccountIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamMember)(nil)).Elem()
}

func (o AccountIamMemberArrayOutput) ToAccountIamMemberArrayOutput() AccountIamMemberArrayOutput {
	return o
}

func (o AccountIamMemberArrayOutput) ToAccountIamMemberArrayOutputWithContext(ctx context.Context) AccountIamMemberArrayOutput {
	return o
}

func (o AccountIamMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AccountIamMember] {
	return pulumix.Output[[]*AccountIamMember]{
		OutputState: o.OutputState,
	}
}

func (o AccountIamMemberArrayOutput) Index(i pulumi.IntInput) AccountIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountIamMember {
		return vs[0].([]*AccountIamMember)[vs[1].(int)]
	}).(AccountIamMemberOutput)
}

type AccountIamMemberMapOutput struct{ *pulumi.OutputState }

func (AccountIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamMember)(nil)).Elem()
}

func (o AccountIamMemberMapOutput) ToAccountIamMemberMapOutput() AccountIamMemberMapOutput {
	return o
}

func (o AccountIamMemberMapOutput) ToAccountIamMemberMapOutputWithContext(ctx context.Context) AccountIamMemberMapOutput {
	return o
}

func (o AccountIamMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountIamMember] {
	return pulumix.Output[map[string]*AccountIamMember]{
		OutputState: o.OutputState,
	}
}

func (o AccountIamMemberMapOutput) MapIndex(k pulumi.StringInput) AccountIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountIamMember {
		return vs[0].(map[string]*AccountIamMember)[vs[1].(string)]
	}).(AccountIamMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamMemberInput)(nil)).Elem(), &AccountIamMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamMemberArrayInput)(nil)).Elem(), AccountIamMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamMemberMapInput)(nil)).Elem(), AccountIamMemberMap{})
	pulumi.RegisterOutputType(AccountIamMemberOutput{})
	pulumi.RegisterOutputType(AccountIamMemberArrayOutput{})
	pulumi.RegisterOutputType(AccountIamMemberMapOutput{})
}
