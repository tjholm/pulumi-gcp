// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.PublicDelegatedPrefixArgs;
import com.pulumi.gcp.compute.inputs.PublicDelegatedPrefixState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a PublicDelegatedPrefix for use with bring your own IP addresses (BYOIP).
 * 
 * To get more information about PublicDelegatedPrefix, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/publicDelegatedPrefixes)
 * * How-to Guides
 *     * [Using bring your own IP](https://cloud.google.com/vpc/docs/using-bring-your-own-ip)
 * 
 * ## Example Usage
 * ### Public Delegated Prefixes Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.PublicAdvertisedPrefix;
 * import com.pulumi.gcp.compute.PublicAdvertisedPrefixArgs;
 * import com.pulumi.gcp.compute.PublicDelegatedPrefix;
 * import com.pulumi.gcp.compute.PublicDelegatedPrefixArgs;
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
 *         var advertised = new PublicAdvertisedPrefix(&#34;advertised&#34;, PublicAdvertisedPrefixArgs.builder()        
 *             .description(&#34;description&#34;)
 *             .dnsVerificationIp(&#34;127.127.0.0&#34;)
 *             .ipCidrRange(&#34;127.127.0.0/16&#34;)
 *             .build());
 * 
 *         var prefixes = new PublicDelegatedPrefix(&#34;prefixes&#34;, PublicDelegatedPrefixArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .description(&#34;my description&#34;)
 *             .ipCidrRange(&#34;127.127.0.0/24&#34;)
 *             .parentPrefix(advertised.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * PublicDelegatedPrefix can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, PublicDelegatedPrefix can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix default projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix")
public class PublicDelegatedPrefix extends com.pulumi.resources.CustomResource {
    /**
     * An optional description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The IPv4 address range, in CIDR format, represented by this public advertised prefix.
     * 
     * ***
     * 
     */
    @Export(name="ipCidrRange", refs={String.class}, tree="[0]")
    private Output<String> ipCidrRange;

    /**
     * @return The IPv4 address range, in CIDR format, represented by this public advertised prefix.
     * 
     * ***
     * 
     */
    public Output<String> ipCidrRange() {
        return this.ipCidrRange;
    }
    /**
     * If true, the prefix will be live migrated.
     * 
     */
    @Export(name="isLiveMigration", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isLiveMigration;

    /**
     * @return If true, the prefix will be live migrated.
     * 
     */
    public Output<Optional<Boolean>> isLiveMigration() {
        return Codegen.optional(this.isLiveMigration);
    }
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
     * 
     */
    @Export(name="parentPrefix", refs={String.class}, tree="[0]")
    private Output<String> parentPrefix;

    /**
     * @return The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
     * 
     */
    public Output<String> parentPrefix() {
        return this.parentPrefix;
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
     * A region where the prefix will reside.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return A region where the prefix will reside.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
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
    public PublicDelegatedPrefix(String name) {
        this(name, PublicDelegatedPrefixArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PublicDelegatedPrefix(String name, PublicDelegatedPrefixArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PublicDelegatedPrefix(String name, PublicDelegatedPrefixArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix", name, args == null ? PublicDelegatedPrefixArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PublicDelegatedPrefix(String name, Output<String> id, @Nullable PublicDelegatedPrefixState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/publicDelegatedPrefix:PublicDelegatedPrefix", name, state, makeResourceOptions(options, id));
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
    public static PublicDelegatedPrefix get(String name, Output<String> id, @Nullable PublicDelegatedPrefixState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PublicDelegatedPrefix(name, id, state, options);
    }
}
