// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RegionTargetHttpsProxyArgs;
import com.pulumi.gcp.compute.inputs.RegionTargetHttpsProxyState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a RegionTargetHttpsProxy resource, which is used by one or more
 * forwarding rules to route incoming HTTPS requests to a URL map.
 * 
 * To get more information about RegionTargetHttpsProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionTargetHttpsProxies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
 * 
 * ## Example Usage
 * ### Region Target Https Proxy Basic
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.RegionSslCertificate;
 * import com.pulumi.gcp.compute.RegionSslCertificateArgs;
 * import com.pulumi.gcp.compute.RegionHealthCheck;
 * import com.pulumi.gcp.compute.RegionHealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.RegionHealthCheckHttpHealthCheckArgs;
 * import com.pulumi.gcp.compute.RegionBackendService;
 * import com.pulumi.gcp.compute.RegionBackendServiceArgs;
 * import com.pulumi.gcp.compute.RegionUrlMap;
 * import com.pulumi.gcp.compute.RegionUrlMapArgs;
 * import com.pulumi.gcp.compute.inputs.RegionUrlMapHostRuleArgs;
 * import com.pulumi.gcp.compute.inputs.RegionUrlMapPathMatcherArgs;
 * import com.pulumi.gcp.compute.RegionTargetHttpsProxy;
 * import com.pulumi.gcp.compute.RegionTargetHttpsProxyArgs;
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
 *         var defaultRegionSslCertificate = new RegionSslCertificate(&#34;defaultRegionSslCertificate&#34;, RegionSslCertificateArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .privateKey(Files.readString(Paths.get(&#34;path/to/private.key&#34;)))
 *             .certificate(Files.readString(Paths.get(&#34;path/to/certificate.crt&#34;)))
 *             .build());
 * 
 *         var defaultRegionHealthCheck = new RegionHealthCheck(&#34;defaultRegionHealthCheck&#34;, RegionHealthCheckArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .httpHealthCheck(RegionHealthCheckHttpHealthCheckArgs.builder()
 *                 .port(80)
 *                 .build())
 *             .build());
 * 
 *         var defaultRegionBackendService = new RegionBackendService(&#34;defaultRegionBackendService&#34;, RegionBackendServiceArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .protocol(&#34;HTTP&#34;)
 *             .loadBalancingScheme(&#34;INTERNAL_MANAGED&#34;)
 *             .timeoutSec(10)
 *             .healthChecks(defaultRegionHealthCheck.id())
 *             .build());
 * 
 *         var defaultRegionUrlMap = new RegionUrlMap(&#34;defaultRegionUrlMap&#34;, RegionUrlMapArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .description(&#34;a description&#34;)
 *             .defaultService(defaultRegionBackendService.id())
 *             .hostRules(RegionUrlMapHostRuleArgs.builder()
 *                 .hosts(&#34;mysite.com&#34;)
 *                 .pathMatcher(&#34;allpaths&#34;)
 *                 .build())
 *             .pathMatchers(RegionUrlMapPathMatcherArgs.builder()
 *                 .name(&#34;allpaths&#34;)
 *                 .defaultService(defaultRegionBackendService.id())
 *                 .pathRules(RegionUrlMapPathMatcherPathRuleArgs.builder()
 *                     .paths(&#34;/*&#34;)
 *                     .service(defaultRegionBackendService.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var defaultRegionTargetHttpsProxy = new RegionTargetHttpsProxy(&#34;defaultRegionTargetHttpsProxy&#34;, RegionTargetHttpsProxyArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .urlMap(defaultRegionUrlMap.id())
 *             .sslCertificates(defaultRegionSslCertificate.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RegionTargetHttpsProxy can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, RegionTargetHttpsProxy can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy default projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy")
public class RegionTargetHttpsProxy extends com.pulumi.resources.CustomResource {
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
     * The Region in which the created target https proxy should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The Region in which the created target https proxy should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * A list of RegionSslCertificate resources that are used to authenticate
     * connections between users and the load balancer. Currently, exactly
     * one SSL certificate must be specified.
     * 
     */
    @Export(name="sslCertificates", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> sslCertificates;

    /**
     * @return A list of RegionSslCertificate resources that are used to authenticate
     * connections between users and the load balancer. Currently, exactly
     * one SSL certificate must be specified.
     * 
     */
    public Output<List<String>> sslCertificates() {
        return this.sslCertificates;
    }
    /**
     * A reference to the Region SslPolicy resource that will be associated with
     * the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
     * resource will not have any SSL policy configured.
     * 
     */
    @Export(name="sslPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sslPolicy;

    /**
     * @return A reference to the Region SslPolicy resource that will be associated with
     * the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
     * resource will not have any SSL policy configured.
     * 
     */
    public Output<Optional<String>> sslPolicy() {
        return Codegen.optional(this.sslPolicy);
    }
    /**
     * A reference to the RegionUrlMap resource that defines the mapping from URL
     * to the RegionBackendService.
     * 
     * ***
     * 
     */
    @Export(name="urlMap", refs={String.class}, tree="[0]")
    private Output<String> urlMap;

    /**
     * @return A reference to the RegionUrlMap resource that defines the mapping from URL
     * to the RegionBackendService.
     * 
     * ***
     * 
     */
    public Output<String> urlMap() {
        return this.urlMap;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionTargetHttpsProxy(String name) {
        this(name, RegionTargetHttpsProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionTargetHttpsProxy(String name, RegionTargetHttpsProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionTargetHttpsProxy(String name, RegionTargetHttpsProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, args == null ? RegionTargetHttpsProxyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RegionTargetHttpsProxy(String name, Output<String> id, @Nullable RegionTargetHttpsProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, state, makeResourceOptions(options, id));
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
    public static RegionTargetHttpsProxy get(String name, Output<String> id, @Nullable RegionTargetHttpsProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionTargetHttpsProxy(name, id, state, options);
    }
}
