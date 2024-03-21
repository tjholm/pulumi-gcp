// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.workstations;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigContainerArgs;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigEncryptionKeyArgs;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigEphemeralDirectoryArgs;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigHostArgs;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigPersistentDirectoryArgs;
import com.pulumi.gcp.workstations.inputs.WorkstationConfigReadinessCheckArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkstationConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkstationConfigArgs Empty = new WorkstationConfigArgs();

    /**
     * Client-specified annotations. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    /**
     * @return Client-specified annotations. This is distinct from labels.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * Container that will be run for each workstation using this configuration when that workstation is started.
     * Structure is documented below.
     * 
     */
    @Import(name="container")
    private @Nullable Output<WorkstationConfigContainerArgs> container;

    /**
     * @return Container that will be run for each workstation using this configuration when that workstation is started.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkstationConfigContainerArgs>> container() {
        return Optional.ofNullable(this.container);
    }

    /**
     * Disables support for plain TCP connections in the workstation. By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
     * 
     */
    @Import(name="disableTcpConnections")
    private @Nullable Output<Boolean> disableTcpConnections;

    /**
     * @return Disables support for plain TCP connections in the workstation. By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
     * 
     */
    public Optional<Output<Boolean>> disableTcpConnections() {
        return Optional.ofNullable(this.disableTcpConnections);
    }

    /**
     * Human-readable name for this resource.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human-readable name for this resource.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
     * 
     */
    @Import(name="enableAuditAgent")
    private @Nullable Output<Boolean> enableAuditAgent;

    /**
     * @return Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
     * 
     */
    public Optional<Output<Boolean>> enableAuditAgent() {
        return Optional.ofNullable(this.enableAuditAgent);
    }

    /**
     * Encrypts resources of this workstation configuration using a customer-managed encryption key.
     * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
     * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
     * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
     * Structure is documented below.
     * 
     */
    @Import(name="encryptionKey")
    private @Nullable Output<WorkstationConfigEncryptionKeyArgs> encryptionKey;

    /**
     * @return Encrypts resources of this workstation configuration using a customer-managed encryption key.
     * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
     * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
     * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkstationConfigEncryptionKeyArgs>> encryptionKey() {
        return Optional.ofNullable(this.encryptionKey);
    }

    /**
     * Ephemeral directories which won&#39;t persist across workstation sessions.
     * Structure is documented below.
     * 
     */
    @Import(name="ephemeralDirectories")
    private @Nullable Output<List<WorkstationConfigEphemeralDirectoryArgs>> ephemeralDirectories;

    /**
     * @return Ephemeral directories which won&#39;t persist across workstation sessions.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<WorkstationConfigEphemeralDirectoryArgs>>> ephemeralDirectories() {
        return Optional.ofNullable(this.ephemeralDirectories);
    }

    /**
     * Runtime host for a workstation.
     * Structure is documented below.
     * 
     */
    @Import(name="host")
    private @Nullable Output<WorkstationConfigHostArgs> host;

    /**
     * @return Runtime host for a workstation.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkstationConfigHostArgs>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * How long to wait before automatically stopping an instance that hasn&#39;t recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Import(name="idleTimeout")
    private @Nullable Output<String> idleTimeout;

    /**
     * @return How long to wait before automatically stopping an instance that hasn&#39;t recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<Output<String>> idleTimeout() {
        return Optional.ofNullable(this.idleTimeout);
    }

    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location where the workstation cluster config should reside.
     * 
     * ***
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return The location where the workstation cluster config should reside.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * Directories to persist across workstation sessions.
     * Structure is documented below.
     * 
     */
    @Import(name="persistentDirectories")
    private @Nullable Output<List<WorkstationConfigPersistentDirectoryArgs>> persistentDirectories;

    /**
     * @return Directories to persist across workstation sessions.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<WorkstationConfigPersistentDirectoryArgs>>> persistentDirectories() {
        return Optional.ofNullable(this.persistentDirectories);
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
     * Readiness checks to be performed on a workstation.
     * Structure is documented below.
     * 
     */
    @Import(name="readinessChecks")
    private @Nullable Output<List<WorkstationConfigReadinessCheckArgs>> readinessChecks;

    /**
     * @return Readiness checks to be performed on a workstation.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<WorkstationConfigReadinessCheckArgs>>> readinessChecks() {
        return Optional.ofNullable(this.readinessChecks);
    }

    /**
     * Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster&#39;s region must be specified—for example, `[&#39;us-central1-a&#39;, &#39;us-central1-f&#39;]`.
     * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
     * 
     */
    @Import(name="replicaZones")
    private @Nullable Output<List<String>> replicaZones;

    /**
     * @return Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster&#39;s region must be specified—for example, `[&#39;us-central1-a&#39;, &#39;us-central1-f&#39;]`.
     * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
     * 
     */
    public Optional<Output<List<String>>> replicaZones() {
        return Optional.ofNullable(this.replicaZones);
    }

    /**
     * How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryption_key` is set. Defaults to 12 hours.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Import(name="runningTimeout")
    private @Nullable Output<String> runningTimeout;

    /**
     * @return How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryption_key` is set. Defaults to 12 hours.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<Output<String>> runningTimeout() {
        return Optional.ofNullable(this.runningTimeout);
    }

    /**
     * The ID of the parent workstation cluster.
     * 
     */
    @Import(name="workstationClusterId", required=true)
    private Output<String> workstationClusterId;

    /**
     * @return The ID of the parent workstation cluster.
     * 
     */
    public Output<String> workstationClusterId() {
        return this.workstationClusterId;
    }

    /**
     * The ID to be assigned to the workstation cluster config.
     * 
     */
    @Import(name="workstationConfigId", required=true)
    private Output<String> workstationConfigId;

    /**
     * @return The ID to be assigned to the workstation cluster config.
     * 
     */
    public Output<String> workstationConfigId() {
        return this.workstationConfigId;
    }

    private WorkstationConfigArgs() {}

    private WorkstationConfigArgs(WorkstationConfigArgs $) {
        this.annotations = $.annotations;
        this.container = $.container;
        this.disableTcpConnections = $.disableTcpConnections;
        this.displayName = $.displayName;
        this.enableAuditAgent = $.enableAuditAgent;
        this.encryptionKey = $.encryptionKey;
        this.ephemeralDirectories = $.ephemeralDirectories;
        this.host = $.host;
        this.idleTimeout = $.idleTimeout;
        this.labels = $.labels;
        this.location = $.location;
        this.persistentDirectories = $.persistentDirectories;
        this.project = $.project;
        this.readinessChecks = $.readinessChecks;
        this.replicaZones = $.replicaZones;
        this.runningTimeout = $.runningTimeout;
        this.workstationClusterId = $.workstationClusterId;
        this.workstationConfigId = $.workstationConfigId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkstationConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkstationConfigArgs $;

        public Builder() {
            $ = new WorkstationConfigArgs();
        }

        public Builder(WorkstationConfigArgs defaults) {
            $ = new WorkstationConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param annotations Client-specified annotations. This is distinct from labels.
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
         * @param annotations Client-specified annotations. This is distinct from labels.
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
         * @param container Container that will be run for each workstation using this configuration when that workstation is started.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder container(@Nullable Output<WorkstationConfigContainerArgs> container) {
            $.container = container;
            return this;
        }

        /**
         * @param container Container that will be run for each workstation using this configuration when that workstation is started.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder container(WorkstationConfigContainerArgs container) {
            return container(Output.of(container));
        }

        /**
         * @param disableTcpConnections Disables support for plain TCP connections in the workstation. By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
         * 
         * @return builder
         * 
         */
        public Builder disableTcpConnections(@Nullable Output<Boolean> disableTcpConnections) {
            $.disableTcpConnections = disableTcpConnections;
            return this;
        }

        /**
         * @param disableTcpConnections Disables support for plain TCP connections in the workstation. By default the service supports TCP connections via a websocket relay. Setting this option to true disables that relay, which prevents the usage of services that require plain tcp connections, such as ssh. When enabled, all communication must occur over https or wss.
         * 
         * @return builder
         * 
         */
        public Builder disableTcpConnections(Boolean disableTcpConnections) {
            return disableTcpConnections(Output.of(disableTcpConnections));
        }

        /**
         * @param displayName Human-readable name for this resource.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human-readable name for this resource.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param enableAuditAgent Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
         * 
         * @return builder
         * 
         */
        public Builder enableAuditAgent(@Nullable Output<Boolean> enableAuditAgent) {
            $.enableAuditAgent = enableAuditAgent;
            return this;
        }

        /**
         * @param enableAuditAgent Whether to enable Linux `auditd` logging on the workstation. When enabled, a service account must also be specified that has `logging.buckets.write` permission on the project. Operating system audit logging is distinct from Cloud Audit Logs.
         * 
         * @return builder
         * 
         */
        public Builder enableAuditAgent(Boolean enableAuditAgent) {
            return enableAuditAgent(Output.of(enableAuditAgent));
        }

        /**
         * @param encryptionKey Encrypts resources of this workstation configuration using a customer-managed encryption key.
         * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
         * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
         * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionKey(@Nullable Output<WorkstationConfigEncryptionKeyArgs> encryptionKey) {
            $.encryptionKey = encryptionKey;
            return this;
        }

        /**
         * @param encryptionKey Encrypts resources of this workstation configuration using a customer-managed encryption key.
         * If specified, the boot disk of the Compute Engine instance and the persistent disk are encrypted using this encryption key. If this field is not set, the disks are encrypted using a generated key. Customer-managed encryption keys do not protect disk metadata.
         * If the customer-managed encryption key is rotated, when the workstation instance is stopped, the system attempts to recreate the persistent disk with the new version of the key. Be sure to keep older versions of the key until the persistent disk is recreated. Otherwise, data on the persistent disk will be lost.
         * If the encryption key is revoked, the workstation session will automatically be stopped within 7 hours.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionKey(WorkstationConfigEncryptionKeyArgs encryptionKey) {
            return encryptionKey(Output.of(encryptionKey));
        }

        /**
         * @param ephemeralDirectories Ephemeral directories which won&#39;t persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ephemeralDirectories(@Nullable Output<List<WorkstationConfigEphemeralDirectoryArgs>> ephemeralDirectories) {
            $.ephemeralDirectories = ephemeralDirectories;
            return this;
        }

        /**
         * @param ephemeralDirectories Ephemeral directories which won&#39;t persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ephemeralDirectories(List<WorkstationConfigEphemeralDirectoryArgs> ephemeralDirectories) {
            return ephemeralDirectories(Output.of(ephemeralDirectories));
        }

        /**
         * @param ephemeralDirectories Ephemeral directories which won&#39;t persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ephemeralDirectories(WorkstationConfigEphemeralDirectoryArgs... ephemeralDirectories) {
            return ephemeralDirectories(List.of(ephemeralDirectories));
        }

        /**
         * @param host Runtime host for a workstation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<WorkstationConfigHostArgs> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Runtime host for a workstation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder host(WorkstationConfigHostArgs host) {
            return host(Output.of(host));
        }

        /**
         * @param idleTimeout How long to wait before automatically stopping an instance that hasn&#39;t recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
         * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(@Nullable Output<String> idleTimeout) {
            $.idleTimeout = idleTimeout;
            return this;
        }

        /**
         * @param idleTimeout How long to wait before automatically stopping an instance that hasn&#39;t recently received any user traffic. A value of 0 indicates that this instance should never time out from idleness. Defaults to 20 minutes.
         * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder idleTimeout(String idleTimeout) {
            return idleTimeout(Output.of(idleTimeout));
        }

        /**
         * @param labels Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
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
         * @param labels Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
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
         * @param location The location where the workstation cluster config should reside.
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
         * @param location The location where the workstation cluster config should reside.
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
         * @param persistentDirectories Directories to persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder persistentDirectories(@Nullable Output<List<WorkstationConfigPersistentDirectoryArgs>> persistentDirectories) {
            $.persistentDirectories = persistentDirectories;
            return this;
        }

        /**
         * @param persistentDirectories Directories to persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder persistentDirectories(List<WorkstationConfigPersistentDirectoryArgs> persistentDirectories) {
            return persistentDirectories(Output.of(persistentDirectories));
        }

        /**
         * @param persistentDirectories Directories to persist across workstation sessions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder persistentDirectories(WorkstationConfigPersistentDirectoryArgs... persistentDirectories) {
            return persistentDirectories(List.of(persistentDirectories));
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
         * @param readinessChecks Readiness checks to be performed on a workstation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder readinessChecks(@Nullable Output<List<WorkstationConfigReadinessCheckArgs>> readinessChecks) {
            $.readinessChecks = readinessChecks;
            return this;
        }

        /**
         * @param readinessChecks Readiness checks to be performed on a workstation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder readinessChecks(List<WorkstationConfigReadinessCheckArgs> readinessChecks) {
            return readinessChecks(Output.of(readinessChecks));
        }

        /**
         * @param readinessChecks Readiness checks to be performed on a workstation.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder readinessChecks(WorkstationConfigReadinessCheckArgs... readinessChecks) {
            return readinessChecks(List.of(readinessChecks));
        }

        /**
         * @param replicaZones Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster&#39;s region must be specified—for example, `[&#39;us-central1-a&#39;, &#39;us-central1-f&#39;]`.
         * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
         * 
         * @return builder
         * 
         */
        public Builder replicaZones(@Nullable Output<List<String>> replicaZones) {
            $.replicaZones = replicaZones;
            return this;
        }

        /**
         * @param replicaZones Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster&#39;s region must be specified—for example, `[&#39;us-central1-a&#39;, &#39;us-central1-f&#39;]`.
         * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
         * 
         * @return builder
         * 
         */
        public Builder replicaZones(List<String> replicaZones) {
            return replicaZones(Output.of(replicaZones));
        }

        /**
         * @param replicaZones Specifies the zones used to replicate the VM and disk resources within the region. If set, exactly two zones within the workstation cluster&#39;s region must be specified—for example, `[&#39;us-central1-a&#39;, &#39;us-central1-f&#39;]`.
         * If this field is empty, two default zones within the region are used. Immutable after the workstation configuration is created.
         * 
         * @return builder
         * 
         */
        public Builder replicaZones(String... replicaZones) {
            return replicaZones(List.of(replicaZones));
        }

        /**
         * @param runningTimeout How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryption_key` is set. Defaults to 12 hours.
         * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder runningTimeout(@Nullable Output<String> runningTimeout) {
            $.runningTimeout = runningTimeout;
            return this;
        }

        /**
         * @param runningTimeout How long to wait before automatically stopping a workstation after it was started. A value of 0 indicates that workstations using this configuration should never time out from running duration. Must be greater than 0 and less than 24 hours if `encryption_key` is set. Defaults to 12 hours.
         * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder runningTimeout(String runningTimeout) {
            return runningTimeout(Output.of(runningTimeout));
        }

        /**
         * @param workstationClusterId The ID of the parent workstation cluster.
         * 
         * @return builder
         * 
         */
        public Builder workstationClusterId(Output<String> workstationClusterId) {
            $.workstationClusterId = workstationClusterId;
            return this;
        }

        /**
         * @param workstationClusterId The ID of the parent workstation cluster.
         * 
         * @return builder
         * 
         */
        public Builder workstationClusterId(String workstationClusterId) {
            return workstationClusterId(Output.of(workstationClusterId));
        }

        /**
         * @param workstationConfigId The ID to be assigned to the workstation cluster config.
         * 
         * @return builder
         * 
         */
        public Builder workstationConfigId(Output<String> workstationConfigId) {
            $.workstationConfigId = workstationConfigId;
            return this;
        }

        /**
         * @param workstationConfigId The ID to be assigned to the workstation cluster config.
         * 
         * @return builder
         * 
         */
        public Builder workstationConfigId(String workstationConfigId) {
            return workstationConfigId(Output.of(workstationConfigId));
        }

        public WorkstationConfigArgs build() {
            if ($.location == null) {
                throw new MissingRequiredPropertyException("WorkstationConfigArgs", "location");
            }
            if ($.workstationClusterId == null) {
                throw new MissingRequiredPropertyException("WorkstationConfigArgs", "workstationClusterId");
            }
            if ($.workstationConfigId == null) {
                throw new MissingRequiredPropertyException("WorkstationConfigArgs", "workstationConfigId");
            }
            return $;
        }
    }

}
