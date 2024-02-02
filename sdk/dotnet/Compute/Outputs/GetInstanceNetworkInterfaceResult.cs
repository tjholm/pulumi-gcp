// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class GetInstanceNetworkInterfaceResult
    {
        /// <summary>
        /// Access configurations, i.e. IPs via which this
        /// instance can be accessed via the Internet. Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkInterfaceAccessConfigResult> AccessConfigs;
        /// <summary>
        /// An array of alias IP ranges for this network interface. Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkInterfaceAliasIpRangeResult> AliasIpRanges;
        /// <summary>
        /// The prefix length of the primary internal IPv6 range.
        /// </summary>
        public readonly int InternalIpv6PrefixLength;
        /// <summary>
        /// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkInterfaceIpv6AccessConfigResult> Ipv6AccessConfigs;
        /// <summary>
        /// One of EXTERNAL, INTERNAL to indicate whether the IP can be accessed from the Internet. This field is always inherited from its subnetwork.
        /// </summary>
        public readonly string Ipv6AccessType;
        /// <summary>
        /// An IPv6 internal network address for this network interface. If not specified, Google Cloud will automatically assign an internal IPv6 address from the instance's subnetwork.
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// The name of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name or self_link of the network attached to this interface.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Beta The URL of the network attachment to this interface.
        /// </summary>
        public readonly string NetworkAttachment;
        /// <summary>
        /// The private IP address assigned to the instance.
        /// </summary>
        public readonly string NetworkIp;
        /// <summary>
        /// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
        /// </summary>
        public readonly string NicType;
        /// <summary>
        /// The networking queue count that's specified by users for the network interface. Both Rx and Tx queues will be set to this number. It will be empty if not specified.
        /// </summary>
        public readonly int QueueCount;
        /// <summary>
        /// A full or partial URL to a security policy to add to this instance. If this field is set to an empty string it will remove the associated security policy.
        /// </summary>
        public readonly string SecurityPolicy;
        /// <summary>
        /// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
        /// </summary>
        public readonly string StackType;
        /// <summary>
        /// The name or self_link of the subnetwork attached to this interface.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// The project in which the subnetwork belongs.
        /// </summary>
        public readonly string SubnetworkProject;

        [OutputConstructor]
        private GetInstanceNetworkInterfaceResult(
            ImmutableArray<Outputs.GetInstanceNetworkInterfaceAccessConfigResult> accessConfigs,

            ImmutableArray<Outputs.GetInstanceNetworkInterfaceAliasIpRangeResult> aliasIpRanges,

            int internalIpv6PrefixLength,

            ImmutableArray<Outputs.GetInstanceNetworkInterfaceIpv6AccessConfigResult> ipv6AccessConfigs,

            string ipv6AccessType,

            string ipv6Address,

            string name,

            string network,

            string networkAttachment,

            string networkIp,

            string nicType,

            int queueCount,

            string securityPolicy,

            string stackType,

            string subnetwork,

            string subnetworkProject)
        {
            AccessConfigs = accessConfigs;
            AliasIpRanges = aliasIpRanges;
            InternalIpv6PrefixLength = internalIpv6PrefixLength;
            Ipv6AccessConfigs = ipv6AccessConfigs;
            Ipv6AccessType = ipv6AccessType;
            Ipv6Address = ipv6Address;
            Name = name;
            Network = network;
            NetworkAttachment = networkAttachment;
            NetworkIp = networkIp;
            NicType = nicType;
            QueueCount = queueCount;
            SecurityPolicy = securityPolicy;
            StackType = stackType;
            Subnetwork = subnetwork;
            SubnetworkProject = subnetworkProject;
        }
    }
}
