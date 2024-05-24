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
    public sealed class ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList
    {
        /// <summary>
        /// Full machine-type names, e.g. `"n1-standard-16"`.
        /// </summary>
        public readonly ImmutableArray<string> MachineTypes;
        /// <summary>
        /// Preference of this instance selection. A lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
        /// 
        /// - - -
        /// </summary>
        public readonly int? Rank;

        [OutputConstructor]
        private ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList(
            ImmutableArray<string> machineTypes,

            int? rank)
        {
            MachineTypes = machineTypes;
            Rank = rank;
        }
    }
}
