// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.RegionNetworkFirewallPolicyRuleMatchLayer4Config;
import com.pulumi.gcp.compute.outputs.RegionNetworkFirewallPolicyRuleMatchSrcSecureTag;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RegionNetworkFirewallPolicyRuleMatch {
    /**
     * @return Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
     * 
     */
    private @Nullable List<String> destAddressGroups;
    /**
     * @return Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.
     * 
     */
    private @Nullable List<String> destFqdns;
    /**
     * @return CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
     * 
     */
    private @Nullable List<String> destIpRanges;
    /**
     * @return The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.
     * 
     */
    private @Nullable List<String> destRegionCodes;
    /**
     * @return Name of the Google Cloud Threat Intelligence list.
     * 
     */
    private @Nullable List<String> destThreatIntelligences;
    /**
     * @return Pairs of IP protocols and ports that the rule should match.
     * 
     */
    private List<RegionNetworkFirewallPolicyRuleMatchLayer4Config> layer4Configs;
    /**
     * @return Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
     * 
     */
    private @Nullable List<String> srcAddressGroups;
    /**
     * @return Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.
     * 
     */
    private @Nullable List<String> srcFqdns;
    /**
     * @return CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
     * 
     */
    private @Nullable List<String> srcIpRanges;
    /**
     * @return The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.
     * 
     */
    private @Nullable List<String> srcRegionCodes;
    /**
     * @return List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the &lt;code&gt;srcSecureTag&lt;/code&gt; are INEFFECTIVE, and there is no &lt;code&gt;srcIpRange&lt;/code&gt;, this rule will be ignored. Maximum number of source tag values allowed is 256.
     * 
     */
    private @Nullable List<RegionNetworkFirewallPolicyRuleMatchSrcSecureTag> srcSecureTags;
    /**
     * @return Name of the Google Cloud Threat Intelligence list.
     * 
     */
    private @Nullable List<String> srcThreatIntelligences;

    private RegionNetworkFirewallPolicyRuleMatch() {}
    /**
     * @return Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
     * 
     */
    public List<String> destAddressGroups() {
        return this.destAddressGroups == null ? List.of() : this.destAddressGroups;
    }
    /**
     * @return Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.
     * 
     */
    public List<String> destFqdns() {
        return this.destFqdns == null ? List.of() : this.destFqdns;
    }
    /**
     * @return CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
     * 
     */
    public List<String> destIpRanges() {
        return this.destIpRanges == null ? List.of() : this.destIpRanges;
    }
    /**
     * @return The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.
     * 
     */
    public List<String> destRegionCodes() {
        return this.destRegionCodes == null ? List.of() : this.destRegionCodes;
    }
    /**
     * @return Name of the Google Cloud Threat Intelligence list.
     * 
     */
    public List<String> destThreatIntelligences() {
        return this.destThreatIntelligences == null ? List.of() : this.destThreatIntelligences;
    }
    /**
     * @return Pairs of IP protocols and ports that the rule should match.
     * 
     */
    public List<RegionNetworkFirewallPolicyRuleMatchLayer4Config> layer4Configs() {
        return this.layer4Configs;
    }
    /**
     * @return Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
     * 
     */
    public List<String> srcAddressGroups() {
        return this.srcAddressGroups == null ? List.of() : this.srcAddressGroups;
    }
    /**
     * @return Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.
     * 
     */
    public List<String> srcFqdns() {
        return this.srcFqdns == null ? List.of() : this.srcFqdns;
    }
    /**
     * @return CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
     * 
     */
    public List<String> srcIpRanges() {
        return this.srcIpRanges == null ? List.of() : this.srcIpRanges;
    }
    /**
     * @return The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.
     * 
     */
    public List<String> srcRegionCodes() {
        return this.srcRegionCodes == null ? List.of() : this.srcRegionCodes;
    }
    /**
     * @return List of secure tag values, which should be matched at the source of the traffic. For INGRESS rule, if all the &lt;code&gt;srcSecureTag&lt;/code&gt; are INEFFECTIVE, and there is no &lt;code&gt;srcIpRange&lt;/code&gt;, this rule will be ignored. Maximum number of source tag values allowed is 256.
     * 
     */
    public List<RegionNetworkFirewallPolicyRuleMatchSrcSecureTag> srcSecureTags() {
        return this.srcSecureTags == null ? List.of() : this.srcSecureTags;
    }
    /**
     * @return Name of the Google Cloud Threat Intelligence list.
     * 
     */
    public List<String> srcThreatIntelligences() {
        return this.srcThreatIntelligences == null ? List.of() : this.srcThreatIntelligences;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionNetworkFirewallPolicyRuleMatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> destAddressGroups;
        private @Nullable List<String> destFqdns;
        private @Nullable List<String> destIpRanges;
        private @Nullable List<String> destRegionCodes;
        private @Nullable List<String> destThreatIntelligences;
        private List<RegionNetworkFirewallPolicyRuleMatchLayer4Config> layer4Configs;
        private @Nullable List<String> srcAddressGroups;
        private @Nullable List<String> srcFqdns;
        private @Nullable List<String> srcIpRanges;
        private @Nullable List<String> srcRegionCodes;
        private @Nullable List<RegionNetworkFirewallPolicyRuleMatchSrcSecureTag> srcSecureTags;
        private @Nullable List<String> srcThreatIntelligences;
        public Builder() {}
        public Builder(RegionNetworkFirewallPolicyRuleMatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.destAddressGroups = defaults.destAddressGroups;
    	      this.destFqdns = defaults.destFqdns;
    	      this.destIpRanges = defaults.destIpRanges;
    	      this.destRegionCodes = defaults.destRegionCodes;
    	      this.destThreatIntelligences = defaults.destThreatIntelligences;
    	      this.layer4Configs = defaults.layer4Configs;
    	      this.srcAddressGroups = defaults.srcAddressGroups;
    	      this.srcFqdns = defaults.srcFqdns;
    	      this.srcIpRanges = defaults.srcIpRanges;
    	      this.srcRegionCodes = defaults.srcRegionCodes;
    	      this.srcSecureTags = defaults.srcSecureTags;
    	      this.srcThreatIntelligences = defaults.srcThreatIntelligences;
        }

        @CustomType.Setter
        public Builder destAddressGroups(@Nullable List<String> destAddressGroups) {
            this.destAddressGroups = destAddressGroups;
            return this;
        }
        public Builder destAddressGroups(String... destAddressGroups) {
            return destAddressGroups(List.of(destAddressGroups));
        }
        @CustomType.Setter
        public Builder destFqdns(@Nullable List<String> destFqdns) {
            this.destFqdns = destFqdns;
            return this;
        }
        public Builder destFqdns(String... destFqdns) {
            return destFqdns(List.of(destFqdns));
        }
        @CustomType.Setter
        public Builder destIpRanges(@Nullable List<String> destIpRanges) {
            this.destIpRanges = destIpRanges;
            return this;
        }
        public Builder destIpRanges(String... destIpRanges) {
            return destIpRanges(List.of(destIpRanges));
        }
        @CustomType.Setter
        public Builder destRegionCodes(@Nullable List<String> destRegionCodes) {
            this.destRegionCodes = destRegionCodes;
            return this;
        }
        public Builder destRegionCodes(String... destRegionCodes) {
            return destRegionCodes(List.of(destRegionCodes));
        }
        @CustomType.Setter
        public Builder destThreatIntelligences(@Nullable List<String> destThreatIntelligences) {
            this.destThreatIntelligences = destThreatIntelligences;
            return this;
        }
        public Builder destThreatIntelligences(String... destThreatIntelligences) {
            return destThreatIntelligences(List.of(destThreatIntelligences));
        }
        @CustomType.Setter
        public Builder layer4Configs(List<RegionNetworkFirewallPolicyRuleMatchLayer4Config> layer4Configs) {
            this.layer4Configs = Objects.requireNonNull(layer4Configs);
            return this;
        }
        public Builder layer4Configs(RegionNetworkFirewallPolicyRuleMatchLayer4Config... layer4Configs) {
            return layer4Configs(List.of(layer4Configs));
        }
        @CustomType.Setter
        public Builder srcAddressGroups(@Nullable List<String> srcAddressGroups) {
            this.srcAddressGroups = srcAddressGroups;
            return this;
        }
        public Builder srcAddressGroups(String... srcAddressGroups) {
            return srcAddressGroups(List.of(srcAddressGroups));
        }
        @CustomType.Setter
        public Builder srcFqdns(@Nullable List<String> srcFqdns) {
            this.srcFqdns = srcFqdns;
            return this;
        }
        public Builder srcFqdns(String... srcFqdns) {
            return srcFqdns(List.of(srcFqdns));
        }
        @CustomType.Setter
        public Builder srcIpRanges(@Nullable List<String> srcIpRanges) {
            this.srcIpRanges = srcIpRanges;
            return this;
        }
        public Builder srcIpRanges(String... srcIpRanges) {
            return srcIpRanges(List.of(srcIpRanges));
        }
        @CustomType.Setter
        public Builder srcRegionCodes(@Nullable List<String> srcRegionCodes) {
            this.srcRegionCodes = srcRegionCodes;
            return this;
        }
        public Builder srcRegionCodes(String... srcRegionCodes) {
            return srcRegionCodes(List.of(srcRegionCodes));
        }
        @CustomType.Setter
        public Builder srcSecureTags(@Nullable List<RegionNetworkFirewallPolicyRuleMatchSrcSecureTag> srcSecureTags) {
            this.srcSecureTags = srcSecureTags;
            return this;
        }
        public Builder srcSecureTags(RegionNetworkFirewallPolicyRuleMatchSrcSecureTag... srcSecureTags) {
            return srcSecureTags(List.of(srcSecureTags));
        }
        @CustomType.Setter
        public Builder srcThreatIntelligences(@Nullable List<String> srcThreatIntelligences) {
            this.srcThreatIntelligences = srcThreatIntelligences;
            return this;
        }
        public Builder srcThreatIntelligences(String... srcThreatIntelligences) {
            return srcThreatIntelligences(List.of(srcThreatIntelligences));
        }
        public RegionNetworkFirewallPolicyRuleMatch build() {
            final var o = new RegionNetworkFirewallPolicyRuleMatch();
            o.destAddressGroups = destAddressGroups;
            o.destFqdns = destFqdns;
            o.destIpRanges = destIpRanges;
            o.destRegionCodes = destRegionCodes;
            o.destThreatIntelligences = destThreatIntelligences;
            o.layer4Configs = layer4Configs;
            o.srcAddressGroups = srcAddressGroups;
            o.srcFqdns = srcFqdns;
            o.srcIpRanges = srcIpRanges;
            o.srcRegionCodes = srcRegionCodes;
            o.srcSecureTags = srcSecureTags;
            o.srcThreatIntelligences = srcThreatIntelligences;
            return o;
        }
    }
}
