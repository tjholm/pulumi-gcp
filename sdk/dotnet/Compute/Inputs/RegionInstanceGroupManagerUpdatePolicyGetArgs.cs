// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionInstanceGroupManagerUpdatePolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance redistribution policy for regional managed instance groups. Valid values are: `"PROACTIVE"`, `"NONE"`. If `PROACTIVE` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If `NONE`, proactive redistribution is disabled.
        /// </summary>
        [Input("instanceRedistributionType")]
        public Input<string>? InstanceRedistributionType { get; set; }

        /// <summary>
        /// , The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with `max_surge_percent`. It has to be either 0 or at least equal to the number of zones.  If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
        /// </summary>
        [Input("maxSurgeFixed")]
        public Input<int>? MaxSurgeFixed { get; set; }

        /// <summary>
        /// , The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with `max_surge_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
        /// </summary>
        [Input("maxSurgePercent")]
        public Input<int>? MaxSurgePercent { get; set; }

        /// <summary>
        /// , The maximum number of instances that can be unavailable during the update process. Conflicts with `max_unavailable_percent`. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of `max_unavailable_fixed` or `max_surge_fixed` must be greater than 0.
        /// </summary>
        [Input("maxUnavailableFixed")]
        public Input<int>? MaxUnavailableFixed { get; set; }

        /// <summary>
        /// , The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with `max_unavailable_fixed`. Percent value is only allowed for regional managed instance groups with size at least 10.
        /// </summary>
        [Input("maxUnavailablePercent")]
        public Input<int>? MaxUnavailablePercent { get; set; }

        /// <summary>
        /// , Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]
        /// </summary>
        [Input("minReadySec")]
        public Input<int>? MinReadySec { get; set; }

        /// <summary>
        /// Minimal action to be taken on an instance. You can specify either `REFRESH` to update without stopping instances, `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `REFRESH`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
        /// </summary>
        [Input("minimalAction", required: true)]
        public Input<string> MinimalAction { get; set; } = null!;

        /// <summary>
        /// Most disruptive action that is allowed to be taken on an instance. You can specify either NONE to forbid any actions, REFRESH to allow actions that do not need instance restart, RESTART to allow actions that can be applied without instance replacing or REPLACE to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all.
        /// </summary>
        [Input("mostDisruptiveAllowedAction")]
        public Input<string>? MostDisruptiveAllowedAction { get; set; }

        /// <summary>
        /// , The instance replacement method for managed instance groups. Valid values are: "RECREATE", "SUBSTITUTE". If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
        /// - - -
        /// </summary>
        [Input("replacementMethod")]
        public Input<string>? ReplacementMethod { get; set; }

        /// <summary>
        /// The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RegionInstanceGroupManagerUpdatePolicyGetArgs()
        {
        }
        public static new RegionInstanceGroupManagerUpdatePolicyGetArgs Empty => new RegionInstanceGroupManagerUpdatePolicyGetArgs();
    }
}
