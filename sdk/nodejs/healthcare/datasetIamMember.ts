// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:
 *
 * * `gcp.healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
 * * `gcp.healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
 * * `gcp.healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
 *
 * > **Note:** `gcp.healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.DatasetIamBinding` and `gcp.healthcare.DatasetIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.healthcare.DatasetIamBinding` resources **can be** used in conjunction with `gcp.healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_healthcare\_dataset\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const dataset = new gcp.healthcare.DatasetIamPolicy("dataset", {
 *     datasetId: "your-dataset-id",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_healthcare\_dataset\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dataset = new gcp.healthcare.DatasetIamBinding("dataset", {
 *     datasetId: "your-dataset-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_healthcare\_dataset\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dataset = new gcp.healthcare.DatasetIamMember("dataset", {
 *     datasetId: "your-dataset-id",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## Import
 *
 * ### Importing IAM policies IAM policy imports use the identifier of the Healthcase Dataset resource. For example* `"{{project_id}}/{{location}}/{{dataset}}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
 *
 *  id = "{{project_id}}/{{location}}/{{dataset}}"
 *
 *  to = google_healthcare_dataset_iam_policy.default } The `pulumi import` command can also be used
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/datasetIamMember:DatasetIamMember default {{project_id}}/{{location}}/{{dataset}}
 * ```
 */
export class DatasetIamMember extends pulumi.CustomResource {
    /**
     * Get an existing DatasetIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetIamMemberState, opts?: pulumi.CustomResourceOptions): DatasetIamMember {
        return new DatasetIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/datasetIamMember:DatasetIamMember';

    /**
     * Returns true if the given object is an instance of DatasetIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.healthcare.DatasetIamMemberCondition | undefined>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    public readonly datasetId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the dataset's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a DatasetIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetIamMemberArgs | DatasetIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatasetIamMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["datasetId"] = state ? state.datasetId : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DatasetIamMemberArgs | undefined;
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["datasetId"] = args ? args.datasetId : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatasetIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetIamMember resources.
 */
export interface DatasetIamMemberState {
    condition?: pulumi.Input<inputs.healthcare.DatasetIamMemberCondition>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    datasetId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the dataset's IAM policy.
     */
    etag?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatasetIamMember resource.
 */
export interface DatasetIamMemberArgs {
    condition?: pulumi.Input<inputs.healthcare.DatasetIamMemberCondition>;
    /**
     * The dataset ID, in the form
     * `{project_id}/{location_name}/{dataset_name}` or
     * `{location_name}/{dataset_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    datasetId: pulumi.Input<string>;
    member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
