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
    /// Each network has its own firewall controlling access to and from the
    /// instances.
    /// 
    /// All traffic to instances, even from other instances, is blocked by the
    /// firewall unless firewall rules are created to allow it.
    /// 
    /// The default network has automatically created firewall rules that are
    /// shown in default firewall rules. No manually created network has
    /// automatically created firewall rules except for a default "allow" rule for
    /// outgoing traffic and a default "deny" for incoming traffic. For all
    /// networks except the default network, you must create any firewall rules
    /// you need.
    /// 
    /// To get more information about Firewall, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/firewalls)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)
    /// 
    /// ## Example Usage
    /// ### Firewall Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultNetwork = new Gcp.Compute.Network("defaultNetwork", new Gcp.Compute.NetworkArgs
    ///         {
    ///         });
    ///         var defaultFirewall = new Gcp.Compute.Firewall("defaultFirewall", new Gcp.Compute.FirewallArgs
    ///         {
    ///             Network = defaultNetwork.Name,
    ///             Allows = 
    ///             {
    ///                 new Gcp.Compute.Inputs.FirewallAllowArgs
    ///                 {
    ///                     Protocol = "icmp",
    ///                 },
    ///                 new Gcp.Compute.Inputs.FirewallAllowArgs
    ///                 {
    ///                     Protocol = "tcp",
    ///                     Ports = 
    ///                     {
    ///                         "80",
    ///                         "8080",
    ///                         "1000-2000",
    ///                     },
    ///                 },
    ///             },
    ///             SourceTags = 
    ///             {
    ///                 "web",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Firewall With Target Tags
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rules = new Gcp.Compute.Firewall("rules", new Gcp.Compute.FirewallArgs
    ///         {
    ///             Allows = 
    ///             {
    ///                 new Gcp.Compute.Inputs.FirewallAllowArgs
    ///                 {
    ///                     Ports = 
    ///                     {
    ///                         "80",
    ///                         "8080",
    ///                         "1000-2000",
    ///                     },
    ///                     Protocol = "tcp",
    ///                 },
    ///             },
    ///             Description = "Creates firewall rule targeting tagged instances",
    ///             Network = "default",
    ///             Project = "my-project-name",
    ///             SourceTags = 
    ///             {
    ///                 "foo",
    ///             },
    ///             TargetTags = 
    ///             {
    ///                 "web",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/firewall:Firewall default projects/{{project}}/global/firewalls/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/firewall:Firewall default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/firewall:Firewall default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/firewall:Firewall")]
    public partial class Firewall : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of ALLOW rules specified by this firewall. Each rule
        /// specifies a protocol and port-range tuple that describes a permitted
        /// connection.
        /// Structure is documented below.
        /// </summary>
        [Output("allows")]
        public Output<ImmutableArray<Outputs.FirewallAllow>> Allows { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// The list of DENY rules specified by this firewall. Each rule specifies
        /// a protocol and port-range tuple that describes a denied connection.
        /// Structure is documented below.
        /// </summary>
        [Output("denies")]
        public Output<ImmutableArray<Outputs.FirewallDeny>> Denies { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If destination ranges are specified, the firewall will apply only to
        /// traffic that has destination IP address in these ranges. These ranges
        /// must be expressed in CIDR format. Only IPv4 is supported.
        /// </summary>
        [Output("destinationRanges")]
        public Output<ImmutableArray<string>> DestinationRanges { get; private set; } = null!;

        /// <summary>
        /// Direction of traffic to which this firewall applies; default is
        /// INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
        /// destinationRanges; For EGRESS traffic, it is NOT supported to specify
        /// `source_ranges` OR `source_tags`. For INGRESS traffic, one of `source_ranges`,
        /// `source_tags` or `source_service_accounts` is required.
        /// Possible values are `INGRESS` and `EGRESS`.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Denotes whether the firewall rule is disabled, i.e not applied to the
        /// network it is associated with. When set to true, the firewall rule is
        /// not enforced and the network behaves as if it did not exist. If this
        /// is unspecified, the firewall rule will be enabled.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// This field denotes whether to enable logging for a particular firewall rule.
        /// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
        /// </summary>
        [Output("enableLogging")]
        public Output<bool> EnableLogging { get; private set; } = null!;

        /// <summary>
        /// This field denotes the logging options for a particular firewall rule.
        /// If defined, logging is enabled, and logs will be exported to Cloud Logging.
        /// Structure is documented below.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.FirewallLogConfig?> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name or self_link of the network to attach this firewall to.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Priority for this rule. This is an integer between 0 and 65535, both
        /// inclusive. When not specified, the value assumed is 1000. Relative
        /// priorities determine precedence of conflicting rules. Lower value of
        /// priority implies higher precedence (eg, a rule with priority 0 has
        /// higher precedence than a rule with priority 1). DENY rules take
        /// precedence over ALLOW rules having equal priority.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// If source ranges are specified, the firewall will apply only to
        /// traffic that has source IP address in these ranges. These ranges must
        /// be expressed in CIDR format. One or both of sourceRanges and
        /// sourceTags may be set. If both properties are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP that belongs to a tag listed in the sourceTags property. The
        /// connection does not need to match both properties for the firewall to
        /// apply. Only IPv4 is supported.
        /// </summary>
        [Output("sourceRanges")]
        public Output<ImmutableArray<string>> SourceRanges { get; private set; } = null!;

        /// <summary>
        /// If source service accounts are specified, the firewall will apply only
        /// to traffic originating from an instance with a service account in this
        /// list. Source service accounts cannot be used to control traffic to an
        /// instance's external IP address because service accounts are associated
        /// with an instance, not an IP address. sourceRanges can be set at the
        /// same time as sourceServiceAccounts. If both are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP belongs to an instance with service account listed in
        /// sourceServiceAccount. The connection does not need to match both
        /// properties for the firewall to apply. sourceServiceAccounts cannot be
        /// used at the same time as sourceTags or targetTags.
        /// </summary>
        [Output("sourceServiceAccounts")]
        public Output<ImmutableArray<string>> SourceServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// If source tags are specified, the firewall will apply only to traffic
        /// with source IP that belongs to a tag listed in source tags. Source
        /// tags cannot be used to control traffic to an instance's external IP
        /// address. Because tags are associated with an instance, not an IP
        /// address. One or both of sourceRanges and sourceTags may be set. If
        /// both properties are set, the firewall will apply to traffic that has
        /// source IP address within sourceRanges OR the source IP that belongs to
        /// a tag listed in the sourceTags property. The connection does not need
        /// to match both properties for the firewall to apply.
        /// </summary>
        [Output("sourceTags")]
        public Output<ImmutableArray<string>> SourceTags { get; private set; } = null!;

        /// <summary>
        /// A list of service accounts indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// targetServiceAccounts cannot be used at the same time as targetTags or
        /// sourceTags. If neither targetServiceAccounts nor targetTags are
        /// specified, the firewall rule applies to all instances on the specified
        /// network.
        /// </summary>
        [Output("targetServiceAccounts")]
        public Output<ImmutableArray<string>> TargetServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// A list of instance tags indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// If no targetTags are specified, the firewall rule applies to all
        /// instances on the specified network.
        /// </summary>
        [Output("targetTags")]
        public Output<ImmutableArray<string>> TargetTags { get; private set; } = null!;


        /// <summary>
        /// Create a Firewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Firewall(string name, FirewallArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/firewall:Firewall", name, args ?? new FirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Firewall(string name, Input<string> id, FirewallState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/firewall:Firewall", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Firewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Firewall Get(string name, Input<string> id, FirewallState? state = null, CustomResourceOptions? options = null)
        {
            return new Firewall(name, id, state, options);
        }
    }

    public sealed class FirewallArgs : Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<Inputs.FirewallAllowArgs>? _allows;

        /// <summary>
        /// The list of ALLOW rules specified by this firewall. Each rule
        /// specifies a protocol and port-range tuple that describes a permitted
        /// connection.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAllowArgs> Allows
        {
            get => _allows ?? (_allows = new InputList<Inputs.FirewallAllowArgs>());
            set => _allows = value;
        }

        [Input("denies")]
        private InputList<Inputs.FirewallDenyArgs>? _denies;

        /// <summary>
        /// The list of DENY rules specified by this firewall. Each rule specifies
        /// a protocol and port-range tuple that describes a denied connection.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FirewallDenyArgs> Denies
        {
            get => _denies ?? (_denies = new InputList<Inputs.FirewallDenyArgs>());
            set => _denies = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationRanges")]
        private InputList<string>? _destinationRanges;

        /// <summary>
        /// If destination ranges are specified, the firewall will apply only to
        /// traffic that has destination IP address in these ranges. These ranges
        /// must be expressed in CIDR format. Only IPv4 is supported.
        /// </summary>
        public InputList<string> DestinationRanges
        {
            get => _destinationRanges ?? (_destinationRanges = new InputList<string>());
            set => _destinationRanges = value;
        }

        /// <summary>
        /// Direction of traffic to which this firewall applies; default is
        /// INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
        /// destinationRanges; For EGRESS traffic, it is NOT supported to specify
        /// `source_ranges` OR `source_tags`. For INGRESS traffic, one of `source_ranges`,
        /// `source_tags` or `source_service_accounts` is required.
        /// Possible values are `INGRESS` and `EGRESS`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Denotes whether the firewall rule is disabled, i.e not applied to the
        /// network it is associated with. When set to true, the firewall rule is
        /// not enforced and the network behaves as if it did not exist. If this
        /// is unspecified, the firewall rule will be enabled.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// This field denotes whether to enable logging for a particular firewall rule.
        /// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// This field denotes the logging options for a particular firewall rule.
        /// If defined, logging is enabled, and logs will be exported to Cloud Logging.
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.FirewallLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or self_link of the network to attach this firewall to.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// Priority for this rule. This is an integer between 0 and 65535, both
        /// inclusive. When not specified, the value assumed is 1000. Relative
        /// priorities determine precedence of conflicting rules. Lower value of
        /// priority implies higher precedence (eg, a rule with priority 0 has
        /// higher precedence than a rule with priority 1). DENY rules take
        /// precedence over ALLOW rules having equal priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceRanges")]
        private InputList<string>? _sourceRanges;

        /// <summary>
        /// If source ranges are specified, the firewall will apply only to
        /// traffic that has source IP address in these ranges. These ranges must
        /// be expressed in CIDR format. One or both of sourceRanges and
        /// sourceTags may be set. If both properties are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP that belongs to a tag listed in the sourceTags property. The
        /// connection does not need to match both properties for the firewall to
        /// apply. Only IPv4 is supported.
        /// </summary>
        public InputList<string> SourceRanges
        {
            get => _sourceRanges ?? (_sourceRanges = new InputList<string>());
            set => _sourceRanges = value;
        }

        [Input("sourceServiceAccounts")]
        private InputList<string>? _sourceServiceAccounts;

        /// <summary>
        /// If source service accounts are specified, the firewall will apply only
        /// to traffic originating from an instance with a service account in this
        /// list. Source service accounts cannot be used to control traffic to an
        /// instance's external IP address because service accounts are associated
        /// with an instance, not an IP address. sourceRanges can be set at the
        /// same time as sourceServiceAccounts. If both are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP belongs to an instance with service account listed in
        /// sourceServiceAccount. The connection does not need to match both
        /// properties for the firewall to apply. sourceServiceAccounts cannot be
        /// used at the same time as sourceTags or targetTags.
        /// </summary>
        public InputList<string> SourceServiceAccounts
        {
            get => _sourceServiceAccounts ?? (_sourceServiceAccounts = new InputList<string>());
            set => _sourceServiceAccounts = value;
        }

        [Input("sourceTags")]
        private InputList<string>? _sourceTags;

        /// <summary>
        /// If source tags are specified, the firewall will apply only to traffic
        /// with source IP that belongs to a tag listed in source tags. Source
        /// tags cannot be used to control traffic to an instance's external IP
        /// address. Because tags are associated with an instance, not an IP
        /// address. One or both of sourceRanges and sourceTags may be set. If
        /// both properties are set, the firewall will apply to traffic that has
        /// source IP address within sourceRanges OR the source IP that belongs to
        /// a tag listed in the sourceTags property. The connection does not need
        /// to match both properties for the firewall to apply.
        /// </summary>
        public InputList<string> SourceTags
        {
            get => _sourceTags ?? (_sourceTags = new InputList<string>());
            set => _sourceTags = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// targetServiceAccounts cannot be used at the same time as targetTags or
        /// sourceTags. If neither targetServiceAccounts nor targetTags are
        /// specified, the firewall rule applies to all instances on the specified
        /// network.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        [Input("targetTags")]
        private InputList<string>? _targetTags;

        /// <summary>
        /// A list of instance tags indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// If no targetTags are specified, the firewall rule applies to all
        /// instances on the specified network.
        /// </summary>
        public InputList<string> TargetTags
        {
            get => _targetTags ?? (_targetTags = new InputList<string>());
            set => _targetTags = value;
        }

        public FirewallArgs()
        {
        }
    }

    public sealed class FirewallState : Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<Inputs.FirewallAllowGetArgs>? _allows;

        /// <summary>
        /// The list of ALLOW rules specified by this firewall. Each rule
        /// specifies a protocol and port-range tuple that describes a permitted
        /// connection.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAllowGetArgs> Allows
        {
            get => _allows ?? (_allows = new InputList<Inputs.FirewallAllowGetArgs>());
            set => _allows = value;
        }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("denies")]
        private InputList<Inputs.FirewallDenyGetArgs>? _denies;

        /// <summary>
        /// The list of DENY rules specified by this firewall. Each rule specifies
        /// a protocol and port-range tuple that describes a denied connection.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FirewallDenyGetArgs> Denies
        {
            get => _denies ?? (_denies = new InputList<Inputs.FirewallDenyGetArgs>());
            set => _denies = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationRanges")]
        private InputList<string>? _destinationRanges;

        /// <summary>
        /// If destination ranges are specified, the firewall will apply only to
        /// traffic that has destination IP address in these ranges. These ranges
        /// must be expressed in CIDR format. Only IPv4 is supported.
        /// </summary>
        public InputList<string> DestinationRanges
        {
            get => _destinationRanges ?? (_destinationRanges = new InputList<string>());
            set => _destinationRanges = value;
        }

        /// <summary>
        /// Direction of traffic to which this firewall applies; default is
        /// INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
        /// destinationRanges; For EGRESS traffic, it is NOT supported to specify
        /// `source_ranges` OR `source_tags`. For INGRESS traffic, one of `source_ranges`,
        /// `source_tags` or `source_service_accounts` is required.
        /// Possible values are `INGRESS` and `EGRESS`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Denotes whether the firewall rule is disabled, i.e not applied to the
        /// network it is associated with. When set to true, the firewall rule is
        /// not enforced and the network behaves as if it did not exist. If this
        /// is unspecified, the firewall rule will be enabled.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// This field denotes whether to enable logging for a particular firewall rule.
        /// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// This field denotes the logging options for a particular firewall rule.
        /// If defined, logging is enabled, and logs will be exported to Cloud Logging.
        /// Structure is documented below.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.FirewallLogConfigGetArgs>? LogConfig { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or self_link of the network to attach this firewall to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Priority for this rule. This is an integer between 0 and 65535, both
        /// inclusive. When not specified, the value assumed is 1000. Relative
        /// priorities determine precedence of conflicting rules. Lower value of
        /// priority implies higher precedence (eg, a rule with priority 0 has
        /// higher precedence than a rule with priority 1). DENY rules take
        /// precedence over ALLOW rules having equal priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("sourceRanges")]
        private InputList<string>? _sourceRanges;

        /// <summary>
        /// If source ranges are specified, the firewall will apply only to
        /// traffic that has source IP address in these ranges. These ranges must
        /// be expressed in CIDR format. One or both of sourceRanges and
        /// sourceTags may be set. If both properties are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP that belongs to a tag listed in the sourceTags property. The
        /// connection does not need to match both properties for the firewall to
        /// apply. Only IPv4 is supported.
        /// </summary>
        public InputList<string> SourceRanges
        {
            get => _sourceRanges ?? (_sourceRanges = new InputList<string>());
            set => _sourceRanges = value;
        }

        [Input("sourceServiceAccounts")]
        private InputList<string>? _sourceServiceAccounts;

        /// <summary>
        /// If source service accounts are specified, the firewall will apply only
        /// to traffic originating from an instance with a service account in this
        /// list. Source service accounts cannot be used to control traffic to an
        /// instance's external IP address because service accounts are associated
        /// with an instance, not an IP address. sourceRanges can be set at the
        /// same time as sourceServiceAccounts. If both are set, the firewall will
        /// apply to traffic that has source IP address within sourceRanges OR the
        /// source IP belongs to an instance with service account listed in
        /// sourceServiceAccount. The connection does not need to match both
        /// properties for the firewall to apply. sourceServiceAccounts cannot be
        /// used at the same time as sourceTags or targetTags.
        /// </summary>
        public InputList<string> SourceServiceAccounts
        {
            get => _sourceServiceAccounts ?? (_sourceServiceAccounts = new InputList<string>());
            set => _sourceServiceAccounts = value;
        }

        [Input("sourceTags")]
        private InputList<string>? _sourceTags;

        /// <summary>
        /// If source tags are specified, the firewall will apply only to traffic
        /// with source IP that belongs to a tag listed in source tags. Source
        /// tags cannot be used to control traffic to an instance's external IP
        /// address. Because tags are associated with an instance, not an IP
        /// address. One or both of sourceRanges and sourceTags may be set. If
        /// both properties are set, the firewall will apply to traffic that has
        /// source IP address within sourceRanges OR the source IP that belongs to
        /// a tag listed in the sourceTags property. The connection does not need
        /// to match both properties for the firewall to apply.
        /// </summary>
        public InputList<string> SourceTags
        {
            get => _sourceTags ?? (_sourceTags = new InputList<string>());
            set => _sourceTags = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// targetServiceAccounts cannot be used at the same time as targetTags or
        /// sourceTags. If neither targetServiceAccounts nor targetTags are
        /// specified, the firewall rule applies to all instances on the specified
        /// network.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        [Input("targetTags")]
        private InputList<string>? _targetTags;

        /// <summary>
        /// A list of instance tags indicating sets of instances located in the
        /// network that may make network connections as specified in allowed[].
        /// If no targetTags are specified, the firewall rule applies to all
        /// instances on the specified network.
        /// </summary>
        public InputList<string> TargetTags
        {
            get => _targetTags ?? (_targetTags = new InputList<string>());
            set => _targetTags = value;
        }

        public FirewallState()
        {
        }
    }
}
