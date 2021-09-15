// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.
 *
 * For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultFirewallPolicy = new gcp.compute.FirewallPolicy("defaultFirewallPolicy", {
 *     parent: "organizations/12345",
 *     shortName: "my-policy",
 *     description: "Example Resource",
 * });
 * const defaultFirewallPolicyAssociation = new gcp.compute.FirewallPolicyAssociation("defaultFirewallPolicyAssociation", {
 *     firewallPolicy: defaultFirewallPolicy.id,
 *     attachmentTarget: google_folder.folder.name,
 * });
 * ```
 *
 * ## Import
 *
 * FirewallPolicyAssociation can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
 * ```
 */
export class FirewallPolicyAssociation extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicyAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallPolicyAssociationState, opts?: pulumi.CustomResourceOptions): FirewallPolicyAssociation {
        return new FirewallPolicyAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation';

    /**
     * Returns true if the given object is an instance of FirewallPolicyAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicyAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicyAssociation.__pulumiType;
    }

    /**
     * The target that the firewall policy is attached to.
     */
    public readonly attachmentTarget!: pulumi.Output<string>;
    /**
     * The firewall policy ID of the association.
     */
    public readonly firewallPolicy!: pulumi.Output<string>;
    /**
     * The name for an association.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The short name of the firewall policy of the association.
     */
    public /*out*/ readonly shortName!: pulumi.Output<string>;

    /**
     * Create a FirewallPolicyAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicyAssociationArgs | FirewallPolicyAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallPolicyAssociationState | undefined;
            inputs["attachmentTarget"] = state ? state.attachmentTarget : undefined;
            inputs["firewallPolicy"] = state ? state.firewallPolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["shortName"] = state ? state.shortName : undefined;
        } else {
            const args = argsOrState as FirewallPolicyAssociationArgs | undefined;
            if ((!args || args.attachmentTarget === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attachmentTarget'");
            }
            if ((!args || args.firewallPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewallPolicy'");
            }
            inputs["attachmentTarget"] = args ? args.attachmentTarget : undefined;
            inputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["shortName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallPolicyAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallPolicyAssociation resources.
 */
export interface FirewallPolicyAssociationState {
    /**
     * The target that the firewall policy is attached to.
     */
    attachmentTarget?: pulumi.Input<string>;
    /**
     * The firewall policy ID of the association.
     */
    firewallPolicy?: pulumi.Input<string>;
    /**
     * The name for an association.
     */
    name?: pulumi.Input<string>;
    /**
     * The short name of the firewall policy of the association.
     */
    shortName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallPolicyAssociation resource.
 */
export interface FirewallPolicyAssociationArgs {
    /**
     * The target that the firewall policy is attached to.
     */
    attachmentTarget: pulumi.Input<string>;
    /**
     * The firewall policy ID of the association.
     */
    firewallPolicy: pulumi.Input<string>;
    /**
     * The name for an association.
     */
    name?: pulumi.Input<string>;
}
