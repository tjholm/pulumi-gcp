// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides access to available Kubernetes versions in a location for a given project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := container.GetAwsVersions(ctx, &container.GetAwsVersionsArgs{
//				Location: pulumi.StringRef("us-west1"),
//				Project:  pulumi.StringRef("my-project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstAvailableVersion", data.Google_container_aws_versions.Versions.Valid_versions[0])
//			return nil
//		})
//	}
//
// ```
func GetAwsVersions(ctx *pulumi.Context, args *GetAwsVersionsArgs, opts ...pulumi.InvokeOption) (*GetAwsVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAwsVersionsResult
	err := ctx.Invoke("gcp:container/getAwsVersions:getAwsVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsVersions.
type GetAwsVersionsArgs struct {
	// The location to list versions for.
	Location *string `pulumi:"location"`
	// ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getAwsVersions.
type GetAwsVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// A list of AWS regions that are available for use with this project and GCP location.
	SupportedRegions []string `pulumi:"supportedRegions"`
	// A list of versions available for use with this project and location.
	ValidVersions []string `pulumi:"validVersions"`
}

func GetAwsVersionsOutput(ctx *pulumi.Context, args GetAwsVersionsOutputArgs, opts ...pulumi.InvokeOption) GetAwsVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAwsVersionsResult, error) {
			args := v.(GetAwsVersionsArgs)
			r, err := GetAwsVersions(ctx, &args, opts...)
			var s GetAwsVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAwsVersionsResultOutput)
}

// A collection of arguments for invoking getAwsVersions.
type GetAwsVersionsOutputArgs struct {
	// The location to list versions for.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetAwsVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getAwsVersions.
type GetAwsVersionsResultOutput struct{ *pulumi.OutputState }

func (GetAwsVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsVersionsResult)(nil)).Elem()
}

func (o GetAwsVersionsResultOutput) ToGetAwsVersionsResultOutput() GetAwsVersionsResultOutput {
	return o
}

func (o GetAwsVersionsResultOutput) ToGetAwsVersionsResultOutputWithContext(ctx context.Context) GetAwsVersionsResultOutput {
	return o
}

func (o GetAwsVersionsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetAwsVersionsResult] {
	return pulumix.Output[GetAwsVersionsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetAwsVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAwsVersionsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsVersionsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GetAwsVersionsResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsVersionsResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// A list of AWS regions that are available for use with this project and GCP location.
func (o GetAwsVersionsResultOutput) SupportedRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAwsVersionsResult) []string { return v.SupportedRegions }).(pulumi.StringArrayOutput)
}

// A list of versions available for use with this project and location.
func (o GetAwsVersionsResultOutput) ValidVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAwsVersionsResult) []string { return v.ValidVersions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAwsVersionsResultOutput{})
}
