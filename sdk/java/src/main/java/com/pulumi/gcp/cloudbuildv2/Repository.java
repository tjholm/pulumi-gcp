// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuildv2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudbuildv2.RepositoryArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.RepositoryState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A repository associated to a parent connection.
 * 
 * To get more information about Repository, see:
 * 
 * * [API documentation](https://cloud.google.com/build/docs/api/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/build/docs)
 * 
 * ## Example Usage
 * 
 * ### Cloudbuildv2 Repository Ghe Doc
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.secretmanager.Secret;
 * import com.pulumi.gcp.secretmanager.SecretArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationAutoArgs;
 * import com.pulumi.gcp.secretmanager.SecretVersion;
 * import com.pulumi.gcp.secretmanager.SecretVersionArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.secretmanager.SecretIamPolicy;
 * import com.pulumi.gcp.secretmanager.SecretIamPolicyArgs;
 * import com.pulumi.gcp.cloudbuildv2.Connection;
 * import com.pulumi.gcp.cloudbuildv2.ConnectionArgs;
 * import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionGithubEnterpriseConfigArgs;
 * import com.pulumi.gcp.cloudbuildv2.Repository;
 * import com.pulumi.gcp.cloudbuildv2.RepositoryArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var private_key_secret = new Secret("private-key-secret", SecretArgs.builder()
 *             .secretId("ghe-pk-secret")
 *             .replication(SecretReplicationArgs.builder()
 *                 .auto()
 *                 .build())
 *             .build());
 * 
 *         var private_key_secret_version = new SecretVersion("private-key-secret-version", SecretVersionArgs.builder()
 *             .secret(private_key_secret.id())
 *             .secretData(StdFunctions.file(FileArgs.builder()
 *                 .input("private-key.pem")
 *                 .build()).result())
 *             .build());
 * 
 *         var webhook_secret_secret = new Secret("webhook-secret-secret", SecretArgs.builder()
 *             .secretId("github-token-secret")
 *             .replication(SecretReplicationArgs.builder()
 *                 .auto()
 *                 .build())
 *             .build());
 * 
 *         var webhook_secret_secret_version = new SecretVersion("webhook-secret-secret-version", SecretVersionArgs.builder()
 *             .secret(webhook_secret_secret.id())
 *             .secretData("<webhook-secret-data>")
 *             .build());
 * 
 *         final var p4sa-secretAccessor = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/secretmanager.secretAccessor")
 *                 .members("serviceAccount:service-123456789}{@literal @}{@code gcp-sa-cloudbuild.iam.gserviceaccount.com")
 *                 .build())
 *             .build());
 * 
 *         var policy_pk = new SecretIamPolicy("policy-pk", SecretIamPolicyArgs.builder()
 *             .secretId(private_key_secret.secretId())
 *             .policyData(p4sa_secretAccessor.policyData())
 *             .build());
 * 
 *         var policy_whs = new SecretIamPolicy("policy-whs", SecretIamPolicyArgs.builder()
 *             .secretId(webhook_secret_secret.secretId())
 *             .policyData(p4sa_secretAccessor.policyData())
 *             .build());
 * 
 *         var my_connection = new Connection("my-connection", ConnectionArgs.builder()
 *             .location("us-central1")
 *             .name("my-terraform-ghe-connection")
 *             .githubEnterpriseConfig(ConnectionGithubEnterpriseConfigArgs.builder()
 *                 .hostUri("https://ghe.com")
 *                 .privateKeySecretVersion(private_key_secret_version.id())
 *                 .webhookSecretSecretVersion(webhook_secret_secret_version.id())
 *                 .appId(200)
 *                 .appSlug("gcb-app")
 *                 .appInstallationId(300)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     policy_pk,
 *                     policy_whs)
 *                 .build());
 * 
 *         var my_repository = new Repository("my-repository", RepositoryArgs.builder()
 *             .name("my-terraform-ghe-repo")
 *             .location("us-central1")
 *             .parentConnection(my_connection.name())
 *             .remoteUri("https://ghe.com/hashicorp/terraform-provider-google.git")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Cloudbuildv2 Repository Github Doc
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.secretmanager.Secret;
 * import com.pulumi.gcp.secretmanager.SecretArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationAutoArgs;
 * import com.pulumi.gcp.secretmanager.SecretVersion;
 * import com.pulumi.gcp.secretmanager.SecretVersionArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.secretmanager.SecretIamPolicy;
 * import com.pulumi.gcp.secretmanager.SecretIamPolicyArgs;
 * import com.pulumi.gcp.cloudbuildv2.Connection;
 * import com.pulumi.gcp.cloudbuildv2.ConnectionArgs;
 * import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionGithubConfigArgs;
 * import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionGithubConfigAuthorizerCredentialArgs;
 * import com.pulumi.gcp.cloudbuildv2.Repository;
 * import com.pulumi.gcp.cloudbuildv2.RepositoryArgs;
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
 *         var github_token_secret = new Secret("github-token-secret", SecretArgs.builder()
 *             .secretId("github-token-secret")
 *             .replication(SecretReplicationArgs.builder()
 *                 .auto()
 *                 .build())
 *             .build());
 * 
 *         var github_token_secret_version = new SecretVersion("github-token-secret-version", SecretVersionArgs.builder()
 *             .secret(github_token_secret.id())
 *             .secretData(StdFunctions.file(FileArgs.builder()
 *                 .input("my-github-token.txt")
 *                 .build()).result())
 *             .build());
 * 
 *         final var p4sa-secretAccessor = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/secretmanager.secretAccessor")
 *                 .members("serviceAccount:service-123456789}{@literal @}{@code gcp-sa-cloudbuild.iam.gserviceaccount.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new SecretIamPolicy("policy", SecretIamPolicyArgs.builder()
 *             .secretId(github_token_secret.secretId())
 *             .policyData(p4sa_secretAccessor.policyData())
 *             .build());
 * 
 *         var my_connection = new Connection("my-connection", ConnectionArgs.builder()
 *             .location("us-central1")
 *             .name("my-connection")
 *             .githubConfig(ConnectionGithubConfigArgs.builder()
 *                 .appInstallationId(123123)
 *                 .authorizerCredential(ConnectionGithubConfigAuthorizerCredentialArgs.builder()
 *                     .oauthTokenSecretVersion(github_token_secret_version.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var my_repository = new Repository("my-repository", RepositoryArgs.builder()
 *             .location("us-central1")
 *             .name("my-repo")
 *             .parentConnection(my_connection.name())
 *             .remoteUri("https://github.com/myuser/myrepo.git")
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
 * Repository can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/repositories/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{parent_connection}}/{{name}}`
 * 
 * * `{{location}}/{{parent_connection}}/{{name}}`
 * 
 * When using the `pulumi import` command, Repository can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:cloudbuildv2/repository:Repository default projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/repositories/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:cloudbuildv2/repository:Repository default {{project}}/{{location}}/{{parent_connection}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:cloudbuildv2/repository:Repository default {{location}}/{{parent_connection}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudbuildv2/repository:Repository")
public class Repository extends com.pulumi.resources.CustomResource {
    /**
     * Allows clients to store small amounts of arbitrary data.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Allows clients to store small amounts of arbitrary data.
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
    }
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Name of the repository.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the repository.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The connection for the resource
     * 
     * ***
     * 
     */
    @Export(name="parentConnection", refs={String.class}, tree="[0]")
    private Output<String> parentConnection;

    /**
     * @return The connection for the resource
     * 
     * ***
     * 
     */
    public Output<String> parentConnection() {
        return this.parentConnection;
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
     * Required. Git Clone HTTPS URI.
     * 
     */
    @Export(name="remoteUri", refs={String.class}, tree="[0]")
    private Output<String> remoteUri;

    /**
     * @return Required. Git Clone HTTPS URI.
     * 
     */
    public Output<String> remoteUri() {
        return this.remoteUri;
    }
    /**
     * Output only. Server assigned timestamp for when the connection was updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Server assigned timestamp for when the connection was updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Repository(java.lang.String name) {
        this(name, RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Repository(java.lang.String name, RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Repository(java.lang.String name, RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudbuildv2/repository:Repository", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Repository(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudbuildv2/repository:Repository", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryArgs makeArgs(RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryArgs.Empty : args;
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
    public static Repository get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Repository(name, id, state, options);
    }
}
