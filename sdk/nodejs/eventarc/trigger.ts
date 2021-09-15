// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
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
 *                 args: ["arrgs"],
 *             }],
 *             containerConcurrency: 50,
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
     * Output only. The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Required. Destination specifies where the events should be sent to.
     */
    public readonly destination!: pulumi.Output<outputs.eventarc.TriggerDestination>;
    /**
     * Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create
     * requests to ensure the client has an up-to-date value before proceeding.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
     */
    public readonly matchingCriterias!: pulumi.Output<outputs.eventarc.TriggerMatchingCriteria[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    public readonly serviceAccount!: pulumi.Output<string | undefined>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    public readonly transports!: pulumi.Output<outputs.eventarc.TriggerTransport[]>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["matchingCriterias"] = state ? state.matchingCriterias : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["transports"] = state ? state.transports : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
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
            inputs["destination"] = args ? args.destination : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["matchingCriterias"] = args ? args.matchingCriterias : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["transports"] = args ? args.transports : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * Output only. The creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Required. Destination specifies where the events should be sent to.
     */
    destination?: pulumi.Input<inputs.eventarc.TriggerDestination>;
    /**
     * Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create
     * requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
     */
    matchingCriterias?: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerMatchingCriteria>[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    transports?: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerTransport>[]>;
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
     * Required. Destination specifies where the events should be sent to.
     */
    destination: pulumi.Input<inputs.eventarc.TriggerDestination>;
    /**
     * Optional. User labels attached to the triggers that can be used to group resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
     */
    matchingCriterias: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerMatchingCriteria>[]>;
    /**
     * Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
     */
    name?: pulumi.Input<string>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
     */
    transports?: pulumi.Input<pulumi.Input<inputs.eventarc.TriggerTransport>[]>;
}
