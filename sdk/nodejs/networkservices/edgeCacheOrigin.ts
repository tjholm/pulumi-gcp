// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * EdgeCacheOrigin represents a HTTP-reachable backend for an EdgeCacheService.
 *
 * To get more information about EdgeCacheOrigin, see:
 *
 * * [API documentation](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheOrigins)
 *
 * ## Example Usage
 * ### Network Services Edge Cache Origin Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networkservices.EdgeCacheOrigin("default", {
 *     description: "The default bucket for media edge test",
 *     originAddress: "gs://media-edge-default",
 * });
 * ```
 * ### Network Services Edge Cache Origin Advanced
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const fallback = new gcp.networkservices.EdgeCacheOrigin("fallback", {
 *     originAddress: "fallback.example.com",
 *     description: "The default bucket for media edge test",
 *     maxAttempts: 3,
 *     protocol: "HTTP",
 *     port: 80,
 *     retryConditions: [
 *         "CONNECT_FAILURE",
 *         "NOT_FOUND",
 *         "HTTP_5XX",
 *         "FORBIDDEN",
 *     ],
 *     timeout: {
 *         connectTimeout: "10s",
 *         maxAttemptsTimeout: "20s",
 *         responseTimeout: "60s",
 *         readTimeout: "5s",
 *     },
 *     originOverrideAction: {
 *         urlRewrite: {
 *             hostRewrite: "example.com",
 *         },
 *         headerAction: {
 *             requestHeadersToAdds: [{
 *                 headerName: "x-header",
 *                 headerValue: "value",
 *                 replace: true,
 *             }],
 *         },
 *     },
 *     originRedirect: {
 *         redirectConditions: [
 *             "MOVED_PERMANENTLY",
 *             "FOUND",
 *             "SEE_OTHER",
 *             "TEMPORARY_REDIRECT",
 *             "PERMANENT_REDIRECT",
 *         ],
 *     },
 * });
 * const _default = new gcp.networkservices.EdgeCacheOrigin("default", {
 *     originAddress: "gs://media-edge-default",
 *     failoverOrigin: fallback.id,
 *     description: "The default bucket for media edge test",
 *     maxAttempts: 2,
 *     labels: {
 *         a: "b",
 *     },
 *     timeout: {
 *         connectTimeout: "10s",
 *     },
 * });
 * ```
 * ### Network Services Edge Cache Origin V4auth
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const secret_basic = new gcp.secretmanager.Secret("secret-basic", {
 *     secretId: "secret-name",
 *     replication: {
 *         auto: {},
 *     },
 * });
 * const secret_version_basic = new gcp.secretmanager.SecretVersion("secret-version-basic", {
 *     secret: secret_basic.id,
 *     secretData: "secret-data",
 * });
 * const _default = new gcp.networkservices.EdgeCacheOrigin("default", {
 *     originAddress: "gs://media-edge-default",
 *     description: "The default bucket for V4 authentication",
 *     awsV4Authentication: {
 *         accessKeyId: "ACCESSKEYID",
 *         secretAccessKeyVersion: secret_version_basic.id,
 *         originRegion: "auto",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * EdgeCacheOrigin can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin default {{name}}
 * ```
 */
export class EdgeCacheOrigin extends pulumi.CustomResource {
    /**
     * Get an existing EdgeCacheOrigin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EdgeCacheOriginState, opts?: pulumi.CustomResourceOptions): EdgeCacheOrigin {
        return new EdgeCacheOrigin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkservices/edgeCacheOrigin:EdgeCacheOrigin';

    /**
     * Returns true if the given object is an instance of EdgeCacheOrigin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EdgeCacheOrigin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EdgeCacheOrigin.__pulumiType;
    }

    /**
     * Enable AWS Signature Version 4 origin authentication.
     * Structure is documented below.
     */
    public readonly awsV4Authentication!: pulumi.Output<outputs.networkservices.EdgeCacheOriginAwsV4Authentication | undefined>;
    /**
     * A human-readable description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Origin resource to try when the current origin cannot be reached.
     * After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
     * The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
     * A reference to a Topic resource.
     */
    public readonly failoverOrigin!: pulumi.Output<string | undefined>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
     * Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
     * retryConditions and failoverOrigin to control its own cache fill failures.
     * The total number of allowed attempts to cache fill across this and failover origins is limited to four.
     * The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
     * The last valid, non-retried response from all origins will be returned to the client.
     * If no origin returns a valid response, an HTTP 502 will be returned to the client.
     * Defaults to 1. Must be a value greater than 0 and less than 4.
     */
    public readonly maxAttempts!: pulumi.Output<number | undefined>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
     * This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com, IPv4: 35.218.1.1, IPv6: 2607:f8b0:4012:809::200e, Cloud Storage: gs://bucketname
     * When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.  It must not contain a protocol (e.g., https://) and it must not contain any slashes.
     * If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
     */
    public readonly originAddress!: pulumi.Output<string>;
    /**
     * The override actions, including url rewrites and header
     * additions, for requests that use this origin.
     * Structure is documented below.
     */
    public readonly originOverrideAction!: pulumi.Output<outputs.networkservices.EdgeCacheOriginOriginOverrideAction | undefined>;
    /**
     * Follow redirects from this origin.
     * Structure is documented below.
     */
    public readonly originRedirect!: pulumi.Output<outputs.networkservices.EdgeCacheOriginOriginRedirect | undefined>;
    /**
     * The port to connect to the origin on.
     * Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
     * When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
     * Possible values are: `HTTP2`, `HTTPS`, `HTTP`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Specifies one or more retry conditions for the configured origin.
     * If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
     * the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
     * The default retryCondition is "CONNECT_FAILURE".
     * retryConditions apply to this origin, and not subsequent failoverOrigin(s),
     * which may specify their own retryConditions and maxAttempts.
     * Valid values are:
     * - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
     * - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
     * - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
     * - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
     * - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
     * - FORBIDDEN: Retry if the origin returns a HTTP 403 (Forbidden).
     * Each value may be one of: `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, `NOT_FOUND`, `FORBIDDEN`.
     */
    public readonly retryConditions!: pulumi.Output<string[]>;
    /**
     * The connection and HTTP timeout configuration for this origin.
     * Structure is documented below.
     */
    public readonly timeout!: pulumi.Output<outputs.networkservices.EdgeCacheOriginTimeout | undefined>;

    /**
     * Create a EdgeCacheOrigin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EdgeCacheOriginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EdgeCacheOriginArgs | EdgeCacheOriginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EdgeCacheOriginState | undefined;
            resourceInputs["awsV4Authentication"] = state ? state.awsV4Authentication : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["failoverOrigin"] = state ? state.failoverOrigin : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maxAttempts"] = state ? state.maxAttempts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["originAddress"] = state ? state.originAddress : undefined;
            resourceInputs["originOverrideAction"] = state ? state.originOverrideAction : undefined;
            resourceInputs["originRedirect"] = state ? state.originRedirect : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["retryConditions"] = state ? state.retryConditions : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as EdgeCacheOriginArgs | undefined;
            if ((!args || args.originAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'originAddress'");
            }
            resourceInputs["awsV4Authentication"] = args ? args.awsV4Authentication : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["failoverOrigin"] = args ? args.failoverOrigin : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maxAttempts"] = args ? args.maxAttempts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["originAddress"] = args ? args.originAddress : undefined;
            resourceInputs["originOverrideAction"] = args ? args.originOverrideAction : undefined;
            resourceInputs["originRedirect"] = args ? args.originRedirect : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["retryConditions"] = args ? args.retryConditions : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EdgeCacheOrigin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EdgeCacheOrigin resources.
 */
export interface EdgeCacheOriginState {
    /**
     * Enable AWS Signature Version 4 origin authentication.
     * Structure is documented below.
     */
    awsV4Authentication?: pulumi.Input<inputs.networkservices.EdgeCacheOriginAwsV4Authentication>;
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The Origin resource to try when the current origin cannot be reached.
     * After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
     * The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
     * A reference to a Topic resource.
     */
    failoverOrigin?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
     * Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
     * retryConditions and failoverOrigin to control its own cache fill failures.
     * The total number of allowed attempts to cache fill across this and failover origins is limited to four.
     * The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
     * The last valid, non-retried response from all origins will be returned to the client.
     * If no origin returns a valid response, an HTTP 502 will be returned to the client.
     * Defaults to 1. Must be a value greater than 0 and less than 4.
     */
    maxAttempts?: pulumi.Input<number>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
     * This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com, IPv4: 35.218.1.1, IPv6: 2607:f8b0:4012:809::200e, Cloud Storage: gs://bucketname
     * When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.  It must not contain a protocol (e.g., https://) and it must not contain any slashes.
     * If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
     */
    originAddress?: pulumi.Input<string>;
    /**
     * The override actions, including url rewrites and header
     * additions, for requests that use this origin.
     * Structure is documented below.
     */
    originOverrideAction?: pulumi.Input<inputs.networkservices.EdgeCacheOriginOriginOverrideAction>;
    /**
     * Follow redirects from this origin.
     * Structure is documented below.
     */
    originRedirect?: pulumi.Input<inputs.networkservices.EdgeCacheOriginOriginRedirect>;
    /**
     * The port to connect to the origin on.
     * Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
     * When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
     * Possible values are: `HTTP2`, `HTTPS`, `HTTP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Specifies one or more retry conditions for the configured origin.
     * If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
     * the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
     * The default retryCondition is "CONNECT_FAILURE".
     * retryConditions apply to this origin, and not subsequent failoverOrigin(s),
     * which may specify their own retryConditions and maxAttempts.
     * Valid values are:
     * - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
     * - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
     * - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
     * - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
     * - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
     * - FORBIDDEN: Retry if the origin returns a HTTP 403 (Forbidden).
     * Each value may be one of: `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, `NOT_FOUND`, `FORBIDDEN`.
     */
    retryConditions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection and HTTP timeout configuration for this origin.
     * Structure is documented below.
     */
    timeout?: pulumi.Input<inputs.networkservices.EdgeCacheOriginTimeout>;
}

/**
 * The set of arguments for constructing a EdgeCacheOrigin resource.
 */
export interface EdgeCacheOriginArgs {
    /**
     * Enable AWS Signature Version 4 origin authentication.
     * Structure is documented below.
     */
    awsV4Authentication?: pulumi.Input<inputs.networkservices.EdgeCacheOriginAwsV4Authentication>;
    /**
     * A human-readable description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The Origin resource to try when the current origin cannot be reached.
     * After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
     * The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
     * A reference to a Topic resource.
     */
    failoverOrigin?: pulumi.Input<string>;
    /**
     * Set of label tags associated with the EdgeCache resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
     * Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
     * retryConditions and failoverOrigin to control its own cache fill failures.
     * The total number of allowed attempts to cache fill across this and failover origins is limited to four.
     * The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
     * The last valid, non-retried response from all origins will be returned to the client.
     * If no origin returns a valid response, an HTTP 502 will be returned to the client.
     * Defaults to 1. Must be a value greater than 0 and less than 4.
     */
    maxAttempts?: pulumi.Input<number>;
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
     * This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com, IPv4: 35.218.1.1, IPv6: 2607:f8b0:4012:809::200e, Cloud Storage: gs://bucketname
     * When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.  It must not contain a protocol (e.g., https://) and it must not contain any slashes.
     * If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
     */
    originAddress: pulumi.Input<string>;
    /**
     * The override actions, including url rewrites and header
     * additions, for requests that use this origin.
     * Structure is documented below.
     */
    originOverrideAction?: pulumi.Input<inputs.networkservices.EdgeCacheOriginOriginOverrideAction>;
    /**
     * Follow redirects from this origin.
     * Structure is documented below.
     */
    originRedirect?: pulumi.Input<inputs.networkservices.EdgeCacheOriginOriginRedirect>;
    /**
     * The port to connect to the origin on.
     * Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
     * When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server.
     * Possible values are: `HTTP2`, `HTTPS`, `HTTP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Specifies one or more retry conditions for the configured origin.
     * If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
     * the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
     * The default retryCondition is "CONNECT_FAILURE".
     * retryConditions apply to this origin, and not subsequent failoverOrigin(s),
     * which may specify their own retryConditions and maxAttempts.
     * Valid values are:
     * - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
     * - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
     * - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
     * - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
     * - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet.
     * - FORBIDDEN: Retry if the origin returns a HTTP 403 (Forbidden).
     * Each value may be one of: `CONNECT_FAILURE`, `HTTP_5XX`, `GATEWAY_ERROR`, `RETRIABLE_4XX`, `NOT_FOUND`, `FORBIDDEN`.
     */
    retryConditions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The connection and HTTP timeout configuration for this origin.
     * Structure is documented below.
     */
    timeout?: pulumi.Input<inputs.networkservices.EdgeCacheOriginTimeout>;
}
