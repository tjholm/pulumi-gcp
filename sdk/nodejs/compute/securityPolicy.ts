// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Security Policy defines an IP blacklist or whitelist that protects load balanced Google Cloud services by denying or permitting traffic from specified IP ranges. For more information
 * see the [official documentation](https://cloud.google.com/armor/docs/configure-security-policies)
 * and the [API](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies).
 *
 * Security Policy is used by google_compute_backend_service.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = new gcp.compute.SecurityPolicy("policy", {rules: [
 *     {
 *         action: "deny(403)",
 *         description: "Deny access to IPs in 9.9.9.0/24",
 *         match: {
 *             config: {
 *                 srcIpRanges: ["9.9.9.0/24"],
 *             },
 *             versionedExpr: "SRC_IPS_V1",
 *         },
 *         priority: 1000,
 *     },
 *     {
 *         action: "allow",
 *         description: "default rule",
 *         match: {
 *             config: {
 *                 srcIpRanges: ["*"],
 *             },
 *             versionedExpr: "SRC_IPS_V1",
 *         },
 *         priority: 2147483647,
 *     },
 * ]});
 * ```
 * ### With ReCAPTCHA Configuration Options
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.recaptcha.EnterpriseKey("primary", {
 *     displayName: "display-name",
 *     labels: {
 *         "label-one": "value-one",
 *     },
 *     project: "my-project-name",
 *     webSettings: {
 *         integrationType: "INVISIBLE",
 *         allowAllDomains: true,
 *         allowedDomains: ["localhost"],
 *     },
 * });
 * const policy = new gcp.compute.SecurityPolicy("policy", {
 *     description: "basic security policy",
 *     type: "CLOUD_ARMOR",
 *     recaptchaOptionsConfig: {
 *         redirectSiteKey: primary.name,
 *     },
 * });
 * ```
 * ### With Header Actions
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = new gcp.compute.SecurityPolicy("policy", {rules: [
 *     {
 *         action: "allow",
 *         description: "default rule",
 *         match: {
 *             config: {
 *                 srcIpRanges: ["*"],
 *             },
 *             versionedExpr: "SRC_IPS_V1",
 *         },
 *         priority: 2147483647,
 *     },
 *     {
 *         action: "allow",
 *         headerAction: {
 *             requestHeadersToAdds: [
 *                 {
 *                     headerName: "reCAPTCHA-Warning",
 *                     headerValue: "high",
 *                 },
 *                 {
 *                     headerName: "X-Resource",
 *                     headerValue: "test",
 *                 },
 *             ],
 *         },
 *         match: {
 *             expr: {
 *                 expression: "request.path.matches(\"/login.html\") && token.recaptcha_session.score < 0.2",
 *             },
 *         },
 *         priority: 1000,
 *     },
 * ]});
 * ```
 * ### With EnforceOnKey Value As Empty String
 * A scenario example that won't cause any conflict between `enforceOnKey` and `enforceOnKeyConfigs`, because `enforceOnKey` was specified as an empty string:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = new gcp.compute.SecurityPolicy("policy", {
 *     description: "throttle rule with enforce_on_key_configs",
 *     rules: [{
 *         action: "throttle",
 *         description: "default rule",
 *         match: {
 *             config: {
 *                 srcIpRanges: ["*"],
 *             },
 *             versionedExpr: "SRC_IPS_V1",
 *         },
 *         priority: 2147483647,
 *         rateLimitOptions: {
 *             conformAction: "allow",
 *             enforceOnKey: "",
 *             enforceOnKeyConfigs: [{
 *                 enforceOnKeyType: "IP",
 *             }],
 *             exceedAction: "redirect",
 *             exceedRedirectOptions: {
 *                 target: "<https://www.example.com>",
 *                 type: "EXTERNAL_302",
 *             },
 *             rateLimitThreshold: {
 *                 count: 10,
 *                 intervalSec: 60,
 *             },
 *         },
 *     }],
 * });
 * ```
 */
export class SecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityPolicyState, opts?: pulumi.CustomResourceOptions): SecurityPolicy {
        return new SecurityPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/securityPolicy:SecurityPolicy';

    /**
     * Returns true if the given object is an instance of SecurityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityPolicy.__pulumiType;
    }

    /**
     * Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
     */
    public readonly adaptiveProtectionConfig!: pulumi.Output<outputs.compute.SecurityPolicyAdaptiveProtectionConfig | undefined>;
    /**
     * [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
     * Structure is documented below.
     */
    public readonly advancedOptionsConfig!: pulumi.Output<outputs.compute.SecurityPolicyAdvancedOptionsConfig>;
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Fingerprint of this resource.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The name of the security policy.
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
     */
    public readonly recaptchaOptionsConfig!: pulumi.Output<outputs.compute.SecurityPolicyRecaptchaOptionsConfig | undefined>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.compute.SecurityPolicyRule[]>;
    /**
     * The URI of the created resourc
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The type indicates the intended use of the security policy. This field can be set only at resource creation time.
     * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
     * They filter requests before they hit the origin servers.
     * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
     * (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
     * They filter requests before the request is served from Google's cache.
     * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
     * managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPolicyArgs | SecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityPolicyState | undefined;
            resourceInputs["adaptiveProtectionConfig"] = state ? state.adaptiveProtectionConfig : undefined;
            resourceInputs["advancedOptionsConfig"] = state ? state.advancedOptionsConfig : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["recaptchaOptionsConfig"] = state ? state.recaptchaOptionsConfig : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecurityPolicyArgs | undefined;
            resourceInputs["adaptiveProtectionConfig"] = args ? args.adaptiveProtectionConfig : undefined;
            resourceInputs["advancedOptionsConfig"] = args ? args.advancedOptionsConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["recaptchaOptionsConfig"] = args ? args.recaptchaOptionsConfig : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityPolicy resources.
 */
export interface SecurityPolicyState {
    /**
     * Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
     */
    adaptiveProtectionConfig?: pulumi.Input<inputs.compute.SecurityPolicyAdaptiveProtectionConfig>;
    /**
     * [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
     * Structure is documented below.
     */
    advancedOptionsConfig?: pulumi.Input<inputs.compute.SecurityPolicyAdvancedOptionsConfig>;
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    description?: pulumi.Input<string>;
    /**
     * Fingerprint of this resource.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The name of the security policy.
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
     */
    recaptchaOptionsConfig?: pulumi.Input<inputs.compute.SecurityPolicyRecaptchaOptionsConfig>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.SecurityPolicyRule>[]>;
    /**
     * The URI of the created resourc
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The type indicates the intended use of the security policy. This field can be set only at resource creation time.
     * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
     * They filter requests before they hit the origin servers.
     * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
     * (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
     * They filter requests before the request is served from Google's cache.
     * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
     * managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityPolicy resource.
 */
export interface SecurityPolicyArgs {
    /**
     * Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
     */
    adaptiveProtectionConfig?: pulumi.Input<inputs.compute.SecurityPolicyAdaptiveProtectionConfig>;
    /**
     * [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
     * Structure is documented below.
     */
    advancedOptionsConfig?: pulumi.Input<inputs.compute.SecurityPolicyAdvancedOptionsConfig>;
    /**
     * An optional description of this security policy. Max size is 2048.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the security policy.
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
     */
    recaptchaOptionsConfig?: pulumi.Input<inputs.compute.SecurityPolicyRecaptchaOptionsConfig>;
    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
     * security policy, a default rule with action "allow" will be added. Structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.SecurityPolicyRule>[]>;
    /**
     * The type indicates the intended use of the security policy. This field can be set only at resource creation time.
     * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
     * They filter requests before they hit the origin servers.
     * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
     * (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
     * They filter requests before the request is served from Google's cache.
     * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
     * managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
     */
    type?: pulumi.Input<string>;
}
