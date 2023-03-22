// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Apigee NAT (network address translation) address. A NAT address is a static external IP address used for Internet egress traffic. This is not avaible for Apigee hybrid.
// Apigee NAT addresses are not automatically activated because they might require explicit allow entries on the target systems first. See https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances.natAddresses/activate
//
// To get more information about NatAddress, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances.natAddresses)
// * How-to Guides
//   - [Provisioning NAT IPs](https://cloud.google.com/apigee/docs/api-platform/security/nat-provisioning)
//
// ## Example Usage
//
// ## Import
//
// # NatAddress can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigee/natAddress:NatAddress default {{instance_id}}/natAddresses/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigee/natAddress:NatAddress default {{instance_id}}/{{name}}
//
// ```
type NatAddress struct {
	pulumi.CustomResourceState

	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The allocated NAT IP address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Resource ID of the NAT address.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the NAT IP address.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewNatAddress registers a new resource with the given unique name, arguments, and options.
func NewNatAddress(ctx *pulumi.Context,
	name string, args *NatAddressArgs, opts ...pulumi.ResourceOption) (*NatAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource NatAddress
	err := ctx.RegisterResource("gcp:apigee/natAddress:NatAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatAddress gets an existing NatAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatAddressState, opts ...pulumi.ResourceOption) (*NatAddress, error) {
	var resource NatAddress
	err := ctx.ReadResource("gcp:apigee/natAddress:NatAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatAddress resources.
type natAddressState struct {
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	InstanceId *string `pulumi:"instanceId"`
	// The allocated NAT IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// Resource ID of the NAT address.
	Name *string `pulumi:"name"`
	// State of the NAT IP address.
	State *string `pulumi:"state"`
}

type NatAddressState struct {
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	InstanceId pulumi.StringPtrInput
	// The allocated NAT IP address.
	IpAddress pulumi.StringPtrInput
	// Resource ID of the NAT address.
	Name pulumi.StringPtrInput
	// State of the NAT IP address.
	State pulumi.StringPtrInput
}

func (NatAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*natAddressState)(nil)).Elem()
}

type natAddressArgs struct {
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	InstanceId string `pulumi:"instanceId"`
	// Resource ID of the NAT address.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a NatAddress resource.
type NatAddressArgs struct {
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	InstanceId pulumi.StringInput
	// Resource ID of the NAT address.
	Name pulumi.StringPtrInput
}

func (NatAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natAddressArgs)(nil)).Elem()
}

type NatAddressInput interface {
	pulumi.Input

	ToNatAddressOutput() NatAddressOutput
	ToNatAddressOutputWithContext(ctx context.Context) NatAddressOutput
}

func (*NatAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**NatAddress)(nil)).Elem()
}

func (i *NatAddress) ToNatAddressOutput() NatAddressOutput {
	return i.ToNatAddressOutputWithContext(context.Background())
}

func (i *NatAddress) ToNatAddressOutputWithContext(ctx context.Context) NatAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatAddressOutput)
}

// NatAddressArrayInput is an input type that accepts NatAddressArray and NatAddressArrayOutput values.
// You can construct a concrete instance of `NatAddressArrayInput` via:
//
//	NatAddressArray{ NatAddressArgs{...} }
type NatAddressArrayInput interface {
	pulumi.Input

	ToNatAddressArrayOutput() NatAddressArrayOutput
	ToNatAddressArrayOutputWithContext(context.Context) NatAddressArrayOutput
}

type NatAddressArray []NatAddressInput

func (NatAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatAddress)(nil)).Elem()
}

func (i NatAddressArray) ToNatAddressArrayOutput() NatAddressArrayOutput {
	return i.ToNatAddressArrayOutputWithContext(context.Background())
}

func (i NatAddressArray) ToNatAddressArrayOutputWithContext(ctx context.Context) NatAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatAddressArrayOutput)
}

// NatAddressMapInput is an input type that accepts NatAddressMap and NatAddressMapOutput values.
// You can construct a concrete instance of `NatAddressMapInput` via:
//
//	NatAddressMap{ "key": NatAddressArgs{...} }
type NatAddressMapInput interface {
	pulumi.Input

	ToNatAddressMapOutput() NatAddressMapOutput
	ToNatAddressMapOutputWithContext(context.Context) NatAddressMapOutput
}

type NatAddressMap map[string]NatAddressInput

func (NatAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatAddress)(nil)).Elem()
}

func (i NatAddressMap) ToNatAddressMapOutput() NatAddressMapOutput {
	return i.ToNatAddressMapOutputWithContext(context.Background())
}

func (i NatAddressMap) ToNatAddressMapOutputWithContext(ctx context.Context) NatAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatAddressMapOutput)
}

type NatAddressOutput struct{ *pulumi.OutputState }

func (NatAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatAddress)(nil)).Elem()
}

func (o NatAddressOutput) ToNatAddressOutput() NatAddressOutput {
	return o
}

func (o NatAddressOutput) ToNatAddressOutputWithContext(ctx context.Context) NatAddressOutput {
	return o
}

// The Apigee instance associated with the Apigee environment,
// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
func (o NatAddressOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatAddress) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The allocated NAT IP address.
func (o NatAddressOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NatAddress) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Resource ID of the NAT address.
func (o NatAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NatAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the NAT IP address.
func (o NatAddressOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NatAddress) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type NatAddressArrayOutput struct{ *pulumi.OutputState }

func (NatAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatAddress)(nil)).Elem()
}

func (o NatAddressArrayOutput) ToNatAddressArrayOutput() NatAddressArrayOutput {
	return o
}

func (o NatAddressArrayOutput) ToNatAddressArrayOutputWithContext(ctx context.Context) NatAddressArrayOutput {
	return o
}

func (o NatAddressArrayOutput) Index(i pulumi.IntInput) NatAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NatAddress {
		return vs[0].([]*NatAddress)[vs[1].(int)]
	}).(NatAddressOutput)
}

type NatAddressMapOutput struct{ *pulumi.OutputState }

func (NatAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatAddress)(nil)).Elem()
}

func (o NatAddressMapOutput) ToNatAddressMapOutput() NatAddressMapOutput {
	return o
}

func (o NatAddressMapOutput) ToNatAddressMapOutputWithContext(ctx context.Context) NatAddressMapOutput {
	return o
}

func (o NatAddressMapOutput) MapIndex(k pulumi.StringInput) NatAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NatAddress {
		return vs[0].(map[string]*NatAddress)[vs[1].(string)]
	}).(NatAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatAddressInput)(nil)).Elem(), &NatAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatAddressArrayInput)(nil)).Elem(), NatAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatAddressMapInput)(nil)).Elem(), NatAddressMap{})
	pulumi.RegisterOutputType(NatAddressOutput{})
	pulumi.RegisterOutputType(NatAddressArrayOutput{})
	pulumi.RegisterOutputType(NatAddressMapOutput{})
}
