// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firestore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firestore.IndexArgs;
import com.pulumi.gcp.firestore.inputs.IndexState;
import com.pulumi.gcp.firestore.outputs.IndexField;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Cloud Firestore indexes enable simple and complex queries against documents in a database.
 *  This resource manages composite indexes and not single
 * field indexes.
 * 
 * To get more information about Index, see:
 * 
 * * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/firestore/docs/query-data/indexing)
 * 
 * &gt; **Warning:** This resource creates a Firestore Index on a project that already has
 * a Firestore database. If you haven&#39;t already created it, you may
 * create a `gcp.firestore.Database` resource and `location_id` set
 * to your chosen location. If you wish to use App Engine, you may
 * instead create a `gcp.appengine.Application` resource with
 * `database_type` set to `&#34;CLOUD_FIRESTORE&#34;`. Your Firestore location
 * will be the same as the App Engine location specified.
 * 
 * ## Example Usage
 * ### Firestore Index Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firestore.Index;
 * import com.pulumi.gcp.firestore.IndexArgs;
 * import com.pulumi.gcp.firestore.inputs.IndexFieldArgs;
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
 *         var my_index = new Index(&#34;my-index&#34;, IndexArgs.builder()        
 *             .collection(&#34;chatrooms&#34;)
 *             .fields(            
 *                 IndexFieldArgs.builder()
 *                     .fieldPath(&#34;name&#34;)
 *                     .order(&#34;ASCENDING&#34;)
 *                     .build(),
 *                 IndexFieldArgs.builder()
 *                     .fieldPath(&#34;description&#34;)
 *                     .order(&#34;DESCENDING&#34;)
 *                     .build())
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Firestore Index Datastore Mode
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firestore.Index;
 * import com.pulumi.gcp.firestore.IndexArgs;
 * import com.pulumi.gcp.firestore.inputs.IndexFieldArgs;
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
 *         var my_datastore_mode_index = new Index(&#34;my-datastore-mode-index&#34;, IndexArgs.builder()        
 *             .apiScope(&#34;DATASTORE_MODE_API&#34;)
 *             .collection(&#34;chatrooms&#34;)
 *             .fields(            
 *                 IndexFieldArgs.builder()
 *                     .fieldPath(&#34;name&#34;)
 *                     .order(&#34;ASCENDING&#34;)
 *                     .build(),
 *                 IndexFieldArgs.builder()
 *                     .fieldPath(&#34;description&#34;)
 *                     .order(&#34;DESCENDING&#34;)
 *                     .build())
 *             .project(&#34;my-project-name&#34;)
 *             .queryScope(&#34;COLLECTION_RECURSIVE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Index can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Index can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:firestore/index:Index default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:firestore/index:Index")
public class Index extends com.pulumi.resources.CustomResource {
    /**
     * The API scope at which a query is run.
     * Default value is `ANY_API`.
     * Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
     * 
     */
    @Export(name="apiScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiScope;

    /**
     * @return The API scope at which a query is run.
     * Default value is `ANY_API`.
     * Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
     * 
     */
    public Output<Optional<String>> apiScope() {
        return Codegen.optional(this.apiScope);
    }
    /**
     * The collection being indexed.
     * 
     */
    @Export(name="collection", refs={String.class}, tree="[0]")
    private Output<String> collection;

    /**
     * @return The collection being indexed.
     * 
     */
    public Output<String> collection() {
        return this.collection;
    }
    /**
     * The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    @Export(name="database", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> database;

    /**
     * @return The Firestore database id. Defaults to `&#34;(default)&#34;`.
     * 
     */
    public Output<Optional<String>> database() {
        return Codegen.optional(this.database);
    }
    /**
     * The fields supported by this index. The last field entry is always for
     * the field path `__name__`. If, on creation, `__name__` was not
     * specified as the last field, it will be added automatically with the
     * same direction as that of the last field defined. If the final field
     * in a composite index is not directional, the `__name__` will be
     * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
     * Structure is documented below.
     * 
     */
    @Export(name="fields", refs={List.class,IndexField.class}, tree="[0,1]")
    private Output<List<IndexField>> fields;

    /**
     * @return The fields supported by this index. The last field entry is always for
     * the field path `__name__`. If, on creation, `__name__` was not
     * specified as the last field, it will be added automatically with the
     * same direction as that of the last field defined. If the final field
     * in a composite index is not directional, the `__name__` will be
     * ordered `&#34;ASCENDING&#34;` (unless explicitly specified otherwise).
     * Structure is documented below.
     * 
     */
    public Output<List<IndexField>> fields() {
        return this.fields;
    }
    /**
     * A server defined name for this index. Format:
     * `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A server defined name for this index. Format:
     * `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}`
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
     * The scope at which a query is run.
     * Default value is `COLLECTION`.
     * Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
     * 
     */
    @Export(name="queryScope", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> queryScope;

    /**
     * @return The scope at which a query is run.
     * Default value is `COLLECTION`.
     * Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
     * 
     */
    public Output<Optional<String>> queryScope() {
        return Codegen.optional(this.queryScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Index(String name) {
        this(name, IndexArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Index(String name, IndexArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Index(String name, IndexArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firestore/index:Index", name, args == null ? IndexArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Index(String name, Output<String> id, @Nullable IndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firestore/index:Index", name, state, makeResourceOptions(options, id));
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
    public static Index get(String name, Output<String> id, @Nullable IndexState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Index(name, id, state, options);
    }
}
