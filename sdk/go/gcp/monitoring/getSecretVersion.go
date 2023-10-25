// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the value and metadata from a Secret Manager secret version. For more information see the official documentation datasource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secretmanager.LookupSecretVersion(ctx, &secretmanager.LookupSecretVersionArgs{
//				Secret: "my-secret",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: gcp.monitoring.getSecretVersion has been deprecated in favor of gcp.secretmanager.getSecretVersion
func GetSecretVersion(ctx *pulumi.Context, args *GetSecretVersionArgs, opts ...pulumi.InvokeOption) (*GetSecretVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSecretVersionResult
	err := ctx.Invoke("gcp:monitoring/getSecretVersion:getSecretVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretVersion.
type GetSecretVersionArgs struct {
	// The project to get the secret version for. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The secret to get the secret version for.
	Secret string `pulumi:"secret"`
	// The version of the secret to get. If it
	// is not provided, the latest version is retrieved.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getSecretVersion.
type GetSecretVersionResult struct {
	// The time at which the Secret was created.
	CreateTime string `pulumi:"createTime"`
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime string `pulumi:"destroyTime"`
	// True if the current state of the SecretVersion is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The resource name of the SecretVersion. Format:
	// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	Secret  string `pulumi:"secret"`
	// The secret data. No larger than 64KiB.
	SecretData string `pulumi:"secretData"`
	Version    string `pulumi:"version"`
}

func GetSecretVersionOutput(ctx *pulumi.Context, args GetSecretVersionOutputArgs, opts ...pulumi.InvokeOption) GetSecretVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretVersionResult, error) {
			args := v.(GetSecretVersionArgs)
			r, err := GetSecretVersion(ctx, &args, opts...)
			var s GetSecretVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretVersionResultOutput)
}

// A collection of arguments for invoking getSecretVersion.
type GetSecretVersionOutputArgs struct {
	// The project to get the secret version for. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The secret to get the secret version for.
	Secret pulumi.StringInput `pulumi:"secret"`
	// The version of the secret to get. If it
	// is not provided, the latest version is retrieved.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (GetSecretVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionArgs)(nil)).Elem()
}

// A collection of values returned by getSecretVersion.
type GetSecretVersionResultOutput struct{ *pulumi.OutputState }

func (GetSecretVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionResult)(nil)).Elem()
}

func (o GetSecretVersionResultOutput) ToGetSecretVersionResultOutput() GetSecretVersionResultOutput {
	return o
}

func (o GetSecretVersionResultOutput) ToGetSecretVersionResultOutputWithContext(ctx context.Context) GetSecretVersionResultOutput {
	return o
}

func (o GetSecretVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSecretVersionResult] {
	return pulumix.Output[GetSecretVersionResult]{
		OutputState: o.OutputState,
	}
}

// The time at which the Secret was created.
func (o GetSecretVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which the Secret was destroyed. Only present if state is DESTROYED.
func (o GetSecretVersionResultOutput) DestroyTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.DestroyTime }).(pulumi.StringOutput)
}

// True if the current state of the SecretVersion is enabled.
func (o GetSecretVersionResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetSecretVersionResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name of the SecretVersion. Format:
// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
func (o GetSecretVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Secret }).(pulumi.StringOutput)
}

// The secret data. No larger than 64KiB.
func (o GetSecretVersionResultOutput) SecretData() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.SecretData }).(pulumi.StringOutput)
}

func (o GetSecretVersionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretVersionResultOutput{})
}
