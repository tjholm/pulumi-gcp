// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve information about a set of projects based on a filter. See the
// [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
// for more details.
//
// ## Example Usage
// ### Searching For Projects About To Be Deleted In An Org
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			my_org_projects, err := projects.GetProject(ctx, &projects.GetProjectArgs{
//				Filter: "parent.id:012345678910 lifecycleState:DELETE_REQUESTED",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = organizations.LookupProject(ctx, &organizations.LookupProjectArgs{
//				ProjectId: pulumi.StringRef(my_org_projects.Projects[0].ProjectId),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProject(ctx *pulumi.Context, args *GetProjectArgs, opts ...pulumi.InvokeOption) (*GetProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectResult
	err := ctx.Invoke("gcp:projects/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type GetProjectArgs struct {
	// A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
	Filter string `pulumi:"filter"`
}

// A collection of values returned by getProject.
type GetProjectResult struct {
	Filter string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of projects matching the provided filter. Structure is defined below.
	Projects []GetProjectProject `pulumi:"projects"`
}

func GetProjectOutput(ctx *pulumi.Context, args GetProjectOutputArgs, opts ...pulumi.InvokeOption) GetProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectResult, error) {
			args := v.(GetProjectArgs)
			r, err := GetProject(ctx, &args, opts...)
			var s GetProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type GetProjectOutputArgs struct {
	// A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
	Filter pulumi.StringInput `pulumi:"filter"`
}

func (GetProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type GetProjectResultOutput struct{ *pulumi.OutputState }

func (GetProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectResult)(nil)).Elem()
}

func (o GetProjectResultOutput) ToGetProjectResultOutput() GetProjectResultOutput {
	return o
}

func (o GetProjectResultOutput) ToGetProjectResultOutputWithContext(ctx context.Context) GetProjectResultOutput {
	return o
}

func (o GetProjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetProjectResult] {
	return pulumix.Output[GetProjectResult]{
		OutputState: o.OutputState,
	}
}

func (o GetProjectResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of projects matching the provided filter. Structure is defined below.
func (o GetProjectResultOutput) Projects() GetProjectProjectArrayOutput {
	return o.ApplyT(func(v GetProjectResult) []GetProjectProject { return v.Projects }).(GetProjectProjectArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectResultOutput{})
}
