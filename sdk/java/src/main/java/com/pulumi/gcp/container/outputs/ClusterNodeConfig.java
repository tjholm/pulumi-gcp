// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigEphemeralStorageConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigGcfsConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigGuestAccelerator;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigGvnic;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigKubeletConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigLinuxNodeConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigReservationAffinity;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigSandboxConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigShieldedInstanceConfig;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigTaint;
import com.pulumi.gcp.container.outputs.ClusterNodeConfigWorkloadMetadataConfig;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterNodeConfig {
    /**
     * @return The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool. This should be of the form projects/[KEY_PROJECT_ID]/locations/[LOCATION]/keyRings/[RING_NAME]/cryptoKeys/[KEY_NAME]. For more information about protecting resources with Cloud KMS Keys please see: https://cloud.google.com/compute/docs/disks/customer-managed-encryption
     * 
     */
    private @Nullable String bootDiskKmsKey;
    /**
     * @return Size of the disk attached to each node, specified
     * in GB. The smallest allowed disk size is 10GB. Defaults to 100GB.
     * 
     */
    private @Nullable Integer diskSizeGb;
    /**
     * @return Type of the disk attached to each node
     * (e.g. &#39;pd-standard&#39;, &#39;pd-balanced&#39; or &#39;pd-ssd&#39;). If unspecified, the default disk type is &#39;pd-standard&#39;
     * 
     */
    private @Nullable String diskType;
    /**
     * @return Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigEphemeralStorageConfig ephemeralStorageConfig;
    /**
     * @return Parameters for the Google Container Filesystem (GCFS).
     * If unspecified, GCFS will not be enabled on the node pool. When enabling this feature you must specify `image_type = &#34;COS_CONTAINERD&#34;` and `node_version` from GKE versions 1.19 or later to use it.
     * For GKE versions 1.19, 1.20, and 1.21, the recommended minimum `node_version` would be 1.19.15-gke.1300, 1.20.11-gke.1300, and 1.21.5-gke.1300 respectively.
     * A `machine_type` that has more than 16 GiB of memory is also recommended.
     * GCFS must be enabled in order to use [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming).
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigGcfsConfig gcfsConfig;
    private @Nullable List<ClusterNodeConfigGuestAccelerator> guestAccelerators;
    /**
     * @return Google Virtual NIC (gVNIC) is a virtual network interface.
     * Installing the gVNIC driver allows for more efficient traffic transmission across the Google network infrastructure.
     * gVNIC is an alternative to the virtIO-based ethernet driver. GKE nodes must use a Container-Optimized OS node image.
     * GKE node version 1.15.11-gke.15 or later
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigGvnic gvnic;
    /**
     * @return The image type to use for this node. Note that changing the image type
     * will delete and recreate all nodes in the node pool.
     * 
     */
    private @Nullable String imageType;
    /**
     * @return Kubelet configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigKubeletConfig kubeletConfig;
    /**
     * @return The Kubernetes labels (key/value pairs) to be applied to each node. The kubernetes.io/ and k8s.io/ prefixes are
     * reserved by Kubernetes Core components and cannot be specified.
     * 
     */
    private @Nullable Map<String,String> labels;
    /**
     * @return Linux node configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
     * Note that validations happen all server side. All attributes are optional.
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigLinuxNodeConfig linuxNodeConfig;
    /**
     * @return The amount of local SSD disks that will be
     * attached to each cluster node. Defaults to 0.
     * 
     */
    private @Nullable Integer localSsdCount;
    /**
     * @return Parameter for specifying the type of logging agent used in a node pool. This will override any cluster-wide default value. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
     * 
     */
    private @Nullable String loggingVariant;
    /**
     * @return The name of a Google Compute Engine machine type.
     * Defaults to `e2-medium`. To create a custom machine type, value should be set as specified
     * [here](https://cloud.google.com/compute/docs/reference/latest/instances#machineType).
     * 
     */
    private @Nullable String machineType;
    private @Nullable Map<String,String> metadata;
    /**
     * @return Minimum CPU platform to be used by this instance.
     * The instance may be scheduled on the specified or newer CPU platform. Applicable
     * values are the friendly names of CPU platforms, such as `Intel Haswell`. See the
     * [official documentation](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
     * for more information.
     * 
     */
    private @Nullable String minCpuPlatform;
    /**
     * @return Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on [sole tenant nodes](https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes).
     * 
     */
    private @Nullable String nodeGroup;
    /**
     * @return The set of Google API scopes to be made available
     * on all of the node VMs under the &#34;default&#34; service account.
     * Use the &#34;https://www.googleapis.com/auth/cloud-platform&#34; scope to grant access to all APIs. It is recommended that you set `service_account` to a non-default service account and grant IAM roles to that service account for only the resources that it needs.
     * 
     */
    private @Nullable List<String> oauthScopes;
    /**
     * @return A boolean that represents whether or not the underlying node VMs
     * are preemptible. See the [official documentation](https://cloud.google.com/container-engine/docs/preemptible-vm)
     * for more information. Defaults to false.
     * 
     */
    private @Nullable Boolean preemptible;
    /**
     * @return The configuration of the desired reservation which instances could take capacity from. Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigReservationAffinity reservationAffinity;
    /**
     * @return The GCP labels (key/value pairs) to be applied to each node. Refer [here](https://cloud.google.com/kubernetes-engine/docs/how-to/creating-managing-labels)
     * for how these labels are applied to clusters, node pools and nodes.
     * 
     */
    private @Nullable Map<String,String> resourceLabels;
    /**
     * @return ) [GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/how-to/sandbox-pods) configuration. When enabling this feature you must specify `image_type = &#34;COS_CONTAINERD&#34;` and `node_version = &#34;1.12.7-gke.17&#34;` or later to use it.
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigSandboxConfig sandboxConfig;
    /**
     * @return The service account to be used by the Node VMs.
     * If not specified, the &#34;default&#34; service account is used.
     * 
     */
    private @Nullable String serviceAccount;
    /**
     * @return Shielded Instance options. Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigShieldedInstanceConfig shieldedInstanceConfig;
    /**
     * @return A boolean that represents whether the underlying node VMs are spot.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms)
     * for more information. Defaults to false.
     * 
     */
    private @Nullable Boolean spot;
    /**
     * @return The list of instance tags applied to all nodes. Tags are used to identify
     * valid sources or targets for network firewalls.
     * 
     */
    private @Nullable List<String> tags;
    private @Nullable List<ClusterNodeConfigTaint> taints;
    /**
     * @return Metadata configuration to expose to workloads on the node pool.
     * Structure is documented below.
     * 
     */
    private @Nullable ClusterNodeConfigWorkloadMetadataConfig workloadMetadataConfig;

    private ClusterNodeConfig() {}
    /**
     * @return The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool. This should be of the form projects/[KEY_PROJECT_ID]/locations/[LOCATION]/keyRings/[RING_NAME]/cryptoKeys/[KEY_NAME]. For more information about protecting resources with Cloud KMS Keys please see: https://cloud.google.com/compute/docs/disks/customer-managed-encryption
     * 
     */
    public Optional<String> bootDiskKmsKey() {
        return Optional.ofNullable(this.bootDiskKmsKey);
    }
    /**
     * @return Size of the disk attached to each node, specified
     * in GB. The smallest allowed disk size is 10GB. Defaults to 100GB.
     * 
     */
    public Optional<Integer> diskSizeGb() {
        return Optional.ofNullable(this.diskSizeGb);
    }
    /**
     * @return Type of the disk attached to each node
     * (e.g. &#39;pd-standard&#39;, &#39;pd-balanced&#39; or &#39;pd-ssd&#39;). If unspecified, the default disk type is &#39;pd-standard&#39;
     * 
     */
    public Optional<String> diskType() {
        return Optional.ofNullable(this.diskType);
    }
    /**
     * @return Parameters for the ephemeral storage filesystem. If unspecified, ephemeral storage is backed by the boot disk. Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigEphemeralStorageConfig> ephemeralStorageConfig() {
        return Optional.ofNullable(this.ephemeralStorageConfig);
    }
    /**
     * @return Parameters for the Google Container Filesystem (GCFS).
     * If unspecified, GCFS will not be enabled on the node pool. When enabling this feature you must specify `image_type = &#34;COS_CONTAINERD&#34;` and `node_version` from GKE versions 1.19 or later to use it.
     * For GKE versions 1.19, 1.20, and 1.21, the recommended minimum `node_version` would be 1.19.15-gke.1300, 1.20.11-gke.1300, and 1.21.5-gke.1300 respectively.
     * A `machine_type` that has more than 16 GiB of memory is also recommended.
     * GCFS must be enabled in order to use [image streaming](https://cloud.google.com/kubernetes-engine/docs/how-to/image-streaming).
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigGcfsConfig> gcfsConfig() {
        return Optional.ofNullable(this.gcfsConfig);
    }
    public List<ClusterNodeConfigGuestAccelerator> guestAccelerators() {
        return this.guestAccelerators == null ? List.of() : this.guestAccelerators;
    }
    /**
     * @return Google Virtual NIC (gVNIC) is a virtual network interface.
     * Installing the gVNIC driver allows for more efficient traffic transmission across the Google network infrastructure.
     * gVNIC is an alternative to the virtIO-based ethernet driver. GKE nodes must use a Container-Optimized OS node image.
     * GKE node version 1.15.11-gke.15 or later
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigGvnic> gvnic() {
        return Optional.ofNullable(this.gvnic);
    }
    /**
     * @return The image type to use for this node. Note that changing the image type
     * will delete and recreate all nodes in the node pool.
     * 
     */
    public Optional<String> imageType() {
        return Optional.ofNullable(this.imageType);
    }
    /**
     * @return Kubelet configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigKubeletConfig> kubeletConfig() {
        return Optional.ofNullable(this.kubeletConfig);
    }
    /**
     * @return The Kubernetes labels (key/value pairs) to be applied to each node. The kubernetes.io/ and k8s.io/ prefixes are
     * reserved by Kubernetes Core components and cannot be specified.
     * 
     */
    public Map<String,String> labels() {
        return this.labels == null ? Map.of() : this.labels;
    }
    /**
     * @return Linux node configuration, currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
     * Note that validations happen all server side. All attributes are optional.
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigLinuxNodeConfig> linuxNodeConfig() {
        return Optional.ofNullable(this.linuxNodeConfig);
    }
    /**
     * @return The amount of local SSD disks that will be
     * attached to each cluster node. Defaults to 0.
     * 
     */
    public Optional<Integer> localSsdCount() {
        return Optional.ofNullable(this.localSsdCount);
    }
    /**
     * @return Parameter for specifying the type of logging agent used in a node pool. This will override any cluster-wide default value. Valid values include DEFAULT and MAX_THROUGHPUT. See [Increasing logging agent throughput](https://cloud.google.com/stackdriver/docs/solutions/gke/managing-logs#throughput) for more information.
     * 
     */
    public Optional<String> loggingVariant() {
        return Optional.ofNullable(this.loggingVariant);
    }
    /**
     * @return The name of a Google Compute Engine machine type.
     * Defaults to `e2-medium`. To create a custom machine type, value should be set as specified
     * [here](https://cloud.google.com/compute/docs/reference/latest/instances#machineType).
     * 
     */
    public Optional<String> machineType() {
        return Optional.ofNullable(this.machineType);
    }
    public Map<String,String> metadata() {
        return this.metadata == null ? Map.of() : this.metadata;
    }
    /**
     * @return Minimum CPU platform to be used by this instance.
     * The instance may be scheduled on the specified or newer CPU platform. Applicable
     * values are the friendly names of CPU platforms, such as `Intel Haswell`. See the
     * [official documentation](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
     * for more information.
     * 
     */
    public Optional<String> minCpuPlatform() {
        return Optional.ofNullable(this.minCpuPlatform);
    }
    /**
     * @return Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on [sole tenant nodes](https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes).
     * 
     */
    public Optional<String> nodeGroup() {
        return Optional.ofNullable(this.nodeGroup);
    }
    /**
     * @return The set of Google API scopes to be made available
     * on all of the node VMs under the &#34;default&#34; service account.
     * Use the &#34;https://www.googleapis.com/auth/cloud-platform&#34; scope to grant access to all APIs. It is recommended that you set `service_account` to a non-default service account and grant IAM roles to that service account for only the resources that it needs.
     * 
     */
    public List<String> oauthScopes() {
        return this.oauthScopes == null ? List.of() : this.oauthScopes;
    }
    /**
     * @return A boolean that represents whether or not the underlying node VMs
     * are preemptible. See the [official documentation](https://cloud.google.com/container-engine/docs/preemptible-vm)
     * for more information. Defaults to false.
     * 
     */
    public Optional<Boolean> preemptible() {
        return Optional.ofNullable(this.preemptible);
    }
    /**
     * @return The configuration of the desired reservation which instances could take capacity from. Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigReservationAffinity> reservationAffinity() {
        return Optional.ofNullable(this.reservationAffinity);
    }
    /**
     * @return The GCP labels (key/value pairs) to be applied to each node. Refer [here](https://cloud.google.com/kubernetes-engine/docs/how-to/creating-managing-labels)
     * for how these labels are applied to clusters, node pools and nodes.
     * 
     */
    public Map<String,String> resourceLabels() {
        return this.resourceLabels == null ? Map.of() : this.resourceLabels;
    }
    /**
     * @return ) [GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/how-to/sandbox-pods) configuration. When enabling this feature you must specify `image_type = &#34;COS_CONTAINERD&#34;` and `node_version = &#34;1.12.7-gke.17&#34;` or later to use it.
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigSandboxConfig> sandboxConfig() {
        return Optional.ofNullable(this.sandboxConfig);
    }
    /**
     * @return The service account to be used by the Node VMs.
     * If not specified, the &#34;default&#34; service account is used.
     * 
     */
    public Optional<String> serviceAccount() {
        return Optional.ofNullable(this.serviceAccount);
    }
    /**
     * @return Shielded Instance options. Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigShieldedInstanceConfig> shieldedInstanceConfig() {
        return Optional.ofNullable(this.shieldedInstanceConfig);
    }
    /**
     * @return A boolean that represents whether the underlying node VMs are spot.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms)
     * for more information. Defaults to false.
     * 
     */
    public Optional<Boolean> spot() {
        return Optional.ofNullable(this.spot);
    }
    /**
     * @return The list of instance tags applied to all nodes. Tags are used to identify
     * valid sources or targets for network firewalls.
     * 
     */
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    public List<ClusterNodeConfigTaint> taints() {
        return this.taints == null ? List.of() : this.taints;
    }
    /**
     * @return Metadata configuration to expose to workloads on the node pool.
     * Structure is documented below.
     * 
     */
    public Optional<ClusterNodeConfigWorkloadMetadataConfig> workloadMetadataConfig() {
        return Optional.ofNullable(this.workloadMetadataConfig);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterNodeConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String bootDiskKmsKey;
        private @Nullable Integer diskSizeGb;
        private @Nullable String diskType;
        private @Nullable ClusterNodeConfigEphemeralStorageConfig ephemeralStorageConfig;
        private @Nullable ClusterNodeConfigGcfsConfig gcfsConfig;
        private @Nullable List<ClusterNodeConfigGuestAccelerator> guestAccelerators;
        private @Nullable ClusterNodeConfigGvnic gvnic;
        private @Nullable String imageType;
        private @Nullable ClusterNodeConfigKubeletConfig kubeletConfig;
        private @Nullable Map<String,String> labels;
        private @Nullable ClusterNodeConfigLinuxNodeConfig linuxNodeConfig;
        private @Nullable Integer localSsdCount;
        private @Nullable String loggingVariant;
        private @Nullable String machineType;
        private @Nullable Map<String,String> metadata;
        private @Nullable String minCpuPlatform;
        private @Nullable String nodeGroup;
        private @Nullable List<String> oauthScopes;
        private @Nullable Boolean preemptible;
        private @Nullable ClusterNodeConfigReservationAffinity reservationAffinity;
        private @Nullable Map<String,String> resourceLabels;
        private @Nullable ClusterNodeConfigSandboxConfig sandboxConfig;
        private @Nullable String serviceAccount;
        private @Nullable ClusterNodeConfigShieldedInstanceConfig shieldedInstanceConfig;
        private @Nullable Boolean spot;
        private @Nullable List<String> tags;
        private @Nullable List<ClusterNodeConfigTaint> taints;
        private @Nullable ClusterNodeConfigWorkloadMetadataConfig workloadMetadataConfig;
        public Builder() {}
        public Builder(ClusterNodeConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bootDiskKmsKey = defaults.bootDiskKmsKey;
    	      this.diskSizeGb = defaults.diskSizeGb;
    	      this.diskType = defaults.diskType;
    	      this.ephemeralStorageConfig = defaults.ephemeralStorageConfig;
    	      this.gcfsConfig = defaults.gcfsConfig;
    	      this.guestAccelerators = defaults.guestAccelerators;
    	      this.gvnic = defaults.gvnic;
    	      this.imageType = defaults.imageType;
    	      this.kubeletConfig = defaults.kubeletConfig;
    	      this.labels = defaults.labels;
    	      this.linuxNodeConfig = defaults.linuxNodeConfig;
    	      this.localSsdCount = defaults.localSsdCount;
    	      this.loggingVariant = defaults.loggingVariant;
    	      this.machineType = defaults.machineType;
    	      this.metadata = defaults.metadata;
    	      this.minCpuPlatform = defaults.minCpuPlatform;
    	      this.nodeGroup = defaults.nodeGroup;
    	      this.oauthScopes = defaults.oauthScopes;
    	      this.preemptible = defaults.preemptible;
    	      this.reservationAffinity = defaults.reservationAffinity;
    	      this.resourceLabels = defaults.resourceLabels;
    	      this.sandboxConfig = defaults.sandboxConfig;
    	      this.serviceAccount = defaults.serviceAccount;
    	      this.shieldedInstanceConfig = defaults.shieldedInstanceConfig;
    	      this.spot = defaults.spot;
    	      this.tags = defaults.tags;
    	      this.taints = defaults.taints;
    	      this.workloadMetadataConfig = defaults.workloadMetadataConfig;
        }

        @CustomType.Setter
        public Builder bootDiskKmsKey(@Nullable String bootDiskKmsKey) {
            this.bootDiskKmsKey = bootDiskKmsKey;
            return this;
        }
        @CustomType.Setter
        public Builder diskSizeGb(@Nullable Integer diskSizeGb) {
            this.diskSizeGb = diskSizeGb;
            return this;
        }
        @CustomType.Setter
        public Builder diskType(@Nullable String diskType) {
            this.diskType = diskType;
            return this;
        }
        @CustomType.Setter
        public Builder ephemeralStorageConfig(@Nullable ClusterNodeConfigEphemeralStorageConfig ephemeralStorageConfig) {
            this.ephemeralStorageConfig = ephemeralStorageConfig;
            return this;
        }
        @CustomType.Setter
        public Builder gcfsConfig(@Nullable ClusterNodeConfigGcfsConfig gcfsConfig) {
            this.gcfsConfig = gcfsConfig;
            return this;
        }
        @CustomType.Setter
        public Builder guestAccelerators(@Nullable List<ClusterNodeConfigGuestAccelerator> guestAccelerators) {
            this.guestAccelerators = guestAccelerators;
            return this;
        }
        public Builder guestAccelerators(ClusterNodeConfigGuestAccelerator... guestAccelerators) {
            return guestAccelerators(List.of(guestAccelerators));
        }
        @CustomType.Setter
        public Builder gvnic(@Nullable ClusterNodeConfigGvnic gvnic) {
            this.gvnic = gvnic;
            return this;
        }
        @CustomType.Setter
        public Builder imageType(@Nullable String imageType) {
            this.imageType = imageType;
            return this;
        }
        @CustomType.Setter
        public Builder kubeletConfig(@Nullable ClusterNodeConfigKubeletConfig kubeletConfig) {
            this.kubeletConfig = kubeletConfig;
            return this;
        }
        @CustomType.Setter
        public Builder labels(@Nullable Map<String,String> labels) {
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder linuxNodeConfig(@Nullable ClusterNodeConfigLinuxNodeConfig linuxNodeConfig) {
            this.linuxNodeConfig = linuxNodeConfig;
            return this;
        }
        @CustomType.Setter
        public Builder localSsdCount(@Nullable Integer localSsdCount) {
            this.localSsdCount = localSsdCount;
            return this;
        }
        @CustomType.Setter
        public Builder loggingVariant(@Nullable String loggingVariant) {
            this.loggingVariant = loggingVariant;
            return this;
        }
        @CustomType.Setter
        public Builder machineType(@Nullable String machineType) {
            this.machineType = machineType;
            return this;
        }
        @CustomType.Setter
        public Builder metadata(@Nullable Map<String,String> metadata) {
            this.metadata = metadata;
            return this;
        }
        @CustomType.Setter
        public Builder minCpuPlatform(@Nullable String minCpuPlatform) {
            this.minCpuPlatform = minCpuPlatform;
            return this;
        }
        @CustomType.Setter
        public Builder nodeGroup(@Nullable String nodeGroup) {
            this.nodeGroup = nodeGroup;
            return this;
        }
        @CustomType.Setter
        public Builder oauthScopes(@Nullable List<String> oauthScopes) {
            this.oauthScopes = oauthScopes;
            return this;
        }
        public Builder oauthScopes(String... oauthScopes) {
            return oauthScopes(List.of(oauthScopes));
        }
        @CustomType.Setter
        public Builder preemptible(@Nullable Boolean preemptible) {
            this.preemptible = preemptible;
            return this;
        }
        @CustomType.Setter
        public Builder reservationAffinity(@Nullable ClusterNodeConfigReservationAffinity reservationAffinity) {
            this.reservationAffinity = reservationAffinity;
            return this;
        }
        @CustomType.Setter
        public Builder resourceLabels(@Nullable Map<String,String> resourceLabels) {
            this.resourceLabels = resourceLabels;
            return this;
        }
        @CustomType.Setter
        public Builder sandboxConfig(@Nullable ClusterNodeConfigSandboxConfig sandboxConfig) {
            this.sandboxConfig = sandboxConfig;
            return this;
        }
        @CustomType.Setter
        public Builder serviceAccount(@Nullable String serviceAccount) {
            this.serviceAccount = serviceAccount;
            return this;
        }
        @CustomType.Setter
        public Builder shieldedInstanceConfig(@Nullable ClusterNodeConfigShieldedInstanceConfig shieldedInstanceConfig) {
            this.shieldedInstanceConfig = shieldedInstanceConfig;
            return this;
        }
        @CustomType.Setter
        public Builder spot(@Nullable Boolean spot) {
            this.spot = spot;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable List<String> tags) {
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder taints(@Nullable List<ClusterNodeConfigTaint> taints) {
            this.taints = taints;
            return this;
        }
        public Builder taints(ClusterNodeConfigTaint... taints) {
            return taints(List.of(taints));
        }
        @CustomType.Setter
        public Builder workloadMetadataConfig(@Nullable ClusterNodeConfigWorkloadMetadataConfig workloadMetadataConfig) {
            this.workloadMetadataConfig = workloadMetadataConfig;
            return this;
        }
        public ClusterNodeConfig build() {
            final var o = new ClusterNodeConfig();
            o.bootDiskKmsKey = bootDiskKmsKey;
            o.diskSizeGb = diskSizeGb;
            o.diskType = diskType;
            o.ephemeralStorageConfig = ephemeralStorageConfig;
            o.gcfsConfig = gcfsConfig;
            o.guestAccelerators = guestAccelerators;
            o.gvnic = gvnic;
            o.imageType = imageType;
            o.kubeletConfig = kubeletConfig;
            o.labels = labels;
            o.linuxNodeConfig = linuxNodeConfig;
            o.localSsdCount = localSsdCount;
            o.loggingVariant = loggingVariant;
            o.machineType = machineType;
            o.metadata = metadata;
            o.minCpuPlatform = minCpuPlatform;
            o.nodeGroup = nodeGroup;
            o.oauthScopes = oauthScopes;
            o.preemptible = preemptible;
            o.reservationAffinity = reservationAffinity;
            o.resourceLabels = resourceLabels;
            o.sandboxConfig = sandboxConfig;
            o.serviceAccount = serviceAccount;
            o.shieldedInstanceConfig = shieldedInstanceConfig;
            o.spot = spot;
            o.tags = tags;
            o.taints = taints;
            o.workloadMetadataConfig = workloadMetadataConfig;
            return o;
        }
    }
}
