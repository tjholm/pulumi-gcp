// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.TargetServerArgs;
import com.pulumi.gcp.apigee.inputs.TargetServerState;
import com.pulumi.gcp.apigee.outputs.TargetServerSSlInfo;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * TargetServer configuration. TargetServers are used to decouple a proxy TargetEndpoint HTTPTargetConnections from concrete URLs for backend services.
 * 
 * To get more information about TargetServer, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.targetservers/create)
 * * How-to Guides
 *     * [Load balancing across backend servers](https://cloud.google.com/apigee/docs/api-platform/deploy/load-balancing-across-backend-servers)
 * 
 * ## Example Usage
 * 
 * ### Apigee Target Server Test Basic
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
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.apigee.Organization;
 * import com.pulumi.gcp.apigee.OrganizationArgs;
 * import com.pulumi.gcp.apigee.Environment;
 * import com.pulumi.gcp.apigee.EnvironmentArgs;
 * import com.pulumi.gcp.apigee.TargetServer;
 * import com.pulumi.gcp.apigee.TargetServerArgs;
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
 *         var project = new Project("project", ProjectArgs.builder()
 *             .projectId("my-project")
 *             .name("my-project")
 *             .orgId("123456789")
 *             .billingAccount("000000-0000000-0000000-000000")
 *             .build());
 * 
 *         var apigee = new Service("apigee", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("apigee.googleapis.com")
 *             .build());
 * 
 *         var servicenetworking = new Service("servicenetworking", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("servicenetworking.googleapis.com")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(apigee)
 *                 .build());
 * 
 *         var compute = new Service("compute", ServiceArgs.builder()
 *             .project(project.projectId())
 *             .service("compute.googleapis.com")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(servicenetworking)
 *                 .build());
 * 
 *         var apigeeNetwork = new Network("apigeeNetwork", NetworkArgs.builder()
 *             .name("apigee-network")
 *             .project(project.projectId())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(compute)
 *                 .build());
 * 
 *         var apigeeRange = new GlobalAddress("apigeeRange", GlobalAddressArgs.builder()
 *             .name("apigee-range")
 *             .purpose("VPC_PEERING")
 *             .addressType("INTERNAL")
 *             .prefixLength(16)
 *             .network(apigeeNetwork.id())
 *             .project(project.projectId())
 *             .build());
 * 
 *         var apigeeVpcConnection = new Connection("apigeeVpcConnection", ConnectionArgs.builder()
 *             .network(apigeeNetwork.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(apigeeRange.name())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(servicenetworking)
 *                 .build());
 * 
 *         var apigeeOrg = new Organization("apigeeOrg", OrganizationArgs.builder()
 *             .analyticsRegion("us-central1")
 *             .projectId(project.projectId())
 *             .authorizedNetwork(apigeeNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     apigeeVpcConnection,
 *                     apigee)
 *                 .build());
 * 
 *         var apigeeEnvironment = new Environment("apigeeEnvironment", EnvironmentArgs.builder()
 *             .orgId(apigeeOrg.id())
 *             .name("my-environment-name")
 *             .description("Apigee Environment")
 *             .displayName("environment-1")
 *             .build());
 * 
 *         var apigeeTargetServer = new TargetServer("apigeeTargetServer", TargetServerArgs.builder()
 *             .name("my-target-server")
 *             .description("Apigee Target Server")
 *             .protocol("HTTP")
 *             .host("abc.foo.com")
 *             .port(8080)
 *             .envId(apigeeEnvironment.id())
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
 * TargetServer can be imported using any of these accepted formats:
 * 
 * * `{{env_id}}/targetservers/{{name}}`
 * 
 * * `{{env_id}}/{{name}}`
 * 
 * When using the `pulumi import` command, TargetServer can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/targetservers/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/targetServer:TargetServer")
public class TargetServer extends com.pulumi.resources.CustomResource {
    /**
     * A human-readable description of this TargetServer.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description of this TargetServer.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Apigee environment group associated with the Apigee environment,
     * in the format `organizations/{{org_name}}/environments/{{env_name}}`.
     * 
     * ***
     * 
     */
    @Export(name="envId", refs={String.class}, tree="[0]")
    private Output<String> envId;

    /**
     * @return The Apigee environment group associated with the Apigee environment,
     * in the format `organizations/{{org_name}}/environments/{{env_name}}`.
     * 
     * ***
     * 
     */
    public Output<String> envId() {
        return this.envId;
    }
    /**
     * The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
     * 
     */
    @Export(name="isEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isEnabled;

    /**
     * @return Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> isEnabled() {
        return Codegen.optional(this.isEnabled);
    }
    /**
     * The resource id of this reference. Values must match the regular expression [\w\s-.]+.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource id of this reference. Values must match the regular expression [\w\s-.]+.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * Immutable. The protocol used by this TargetServer.
     * Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return Immutable. The protocol used by this TargetServer.
     * Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="sSlInfo", refs={TargetServerSSlInfo.class}, tree="[0]")
    private Output</* @Nullable */ TargetServerSSlInfo> sSlInfo;

    /**
     * @return Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<TargetServerSSlInfo>> sSlInfo() {
        return Codegen.optional(this.sSlInfo);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TargetServer(java.lang.String name) {
        this(name, TargetServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TargetServer(java.lang.String name, TargetServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TargetServer(java.lang.String name, TargetServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/targetServer:TargetServer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TargetServer(java.lang.String name, Output<java.lang.String> id, @Nullable TargetServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/targetServer:TargetServer", name, state, makeResourceOptions(options, id), false);
    }

    private static TargetServerArgs makeArgs(TargetServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TargetServerArgs.Empty : args;
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
    public static TargetServer get(java.lang.String name, Output<java.lang.String> id, @Nullable TargetServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TargetServer(name, id, state, options);
    }
}
