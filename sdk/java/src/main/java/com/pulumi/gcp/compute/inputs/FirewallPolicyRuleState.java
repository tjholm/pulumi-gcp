// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.FirewallPolicyRuleMatchArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FirewallPolicyRuleState extends com.pulumi.resources.ResourceArgs {

    public static final FirewallPolicyRuleState Empty = new FirewallPolicyRuleState();

    /**
     * The Action to perform when the client connection triggers the rule. Can currently be either &#34;allow&#34; or &#34;deny()&#34; where valid values for status are 403, 404, and 502.
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return The Action to perform when the client connection triggers the rule. Can currently be either &#34;allow&#34; or &#34;deny()&#34; where valid values for status are 403, 404, and 502.
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * An optional description for this resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description for this resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The direction in which this rule applies. Possible values: INGRESS, EGRESS
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return The direction in which this rule applies. Possible values: INGRESS, EGRESS
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on &#34;goto_next&#34; rules.
     * 
     */
    @Import(name="enableLogging")
    private @Nullable Output<Boolean> enableLogging;

    /**
     * @return Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on &#34;goto_next&#34; rules.
     * 
     */
    public Optional<Output<Boolean>> enableLogging() {
        return Optional.ofNullable(this.enableLogging);
    }

    /**
     * The firewall policy of the resource.
     * 
     */
    @Import(name="firewallPolicy")
    private @Nullable Output<String> firewallPolicy;

    /**
     * @return The firewall policy of the resource.
     * 
     */
    public Optional<Output<String>> firewallPolicy() {
        return Optional.ofNullable(this.firewallPolicy);
    }

    /**
     * Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     * 
     */
    @Import(name="kind")
    private @Nullable Output<String> kind;

    /**
     * @return Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
     * 
     */
    public Optional<Output<String>> kind() {
        return Optional.ofNullable(this.kind);
    }

    /**
     * A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced. Structure is documented below.
     * 
     */
    @Import(name="match")
    private @Nullable Output<FirewallPolicyRuleMatchArgs> match;

    /**
     * @return A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced. Structure is documented below.
     * 
     */
    public Optional<Output<FirewallPolicyRuleMatchArgs>> match() {
        return Optional.ofNullable(this.match);
    }

    /**
     * An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * Calculation of the complexity of a single firewall policy rule.
     * 
     */
    @Import(name="ruleTupleCount")
    private @Nullable Output<Integer> ruleTupleCount;

    /**
     * @return Calculation of the complexity of a single firewall policy rule.
     * 
     */
    public Optional<Output<Integer>> ruleTupleCount() {
        return Optional.ofNullable(this.ruleTupleCount);
    }

    /**
     * A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
     * 
     */
    @Import(name="targetResources")
    private @Nullable Output<List<String>> targetResources;

    /**
     * @return A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
     * 
     */
    public Optional<Output<List<String>>> targetResources() {
        return Optional.ofNullable(this.targetResources);
    }

    /**
     * A list of service accounts indicating the sets of instances that are applied with this rule.
     * 
     */
    @Import(name="targetServiceAccounts")
    private @Nullable Output<List<String>> targetServiceAccounts;

    /**
     * @return A list of service accounts indicating the sets of instances that are applied with this rule.
     * 
     */
    public Optional<Output<List<String>>> targetServiceAccounts() {
        return Optional.ofNullable(this.targetServiceAccounts);
    }

    private FirewallPolicyRuleState() {}

    private FirewallPolicyRuleState(FirewallPolicyRuleState $) {
        this.action = $.action;
        this.description = $.description;
        this.direction = $.direction;
        this.disabled = $.disabled;
        this.enableLogging = $.enableLogging;
        this.firewallPolicy = $.firewallPolicy;
        this.kind = $.kind;
        this.match = $.match;
        this.priority = $.priority;
        this.ruleTupleCount = $.ruleTupleCount;
        this.targetResources = $.targetResources;
        this.targetServiceAccounts = $.targetServiceAccounts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FirewallPolicyRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FirewallPolicyRuleState $;

        public Builder() {
            $ = new FirewallPolicyRuleState();
        }

        public Builder(FirewallPolicyRuleState defaults) {
            $ = new FirewallPolicyRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The Action to perform when the client connection triggers the rule. Can currently be either &#34;allow&#34; or &#34;deny()&#34; where valid values for status are 403, 404, and 502.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The Action to perform when the client connection triggers the rule. Can currently be either &#34;allow&#34; or &#34;deny()&#34; where valid values for status are 403, 404, and 502.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param description An optional description for this resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description for this resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param direction The direction in which this rule applies. Possible values: INGRESS, EGRESS
         * 
         * @return builder
         * 
         */
        public Builder direction(@Nullable Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction The direction in which this rule applies. Possible values: INGRESS, EGRESS
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param disabled Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param enableLogging Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on &#34;goto_next&#34; rules.
         * 
         * @return builder
         * 
         */
        public Builder enableLogging(@Nullable Output<Boolean> enableLogging) {
            $.enableLogging = enableLogging;
            return this;
        }

        /**
         * @param enableLogging Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on &#34;goto_next&#34; rules.
         * 
         * @return builder
         * 
         */
        public Builder enableLogging(Boolean enableLogging) {
            return enableLogging(Output.of(enableLogging));
        }

        /**
         * @param firewallPolicy The firewall policy of the resource.
         * 
         * @return builder
         * 
         */
        public Builder firewallPolicy(@Nullable Output<String> firewallPolicy) {
            $.firewallPolicy = firewallPolicy;
            return this;
        }

        /**
         * @param firewallPolicy The firewall policy of the resource.
         * 
         * @return builder
         * 
         */
        public Builder firewallPolicy(String firewallPolicy) {
            return firewallPolicy(Output.of(firewallPolicy));
        }

        /**
         * @param kind Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
         * 
         * @return builder
         * 
         */
        public Builder kind(@Nullable Output<String> kind) {
            $.kind = kind;
            return this;
        }

        /**
         * @param kind Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
         * 
         * @return builder
         * 
         */
        public Builder kind(String kind) {
            return kind(Output.of(kind));
        }

        /**
         * @param match A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder match(@Nullable Output<FirewallPolicyRuleMatchArgs> match) {
            $.match = match;
            return this;
        }

        /**
         * @param match A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding &#39;action&#39; is enforced. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder match(FirewallPolicyRuleMatchArgs match) {
            return match(Output.of(match));
        }

        /**
         * @param priority An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param ruleTupleCount Calculation of the complexity of a single firewall policy rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleTupleCount(@Nullable Output<Integer> ruleTupleCount) {
            $.ruleTupleCount = ruleTupleCount;
            return this;
        }

        /**
         * @param ruleTupleCount Calculation of the complexity of a single firewall policy rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleTupleCount(Integer ruleTupleCount) {
            return ruleTupleCount(Output.of(ruleTupleCount));
        }

        /**
         * @param targetResources A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
         * 
         * @return builder
         * 
         */
        public Builder targetResources(@Nullable Output<List<String>> targetResources) {
            $.targetResources = targetResources;
            return this;
        }

        /**
         * @param targetResources A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
         * 
         * @return builder
         * 
         */
        public Builder targetResources(List<String> targetResources) {
            return targetResources(Output.of(targetResources));
        }

        /**
         * @param targetResources A list of network resource URLs to which this rule applies. This field allows you to control which network&#39;s VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
         * 
         * @return builder
         * 
         */
        public Builder targetResources(String... targetResources) {
            return targetResources(List.of(targetResources));
        }

        /**
         * @param targetServiceAccounts A list of service accounts indicating the sets of instances that are applied with this rule.
         * 
         * @return builder
         * 
         */
        public Builder targetServiceAccounts(@Nullable Output<List<String>> targetServiceAccounts) {
            $.targetServiceAccounts = targetServiceAccounts;
            return this;
        }

        /**
         * @param targetServiceAccounts A list of service accounts indicating the sets of instances that are applied with this rule.
         * 
         * @return builder
         * 
         */
        public Builder targetServiceAccounts(List<String> targetServiceAccounts) {
            return targetServiceAccounts(Output.of(targetServiceAccounts));
        }

        /**
         * @param targetServiceAccounts A list of service accounts indicating the sets of instances that are applied with this rule.
         * 
         * @return builder
         * 
         */
        public Builder targetServiceAccounts(String... targetServiceAccounts) {
            return targetServiceAccounts(List.of(targetServiceAccounts));
        }

        public FirewallPolicyRuleState build() {
            return $;
        }
    }

}
