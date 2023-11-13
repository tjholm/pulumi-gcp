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
    /// Represents a VPN gateway managed outside of GCP.
    /// 
    /// To get more information about ExternalVpnGateway, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)
    /// 
    /// ## Example Usage
    /// ### External Vpn Gateway
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network = new Gcp.Compute.Network("network", new()
    ///     {
    ///         RoutingMode = "GLOBAL",
    ///         AutoCreateSubnetworks = false,
    ///     });
    /// 
    ///     var haGateway = new Gcp.Compute.HaVpnGateway("haGateway", new()
    ///     {
    ///         Region = "us-central1",
    ///         Network = network.Id,
    ///     });
    /// 
    ///     var externalGateway = new Gcp.Compute.ExternalVpnGateway("externalGateway", new()
    ///     {
    ///         RedundancyType = "SINGLE_IP_INTERNALLY_REDUNDANT",
    ///         Description = "An externally managed VPN gateway",
    ///         Interfaces = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.ExternalVpnGatewayInterfaceArgs
    ///             {
    ///                 Id = 0,
    ///                 IpAddress = "8.8.8.8",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var networkSubnet1 = new Gcp.Compute.Subnetwork("networkSubnet1", new()
    ///     {
    ///         IpCidrRange = "10.0.1.0/24",
    ///         Region = "us-central1",
    ///         Network = network.Id,
    ///     });
    /// 
    ///     var networkSubnet2 = new Gcp.Compute.Subnetwork("networkSubnet2", new()
    ///     {
    ///         IpCidrRange = "10.0.2.0/24",
    ///         Region = "us-west1",
    ///         Network = network.Id,
    ///     });
    /// 
    ///     var router1 = new Gcp.Compute.Router("router1", new()
    ///     {
    ///         Network = network.Name,
    ///         Bgp = new Gcp.Compute.Inputs.RouterBgpArgs
    ///         {
    ///             Asn = 64514,
    ///         },
    ///     });
    /// 
    ///     var tunnel1 = new Gcp.Compute.VPNTunnel("tunnel1", new()
    ///     {
    ///         Region = "us-central1",
    ///         VpnGateway = haGateway.Id,
    ///         PeerExternalGateway = externalGateway.Id,
    ///         PeerExternalGatewayInterface = 0,
    ///         SharedSecret = "a secret message",
    ///         Router = router1.Id,
    ///         VpnGatewayInterface = 0,
    ///     });
    /// 
    ///     var tunnel2 = new Gcp.Compute.VPNTunnel("tunnel2", new()
    ///     {
    ///         Region = "us-central1",
    ///         VpnGateway = haGateway.Id,
    ///         PeerExternalGateway = externalGateway.Id,
    ///         PeerExternalGatewayInterface = 0,
    ///         SharedSecret = "a secret message",
    ///         Router = router1.Id.Apply(id =&gt; $" {id}"),
    ///         VpnGatewayInterface = 1,
    ///     });
    /// 
    ///     var router1Interface1 = new Gcp.Compute.RouterInterface("router1Interface1", new()
    ///     {
    ///         Router = router1.Name,
    ///         Region = "us-central1",
    ///         IpRange = "169.254.0.1/30",
    ///         VpnTunnel = tunnel1.Name,
    ///     });
    /// 
    ///     var router1Peer1 = new Gcp.Compute.RouterPeer("router1Peer1", new()
    ///     {
    ///         Router = router1.Name,
    ///         Region = "us-central1",
    ///         PeerIpAddress = "169.254.0.2",
    ///         PeerAsn = 64515,
    ///         AdvertisedRoutePriority = 100,
    ///         Interface = router1Interface1.Name,
    ///     });
    /// 
    ///     var router1Interface2 = new Gcp.Compute.RouterInterface("router1Interface2", new()
    ///     {
    ///         Router = router1.Name,
    ///         Region = "us-central1",
    ///         IpRange = "169.254.1.1/30",
    ///         VpnTunnel = tunnel2.Name,
    ///     });
    /// 
    ///     var router1Peer2 = new Gcp.Compute.RouterPeer("router1Peer2", new()
    ///     {
    ///         Router = router1.Name,
    ///         Region = "us-central1",
    ///         PeerIpAddress = "169.254.1.2",
    ///         PeerAsn = 64515,
    ///         AdvertisedRoutePriority = 100,
    ///         Interface = router1Interface2.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ExternalVpnGateway can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/externalVpnGateway:ExternalVpnGateway")]
    public partial class ExternalVpnGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.ExternalVpnGatewayInterface>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource.  Used
        /// internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels for the external VPN gateway resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        /// </summary>
        [Output("redundancyType")]
        public Output<string?> RedundancyType { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalVpnGateway(string name, ExternalVpnGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args ?? new ExternalVpnGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalVpnGateway(string name, Input<string> id, ExternalVpnGatewayState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExternalVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalVpnGateway Get(string name, Input<string> id, ExternalVpnGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalVpnGateway(name, id, state, options);
        }
    }

    public sealed class ExternalVpnGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfaceArgs>? _interfaces;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfaceArgs>());
            set => _interfaces = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for the external VPN gateway resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        public ExternalVpnGatewayArgs()
        {
        }
        public static new ExternalVpnGatewayArgs Empty => new ExternalVpnGatewayArgs();
    }

    public sealed class ExternalVpnGatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("interfaces")]
        private InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs>? _interfaces;

        /// <summary>
        /// A list of interfaces on this external VPN gateway.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.ExternalVpnGatewayInterfaceGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource.  Used
        /// internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for the external VPN gateway resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Indicates the redundancy type of this external VPN gateway
        /// Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public ExternalVpnGatewayState()
        {
        }
        public static new ExternalVpnGatewayState Empty => new ExternalVpnGatewayState();
    }
}
