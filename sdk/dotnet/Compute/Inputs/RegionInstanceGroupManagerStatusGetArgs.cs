// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceGroupManagerStatusGetArgs : Pulumi.ResourceArgs
    {
        [Input("allInstancesConfigs")]
        private InputList<Inputs.RegionInstanceGroupManagerStatusAllInstancesConfigGetArgs>? _allInstancesConfigs;

        /// <summary>
        /// )
        /// Properties to set on all instances in the group. After setting
        /// allInstancesConfig on the group, you must update the group's instances to
        /// apply the configuration.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerStatusAllInstancesConfigGetArgs> AllInstancesConfigs
        {
            get => _allInstancesConfigs ?? (_allInstancesConfigs = new InputList<Inputs.RegionInstanceGroupManagerStatusAllInstancesConfigGetArgs>());
            set => _allInstancesConfigs = value;
        }

        /// <summary>
        /// A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified.
        /// </summary>
        [Input("isStable")]
        public Input<bool>? IsStable { get; set; }

        [Input("statefuls")]
        private InputList<Inputs.RegionInstanceGroupManagerStatusStatefulGetArgs>? _statefuls;

        /// <summary>
        /// Stateful status of the given Instance Group Manager.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerStatusStatefulGetArgs> Statefuls
        {
            get => _statefuls ?? (_statefuls = new InputList<Inputs.RegionInstanceGroupManagerStatusStatefulGetArgs>());
            set => _statefuls = value;
        }

        [Input("versionTargets")]
        private InputList<Inputs.RegionInstanceGroupManagerStatusVersionTargetGetArgs>? _versionTargets;

        /// <summary>
        /// A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances' target version are specified by version field on Instance Group Manager.
        /// </summary>
        public InputList<Inputs.RegionInstanceGroupManagerStatusVersionTargetGetArgs> VersionTargets
        {
            get => _versionTargets ?? (_versionTargets = new InputList<Inputs.RegionInstanceGroupManagerStatusVersionTargetGetArgs>());
            set => _versionTargets = value;
        }

        public RegionInstanceGroupManagerStatusGetArgs()
        {
        }
    }
}
