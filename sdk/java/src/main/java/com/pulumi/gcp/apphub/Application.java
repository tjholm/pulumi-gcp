// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apphub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apphub.ApplicationArgs;
import com.pulumi.gcp.apphub.inputs.ApplicationState;
import com.pulumi.gcp.apphub.outputs.ApplicationAttributes;
import com.pulumi.gcp.apphub.outputs.ApplicationScope;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Application is a functional grouping of Services and Workloads that helps achieve a desired end-to-end business functionality. Services and Workloads are owned by the Application.
 * 
 * ## Example Usage
 * 
 * ### Application Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apphub.Application;
 * import com.pulumi.gcp.apphub.ApplicationArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationScopeArgs;
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
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .location("us-east1")
 *             .applicationId("example-application")
 *             .scope(ApplicationScopeArgs.builder()
 *                 .type("REGIONAL")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Application Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apphub.Application;
 * import com.pulumi.gcp.apphub.ApplicationArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationScopeArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationAttributesArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationAttributesEnvironmentArgs;
 * import com.pulumi.gcp.apphub.inputs.ApplicationAttributesCriticalityArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var example2 = new Application("example2", ApplicationArgs.builder()
 *             .location("us-east1")
 *             .applicationId("example-application")
 *             .displayName("Application Full")
 *             .scope(ApplicationScopeArgs.builder()
 *                 .type("REGIONAL")
 *                 .build())
 *             .description("Application for testing")
 *             .attributes(ApplicationAttributesArgs.builder()
 *                 .environment(ApplicationAttributesEnvironmentArgs.builder()
 *                     .type("STAGING")
 *                     .build())
 *                 .criticality(ApplicationAttributesCriticalityArgs.builder()
 *                     .type("MISSION_CRITICAL")
 *                     .build())
 *                 .businessOwners(ApplicationAttributesBusinessOwnerArgs.builder()
 *                     .displayName("Alice")
 *                     .email("alice}{@literal @}{@code google.com")
 *                     .build())
 *                 .developerOwners(ApplicationAttributesDeveloperOwnerArgs.builder()
 *                     .displayName("Bob")
 *                     .email("bob}{@literal @}{@code google.com")
 *                     .build())
 *                 .operatorOwners(ApplicationAttributesOperatorOwnerArgs.builder()
 *                     .displayName("Charlie")
 *                     .email("charlie}{@literal @}{@code google.com")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Application can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/applications/{{application_id}}`
 * 
 * * `{{project}}/{{location}}/{{application_id}}`
 * 
 * * `{{location}}/{{application_id}}`
 * 
 * When using the `pulumi import` command, Application can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:apphub/application:Application default projects/{{project}}/locations/{{location}}/applications/{{application_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apphub/application:Application default {{project}}/{{location}}/{{application_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apphub/application:Application default {{location}}/{{application_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:apphub/application:Application")
public class Application extends com.pulumi.resources.CustomResource {
    /**
     * Required. The Application identifier.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return Required. The Application identifier.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * Consumer provided attributes.
     * 
     */
    @Export(name="attributes", refs={ApplicationAttributes.class}, tree="[0]")
    private Output</* @Nullable */ ApplicationAttributes> attributes;

    /**
     * @return Consumer provided attributes.
     * 
     */
    public Output<Optional<ApplicationAttributes>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * Output only. Create time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Create time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. User-defined description of an Application.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. User-defined description of an Application.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Optional. User-defined name for the Application.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Optional. User-defined name for the Application.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Part of `parent`. See documentation of `projectsId`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Part of `parent`. See documentation of `projectsId`.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Identifier. The resource name of an Application. Format:
     * &#34;projects/{host-project-id}/locations/{location}/applications/{application-id}&#34;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Identifier. The resource name of an Application. Format:
     * &#34;projects/{host-project-id}/locations/{location}/applications/{application-id}&#34;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * Scope of an application.
     * Structure is documented below.
     * 
     */
    @Export(name="scope", refs={ApplicationScope.class}, tree="[0]")
    private Output<ApplicationScope> scope;

    /**
     * @return Scope of an application.
     * Structure is documented below.
     * 
     */
    public Output<ApplicationScope> scope() {
        return this.scope;
    }
    /**
     * Output only. Application state.
     * Possible values:
     * STATE_UNSPECIFIED
     * CREATING
     * ACTIVE
     * DELETING
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. Application state.
     * Possible values:
     * STATE_UNSPECIFIED
     * CREATING
     * ACTIVE
     * DELETING
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Output only. A universally unique identifier (in UUID4 format) for the `Application`.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. A universally unique identifier (in UUID4 format) for the `Application`.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. Update time.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Update time.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Application(java.lang.String name) {
        this(name, ApplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Application(java.lang.String name, ApplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Application(java.lang.String name, ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apphub/application:Application", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Application(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apphub/application:Application", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationArgs makeArgs(ApplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Application get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Application(name, id, state, options);
    }
}
