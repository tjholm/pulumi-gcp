// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class InstanceFromTemplateNetworkPerformanceConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceFromTemplateNetworkPerformanceConfigArgs Empty = new InstanceFromTemplateNetworkPerformanceConfigArgs();

    /**
     * The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT
     * 
     */
    @Import(name="totalEgressBandwidthTier", required=true)
    private Output<String> totalEgressBandwidthTier;

    /**
     * @return The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT
     * 
     */
    public Output<String> totalEgressBandwidthTier() {
        return this.totalEgressBandwidthTier;
    }

    private InstanceFromTemplateNetworkPerformanceConfigArgs() {}

    private InstanceFromTemplateNetworkPerformanceConfigArgs(InstanceFromTemplateNetworkPerformanceConfigArgs $) {
        this.totalEgressBandwidthTier = $.totalEgressBandwidthTier;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceFromTemplateNetworkPerformanceConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceFromTemplateNetworkPerformanceConfigArgs $;

        public Builder() {
            $ = new InstanceFromTemplateNetworkPerformanceConfigArgs();
        }

        public Builder(InstanceFromTemplateNetworkPerformanceConfigArgs defaults) {
            $ = new InstanceFromTemplateNetworkPerformanceConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param totalEgressBandwidthTier The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT
         * 
         * @return builder
         * 
         */
        public Builder totalEgressBandwidthTier(Output<String> totalEgressBandwidthTier) {
            $.totalEgressBandwidthTier = totalEgressBandwidthTier;
            return this;
        }

        /**
         * @param totalEgressBandwidthTier The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT
         * 
         * @return builder
         * 
         */
        public Builder totalEgressBandwidthTier(String totalEgressBandwidthTier) {
            return totalEgressBandwidthTier(Output.of(totalEgressBandwidthTier));
        }

        public InstanceFromTemplateNetworkPerformanceConfigArgs build() {
            if ($.totalEgressBandwidthTier == null) {
                throw new MissingRequiredPropertyException("InstanceFromTemplateNetworkPerformanceConfigArgs", "totalEgressBandwidthTier");
            }
            return $;
        }
    }

}
