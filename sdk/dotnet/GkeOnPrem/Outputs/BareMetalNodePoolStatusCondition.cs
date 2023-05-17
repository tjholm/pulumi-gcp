// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeOnPrem.Outputs
{

    [OutputType]
    public sealed class BareMetalNodePoolStatusCondition
    {
        /// <summary>
        /// (Output)
        /// Last time the condition transit from one status to another.
        /// </summary>
        public readonly string? LastTransitionTime;
        /// <summary>
        /// Human-readable message indicating details about last transition.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Machine-readable message indicating details about last transition.
        /// </summary>
        public readonly string? Reason;
        /// <summary>
        /// (Output)
        /// The lifecycle state of the condition.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Type of the condition.
        /// (e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private BareMetalNodePoolStatusCondition(
            string? lastTransitionTime,

            string? message,

            string? reason,

            string? state,

            string? type)
        {
            LastTransitionTime = lastTransitionTime;
            Message = message;
            Reason = reason;
            State = state;
            Type = type;
        }
    }
}
