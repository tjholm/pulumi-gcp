// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewSyncConfigArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiFeatureOnlineStoreFeatureviewState extends com.pulumi.resources.ResourceArgs {

    public static final AiFeatureOnlineStoreFeatureviewState Empty = new AiFeatureOnlineStoreFeatureviewState();

    /**
     * Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    @Import(name="bigQuerySource")
    private @Nullable Output<AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs> bigQuerySource;

    /**
     * @return Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs>> bigQuerySource() {
        return Optional.ofNullable(this.bigQuerySource);
    }

    /**
     * The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
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
     * The name of the FeatureOnlineStore to use for the featureview.
     * 
     */
    @Import(name="featureOnlineStore")
    private @Nullable Output<String> featureOnlineStore;

    /**
     * @return The name of the FeatureOnlineStore to use for the featureview.
     * 
     */
    public Optional<Output<String>> featureOnlineStore() {
        return Optional.ofNullable(this.featureOnlineStore);
    }

    /**
     * A set of key/value label pairs to assign to this FeatureView.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this FeatureView.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * The region for the resource. It should be the same as the featureonlinestore region.
     * 
     * ***
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region for the resource. It should be the same as the featureonlinestore region.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
     * Structure is documented below.
     * 
     */
    @Import(name="syncConfig")
    private @Nullable Output<AiFeatureOnlineStoreFeatureviewSyncConfigArgs> syncConfig;

    /**
     * @return Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AiFeatureOnlineStoreFeatureviewSyncConfigArgs>> syncConfig() {
        return Optional.ofNullable(this.syncConfig);
    }

    /**
     * The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    /**
     * Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
     * Structure is documented below.
     * 
     */
    @Import(name="vectorSearchConfig")
    private @Nullable Output<AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs> vectorSearchConfig;

    /**
     * @return Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs>> vectorSearchConfig() {
        return Optional.ofNullable(this.vectorSearchConfig);
    }

    private AiFeatureOnlineStoreFeatureviewState() {}

    private AiFeatureOnlineStoreFeatureviewState(AiFeatureOnlineStoreFeatureviewState $) {
        this.bigQuerySource = $.bigQuerySource;
        this.createTime = $.createTime;
        this.effectiveLabels = $.effectiveLabels;
        this.featureOnlineStore = $.featureOnlineStore;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.region = $.region;
        this.syncConfig = $.syncConfig;
        this.updateTime = $.updateTime;
        this.vectorSearchConfig = $.vectorSearchConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiFeatureOnlineStoreFeatureviewState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiFeatureOnlineStoreFeatureviewState $;

        public Builder() {
            $ = new AiFeatureOnlineStoreFeatureviewState();
        }

        public Builder(AiFeatureOnlineStoreFeatureviewState defaults) {
            $ = new AiFeatureOnlineStoreFeatureviewState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bigQuerySource Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder bigQuerySource(@Nullable Output<AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs> bigQuerySource) {
            $.bigQuerySource = bigQuerySource;
            return this;
        }

        /**
         * @param bigQuerySource Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder bigQuerySource(AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs bigQuerySource) {
            return bigQuerySource(Output.of(bigQuerySource));
        }

        /**
         * @param createTime The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
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
         * @param featureOnlineStore The name of the FeatureOnlineStore to use for the featureview.
         * 
         * @return builder
         * 
         */
        public Builder featureOnlineStore(@Nullable Output<String> featureOnlineStore) {
            $.featureOnlineStore = featureOnlineStore;
            return this;
        }

        /**
         * @param featureOnlineStore The name of the FeatureOnlineStore to use for the featureview.
         * 
         * @return builder
         * 
         */
        public Builder featureOnlineStore(String featureOnlineStore) {
            return featureOnlineStore(Output.of(featureOnlineStore));
        }

        /**
         * @param labels A set of key/value label pairs to assign to this FeatureView.
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
         * @param labels A set of key/value label pairs to assign to this FeatureView.
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
         * @param name Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param region The region for the resource. It should be the same as the featureonlinestore region.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region for the resource. It should be the same as the featureonlinestore region.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param syncConfig Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder syncConfig(@Nullable Output<AiFeatureOnlineStoreFeatureviewSyncConfigArgs> syncConfig) {
            $.syncConfig = syncConfig;
            return this;
        }

        /**
         * @param syncConfig Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder syncConfig(AiFeatureOnlineStoreFeatureviewSyncConfigArgs syncConfig) {
            return syncConfig(Output.of(syncConfig));
        }

        /**
         * @param updateTime The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        /**
         * @param vectorSearchConfig Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder vectorSearchConfig(@Nullable Output<AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs> vectorSearchConfig) {
            $.vectorSearchConfig = vectorSearchConfig;
            return this;
        }

        /**
         * @param vectorSearchConfig Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder vectorSearchConfig(AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs vectorSearchConfig) {
            return vectorSearchConfig(Output.of(vectorSearchConfig));
        }

        public AiFeatureOnlineStoreFeatureviewState build() {
            return $;
        }
    }

}
