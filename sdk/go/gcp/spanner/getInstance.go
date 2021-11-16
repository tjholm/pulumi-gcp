// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a spanner instance from Google Cloud by its name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/spanner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := spanner.LookupInstance(ctx, &spanner.LookupInstanceArgs{
// 			Name: "bar",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("gcp:spanner/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	Config      *string `pulumi:"config"`
	DisplayName *string `pulumi:"displayName"`
	// The name of the spanner instance.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	Config       *string `pulumi:"config"`
	DisplayName  *string `pulumi:"displayName"`
	ForceDestroy bool    `pulumi:"forceDestroy"`
	// The provider-assigned unique ID for this managed resource.
	Id              string            `pulumi:"id"`
	Labels          map[string]string `pulumi:"labels"`
	Name            string            `pulumi:"name"`
	NumNodes        int               `pulumi:"numNodes"`
	ProcessingUnits int               `pulumi:"processingUnits"`
	Project         *string           `pulumi:"project"`
	State           string            `pulumi:"state"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type LookupInstanceOutputArgs struct {
	Config      pulumi.StringPtrInput `pulumi:"config"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The name of the spanner instance.
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Config }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) NumNodes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.NumNodes }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) ProcessingUnits() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.ProcessingUnits }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
