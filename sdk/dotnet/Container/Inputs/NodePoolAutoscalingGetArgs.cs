// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class NodePoolAutoscalingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location policy specifies the algorithm used when scaling-up the node pool. \
        /// "BALANCED" - Is a best effort policy that aims to balance the sizes of available zones. \
        /// "ANY" - Instructs the cluster autoscaler to prioritize utilization of unused reservations,
        /// and reduce preemption risk for Spot VMs.
        /// </summary>
        [Input("locationPolicy")]
        public Input<string>? LocationPolicy { get; set; }

        /// <summary>
        /// Maximum number of nodes per zone in the NodePool.
        /// Must be &gt;= min_node_count. Cannot be used with total limits.
        /// </summary>
        [Input("maxNodeCount")]
        public Input<int>? MaxNodeCount { get; set; }

        /// <summary>
        /// Minimum number of nodes per zone in the NodePool.
        /// Must be &gt;=0 and &lt;= `max_node_count`. Cannot be used with total limits.
        /// </summary>
        [Input("minNodeCount")]
        public Input<int>? MinNodeCount { get; set; }

        /// <summary>
        /// Total maximum number of nodes in the NodePool.
        /// Must be &gt;= total_min_node_count. Cannot be used with per zone limits.
        /// </summary>
        [Input("totalMaxNodeCount")]
        public Input<int>? TotalMaxNodeCount { get; set; }

        /// <summary>
        /// Total minimum number of nodes in the NodePool.
        /// Must be &gt;=0 and &lt;= `total_max_node_count`. Cannot be used with per zone limits.
        /// </summary>
        [Input("totalMinNodeCount")]
        public Input<int>? TotalMinNodeCount { get; set; }

        public NodePoolAutoscalingGetArgs()
        {
        }
        public static new NodePoolAutoscalingGetArgs Empty => new NodePoolAutoscalingGetArgs();
    }
}
