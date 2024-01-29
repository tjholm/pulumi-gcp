// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Region Backend Service defines a regionally-scoped group of virtual
 * machines that will serve traffic for load balancing.
 *
 * To get more information about RegionBackendService, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionBackendServices)
 * * How-to Guides
 *     * [Internal TCP/UDP Load Balancing](https://cloud.google.com/compute/docs/load-balancing/internal/)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * RegionBackendService can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/backendServices/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, RegionBackendService can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default projects/{{project}}/regions/{{region}}/backendServices/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionBackendService:RegionBackendService default {{name}}
 * ```
 */
export class RegionBackendService extends pulumi.CustomResource {
    /**
     * Get an existing RegionBackendService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionBackendServiceState, opts?: pulumi.CustomResourceOptions): RegionBackendService {
        return new RegionBackendService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionBackendService:RegionBackendService';

    /**
     * Returns true if the given object is an instance of RegionBackendService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionBackendService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionBackendService.__pulumiType;
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
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    public readonly backends!: pulumi.Output<outputs.compute.RegionBackendServiceBackend[] | undefined>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    public readonly cdnPolicy!: pulumi.Output<outputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    public readonly circuitBreakers!: pulumi.Output<outputs.compute.RegionBackendServiceCircuitBreakers | undefined>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    public readonly connectionDrainingTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Connection Tracking configuration for this BackendService.
     * This is available only for Layer 4 Internal Load Balancing and
     * Network Load Balancing.
     * Structure is documented below.
     */
    public readonly connectionTrackingPolicy!: pulumi.Output<outputs.compute.RegionBackendServiceConnectionTrackingPolicy | undefined>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    public readonly consistentHash!: pulumi.Output<outputs.compute.RegionBackendServiceConsistentHash | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    public readonly enableCdn!: pulumi.Output<boolean | undefined>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    public readonly failoverPolicy!: pulumi.Output<outputs.compute.RegionBackendServiceFailoverPolicy | undefined>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this
     * object. This field is used in optimistic locking.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    public readonly healthChecks!: pulumi.Output<string | undefined>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    public readonly iap!: pulumi.Output<outputs.compute.RegionBackendServiceIap | undefined>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s). For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `INTERNAL`.
     * Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL`, `INTERNAL_MANAGED`.
     */
    public readonly loadBalancingScheme!: pulumi.Output<string | undefined>;
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
    public readonly logConfig!: pulumi.Output<outputs.compute.RegionBackendServiceLogConfig>;
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
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    public readonly outlierDetection!: pulumi.Output<outputs.compute.RegionBackendServiceOutlierDetection | undefined>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    public readonly portName!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, `UNSPECIFIED`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The security policy associated with this backend service.
     */
    public readonly securityPolicy!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`, `CLIENT_IP_NO_DESTINATION`.
     */
    public readonly sessionAffinity!: pulumi.Output<string>;
    /**
     * Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S) load balancing.
     * Structure is documented below.
     */
    public readonly subsetting!: pulumi.Output<outputs.compute.RegionBackendServiceSubsetting | undefined>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    public readonly timeoutSec!: pulumi.Output<number>;

    /**
     * Create a RegionBackendService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegionBackendServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionBackendServiceArgs | RegionBackendServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionBackendServiceState | undefined;
            resourceInputs["affinityCookieTtlSec"] = state ? state.affinityCookieTtlSec : undefined;
            resourceInputs["backends"] = state ? state.backends : undefined;
            resourceInputs["cdnPolicy"] = state ? state.cdnPolicy : undefined;
            resourceInputs["circuitBreakers"] = state ? state.circuitBreakers : undefined;
            resourceInputs["connectionDrainingTimeoutSec"] = state ? state.connectionDrainingTimeoutSec : undefined;
            resourceInputs["connectionTrackingPolicy"] = state ? state.connectionTrackingPolicy : undefined;
            resourceInputs["consistentHash"] = state ? state.consistentHash : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableCdn"] = state ? state.enableCdn : undefined;
            resourceInputs["failoverPolicy"] = state ? state.failoverPolicy : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["healthChecks"] = state ? state.healthChecks : undefined;
            resourceInputs["iap"] = state ? state.iap : undefined;
            resourceInputs["loadBalancingScheme"] = state ? state.loadBalancingScheme : undefined;
            resourceInputs["localityLbPolicy"] = state ? state.localityLbPolicy : undefined;
            resourceInputs["logConfig"] = state ? state.logConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["outlierDetection"] = state ? state.outlierDetection : undefined;
            resourceInputs["portName"] = state ? state.portName : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityPolicy"] = state ? state.securityPolicy : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["sessionAffinity"] = state ? state.sessionAffinity : undefined;
            resourceInputs["subsetting"] = state ? state.subsetting : undefined;
            resourceInputs["timeoutSec"] = state ? state.timeoutSec : undefined;
        } else {
            const args = argsOrState as RegionBackendServiceArgs | undefined;
            resourceInputs["affinityCookieTtlSec"] = args ? args.affinityCookieTtlSec : undefined;
            resourceInputs["backends"] = args ? args.backends : undefined;
            resourceInputs["cdnPolicy"] = args ? args.cdnPolicy : undefined;
            resourceInputs["circuitBreakers"] = args ? args.circuitBreakers : undefined;
            resourceInputs["connectionDrainingTimeoutSec"] = args ? args.connectionDrainingTimeoutSec : undefined;
            resourceInputs["connectionTrackingPolicy"] = args ? args.connectionTrackingPolicy : undefined;
            resourceInputs["consistentHash"] = args ? args.consistentHash : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableCdn"] = args ? args.enableCdn : undefined;
            resourceInputs["failoverPolicy"] = args ? args.failoverPolicy : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["iap"] = args ? args.iap : undefined;
            resourceInputs["loadBalancingScheme"] = args ? args.loadBalancingScheme : undefined;
            resourceInputs["localityLbPolicy"] = args ? args.localityLbPolicy : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["outlierDetection"] = args ? args.outlierDetection : undefined;
            resourceInputs["portName"] = args ? args.portName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityPolicy"] = args ? args.securityPolicy : undefined;
            resourceInputs["sessionAffinity"] = args ? args.sessionAffinity : undefined;
            resourceInputs["subsetting"] = args ? args.subsetting : undefined;
            resourceInputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionBackendService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionBackendService resources.
 */
export interface RegionBackendServiceState {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.RegionBackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.RegionBackendServiceCircuitBreakers>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Connection Tracking configuration for this BackendService.
     * This is available only for Layer 4 Internal Load Balancing and
     * Network Load Balancing.
     * Structure is documented below.
     */
    connectionTrackingPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceConnectionTrackingPolicy>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    consistentHash?: pulumi.Input<inputs.compute.RegionBackendServiceConsistentHash>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    failoverPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceFailoverPolicy>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this
     * object. This field is used in optimistic locking.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.RegionBackendServiceIap>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s). For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `INTERNAL`.
     * Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL`, `INTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
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
    logConfig?: pulumi.Input<inputs.compute.RegionBackendServiceLogConfig>;
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
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    network?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.RegionBackendServiceOutlierDetection>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The security policy associated with this backend service.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`, `CLIENT_IP_NO_DESTINATION`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S) load balancing.
     * Structure is documented below.
     */
    subsetting?: pulumi.Input<inputs.compute.RegionBackendServiceSubsetting>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RegionBackendService resource.
 */
export interface RegionBackendServiceArgs {
    /**
     * Lifetime of cookies in seconds if sessionAffinity is
     * GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
     * only until the end of the browser session (or equivalent). The
     * maximum allowed value for TTL is one day.
     * When the load balancing scheme is INTERNAL, this field is not used.
     */
    affinityCookieTtlSec?: pulumi.Input<number>;
    /**
     * The set of backends that serve this RegionBackendService.
     * Structure is documented below.
     */
    backends?: pulumi.Input<pulumi.Input<inputs.compute.RegionBackendServiceBackend>[]>;
    /**
     * Cloud CDN configuration for this BackendService.
     * Structure is documented below.
     */
    cdnPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceCdnPolicy>;
    /**
     * Settings controlling the volume of connections to a backend service. This field
     * is applicable only when the `loadBalancingScheme` is set to INTERNAL_MANAGED
     * and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    circuitBreakers?: pulumi.Input<inputs.compute.RegionBackendServiceCircuitBreakers>;
    /**
     * Time for which instance will be drained (not accept new
     * connections, but still work to finish started).
     */
    connectionDrainingTimeoutSec?: pulumi.Input<number>;
    /**
     * Connection Tracking configuration for this BackendService.
     * This is available only for Layer 4 Internal Load Balancing and
     * Network Load Balancing.
     * Structure is documented below.
     */
    connectionTrackingPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceConnectionTrackingPolicy>;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session
     * affinity based on HTTP headers, cookies or other properties. This load balancing
     * policy is applicable only for HTTP connections. The affinity to a particular
     * destination host will be lost when one or more hosts are added/removed from the
     * destination service. This field specifies parameters that control consistent
     * hashing.
     * This field only applies when all of the following are true -
     */
    consistentHash?: pulumi.Input<inputs.compute.RegionBackendServiceConsistentHash>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * If true, enable Cloud CDN for this RegionBackendService.
     */
    enableCdn?: pulumi.Input<boolean>;
    /**
     * Policy for failovers.
     * Structure is documented below.
     */
    failoverPolicy?: pulumi.Input<inputs.compute.RegionBackendServiceFailoverPolicy>;
    /**
     * The set of URLs to HealthCheck resources for health checking
     * this RegionBackendService. Currently at most one health
     * check can be specified.
     * A health check must be specified unless the backend service uses an internet
     * or serverless NEG as a backend.
     */
    healthChecks?: pulumi.Input<string>;
    /**
     * Settings for enabling Cloud Identity Aware Proxy
     * Structure is documented below.
     */
    iap?: pulumi.Input<inputs.compute.RegionBackendServiceIap>;
    /**
     * Indicates what kind of load balancing this regional backend service
     * will be used for. A backend service created for one type of load
     * balancing cannot be used with the other(s). For more information, refer to
     * [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service).
     * Default value is `INTERNAL`.
     * Possible values are: `EXTERNAL`, `EXTERNAL_MANAGED`, `INTERNAL`, `INTERNAL_MANAGED`.
     */
    loadBalancingScheme?: pulumi.Input<string>;
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
    logConfig?: pulumi.Input<inputs.compute.RegionBackendServiceLogConfig>;
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
     * The URL of the network to which this backend service belongs.
     * This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    network?: pulumi.Input<string>;
    /**
     * Settings controlling eviction of unhealthy hosts from the load balancing pool.
     * This field is applicable only when the `loadBalancingScheme` is set
     * to INTERNAL_MANAGED and the `protocol` is set to HTTP, HTTPS, or HTTP2.
     * Structure is documented below.
     */
    outlierDetection?: pulumi.Input<inputs.compute.RegionBackendServiceOutlierDetection>;
    /**
     * A named port on a backend instance group representing the port for
     * communication to the backend VMs in that group. Required when the
     * loadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
     * and the backends are instance groups. The named port must be defined on each
     * backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
     * default of "http" if not given.
     * Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
     */
    portName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol this RegionBackendService uses to communicate with backends.
     * The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
     * types and may result in errors if used with the GA API.
     * Possible values are: `HTTP`, `HTTPS`, `HTTP2`, `SSL`, `TCP`, `UDP`, `GRPC`, `UNSPECIFIED`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The security policy associated with this backend service.
     */
    securityPolicy?: pulumi.Input<string>;
    /**
     * Type of session affinity to use. The default is NONE. Session affinity is
     * not applicable if the protocol is UDP.
     * Possible values are: `NONE`, `CLIENT_IP`, `CLIENT_IP_PORT_PROTO`, `CLIENT_IP_PROTO`, `GENERATED_COOKIE`, `HEADER_FIELD`, `HTTP_COOKIE`, `CLIENT_IP_NO_DESTINATION`.
     */
    sessionAffinity?: pulumi.Input<string>;
    /**
     * Subsetting configuration for this BackendService. Currently this is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S) load balancing.
     * Structure is documented below.
     */
    subsetting?: pulumi.Input<inputs.compute.RegionBackendServiceSubsetting>;
    /**
     * How many seconds to wait for the backend before considering it a
     * failed request. Default is 30 seconds. Valid range is [1, 86400].
     */
    timeoutSec?: pulumi.Input<number>;
}
