// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Access Context Manager (VPC Service Controls) AccessPolicy. Each of these resources serves a different use case:
 *
 * * `gcp.accesscontextmanager.AccessPolicyIamPolicy`: Authoritative. Sets the IAM policy for the accesspolicy and replaces any existing policy already attached.
 * * `gcp.accesscontextmanager.AccessPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the accesspolicy are preserved.
 * * `gcp.accesscontextmanager.AccessPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the accesspolicy are preserved.
 *
 * > **Note:** `gcp.accesscontextmanager.AccessPolicyIamPolicy` **cannot** be used in conjunction with `gcp.accesscontextmanager.AccessPolicyIamBinding` and `gcp.accesscontextmanager.AccessPolicyIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.accesscontextmanager.AccessPolicyIamBinding` resources **can be** used in conjunction with `gcp.accesscontextmanager.AccessPolicyIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_access\_context\_manager\_access\_policy\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/accesscontextmanager.policyAdmin",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.accesscontextmanager.AccessPolicyIamPolicy("policy", {policyData: admin.then(admin => admin.policyData)});
 * ```
 *
 * ## google\_access\_context\_manager\_access\_policy\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.accesscontextmanager.AccessPolicyIamBinding("binding", {
 *     role: "roles/accesscontextmanager.policyAdmin",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_access\_context\_manager\_access\_policy\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.accesscontextmanager.AccessPolicyIamMember("member", {
 *     role: "roles/accesscontextmanager.policyAdmin",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* accessPolicies/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Access Context Manager (VPC Service Controls) accesspolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/accessPolicyIamMember:AccessPolicyIamMember editor "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/accessPolicyIamMember:AccessPolicyIamMember editor "accessPolicies/{{access_policy}} roles/accesscontextmanager.policyAdmin"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/accessPolicyIamMember:AccessPolicyIamMember editor accessPolicies/{{access_policy}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class AccessPolicyIamMember extends pulumi.CustomResource {
    /**
     * Get an existing AccessPolicyIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPolicyIamMemberState, opts?: pulumi.CustomResourceOptions): AccessPolicyIamMember {
        return new AccessPolicyIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/accessPolicyIamMember:AccessPolicyIamMember';

    /**
     * Returns true if the given object is an instance of AccessPolicyIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPolicyIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPolicyIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.accesscontextmanager.AccessPolicyIamMemberCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a AccessPolicyIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPolicyIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPolicyIamMemberArgs | AccessPolicyIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPolicyIamMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as AccessPolicyIamMemberArgs | undefined;
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPolicyIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPolicyIamMember resources.
 */
export interface AccessPolicyIamMemberState {
    condition?: pulumi.Input<inputs.accesscontextmanager.AccessPolicyIamMemberCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessPolicyIamMember resource.
 */
export interface AccessPolicyIamMemberArgs {
    condition?: pulumi.Input<inputs.accesscontextmanager.AccessPolicyIamMemberCondition>;
    member: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.accesscontextmanager.AccessPolicyIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
