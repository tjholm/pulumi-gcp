// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicenetworking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a private VPC connection with a GCP service provider. For more information see
// [the official documentation](https://cloud.google.com/vpc/docs/configure-private-services-access#creating-connection)
// and
// [API](https://cloud.google.com/service-infrastructure/docs/service-networking/reference/rest/v1/services.connections).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			peeringNetwork, err := compute.NewNetwork(ctx, "peeringNetwork", nil)
//			if err != nil {
//				return err
//			}
//			privateIpAlloc, err := compute.NewGlobalAddress(ctx, "privateIpAlloc", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(16),
//				Network:      peeringNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = servicenetworking.NewConnection(ctx, "default", &servicenetworking.ConnectionArgs{
//				Network: peeringNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					privateIpAlloc.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetworkPeeringRoutesConfig(ctx, "peeringRoutes", &compute.NetworkPeeringRoutesConfigArgs{
//				Peering:            _default.Peering,
//				Network:            peeringNetwork.Name,
//				ImportCustomRoutes: pulumi.Bool(true),
//				ExportCustomRoutes: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # ServiceNetworkingConnection can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:servicenetworking/connection:Connection peering_connection {{peering-network}}:{{service}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:servicenetworking/connection:Connection peering_connection /projects/{{project}}/global/networks/{{peering-network}}:{{service}}
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringOutput `pulumi:"network"`
	// (Computed) The name of the VPC Network Peering connection that was created by the service producer.
	Peering pulumi.StringOutput `pulumi:"peering"`
	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges pulumi.StringArrayOutput `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.ReservedPeeringRanges == nil {
		return nil, errors.New("invalid value for required argument 'ReservedPeeringRanges'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("gcp:servicenetworking/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("gcp:servicenetworking/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Name of VPC network connected with service producers using VPC peering.
	Network *string `pulumi:"network"`
	// (Computed) The name of the VPC Network Peering connection that was created by the service producer.
	Peering *string `pulumi:"peering"`
	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges []string `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service *string `pulumi:"service"`
}

type ConnectionState struct {
	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringPtrInput
	// (Computed) The name of the VPC Network Peering connection that was created by the service producer.
	Peering pulumi.StringPtrInput
	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges pulumi.StringArrayInput
	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Name of VPC network connected with service producers using VPC peering.
	Network string `pulumi:"network"`
	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges []string `pulumi:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Name of VPC network connected with service producers using VPC peering.
	Network pulumi.StringInput
	// Named IP address range(s) of PEERING type reserved for
	// this service provider. Note that invoking this method with a different range when connection
	// is already established will not reallocate already provisioned service producer subnetworks.
	ReservedPeeringRanges pulumi.StringArrayInput
	// Provider peering service that is managing peering connectivity for a
	// service provider organization. For Google services that support this functionality it is
	// 'servicenetworking.googleapis.com'.
	Service pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: i.ToConnectionOutputWithContext(ctx).OutputState,
	}
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

func (i ConnectionArray) ToOutput(ctx context.Context) pulumix.Output[[]*Connection] {
	return pulumix.Output[[]*Connection]{
		OutputState: i.ToConnectionArrayOutputWithContext(ctx).OutputState,
	}
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

func (i ConnectionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Connection] {
	return pulumix.Output[map[string]*Connection]{
		OutputState: i.ToConnectionMapOutputWithContext(ctx).OutputState,
	}
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*Connection] {
	return pulumix.Output[*Connection]{
		OutputState: o.OutputState,
	}
}

// Name of VPC network connected with service producers using VPC peering.
func (o ConnectionOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// (Computed) The name of the VPC Network Peering connection that was created by the service producer.
func (o ConnectionOutput) Peering() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Peering }).(pulumi.StringOutput)
}

// Named IP address range(s) of PEERING type reserved for
// this service provider. Note that invoking this method with a different range when connection
// is already established will not reallocate already provisioned service producer subnetworks.
func (o ConnectionOutput) ReservedPeeringRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringArrayOutput { return v.ReservedPeeringRanges }).(pulumi.StringArrayOutput)
}

// Provider peering service that is managing peering connectivity for a
// service provider organization. For Google services that support this functionality it is
// 'servicenetworking.googleapis.com'.
func (o ConnectionOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Connection] {
	return pulumix.Output[[]*Connection]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Connection] {
	return pulumix.Output[map[string]*Connection]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
