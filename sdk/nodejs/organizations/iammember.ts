// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Four different resources help you manage your IAM policy for a organization. Each of these resources serves a different use case:
 *
 * * `gcp.organizations.IAMPolicy`: Authoritative. Sets the IAM policy for the organization and replaces any existing policy already attached.
 * * `gcp.organizations.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the organization are preserved.
 * * `gcp.organizations.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the organization are preserved.
 * * `gcp.organizations.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
 *
 * > **Note:** `gcp.organizations.IAMPolicy` **cannot** be used in conjunction with `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.organizations.IAMBinding` resources **can be** used in conjunction with `gcp.organizations.IAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_organization\_iam\_policy
 *
 * !> **Warning:** New organizations have several default policies which will,
 *    without extreme caution, be **overwritten** by use of this resource.
 *    The safest alternative is to use multiple `gcp.organizations.IAMBinding`
 *    resources. This resource makes it easy to remove your own access to
 *    an organization, which will require a call to Google Support to have
 *    fixed, and can take multiple days to resolve.
 *
 *    In general, this resource should only be used with organizations
 *    fully managed by this provider.I f you do use this resource,
 *    the best way to be sure that you are not making dangerous changes is to start
 *    by **importing** your existing policy, and examining the diff very closely.
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
 * const organization = new gcp.organizations.IAMPolicy("organization", {
 *     orgId: "1234567890",
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
 *         condition: {
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *             title: "expires_after_2019_12_31",
 *         },
 *         members: ["user:jane@example.com"],
 *         role: "roles/editor",
 *     }],
 * });
 * const organization = new gcp.organizations.IAMPolicy("organization", {
 *     orgId: "1234567890",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_organization\_iam\_binding
 *
 * > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your organization.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organization = new gcp.organizations.IAMBinding("organization", {
 *     members: ["user:jane@example.com"],
 *     orgId: "1234567890",
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organization = new gcp.organizations.IAMBinding("organization", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     members: ["user:jane@example.com"],
 *     orgId: "1234567890",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_organization\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organization = new gcp.organizations.IAMMember("organization", {
 *     member: "user:jane@example.com",
 *     orgId: "1234567890",
 *     role: "roles/editor",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organization = new gcp.organizations.IAMMember("organization", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     member: "user:jane@example.com",
 *     orgId: "1234567890",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_organization\_iam\_audit\_config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organization = new gcp.organizations.IamAuditConfig("organization", {
 *     auditLogConfigs: [
 *         {
 *             logType: "ADMIN_READ",
 *         },
 *         {
 *             exemptedMembers: ["user:joebloggs@hashicorp.com"],
 *             logType: "DATA_READ",
 *         },
 *     ],
 *     orgId: "1234567890",
 *     service: "allServices",
 * });
 * ```
 *
 * ## Import
 *
 * ### Importing Audit Configs An audit config can be imported into a `google_organization_iam_audit_config` resource using the resource's `org_id` and the `service`, e.g* `"{{org_id}} foo.googleapis.com"` An `import` block (Terraform v1.5.0 and later) can be used to import audit configstf import {
 *
 *  id = "{{org_id}} foo.googleapis.com"
 *
 *  to = google_organization_iam_audit_config.default } The `pulumi import` command can also be used
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iAMMember:IAMMember default "{{org_id}} foo.googleapis.com"
 * ```
 */
export class IAMMember extends pulumi.CustomResource {
    /**
     * Get an existing IAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMMemberState, opts?: pulumi.CustomResourceOptions): IAMMember {
        return new IAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/iAMMember:IAMMember';

    /**
     * Returns true if the given object is an instance of IAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMMember.__pulumiType;
    }

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.organizations.IAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The organization id of the target organization.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a IAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMMemberArgs | IAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAMMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as IAMMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IAMMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMMember resources.
 */
export interface IAMMemberState {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.organizations.IAMMemberCondition>;
    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    etag?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * The organization id of the target organization.
     */
    orgId?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMMember resource.
 */
export interface IAMMemberArgs {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.organizations.IAMMemberCondition>;
    member: pulumi.Input<string>;
    /**
     * The organization id of the target organization.
     */
    orgId: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     */
    role: pulumi.Input<string>;
}
