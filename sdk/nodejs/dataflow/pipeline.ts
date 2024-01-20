// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The main pipeline entity and all the necessary metadata for launching and managing linked jobs.
 *
 * To get more information about Pipeline, see:
 *
 * * [API documentation](https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dataflow)
 *
 * ## Example Usage
 * ### Data Pipeline Pipeline
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serviceAccount = new gcp.serviceaccount.Account("serviceAccount", {
 *     accountId: "my-account",
 *     displayName: "Service Account",
 * });
 * const primary = new gcp.dataflow.Pipeline("primary", {
 *     displayName: "my-pipeline",
 *     type: "PIPELINE_TYPE_BATCH",
 *     state: "STATE_ACTIVE",
 *     region: "us-central1",
 *     workload: {
 *         dataflowLaunchTemplateRequest: {
 *             projectId: "my-project",
 *             gcsPath: "gs://my-bucket/path",
 *             launchParameters: {
 *                 jobName: "my-job",
 *                 parameters: {
 *                     name: "wrench",
 *                 },
 *                 environment: {
 *                     numWorkers: 5,
 *                     maxWorkers: 5,
 *                     zone: "us-centra1-a",
 *                     serviceAccountEmail: serviceAccount.email,
 *                     network: "default",
 *                     tempLocation: "gs://my-bucket/tmp_dir",
 *                     bypassTempDirValidation: false,
 *                     machineType: "E2",
 *                     additionalUserLabels: {
 *                         context: "test",
 *                     },
 *                     workerRegion: "us-central1",
 *                     workerZone: "us-central1-a",
 *                     enableStreamingEngine: false,
 *                 },
 *                 update: false,
 *                 transformNameMapping: {
 *                     name: "wrench",
 *                 },
 *             },
 *             location: "us-central1",
 *         },
 *     },
 *     scheduleInfo: {
 *         schedule: "* *&#47;2 * * *",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Pipeline can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/pipelines/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Pipeline using one of the formats above. For exampletf import {
 *
 *  id = "projects/{{project}}/locations/{{region}}/pipelines/{{name}}"
 *
 *  to = google_data_pipeline_pipeline.default }
 *
 * ```sh
 *  $ pulumi import gcp:dataflow/pipeline:Pipeline When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Pipeline can be imported using one of the formats above. For example
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataflow/pipeline:Pipeline default projects/{{project}}/locations/{{region}}/pipelines/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{name}}
 * ```
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataflow/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Number of jobs.
     */
    public /*out*/ readonly jobCount!: pulumi.Output<number>;
    /**
     * The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly lastUpdateTime!: pulumi.Output<string>;
    /**
     * "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
     * "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
     * "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
     * "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly pipelineSources!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A reference to the region
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
     * Structure is documented below.
     */
    public readonly scheduleInfo!: pulumi.Output<outputs.dataflow.PipelineScheduleInfo | undefined>;
    /**
     * Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
     */
    public readonly schedulerServiceAccountEmail!: pulumi.Output<string>;
    /**
     * The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
     * Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
     *
     *
     * - - -
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
     * Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Workload information for creating new jobs.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
     * Structure is documented below.
     */
    public readonly workload!: pulumi.Output<outputs.dataflow.PipelineWorkload | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["jobCount"] = state ? state.jobCount : undefined;
            resourceInputs["lastUpdateTime"] = state ? state.lastUpdateTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pipelineSources"] = state ? state.pipelineSources : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["scheduleInfo"] = state ? state.scheduleInfo : undefined;
            resourceInputs["schedulerServiceAccountEmail"] = state ? state.schedulerServiceAccountEmail : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["workload"] = state ? state.workload : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if ((!args || args.state === undefined) && !opts.urn) {
                throw new Error("Missing required property 'state'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pipelineSources"] = args ? args.pipelineSources : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["scheduleInfo"] = args ? args.scheduleInfo : undefined;
            resourceInputs["schedulerServiceAccountEmail"] = args ? args.schedulerServiceAccountEmail : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["workload"] = args ? args.workload : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["jobCount"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    createTime?: pulumi.Input<string>;
    /**
     * The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
     */
    displayName?: pulumi.Input<string>;
    /**
     * Number of jobs.
     */
    jobCount?: pulumi.Input<number>;
    /**
     * The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    lastUpdateTime?: pulumi.Input<string>;
    /**
     * "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
     * "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
     * "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
     * "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
     */
    name?: pulumi.Input<string>;
    /**
     * The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    pipelineSources?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A reference to the region
     */
    region?: pulumi.Input<string>;
    /**
     * Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
     * Structure is documented below.
     */
    scheduleInfo?: pulumi.Input<inputs.dataflow.PipelineScheduleInfo>;
    /**
     * Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
     */
    schedulerServiceAccountEmail?: pulumi.Input<string>;
    /**
     * The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
     * Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
     *
     *
     * - - -
     */
    state?: pulumi.Input<string>;
    /**
     * The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
     * Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
     */
    type?: pulumi.Input<string>;
    /**
     * Workload information for creating new jobs.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
     * Structure is documented below.
     */
    workload?: pulumi.Input<inputs.dataflow.PipelineWorkload>;
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
     */
    displayName?: pulumi.Input<string>;
    /**
     * "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
     * "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
     * "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
     * "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
     */
    name?: pulumi.Input<string>;
    /**
     * The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    pipelineSources?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A reference to the region
     */
    region?: pulumi.Input<string>;
    /**
     * Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
     * Structure is documented below.
     */
    scheduleInfo?: pulumi.Input<inputs.dataflow.PipelineScheduleInfo>;
    /**
     * Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
     */
    schedulerServiceAccountEmail?: pulumi.Input<string>;
    /**
     * The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
     * Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
     *
     *
     * - - -
     */
    state: pulumi.Input<string>;
    /**
     * The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
     * Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
     */
    type: pulumi.Input<string>;
    /**
     * Workload information for creating new jobs.
     * https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
     * Structure is documented below.
     */
    workload?: pulumi.Input<inputs.dataflow.PipelineWorkload>;
}
