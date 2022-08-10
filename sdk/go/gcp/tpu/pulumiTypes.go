// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tpu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodeNetworkEndpoint struct {
	IpAddress *string `pulumi:"ipAddress"`
	Port      *int    `pulumi:"port"`
}

// NodeNetworkEndpointInput is an input type that accepts NodeNetworkEndpointArgs and NodeNetworkEndpointOutput values.
// You can construct a concrete instance of `NodeNetworkEndpointInput` via:
//
//	NodeNetworkEndpointArgs{...}
type NodeNetworkEndpointInput interface {
	pulumi.Input

	ToNodeNetworkEndpointOutput() NodeNetworkEndpointOutput
	ToNodeNetworkEndpointOutputWithContext(context.Context) NodeNetworkEndpointOutput
}

type NodeNetworkEndpointArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	Port      pulumi.IntPtrInput    `pulumi:"port"`
}

func (NodeNetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeNetworkEndpoint)(nil)).Elem()
}

func (i NodeNetworkEndpointArgs) ToNodeNetworkEndpointOutput() NodeNetworkEndpointOutput {
	return i.ToNodeNetworkEndpointOutputWithContext(context.Background())
}

func (i NodeNetworkEndpointArgs) ToNodeNetworkEndpointOutputWithContext(ctx context.Context) NodeNetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeNetworkEndpointOutput)
}

// NodeNetworkEndpointArrayInput is an input type that accepts NodeNetworkEndpointArray and NodeNetworkEndpointArrayOutput values.
// You can construct a concrete instance of `NodeNetworkEndpointArrayInput` via:
//
//	NodeNetworkEndpointArray{ NodeNetworkEndpointArgs{...} }
type NodeNetworkEndpointArrayInput interface {
	pulumi.Input

	ToNodeNetworkEndpointArrayOutput() NodeNetworkEndpointArrayOutput
	ToNodeNetworkEndpointArrayOutputWithContext(context.Context) NodeNetworkEndpointArrayOutput
}

type NodeNetworkEndpointArray []NodeNetworkEndpointInput

func (NodeNetworkEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeNetworkEndpoint)(nil)).Elem()
}

func (i NodeNetworkEndpointArray) ToNodeNetworkEndpointArrayOutput() NodeNetworkEndpointArrayOutput {
	return i.ToNodeNetworkEndpointArrayOutputWithContext(context.Background())
}

func (i NodeNetworkEndpointArray) ToNodeNetworkEndpointArrayOutputWithContext(ctx context.Context) NodeNetworkEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeNetworkEndpointArrayOutput)
}

type NodeNetworkEndpointOutput struct{ *pulumi.OutputState }

func (NodeNetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeNetworkEndpoint)(nil)).Elem()
}

func (o NodeNetworkEndpointOutput) ToNodeNetworkEndpointOutput() NodeNetworkEndpointOutput {
	return o
}

func (o NodeNetworkEndpointOutput) ToNodeNetworkEndpointOutputWithContext(ctx context.Context) NodeNetworkEndpointOutput {
	return o
}

func (o NodeNetworkEndpointOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeNetworkEndpoint) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o NodeNetworkEndpointOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NodeNetworkEndpoint) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type NodeNetworkEndpointArrayOutput struct{ *pulumi.OutputState }

func (NodeNetworkEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeNetworkEndpoint)(nil)).Elem()
}

func (o NodeNetworkEndpointArrayOutput) ToNodeNetworkEndpointArrayOutput() NodeNetworkEndpointArrayOutput {
	return o
}

func (o NodeNetworkEndpointArrayOutput) ToNodeNetworkEndpointArrayOutputWithContext(ctx context.Context) NodeNetworkEndpointArrayOutput {
	return o
}

func (o NodeNetworkEndpointArrayOutput) Index(i pulumi.IntInput) NodeNetworkEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeNetworkEndpoint {
		return vs[0].([]NodeNetworkEndpoint)[vs[1].(int)]
	}).(NodeNetworkEndpointOutput)
}

type NodeSchedulingConfig struct {
	// Defines whether the TPU instance is preemptible.
	Preemptible bool `pulumi:"preemptible"`
}

// NodeSchedulingConfigInput is an input type that accepts NodeSchedulingConfigArgs and NodeSchedulingConfigOutput values.
// You can construct a concrete instance of `NodeSchedulingConfigInput` via:
//
//	NodeSchedulingConfigArgs{...}
type NodeSchedulingConfigInput interface {
	pulumi.Input

	ToNodeSchedulingConfigOutput() NodeSchedulingConfigOutput
	ToNodeSchedulingConfigOutputWithContext(context.Context) NodeSchedulingConfigOutput
}

type NodeSchedulingConfigArgs struct {
	// Defines whether the TPU instance is preemptible.
	Preemptible pulumi.BoolInput `pulumi:"preemptible"`
}

func (NodeSchedulingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSchedulingConfig)(nil)).Elem()
}

func (i NodeSchedulingConfigArgs) ToNodeSchedulingConfigOutput() NodeSchedulingConfigOutput {
	return i.ToNodeSchedulingConfigOutputWithContext(context.Background())
}

func (i NodeSchedulingConfigArgs) ToNodeSchedulingConfigOutputWithContext(ctx context.Context) NodeSchedulingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeSchedulingConfigOutput)
}

func (i NodeSchedulingConfigArgs) ToNodeSchedulingConfigPtrOutput() NodeSchedulingConfigPtrOutput {
	return i.ToNodeSchedulingConfigPtrOutputWithContext(context.Background())
}

func (i NodeSchedulingConfigArgs) ToNodeSchedulingConfigPtrOutputWithContext(ctx context.Context) NodeSchedulingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeSchedulingConfigOutput).ToNodeSchedulingConfigPtrOutputWithContext(ctx)
}

// NodeSchedulingConfigPtrInput is an input type that accepts NodeSchedulingConfigArgs, NodeSchedulingConfigPtr and NodeSchedulingConfigPtrOutput values.
// You can construct a concrete instance of `NodeSchedulingConfigPtrInput` via:
//
//	        NodeSchedulingConfigArgs{...}
//
//	or:
//
//	        nil
type NodeSchedulingConfigPtrInput interface {
	pulumi.Input

	ToNodeSchedulingConfigPtrOutput() NodeSchedulingConfigPtrOutput
	ToNodeSchedulingConfigPtrOutputWithContext(context.Context) NodeSchedulingConfigPtrOutput
}

type nodeSchedulingConfigPtrType NodeSchedulingConfigArgs

func NodeSchedulingConfigPtr(v *NodeSchedulingConfigArgs) NodeSchedulingConfigPtrInput {
	return (*nodeSchedulingConfigPtrType)(v)
}

func (*nodeSchedulingConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeSchedulingConfig)(nil)).Elem()
}

func (i *nodeSchedulingConfigPtrType) ToNodeSchedulingConfigPtrOutput() NodeSchedulingConfigPtrOutput {
	return i.ToNodeSchedulingConfigPtrOutputWithContext(context.Background())
}

func (i *nodeSchedulingConfigPtrType) ToNodeSchedulingConfigPtrOutputWithContext(ctx context.Context) NodeSchedulingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeSchedulingConfigPtrOutput)
}

type NodeSchedulingConfigOutput struct{ *pulumi.OutputState }

func (NodeSchedulingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSchedulingConfig)(nil)).Elem()
}

func (o NodeSchedulingConfigOutput) ToNodeSchedulingConfigOutput() NodeSchedulingConfigOutput {
	return o
}

func (o NodeSchedulingConfigOutput) ToNodeSchedulingConfigOutputWithContext(ctx context.Context) NodeSchedulingConfigOutput {
	return o
}

func (o NodeSchedulingConfigOutput) ToNodeSchedulingConfigPtrOutput() NodeSchedulingConfigPtrOutput {
	return o.ToNodeSchedulingConfigPtrOutputWithContext(context.Background())
}

func (o NodeSchedulingConfigOutput) ToNodeSchedulingConfigPtrOutputWithContext(ctx context.Context) NodeSchedulingConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeSchedulingConfig) *NodeSchedulingConfig {
		return &v
	}).(NodeSchedulingConfigPtrOutput)
}

// Defines whether the TPU instance is preemptible.
func (o NodeSchedulingConfigOutput) Preemptible() pulumi.BoolOutput {
	return o.ApplyT(func(v NodeSchedulingConfig) bool { return v.Preemptible }).(pulumi.BoolOutput)
}

type NodeSchedulingConfigPtrOutput struct{ *pulumi.OutputState }

func (NodeSchedulingConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeSchedulingConfig)(nil)).Elem()
}

func (o NodeSchedulingConfigPtrOutput) ToNodeSchedulingConfigPtrOutput() NodeSchedulingConfigPtrOutput {
	return o
}

func (o NodeSchedulingConfigPtrOutput) ToNodeSchedulingConfigPtrOutputWithContext(ctx context.Context) NodeSchedulingConfigPtrOutput {
	return o
}

func (o NodeSchedulingConfigPtrOutput) Elem() NodeSchedulingConfigOutput {
	return o.ApplyT(func(v *NodeSchedulingConfig) NodeSchedulingConfig {
		if v != nil {
			return *v
		}
		var ret NodeSchedulingConfig
		return ret
	}).(NodeSchedulingConfigOutput)
}

// Defines whether the TPU instance is preemptible.
func (o NodeSchedulingConfigPtrOutput) Preemptible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeSchedulingConfig) *bool {
		if v == nil {
			return nil
		}
		return &v.Preemptible
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeNetworkEndpointInput)(nil)).Elem(), NodeNetworkEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeNetworkEndpointArrayInput)(nil)).Elem(), NodeNetworkEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSchedulingConfigInput)(nil)).Elem(), NodeSchedulingConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSchedulingConfigPtrInput)(nil)).Elem(), NodeSchedulingConfigArgs{})
	pulumi.RegisterOutputType(NodeNetworkEndpointOutput{})
	pulumi.RegisterOutputType(NodeNetworkEndpointArrayOutput{})
	pulumi.RegisterOutputType(NodeSchedulingConfigOutput{})
	pulumi.RegisterOutputType(NodeSchedulingConfigPtrOutput{})
}
