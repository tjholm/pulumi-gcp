// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.datacatalog.EntryArgs;
import com.pulumi.gcp.datacatalog.inputs.EntryState;
import com.pulumi.gcp.datacatalog.outputs.EntryBigqueryDateShardedSpec;
import com.pulumi.gcp.datacatalog.outputs.EntryBigqueryTableSpec;
import com.pulumi.gcp.datacatalog.outputs.EntryGcsFilesetSpec;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Entry Metadata. A Data Catalog Entry resource represents another resource in Google Cloud Platform
 * (such as a BigQuery dataset or a Pub/Sub topic) or outside of Google Cloud Platform. Clients can use
 * the linkedResource field in the Entry resource to refer to the original resource ID of the source system.
 * 
 * An Entry resource contains resource details, such as its schema. An Entry can also be used to attach
 * flexible metadata, such as a Tag.
 * 
 * To get more information about Entry, see:
 * 
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 * 
 * ## Example Usage
 * 
 * ### Data Catalog Entry Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.EntryGroup;
 * import com.pulumi.gcp.datacatalog.EntryGroupArgs;
 * import com.pulumi.gcp.datacatalog.Entry;
 * import com.pulumi.gcp.datacatalog.EntryArgs;
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
 *         var entryGroup = new EntryGroup("entryGroup", EntryGroupArgs.builder()
 *             .entryGroupId("my_group")
 *             .build());
 * 
 *         var basicEntry = new Entry("basicEntry", EntryArgs.builder()
 *             .entryGroup(entryGroup.id())
 *             .entryId("my_entry")
 *             .userSpecifiedType("my_custom_type")
 *             .userSpecifiedSystem("SomethingExternal")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Catalog Entry Fileset
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.EntryGroup;
 * import com.pulumi.gcp.datacatalog.EntryGroupArgs;
 * import com.pulumi.gcp.datacatalog.Entry;
 * import com.pulumi.gcp.datacatalog.EntryArgs;
 * import com.pulumi.gcp.datacatalog.inputs.EntryGcsFilesetSpecArgs;
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
 *         var entryGroup = new EntryGroup("entryGroup", EntryGroupArgs.builder()
 *             .entryGroupId("my_group")
 *             .build());
 * 
 *         var basicEntry = new Entry("basicEntry", EntryArgs.builder()
 *             .entryGroup(entryGroup.id())
 *             .entryId("my_entry")
 *             .type("FILESET")
 *             .gcsFilesetSpec(EntryGcsFilesetSpecArgs.builder()
 *                 .filePatterns("gs://fake_bucket/dir/*")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Catalog Entry Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.EntryGroup;
 * import com.pulumi.gcp.datacatalog.EntryGroupArgs;
 * import com.pulumi.gcp.datacatalog.Entry;
 * import com.pulumi.gcp.datacatalog.EntryArgs;
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
 *         var entryGroup = new EntryGroup("entryGroup", EntryGroupArgs.builder()
 *             .entryGroupId("my_group")
 *             .build());
 * 
 *         var basicEntry = new Entry("basicEntry", EntryArgs.builder()
 *             .entryGroup(entryGroup.id())
 *             .entryId("my_entry")
 *             .userSpecifiedType("my_user_specified_type")
 *             .userSpecifiedSystem("Something_custom")
 *             .linkedResource("my/linked/resource")
 *             .displayName("my custom type entry")
 *             .description("a custom type entry for a user specified system")
 *             .schema("""
 * {
 *   "columns": [
 *     {
 *       "column": "first_name",
 *       "description": "First name",
 *       "mode": "REQUIRED",
 *       "type": "STRING"
 *     },
 *     {
 *       "column": "last_name",
 *       "description": "Last name",
 *       "mode": "REQUIRED",
 *       "type": "STRING"
 *     },
 *     {
 *       "column": "address",
 *       "description": "Address",
 *       "mode": "REPEATED",
 *       "subcolumns": [
 *         {
 *           "column": "city",
 *           "description": "City",
 *           "mode": "NULLABLE",
 *           "type": "STRING"
 *         },
 *         {
 *           "column": "state",
 *           "description": "State",
 *           "mode": "NULLABLE",
 *           "type": "STRING"
 *         }
 *       ],
 *       "type": "RECORD"
 *     }
 *   ]
 * }
 *             """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Entry can be imported using any of these accepted formats:
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Entry can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:datacatalog/entry:Entry default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:datacatalog/entry:Entry")
public class Entry extends com.pulumi.resources.CustomResource {
    /**
     * Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.
     * Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
     * Structure is documented below.
     * 
     */
    @Export(name="bigqueryDateShardedSpecs", refs={List.class,EntryBigqueryDateShardedSpec.class}, tree="[0,1]")
    private Output<List<EntryBigqueryDateShardedSpec>> bigqueryDateShardedSpecs;

    /**
     * @return Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD.
     * Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
     * Structure is documented below.
     * 
     */
    public Output<List<EntryBigqueryDateShardedSpec>> bigqueryDateShardedSpecs() {
        return this.bigqueryDateShardedSpecs;
    }
    /**
     * Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
     * Structure is documented below.
     * 
     */
    @Export(name="bigqueryTableSpecs", refs={List.class,EntryBigqueryTableSpec.class}, tree="[0,1]")
    private Output<List<EntryBigqueryTableSpec>> bigqueryTableSpecs;

    /**
     * @return Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
     * Structure is documented below.
     * 
     */
    public Output<List<EntryBigqueryTableSpec>> bigqueryTableSpecs() {
        return this.bigqueryTableSpecs;
    }
    /**
     * Entry description, which can consist of several sentences or paragraphs that describe entry contents.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Entry description, which can consist of several sentences or paragraphs that describe entry contents.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Display information such as title and description. A short name to identify the entry,
     * for example, &#34;Analytics Data - Jan 2011&#34;.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Display information such as title and description. A short name to identify the entry,
     * for example, &#34;Analytics Data - Jan 2011&#34;.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The name of the entry group this entry is in.
     * 
     */
    @Export(name="entryGroup", refs={String.class}, tree="[0]")
    private Output<String> entryGroup;

    /**
     * @return The name of the entry group this entry is in.
     * 
     */
    public Output<String> entryGroup() {
        return this.entryGroup;
    }
    /**
     * The id of the entry to create.
     * 
     * ***
     * 
     */
    @Export(name="entryId", refs={String.class}, tree="[0]")
    private Output<String> entryId;

    /**
     * @return The id of the entry to create.
     * 
     * ***
     * 
     */
    public Output<String> entryId() {
        return this.entryId;
    }
    /**
     * Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
     * Structure is documented below.
     * 
     */
    @Export(name="gcsFilesetSpec", refs={EntryGcsFilesetSpec.class}, tree="[0]")
    private Output</* @Nullable */ EntryGcsFilesetSpec> gcsFilesetSpec;

    /**
     * @return Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
     * Structure is documented below.
     * 
     */
    public Output<Optional<EntryGcsFilesetSpec>> gcsFilesetSpec() {
        return Codegen.optional(this.gcsFilesetSpec);
    }
    /**
     * This field indicates the entry&#39;s source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
     * 
     */
    @Export(name="integratedSystem", refs={String.class}, tree="[0]")
    private Output<String> integratedSystem;

    /**
     * @return This field indicates the entry&#39;s source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
     * 
     */
    public Output<String> integratedSystem() {
        return this.integratedSystem;
    }
    /**
     * The resource this metadata entry refers to.
     * For Google Cloud Platform resources, linkedResource is the full name of the resource.
     * For example, the linkedResource for a table resource from BigQuery is:
     * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
     * Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
     * this field is optional and defaults to an empty string.
     * 
     */
    @Export(name="linkedResource", refs={String.class}, tree="[0]")
    private Output<String> linkedResource;

    /**
     * @return The resource this metadata entry refers to.
     * For Google Cloud Platform resources, linkedResource is the full name of the resource.
     * For example, the linkedResource for a table resource from BigQuery is:
     * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
     * Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
     * this field is optional and defaults to an empty string.
     * 
     */
    public Output<String> linkedResource() {
        return this.linkedResource;
    }
    /**
     * The Data Catalog resource name of the entry in URL format.
     * Example: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.
     * Note that this Entry and its child resources may not actually be stored in the location in this name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The Data Catalog resource name of the entry in URL format.
     * Example: projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}.
     * Note that this Entry and its child resources may not actually be stored in the location in this name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
     * attached to it. See
     * https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
     * for what fields this schema can contain.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> schema;

    /**
     * @return Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
     * attached to it. See
     * https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
     * for what fields this schema can contain.
     * 
     */
    public Output<Optional<String>> schema() {
        return Codegen.optional(this.schema);
    }
    /**
     * The type of the entry. Only used for Entries with types in the EntryType enum.
     * Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
     * Possible values are: `FILESET`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of the entry. Only used for Entries with types in the EntryType enum.
     * Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
     * Possible values are: `FILESET`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * This field indicates the entry&#39;s source system that Data Catalog does not integrate with.
     * userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
     * and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     * 
     */
    @Export(name="userSpecifiedSystem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userSpecifiedSystem;

    /**
     * @return This field indicates the entry&#39;s source system that Data Catalog does not integrate with.
     * userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
     * and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     * 
     */
    public Output<Optional<String>> userSpecifiedSystem() {
        return Codegen.optional(this.userSpecifiedSystem);
    }
    /**
     * Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
     * When creating an entry, users should check the enum values first, if nothing matches the entry
     * to be created, then provide a custom value, for example &#34;my_special_type&#34;.
     * userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
     * numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     * 
     */
    @Export(name="userSpecifiedType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userSpecifiedType;

    /**
     * @return Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
     * When creating an entry, users should check the enum values first, if nothing matches the entry
     * to be created, then provide a custom value, for example &#34;my_special_type&#34;.
     * userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
     * numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     * 
     */
    public Output<Optional<String>> userSpecifiedType() {
        return Codegen.optional(this.userSpecifiedType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Entry(String name) {
        this(name, EntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Entry(String name, EntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Entry(String name, EntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/entry:Entry", name, args == null ? EntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Entry(String name, Output<String> id, @Nullable EntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/entry:Entry", name, state, makeResourceOptions(options, id));
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
    public static Entry get(String name, Output<String> id, @Nullable EntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Entry(name, id, state, options);
    }
}
