// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * UrlMaps are used to route requests to a backend service based on rules
 * that you define for the host and path of an incoming URL.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * RegionUrlMap can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/urlMaps/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, RegionUrlMap can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default projects/{{project}}/regions/{{region}}/urlMaps/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/regionUrlMap:RegionUrlMap default {{name}}
 * ```
 */
export class RegionUrlMap extends pulumi.CustomResource {
    /**
     * Get an existing RegionUrlMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegionUrlMapState, opts?: pulumi.CustomResourceOptions): RegionUrlMap {
        return new RegionUrlMap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/regionUrlMap:RegionUrlMap';

    /**
     * Returns true if the given object is an instance of RegionUrlMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionUrlMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionUrlMap.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.
     * defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
     * Structure is documented below.
     */
    public readonly defaultRouteAction!: pulumi.Output<outputs.compute.RegionUrlMapDefaultRouteAction | undefined>;
    /**
     * The full or partial URL of the defaultService resource to which traffic is directed if
     * none of the hostRules match. If defaultRouteAction is additionally specified, advanced
     * routing actions like URL Rewrites, etc. take effect prior to sending the request to the
     * backend. However, if defaultService is specified, defaultRouteAction cannot contain any
     * weightedBackendServices. Conversely, if routeAction specifies any
     * weightedBackendServices, service must not be specified.  Only one of defaultService,
     * defaultUrlRedirect or defaultRouteAction.weightedBackendService must be set.
     */
    public readonly defaultService!: pulumi.Output<string | undefined>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    public readonly defaultUrlRedirect!: pulumi.Output<outputs.compute.RegionUrlMapDefaultUrlRedirect | undefined>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Fingerprint of this resource. This field is used internally during
     * updates of this resource.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    public readonly hostRules!: pulumi.Output<outputs.compute.RegionUrlMapHostRule[] | undefined>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly mapId!: pulumi.Output<number>;
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
     * The list of named PathMatchers to use against the URL.
     * Structure is documented below.
     */
    public readonly pathMatchers!: pulumi.Output<outputs.compute.RegionUrlMapPathMatcher[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the url map should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The list of expected URL mappings. Requests to update this UrlMap will
     * succeed only if all of the test cases pass.
     * Structure is documented below.
     */
    public readonly tests!: pulumi.Output<outputs.compute.RegionUrlMapTest[] | undefined>;

    /**
     * Create a RegionUrlMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegionUrlMapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegionUrlMapArgs | RegionUrlMapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegionUrlMapState | undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["defaultRouteAction"] = state ? state.defaultRouteAction : undefined;
            resourceInputs["defaultService"] = state ? state.defaultService : undefined;
            resourceInputs["defaultUrlRedirect"] = state ? state.defaultUrlRedirect : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["hostRules"] = state ? state.hostRules : undefined;
            resourceInputs["mapId"] = state ? state.mapId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pathMatchers"] = state ? state.pathMatchers : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["tests"] = state ? state.tests : undefined;
        } else {
            const args = argsOrState as RegionUrlMapArgs | undefined;
            resourceInputs["defaultRouteAction"] = args ? args.defaultRouteAction : undefined;
            resourceInputs["defaultService"] = args ? args.defaultService : undefined;
            resourceInputs["defaultUrlRedirect"] = args ? args.defaultUrlRedirect : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostRules"] = args ? args.hostRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pathMatchers"] = args ? args.pathMatchers : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tests"] = args ? args.tests : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["mapId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionUrlMap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegionUrlMap resources.
 */
export interface RegionUrlMapState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.
     * defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
     * Structure is documented below.
     */
    defaultRouteAction?: pulumi.Input<inputs.compute.RegionUrlMapDefaultRouteAction>;
    /**
     * The full or partial URL of the defaultService resource to which traffic is directed if
     * none of the hostRules match. If defaultRouteAction is additionally specified, advanced
     * routing actions like URL Rewrites, etc. take effect prior to sending the request to the
     * backend. However, if defaultService is specified, defaultRouteAction cannot contain any
     * weightedBackendServices. Conversely, if routeAction specifies any
     * weightedBackendServices, service must not be specified.  Only one of defaultService,
     * defaultUrlRedirect or defaultRouteAction.weightedBackendService must be set.
     */
    defaultService?: pulumi.Input<string>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    defaultUrlRedirect?: pulumi.Input<inputs.compute.RegionUrlMapDefaultUrlRedirect>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. This field is used internally during
     * updates of this resource.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    hostRules?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapHostRule>[]>;
    /**
     * The unique identifier for the resource.
     */
    mapId?: pulumi.Input<number>;
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
     * The list of named PathMatchers to use against the URL.
     * Structure is documented below.
     */
    pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The Region in which the url map should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The list of expected URL mappings. Requests to update this UrlMap will
     * succeed only if all of the test cases pass.
     * Structure is documented below.
     */
    tests?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapTest>[]>;
}

/**
 * The set of arguments for constructing a RegionUrlMap resource.
 */
export interface RegionUrlMapArgs {
    /**
     * defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set. Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
     * Only one of defaultRouteAction or defaultUrlRedirect must be set.
     * URL maps for Classic external HTTP(S) load balancers only support the urlRewrite action within defaultRouteAction.
     * defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
     * Structure is documented below.
     */
    defaultRouteAction?: pulumi.Input<inputs.compute.RegionUrlMapDefaultRouteAction>;
    /**
     * The full or partial URL of the defaultService resource to which traffic is directed if
     * none of the hostRules match. If defaultRouteAction is additionally specified, advanced
     * routing actions like URL Rewrites, etc. take effect prior to sending the request to the
     * backend. However, if defaultService is specified, defaultRouteAction cannot contain any
     * weightedBackendServices. Conversely, if routeAction specifies any
     * weightedBackendServices, service must not be specified.  Only one of defaultService,
     * defaultUrlRedirect or defaultRouteAction.weightedBackendService must be set.
     */
    defaultService?: pulumi.Input<string>;
    /**
     * When none of the specified hostRules match, the request is redirected to a URL specified
     * by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
     * defaultRouteAction must not be set.
     * Structure is documented below.
     */
    defaultUrlRedirect?: pulumi.Input<inputs.compute.RegionUrlMapDefaultUrlRedirect>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The list of HostRules to use against the URL.
     * Structure is documented below.
     */
    hostRules?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapHostRule>[]>;
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
     * The list of named PathMatchers to use against the URL.
     * Structure is documented below.
     */
    pathMatchers?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapPathMatcher>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The Region in which the url map should reside.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The list of expected URL mappings. Requests to update this UrlMap will
     * succeed only if all of the test cases pass.
     * Structure is documented below.
     */
    tests?: pulumi.Input<pulumi.Input<inputs.compute.RegionUrlMapTest>[]>;
}
