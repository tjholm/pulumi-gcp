// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.notebooks.LocationArgs;
import com.pulumi.gcp.notebooks.inputs.LocationState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Represents a Location resource.
 * 
 * ## Import
 * 
 * Location can be imported using any of these accepted formats* `projects/{{project}}/locations/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Location can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default projects/{{project}}/locations/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/location:Location default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:notebooks/location:Location")
public class Location extends com.pulumi.resources.CustomResource {
    /**
     * Name of the Location resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Location resource.
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
    public Location(String name) {
        this(name, LocationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Location(String name, @Nullable LocationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Location(String name, @Nullable LocationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:notebooks/location:Location", name, args == null ? LocationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Location(String name, Output<String> id, @Nullable LocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:notebooks/location:Location", name, state, makeResourceOptions(options, id));
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
    public static Location get(String name, Output<String> id, @Nullable LocationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Location(name, id, state, options);
    }
}
