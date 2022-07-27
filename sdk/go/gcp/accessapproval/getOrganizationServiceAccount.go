// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accessapproval

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the email address of an organization's Access Approval service account.
//
// Each Google Cloud organization has a unique service account used by Access Approval.
// When using Access Approval with a
// [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
// this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
// Cloud KMS key used to sign approvals.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accessapproval"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		serviceAccount, err := accessapproval.GetOrganizationServiceAccount(ctx, &accessapproval.GetOrganizationServiceAccountArgs{
// 			OrganizationId: "my-organization",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKeyIAMMember(ctx, "iam", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Crypto_key.Id),
// 			Role:        pulumi.String("roles/cloudkms.signerVerifier"),
// 			Member:      pulumi.String(fmt.Sprintf("serviceAccount:%v", serviceAccount.AccountEmail)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganizationServiceAccount(ctx *pulumi.Context, args *GetOrganizationServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetOrganizationServiceAccountResult, error) {
	var rv GetOrganizationServiceAccountResult
	err := ctx.Invoke("gcp:accessapproval/getOrganizationServiceAccount:getOrganizationServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationServiceAccount.
type GetOrganizationServiceAccountArgs struct {
	// The organization ID the service account was created for.
	OrganizationId string `pulumi:"organizationId"`
}

// A collection of values returned by getOrganizationServiceAccount.
type GetOrganizationServiceAccountResult struct {
	// The email address of the service account. This value is
	// often used to refer to the service account in order to grant IAM permissions.
	AccountEmail string `pulumi:"accountEmail"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Access Approval service account resource name. Format is "organizations/{organization_id}/serviceAccount".
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
}

func GetOrganizationServiceAccountOutput(ctx *pulumi.Context, args GetOrganizationServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationServiceAccountResult, error) {
			args := v.(GetOrganizationServiceAccountArgs)
			r, err := GetOrganizationServiceAccount(ctx, &args, opts...)
			var s GetOrganizationServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationServiceAccountResultOutput)
}

// A collection of arguments for invoking getOrganizationServiceAccount.
type GetOrganizationServiceAccountOutputArgs struct {
	// The organization ID the service account was created for.
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (GetOrganizationServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationServiceAccount.
type GetOrganizationServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationServiceAccountResult)(nil)).Elem()
}

func (o GetOrganizationServiceAccountResultOutput) ToGetOrganizationServiceAccountResultOutput() GetOrganizationServiceAccountResultOutput {
	return o
}

func (o GetOrganizationServiceAccountResultOutput) ToGetOrganizationServiceAccountResultOutputWithContext(ctx context.Context) GetOrganizationServiceAccountResultOutput {
	return o
}

// The email address of the service account. This value is
// often used to refer to the service account in order to grant IAM permissions.
func (o GetOrganizationServiceAccountResultOutput) AccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationServiceAccountResult) string { return v.AccountEmail }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Access Approval service account resource name. Format is "organizations/{organization_id}/serviceAccount".
func (o GetOrganizationServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOrganizationServiceAccountResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationServiceAccountResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationServiceAccountResultOutput{})
}
