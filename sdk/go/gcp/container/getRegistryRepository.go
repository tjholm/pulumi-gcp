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
//			foo, err := container.GetRegistryRepository(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("gcrLocation", foo.RepositoryUrl)
//			return nil
//		})
//	}
//
// ```
func GetRegistryRepository(ctx *pulumi.Context, args *GetRegistryRepositoryArgs, opts ...pulumi.InvokeOption) (*GetRegistryRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRegistryRepositoryResult
	err := ctx.Invoke("gcp:container/getRegistryRepository:getRegistryRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryRepository.
type GetRegistryRepositoryArgs struct {
	// The project ID that this repository is attached to.  If not provided, provider project will be used instead.
	Project *string `pulumi:"project"`
	// The GCR region to use.  As of this writing, one of `asia`, `eu`, and `us`.  See [the documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling) for additional information.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getRegistryRepository.
type GetRegistryRepositoryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Project string  `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The URL at which the repository can be accessed.
	RepositoryUrl string `pulumi:"repositoryUrl"`
}

func GetRegistryRepositoryOutput(ctx *pulumi.Context, args GetRegistryRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetRegistryRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistryRepositoryResult, error) {
			args := v.(GetRegistryRepositoryArgs)
			r, err := GetRegistryRepository(ctx, &args, opts...)
			var s GetRegistryRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistryRepositoryResultOutput)
}

// A collection of arguments for invoking getRegistryRepository.
type GetRegistryRepositoryOutputArgs struct {
	// The project ID that this repository is attached to.  If not provided, provider project will be used instead.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The GCR region to use.  As of this writing, one of `asia`, `eu`, and `us`.  See [the documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling) for additional information.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetRegistryRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRegistryRepository.
type GetRegistryRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetRegistryRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistryRepositoryResult)(nil)).Elem()
}

func (o GetRegistryRepositoryResultOutput) ToGetRegistryRepositoryResultOutput() GetRegistryRepositoryResultOutput {
	return o
}

func (o GetRegistryRepositoryResultOutput) ToGetRegistryRepositoryResultOutputWithContext(ctx context.Context) GetRegistryRepositoryResultOutput {
	return o
}

func (o GetRegistryRepositoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRegistryRepositoryResult] {
	return pulumix.Output[GetRegistryRepositoryResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegistryRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegistryRepositoryResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryRepositoryResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetRegistryRepositoryResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistryRepositoryResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The URL at which the repository can be accessed.
func (o GetRegistryRepositoryResultOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegistryRepositoryResult) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistryRepositoryResultOutput{})
}
