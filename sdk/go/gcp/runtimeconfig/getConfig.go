// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfig(ctx *pulumi.Context, args *LookupConfigArgs, opts ...pulumi.InvokeOption) (*LookupConfigResult, error) {
	var rv LookupConfigResult
	err := ctx.Invoke("gcp:runtimeconfig/getConfig:getConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfig.
type LookupConfigArgs struct {
	// The name of the Runtime Configurator configuration.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getConfig.
type LookupConfigResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

func LookupConfigOutput(ctx *pulumi.Context, args LookupConfigOutputArgs, opts ...pulumi.InvokeOption) LookupConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigResult, error) {
			args := v.(LookupConfigArgs)
			r, err := LookupConfig(ctx, &args, opts...)
			return *r, err
		}).(LookupConfigResultOutput)
}

// A collection of arguments for invoking getConfig.
type LookupConfigOutputArgs struct {
	// The name of the Runtime Configurator configuration.
	Name pulumi.StringInput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigArgs)(nil)).Elem()
}

// A collection of values returned by getConfig.
type LookupConfigResultOutput struct{ *pulumi.OutputState }

func (LookupConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigResult)(nil)).Elem()
}

func (o LookupConfigResultOutput) ToLookupConfigResultOutput() LookupConfigResultOutput {
	return o
}

func (o LookupConfigResultOutput) ToLookupConfigResultOutputWithContext(ctx context.Context) LookupConfigResultOutput {
	return o
}

func (o LookupConfigResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigResultOutput{})
}
