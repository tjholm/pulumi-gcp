// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuildv2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudbuildv2.ConnectionArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionState;
import com.pulumi.gcp.cloudbuildv2.outputs.ConnectionGithubConfig;
import com.pulumi.gcp.cloudbuildv2.outputs.ConnectionGithubEnterpriseConfig;
import com.pulumi.gcp.cloudbuildv2.outputs.ConnectionInstallationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Beta only: The Cloudbuildv2 Connection resource
 * 
 * ## Example Usage
 * ### GitHub Connection
 * Creates a Connection to github.com
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.secretmanager.Secret;
 * import com.pulumi.gcp.secretmanager.SecretArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationArgs;
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
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var github_token_secret = new Secret(&#34;github-token-secret&#34;, SecretArgs.builder()        
 *             .secretId(&#34;github-token-secret&#34;)
 *             .replication(SecretReplicationArgs.builder()
 *                 .automatic(true)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var github_token_secret_version = new SecretVersion(&#34;github-token-secret-version&#34;, SecretVersionArgs.builder()        
 *             .secret(github_token_secret.id())
 *             .secretData(Files.readString(Paths.get(&#34;my-github-token.txt&#34;)))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         final var p4sa-secretAccessor = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role(&#34;roles/secretmanager.secretAccessor&#34;)
 *                 .members(&#34;serviceAccount:service-123456789@gcp-sa-cloudbuild.iam.gserviceaccount.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new SecretIamPolicy(&#34;policy&#34;, SecretIamPolicyArgs.builder()        
 *             .secretId(github_token_secret.secretId())
 *             .policyData(p4sa_secretAccessor.policyData())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var my_connection = new Connection(&#34;my-connection&#34;, ConnectionArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .githubConfig(ConnectionGithubConfigArgs.builder()
 *                 .appInstallationId(123123)
 *                 .authorizerCredential(ConnectionGithubConfigAuthorizerCredentialArgs.builder()
 *                     .oauthTokenSecretVersion(github_token_secret_version.id())
 *                     .build())
 *                 .build())
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
 * Connection can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:cloudbuildv2/connection:Connection default projects/{{project}}/locations/{{location}}/connections/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudbuildv2/connection:Connection default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudbuildv2/connection:Connection default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudbuildv2/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    /**
     * Allows clients to store small amounts of arbitrary data.
     * 
     */
    @Export(name="annotations", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Allows clients to store small amounts of arbitrary data.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     * 
     */
    @Export(name="disabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Configuration for connections to github.com.
     * 
     */
    @Export(name="githubConfig", type=ConnectionGithubConfig.class, parameters={})
    private Output</* @Nullable */ ConnectionGithubConfig> githubConfig;

    /**
     * @return Configuration for connections to github.com.
     * 
     */
    public Output<Optional<ConnectionGithubConfig>> githubConfig() {
        return Codegen.optional(this.githubConfig);
    }
    /**
     * Configuration for connections to an instance of GitHub Enterprise.
     * 
     */
    @Export(name="githubEnterpriseConfig", type=ConnectionGithubEnterpriseConfig.class, parameters={})
    private Output</* @Nullable */ ConnectionGithubEnterpriseConfig> githubEnterpriseConfig;

    /**
     * @return Configuration for connections to an instance of GitHub Enterprise.
     * 
     */
    public Output<Optional<ConnectionGithubEnterpriseConfig>> githubEnterpriseConfig() {
        return Codegen.optional(this.githubEnterpriseConfig);
    }
    /**
     * Output only. Installation state of the Connection.
     * 
     */
    @Export(name="installationStates", type=List.class, parameters={ConnectionInstallationState.class})
    private Output<List<ConnectionInstallationState>> installationStates;

    /**
     * @return Output only. Installation state of the Connection.
     * 
     */
    public Output<List<ConnectionInstallationState>> installationStates() {
        return this.installationStates;
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project for the resource
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Output only. Set to true when the connection is being set up or updated in the background.
     * 
     */
    @Export(name="reconciling", type=Boolean.class, parameters={})
    private Output<Boolean> reconciling;

    /**
     * @return Output only. Set to true when the connection is being set up or updated in the background.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * Output only. Server assigned timestamp for when the connection was updated.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
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
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudbuildv2/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudbuildv2/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
