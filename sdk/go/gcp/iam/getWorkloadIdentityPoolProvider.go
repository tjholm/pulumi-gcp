// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a IAM workload identity provider from Google Cloud by its id.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = iam.LookupWorkloadIdentityPoolProvider(ctx, &iam.LookupWorkloadIdentityPoolProviderArgs{
//				WorkloadIdentityPoolId:         "foo-pool",
//				WorkloadIdentityPoolProviderId: "bar-provider",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWorkloadIdentityPoolProvider(ctx *pulumi.Context, args *LookupWorkloadIdentityPoolProviderArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadIdentityPoolProviderResult, error) {
	var rv LookupWorkloadIdentityPoolProviderResult
	err := ctx.Invoke("gcp:iam/getWorkloadIdentityPoolProvider:getWorkloadIdentityPoolProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkloadIdentityPoolProvider.
type LookupWorkloadIdentityPoolProviderArgs struct {
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The id of the pool which is the
	// final component of the pool resource name.
	WorkloadIdentityPoolId string `pulumi:"workloadIdentityPoolId"`
	// The id of the provider which is the
	// final component of the resource name.
	WorkloadIdentityPoolProviderId string `pulumi:"workloadIdentityPoolProviderId"`
}

// A collection of values returned by getWorkloadIdentityPoolProvider.
type LookupWorkloadIdentityPoolProviderResult struct {
	AttributeCondition string                              `pulumi:"attributeCondition"`
	AttributeMapping   map[string]string                   `pulumi:"attributeMapping"`
	Aws                []GetWorkloadIdentityPoolProviderAw `pulumi:"aws"`
	Description        string                              `pulumi:"description"`
	Disabled           bool                                `pulumi:"disabled"`
	DisplayName        string                              `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                                `pulumi:"id"`
	Name                           string                                `pulumi:"name"`
	Oidcs                          []GetWorkloadIdentityPoolProviderOidc `pulumi:"oidcs"`
	Project                        *string                               `pulumi:"project"`
	State                          string                                `pulumi:"state"`
	WorkloadIdentityPoolId         string                                `pulumi:"workloadIdentityPoolId"`
	WorkloadIdentityPoolProviderId string                                `pulumi:"workloadIdentityPoolProviderId"`
}

func LookupWorkloadIdentityPoolProviderOutput(ctx *pulumi.Context, args LookupWorkloadIdentityPoolProviderOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadIdentityPoolProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadIdentityPoolProviderResult, error) {
			args := v.(LookupWorkloadIdentityPoolProviderArgs)
			r, err := LookupWorkloadIdentityPoolProvider(ctx, &args, opts...)
			var s LookupWorkloadIdentityPoolProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadIdentityPoolProviderResultOutput)
}

// A collection of arguments for invoking getWorkloadIdentityPoolProvider.
type LookupWorkloadIdentityPoolProviderOutputArgs struct {
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The id of the pool which is the
	// final component of the pool resource name.
	WorkloadIdentityPoolId pulumi.StringInput `pulumi:"workloadIdentityPoolId"`
	// The id of the provider which is the
	// final component of the resource name.
	WorkloadIdentityPoolProviderId pulumi.StringInput `pulumi:"workloadIdentityPoolProviderId"`
}

func (LookupWorkloadIdentityPoolProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadIdentityPoolProviderArgs)(nil)).Elem()
}

// A collection of values returned by getWorkloadIdentityPoolProvider.
type LookupWorkloadIdentityPoolProviderResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadIdentityPoolProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadIdentityPoolProviderResult)(nil)).Elem()
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) ToLookupWorkloadIdentityPoolProviderResultOutput() LookupWorkloadIdentityPoolProviderResultOutput {
	return o
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) ToLookupWorkloadIdentityPoolProviderResultOutputWithContext(ctx context.Context) LookupWorkloadIdentityPoolProviderResultOutput {
	return o
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) AttributeCondition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.AttributeCondition }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) AttributeMapping() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) map[string]string { return v.AttributeMapping }).(pulumi.StringMapOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Aws() GetWorkloadIdentityPoolProviderAwArrayOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) []GetWorkloadIdentityPoolProviderAw { return v.Aws }).(GetWorkloadIdentityPoolProviderAwArrayOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWorkloadIdentityPoolProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Oidcs() GetWorkloadIdentityPoolProviderOidcArrayOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) []GetWorkloadIdentityPoolProviderOidc { return v.Oidcs }).(GetWorkloadIdentityPoolProviderOidcArrayOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) WorkloadIdentityPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.WorkloadIdentityPoolId }).(pulumi.StringOutput)
}

func (o LookupWorkloadIdentityPoolProviderResultOutput) WorkloadIdentityPoolProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadIdentityPoolProviderResult) string { return v.WorkloadIdentityPoolProviderId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadIdentityPoolProviderResultOutput{})
}
