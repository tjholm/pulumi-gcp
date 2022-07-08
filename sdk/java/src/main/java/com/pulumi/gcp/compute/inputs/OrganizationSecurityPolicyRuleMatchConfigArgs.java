// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationSecurityPolicyRuleMatchConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationSecurityPolicyRuleMatchConfigArgs Empty = new OrganizationSecurityPolicyRuleMatchConfigArgs();

    /**
     * Destination IP address range in CIDR format. Required for
     * EGRESS rules.
     * 
     */
    @Import(name="destIpRanges")
    private @Nullable Output<List<String>> destIpRanges;

    /**
     * @return Destination IP address range in CIDR format. Required for
     * EGRESS rules.
     * 
     */
    public Optional<Output<List<String>>> destIpRanges() {
        return Optional.ofNullable(this.destIpRanges);
    }

    /**
     * Pairs of IP protocols and ports that the rule should match.
     * Structure is documented below.
     * 
     */
    @Import(name="layer4Configs", required=true)
    private Output<List<OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs>> layer4Configs;

    /**
     * @return Pairs of IP protocols and ports that the rule should match.
     * Structure is documented below.
     * 
     */
    public Output<List<OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs>> layer4Configs() {
        return this.layer4Configs;
    }

    /**
     * Source IP address range in CIDR format. Required for
     * INGRESS rules.
     * 
     */
    @Import(name="srcIpRanges")
    private @Nullable Output<List<String>> srcIpRanges;

    /**
     * @return Source IP address range in CIDR format. Required for
     * INGRESS rules.
     * 
     */
    public Optional<Output<List<String>>> srcIpRanges() {
        return Optional.ofNullable(this.srcIpRanges);
    }

    private OrganizationSecurityPolicyRuleMatchConfigArgs() {}

    private OrganizationSecurityPolicyRuleMatchConfigArgs(OrganizationSecurityPolicyRuleMatchConfigArgs $) {
        this.destIpRanges = $.destIpRanges;
        this.layer4Configs = $.layer4Configs;
        this.srcIpRanges = $.srcIpRanges;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationSecurityPolicyRuleMatchConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationSecurityPolicyRuleMatchConfigArgs $;

        public Builder() {
            $ = new OrganizationSecurityPolicyRuleMatchConfigArgs();
        }

        public Builder(OrganizationSecurityPolicyRuleMatchConfigArgs defaults) {
            $ = new OrganizationSecurityPolicyRuleMatchConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param destIpRanges Destination IP address range in CIDR format. Required for
         * EGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder destIpRanges(@Nullable Output<List<String>> destIpRanges) {
            $.destIpRanges = destIpRanges;
            return this;
        }

        /**
         * @param destIpRanges Destination IP address range in CIDR format. Required for
         * EGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder destIpRanges(List<String> destIpRanges) {
            return destIpRanges(Output.of(destIpRanges));
        }

        /**
         * @param destIpRanges Destination IP address range in CIDR format. Required for
         * EGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder destIpRanges(String... destIpRanges) {
            return destIpRanges(List.of(destIpRanges));
        }

        /**
         * @param layer4Configs Pairs of IP protocols and ports that the rule should match.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder layer4Configs(Output<List<OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs>> layer4Configs) {
            $.layer4Configs = layer4Configs;
            return this;
        }

        /**
         * @param layer4Configs Pairs of IP protocols and ports that the rule should match.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder layer4Configs(List<OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs> layer4Configs) {
            return layer4Configs(Output.of(layer4Configs));
        }

        /**
         * @param layer4Configs Pairs of IP protocols and ports that the rule should match.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder layer4Configs(OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs... layer4Configs) {
            return layer4Configs(List.of(layer4Configs));
        }

        /**
         * @param srcIpRanges Source IP address range in CIDR format. Required for
         * INGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder srcIpRanges(@Nullable Output<List<String>> srcIpRanges) {
            $.srcIpRanges = srcIpRanges;
            return this;
        }

        /**
         * @param srcIpRanges Source IP address range in CIDR format. Required for
         * INGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder srcIpRanges(List<String> srcIpRanges) {
            return srcIpRanges(Output.of(srcIpRanges));
        }

        /**
         * @param srcIpRanges Source IP address range in CIDR format. Required for
         * INGRESS rules.
         * 
         * @return builder
         * 
         */
        public Builder srcIpRanges(String... srcIpRanges) {
            return srcIpRanges(List.of(srcIpRanges));
        }

        public OrganizationSecurityPolicyRuleMatchConfigArgs build() {
            $.layer4Configs = Objects.requireNonNull($.layer4Configs, "expected parameter 'layer4Configs' to be non-null");
            return $;
        }
    }

}
