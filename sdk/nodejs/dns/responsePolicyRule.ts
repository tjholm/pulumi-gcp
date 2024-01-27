// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Response Policy Rule is a selector that applies its behavior to queries that match the selector.
 * Selectors are DNS names, which may be wildcards or exact matches.
 * Each DNS query subject to a Response Policy matches at most one ResponsePolicyRule,
 * as identified by the dnsName field with the longest matching suffix.
 *
 * ## Example Usage
 * ### Dns Response Policy Rule Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const network_1 = new gcp.compute.Network("network-1", {autoCreateSubnetworks: false});
 * const network_2 = new gcp.compute.Network("network-2", {autoCreateSubnetworks: false});
 * const response_policy = new gcp.dns.ResponsePolicy("response-policy", {
 *     responsePolicyName: "example-response-policy",
 *     networks: [
 *         {
 *             networkUrl: network_1.id,
 *         },
 *         {
 *             networkUrl: network_2.id,
 *         },
 *     ],
 * });
 * const example_response_policy_rule = new gcp.dns.ResponsePolicyRule("example-response-policy-rule", {
 *     responsePolicy: response_policy.responsePolicyName,
 *     ruleName: "example-rule",
 *     dnsName: "dns.example.com.",
 *     localData: {
 *         localDatas: [{
 *             name: "dns.example.com.",
 *             type: "A",
 *             ttl: 300,
 *             rrdatas: ["192.0.2.91"],
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ResponsePolicyRule can be imported using any of these accepted formats* `projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}` * `{{project}}/{{response_policy}}/{{rule_name}}` * `{{response_policy}}/{{rule_name}}` When using the `pulumi import` command, ResponsePolicyRule can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:dns/responsePolicyRule:ResponsePolicyRule default projects/{{project}}/responsePolicies/{{response_policy}}/rules/{{rule_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/responsePolicyRule:ResponsePolicyRule default {{project}}/{{response_policy}}/{{rule_name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dns/responsePolicyRule:ResponsePolicyRule default {{response_policy}}/{{rule_name}}
 * ```
 */
export class ResponsePolicyRule extends pulumi.CustomResource {
    /**
     * Get an existing ResponsePolicyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResponsePolicyRuleState, opts?: pulumi.CustomResourceOptions): ResponsePolicyRule {
        return new ResponsePolicyRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dns/responsePolicyRule:ResponsePolicyRule';

    /**
     * Returns true if the given object is an instance of ResponsePolicyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResponsePolicyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResponsePolicyRule.__pulumiType;
    }

    /**
     * (Optional, Beta)
     * Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'
     */
    public readonly behavior!: pulumi.Output<string | undefined>;
    /**
     * The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     */
    public readonly dnsName!: pulumi.Output<string>;
    /**
     * Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
     * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     * Structure is documented below.
     */
    public readonly localData!: pulumi.Output<outputs.dns.ResponsePolicyRuleLocalData | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Identifies the response policy addressed by this request.
     *
     *
     * - - -
     */
    public readonly responsePolicy!: pulumi.Output<string>;
    /**
     * An identifier for this rule. Must be unique with the ResponsePolicy.
     */
    public readonly ruleName!: pulumi.Output<string>;

    /**
     * Create a ResponsePolicyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponsePolicyRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResponsePolicyRuleArgs | ResponsePolicyRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResponsePolicyRuleState | undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["localData"] = state ? state.localData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["responsePolicy"] = state ? state.responsePolicy : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
        } else {
            const args = argsOrState as ResponsePolicyRuleArgs | undefined;
            if ((!args || args.dnsName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsName'");
            }
            if ((!args || args.responsePolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responsePolicy'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["dnsName"] = args ? args.dnsName : undefined;
            resourceInputs["localData"] = args ? args.localData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["responsePolicy"] = args ? args.responsePolicy : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResponsePolicyRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResponsePolicyRule resources.
 */
export interface ResponsePolicyRuleState {
    /**
     * (Optional, Beta)
     * Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'
     */
    behavior?: pulumi.Input<string>;
    /**
     * The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
     * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     * Structure is documented below.
     */
    localData?: pulumi.Input<inputs.dns.ResponsePolicyRuleLocalData>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Identifies the response policy addressed by this request.
     *
     *
     * - - -
     */
    responsePolicy?: pulumi.Input<string>;
    /**
     * An identifier for this rule. Must be unique with the ResponsePolicy.
     */
    ruleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResponsePolicyRule resource.
 */
export interface ResponsePolicyRuleArgs {
    /**
     * (Optional, Beta)
     * Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'
     */
    behavior?: pulumi.Input<string>;
    /**
     * The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
     */
    dnsName: pulumi.Input<string>;
    /**
     * Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
     * in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
     * Structure is documented below.
     */
    localData?: pulumi.Input<inputs.dns.ResponsePolicyRuleLocalData>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Identifies the response policy addressed by this request.
     *
     *
     * - - -
     */
    responsePolicy: pulumi.Input<string>;
    /**
     * An identifier for this rule. Must be unique with the ResponsePolicy.
     */
    ruleName: pulumi.Input<string>;
}
