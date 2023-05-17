// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Compute NetworkFirewallPolicyRule resource
 *
 * ## Example Usage
 * ### Global
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicGlobalNetworksecurityAddressGroup = new gcp.networksecurity.AddressGroup("basicGlobalNetworksecurityAddressGroup", {
 *     parent: "projects/my-project-name",
 *     description: "Sample global networksecurity_address_group",
 *     location: "global",
 *     items: ["208.80.154.224/32"],
 *     type: "IPV4",
 *     capacity: 100,
 * });
 * const basicNetworkFirewallPolicy = new gcp.compute.NetworkFirewallPolicy("basicNetworkFirewallPolicy", {
 *     description: "Sample global network firewall policy",
 *     project: "my-project-name",
 * });
 * const basicNetwork = new gcp.compute.Network("basicNetwork", {});
 * const basicKey = new gcp.tags.TagKey("basicKey", {
 *     description: "For keyname resources.",
 *     parent: "organizations/123456789",
 *     purpose: "GCE_FIREWALL",
 *     shortName: "tagkey",
 *     purposeData: {
 *         network: pulumi.interpolate`my-project-name/${basicNetwork.name}`,
 *     },
 * });
 * const basicValue = new gcp.tags.TagValue("basicValue", {
 *     description: "For valuename resources.",
 *     parent: pulumi.interpolate`tagKeys/${basicKey.name}`,
 *     shortName: "tagvalue",
 * });
 * const primary = new gcp.compute.NetworkFirewallPolicyRule("primary", {
 *     action: "allow",
 *     description: "This is a simple rule description",
 *     direction: "INGRESS",
 *     disabled: false,
 *     enableLogging: true,
 *     firewallPolicy: basicNetworkFirewallPolicy.name,
 *     priority: 1000,
 *     ruleName: "test-rule",
 *     targetServiceAccounts: ["my@service-account.com"],
 *     match: {
 *         srcIpRanges: ["10.100.0.1/32"],
 *         srcFqdns: ["google.com"],
 *         srcRegionCodes: ["US"],
 *         srcThreatIntelligences: ["iplist-known-malicious-ips"],
 *         srcSecureTags: [{
 *             name: pulumi.interpolate`tagValues/${basicValue.name}`,
 *         }],
 *         layer4Configs: [{
 *             ipProtocol: "all",
 *         }],
 *         srcAddressGroups: [basicGlobalNetworksecurityAddressGroup.id],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * NetworkFirewallPolicyRule can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkFirewallPolicyRule:NetworkFirewallPolicyRule default projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkFirewallPolicyRule:NetworkFirewallPolicyRule default {{project}}/{{firewall_policy}}/{{priority}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/networkFirewallPolicyRule:NetworkFirewallPolicyRule default {{firewall_policy}}/{{priority}}
 * ```
 */
export class NetworkFirewallPolicyRule extends pulumi.CustomResource {
    /**
     * Get an existing NetworkFirewallPolicyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkFirewallPolicyRuleState, opts?: pulumi.CustomResourceOptions): NetworkFirewallPolicyRule {
        return new NetworkFirewallPolicyRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkFirewallPolicyRule:NetworkFirewallPolicyRule';

    /**
     * Returns true if the given object is an instance of NetworkFirewallPolicyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkFirewallPolicyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkFirewallPolicyRule.__pulumiType;
    }

    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * An optional description for this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The direction in which this rule applies. Possible values: INGRESS, EGRESS
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
     */
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    /**
     * The firewall policy of the resource.
     */
    public readonly firewallPolicy!: pulumi.Output<string>;
    /**
     * Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     */
    public readonly match!: pulumi.Output<outputs.compute.NetworkFirewallPolicyRuleMatch>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The project for the resource
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * An optional name for the rule. This field is not a unique identifier and can be updated.
     */
    public readonly ruleName!: pulumi.Output<string | undefined>;
    /**
     * Calculation of the complexity of a single firewall policy rule.
     */
    public /*out*/ readonly ruleTupleCount!: pulumi.Output<number>;
    /**
     * A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
     */
    public readonly targetSecureTags!: pulumi.Output<outputs.compute.NetworkFirewallPolicyRuleTargetSecureTag[] | undefined>;
    /**
     * A list of service accounts indicating the sets of instances that are applied with this rule.
     */
    public readonly targetServiceAccounts!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NetworkFirewallPolicyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkFirewallPolicyRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkFirewallPolicyRuleArgs | NetworkFirewallPolicyRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkFirewallPolicyRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["enableLogging"] = state ? state.enableLogging : undefined;
            resourceInputs["firewallPolicy"] = state ? state.firewallPolicy : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["match"] = state ? state.match : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["ruleTupleCount"] = state ? state.ruleTupleCount : undefined;
            resourceInputs["targetSecureTags"] = state ? state.targetSecureTags : undefined;
            resourceInputs["targetServiceAccounts"] = state ? state.targetServiceAccounts : undefined;
        } else {
            const args = argsOrState as NetworkFirewallPolicyRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.firewallPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewallPolicy'");
            }
            if ((!args || args.match === undefined) && !opts.urn) {
                throw new Error("Missing required property 'match'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["enableLogging"] = args ? args.enableLogging : undefined;
            resourceInputs["firewallPolicy"] = args ? args.firewallPolicy : undefined;
            resourceInputs["match"] = args ? args.match : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["targetSecureTags"] = args ? args.targetSecureTags : undefined;
            resourceInputs["targetServiceAccounts"] = args ? args.targetServiceAccounts : undefined;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["ruleTupleCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkFirewallPolicyRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkFirewallPolicyRule resources.
 */
export interface NetworkFirewallPolicyRuleState {
    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
     */
    action?: pulumi.Input<string>;
    /**
     * An optional description for this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction in which this rule applies. Possible values: INGRESS, EGRESS
     */
    direction?: pulumi.Input<string>;
    /**
     * Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * The firewall policy of the resource.
     */
    firewallPolicy?: pulumi.Input<string>;
    /**
     * Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     */
    kind?: pulumi.Input<string>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     */
    match?: pulumi.Input<inputs.compute.NetworkFirewallPolicyRuleMatch>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * An optional name for the rule. This field is not a unique identifier and can be updated.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Calculation of the complexity of a single firewall policy rule.
     */
    ruleTupleCount?: pulumi.Input<number>;
    /**
     * A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
     */
    targetSecureTags?: pulumi.Input<pulumi.Input<inputs.compute.NetworkFirewallPolicyRuleTargetSecureTag>[]>;
    /**
     * A list of service accounts indicating the sets of instances that are applied with this rule.
     */
    targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NetworkFirewallPolicyRule resource.
 */
export interface NetworkFirewallPolicyRuleArgs {
    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
     */
    action: pulumi.Input<string>;
    /**
     * An optional description for this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction in which this rule applies. Possible values: INGRESS, EGRESS
     */
    direction: pulumi.Input<string>;
    /**
     * Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * The firewall policy of the resource.
     */
    firewallPolicy: pulumi.Input<string>;
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
     */
    match: pulumi.Input<inputs.compute.NetworkFirewallPolicyRuleMatch>;
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     */
    priority: pulumi.Input<number>;
    /**
     * The project for the resource
     */
    project?: pulumi.Input<string>;
    /**
     * An optional name for the rule. This field is not a unique identifier and can be updated.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
     */
    targetSecureTags?: pulumi.Input<pulumi.Input<inputs.compute.NetworkFirewallPolicyRuleTargetSecureTag>[]>;
    /**
     * A list of service accounts indicating the sets of instances that are applied with this rule.
     */
    targetServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
}
