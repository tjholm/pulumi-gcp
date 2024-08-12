// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataform.RepositoryWorkflowConfigArgs;
import com.pulumi.gcp.dataform.inputs.RepositoryWorkflowConfigState;
import com.pulumi.gcp.dataform.outputs.RepositoryWorkflowConfigInvocationConfig;
import com.pulumi.gcp.dataform.outputs.RepositoryWorkflowConfigRecentScheduledExecutionRecord;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Dataform Repository Workflow Config
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.sourcerepo.Repository;
 * import com.pulumi.gcp.sourcerepo.RepositoryArgs;
 * import com.pulumi.gcp.secretmanager.Secret;
 * import com.pulumi.gcp.secretmanager.SecretArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationAutoArgs;
 * import com.pulumi.gcp.secretmanager.SecretVersion;
 * import com.pulumi.gcp.secretmanager.SecretVersionArgs;
 * import com.pulumi.gcp.dataform.Repository;
 * import com.pulumi.gcp.dataform.RepositoryArgs;
 * import com.pulumi.gcp.dataform.inputs.RepositoryGitRemoteSettingsArgs;
 * import com.pulumi.gcp.dataform.inputs.RepositoryWorkspaceCompilationOverridesArgs;
 * import com.pulumi.gcp.dataform.RepositoryReleaseConfig;
 * import com.pulumi.gcp.dataform.RepositoryReleaseConfigArgs;
 * import com.pulumi.gcp.dataform.inputs.RepositoryReleaseConfigCodeCompilationConfigArgs;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.dataform.RepositoryWorkflowConfig;
 * import com.pulumi.gcp.dataform.RepositoryWorkflowConfigArgs;
 * import com.pulumi.gcp.dataform.inputs.RepositoryWorkflowConfigInvocationConfigArgs;
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
 *         var gitRepository = new Repository("gitRepository", RepositoryArgs.builder()
 *             .name("my/repository")
 *             .build());
 * 
 *         var secret = new Secret("secret", SecretArgs.builder()
 *             .secretId("my_secret")
 *             .replication(SecretReplicationArgs.builder()
 *                 .auto()
 *                 .build())
 *             .build());
 * 
 *         var secretVersion = new SecretVersion("secretVersion", SecretVersionArgs.builder()
 *             .secret(secret.id())
 *             .secretData("secret-data")
 *             .build());
 * 
 *         var repository = new Repository("repository", RepositoryArgs.builder()
 *             .name("dataform_repository")
 *             .region("us-central1")
 *             .gitRemoteSettings(RepositoryGitRemoteSettingsArgs.builder()
 *                 .url(gitRepository.url())
 *                 .defaultBranch("main")
 *                 .authenticationTokenSecretVersion(secretVersion.id())
 *                 .build())
 *             .workspaceCompilationOverrides(RepositoryWorkspaceCompilationOverridesArgs.builder()
 *                 .defaultDatabase("database")
 *                 .schemaSuffix("_suffix")
 *                 .tablePrefix("prefix_")
 *                 .build())
 *             .build());
 * 
 *         var releaseConfig = new RepositoryReleaseConfig("releaseConfig", RepositoryReleaseConfigArgs.builder()
 *             .project(repository.project())
 *             .region(repository.region())
 *             .repository(repository.name())
 *             .name("my_release")
 *             .gitCommitish("main")
 *             .cronSchedule("0 7 * * *")
 *             .timeZone("America/New_York")
 *             .codeCompilationConfig(RepositoryReleaseConfigCodeCompilationConfigArgs.builder()
 *                 .defaultDatabase("gcp-example-project")
 *                 .defaultSchema("example-dataset")
 *                 .defaultLocation("us-central1")
 *                 .assertionSchema("example-assertion-dataset")
 *                 .databaseSuffix("")
 *                 .schemaSuffix("")
 *                 .tablePrefix("")
 *                 .vars(Map.of("var1", "value"))
 *                 .build())
 *             .build());
 * 
 *         var dataformSa = new Account("dataformSa", AccountArgs.builder()
 *             .accountId("dataform-sa")
 *             .displayName("Dataform Service Account")
 *             .build());
 * 
 *         var workflow = new RepositoryWorkflowConfig("workflow", RepositoryWorkflowConfigArgs.builder()
 *             .project(repository.project())
 *             .region(repository.region())
 *             .repository(repository.name())
 *             .name("my_workflow")
 *             .releaseConfig(releaseConfig.id())
 *             .invocationConfig(RepositoryWorkflowConfigInvocationConfigArgs.builder()
 *                 .includedTargets(                
 *                     RepositoryWorkflowConfigInvocationConfigIncludedTargetArgs.builder()
 *                         .database("gcp-example-project")
 *                         .schema("example-dataset")
 *                         .name("target_1")
 *                         .build(),
 *                     RepositoryWorkflowConfigInvocationConfigIncludedTargetArgs.builder()
 *                         .database("gcp-example-project")
 *                         .schema("example-dataset")
 *                         .name("target_2")
 *                         .build())
 *                 .includedTags("tag_1")
 *                 .transitiveDependenciesIncluded(true)
 *                 .transitiveDependentsIncluded(true)
 *                 .fullyRefreshIncrementalTablesEnabled(false)
 *                 .serviceAccount(dataformSa.email())
 *                 .build())
 *             .cronSchedule("0 7 * * *")
 *             .timeZone("America/New_York")
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
 * RepositoryWorkflowConfig can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{repository}}/{{name}}`
 * 
 * * `{{region}}/{{repository}}/{{name}}`
 * 
 * * `{{repository}}/{{name}}`
 * 
 * When using the `pulumi import` command, RepositoryWorkflowConfig can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{project}}/{{region}}/{{repository}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{region}}/{{repository}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{repository}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig")
public class RepositoryWorkflowConfig extends com.pulumi.resources.CustomResource {
    /**
     * Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     * 
     */
    @Export(name="cronSchedule", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cronSchedule;

    /**
     * @return Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     * 
     */
    public Output<Optional<String>> cronSchedule() {
        return Codegen.optional(this.cronSchedule);
    }
    /**
     * Optional. If left unset, a default InvocationConfig will be used.
     * Structure is documented below.
     * 
     */
    @Export(name="invocationConfig", refs={RepositoryWorkflowConfigInvocationConfig.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryWorkflowConfigInvocationConfig> invocationConfig;

    /**
     * @return Optional. If left unset, a default InvocationConfig will be used.
     * Structure is documented below.
     * 
     */
    public Output<Optional<RepositoryWorkflowConfigInvocationConfig>> invocationConfig() {
        return Codegen.optional(this.invocationConfig);
    }
    /**
     * The workflow&#39;s name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The workflow&#39;s name.
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
     * Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
     * Structure is documented below.
     * 
     */
    @Export(name="recentScheduledExecutionRecords", refs={List.class,RepositoryWorkflowConfigRecentScheduledExecutionRecord.class}, tree="[0,1]")
    private Output<List<RepositoryWorkflowConfigRecentScheduledExecutionRecord>> recentScheduledExecutionRecords;

    /**
     * @return Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
     * Structure is documented below.
     * 
     */
    public Output<List<RepositoryWorkflowConfigRecentScheduledExecutionRecord>> recentScheduledExecutionRecords() {
        return this.recentScheduledExecutionRecords;
    }
    /**
     * A reference to the region
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return A reference to the region
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
     * 
     * ***
     * 
     */
    @Export(name="releaseConfig", refs={String.class}, tree="[0]")
    private Output<String> releaseConfig;

    /**
     * @return The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
     * 
     * ***
     * 
     */
    public Output<String> releaseConfig() {
        return this.releaseConfig;
    }
    /**
     * A reference to the Dataform repository
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> repository;

    /**
     * @return A reference to the Dataform repository
     * 
     */
    public Output<Optional<String>> repository() {
        return Codegen.optional(this.repository);
    }
    /**
     * Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     * 
     */
    @Export(name="timeZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timeZone;

    /**
     * @return Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     * 
     */
    public Output<Optional<String>> timeZone() {
        return Codegen.optional(this.timeZone);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryWorkflowConfig(java.lang.String name) {
        this(name, RepositoryWorkflowConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryWorkflowConfig(java.lang.String name, RepositoryWorkflowConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryWorkflowConfig(java.lang.String name, RepositoryWorkflowConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RepositoryWorkflowConfig(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryWorkflowConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryWorkflowConfigArgs makeArgs(RepositoryWorkflowConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryWorkflowConfigArgs.Empty : args;
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
    public static RepositoryWorkflowConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryWorkflowConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryWorkflowConfig(name, id, state, options);
    }
}
