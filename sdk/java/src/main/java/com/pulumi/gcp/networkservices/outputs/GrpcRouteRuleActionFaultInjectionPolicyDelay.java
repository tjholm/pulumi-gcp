// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GrpcRouteRuleActionFaultInjectionPolicyDelay {
    /**
     * @return Specify a fixed delay before forwarding the request.
     * 
     */
    private @Nullable String fixedDelay;
    /**
     * @return The percentage of traffic on which delay will be injected.
     * 
     */
    private @Nullable Integer percentage;

    private GrpcRouteRuleActionFaultInjectionPolicyDelay() {}
    /**
     * @return Specify a fixed delay before forwarding the request.
     * 
     */
    public Optional<String> fixedDelay() {
        return Optional.ofNullable(this.fixedDelay);
    }
    /**
     * @return The percentage of traffic on which delay will be injected.
     * 
     */
    public Optional<Integer> percentage() {
        return Optional.ofNullable(this.percentage);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GrpcRouteRuleActionFaultInjectionPolicyDelay defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String fixedDelay;
        private @Nullable Integer percentage;
        public Builder() {}
        public Builder(GrpcRouteRuleActionFaultInjectionPolicyDelay defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fixedDelay = defaults.fixedDelay;
    	      this.percentage = defaults.percentage;
        }

        @CustomType.Setter
        public Builder fixedDelay(@Nullable String fixedDelay) {
            this.fixedDelay = fixedDelay;
            return this;
        }
        @CustomType.Setter
        public Builder percentage(@Nullable Integer percentage) {
            this.percentage = percentage;
            return this;
        }
        public GrpcRouteRuleActionFaultInjectionPolicyDelay build() {
            final var o = new GrpcRouteRuleActionFaultInjectionPolicyDelay();
            o.fixedDelay = fixedDelay;
            o.percentage = percentage;
            return o;
        }
    }
}
