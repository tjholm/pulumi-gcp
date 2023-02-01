// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterNodePoolNodeConfigWorkloadMetadataConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How to expose the node metadata to the workload running on the node.
        /// Accepted values are:
        /// * MODE_UNSPECIFIED: Not Set
        /// * GCE_METADATA: Expose all Compute Engine metadata to pods.
        /// * GKE_METADATA: Run the GKE Metadata Server on this node. The GKE Metadata Server exposes a metadata API to workloads that is compatible with the V1 Compute Metadata APIs exposed by the Compute Engine and App Engine Metadata Servers. This feature can only be enabled if [workload identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) is enabled at the cluster level.
        /// </summary>
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public ClusterNodePoolNodeConfigWorkloadMetadataConfigArgs()
        {
        }
        public static new ClusterNodePoolNodeConfigWorkloadMetadataConfigArgs Empty => new ClusterNodePoolNodeConfigWorkloadMetadataConfigArgs();
    }
}
