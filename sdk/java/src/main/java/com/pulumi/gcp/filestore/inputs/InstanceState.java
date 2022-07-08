// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.filestore.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.filestore.inputs.InstanceFileSharesArgs;
import com.pulumi.gcp.filestore.inputs.InstanceNetworkArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceState Empty = new InstanceState();

    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * A description of the instance.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the instance.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     * 
     */
    @Import(name="fileShares")
    private @Nullable Output<InstanceFileSharesArgs> fileShares;

    /**
     * @return File system shares on the instance. For this version, only a
     * single file share is supported.
     * Structure is documented below.
     * 
     */
    public Optional<Output<InstanceFileSharesArgs>> fileShares() {
        return Optional.ofNullable(this.fileShares);
    }

    /**
     * KMS key name used for data encryption.
     * 
     */
    @Import(name="kmsKeyName")
    private @Nullable Output<String> kmsKeyName;

    /**
     * @return KMS key name used for data encryption.
     * 
     */
    public Optional<Output<String>> kmsKeyName() {
        return Optional.ofNullable(this.kmsKeyName);
    }

    /**
     * Resource labels to represent user-provided metadata.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Resource labels to represent user-provided metadata.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The name of the fileshare (16 characters or less)
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the fileshare (16 characters or less)
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     * 
     */
    @Import(name="networks")
    private @Nullable Output<List<InstanceNetworkArgs>> networks;

    /**
     * @return VPC networks to which the instance is connected. For this version,
     * only a single network is supported.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<InstanceNetworkArgs>>> networks() {
        return Optional.ofNullable(this.networks);
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
     * The service tier of the instance.
     * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
     * 
     */
    @Import(name="tier")
    private @Nullable Output<String> tier;

    /**
     * @return The service tier of the instance.
     * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
     * 
     */
    public Optional<Output<String>> tier() {
        return Optional.ofNullable(this.tier);
    }

    /**
     * - 
     * (Optional, Deprecated)
     * The name of the Filestore zone of the instance.
     * 
     * @deprecated
     * Deprecated in favor of location.
     * 
     */
    @Deprecated /* Deprecated in favor of location. */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return -
     * (Optional, Deprecated)
     * The name of the Filestore zone of the instance.
     * 
     * @deprecated
     * Deprecated in favor of location.
     * 
     */
    @Deprecated /* Deprecated in favor of location. */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private InstanceState() {}

    private InstanceState(InstanceState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.etag = $.etag;
        this.fileShares = $.fileShares;
        this.kmsKeyName = $.kmsKeyName;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.networks = $.networks;
        this.project = $.project;
        this.tier = $.tier;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceState $;

        public Builder() {
            $ = new InstanceState();
        }

        public Builder(InstanceState defaults) {
            $ = new InstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime Creation timestamp in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Creation timestamp in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description A description of the instance.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the instance.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param etag Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param fileShares File system shares on the instance. For this version, only a
         * single file share is supported.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder fileShares(@Nullable Output<InstanceFileSharesArgs> fileShares) {
            $.fileShares = fileShares;
            return this;
        }

        /**
         * @param fileShares File system shares on the instance. For this version, only a
         * single file share is supported.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder fileShares(InstanceFileSharesArgs fileShares) {
            return fileShares(Output.of(fileShares));
        }

        /**
         * @param kmsKeyName KMS key name used for data encryption.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(@Nullable Output<String> kmsKeyName) {
            $.kmsKeyName = kmsKeyName;
            return this;
        }

        /**
         * @param kmsKeyName KMS key name used for data encryption.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(String kmsKeyName) {
            return kmsKeyName(Output.of(kmsKeyName));
        }

        /**
         * @param labels Resource labels to represent user-provided metadata.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Resource labels to represent user-provided metadata.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name The name of the fileshare (16 characters or less)
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the fileshare (16 characters or less)
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networks VPC networks to which the instance is connected. For this version,
         * only a single network is supported.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder networks(@Nullable Output<List<InstanceNetworkArgs>> networks) {
            $.networks = networks;
            return this;
        }

        /**
         * @param networks VPC networks to which the instance is connected. For this version,
         * only a single network is supported.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder networks(List<InstanceNetworkArgs> networks) {
            return networks(Output.of(networks));
        }

        /**
         * @param networks VPC networks to which the instance is connected. For this version,
         * only a single network is supported.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder networks(InstanceNetworkArgs... networks) {
            return networks(List.of(networks));
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
         * @param tier The service tier of the instance.
         * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
         * 
         * @return builder
         * 
         */
        public Builder tier(@Nullable Output<String> tier) {
            $.tier = tier;
            return this;
        }

        /**
         * @param tier The service tier of the instance.
         * Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE
         * 
         * @return builder
         * 
         */
        public Builder tier(String tier) {
            return tier(Output.of(tier));
        }

        /**
         * @param zone -
         * (Optional, Deprecated)
         * The name of the Filestore zone of the instance.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated in favor of location.
         * 
         */
        @Deprecated /* Deprecated in favor of location. */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone -
         * (Optional, Deprecated)
         * The name of the Filestore zone of the instance.
         * 
         * @return builder
         * 
         * @deprecated
         * Deprecated in favor of location.
         * 
         */
        @Deprecated /* Deprecated in favor of location. */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public InstanceState build() {
            return $;
        }
    }

}
