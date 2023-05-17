// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Cloud Router interface. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/routers).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const foobar = new gcp.compute.RouterInterface("foobar", {
 *     ipRange: "169.254.1.1/30",
 *     region: "us-central1",
 *     router: "router-1",
 *     vpnTunnel: "tunnel-1",
 * });
 * ```
 *
 * ## Import
 *
 * Router interfaces can be imported using the `project` (optional), `region`, `router`, and `name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerInterface:RouterInterface foobar my-project/us-central1/router-1/interface-1
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/routerInterface:RouterInterface foobar us-central1/router-1/interface-1
 * ```
 */
export class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceState, opts?: pulumi.CustomResourceOptions): RouterInterface {
        return new RouterInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/routerInterface:RouterInterface';

    /**
     * Returns true if the given object is an instance of RouterInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterInterface.__pulumiType;
    }

    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    public readonly interconnectAttachment!: pulumi.Output<string | undefined>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    public readonly ipRange!: pulumi.Output<string>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The regional private internal IP address that is used
     * to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
     */
    public readonly privateIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which this interface's routerbelongs.
     * If it is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the interface that is redundant to
     * this interface. Changing this forces a new interface to be created.
     */
    public readonly redundantInterface!: pulumi.Output<string>;
    /**
     * The region this interface's router sits in.
     * If not specified, the project region will be used. Changing this forces a new interface to be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * The URI of the subnetwork resource that this interface
     * belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    public readonly subnetwork!: pulumi.Output<string | undefined>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    public readonly vpnTunnel!: pulumi.Output<string | undefined>;

    /**
     * Create a RouterInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterInterfaceArgs | RouterInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterInterfaceState | undefined;
            resourceInputs["interconnectAttachment"] = state ? state.interconnectAttachment : undefined;
            resourceInputs["ipRange"] = state ? state.ipRange : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["redundantInterface"] = state ? state.redundantInterface : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["router"] = state ? state.router : undefined;
            resourceInputs["subnetwork"] = state ? state.subnetwork : undefined;
            resourceInputs["vpnTunnel"] = state ? state.vpnTunnel : undefined;
        } else {
            const args = argsOrState as RouterInterfaceArgs | undefined;
            if ((!args || args.router === undefined) && !opts.urn) {
                throw new Error("Missing required property 'router'");
            }
            resourceInputs["interconnectAttachment"] = args ? args.interconnectAttachment : undefined;
            resourceInputs["ipRange"] = args ? args.ipRange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["redundantInterface"] = args ? args.redundantInterface : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["router"] = args ? args.router : undefined;
            resourceInputs["subnetwork"] = args ? args.subnetwork : undefined;
            resourceInputs["vpnTunnel"] = args ? args.vpnTunnel : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouterInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterInterface resources.
 */
export interface RouterInterfaceState {
    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    interconnectAttachment?: pulumi.Input<string>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The regional private internal IP address that is used
     * to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the project in which this interface's routerbelongs.
     * If it is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the interface that is redundant to
     * this interface. Changing this forces a new interface to be created.
     */
    redundantInterface?: pulumi.Input<string>;
    /**
     * The region this interface's router sits in.
     * If not specified, the project region will be used. Changing this forces a new interface to be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    router?: pulumi.Input<string>;
    /**
     * The URI of the subnetwork resource that this interface
     * belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    subnetwork?: pulumi.Input<string>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    vpnTunnel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterInterface resource.
 */
export interface RouterInterfaceArgs {
    /**
     * The name or resource link to the
     * VLAN interconnect for this interface. Changing this forces a new interface to
     * be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    interconnectAttachment?: pulumi.Input<string>;
    /**
     * IP address and range of the interface. The IP range must be
     * in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
     */
    ipRange?: pulumi.Input<string>;
    /**
     * A unique name for the interface, required by GCE. Changing
     * this forces a new interface to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The regional private internal IP address that is used
     * to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The ID of the project in which this interface's routerbelongs.
     * If it is not provided, the provider project is used. Changing this forces a new interface to be created.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the interface that is redundant to
     * this interface. Changing this forces a new interface to be created.
     */
    redundantInterface?: pulumi.Input<string>;
    /**
     * The region this interface's router sits in.
     * If not specified, the project region will be used. Changing this forces a new interface to be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The name of the router this interface will be attached to.
     * Changing this forces a new interface to be created.
     */
    router: pulumi.Input<string>;
    /**
     * The URI of the subnetwork resource that this interface
     * belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    subnetwork?: pulumi.Input<string>;
    /**
     * The name or resource link to the VPN tunnel this
     * interface will be linked to. Changing this forces a new interface to be created. Only
     * one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
     */
    vpnTunnel?: pulumi.Input<string>;
}
