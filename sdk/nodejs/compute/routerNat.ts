// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A NAT service created in a router.
 *
 * To get more information about RouterNat, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
 * * How-to Guides
 *     * [Google Cloud Router](https://cloud.google.com/router/docs/)
 *
 * ## Example Usage
 * ### Router Nat Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const net = new gcp.compute.Network("net", {});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     network: net.id,
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 * });
 * const router = new gcp.compute.Router("router", {
 *     region: subnet.region,
 *     network: net.id,
 *     bgp: {
 *         asn: 64514,
 *     },
 * });
 * const nat = new gcp.compute.RouterNat("nat", {
 *     router: router.name,
 *     region: router.region,
 *     natIpAllocateOption: "AUTO_ONLY",
 *     sourceSubnetworkIpRangesToNat: "ALL_SUBNETWORKS_ALL_IP_RANGES",
 *     logConfig: {
 *         enable: true,
 *         filter: "ERRORS_ONLY",
 *     },
 * });
 * ```
 * ### Router Nat Manual Ips
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const net = new gcp.compute.Network("net", {});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     network: net.id,
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 * });
 * const router = new gcp.compute.Router("router", {
 *     region: subnet.region,
 *     network: net.id,
 * });
 * const address: gcp.compute.Address[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     address.push(new gcp.compute.Address(`address-${range.value}`, {region: subnet.region}));
 * }
 * const natManual = new gcp.compute.RouterNat("natManual", {
 *     router: router.name,
 *     region: router.region,
 *     natIpAllocateOption: "MANUAL_ONLY",
 *     natIps: address.map(__item => __item.selfLink),
 *     sourceSubnetworkIpRangesToNat: "LIST_OF_SUBNETWORKS",
 *     subnetworks: [{
 *         name: subnet.id,
 *         sourceIpRangesToNats: ["ALL_IP_RANGES"],
 *     }],
 * });
 * ```
 * ### Router Nat Rules
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const net = new gcp.compute.Network("net", {autoCreateSubnetworks: false});
 * const subnet = new gcp.compute.Subnetwork("subnet", {
 *     network: net.id,
 *     ipCidrRange: "10.0.0.0/16",
 *     region: "us-central1",
 * });
 * const router = new gcp.compute.Router("router", {
 *     region: subnet.region,
 *     network: net.id,
 * });
 * const addr1 = new gcp.compute.Address("addr1", {region: subnet.region});
 * const addr2 = new gcp.compute.Address("addr2", {region: subnet.region});
 * const addr3 = new gcp.compute.Address("addr3", {region: subnet.region});
 * const natRules = new gcp.compute.RouterNat("natRules", {
 *     router: router.name,
 *     region: router.region,
 *     natIpAllocateOption: "MANUAL_ONLY",
 *     natIps: [addr1.selfLink],
 *     sourceSubnetworkIpRangesToNat: "LIST_OF_SUBNETWORKS",
 *     subnetworks: [{
 *         name: subnet.id,
 *         sourceIpRangesToNats: ["ALL_IP_RANGES"],
 *     }],
 *     rules: [{
 *         ruleNumber: 100,
 *         description: "nat rules example",
 *         match: "inIpRange(destination.ip, '1.1.0.0/16') || inIpRange(destination.ip, '2.2.0.0/16')",
 *         action: {
 *             sourceNatActiveIps: [
 *                 addr2.selfLink,
 *                 addr3.selfLink,
 *             ],
 *         },
 *     }],
 *     enableEndpointIndependentMapping: false,
 * });
 * ```
 *
 * ## Import
 *
 * RouterNat can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default projects/{{project}}/regions/{{region}}/routers/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{project}}/{{region}}/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{region}}/{{router}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerNat:RouterNat default {{router}}/{{name}}
 * ```
 */
export class RouterNat extends pulumi.CustomResource {
    /**
     * Get an existing RouterNat resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterNatState, opts?: pulumi.CustomResourceOptions): RouterNat {
        return new RouterNat(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/routerNat:RouterNat';

    /**
     * Returns true if the given object is an instance of RouterNat.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterNat {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterNat.__pulumiType;
    }

    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    public readonly drainNatIps!: pulumi.Output<string[] | undefined>;
    /**
     * Enable Dynamic Port Allocation.
     * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
     * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
     * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
     * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
     * Mutually exclusive with enableEndpointIndependentMapping.
     */
    public readonly enableDynamicPortAllocation!: pulumi.Output<boolean>;
    /**
     * Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
     * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
     */
    public readonly enableEndpointIndependentMapping!: pulumi.Output<boolean | undefined>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    public readonly icmpIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.RouterNatLogConfig | undefined>;
    /**
     * Maximum number of ports allocated to a VM from this NAT.
     * This field can only be set when enableDynamicPortAllocation is enabled.
     */
    public readonly maxPortsPerVm!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    public readonly minPortsPerVm!: pulumi.Output<number | undefined>;
    /**
     * Name of the NAT service. The name must be 1-63 characters long and
     * comply with RFC1035.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are: `MANUAL_ONLY`, `AUTO_ONLY`.
     */
    public readonly natIpAllocateOption!: pulumi.Output<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    public readonly natIps!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where the router and NAT reside.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     *
     *
     * - - -
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * A list of rules associated with this NAT.
     * Structure is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.compute.RouterNatRule[] | undefined>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are: `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, `LIST_OF_SUBNETWORKS`.
     */
    public readonly sourceSubnetworkIpRangesToNat!: pulumi.Output<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    public readonly subnetworks!: pulumi.Output<outputs.compute.RouterNatSubnetwork[] | undefined>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    public readonly tcpEstablishedIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
     * Defaults to 120s if not set.
     */
    public readonly tcpTimeWaitTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    public readonly tcpTransitoryIdleTimeoutSec!: pulumi.Output<number | undefined>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    public readonly udpIdleTimeoutSec!: pulumi.Output<number | undefined>;

    /**
     * Create a RouterNat resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterNatArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterNatArgs | RouterNatState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterNatState | undefined;
            resourceInputs["drainNatIps"] = state ? state.drainNatIps : undefined;
            resourceInputs["enableDynamicPortAllocation"] = state ? state.enableDynamicPortAllocation : undefined;
            resourceInputs["enableEndpointIndependentMapping"] = state ? state.enableEndpointIndependentMapping : undefined;
            resourceInputs["icmpIdleTimeoutSec"] = state ? state.icmpIdleTimeoutSec : undefined;
            resourceInputs["logConfig"] = state ? state.logConfig : undefined;
            resourceInputs["maxPortsPerVm"] = state ? state.maxPortsPerVm : undefined;
            resourceInputs["minPortsPerVm"] = state ? state.minPortsPerVm : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["natIpAllocateOption"] = state ? state.natIpAllocateOption : undefined;
            resourceInputs["natIps"] = state ? state.natIps : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["router"] = state ? state.router : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["sourceSubnetworkIpRangesToNat"] = state ? state.sourceSubnetworkIpRangesToNat : undefined;
            resourceInputs["subnetworks"] = state ? state.subnetworks : undefined;
            resourceInputs["tcpEstablishedIdleTimeoutSec"] = state ? state.tcpEstablishedIdleTimeoutSec : undefined;
            resourceInputs["tcpTimeWaitTimeoutSec"] = state ? state.tcpTimeWaitTimeoutSec : undefined;
            resourceInputs["tcpTransitoryIdleTimeoutSec"] = state ? state.tcpTransitoryIdleTimeoutSec : undefined;
            resourceInputs["udpIdleTimeoutSec"] = state ? state.udpIdleTimeoutSec : undefined;
        } else {
            const args = argsOrState as RouterNatArgs | undefined;
            if ((!args || args.natIpAllocateOption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natIpAllocateOption'");
            }
            if ((!args || args.router === undefined) && !opts.urn) {
                throw new Error("Missing required property 'router'");
            }
            if ((!args || args.sourceSubnetworkIpRangesToNat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceSubnetworkIpRangesToNat'");
            }
            resourceInputs["drainNatIps"] = args ? args.drainNatIps : undefined;
            resourceInputs["enableDynamicPortAllocation"] = args ? args.enableDynamicPortAllocation : undefined;
            resourceInputs["enableEndpointIndependentMapping"] = args ? args.enableEndpointIndependentMapping : undefined;
            resourceInputs["icmpIdleTimeoutSec"] = args ? args.icmpIdleTimeoutSec : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["maxPortsPerVm"] = args ? args.maxPortsPerVm : undefined;
            resourceInputs["minPortsPerVm"] = args ? args.minPortsPerVm : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["natIpAllocateOption"] = args ? args.natIpAllocateOption : undefined;
            resourceInputs["natIps"] = args ? args.natIps : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["router"] = args ? args.router : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["sourceSubnetworkIpRangesToNat"] = args ? args.sourceSubnetworkIpRangesToNat : undefined;
            resourceInputs["subnetworks"] = args ? args.subnetworks : undefined;
            resourceInputs["tcpEstablishedIdleTimeoutSec"] = args ? args.tcpEstablishedIdleTimeoutSec : undefined;
            resourceInputs["tcpTimeWaitTimeoutSec"] = args ? args.tcpTimeWaitTimeoutSec : undefined;
            resourceInputs["tcpTransitoryIdleTimeoutSec"] = args ? args.tcpTransitoryIdleTimeoutSec : undefined;
            resourceInputs["udpIdleTimeoutSec"] = args ? args.udpIdleTimeoutSec : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterNat.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterNat resources.
 */
export interface RouterNatState {
    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    drainNatIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable Dynamic Port Allocation.
     * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
     * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
     * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
     * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
     * Mutually exclusive with enableEndpointIndependentMapping.
     */
    enableDynamicPortAllocation?: pulumi.Input<boolean>;
    /**
     * Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
     * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
     */
    enableEndpointIndependentMapping?: pulumi.Input<boolean>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    icmpIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.RouterNatLogConfig>;
    /**
     * Maximum number of ports allocated to a VM from this NAT.
     * This field can only be set when enableDynamicPortAllocation is enabled.
     */
    maxPortsPerVm?: pulumi.Input<number>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    minPortsPerVm?: pulumi.Input<number>;
    /**
     * Name of the NAT service. The name must be 1-63 characters long and
     * comply with RFC1035.
     */
    name?: pulumi.Input<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are: `MANUAL_ONLY`, `AUTO_ONLY`.
     */
    natIpAllocateOption?: pulumi.Input<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    natIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region where the router and NAT reside.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     *
     *
     * - - -
     */
    router?: pulumi.Input<string>;
    /**
     * A list of rules associated with this NAT.
     * Structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatRule>[]>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are: `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, `LIST_OF_SUBNETWORKS`.
     */
    sourceSubnetworkIpRangesToNat?: pulumi.Input<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    subnetworks?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatSubnetwork>[]>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
     * Defaults to 120s if not set.
     */
    tcpTimeWaitTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    udpIdleTimeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RouterNat resource.
 */
export interface RouterNatArgs {
    /**
     * A list of URLs of the IP resources to be drained. These IPs must be
     * valid static external IPs that have been assigned to the NAT.
     */
    drainNatIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable Dynamic Port Allocation.
     * If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
     * If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
     * If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
     * If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
     * Mutually exclusive with enableEndpointIndependentMapping.
     */
    enableDynamicPortAllocation?: pulumi.Input<boolean>;
    /**
     * Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
     * see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
     */
    enableEndpointIndependentMapping?: pulumi.Input<boolean>;
    /**
     * Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
     */
    icmpIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Configuration for logging on NAT
     * Structure is documented below.
     */
    logConfig?: pulumi.Input<inputs.compute.RouterNatLogConfig>;
    /**
     * Maximum number of ports allocated to a VM from this NAT.
     * This field can only be set when enableDynamicPortAllocation is enabled.
     */
    maxPortsPerVm?: pulumi.Input<number>;
    /**
     * Minimum number of ports allocated to a VM from this NAT.
     */
    minPortsPerVm?: pulumi.Input<number>;
    /**
     * Name of the NAT service. The name must be 1-63 characters long and
     * comply with RFC1035.
     */
    name?: pulumi.Input<string>;
    /**
     * How external IPs should be allocated for this NAT. Valid values are
     * `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
     * Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
     * Possible values are: `MANUAL_ONLY`, `AUTO_ONLY`.
     */
    natIpAllocateOption: pulumi.Input<string>;
    /**
     * Self-links of NAT IPs. Only valid if natIpAllocateOption
     * is set to MANUAL_ONLY.
     */
    natIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region where the router and NAT reside.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of the Cloud Router in which this NAT will be configured.
     *
     *
     * - - -
     */
    router: pulumi.Input<string>;
    /**
     * A list of rules associated with this NAT.
     * Structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatRule>[]>;
    /**
     * How NAT should be configured per Subnetwork.
     * If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
     * IP ranges in every Subnetwork are allowed to Nat.
     * If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
     * ranges in every Subnetwork are allowed to Nat.
     * `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
     * (specified in the field subnetwork below). Note that if this field
     * contains ALL_SUBNETWORKS_ALL_IP_RANGES or
     * ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
     * other RouterNat section in any Router for this network in this region.
     * Possible values are: `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, `LIST_OF_SUBNETWORKS`.
     */
    sourceSubnetworkIpRangesToNat: pulumi.Input<string>;
    /**
     * One or more subnetwork NAT configurations. Only used if
     * `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`
     * Structure is documented below.
     */
    subnetworks?: pulumi.Input<pulumi.Input<inputs.compute.RouterNatSubnetwork>[]>;
    /**
     * Timeout (in seconds) for TCP established connections.
     * Defaults to 1200s if not set.
     */
    tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
     * Defaults to 120s if not set.
     */
    tcpTimeWaitTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for TCP transitory connections.
     * Defaults to 30s if not set.
     */
    tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    /**
     * Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
     */
    udpIdleTimeoutSec?: pulumi.Input<number>;
}
