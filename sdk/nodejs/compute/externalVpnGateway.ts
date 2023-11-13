// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a VPN gateway managed outside of GCP.
 *
 * To get more information about ExternalVpnGateway, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)
 *
 * ## Example Usage
 * ### External Vpn Gateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const network = new gcp.compute.Network("network", {
 *     routingMode: "GLOBAL",
 *     autoCreateSubnetworks: false,
 * });
 * const haGateway = new gcp.compute.HaVpnGateway("haGateway", {
 *     region: "us-central1",
 *     network: network.id,
 * });
 * const externalGateway = new gcp.compute.ExternalVpnGateway("externalGateway", {
 *     redundancyType: "SINGLE_IP_INTERNALLY_REDUNDANT",
 *     description: "An externally managed VPN gateway",
 *     interfaces: [{
 *         id: 0,
 *         ipAddress: "8.8.8.8",
 *     }],
 * });
 * const networkSubnet1 = new gcp.compute.Subnetwork("networkSubnet1", {
 *     ipCidrRange: "10.0.1.0/24",
 *     region: "us-central1",
 *     network: network.id,
 * });
 * const networkSubnet2 = new gcp.compute.Subnetwork("networkSubnet2", {
 *     ipCidrRange: "10.0.2.0/24",
 *     region: "us-west1",
 *     network: network.id,
 * });
 * const router1 = new gcp.compute.Router("router1", {
 *     network: network.name,
 *     bgp: {
 *         asn: 64514,
 *     },
 * });
 * const tunnel1 = new gcp.compute.VPNTunnel("tunnel1", {
 *     region: "us-central1",
 *     vpnGateway: haGateway.id,
 *     peerExternalGateway: externalGateway.id,
 *     peerExternalGatewayInterface: 0,
 *     sharedSecret: "a secret message",
 *     router: router1.id,
 *     vpnGatewayInterface: 0,
 * });
 * const tunnel2 = new gcp.compute.VPNTunnel("tunnel2", {
 *     region: "us-central1",
 *     vpnGateway: haGateway.id,
 *     peerExternalGateway: externalGateway.id,
 *     peerExternalGatewayInterface: 0,
 *     sharedSecret: "a secret message",
 *     router: pulumi.interpolate` ${router1.id}`,
 *     vpnGatewayInterface: 1,
 * });
 * const router1Interface1 = new gcp.compute.RouterInterface("router1Interface1", {
 *     router: router1.name,
 *     region: "us-central1",
 *     ipRange: "169.254.0.1/30",
 *     vpnTunnel: tunnel1.name,
 * });
 * const router1Peer1 = new gcp.compute.RouterPeer("router1Peer1", {
 *     router: router1.name,
 *     region: "us-central1",
 *     peerIpAddress: "169.254.0.2",
 *     peerAsn: 64515,
 *     advertisedRoutePriority: 100,
 *     "interface": router1Interface1.name,
 * });
 * const router1Interface2 = new gcp.compute.RouterInterface("router1Interface2", {
 *     router: router1.name,
 *     region: "us-central1",
 *     ipRange: "169.254.1.1/30",
 *     vpnTunnel: tunnel2.name,
 * });
 * const router1Peer2 = new gcp.compute.RouterPeer("router1Peer2", {
 *     router: router1.name,
 *     region: "us-central1",
 *     peerIpAddress: "169.254.1.2",
 *     peerAsn: 64515,
 *     advertisedRoutePriority: 100,
 *     "interface": router1Interface2.name,
 * });
 * ```
 *
 * ## Import
 *
 * ExternalVpnGateway can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
 * ```
 */
export class ExternalVpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing ExternalVpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalVpnGatewayState, opts?: pulumi.CustomResourceOptions): ExternalVpnGateway {
        return new ExternalVpnGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/externalVpnGateway:ExternalVpnGateway';

    /**
     * Returns true if the given object is an instance of ExternalVpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalVpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalVpnGateway.__pulumiType;
    }

    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of interfaces on this external VPN gateway.
     * Structure is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.compute.ExternalVpnGatewayInterface[] | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource.  Used
     * internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels for the external VPN gateway resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Indicates the redundancy type of this external VPN gateway
     * Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
     */
    public readonly redundancyType!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a ExternalVpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExternalVpnGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalVpnGatewayArgs | ExternalVpnGatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalVpnGatewayState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["redundancyType"] = state ? state.redundancyType : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as ExternalVpnGatewayArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["redundancyType"] = args ? args.redundancyType : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ExternalVpnGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalVpnGateway resources.
 */
export interface ExternalVpnGatewayState {
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of interfaces on this external VPN gateway.
     * Structure is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.compute.ExternalVpnGatewayInterface>[]>;
    /**
     * The fingerprint used for optimistic locking of this resource.  Used
     * internally during updates.
     */
    labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels for the external VPN gateway resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the redundancy type of this external VPN gateway
     * Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
     */
    redundancyType?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalVpnGateway resource.
 */
export interface ExternalVpnGatewayArgs {
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of interfaces on this external VPN gateway.
     * Structure is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.compute.ExternalVpnGatewayInterface>[]>;
    /**
     * Labels for the external VPN gateway resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Indicates the redundancy type of this external VPN gateway
     * Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
     */
    redundancyType?: pulumi.Input<string>;
}
