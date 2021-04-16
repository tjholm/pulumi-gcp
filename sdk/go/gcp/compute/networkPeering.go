// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a network peering within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/networks).
//
// > Both network must create a peering with each other for the peering
// to be functional.
//
// > Subnets IP ranges across peered VPC networks cannot overlap.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewNetwork(ctx, "_default", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		other, err := compute.NewNetwork(ctx, "other", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkPeering(ctx, "peering1", &compute.NetworkPeeringArgs{
// 			Network:     _default.ID(),
// 			PeerNetwork: other.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkPeering(ctx, "peering2", &compute.NetworkPeeringArgs{
// 			Network:     other.ID(),
// 			PeerNetwork: _default.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// VPC network peerings can be imported using the name and project of the primary network the peering exists in and the name of the network peering
//
// ```sh
//  $ pulumi import gcp:compute/networkPeering:NetworkPeering peering_network project-name/network-name/peering-name
// ```
type NetworkPeering struct {
	pulumi.CustomResourceState

	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrOutput `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrOutput `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary network of the peering.
	Network pulumi.StringOutput `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringOutput `pulumi:"peerNetwork"`
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State pulumi.StringOutput `pulumi:"state"`
	// Details about the current state of the peering.
	StateDetails pulumi.StringOutput `pulumi:"stateDetails"`
}

// NewNetworkPeering registers a new resource with the given unique name, arguments, and options.
func NewNetworkPeering(ctx *pulumi.Context,
	name string, args *NetworkPeeringArgs, opts ...pulumi.ResourceOption) (*NetworkPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.PeerNetwork == nil {
		return nil, errors.New("invalid value for required argument 'PeerNetwork'")
	}
	var resource NetworkPeering
	err := ctx.RegisterResource("gcp:compute/networkPeering:NetworkPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPeering gets an existing NetworkPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPeeringState, opts ...pulumi.ResourceOption) (*NetworkPeering, error) {
	var resource NetworkPeering
	err := ctx.ReadResource("gcp:compute/networkPeering:NetworkPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPeering resources.
type networkPeeringState struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp *bool `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp *bool `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name *string `pulumi:"name"`
	// The primary network of the peering.
	Network *string `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork *string `pulumi:"peerNetwork"`
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State *string `pulumi:"state"`
	// Details about the current state of the peering.
	StateDetails *string `pulumi:"stateDetails"`
}

type NetworkPeeringState struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Name of the peering.
	Name pulumi.StringPtrInput
	// The primary network of the peering.
	Network pulumi.StringPtrInput
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringPtrInput
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State pulumi.StringPtrInput
	// Details about the current state of the peering.
	StateDetails pulumi.StringPtrInput
}

func (NetworkPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringState)(nil)).Elem()
}

type networkPeeringArgs struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp *bool `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp *bool `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name *string `pulumi:"name"`
	// The primary network of the peering.
	Network string `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork string `pulumi:"peerNetwork"`
}

// The set of arguments for constructing a NetworkPeering resource.
type NetworkPeeringArgs struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Name of the peering.
	Name pulumi.StringPtrInput
	// The primary network of the peering.
	Network pulumi.StringInput
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringInput
}

func (NetworkPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringArgs)(nil)).Elem()
}

type NetworkPeeringInput interface {
	pulumi.Input

	ToNetworkPeeringOutput() NetworkPeeringOutput
	ToNetworkPeeringOutputWithContext(ctx context.Context) NetworkPeeringOutput
}

func (*NetworkPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPeering)(nil))
}

func (i *NetworkPeering) ToNetworkPeeringOutput() NetworkPeeringOutput {
	return i.ToNetworkPeeringOutputWithContext(context.Background())
}

func (i *NetworkPeering) ToNetworkPeeringOutputWithContext(ctx context.Context) NetworkPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringOutput)
}

func (i *NetworkPeering) ToNetworkPeeringPtrOutput() NetworkPeeringPtrOutput {
	return i.ToNetworkPeeringPtrOutputWithContext(context.Background())
}

func (i *NetworkPeering) ToNetworkPeeringPtrOutputWithContext(ctx context.Context) NetworkPeeringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringPtrOutput)
}

type NetworkPeeringPtrInput interface {
	pulumi.Input

	ToNetworkPeeringPtrOutput() NetworkPeeringPtrOutput
	ToNetworkPeeringPtrOutputWithContext(ctx context.Context) NetworkPeeringPtrOutput
}

type networkPeeringPtrType NetworkPeeringArgs

func (*networkPeeringPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPeering)(nil))
}

func (i *networkPeeringPtrType) ToNetworkPeeringPtrOutput() NetworkPeeringPtrOutput {
	return i.ToNetworkPeeringPtrOutputWithContext(context.Background())
}

func (i *networkPeeringPtrType) ToNetworkPeeringPtrOutputWithContext(ctx context.Context) NetworkPeeringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringPtrOutput)
}

// NetworkPeeringArrayInput is an input type that accepts NetworkPeeringArray and NetworkPeeringArrayOutput values.
// You can construct a concrete instance of `NetworkPeeringArrayInput` via:
//
//          NetworkPeeringArray{ NetworkPeeringArgs{...} }
type NetworkPeeringArrayInput interface {
	pulumi.Input

	ToNetworkPeeringArrayOutput() NetworkPeeringArrayOutput
	ToNetworkPeeringArrayOutputWithContext(context.Context) NetworkPeeringArrayOutput
}

type NetworkPeeringArray []NetworkPeeringInput

func (NetworkPeeringArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkPeering)(nil))
}

func (i NetworkPeeringArray) ToNetworkPeeringArrayOutput() NetworkPeeringArrayOutput {
	return i.ToNetworkPeeringArrayOutputWithContext(context.Background())
}

func (i NetworkPeeringArray) ToNetworkPeeringArrayOutputWithContext(ctx context.Context) NetworkPeeringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringArrayOutput)
}

// NetworkPeeringMapInput is an input type that accepts NetworkPeeringMap and NetworkPeeringMapOutput values.
// You can construct a concrete instance of `NetworkPeeringMapInput` via:
//
//          NetworkPeeringMap{ "key": NetworkPeeringArgs{...} }
type NetworkPeeringMapInput interface {
	pulumi.Input

	ToNetworkPeeringMapOutput() NetworkPeeringMapOutput
	ToNetworkPeeringMapOutputWithContext(context.Context) NetworkPeeringMapOutput
}

type NetworkPeeringMap map[string]NetworkPeeringInput

func (NetworkPeeringMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkPeering)(nil))
}

func (i NetworkPeeringMap) ToNetworkPeeringMapOutput() NetworkPeeringMapOutput {
	return i.ToNetworkPeeringMapOutputWithContext(context.Background())
}

func (i NetworkPeeringMap) ToNetworkPeeringMapOutputWithContext(ctx context.Context) NetworkPeeringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPeeringMapOutput)
}

type NetworkPeeringOutput struct {
	*pulumi.OutputState
}

func (NetworkPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPeering)(nil))
}

func (o NetworkPeeringOutput) ToNetworkPeeringOutput() NetworkPeeringOutput {
	return o
}

func (o NetworkPeeringOutput) ToNetworkPeeringOutputWithContext(ctx context.Context) NetworkPeeringOutput {
	return o
}

func (o NetworkPeeringOutput) ToNetworkPeeringPtrOutput() NetworkPeeringPtrOutput {
	return o.ToNetworkPeeringPtrOutputWithContext(context.Background())
}

func (o NetworkPeeringOutput) ToNetworkPeeringPtrOutputWithContext(ctx context.Context) NetworkPeeringPtrOutput {
	return o.ApplyT(func(v NetworkPeering) *NetworkPeering {
		return &v
	}).(NetworkPeeringPtrOutput)
}

type NetworkPeeringPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkPeeringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPeering)(nil))
}

func (o NetworkPeeringPtrOutput) ToNetworkPeeringPtrOutput() NetworkPeeringPtrOutput {
	return o
}

func (o NetworkPeeringPtrOutput) ToNetworkPeeringPtrOutputWithContext(ctx context.Context) NetworkPeeringPtrOutput {
	return o
}

type NetworkPeeringArrayOutput struct{ *pulumi.OutputState }

func (NetworkPeeringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkPeering)(nil))
}

func (o NetworkPeeringArrayOutput) ToNetworkPeeringArrayOutput() NetworkPeeringArrayOutput {
	return o
}

func (o NetworkPeeringArrayOutput) ToNetworkPeeringArrayOutputWithContext(ctx context.Context) NetworkPeeringArrayOutput {
	return o
}

func (o NetworkPeeringArrayOutput) Index(i pulumi.IntInput) NetworkPeeringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkPeering {
		return vs[0].([]NetworkPeering)[vs[1].(int)]
	}).(NetworkPeeringOutput)
}

type NetworkPeeringMapOutput struct{ *pulumi.OutputState }

func (NetworkPeeringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkPeering)(nil))
}

func (o NetworkPeeringMapOutput) ToNetworkPeeringMapOutput() NetworkPeeringMapOutput {
	return o
}

func (o NetworkPeeringMapOutput) ToNetworkPeeringMapOutputWithContext(ctx context.Context) NetworkPeeringMapOutput {
	return o
}

func (o NetworkPeeringMapOutput) MapIndex(k pulumi.StringInput) NetworkPeeringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkPeering {
		return vs[0].(map[string]NetworkPeering)[vs[1].(string)]
	}).(NetworkPeeringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkPeeringOutput{})
	pulumi.RegisterOutputType(NetworkPeeringPtrOutput{})
	pulumi.RegisterOutputType(NetworkPeeringArrayOutput{})
	pulumi.RegisterOutputType(NetworkPeeringMapOutput{})
}
