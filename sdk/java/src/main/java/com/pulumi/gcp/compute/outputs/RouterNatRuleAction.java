// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RouterNatRuleAction {
    /**
     * @return A list of URLs of the IP resources used for this NAT rule.
     * These IP addresses must be valid static external IP addresses assigned to the project.
     * This field is used for public NAT.
     * 
     */
    private @Nullable List<String> sourceNatActiveIps;
    /**
     * @return A list of URLs of the subnetworks used as source ranges for this NAT Rule.
     * These subnetworks must have purpose set to PRIVATE_NAT.
     * This field is used for private NAT.
     * 
     */
    private @Nullable List<String> sourceNatActiveRanges;
    /**
     * @return A list of URLs of the IP resources to be drained.
     * These IPs must be valid static external IPs that have been assigned to the NAT.
     * These IPs should be used for updating/patching a NAT rule only.
     * This field is used for public NAT.
     * 
     */
    private @Nullable List<String> sourceNatDrainIps;
    /**
     * @return A list of URLs of subnetworks representing source ranges to be drained.
     * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
     * This field is used for private NAT.
     * 
     */
    private @Nullable List<String> sourceNatDrainRanges;

    private RouterNatRuleAction() {}
    /**
     * @return A list of URLs of the IP resources used for this NAT rule.
     * These IP addresses must be valid static external IP addresses assigned to the project.
     * This field is used for public NAT.
     * 
     */
    public List<String> sourceNatActiveIps() {
        return this.sourceNatActiveIps == null ? List.of() : this.sourceNatActiveIps;
    }
    /**
     * @return A list of URLs of the subnetworks used as source ranges for this NAT Rule.
     * These subnetworks must have purpose set to PRIVATE_NAT.
     * This field is used for private NAT.
     * 
     */
    public List<String> sourceNatActiveRanges() {
        return this.sourceNatActiveRanges == null ? List.of() : this.sourceNatActiveRanges;
    }
    /**
     * @return A list of URLs of the IP resources to be drained.
     * These IPs must be valid static external IPs that have been assigned to the NAT.
     * These IPs should be used for updating/patching a NAT rule only.
     * This field is used for public NAT.
     * 
     */
    public List<String> sourceNatDrainIps() {
        return this.sourceNatDrainIps == null ? List.of() : this.sourceNatDrainIps;
    }
    /**
     * @return A list of URLs of subnetworks representing source ranges to be drained.
     * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
     * This field is used for private NAT.
     * 
     */
    public List<String> sourceNatDrainRanges() {
        return this.sourceNatDrainRanges == null ? List.of() : this.sourceNatDrainRanges;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RouterNatRuleAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> sourceNatActiveIps;
        private @Nullable List<String> sourceNatActiveRanges;
        private @Nullable List<String> sourceNatDrainIps;
        private @Nullable List<String> sourceNatDrainRanges;
        public Builder() {}
        public Builder(RouterNatRuleAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.sourceNatActiveIps = defaults.sourceNatActiveIps;
    	      this.sourceNatActiveRanges = defaults.sourceNatActiveRanges;
    	      this.sourceNatDrainIps = defaults.sourceNatDrainIps;
    	      this.sourceNatDrainRanges = defaults.sourceNatDrainRanges;
        }

        @CustomType.Setter
        public Builder sourceNatActiveIps(@Nullable List<String> sourceNatActiveIps) {

            this.sourceNatActiveIps = sourceNatActiveIps;
            return this;
        }
        public Builder sourceNatActiveIps(String... sourceNatActiveIps) {
            return sourceNatActiveIps(List.of(sourceNatActiveIps));
        }
        @CustomType.Setter
        public Builder sourceNatActiveRanges(@Nullable List<String> sourceNatActiveRanges) {

            this.sourceNatActiveRanges = sourceNatActiveRanges;
            return this;
        }
        public Builder sourceNatActiveRanges(String... sourceNatActiveRanges) {
            return sourceNatActiveRanges(List.of(sourceNatActiveRanges));
        }
        @CustomType.Setter
        public Builder sourceNatDrainIps(@Nullable List<String> sourceNatDrainIps) {

            this.sourceNatDrainIps = sourceNatDrainIps;
            return this;
        }
        public Builder sourceNatDrainIps(String... sourceNatDrainIps) {
            return sourceNatDrainIps(List.of(sourceNatDrainIps));
        }
        @CustomType.Setter
        public Builder sourceNatDrainRanges(@Nullable List<String> sourceNatDrainRanges) {

            this.sourceNatDrainRanges = sourceNatDrainRanges;
            return this;
        }
        public Builder sourceNatDrainRanges(String... sourceNatDrainRanges) {
            return sourceNatDrainRanges(List.of(sourceNatDrainRanges));
        }
        public RouterNatRuleAction build() {
            final var _resultValue = new RouterNatRuleAction();
            _resultValue.sourceNatActiveIps = sourceNatActiveIps;
            _resultValue.sourceNatActiveRanges = sourceNatActiveRanges;
            _resultValue.sourceNatDrainIps = sourceNatDrainIps;
            _resultValue.sourceNatDrainRanges = sourceNatDrainRanges;
            return _resultValue;
        }
    }
}
