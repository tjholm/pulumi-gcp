// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve the default App Engine service account for the specified project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/appengine"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := appengine.GetDefaultServiceAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("defaultAccount", _default.Email)
//			return nil
//		})
//	}
//
// ```
func GetDefaultServiceAccount(ctx *pulumi.Context, args *GetDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetDefaultServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDefaultServiceAccountResult
	err := ctx.Invoke("gcp:appengine/getDefaultServiceAccount:getDefaultServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResult struct {
	// The display name for the service account.
	DisplayName string `pulumi:"displayName"`
	// Email address of the default service account used by App Engine in this project.
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
	Member string `pulumi:"member"`
	// The fully-qualified name of the service account.
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId string `pulumi:"uniqueId"`
}

func GetDefaultServiceAccountOutput(ctx *pulumi.Context, args GetDefaultServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetDefaultServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDefaultServiceAccountResult, error) {
			args := v.(GetDefaultServiceAccountArgs)
			r, err := GetDefaultServiceAccount(ctx, &args, opts...)
			var s GetDefaultServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDefaultServiceAccountResultOutput)
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountOutputArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetDefaultServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetDefaultServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultServiceAccountResult)(nil)).Elem()
}

func (o GetDefaultServiceAccountResultOutput) ToGetDefaultServiceAccountResultOutput() GetDefaultServiceAccountResultOutput {
	return o
}

func (o GetDefaultServiceAccountResultOutput) ToGetDefaultServiceAccountResultOutputWithContext(ctx context.Context) GetDefaultServiceAccountResultOutput {
	return o
}

func (o GetDefaultServiceAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDefaultServiceAccountResult] {
	return pulumix.Output[GetDefaultServiceAccountResult]{
		OutputState: o.OutputState,
	}
}

// The display name for the service account.
func (o GetDefaultServiceAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Email address of the default service account used by App Engine in this project.
func (o GetDefaultServiceAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDefaultServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
func (o GetDefaultServiceAccountResultOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Member }).(pulumi.StringOutput)
}

// The fully-qualified name of the service account.
func (o GetDefaultServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDefaultServiceAccountResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Project }).(pulumi.StringOutput)
}

// The unique id of the service account.
func (o GetDefaultServiceAccountResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDefaultServiceAccountResultOutput{})
}
