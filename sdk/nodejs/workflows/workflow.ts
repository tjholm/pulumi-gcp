// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Workflow program to be executed by Workflows.
 *
 * To get more information about Workflow, see:
 *
 * * [API documentation](https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows)
 * * How-to Guides
 *     * [Managing Workflows](https://cloud.google.com/workflows/docs/creating-updating-workflow)
 *
 * ## Example Usage
 * ### Workflow Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const testAccount = new gcp.serviceaccount.Account("testAccount", {
 *     accountId: "my-account",
 *     displayName: "Test Service Account",
 * });
 * const example = new gcp.workflows.Workflow("example", {
 *     region: "us-central1",
 *     description: "Magic",
 *     serviceAccount: testAccount.id,
 *     labels: {
 *         env: "test",
 *     },
 *     sourceContents: `# This is a sample workflow. You can replace it with your source code.
 * #
 * # This workflow does the following:
 * # - reads current time and date information from an external API and stores
 * #   the response in currentTime variable
 * # - retrieves a list of Wikipedia articles related to the day of the week
 * #   from currentTime
 * # - returns the list of articles as an output of the workflow
 * #
 * # Note: In Terraform you need to escape the $$ or it will cause errors.
 *
 * - getCurrentTime:
 *     call: http.get
 *     args:
 *         url: https://timeapi.io/api/Time/current/zone?timeZone=Europe/Amsterdam
 *     result: currentTime
 * - readWikipedia:
 *     call: http.get
 *     args:
 *         url: https://en.wikipedia.org/w/api.php
 *         query:
 *             action: opensearch
 *             search: ${currentTime.body.dayOfWeek}
 *     result: wikiResult
 * - returnOutput:
 *     return: ${wikiResult.body[1]}
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowState, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:workflows/workflow:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The KMS key used to encrypt workflow and execution data.
     * Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
     */
    public readonly cryptoKeyName!: pulumi.Output<string | undefined>;
    /**
     * Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A set of key/value label pairs to assign to this Workflow.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the Workflow.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and name are unspecified, a random value is chosen for the name.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The region of the workflow.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The revision of the workflow. A new one is generated if the service account or source contents is changed.
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * Name of the service account associated with the latest workflow version. This service
     * account represents the identity of the workflow and determines what permissions the workflow has.
     * Format: projects/{project}/serviceAccounts/{account} or {account}.
     * Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
     * The {account} value can be the email address or the uniqueId of the service account.
     * If not provided, workflow will use the project's default service account.
     * Modifying this field for an existing workflow results in a new workflow revision.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Workflow code to be executed. The size limit is 32KB.
     */
    public readonly sourceContents!: pulumi.Output<string | undefined>;
    /**
     * State of the workflow deployment.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowArgs | WorkflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["cryptoKeyName"] = state ? state.cryptoKeyName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["revisionId"] = state ? state.revisionId : undefined;
            resourceInputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            resourceInputs["sourceContents"] = state ? state.sourceContents : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as WorkflowArgs | undefined;
            resourceInputs["cryptoKeyName"] = args ? args.cryptoKeyName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["sourceContents"] = args ? args.sourceContents : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Workflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workflow resources.
 */
export interface WorkflowState {
    /**
     * The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The KMS key used to encrypt workflow and execution data.
     * Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
     */
    cryptoKeyName?: pulumi.Input<string>;
    /**
     * Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of key/value label pairs to assign to this Workflow.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Workflow.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and name are unspecified, a random value is chosen for the name.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The region of the workflow.
     */
    region?: pulumi.Input<string>;
    /**
     * The revision of the workflow. A new one is generated if the service account or source contents is changed.
     */
    revisionId?: pulumi.Input<string>;
    /**
     * Name of the service account associated with the latest workflow version. This service
     * account represents the identity of the workflow and determines what permissions the workflow has.
     * Format: projects/{project}/serviceAccounts/{account} or {account}.
     * Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
     * The {account} value can be the email address or the uniqueId of the service account.
     * If not provided, workflow will use the project's default service account.
     * Modifying this field for an existing workflow results in a new workflow revision.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Workflow code to be executed. The size limit is 32KB.
     */
    sourceContents?: pulumi.Input<string>;
    /**
     * State of the workflow deployment.
     */
    state?: pulumi.Input<string>;
    /**
     * The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * The KMS key used to encrypt workflow and execution data.
     * Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
     */
    cryptoKeyName?: pulumi.Input<string>;
    /**
     * Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
     */
    description?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to this Workflow.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Workflow.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and name are unspecified, a random value is chosen for the name.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the workflow.
     */
    region?: pulumi.Input<string>;
    /**
     * Name of the service account associated with the latest workflow version. This service
     * account represents the identity of the workflow and determines what permissions the workflow has.
     * Format: projects/{project}/serviceAccounts/{account} or {account}.
     * Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
     * The {account} value can be the email address or the uniqueId of the service account.
     * If not provided, workflow will use the project's default service account.
     * Modifying this field for an existing workflow results in a new workflow revision.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Workflow code to be executed. The size limit is 32KB.
     */
    sourceContents?: pulumi.Input<string>;
}
