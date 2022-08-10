// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
//
// The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
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
//			debian, err := container.GetRegistryImage(ctx, &container.GetRegistryImageArgs{
//				Name: "debian",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("gcrLocation", debian.ImageUrl)
//			return nil
//		})
//	}
//
// ```
func GetRegistryImage(ctx *pulumi.Context, args *GetRegistryImageArgs, opts ...pulumi.InvokeOption) (*GetRegistryImageResult, error) {
	var rv GetRegistryImageResult
	err := ctx.Invoke("gcp:container/getRegistryImage:getRegistryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryImage.
type GetRegistryImageArgs struct {
	Digest  *string `pulumi:"digest"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	Tag     *string `pulumi:"tag"`
}

// A collection of values returned by getRegistryImage.
type GetRegistryImageResult struct {
	Digest *string `pulumi:"digest"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	ImageUrl string  `pulumi:"imageUrl"`
	Name     string  `pulumi:"name"`
	Project  string  `pulumi:"project"`
	Region   *string `pulumi:"region"`
	Tag      *string `pulumi:"tag"`
}

func GetRegistryImageOutput(ctx *pulumi.Context, args GetRegistryImageOutputArgs, opts ...pulumi.InvokeOption) GetRegistryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryImageResult, error) {
			args := v.(GetRegistryImageArgs)
			r, err := GetRegistryImage(ctx, &args, opts...)
			var s GetRegistryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryImageResultOutput)
}

// A collection of arguments for invoking getRegistryImage.
type GetRegistryImageOutputArgs struct {
	Digest  pulumi.StringPtrInput `pulumi:"digest"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
	Tag     pulumi.StringPtrInput `pulumi:"tag"`
}

func (GetRegistryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryImageArgs)(nil)).Elem()
}

// A collection of values returned by getRegistryImage.
type GetRegistryImageResultOutput struct{ *pulumi.OutputState }

func (GetRegistryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryImageResult)(nil)).Elem()
}

func (o GetRegistryImageResultOutput) ToGetRegistryImageResultOutput() GetRegistryImageResultOutput {
	return o
}

func (o GetRegistryImageResultOutput) ToGetRegistryImageResultOutputWithContext(ctx context.Context) GetRegistryImageResultOutput {
	return o
}

func (o GetRegistryImageResultOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryImageResult) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegistryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegistryImageResultOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryImageResult) string { return v.ImageUrl }).(pulumi.StringOutput)
}

func (o GetRegistryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetRegistryImageResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryImageResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetRegistryImageResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryImageResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetRegistryImageResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryImageResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryImageResultOutput{})
}
