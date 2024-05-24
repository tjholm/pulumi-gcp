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
    public sealed class ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfig
    {
        /// <summary>
        /// The Compute Engine accelerator (GPU) configuration for these instances. Can be specified 
        /// multiple times.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfigAccelerator> Accelerators;
        /// <summary>
        /// Disk Config
        /// </summary>
        public readonly Outputs.ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfigDiskConfig? DiskConfig;
        /// <summary>
        /// List of auxiliary node group instance names which have been assigned to the cluster.
        /// </summary>
        public readonly ImmutableArray<string> InstanceNames;
        /// <summary>
        /// The name of a Google Compute Engine machine type
        /// to create for the node group. If not specified, GCP will default to a predetermined
        /// computed value (currently `n1-standard-4`).
        /// </summary>
        public readonly string? MachineType;
        /// <summary>
        /// The name of a minimum generation of CPU family
        /// for the node group. If not specified, GCP will default to a predetermined computed value
        /// for each zone. See [the guide](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
        /// for details about which CPU families are available (and defaulted) for each zone.
        /// </summary>
        public readonly string? MinCpuPlatform;
        /// <summary>
        /// Specifies the number of master nodes to create.
        /// Please set a number greater than 0. Node Group must have at least 1 instance.
        /// </summary>
        public readonly int? NumInstances;

        [OutputConstructor]
        private ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfig(
            ImmutableArray<Outputs.ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfigAccelerator> accelerators,

            Outputs.ClusterClusterConfigAuxiliaryNodeGroupNodeGroupNodeGroupConfigDiskConfig? diskConfig,

            ImmutableArray<string> instanceNames,

            string? machineType,

            string? minCpuPlatform,

            int? numInstances)
        {
            Accelerators = accelerators;
            DiskConfig = diskConfig;
            InstanceNames = instanceNames;
            MachineType = machineType;
            MinCpuPlatform = minCpuPlatform;
            NumInstances = numInstances;
        }
    }
}
