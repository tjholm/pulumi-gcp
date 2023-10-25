// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The PrivateConnection resource is used to establish private connectivity between Datastream and a customer's network.
 *
 * To get more information about PrivateConnection, see:
 *
 * * [API documentation](https://cloud.google.com/datastream/docs/reference/rest/v1/projects.locations.privateConnections)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/datastream/docs/create-a-private-connectivity-configuration)
 *
 * ## Example Usage
 * ### Datastream Private Connection Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {});
 * const defaultPrivateConnection = new gcp.datastream.PrivateConnection("defaultPrivateConnection", {
 *     displayName: "Connection profile",
 *     location: "us-central1",
 *     privateConnectionId: "my-connection",
 *     labels: {
 *         key: "value",
 *     },
 *     vpcPeeringConfig: {
 *         vpc: defaultNetwork.id,
 *         subnet: "10.0.0.0/29",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * PrivateConnection can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default {{project}}/{{location}}/{{private_connection_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastream/privateConnection:PrivateConnection default {{location}}/{{private_connection_id}}
 * ```
 */
export class PrivateConnection extends pulumi.CustomResource {
    /**
     * Get an existing PrivateConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateConnectionState, opts?: pulumi.CustomResourceOptions): PrivateConnection {
        return new PrivateConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datastream/privateConnection:PrivateConnection';

    /**
     * Returns true if the given object is an instance of PrivateConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateConnection.__pulumiType;
    }

    /**
     * Display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The PrivateConnection error in case of failure.
     * Structure is documented below.
     */
    public /*out*/ readonly errors!: pulumi.Output<outputs.datastream.PrivateConnectionError[]>;
    /**
     * Labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the location this private connection is located in.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource's name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The private connectivity identifier.
     */
    public readonly privateConnectionId!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * State of the PrivateConnection.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The VPC Peering configuration is used to create VPC peering
     * between Datastream and the consumer's VPC.
     * Structure is documented below.
     */
    public readonly vpcPeeringConfig!: pulumi.Output<outputs.datastream.PrivateConnectionVpcPeeringConfig>;

    /**
     * Create a PrivateConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateConnectionArgs | PrivateConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateConnectionState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["errors"] = state ? state.errors : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateConnectionId"] = state ? state.privateConnectionId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["vpcPeeringConfig"] = state ? state.vpcPeeringConfig : undefined;
        } else {
            const args = argsOrState as PrivateConnectionArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.privateConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateConnectionId'");
            }
            if ((!args || args.vpcPeeringConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcPeeringConfig'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["privateConnectionId"] = args ? args.privateConnectionId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["vpcPeeringConfig"] = args ? args.vpcPeeringConfig : undefined;
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateConnection resources.
 */
export interface PrivateConnectionState {
    /**
     * Display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The PrivateConnection error in case of failure.
     * Structure is documented below.
     */
    errors?: pulumi.Input<pulumi.Input<inputs.datastream.PrivateConnectionError>[]>;
    /**
     * Labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this private connection is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The resource's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The private connectivity identifier.
     */
    privateConnectionId?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * State of the PrivateConnection.
     */
    state?: pulumi.Input<string>;
    /**
     * The VPC Peering configuration is used to create VPC peering
     * between Datastream and the consumer's VPC.
     * Structure is documented below.
     */
    vpcPeeringConfig?: pulumi.Input<inputs.datastream.PrivateConnectionVpcPeeringConfig>;
}

/**
 * The set of arguments for constructing a PrivateConnection resource.
 */
export interface PrivateConnectionArgs {
    /**
     * Display name.
     */
    displayName: pulumi.Input<string>;
    /**
     * Labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this private connection is located in.
     */
    location: pulumi.Input<string>;
    /**
     * The private connectivity identifier.
     */
    privateConnectionId: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The VPC Peering configuration is used to create VPC peering
     * between Datastream and the consumer's VPC.
     * Structure is documented below.
     */
    vpcPeeringConfig: pulumi.Input<inputs.datastream.PrivateConnectionVpcPeeringConfig>;
}
