// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a organization-level logging sink. For more information see:
 * * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/organizations.sinks)
 * * How-to Guides
 *     * [Exporting Logs](https://cloud.google.com/logging/docs/export)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const log_bucket = new gcp.storage.Bucket("log-bucket", {location: "US"});
 * const my_sink = new gcp.logging.OrganizationSink("my-sink", {
 *     description: "some explanation on what this is",
 *     orgId: "123456789",
 *     destination: pulumi.interpolate`storage.googleapis.com/${log_bucket.name}`,
 *     filter: "resource.type = gce_instance AND severity >= WARNING",
 * });
 * const log_writer = new gcp.projects.IAMMember("log-writer", {
 *     project: "your-project-id",
 *     role: "roles/storage.objectCreator",
 *     member: my_sink.writerIdentity,
 * });
 * ```
 *
 * ## Import
 *
 * Organization-level logging sinks can be imported using this format* `organizations/{{organization_id}}/sinks/{{sink_id}}` When using the `pulumi import` command, organization-level logging sinks can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:logging/organizationSink:OrganizationSink default organizations/{{organization_id}}/sinks/{{sink_id}}
 * ```
 */
export class OrganizationSink extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSinkState, opts?: pulumi.CustomResourceOptions): OrganizationSink {
        return new OrganizationSink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/organizationSink:OrganizationSink';

    /**
     * Returns true if the given object is an instance of OrganizationSink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSink.__pulumiType;
    }

    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    public readonly bigqueryOptions!: pulumi.Output<outputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * A description of this sink. The maximum length of the description is 8000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
     *
     * - `storage.googleapis.com/[GCS_BUCKET]`
     * - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
     * - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]/locations/global/buckets/[BUCKET_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]`
     *
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * If set to True, then this sink is disabled and it does not export any log entries.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    public readonly exclusions!: pulumi.Output<outputs.logging.OrganizationSinkExclusion[] | undefined>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    public readonly includeChildren!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the logging sink.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity!: pulumi.Output<string>;

    /**
     * Create a OrganizationSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSinkArgs | OrganizationSinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationSinkState | undefined;
            resourceInputs["bigqueryOptions"] = state ? state.bigqueryOptions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["exclusions"] = state ? state.exclusions : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["includeChildren"] = state ? state.includeChildren : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as OrganizationSinkArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["bigqueryOptions"] = args ? args.bigqueryOptions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["exclusions"] = args ? args.exclusions : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["includeChildren"] = args ? args.includeChildren : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["writerIdentity"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationSink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSink resources.
 */
export interface OrganizationSinkState {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    bigqueryOptions?: pulumi.Input<inputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * A description of this sink. The maximum length of the description is 8000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
     *
     * - `storage.googleapis.com/[GCS_BUCKET]`
     * - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
     * - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]/locations/global/buckets/[BUCKET_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]`
     *
     * The writer associated with the sink must have access to write to the above resource.
     */
    destination?: pulumi.Input<string>;
    /**
     * If set to True, then this sink is disabled and it does not export any log entries.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    exclusions?: pulumi.Input<pulumi.Input<inputs.logging.OrganizationSinkExclusion>[]>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationSink resource.
 */
export interface OrganizationSinkArgs {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    bigqueryOptions?: pulumi.Input<inputs.logging.OrganizationSinkBigqueryOptions>;
    /**
     * A description of this sink. The maximum length of the description is 8000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:
     *
     * - `storage.googleapis.com/[GCS_BUCKET]`
     * - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
     * - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]/locations/global/buckets/[BUCKET_ID]`
     * - `logging.googleapis.com/projects/[PROJECT_ID]`
     *
     * The writer associated with the sink must have access to write to the above resource.
     */
    destination: pulumi.Input<string>;
    /**
     * If set to True, then this sink is disabled and it does not export any log entries.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    exclusions?: pulumi.Input<pulumi.Input<inputs.logging.OrganizationSinkExclusion>[]>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Whether or not to include children organizations in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
     */
    includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    name?: pulumi.Input<string>;
    /**
     * The numeric ID of the organization to be exported to the sink.
     */
    orgId: pulumi.Input<string>;
}
