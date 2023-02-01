// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.sql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.sql.SourceRepresentationInstanceArgs;
import com.pulumi.gcp.sql.inputs.SourceRepresentationInstanceState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Sql Source Representation Instance Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.sql.SourceRepresentationInstance;
 * import com.pulumi.gcp.sql.SourceRepresentationInstanceArgs;
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
 *         var instance = new SourceRepresentationInstance(&#34;instance&#34;, SourceRepresentationInstanceArgs.builder()        
 *             .databaseVersion(&#34;MYSQL_8_0&#34;)
 *             .dumpFilePath(&#34;gs://replica-bucket/source-database.sql.gz&#34;)
 *             .host(&#34;10.20.30.40&#34;)
 *             .password(&#34;password-for-the-user&#34;)
 *             .port(3306)
 *             .region(&#34;us-central1&#34;)
 *             .username(&#34;some-user&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SourceRepresentationInstance can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default projects/{{project}}/instances/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance")
public class SourceRepresentationInstance extends com.pulumi.resources.CustomResource {
    /**
     * The CA certificate on the external server. Include only if SSL/TLS is used on the external server.
     * 
     */
    @Export(name="caCertificate", type=String.class, parameters={})
    private Output</* @Nullable */ String> caCertificate;

    /**
     * @return The CA certificate on the external server. Include only if SSL/TLS is used on the external server.
     * 
     */
    public Output<Optional<String>> caCertificate() {
        return Codegen.optional(this.caCertificate);
    }
    /**
     * The client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
     * 
     */
    @Export(name="clientCertificate", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientCertificate;

    /**
     * @return The client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
     * 
     */
    public Output<Optional<String>> clientCertificate() {
        return Codegen.optional(this.clientCertificate);
    }
    /**
     * The private key file for the client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
     * 
     */
    @Export(name="clientKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientKey;

    /**
     * @return The private key file for the client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
     * 
     */
    public Output<Optional<String>> clientKey() {
        return Codegen.optional(this.clientKey);
    }
    /**
     * The MySQL version running on your source database server.
     * Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
     * 
     */
    @Export(name="databaseVersion", type=String.class, parameters={})
    private Output<String> databaseVersion;

    /**
     * @return The MySQL version running on your source database server.
     * Possible values are `MYSQL_5_5`, `MYSQL_5_6`, `MYSQL_5_7`, and `MYSQL_8_0`.
     * 
     */
    public Output<String> databaseVersion() {
        return this.databaseVersion;
    }
    /**
     * A file in the bucket that contains the data from the external server.
     * 
     */
    @Export(name="dumpFilePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> dumpFilePath;

    /**
     * @return A file in the bucket that contains the data from the external server.
     * 
     */
    public Output<Optional<String>> dumpFilePath() {
        return Codegen.optional(this.dumpFilePath);
    }
    /**
     * The externally accessible IPv4 address for the source database server.
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output<String> host;

    /**
     * @return The externally accessible IPv4 address for the source database server.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The name of the source representation instance. Use any valid Cloud SQL instance name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the source representation instance. Use any valid Cloud SQL instance name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password for the replication user account.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return The password for the replication user account.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The externally accessible port for the source database server.
     * Defaults to 3306.
     * 
     */
    @Export(name="port", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> port;

    /**
     * @return The externally accessible port for the source database server.
     * Defaults to 3306.
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
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
     * The Region in which the created instance should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The Region in which the created instance should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The replication user account on the external server.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output</* @Nullable */ String> username;

    /**
     * @return The replication user account on the external server.
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SourceRepresentationInstance(String name) {
        this(name, SourceRepresentationInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SourceRepresentationInstance(String name, SourceRepresentationInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SourceRepresentationInstance(String name, SourceRepresentationInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, args == null ? SourceRepresentationInstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SourceRepresentationInstance(String name, Output<String> id, @Nullable SourceRepresentationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:sql/sourceRepresentationInstance:SourceRepresentationInstance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static SourceRepresentationInstance get(String name, Output<String> id, @Nullable SourceRepresentationInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SourceRepresentationInstance(name, id, state, options);
    }
}
