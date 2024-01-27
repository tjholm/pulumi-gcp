// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigtable;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigtable.TableArgs;
import com.pulumi.gcp.bigtable.inputs.TableState;
import com.pulumi.gcp.bigtable.outputs.TableColumnFamily;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Google Cloud Bigtable table inside an instance. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigtable.Instance;
 * import com.pulumi.gcp.bigtable.InstanceArgs;
 * import com.pulumi.gcp.bigtable.inputs.InstanceClusterArgs;
 * import com.pulumi.gcp.bigtable.Table;
 * import com.pulumi.gcp.bigtable.TableArgs;
 * import com.pulumi.gcp.bigtable.inputs.TableColumnFamilyArgs;
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
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .clusters(InstanceClusterArgs.builder()
 *                 .clusterId(&#34;tf-instance-cluster&#34;)
 *                 .zone(&#34;us-central1-b&#34;)
 *                 .numNodes(3)
 *                 .storageType(&#34;HDD&#34;)
 *                 .build())
 *             .build());
 * 
 *         var table = new Table(&#34;table&#34;, TableArgs.builder()        
 *             .instanceName(instance.name())
 *             .splitKeys(            
 *                 &#34;a&#34;,
 *                 &#34;b&#34;,
 *                 &#34;c&#34;)
 *             .columnFamilies(            
 *                 TableColumnFamilyArgs.builder()
 *                     .family(&#34;family-first&#34;)
 *                     .build(),
 *                 TableColumnFamilyArgs.builder()
 *                     .family(&#34;family-second&#34;)
 *                     .build())
 *             .changeStreamRetention(&#34;24h0m0s&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * -&gt; **Fields affected by import** The following fields can&#39;t be read and will show diffs if set in config when imported`split_keys` Bigtable Tables can be imported using any of these accepted formats* `projects/{{project}}/instances/{{instance_name}}/tables/{{name}}` * `{{project}}/{{instance_name}}/{{name}}` * `{{instance_name}}/{{name}}` When using the `pulumi import` command, Bigtable Tables can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default {{project}}/{{instance_name}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default {{instance_name}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigtable/table:Table")
public class Table extends com.pulumi.resources.CustomResource {
    /**
     * Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
     * 
     * ***
     * 
     */
    @Export(name="changeStreamRetention", refs={String.class}, tree="[0]")
    private Output<String> changeStreamRetention;

    /**
     * @return Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
     * 
     * ***
     * 
     */
    public Output<String> changeStreamRetention() {
        return this.changeStreamRetention;
    }
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     * 
     */
    @Export(name="columnFamilies", refs={List.class,TableColumnFamily.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TableColumnFamily>> columnFamilies;

    /**
     * @return A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     * 
     */
    public Output<Optional<List<TableColumnFamily>>> columnFamilies() {
        return Codegen.optional(this.columnFamilies);
    }
    /**
     * A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     * 
     */
    @Export(name="deletionProtection", refs={String.class}, tree="[0]")
    private Output<String> deletionProtection;

    /**
     * @return A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     * 
     */
    public Output<String> deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * The name of the Bigtable instance.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return The name of the Bigtable instance.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * A list of predefined keys to split the table on.
     * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     * 
     */
    @Export(name="splitKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> splitKeys;

    /**
     * @return A list of predefined keys to split the table on.
     * !&gt; **Warning:** Modifying the `split_keys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     * 
     */
    public Output<Optional<List<String>>> splitKeys() {
        return Codegen.optional(this.splitKeys);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Table(String name) {
        this(name, TableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Table(String name, TableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Table(String name, TableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/table:Table", name, args == null ? TableArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Table(String name, Output<String> id, @Nullable TableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigtable/table:Table", name, state, makeResourceOptions(options, id));
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
    public static Table get(String name, Output<String> id, @Nullable TableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Table(name, id, state, options);
    }
}
