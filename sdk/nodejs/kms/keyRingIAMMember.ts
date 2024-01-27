// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
 *
 * * `gcp.kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
 * * `gcp.kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
 * * `gcp.kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
 *
 * > **Note:** `gcp.kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `gcp.kms.KeyRingIAMBinding` and `gcp.kms.KeyRingIAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.kms.KeyRingIAMBinding` resources **can be** used in conjunction with `gcp.kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_kms\_key\_ring\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const keyRing = new gcp.kms.KeyRingIAMPolicy("keyRing", {
 *     keyRingId: keyring.id,
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
 * const keyring = new gcp.kms.KeyRing("keyring", {location: "global"});
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const keyRing = new gcp.kms.KeyRingIAMPolicy("keyRing", {
 *     keyRingId: keyring.id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_kms\_key\_ring\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMBinding("keyRing", {
 *     keyRingId: "your-key-ring-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/cloudkms.admin",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMBinding("keyRing", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     keyRingId: "your-key-ring-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/cloudkms.admin",
 * });
 * ```
 *
 * ## google\_kms\_key\_ring\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMMember("keyRing", {
 *     keyRingId: "your-key-ring-id",
 *     member: "user:jane@example.com",
 *     role: "roles/cloudkms.admin",
 * });
 * ```
 *
 * With IAM Conditions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRingIAMMember("keyRing", {
 *     condition: {
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         title: "expires_after_2019_12_31",
 *     },
 *     keyRingId: "your-key-ring-id",
 *     member: "user:jane@example.com",
 *     role: "roles/cloudkms.admin",
 * });
 * ```
 *
 * ## Import
 *
 * ### Importing IAM policies IAM policy imports use the identifier of the Cloud KMS key ring only. For example* `{{project_id}}/{{location}}/{{key_ring_name}}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
 *
 *  id = "{{project_id}}/{{location}}/{{key_ring_name}}"
 *
 *  to = google_kms_key_ring_iam_policy.default } The `pulumi import` command can also be used
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMMember:KeyRingIAMMember default {{project_id}}/{{location}}/{{key_ring_name}}
 * ```
 */
export class KeyRingIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing KeyRingIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyRingIAMMemberState, opts?: pulumi.CustomResourceOptions): KeyRingIAMMember {
        return new KeyRingIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/keyRingIAMMember:KeyRingIAMMember';

    /**
     * Returns true if the given object is an instance of KeyRingIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyRingIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyRingIAMMember.__pulumiType;
    }

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.kms.KeyRingIAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
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
    public readonly keyRingId!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a KeyRingIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyRingIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyRingIAMMemberArgs | KeyRingIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyRingIAMMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["keyRingId"] = state ? state.keyRingId : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as KeyRingIAMMemberArgs | undefined;
            if ((!args || args.keyRingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRingId'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["keyRingId"] = args ? args.keyRingId : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeyRingIAMMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyRingIAMMember resources.
 */
export interface KeyRingIAMMemberState {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.kms.KeyRingIAMMemberCondition>;
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
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
    keyRingId?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyRingIAMMember resource.
 */
export interface KeyRingIAMMemberArgs {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.kms.KeyRingIAMMemberCondition>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
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
    keyRingId: pulumi.Input<string>;
    member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
