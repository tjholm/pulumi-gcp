// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.datacatalog.TaxonomyArgs;
import com.pulumi.gcp.datacatalog.inputs.TaxonomyState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Data Catalog Taxonomy Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datacatalog.Taxonomy;
 * import com.pulumi.gcp.datacatalog.TaxonomyArgs;
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
 *         var basicTaxonomy = new Taxonomy(&#34;basicTaxonomy&#34;, TaxonomyArgs.builder()        
 *             .region(&#34;us&#34;)
 *             .displayName(&#34;my_display_name&#34;)
 *             .description(&#34;A collection of policy tags&#34;)
 *             .activatedPolicyTypes(&#34;FINE_GRAINED_ACCESS_CONTROL&#34;)
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
 * Taxonomy can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:datacatalog/taxonomy:Taxonomy default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:datacatalog/taxonomy:Taxonomy")
public class Taxonomy extends com.pulumi.resources.CustomResource {
    /**
     * A list of policy types that are activated for this taxonomy. If not set,
     * defaults to an empty list.
     * Each value may be one of `POLICY_TYPE_UNSPECIFIED` and `FINE_GRAINED_ACCESS_CONTROL`.
     * 
     */
    @Export(name="activatedPolicyTypes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> activatedPolicyTypes;

    /**
     * @return A list of policy types that are activated for this taxonomy. If not set,
     * defaults to an empty list.
     * Each value may be one of `POLICY_TYPE_UNSPECIFIED` and `FINE_GRAINED_ACCESS_CONTROL`.
     * 
     */
    public Output<Optional<List<String>>> activatedPolicyTypes() {
        return Codegen.optional(this.activatedPolicyTypes);
    }
    /**
     * Description of this taxonomy. It must: contain only unicode characters,
     * tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
     * long when encoded in UTF-8. If not set, defaults to an empty description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of this taxonomy. It must: contain only unicode characters,
     * tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
     * long when encoded in UTF-8. If not set, defaults to an empty description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User defined name of this taxonomy.
     * It must: contain only unicode letters, numbers, underscores, dashes
     * and spaces; not start or end with spaces; and be at most 200 bytes
     * long when encoded in UTF-8.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return User defined name of this taxonomy.
     * It must: contain only unicode letters, numbers, underscores, dashes
     * and spaces; not start or end with spaces; and be at most 200 bytes
     * long when encoded in UTF-8.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Resource name of this taxonomy, whose format is:
     * &#34;projects/{project}/locations/{region}/taxonomies/{taxonomy}&#34;.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Resource name of this taxonomy, whose format is:
     * &#34;projects/{project}/locations/{region}/taxonomies/{taxonomy}&#34;.
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
     * Taxonomy location region.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return Taxonomy location region.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Taxonomy(String name) {
        this(name, TaxonomyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Taxonomy(String name, TaxonomyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Taxonomy(String name, TaxonomyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/taxonomy:Taxonomy", name, args == null ? TaxonomyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Taxonomy(String name, Output<String> id, @Nullable TaxonomyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:datacatalog/taxonomy:Taxonomy", name, state, makeResourceOptions(options, id));
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
    public static Taxonomy get(String name, Output<String> id, @Nullable TaxonomyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Taxonomy(name, id, state, options);
    }
}
