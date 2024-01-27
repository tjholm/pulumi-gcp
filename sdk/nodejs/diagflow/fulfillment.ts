// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * By default, your agent responds to a matched intent with a static response. If you're using one of the integration options, you can provide a more dynamic response by using fulfillment. When you enable fulfillment for an intent, Dialogflow responds to that intent by calling a service that you define. For example, if an end-user wants to schedule a haircut on Friday, your service can check your database and respond to the end-user with availability information for Friday.
 *
 * To get more information about Fulfillment, see:
 *
 * * [API documentation](https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/projects.agent/getFulfillment)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/es/docs/fulfillment-overview)
 *
 * ## Example Usage
 * ### Dialogflow Fulfillment Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicAgent = new gcp.diagflow.Agent("basicAgent", {
 *     displayName: "example_agent",
 *     defaultLanguageCode: "en",
 *     timeZone: "America/New_York",
 * });
 * const basicFulfillment = new gcp.diagflow.Fulfillment("basicFulfillment", {
 *     displayName: "basic-fulfillment",
 *     enabled: true,
 *     genericWebService: {
 *         uri: "https://google.com",
 *         username: "admin",
 *         password: "password",
 *         requestHeaders: {
 *             name: "wrench",
 *         },
 *     },
 * }, {
 *     dependsOn: [basicAgent],
 * });
 * ```
 *
 * ## Import
 *
 * Fulfillment can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Fulfillment can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:diagflow/fulfillment:Fulfillment default {{name}}
 * ```
 */
export class Fulfillment extends pulumi.CustomResource {
    /**
     * Get an existing Fulfillment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FulfillmentState, opts?: pulumi.CustomResourceOptions): Fulfillment {
        return new Fulfillment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/fulfillment:Fulfillment';

    /**
     * Returns true if the given object is an instance of Fulfillment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fulfillment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fulfillment.__pulumiType;
    }

    /**
     * The human-readable name of the fulfillment, unique within the agent.
     *
     *
     * - - -
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether fulfillment is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     */
    public readonly features!: pulumi.Output<outputs.diagflow.FulfillmentFeature[] | undefined>;
    /**
     * Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     */
    public readonly genericWebService!: pulumi.Output<outputs.diagflow.FulfillmentGenericWebService | undefined>;
    /**
     * The unique identifier of the fulfillment.
     * Format: projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Fulfillment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FulfillmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FulfillmentArgs | FulfillmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FulfillmentState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["genericWebService"] = state ? state.genericWebService : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as FulfillmentArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["genericWebService"] = args ? args.genericWebService : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Fulfillment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fulfillment resources.
 */
export interface FulfillmentState {
    /**
     * The human-readable name of the fulfillment, unique within the agent.
     *
     *
     * - - -
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether fulfillment is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.diagflow.FulfillmentFeature>[]>;
    /**
     * Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     */
    genericWebService?: pulumi.Input<inputs.diagflow.FulfillmentGenericWebService>;
    /**
     * The unique identifier of the fulfillment.
     * Format: projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fulfillment resource.
 */
export interface FulfillmentArgs {
    /**
     * The human-readable name of the fulfillment, unique within the agent.
     *
     *
     * - - -
     */
    displayName: pulumi.Input<string>;
    /**
     * Whether fulfillment is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     */
    features?: pulumi.Input<pulumi.Input<inputs.diagflow.FulfillmentFeature>[]>;
    /**
     * Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     */
    genericWebService?: pulumi.Input<inputs.diagflow.FulfillmentGenericWebService>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
