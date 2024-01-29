// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouterNatRuleActionArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouterNatRuleActionArgs Empty = new RouterNatRuleActionArgs();

    /**
     * A list of URLs of the IP resources used for this NAT rule.
     * These IP addresses must be valid static external IP addresses assigned to the project.
     * This field is used for public NAT.
     * 
     */
    @Import(name="sourceNatActiveIps")
    private @Nullable Output<List<String>> sourceNatActiveIps;

    /**
     * @return A list of URLs of the IP resources used for this NAT rule.
     * These IP addresses must be valid static external IP addresses assigned to the project.
     * This field is used for public NAT.
     * 
     */
    public Optional<Output<List<String>>> sourceNatActiveIps() {
        return Optional.ofNullable(this.sourceNatActiveIps);
    }

    /**
     * A list of URLs of the subnetworks used as source ranges for this NAT Rule.
     * These subnetworks must have purpose set to PRIVATE_NAT.
     * This field is used for private NAT.
     * 
     */
    @Import(name="sourceNatActiveRanges")
    private @Nullable Output<List<String>> sourceNatActiveRanges;

    /**
     * @return A list of URLs of the subnetworks used as source ranges for this NAT Rule.
     * These subnetworks must have purpose set to PRIVATE_NAT.
     * This field is used for private NAT.
     * 
     */
    public Optional<Output<List<String>>> sourceNatActiveRanges() {
        return Optional.ofNullable(this.sourceNatActiveRanges);
    }

    /**
     * A list of URLs of the IP resources to be drained.
     * These IPs must be valid static external IPs that have been assigned to the NAT.
     * These IPs should be used for updating/patching a NAT rule only.
     * This field is used for public NAT.
     * 
     */
    @Import(name="sourceNatDrainIps")
    private @Nullable Output<List<String>> sourceNatDrainIps;

    /**
     * @return A list of URLs of the IP resources to be drained.
     * These IPs must be valid static external IPs that have been assigned to the NAT.
     * These IPs should be used for updating/patching a NAT rule only.
     * This field is used for public NAT.
     * 
     */
    public Optional<Output<List<String>>> sourceNatDrainIps() {
        return Optional.ofNullable(this.sourceNatDrainIps);
    }

    /**
     * A list of URLs of subnetworks representing source ranges to be drained.
     * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
     * This field is used for private NAT.
     * 
     */
    @Import(name="sourceNatDrainRanges")
    private @Nullable Output<List<String>> sourceNatDrainRanges;

    /**
     * @return A list of URLs of subnetworks representing source ranges to be drained.
     * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
     * This field is used for private NAT.
     * 
     */
    public Optional<Output<List<String>>> sourceNatDrainRanges() {
        return Optional.ofNullable(this.sourceNatDrainRanges);
    }

    private RouterNatRuleActionArgs() {}

    private RouterNatRuleActionArgs(RouterNatRuleActionArgs $) {
        this.sourceNatActiveIps = $.sourceNatActiveIps;
        this.sourceNatActiveRanges = $.sourceNatActiveRanges;
        this.sourceNatDrainIps = $.sourceNatDrainIps;
        this.sourceNatDrainRanges = $.sourceNatDrainRanges;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouterNatRuleActionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouterNatRuleActionArgs $;

        public Builder() {
            $ = new RouterNatRuleActionArgs();
        }

        public Builder(RouterNatRuleActionArgs defaults) {
            $ = new RouterNatRuleActionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param sourceNatActiveIps A list of URLs of the IP resources used for this NAT rule.
         * These IP addresses must be valid static external IP addresses assigned to the project.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveIps(@Nullable Output<List<String>> sourceNatActiveIps) {
            $.sourceNatActiveIps = sourceNatActiveIps;
            return this;
        }

        /**
         * @param sourceNatActiveIps A list of URLs of the IP resources used for this NAT rule.
         * These IP addresses must be valid static external IP addresses assigned to the project.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveIps(List<String> sourceNatActiveIps) {
            return sourceNatActiveIps(Output.of(sourceNatActiveIps));
        }

        /**
         * @param sourceNatActiveIps A list of URLs of the IP resources used for this NAT rule.
         * These IP addresses must be valid static external IP addresses assigned to the project.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveIps(String... sourceNatActiveIps) {
            return sourceNatActiveIps(List.of(sourceNatActiveIps));
        }

        /**
         * @param sourceNatActiveRanges A list of URLs of the subnetworks used as source ranges for this NAT Rule.
         * These subnetworks must have purpose set to PRIVATE_NAT.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveRanges(@Nullable Output<List<String>> sourceNatActiveRanges) {
            $.sourceNatActiveRanges = sourceNatActiveRanges;
            return this;
        }

        /**
         * @param sourceNatActiveRanges A list of URLs of the subnetworks used as source ranges for this NAT Rule.
         * These subnetworks must have purpose set to PRIVATE_NAT.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveRanges(List<String> sourceNatActiveRanges) {
            return sourceNatActiveRanges(Output.of(sourceNatActiveRanges));
        }

        /**
         * @param sourceNatActiveRanges A list of URLs of the subnetworks used as source ranges for this NAT Rule.
         * These subnetworks must have purpose set to PRIVATE_NAT.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatActiveRanges(String... sourceNatActiveRanges) {
            return sourceNatActiveRanges(List.of(sourceNatActiveRanges));
        }

        /**
         * @param sourceNatDrainIps A list of URLs of the IP resources to be drained.
         * These IPs must be valid static external IPs that have been assigned to the NAT.
         * These IPs should be used for updating/patching a NAT rule only.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainIps(@Nullable Output<List<String>> sourceNatDrainIps) {
            $.sourceNatDrainIps = sourceNatDrainIps;
            return this;
        }

        /**
         * @param sourceNatDrainIps A list of URLs of the IP resources to be drained.
         * These IPs must be valid static external IPs that have been assigned to the NAT.
         * These IPs should be used for updating/patching a NAT rule only.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainIps(List<String> sourceNatDrainIps) {
            return sourceNatDrainIps(Output.of(sourceNatDrainIps));
        }

        /**
         * @param sourceNatDrainIps A list of URLs of the IP resources to be drained.
         * These IPs must be valid static external IPs that have been assigned to the NAT.
         * These IPs should be used for updating/patching a NAT rule only.
         * This field is used for public NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainIps(String... sourceNatDrainIps) {
            return sourceNatDrainIps(List.of(sourceNatDrainIps));
        }

        /**
         * @param sourceNatDrainRanges A list of URLs of subnetworks representing source ranges to be drained.
         * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainRanges(@Nullable Output<List<String>> sourceNatDrainRanges) {
            $.sourceNatDrainRanges = sourceNatDrainRanges;
            return this;
        }

        /**
         * @param sourceNatDrainRanges A list of URLs of subnetworks representing source ranges to be drained.
         * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainRanges(List<String> sourceNatDrainRanges) {
            return sourceNatDrainRanges(Output.of(sourceNatDrainRanges));
        }

        /**
         * @param sourceNatDrainRanges A list of URLs of subnetworks representing source ranges to be drained.
         * This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
         * This field is used for private NAT.
         * 
         * @return builder
         * 
         */
        public Builder sourceNatDrainRanges(String... sourceNatDrainRanges) {
            return sourceNatDrainRanges(List.of(sourceNatDrainRanges));
        }

        public RouterNatRuleActionArgs build() {
            return $;
        }
    }

}
