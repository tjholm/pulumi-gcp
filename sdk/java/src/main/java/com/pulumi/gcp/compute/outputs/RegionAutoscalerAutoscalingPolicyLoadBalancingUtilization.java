// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.util.Objects;

@CustomType
public final class RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization {
    /**
     * @return Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     * 
     */
    private final Double target;

    @CustomType.Constructor
    private RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(@CustomType.Parameter("target") Double target) {
        this.target = target;
    }

    /**
     * @return Fraction of backend capacity utilization (set in HTTP(s) load
     * balancing configuration) that autoscaler should maintain. Must
     * be a positive float value. If not defined, the default is 0.8.
     * 
     */
    public Double target() {
        return this.target;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Double target;

        public Builder() {
    	      // Empty
        }

        public Builder(RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.target = defaults.target;
        }

        public Builder target(Double target) {
            this.target = Objects.requireNonNull(target);
            return this;
        }        public RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization build() {
            return new RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(target);
        }
    }
}
