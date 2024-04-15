// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Compute Engine Snapshot. Each of these resources serves a different use case:
 *
 * * `gcp.compute.SnapshotIamPolicy`: Authoritative. Sets the IAM policy for the snapshot and replaces any existing policy already attached.
 * * `gcp.compute.SnapshotIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the snapshot are preserved.
 * * `gcp.compute.SnapshotIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the snapshot are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.compute.SnapshotIamPolicy`: Retrieves the IAM policy for the snapshot
 *
 * > **Note:** `gcp.compute.SnapshotIamPolicy` **cannot** be used in conjunction with `gcp.compute.SnapshotIamBinding` and `gcp.compute.SnapshotIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.compute.SnapshotIamBinding` resources **can be** used in conjunction with `gcp.compute.SnapshotIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_compute\_snapshot\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
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
 * const policy = new gcp.compute.SnapshotIamPolicy("policy", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_compute\_snapshot\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.compute.SnapshotIamBinding("binding", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_compute\_snapshot\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.compute.SnapshotIamMember("member", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_compute\_snapshot\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
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
 * const policy = new gcp.compute.SnapshotIamPolicy("policy", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_compute\_snapshot\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.compute.SnapshotIamBinding("binding", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_compute\_snapshot\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.compute.SnapshotIamMember("member", {
 *     project: snapshot.project,
 *     name: snapshot.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms:
 *
 * * projects/{{project}}/global/snapshots/{{name}}
 *
 * * {{project}}/{{name}}
 *
 * * {{name}}
 *
 * Any variables not passed in the import command will be taken from the provider configuration.
 *
 * Compute Engine snapshot IAM resources can be imported using the resource identifiers, role, and member.
 *
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 * $ pulumi import gcp:compute/snapshotIamPolicy:SnapshotIamPolicy editor "projects/{{project}}/global/snapshots/{{snapshot}} roles/viewer user:jane@example.com"
 * ```
 *
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 *
 * ```sh
 * $ pulumi import gcp:compute/snapshotIamPolicy:SnapshotIamPolicy editor "projects/{{project}}/global/snapshots/{{snapshot}} roles/viewer"
 * ```
 *
 * IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 * $ pulumi import gcp:compute/snapshotIamPolicy:SnapshotIamPolicy editor projects/{{project}}/global/snapshots/{{snapshot}}
 * ```
 *
 * -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
 *
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class SnapshotIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotIamPolicyState, opts?: pulumi.CustomResourceOptions): SnapshotIamPolicy {
        return new SnapshotIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/snapshotIamPolicy:SnapshotIamPolicy';

    /**
     * Returns true if the given object is an instance of SnapshotIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotIamPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
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
     * Create a SnapshotIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotIamPolicyArgs | SnapshotIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotIamPolicyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as SnapshotIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnapshotIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotIamPolicy resources.
 */
export interface SnapshotIamPolicyState {
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
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
}

/**
 * The set of arguments for constructing a SnapshotIamPolicy resource.
 */
export interface SnapshotIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
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
}
