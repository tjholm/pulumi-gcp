// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a Route resource.
 *
 * A route is a rule that specifies how certain packets should be handled by
 * the virtual network. Routes are associated with virtual machines by tag,
 * and the set of routes for a particular virtual machine is called its
 * routing table. For each packet leaving a virtual machine, the system
 * searches that virtual machine's routing table for a single best matching
 * route.
 *
 * Routes match packets by destination IP address, preferring smaller or more
 * specific ranges over larger ones. If there is a tie, the system selects
 * the route with the smallest priority value. If there is still a tie, it
 * uses the layer three and four packet headers to select just one of the
 * remaining matching routes. The packet is then forwarded as specified by
 * the nextHop field of the winning route -- either to another virtual
 * machine destination, a virtual machine gateway or a Compute
 * Engine-operated gateway. Packets that do not match any route in the
 * sending virtual machine's routing table will be dropped.
 *
 * A Route resource must have exactly one specification of either
 * nextHopGateway, nextHopInstance, nextHopIp, nextHopVpnTunnel, or
 * nextHopIlb.
 *
 * To get more information about Route, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routes)
 * * How-to Guides
 *     * [Using Routes](https://cloud.google.com/vpc/docs/using-routes)
 *
 * ## Example Usage
 * ### Route Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultRoute = new gcp.compute.Route("defaultRoute", {
 *     destRange: "15.0.0.0/24",
 *     network: defaultNetwork.name,
 *     nextHopIp: "10.132.1.5",
 *     priority: 100,
 * });
 * ```
 *
 * ## Import
 *
 * Route can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default projects/{{project}}/global/routes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default {{name}}
 * ```
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    public readonly destRange!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network that this route applies to.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    public readonly nextHopGateway!: pulumi.Output<string | undefined>;
    /**
     * The IP address or URL to a forwarding rule of type
     * loadBalancingScheme=INTERNAL that should handle matching
     * packets.
     * With the GA provider you can only specify the forwarding
     * rule as a partial or full URL. For example, the following
     * are all valid values:
     * * 10.128.0.56
     * * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * * regions/region/forwardingRules/forwardingRule
     * When the beta provider, you can also specify the IP address
     * of a forwarding rule from the same VPC or any peered VPC.
     * Note that this can only be used when the destinationRange is
     * a public (non-RFC 1918) IP CIDR range.
     */
    public readonly nextHopIlb!: pulumi.Output<string | undefined>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    public readonly nextHopInstance!: pulumi.Output<string | undefined>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    public readonly nextHopInstanceZone!: pulumi.Output<string>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    public readonly nextHopIp!: pulumi.Output<string>;
    /**
     * URL to a Network that should handle matching packets.
     */
    public /*out*/ readonly nextHopNetwork!: pulumi.Output<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    public readonly nextHopVpnTunnel!: pulumi.Output<string | undefined>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destRange"] = state ? state.destRange : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["nextHopGateway"] = state ? state.nextHopGateway : undefined;
            resourceInputs["nextHopIlb"] = state ? state.nextHopIlb : undefined;
            resourceInputs["nextHopInstance"] = state ? state.nextHopInstance : undefined;
            resourceInputs["nextHopInstanceZone"] = state ? state.nextHopInstanceZone : undefined;
            resourceInputs["nextHopIp"] = state ? state.nextHopIp : undefined;
            resourceInputs["nextHopNetwork"] = state ? state.nextHopNetwork : undefined;
            resourceInputs["nextHopVpnTunnel"] = state ? state.nextHopVpnTunnel : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if ((!args || args.destRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destRange'");
            }
            if ((!args || args.network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'network'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destRange"] = args ? args.destRange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nextHopGateway"] = args ? args.nextHopGateway : undefined;
            resourceInputs["nextHopIlb"] = args ? args.nextHopIlb : undefined;
            resourceInputs["nextHopInstance"] = args ? args.nextHopInstance : undefined;
            resourceInputs["nextHopInstanceZone"] = args ? args.nextHopInstanceZone : undefined;
            resourceInputs["nextHopIp"] = args ? args.nextHopIp : undefined;
            resourceInputs["nextHopVpnTunnel"] = args ? args.nextHopVpnTunnel : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["nextHopNetwork"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Route.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    destRange?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The network that this route applies to.
     */
    network?: pulumi.Input<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    nextHopGateway?: pulumi.Input<string>;
    /**
     * The IP address or URL to a forwarding rule of type
     * loadBalancingScheme=INTERNAL that should handle matching
     * packets.
     * With the GA provider you can only specify the forwarding
     * rule as a partial or full URL. For example, the following
     * are all valid values:
     * * 10.128.0.56
     * * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * * regions/region/forwardingRules/forwardingRule
     * When the beta provider, you can also specify the IP address
     * of a forwarding rule from the same VPC or any peered VPC.
     * Note that this can only be used when the destinationRange is
     * a public (non-RFC 1918) IP CIDR range.
     */
    nextHopIlb?: pulumi.Input<string>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    nextHopInstance?: pulumi.Input<string>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    nextHopInstanceZone?: pulumi.Input<string>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    nextHopIp?: pulumi.Input<string>;
    /**
     * URL to a Network that should handle matching packets.
     */
    nextHopNetwork?: pulumi.Input<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    nextHopVpnTunnel?: pulumi.Input<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     */
    destRange: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The network that this route applies to.
     */
    network: pulumi.Input<string>;
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     */
    nextHopGateway?: pulumi.Input<string>;
    /**
     * The IP address or URL to a forwarding rule of type
     * loadBalancingScheme=INTERNAL that should handle matching
     * packets.
     * With the GA provider you can only specify the forwarding
     * rule as a partial or full URL. For example, the following
     * are all valid values:
     * * 10.128.0.56
     * * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * * regions/region/forwardingRules/forwardingRule
     * When the beta provider, you can also specify the IP address
     * of a forwarding rule from the same VPC or any peered VPC.
     * Note that this can only be used when the destinationRange is
     * a public (non-RFC 1918) IP CIDR range.
     */
    nextHopIlb?: pulumi.Input<string>;
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `nextHopInstanceZone`.
     */
    nextHopInstance?: pulumi.Input<string>;
    /**
     * (Optional when `nextHopInstance` is
     * specified)  The zone of the instance specified in
     * `nextHopInstance`.  Omit if `nextHopInstance` is specified as
     * a URL.
     */
    nextHopInstanceZone?: pulumi.Input<string>;
    /**
     * Network IP address of an instance that should handle matching packets.
     */
    nextHopIp?: pulumi.Input<string>;
    /**
     * URL to a VpnTunnel that should handle matching packets.
     */
    nextHopVpnTunnel?: pulumi.Input<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     */
    priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
