// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.TargetGrpcProxyArgs;
import com.pulumi.gcp.compute.inputs.TargetGrpcProxyState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a Target gRPC Proxy resource. A target gRPC proxy is a component
 * of load balancers intended for load balancing gRPC traffic. Global forwarding
 * rules reference a target gRPC proxy. The Target gRPC Proxy references
 * a URL map which specifies how traffic routes to gRPC backend services.
 * 
 * To get more information about TargetGrpcProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetGrpcProxies)
 * * How-to Guides
 *     * [Using Target gRPC Proxies](https://cloud.google.com/traffic-director/docs/proxyless-overview)
 * 
 * ## Example Usage
 * ### Target Grpc Proxy Basic
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.HealthCheck;
 * import com.pulumi.gcp.compute.HealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.HealthCheckGrpcHealthCheckArgs;
 * import com.pulumi.gcp.compute.BackendService;
 * import com.pulumi.gcp.compute.BackendServiceArgs;
 * import com.pulumi.gcp.compute.URLMap;
 * import com.pulumi.gcp.compute.URLMapArgs;
 * import com.pulumi.gcp.compute.inputs.URLMapHostRuleArgs;
 * import com.pulumi.gcp.compute.inputs.URLMapPathMatcherArgs;
 * import com.pulumi.gcp.compute.inputs.URLMapTestArgs;
 * import com.pulumi.gcp.compute.TargetGrpcProxy;
 * import com.pulumi.gcp.compute.TargetGrpcProxyArgs;
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
 *             .grpcHealthCheck(HealthCheckGrpcHealthCheckArgs.builder()
 *                 .portName(&#34;health-check-port&#34;)
 *                 .portSpecification(&#34;USE_NAMED_PORT&#34;)
 *                 .grpcServiceName(&#34;testservice&#34;)
 *                 .build())
 *             .build());
 * 
 *         var home = new BackendService(&#34;home&#34;, BackendServiceArgs.builder()        
 *             .portName(&#34;grpc&#34;)
 *             .protocol(&#34;GRPC&#34;)
 *             .timeoutSec(10)
 *             .healthChecks(defaultHealthCheck.id())
 *             .loadBalancingScheme(&#34;INTERNAL_SELF_MANAGED&#34;)
 *             .build());
 * 
 *         var urlmap = new URLMap(&#34;urlmap&#34;, URLMapArgs.builder()        
 *             .description(&#34;a description&#34;)
 *             .defaultService(home.id())
 *             .hostRules(URLMapHostRuleArgs.builder()
 *                 .hosts(&#34;mysite.com&#34;)
 *                 .pathMatcher(&#34;allpaths&#34;)
 *                 .build())
 *             .pathMatchers(URLMapPathMatcherArgs.builder()
 *                 .name(&#34;allpaths&#34;)
 *                 .defaultService(home.id())
 *                 .routeRules(URLMapPathMatcherRouteRuleArgs.builder()
 *                     .priority(1)
 *                     .headerAction(URLMapPathMatcherRouteRuleHeaderActionArgs.builder()
 *                         .requestHeadersToRemoves(&#34;RemoveMe2&#34;)
 *                         .requestHeadersToAdds(URLMapPathMatcherRouteRuleHeaderActionRequestHeadersToAddArgs.builder()
 *                             .headerName(&#34;AddSomethingElse&#34;)
 *                             .headerValue(&#34;MyOtherValue&#34;)
 *                             .replace(true)
 *                             .build())
 *                         .responseHeadersToRemoves(&#34;RemoveMe3&#34;)
 *                         .responseHeadersToAdds(URLMapPathMatcherRouteRuleHeaderActionResponseHeadersToAddArgs.builder()
 *                             .headerName(&#34;AddMe&#34;)
 *                             .headerValue(&#34;MyValue&#34;)
 *                             .replace(false)
 *                             .build())
 *                         .build())
 *                     .matchRules(URLMapPathMatcherRouteRuleMatchRuleArgs.builder()
 *                         .fullPathMatch(&#34;a full path&#34;)
 *                         .headerMatches(URLMapPathMatcherRouteRuleMatchRuleHeaderMatchArgs.builder()
 *                             .headerName(&#34;someheader&#34;)
 *                             .exactMatch(&#34;match this exactly&#34;)
 *                             .invertMatch(true)
 *                             .build())
 *                         .ignoreCase(true)
 *                         .metadataFilters(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterArgs.builder()
 *                             .filterMatchCriteria(&#34;MATCH_ANY&#34;)
 *                             .filterLabels(URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelArgs.builder()
 *                                 .name(&#34;PLANET&#34;)
 *                                 .value(&#34;MARS&#34;)
 *                                 .build())
 *                             .build())
 *                         .queryParameterMatches(URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatchArgs.builder()
 *                             .name(&#34;a query parameter&#34;)
 *                             .presentMatch(true)
 *                             .build())
 *                         .build())
 *                     .urlRedirect(URLMapPathMatcherRouteRuleUrlRedirectArgs.builder()
 *                         .hostRedirect(&#34;A host&#34;)
 *                         .httpsRedirect(false)
 *                         .pathRedirect(&#34;some/path&#34;)
 *                         .redirectResponseCode(&#34;TEMPORARY_REDIRECT&#34;)
 *                         .stripQuery(true)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .tests(URLMapTestArgs.builder()
 *                 .service(home.id())
 *                 .host(&#34;hi.com&#34;)
 *                 .path(&#34;/home&#34;)
 *                 .build())
 *             .build());
 * 
 *         var defaultTargetGrpcProxy = new TargetGrpcProxy(&#34;defaultTargetGrpcProxy&#34;, TargetGrpcProxyArgs.builder()        
 *             .urlMap(urlmap.id())
 *             .validateForProxyless(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * TargetGrpcProxy can be imported using any of these accepted formats* `projects/{{project}}/global/targetGrpcProxies/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, TargetGrpcProxy can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default projects/{{project}}/global/targetGrpcProxies/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/targetGrpcProxy:TargetGrpcProxy")
public class TargetGrpcProxy extends com.pulumi.resources.CustomResource {
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
     * Fingerprint of this resource. A hash of the contents stored in
     * this object. This field is used in optimistic locking. This field
     * will be ignored when inserting a TargetGrpcProxy. An up-to-date
     * fingerprint must be provided in order to patch/update the
     * TargetGrpcProxy; otherwise, the request will fail with error
     * 412 conditionNotMet. To see the latest fingerprint, make a get()
     * request to retrieve the TargetGrpcProxy. A base64-encoded string.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return Fingerprint of this resource. A hash of the contents stored in
     * this object. This field is used in optimistic locking. This field
     * will be ignored when inserting a TargetGrpcProxy. An up-to-date
     * fingerprint must be provided in order to patch/update the
     * TargetGrpcProxy; otherwise, the request will fail with error
     * 412 conditionNotMet. To see the latest fingerprint, make a get()
     * request to retrieve the TargetGrpcProxy. A base64-encoded string.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * Name of the resource. Provided by the client when the resource
     * is created. The name must be 1-63 characters long, and comply
     * with RFC1035. Specifically, the name must be 1-63 characters long
     * and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. Provided by the client when the resource
     * is created. The name must be 1-63 characters long, and comply
     * with RFC1035. Specifically, the name must be 1-63 characters long
     * and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     * ***
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
     * Server-defined URL with id for the resource.
     * 
     */
    @Export(name="selfLinkWithId", refs={String.class}, tree="[0]")
    private Output<String> selfLinkWithId;

    /**
     * @return Server-defined URL with id for the resource.
     * 
     */
    public Output<String> selfLinkWithId() {
        return this.selfLinkWithId;
    }
    /**
     * URL to the UrlMap resource that defines the mapping from URL to
     * the BackendService. The protocol field in the BackendService
     * must be set to GRPC.
     * 
     */
    @Export(name="urlMap", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> urlMap;

    /**
     * @return URL to the UrlMap resource that defines the mapping from URL to
     * the BackendService. The protocol field in the BackendService
     * must be set to GRPC.
     * 
     */
    public Output<Optional<String>> urlMap() {
        return Codegen.optional(this.urlMap);
    }
    /**
     * If true, indicates that the BackendServices referenced by
     * the urlMap may be accessed by gRPC applications without using
     * a sidecar proxy. This will enable configuration checks on urlMap
     * and its referenced BackendServices to not allow unsupported features.
     * A gRPC application must use &#34;xds:///&#34; scheme in the target URI
     * of the service it is connecting to. If false, indicates that the
     * BackendServices referenced by the urlMap will be accessed by gRPC
     * applications via a sidecar proxy. In this case, a gRPC application
     * must not use &#34;xds:///&#34; scheme in the target URI of the service
     * it is connecting to
     * 
     */
    @Export(name="validateForProxyless", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> validateForProxyless;

    /**
     * @return If true, indicates that the BackendServices referenced by
     * the urlMap may be accessed by gRPC applications without using
     * a sidecar proxy. This will enable configuration checks on urlMap
     * and its referenced BackendServices to not allow unsupported features.
     * A gRPC application must use &#34;xds:///&#34; scheme in the target URI
     * of the service it is connecting to. If false, indicates that the
     * BackendServices referenced by the urlMap will be accessed by gRPC
     * applications via a sidecar proxy. In this case, a gRPC application
     * must not use &#34;xds:///&#34; scheme in the target URI of the service
     * it is connecting to
     * 
     */
    public Output<Optional<Boolean>> validateForProxyless() {
        return Codegen.optional(this.validateForProxyless);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TargetGrpcProxy(String name) {
        this(name, TargetGrpcProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TargetGrpcProxy(String name, @Nullable TargetGrpcProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TargetGrpcProxy(String name, @Nullable TargetGrpcProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, args == null ? TargetGrpcProxyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TargetGrpcProxy(String name, Output<String> id, @Nullable TargetGrpcProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/targetGrpcProxy:TargetGrpcProxy", name, state, makeResourceOptions(options, id));
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
    public static TargetGrpcProxy get(String name, Output<String> id, @Nullable TargetGrpcProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TargetGrpcProxy(name, id, state, options);
    }
}
