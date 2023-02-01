// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegionAutoscalerAutoscalingPolicyMetric {
    private @Nullable String filter;
    /**
     * @return The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     * 
     */
    private String name;
    private @Nullable Double singleInstanceAssignment;
    /**
     * @return The target value of the metric that autoscaler should
     * maintain. This must be a positive value. A utilization
     * metric scales number of virtual machines handling requests
     * to increase or decrease proportionally to the metric.
     * For example, a good metric to use as a utilizationTarget is
     * www.googleapis.com/compute/instance/network/received_bytes_count.
     * The autoscaler will work to keep this value constant for each
     * of the instances.
     * 
     */
    private @Nullable Double target;
    /**
     * @return Defines how target utilization value is expressed for a
     * Stackdriver Monitoring metric.
     * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
     * 
     */
    private @Nullable String type;

    private RegionAutoscalerAutoscalingPolicyMetric() {}
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }
    /**
     * @return The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     * 
     */
    public String name() {
        return this.name;
    }
    public Optional<Double> singleInstanceAssignment() {
        return Optional.ofNullable(this.singleInstanceAssignment);
    }
    /**
     * @return The target value of the metric that autoscaler should
     * maintain. This must be a positive value. A utilization
     * metric scales number of virtual machines handling requests
     * to increase or decrease proportionally to the metric.
     * For example, a good metric to use as a utilizationTarget is
     * www.googleapis.com/compute/instance/network/received_bytes_count.
     * The autoscaler will work to keep this value constant for each
     * of the instances.
     * 
     */
    public Optional<Double> target() {
        return Optional.ofNullable(this.target);
    }
    /**
     * @return Defines how target utilization value is expressed for a
     * Stackdriver Monitoring metric.
     * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionAutoscalerAutoscalingPolicyMetric defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String filter;
        private String name;
        private @Nullable Double singleInstanceAssignment;
        private @Nullable Double target;
        private @Nullable String type;
        public Builder() {}
        public Builder(RegionAutoscalerAutoscalingPolicyMetric defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filter = defaults.filter;
    	      this.name = defaults.name;
    	      this.singleInstanceAssignment = defaults.singleInstanceAssignment;
    	      this.target = defaults.target;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder filter(@Nullable String filter) {
            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder singleInstanceAssignment(@Nullable Double singleInstanceAssignment) {
            this.singleInstanceAssignment = singleInstanceAssignment;
            return this;
        }
        @CustomType.Setter
        public Builder target(@Nullable Double target) {
            this.target = target;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public RegionAutoscalerAutoscalingPolicyMetric build() {
            final var o = new RegionAutoscalerAutoscalingPolicyMetric();
            o.filter = filter;
            o.name = name;
            o.singleInstanceAssignment = singleInstanceAssignment;
            o.target = target;
            o.type = type;
            return o;
        }
    }
}
