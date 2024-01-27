// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceNetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessConfigs")]
        private InputList<Inputs.InstanceNetworkInterfaceAccessConfigArgs>? _accessConfigs;
        public InputList<Inputs.InstanceNetworkInterfaceAccessConfigArgs> AccessConfigs
        {
            get => _accessConfigs ?? (_accessConfigs = new InputList<Inputs.InstanceNetworkInterfaceAccessConfigArgs>());
            set => _accessConfigs = value;
        }

        [Input("aliasIpRanges")]
        private InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeArgs>? _aliasIpRanges;

        /// <summary>
        /// An
        /// array of alias IP ranges for this network interface. Can only be specified for network
        /// interfaces on subnet-mode networks. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeArgs> AliasIpRanges
        {
            get => _aliasIpRanges ?? (_aliasIpRanges = new InputList<Inputs.InstanceNetworkInterfaceAliasIpRangeArgs>());
            set => _aliasIpRanges = value;
        }

        [Input("internalIpv6PrefixLength")]
        public Input<int>? InternalIpv6PrefixLength { get; set; }

        [Input("ipv6AccessConfigs")]
        private InputList<Inputs.InstanceNetworkInterfaceIpv6AccessConfigArgs>? _ipv6AccessConfigs;

        /// <summary>
        /// An array of IPv6 access configurations for this interface.
        /// Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig
        /// specified, then this instance will have no external IPv6 Internet access. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceNetworkInterfaceIpv6AccessConfigArgs> Ipv6AccessConfigs
        {
            get => _ipv6AccessConfigs ?? (_ipv6AccessConfigs = new InputList<Inputs.InstanceNetworkInterfaceIpv6AccessConfigArgs>());
            set => _ipv6AccessConfigs = value;
        }

        /// <summary>
        /// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet.
        /// This field is always inherited from its subnetwork.
        /// </summary>
        [Input("ipv6AccessType")]
        public Input<string>? Ipv6AccessType { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or self_link of the network to attach this interface to.
        /// Either `network` or `subnetwork` must be provided. If network isn't provided it will
        /// be inferred from the subnetwork.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Beta The URL of the network attachment that this interface should connect to in the following format: `projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}`.
        /// </summary>
        [Input("networkAttachment")]
        public Input<string>? NetworkAttachment { get; set; }

        /// <summary>
        /// The private IP address to assign to the instance. If
        /// empty, the address will be automatically assigned.
        /// </summary>
        [Input("networkIp")]
        public Input<string>? NetworkIp { get; set; }

        /// <summary>
        /// The type of vNIC to be used on this interface. Possible values: GVNIC, VIRTIO_NET.
        /// </summary>
        [Input("nicType")]
        public Input<string>? NicType { get; set; }

        /// <summary>
        /// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
        /// </summary>
        [Input("queueCount")]
        public Input<int>? QueueCount { get; set; }

        /// <summary>
        /// Beta A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
        /// </summary>
        [Input("securityPolicy")]
        public Input<string>? SecurityPolicy { get; set; }

        /// <summary>
        /// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. Values are IPV4_IPV6 or IPV4_ONLY. If not specified, IPV4_ONLY will be used.
        /// </summary>
        [Input("stackType")]
        public Input<string>? StackType { get; set; }

        /// <summary>
        /// The name or self_link of the subnetwork to attach this
        /// interface to. Either `network` or `subnetwork` must be provided. If network isn't provided
        /// it will be inferred from the subnetwork. The subnetwork must exist in the same region this
        /// instance will be created in. If the network resource is in
        /// [legacy](https://cloud.google.com/vpc/docs/legacy) mode, do not specify this field. If the
        /// network is in auto subnet mode, specifying the subnetwork is optional. If the network is
        /// in custom subnet mode, specifying the subnetwork is required.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// The project in which the subnetwork belongs.
        /// If the `subnetwork` is a self_link, this field is ignored in favor of the project
        /// defined in the subnetwork self_link. If the `subnetwork` is a name and this
        /// field is not provided, the provider project is used.
        /// </summary>
        [Input("subnetworkProject")]
        public Input<string>? SubnetworkProject { get; set; }

        public InstanceNetworkInterfaceArgs()
        {
        }
        public static new InstanceNetworkInterfaceArgs Empty => new InstanceNetworkInterfaceArgs();
    }
}
