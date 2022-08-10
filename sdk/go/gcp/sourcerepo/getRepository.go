// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get infomation about an existing Google Cloud Source Repository.
// For more information see [the official documentation](https://cloud.google.com/source-repositories)
// and
// [API](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sourcerepo"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sourcerepo.LookupRepository(ctx, &sourcerepo.LookupRepositoryArgs{
//				Name: "my-repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	var rv LookupRepositoryResult
	err := ctx.Invoke("gcp:sourcerepo/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryArgs struct {
	// Resource name of the repository. The repo name may contain slashes. eg, `name/with/slash`
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getRepository.
type LookupRepositoryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string                      `pulumi:"id"`
	Name          string                      `pulumi:"name"`
	Project       *string                     `pulumi:"project"`
	PubsubConfigs []GetRepositoryPubsubConfig `pulumi:"pubsubConfigs"`
	Size          int                         `pulumi:"size"`
	Url           string                      `pulumi:"url"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResult, error) {
			args := v.(LookupRepositoryArgs)
			r, err := LookupRepository(ctx, &args, opts...)
			var s LookupRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRepositoryResultOutput)
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryOutputArgs struct {
	// Resource name of the repository. The repo name may contain slashes. eg, `name/with/slash`
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRepository.
type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRepositoryResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupRepositoryResultOutput) PubsubConfigs() GetRepositoryPubsubConfigArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []GetRepositoryPubsubConfig { return v.PubsubConfigs }).(GetRepositoryPubsubConfigArrayOutput)
}

func (o LookupRepositoryResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryResult) int { return v.Size }).(pulumi.IntOutput)
}

func (o LookupRepositoryResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}
