// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.TargetTCPProxyArgs;
import com.pulumi.gcp.compute.inputs.TargetTCPProxyState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a TargetTcpProxy resource, which is used by one or more
 * global forwarding rule to route incoming TCP requests to a Backend
 * service.
 * 
 * To get more information about TargetTcpProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetTcpProxies)
 * * How-to Guides
 *     * [Setting Up TCP proxy for Google Cloud Load Balancing](https://cloud.google.com/compute/docs/load-balancing/tcp-ssl/tcp-proxy)
 * 
 * ## Example Usage
 * ### Target Tcp Proxy Basic
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.HealthCheck;
 * import com.pulumi.gcp.compute.HealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.HealthCheckTcpHealthCheckArgs;
 * import com.pulumi.gcp.compute.BackendService;
 * import com.pulumi.gcp.compute.BackendServiceArgs;
 * import com.pulumi.gcp.compute.TargetTCPProxy;
 * import com.pulumi.gcp.compute.TargetTCPProxyArgs;
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
 *         var defaultHealthCheck = new HealthCheck(&#34;defaultHealthCheck&#34;, HealthCheckArgs.builder()        
 *             .timeoutSec(1)
 *             .checkIntervalSec(1)
 *             .tcpHealthCheck(HealthCheckTcpHealthCheckArgs.builder()
 *                 .port(&#34;443&#34;)
 *                 .build())
 *             .build());
 * 
 *         var defaultBackendService = new BackendService(&#34;defaultBackendService&#34;, BackendServiceArgs.builder()        
 *             .protocol(&#34;TCP&#34;)
 *             .timeoutSec(10)
 *             .healthChecks(defaultHealthCheck.id())
 *             .build());
 * 
 *         var defaultTargetTCPProxy = new TargetTCPProxy(&#34;defaultTargetTCPProxy&#34;, TargetTCPProxyArgs.builder()        
 *             .backendService(defaultBackendService.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * TargetTcpProxy can be imported using any of these accepted formats* `projects/{{project}}/global/targetTcpProxies/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, TargetTcpProxy can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default projects/{{project}}/global/targetTcpProxies/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetTCPProxy:TargetTCPProxy default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/targetTCPProxy:TargetTCPProxy")
public class TargetTCPProxy extends com.pulumi.resources.CustomResource {
    /**
     * A reference to the BackendService resource.
     * 
     * ***
     * 
     */
    @Export(name="backendService", refs={String.class}, tree="[0]")
    private Output<String> backendService;

    /**
     * @return A reference to the BackendService resource.
     * 
     * ***
     * 
     */
    public Output<String> backendService() {
        return this.backendService;
    }
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTimestamp", refs={String.class}, tree="[0]")
    private Output<String> creationTimestamp;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Output<String> creationTimestamp() {
        return this.creationTimestamp;
    }
    /**
     * An optional description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
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
     * This field only applies when the forwarding rule that references
     * this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * 
     */
    @Export(name="proxyBind", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> proxyBind;

    /**
     * @return This field only applies when the forwarding rule that references
     * this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * 
     */
    public Output<Boolean> proxyBind() {
        return this.proxyBind;
    }
    /**
     * Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are: `NONE`, `PROXY_V1`.
     * 
     */
    @Export(name="proxyHeader", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxyHeader;

    /**
     * @return Specifies the type of proxy header to append before sending data to
     * the backend.
     * Default value is `NONE`.
     * Possible values are: `NONE`, `PROXY_V1`.
     * 
     */
    public Output<Optional<String>> proxyHeader() {
        return Codegen.optional(this.proxyHeader);
    }
    /**
     * The unique identifier for the resource.
     * 
     */
    @Export(name="proxyId", refs={Integer.class}, tree="[0]")
    private Output<Integer> proxyId;

    /**
     * @return The unique identifier for the resource.
     * 
     */
    public Output<Integer> proxyId() {
        return this.proxyId;
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TargetTCPProxy(String name) {
        this(name, TargetTCPProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TargetTCPProxy(String name, TargetTCPProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TargetTCPProxy(String name, TargetTCPProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetTCPProxy:TargetTCPProxy", name, args == null ? TargetTCPProxyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TargetTCPProxy(String name, Output<String> id, @Nullable TargetTCPProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetTCPProxy:TargetTCPProxy", name, state, makeResourceOptions(options, id));
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
    public static TargetTCPProxy get(String name, Output<String> id, @Nullable TargetTCPProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TargetTCPProxy(name, id, state, options);
    }
}
