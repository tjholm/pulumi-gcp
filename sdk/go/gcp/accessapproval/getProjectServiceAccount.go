// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accessapproval

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the email address of a project's Access Approval service account.
//
// Each Google Cloud project has a unique service account used by Access Approval.
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
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/accessapproval"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			serviceAccount, err := accessapproval.GetProjectServiceAccount(ctx, &accessapproval.GetProjectServiceAccountArgs{
//				ProjectId: "my-project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewCryptoKeyIAMMember(ctx, "iam", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: pulumi.Any(google_kms_crypto_key.Crypto_key.Id),
//				Role:        pulumi.String("roles/cloudkms.signerVerifier"),
//				Member:      pulumi.String(fmt.Sprintf("serviceAccount:%v", serviceAccount.AccountEmail)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjectServiceAccount(ctx *pulumi.Context, args *GetProjectServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetProjectServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectServiceAccountResult
	err := ctx.Invoke("gcp:accessapproval/getProjectServiceAccount:getProjectServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectServiceAccount.
type GetProjectServiceAccountArgs struct {
	// The project ID the service account was created for.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getProjectServiceAccount.
type GetProjectServiceAccountResult struct {
	// The email address of the service account. This value is
	// often used to refer to the service account in order to grant IAM permissions.
	AccountEmail string `pulumi:"accountEmail"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Access Approval service account resource name. Format is "projects/{project_id}/serviceAccount".
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
}

func GetProjectServiceAccountOutput(ctx *pulumi.Context, args GetProjectServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetProjectServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectServiceAccountResult, error) {
			args := v.(GetProjectServiceAccountArgs)
			r, err := GetProjectServiceAccount(ctx, &args, opts...)
			var s GetProjectServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectServiceAccountResultOutput)
}

// A collection of arguments for invoking getProjectServiceAccount.
type GetProjectServiceAccountOutputArgs struct {
	// The project ID the service account was created for.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetProjectServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getProjectServiceAccount.
type GetProjectServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetProjectServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectServiceAccountResult)(nil)).Elem()
}

func (o GetProjectServiceAccountResultOutput) ToGetProjectServiceAccountResultOutput() GetProjectServiceAccountResultOutput {
	return o
}

func (o GetProjectServiceAccountResultOutput) ToGetProjectServiceAccountResultOutputWithContext(ctx context.Context) GetProjectServiceAccountResultOutput {
	return o
}

func (o GetProjectServiceAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProjectServiceAccountResult] {
	return pulumix.Output[GetProjectServiceAccountResult]{
		OutputState: o.OutputState,
	}
}

// The email address of the service account. This value is
// often used to refer to the service account in order to grant IAM permissions.
func (o GetProjectServiceAccountResultOutput) AccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectServiceAccountResult) string { return v.AccountEmail }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Access Approval service account resource name. Format is "projects/{project_id}/serviceAccount".
func (o GetProjectServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetProjectServiceAccountResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectServiceAccountResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectServiceAccountResultOutput{})
}
