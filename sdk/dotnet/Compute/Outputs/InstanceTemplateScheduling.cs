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
    public sealed class InstanceTemplateScheduling
    {
        /// <summary>
        /// Specifies whether the instance should be
        /// automatically restarted if it is terminated by Compute Engine (not
        /// terminated by a user). This defaults to true.
        /// </summary>
        public readonly bool? AutomaticRestart;
        /// <summary>
        /// Describe the type of termination action for `SPOT` VM. Can be `STOP` or `DELETE`.  Read more on [here](https://cloud.google.com/compute/docs/instances/create-use-spot)
        /// </summary>
        public readonly string? InstanceTerminationAction;
        /// <summary>
        /// Specifies the maximum amount of time a Local Ssd Vm should wait while
        ///   recovery of the Local Ssd state is attempted. Its value should be in
        ///   between 0 and 168 hours with hour granularity and the default value being 1
        ///   hour.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTemplateSchedulingLocalSsdRecoveryTimeout> LocalSsdRecoveryTimeouts;
        /// <summary>
        /// Specifies the frequency of planned maintenance events. The accepted values are: `PERIODIC`.
        /// </summary>
        public readonly string? MaintenanceInterval;
        /// <summary>
        /// The duration of the instance. Instance will run and be terminated after then, the termination action could be defined in `instance_termination_action`. Only support `DELETE` `instance_termination_action` at this point. Structure is documented below.
        /// &lt;a name="nested_max_run_duration"&gt;&lt;/a&gt;The `max_run_duration` block supports:
        /// </summary>
        public readonly Outputs.InstanceTemplateSchedulingMaxRunDuration? MaxRunDuration;
        /// <summary>
        /// Minimum number of cpus for the instance.
        /// </summary>
        public readonly int? MinNodeCpus;
        /// <summary>
        /// Specifies node affinities or anti-affinities
        /// to determine which sole-tenant nodes your instances and managed instance
        /// groups will use as host systems. Read more on sole-tenant node creation
        /// [here](https://cloud.google.com/compute/docs/nodes/create-nodes).
        /// Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTemplateSchedulingNodeAffinity> NodeAffinities;
        /// <summary>
        /// Defines the maintenance behavior for this
        /// instance.
        /// </summary>
        public readonly string? OnHostMaintenance;
        /// <summary>
        /// Allows instance to be preempted. This defaults to
        /// false. Read more on this
        /// [here](https://cloud.google.com/compute/docs/instances/preemptible).
        /// </summary>
        public readonly bool? Preemptible;
        /// <summary>
        /// Describe the type of preemptible VM. This field accepts the value `STANDARD` or `SPOT`. If the value is `STANDARD`, there will be no discount. If this   is set to `SPOT`,
        /// `preemptible` should be `true` and `automatic_restart` should be
        /// `false`. For more info about
        /// `SPOT`, read [here](https://cloud.google.com/compute/docs/instances/spot)
        /// </summary>
        public readonly string? ProvisioningModel;

        [OutputConstructor]
        private InstanceTemplateScheduling(
            bool? automaticRestart,

            string? instanceTerminationAction,

            ImmutableArray<Outputs.InstanceTemplateSchedulingLocalSsdRecoveryTimeout> localSsdRecoveryTimeouts,

            string? maintenanceInterval,

            Outputs.InstanceTemplateSchedulingMaxRunDuration? maxRunDuration,

            int? minNodeCpus,

            ImmutableArray<Outputs.InstanceTemplateSchedulingNodeAffinity> nodeAffinities,

            string? onHostMaintenance,

            bool? preemptible,

            string? provisioningModel)
        {
            AutomaticRestart = automaticRestart;
            InstanceTerminationAction = instanceTerminationAction;
            LocalSsdRecoveryTimeouts = localSsdRecoveryTimeouts;
            MaintenanceInterval = maintenanceInterval;
            MaxRunDuration = maxRunDuration;
            MinNodeCpus = minNodeCpus;
            NodeAffinities = nodeAffinities;
            OnHostMaintenance = onHostMaintenance;
            Preemptible = preemptible;
            ProvisioningModel = provisioningModel;
        }
    }
}
