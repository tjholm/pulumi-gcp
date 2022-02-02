// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:
 *
 * * `gcp.bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
 * * `gcp.bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
 * * `gcp.bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
 *
 * > **Note:** `gcp.bigquery.IamPolicy` **cannot** be used in conjunction with `gcp.bigquery.IamBinding` and `gcp.bigquery.IamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.bigquery.IamBinding` resources **can be** used in conjunction with `gcp.bigquery.IamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_bigquery\_table\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/bigquery.dataOwner",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.bigquery.IamPolicy("policy", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/bigquery.dataOwner",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.bigquery.IamPolicy("policy", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * ## google\_bigquery\_table\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.bigquery.IamBinding("binding", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     role: "roles/bigquery.dataOwner",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.bigquery.IamBinding("binding", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     role: "roles/bigquery.dataOwner",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * ## google\_bigquery\_table\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.bigquery.IamMember("member", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     role: "roles/bigquery.dataOwner",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.bigquery.IamMember("member", {
 *     project: google_bigquery_table.test.project,
 *     datasetId: google_bigquery_table.test.dataset_id,
 *     tableId: google_bigquery_table.test.table_id,
 *     role: "roles/bigquery.dataOwner",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} * {{project}}/{{dataset_id}}/{{table_id}} * {{dataset_id}}/{{table_id}} * {{table_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery table IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class IamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamPolicyState, opts?: pulumi.CustomResourceOptions): IamPolicy {
        return new IamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/iamPolicy:IamPolicy';

    /**
     * Returns true if the given object is an instance of IamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamPolicy.__pulumiType;
    }

    public readonly datasetId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly tableId!: pulumi.Output<string>;

    /**
     * Create a IamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamPolicyArgs | IamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamPolicyState | undefined;
            resourceInputs["datasetId"] = state ? state.datasetId : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["tableId"] = state ? state.tableId : undefined;
        } else {
            const args = argsOrState as IamPolicyArgs | undefined;
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.tableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableId'");
            }
            resourceInputs["datasetId"] = args ? args.datasetId : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tableId"] = args ? args.tableId : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamPolicy resources.
 */
export interface IamPolicyState {
    datasetId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    tableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamPolicy resource.
 */
export interface IamPolicyArgs {
    datasetId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    tableId: pulumi.Input<string>;
}
