// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs;
import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigGceClusterConfigReservationAffinityArgs;
import com.pulumi.gcp.dataproc.inputs.ClusterClusterConfigGceClusterConfigShieldedInstanceConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterClusterConfigGceClusterConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterClusterConfigGceClusterConfigArgs Empty = new ClusterClusterConfigGceClusterConfigArgs();

    /**
     * By default, clusters are not restricted to internal IP addresses,
     * and will have ephemeral external IP addresses assigned to each instance. If set to true, all
     * instances in the cluster will only have internal IP addresses. Note: Private Google Access
     * (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
     * will be launched in.
     * 
     */
    @Import(name="internalIpOnly")
    private @Nullable Output<Boolean> internalIpOnly;

    /**
     * @return By default, clusters are not restricted to internal IP addresses,
     * and will have ephemeral external IP addresses assigned to each instance. If set to true, all
     * instances in the cluster will only have internal IP addresses. Note: Private Google Access
     * (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
     * will be launched in.
     * 
     */
    public Optional<Output<Boolean>> internalIpOnly() {
        return Optional.ofNullable(this.internalIpOnly);
    }

    /**
     * A map of the Compute Engine metadata entries to add to all instances
     * (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return A map of the Compute Engine metadata entries to add to all instances
     * (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * The name or self_link of the Google Compute Engine
     * network to the cluster will be part of. Conflicts with `subnetwork`.
     * If neither is specified, this defaults to the &#34;default&#34; network.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The name or self_link of the Google Compute Engine
     * network to the cluster will be part of. Conflicts with `subnetwork`.
     * If neither is specified, this defaults to the &#34;default&#34; network.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * Node Group Affinity for sole-tenant clusters.
     * 
     */
    @Import(name="nodeGroupAffinity")
    private @Nullable Output<ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs> nodeGroupAffinity;

    /**
     * @return Node Group Affinity for sole-tenant clusters.
     * 
     */
    public Optional<Output<ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs>> nodeGroupAffinity() {
        return Optional.ofNullable(this.nodeGroupAffinity);
    }

    /**
     * Reservation Affinity for consuming zonal reservation.
     * 
     */
    @Import(name="reservationAffinity")
    private @Nullable Output<ClusterClusterConfigGceClusterConfigReservationAffinityArgs> reservationAffinity;

    /**
     * @return Reservation Affinity for consuming zonal reservation.
     * 
     */
    public Optional<Output<ClusterClusterConfigGceClusterConfigReservationAffinityArgs>> reservationAffinity() {
        return Optional.ofNullable(this.reservationAffinity);
    }

    /**
     * The service account to be used by the Node VMs.
     * If not specified, the &#34;default&#34; service account is used.
     * 
     */
    @Import(name="serviceAccount")
    private @Nullable Output<String> serviceAccount;

    /**
     * @return The service account to be used by the Node VMs.
     * If not specified, the &#34;default&#34; service account is used.
     * 
     */
    public Optional<Output<String>> serviceAccount() {
        return Optional.ofNullable(this.serviceAccount);
    }

    /**
     * The set of Google API scopes
     * to be made available on all of the node VMs under the `service_account`
     * specified. Both OAuth2 URLs and gcloud
     * short names are supported. To allow full access to all Cloud APIs, use the
     * `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
     * 
     */
    @Import(name="serviceAccountScopes")
    private @Nullable Output<List<String>> serviceAccountScopes;

    /**
     * @return The set of Google API scopes
     * to be made available on all of the node VMs under the `service_account`
     * specified. Both OAuth2 URLs and gcloud
     * short names are supported. To allow full access to all Cloud APIs, use the
     * `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
     * 
     */
    public Optional<Output<List<String>>> serviceAccountScopes() {
        return Optional.ofNullable(this.serviceAccountScopes);
    }

    /**
     * Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
     * 
     */
    @Import(name="shieldedInstanceConfig")
    private @Nullable Output<ClusterClusterConfigGceClusterConfigShieldedInstanceConfigArgs> shieldedInstanceConfig;

    /**
     * @return Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
     * 
     */
    public Optional<Output<ClusterClusterConfigGceClusterConfigShieldedInstanceConfigArgs>> shieldedInstanceConfig() {
        return Optional.ofNullable(this.shieldedInstanceConfig);
    }

    /**
     * The name or self_link of the Google Compute Engine
     * subnetwork the cluster will be part of. Conflicts with `network`.
     * 
     */
    @Import(name="subnetwork")
    private @Nullable Output<String> subnetwork;

    /**
     * @return The name or self_link of the Google Compute Engine
     * subnetwork the cluster will be part of. Conflicts with `network`.
     * 
     */
    public Optional<Output<String>> subnetwork() {
        return Optional.ofNullable(this.subnetwork);
    }

    /**
     * The list of instance tags applied to instances in the cluster.
     * Tags are used to identify valid sources or targets for network firewalls.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The list of instance tags applied to instances in the cluster.
     * Tags are used to identify valid sources or targets for network firewalls.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The GCP zone where your data is stored and used (i.e. where
     * the master and the worker nodes will be created in). If `region` is set to &#39;global&#39; (default)
     * then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
     * to determine this automatically for you.
     * Note: This setting additionally determines and restricts
     * which computing resources are available for use with other configs such as
     * `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return The GCP zone where your data is stored and used (i.e. where
     * the master and the worker nodes will be created in). If `region` is set to &#39;global&#39; (default)
     * then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
     * to determine this automatically for you.
     * Note: This setting additionally determines and restricts
     * which computing resources are available for use with other configs such as
     * `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private ClusterClusterConfigGceClusterConfigArgs() {}

    private ClusterClusterConfigGceClusterConfigArgs(ClusterClusterConfigGceClusterConfigArgs $) {
        this.internalIpOnly = $.internalIpOnly;
        this.metadata = $.metadata;
        this.network = $.network;
        this.nodeGroupAffinity = $.nodeGroupAffinity;
        this.reservationAffinity = $.reservationAffinity;
        this.serviceAccount = $.serviceAccount;
        this.serviceAccountScopes = $.serviceAccountScopes;
        this.shieldedInstanceConfig = $.shieldedInstanceConfig;
        this.subnetwork = $.subnetwork;
        this.tags = $.tags;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterClusterConfigGceClusterConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterClusterConfigGceClusterConfigArgs $;

        public Builder() {
            $ = new ClusterClusterConfigGceClusterConfigArgs();
        }

        public Builder(ClusterClusterConfigGceClusterConfigArgs defaults) {
            $ = new ClusterClusterConfigGceClusterConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param internalIpOnly By default, clusters are not restricted to internal IP addresses,
         * and will have ephemeral external IP addresses assigned to each instance. If set to true, all
         * instances in the cluster will only have internal IP addresses. Note: Private Google Access
         * (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
         * will be launched in.
         * 
         * @return builder
         * 
         */
        public Builder internalIpOnly(@Nullable Output<Boolean> internalIpOnly) {
            $.internalIpOnly = internalIpOnly;
            return this;
        }

        /**
         * @param internalIpOnly By default, clusters are not restricted to internal IP addresses,
         * and will have ephemeral external IP addresses assigned to each instance. If set to true, all
         * instances in the cluster will only have internal IP addresses. Note: Private Google Access
         * (also known as `privateIpGoogleAccess`) must be enabled on the subnetwork that the cluster
         * will be launched in.
         * 
         * @return builder
         * 
         */
        public Builder internalIpOnly(Boolean internalIpOnly) {
            return internalIpOnly(Output.of(internalIpOnly));
        }

        /**
         * @param metadata A map of the Compute Engine metadata entries to add to all instances
         * (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata A map of the Compute Engine metadata entries to add to all instances
         * (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param network The name or self_link of the Google Compute Engine
         * network to the cluster will be part of. Conflicts with `subnetwork`.
         * If neither is specified, this defaults to the &#34;default&#34; network.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The name or self_link of the Google Compute Engine
         * network to the cluster will be part of. Conflicts with `subnetwork`.
         * If neither is specified, this defaults to the &#34;default&#34; network.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param nodeGroupAffinity Node Group Affinity for sole-tenant clusters.
         * 
         * @return builder
         * 
         */
        public Builder nodeGroupAffinity(@Nullable Output<ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs> nodeGroupAffinity) {
            $.nodeGroupAffinity = nodeGroupAffinity;
            return this;
        }

        /**
         * @param nodeGroupAffinity Node Group Affinity for sole-tenant clusters.
         * 
         * @return builder
         * 
         */
        public Builder nodeGroupAffinity(ClusterClusterConfigGceClusterConfigNodeGroupAffinityArgs nodeGroupAffinity) {
            return nodeGroupAffinity(Output.of(nodeGroupAffinity));
        }

        /**
         * @param reservationAffinity Reservation Affinity for consuming zonal reservation.
         * 
         * @return builder
         * 
         */
        public Builder reservationAffinity(@Nullable Output<ClusterClusterConfigGceClusterConfigReservationAffinityArgs> reservationAffinity) {
            $.reservationAffinity = reservationAffinity;
            return this;
        }

        /**
         * @param reservationAffinity Reservation Affinity for consuming zonal reservation.
         * 
         * @return builder
         * 
         */
        public Builder reservationAffinity(ClusterClusterConfigGceClusterConfigReservationAffinityArgs reservationAffinity) {
            return reservationAffinity(Output.of(reservationAffinity));
        }

        /**
         * @param serviceAccount The service account to be used by the Node VMs.
         * If not specified, the &#34;default&#34; service account is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(@Nullable Output<String> serviceAccount) {
            $.serviceAccount = serviceAccount;
            return this;
        }

        /**
         * @param serviceAccount The service account to be used by the Node VMs.
         * If not specified, the &#34;default&#34; service account is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceAccount(String serviceAccount) {
            return serviceAccount(Output.of(serviceAccount));
        }

        /**
         * @param serviceAccountScopes The set of Google API scopes
         * to be made available on all of the node VMs under the `service_account`
         * specified. Both OAuth2 URLs and gcloud
         * short names are supported. To allow full access to all Cloud APIs, use the
         * `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountScopes(@Nullable Output<List<String>> serviceAccountScopes) {
            $.serviceAccountScopes = serviceAccountScopes;
            return this;
        }

        /**
         * @param serviceAccountScopes The set of Google API scopes
         * to be made available on all of the node VMs under the `service_account`
         * specified. Both OAuth2 URLs and gcloud
         * short names are supported. To allow full access to all Cloud APIs, use the
         * `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountScopes(List<String> serviceAccountScopes) {
            return serviceAccountScopes(Output.of(serviceAccountScopes));
        }

        /**
         * @param serviceAccountScopes The set of Google API scopes
         * to be made available on all of the node VMs under the `service_account`
         * specified. Both OAuth2 URLs and gcloud
         * short names are supported. To allow full access to all Cloud APIs, use the
         * `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
         * 
         * @return builder
         * 
         */
        public Builder serviceAccountScopes(String... serviceAccountScopes) {
            return serviceAccountScopes(List.of(serviceAccountScopes));
        }

        /**
         * @param shieldedInstanceConfig Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
         * 
         * @return builder
         * 
         */
        public Builder shieldedInstanceConfig(@Nullable Output<ClusterClusterConfigGceClusterConfigShieldedInstanceConfigArgs> shieldedInstanceConfig) {
            $.shieldedInstanceConfig = shieldedInstanceConfig;
            return this;
        }

        /**
         * @param shieldedInstanceConfig Shielded Instance Config for clusters using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
         * 
         * @return builder
         * 
         */
        public Builder shieldedInstanceConfig(ClusterClusterConfigGceClusterConfigShieldedInstanceConfigArgs shieldedInstanceConfig) {
            return shieldedInstanceConfig(Output.of(shieldedInstanceConfig));
        }

        /**
         * @param subnetwork The name or self_link of the Google Compute Engine
         * subnetwork the cluster will be part of. Conflicts with `network`.
         * 
         * @return builder
         * 
         */
        public Builder subnetwork(@Nullable Output<String> subnetwork) {
            $.subnetwork = subnetwork;
            return this;
        }

        /**
         * @param subnetwork The name or self_link of the Google Compute Engine
         * subnetwork the cluster will be part of. Conflicts with `network`.
         * 
         * @return builder
         * 
         */
        public Builder subnetwork(String subnetwork) {
            return subnetwork(Output.of(subnetwork));
        }

        /**
         * @param tags The list of instance tags applied to instances in the cluster.
         * Tags are used to identify valid sources or targets for network firewalls.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The list of instance tags applied to instances in the cluster.
         * Tags are used to identify valid sources or targets for network firewalls.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The list of instance tags applied to instances in the cluster.
         * Tags are used to identify valid sources or targets for network firewalls.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param zone The GCP zone where your data is stored and used (i.e. where
         * the master and the worker nodes will be created in). If `region` is set to &#39;global&#39; (default)
         * then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
         * to determine this automatically for you.
         * Note: This setting additionally determines and restricts
         * which computing resources are available for use with other configs such as
         * `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The GCP zone where your data is stored and used (i.e. where
         * the master and the worker nodes will be created in). If `region` is set to &#39;global&#39; (default)
         * then `zone` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone)
         * to determine this automatically for you.
         * Note: This setting additionally determines and restricts
         * which computing resources are available for use with other configs such as
         * `cluster_config.master_config.machine_type` and `cluster_config.worker_config.machine_type`.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public ClusterClusterConfigGceClusterConfigArgs build() {
            return $;
        }
    }

}
