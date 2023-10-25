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
 * * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
 * * How-to Guides
 *     * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)
 *
 * > **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
 * by using the `gcp.projects.ServiceIdentity` resource.
 *
 * ## Example Usage
 * ### Pubsub Subscription Push
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleTopic = new gcp.pubsub.Topic("exampleTopic", {});
 * const exampleSubscription = new gcp.pubsub.Subscription("exampleSubscription", {
 *     topic: exampleTopic.name,
 *     ackDeadlineSeconds: 20,
 *     labels: {
 *         foo: "bar",
 *     },
 *     pushConfig: {
 *         pushEndpoint: "https://example.com/push",
 *         attributes: {
 *             "x-goog-version": "v1",
 *         },
 *     },
 * });
 * ```
 * ### Pubsub Subscription Pull
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleTopic = new gcp.pubsub.Topic("exampleTopic", {});
 * const exampleSubscription = new gcp.pubsub.Subscription("exampleSubscription", {
 *     topic: exampleTopic.name,
 *     labels: {
 *         foo: "bar",
 *     },
 *     messageRetentionDuration: "1200s",
 *     retainAckedMessages: true,
 *     ackDeadlineSeconds: 20,
 *     expirationPolicy: {
 *         ttl: "300000.5s",
 *     },
 *     retryPolicy: {
 *         minimumBackoff: "10s",
 *     },
 *     enableMessageOrdering: false,
 * });
 * ```
 * ### Pubsub Subscription Different Project
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleTopic = new gcp.pubsub.Topic("exampleTopic", {project: "topic-project"});
 * const exampleSubscription = new gcp.pubsub.Subscription("exampleSubscription", {
 *     project: "subscription-project",
 *     topic: exampleTopic.name,
 * });
 * ```
 * ### Pubsub Subscription Dead Letter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleTopic = new gcp.pubsub.Topic("exampleTopic", {});
 * const exampleDeadLetter = new gcp.pubsub.Topic("exampleDeadLetter", {});
 * const exampleSubscription = new gcp.pubsub.Subscription("exampleSubscription", {
 *     topic: exampleTopic.name,
 *     deadLetterPolicy: {
 *         deadLetterTopic: exampleDeadLetter.id,
 *         maxDeliveryAttempts: 10,
 *     },
 * });
 * ```
 * ### Pubsub Subscription Push Bq
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const exampleTopic = new gcp.pubsub.Topic("exampleTopic", {});
 * const project = gcp.organizations.getProject({});
 * const viewer = new gcp.projects.IAMMember("viewer", {
 *     project: project.then(project => project.projectId),
 *     role: "roles/bigquery.metadataViewer",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-pubsub.iam.gserviceaccount.com`),
 * });
 * const editor = new gcp.projects.IAMMember("editor", {
 *     project: project.then(project => project.projectId),
 *     role: "roles/bigquery.dataEditor",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-pubsub.iam.gserviceaccount.com`),
 * });
 * const testDataset = new gcp.bigquery.Dataset("testDataset", {datasetId: "example_dataset"});
 * const testTable = new gcp.bigquery.Table("testTable", {
 *     deletionProtection: false,
 *     tableId: "example_table",
 *     datasetId: testDataset.datasetId,
 *     schema: `[
 *   {
 *     "name": "data",
 *     "type": "STRING",
 *     "mode": "NULLABLE",
 *     "description": "The data"
 *   }
 * ]
 * `,
 * });
 * const exampleSubscription = new gcp.pubsub.Subscription("exampleSubscription", {
 *     topic: exampleTopic.name,
 *     bigqueryConfig: {
 *         table: pulumi.interpolate`${testTable.project}.${testTable.datasetId}.${testTable.tableId}`,
 *     },
 * }, {
 *     dependsOn: [
 *         viewer,
 *         editor,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Subscription can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/subscription:Subscription default projects/{{project}}/subscriptions/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/subscription:Subscription default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:pubsub/subscription:Subscription default {{name}}
 * ```
 */
export class Subscription extends pulumi.CustomResource {
    /**
     * Get an existing Subscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionState, opts?: pulumi.CustomResourceOptions): Subscription {
        return new Subscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/subscription:Subscription';

    /**
     * Returns true if the given object is an instance of Subscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subscription.__pulumiType;
    }

    /**
     * This value is the maximum time after a subscriber receives a message
     * before the subscriber should acknowledge the message. After message
     * delivery but before the ack deadline expires and before the message is
     * acknowledged, it is an outstanding message and will not be delivered
     * again during that time (on a best-effort basis).
     * For pull subscriptions, this value is used as the initial value for
     * the ack deadline. To override this value for a given message, call
     * subscriptions.modifyAckDeadline with the corresponding ackId if using
     * pull. The minimum custom deadline you can specify is 10 seconds. The
     * maximum custom deadline you can specify is 600 seconds (10 minutes).
     * If this parameter is 0, a default value of 10 seconds is used.
     * For push delivery, this value is also used to set the request timeout
     * for the call to the push endpoint.
     * If the subscriber never acknowledges the message, the Pub/Sub system
     * will eventually redeliver the message.
     */
    public readonly ackDeadlineSeconds!: pulumi.Output<number>;
    /**
     * If delivery to BigQuery is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    public readonly bigqueryConfig!: pulumi.Output<outputs.pubsub.SubscriptionBigqueryConfig | undefined>;
    /**
     * If delivery to Cloud Storage is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    public readonly cloudStorageConfig!: pulumi.Output<outputs.pubsub.SubscriptionCloudStorageConfig | undefined>;
    /**
     * A policy that specifies the conditions for dead lettering messages in
     * this subscription. If deadLetterPolicy is not set, dead lettering
     * is disabled.
     * The Cloud Pub/Sub service account associated with this subscription's
     * parent project (i.e.,
     * service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
     * permission to Acknowledge() messages on this subscription.
     * Structure is documented below.
     */
    public readonly deadLetterPolicy!: pulumi.Output<outputs.pubsub.SubscriptionDeadLetterPolicy | undefined>;
    /**
     * If `true`, Pub/Sub provides the following guarantees for the delivery
     * of a message with a given value of messageId on this Subscriptions':
     * - The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
     * - An acknowledged message will not be resent to a subscriber.
     * Note that subscribers may still receive multiple copies of a message when `enableExactlyOnceDelivery`
     * is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values
     */
    public readonly enableExactlyOnceDelivery!: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
     * the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
     * may be delivered in any order.
     */
    public readonly enableMessageOrdering!: pulumi.Output<boolean | undefined>;
    /**
     * A policy that specifies the conditions for this subscription's expiration.
     * A subscription is considered active as long as any connected subscriber
     * is successfully consuming messages from the subscription or is issuing
     * operations on the subscription. If expirationPolicy is not set, a default
     * policy with ttl of 31 days will be used.  If it is set but ttl is "", the
     * resource never expires.  The minimum allowed value for expirationPolicy.ttl
     * is 1 day.
     * Structure is documented below.
     */
    public readonly expirationPolicy!: pulumi.Output<outputs.pubsub.SubscriptionExpirationPolicy>;
    /**
     * The subscription only delivers the messages that match the filter.
     * Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
     * by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
     * you can't modify the filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to this Subscription.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * How long to retain unacknowledged messages in the subscription's
     * backlog, from the moment a message is published. If
     * retainAckedMessages is true, then this also configures the retention
     * of acknowledged messages, and thus configures how far back in time a
     * subscriptions.seek can be done. Defaults to 7 days. Cannot be more
     * than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
     * A duration in seconds with up to nine fractional digits, terminated
     * by 's'. Example: `"600.5s"`.
     */
    public readonly messageRetentionDuration!: pulumi.Output<string | undefined>;
    /**
     * Name of the subscription.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * If push delivery is used with this subscription, this field is used to
     * configure it. An empty pushConfig signifies that the subscriber will
     * pull and ack messages using API methods.
     * Structure is documented below.
     */
    public readonly pushConfig!: pulumi.Output<outputs.pubsub.SubscriptionPushConfig | undefined>;
    /**
     * Indicates whether to retain acknowledged messages. If `true`, then
     * messages are not expunged from the subscription's backlog, even if
     * they are acknowledged, until they fall out of the
     * messageRetentionDuration window.
     */
    public readonly retainAckedMessages!: pulumi.Output<boolean | undefined>;
    /**
     * A policy that specifies how Pub/Sub retries message delivery for this subscription.
     * If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
     * RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
     * Structure is documented below.
     */
    public readonly retryPolicy!: pulumi.Output<outputs.pubsub.SubscriptionRetryPolicy | undefined>;
    /**
     * A reference to a Topic resource.
     *
     *
     * - - -
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a Subscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionArgs | SubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionState | undefined;
            resourceInputs["ackDeadlineSeconds"] = state ? state.ackDeadlineSeconds : undefined;
            resourceInputs["bigqueryConfig"] = state ? state.bigqueryConfig : undefined;
            resourceInputs["cloudStorageConfig"] = state ? state.cloudStorageConfig : undefined;
            resourceInputs["deadLetterPolicy"] = state ? state.deadLetterPolicy : undefined;
            resourceInputs["enableExactlyOnceDelivery"] = state ? state.enableExactlyOnceDelivery : undefined;
            resourceInputs["enableMessageOrdering"] = state ? state.enableMessageOrdering : undefined;
            resourceInputs["expirationPolicy"] = state ? state.expirationPolicy : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["messageRetentionDuration"] = state ? state.messageRetentionDuration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pushConfig"] = state ? state.pushConfig : undefined;
            resourceInputs["retainAckedMessages"] = state ? state.retainAckedMessages : undefined;
            resourceInputs["retryPolicy"] = state ? state.retryPolicy : undefined;
            resourceInputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as SubscriptionArgs | undefined;
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            resourceInputs["ackDeadlineSeconds"] = args ? args.ackDeadlineSeconds : undefined;
            resourceInputs["bigqueryConfig"] = args ? args.bigqueryConfig : undefined;
            resourceInputs["cloudStorageConfig"] = args ? args.cloudStorageConfig : undefined;
            resourceInputs["deadLetterPolicy"] = args ? args.deadLetterPolicy : undefined;
            resourceInputs["enableExactlyOnceDelivery"] = args ? args.enableExactlyOnceDelivery : undefined;
            resourceInputs["enableMessageOrdering"] = args ? args.enableMessageOrdering : undefined;
            resourceInputs["expirationPolicy"] = args ? args.expirationPolicy : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["messageRetentionDuration"] = args ? args.messageRetentionDuration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pushConfig"] = args ? args.pushConfig : undefined;
            resourceInputs["retainAckedMessages"] = args ? args.retainAckedMessages : undefined;
            resourceInputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            resourceInputs["topic"] = args ? args.topic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subscription resources.
 */
export interface SubscriptionState {
    /**
     * This value is the maximum time after a subscriber receives a message
     * before the subscriber should acknowledge the message. After message
     * delivery but before the ack deadline expires and before the message is
     * acknowledged, it is an outstanding message and will not be delivered
     * again during that time (on a best-effort basis).
     * For pull subscriptions, this value is used as the initial value for
     * the ack deadline. To override this value for a given message, call
     * subscriptions.modifyAckDeadline with the corresponding ackId if using
     * pull. The minimum custom deadline you can specify is 10 seconds. The
     * maximum custom deadline you can specify is 600 seconds (10 minutes).
     * If this parameter is 0, a default value of 10 seconds is used.
     * For push delivery, this value is also used to set the request timeout
     * for the call to the push endpoint.
     * If the subscriber never acknowledges the message, the Pub/Sub system
     * will eventually redeliver the message.
     */
    ackDeadlineSeconds?: pulumi.Input<number>;
    /**
     * If delivery to BigQuery is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    bigqueryConfig?: pulumi.Input<inputs.pubsub.SubscriptionBigqueryConfig>;
    /**
     * If delivery to Cloud Storage is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    cloudStorageConfig?: pulumi.Input<inputs.pubsub.SubscriptionCloudStorageConfig>;
    /**
     * A policy that specifies the conditions for dead lettering messages in
     * this subscription. If deadLetterPolicy is not set, dead lettering
     * is disabled.
     * The Cloud Pub/Sub service account associated with this subscription's
     * parent project (i.e.,
     * service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
     * permission to Acknowledge() messages on this subscription.
     * Structure is documented below.
     */
    deadLetterPolicy?: pulumi.Input<inputs.pubsub.SubscriptionDeadLetterPolicy>;
    /**
     * If `true`, Pub/Sub provides the following guarantees for the delivery
     * of a message with a given value of messageId on this Subscriptions':
     * - The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
     * - An acknowledged message will not be resent to a subscriber.
     * Note that subscribers may still receive multiple copies of a message when `enableExactlyOnceDelivery`
     * is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values
     */
    enableExactlyOnceDelivery?: pulumi.Input<boolean>;
    /**
     * If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
     * the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
     * may be delivered in any order.
     */
    enableMessageOrdering?: pulumi.Input<boolean>;
    /**
     * A policy that specifies the conditions for this subscription's expiration.
     * A subscription is considered active as long as any connected subscriber
     * is successfully consuming messages from the subscription or is issuing
     * operations on the subscription. If expirationPolicy is not set, a default
     * policy with ttl of 31 days will be used.  If it is set but ttl is "", the
     * resource never expires.  The minimum allowed value for expirationPolicy.ttl
     * is 1 day.
     * Structure is documented below.
     */
    expirationPolicy?: pulumi.Input<inputs.pubsub.SubscriptionExpirationPolicy>;
    /**
     * The subscription only delivers the messages that match the filter.
     * Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
     * by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
     * you can't modify the filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to this Subscription.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * How long to retain unacknowledged messages in the subscription's
     * backlog, from the moment a message is published. If
     * retainAckedMessages is true, then this also configures the retention
     * of acknowledged messages, and thus configures how far back in time a
     * subscriptions.seek can be done. Defaults to 7 days. Cannot be more
     * than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
     * A duration in seconds with up to nine fractional digits, terminated
     * by 's'. Example: `"600.5s"`.
     */
    messageRetentionDuration?: pulumi.Input<string>;
    /**
     * Name of the subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * If push delivery is used with this subscription, this field is used to
     * configure it. An empty pushConfig signifies that the subscriber will
     * pull and ack messages using API methods.
     * Structure is documented below.
     */
    pushConfig?: pulumi.Input<inputs.pubsub.SubscriptionPushConfig>;
    /**
     * Indicates whether to retain acknowledged messages. If `true`, then
     * messages are not expunged from the subscription's backlog, even if
     * they are acknowledged, until they fall out of the
     * messageRetentionDuration window.
     */
    retainAckedMessages?: pulumi.Input<boolean>;
    /**
     * A policy that specifies how Pub/Sub retries message delivery for this subscription.
     * If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
     * RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
     * Structure is documented below.
     */
    retryPolicy?: pulumi.Input<inputs.pubsub.SubscriptionRetryPolicy>;
    /**
     * A reference to a Topic resource.
     *
     *
     * - - -
     */
    topic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subscription resource.
 */
export interface SubscriptionArgs {
    /**
     * This value is the maximum time after a subscriber receives a message
     * before the subscriber should acknowledge the message. After message
     * delivery but before the ack deadline expires and before the message is
     * acknowledged, it is an outstanding message and will not be delivered
     * again during that time (on a best-effort basis).
     * For pull subscriptions, this value is used as the initial value for
     * the ack deadline. To override this value for a given message, call
     * subscriptions.modifyAckDeadline with the corresponding ackId if using
     * pull. The minimum custom deadline you can specify is 10 seconds. The
     * maximum custom deadline you can specify is 600 seconds (10 minutes).
     * If this parameter is 0, a default value of 10 seconds is used.
     * For push delivery, this value is also used to set the request timeout
     * for the call to the push endpoint.
     * If the subscriber never acknowledges the message, the Pub/Sub system
     * will eventually redeliver the message.
     */
    ackDeadlineSeconds?: pulumi.Input<number>;
    /**
     * If delivery to BigQuery is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    bigqueryConfig?: pulumi.Input<inputs.pubsub.SubscriptionBigqueryConfig>;
    /**
     * If delivery to Cloud Storage is used with this subscription, this field is used to configure it.
     * Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
     * If all three are empty, then the subscriber will pull and ack messages using API methods.
     * Structure is documented below.
     */
    cloudStorageConfig?: pulumi.Input<inputs.pubsub.SubscriptionCloudStorageConfig>;
    /**
     * A policy that specifies the conditions for dead lettering messages in
     * this subscription. If deadLetterPolicy is not set, dead lettering
     * is disabled.
     * The Cloud Pub/Sub service account associated with this subscription's
     * parent project (i.e.,
     * service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
     * permission to Acknowledge() messages on this subscription.
     * Structure is documented below.
     */
    deadLetterPolicy?: pulumi.Input<inputs.pubsub.SubscriptionDeadLetterPolicy>;
    /**
     * If `true`, Pub/Sub provides the following guarantees for the delivery
     * of a message with a given value of messageId on this Subscriptions':
     * - The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
     * - An acknowledged message will not be resent to a subscriber.
     * Note that subscribers may still receive multiple copies of a message when `enableExactlyOnceDelivery`
     * is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values
     */
    enableExactlyOnceDelivery?: pulumi.Input<boolean>;
    /**
     * If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
     * the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
     * may be delivered in any order.
     */
    enableMessageOrdering?: pulumi.Input<boolean>;
    /**
     * A policy that specifies the conditions for this subscription's expiration.
     * A subscription is considered active as long as any connected subscriber
     * is successfully consuming messages from the subscription or is issuing
     * operations on the subscription. If expirationPolicy is not set, a default
     * policy with ttl of 31 days will be used.  If it is set but ttl is "", the
     * resource never expires.  The minimum allowed value for expirationPolicy.ttl
     * is 1 day.
     * Structure is documented below.
     */
    expirationPolicy?: pulumi.Input<inputs.pubsub.SubscriptionExpirationPolicy>;
    /**
     * The subscription only delivers the messages that match the filter.
     * Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
     * by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
     * you can't modify the filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to this Subscription.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * How long to retain unacknowledged messages in the subscription's
     * backlog, from the moment a message is published. If
     * retainAckedMessages is true, then this also configures the retention
     * of acknowledged messages, and thus configures how far back in time a
     * subscriptions.seek can be done. Defaults to 7 days. Cannot be more
     * than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
     * A duration in seconds with up to nine fractional digits, terminated
     * by 's'. Example: `"600.5s"`.
     */
    messageRetentionDuration?: pulumi.Input<string>;
    /**
     * Name of the subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * If push delivery is used with this subscription, this field is used to
     * configure it. An empty pushConfig signifies that the subscriber will
     * pull and ack messages using API methods.
     * Structure is documented below.
     */
    pushConfig?: pulumi.Input<inputs.pubsub.SubscriptionPushConfig>;
    /**
     * Indicates whether to retain acknowledged messages. If `true`, then
     * messages are not expunged from the subscription's backlog, even if
     * they are acknowledged, until they fall out of the
     * messageRetentionDuration window.
     */
    retainAckedMessages?: pulumi.Input<boolean>;
    /**
     * A policy that specifies how Pub/Sub retries message delivery for this subscription.
     * If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
     * RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
     * Structure is documented below.
     */
    retryPolicy?: pulumi.Input<inputs.pubsub.SubscriptionRetryPolicy>;
    /**
     * A reference to a Topic resource.
     *
     *
     * - - -
     */
    topic: pulumi.Input<string>;
}
