// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplatePlacementManagedClusterConfigLifecycleConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The time when cluster will be auto-deleted (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        [Input("autoDeleteTime")]
        public Input<string>? AutoDeleteTime { get; set; }

        /// <summary>
        /// Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        [Input("autoDeleteTtl")]
        public Input<string>? AutoDeleteTtl { get; set; }

        /// <summary>
        /// Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json).
        /// </summary>
        [Input("idleDeleteTtl")]
        public Input<string>? IdleDeleteTtl { get; set; }

        /// <summary>
        /// -
        /// Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        [Input("idleStartTime")]
        public Input<string>? IdleStartTime { get; set; }

        public WorkflowTemplatePlacementManagedClusterConfigLifecycleConfigGetArgs()
        {
        }
    }
}
