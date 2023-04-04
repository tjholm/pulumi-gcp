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
    public sealed class ClusterNodePoolNodeConfig
    {
        /// <summary>
        /// Specifies options for controlling
        /// advanced machine features. Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigAdvancedMachineFeatures? AdvancedMachineFeatures;
        /// <summary>
        /// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool. This should be of the form projects/[KEY_PROJECT_ID]/locations/[LOCATION]/keyRings/[RING_NAME]/cryptoKeys/[KEY_NAME]. For more information about protecting resources with Cloud KMS Keys please see: &lt;https://cloud.google.com/compute/docs/disks/customer-managed-encryption&gt;
        /// </summary>
        public readonly string? BootDiskKmsKey;
        /// <summary>
        /// Size of the disk attached to each node, specified
        /// in GB. The smallest allowed disk size is 10GB. Defaults to 100GB.
        /// </summary>
        public readonly int? DiskSizeGb;
        /// <summary>
        /// Type of the disk attached to each node
        /// (e.g. 'pd-standard', 'pd-balanced' or 'pd-ssd'). If unspecified, the default disk type is 'pd-standard'
        /// </summary>
        public readonly string? DiskType;
        /// <summary>
        /// Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigEphemeralStorageConfig? EphemeralStorageConfig;
        /// <summary>
        /// Parameters for the Google Container Filesystem (GCFS).
        /// If unspecified, GCFS will not be enabled on the node pool. When enabling this feature you must specify `image_type = "COS_CONTAINERD"` and `node_version` from GKE versions 1.19 or later to use it.
        /// For GKE versions 1.19, 1.20, and 1.21, the recommended minimum `node_version` would be 1.19.15-gke.1300, 1.20.11-gke.1300, and 1.21.5-gke.1300 respectively.
        /// A `machine_type` that has more than 16 GiB of memory is also recommended.
        /// GCFS must be enabled in order to use [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming).
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigGcfsConfig? GcfsConfig;
        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance.
        /// Structure documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterNodePoolNodeConfigGuestAccelerator> GuestAccelerators;
        /// <summary>
        /// Google Virtual NIC (gVNIC) is a virtual network interface.
        /// Installing the gVNIC driver allows for more efficient traffic transmission across the Google network infrastructure.
        /// gVNIC is an alternative to the virtIO-based ethernet driver. GKE nodes must use a Container-Optimized OS node image.
        /// GKE node version 1.15.11-gke.15 or later
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigGvnic? Gvnic;
        /// <summary>
        /// The image type to use for this node. Note that changing the image type
        /// will delete and recreate all nodes in the node pool.
        /// </summary>
        public readonly string? ImageType;
        /// <summary>
        /// Kubelet configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigKubeletConfig? KubeletConfig;
        /// <summary>
        /// The Kubernetes labels (key/value pairs) to be applied to each node. The kubernetes.io/ and k8s.io/ prefixes are
        /// reserved by Kubernetes Core components and cannot be specified.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;
        /// <summary>
        /// Linux node configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
        /// Note that validations happen all server side. All attributes are optional.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigLinuxNodeConfig? LinuxNodeConfig;
        /// <summary>
        /// Parameters for the local NVMe SSDs. Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig? LocalNvmeSsdBlockConfig;
        /// <summary>
        /// The amount of local SSD disks that will be
        /// attached to each cluster node. Defaults to 0.
        /// </summary>
        public readonly int? LocalSsdCount;
        /// <summary>
        /// Parameter for specifying the type of logging agent used in a node pool. This will override any cluster-wide default value. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
        /// </summary>
        public readonly string? LoggingVariant;
        /// <summary>
        /// The name of a Google Compute Engine machine type.
        /// Defaults to `e2-medium`. To create a custom machine type, value should be set as specified
        /// [here](https://cloud.google.com/compute/docs/reference/latest/instances#machineType).
        /// </summary>
        public readonly string? MachineType;
        /// <summary>
        /// The metadata key/value pairs assigned to instances in
        /// the cluster. From GKE `1.12` onwards, `disable-legacy-endpoints` is set to
        /// `true` by the API; if `metadata` is set but that default value is not
        /// included, the provider will attempt to unset the value. To avoid this, set the
        /// value in your config.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Metadata;
        /// <summary>
        /// Minimum CPU platform to be used by this instance.
        /// The instance may be scheduled on the specified or newer CPU platform. Applicable
        /// values are the friendly names of CPU platforms, such as `Intel Haswell`. See the
        /// [official documentation](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
        /// for more information.
        /// </summary>
        public readonly string? MinCpuPlatform;
        /// <summary>
        /// Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on [sole tenant nodes](https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes).
        /// </summary>
        public readonly string? NodeGroup;
        /// <summary>
        /// The set of Google API scopes to be made available
        /// on all of the node VMs under the "default" service account.
        /// Use the "https://www.googleapis.com/auth/cloud-platform" scope to grant access to all APIs. It is recommended that you set `service_account` to a non-default service account and grant IAM roles to that service account for only the resources that it needs.
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;
        /// <summary>
        /// A boolean that represents whether or not the underlying node VMs
        /// are preemptible. See the [official documentation](https://cloud.google.com/container-engine/docs/preemptible-vm)
        /// for more information. Defaults to false.
        /// </summary>
        public readonly bool? Preemptible;
        /// <summary>
        /// The configuration of the desired reservation which instances could take capacity from. Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigReservationAffinity? ReservationAffinity;
        /// <summary>
        /// The GCP labels (key/value pairs) to be applied to each node. Refer [here](https://cloud.google.com/kubernetes-engine/docs/how-to/creating-managing-labels)
        /// for how these labels are applied to clusters, node pools and nodes.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ResourceLabels;
        public readonly Outputs.ClusterNodePoolNodeConfigSandboxConfig? SandboxConfig;
        /// <summary>
        /// The service account to be used by the Node VMs.
        /// If not specified, the "default" service account is used.
        /// </summary>
        public readonly string? ServiceAccount;
        /// <summary>
        /// Shielded Instance options. Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigShieldedInstanceConfig? ShieldedInstanceConfig;
        /// <summary>
        /// A boolean that represents whether the underlying node VMs are spot.
        /// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms)
        /// for more information. Defaults to false.
        /// </summary>
        public readonly bool? Spot;
        /// <summary>
        /// The list of instance tags applied to all nodes. Tags are used to identify
        /// valid sources or targets for network firewalls.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// A list of [Kubernetes taints](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
        /// to apply to nodes. GKE's API can only set this field on cluster creation.
        /// However, GKE will add taints to your nodes if you enable certain features such
        /// as GPUs. If this field is set, any diffs on this field will cause the provider to
        /// recreate the underlying resource. Taint values can be updated safely in
        /// Kubernetes (eg. through `kubectl`), and it's recommended that you do not use
        /// this field to manage taints. If you do, `lifecycle.ignore_changes` is
        /// recommended. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterNodePoolNodeConfigTaint> Taints;
        /// <summary>
        /// Metadata configuration to expose to workloads on the node pool.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ClusterNodePoolNodeConfigWorkloadMetadataConfig? WorkloadMetadataConfig;

        [OutputConstructor]
        private ClusterNodePoolNodeConfig(
            Outputs.ClusterNodePoolNodeConfigAdvancedMachineFeatures? advancedMachineFeatures,

            string? bootDiskKmsKey,

            int? diskSizeGb,

            string? diskType,

            Outputs.ClusterNodePoolNodeConfigEphemeralStorageConfig? ephemeralStorageConfig,

            Outputs.ClusterNodePoolNodeConfigGcfsConfig? gcfsConfig,

            ImmutableArray<Outputs.ClusterNodePoolNodeConfigGuestAccelerator> guestAccelerators,

            Outputs.ClusterNodePoolNodeConfigGvnic? gvnic,

            string? imageType,

            Outputs.ClusterNodePoolNodeConfigKubeletConfig? kubeletConfig,

            ImmutableDictionary<string, string>? labels,

            Outputs.ClusterNodePoolNodeConfigLinuxNodeConfig? linuxNodeConfig,

            Outputs.ClusterNodePoolNodeConfigLocalNvmeSsdBlockConfig? localNvmeSsdBlockConfig,

            int? localSsdCount,

            string? loggingVariant,

            string? machineType,

            ImmutableDictionary<string, string>? metadata,

            string? minCpuPlatform,

            string? nodeGroup,

            ImmutableArray<string> oauthScopes,

            bool? preemptible,

            Outputs.ClusterNodePoolNodeConfigReservationAffinity? reservationAffinity,

            ImmutableDictionary<string, string>? resourceLabels,

            Outputs.ClusterNodePoolNodeConfigSandboxConfig? sandboxConfig,

            string? serviceAccount,

            Outputs.ClusterNodePoolNodeConfigShieldedInstanceConfig? shieldedInstanceConfig,

            bool? spot,

            ImmutableArray<string> tags,

            ImmutableArray<Outputs.ClusterNodePoolNodeConfigTaint> taints,

            Outputs.ClusterNodePoolNodeConfigWorkloadMetadataConfig? workloadMetadataConfig)
        {
            AdvancedMachineFeatures = advancedMachineFeatures;
            BootDiskKmsKey = bootDiskKmsKey;
            DiskSizeGb = diskSizeGb;
            DiskType = diskType;
            EphemeralStorageConfig = ephemeralStorageConfig;
            GcfsConfig = gcfsConfig;
            GuestAccelerators = guestAccelerators;
            Gvnic = gvnic;
            ImageType = imageType;
            KubeletConfig = kubeletConfig;
            Labels = labels;
            LinuxNodeConfig = linuxNodeConfig;
            LocalNvmeSsdBlockConfig = localNvmeSsdBlockConfig;
            LocalSsdCount = localSsdCount;
            LoggingVariant = loggingVariant;
            MachineType = machineType;
            Metadata = metadata;
            MinCpuPlatform = minCpuPlatform;
            NodeGroup = nodeGroup;
            OauthScopes = oauthScopes;
            Preemptible = preemptible;
            ReservationAffinity = reservationAffinity;
            ResourceLabels = resourceLabels;
            SandboxConfig = sandboxConfig;
            ServiceAccount = serviceAccount;
            ShieldedInstanceConfig = shieldedInstanceConfig;
            Spot = spot;
            Tags = tags;
            Taints = taints;
            WorkloadMetadataConfig = workloadMetadataConfig;
        }
    }
}
