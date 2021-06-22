// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget
    {
        /// <summary>
        /// Optional. A namespace within the GKE cluster to deploy into.
        /// </summary>
        public readonly string? ClusterNamespace;
        /// <summary>
        /// Optional. The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'
        /// </summary>
        public readonly string? TargetGkeCluster;

        [OutputConstructor]
        private WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(
            string? clusterNamespace,

            string? targetGkeCluster)
        {
            ClusterNamespace = clusterNamespace;
            TargetGkeCluster = targetGkeCluster;
        }
    }
}
