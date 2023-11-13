// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiFeatureStoreEntityTypeState extends com.pulumi.resources.ResourceArgs {

    public static final AiFeatureStoreEntityTypeState Empty = new AiFeatureStoreEntityTypeState();

    /**
     * The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Optional. Description of the EntityType.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Optional. Description of the EntityType.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * Used to perform consistent read-modify-write updates.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return Used to perform consistent read-modify-write updates.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
     * 
     * ***
     * 
     */
    @Import(name="featurestore")
    private @Nullable Output<String> featurestore;

    /**
     * @return The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> featurestore() {
        return Optional.ofNullable(this.featurestore);
    }

    /**
     * A set of key/value label pairs to assign to this EntityType.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this EntityType.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The default monitoring configuration for all Features under this EntityType.
     * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
     * Structure is documented below.
     * 
     */
    @Import(name="monitoringConfig")
    private @Nullable Output<AiFeatureStoreEntityTypeMonitoringConfigArgs> monitoringConfig;

    /**
     * @return The default monitoring configuration for all Features under this EntityType.
     * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AiFeatureStoreEntityTypeMonitoringConfigArgs>> monitoringConfig() {
        return Optional.ofNullable(this.monitoringConfig);
    }

    /**
     * The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline
     * storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays
     * since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
     * 
     */
    @Import(name="offlineStorageTtlDays")
    private @Nullable Output<Integer> offlineStorageTtlDays;

    /**
     * @return Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline
     * storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays
     * since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
     * 
     */
    public Optional<Output<Integer>> offlineStorageTtlDays() {
        return Optional.ofNullable(this.offlineStorageTtlDays);
    }

    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * The region of the EntityType.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region of the EntityType.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private AiFeatureStoreEntityTypeState() {}

    private AiFeatureStoreEntityTypeState(AiFeatureStoreEntityTypeState $) {
        this.createTime = $.createTime;
        this.description = $.description;
        this.effectiveLabels = $.effectiveLabels;
        this.etag = $.etag;
        this.featurestore = $.featurestore;
        this.labels = $.labels;
        this.monitoringConfig = $.monitoringConfig;
        this.name = $.name;
        this.offlineStorageTtlDays = $.offlineStorageTtlDays;
        this.pulumiLabels = $.pulumiLabels;
        this.region = $.region;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiFeatureStoreEntityTypeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiFeatureStoreEntityTypeState $;

        public Builder() {
            $ = new AiFeatureStoreEntityTypeState();
        }

        public Builder(AiFeatureStoreEntityTypeState defaults) {
            $ = new AiFeatureStoreEntityTypeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description Optional. Description of the EntityType.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Optional. Description of the EntityType.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param etag Used to perform consistent read-modify-write updates.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag Used to perform consistent read-modify-write updates.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param featurestore The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder featurestore(@Nullable Output<String> featurestore) {
            $.featurestore = featurestore;
            return this;
        }

        /**
         * @param featurestore The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder featurestore(String featurestore) {
            return featurestore(Output.of(featurestore));
        }

        /**
         * @param labels A set of key/value label pairs to assign to this EntityType.
         * 
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
         * @param labels A set of key/value label pairs to assign to this EntityType.
         * 
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
         * @param monitoringConfig The default monitoring configuration for all Features under this EntityType.
         * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder monitoringConfig(@Nullable Output<AiFeatureStoreEntityTypeMonitoringConfigArgs> monitoringConfig) {
            $.monitoringConfig = monitoringConfig;
            return this;
        }

        /**
         * @param monitoringConfig The default monitoring configuration for all Features under this EntityType.
         * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder monitoringConfig(AiFeatureStoreEntityTypeMonitoringConfigArgs monitoringConfig) {
            return monitoringConfig(Output.of(monitoringConfig));
        }

        /**
         * @param name The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param offlineStorageTtlDays Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline
         * storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays
         * since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
         * 
         * @return builder
         * 
         */
        public Builder offlineStorageTtlDays(@Nullable Output<Integer> offlineStorageTtlDays) {
            $.offlineStorageTtlDays = offlineStorageTtlDays;
            return this;
        }

        /**
         * @param offlineStorageTtlDays Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline
         * storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays
         * since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
         * 
         * @return builder
         * 
         */
        public Builder offlineStorageTtlDays(Integer offlineStorageTtlDays) {
            return offlineStorageTtlDays(Output.of(offlineStorageTtlDays));
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param region The region of the EntityType.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region of the EntityType.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param updateTime The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public AiFeatureStoreEntityTypeState build() {
            return $;
        }
    }

}
