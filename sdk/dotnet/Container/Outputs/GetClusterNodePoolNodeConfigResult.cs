// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class GetClusterNodePoolNodeConfigResult
    {
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigAdvancedMachineFeatureResult> AdvancedMachineFeatures;
        public readonly string BootDiskKmsKey;
        public readonly int DiskSizeGb;
        public readonly string DiskType;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigEphemeralStorageConfigResult> EphemeralStorageConfigs;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGcfsConfigResult> GcfsConfigs;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGuestAcceleratorResult> GuestAccelerators;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGvnicResult> Gvnics;
        public readonly string ImageType;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigKubeletConfigResult> KubeletConfigs;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigLinuxNodeConfigResult> LinuxNodeConfigs;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigResult> LocalNvmeSsdBlockConfigs;
        public readonly int LocalSsdCount;
        public readonly string LoggingVariant;
        public readonly string MachineType;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string MinCpuPlatform;
        public readonly string NodeGroup;
        public readonly ImmutableArray<string> OauthScopes;
        public readonly bool Preemptible;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigReservationAffinityResult> ReservationAffinities;
        public readonly ImmutableDictionary<string, string> ResourceLabels;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigSandboxConfigResult> SandboxConfigs;
        public readonly string ServiceAccount;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigShieldedInstanceConfigResult> ShieldedInstanceConfigs;
        public readonly bool Spot;
        public readonly ImmutableArray<string> Tags;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigTaintResult> Taints;
        public readonly ImmutableArray<Outputs.GetClusterNodePoolNodeConfigWorkloadMetadataConfigResult> WorkloadMetadataConfigs;

        [OutputConstructor]
        private GetClusterNodePoolNodeConfigResult(
            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigAdvancedMachineFeatureResult> advancedMachineFeatures,

            string bootDiskKmsKey,

            int diskSizeGb,

            string diskType,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigEphemeralStorageConfigResult> ephemeralStorageConfigs,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGcfsConfigResult> gcfsConfigs,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGuestAcceleratorResult> guestAccelerators,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigGvnicResult> gvnics,

            string imageType,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigKubeletConfigResult> kubeletConfigs,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigLinuxNodeConfigResult> linuxNodeConfigs,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigResult> localNvmeSsdBlockConfigs,

            int localSsdCount,

            string loggingVariant,

            string machineType,

            ImmutableDictionary<string, string> metadata,

            string minCpuPlatform,

            string nodeGroup,

            ImmutableArray<string> oauthScopes,

            bool preemptible,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigReservationAffinityResult> reservationAffinities,

            ImmutableDictionary<string, string> resourceLabels,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigSandboxConfigResult> sandboxConfigs,

            string serviceAccount,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigShieldedInstanceConfigResult> shieldedInstanceConfigs,

            bool spot,

            ImmutableArray<string> tags,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigTaintResult> taints,

            ImmutableArray<Outputs.GetClusterNodePoolNodeConfigWorkloadMetadataConfigResult> workloadMetadataConfigs)
        {
            AdvancedMachineFeatures = advancedMachineFeatures;
            BootDiskKmsKey = bootDiskKmsKey;
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            EphemeralStorageConfigs = ephemeralStorageConfigs;
            GcfsConfigs = gcfsConfigs;
            GuestAccelerators = guestAccelerators;
            Gvnics = gvnics;
            ImageType = imageType;
            KubeletConfigs = kubeletConfigs;
            Labels = labels;
            LinuxNodeConfigs = linuxNodeConfigs;
            LocalNvmeSsdBlockConfigs = localNvmeSsdBlockConfigs;
            LocalSsdCount = localSsdCount;
            LoggingVariant = loggingVariant;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            NodeGroup = nodeGroup;
            OauthScopes = oauthScopes;
            Preemptible = preemptible;
            ReservationAffinities = reservationAffinities;
            ResourceLabels = resourceLabels;
            SandboxConfigs = sandboxConfigs;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Spot = spot;
            Tags = tags;
            Taints = taints;
            WorkloadMetadataConfigs = workloadMetadataConfigs;
        }
    }
}
