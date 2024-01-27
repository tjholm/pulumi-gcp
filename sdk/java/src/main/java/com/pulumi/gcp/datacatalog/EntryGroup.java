// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.datacatalog.EntryGroupArgs;
import com.pulumi.gcp.datacatalog.inputs.EntryGroupState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.
 * 
 * To get more information about EntryGroup, see:
 * 
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 * 
 * ## Example Usage
 * ### Data Catalog Entry Group Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.EntryGroup;
 * import com.pulumi.gcp.datacatalog.EntryGroupArgs;
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
 *         var basicEntryGroup = new EntryGroup(&#34;basicEntryGroup&#34;, EntryGroupArgs.builder()        
 *             .entryGroupId(&#34;my_group&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Data Catalog Entry Group Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.EntryGroup;
 * import com.pulumi.gcp.datacatalog.EntryGroupArgs;
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
 *         var basicEntryGroup = new EntryGroup(&#34;basicEntryGroup&#34;, EntryGroupArgs.builder()        
 *             .description(&#34;example entry group&#34;)
 *             .displayName(&#34;entry group&#34;)
 *             .entryGroupId(&#34;my_group&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EntryGroup can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, EntryGroup can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:datacatalog/entryGroup:EntryGroup default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:datacatalog/entryGroup:EntryGroup")
public class EntryGroup extends com.pulumi.resources.CustomResource {
    /**
     * Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A short name to identify the entry group, for example, &#34;analytics data - jan 2011&#34;.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return A short name to identify the entry group, for example, &#34;analytics data - jan 2011&#34;.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The id of the entry group to create. The id must begin with a letter or underscore,
     * contain only English letters, numbers and underscores, and be at most 64 characters.
     * 
     * ***
     * 
     */
    @Export(name="entryGroupId", refs={String.class}, tree="[0]")
    private Output<String> entryGroupId;

    /**
     * @return The id of the entry group to create. The id must begin with a letter or underscore,
     * contain only English letters, numbers and underscores, and be at most 64 characters.
     * 
     * ***
     * 
     */
    public Output<String> entryGroupId() {
        return this.entryGroupId;
    }
    /**
     * The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
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
     * EntryGroup location region.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return EntryGroup location region.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EntryGroup(String name) {
        this(name, EntryGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EntryGroup(String name, EntryGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EntryGroup(String name, EntryGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/entryGroup:EntryGroup", name, args == null ? EntryGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EntryGroup(String name, Output<String> id, @Nullable EntryGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/entryGroup:EntryGroup", name, state, makeResourceOptions(options, id));
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
    public static EntryGroup get(String name, Output<String> id, @Nullable EntryGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EntryGroup(name, id, state, options);
    }
}
