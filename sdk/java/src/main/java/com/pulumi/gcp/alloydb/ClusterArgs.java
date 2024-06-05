// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.alloydb.inputs.ClusterAutomatedBackupPolicyArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterContinuousBackupConfigArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterEncryptionConfigArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterMaintenanceUpdatePolicyArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterNetworkConfigArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterPscConfigArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterRestoreBackupSourceArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterRestoreContinuousBackupSourceArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterSecondaryConfigArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterArgs Empty = new ClusterArgs();

    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    /**
     * @return Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
     * Structure is documented below.
     * 
     */
    @Import(name="automatedBackupPolicy")
    private @Nullable Output<ClusterAutomatedBackupPolicyArgs> automatedBackupPolicy;

    /**
     * @return The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterAutomatedBackupPolicyArgs>> automatedBackupPolicy() {
        return Optional.ofNullable(this.automatedBackupPolicy);
    }

    /**
     * The ID of the alloydb cluster.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return The ID of the alloydb cluster.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * The type of cluster. If not set, defaults to PRIMARY.
     * Default value is `PRIMARY`.
     * Possible values are: `PRIMARY`, `SECONDARY`.
     * 
     */
    @Import(name="clusterType")
    private @Nullable Output<String> clusterType;

    /**
     * @return The type of cluster. If not set, defaults to PRIMARY.
     * Default value is `PRIMARY`.
     * Possible values are: `PRIMARY`, `SECONDARY`.
     * 
     */
    public Optional<Output<String>> clusterType() {
        return Optional.ofNullable(this.clusterType);
    }

    /**
     * The continuous backup config for this cluster.
     * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
     * Structure is documented below.
     * 
     */
    @Import(name="continuousBackupConfig")
    private @Nullable Output<ClusterContinuousBackupConfigArgs> continuousBackupConfig;

    /**
     * @return The continuous backup config for this cluster.
     * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterContinuousBackupConfigArgs>> continuousBackupConfig() {
        return Optional.ofNullable(this.continuousBackupConfig);
    }

    /**
     * The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
     * 
     */
    @Import(name="databaseVersion")
    private @Nullable Output<String> databaseVersion;

    /**
     * @return The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
     * 
     */
    public Optional<Output<String>> databaseVersion() {
        return Optional.ofNullable(this.databaseVersion);
    }

    /**
     * Policy to determine if the cluster should be deleted forcefully.
     * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
     * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
     * 
     */
    @Import(name="deletionPolicy")
    private @Nullable Output<String> deletionPolicy;

    /**
     * @return Policy to determine if the cluster should be deleted forcefully.
     * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
     * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
     * 
     */
    public Optional<Output<String>> deletionPolicy() {
        return Optional.ofNullable(this.deletionPolicy);
    }

    /**
     * User-settable and human-readable display name for the Cluster.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return User-settable and human-readable display name for the Cluster.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     * 
     */
    @Import(name="encryptionConfig")
    private @Nullable Output<ClusterEncryptionConfigArgs> encryptionConfig;

    /**
     * @return EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterEncryptionConfigArgs>> encryptionConfig() {
        return Optional.ofNullable(this.encryptionConfig);
    }

    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * Initial user to setup during cluster creation.
     * Structure is documented below.
     * 
     */
    @Import(name="initialUser")
    private @Nullable Output<ClusterInitialUserArgs> initialUser;

    /**
     * @return Initial user to setup during cluster creation.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterInitialUserArgs>> initialUser() {
        return Optional.ofNullable(this.initialUser);
    }

    /**
     * User-defined labels for the alloydb cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return User-defined labels for the alloydb cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location where the alloydb cluster should reside.
     * 
     * ***
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return The location where the alloydb cluster should reside.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * MaintenanceUpdatePolicy defines the policy for system updates.
     * Structure is documented below.
     * 
     */
    @Import(name="maintenanceUpdatePolicy")
    private @Nullable Output<ClusterMaintenanceUpdatePolicyArgs> maintenanceUpdatePolicy;

    /**
     * @return MaintenanceUpdatePolicy defines the policy for system updates.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterMaintenanceUpdatePolicyArgs>> maintenanceUpdatePolicy() {
        return Optional.ofNullable(this.maintenanceUpdatePolicy);
    }

    /**
     * (Optional, Deprecated)
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     * @deprecated
     * `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     */
    @Deprecated /* `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration. */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return (Optional, Deprecated)
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     * @deprecated
     * `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     */
    @Deprecated /* `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration. */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * Metadata related to network configuration.
     * Structure is documented below.
     * 
     */
    @Import(name="networkConfig")
    private @Nullable Output<ClusterNetworkConfigArgs> networkConfig;

    /**
     * @return Metadata related to network configuration.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterNetworkConfigArgs>> networkConfig() {
        return Optional.ofNullable(this.networkConfig);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Configuration for Private Service Connect (PSC) for the cluster.
     * Structure is documented below.
     * 
     */
    @Import(name="pscConfig")
    private @Nullable Output<ClusterPscConfigArgs> pscConfig;

    /**
     * @return Configuration for Private Service Connect (PSC) for the cluster.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterPscConfigArgs>> pscConfig() {
        return Optional.ofNullable(this.pscConfig);
    }

    /**
     * The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    @Import(name="restoreBackupSource")
    private @Nullable Output<ClusterRestoreBackupSourceArgs> restoreBackupSource;

    /**
     * @return The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterRestoreBackupSourceArgs>> restoreBackupSource() {
        return Optional.ofNullable(this.restoreBackupSource);
    }

    /**
     * The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    @Import(name="restoreContinuousBackupSource")
    private @Nullable Output<ClusterRestoreContinuousBackupSourceArgs> restoreContinuousBackupSource;

    /**
     * @return The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterRestoreContinuousBackupSourceArgs>> restoreContinuousBackupSource() {
        return Optional.ofNullable(this.restoreContinuousBackupSource);
    }

    /**
     * Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
     * Structure is documented below.
     * 
     */
    @Import(name="secondaryConfig")
    private @Nullable Output<ClusterSecondaryConfigArgs> secondaryConfig;

    /**
     * @return Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClusterSecondaryConfigArgs>> secondaryConfig() {
        return Optional.ofNullable(this.secondaryConfig);
    }

    private ClusterArgs() {}

    private ClusterArgs(ClusterArgs $) {
        this.annotations = $.annotations;
        this.automatedBackupPolicy = $.automatedBackupPolicy;
        this.clusterId = $.clusterId;
        this.clusterType = $.clusterType;
        this.continuousBackupConfig = $.continuousBackupConfig;
        this.databaseVersion = $.databaseVersion;
        this.deletionPolicy = $.deletionPolicy;
        this.displayName = $.displayName;
        this.encryptionConfig = $.encryptionConfig;
        this.etag = $.etag;
        this.initialUser = $.initialUser;
        this.labels = $.labels;
        this.location = $.location;
        this.maintenanceUpdatePolicy = $.maintenanceUpdatePolicy;
        this.network = $.network;
        this.networkConfig = $.networkConfig;
        this.project = $.project;
        this.pscConfig = $.pscConfig;
        this.restoreBackupSource = $.restoreBackupSource;
        this.restoreContinuousBackupSource = $.restoreContinuousBackupSource;
        this.secondaryConfig = $.secondaryConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterArgs $;

        public Builder() {
            $ = new ClusterArgs();
        }

        public Builder(ClusterArgs defaults) {
            $ = new ClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param annotations Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
         * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
         * 
         * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
         * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        /**
         * @param annotations Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
         * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
         * 
         * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
         * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
        }

        /**
         * @param automatedBackupPolicy The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder automatedBackupPolicy(@Nullable Output<ClusterAutomatedBackupPolicyArgs> automatedBackupPolicy) {
            $.automatedBackupPolicy = automatedBackupPolicy;
            return this;
        }

        /**
         * @param automatedBackupPolicy The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder automatedBackupPolicy(ClusterAutomatedBackupPolicyArgs automatedBackupPolicy) {
            return automatedBackupPolicy(Output.of(automatedBackupPolicy));
        }

        /**
         * @param clusterId The ID of the alloydb cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The ID of the alloydb cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param clusterType The type of cluster. If not set, defaults to PRIMARY.
         * Default value is `PRIMARY`.
         * Possible values are: `PRIMARY`, `SECONDARY`.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(@Nullable Output<String> clusterType) {
            $.clusterType = clusterType;
            return this;
        }

        /**
         * @param clusterType The type of cluster. If not set, defaults to PRIMARY.
         * Default value is `PRIMARY`.
         * Possible values are: `PRIMARY`, `SECONDARY`.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(String clusterType) {
            return clusterType(Output.of(clusterType));
        }

        /**
         * @param continuousBackupConfig The continuous backup config for this cluster.
         * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder continuousBackupConfig(@Nullable Output<ClusterContinuousBackupConfigArgs> continuousBackupConfig) {
            $.continuousBackupConfig = continuousBackupConfig;
            return this;
        }

        /**
         * @param continuousBackupConfig The continuous backup config for this cluster.
         * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder continuousBackupConfig(ClusterContinuousBackupConfigArgs continuousBackupConfig) {
            return continuousBackupConfig(Output.of(continuousBackupConfig));
        }

        /**
         * @param databaseVersion The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
         * 
         * @return builder
         * 
         */
        public Builder databaseVersion(@Nullable Output<String> databaseVersion) {
            $.databaseVersion = databaseVersion;
            return this;
        }

        /**
         * @param databaseVersion The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
         * 
         * @return builder
         * 
         */
        public Builder databaseVersion(String databaseVersion) {
            return databaseVersion(Output.of(databaseVersion));
        }

        /**
         * @param deletionPolicy Policy to determine if the cluster should be deleted forcefully.
         * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
         * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
         * 
         * @return builder
         * 
         */
        public Builder deletionPolicy(@Nullable Output<String> deletionPolicy) {
            $.deletionPolicy = deletionPolicy;
            return this;
        }

        /**
         * @param deletionPolicy Policy to determine if the cluster should be deleted forcefully.
         * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
         * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
         * 
         * @return builder
         * 
         */
        public Builder deletionPolicy(String deletionPolicy) {
            return deletionPolicy(Output.of(deletionPolicy));
        }

        /**
         * @param displayName User-settable and human-readable display name for the Cluster.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName User-settable and human-readable display name for the Cluster.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param encryptionConfig EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfig(@Nullable Output<ClusterEncryptionConfigArgs> encryptionConfig) {
            $.encryptionConfig = encryptionConfig;
            return this;
        }

        /**
         * @param encryptionConfig EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfig(ClusterEncryptionConfigArgs encryptionConfig) {
            return encryptionConfig(Output.of(encryptionConfig));
        }

        /**
         * @param etag For Resource freshness validation (https://google.aip.dev/154)
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag For Resource freshness validation (https://google.aip.dev/154)
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param initialUser Initial user to setup during cluster creation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder initialUser(@Nullable Output<ClusterInitialUserArgs> initialUser) {
            $.initialUser = initialUser;
            return this;
        }

        /**
         * @param initialUser Initial user to setup during cluster creation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder initialUser(ClusterInitialUserArgs initialUser) {
            return initialUser(Output.of(initialUser));
        }

        /**
         * @param labels User-defined labels for the alloydb cluster.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined labels for the alloydb cluster.
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location where the alloydb cluster should reside.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location where the alloydb cluster should reside.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param maintenanceUpdatePolicy MaintenanceUpdatePolicy defines the policy for system updates.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceUpdatePolicy(@Nullable Output<ClusterMaintenanceUpdatePolicyArgs> maintenanceUpdatePolicy) {
            $.maintenanceUpdatePolicy = maintenanceUpdatePolicy;
            return this;
        }

        /**
         * @param maintenanceUpdatePolicy MaintenanceUpdatePolicy defines the policy for system updates.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceUpdatePolicy(ClusterMaintenanceUpdatePolicyArgs maintenanceUpdatePolicy) {
            return maintenanceUpdatePolicy(Output.of(maintenanceUpdatePolicy));
        }

        /**
         * @param network (Optional, Deprecated)
         * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
         * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
         * 
         * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
         * 
         * @return builder
         * 
         * @deprecated
         * `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
         * 
         */
        @Deprecated /* `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration. */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network (Optional, Deprecated)
         * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
         * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
         * 
         * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
         * 
         * @return builder
         * 
         * @deprecated
         * `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
         * 
         */
        @Deprecated /* `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration. */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param networkConfig Metadata related to network configuration.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder networkConfig(@Nullable Output<ClusterNetworkConfigArgs> networkConfig) {
            $.networkConfig = networkConfig;
            return this;
        }

        /**
         * @param networkConfig Metadata related to network configuration.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder networkConfig(ClusterNetworkConfigArgs networkConfig) {
            return networkConfig(Output.of(networkConfig));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param pscConfig Configuration for Private Service Connect (PSC) for the cluster.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConfig(@Nullable Output<ClusterPscConfigArgs> pscConfig) {
            $.pscConfig = pscConfig;
            return this;
        }

        /**
         * @param pscConfig Configuration for Private Service Connect (PSC) for the cluster.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder pscConfig(ClusterPscConfigArgs pscConfig) {
            return pscConfig(Output.of(pscConfig));
        }

        /**
         * @param restoreBackupSource The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restoreBackupSource(@Nullable Output<ClusterRestoreBackupSourceArgs> restoreBackupSource) {
            $.restoreBackupSource = restoreBackupSource;
            return this;
        }

        /**
         * @param restoreBackupSource The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restoreBackupSource(ClusterRestoreBackupSourceArgs restoreBackupSource) {
            return restoreBackupSource(Output.of(restoreBackupSource));
        }

        /**
         * @param restoreContinuousBackupSource The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restoreContinuousBackupSource(@Nullable Output<ClusterRestoreContinuousBackupSourceArgs> restoreContinuousBackupSource) {
            $.restoreContinuousBackupSource = restoreContinuousBackupSource;
            return this;
        }

        /**
         * @param restoreContinuousBackupSource The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder restoreContinuousBackupSource(ClusterRestoreContinuousBackupSourceArgs restoreContinuousBackupSource) {
            return restoreContinuousBackupSource(Output.of(restoreContinuousBackupSource));
        }

        /**
         * @param secondaryConfig Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaryConfig(@Nullable Output<ClusterSecondaryConfigArgs> secondaryConfig) {
            $.secondaryConfig = secondaryConfig;
            return this;
        }

        /**
         * @param secondaryConfig Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaryConfig(ClusterSecondaryConfigArgs secondaryConfig) {
            return secondaryConfig(Output.of(secondaryConfig));
        }

        public ClusterArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "clusterId");
            }
            if ($.location == null) {
                throw new MissingRequiredPropertyException("ClusterArgs", "location");
            }
            return $;
        }
    }

}
