// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureOnlineStoreArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreState;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreBigtable;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreDedicatedServingEndpoint;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreEmbeddingManagement;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreOptimized;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Vertex AI Feature Online Store provides a centralized repository for serving ML features and embedding indexes at low latency. The Feature Online Store is a top-level container.
 * 
 * To get more information about FeatureOnlineStore, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 * 
 * ## Example Usage
 * ### Vertex Ai Feature Online Store
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStore;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreBigtableArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreBigtableAutoScalingArgs;
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
 *         var featureOnlineStore = new AiFeatureOnlineStore(&#34;featureOnlineStore&#34;, AiFeatureOnlineStoreArgs.builder()        
 *             .bigtable(AiFeatureOnlineStoreBigtableArgs.builder()
 *                 .autoScaling(AiFeatureOnlineStoreBigtableAutoScalingArgs.builder()
 *                     .cpuUtilizationTarget(50)
 *                     .maxNodeCount(3)
 *                     .minNodeCount(1)
 *                     .build())
 *                 .build())
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featureonlinestore With Beta Fields Optimized
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStore;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreOptimizedArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreDedicatedServingEndpointArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigArgs;
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
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var featureonlinestore = new AiFeatureOnlineStore(&#34;featureonlinestore&#34;, AiFeatureOnlineStoreArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .region(&#34;us-central1&#34;)
 *             .optimized()
 *             .dedicatedServingEndpoint(AiFeatureOnlineStoreDedicatedServingEndpointArgs.builder()
 *                 .privateServiceConnectConfig(AiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigArgs.builder()
 *                     .enablePrivateServiceConnect(true)
 *                     .projectAllowlists(project.applyValue(getProjectResult -&gt; getProjectResult.number()))
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featureonlinestore With Beta Fields Bigtable
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStore;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreBigtableArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreBigtableAutoScalingArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreEmbeddingManagementArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
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
 *         var featureonlinestore = new AiFeatureOnlineStore(&#34;featureonlinestore&#34;, AiFeatureOnlineStoreArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .region(&#34;us-central1&#34;)
 *             .bigtable(AiFeatureOnlineStoreBigtableArgs.builder()
 *                 .autoScaling(AiFeatureOnlineStoreBigtableAutoScalingArgs.builder()
 *                     .minNodeCount(1)
 *                     .maxNodeCount(2)
 *                     .cpuUtilizationTarget(80)
 *                     .build())
 *                 .build())
 *             .embeddingManagement(AiFeatureOnlineStoreEmbeddingManagementArgs.builder()
 *                 .enabled(true)
 *                 .build())
 *             .forceDestroy(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * FeatureOnlineStore can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, FeatureOnlineStore can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore")
public class AiFeatureOnlineStore extends com.pulumi.resources.CustomResource {
    /**
     * Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    @Export(name="bigtable", refs={AiFeatureOnlineStoreBigtable.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreBigtable> bigtable;

    /**
     * @return Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreBigtable>> bigtable() {
        return Codegen.optional(this.bigtable);
    }
    /**
     * The timestamp of when the feature online store was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the feature online store was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
     * Structure is documented below.
     * 
     */
    @Export(name="dedicatedServingEndpoint", refs={AiFeatureOnlineStoreDedicatedServingEndpoint.class}, tree="[0]")
    private Output<AiFeatureOnlineStoreDedicatedServingEndpoint> dedicatedServingEndpoint;

    /**
     * @return The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
     * Structure is documented below.
     * 
     */
    public Output<AiFeatureOnlineStoreDedicatedServingEndpoint> dedicatedServingEndpoint() {
        return this.dedicatedServingEndpoint;
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
     * The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
     * Structure is documented below.
     * 
     */
    @Export(name="embeddingManagement", refs={AiFeatureOnlineStoreEmbeddingManagement.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreEmbeddingManagement> embeddingManagement;

    /**
     * @return The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreEmbeddingManagement>> embeddingManagement() {
        return Codegen.optional(this.embeddingManagement);
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
     * If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The labels with user-defined metadata to organize your feature online stores.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return The labels with user-defined metadata to organize your feature online stores.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
     * 
     */
    @Export(name="optimized", refs={AiFeatureOnlineStoreOptimized.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreOptimized> optimized;

    /**
     * @return Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreOptimized>> optimized() {
        return Codegen.optional(this.optimized);
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
     * The region of feature online store. eg us-central1
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of feature online store. eg us-central1
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The timestamp of when the feature online store was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp of when the feature online store was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiFeatureOnlineStore(String name) {
        this(name, AiFeatureOnlineStoreArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureOnlineStore(String name, @Nullable AiFeatureOnlineStoreArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureOnlineStore(String name, @Nullable AiFeatureOnlineStoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore", name, args == null ? AiFeatureOnlineStoreArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureOnlineStore(String name, Output<String> id, @Nullable AiFeatureOnlineStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore", name, state, makeResourceOptions(options, id));
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
    public static AiFeatureOnlineStore get(String name, Output<String> id, @Nullable AiFeatureOnlineStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureOnlineStore(name, id, state, options);
    }
}
