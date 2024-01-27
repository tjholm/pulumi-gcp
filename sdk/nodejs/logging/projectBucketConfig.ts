// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a project-level logging bucket config. For more information see
 * [the official logging documentation](https://cloud.google.com/logging/docs/) and
 * [Storing Logs](https://cloud.google.com/logging/docs/storage).
 *
 * > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
 *
 * ## Import
 *
 * This resource can be imported using the following format* `projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}` When using the `pulumi import` command, this resource can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
 * ```
 */
export class ProjectBucketConfig extends pulumi.CustomResource {
    /**
     * Get an existing ProjectBucketConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectBucketConfigState, opts?: pulumi.CustomResourceOptions): ProjectBucketConfig {
        return new ProjectBucketConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/projectBucketConfig:ProjectBucketConfig';

    /**
     * Returns true if the given object is an instance of ProjectBucketConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectBucketConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectBucketConfig.__pulumiType;
    }

    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    public readonly bucketId!: pulumi.Output<string>;
    /**
     * The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
     */
    public readonly cmekSettings!: pulumi.Output<outputs.logging.ProjectBucketConfigCmekSettings | undefined>;
    /**
     * Describes this bucket.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
     */
    public readonly enableAnalytics!: pulumi.Output<boolean | undefined>;
    /**
     * A list of indexed fields and related configuration data. Structure is documented below.
     */
    public readonly indexConfigs!: pulumi.Output<outputs.logging.ProjectBucketConfigIndexConfig[] | undefined>;
    /**
     * The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     */
    public /*out*/ readonly lifecycleState!: pulumi.Output<string>;
    /**
     * The location of the bucket.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
     */
    public readonly locked!: pulumi.Output<boolean | undefined>;
    /**
     * The resource name of the CMEK settings.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;

    /**
     * Create a ProjectBucketConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectBucketConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectBucketConfigArgs | ProjectBucketConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectBucketConfigState | undefined;
            resourceInputs["bucketId"] = state ? state.bucketId : undefined;
            resourceInputs["cmekSettings"] = state ? state.cmekSettings : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableAnalytics"] = state ? state.enableAnalytics : undefined;
            resourceInputs["indexConfigs"] = state ? state.indexConfigs : undefined;
            resourceInputs["lifecycleState"] = state ? state.lifecycleState : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["locked"] = state ? state.locked : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["retentionDays"] = state ? state.retentionDays : undefined;
        } else {
            const args = argsOrState as ProjectBucketConfigArgs | undefined;
            if ((!args || args.bucketId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["bucketId"] = args ? args.bucketId : undefined;
            resourceInputs["cmekSettings"] = args ? args.cmekSettings : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableAnalytics"] = args ? args.enableAnalytics : undefined;
            resourceInputs["indexConfigs"] = args ? args.indexConfigs : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["lifecycleState"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectBucketConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectBucketConfig resources.
 */
export interface ProjectBucketConfigState {
    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    bucketId?: pulumi.Input<string>;
    /**
     * The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
     */
    cmekSettings?: pulumi.Input<inputs.logging.ProjectBucketConfigCmekSettings>;
    /**
     * Describes this bucket.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
     */
    enableAnalytics?: pulumi.Input<boolean>;
    /**
     * A list of indexed fields and related configuration data. Structure is documented below.
     */
    indexConfigs?: pulumi.Input<pulumi.Input<inputs.logging.ProjectBucketConfigIndexConfig>[]>;
    /**
     * The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     */
    lifecycleState?: pulumi.Input<string>;
    /**
     * The location of the bucket.
     */
    location?: pulumi.Input<string>;
    /**
     * Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * The resource name of the CMEK settings.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent resource that contains the logging bucket.
     */
    project?: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    retentionDays?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProjectBucketConfig resource.
 */
export interface ProjectBucketConfigArgs {
    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     */
    bucketId: pulumi.Input<string>;
    /**
     * The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
     */
    cmekSettings?: pulumi.Input<inputs.logging.ProjectBucketConfigCmekSettings>;
    /**
     * Describes this bucket.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
     */
    enableAnalytics?: pulumi.Input<boolean>;
    /**
     * A list of indexed fields and related configuration data. Structure is documented below.
     */
    indexConfigs?: pulumi.Input<pulumi.Input<inputs.logging.ProjectBucketConfigIndexConfig>[]>;
    /**
     * The location of the bucket.
     */
    location: pulumi.Input<string>;
    /**
     * Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
     */
    locked?: pulumi.Input<boolean>;
    /**
     * The parent resource that contains the logging bucket.
     */
    project: pulumi.Input<string>;
    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
     */
    retentionDays?: pulumi.Input<number>;
}
