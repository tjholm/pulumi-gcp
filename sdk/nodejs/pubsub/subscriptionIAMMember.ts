// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
 *
 * * `gcp.pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
 * * `gcp.pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
 * * `gcp.pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
 *
 * > **Note:** `gcp.pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `gcp.pubsub.SubscriptionIAMBinding` and `gcp.pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `gcp.pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## gcp.pubsub.SubscriptionIAMPolicy
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
 * const editor = new gcp.pubsub.SubscriptionIAMPolicy("editor", {
 *     subscription: "your-subscription-name",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## gcp.pubsub.SubscriptionIAMBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.pubsub.SubscriptionIAMBinding("editor", {
 *     subscription: "your-subscription-name",
 *     role: "roles/editor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.pubsub.SubscriptionIAMMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.pubsub.SubscriptionIAMMember("editor", {
 *     subscription: "your-subscription-name",
 *     role: "roles/editor",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## gcp.pubsub.SubscriptionIAMBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.pubsub.SubscriptionIAMBinding("editor", {
 *     subscription: "your-subscription-name",
 *     role: "roles/editor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.pubsub.SubscriptionIAMMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.pubsub.SubscriptionIAMMember("editor", {
 *     subscription: "your-subscription-name",
 *     role: "roles/editor",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * ### Importing IAM policies
 *
 * IAM policy imports use the identifier of the Pubsub Subscription resource. For example:
 *
 * * `"projects/{{project_id}}/subscriptions/{{subscription}}"`
 *
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 *
 * tf
 *
 * import {
 *
 *   id = "projects/{{project_id}}/subscriptions/{{subscription}}"
 *
 *   to = google_pubsub_subscription_iam_policy.default
 *
 * }
 *
 * The `pulumi import` command can also be used:
 *
 * ```sh
 * $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember default projects/{{project_id}}/subscriptions/{{subscription}}
 * ```
 */
export class SubscriptionIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing SubscriptionIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscriptionIAMMemberState, opts?: pulumi.CustomResourceOptions): SubscriptionIAMMember {
        return new SubscriptionIAMMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember';

    /**
     * Returns true if the given object is an instance of SubscriptionIAMMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscriptionIAMMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscriptionIAMMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.pubsub.SubscriptionIAMMemberCondition | undefined>;
    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    public readonly member!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    public readonly subscription!: pulumi.Output<string>;

    /**
     * Create a SubscriptionIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscriptionIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscriptionIAMMemberArgs | SubscriptionIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscriptionIAMMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["subscription"] = state ? state.subscription : undefined;
        } else {
            const args = argsOrState as SubscriptionIAMMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.subscription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscription'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["subscription"] = args ? args.subscription : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubscriptionIAMMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscriptionIAMMember resources.
 */
export interface SubscriptionIAMMemberState {
    condition?: pulumi.Input<inputs.pubsub.SubscriptionIAMMemberCondition>;
    /**
     * (Computed) The etag of the subscription's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    member?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    subscription?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubscriptionIAMMember resource.
 */
export interface SubscriptionIAMMemberArgs {
    condition?: pulumi.Input<inputs.pubsub.SubscriptionIAMMemberCondition>;
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     */
    member: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
    /**
     * The subscription name or id to bind to attach IAM policy to.
     */
    subscription: pulumi.Input<string>;
}
