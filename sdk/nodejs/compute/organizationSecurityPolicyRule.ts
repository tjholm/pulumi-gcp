// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * OrganizationSecurityPolicyRule can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule default {{policy_id}}/priority/{{priority}}
 * ```
 */
export class OrganizationSecurityPolicyRule extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSecurityPolicyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSecurityPolicyRuleState, opts?: pulumi.CustomResourceOptions): OrganizationSecurityPolicyRule {
        return new OrganizationSecurityPolicyRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule';

    /**
     * Returns true if the given object is an instance of OrganizationSecurityPolicyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSecurityPolicyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSecurityPolicyRule.__pulumiType;
    }

    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either
     * "allow", "deny" or "gotoNext".
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * A description of the rule.
     * (Optional)
     * A description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The direction in which this rule applies. If unspecified an INGRESS rule is created.
     * Possible values are `INGRESS` and `EGRESS`.
     */
    public readonly direction!: pulumi.Output<string | undefined>;
    /**
     * Denotes whether to enable logging for a particular rule.
     * If logging is enabled, logs will be exported to the
     * configured export destination in Stackdriver.
     */
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     * Structure is documented below.
     */
    public readonly match!: pulumi.Output<outputs.compute.OrganizationSecurityPolicyRuleMatch>;
    /**
     * The ID of the OrganizationSecurityPolicy this rule applies to.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * If set to true, the specified action is not enforced.
     */
    public readonly preview!: pulumi.Output<boolean | undefined>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a value
     * between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
     * highest priority and 2147483647 is the lowest prority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * A list of network resource URLs to which this rule applies.
     * This field allows you to control which network's VMs get
     * this rule. If this field is left blank, all VMs
     * within the organization will receive the rule.
     */
    public readonly targetResources!: pulumi.Output<string[] | undefined>;
    /**
     * A list of service accounts indicating the sets of
     * instances that are applied with this rule.
     */
    public readonly targetServiceAccounts!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OrganizationSecurityPolicyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSecurityPolicyRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSecurityPolicyRuleArgs | OrganizationSecurityPolicyRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationSecurityPolicyRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["enableLogging"] = state ? state.enableLogging : undefined;
            resourceInputs["match"] = state ? state.match : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["preview"] = state ? state.preview : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["targetResources"] = state ? state.targetResources : undefined;
            resourceInputs["targetServiceAccounts"] = state ? state.targetServiceAccounts : undefined;
        } else {
            const args = argsOrState as OrganizationSecurityPolicyRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.match === undefined) && !opts.urn) {
                throw new Error("Missing required property 'match'");
            }
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["enableLogging"] = args ? args.enableLogging : undefined;
            resourceInputs["match"] = args ? args.match : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["preview"] = args ? args.preview : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["targetResources"] = args ? args.targetResources : undefined;
            resourceInputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationSecurityPolicyRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSecurityPolicyRule resources.
 */
export interface OrganizationSecurityPolicyRuleState {
    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either
     * "allow", "deny" or "gotoNext".
     */
    action?: pulumi.Input<string>;
    /**
     * A description of the rule.
     * (Optional)
     * A description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction in which this rule applies. If unspecified an INGRESS rule is created.
     * Possible values are `INGRESS` and `EGRESS`.
     */
    direction?: pulumi.Input<string>;
    /**
     * Denotes whether to enable logging for a particular rule.
     * If logging is enabled, logs will be exported to the
     * configured export destination in Stackdriver.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     * Structure is documented below.
     */
    match?: pulumi.Input<inputs.compute.OrganizationSecurityPolicyRuleMatch>;
    /**
     * The ID of the OrganizationSecurityPolicy this rule applies to.
     */
    policyId?: pulumi.Input<string>;
    /**
     * If set to true, the specified action is not enforced.
     */
    preview?: pulumi.Input<boolean>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a value
     * between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
     * highest priority and 2147483647 is the lowest prority.
     */
    priority?: pulumi.Input<number>;
    /**
     * A list of network resource URLs to which this rule applies.
     * This field allows you to control which network's VMs get
     * this rule. If this field is left blank, all VMs
     * within the organization will receive the rule.
     */
    targetResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating the sets of
     * instances that are applied with this rule.
     */
    targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OrganizationSecurityPolicyRule resource.
 */
export interface OrganizationSecurityPolicyRuleArgs {
    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either
     * "allow", "deny" or "gotoNext".
     */
    action: pulumi.Input<string>;
    /**
     * A description of the rule.
     * (Optional)
     * A description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction in which this rule applies. If unspecified an INGRESS rule is created.
     * Possible values are `INGRESS` and `EGRESS`.
     */
    direction?: pulumi.Input<string>;
    /**
     * Denotes whether to enable logging for a particular rule.
     * If logging is enabled, logs will be exported to the
     * configured export destination in Stackdriver.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     * Structure is documented below.
     */
    match: pulumi.Input<inputs.compute.OrganizationSecurityPolicyRuleMatch>;
    /**
     * The ID of the OrganizationSecurityPolicy this rule applies to.
     */
    policyId: pulumi.Input<string>;
    /**
     * If set to true, the specified action is not enforced.
     */
    preview?: pulumi.Input<boolean>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a value
     * between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
     * highest priority and 2147483647 is the lowest prority.
     */
    priority: pulumi.Input<number>;
    /**
     * A list of network resource URLs to which this rule applies.
     * This field allows you to control which network's VMs get
     * this rule. If this field is left blank, all VMs
     * within the organization will receive the rule.
     */
    targetResources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts indicating the sets of
     * instances that are applied with this rule.
     */
    targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
}
