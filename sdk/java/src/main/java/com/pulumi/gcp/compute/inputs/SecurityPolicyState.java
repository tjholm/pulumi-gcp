// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.SecurityPolicyAdaptiveProtectionConfigArgs;
import com.pulumi.gcp.compute.inputs.SecurityPolicyAdvancedOptionsConfigArgs;
import com.pulumi.gcp.compute.inputs.SecurityPolicyRecaptchaOptionsConfigArgs;
import com.pulumi.gcp.compute.inputs.SecurityPolicyRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityPolicyState Empty = new SecurityPolicyState();

    /**
     * Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
     * 
     */
    @Import(name="adaptiveProtectionConfig")
    private @Nullable Output<SecurityPolicyAdaptiveProtectionConfigArgs> adaptiveProtectionConfig;

    /**
     * @return Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
     * 
     */
    public Optional<Output<SecurityPolicyAdaptiveProtectionConfigArgs>> adaptiveProtectionConfig() {
        return Optional.ofNullable(this.adaptiveProtectionConfig);
    }

    /**
     * [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
     * Structure is documented below.
     * 
     */
    @Import(name="advancedOptionsConfig")
    private @Nullable Output<SecurityPolicyAdvancedOptionsConfigArgs> advancedOptionsConfig;

    /**
     * @return [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
     * Structure is documented below.
     * 
     */
    public Optional<Output<SecurityPolicyAdvancedOptionsConfigArgs>> advancedOptionsConfig() {
        return Optional.ofNullable(this.advancedOptionsConfig);
    }

    /**
     * An optional description of this security policy. Max size is 2048.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description of this security policy. Max size is 2048.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Fingerprint of this resource.
     * 
     */
    @Import(name="fingerprint")
    private @Nullable Output<String> fingerprint;

    /**
     * @return Fingerprint of this resource.
     * 
     */
    public Optional<Output<String>> fingerprint() {
        return Optional.ofNullable(this.fingerprint);
    }

    /**
     * The name of the security policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the security policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
     * 
     */
    @Import(name="recaptchaOptionsConfig")
    private @Nullable Output<SecurityPolicyRecaptchaOptionsConfigArgs> recaptchaOptionsConfig;

    /**
     * @return [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
     * 
     */
    public Optional<Output<SecurityPolicyRecaptchaOptionsConfigArgs>> recaptchaOptionsConfig() {
        return Optional.ofNullable(this.recaptchaOptionsConfig);
    }

    /**
     * The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match &#34;\*&#34;). If no rules are provided when creating a
     * security policy, a default rule with action &#34;allow&#34; will be added. Structure is documented below.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<SecurityPolicyRuleArgs>> rules;

    /**
     * @return The set of rules that belong to this policy. There must always be a default
     * rule (rule with priority 2147483647 and match &#34;\*&#34;). If no rules are provided when creating a
     * security policy, a default rule with action &#34;allow&#34; will be added. Structure is documented below.
     * 
     */
    public Optional<Output<List<SecurityPolicyRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * The URI of the created resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * The type indicates the intended use of the security policy. This field can be set only at resource creation time.
     * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
     *   They filter requests before they hit the origin servers.
     * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
     *   (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
     *   They filter requests before the request is served from Google&#39;s cache.
     * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
     *   managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type indicates the intended use of the security policy. This field can be set only at resource creation time.
     * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
     *   They filter requests before they hit the origin servers.
     * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
     *   (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
     *   They filter requests before the request is served from Google&#39;s cache.
     * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
     *   managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private SecurityPolicyState() {}

    private SecurityPolicyState(SecurityPolicyState $) {
        this.adaptiveProtectionConfig = $.adaptiveProtectionConfig;
        this.advancedOptionsConfig = $.advancedOptionsConfig;
        this.description = $.description;
        this.fingerprint = $.fingerprint;
        this.name = $.name;
        this.project = $.project;
        this.recaptchaOptionsConfig = $.recaptchaOptionsConfig;
        this.rules = $.rules;
        this.selfLink = $.selfLink;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityPolicyState $;

        public Builder() {
            $ = new SecurityPolicyState();
        }

        public Builder(SecurityPolicyState defaults) {
            $ = new SecurityPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param adaptiveProtectionConfig Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder adaptiveProtectionConfig(@Nullable Output<SecurityPolicyAdaptiveProtectionConfigArgs> adaptiveProtectionConfig) {
            $.adaptiveProtectionConfig = adaptiveProtectionConfig;
            return this;
        }

        /**
         * @param adaptiveProtectionConfig Configuration for [Google Cloud Armor Adaptive Protection](https://cloud.google.com/armor/docs/adaptive-protection-overview?hl=en). Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder adaptiveProtectionConfig(SecurityPolicyAdaptiveProtectionConfigArgs adaptiveProtectionConfig) {
            return adaptiveProtectionConfig(Output.of(adaptiveProtectionConfig));
        }

        /**
         * @param advancedOptionsConfig [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder advancedOptionsConfig(@Nullable Output<SecurityPolicyAdvancedOptionsConfigArgs> advancedOptionsConfig) {
            $.advancedOptionsConfig = advancedOptionsConfig;
            return this;
        }

        /**
         * @param advancedOptionsConfig [Advanced Configuration Options](https://cloud.google.com/armor/docs/security-policy-overview#json-parsing).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder advancedOptionsConfig(SecurityPolicyAdvancedOptionsConfigArgs advancedOptionsConfig) {
            return advancedOptionsConfig(Output.of(advancedOptionsConfig));
        }

        /**
         * @param description An optional description of this security policy. Max size is 2048.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description of this security policy. Max size is 2048.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param fingerprint Fingerprint of this resource.
         * 
         * @return builder
         * 
         */
        public Builder fingerprint(@Nullable Output<String> fingerprint) {
            $.fingerprint = fingerprint;
            return this;
        }

        /**
         * @param fingerprint Fingerprint of this resource.
         * 
         * @return builder
         * 
         */
        public Builder fingerprint(String fingerprint) {
            return fingerprint(Output.of(fingerprint));
        }

        /**
         * @param name The name of the security policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the security policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param recaptchaOptionsConfig [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder recaptchaOptionsConfig(@Nullable Output<SecurityPolicyRecaptchaOptionsConfigArgs> recaptchaOptionsConfig) {
            $.recaptchaOptionsConfig = recaptchaOptionsConfig;
            return this;
        }

        /**
         * @param recaptchaOptionsConfig [reCAPTCHA Configuration Options](https://cloud.google.com/armor/docs/configure-security-policies?hl=en#use_a_manual_challenge_to_distinguish_between_human_or_automated_clients). Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder recaptchaOptionsConfig(SecurityPolicyRecaptchaOptionsConfigArgs recaptchaOptionsConfig) {
            return recaptchaOptionsConfig(Output.of(recaptchaOptionsConfig));
        }

        /**
         * @param rules The set of rules that belong to this policy. There must always be a default
         * rule (rule with priority 2147483647 and match &#34;\*&#34;). If no rules are provided when creating a
         * security policy, a default rule with action &#34;allow&#34; will be added. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<SecurityPolicyRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules The set of rules that belong to this policy. There must always be a default
         * rule (rule with priority 2147483647 and match &#34;\*&#34;). If no rules are provided when creating a
         * security policy, a default rule with action &#34;allow&#34; will be added. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<SecurityPolicyRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules The set of rules that belong to this policy. There must always be a default
         * rule (rule with priority 2147483647 and match &#34;\*&#34;). If no rules are provided when creating a
         * security policy, a default rule with action &#34;allow&#34; will be added. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(SecurityPolicyRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param type The type indicates the intended use of the security policy. This field can be set only at resource creation time.
         * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
         *   They filter requests before they hit the origin servers.
         * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
         *   (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
         *   They filter requests before the request is served from Google&#39;s cache.
         * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
         *   managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type indicates the intended use of the security policy. This field can be set only at resource creation time.
         * * CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services.
         *   They filter requests before they hit the origin servers.
         * * CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services
         *   (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
         *   They filter requests before the request is served from Google&#39;s cache.
         * * CLOUD_ARMOR_INTERNAL_SERVICE - Cloud Armor internal service policies can be configured to filter HTTP requests targeting services
         *   managed by Traffic Director in a service mesh. They filter requests before the request is served from the application.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SecurityPolicyState build() {
            return $;
        }
    }

}
