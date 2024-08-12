// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Pub/Sub Schema. Each of these resources serves a different use case:
 *
 * * `gcp.pubsub.SchemaIamPolicy`: Authoritative. Sets the IAM policy for the schema and replaces any existing policy already attached.
 * * `gcp.pubsub.SchemaIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the schema are preserved.
 * * `gcp.pubsub.SchemaIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the schema are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.pubsub.SchemaIamPolicy`: Retrieves the IAM policy for the schema
 *
 * > **Note:** `gcp.pubsub.SchemaIamPolicy` **cannot** be used in conjunction with `gcp.pubsub.SchemaIamBinding` and `gcp.pubsub.SchemaIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.pubsub.SchemaIamBinding` resources **can be** used in conjunction with `gcp.pubsub.SchemaIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## gcp.pubsub.SchemaIamPolicy
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
 * const policy = new gcp.pubsub.SchemaIamPolicy("policy", {
 *     project: example.project,
 *     schema: example.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## gcp.pubsub.SchemaIamBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.pubsub.SchemaIamBinding("binding", {
 *     project: example.project,
 *     schema: example.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.pubsub.SchemaIamMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.pubsub.SchemaIamMember("member", {
 *     project: example.project,
 *     schema: example.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## This resource supports User Project Overrides.
 *
 * - 
 *
 * # IAM policy for Cloud Pub/Sub Schema
 * Three different resources help you manage your IAM policy for Cloud Pub/Sub Schema. Each of these resources serves a different use case:
 *
 * * `gcp.pubsub.SchemaIamPolicy`: Authoritative. Sets the IAM policy for the schema and replaces any existing policy already attached.
 * * `gcp.pubsub.SchemaIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the schema are preserved.
 * * `gcp.pubsub.SchemaIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the schema are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.pubsub.SchemaIamPolicy`: Retrieves the IAM policy for the schema
 *
 * > **Note:** `gcp.pubsub.SchemaIamPolicy` **cannot** be used in conjunction with `gcp.pubsub.SchemaIamBinding` and `gcp.pubsub.SchemaIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.pubsub.SchemaIamBinding` resources **can be** used in conjunction with `gcp.pubsub.SchemaIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## gcp.pubsub.SchemaIamPolicy
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
 * const policy = new gcp.pubsub.SchemaIamPolicy("policy", {
 *     project: example.project,
 *     schema: example.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## gcp.pubsub.SchemaIamBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.pubsub.SchemaIamBinding("binding", {
 *     project: example.project,
 *     schema: example.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.pubsub.SchemaIamMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.pubsub.SchemaIamMember("member", {
 *     project: example.project,
 *     schema: example.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms:
 *
 * * projects/{{project}}/schemas/{{name}}
 *
 * * {{project}}/{{name}}
 *
 * * {{name}}
 *
 * Any variables not passed in the import command will be taken from the provider configuration.
 *
 * Cloud Pub/Sub schema IAM resources can be imported using the resource identifiers, role, and member.
 *
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 * $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor "projects/{{project}}/schemas/{{schema}} roles/viewer user:jane@example.com"
 * ```
 *
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 *
 * ```sh
 * $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor "projects/{{project}}/schemas/{{schema}} roles/viewer"
 * ```
 *
 * IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 * $ pulumi import gcp:pubsub/schemaIamPolicy:SchemaIamPolicy editor projects/{{project}}/schemas/{{schema}}
 * ```
 *
 * -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
 *
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class SchemaIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SchemaIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SchemaIamPolicyState, opts?: pulumi.CustomResourceOptions): SchemaIamPolicy {
        return new SchemaIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:pubsub/schemaIamPolicy:SchemaIamPolicy';

    /**
     * Returns true if the given object is an instance of SchemaIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SchemaIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SchemaIamPolicy.__pulumiType;
    }

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
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly schema!: pulumi.Output<string>;

    /**
     * Create a SchemaIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SchemaIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SchemaIamPolicyArgs | SchemaIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SchemaIamPolicyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
        } else {
            const args = argsOrState as SchemaIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.schema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schema'");
            }
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SchemaIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SchemaIamPolicy resources.
 */
export interface SchemaIamPolicyState {
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
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    schema?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SchemaIamPolicy resource.
 */
export interface SchemaIamPolicyArgs {
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
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    schema: pulumi.Input<string>;
}
