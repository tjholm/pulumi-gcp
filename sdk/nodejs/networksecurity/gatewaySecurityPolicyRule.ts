// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The GatewaySecurityPolicyRule resource is in a nested collection within a GatewaySecurityPolicy and represents
 * a traffic matching condition and associated action to perform.
 *
 * To get more information about GatewaySecurityPolicyRule, see:
 *
 * * [API documentation](https://cloud.google.com/secure-web-proxy/docs/reference/network-security/rest/v1/projects.locations.gatewaySecurityPolicies.rules)
 *
 * ## Example Usage
 * ### Network Security Gateway Security Policy Rules Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultGatewaySecurityPolicy = new gcp.networksecurity.GatewaySecurityPolicy("defaultGatewaySecurityPolicy", {
 *     location: "us-central1",
 *     description: "gateway security policy created to be used as reference by the rule.",
 * });
 * const defaultGatewaySecurityPolicyRule = new gcp.networksecurity.GatewaySecurityPolicyRule("defaultGatewaySecurityPolicyRule", {
 *     location: "us-central1",
 *     gatewaySecurityPolicy: defaultGatewaySecurityPolicy.name,
 *     enabled: true,
 *     description: "my description",
 *     priority: 0,
 *     sessionMatcher: "host() == 'example.com'",
 *     basicProfile: "ALLOW",
 * });
 * ```
 * ### Network Security Gateway Security Policy Rules Advanced
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultGatewaySecurityPolicy = new gcp.networksecurity.GatewaySecurityPolicy("defaultGatewaySecurityPolicy", {
 *     location: "us-central1",
 *     description: "gateway security policy created to be used as reference by the rule.",
 * });
 * const defaultGatewaySecurityPolicyRule = new gcp.networksecurity.GatewaySecurityPolicyRule("defaultGatewaySecurityPolicyRule", {
 *     location: "us-central1",
 *     gatewaySecurityPolicy: defaultGatewaySecurityPolicy.name,
 *     enabled: true,
 *     description: "my description",
 *     priority: 0,
 *     sessionMatcher: "host() == 'example.com'",
 *     applicationMatcher: "request.method == 'POST'",
 *     tlsInspectionEnabled: false,
 *     basicProfile: "ALLOW",
 * });
 * ```
 *
 * ## Import
 *
 * GatewaySecurityPolicyRule can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}` * `{{project}}/{{location}}/{{gateway_security_policy}}/{{name}}` * `{{location}}/{{gateway_security_policy}}/{{name}}` When using the `pulumi import` command, GatewaySecurityPolicyRule can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/gatewaySecurityPolicyRule:GatewaySecurityPolicyRule default projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{gateway_security_policy}}/rules/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/gatewaySecurityPolicyRule:GatewaySecurityPolicyRule default {{project}}/{{location}}/{{gateway_security_policy}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/gatewaySecurityPolicyRule:GatewaySecurityPolicyRule default {{location}}/{{gateway_security_policy}}/{{name}}
 * ```
 */
export class GatewaySecurityPolicyRule extends pulumi.CustomResource {
    /**
     * Get an existing GatewaySecurityPolicyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewaySecurityPolicyRuleState, opts?: pulumi.CustomResourceOptions): GatewaySecurityPolicyRule {
        return new GatewaySecurityPolicyRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networksecurity/gatewaySecurityPolicyRule:GatewaySecurityPolicyRule';

    /**
     * Returns true if the given object is an instance of GatewaySecurityPolicyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewaySecurityPolicyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewaySecurityPolicyRule.__pulumiType;
    }

    /**
     * CEL expression for matching on L7/application level criteria.
     */
    public readonly applicationMatcher!: pulumi.Output<string | undefined>;
    /**
     * Profile which tells what the primitive action should be. Possible values are: * ALLOW * DENY.
     * Possible values are: `BASIC_PROFILE_UNSPECIFIED`, `ALLOW`, `DENY`.
     */
    public readonly basicProfile!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Free-text description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the rule is enforced.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The name of the gatewat security policy this rule belongs to.
     *
     *
     * - - -
     */
    public readonly gatewaySecurityPolicy!: pulumi.Output<string>;
    /**
     * The location of the gateway security policy.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}
     * rule should match the pattern: (^a-z?$).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Priority of the rule. Lower number corresponds to higher precedence.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Server-defined URL of this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * CEL expression for matching on session criteria.
     */
    public readonly sessionMatcher!: pulumi.Output<string>;
    /**
     * Flag to enable TLS inspection of traffic matching on. Can only be true if the
     * parent GatewaySecurityPolicy references a TLSInspectionConfig.
     */
    public readonly tlsInspectionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The timestamp when the resource was updated.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a GatewaySecurityPolicyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewaySecurityPolicyRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewaySecurityPolicyRuleArgs | GatewaySecurityPolicyRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewaySecurityPolicyRuleState | undefined;
            resourceInputs["applicationMatcher"] = state ? state.applicationMatcher : undefined;
            resourceInputs["basicProfile"] = state ? state.basicProfile : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["gatewaySecurityPolicy"] = state ? state.gatewaySecurityPolicy : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["sessionMatcher"] = state ? state.sessionMatcher : undefined;
            resourceInputs["tlsInspectionEnabled"] = state ? state.tlsInspectionEnabled : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as GatewaySecurityPolicyRuleArgs | undefined;
            if ((!args || args.basicProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'basicProfile'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.gatewaySecurityPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewaySecurityPolicy'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.sessionMatcher === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionMatcher'");
            }
            resourceInputs["applicationMatcher"] = args ? args.applicationMatcher : undefined;
            resourceInputs["basicProfile"] = args ? args.basicProfile : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["gatewaySecurityPolicy"] = args ? args.gatewaySecurityPolicy : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sessionMatcher"] = args ? args.sessionMatcher : undefined;
            resourceInputs["tlsInspectionEnabled"] = args ? args.tlsInspectionEnabled : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewaySecurityPolicyRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewaySecurityPolicyRule resources.
 */
export interface GatewaySecurityPolicyRuleState {
    /**
     * CEL expression for matching on L7/application level criteria.
     */
    applicationMatcher?: pulumi.Input<string>;
    /**
     * Profile which tells what the primitive action should be. Possible values are: * ALLOW * DENY.
     * Possible values are: `BASIC_PROFILE_UNSPECIFIED`, `ALLOW`, `DENY`.
     */
    basicProfile?: pulumi.Input<string>;
    /**
     * The timestamp when the resource was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
     */
    createTime?: pulumi.Input<string>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the rule is enforced.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of the gatewat security policy this rule belongs to.
     *
     *
     * - - -
     */
    gatewaySecurityPolicy?: pulumi.Input<string>;
    /**
     * The location of the gateway security policy.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}
     * rule should match the pattern: (^a-z?$).
     */
    name?: pulumi.Input<string>;
    /**
     * Priority of the rule. Lower number corresponds to higher precedence.
     */
    priority?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Server-defined URL of this resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * CEL expression for matching on session criteria.
     */
    sessionMatcher?: pulumi.Input<string>;
    /**
     * Flag to enable TLS inspection of traffic matching on. Can only be true if the
     * parent GatewaySecurityPolicy references a TLSInspectionConfig.
     */
    tlsInspectionEnabled?: pulumi.Input<boolean>;
    /**
     * The timestamp when the resource was updated.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewaySecurityPolicyRule resource.
 */
export interface GatewaySecurityPolicyRuleArgs {
    /**
     * CEL expression for matching on L7/application level criteria.
     */
    applicationMatcher?: pulumi.Input<string>;
    /**
     * Profile which tells what the primitive action should be. Possible values are: * ALLOW * DENY.
     * Possible values are: `BASIC_PROFILE_UNSPECIFIED`, `ALLOW`, `DENY`.
     */
    basicProfile: pulumi.Input<string>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the rule is enforced.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The name of the gatewat security policy this rule belongs to.
     *
     *
     * - - -
     */
    gatewaySecurityPolicy: pulumi.Input<string>;
    /**
     * The location of the gateway security policy.
     */
    location: pulumi.Input<string>;
    /**
     * Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}
     * rule should match the pattern: (^a-z?$).
     */
    name?: pulumi.Input<string>;
    /**
     * Priority of the rule. Lower number corresponds to higher precedence.
     */
    priority: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * CEL expression for matching on session criteria.
     */
    sessionMatcher: pulumi.Input<string>;
    /**
     * Flag to enable TLS inspection of traffic matching on. Can only be true if the
     * parent GatewaySecurityPolicy references a TLSInspectionConfig.
     */
    tlsInspectionEnabled?: pulumi.Input<boolean>;
}
