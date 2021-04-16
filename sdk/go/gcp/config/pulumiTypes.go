// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Batching struct {
	EnableBatching *bool   `pulumi:"enableBatching"`
	SendAfter      *string `pulumi:"sendAfter"`
}

// BatchingInput is an input type that accepts BatchingArgs and BatchingOutput values.
// You can construct a concrete instance of `BatchingInput` via:
//
//          BatchingArgs{...}
type BatchingInput interface {
	pulumi.Input

	ToBatchingOutput() BatchingOutput
	ToBatchingOutputWithContext(context.Context) BatchingOutput
}

type BatchingArgs struct {
	EnableBatching pulumi.BoolPtrInput   `pulumi:"enableBatching"`
	SendAfter      pulumi.StringPtrInput `pulumi:"sendAfter"`
}

func (BatchingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Batching)(nil)).Elem()
}

func (i BatchingArgs) ToBatchingOutput() BatchingOutput {
	return i.ToBatchingOutputWithContext(context.Background())
}

func (i BatchingArgs) ToBatchingOutputWithContext(ctx context.Context) BatchingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchingOutput)
}

type BatchingOutput struct{ *pulumi.OutputState }

func (BatchingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Batching)(nil)).Elem()
}

func (o BatchingOutput) ToBatchingOutput() BatchingOutput {
	return o
}

func (o BatchingOutput) ToBatchingOutputWithContext(ctx context.Context) BatchingOutput {
	return o
}

func (o BatchingOutput) EnableBatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Batching) *bool { return v.EnableBatching }).(pulumi.BoolPtrOutput)
}

func (o BatchingOutput) SendAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Batching) *string { return v.SendAfter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchingOutput{})
}
