// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package composer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides access to available Cloud Composer versions in a region for a given project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/composer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			all, err := composer.GetImageVersions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = composer.NewEnvironment(ctx, "test", &composer.EnvironmentArgs{
//				Region: pulumi.String("us-central1"),
//				Config: &composer.EnvironmentConfigArgs{
//					SoftwareConfig: &composer.EnvironmentConfigSoftwareConfigArgs{
//						ImageVersion: *pulumi.String(all.ImageVersions[0].ImageVersionId),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetImageVersions(ctx *pulumi.Context, args *GetImageVersionsArgs, opts ...pulumi.InvokeOption) (*GetImageVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetImageVersionsResult
	err := ctx.Invoke("gcp:composer/getImageVersions:getImageVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageVersions.
type GetImageVersionsArgs struct {
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location to list versions in.
	// If it is not provider, the provider region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getImageVersions.
type GetImageVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of composer image versions available in the given project and location. Each `imageVersion` contains:
	ImageVersions []GetImageVersionsImageVersion `pulumi:"imageVersions"`
	Project       string                         `pulumi:"project"`
	Region        string                         `pulumi:"region"`
}

func GetImageVersionsOutput(ctx *pulumi.Context, args GetImageVersionsOutputArgs, opts ...pulumi.InvokeOption) GetImageVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageVersionsResult, error) {
			args := v.(GetImageVersionsArgs)
			r, err := GetImageVersions(ctx, &args, opts...)
			var s GetImageVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageVersionsResultOutput)
}

// A collection of arguments for invoking getImageVersions.
type GetImageVersionsOutputArgs struct {
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The location to list versions in.
	// If it is not provider, the provider region is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetImageVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getImageVersions.
type GetImageVersionsResultOutput struct{ *pulumi.OutputState }

func (GetImageVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageVersionsResult)(nil)).Elem()
}

func (o GetImageVersionsResultOutput) ToGetImageVersionsResultOutput() GetImageVersionsResultOutput {
	return o
}

func (o GetImageVersionsResultOutput) ToGetImageVersionsResultOutputWithContext(ctx context.Context) GetImageVersionsResultOutput {
	return o
}

func (o GetImageVersionsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetImageVersionsResult] {
	return pulumix.Output[GetImageVersionsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetImageVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of composer image versions available in the given project and location. Each `imageVersion` contains:
func (o GetImageVersionsResultOutput) ImageVersions() GetImageVersionsImageVersionArrayOutput {
	return o.ApplyT(func(v GetImageVersionsResult) []GetImageVersionsImageVersion { return v.ImageVersions }).(GetImageVersionsImageVersionArrayOutput)
}

func (o GetImageVersionsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageVersionsResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetImageVersionsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageVersionsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageVersionsResultOutput{})
}
