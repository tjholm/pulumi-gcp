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
    public sealed class BareMetalClusterStatus
    {
        /// <summary>
        /// (Output)
        /// ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BareMetalClusterStatusCondition> Conditions;
        /// <summary>
        /// (Output)
        /// Human-friendly representation of the error message from the user cluster
        /// controller. The error message can be temporary as the user cluster
        /// controller creates a cluster or node pool. If the error message persists
        /// for a longer period of time, it can be used to surface error message to
        /// indicate real problems requiring user intervention.
        /// </summary>
        public readonly string? ErrorMessage;

        [OutputConstructor]
        private BareMetalClusterStatus(
            ImmutableArray<Outputs.BareMetalClusterStatusCondition> conditions,

            string? errorMessage)
        {
            Conditions = conditions;
            ErrorMessage = errorMessage;
        }
    }
}
