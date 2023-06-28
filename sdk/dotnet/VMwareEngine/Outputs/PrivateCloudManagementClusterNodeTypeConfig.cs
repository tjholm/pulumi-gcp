// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.VMwareEngine.Outputs
{

    [OutputType]
    public sealed class PrivateCloudManagementClusterNodeTypeConfig
    {
        /// <summary>
        /// Customized number of cores available to each node of the type.
        /// This number must always be one of `nodeType.availableCustomCoreCounts`.
        /// If zero is provided max value from `nodeType.availableCustomCoreCounts` will be used.
        /// This cannot be changed once the PrivateCloud is created.
        /// 
        /// - - -
        /// </summary>
        public readonly int? CustomCoreCount;
        /// <summary>
        /// The number of nodes of this type in the cluster.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        public readonly string NodeTypeId;

        [OutputConstructor]
        private PrivateCloudManagementClusterNodeTypeConfig(
            int? customCoreCount,

            int nodeCount,

            string nodeTypeId)
        {
            CustomCoreCount = customCoreCount;
            NodeCount = nodeCount;
            NodeTypeId = nodeTypeId;
        }
    }
}
