// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the googleProjectIam set of resources.
 *
 * Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
 *
 * * `gcp.serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
 * * `gcp.serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
 * * `gcp.serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
 *
 * > **Note:** `gcp.serviceAccount.IAMPolicy` **cannot** be used in conjunction with `gcp.serviceAccount.IAMBinding` and `gcp.serviceAccount.IAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.serviceAccount.IAMBinding` resources **can be** used in conjunction with `gcp.serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_service\_account\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iam.serviceAccountUser",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const sa = new gcp.serviceaccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can interact with",
 * });
 * const admin_account_iam = new gcp.serviceaccount.IAMPolicy("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_service\_account\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sa = new gcp.serviceaccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceaccount.IAMBinding("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
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
 * const sa = new gcp.serviceaccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that only Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceaccount.IAMBinding("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 *
 * ## google\_service\_account\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.compute.getDefaultServiceAccount({});
 * const sa = new gcp.serviceaccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceaccount.IAMMember("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
 *     member: "user:jane@example.com",
 * });
 * // Allow SA service account use the default GCE account
 * const gce_default_account_iam = new gcp.serviceaccount.IAMMember("gce-default-account-iam", {
 *     serviceAccountId: _default.then(_default => _default.name),
 *     role: "roles/iam.serviceAccountUser",
 *     member: pulumi.interpolate`serviceAccount:${sa.email}`,
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sa = new gcp.serviceaccount.Account("sa", {
 *     accountId: "my-service-account",
 *     displayName: "A service account that Jane can use",
 * });
 * const admin_account_iam = new gcp.serviceaccount.IAMMember("admin-account-iam", {
 *     serviceAccountId: sa.name,
 *     role: "roles/iam.serviceAccountUser",
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
 * Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
 * ```
 */
export class IAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing IAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMBindingState, opts?: pulumi.CustomResourceOptions): IAMBinding {
        return new IAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:serviceAccount/iAMBinding:IAMBinding';

    /**
     * Returns true if the given object is an instance of IAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMBinding.__pulumiType;
    }

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.serviceAccount.IAMBindingCondition | undefined>;
    /**
     * (Computed) The etag of the service account IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;

    /**
     * Create a IAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMBindingArgs | IAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAMBindingState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
        } else {
            const args = argsOrState as IAMBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IAMBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMBinding resources.
 */
export interface IAMBindingState {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.serviceAccount.IAMBindingCondition>;
    /**
     * (Computed) The etag of the service account IAM policy.
     */
    etag?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    serviceAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMBinding resource.
 */
export interface IAMBindingArgs {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.serviceAccount.IAMBindingCondition>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
    /**
     * The fully-qualified name of the service account to apply policy to.
     */
    serviceAccountId: pulumi.Input<string>;
}
