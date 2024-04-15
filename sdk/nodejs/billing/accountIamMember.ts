// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on billing accounts. Each of these resources serves a different use case:
 *
 * * `gcp.billing.AccountIamPolicy`: Authoritative. Sets the IAM policy for the billing accounts and replaces any existing policy already attached.
 * * `gcp.billing.AccountIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
 * * `gcp.billing.AccountIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the billing accounts are preserved.
 *
 * > **Note:** `gcp.billing.AccountIamPolicy` **cannot** be used in conjunction with `gcp.billing.AccountIamBinding` and `gcp.billing.AccountIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the billing account as `gcp.billing.AccountIamPolicy` replaces the entire policy.
 *
 * > **Note:** `gcp.billing.AccountIamBinding` resources **can be** used in conjunction with `gcp.billing.AccountIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_billing\_account\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/billing.viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const editor = new gcp.billing.AccountIamPolicy("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_billing\_account\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.billing.AccountIamBinding("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     role: "roles/billing.viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_billing\_account\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.billing.AccountIamMember("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     role: "roles/billing.viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_billing\_account\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/billing.viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const editor = new gcp.billing.AccountIamPolicy("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_billing\_account\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.billing.AccountIamBinding("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     role: "roles/billing.viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_billing\_account\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.billing.AccountIamMember("editor", {
 *     billingAccountId: "00AA00-000AAA-00AA0A",
 *     role: "roles/billing.viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ### Importing IAM policies
 *
 * IAM policy imports use the `billing_account_id` identifier of the Billing Account resource only. For example:
 *
 * * `{{billing_account_id}}`
 *
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 *
 * tf
 *
 * import {
 *
 *   id = {{billing_account_id}}
 *
 *   to = google_billing_account_iam_policy.default
 *
 * }
 *
 * The `pulumi import` command can also be used:
 *
 * ```sh
 * $ pulumi import gcp:billing/accountIamMember:AccountIamMember default {{billing_account_id}}
 * ```
 */
export class AccountIamMember extends pulumi.CustomResource {
    /**
     * Get an existing AccountIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountIamMemberState, opts?: pulumi.CustomResourceOptions): AccountIamMember {
        return new AccountIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:billing/accountIamMember:AccountIamMember';

    /**
     * Returns true if the given object is an instance of AccountIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountIamMember.__pulumiType;
    }

    /**
     * The billing account id.
     *
     * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
     */
    public readonly billingAccountId!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.billing.AccountIamMemberCondition | undefined>;
    /**
     * (Computed) The etag of the billing account's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     *
     * `gcp.billing.AccountIamPolicy` only:
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a AccountIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountIamMemberArgs | AccountIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountIamMemberState | undefined;
            resourceInputs["billingAccountId"] = state ? state.billingAccountId : undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as AccountIamMemberArgs | undefined;
            if ((!args || args.billingAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccountId'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["billingAccountId"] = args ? args.billingAccountId : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountIamMember resources.
 */
export interface AccountIamMemberState {
    /**
     * The billing account id.
     *
     * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
     */
    billingAccountId?: pulumi.Input<string>;
    condition?: pulumi.Input<inputs.billing.AccountIamMemberCondition>;
    /**
     * (Computed) The etag of the billing account's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     *
     * `gcp.billing.AccountIamPolicy` only:
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountIamMember resource.
 */
export interface AccountIamMemberArgs {
    /**
     * The billing account id.
     *
     * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
     */
    billingAccountId: pulumi.Input<string>;
    condition?: pulumi.Input<inputs.billing.AccountIamMemberCondition>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     *
     * `gcp.billing.AccountIamPolicy` only:
     */
    role: pulumi.Input<string>;
}
