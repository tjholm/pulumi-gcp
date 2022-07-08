// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.binaryauthorization.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.binaryauthorization.inputs.PolicyAdmissionWhitelistPatternArgs;
import com.pulumi.gcp.binaryauthorization.inputs.PolicyClusterAdmissionRuleArgs;
import com.pulumi.gcp.binaryauthorization.inputs.PolicyDefaultAdmissionRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyState extends com.pulumi.resources.ResourceArgs {

    public static final PolicyState Empty = new PolicyState();

    /**
     * A whitelist of image patterns to exclude from admission rules. If an
     * image&#39;s name matches a whitelist pattern, the image&#39;s admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     * 
     */
    @Import(name="admissionWhitelistPatterns")
    private @Nullable Output<List<PolicyAdmissionWhitelistPatternArgs>> admissionWhitelistPatterns;

    /**
     * @return A whitelist of image patterns to exclude from admission rules. If an
     * image&#39;s name matches a whitelist pattern, the image&#39;s admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<PolicyAdmissionWhitelistPatternArgs>>> admissionWhitelistPatterns() {
        return Optional.ofNullable(this.admissionWhitelistPatterns);
    }

    /**
     * Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     * 
     */
    @Import(name="clusterAdmissionRules")
    private @Nullable Output<List<PolicyClusterAdmissionRuleArgs>> clusterAdmissionRules;

    /**
     * @return Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     * 
     */
    public Optional<Output<List<PolicyClusterAdmissionRuleArgs>>> clusterAdmissionRules() {
        return Optional.ofNullable(this.clusterAdmissionRules);
    }

    /**
     * Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     * 
     */
    @Import(name="defaultAdmissionRule")
    private @Nullable Output<PolicyDefaultAdmissionRuleArgs> defaultAdmissionRule;

    /**
     * @return Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     * 
     */
    public Optional<Output<PolicyDefaultAdmissionRuleArgs>> defaultAdmissionRule() {
        return Optional.ofNullable(this.defaultAdmissionRule);
    }

    /**
     * A descriptive comment.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A descriptive comment.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are `ENABLE` and `DISABLE`.
     * 
     */
    @Import(name="globalPolicyEvaluationMode")
    private @Nullable Output<String> globalPolicyEvaluationMode;

    /**
     * @return Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are `ENABLE` and `DISABLE`.
     * 
     */
    public Optional<Output<String>> globalPolicyEvaluationMode() {
        return Optional.ofNullable(this.globalPolicyEvaluationMode);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private PolicyState() {}

    private PolicyState(PolicyState $) {
        this.admissionWhitelistPatterns = $.admissionWhitelistPatterns;
        this.clusterAdmissionRules = $.clusterAdmissionRules;
        this.defaultAdmissionRule = $.defaultAdmissionRule;
        this.description = $.description;
        this.globalPolicyEvaluationMode = $.globalPolicyEvaluationMode;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyState $;

        public Builder() {
            $ = new PolicyState();
        }

        public Builder(PolicyState defaults) {
            $ = new PolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param admissionWhitelistPatterns A whitelist of image patterns to exclude from admission rules. If an
         * image&#39;s name matches a whitelist pattern, the image&#39;s admission
         * requests will always be permitted regardless of your admission rules.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder admissionWhitelistPatterns(@Nullable Output<List<PolicyAdmissionWhitelistPatternArgs>> admissionWhitelistPatterns) {
            $.admissionWhitelistPatterns = admissionWhitelistPatterns;
            return this;
        }

        /**
         * @param admissionWhitelistPatterns A whitelist of image patterns to exclude from admission rules. If an
         * image&#39;s name matches a whitelist pattern, the image&#39;s admission
         * requests will always be permitted regardless of your admission rules.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder admissionWhitelistPatterns(List<PolicyAdmissionWhitelistPatternArgs> admissionWhitelistPatterns) {
            return admissionWhitelistPatterns(Output.of(admissionWhitelistPatterns));
        }

        /**
         * @param admissionWhitelistPatterns A whitelist of image patterns to exclude from admission rules. If an
         * image&#39;s name matches a whitelist pattern, the image&#39;s admission
         * requests will always be permitted regardless of your admission rules.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder admissionWhitelistPatterns(PolicyAdmissionWhitelistPatternArgs... admissionWhitelistPatterns) {
            return admissionWhitelistPatterns(List.of(admissionWhitelistPatterns));
        }

        /**
         * @param clusterAdmissionRules Per-cluster admission rules. An admission rule specifies either that
         * all container images used in a pod creation request must be attested
         * to by one or more attestors, that all pod creations will be allowed,
         * or that all pod creations will be denied. There can be at most one
         * admission rule per cluster spec.
         * 
         * @return builder
         * 
         */
        public Builder clusterAdmissionRules(@Nullable Output<List<PolicyClusterAdmissionRuleArgs>> clusterAdmissionRules) {
            $.clusterAdmissionRules = clusterAdmissionRules;
            return this;
        }

        /**
         * @param clusterAdmissionRules Per-cluster admission rules. An admission rule specifies either that
         * all container images used in a pod creation request must be attested
         * to by one or more attestors, that all pod creations will be allowed,
         * or that all pod creations will be denied. There can be at most one
         * admission rule per cluster spec.
         * 
         * @return builder
         * 
         */
        public Builder clusterAdmissionRules(List<PolicyClusterAdmissionRuleArgs> clusterAdmissionRules) {
            return clusterAdmissionRules(Output.of(clusterAdmissionRules));
        }

        /**
         * @param clusterAdmissionRules Per-cluster admission rules. An admission rule specifies either that
         * all container images used in a pod creation request must be attested
         * to by one or more attestors, that all pod creations will be allowed,
         * or that all pod creations will be denied. There can be at most one
         * admission rule per cluster spec.
         * 
         * @return builder
         * 
         */
        public Builder clusterAdmissionRules(PolicyClusterAdmissionRuleArgs... clusterAdmissionRules) {
            return clusterAdmissionRules(List.of(clusterAdmissionRules));
        }

        /**
         * @param defaultAdmissionRule Default admission rule for a cluster without a per-cluster admission
         * rule.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder defaultAdmissionRule(@Nullable Output<PolicyDefaultAdmissionRuleArgs> defaultAdmissionRule) {
            $.defaultAdmissionRule = defaultAdmissionRule;
            return this;
        }

        /**
         * @param defaultAdmissionRule Default admission rule for a cluster without a per-cluster admission
         * rule.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder defaultAdmissionRule(PolicyDefaultAdmissionRuleArgs defaultAdmissionRule) {
            return defaultAdmissionRule(Output.of(defaultAdmissionRule));
        }

        /**
         * @param description A descriptive comment.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A descriptive comment.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param globalPolicyEvaluationMode Controls the evaluation of a Google-maintained global admission policy
         * for common system-level images. Images not covered by the global
         * policy will be subject to the project admission policy.
         * Possible values are `ENABLE` and `DISABLE`.
         * 
         * @return builder
         * 
         */
        public Builder globalPolicyEvaluationMode(@Nullable Output<String> globalPolicyEvaluationMode) {
            $.globalPolicyEvaluationMode = globalPolicyEvaluationMode;
            return this;
        }

        /**
         * @param globalPolicyEvaluationMode Controls the evaluation of a Google-maintained global admission policy
         * for common system-level images. Images not covered by the global
         * policy will be subject to the project admission policy.
         * Possible values are `ENABLE` and `DISABLE`.
         * 
         * @return builder
         * 
         */
        public Builder globalPolicyEvaluationMode(String globalPolicyEvaluationMode) {
            return globalPolicyEvaluationMode(Output.of(globalPolicyEvaluationMode));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public PolicyState build() {
            return $;
        }
    }

}
