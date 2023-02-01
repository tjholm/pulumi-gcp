// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `key_ring_id`, role, and account e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `key_ring_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/cloudkms.admin"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `key_ring_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam your-project-id/location-name/key-ring-name
 * ```
 */
export class KeyRingIAMBinding extends pulumi.CustomResource {
    /**
     * Get an existing KeyRingIAMBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyRingIAMBindingState, opts?: pulumi.CustomResourceOptions): KeyRingIAMBinding {
        return new KeyRingIAMBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/keyRingIAMBinding:KeyRingIAMBinding';

    /**
     * Returns true if the given object is an instance of KeyRingIAMBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyRingIAMBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyRingIAMBinding.__pulumiType;
    }

    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    public readonly condition!: pulumi.Output<outputs.kms.KeyRingIAMBindingCondition | undefined>;
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly keyRingId!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a KeyRingIAMBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyRingIAMBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyRingIAMBindingArgs | KeyRingIAMBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyRingIAMBindingState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["keyRingId"] = state ? state.keyRingId : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as KeyRingIAMBindingArgs | undefined;
            if ((!args || args.keyRingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyRingId'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["keyRingId"] = args ? args.keyRingId : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeyRingIAMBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyRingIAMBinding resources.
 */
export interface KeyRingIAMBindingState {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.kms.KeyRingIAMBindingCondition>;
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    keyRingId?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyRingIAMBinding resource.
 */
export interface KeyRingIAMBindingArgs {
    /**
     * ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     */
    condition?: pulumi.Input<inputs.kms.KeyRingIAMBindingCondition>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    keyRingId: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
