// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.VMwareEngine.Inputs
{

    public sealed class ClusterNodeTypeConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customized number of cores available to each node of the type.
        /// This number must always be one of `nodeType.availableCustomCoreCounts`.
        /// If zero is provided max value from `nodeType.availableCustomCoreCounts` will be used.
        /// Once the customer is created then corecount cannot be changed.
        /// </summary>
        [Input("customCoreCount")]
        public Input<int>? CustomCoreCount { get; set; }

        /// <summary>
        /// The number of nodes of this type in the cluster.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        [Input("nodeTypeId", required: true)]
        public Input<string> NodeTypeId { get; set; } = null!;

        public ClusterNodeTypeConfigGetArgs()
        {
        }
        public static new ClusterNodeTypeConfigGetArgs Empty => new ClusterNodeTypeConfigGetArgs();
    }
}
