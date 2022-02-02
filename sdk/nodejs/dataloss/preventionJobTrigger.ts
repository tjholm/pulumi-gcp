// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A job trigger configuration.
 *
 * To get more information about JobTrigger, see:
 *
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.jobTriggers)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-job-triggers)
 *
 * ## Example Usage
 * ### Dlp Job Trigger Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.dataloss.PreventionJobTrigger("basic", {
 *     description: "Description",
 *     displayName: "Displayname",
 *     inspectJob: {
 *         actions: [{
 *             saveFindings: {
 *                 outputConfig: {
 *                     table: {
 *                         datasetId: "asdf",
 *                         projectId: "asdf",
 *                     },
 *                 },
 *             },
 *         }],
 *         inspectTemplateName: "fake",
 *         storageConfig: {
 *             cloudStorageOptions: {
 *                 fileSet: {
 *                     url: "gs://mybucket/directory/",
 *                 },
 *             },
 *         },
 *     },
 *     parent: "projects/my-project-name",
 *     triggers: [{
 *         schedule: {
 *             recurrencePeriodDuration: "86400s",
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * JobTrigger can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/jobTriggers/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/{{name}}
 * ```
 */
export class PreventionJobTrigger extends pulumi.CustomResource {
    /**
     * Get an existing PreventionJobTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PreventionJobTriggerState, opts?: pulumi.CustomResourceOptions): PreventionJobTrigger {
        return new PreventionJobTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataloss/preventionJobTrigger:PreventionJobTrigger';

    /**
     * Returns true if the given object is an instance of PreventionJobTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreventionJobTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreventionJobTrigger.__pulumiType;
    }

    /**
     * A description of the job trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User set display name of the job trigger.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Controls what and how to inspect for findings.
     * Structure is documented below.
     */
    public readonly inspectJob!: pulumi.Output<outputs.dataloss.PreventionJobTriggerInspectJob | undefined>;
    /**
     * The timestamp of the last time this trigger executed.
     */
    public /*out*/ readonly lastRunTime!: pulumi.Output<string>;
    /**
     * The name of the Datastore kind.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the trigger, either in the format `projects/{{project}}`
     * or `projects/{{project}}/locations/{{location}}`
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Whether the trigger is currently active.
     * Default value is `HEALTHY`.
     * Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * What event needs to occur for a new job to be started.
     * Structure is documented below.
     */
    public readonly triggers!: pulumi.Output<outputs.dataloss.PreventionJobTriggerTrigger[]>;

    /**
     * Create a PreventionJobTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreventionJobTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PreventionJobTriggerArgs | PreventionJobTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PreventionJobTriggerState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["inspectJob"] = state ? state.inspectJob : undefined;
            resourceInputs["lastRunTime"] = state ? state.lastRunTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as PreventionJobTriggerArgs | undefined;
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            if ((!args || args.triggers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggers'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["inspectJob"] = args ? args.inspectJob : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["lastRunTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PreventionJobTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PreventionJobTrigger resources.
 */
export interface PreventionJobTriggerState {
    /**
     * A description of the job trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * User set display name of the job trigger.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Controls what and how to inspect for findings.
     * Structure is documented below.
     */
    inspectJob?: pulumi.Input<inputs.dataloss.PreventionJobTriggerInspectJob>;
    /**
     * The timestamp of the last time this trigger executed.
     */
    lastRunTime?: pulumi.Input<string>;
    /**
     * The name of the Datastore kind.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the trigger, either in the format `projects/{{project}}`
     * or `projects/{{project}}/locations/{{location}}`
     */
    parent?: pulumi.Input<string>;
    /**
     * Whether the trigger is currently active.
     * Default value is `HEALTHY`.
     * Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
     */
    status?: pulumi.Input<string>;
    /**
     * What event needs to occur for a new job to be started.
     * Structure is documented below.
     */
    triggers?: pulumi.Input<pulumi.Input<inputs.dataloss.PreventionJobTriggerTrigger>[]>;
}

/**
 * The set of arguments for constructing a PreventionJobTrigger resource.
 */
export interface PreventionJobTriggerArgs {
    /**
     * A description of the job trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * User set display name of the job trigger.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Controls what and how to inspect for findings.
     * Structure is documented below.
     */
    inspectJob?: pulumi.Input<inputs.dataloss.PreventionJobTriggerInspectJob>;
    /**
     * The parent of the trigger, either in the format `projects/{{project}}`
     * or `projects/{{project}}/locations/{{location}}`
     */
    parent: pulumi.Input<string>;
    /**
     * Whether the trigger is currently active.
     * Default value is `HEALTHY`.
     * Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
     */
    status?: pulumi.Input<string>;
    /**
     * What event needs to occur for a new job to be started.
     * Structure is documented below.
     */
    triggers: pulumi.Input<pulumi.Input<inputs.dataloss.PreventionJobTriggerTrigger>[]>;
}
