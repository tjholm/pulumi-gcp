// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.identityplatform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.identityplatform.ConfigArgs;
import com.pulumi.gcp.identityplatform.inputs.ConfigState;
import com.pulumi.gcp.identityplatform.outputs.ConfigBlockingFunctions;
import com.pulumi.gcp.identityplatform.outputs.ConfigClient;
import com.pulumi.gcp.identityplatform.outputs.ConfigMfa;
import com.pulumi.gcp.identityplatform.outputs.ConfigMonitoring;
import com.pulumi.gcp.identityplatform.outputs.ConfigMultiTenant;
import com.pulumi.gcp.identityplatform.outputs.ConfigQuota;
import com.pulumi.gcp.identityplatform.outputs.ConfigSignIn;
import com.pulumi.gcp.identityplatform.outputs.ConfigSmsRegionConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Identity Platform configuration for a Cloud project. Identity Platform is an
 * end-to-end authentication system for third-party users to access apps
 * and services.
 * 
 * This entity is created only once during intialization and cannot be deleted,
 * individual Identity Providers may be disabled instead.  This resource may only
 * be created in billing-enabled projects.
 * 
 * To get more information about Config, see:
 * 
 * * [API documentation](https://cloud.google.com/identity-platform/docs/reference/rest/v2/Config)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/identity-platform/docs)
 * 
 * ## Example Usage
 * 
 * ### Identity Platform Config Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.identityplatform.Config;
 * import com.pulumi.gcp.identityplatform.ConfigArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSignInArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSignInAnonymousArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSignInEmailArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSignInPhoneNumberArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSmsRegionConfigArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigSmsRegionConfigAllowlistOnlyArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigBlockingFunctionsArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigBlockingFunctionsForwardInboundCredentialsArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigQuotaArgs;
 * import com.pulumi.gcp.identityplatform.inputs.ConfigQuotaSignUpQuotaConfigArgs;
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
 *         var default_ = new Project("default", ProjectArgs.builder()
 *             .projectId("my-project")
 *             .name("my-project")
 *             .orgId("123456789")
 *             .billingAccount("000000-0000000-0000000-000000")
 *             .labels(Map.of("firebase", "enabled"))
 *             .build());
 * 
 *         var identitytoolkit = new Service("identitytoolkit", ServiceArgs.builder()
 *             .project(default_.projectId())
 *             .service("identitytoolkit.googleapis.com")
 *             .build());
 * 
 *         var defaultConfig = new Config("defaultConfig", ConfigArgs.builder()
 *             .project(default_.projectId())
 *             .autodeleteAnonymousUsers(true)
 *             .signIn(ConfigSignInArgs.builder()
 *                 .allowDuplicateEmails(true)
 *                 .anonymous(ConfigSignInAnonymousArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .email(ConfigSignInEmailArgs.builder()
 *                     .enabled(true)
 *                     .passwordRequired(false)
 *                     .build())
 *                 .phoneNumber(ConfigSignInPhoneNumberArgs.builder()
 *                     .enabled(true)
 *                     .testPhoneNumbers(Map.of("+11231231234", "000000"))
 *                     .build())
 *                 .build())
 *             .smsRegionConfig(ConfigSmsRegionConfigArgs.builder()
 *                 .allowlistOnly(ConfigSmsRegionConfigAllowlistOnlyArgs.builder()
 *                     .allowedRegions(                    
 *                         "US",
 *                         "CA")
 *                     .build())
 *                 .build())
 *             .blockingFunctions(ConfigBlockingFunctionsArgs.builder()
 *                 .triggers(ConfigBlockingFunctionsTriggerArgs.builder()
 *                     .eventType("beforeSignIn")
 *                     .functionUri("https://us-east1-my-project.cloudfunctions.net/before-sign-in")
 *                     .build())
 *                 .forwardInboundCredentials(ConfigBlockingFunctionsForwardInboundCredentialsArgs.builder()
 *                     .refreshToken(true)
 *                     .accessToken(true)
 *                     .idToken(true)
 *                     .build())
 *                 .build())
 *             .quota(ConfigQuotaArgs.builder()
 *                 .signUpQuotaConfig(ConfigQuotaSignUpQuotaConfigArgs.builder()
 *                     .quota(1000)
 *                     .startTime("")
 *                     .quotaDuration("7200s")
 *                     .build())
 *                 .build())
 *             .authorizedDomains(            
 *                 "localhost",
 *                 "my-project.firebaseapp.com",
 *                 "my-project.web.app")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Config can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/config`
 * 
 * * `projects/{{project}}`
 * 
 * * `{{project}}`
 * 
 * When using the `pulumi import` command, Config can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:identityplatform/config:Config default projects/{{project}}/config
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:identityplatform/config:Config default projects/{{project}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:identityplatform/config:Config default {{project}}
 * ```
 * 
 */
@ResourceType(type="gcp:identityplatform/config:Config")
public class Config extends com.pulumi.resources.CustomResource {
    /**
     * List of domains authorized for OAuth redirects.
     * 
     */
    @Export(name="authorizedDomains", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> authorizedDomains;

    /**
     * @return List of domains authorized for OAuth redirects.
     * 
     */
    public Output<List<String>> authorizedDomains() {
        return this.authorizedDomains;
    }
    /**
     * Whether anonymous users will be auto-deleted after a period of 30 days
     * 
     */
    @Export(name="autodeleteAnonymousUsers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autodeleteAnonymousUsers;

    /**
     * @return Whether anonymous users will be auto-deleted after a period of 30 days
     * 
     */
    public Output<Optional<Boolean>> autodeleteAnonymousUsers() {
        return Codegen.optional(this.autodeleteAnonymousUsers);
    }
    /**
     * Configuration related to blocking functions.
     * Structure is documented below.
     * 
     */
    @Export(name="blockingFunctions", refs={ConfigBlockingFunctions.class}, tree="[0]")
    private Output</* @Nullable */ ConfigBlockingFunctions> blockingFunctions;

    /**
     * @return Configuration related to blocking functions.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ConfigBlockingFunctions>> blockingFunctions() {
        return Codegen.optional(this.blockingFunctions);
    }
    /**
     * Options related to how clients making requests on behalf of a project should be configured.
     * Structure is documented below.
     * 
     */
    @Export(name="client", refs={ConfigClient.class}, tree="[0]")
    private Output<ConfigClient> client;

    /**
     * @return Options related to how clients making requests on behalf of a project should be configured.
     * Structure is documented below.
     * 
     */
    public Output<ConfigClient> client() {
        return this.client;
    }
    /**
     * Options related to how clients making requests on behalf of a project should be configured.
     * Structure is documented below.
     * 
     */
    @Export(name="mfa", refs={ConfigMfa.class}, tree="[0]")
    private Output<ConfigMfa> mfa;

    /**
     * @return Options related to how clients making requests on behalf of a project should be configured.
     * Structure is documented below.
     * 
     */
    public Output<ConfigMfa> mfa() {
        return this.mfa;
    }
    /**
     * Configuration related to monitoring project activity.
     * Structure is documented below.
     * 
     */
    @Export(name="monitoring", refs={ConfigMonitoring.class}, tree="[0]")
    private Output<ConfigMonitoring> monitoring;

    /**
     * @return Configuration related to monitoring project activity.
     * Structure is documented below.
     * 
     */
    public Output<ConfigMonitoring> monitoring() {
        return this.monitoring;
    }
    /**
     * Configuration related to multi-tenant functionality.
     * Structure is documented below.
     * 
     */
    @Export(name="multiTenant", refs={ConfigMultiTenant.class}, tree="[0]")
    private Output</* @Nullable */ ConfigMultiTenant> multiTenant;

    /**
     * @return Configuration related to multi-tenant functionality.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ConfigMultiTenant>> multiTenant() {
        return Codegen.optional(this.multiTenant);
    }
    /**
     * The name of the Config resource
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Config resource
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
     * Configuration related to quotas.
     * Structure is documented below.
     * 
     */
    @Export(name="quota", refs={ConfigQuota.class}, tree="[0]")
    private Output</* @Nullable */ ConfigQuota> quota;

    /**
     * @return Configuration related to quotas.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ConfigQuota>> quota() {
        return Codegen.optional(this.quota);
    }
    /**
     * Configuration related to local sign in methods.
     * Structure is documented below.
     * 
     */
    @Export(name="signIn", refs={ConfigSignIn.class}, tree="[0]")
    private Output<ConfigSignIn> signIn;

    /**
     * @return Configuration related to local sign in methods.
     * Structure is documented below.
     * 
     */
    public Output<ConfigSignIn> signIn() {
        return this.signIn;
    }
    /**
     * Configures the regions where users are allowed to send verification SMS for the project or tenant. This is based on the calling code of the destination phone number.
     * Structure is documented below.
     * 
     */
    @Export(name="smsRegionConfig", refs={ConfigSmsRegionConfig.class}, tree="[0]")
    private Output<ConfigSmsRegionConfig> smsRegionConfig;

    /**
     * @return Configures the regions where users are allowed to send verification SMS for the project or tenant. This is based on the calling code of the destination phone number.
     * Structure is documented below.
     * 
     */
    public Output<ConfigSmsRegionConfig> smsRegionConfig() {
        return this.smsRegionConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Config(java.lang.String name) {
        this(name, ConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Config(java.lang.String name, @Nullable ConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Config(java.lang.String name, @Nullable ConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:identityplatform/config:Config", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Config(java.lang.String name, Output<java.lang.String> id, @Nullable ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:identityplatform/config:Config", name, state, makeResourceOptions(options, id), false);
    }

    private static ConfigArgs makeArgs(@Nullable ConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ConfigArgs.Empty : args;
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
    public static Config get(java.lang.String name, Output<java.lang.String> id, @Nullable ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Config(name, id, state, options);
    }
}
