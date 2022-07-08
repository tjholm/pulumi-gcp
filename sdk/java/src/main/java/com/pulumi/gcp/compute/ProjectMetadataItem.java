// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.ProjectMetadataItemArgs;
import com.pulumi.gcp.compute.inputs.ProjectMetadataItemState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a single key/value pair on metadata common to all instances for
 * a project in GCE. Using `gcp.compute.ProjectMetadataItem` lets you
 * manage a single key/value setting in the provider rather than the entire
 * project metadata map.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new ProjectMetadataItem(&#34;default&#34;, ProjectMetadataItemArgs.builder()        
 *             .key(&#34;my_metadata&#34;)
 *             .value(&#34;my_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Project metadata items can be imported using the `key`, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:compute/projectMetadataItem:ProjectMetadataItem default my_metadata
 * ```
 * 
 */
@ResourceType(type="gcp:compute/projectMetadataItem:ProjectMetadataItem")
public class ProjectMetadataItem extends com.pulumi.resources.CustomResource {
    /**
     * The metadata key to set.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The metadata key to set.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * The value to set for the given metadata key.
     * 
     */
    @Export(name="value", type=String.class, parameters={})
    private Output<String> value;

    /**
     * @return The value to set for the given metadata key.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectMetadataItem(String name) {
        this(name, ProjectMetadataItemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectMetadataItem(String name, ProjectMetadataItemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectMetadataItem(String name, ProjectMetadataItemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, args == null ? ProjectMetadataItemArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectMetadataItem(String name, Output<String> id, @Nullable ProjectMetadataItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/projectMetadataItem:ProjectMetadataItem", name, state, makeResourceOptions(options, id));
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
    public static ProjectMetadataItem get(String name, Output<String> id, @Nullable ProjectMetadataItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectMetadataItem(name, id, state, options);
    }
}
