// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a VPN gateway running in GCP. This virtual device is managed
// by Google, but used only by you. This type of VPN Gateway allows for the creation
// of VPN solutions with higher availability than classic Target VPN Gateways.
//
// To get more information about HaVpnGateway, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways)
// * How-to Guides
//     * [Choosing a VPN](https://cloud.google.com/vpn/docs/how-to/choosing-a-vpn)
//     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
//
// ## Example Usage
// ### Ha Vpn Gateway Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network1, err := compute.NewNetwork(ctx, "network1", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewHaVpnGateway(ctx, "haGateway1", &compute.HaVpnGatewayArgs{
// 			Region:  pulumi.String("us-central1"),
// 			Network: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Ha Vpn Gateway Gcp To Gcp
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network1, err := compute.NewNetwork(ctx, "network1", &compute.NetworkArgs{
// 			RoutingMode:           pulumi.String("GLOBAL"),
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		haGateway1, err := compute.NewHaVpnGateway(ctx, "haGateway1", &compute.HaVpnGatewayArgs{
// 			Region:  pulumi.String("us-central1"),
// 			Network: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		network2, err := compute.NewNetwork(ctx, "network2", &compute.NetworkArgs{
// 			RoutingMode:           pulumi.String("GLOBAL"),
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		haGateway2, err := compute.NewHaVpnGateway(ctx, "haGateway2", &compute.HaVpnGatewayArgs{
// 			Region:  pulumi.String("us-central1"),
// 			Network: network2.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network1Subnet1", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.1.0/24"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network1Subnet2", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.2.0/24"),
// 			Region:      pulumi.String("us-west1"),
// 			Network:     network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network2Subnet1", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("192.168.1.0/24"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     network2.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network2Subnet2", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("192.168.2.0/24"),
// 			Region:      pulumi.String("us-east1"),
// 			Network:     network2.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1, err := compute.NewRouter(ctx, "router1", &compute.RouterArgs{
// 			Network: network1.Name,
// 			Bgp: &compute.RouterBgpArgs{
// 				Asn: pulumi.Int(64514),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router2, err := compute.NewRouter(ctx, "router2", &compute.RouterArgs{
// 			Network: network2.Name,
// 			Bgp: &compute.RouterBgpArgs{
// 				Asn: pulumi.Int(64515),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel1, err := compute.NewVPNTunnel(ctx, "tunnel1", &compute.VPNTunnelArgs{
// 			Region:              pulumi.String("us-central1"),
// 			VpnGateway:          haGateway1.ID(),
// 			PeerGcpGateway:      haGateway2.ID(),
// 			SharedSecret:        pulumi.String("a secret message"),
// 			Router:              router1.ID(),
// 			VpnGatewayInterface: pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel2, err := compute.NewVPNTunnel(ctx, "tunnel2", &compute.VPNTunnelArgs{
// 			Region:              pulumi.String("us-central1"),
// 			VpnGateway:          haGateway1.ID(),
// 			PeerGcpGateway:      haGateway2.ID(),
// 			SharedSecret:        pulumi.String("a secret message"),
// 			Router:              router1.ID(),
// 			VpnGatewayInterface: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel3, err := compute.NewVPNTunnel(ctx, "tunnel3", &compute.VPNTunnelArgs{
// 			Region:              pulumi.String("us-central1"),
// 			VpnGateway:          haGateway2.ID(),
// 			PeerGcpGateway:      haGateway1.ID(),
// 			SharedSecret:        pulumi.String("a secret message"),
// 			Router:              router2.ID(),
// 			VpnGatewayInterface: pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel4, err := compute.NewVPNTunnel(ctx, "tunnel4", &compute.VPNTunnelArgs{
// 			Region:              pulumi.String("us-central1"),
// 			VpnGateway:          haGateway2.ID(),
// 			PeerGcpGateway:      haGateway1.ID(),
// 			SharedSecret:        pulumi.String("a secret message"),
// 			Router:              router2.ID(),
// 			VpnGatewayInterface: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1Interface1, err := compute.NewRouterInterface(ctx, "router1Interface1", &compute.RouterInterfaceArgs{
// 			Router:    router1.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.0.1/30"),
// 			VpnTunnel: tunnel1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router1Peer1", &compute.RouterPeerArgs{
// 			Router:                  router1.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.0.2"),
// 			PeerAsn:                 pulumi.Int(64515),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router1Interface1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1Interface2, err := compute.NewRouterInterface(ctx, "router1Interface2", &compute.RouterInterfaceArgs{
// 			Router:    router1.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.1.2/30"),
// 			VpnTunnel: tunnel2.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router1Peer2", &compute.RouterPeerArgs{
// 			Router:                  router1.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.1.1"),
// 			PeerAsn:                 pulumi.Int(64515),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router1Interface2.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router2Interface1, err := compute.NewRouterInterface(ctx, "router2Interface1", &compute.RouterInterfaceArgs{
// 			Router:    router2.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.0.2/30"),
// 			VpnTunnel: tunnel3.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router2Peer1", &compute.RouterPeerArgs{
// 			Router:                  router2.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.0.1"),
// 			PeerAsn:                 pulumi.Int(64514),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router2Interface1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router2Interface2, err := compute.NewRouterInterface(ctx, "router2Interface2", &compute.RouterInterfaceArgs{
// 			Router:    router2.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.1.1/30"),
// 			VpnTunnel: tunnel4.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router2Peer2", &compute.RouterPeerArgs{
// 			Router:                  router2.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.1.2"),
// 			PeerAsn:                 pulumi.Int(64514),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router2Interface2.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Compute Ha Vpn Gateway Encrypted Interconnect
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		address1, err := compute.NewAddress(ctx, "address1", &compute.AddressArgs{
// 			AddressType:  pulumi.String("INTERNAL"),
// 			Purpose:      pulumi.String("IPSEC_INTERCONNECT"),
// 			Address:      pulumi.String("192.168.1.0"),
// 			PrefixLength: pulumi.Int(29),
// 			Network:      network.SelfLink,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router, err := compute.NewRouter(ctx, "router", &compute.RouterArgs{
// 			Network:                     network.Name,
// 			EncryptedInterconnectRouter: pulumi.Bool(true),
// 			Bgp: &compute.RouterBgpArgs{
// 				Asn: pulumi.Int(16550),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		attachment1, err := compute.NewInterconnectAttachment(ctx, "attachment1", &compute.InterconnectAttachmentArgs{
// 			EdgeAvailabilityDomain: pulumi.String("AVAILABILITY_DOMAIN_1"),
// 			Type:                   pulumi.String("PARTNER"),
// 			Router:                 router.ID(),
// 			Encryption:             pulumi.String("IPSEC"),
// 			IpsecInternalAddresses: pulumi.StringArray{
// 				address1.SelfLink,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		address2, err := compute.NewAddress(ctx, "address2", &compute.AddressArgs{
// 			AddressType:  pulumi.String("INTERNAL"),
// 			Purpose:      pulumi.String("IPSEC_INTERCONNECT"),
// 			Address:      pulumi.String("192.168.2.0"),
// 			PrefixLength: pulumi.Int(29),
// 			Network:      network.SelfLink,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		attachment2, err := compute.NewInterconnectAttachment(ctx, "attachment2", &compute.InterconnectAttachmentArgs{
// 			EdgeAvailabilityDomain: pulumi.String("AVAILABILITY_DOMAIN_2"),
// 			Type:                   pulumi.String("PARTNER"),
// 			Router:                 router.ID(),
// 			Encryption:             pulumi.String("IPSEC"),
// 			IpsecInternalAddresses: pulumi.StringArray{
// 				address2.SelfLink,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewHaVpnGateway(ctx, "vpn-gateway", &compute.HaVpnGatewayArgs{
// 			Network: network.ID(),
// 			VpnInterfaces: compute.HaVpnGatewayVpnInterfaceArray{
// 				&compute.HaVpnGatewayVpnInterfaceArgs{
// 					Id:                     pulumi.Int(0),
// 					InterconnectAttachment: attachment1.SelfLink,
// 				},
// 				&compute.HaVpnGatewayVpnInterfaceArgs{
// 					Id:                     pulumi.Int(1),
// 					InterconnectAttachment: attachment2.SelfLink,
// 				},
// 			},
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
// HaVpnGateway can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{name}}
// ```
type HaVpnGateway struct {
	pulumi.CustomResourceState

	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region this gateway should sit in.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VpnInterfaces HaVpnGatewayVpnInterfaceArrayOutput `pulumi:"vpnInterfaces"`
}

// NewHaVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewHaVpnGateway(ctx *pulumi.Context,
	name string, args *HaVpnGatewayArgs, opts ...pulumi.ResourceOption) (*HaVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	var resource HaVpnGateway
	err := ctx.RegisterResource("gcp:compute/haVpnGateway:HaVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHaVpnGateway gets an existing HaVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HaVpnGatewayState, opts ...pulumi.ResourceOption) (*HaVpnGateway, error) {
	var resource HaVpnGateway
	err := ctx.ReadResource("gcp:compute/haVpnGateway:HaVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HaVpnGateway resources.
type haVpnGatewayState struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VpnInterfaces []HaVpnGatewayVpnInterface `pulumi:"vpnInterfaces"`
}

type HaVpnGatewayState struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VpnInterfaces HaVpnGatewayVpnInterfaceArrayInput
}

func (HaVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*haVpnGatewayState)(nil)).Elem()
}

type haVpnGatewayArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region *string `pulumi:"region"`
	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VpnInterfaces []HaVpnGatewayVpnInterface `pulumi:"vpnInterfaces"`
}

// The set of arguments for constructing a HaVpnGateway resource.
type HaVpnGatewayArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region pulumi.StringPtrInput
	// A list of interfaces on this VPN gateway.
	// Structure is documented below.
	VpnInterfaces HaVpnGatewayVpnInterfaceArrayInput
}

func (HaVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*haVpnGatewayArgs)(nil)).Elem()
}

type HaVpnGatewayInput interface {
	pulumi.Input

	ToHaVpnGatewayOutput() HaVpnGatewayOutput
	ToHaVpnGatewayOutputWithContext(ctx context.Context) HaVpnGatewayOutput
}

func (*HaVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVpnGateway)(nil)).Elem()
}

func (i *HaVpnGateway) ToHaVpnGatewayOutput() HaVpnGatewayOutput {
	return i.ToHaVpnGatewayOutputWithContext(context.Background())
}

func (i *HaVpnGateway) ToHaVpnGatewayOutputWithContext(ctx context.Context) HaVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVpnGatewayOutput)
}

// HaVpnGatewayArrayInput is an input type that accepts HaVpnGatewayArray and HaVpnGatewayArrayOutput values.
// You can construct a concrete instance of `HaVpnGatewayArrayInput` via:
//
//          HaVpnGatewayArray{ HaVpnGatewayArgs{...} }
type HaVpnGatewayArrayInput interface {
	pulumi.Input

	ToHaVpnGatewayArrayOutput() HaVpnGatewayArrayOutput
	ToHaVpnGatewayArrayOutputWithContext(context.Context) HaVpnGatewayArrayOutput
}

type HaVpnGatewayArray []HaVpnGatewayInput

func (HaVpnGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVpnGateway)(nil)).Elem()
}

func (i HaVpnGatewayArray) ToHaVpnGatewayArrayOutput() HaVpnGatewayArrayOutput {
	return i.ToHaVpnGatewayArrayOutputWithContext(context.Background())
}

func (i HaVpnGatewayArray) ToHaVpnGatewayArrayOutputWithContext(ctx context.Context) HaVpnGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVpnGatewayArrayOutput)
}

// HaVpnGatewayMapInput is an input type that accepts HaVpnGatewayMap and HaVpnGatewayMapOutput values.
// You can construct a concrete instance of `HaVpnGatewayMapInput` via:
//
//          HaVpnGatewayMap{ "key": HaVpnGatewayArgs{...} }
type HaVpnGatewayMapInput interface {
	pulumi.Input

	ToHaVpnGatewayMapOutput() HaVpnGatewayMapOutput
	ToHaVpnGatewayMapOutputWithContext(context.Context) HaVpnGatewayMapOutput
}

type HaVpnGatewayMap map[string]HaVpnGatewayInput

func (HaVpnGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVpnGateway)(nil)).Elem()
}

func (i HaVpnGatewayMap) ToHaVpnGatewayMapOutput() HaVpnGatewayMapOutput {
	return i.ToHaVpnGatewayMapOutputWithContext(context.Background())
}

func (i HaVpnGatewayMap) ToHaVpnGatewayMapOutputWithContext(ctx context.Context) HaVpnGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVpnGatewayMapOutput)
}

type HaVpnGatewayOutput struct{ *pulumi.OutputState }

func (HaVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVpnGateway)(nil)).Elem()
}

func (o HaVpnGatewayOutput) ToHaVpnGatewayOutput() HaVpnGatewayOutput {
	return o
}

func (o HaVpnGatewayOutput) ToHaVpnGatewayOutputWithContext(ctx context.Context) HaVpnGatewayOutput {
	return o
}

type HaVpnGatewayArrayOutput struct{ *pulumi.OutputState }

func (HaVpnGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVpnGateway)(nil)).Elem()
}

func (o HaVpnGatewayArrayOutput) ToHaVpnGatewayArrayOutput() HaVpnGatewayArrayOutput {
	return o
}

func (o HaVpnGatewayArrayOutput) ToHaVpnGatewayArrayOutputWithContext(ctx context.Context) HaVpnGatewayArrayOutput {
	return o
}

func (o HaVpnGatewayArrayOutput) Index(i pulumi.IntInput) HaVpnGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HaVpnGateway {
		return vs[0].([]*HaVpnGateway)[vs[1].(int)]
	}).(HaVpnGatewayOutput)
}

type HaVpnGatewayMapOutput struct{ *pulumi.OutputState }

func (HaVpnGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVpnGateway)(nil)).Elem()
}

func (o HaVpnGatewayMapOutput) ToHaVpnGatewayMapOutput() HaVpnGatewayMapOutput {
	return o
}

func (o HaVpnGatewayMapOutput) ToHaVpnGatewayMapOutputWithContext(ctx context.Context) HaVpnGatewayMapOutput {
	return o
}

func (o HaVpnGatewayMapOutput) MapIndex(k pulumi.StringInput) HaVpnGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HaVpnGateway {
		return vs[0].(map[string]*HaVpnGateway)[vs[1].(string)]
	}).(HaVpnGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HaVpnGatewayInput)(nil)).Elem(), &HaVpnGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVpnGatewayArrayInput)(nil)).Elem(), HaVpnGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVpnGatewayMapInput)(nil)).Elem(), HaVpnGatewayMap{})
	pulumi.RegisterOutputType(HaVpnGatewayOutput{})
	pulumi.RegisterOutputType(HaVpnGatewayArrayOutput{})
	pulumi.RegisterOutputType(HaVpnGatewayMapOutput{})
}
