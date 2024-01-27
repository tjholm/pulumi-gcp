// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeState;
import com.pulumi.gcp.vertex.outputs.AiFeatureStoreEntityTypeMonitoringConfig;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An entity type is a type of object in a system that needs to be modeled and have stored information about. For example, driver is an entity type, and driver0 is an instance of an entity type driver.
 * 
 * To get more information about FeaturestoreEntitytype, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 * 
 * ## Example Usage
 * ### Vertex Ai Featurestore Entitytype
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEncryptionSpecArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityType;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigImportFeaturesAnalysisArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var featurestore = new AiFeatureStore(&#34;featurestore&#34;, AiFeatureStoreArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .region(&#34;us-central1&#34;)
 *             .onlineServingConfig(AiFeatureStoreOnlineServingConfigArgs.builder()
 *                 .fixedNodeCount(2)
 *                 .build())
 *             .encryptionSpec(AiFeatureStoreEncryptionSpecArgs.builder()
 *                 .kmsKeyName(&#34;kms-name&#34;)
 *                 .build())
 *             .build());
 * 
 *         var entity = new AiFeatureStoreEntityType(&#34;entity&#34;, AiFeatureStoreEntityTypeArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .description(&#34;test description&#34;)
 *             .featurestore(featurestore.id())
 *             .monitoringConfig(AiFeatureStoreEntityTypeMonitoringConfigArgs.builder()
 *                 .snapshotAnalysis(AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs.builder()
 *                     .disabled(false)
 *                     .monitoringIntervalDays(1)
 *                     .stalenessDays(21)
 *                     .build())
 *                 .numericalThresholdConfig(AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs.builder()
 *                     .value(0.8)
 *                     .build())
 *                 .categoricalThresholdConfig(AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs.builder()
 *                     .value(10)
 *                     .build())
 *                 .importFeaturesAnalysis(AiFeatureStoreEntityTypeMonitoringConfigImportFeaturesAnalysisArgs.builder()
 *                     .state(&#34;ENABLED&#34;)
 *                     .anomalyDetectionBaseline(&#34;PREVIOUS_IMPORT_FEATURES_STATS&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featurestore Entitytype With Beta Fields
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEncryptionSpecArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityType;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var featurestore = new AiFeatureStore(&#34;featurestore&#34;, AiFeatureStoreArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .region(&#34;us-central1&#34;)
 *             .onlineServingConfig(AiFeatureStoreOnlineServingConfigArgs.builder()
 *                 .fixedNodeCount(2)
 *                 .build())
 *             .encryptionSpec(AiFeatureStoreEncryptionSpecArgs.builder()
 *                 .kmsKeyName(&#34;kms-name&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var entity = new AiFeatureStoreEntityType(&#34;entity&#34;, AiFeatureStoreEntityTypeArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .featurestore(featurestore.id())
 *             .monitoringConfig(AiFeatureStoreEntityTypeMonitoringConfigArgs.builder()
 *                 .snapshotAnalysis(AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs.builder()
 *                     .disabled(false)
 *                     .monitoringInterval(&#34;86400s&#34;)
 *                     .build())
 *                 .categoricalThresholdConfig(AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs.builder()
 *                     .value(0.3)
 *                     .build())
 *                 .numericalThresholdConfig(AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs.builder()
 *                     .value(0.3)
 *                     .build())
 *                 .build())
 *             .offlineStorageTtlDays(30)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * FeaturestoreEntitytype can be imported using any of these accepted formats* `{{featurestore}}/entityTypes/{{name}}` When using the `pulumi import` command, FeaturestoreEntitytype can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType default {{featurestore}}/entityTypes/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType")
public class AiFeatureStoreEntityType extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the featurestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. Description of the EntityType.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. Description of the EntityType.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Used to perform consistent read-modify-write updates.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return Used to perform consistent read-modify-write updates.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
     * 
     * ***
     * 
     */
    @Export(name="featurestore", refs={String.class}, tree="[0]")
    private Output<String> featurestore;

    /**
     * @return The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}.
     * 
     * ***
     * 
     */
    public Output<String> featurestore() {
        return this.featurestore;
    }
    /**
     * A set of key/value label pairs to assign to this EntityType.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this EntityType.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The default monitoring configuration for all Features under this EntityType.
     * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
     * Structure is documented below.
     * 
     */
    @Export(name="monitoringConfig", refs={AiFeatureStoreEntityTypeMonitoringConfig.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureStoreEntityTypeMonitoringConfig> monitoringConfig;

    /**
     * @return The default monitoring configuration for all Features under this EntityType.
     * If this is populated with [FeaturestoreMonitoringConfig.monitoring_interval] specified, snapshot analysis monitoring is enabled. Otherwise, snapshot analysis monitoring is disabled.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureStoreEntityTypeMonitoringConfig>> monitoringConfig() {
        return Codegen.optional(this.monitoringConfig);
    }
    /**
     * The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the EntityType. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (Optional, Beta)
     * Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
     * 
     */
    @Export(name="offlineStorageTtlDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> offlineStorageTtlDays;

    /**
     * @return (Optional, Beta)
     * Config for data retention policy in offline storage. TTL in days for feature values that will be stored in offline storage. The Feature Store offline storage periodically removes obsolete feature values older than offlineStorageTtlDays since the feature generation time. If unset (or explicitly set to 0), default to 4000 days TTL.
     * 
     */
    public Output<Optional<Integer>> offlineStorageTtlDays() {
        return Codegen.optional(this.offlineStorageTtlDays);
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The region of the EntityType.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of the EntityType.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp of when the featurestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiFeatureStoreEntityType(String name) {
        this(name, AiFeatureStoreEntityTypeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureStoreEntityType(String name, AiFeatureStoreEntityTypeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureStoreEntityType(String name, AiFeatureStoreEntityTypeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType", name, args == null ? AiFeatureStoreEntityTypeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureStoreEntityType(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AiFeatureStoreEntityType get(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureStoreEntityType(name, id, state, options);
    }
}
