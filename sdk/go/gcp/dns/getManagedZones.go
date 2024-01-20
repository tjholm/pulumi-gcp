// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetManagedZones(ctx *pulumi.Context, args *GetManagedZonesArgs, opts ...pulumi.InvokeOption) (*GetManagedZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetManagedZonesResult
	err := ctx.Invoke("gcp:dns/getManagedZones:getManagedZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedZones.
type GetManagedZonesArgs struct {
	ManagedZones []GetManagedZonesManagedZone `pulumi:"managedZones"`
	Project      *string                      `pulumi:"project"`
}

// A collection of values returned by getManagedZones.
type GetManagedZonesResult struct {
	Id           string                       `pulumi:"id"`
	ManagedZones []GetManagedZonesManagedZone `pulumi:"managedZones"`
	Project      *string                      `pulumi:"project"`
}

func GetManagedZonesOutput(ctx *pulumi.Context, args GetManagedZonesOutputArgs, opts ...pulumi.InvokeOption) GetManagedZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetManagedZonesResult, error) {
			args := v.(GetManagedZonesArgs)
			r, err := GetManagedZones(ctx, &args, opts...)
			var s GetManagedZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetManagedZonesResultOutput)
}

// A collection of arguments for invoking getManagedZones.
type GetManagedZonesOutputArgs struct {
	ManagedZones GetManagedZonesManagedZoneArrayInput `pulumi:"managedZones"`
	Project      pulumi.StringPtrInput                `pulumi:"project"`
}

func (GetManagedZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedZonesArgs)(nil)).Elem()
}

// A collection of values returned by getManagedZones.
type GetManagedZonesResultOutput struct{ *pulumi.OutputState }

func (GetManagedZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedZonesResult)(nil)).Elem()
}

func (o GetManagedZonesResultOutput) ToGetManagedZonesResultOutput() GetManagedZonesResultOutput {
	return o
}

func (o GetManagedZonesResultOutput) ToGetManagedZonesResultOutputWithContext(ctx context.Context) GetManagedZonesResultOutput {
	return o
}

func (o GetManagedZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetManagedZonesResultOutput) ManagedZones() GetManagedZonesManagedZoneArrayOutput {
	return o.ApplyT(func(v GetManagedZonesResult) []GetManagedZonesManagedZone { return v.ManagedZones }).(GetManagedZonesManagedZoneArrayOutput)
}

func (o GetManagedZonesResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetManagedZonesResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetManagedZonesResultOutput{})
}
