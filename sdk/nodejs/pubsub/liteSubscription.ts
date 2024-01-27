// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A named resource representing the stream of messages from a single,
 * specific topic, to be delivered to the subscribing application.
 *
 * To get more information about Subscription, see:
 *
 * * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.subscriptions)
 * * How-to Guides
 *     * [Managing Subscriptions](https://cloud.google.com/pubsub/lite/docs/subscriptions)
 *
 * ## Example Usage
 * ### Pubsub Lite Subscription Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const exampleLiteTopic = new gcp.pubsub.LiteTopic("exampleLiteTopic", {
 *     project: project.then(project => project.number),
 *     partitionConfig: {
 *         count: 1,
 *         capacity: {
 *             publishMibPerSec: 4,
 *             subscribeMibPerSec: 8,
 *         },
 *     },
 *     retentionConfig: {
 *         perPartitionBytes: "32212254720",
 *     },
 * });
 * const exampleLiteSubscription = new gcp.pubsub.LiteSubscription("exampleLiteSubscription", {
 *     topic: exampleLiteTopic.name,
 *     deliveryConfig: {
 *         deliveryRequirement: "DELIVER_AFTER_STORED",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Subscription can be imported using any of these accepted formats* `projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}` * `{{project}}/{{zone}}/{{name}}` * `{{zone}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Subscription can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{project}}/{{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{name}}
 * ```
 */
export class LiteSubscription extends pulumi.CustomResource {
    /**
     * Get an existing LiteSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LiteSubscriptionState, opts?: pulumi.CustomResourceOptions): LiteSubscription {
        return new LiteSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/liteSubscription:LiteSubscription';

    /**
     * Returns true if the given object is an instance of LiteSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiteSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiteSubscription.__pulumiType;
    }

    /**
     * The settings for this subscription's message delivery.
     * Structure is documented below.
     */
    public readonly deliveryConfig!: pulumi.Output<outputs.pubsub.LiteSubscriptionDeliveryConfig | undefined>;
    /**
     * Name of the subscription.
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
     * The region of the pubsub lite topic.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * A reference to a Topic resource.
     */
    public readonly topic!: pulumi.Output<string>;
    /**
     * The zone of the pubsub lite topic.
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a LiteSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiteSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LiteSubscriptionArgs | LiteSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LiteSubscriptionState | undefined;
            resourceInputs["deliveryConfig"] = state ? state.deliveryConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as LiteSubscriptionArgs | undefined;
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            resourceInputs["deliveryConfig"] = args ? args.deliveryConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LiteSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LiteSubscription resources.
 */
export interface LiteSubscriptionState {
    /**
     * The settings for this subscription's message delivery.
     * Structure is documented below.
     */
    deliveryConfig?: pulumi.Input<inputs.pubsub.LiteSubscriptionDeliveryConfig>;
    /**
     * Name of the subscription.
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
     * The region of the pubsub lite topic.
     */
    region?: pulumi.Input<string>;
    /**
     * A reference to a Topic resource.
     */
    topic?: pulumi.Input<string>;
    /**
     * The zone of the pubsub lite topic.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LiteSubscription resource.
 */
export interface LiteSubscriptionArgs {
    /**
     * The settings for this subscription's message delivery.
     * Structure is documented below.
     */
    deliveryConfig?: pulumi.Input<inputs.pubsub.LiteSubscriptionDeliveryConfig>;
    /**
     * Name of the subscription.
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
     * The region of the pubsub lite topic.
     */
    region?: pulumi.Input<string>;
    /**
     * A reference to a Topic resource.
     */
    topic: pulumi.Input<string>;
    /**
     * The zone of the pubsub lite topic.
     */
    zone?: pulumi.Input<string>;
}
