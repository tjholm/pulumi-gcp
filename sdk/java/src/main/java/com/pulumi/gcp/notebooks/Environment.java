// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.notebooks.EnvironmentArgs;
import com.pulumi.gcp.notebooks.inputs.EnvironmentState;
import com.pulumi.gcp.notebooks.outputs.EnvironmentContainerImage;
import com.pulumi.gcp.notebooks.outputs.EnvironmentVmImage;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Cloud AI Platform Notebook environment.
 * 
 * To get more information about Environment, see:
 * 
 * * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
 * 
 * ## Example Usage
 * ### Notebook Environment Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.notebooks.Environment;
 * import com.pulumi.gcp.notebooks.EnvironmentArgs;
 * import com.pulumi.gcp.notebooks.inputs.EnvironmentContainerImageArgs;
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
 *         var environment = new Environment(&#34;environment&#34;, EnvironmentArgs.builder()        
 *             .containerImage(EnvironmentContainerImageArgs.builder()
 *                 .repository(&#34;gcr.io/deeplearning-platform-release/base-cpu&#34;)
 *                 .build())
 *             .location(&#34;us-west1-a&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Environment can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/environment:Environment default projects/{{project}}/locations/{{location}}/environments/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/environment:Environment default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:notebooks/environment:Environment default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:notebooks/environment:Environment")
public class Environment extends com.pulumi.resources.CustomResource {
    /**
     * Use a container image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    @Export(name="containerImage", type=EnvironmentContainerImage.class, parameters={})
    private Output</* @Nullable */ EnvironmentContainerImage> containerImage;

    /**
     * @return Use a container image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<EnvironmentContainerImage>> containerImage() {
        return Codegen.optional(this.containerImage);
    }
    /**
     * Instance creation time
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Instance creation time
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A brief description of this environment.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of this environment.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Display name of this environment for the UI.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Display name of this environment for the UI.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * A reference to the zone where the machine resides.
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return A reference to the zone where the machine resides.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The name specified for the Environment instance.
     * Format: projects/{project_id}/locations/{location}/environments/{environmentId}
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name specified for the Environment instance.
     * Format: projects/{project_id}/locations/{location}/environments/{environmentId}
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Path to a Bash script that automatically runs after a notebook instance fully boots up.
     * The path must be a URL or Cloud Storage path. Example: &#34;gs://path-to-file/file-name&#34;
     * 
     */
    @Export(name="postStartupScript", type=String.class, parameters={})
    private Output</* @Nullable */ String> postStartupScript;

    /**
     * @return Path to a Bash script that automatically runs after a notebook instance fully boots up.
     * The path must be a URL or Cloud Storage path. Example: &#34;gs://path-to-file/file-name&#34;
     * 
     */
    public Output<Optional<String>> postStartupScript() {
        return Codegen.optional(this.postStartupScript);
    }
    /**
     * The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The name of the Google Cloud project that this VM image belongs to.
     * Format: projects/{project_id}
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    @Export(name="vmImage", type=EnvironmentVmImage.class, parameters={})
    private Output</* @Nullable */ EnvironmentVmImage> vmImage;

    /**
     * @return Use a Compute Engine VM image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<EnvironmentVmImage>> vmImage() {
        return Codegen.optional(this.vmImage);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Environment(String name) {
        this(name, EnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Environment(String name, EnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Environment(String name, EnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:notebooks/environment:Environment", name, args == null ? EnvironmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Environment(String name, Output<String> id, @Nullable EnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:notebooks/environment:Environment", name, state, makeResourceOptions(options, id));
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
    public static Environment get(String name, Output<String> id, @Nullable EnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Environment(name, id, state, options);
    }
}
