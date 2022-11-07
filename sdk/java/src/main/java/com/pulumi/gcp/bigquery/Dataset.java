// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.DatasetArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetState;
import com.pulumi.gcp.bigquery.outputs.DatasetAccess;
import com.pulumi.gcp.bigquery.outputs.DatasetDefaultEncryptionConfiguration;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Bigquery Dataset Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceAccount.Account;
 * import com.pulumi.gcp.serviceAccount.AccountArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessArgs;
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
 *         var bqowner = new Account(&#34;bqowner&#34;, AccountArgs.builder()        
 *             .accountId(&#34;bqowner&#34;)
 *             .build());
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_dataset&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This is a test description&#34;)
 *             .location(&#34;EU&#34;)
 *             .defaultTableExpirationMs(3600000)
 *             .labels(Map.of(&#34;env&#34;, &#34;default&#34;))
 *             .accesses(            
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;OWNER&#34;)
 *                     .userByEmail(bqowner.email())
 *                     .build(),
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;READER&#34;)
 *                     .domain(&#34;hashicorp.com&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Bigquery Dataset Cmek
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetDefaultEncryptionConfigurationArgs;
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
 *         var keyRing = new KeyRing(&#34;keyRing&#34;, KeyRingArgs.builder()        
 *             .location(&#34;us&#34;)
 *             .build());
 * 
 *         var cryptoKey = new CryptoKey(&#34;cryptoKey&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(keyRing.id())
 *             .build());
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_dataset&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This is a test description&#34;)
 *             .location(&#34;US&#34;)
 *             .defaultTableExpirationMs(3600000)
 *             .defaultEncryptionConfiguration(DatasetDefaultEncryptionConfigurationArgs.builder()
 *                 .kmsKeyName(cryptoKey.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Bigquery Dataset Authorized Dataset
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceAccount.Account;
 * import com.pulumi.gcp.serviceAccount.AccountArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessDatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessDatasetDatasetArgs;
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
 *         var bqowner = new Account(&#34;bqowner&#34;, AccountArgs.builder()        
 *             .accountId(&#34;bqowner&#34;)
 *             .build());
 * 
 *         var public_ = new Dataset(&#34;public&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;public&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This dataset is public&#34;)
 *             .location(&#34;EU&#34;)
 *             .defaultTableExpirationMs(3600000)
 *             .labels(Map.of(&#34;env&#34;, &#34;default&#34;))
 *             .accesses(            
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;OWNER&#34;)
 *                     .userByEmail(bqowner.email())
 *                     .build(),
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;READER&#34;)
 *                     .domain(&#34;hashicorp.com&#34;)
 *                     .build())
 *             .build());
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;private&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This dataset is private&#34;)
 *             .location(&#34;EU&#34;)
 *             .defaultTableExpirationMs(3600000)
 *             .labels(Map.of(&#34;env&#34;, &#34;default&#34;))
 *             .accesses(            
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;OWNER&#34;)
 *                     .userByEmail(bqowner.email())
 *                     .build(),
 *                 DatasetAccessArgs.builder()
 *                     .role(&#34;READER&#34;)
 *                     .domain(&#34;hashicorp.com&#34;)
 *                     .build(),
 *                 DatasetAccessArgs.builder()
 *                     .dataset(DatasetAccessDatasetArgs.builder()
 *                         .dataset(DatasetAccessDatasetDatasetArgs.builder()
 *                             .projectId(public_.project())
 *                             .datasetId(public_.datasetId())
 *                             .build())
 *                         .targetTypes(&#34;VIEWS&#34;)
 *                         .build())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Dataset can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default projects/{{project}}/datasets/{{dataset_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default {{project}}/{{dataset_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default {{dataset_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigquery/dataset:Dataset")
public class Dataset extends com.pulumi.resources.CustomResource {
    /**
     * An array of objects that define dataset access for one or more entities.
     * Structure is documented below.
     * 
     */
    @Export(name="accesses", type=List.class, parameters={DatasetAccess.class})
    private Output<List<DatasetAccess>> accesses;

    /**
     * @return An array of objects that define dataset access for one or more entities.
     * Structure is documented below.
     * 
     */
    public Output<List<DatasetAccess>> accesses() {
        return this.accesses;
    }
    /**
     * The time when this dataset was created, in milliseconds since the epoch.
     * 
     */
    @Export(name="creationTime", type=Integer.class, parameters={})
    private Output<Integer> creationTime;

    /**
     * @return The time when this dataset was created, in milliseconds since the epoch.
     * 
     */
    public Output<Integer> creationTime() {
        return this.creationTime;
    }
    /**
     * The ID of the dataset containing this table.
     * 
     */
    @Export(name="datasetId", type=String.class, parameters={})
    private Output<String> datasetId;

    /**
     * @return The ID of the dataset containing this table.
     * 
     */
    public Output<String> datasetId() {
        return this.datasetId;
    }
    /**
     * The default encryption key for all tables in the dataset. Once this property is set,
     * all newly-created partitioned tables in the dataset will have encryption key set to
     * this value, unless table creation request (or query) overrides the key.
     * Structure is documented below.
     * 
     */
    @Export(name="defaultEncryptionConfiguration", type=DatasetDefaultEncryptionConfiguration.class, parameters={})
    private Output</* @Nullable */ DatasetDefaultEncryptionConfiguration> defaultEncryptionConfiguration;

    /**
     * @return The default encryption key for all tables in the dataset. Once this property is set,
     * all newly-created partitioned tables in the dataset will have encryption key set to
     * this value, unless table creation request (or query) overrides the key.
     * Structure is documented below.
     * 
     */
    public Output<Optional<DatasetDefaultEncryptionConfiguration>> defaultEncryptionConfiguration() {
        return Codegen.optional(this.defaultEncryptionConfiguration);
    }
    /**
     * The default partition expiration for all partitioned tables in
     * the dataset, in milliseconds.
     * 
     */
    @Export(name="defaultPartitionExpirationMs", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultPartitionExpirationMs;

    /**
     * @return The default partition expiration for all partitioned tables in
     * the dataset, in milliseconds.
     * 
     */
    public Output<Optional<Integer>> defaultPartitionExpirationMs() {
        return Codegen.optional(this.defaultPartitionExpirationMs);
    }
    /**
     * The default lifetime of all tables in the dataset, in milliseconds.
     * The minimum value is 3600000 milliseconds (one hour).
     * 
     */
    @Export(name="defaultTableExpirationMs", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultTableExpirationMs;

    /**
     * @return The default lifetime of all tables in the dataset, in milliseconds.
     * The minimum value is 3600000 milliseconds (one hour).
     * 
     */
    public Output<Optional<Integer>> defaultTableExpirationMs() {
        return Codegen.optional(this.defaultTableExpirationMs);
    }
    /**
     * If set to `true`, delete all the tables in the
     * dataset when destroying the resource; otherwise,
     * destroying the resource will fail if tables are present.
     * 
     */
    @Export(name="deleteContentsOnDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deleteContentsOnDestroy;

    /**
     * @return If set to `true`, delete all the tables in the
     * dataset when destroying the resource; otherwise,
     * destroying the resource will fail if tables are present.
     * 
     */
    public Output<Optional<Boolean>> deleteContentsOnDestroy() {
        return Codegen.optional(this.deleteContentsOnDestroy);
    }
    /**
     * A user-friendly description of the dataset
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A user-friendly description of the dataset
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A hash of the resource.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return A hash of the resource.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * A descriptive name for the dataset
     * 
     */
    @Export(name="friendlyName", type=String.class, parameters={})
    private Output</* @Nullable */ String> friendlyName;

    /**
     * @return A descriptive name for the dataset
     * 
     */
    public Output<Optional<String>> friendlyName() {
        return Codegen.optional(this.friendlyName);
    }
    /**
     * The labels associated with this dataset. You can use these to
     * organize and group your datasets
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return The labels associated with this dataset. You can use these to
     * organize and group your datasets
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
     * 
     */
    @Export(name="lastModifiedTime", type=Integer.class, parameters={})
    private Output<Integer> lastModifiedTime;

    /**
     * @return The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
     * 
     */
    public Output<Integer> lastModifiedTime() {
        return this.lastModifiedTime;
    }
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output</* @Nullable */ String> location;

    /**
     * @return The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
     * 
     */
    @Export(name="maxTimeTravelHours", type=String.class, parameters={})
    private Output</* @Nullable */ String> maxTimeTravelHours;

    /**
     * @return Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
     * 
     */
    public Output<Optional<String>> maxTimeTravelHours() {
        return Codegen.optional(this.maxTimeTravelHours);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", type=String.class, parameters={})
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dataset(String name) {
        this(name, DatasetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dataset(String name, DatasetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dataset(String name, DatasetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/dataset:Dataset", name, args == null ? DatasetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Dataset(String name, Output<String> id, @Nullable DatasetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/dataset:Dataset", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static Dataset get(String name, Output<String> id, @Nullable DatasetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dataset(name, id, state, options);
    }
}
