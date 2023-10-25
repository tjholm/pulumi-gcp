// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a Global Address resource. Global addresses are used for
 * HTTP(S) load balancing.
 *
 * To get more information about GlobalAddress, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
 *
 * ## Example Usage
 * ### Global Address Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.compute.GlobalAddress("default", {});
 * ```
 * ### Global Address Private Services Connect
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const network = new gcp.compute.Network("network", {autoCreateSubnetworks: false}, {
 *     provider: google_beta,
 * });
 * const _default = new gcp.compute.GlobalAddress("default", {
 *     addressType: "INTERNAL",
 *     purpose: "PRIVATE_SERVICE_CONNECT",
 *     network: network.id,
 *     address: "100.100.100.105",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * GlobalAddress can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalAddress:GlobalAddress default projects/{{project}}/global/addresses/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{name}}
 * ```
 */
export class GlobalAddress extends pulumi.CustomResource {
    /**
     * Get an existing GlobalAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalAddressState, opts?: pulumi.CustomResourceOptions): GlobalAddress {
        return new GlobalAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/globalAddress:GlobalAddress';

    /**
     * Returns true if the given object is an instance of GlobalAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalAddress.__pulumiType;
    }

    /**
     * The IP address or beginning of the address range represented by this
     * resource. This can be supplied as an input to reserve a specific
     * address or omitted to allow GCP to choose a valid one for you.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The type of the address to reserve.
     * * EXTERNAL indicates public/external single IP address.
     * * INTERNAL indicates internal IP ranges belonging to some network.
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL`.
     */
    public readonly addressType!: pulumi.Output<string | undefined>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP Version that will be used by this address. The default value is `IPV4`.
     * Possible values are: `IPV4`, `IPV6`.
     */
    public readonly ipVersion!: pulumi.Output<string | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
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
     * The URL of the network in which to reserve the IP range. The IP range
     * must be in RFC1918 space. The network cannot be deleted if there are
     * any reserved IP ranges referring to it.
     * This should only be set when using an Internal address.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * The prefix length of the IP range. If not present, it means the
     * address field is a single IP address.
     * This field is not applicable to addresses with addressType=INTERNAL
     * when purpose=PRIVATE_SERVICE_CONNECT
     */
    public readonly prefixLength!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The purpose of the resource. Possible values include:
     * * VPC_PEERING - for peer networks
     * * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
     */
    public readonly purpose!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a GlobalAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalAddressArgs | GlobalAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalAddressState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["addressType"] = state ? state.addressType : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["prefixLength"] = state ? state.prefixLength : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["purpose"] = state ? state.purpose : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as GlobalAddressArgs | undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["addressType"] = args ? args.addressType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["prefixLength"] = args ? args.prefixLength : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["purpose"] = args ? args.purpose : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GlobalAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalAddress resources.
 */
export interface GlobalAddressState {
    /**
     * The IP address or beginning of the address range represented by this
     * resource. This can be supplied as an input to reserve a specific
     * address or omitted to allow GCP to choose a valid one for you.
     */
    address?: pulumi.Input<string>;
    /**
     * The type of the address to reserve.
     * * EXTERNAL indicates public/external single IP address.
     * * INTERNAL indicates internal IP ranges belonging to some network.
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL`.
     */
    addressType?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP Version that will be used by this address. The default value is `IPV4`.
     * Possible values are: `IPV4`, `IPV6`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     */
    labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
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
     * The URL of the network in which to reserve the IP range. The IP range
     * must be in RFC1918 space. The network cannot be deleted if there are
     * any reserved IP ranges referring to it.
     * This should only be set when using an Internal address.
     */
    network?: pulumi.Input<string>;
    /**
     * The prefix length of the IP range. If not present, it means the
     * address field is a single IP address.
     * This field is not applicable to addresses with addressType=INTERNAL
     * when purpose=PRIVATE_SERVICE_CONNECT
     */
    prefixLength?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The purpose of the resource. Possible values include:
     * * VPC_PEERING - for peer networks
     * * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
     */
    purpose?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GlobalAddress resource.
 */
export interface GlobalAddressArgs {
    /**
     * The IP address or beginning of the address range represented by this
     * resource. This can be supplied as an input to reserve a specific
     * address or omitted to allow GCP to choose a valid one for you.
     */
    address?: pulumi.Input<string>;
    /**
     * The type of the address to reserve.
     * * EXTERNAL indicates public/external single IP address.
     * * INTERNAL indicates internal IP ranges belonging to some network.
     * Default value is `EXTERNAL`.
     * Possible values are: `EXTERNAL`, `INTERNAL`.
     */
    addressType?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP Version that will be used by this address. The default value is `IPV4`.
     * Possible values are: `IPV4`, `IPV6`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Labels to apply to this address.  A list of key->value pairs.
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
     * The URL of the network in which to reserve the IP range. The IP range
     * must be in RFC1918 space. The network cannot be deleted if there are
     * any reserved IP ranges referring to it.
     * This should only be set when using an Internal address.
     */
    network?: pulumi.Input<string>;
    /**
     * The prefix length of the IP range. If not present, it means the
     * address field is a single IP address.
     * This field is not applicable to addresses with addressType=INTERNAL
     * when purpose=PRIVATE_SERVICE_CONNECT
     */
    prefixLength?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The purpose of the resource. Possible values include:
     * * VPC_PEERING - for peer networks
     * * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
     */
    purpose?: pulumi.Input<string>;
}
