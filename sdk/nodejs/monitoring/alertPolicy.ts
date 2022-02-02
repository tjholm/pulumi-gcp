// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A description of the conditions under which some aspect of your system is
 * considered to be "unhealthy" and the ways to notify people or services
 * about this state.
 *
 * To get more information about AlertPolicy, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/alerts/)
 *
 * ## Example Usage
 * ### Monitoring Alert Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const alertPolicy = new gcp.monitoring.AlertPolicy("alert_policy", {
 *     combiner: "OR",
 *     conditions: [{
 *         conditionThreshold: {
 *             aggregations: [{
 *                 alignmentPeriod: "60s",
 *                 perSeriesAligner: "ALIGN_RATE",
 *             }],
 *             comparison: "COMPARISON_GT",
 *             duration: "60s",
 *             filter: "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
 *         },
 *         displayName: "test condition",
 *     }],
 *     displayName: "My Alert Policy",
 *     userLabels: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * AlertPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/alertPolicy:AlertPolicy default {{name}}
 * ```
 */
export class AlertPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AlertPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertPolicyState, opts?: pulumi.CustomResourceOptions): AlertPolicy {
        return new AlertPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/alertPolicy:AlertPolicy';

    /**
     * Returns true if the given object is an instance of AlertPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertPolicy.__pulumiType;
    }

    /**
     * Control over how this alert policy's notification channels are notified.
     * Structure is documented below.
     */
    public readonly alertStrategy!: pulumi.Output<outputs.monitoring.AlertPolicyAlertStrategy | undefined>;
    /**
     * How to combine the results of multiple conditions to
     * determine if an incident should be opened.
     * Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
     */
    public readonly combiner!: pulumi.Output<string>;
    /**
     * A list of conditions for the policy. The conditions are combined by
     * AND or OR according to the combiner field. If the combined conditions
     * evaluate to true, then an incident is created. A policy can have from
     * one to six conditions.
     * Structure is documented below.
     */
    public readonly conditions!: pulumi.Output<outputs.monitoring.AlertPolicyCondition[]>;
    /**
     * A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
     * ignored.
     */
    public /*out*/ readonly creationRecords!: pulumi.Output<outputs.monitoring.AlertPolicyCreationRecord[]>;
    /**
     * A short name or phrase used to identify the
     * condition in dashboards, notifications, and
     * incidents. To avoid confusion, don't use the same
     * display name for multiple conditions in the same
     * policy.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Documentation that is included with notifications and incidents related
     * to this policy. Best practice is for the documentation to include information
     * to help responders understand, mitigate, escalate, and correct the underlying
     * problems detected by the alerting policy. Notification channels that have
     * limited capacity might not show this documentation.
     * Structure is documented below.
     */
    public readonly documentation!: pulumi.Output<outputs.monitoring.AlertPolicyDocumentation | undefined>;
    /**
     * Whether or not the policy is enabled. The default is true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * -
     * The unique resource name for this condition.
     * Its syntax is:
     * projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
     * [CONDITION_ID] is assigned by Stackdriver Monitoring when
     * the condition is created as part of a new or updated alerting
     * policy.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Identifies the notification channels to which notifications should be
     * sent when incidents are opened or closed or when new violations occur
     * on an already opened incident. Each element of this array corresponds
     * to the name field in each of the NotificationChannel objects that are
     * returned from the notificationChannels.list method. The syntax of the
     * entries in this field is
     * `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
     */
    public readonly notificationChannels!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * This field is intended to be used for organizing and identifying the AlertPolicy
     * objects.The field can contain up to 64 entries. Each key and value is limited
     * to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
     * can contain only lowercase letters, numerals, underscores, and dashes. Keys
     * must begin with a letter.
     */
    public readonly userLabels!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a AlertPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertPolicyArgs | AlertPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertPolicyState | undefined;
            resourceInputs["alertStrategy"] = state ? state.alertStrategy : undefined;
            resourceInputs["combiner"] = state ? state.combiner : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["creationRecords"] = state ? state.creationRecords : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["documentation"] = state ? state.documentation : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationChannels"] = state ? state.notificationChannels : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["userLabels"] = state ? state.userLabels : undefined;
        } else {
            const args = argsOrState as AlertPolicyArgs | undefined;
            if ((!args || args.combiner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'combiner'");
            }
            if ((!args || args.conditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conditions'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["alertStrategy"] = args ? args.alertStrategy : undefined;
            resourceInputs["combiner"] = args ? args.combiner : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["documentation"] = args ? args.documentation : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["notificationChannels"] = args ? args.notificationChannels : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["userLabels"] = args ? args.userLabels : undefined;
            resourceInputs["creationRecords"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlertPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertPolicy resources.
 */
export interface AlertPolicyState {
    /**
     * Control over how this alert policy's notification channels are notified.
     * Structure is documented below.
     */
    alertStrategy?: pulumi.Input<inputs.monitoring.AlertPolicyAlertStrategy>;
    /**
     * How to combine the results of multiple conditions to
     * determine if an incident should be opened.
     * Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
     */
    combiner?: pulumi.Input<string>;
    /**
     * A list of conditions for the policy. The conditions are combined by
     * AND or OR according to the combiner field. If the combined conditions
     * evaluate to true, then an incident is created. A policy can have from
     * one to six conditions.
     * Structure is documented below.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.monitoring.AlertPolicyCondition>[]>;
    /**
     * A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be
     * ignored.
     */
    creationRecords?: pulumi.Input<pulumi.Input<inputs.monitoring.AlertPolicyCreationRecord>[]>;
    /**
     * A short name or phrase used to identify the
     * condition in dashboards, notifications, and
     * incidents. To avoid confusion, don't use the same
     * display name for multiple conditions in the same
     * policy.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Documentation that is included with notifications and incidents related
     * to this policy. Best practice is for the documentation to include information
     * to help responders understand, mitigate, escalate, and correct the underlying
     * problems detected by the alerting policy. Notification channels that have
     * limited capacity might not show this documentation.
     * Structure is documented below.
     */
    documentation?: pulumi.Input<inputs.monitoring.AlertPolicyDocumentation>;
    /**
     * Whether or not the policy is enabled. The default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * -
     * The unique resource name for this condition.
     * Its syntax is:
     * projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
     * [CONDITION_ID] is assigned by Stackdriver Monitoring when
     * the condition is created as part of a new or updated alerting
     * policy.
     */
    name?: pulumi.Input<string>;
    /**
     * Identifies the notification channels to which notifications should be
     * sent when incidents are opened or closed or when new violations occur
     * on an already opened incident. Each element of this array corresponds
     * to the name field in each of the NotificationChannel objects that are
     * returned from the notificationChannels.list method. The syntax of the
     * entries in this field is
     * `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
     */
    notificationChannels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * This field is intended to be used for organizing and identifying the AlertPolicy
     * objects.The field can contain up to 64 entries. Each key and value is limited
     * to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
     * can contain only lowercase letters, numerals, underscores, and dashes. Keys
     * must begin with a letter.
     */
    userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AlertPolicy resource.
 */
export interface AlertPolicyArgs {
    /**
     * Control over how this alert policy's notification channels are notified.
     * Structure is documented below.
     */
    alertStrategy?: pulumi.Input<inputs.monitoring.AlertPolicyAlertStrategy>;
    /**
     * How to combine the results of multiple conditions to
     * determine if an incident should be opened.
     * Possible values are `AND`, `OR`, and `AND_WITH_MATCHING_RESOURCE`.
     */
    combiner: pulumi.Input<string>;
    /**
     * A list of conditions for the policy. The conditions are combined by
     * AND or OR according to the combiner field. If the combined conditions
     * evaluate to true, then an incident is created. A policy can have from
     * one to six conditions.
     * Structure is documented below.
     */
    conditions: pulumi.Input<pulumi.Input<inputs.monitoring.AlertPolicyCondition>[]>;
    /**
     * A short name or phrase used to identify the
     * condition in dashboards, notifications, and
     * incidents. To avoid confusion, don't use the same
     * display name for multiple conditions in the same
     * policy.
     */
    displayName: pulumi.Input<string>;
    /**
     * Documentation that is included with notifications and incidents related
     * to this policy. Best practice is for the documentation to include information
     * to help responders understand, mitigate, escalate, and correct the underlying
     * problems detected by the alerting policy. Notification channels that have
     * limited capacity might not show this documentation.
     * Structure is documented below.
     */
    documentation?: pulumi.Input<inputs.monitoring.AlertPolicyDocumentation>;
    /**
     * Whether or not the policy is enabled. The default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Identifies the notification channels to which notifications should be
     * sent when incidents are opened or closed or when new violations occur
     * on an already opened incident. Each element of this array corresponds
     * to the name field in each of the NotificationChannel objects that are
     * returned from the notificationChannels.list method. The syntax of the
     * entries in this field is
     * `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`
     */
    notificationChannels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * This field is intended to be used for organizing and identifying the AlertPolicy
     * objects.The field can contain up to 64 entries. Each key and value is limited
     * to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
     * can contain only lowercase letters, numerals, underscores, and dashes. Keys
     * must begin with a letter.
     */
    userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
