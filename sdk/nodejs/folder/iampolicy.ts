// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `folder`, role, and member e.g.
 *
 * ```sh
 *  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `folder` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `folder`.
 *
 * ```sh
 *  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
 * ```
 *
 *  IAM audit config imports use the identifier of the resource in question and the service, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
 *
 * ```sh
 *  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
 * ```
 */
export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:folder/iAMPolicy:IAMPolicy';

    /**
     * Returns true if the given object is an instance of IAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the folder's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    public readonly folder!: pulumi.Output<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the folder. The policy will be
     * merged with any existing policy applied to the folder.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a IAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMPolicyArgs | IAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAMPolicyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["folder"] = state ? state.folder : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if ((!args || args.folder === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folder'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["folder"] = args ? args.folder : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IAMPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    /**
     * (Computed) The etag of the folder's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    folder?: pulumi.Input<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the folder. The policy will be
     * merged with any existing policy applied to the folder.
     */
    policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    /**
     * The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
     */
    folder: pulumi.Input<string>;
    /**
     * The `gcp.organizations.getIAMPolicy` data source that represents
     * the IAM policy that will be applied to the folder. The policy will be
     * merged with any existing policy applied to the folder.
     */
    policyData: pulumi.Input<string>;
}
