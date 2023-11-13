// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Eventarc Trigger resource
 *
 * ## Example Usage
 * ### Basic
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.cloudrun.Service("default", {
 *     location: "europe-west1",
 *     metadata: {
 *         namespace: "my-project-name",
 *     },
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "gcr.io/cloudrun/hello",
 *                 ports: [{
 *                     containerPort: 8080,
 *                 }],
 *             }],
 *             containerConcurrency: 50,
 *             timeoutSeconds: 100,
 *         },
 *     },
 *     traffics: [{
 *         percent: 100,
 *         latestRevision: true,
 *     }],
 * });
 * const primary = new gcp.eventarc.Trigger("primary", {
 *     location: "europe-west1",
 *     matchingCriterias: [{
 *         attribute: "type",
 *         value: "google.cloud.pubsub.topic.v1.messagePublished",
 *     }],
 *     destination: {
 *         cloudRunService: {
 *             service: _default.name,
 *             region: "europe-west1",
 *         },
 *     },
 *     labels: {
 *         foo: "bar",
 *     },
 * });
 * const foo = new gcp.pubsub.Topic("foo", {});
 * ```
 *
 * ## Import
 *
 * Trigger can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/trigger:Trigger default projects/{{project}}/locations/{{location}}/triggers/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/trigger:Trigger default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:eventarc/trigger:Trigger default {{location}}/{{name}}
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:eventarc/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
     */
    public readonly channel!: pulumi.Output<string | undefined>;
    /**
     * Output only. The reason(s) why a trigger is in FAILED state.
     */
    public /*out*/ readonly conditions!: pulumi.Output<{[key: string]: string}>;
    /**
     * Output only. The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Required. Destination specifies where the events should be sent to.
     */
    public readonly destination!: pulumi.Output<outputs.eventarc.TriggerDestination>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
     */
    public readonly eventDataContentType!: pulumi.Output<string>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
     */
    public readonly matchingCriterias!: pulumi.Output<outputs.eventarc.TriggerMatchingCriteria[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: any}>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    public readonly serviceAccount!: pulumi.Output<string | undefined>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    public readonly transport!: pulumi.Output<outputs.eventarc.TriggerTransport>;
    /**
     * Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Output only. The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            resourceInputs["channel"] = state ? state.channel : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["eventDataContentType"] = state ? state.eventDataContentType : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["matchingCriterias"] = state ? state.matchingCriterias : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            resourceInputs["transport"] = state ? state.transport : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.matchingCriterias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'matchingCriterias'");
            }
            resourceInputs["channel"] = args ? args.channel : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["eventDataContentType"] = args ? args.eventDataContentType : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["matchingCriterias"] = args ? args.matchingCriterias : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["transport"] = args ? args.transport : undefined;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Trigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
     */
    channel?: pulumi.Input<string>;
    /**
     * Output only. The reason(s) why a trigger is in FAILED state.
     */
    conditions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Output only. The creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Required. Destination specifies where the events should be sent to.
     */
    destination?: pulumi.Input<inputs.eventarc.TriggerDestination>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
     */
    eventDataContentType?: pulumi.Input<string>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
     */
    matchingCriterias?: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerMatchingCriteria>[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    transport?: pulumi.Input<inputs.eventarc.TriggerTransport>;
    /**
     * Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    uid?: pulumi.Input<string>;
    /**
     * Output only. The last-modified time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
     */
    channel?: pulumi.Input<string>;
    /**
     * Required. Destination specifies where the events should be sent to.
     */
    destination: pulumi.Input<inputs.eventarc.TriggerDestination>;
    /**
     * Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
     */
    eventDataContentType?: pulumi.Input<string>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
     */
    matchingCriterias: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerMatchingCriteria>[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    transport?: pulumi.Input<inputs.eventarc.TriggerTransport>;
}
