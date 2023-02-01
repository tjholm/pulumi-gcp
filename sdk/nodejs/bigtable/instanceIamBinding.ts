// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
 *
 * * `gcp.bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `gcp.bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `gcp.bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 *
 * > **Note:** `gcp.bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.InstanceIamBinding` and `gcp.bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `gcp.bigtable.InstanceIamPolicy` replaces the entire policy.
 *
 * > **Note:** `gcp.bigtable.InstanceIamBinding` resources **can be** used in conjunction with `gcp.bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_bigtable\_instance\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/bigtable.user",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const editor = new gcp.bigtable.InstanceIamPolicy("editor", {
 *     project: "your-project",
 *     instance: "your-bigtable-instance",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_bigtable\_instance\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.bigtable.InstanceIamBinding("editor", {
 *     instance: "your-bigtable-instance",
 *     members: ["user:jane@example.com"],
 *     role: "roles/bigtable.user",
 * });
 * ```
 *
 * ## google\_bigtable\_instance\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const editor = new gcp.bigtable.InstanceIamMember("editor", {
 *     instance: "your-bigtable-instance",
 *     member: "user:jane@example.com",
 *     role: "roles/bigtable.user",
 * });
 * ```
 *
 * ## Import
 *
 * Instance IAM resources can be imported using the project, instance name, role and/or member.
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance}"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance} roles/editor"
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/instanceIamBinding:InstanceIamBinding editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class InstanceIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceIamBindingState, opts?: pulumi.CustomResourceOptions): InstanceIamBinding {
        return new InstanceIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/instanceIamBinding:InstanceIamBinding';

    /**
     * Returns true if the given object is an instance of InstanceIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.bigtable.InstanceIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the instances's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    public readonly instance!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a InstanceIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceIamBindingArgs | InstanceIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceIamBindingState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["instance"] = state ? state.instance : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as InstanceIamBindingArgs | undefined;
            if ((!args || args.instance === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instance'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["instance"] = args ? args.instance : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceIamBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceIamBinding resources.
 */
export interface InstanceIamBindingState {
    condition?: pulumi.Input<inputs.bigtable.InstanceIamBindingCondition>;
    /**
     * (Computed) The etag of the instances's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    instance?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceIamBinding resource.
 */
export interface InstanceIamBindingArgs {
    condition?: pulumi.Input<inputs.bigtable.InstanceIamBindingCondition>;
    /**
     * The name or relative resource id of the instance to manage IAM policies for.
     */
    instance: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     */
    role: pulumi.Input<string>;
}
