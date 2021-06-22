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
    public sealed class WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
    {
        /// <summary>
        /// Optional. The Compute Engine accelerator configuration for these instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerator> Accelerators;
        /// <summary>
        /// Optional. Disk option config settings.
        /// </summary>
        public readonly Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig? DiskConfig;
        /// <summary>
        /// Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.
        /// </summary>
        public readonly string? Image;
        /// <summary>
        /// -
        /// Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.
        /// </summary>
        public readonly ImmutableArray<string> InstanceNames;
        /// <summary>
        /// -
        /// Output only. Specifies that this instance group contains preemptible instances.
        /// </summary>
        public readonly bool? IsPreemptible;
        /// <summary>
        /// Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/(https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.
        /// </summary>
        public readonly string? MachineType;
        /// <summary>
        /// -
        /// Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig> ManagedGroupConfigs;
        /// <summary>
        /// Optional. Specifies the minimum cpu platform for the Instance Group. See (https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
        /// </summary>
        public readonly string? MinCpuPlatform;
        /// <summary>
        /// Optional. The number of VM instances in the instance group. For master instance groups, must be set to 1.
        /// </summary>
        public readonly int? NumInstances;
        /// <summary>
        /// Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE
        /// </summary>
        public readonly string? Preemptibility;

        [OutputConstructor]
        private WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(
            ImmutableArray<Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerator> accelerators,

            Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig? diskConfig,

            string? image,

            ImmutableArray<string> instanceNames,

            bool? isPreemptible,

            string? machineType,

            ImmutableArray<Outputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig> managedGroupConfigs,

            string? minCpuPlatform,

            int? numInstances,

            string? preemptibility)
        {
            Accelerators = accelerators;
            DiskConfig = diskConfig;
            Image = image;
            InstanceNames = instanceNames;
            IsPreemptible = isPreemptible;
            MachineType = machineType;
            ManagedGroupConfigs = managedGroupConfigs;
            MinCpuPlatform = minCpuPlatform;
            NumInstances = numInstances;
            Preemptibility = preemptibility;
        }
    }
}
