// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.identityplatform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.identityplatform.ProjectDefaultConfigArgs;
import com.pulumi.gcp.identityplatform.inputs.ProjectDefaultConfigState;
import com.pulumi.gcp.identityplatform.outputs.ProjectDefaultConfigSignIn;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **Warning:** `gcp.identityplatform.Config` is deprecated and will be removed in the next major release of the provider. Use the `gcp.identityplatform.Config` resource instead. It contains a more comprehensive list of fields, and was created before `gcp.identityplatform.ProjectDefaultConfig` was added.
 * 
 * There is no persistent data associated with this resource.
 * 
 * &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
 * you must specify a `billing_project` and set `user_project_override` to true
 * in the provider configuration. Otherwise the ACM API will return a 403 error.
 * Your account must have the `serviceusage.services.use` permission on the
 * `billing_project` you defined.
 * 
 * ## Example Usage
 * ### Identity Platform Project Default Config
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.identityplatform.ProjectDefaultConfig;
 * import com.pulumi.gcp.identityplatform.ProjectDefaultConfigArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ProjectDefaultConfigSignInArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ProjectDefaultConfigSignInAnonymousArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ProjectDefaultConfigSignInEmailArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ProjectDefaultConfigSignInPhoneNumberArgs;
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
 *         var default_ = new ProjectDefaultConfig(&#34;default&#34;, ProjectDefaultConfigArgs.builder()        
 *             .signIn(ProjectDefaultConfigSignInArgs.builder()
 *                 .allowDuplicateEmails(true)
 *                 .anonymous(ProjectDefaultConfigSignInAnonymousArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .email(ProjectDefaultConfigSignInEmailArgs.builder()
 *                     .enabled(true)
 *                     .passwordRequired(false)
 *                     .build())
 *                 .phoneNumber(ProjectDefaultConfigSignInPhoneNumberArgs.builder()
 *                     .enabled(true)
 *                     .testPhoneNumbers(Map.of(&#34;+11231231234&#34;, &#34;000000&#34;))
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ProjectDefaultConfig can be imported using any of these accepted formats* `projects/{{project}}/config/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, ProjectDefaultConfig can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig default projects/{{project}}/config/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig")
public class ProjectDefaultConfig extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Config resource. Example: &#34;projects/my-awesome-project/config&#34;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Config resource. Example: &#34;projects/my-awesome-project/config&#34;
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
     * Configuration related to local sign in methods.
     * Structure is documented below.
     * 
     */
    @Export(name="signIn", refs={ProjectDefaultConfigSignIn.class}, tree="[0]")
    private Output</* @Nullable */ ProjectDefaultConfigSignIn> signIn;

    /**
     * @return Configuration related to local sign in methods.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ProjectDefaultConfigSignIn>> signIn() {
        return Codegen.optional(this.signIn);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectDefaultConfig(String name) {
        this(name, ProjectDefaultConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectDefaultConfig(String name, @Nullable ProjectDefaultConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectDefaultConfig(String name, @Nullable ProjectDefaultConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig", name, args == null ? ProjectDefaultConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectDefaultConfig(String name, Output<String> id, @Nullable ProjectDefaultConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:identityplatform/projectDefaultConfig:ProjectDefaultConfig", name, state, makeResourceOptions(options, id));
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
    public static ProjectDefaultConfig get(String name, Output<String> id, @Nullable ProjectDefaultConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectDefaultConfig(name, id, state, options);
    }
}
