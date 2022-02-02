// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Data catalog EntryGroup. Each of these resources serves a different use case:
 *
 * * `gcp.datacatalog.EntryGroupIamPolicy`: Authoritative. Sets the IAM policy for the entrygroup and replaces any existing policy already attached.
 * * `gcp.datacatalog.EntryGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the entrygroup are preserved.
 * * `gcp.datacatalog.EntryGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the entrygroup are preserved.
 *
 * > **Note:** `gcp.datacatalog.EntryGroupIamPolicy` **cannot** be used in conjunction with `gcp.datacatalog.EntryGroupIamBinding` and `gcp.datacatalog.EntryGroupIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.datacatalog.EntryGroupIamBinding` resources **can be** used in conjunction with `gcp.datacatalog.EntryGroupIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_data\_catalog\_entry\_group\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.datacatalog.EntryGroupIamPolicy("policy", {
 *     entryGroup: google_data_catalog_entry_group.basic_entry_group.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_data\_catalog\_entry\_group\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.datacatalog.EntryGroupIamBinding("binding", {
 *     entryGroup: google_data_catalog_entry_group.basic_entry_group.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_data\_catalog\_entry\_group\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.datacatalog.EntryGroupIamMember("member", {
 *     entryGroup: google_data_catalog_entry_group.basic_entry_group.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} * {{project}}/{{region}}/{{entry_group}} * {{region}}/{{entry_group}} * {{entry_group}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog entrygroup IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/entryGroupIamMember:EntryGroupIamMember editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/entryGroupIamMember:EntryGroupIamMember editor "projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:datacatalog/entryGroupIamMember:EntryGroupIamMember editor projects/{{project}}/locations/{{region}}/entryGroups/{{entry_group}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class EntryGroupIamMember extends pulumi.CustomResource {
    /**
     * Get an existing EntryGroupIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntryGroupIamMemberState, opts?: pulumi.CustomResourceOptions): EntryGroupIamMember {
        return new EntryGroupIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/entryGroupIamMember:EntryGroupIamMember';

    /**
     * Returns true if the given object is an instance of EntryGroupIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntryGroupIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntryGroupIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.datacatalog.EntryGroupIamMemberCondition | undefined>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly entryGroup!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a EntryGroupIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntryGroupIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntryGroupIamMemberArgs | EntryGroupIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntryGroupIamMemberState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["entryGroup"] = state ? state.entryGroup : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as EntryGroupIamMemberArgs | undefined;
            if ((!args || args.entryGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entryGroup'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["entryGroup"] = args ? args.entryGroup : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntryGroupIamMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EntryGroupIamMember resources.
 */
export interface EntryGroupIamMemberState {
    condition?: pulumi.Input<inputs.datacatalog.EntryGroupIamMemberCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    entryGroup?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    member?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EntryGroupIamMember resource.
 */
export interface EntryGroupIamMemberArgs {
    condition?: pulumi.Input<inputs.datacatalog.EntryGroupIamMemberCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    entryGroup: pulumi.Input<string>;
    member: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.datacatalog.EntryGroupIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
