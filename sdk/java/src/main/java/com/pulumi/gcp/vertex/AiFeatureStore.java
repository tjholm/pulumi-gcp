// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreState;
import com.pulumi.gcp.vertex.outputs.AiFeatureStoreEncryptionSpec;
import com.pulumi.gcp.vertex.outputs.AiFeatureStoreOnlineServingConfig;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A collection of DataItems and Annotations on them.
 * 
 * To get more information about Featurestore, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 * 
 * ## Example Usage
 * ### Vertex Ai Featurestore
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEncryptionSpecArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
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
 *             .encryptionSpec(AiFeatureStoreEncryptionSpecArgs.builder()
 *                 .kmsKeyName(&#34;kms-name&#34;)
 *                 .build())
 *             .forceDestroy(true)
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .onlineServingConfig(AiFeatureStoreOnlineServingConfigArgs.builder()
 *                 .fixedNodeCount(2)
 *                 .build())
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featurestore With Beta Fields
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
 *             .onlineStorageTtlDays(30)
 *             .forceDestroy(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featurestore Scaling
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureStore;
 * import com.pulumi.gcp.vertex.AiFeatureStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEncryptionSpecArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureStoreOnlineServingConfigScalingArgs;
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
 *             .encryptionSpec(AiFeatureStoreEncryptionSpecArgs.builder()
 *                 .kmsKeyName(&#34;kms-name&#34;)
 *                 .build())
 *             .forceDestroy(true)
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .onlineServingConfig(AiFeatureStoreOnlineServingConfigArgs.builder()
 *                 .scaling(AiFeatureStoreOnlineServingConfigScalingArgs.builder()
 *                     .maxNodeCount(10)
 *                     .minNodeCount(2)
 *                     .build())
 *                 .build())
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Featurestore can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default projects/{{project}}/locations/{{region}}/featurestores/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureStore:AiFeatureStore default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureStore:AiFeatureStore")
public class AiFeatureStore extends com.pulumi.resources.CustomResource {
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
     * If set, both of the online and offline data storage will be secured by this key.
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionSpec", refs={AiFeatureStoreEncryptionSpec.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureStoreEncryptionSpec> encryptionSpec;

    /**
     * @return If set, both of the online and offline data storage will be secured by this key.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureStoreEncryptionSpec>> encryptionSpec() {
        return Codegen.optional(this.encryptionSpec);
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
     * If set to true, any EntityTypes and Features for this Featurestore will also be deleted
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return If set to true, any EntityTypes and Features for this Featurestore will also be deleted
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * A set of key/value label pairs to assign to this Featurestore.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this Featurestore.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Config for online serving resources.
     * Structure is documented below.
     * 
     */
    @Export(name="onlineServingConfig", refs={AiFeatureStoreOnlineServingConfig.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureStoreOnlineServingConfig> onlineServingConfig;

    /**
     * @return Config for online serving resources.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureStoreOnlineServingConfig>> onlineServingConfig() {
        return Codegen.optional(this.onlineServingConfig);
    }
    /**
     * TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
     * periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
     * that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
     * featurestore. If not set, default to 4000 days
     * 
     */
    @Export(name="onlineStorageTtlDays", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> onlineStorageTtlDays;

    /**
     * @return TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
     * periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
     * that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
     * featurestore. If not set, default to 4000 days
     * 
     */
    public Output<Optional<Integer>> onlineStorageTtlDays() {
        return Codegen.optional(this.onlineStorageTtlDays);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
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
     * The region of the dataset. eg us-central1
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of the dataset. eg us-central1
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
    public AiFeatureStore(String name) {
        this(name, AiFeatureStoreArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureStore(String name, @Nullable AiFeatureStoreArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureStore(String name, @Nullable AiFeatureStoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStore:AiFeatureStore", name, args == null ? AiFeatureStoreArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureStore(String name, Output<String> id, @Nullable AiFeatureStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStore:AiFeatureStore", name, state, makeResourceOptions(options, id));
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
    public static AiFeatureStore get(String name, Output<String> id, @Nullable AiFeatureStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureStore(name, id, state, options);
    }
}
