// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigCloudrunConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigConfigConnectorConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigDnsCacheConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGcpFilestoreCsiDriverConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigGkeBackupAgentConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigHorizontalPodAutoscalingArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigHttpLoadBalancingArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigIstioConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigKalmConfigArgs;
import com.pulumi.gcp.container.inputs.ClusterAddonsConfigNetworkPolicyConfigArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterAddonsConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterAddonsConfigArgs Empty = new ClusterAddonsConfigArgs();

    /**
     * . Structure is documented below.
     * 
     */
    @Import(name="cloudrunConfig")
    private @Nullable Output<ClusterAddonsConfigCloudrunConfigArgs> cloudrunConfig;

    /**
     * @return . Structure is documented below.
     * 
     */
    public Optional<Output<ClusterAddonsConfigCloudrunConfigArgs>> cloudrunConfig() {
        return Optional.ofNullable(this.cloudrunConfig);
    }

    /**
     * .
     * The status of the ConfigConnector addon. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    @Import(name="configConnectorConfig")
    private @Nullable Output<ClusterAddonsConfigConfigConnectorConfigArgs> configConnectorConfig;

    /**
     * @return .
     * The status of the ConfigConnector addon. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigConfigConnectorConfigArgs>> configConnectorConfig() {
        return Optional.ofNullable(this.configConnectorConfig);
    }

    /**
     * .
     * The status of the NodeLocal DNSCache addon. It is disabled by default.
     * Set `enabled = true` to enable.
     * 
     */
    @Import(name="dnsCacheConfig")
    private @Nullable Output<ClusterAddonsConfigDnsCacheConfigArgs> dnsCacheConfig;

    /**
     * @return .
     * The status of the NodeLocal DNSCache addon. It is disabled by default.
     * Set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigDnsCacheConfigArgs>> dnsCacheConfig() {
        return Optional.ofNullable(this.dnsCacheConfig);
    }

    /**
     * .
     * Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set `enabled = true` to enable.
     * 
     */
    @Import(name="gcePersistentDiskCsiDriverConfig")
    private @Nullable Output<ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArgs> gcePersistentDiskCsiDriverConfig;

    /**
     * @return .
     * Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArgs>> gcePersistentDiskCsiDriverConfig() {
        return Optional.ofNullable(this.gcePersistentDiskCsiDriverConfig);
    }

    /**
     * The status of the Filestore CSI driver addon,
     * which allows the usage of filestore instance as volumes.
     * It is disabled by default; set `enabled = true` to enable.
     * 
     */
    @Import(name="gcpFilestoreCsiDriverConfig")
    private @Nullable Output<ClusterAddonsConfigGcpFilestoreCsiDriverConfigArgs> gcpFilestoreCsiDriverConfig;

    /**
     * @return The status of the Filestore CSI driver addon,
     * which allows the usage of filestore instance as volumes.
     * It is disabled by default; set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigGcpFilestoreCsiDriverConfigArgs>> gcpFilestoreCsiDriverConfig() {
        return Optional.ofNullable(this.gcpFilestoreCsiDriverConfig);
    }

    /**
     * ).
     * The status of the Backup for GKE agent addon. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    @Import(name="gkeBackupAgentConfig")
    private @Nullable Output<ClusterAddonsConfigGkeBackupAgentConfigArgs> gkeBackupAgentConfig;

    /**
     * @return ).
     * The status of the Backup for GKE agent addon. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigGkeBackupAgentConfigArgs>> gkeBackupAgentConfig() {
        return Optional.ofNullable(this.gkeBackupAgentConfig);
    }

    /**
     * The status of the Horizontal Pod Autoscaling
     * addon, which increases or decreases the number of replica pods a replication controller
     * has based on the resource usage of the existing pods.
     * It is enabled by default;
     * set `disabled = true` to disable.
     * 
     */
    @Import(name="horizontalPodAutoscaling")
    private @Nullable Output<ClusterAddonsConfigHorizontalPodAutoscalingArgs> horizontalPodAutoscaling;

    /**
     * @return The status of the Horizontal Pod Autoscaling
     * addon, which increases or decreases the number of replica pods a replication controller
     * has based on the resource usage of the existing pods.
     * It is enabled by default;
     * set `disabled = true` to disable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigHorizontalPodAutoscalingArgs>> horizontalPodAutoscaling() {
        return Optional.ofNullable(this.horizontalPodAutoscaling);
    }

    /**
     * The status of the HTTP (L7) load balancing
     * controller addon, which makes it easy to set up HTTP load balancers for services in a
     * cluster. It is enabled by default; set `disabled = true` to disable.
     * 
     */
    @Import(name="httpLoadBalancing")
    private @Nullable Output<ClusterAddonsConfigHttpLoadBalancingArgs> httpLoadBalancing;

    /**
     * @return The status of the HTTP (L7) load balancing
     * controller addon, which makes it easy to set up HTTP load balancers for services in a
     * cluster. It is enabled by default; set `disabled = true` to disable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigHttpLoadBalancingArgs>> httpLoadBalancing() {
        return Optional.ofNullable(this.httpLoadBalancing);
    }

    /**
     * .
     * Structure is documented below.
     * 
     */
    @Import(name="istioConfig")
    private @Nullable Output<ClusterAddonsConfigIstioConfigArgs> istioConfig;

    /**
     * @return .
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterAddonsConfigIstioConfigArgs>> istioConfig() {
        return Optional.ofNullable(this.istioConfig);
    }

    /**
     * .
     * Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    @Import(name="kalmConfig")
    private @Nullable Output<ClusterAddonsConfigKalmConfigArgs> kalmConfig;

    /**
     * @return .
     * Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set `enabled = true` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigKalmConfigArgs>> kalmConfig() {
        return Optional.ofNullable(this.kalmConfig);
    }

    /**
     * Whether we should enable the network policy addon
     * for the master.  This must be enabled in order to enable network policy for the nodes.
     * To enable this, you must also define a `network_policy` block,
     * otherwise nothing will happen.
     * It can only be disabled if the nodes already do not have network policies enabled.
     * Defaults to disabled; set `disabled = false` to enable.
     * 
     */
    @Import(name="networkPolicyConfig")
    private @Nullable Output<ClusterAddonsConfigNetworkPolicyConfigArgs> networkPolicyConfig;

    /**
     * @return Whether we should enable the network policy addon
     * for the master.  This must be enabled in order to enable network policy for the nodes.
     * To enable this, you must also define a `network_policy` block,
     * otherwise nothing will happen.
     * It can only be disabled if the nodes already do not have network policies enabled.
     * Defaults to disabled; set `disabled = false` to enable.
     * 
     */
    public Optional<Output<ClusterAddonsConfigNetworkPolicyConfigArgs>> networkPolicyConfig() {
        return Optional.ofNullable(this.networkPolicyConfig);
    }

    private ClusterAddonsConfigArgs() {}

    private ClusterAddonsConfigArgs(ClusterAddonsConfigArgs $) {
        this.cloudrunConfig = $.cloudrunConfig;
        this.configConnectorConfig = $.configConnectorConfig;
        this.dnsCacheConfig = $.dnsCacheConfig;
        this.gcePersistentDiskCsiDriverConfig = $.gcePersistentDiskCsiDriverConfig;
        this.gcpFilestoreCsiDriverConfig = $.gcpFilestoreCsiDriverConfig;
        this.gkeBackupAgentConfig = $.gkeBackupAgentConfig;
        this.horizontalPodAutoscaling = $.horizontalPodAutoscaling;
        this.httpLoadBalancing = $.httpLoadBalancing;
        this.istioConfig = $.istioConfig;
        this.kalmConfig = $.kalmConfig;
        this.networkPolicyConfig = $.networkPolicyConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterAddonsConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterAddonsConfigArgs $;

        public Builder() {
            $ = new ClusterAddonsConfigArgs();
        }

        public Builder(ClusterAddonsConfigArgs defaults) {
            $ = new ClusterAddonsConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cloudrunConfig . Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudrunConfig(@Nullable Output<ClusterAddonsConfigCloudrunConfigArgs> cloudrunConfig) {
            $.cloudrunConfig = cloudrunConfig;
            return this;
        }

        /**
         * @param cloudrunConfig . Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudrunConfig(ClusterAddonsConfigCloudrunConfigArgs cloudrunConfig) {
            return cloudrunConfig(Output.of(cloudrunConfig));
        }

        /**
         * @param configConnectorConfig .
         * The status of the ConfigConnector addon. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder configConnectorConfig(@Nullable Output<ClusterAddonsConfigConfigConnectorConfigArgs> configConnectorConfig) {
            $.configConnectorConfig = configConnectorConfig;
            return this;
        }

        /**
         * @param configConnectorConfig .
         * The status of the ConfigConnector addon. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder configConnectorConfig(ClusterAddonsConfigConfigConnectorConfigArgs configConnectorConfig) {
            return configConnectorConfig(Output.of(configConnectorConfig));
        }

        /**
         * @param dnsCacheConfig .
         * The status of the NodeLocal DNSCache addon. It is disabled by default.
         * Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder dnsCacheConfig(@Nullable Output<ClusterAddonsConfigDnsCacheConfigArgs> dnsCacheConfig) {
            $.dnsCacheConfig = dnsCacheConfig;
            return this;
        }

        /**
         * @param dnsCacheConfig .
         * The status of the NodeLocal DNSCache addon. It is disabled by default.
         * Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder dnsCacheConfig(ClusterAddonsConfigDnsCacheConfigArgs dnsCacheConfig) {
            return dnsCacheConfig(Output.of(dnsCacheConfig));
        }

        /**
         * @param gcePersistentDiskCsiDriverConfig .
         * Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gcePersistentDiskCsiDriverConfig(@Nullable Output<ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArgs> gcePersistentDiskCsiDriverConfig) {
            $.gcePersistentDiskCsiDriverConfig = gcePersistentDiskCsiDriverConfig;
            return this;
        }

        /**
         * @param gcePersistentDiskCsiDriverConfig .
         * Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver. Defaults to disabled; set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gcePersistentDiskCsiDriverConfig(ClusterAddonsConfigGcePersistentDiskCsiDriverConfigArgs gcePersistentDiskCsiDriverConfig) {
            return gcePersistentDiskCsiDriverConfig(Output.of(gcePersistentDiskCsiDriverConfig));
        }

        /**
         * @param gcpFilestoreCsiDriverConfig The status of the Filestore CSI driver addon,
         * which allows the usage of filestore instance as volumes.
         * It is disabled by default; set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gcpFilestoreCsiDriverConfig(@Nullable Output<ClusterAddonsConfigGcpFilestoreCsiDriverConfigArgs> gcpFilestoreCsiDriverConfig) {
            $.gcpFilestoreCsiDriverConfig = gcpFilestoreCsiDriverConfig;
            return this;
        }

        /**
         * @param gcpFilestoreCsiDriverConfig The status of the Filestore CSI driver addon,
         * which allows the usage of filestore instance as volumes.
         * It is disabled by default; set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gcpFilestoreCsiDriverConfig(ClusterAddonsConfigGcpFilestoreCsiDriverConfigArgs gcpFilestoreCsiDriverConfig) {
            return gcpFilestoreCsiDriverConfig(Output.of(gcpFilestoreCsiDriverConfig));
        }

        /**
         * @param gkeBackupAgentConfig ).
         * The status of the Backup for GKE agent addon. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gkeBackupAgentConfig(@Nullable Output<ClusterAddonsConfigGkeBackupAgentConfigArgs> gkeBackupAgentConfig) {
            $.gkeBackupAgentConfig = gkeBackupAgentConfig;
            return this;
        }

        /**
         * @param gkeBackupAgentConfig ).
         * The status of the Backup for GKE agent addon. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder gkeBackupAgentConfig(ClusterAddonsConfigGkeBackupAgentConfigArgs gkeBackupAgentConfig) {
            return gkeBackupAgentConfig(Output.of(gkeBackupAgentConfig));
        }

        /**
         * @param horizontalPodAutoscaling The status of the Horizontal Pod Autoscaling
         * addon, which increases or decreases the number of replica pods a replication controller
         * has based on the resource usage of the existing pods.
         * It is enabled by default;
         * set `disabled = true` to disable.
         * 
         * @return builder
         * 
         */
        public Builder horizontalPodAutoscaling(@Nullable Output<ClusterAddonsConfigHorizontalPodAutoscalingArgs> horizontalPodAutoscaling) {
            $.horizontalPodAutoscaling = horizontalPodAutoscaling;
            return this;
        }

        /**
         * @param horizontalPodAutoscaling The status of the Horizontal Pod Autoscaling
         * addon, which increases or decreases the number of replica pods a replication controller
         * has based on the resource usage of the existing pods.
         * It is enabled by default;
         * set `disabled = true` to disable.
         * 
         * @return builder
         * 
         */
        public Builder horizontalPodAutoscaling(ClusterAddonsConfigHorizontalPodAutoscalingArgs horizontalPodAutoscaling) {
            return horizontalPodAutoscaling(Output.of(horizontalPodAutoscaling));
        }

        /**
         * @param httpLoadBalancing The status of the HTTP (L7) load balancing
         * controller addon, which makes it easy to set up HTTP load balancers for services in a
         * cluster. It is enabled by default; set `disabled = true` to disable.
         * 
         * @return builder
         * 
         */
        public Builder httpLoadBalancing(@Nullable Output<ClusterAddonsConfigHttpLoadBalancingArgs> httpLoadBalancing) {
            $.httpLoadBalancing = httpLoadBalancing;
            return this;
        }

        /**
         * @param httpLoadBalancing The status of the HTTP (L7) load balancing
         * controller addon, which makes it easy to set up HTTP load balancers for services in a
         * cluster. It is enabled by default; set `disabled = true` to disable.
         * 
         * @return builder
         * 
         */
        public Builder httpLoadBalancing(ClusterAddonsConfigHttpLoadBalancingArgs httpLoadBalancing) {
            return httpLoadBalancing(Output.of(httpLoadBalancing));
        }

        /**
         * @param istioConfig .
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder istioConfig(@Nullable Output<ClusterAddonsConfigIstioConfigArgs> istioConfig) {
            $.istioConfig = istioConfig;
            return this;
        }

        /**
         * @param istioConfig .
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder istioConfig(ClusterAddonsConfigIstioConfigArgs istioConfig) {
            return istioConfig(Output.of(istioConfig));
        }

        /**
         * @param kalmConfig .
         * Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder kalmConfig(@Nullable Output<ClusterAddonsConfigKalmConfigArgs> kalmConfig) {
            $.kalmConfig = kalmConfig;
            return this;
        }

        /**
         * @param kalmConfig .
         * Configuration for the KALM addon, which manages the lifecycle of k8s. It is disabled by default; Set `enabled = true` to enable.
         * 
         * @return builder
         * 
         */
        public Builder kalmConfig(ClusterAddonsConfigKalmConfigArgs kalmConfig) {
            return kalmConfig(Output.of(kalmConfig));
        }

        /**
         * @param networkPolicyConfig Whether we should enable the network policy addon
         * for the master.  This must be enabled in order to enable network policy for the nodes.
         * To enable this, you must also define a `network_policy` block,
         * otherwise nothing will happen.
         * It can only be disabled if the nodes already do not have network policies enabled.
         * Defaults to disabled; set `disabled = false` to enable.
         * 
         * @return builder
         * 
         */
        public Builder networkPolicyConfig(@Nullable Output<ClusterAddonsConfigNetworkPolicyConfigArgs> networkPolicyConfig) {
            $.networkPolicyConfig = networkPolicyConfig;
            return this;
        }

        /**
         * @param networkPolicyConfig Whether we should enable the network policy addon
         * for the master.  This must be enabled in order to enable network policy for the nodes.
         * To enable this, you must also define a `network_policy` block,
         * otherwise nothing will happen.
         * It can only be disabled if the nodes already do not have network policies enabled.
         * Defaults to disabled; set `disabled = false` to enable.
         * 
         * @return builder
         * 
         */
        public Builder networkPolicyConfig(ClusterAddonsConfigNetworkPolicyConfigArgs networkPolicyConfig) {
            return networkPolicyConfig(Output.of(networkPolicyConfig));
        }

        public ClusterAddonsConfigArgs build() {
            return $;
        }
    }

}
