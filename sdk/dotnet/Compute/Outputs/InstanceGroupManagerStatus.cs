// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class InstanceGroupManagerStatus
    {
        /// <summary>
        /// A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
        /// </summary>
        public readonly bool? IsStable;
        /// <summary>
        /// Stateful status of the given Instance Group Manager.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceGroupManagerStatusStateful> Statefuls;
        /// <summary>
        /// A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances' target version are specified by version field on Instance Group Manager.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceGroupManagerStatusVersionTarget> VersionTargets;

        [OutputConstructor]
        private InstanceGroupManagerStatus(
            bool? isStable,

            ImmutableArray<Outputs.InstanceGroupManagerStatusStateful> statefuls,

            ImmutableArray<Outputs.InstanceGroupManagerStatusVersionTarget> versionTargets)
        {
            IsStable = isStable;
            Statefuls = statefuls;
            VersionTargets = versionTargets;
        }
    }
}
