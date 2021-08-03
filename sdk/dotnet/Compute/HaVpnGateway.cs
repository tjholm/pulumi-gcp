// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents a VPN gateway running in GCP. This virtual device is managed
    /// by Google, but used only by you. This type of VPN Gateway allows for the creation
    /// of VPN solutions with higher availability than classic Target VPN Gateways.
    /// 
    /// To get more information about HaVpnGateway, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways)
    /// * How-to Guides
    ///     * [Choosing a VPN](https://cloud.google.com/vpn/docs/how-to/choosing-a-vpn)
    ///     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
    /// 
    /// ## Example Usage
    /// ### Ha Vpn Gateway Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network1 = new Gcp.Compute.Network("network1", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var haGateway1 = new Gcp.Compute.HaVpnGateway("haGateway1", new Gcp.Compute.HaVpnGatewayArgs
    ///         {
    ///             Region = "us-central1",
    ///             Network = network1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Ha Vpn Gateway Gcp To Gcp
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network1 = new Gcp.Compute.Network("network1", new Gcp.Compute.NetworkArgs
    ///         {
    ///             RoutingMode = "GLOBAL",
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var haGateway1 = new Gcp.Compute.HaVpnGateway("haGateway1", new Gcp.Compute.HaVpnGatewayArgs
    ///         {
    ///             Region = "us-central1",
    ///             Network = network1.Id,
    ///         });
    ///         var network2 = new Gcp.Compute.Network("network2", new Gcp.Compute.NetworkArgs
    ///         {
    ///             RoutingMode = "GLOBAL",
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var haGateway2 = new Gcp.Compute.HaVpnGateway("haGateway2", new Gcp.Compute.HaVpnGatewayArgs
    ///         {
    ///             Region = "us-central1",
    ///             Network = network2.Id,
    ///         });
    ///         var network1Subnet1 = new Gcp.Compute.Subnetwork("network1Subnet1", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "10.0.1.0/24",
    ///             Region = "us-central1",
    ///             Network = network1.Id,
    ///         });
    ///         var network1Subnet2 = new Gcp.Compute.Subnetwork("network1Subnet2", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "10.0.2.0/24",
    ///             Region = "us-west1",
    ///             Network = network1.Id,
    ///         });
    ///         var network2Subnet1 = new Gcp.Compute.Subnetwork("network2Subnet1", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "192.168.1.0/24",
    ///             Region = "us-central1",
    ///             Network = network2.Id,
    ///         });
    ///         var network2Subnet2 = new Gcp.Compute.Subnetwork("network2Subnet2", new Gcp.Compute.SubnetworkArgs
    ///         {
    ///             IpCidrRange = "192.168.2.0/24",
    ///             Region = "us-east1",
    ///             Network = network2.Id,
    ///         });
    ///         var router1 = new Gcp.Compute.Router("router1", new Gcp.Compute.RouterArgs
    ///         {
    ///             Network = network1.Name,
    ///             Bgp = new Gcp.Compute.Inputs.RouterBgpArgs
    ///             {
    ///                 Asn = 64514,
    ///             },
    ///         });
    ///         var router2 = new Gcp.Compute.Router("router2", new Gcp.Compute.RouterArgs
    ///         {
    ///             Network = network2.Name,
    ///             Bgp = new Gcp.Compute.Inputs.RouterBgpArgs
    ///             {
    ///                 Asn = 64515,
    ///             },
    ///         });
    ///         var tunnel1 = new Gcp.Compute.VPNTunnel("tunnel1", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             Region = "us-central1",
    ///             VpnGateway = haGateway1.Id,
    ///             PeerGcpGateway = haGateway2.Id,
    ///             SharedSecret = "a secret message",
    ///             Router = router1.Id,
    ///             VpnGatewayInterface = 0,
    ///         });
    ///         var tunnel2 = new Gcp.Compute.VPNTunnel("tunnel2", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             Region = "us-central1",
    ///             VpnGateway = haGateway1.Id,
    ///             PeerGcpGateway = haGateway2.Id,
    ///             SharedSecret = "a secret message",
    ///             Router = router1.Id,
    ///             VpnGatewayInterface = 1,
    ///         });
    ///         var tunnel3 = new Gcp.Compute.VPNTunnel("tunnel3", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             Region = "us-central1",
    ///             VpnGateway = haGateway2.Id,
    ///             PeerGcpGateway = haGateway1.Id,
    ///             SharedSecret = "a secret message",
    ///             Router = router2.Id,
    ///             VpnGatewayInterface = 0,
    ///         });
    ///         var tunnel4 = new Gcp.Compute.VPNTunnel("tunnel4", new Gcp.Compute.VPNTunnelArgs
    ///         {
    ///             Region = "us-central1",
    ///             VpnGateway = haGateway2.Id,
    ///             PeerGcpGateway = haGateway1.Id,
    ///             SharedSecret = "a secret message",
    ///             Router = router2.Id,
    ///             VpnGatewayInterface = 1,
    ///         });
    ///         var router1Interface1 = new Gcp.Compute.RouterInterface("router1Interface1", new Gcp.Compute.RouterInterfaceArgs
    ///         {
    ///             Router = router1.Name,
    ///             Region = "us-central1",
    ///             IpRange = "169.254.0.1/30",
    ///             VpnTunnel = tunnel1.Name,
    ///         });
    ///         var router1Peer1 = new Gcp.Compute.RouterPeer("router1Peer1", new Gcp.Compute.RouterPeerArgs
    ///         {
    ///             Router = router1.Name,
    ///             Region = "us-central1",
    ///             PeerIpAddress = "169.254.0.2",
    ///             PeerAsn = 64515,
    ///             AdvertisedRoutePriority = 100,
    ///             Interface = router1Interface1.Name,
    ///         });
    ///         var router1Interface2 = new Gcp.Compute.RouterInterface("router1Interface2", new Gcp.Compute.RouterInterfaceArgs
    ///         {
    ///             Router = router1.Name,
    ///             Region = "us-central1",
    ///             IpRange = "169.254.1.1/30",
    ///             VpnTunnel = tunnel2.Name,
    ///         });
    ///         var router1Peer2 = new Gcp.Compute.RouterPeer("router1Peer2", new Gcp.Compute.RouterPeerArgs
    ///         {
    ///             Router = router1.Name,
    ///             Region = "us-central1",
    ///             PeerIpAddress = "169.254.1.2",
    ///             PeerAsn = 64515,
    ///             AdvertisedRoutePriority = 100,
    ///             Interface = router1Interface2.Name,
    ///         });
    ///         var router2Interface1 = new Gcp.Compute.RouterInterface("router2Interface1", new Gcp.Compute.RouterInterfaceArgs
    ///         {
    ///             Router = router2.Name,
    ///             Region = "us-central1",
    ///             IpRange = "169.254.0.1/30",
    ///             VpnTunnel = tunnel3.Name,
    ///         });
    ///         var router2Peer1 = new Gcp.Compute.RouterPeer("router2Peer1", new Gcp.Compute.RouterPeerArgs
    ///         {
    ///             Router = router2.Name,
    ///             Region = "us-central1",
    ///             PeerIpAddress = "169.254.0.2",
    ///             PeerAsn = 64514,
    ///             AdvertisedRoutePriority = 100,
    ///             Interface = router2Interface1.Name,
    ///         });
    ///         var router2Interface2 = new Gcp.Compute.RouterInterface("router2Interface2", new Gcp.Compute.RouterInterfaceArgs
    ///         {
    ///             Router = router2.Name,
    ///             Region = "us-central1",
    ///             IpRange = "169.254.1.1/30",
    ///             VpnTunnel = tunnel4.Name,
    ///         });
    ///         var router2Peer2 = new Gcp.Compute.RouterPeer("router2Peer2", new Gcp.Compute.RouterPeerArgs
    ///         {
    ///             Router = router2.Name,
    ///             Region = "us-central1",
    ///             PeerIpAddress = "169.254.1.2",
    ///             PeerAsn = 64514,
    ///             AdvertisedRoutePriority = 100,
    ///             Interface = router2Interface2.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Compute Ha Vpn Gateway Encrypted Interconnect
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var network = new Gcp.Compute.Network("network", new Gcp.Compute.NetworkArgs
    ///         {
    ///             AutoCreateSubnetworks = false,
    ///         });
    ///         var address1 = new Gcp.Compute.Address("address1", new Gcp.Compute.AddressArgs
    ///         {
    ///             AddressType = "INTERNAL",
    ///             Purpose = "IPSEC_INTERCONNECT",
    ///             Address = "192.168.1.0",
    ///             PrefixLength = 29,
    ///             Network = network.SelfLink,
    ///         });
    ///         var router = new Gcp.Compute.Router("router", new Gcp.Compute.RouterArgs
    ///         {
    ///             Network = network.Name,
    ///             EncryptedInterconnectRouter = true,
    ///             Bgp = new Gcp.Compute.Inputs.RouterBgpArgs
    ///             {
    ///                 Asn = 16550,
    ///             },
    ///         });
    ///         var attachment1 = new Gcp.Compute.InterconnectAttachment("attachment1", new Gcp.Compute.InterconnectAttachmentArgs
    ///         {
    ///             EdgeAvailabilityDomain = "AVAILABILITY_DOMAIN_1",
    ///             Type = "PARTNER",
    ///             Router = router.Id,
    ///             Encryption = "IPSEC",
    ///             IpsecInternalAddresses = 
    ///             {
    ///                 address1.SelfLink,
    ///             },
    ///         });
    ///         var address2 = new Gcp.Compute.Address("address2", new Gcp.Compute.AddressArgs
    ///         {
    ///             AddressType = "INTERNAL",
    ///             Purpose = "IPSEC_INTERCONNECT",
    ///             Address = "192.168.2.0",
    ///             PrefixLength = 29,
    ///             Network = network.SelfLink,
    ///         });
    ///         var attachment2 = new Gcp.Compute.InterconnectAttachment("attachment2", new Gcp.Compute.InterconnectAttachmentArgs
    ///         {
    ///             EdgeAvailabilityDomain = "AVAILABILITY_DOMAIN_2",
    ///             Type = "PARTNER",
    ///             Router = router.Id,
    ///             Encryption = "IPSEC",
    ///             IpsecInternalAddresses = 
    ///             {
    ///                 address2.SelfLink,
    ///             },
    ///         });
    ///         var vpn_gateway = new Gcp.Compute.HaVpnGateway("vpn-gateway", new Gcp.Compute.HaVpnGatewayArgs
    ///         {
    ///             Network = network.Id,
    ///             VpnInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.HaVpnGatewayVpnInterfaceArgs
    ///                 {
    ///                     Id = 0,
    ///                     InterconnectAttachment = attachment1.SelfLink,
    ///                 },
    ///                 new Gcp.Compute.Inputs.HaVpnGatewayVpnInterfaceArgs
    ///                 {
    ///                     Id = 1,
    ///                     InterconnectAttachment = attachment2.SelfLink,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// HaVpnGateway can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/haVpnGateway:HaVpnGateway default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/haVpnGateway:HaVpnGateway")]
    public partial class HaVpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A list of interfaces on this VPN gateway.
        /// Structure is documented below.
        /// </summary>
        [Output("vpnInterfaces")]
        public Output<ImmutableArray<Outputs.HaVpnGatewayVpnInterface>> VpnInterfaces { get; private set; } = null!;


        /// <summary>
        /// Create a HaVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HaVpnGateway(string name, HaVpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/haVpnGateway:HaVpnGateway", name, args ?? new HaVpnGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HaVpnGateway(string name, Input<string> id, HaVpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/haVpnGateway:HaVpnGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HaVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HaVpnGateway Get(string name, Input<string> id, HaVpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new HaVpnGateway(name, id, state, options);
        }
    }

    public sealed class HaVpnGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("vpnInterfaces")]
        private InputList<Inputs.HaVpnGatewayVpnInterfaceArgs>? _vpnInterfaces;

        /// <summary>
        /// A list of interfaces on this VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.HaVpnGatewayVpnInterfaceArgs> VpnInterfaces
        {
            get => _vpnInterfaces ?? (_vpnInterfaces = new InputList<Inputs.HaVpnGatewayVpnInterfaceArgs>());
            set => _vpnInterfaces = value;
        }

        public HaVpnGatewayArgs()
        {
        }
    }

    public sealed class HaVpnGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network this VPN gateway is accepting traffic for.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this gateway should sit in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("vpnInterfaces")]
        private InputList<Inputs.HaVpnGatewayVpnInterfaceGetArgs>? _vpnInterfaces;

        /// <summary>
        /// A list of interfaces on this VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.HaVpnGatewayVpnInterfaceGetArgs> VpnInterfaces
        {
            get => _vpnInterfaces ?? (_vpnInterfaces = new InputList<Inputs.HaVpnGatewayVpnInterfaceGetArgs>());
            set => _vpnInterfaces = value;
        }

        public HaVpnGatewayState()
        {
        }
    }
}
