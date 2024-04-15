// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Secret Manager Secret. Each of these resources serves a different use case:
 *
 * * `gcp.secretmanager.SecretIamPolicy`: Authoritative. Sets the IAM policy for the secret and replaces any existing policy already attached.
 * * `gcp.secretmanager.SecretIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the secret are preserved.
 * * `gcp.secretmanager.SecretIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the secret are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.secretmanager.SecretIamPolicy`: Retrieves the IAM policy for the secret
 *
 * > **Note:** `gcp.secretmanager.SecretIamPolicy` **cannot** be used in conjunction with `gcp.secretmanager.SecretIamBinding` and `gcp.secretmanager.SecretIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.secretmanager.SecretIamBinding` resources **can be** used in conjunction with `gcp.secretmanager.SecretIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_secret\_manager\_secret\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/secretmanager.secretAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.secretmanager.SecretIamPolicy("policy", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_secret\_manager\_secret\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.secretmanager.SecretIamBinding("binding", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     role: "roles/secretmanager.secretAccessor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_secret\_manager\_secret\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.secretmanager.SecretIamMember("member", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     role: "roles/secretmanager.secretAccessor",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_secret\_manager\_secret\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/secretmanager.secretAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.secretmanager.SecretIamPolicy("policy", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_secret\_manager\_secret\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.secretmanager.SecretIamBinding("binding", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     role: "roles/secretmanager.secretAccessor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_secret\_manager\_secret\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.secretmanager.SecretIamMember("member", {
 *     project: secret_basic.project,
 *     secretId: secret_basic.secretId,
 *     role: "roles/secretmanager.secretAccessor",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms:
 *
 * * projects/{{project}}/secrets/{{secret_id}}
 *
 * * {{project}}/{{secret_id}}
 *
 * * {{secret_id}}
 *
 * Any variables not passed in the import command will be taken from the provider configuration.
 *
 * Secret Manager secret IAM resources can be imported using the resource identifiers, role, and member.
 *
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 * $ pulumi import gcp:secretmanager/secretIamPolicy:SecretIamPolicy editor "projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor user:jane@example.com"
 * ```
 *
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 *
 * ```sh
 * $ pulumi import gcp:secretmanager/secretIamPolicy:SecretIamPolicy editor "projects/{{project}}/secrets/{{secret_id}} roles/secretmanager.secretAccessor"
 * ```
 *
 * IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 * $ pulumi import gcp:secretmanager/secretIamPolicy:SecretIamPolicy editor projects/{{project}}/secrets/{{secret_id}}
 * ```
 *
 * -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
 *
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class SecretIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SecretIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretIamPolicyState, opts?: pulumi.CustomResourceOptions): SecretIamPolicy {
        return new SecretIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:secretmanager/secretIamPolicy:SecretIamPolicy';

    /**
     * Returns true if the given object is an instance of SecretIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretIamPolicy.__pulumiType;
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
    public readonly secretId!: pulumi.Output<string>;

    /**
     * Create a SecretIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretIamPolicyArgs | SecretIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretIamPolicyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["secretId"] = state ? state.secretId : undefined;
        } else {
            const args = argsOrState as SecretIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.secretId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretId'");
            }
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["secretId"] = args ? args.secretId : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretIamPolicy resources.
 */
export interface SecretIamPolicyState {
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
    secretId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretIamPolicy resource.
 */
export interface SecretIamPolicyArgs {
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
    secretId: pulumi.Input<string>;
}
