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
//	$ pulumi import gcp:billing/accountIamPolicy:AccountIamPolicy binding "your-billing-account-id"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamPolicy:AccountIamPolicy binding "your-billing-account-id roles/billing.user"
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:billing/accountIamPolicy:AccountIamPolicy binding "your-billing-account-id roles/billing.user user:jane@example.com"
//
// ```
//
//	-> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `organizations/my-org-id/roles/my-custom-role`.
type AccountIamPolicy struct {
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
	BillingAccountId pulumi.StringOutput `pulumi:"billingAccountId"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewAccountIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccountIamPolicy(ctx *pulumi.Context,
	name string, args *AccountIamPolicyArgs, opts ...pulumi.ResourceOption) (*AccountIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountIamPolicy
	err := ctx.RegisterResource("gcp:billing/accountIamPolicy:AccountIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountIamPolicy gets an existing AccountIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountIamPolicyState, opts ...pulumi.ResourceOption) (*AccountIamPolicy, error) {
	var resource AccountIamPolicy
	err := ctx.ReadResource("gcp:billing/accountIamPolicy:AccountIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountIamPolicy resources.
type accountIamPolicyState struct {
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
	BillingAccountId *string `pulumi:"billingAccountId"`
	// (Computed) The etag of the billing account's IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData *string `pulumi:"policyData"`
}

type AccountIamPolicyState struct {
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
	// (Computed) The etag of the billing account's IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringPtrInput
}

func (AccountIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamPolicyState)(nil)).Elem()
}

type accountIamPolicyArgs struct {
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
	BillingAccountId string `pulumi:"billingAccountId"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a AccountIamPolicy resource.
type AccountIamPolicyArgs struct {
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
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	//
	// ***
	PolicyData pulumi.StringInput
}

func (AccountIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountIamPolicyArgs)(nil)).Elem()
}

type AccountIamPolicyInput interface {
	pulumi.Input

	ToAccountIamPolicyOutput() AccountIamPolicyOutput
	ToAccountIamPolicyOutputWithContext(ctx context.Context) AccountIamPolicyOutput
}

func (*AccountIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamPolicy)(nil)).Elem()
}

func (i *AccountIamPolicy) ToAccountIamPolicyOutput() AccountIamPolicyOutput {
	return i.ToAccountIamPolicyOutputWithContext(context.Background())
}

func (i *AccountIamPolicy) ToAccountIamPolicyOutputWithContext(ctx context.Context) AccountIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamPolicyOutput)
}

func (i *AccountIamPolicy) ToOutput(ctx context.Context) pulumix.Output[*AccountIamPolicy] {
	return pulumix.Output[*AccountIamPolicy]{
		OutputState: i.ToAccountIamPolicyOutputWithContext(ctx).OutputState,
	}
}

// AccountIamPolicyArrayInput is an input type that accepts AccountIamPolicyArray and AccountIamPolicyArrayOutput values.
// You can construct a concrete instance of `AccountIamPolicyArrayInput` via:
//
//	AccountIamPolicyArray{ AccountIamPolicyArgs{...} }
type AccountIamPolicyArrayInput interface {
	pulumi.Input

	ToAccountIamPolicyArrayOutput() AccountIamPolicyArrayOutput
	ToAccountIamPolicyArrayOutputWithContext(context.Context) AccountIamPolicyArrayOutput
}

type AccountIamPolicyArray []AccountIamPolicyInput

func (AccountIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamPolicy)(nil)).Elem()
}

func (i AccountIamPolicyArray) ToAccountIamPolicyArrayOutput() AccountIamPolicyArrayOutput {
	return i.ToAccountIamPolicyArrayOutputWithContext(context.Background())
}

func (i AccountIamPolicyArray) ToAccountIamPolicyArrayOutputWithContext(ctx context.Context) AccountIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamPolicyArrayOutput)
}

func (i AccountIamPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*AccountIamPolicy] {
	return pulumix.Output[[]*AccountIamPolicy]{
		OutputState: i.ToAccountIamPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// AccountIamPolicyMapInput is an input type that accepts AccountIamPolicyMap and AccountIamPolicyMapOutput values.
// You can construct a concrete instance of `AccountIamPolicyMapInput` via:
//
//	AccountIamPolicyMap{ "key": AccountIamPolicyArgs{...} }
type AccountIamPolicyMapInput interface {
	pulumi.Input

	ToAccountIamPolicyMapOutput() AccountIamPolicyMapOutput
	ToAccountIamPolicyMapOutputWithContext(context.Context) AccountIamPolicyMapOutput
}

type AccountIamPolicyMap map[string]AccountIamPolicyInput

func (AccountIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamPolicy)(nil)).Elem()
}

func (i AccountIamPolicyMap) ToAccountIamPolicyMapOutput() AccountIamPolicyMapOutput {
	return i.ToAccountIamPolicyMapOutputWithContext(context.Background())
}

func (i AccountIamPolicyMap) ToAccountIamPolicyMapOutputWithContext(ctx context.Context) AccountIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountIamPolicyMapOutput)
}

func (i AccountIamPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountIamPolicy] {
	return pulumix.Output[map[string]*AccountIamPolicy]{
		OutputState: i.ToAccountIamPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type AccountIamPolicyOutput struct{ *pulumi.OutputState }

func (AccountIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountIamPolicy)(nil)).Elem()
}

func (o AccountIamPolicyOutput) ToAccountIamPolicyOutput() AccountIamPolicyOutput {
	return o
}

func (o AccountIamPolicyOutput) ToAccountIamPolicyOutputWithContext(ctx context.Context) AccountIamPolicyOutput {
	return o
}

func (o AccountIamPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*AccountIamPolicy] {
	return pulumix.Output[*AccountIamPolicy]{
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
func (o AccountIamPolicyOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamPolicy) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

// (Computed) The etag of the billing account's IAM policy.
func (o AccountIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The policy data generated by a `organizations.getIAMPolicy` data source.
//
// ***
func (o AccountIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type AccountIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (AccountIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountIamPolicy)(nil)).Elem()
}

func (o AccountIamPolicyArrayOutput) ToAccountIamPolicyArrayOutput() AccountIamPolicyArrayOutput {
	return o
}

func (o AccountIamPolicyArrayOutput) ToAccountIamPolicyArrayOutputWithContext(ctx context.Context) AccountIamPolicyArrayOutput {
	return o
}

func (o AccountIamPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AccountIamPolicy] {
	return pulumix.Output[[]*AccountIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o AccountIamPolicyArrayOutput) Index(i pulumi.IntInput) AccountIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountIamPolicy {
		return vs[0].([]*AccountIamPolicy)[vs[1].(int)]
	}).(AccountIamPolicyOutput)
}

type AccountIamPolicyMapOutput struct{ *pulumi.OutputState }

func (AccountIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountIamPolicy)(nil)).Elem()
}

func (o AccountIamPolicyMapOutput) ToAccountIamPolicyMapOutput() AccountIamPolicyMapOutput {
	return o
}

func (o AccountIamPolicyMapOutput) ToAccountIamPolicyMapOutputWithContext(ctx context.Context) AccountIamPolicyMapOutput {
	return o
}

func (o AccountIamPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AccountIamPolicy] {
	return pulumix.Output[map[string]*AccountIamPolicy]{
		OutputState: o.OutputState,
	}
}

func (o AccountIamPolicyMapOutput) MapIndex(k pulumi.StringInput) AccountIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountIamPolicy {
		return vs[0].(map[string]*AccountIamPolicy)[vs[1].(string)]
	}).(AccountIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamPolicyInput)(nil)).Elem(), &AccountIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamPolicyArrayInput)(nil)).Elem(), AccountIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountIamPolicyMapInput)(nil)).Elem(), AccountIamPolicyMap{})
	pulumi.RegisterOutputType(AccountIamPolicyOutput{})
	pulumi.RegisterOutputType(AccountIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(AccountIamPolicyMapOutput{})
}
