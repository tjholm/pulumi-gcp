// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterPrivateClusterConfig
    {
        /// <summary>
        /// When `true`, the cluster's private
        /// endpoint is used as the cluster endpoint and access through the public endpoint
        /// is disabled. When `false`, either endpoint can be used. This field only applies
        /// to private clusters, when `enable_private_nodes` is `true`.
        /// </summary>
        public readonly bool? EnablePrivateEndpoint;
        /// <summary>
        /// Enables the private cluster feature,
        /// creating a private endpoint on the cluster. In a private cluster, nodes only
        /// have RFC 1918 private addresses and communicate with the master's private
        /// endpoint via private networking.
        /// </summary>
        public readonly bool? EnablePrivateNodes;
        public readonly Outputs.ClusterPrivateClusterConfigMasterGlobalAccessConfig? MasterGlobalAccessConfig;
        /// <summary>
        /// The IP range in CIDR notation to use for
        /// the hosted master network. This range will be used for assigning private IP
        /// addresses to the cluster master(s) and the ILB VIP. This range must not overlap
        /// with any other ranges in use within the cluster's network, and it must be a /28
        /// subnet. See [Private Cluster Limitations](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters#req_res_lim)
        /// for more details. This field only applies to private clusters, when
        /// `enable_private_nodes` is `true`.
        /// </summary>
        public readonly string? MasterIpv4CidrBlock;
        /// <summary>
        /// The name of the peering between this cluster and the Google owned VPC.
        /// </summary>
        public readonly string? PeeringName;
        /// <summary>
        /// The internal IP address of this cluster's master endpoint.
        /// </summary>
        public readonly string? PrivateEndpoint;
        /// <summary>
        /// Subnetwork in cluster's network where master's endpoint will be provisioned.
        /// </summary>
        public readonly string? PrivateEndpointSubnetwork;
        /// <summary>
        /// The external IP address of this cluster's master endpoint.
        /// </summary>
        public readonly string? PublicEndpoint;

        [OutputConstructor]
        private ClusterPrivateClusterConfig(
            bool? enablePrivateEndpoint,

            bool? enablePrivateNodes,

            Outputs.ClusterPrivateClusterConfigMasterGlobalAccessConfig? masterGlobalAccessConfig,

            string? masterIpv4CidrBlock,

            string? peeringName,

            string? privateEndpoint,

            string? privateEndpointSubnetwork,

            string? publicEndpoint)
        {
            EnablePrivateEndpoint = enablePrivateEndpoint;
            EnablePrivateNodes = enablePrivateNodes;
            MasterGlobalAccessConfig = masterGlobalAccessConfig;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            PeeringName = peeringName;
            PrivateEndpoint = privateEndpoint;
            PrivateEndpointSubnetwork = privateEndpointSubnetwork;
            PublicEndpoint = publicEndpoint;
        }
    }
}
