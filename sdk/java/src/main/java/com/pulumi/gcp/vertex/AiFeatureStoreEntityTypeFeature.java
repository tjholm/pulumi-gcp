// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeFeatureArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeFeatureState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Feature Metadata information that describes an attribute of an entity type. For example, apple is an entity type, and color is a feature that describes apple.
 * 
 * To get more information about FeaturestoreEntitytypeFeature, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 * 
 * ## Example Usage
 * ### Vertex Ai Featurestore Entitytype Feature
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityType;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeFeature;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeFeatureArgs;
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
 *             .build());
 * 
 *         var entity = new AiFeatureStoreEntityType(&#34;entity&#34;, AiFeatureStoreEntityTypeArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .featurestore(featurestore.id())
 *             .build());
 * 
 *         var feature = new AiFeatureStoreEntityTypeFeature(&#34;feature&#34;, AiFeatureStoreEntityTypeFeatureArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .entitytype(entity.id())
 *             .valueType(&#34;INT64_ARRAY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featurestore Entitytype Feature With Beta Fields
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityType;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeFeature;
 * import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeFeatureArgs;
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
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var feature = new AiFeatureStoreEntityTypeFeature(&#34;feature&#34;, AiFeatureStoreEntityTypeFeatureArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .entitytype(entity.id())
 *             .valueType(&#34;INT64_ARRAY&#34;)
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
 * FeaturestoreEntitytypeFeature can be imported using any of these accepted formats* `{{entitytype}}/features/{{name}}` When using the `pulumi import` command, FeaturestoreEntitytypeFeature can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature default {{entitytype}}/features/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature")
public class AiFeatureStoreEntityTypeFeature extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp of when the entity type was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the entity type was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the feature.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the feature.
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
     * The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
     * 
     * ***
     * 
     */
    @Export(name="entitytype", refs={String.class}, tree="[0]")
    private Output<String> entitytype;

    /**
     * @return The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
     * 
     * ***
     * 
     */
    public Output<String> entitytype() {
        return this.entitytype;
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
     * A set of key/value label pairs to assign to the feature.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to the feature.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The region of the feature
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of the feature
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The timestamp when the entity type was most recently updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp when the entity type was most recently updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
     * 
     */
    @Export(name="valueType", refs={String.class}, tree="[0]")
    private Output<String> valueType;

    /**
     * @return Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
     * 
     */
    public Output<String> valueType() {
        return this.valueType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiFeatureStoreEntityTypeFeature(String name) {
        this(name, AiFeatureStoreEntityTypeFeatureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureStoreEntityTypeFeature(String name, AiFeatureStoreEntityTypeFeatureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureStoreEntityTypeFeature(String name, AiFeatureStoreEntityTypeFeatureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, args == null ? AiFeatureStoreEntityTypeFeatureArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureStoreEntityTypeFeature(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeFeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, state, makeResourceOptions(options, id));
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
    public static AiFeatureStoreEntityTypeFeature get(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeFeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureStoreEntityTypeFeature(name, id, state, options);
    }
}
