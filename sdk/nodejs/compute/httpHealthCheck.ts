// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An HttpHealthCheck resource. This resource defines a template for how
 * individual VMs should be checked for health, via HTTP.
 *
 * > **Note:** gcp.compute.HttpHealthCheck is a legacy health check.
 * The newer [gcp.compute.HealthCheck](https://www.terraform.io/docs/providers/google/r/compute_health_check.html)
 * should be preferred for all uses except
 * [Network Load Balancers](https://cloud.google.com/compute/docs/load-balancing/network/)
 * which still require the legacy version.
 *
 * To get more information about HttpHealthCheck, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/httpHealthChecks)
 * * How-to Guides
 *     * [Adding Health Checks](https://cloud.google.com/compute/docs/load-balancing/health-checks#legacy_health_checks)
 *
 * ## Example Usage
 * ### Http Health Check Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/health_check",
 *     timeoutSec: 1,
 * });
 * ```
 *
 * ## Import
 *
 * HttpHealthCheck can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/httpHealthCheck:HttpHealthCheck default projects/{{project}}/global/httpHealthChecks/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/httpHealthCheck:HttpHealthCheck default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/httpHealthCheck:HttpHealthCheck default {{name}}
 * ```
 */
export class HttpHealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing HttpHealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HttpHealthCheckState, opts?: pulumi.CustomResourceOptions): HttpHealthCheck {
        return new HttpHealthCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/httpHealthCheck:HttpHealthCheck';

    /**
     * Returns true if the given object is an instance of HttpHealthCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HttpHealthCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HttpHealthCheck.__pulumiType;
    }

    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    public readonly checkIntervalSec!: pulumi.Output<number | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    public readonly healthyThreshold!: pulumi.Output<number | undefined>;
    /**
     * The value of the host header in the HTTP health check request. If
     * left empty (default value), the public IP on behalf of which this
     * health check is performed will be used.
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The TCP port number for the HTTP health check request.
     * The default value is 80.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The request path of the HTTP health check request.
     * The default value is /.
     */
    public readonly requestPath!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    public readonly timeoutSec!: pulumi.Output<number | undefined>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    public readonly unhealthyThreshold!: pulumi.Output<number | undefined>;

    /**
     * Create a HttpHealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HttpHealthCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HttpHealthCheckArgs | HttpHealthCheckState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HttpHealthCheckState | undefined;
            resourceInputs["checkIntervalSec"] = state ? state.checkIntervalSec : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["requestPath"] = state ? state.requestPath : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["timeoutSec"] = state ? state.timeoutSec : undefined;
            resourceInputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HttpHealthCheckArgs | undefined;
            resourceInputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestPath"] = args ? args.requestPath : undefined;
            resourceInputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            resourceInputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HttpHealthCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HttpHealthCheck resources.
 */
export interface HttpHealthCheckState {
    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    checkIntervalSec?: pulumi.Input<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The value of the host header in the HTTP health check request. If
     * left empty (default value), the public IP on behalf of which this
     * health check is performed will be used.
     */
    host?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The TCP port number for the HTTP health check request.
     * The default value is 80.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The request path of the HTTP health check request.
     * The default value is /.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    timeoutSec?: pulumi.Input<number>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HttpHealthCheck resource.
 */
export interface HttpHealthCheckArgs {
    /**
     * How often (in seconds) to send a health check. The default value is 5
     * seconds.
     */
    checkIntervalSec?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many
     * consecutive successes. The default value is 2.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The value of the host header in the HTTP health check request. If
     * left empty (default value), the public IP on behalf of which this
     * health check is performed will be used.
     */
    host?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The TCP port number for the HTTP health check request.
     * The default value is 80.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The request path of the HTTP health check request.
     * The default value is /.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * How long (in seconds) to wait before claiming failure.
     * The default value is 5 seconds.  It is invalid for timeoutSec to have
     * greater value than checkIntervalSec.
     */
    timeoutSec?: pulumi.Input<number>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many
     * consecutive failures. The default value is 2.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}
