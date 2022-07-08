// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.ProjectDefaultNetworkTierArgs;
import com.pulumi.gcp.compute.inputs.ProjectDefaultNetworkTierState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Configures the Google Compute Engine
 * [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
 * for a project.
 * 
 * For more information, see,
 * [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).
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
 *         var default_ = new ProjectDefaultNetworkTier(&#34;default&#34;, ProjectDefaultNetworkTierArgs.builder()        
 *             .networkTier(&#34;PREMIUM&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using the project ID
 * 
 * ```sh
 *  $ pulumi import gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier default project-id`
 * ```
 * 
 */
@ResourceType(type="gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier")
public class ProjectDefaultNetworkTier extends com.pulumi.resources.CustomResource {
    /**
     * The default network tier to be configured for the project.
     * This field can take the following values: `PREMIUM` or `STANDARD`.
     * 
     */
    @Export(name="networkTier", type=String.class, parameters={})
    private Output<String> networkTier;

    /**
     * @return The default network tier to be configured for the project.
     * This field can take the following values: `PREMIUM` or `STANDARD`.
     * 
     */
    public Output<String> networkTier() {
        return this.networkTier;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectDefaultNetworkTier(String name) {
        this(name, ProjectDefaultNetworkTierArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectDefaultNetworkTier(String name, ProjectDefaultNetworkTierArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectDefaultNetworkTier(String name, ProjectDefaultNetworkTierArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, args == null ? ProjectDefaultNetworkTierArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectDefaultNetworkTier(String name, Output<String> id, @Nullable ProjectDefaultNetworkTierState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, state, makeResourceOptions(options, id));
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
    public static ProjectDefaultNetworkTier get(String name, Output<String> id, @Nullable ProjectDefaultNetworkTierState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectDefaultNetworkTier(name, id, state, options);
    }
}
