// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a Google Cloud Billing Subaccount.
//
// !> **WARNING:** Deleting this resource will not delete or close the billing subaccount.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/billing"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := billing.NewSubAccount(ctx, "subaccount", &billing.SubAccountArgs{
// 			DisplayName:          pulumi.String("My Billing Account"),
// 			MasterBillingAccount: pulumi.String("012345-567890-ABCDEF"),
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
// Billing Subaccounts can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:billing/subAccount:SubAccount default billingAccounts/{billing_account_id}
// ```
type SubAccount struct {
	pulumi.CustomResourceState

	// The billing account id.
	BillingAccountId pulumi.StringOutput `pulumi:"billingAccountId"`
	// If set to "RENAME_ON_DESTROY" the billing account displayName
	// will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	// Default is "".
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// The display name of the billing account.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the master billing account that the subaccount
	// will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	MasterBillingAccount pulumi.StringOutput `pulumi:"masterBillingAccount"`
	// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// `true` if the billing account is open, `false` if the billing account is closed.
	Open pulumi.BoolOutput `pulumi:"open"`
}

// NewSubAccount registers a new resource with the given unique name, arguments, and options.
func NewSubAccount(ctx *pulumi.Context,
	name string, args *SubAccountArgs, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MasterBillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'MasterBillingAccount'")
	}
	var resource SubAccount
	err := ctx.RegisterResource("gcp:billing/subAccount:SubAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubAccount gets an existing SubAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountState, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	var resource SubAccount
	err := ctx.ReadResource("gcp:billing/subAccount:SubAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubAccount resources.
type subAccountState struct {
	// The billing account id.
	BillingAccountId *string `pulumi:"billingAccountId"`
	// If set to "RENAME_ON_DESTROY" the billing account displayName
	// will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	// Default is "".
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The display name of the billing account.
	DisplayName *string `pulumi:"displayName"`
	// The name of the master billing account that the subaccount
	// will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	MasterBillingAccount *string `pulumi:"masterBillingAccount"`
	// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
	Name *string `pulumi:"name"`
	// `true` if the billing account is open, `false` if the billing account is closed.
	Open *bool `pulumi:"open"`
}

type SubAccountState struct {
	// The billing account id.
	BillingAccountId pulumi.StringPtrInput
	// If set to "RENAME_ON_DESTROY" the billing account displayName
	// will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	// Default is "".
	DeletionPolicy pulumi.StringPtrInput
	// The display name of the billing account.
	DisplayName pulumi.StringPtrInput
	// The name of the master billing account that the subaccount
	// will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	MasterBillingAccount pulumi.StringPtrInput
	// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
	Name pulumi.StringPtrInput
	// `true` if the billing account is open, `false` if the billing account is closed.
	Open pulumi.BoolPtrInput
}

func (SubAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountState)(nil)).Elem()
}

type subAccountArgs struct {
	// If set to "RENAME_ON_DESTROY" the billing account displayName
	// will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	// Default is "".
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The display name of the billing account.
	DisplayName string `pulumi:"displayName"`
	// The name of the master billing account that the subaccount
	// will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	MasterBillingAccount string `pulumi:"masterBillingAccount"`
}

// The set of arguments for constructing a SubAccount resource.
type SubAccountArgs struct {
	// If set to "RENAME_ON_DESTROY" the billing account displayName
	// will be changed to "Destroyed" along with a timestamp.  If set to "" this will not occur.
	// Default is "".
	DeletionPolicy pulumi.StringPtrInput
	// The display name of the billing account.
	DisplayName pulumi.StringInput
	// The name of the master billing account that the subaccount
	// will be created under in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
	MasterBillingAccount pulumi.StringInput
}

func (SubAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountArgs)(nil)).Elem()
}

type SubAccountInput interface {
	pulumi.Input

	ToSubAccountOutput() SubAccountOutput
	ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput
}

func (*SubAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (i *SubAccount) ToSubAccountOutput() SubAccountOutput {
	return i.ToSubAccountOutputWithContext(context.Background())
}

func (i *SubAccount) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountOutput)
}

// SubAccountArrayInput is an input type that accepts SubAccountArray and SubAccountArrayOutput values.
// You can construct a concrete instance of `SubAccountArrayInput` via:
//
//          SubAccountArray{ SubAccountArgs{...} }
type SubAccountArrayInput interface {
	pulumi.Input

	ToSubAccountArrayOutput() SubAccountArrayOutput
	ToSubAccountArrayOutputWithContext(context.Context) SubAccountArrayOutput
}

type SubAccountArray []SubAccountInput

func (SubAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubAccount)(nil)).Elem()
}

func (i SubAccountArray) ToSubAccountArrayOutput() SubAccountArrayOutput {
	return i.ToSubAccountArrayOutputWithContext(context.Background())
}

func (i SubAccountArray) ToSubAccountArrayOutputWithContext(ctx context.Context) SubAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountArrayOutput)
}

// SubAccountMapInput is an input type that accepts SubAccountMap and SubAccountMapOutput values.
// You can construct a concrete instance of `SubAccountMapInput` via:
//
//          SubAccountMap{ "key": SubAccountArgs{...} }
type SubAccountMapInput interface {
	pulumi.Input

	ToSubAccountMapOutput() SubAccountMapOutput
	ToSubAccountMapOutputWithContext(context.Context) SubAccountMapOutput
}

type SubAccountMap map[string]SubAccountInput

func (SubAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubAccount)(nil)).Elem()
}

func (i SubAccountMap) ToSubAccountMapOutput() SubAccountMapOutput {
	return i.ToSubAccountMapOutputWithContext(context.Background())
}

func (i SubAccountMap) ToSubAccountMapOutputWithContext(ctx context.Context) SubAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountMapOutput)
}

type SubAccountOutput struct{ *pulumi.OutputState }

func (SubAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (o SubAccountOutput) ToSubAccountOutput() SubAccountOutput {
	return o
}

func (o SubAccountOutput) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return o
}

type SubAccountArrayOutput struct{ *pulumi.OutputState }

func (SubAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubAccount)(nil)).Elem()
}

func (o SubAccountArrayOutput) ToSubAccountArrayOutput() SubAccountArrayOutput {
	return o
}

func (o SubAccountArrayOutput) ToSubAccountArrayOutputWithContext(ctx context.Context) SubAccountArrayOutput {
	return o
}

func (o SubAccountArrayOutput) Index(i pulumi.IntInput) SubAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubAccount {
		return vs[0].([]*SubAccount)[vs[1].(int)]
	}).(SubAccountOutput)
}

type SubAccountMapOutput struct{ *pulumi.OutputState }

func (SubAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubAccount)(nil)).Elem()
}

func (o SubAccountMapOutput) ToSubAccountMapOutput() SubAccountMapOutput {
	return o
}

func (o SubAccountMapOutput) ToSubAccountMapOutputWithContext(ctx context.Context) SubAccountMapOutput {
	return o
}

func (o SubAccountMapOutput) MapIndex(k pulumi.StringInput) SubAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubAccount {
		return vs[0].(map[string]*SubAccount)[vs[1].(string)]
	}).(SubAccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubAccountInput)(nil)).Elem(), &SubAccount{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubAccountArrayInput)(nil)).Elem(), SubAccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubAccountMapInput)(nil)).Elem(), SubAccountMap{})
	pulumi.RegisterOutputType(SubAccountOutput{})
	pulumi.RegisterOutputType(SubAccountArrayOutput{})
	pulumi.RegisterOutputType(SubAccountMapOutput{})
}
