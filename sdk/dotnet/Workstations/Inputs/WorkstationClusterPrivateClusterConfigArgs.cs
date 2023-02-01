// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Inputs
{

    public sealed class WorkstationClusterPrivateClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname for the workstation cluster.
        /// This field will be populated only when private endpoint is enabled.
        /// To access workstations in the cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment.
        /// </summary>
        [Input("clusterHostname")]
        public Input<string>? ClusterHostname { get; set; }

        /// <summary>
        /// Whether Workstations endpoint is private.
        /// </summary>
        [Input("enablePrivateEndpoint", required: true)]
        public Input<bool> EnablePrivateEndpoint { get; set; } = null!;

        /// <summary>
        /// Service attachment URI for the workstation cluster.
        /// The service attachemnt is created when private endpoint is enabled.
        /// To access workstations in the cluster, configure access to the managed service using (Private Service Connect)[https://cloud.google.com/vpc/docs/configure-private-service-connect-services].
        /// </summary>
        [Input("serviceAttachmentUri")]
        public Input<string>? ServiceAttachmentUri { get; set; }

        public WorkstationClusterPrivateClusterConfigArgs()
        {
        }
        public static new WorkstationClusterPrivateClusterConfigArgs Empty => new WorkstationClusterPrivateClusterConfigArgs();
    }
}
