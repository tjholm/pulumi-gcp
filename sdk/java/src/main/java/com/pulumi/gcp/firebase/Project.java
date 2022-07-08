// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firebase.ProjectArgs;
import com.pulumi.gcp.firebase.inputs.ProjectState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * A Google Cloud Firebase instance. This enables Firebase resources on a given google project.
 * Since a FirebaseProject is actually also a GCP Project, a FirebaseProject uses underlying GCP
 * identifiers (most importantly, the projectId) as its own for easy interop with GCP APIs.
 * 
 * Once Firebase has been added to a Google Project it cannot be removed.
 * 
 * To get more information about Project, see:
 * 
 * * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects)
 * * How-to Guides
 *     * [Official Documentation](https://firebase.google.com/)
 * 
 * ## Example Usage
 * ### Firebase Project Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var defaultProject = new Project(&#34;defaultProject&#34;, ProjectArgs.builder()        
 *             .projectId(&#34;tf-test&#34;)
 *             .orgId(&#34;123456789&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var defaultFirebase_projectProject = new Project(&#34;defaultFirebase/projectProject&#34;, ProjectArgs.builder()        
 *             .project(defaultProject.projectId())
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
 * Project can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/project:Project default projects/{{project}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/project:Project default {{project}}
 * ```
 * 
 */
@ResourceType(type="gcp:firebase/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    /**
     * The GCP project display name
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The GCP project display name
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
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
     * The number of the google project that firebase is enabled on.
     * 
     */
    @Export(name="projectNumber", type=String.class, parameters={})
    private Output<String> projectNumber;

    /**
     * @return The number of the google project that firebase is enabled on.
     * 
     */
    public Output<String> projectNumber() {
        return this.projectNumber;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, @Nullable ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, @Nullable ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/project:Project", name, state, makeResourceOptions(options, id));
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
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
