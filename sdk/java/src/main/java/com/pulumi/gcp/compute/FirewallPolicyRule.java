// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.FirewallPolicyRuleArgs;
import com.pulumi.gcp.compute.inputs.FirewallPolicyRuleState;
import com.pulumi.gcp.compute.outputs.FirewallPolicyRuleMatch;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Compute FirewallPolicyRule resource
 * 
 * ## Example Usage
 * 
 * ### Basic_fir_sec_rule
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.networksecurity.AddressGroup;
 * import com.pulumi.gcp.networksecurity.AddressGroupArgs;
 * import com.pulumi.gcp.organizations.Folder;
 * import com.pulumi.gcp.organizations.FolderArgs;
 * import com.pulumi.gcp.compute.FirewallPolicy;
 * import com.pulumi.gcp.compute.FirewallPolicyArgs;
 * import com.pulumi.gcp.compute.FirewallPolicyRule;
 * import com.pulumi.gcp.compute.FirewallPolicyRuleArgs;
 * import com.pulumi.gcp.compute.inputs.FirewallPolicyRuleMatchArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var basicGlobalNetworksecurityAddressGroup = new AddressGroup("basicGlobalNetworksecurityAddressGroup", AddressGroupArgs.builder()
 *             .name("policy")
 *             .parent("organizations/123456789")
 *             .description("Sample global networksecurity_address_group")
 *             .location("global")
 *             .items("208.80.154.224/32")
 *             .type("IPV4")
 *             .capacity(100)
 *             .build());
 * 
 *         var folder = new Folder("folder", FolderArgs.builder()
 *             .displayName("policy")
 *             .parent("organizations/123456789")
 *             .build());
 * 
 *         var default_ = new FirewallPolicy("default", FirewallPolicyArgs.builder()
 *             .parent(folder.id())
 *             .shortName("policy")
 *             .description("Resource created for Terraform acceptance testing")
 *             .build());
 * 
 *         var primary = new FirewallPolicyRule("primary", FirewallPolicyRuleArgs.builder()
 *             .firewallPolicy(default_.name())
 *             .description("Resource created for Terraform acceptance testing")
 *             .priority(9000)
 *             .enableLogging(true)
 *             .action("allow")
 *             .direction("EGRESS")
 *             .disabled(false)
 *             .match(FirewallPolicyRuleMatchArgs.builder()
 *                 .layer4Configs(                
 *                     FirewallPolicyRuleMatchLayer4ConfigArgs.builder()
 *                         .ipProtocol("tcp")
 *                         .ports(8080)
 *                         .build(),
 *                     FirewallPolicyRuleMatchLayer4ConfigArgs.builder()
 *                         .ipProtocol("udp")
 *                         .ports(22)
 *                         .build())
 *                 .destIpRanges("11.100.0.1/32")
 *                 .destFqdns()
 *                 .destRegionCodes("US")
 *                 .destThreatIntelligences("iplist-known-malicious-ips")
 *                 .srcAddressGroups()
 *                 .destAddressGroups(basicGlobalNetworksecurityAddressGroup.id())
 *                 .build())
 *             .targetServiceAccounts("my{@literal @}service-account.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * FirewallPolicyRule can be imported using any of these accepted formats:
 * 
 * * `locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}`
 * 
 * * `{{firewall_policy}}/{{priority}}`
 * 
 * When using the `pulumi import` command, FirewallPolicyRule can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default {{firewall_policy}}/{{priority}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/firewallPolicyRule:FirewallPolicyRule")
public class FirewallPolicyRule extends com.pulumi.resources.CustomResource {
    /**
     * The Action to perform when the client connection triggers the rule. Valid actions are &#34;allow&#34;, &#34;deny&#34;, &#34;goto_next&#34; and &#34;apply_security_profile_group&#34;.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return The Action to perform when the client connection triggers the rule. Valid actions are &#34;allow&#34;, &#34;deny&#34;, &#34;goto_next&#34; and &#34;apply_security_profile_group&#34;.
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * An optional description for this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description for this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The direction in which this rule applies. Possible values: INGRESS, EGRESS
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return The direction in which this rule applies. Possible values: INGRESS, EGRESS
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and
     * traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and
     * traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured
     * export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on
     * &#34;goto_next&#34; rules.
     * 
     */
    @Export(name="enableLogging", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableLogging;

    /**
     * @return Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured
     * export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on
     * &#34;goto_next&#34; rules.
     * 
     */
    public Output<Optional<Boolean>> enableLogging() {
        return Codegen.optional(this.enableLogging);
    }
    /**
     * The firewall policy of the resource.
     * 
     */
    @Export(name="firewallPolicy", refs={String.class}, tree="[0]")
    private Output<String> firewallPolicy;

    /**
     * @return The firewall policy of the resource.
     * 
     */
    public Output<String> firewallPolicy() {
        return this.firewallPolicy;
    }
    /**
     * Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output<String> kind;

    /**
     * @return Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     * 
     */
    public Output<String> kind() {
        return this.kind;
    }
    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced.
     * 
     */
    @Export(name="match", refs={FirewallPolicyRuleMatch.class}, tree="[0]")
    private Output<FirewallPolicyRuleMatch> match;

    /**
     * @return A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced.
     * 
     */
    public Output<FirewallPolicyRuleMatch> match() {
        return this.match;
    }
    /**
     * An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output<Integer> priority;

    /**
     * @return An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * Calculation of the complexity of a single firewall policy rule.
     * 
     */
    @Export(name="ruleTupleCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> ruleTupleCount;

    /**
     * @return Calculation of the complexity of a single firewall policy rule.
     * 
     */
    public Output<Integer> ruleTupleCount() {
        return this.ruleTupleCount;
    }
    /**
     * A fully-qualified URL of a SecurityProfileGroup resource. Example:
     * https://networksecurity.googleapis.com/v1/organizations/{organizationId}/locations/global/securityProfileGroups/my-security-profile-group.
     * It must be specified if action = &#39;apply_security_profile_group&#39; and cannot be specified for other actions.
     * 
     */
    @Export(name="securityProfileGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityProfileGroup;

    /**
     * @return A fully-qualified URL of a SecurityProfileGroup resource. Example:
     * https://networksecurity.googleapis.com/v1/organizations/{organizationId}/locations/global/securityProfileGroups/my-security-profile-group.
     * It must be specified if action = &#39;apply_security_profile_group&#39; and cannot be specified for other actions.
     * 
     */
    public Output<Optional<String>> securityProfileGroup() {
        return Codegen.optional(this.securityProfileGroup);
    }
    /**
     * A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get
     * this rule. If this field is left blank, all VMs within the organization will receive the rule.
     * 
     */
    @Export(name="targetResources", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> targetResources;

    /**
     * @return A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get
     * this rule. If this field is left blank, all VMs within the organization will receive the rule.
     * 
     */
    public Output<Optional<List<String>>> targetResources() {
        return Codegen.optional(this.targetResources);
    }
    /**
     * A list of service accounts indicating the sets of instances that are applied with this rule.
     * 
     */
    @Export(name="targetServiceAccounts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> targetServiceAccounts;

    /**
     * @return A list of service accounts indicating the sets of instances that are applied with this rule.
     * 
     */
    public Output<Optional<List<String>>> targetServiceAccounts() {
        return Codegen.optional(this.targetServiceAccounts);
    }
    /**
     * Boolean flag indicating if the traffic should be TLS decrypted. It can be set only if action =
     * &#39;apply_security_profile_group&#39; and cannot be set for other actions.
     * 
     */
    @Export(name="tlsInspect", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tlsInspect;

    /**
     * @return Boolean flag indicating if the traffic should be TLS decrypted. It can be set only if action =
     * &#39;apply_security_profile_group&#39; and cannot be set for other actions.
     * 
     */
    public Output<Optional<Boolean>> tlsInspect() {
        return Codegen.optional(this.tlsInspect);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FirewallPolicyRule(String name) {
        this(name, FirewallPolicyRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FirewallPolicyRule(String name, FirewallPolicyRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FirewallPolicyRule(String name, FirewallPolicyRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewallPolicyRule:FirewallPolicyRule", name, args == null ? FirewallPolicyRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FirewallPolicyRule(String name, Output<String> id, @Nullable FirewallPolicyRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewallPolicyRule:FirewallPolicyRule", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static FirewallPolicyRule get(String name, Output<String> id, @Nullable FirewallPolicyRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FirewallPolicyRule(name, id, state, options);
    }
}
