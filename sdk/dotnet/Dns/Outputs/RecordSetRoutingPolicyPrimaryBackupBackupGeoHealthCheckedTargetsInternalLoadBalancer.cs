// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Outputs
{

    [OutputType]
    public sealed class RecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancer
    {
        /// <summary>
        /// The frontend IP address of the load balancer.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
        /// </summary>
        public readonly string LoadBalancerType;
        /// <summary>
        /// The fully qualified url of the network in which the load balancer belongs. This should be formatted like `projects/{project}/global/networks/{network}` or `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`.
        /// </summary>
        public readonly string NetworkUrl;
        /// <summary>
        /// The configured port of the load balancer.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// The ID of the project in which the load balancer belongs.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The region of the load balancer. Only needed for regional load balancers.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private RecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancer(
            string ipAddress,

            string ipProtocol,

            string loadBalancerType,

            string networkUrl,

            string port,

            string project,

            string? region)
        {
            IpAddress = ipAddress;
            IpProtocol = ipProtocol;
            LoadBalancerType = loadBalancerType;
            NetworkUrl = networkUrl;
            Port = port;
            Project = project;
            Region = region;
        }
    }
}
