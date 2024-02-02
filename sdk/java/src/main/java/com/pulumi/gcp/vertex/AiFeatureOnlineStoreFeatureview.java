// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureOnlineStoreFeatureviewArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewState;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreFeatureviewBigQuerySource;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreFeatureviewSyncConfig;
import com.pulumi.gcp.vertex.outputs.AiFeatureOnlineStoreFeatureviewVectorSearchConfig;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * FeatureView is representation of values that the FeatureOnlineStore will serve based on its syncConfig.
 * 
 * To get more information about FeatureOnlineStoreFeatureview, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores.featureViews)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
 * 
 * ## Example Usage
 * ### Vertex Ai Featureonlinestore Featureview
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
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreFeatureview;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreFeatureviewArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewSyncConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
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
 *             .build());
 * 
 *         var tf_test_dataset = new Dataset(&#34;tf-test-dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_feature_view&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This is a test description&#34;)
 *             .location(&#34;US&#34;)
 *             .build());
 * 
 *         var tf_test_table = new Table(&#34;tf-test-table&#34;, TableArgs.builder()        
 *             .deletionProtection(false)
 *             .datasetId(tf_test_dataset.datasetId())
 *             .tableId(&#34;example_feature_view&#34;)
 *             .schema(&#34;&#34;&#34;
 *   [
 *   {
 *     &#34;name&#34;: &#34;entity_id&#34;,
 *     &#34;mode&#34;: &#34;NULLABLE&#34;,
 *     &#34;type&#34;: &#34;STRING&#34;,
 *     &#34;description&#34;: &#34;Test default entity_id&#34;
 *   },
 *     {
 *     &#34;name&#34;: &#34;test_entity_column&#34;,
 *     &#34;mode&#34;: &#34;NULLABLE&#34;,
 *     &#34;type&#34;: &#34;STRING&#34;,
 *     &#34;description&#34;: &#34;test secondary entity column&#34;
 *   },
 *   {
 *     &#34;name&#34;: &#34;feature_timestamp&#34;,
 *     &#34;mode&#34;: &#34;NULLABLE&#34;,
 *     &#34;type&#34;: &#34;TIMESTAMP&#34;,
 *     &#34;description&#34;: &#34;Default timestamp value&#34;
 *   }
 * ]
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var featureview = new AiFeatureOnlineStoreFeatureview(&#34;featureview&#34;, AiFeatureOnlineStoreFeatureviewArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .featureOnlineStore(featureonlinestore.name())
 *             .syncConfig(AiFeatureOnlineStoreFeatureviewSyncConfigArgs.builder()
 *                 .cron(&#34;0 0 * * *&#34;)
 *                 .build())
 *             .bigQuerySource(AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs.builder()
 *                 .uri(Output.tuple(tf_test_table.project(), tf_test_table.datasetId(), tf_test_table.tableId()).applyValue(values -&gt; {
 *                     var project = values.t1;
 *                     var datasetId = values.t2;
 *                     var tableId = values.t3;
 *                     return String.format(&#34;bq://%s.%s.%s&#34;, project,datasetId,tableId);
 *                 }))
 *                 .entityIdColumns(&#34;test_entity_column&#34;)
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Featureonlinestore Featureview With Vector Search
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
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreFeatureview;
 * import com.pulumi.gcp.vertex.AiFeatureOnlineStoreFeatureviewArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewSyncConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs;
 * import com.pulumi.gcp.vertex.inputs.AiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigArgs;
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
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var tf_test_dataset = new Dataset(&#34;tf-test-dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_feature_view_vector_search&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This is a test description&#34;)
 *             .location(&#34;US&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var tf_test_table = new Table(&#34;tf-test-table&#34;, TableArgs.builder()        
 *             .deletionProtection(false)
 *             .datasetId(tf_test_dataset.datasetId())
 *             .tableId(&#34;example_feature_view_vector_search&#34;)
 *             .schema(&#34;&#34;&#34;
 * [
 * {
 *   &#34;name&#34;: &#34;test_primary_id&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;STRING&#34;,
 *   &#34;description&#34;: &#34;primary test id&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;embedding&#34;,
 *   &#34;mode&#34;: &#34;REPEATED&#34;,
 *   &#34;type&#34;: &#34;FLOAT&#34;,
 *   &#34;description&#34;: &#34;embedding column for primary_id column&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;country&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;STRING&#34;,
 *   &#34;description&#34;: &#34;country&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;test_crowding_column&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;INTEGER&#34;,
 *   &#34;description&#34;: &#34;test crowding column&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;entity_id&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;STRING&#34;,
 *   &#34;description&#34;: &#34;Test default entity_id&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;test_entity_column&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;STRING&#34;,
 *   &#34;description&#34;: &#34;test secondary entity column&#34;
 * },
 * {
 *   &#34;name&#34;: &#34;feature_timestamp&#34;,
 *   &#34;mode&#34;: &#34;NULLABLE&#34;,
 *   &#34;type&#34;: &#34;TIMESTAMP&#34;,
 *   &#34;description&#34;: &#34;Default timestamp value&#34;
 * }
 * ]
 *             &#34;&#34;&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var featureviewVectorSearch = new AiFeatureOnlineStoreFeatureview(&#34;featureviewVectorSearch&#34;, AiFeatureOnlineStoreFeatureviewArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .featureOnlineStore(featureonlinestore.name())
 *             .syncConfig(AiFeatureOnlineStoreFeatureviewSyncConfigArgs.builder()
 *                 .cron(&#34;0 0 * * *&#34;)
 *                 .build())
 *             .bigQuerySource(AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs.builder()
 *                 .uri(Output.tuple(tf_test_table.project(), tf_test_table.datasetId(), tf_test_table.tableId()).applyValue(values -&gt; {
 *                     var project = values.t1;
 *                     var datasetId = values.t2;
 *                     var tableId = values.t3;
 *                     return String.format(&#34;bq://%s.%s.%s&#34;, project,datasetId,tableId);
 *                 }))
 *                 .entityIdColumns(&#34;test_entity_column&#34;)
 *                 .build())
 *             .vectorSearchConfig(AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs.builder()
 *                 .embeddingColumn(&#34;embedding&#34;)
 *                 .filterColumns(&#34;country&#34;)
 *                 .crowdingColumn(&#34;test_crowding_column&#34;)
 *                 .distanceMeasureType(&#34;DOT_PRODUCT_DISTANCE&#34;)
 *                 .treeAhConfig(AiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigArgs.builder()
 *                     .leafNodeEmbeddingCount(&#34;1000&#34;)
 *                     .build())
 *                 .embeddingDimension(&#34;2&#34;)
 *                 .build())
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
 * FeatureOnlineStoreFeatureview can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}` * `{{project}}/{{region}}/{{feature_online_store}}/{{name}}` * `{{region}}/{{feature_online_store}}/{{name}}` * `{{feature_online_store}}/{{name}}` When using the `pulumi import` command, FeatureOnlineStoreFeatureview can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{project}}/{{region}}/{{feature_online_store}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{region}}/{{feature_online_store}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{feature_online_store}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview")
public class AiFeatureOnlineStoreFeatureview extends com.pulumi.resources.CustomResource {
    /**
     * Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    @Export(name="bigQuerySource", refs={AiFeatureOnlineStoreFeatureviewBigQuerySource.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreFeatureviewBigQuerySource> bigQuerySource;

    /**
     * @return Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreFeatureviewBigQuerySource>> bigQuerySource() {
        return Codegen.optional(this.bigQuerySource);
    }
    /**
     * The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the featureOnlinestore was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
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
     * The name of the FeatureOnlineStore to use for the featureview.
     * 
     */
    @Export(name="featureOnlineStore", refs={String.class}, tree="[0]")
    private Output<String> featureOnlineStore;

    /**
     * @return The name of the FeatureOnlineStore to use for the featureview.
     * 
     */
    public Output<String> featureOnlineStore() {
        return this.featureOnlineStore;
    }
    /**
     * A set of key/value label pairs to assign to this FeatureView.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this FeatureView.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The region for the resource. It should be the same as the featureonlinestore region.
     * 
     * ***
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region for the resource. It should be the same as the featureonlinestore region.
     * 
     * ***
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
     * Structure is documented below.
     * 
     */
    @Export(name="syncConfig", refs={AiFeatureOnlineStoreFeatureviewSyncConfig.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreFeatureviewSyncConfig> syncConfig;

    /**
     * @return Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreFeatureviewSyncConfig>> syncConfig() {
        return Codegen.optional(this.syncConfig);
    }
    /**
     * The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
     * Structure is documented below.
     * 
     */
    @Export(name="vectorSearchConfig", refs={AiFeatureOnlineStoreFeatureviewVectorSearchConfig.class}, tree="[0]")
    private Output</* @Nullable */ AiFeatureOnlineStoreFeatureviewVectorSearchConfig> vectorSearchConfig;

    /**
     * @return Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AiFeatureOnlineStoreFeatureviewVectorSearchConfig>> vectorSearchConfig() {
        return Codegen.optional(this.vectorSearchConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiFeatureOnlineStoreFeatureview(String name) {
        this(name, AiFeatureOnlineStoreFeatureviewArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureOnlineStoreFeatureview(String name, AiFeatureOnlineStoreFeatureviewArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureOnlineStoreFeatureview(String name, AiFeatureOnlineStoreFeatureviewArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview", name, args == null ? AiFeatureOnlineStoreFeatureviewArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureOnlineStoreFeatureview(String name, Output<String> id, @Nullable AiFeatureOnlineStoreFeatureviewState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview", name, state, makeResourceOptions(options, id));
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
    public static AiFeatureOnlineStoreFeatureview get(String name, Output<String> id, @Nullable AiFeatureOnlineStoreFeatureviewState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureOnlineStoreFeatureview(name, id, state, options);
    }
}
