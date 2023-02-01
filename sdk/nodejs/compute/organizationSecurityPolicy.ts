// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * OrganizationSecurityPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default locations/global/securityPolicies/{{policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default {{policy_id}}
 * ```
 */
export class OrganizationSecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSecurityPolicyState, opts?: pulumi.CustomResourceOptions): OrganizationSecurityPolicy {
        return new OrganizationSecurityPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy';

    /**
     * Returns true if the given object is an instance of OrganizationSecurityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSecurityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSecurityPolicy.__pulumiType;
    }

    /**
     * A textual description for the organization security policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A textual name of the security policy.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. This field is used internally during
     * updates of this resource.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id} or folders/{folder_id}
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    public /*out*/ readonly policyId!: pulumi.Output<string>;
    /**
     * The type indicates the intended use of the security policy.
     * For organization security policies, the only supported type
     * is "FIREWALL".
     * Default value is `FIREWALL`.
     * Possible values are `FIREWALL`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a OrganizationSecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSecurityPolicyArgs | OrganizationSecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationSecurityPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as OrganizationSecurityPolicyArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["policyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationSecurityPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSecurityPolicy resources.
 */
export interface OrganizationSecurityPolicyState {
    /**
     * A textual description for the organization security policy.
     */
    description?: pulumi.Input<string>;
    /**
     * A textual name of the security policy.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource. This field is used internally during
     * updates of this resource.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id} or folders/{folder_id}
     */
    parent?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     */
    policyId?: pulumi.Input<string>;
    /**
     * The type indicates the intended use of the security policy.
     * For organization security policies, the only supported type
     * is "FIREWALL".
     * Default value is `FIREWALL`.
     * Possible values are `FIREWALL`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationSecurityPolicy resource.
 */
export interface OrganizationSecurityPolicyArgs {
    /**
     * A textual description for the organization security policy.
     */
    description?: pulumi.Input<string>;
    /**
     * A textual name of the security policy.
     */
    displayName: pulumi.Input<string>;
    /**
     * The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id} or folders/{folder_id}
     */
    parent: pulumi.Input<string>;
    /**
     * The type indicates the intended use of the security policy.
     * For organization security policies, the only supported type
     * is "FIREWALL".
     * Default value is `FIREWALL`.
     * Possible values are `FIREWALL`.
     */
    type?: pulumi.Input<string>;
}
