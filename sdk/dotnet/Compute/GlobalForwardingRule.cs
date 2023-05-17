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
    /// Represents a GlobalForwardingRule resource. Global forwarding rules are
    /// used to forward traffic to the correct load balancer for HTTP load
    /// balancing. Global forwarding rules can only be used for HTTP load
    /// balancing.
    /// 
    /// For more information, see
    /// https://cloud.google.com/compute/docs/load-balancing/http/
    /// 
    /// ## Example Usage
    /// ### Global Forwarding Rule External Managed
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultBackendService = new Gcp.Compute.BackendService("defaultBackendService", new()
    ///     {
    ///         PortName = "http",
    ///         Protocol = "HTTP",
    ///         TimeoutSec = 10,
    ///         LoadBalancingScheme = "EXTERNAL_MANAGED",
    ///     });
    /// 
    ///     var defaultURLMap = new Gcp.Compute.URLMap("defaultURLMap", new()
    ///     {
    ///         Description = "a description",
    ///         DefaultService = defaultBackendService.Id,
    ///         HostRules = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.URLMapHostRuleArgs
    ///             {
    ///                 Hosts = new[]
    ///                 {
    ///                     "mysite.com",
    ///                 },
    ///                 PathMatcher = "allpaths",
    ///             },
    ///         },
    ///         PathMatchers = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.URLMapPathMatcherArgs
    ///             {
    ///                 Name = "allpaths",
    ///                 DefaultService = defaultBackendService.Id,
    ///                 PathRules = new[]
    ///                 {
    ///                     new Gcp.Compute.Inputs.URLMapPathMatcherPathRuleArgs
    ///                     {
    ///                         Paths = new[]
    ///                         {
    ///                             "/*",
    ///                         },
    ///                         Service = defaultBackendService.Id,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultTargetHttpProxy = new Gcp.Compute.TargetHttpProxy("defaultTargetHttpProxy", new()
    ///     {
    ///         Description = "a description",
    ///         UrlMap = defaultURLMap.Id,
    ///     });
    /// 
    ///     var defaultGlobalForwardingRule = new Gcp.Compute.GlobalForwardingRule("defaultGlobalForwardingRule", new()
    ///     {
    ///         Target = defaultTargetHttpProxy.Id,
    ///         PortRange = "80",
    ///         LoadBalancingScheme = "EXTERNAL_MANAGED",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Private Service Connect Google Apis
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
    ///         Project = "my-project-name",
    ///         AutoCreateSubnetworks = false,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var vpcSubnetwork = new Gcp.Compute.Subnetwork("vpcSubnetwork", new()
    ///     {
    ///         Project = network.Project,
    ///         IpCidrRange = "10.2.0.0/16",
    ///         Region = "us-central1",
    ///         Network = network.Id,
    ///         PrivateIpGoogleAccess = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultGlobalAddress = new Gcp.Compute.GlobalAddress("defaultGlobalAddress", new()
    ///     {
    ///         Project = network.Project,
    ///         AddressType = "INTERNAL",
    ///         Purpose = "PRIVATE_SERVICE_CONNECT",
    ///         Network = network.Id,
    ///         Address = "100.100.100.106",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultGlobalForwardingRule = new Gcp.Compute.GlobalForwardingRule("defaultGlobalForwardingRule", new()
    ///     {
    ///         Project = network.Project,
    ///         Target = "all-apis",
    ///         Network = network.Id,
    ///         IpAddress = defaultGlobalAddress.Id,
    ///         LoadBalancingScheme = "",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GlobalForwardingRule can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default projects/{{project}}/global/forwardingRules/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/globalForwardingRule:GlobalForwardingRule default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/globalForwardingRule:GlobalForwardingRule")]
    public partial class GlobalForwardingRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
        /// </summary>
        [Output("allowPscGlobalAccess")]
        public Output<bool?> AllowPscGlobalAccess { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
        /// </summary>
        [Output("baseForwardingRule")]
        public Output<string> BaseForwardingRule { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// IP address for which this forwarding rule accepts traffic. When a client
        /// sends traffic to this IP address, the forwarding rule directs the traffic
        /// to the referenced `target`.
        /// While creating a forwarding rule, specifying an `IPAddress` is
        /// required under the following circumstances:
        /// * When the `target` is set to `targetGrpcProxy` and
        /// `validateForProxyless` is set to `true`, the
        /// `IPAddress` should be set to `0.0.0.0`.
        /// * When the `target` is a Private Service Connect Google APIs
        /// bundle, you must specify an `IPAddress`.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The IP protocol to which this rule applies.
        /// For protocol forwarding, valid
        /// options are `TCP`, `UDP`, `ESP`,
        /// `AH`, `SCTP`, `ICMP` and
        /// `L3_DEFAULT`.
        /// The valid IP protocols are different for different load balancing products
        /// as described in [Load balancing
        /// features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
        /// Possible values are: `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, `ICMP`.
        /// </summary>
        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// The IP Version that will be used by this global forwarding rule.
        /// Possible values are: `IPV4`, `IPV6`.
        /// </summary>
        [Output("ipVersion")]
        public Output<string?> IpVersion { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource.  Used
        /// internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this forwarding rule.  A list of key-&gt;value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Specifies the forwarding rule type.
        /// For more information about forwarding rules, refer to
        /// [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts).
        /// Default value is `EXTERNAL`.
        /// Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL_SELF_MANAGED`.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string?> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// Structure is documented below.
        /// </summary>
        [Output("metadataFilters")]
        public Output<ImmutableArray<Outputs.GlobalForwardingRuleMetadataFilter>> MetadataFilters { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-63 characters long, and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
        /// Specifically, the name must be 1-63 characters long and match the regular
        /// expression `a-z?` which means the first
        /// character must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, the forwarding rule name must be a 1-20 characters string with
        /// lowercase letters and numbers and must start with a letter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This field is not used for external load balancing.
        /// For Internal TCP/UDP Load Balancing, this field identifies the network that
        /// the load balanced IP should belong to for this Forwarding Rule.
        /// If the subnetwork is specified, the network of the subnetwork will be used.
        /// If neither subnetwork nor this field is specified, the default network will
        /// be used.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, a network must be provided.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// This field can only be used:
        /// * If `IPProtocol` is one of TCP, UDP, or SCTP.
        /// * By backend service-based network load balancers, target pool-based
        /// network load balancers, internal proxy load balancers, external proxy load
        /// balancers, Traffic Director, external protocol forwarding, and Classic VPN.
        /// Some products have restrictions on what ports can be used. See
        /// [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
        /// for details.
        /// </summary>
        [Output("portRange")]
        public Output<string?> PortRange { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        [Output("pscConnectionId")]
        public Output<string> PscConnectionId { get; private set; } = null!;

        /// <summary>
        /// The PSC connection status of the PSC Forwarding Rule. Possible values: `STATUS_UNSPECIFIED`, `PENDING`, `ACCEPTED`, `REJECTED`, `CLOSED`
        /// </summary>
        [Output("pscConnectionStatus")]
        public Output<string> PscConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
        /// </summary>
        [Output("sourceIpRanges")]
        public Output<ImmutableArray<string>> SourceIpRanges { get; private set; } = null!;

        /// <summary>
        /// The URL of the target resource to receive the matched traffic.  For
        /// regional forwarding rules, this target must be in the same region as the
        /// forwarding rule. For global forwarding rules, this target must be a global
        /// load balancing resource.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// *  For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
        /// *  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:
        /// *  `vpc-sc` - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).
        /// *  `all-apis` - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalForwardingRule(string name, GlobalForwardingRuleArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, args ?? new GlobalForwardingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalForwardingRule(string name, Input<string> id, GlobalForwardingRuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/globalForwardingRule:GlobalForwardingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalForwardingRule Get(string name, Input<string> id, GlobalForwardingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalForwardingRule(name, id, state, options);
        }
    }

    public sealed class GlobalForwardingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
        /// </summary>
        [Input("allowPscGlobalAccess")]
        public Input<bool>? AllowPscGlobalAccess { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP address for which this forwarding rule accepts traffic. When a client
        /// sends traffic to this IP address, the forwarding rule directs the traffic
        /// to the referenced `target`.
        /// While creating a forwarding rule, specifying an `IPAddress` is
        /// required under the following circumstances:
        /// * When the `target` is set to `targetGrpcProxy` and
        /// `validateForProxyless` is set to `true`, the
        /// `IPAddress` should be set to `0.0.0.0`.
        /// * When the `target` is a Private Service Connect Google APIs
        /// bundle, you must specify an `IPAddress`.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP protocol to which this rule applies.
        /// For protocol forwarding, valid
        /// options are `TCP`, `UDP`, `ESP`,
        /// `AH`, `SCTP`, `ICMP` and
        /// `L3_DEFAULT`.
        /// The valid IP protocols are different for different load balancing products
        /// as described in [Load balancing
        /// features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
        /// Possible values are: `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, `ICMP`.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The IP Version that will be used by this global forwarding rule.
        /// Possible values are: `IPV4`, `IPV6`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule.  A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the forwarding rule type.
        /// For more information about forwarding rules, refer to
        /// [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts).
        /// Default value is `EXTERNAL`.
        /// Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL_SELF_MANAGED`.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.GlobalForwardingRuleMetadataFilterArgs>? _metadataFilters;

        /// <summary>
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GlobalForwardingRuleMetadataFilterArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.GlobalForwardingRuleMetadataFilterArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-63 characters long, and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
        /// Specifically, the name must be 1-63 characters long and match the regular
        /// expression `a-z?` which means the first
        /// character must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, the forwarding rule name must be a 1-20 characters string with
        /// lowercase letters and numbers and must start with a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is not used for external load balancing.
        /// For Internal TCP/UDP Load Balancing, this field identifies the network that
        /// the load balanced IP should belong to for this Forwarding Rule.
        /// If the subnetwork is specified, the network of the subnetwork will be used.
        /// If neither subnetwork nor this field is specified, the default network will
        /// be used.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, a network must be provided.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// This field can only be used:
        /// * If `IPProtocol` is one of TCP, UDP, or SCTP.
        /// * By backend service-based network load balancers, target pool-based
        /// network load balancers, internal proxy load balancers, external proxy load
        /// balancers, Traffic Director, external protocol forwarding, and Classic VPN.
        /// Some products have restrictions on what ports can be used. See
        /// [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
        /// for details.
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceIpRanges")]
        private InputList<string>? _sourceIpRanges;

        /// <summary>
        /// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
        /// </summary>
        public InputList<string> SourceIpRanges
        {
            get => _sourceIpRanges ?? (_sourceIpRanges = new InputList<string>());
            set => _sourceIpRanges = value;
        }

        /// <summary>
        /// The URL of the target resource to receive the matched traffic.  For
        /// regional forwarding rules, this target must be in the same region as the
        /// forwarding rule. For global forwarding rules, this target must be a global
        /// load balancing resource.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// *  For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
        /// *  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:
        /// *  `vpc-sc` - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).
        /// *  `all-apis` - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        public GlobalForwardingRuleArgs()
        {
        }
        public static new GlobalForwardingRuleArgs Empty => new GlobalForwardingRuleArgs();
    }

    public sealed class GlobalForwardingRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
        /// </summary>
        [Input("allowPscGlobalAccess")]
        public Input<bool>? AllowPscGlobalAccess { get; set; }

        /// <summary>
        /// [Output Only] The URL for the corresponding base Forwarding Rule. By base Forwarding Rule, we mean the Forwarding Rule that has the same IP address, protocol, and port settings with the current Forwarding Rule, but without sourceIPRanges specified. Always empty if the current Forwarding Rule does not have sourceIPRanges specified.
        /// </summary>
        [Input("baseForwardingRule")]
        public Input<string>? BaseForwardingRule { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP address for which this forwarding rule accepts traffic. When a client
        /// sends traffic to this IP address, the forwarding rule directs the traffic
        /// to the referenced `target`.
        /// While creating a forwarding rule, specifying an `IPAddress` is
        /// required under the following circumstances:
        /// * When the `target` is set to `targetGrpcProxy` and
        /// `validateForProxyless` is set to `true`, the
        /// `IPAddress` should be set to `0.0.0.0`.
        /// * When the `target` is a Private Service Connect Google APIs
        /// bundle, you must specify an `IPAddress`.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The IP protocol to which this rule applies.
        /// For protocol forwarding, valid
        /// options are `TCP`, `UDP`, `ESP`,
        /// `AH`, `SCTP`, `ICMP` and
        /// `L3_DEFAULT`.
        /// The valid IP protocols are different for different load balancing products
        /// as described in [Load balancing
        /// features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
        /// Possible values are: `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, `ICMP`.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The IP Version that will be used by this global forwarding rule.
        /// Possible values are: `IPV4`, `IPV6`.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource.  Used
        /// internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this forwarding rule.  A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the forwarding rule type.
        /// For more information about forwarding rules, refer to
        /// [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts).
        /// Default value is `EXTERNAL`.
        /// Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL_SELF_MANAGED`.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.GlobalForwardingRuleMetadataFilterGetArgs>? _metadataFilters;

        /// <summary>
        /// Opaque filter criteria used by Loadbalancer to restrict routing
        /// configuration to a limited set xDS compliant clients. In their xDS
        /// requests to Loadbalancer, xDS clients present node metadata. If a
        /// match takes place, the relevant routing configuration is made available
        /// to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set
        /// to MATCH_ANY, at least one of the filterLabels must match the
        /// corresponding label provided in the metadata. If its filterMatchCriteria
        /// is set to MATCH_ALL, then all of its filterLabels must match with
        /// corresponding labels in the provided metadata.
        /// metadataFilters specified here can be overridden by those specified in
        /// the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their
        /// loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.GlobalForwardingRuleMetadataFilterGetArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.GlobalForwardingRuleMetadataFilterGetArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-63 characters long, and comply with
        /// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
        /// Specifically, the name must be 1-63 characters long and match the regular
        /// expression `a-z?` which means the first
        /// character must be a lowercase letter, and all following characters must
        /// be a dash, lowercase letter, or digit, except the last character, which
        /// cannot be a dash.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, the forwarding rule name must be a 1-20 characters string with
        /// lowercase letters and numbers and must start with a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is not used for external load balancing.
        /// For Internal TCP/UDP Load Balancing, this field identifies the network that
        /// the load balanced IP should belong to for this Forwarding Rule.
        /// If the subnetwork is specified, the network of the subnetwork will be used.
        /// If neither subnetwork nor this field is specified, the default network will
        /// be used.
        /// For Private Service Connect forwarding rules that forward traffic to Google
        /// APIs, a network must be provided.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// This field can only be used:
        /// * If `IPProtocol` is one of TCP, UDP, or SCTP.
        /// * By backend service-based network load balancers, target pool-based
        /// network load balancers, internal proxy load balancers, external proxy load
        /// balancers, Traffic Director, external protocol forwarding, and Classic VPN.
        /// Some products have restrictions on what ports can be used. See
        /// [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
        /// for details.
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        [Input("pscConnectionId")]
        public Input<string>? PscConnectionId { get; set; }

        /// <summary>
        /// The PSC connection status of the PSC Forwarding Rule. Possible values: `STATUS_UNSPECIFIED`, `PENDING`, `ACCEPTED`, `REJECTED`, `CLOSED`
        /// </summary>
        [Input("pscConnectionStatus")]
        public Input<string>? PscConnectionStatus { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("sourceIpRanges")]
        private InputList<string>? _sourceIpRanges;

        /// <summary>
        /// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here. Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
        /// </summary>
        public InputList<string> SourceIpRanges
        {
            get => _sourceIpRanges ?? (_sourceIpRanges = new InputList<string>());
            set => _sourceIpRanges = value;
        }

        /// <summary>
        /// The URL of the target resource to receive the matched traffic.  For
        /// regional forwarding rules, this target must be in the same region as the
        /// forwarding rule. For global forwarding rules, this target must be a global
        /// load balancing resource.
        /// The forwarded traffic must be of a type appropriate to the target object.
        /// *  For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
        /// *  For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:
        /// *  `vpc-sc` - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).
        /// *  `all-apis` - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public GlobalForwardingRuleState()
        {
        }
        public static new GlobalForwardingRuleState Empty => new GlobalForwardingRuleState();
    }
}
