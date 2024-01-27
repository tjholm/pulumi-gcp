// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class InstanceTemplateNetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessConfigs")]
        private InputList<Inputs.InstanceTemplateNetworkInterfaceAccessConfigArgs>? _accessConfigs;

        /// <summary>
        /// Access configurations, i.e. IPs via which this
        /// instance can be accessed via the Internet. Omit to ensure that the instance
        /// is not accessible from the Internet (this means that ssh provisioners will
        /// not work unless you can send traffic to the instance's
        /// network (e.g. via tunnel or because it is running on another cloud instance
        /// on that network). This block can be repeated multiple times. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateNetworkInterfaceAccessConfigArgs> AccessConfigs
        {
            get => _accessConfigs ?? (_accessConfigs = new InputList<Inputs.InstanceTemplateNetworkInterfaceAccessConfigArgs>());
            set => _accessConfigs = value;
        }

        [Input("aliasIpRanges")]
        private InputList<Inputs.InstanceTemplateNetworkInterfaceAliasIpRangeArgs>? _aliasIpRanges;

        /// <summary>
        /// An
        /// array of alias IP ranges for this network interface. Can only be specified for network
        /// interfaces on subnet-mode networks. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateNetworkInterfaceAliasIpRangeArgs> AliasIpRanges
        {
            get => _aliasIpRanges ?? (_aliasIpRanges = new InputList<Inputs.InstanceTemplateNetworkInterfaceAliasIpRangeArgs>());
            set => _aliasIpRanges = value;
        }

        [Input("internalIpv6PrefixLength")]
        public Input<int>? InternalIpv6PrefixLength { get; set; }

        [Input("ipv6AccessConfigs")]
        private InputList<Inputs.InstanceTemplateNetworkInterfaceIpv6AccessConfigArgs>? _ipv6AccessConfigs;

        /// <summary>
        /// An array of IPv6 access configurations for this interface.
        /// Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig
        /// specified, then this instance will have no external IPv6 Internet access. Structure documented below.
        /// </summary>
        public InputList<Inputs.InstanceTemplateNetworkInterfaceIpv6AccessConfigArgs> Ipv6AccessConfigs
        {
            get => _ipv6AccessConfigs ?? (_ipv6AccessConfigs = new InputList<Inputs.InstanceTemplateNetworkInterfaceIpv6AccessConfigArgs>());
            set => _ipv6AccessConfigs = value;
        }

        [Input("ipv6AccessType")]
        public Input<string>? Ipv6AccessType { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        /// <summary>
        /// The name of the instance template. If you leave
        /// this blank, the provider will auto-generate a unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name or self_link of the network to attach this interface to.
        /// Use `network` attribute for Legacy or Auto subnetted networks and
        /// `subnetwork` for custom subnetted networks.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.
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
        /// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. Values are IPV4_IPV6 or IPV4_ONLY. If not specified, IPV4_ONLY will be used.
        /// </summary>
        [Input("stackType")]
        public Input<string>? StackType { get; set; }

        /// <summary>
        /// the name of the subnetwork to attach this interface
        /// to. The subnetwork must exist in the same `region` this instance will be
        /// created in. Either `network` or `subnetwork` must be provided.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        /// <summary>
        /// The ID of the project in which the subnetwork belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("subnetworkProject")]
        public Input<string>? SubnetworkProject { get; set; }

        public InstanceTemplateNetworkInterfaceArgs()
        {
        }
        public static new InstanceTemplateNetworkInterfaceArgs Empty => new InstanceTemplateNetworkInterfaceArgs();
    }
}
