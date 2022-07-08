// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.pubsub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.pubsub.inputs.LiteTopicPartitionConfigArgs;
import com.pulumi.gcp.pubsub.inputs.LiteTopicReservationConfigArgs;
import com.pulumi.gcp.pubsub.inputs.LiteTopicRetentionConfigArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LiteTopicArgs extends com.pulumi.resources.ResourceArgs {

    public static final LiteTopicArgs Empty = new LiteTopicArgs();

    /**
     * Name of the topic.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the topic.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The settings for this topic&#39;s partitions.
     * Structure is documented below.
     * 
     */
    @Import(name="partitionConfig")
    private @Nullable Output<LiteTopicPartitionConfigArgs> partitionConfig;

    /**
     * @return The settings for this topic&#39;s partitions.
     * Structure is documented below.
     * 
     */
    public Optional<Output<LiteTopicPartitionConfigArgs>> partitionConfig() {
        return Optional.ofNullable(this.partitionConfig);
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
     * The region of the pubsub lite topic.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region of the pubsub lite topic.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The settings for this topic&#39;s Reservation usage.
     * Structure is documented below.
     * 
     */
    @Import(name="reservationConfig")
    private @Nullable Output<LiteTopicReservationConfigArgs> reservationConfig;

    /**
     * @return The settings for this topic&#39;s Reservation usage.
     * Structure is documented below.
     * 
     */
    public Optional<Output<LiteTopicReservationConfigArgs>> reservationConfig() {
        return Optional.ofNullable(this.reservationConfig);
    }

    /**
     * The settings for a topic&#39;s message retention.
     * Structure is documented below.
     * 
     */
    @Import(name="retentionConfig")
    private @Nullable Output<LiteTopicRetentionConfigArgs> retentionConfig;

    /**
     * @return The settings for a topic&#39;s message retention.
     * Structure is documented below.
     * 
     */
    public Optional<Output<LiteTopicRetentionConfigArgs>> retentionConfig() {
        return Optional.ofNullable(this.retentionConfig);
    }

    /**
     * The zone of the pubsub lite topic.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return The zone of the pubsub lite topic.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private LiteTopicArgs() {}

    private LiteTopicArgs(LiteTopicArgs $) {
        this.name = $.name;
        this.partitionConfig = $.partitionConfig;
        this.project = $.project;
        this.region = $.region;
        this.reservationConfig = $.reservationConfig;
        this.retentionConfig = $.retentionConfig;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LiteTopicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LiteTopicArgs $;

        public Builder() {
            $ = new LiteTopicArgs();
        }

        public Builder(LiteTopicArgs defaults) {
            $ = new LiteTopicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param partitionConfig The settings for this topic&#39;s partitions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder partitionConfig(@Nullable Output<LiteTopicPartitionConfigArgs> partitionConfig) {
            $.partitionConfig = partitionConfig;
            return this;
        }

        /**
         * @param partitionConfig The settings for this topic&#39;s partitions.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder partitionConfig(LiteTopicPartitionConfigArgs partitionConfig) {
            return partitionConfig(Output.of(partitionConfig));
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
         * @param region The region of the pubsub lite topic.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region of the pubsub lite topic.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param reservationConfig The settings for this topic&#39;s Reservation usage.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder reservationConfig(@Nullable Output<LiteTopicReservationConfigArgs> reservationConfig) {
            $.reservationConfig = reservationConfig;
            return this;
        }

        /**
         * @param reservationConfig The settings for this topic&#39;s Reservation usage.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder reservationConfig(LiteTopicReservationConfigArgs reservationConfig) {
            return reservationConfig(Output.of(reservationConfig));
        }

        /**
         * @param retentionConfig The settings for a topic&#39;s message retention.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder retentionConfig(@Nullable Output<LiteTopicRetentionConfigArgs> retentionConfig) {
            $.retentionConfig = retentionConfig;
            return this;
        }

        /**
         * @param retentionConfig The settings for a topic&#39;s message retention.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder retentionConfig(LiteTopicRetentionConfigArgs retentionConfig) {
            return retentionConfig(Output.of(retentionConfig));
        }

        /**
         * @param zone The zone of the pubsub lite topic.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The zone of the pubsub lite topic.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public LiteTopicArgs build() {
            return $;
        }
    }

}
