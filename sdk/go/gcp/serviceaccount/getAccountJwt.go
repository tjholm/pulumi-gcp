// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides a [self-signed JWT](https://cloud.google.com/iam/docs/create-short-lived-credentials-direct#sa-credentials-jwt).  Tokens issued from this data source are typically used to call external services that accept JWTs for authentication.
//
// ## Example Usage
//
// Note: in order to use the following, the caller must have _at least_ `roles/iam.serviceAccountTokenCreator` on the `targetServiceAccount`.
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"foo": "bar",
//				"sub": "subject",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			foo, err := serviceAccount.GetAccountJwt(ctx, &serviceaccount.GetAccountJwtArgs{
//				TargetServiceAccount: "impersonated-account@project.iam.gserviceaccount.com",
//				Payload:              json0,
//				ExpiresIn:            pulumi.IntRef(60),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("jwt", foo.Jwt)
//			return nil
//		})
//	}
//
// ```
func GetAccountJwt(ctx *pulumi.Context, args *GetAccountJwtArgs, opts ...pulumi.InvokeOption) (*GetAccountJwtResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccountJwtResult
	err := ctx.Invoke("gcp:serviceAccount/getAccountJwt:getAccountJwt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountJwt.
type GetAccountJwtArgs struct {
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.
	Delegates []string `pulumi:"delegates"`
	// Number of seconds until the JWT expires. If set and non-zero an `exp` claim will be added to the payload derived from the current timestamp plus expiresIn seconds.
	ExpiresIn *int `pulumi:"expiresIn"`
	// The JSON-encoded JWT claims set to include in the self-signed JWT.
	Payload string `pulumi:"payload"`
	// The email of the service account that will sign the JWT.
	TargetServiceAccount string `pulumi:"targetServiceAccount"`
}

// A collection of values returned by getAccountJwt.
type GetAccountJwtResult struct {
	Delegates []string `pulumi:"delegates"`
	ExpiresIn *int     `pulumi:"expiresIn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The signed JWT containing the JWT Claims Set from the `payload`.
	Jwt                  string `pulumi:"jwt"`
	Payload              string `pulumi:"payload"`
	TargetServiceAccount string `pulumi:"targetServiceAccount"`
}

func GetAccountJwtOutput(ctx *pulumi.Context, args GetAccountJwtOutputArgs, opts ...pulumi.InvokeOption) GetAccountJwtResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccountJwtResult, error) {
			args := v.(GetAccountJwtArgs)
			r, err := GetAccountJwt(ctx, &args, opts...)
			var s GetAccountJwtResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccountJwtResultOutput)
}

// A collection of arguments for invoking getAccountJwt.
type GetAccountJwtOutputArgs struct {
	// Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.
	Delegates pulumi.StringArrayInput `pulumi:"delegates"`
	// Number of seconds until the JWT expires. If set and non-zero an `exp` claim will be added to the payload derived from the current timestamp plus expiresIn seconds.
	ExpiresIn pulumi.IntPtrInput `pulumi:"expiresIn"`
	// The JSON-encoded JWT claims set to include in the self-signed JWT.
	Payload pulumi.StringInput `pulumi:"payload"`
	// The email of the service account that will sign the JWT.
	TargetServiceAccount pulumi.StringInput `pulumi:"targetServiceAccount"`
}

func (GetAccountJwtOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountJwtArgs)(nil)).Elem()
}

// A collection of values returned by getAccountJwt.
type GetAccountJwtResultOutput struct{ *pulumi.OutputState }

func (GetAccountJwtResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccountJwtResult)(nil)).Elem()
}

func (o GetAccountJwtResultOutput) ToGetAccountJwtResultOutput() GetAccountJwtResultOutput {
	return o
}

func (o GetAccountJwtResultOutput) ToGetAccountJwtResultOutputWithContext(ctx context.Context) GetAccountJwtResultOutput {
	return o
}

func (o GetAccountJwtResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetAccountJwtResult] {
	return pulumix.Output[GetAccountJwtResult]{
		OutputState: o.OutputState,
	}
}

func (o GetAccountJwtResultOutput) Delegates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccountJwtResult) []string { return v.Delegates }).(pulumi.StringArrayOutput)
}

func (o GetAccountJwtResultOutput) ExpiresIn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAccountJwtResult) *int { return v.ExpiresIn }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccountJwtResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountJwtResult) string { return v.Id }).(pulumi.StringOutput)
}

// The signed JWT containing the JWT Claims Set from the `payload`.
func (o GetAccountJwtResultOutput) Jwt() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountJwtResult) string { return v.Jwt }).(pulumi.StringOutput)
}

func (o GetAccountJwtResultOutput) Payload() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountJwtResult) string { return v.Payload }).(pulumi.StringOutput)
}

func (o GetAccountJwtResultOutput) TargetServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccountJwtResult) string { return v.TargetServiceAccount }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccountJwtResultOutput{})
}
