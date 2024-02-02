// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Backend Service defines a group of virtual machines that will serve
 * traffic for load balancing. This resource is a global backend service,
 * appropriate for external load balancing or self-managed internal load balancing.
 * For managed internal load balancing, use a regional backend service instead.
 *
 * Currently self-managed internal load balancing is only available in beta.
 *
 * To get more information about BackendService, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendServices)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
 *
 * ## Example Usage
 * ### Backend Service External Iap
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.BackendService("default", {
 *     iap: {
 *         oauth2ClientId: "abc",
 *         oauth2ClientSecret: "xyz",
 *     },
 *     loadBalancingScheme: "EXTERNAL",
 *     protocol: "HTTP",
 * });
 * ```
 * ### Backend Service Cache Include Http Headers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.BackendService("default", {
 *     cdnPolicy: {
 *         cacheKeyPolicy: {
 *             includeHost: true,
 *             includeHttpHeaders: ["X-My-Header-Field"],
 *             includeProtocol: true,
 *             includeQueryString: true,
 *         },
 *         cacheMode: "USE_ORIGIN_HEADERS",
 *     },
 *     enableCdn: true,
 * });
 * ```
 * ### Backend Service Cache Include Named Cookies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.BackendService("default", {
 *     cdnPolicy: {
 *         cacheKeyPolicy: {
 *             includeHost: true,
 *             includeNamedCookies: [
 *                 "__next_preview_data",
 *                 "__prerender_bypass",
 *             ],
 *             includeProtocol: true,
 *             includeQueryString: true,
 *         },
 *         cacheMode: "CACHE_ALL_STATIC",
 *         clientTtl: 7200,
 *         defaultTtl: 3600,
 *         maxTtl: 10800,
 *     },
 *     enableCdn: true,
 * });
 * ```
 * ### Backend Service Network Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const externalProxy = new gcp.compute.GlobalNetworkEndpointGroup("externalProxy", {
 *     networkEndpointType: "INTERNET_FQDN_PORT",
 *     defaultPort: 443,
 * }, {
 *     provider: google_beta,
 * });
 * const proxy = new gcp.compute.GlobalNetworkEndpoint("proxy", {
 *     globalNetworkEndpointGroup: externalProxy.id,
 *     fqdn: "test.example.com",
 *     port: externalProxy.defaultPort,
 * }, {
 *     provider: google_beta,
 * });
 * const _default = new gcp.compute.BackendService("default", {
 *     enableCdn: true,
 *     timeoutSec: 10,
 *     connectionDrainingTimeoutSec: 10,
 *     customRequestHeaders: [proxy.fqdn.apply(fqdn => `host: ${fqdn}`)],
 *     customResponseHeaders: ["X-Cache-Hit: {cdn_cache_status}"],
 *     backends: [{
 *         group: externalProxy.id,
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * BackendService can be imported using any of these accepted formats* `projects/{{project}}/global/backendServices/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, BackendService can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:compute/backendService:BackendService default projects/{{project}}/global/backendServices/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/backendService:BackendService default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/backendService:BackendService default {{name}}
 * ```
 */
export class BackendService extends pulumi.CustomResource {
    /**
     * Get an existing BackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendServiceState, opts?: pulumi.CustomResourceOptions): BackendService {
        return new BackendService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/backendService:BackendService';

    /**
     * Returns true if the given object is an instance of BackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendService.__pulumiType;
    }

    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    public readonly affinityCookieTtlSec!: pulumi.Output<number | undefined>;
    /**
     * The set of backends that serve this BackendService.
     * Structure is documented below.
     */
    public readonly backends!: pulumi.Output<outputs.compute.BackendServiceBackend[] | undefined>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.BackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    public readonly circuitBreakers!: pulumi.Output<outputs.compute.BackendServiceCircuitBreakers | undefined>;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     * Possible values are: `AUTOMATIC`, `DISABLED`.
     */
    public readonly compressionMode!: pulumi.Output<string | undefined>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    public readonly connectionDrainingTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing. This field only applies if the loadBalancingScheme is set to
     * INTERNAL_SELF_MANAGED. This field is only applicable when localityLbPolicy is
     * set to MAGLEV or RING_HASH.
     * Structure is documented below.
     */
    public readonly consistentHash!: pulumi.Output<outputs.compute.BackendServiceConsistentHash | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * requests.
     */
    public readonly customRequestHeaders!: pulumi.Output<string[] | undefined>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * responses.
     */
    public readonly customResponseHeaders!: pulumi.Output<string[] | undefined>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The resource URL for the edge security policy associated with this backend service.
     */
    public readonly edgeSecurityPolicy!: pulumi.Output<string | undefined>;
    /**
     * If true, enable Cloud CDN for this BackendService.
     */
    public readonly enableCdn!: pulumi.Output<boolean | undefined>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this
     * object. This field is used in optimistic locking.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    public /*out*/ readonly generatedId!: pulumi.Output<number>;
    /**
     * The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
     * for health checking this BackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     * For internal load balancing, a URL to a HealthCheck resource must be specified instead.
     */
    public readonly healthChecks!: pulumi.Output<string | undefined>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    public readonly iap!: pulumi.Output<outputs.compute.BackendServiceIap | undefined>;
    /**
     * Indicates whether the backend service will be used with internal or
     * external load balancing. A backend service created for one type of
     * load balancing cannot be used with the other. For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL_SELF_MANAGED`, `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
    /**
     * A list of locality load balancing policies to be used in order of
     * preference. Either the policy or the customPolicy field should be set.
     * Overrides any value set in the localityLbPolicy field.
     * localityLbPolicies is only supported when the BackendService is referenced
     * by a URL Map that is referenced by a target gRPC proxy that has the
     * validateForProxyless field set to true.
     * Structure is documented below.
     */
    public readonly localityLbPolicies!: pulumi.Output<outputs.compute.BackendServiceLocalityLbPolicy[] | undefined>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are:
     */
    public readonly localityLbPolicy!: pulumi.Output<string | undefined>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.BackendServiceLogConfig>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * Applicable backend service types can be a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.
     * Structure is documented below.
     */
    public readonly outlierDetection!: pulumi.Output<outputs.compute.BackendServiceOutlierDetection | undefined>;
    /**
     * Name of backend port. The same name should appear in the instance
     * groups referenced by this service. Required when the load balancing
     * scheme is EXTERNAL.
     */
    public readonly portName!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The protocol this BackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API. **NOTE**: With protocol “UNSPECIFIED”,
     * the backend service can be used by Layer 4 Internal Load Balancing or Network Load Balancing
     * with TCP/UDP/L3_DEFAULT Forwarding Rule protocol.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `TCP`, `SSL`, `GRPC`, `UNSPECIFIED`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The security policy associated with this backend service.
     */
    public readonly securityPolicy!: pulumi.Output<string | undefined>;
    /**
     * The security settings that apply to this backend service. This field is applicable to either
     * a regional backend service with the serviceProtocol set to HTTP, HTTPS, or HTTP2, and
     * loadBalancingScheme set to INTERNAL_MANAGED; or a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    public readonly securitySettings!: pulumi.Output<outputs.compute.BackendServiceSecuritySettings | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`.
     */
    public readonly sessionAffinity!: pulumi.Output<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a BackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BackendServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendServiceArgs | BackendServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendServiceState | undefined;
            resourceInputs["affinityCookieTtlSec"] = state ? state.affinityCookieTtlSec : undefined;
            resourceInputs["backends"] = state ? state.backends : undefined;
            resourceInputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            resourceInputs["circuitBreakers"] = state ? state.circuitBreakers : undefined;
            resourceInputs["compressionMode"] = state ? state.compressionMode : undefined;
            resourceInputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            resourceInputs["consistentHash"] = state ? state.consistentHash : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["customRequestHeaders"] = state ? state.customRequestHeaders : undefined;
            resourceInputs["customResponseHeaders"] = state ? state.customResponseHeaders : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeSecurityPolicy"] = state ? state.edgeSecurityPolicy : undefined;
            resourceInputs["enableCdn"] = state ? state.enableCdn : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["generatedId"] = state ? state.generatedId : undefined;
            resourceInputs["healthChecks"] = state ? state.healthChecks : undefined;
            resourceInputs["iap"] = state ? state.iap : undefined;
            resourceInputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            resourceInputs["localityLbPolicies"] = state ? state.localityLbPolicies : undefined;
            resourceInputs["localityLbPolicy"] = state ? state.localityLbPolicy : undefined;
            resourceInputs["logConfig"] = state ? state.logConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outlierDetection"] = state ? state.outlierDetection : undefined;
            resourceInputs["portName"] = state ? state.portName : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["securityPolicy"] = state ? state.securityPolicy : undefined;
            resourceInputs["securitySettings"] = state ? state.securitySettings : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            resourceInputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as BackendServiceArgs | undefined;
            resourceInputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            resourceInputs["backends"] = args ? args.backends : undefined;
            resourceInputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            resourceInputs["circuitBreakers"] = args ? args.circuitBreakers : undefined;
            resourceInputs["compressionMode"] = args ? args.compressionMode : undefined;
            resourceInputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            resourceInputs["consistentHash"] = args ? args.consistentHash : undefined;
            resourceInputs["customRequestHeaders"] = args ? args.customRequestHeaders : undefined;
            resourceInputs["customResponseHeaders"] = args ? args.customResponseHeaders : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeSecurityPolicy"] = args ? args.edgeSecurityPolicy : undefined;
            resourceInputs["enableCdn"] = args ? args.enableCdn : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["iap"] = args ? args.iap : undefined;
            resourceInputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            resourceInputs["localityLbPolicies"] = args ? args.localityLbPolicies : undefined;
            resourceInputs["localityLbPolicy"] = args ? args.localityLbPolicy : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outlierDetection"] = args ? args.outlierDetection : undefined;
            resourceInputs["portName"] = args ? args.portName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            resourceInputs["securitySettings"] = args ? args.securitySettings : undefined;
            resourceInputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            resourceInputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["generatedId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackendService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendService resources.
 */
export interface BackendServiceState {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this BackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.BackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.BackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.BackendServiceCircuitBreakers>;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     * Possible values are: `AUTOMATIC`, `DISABLED`.
     */
    compressionMode?: pulumi.Input<string>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing. This field only applies if the loadBalancingScheme is set to
     * INTERNAL_SELF_MANAGED. This field is only applicable when localityLbPolicy is
     * set to MAGLEV or RING_HASH.
     * Structure is documented below.
     */
    consistentHash?: pulumi.Input<inputs.compute.BackendServiceConsistentHash>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * requests.
     */
    customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * responses.
     */
    customResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The resource URL for the edge security policy associated with this backend service.
     */
    edgeSecurityPolicy?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this BackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this
     * object. This field is used in optimistic locking.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    generatedId?: pulumi.Input<number>;
    /**
     * The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
     * for health checking this BackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     * For internal load balancing, a URL to a HealthCheck resource must be specified instead.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.BackendServiceIap>;
    /**
     * Indicates whether the backend service will be used with internal or
     * external load balancing. A backend service created for one type of
     * load balancing cannot be used with the other. For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL_SELF_MANAGED`, `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
    /**
     * A list of locality load balancing policies to be used in order of
     * preference. Either the policy or the customPolicy field should be set.
     * Overrides any value set in the localityLbPolicy field.
     * localityLbPolicies is only supported when the BackendService is referenced
     * by a URL Map that is referenced by a target gRPC proxy that has the
     * validateForProxyless field set to true.
     * Structure is documented below.
     */
    localityLbPolicies?: pulumi.Input<pulumi.Input<inputs.compute.BackendServiceLocalityLbPolicy>[]>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are:
     */
    localityLbPolicy?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.BackendServiceLogConfig>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * Applicable backend service types can be a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.BackendServiceOutlierDetection>;
    /**
     * Name of backend port. The same name should appear in the instance
     * groups referenced by this service. Required when the load balancing
     * scheme is EXTERNAL.
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this BackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API. **NOTE**: With protocol “UNSPECIFIED”,
     * the backend service can be used by Layer 4 Internal Load Balancing or Network Load Balancing
     * with TCP/UDP/L3_DEFAULT Forwarding Rule protocol.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `TCP`, `SSL`, `GRPC`, `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The security policy associated with this backend service.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * The security settings that apply to this backend service. This field is applicable to either
     * a regional backend service with the serviceProtocol set to HTTP, HTTPS, or HTTP2, and
     * loadBalancingScheme set to INTERNAL_MANAGED; or a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    securitySettings?: pulumi.Input<inputs.compute.BackendServiceSecuritySettings>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BackendService resource.
 */
export interface BackendServiceArgs {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this BackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.BackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.BackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the loadBalancingScheme is set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.BackendServiceCircuitBreakers>;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     * Possible values are: `AUTOMATIC`, `DISABLED`.
     */
    compressionMode?: pulumi.Input<string>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing. This field only applies if the loadBalancingScheme is set to
     * INTERNAL_SELF_MANAGED. This field is only applicable when localityLbPolicy is
     * set to MAGLEV or RING_HASH.
     * Structure is documented below.
     */
    consistentHash?: pulumi.Input<inputs.compute.BackendServiceConsistentHash>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * requests.
     */
    customRequestHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Headers that the HTTP/S load balancer should add to proxied
     * responses.
     */
    customResponseHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The resource URL for the edge security policy associated with this backend service.
     */
    edgeSecurityPolicy?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this BackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource
     * for health checking this BackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     * For internal load balancing, a URL to a HealthCheck resource must be specified instead.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.BackendServiceIap>;
    /**
     * Indicates whether the backend service will be used with internal or
     * external load balancing. A backend service created for one type of
     * load balancing cannot be used with the other. For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL_SELF_MANAGED`, `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
    /**
     * A list of locality load balancing policies to be used in order of
     * preference. Either the policy or the customPolicy field should be set.
     * Overrides any value set in the localityLbPolicy field.
     * localityLbPolicies is only supported when the BackendService is referenced
     * by a URL Map that is referenced by a target gRPC proxy that has the
     * validateForProxyless field set to true.
     * Structure is documented below.
     */
    localityLbPolicies?: pulumi.Input<pulumi.Input<inputs.compute.BackendServiceLocalityLbPolicy>[]>;
    /**
     * The load balancing algorithm used within the scope of the locality.
     * The possible values are:
     */
    localityLbPolicy?: pulumi.Input<string>;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service.
     * If logging is enabled, logs will be exported to Stackdriver.
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.BackendServiceLogConfig>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * Applicable backend service types can be a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL_MANAGED.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.BackendServiceOutlierDetection>;
    /**
     * Name of backend port. The same name should appear in the instance
     * groups referenced by this service. Required when the load balancing
     * scheme is EXTERNAL.
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this BackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API. **NOTE**: With protocol “UNSPECIFIED”,
     * the backend service can be used by Layer 4 Internal Load Balancing or Network Load Balancing
     * with TCP/UDP/L3_DEFAULT Forwarding Rule protocol.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `TCP`, `SSL`, `GRPC`, `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The security policy associated with this backend service.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * The security settings that apply to this backend service. This field is applicable to either
     * a regional backend service with the serviceProtocol set to HTTP, HTTPS, or HTTP2, and
     * loadBalancingScheme set to INTERNAL_MANAGED; or a global backend service with the
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     */
    securitySettings?: pulumi.Input<inputs.compute.BackendServiceSecuritySettings>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}
