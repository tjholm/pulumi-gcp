// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AutoscalarAutoscalingPolicyMetricArgs extends com.pulumi.resources.ResourceArgs {

    public static final AutoscalarAutoscalingPolicyMetricArgs Empty = new AutoscalarAutoscalingPolicyMetricArgs();

    @Import(name="filter")
    private @Nullable Output<String> filter;

    public Optional<Output<String>> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The identifier (type) of the Stackdriver Monitoring metric.
     * The metric cannot have negative values.
     * The metric must have a value type of INT64 or DOUBLE.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    @Import(name="singleInstanceAssignment")
    private @Nullable Output<Double> singleInstanceAssignment;

    public Optional<Output<Double>> singleInstanceAssignment() {
        return Optional.ofNullable(this.singleInstanceAssignment);
    }

    /**
     * The target value of the metric that autoscaler should
     * maintain. This must be a positive value. A utilization
     * metric scales number of virtual machines handling requests
     * to increase or decrease proportionally to the metric.
     * For example, a good metric to use as a utilizationTarget is
     * www.googleapis.com/compute/instance/network/received_bytes_count.
     * The autoscaler will work to keep this value constant for each
     * of the instances.
     * 
     */
    @Import(name="target")
    private @Nullable Output<Double> target;

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
    public Optional<Output<Double>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * Defines how target utilization value is expressed for a
     * Stackdriver Monitoring metric.
     * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Defines how target utilization value is expressed for a
     * Stackdriver Monitoring metric.
     * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private AutoscalarAutoscalingPolicyMetricArgs() {}

    private AutoscalarAutoscalingPolicyMetricArgs(AutoscalarAutoscalingPolicyMetricArgs $) {
        this.filter = $.filter;
        this.name = $.name;
        this.singleInstanceAssignment = $.singleInstanceAssignment;
        this.target = $.target;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AutoscalarAutoscalingPolicyMetricArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AutoscalarAutoscalingPolicyMetricArgs $;

        public Builder() {
            $ = new AutoscalarAutoscalingPolicyMetricArgs();
        }

        public Builder(AutoscalarAutoscalingPolicyMetricArgs defaults) {
            $ = new AutoscalarAutoscalingPolicyMetricArgs(Objects.requireNonNull(defaults));
        }

        public Builder filter(@Nullable Output<String> filter) {
            $.filter = filter;
            return this;
        }

        public Builder filter(String filter) {
            return filter(Output.of(filter));
        }

        /**
         * @param name The identifier (type) of the Stackdriver Monitoring metric.
         * The metric cannot have negative values.
         * The metric must have a value type of INT64 or DOUBLE.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The identifier (type) of the Stackdriver Monitoring metric.
         * The metric cannot have negative values.
         * The metric must have a value type of INT64 or DOUBLE.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder singleInstanceAssignment(@Nullable Output<Double> singleInstanceAssignment) {
            $.singleInstanceAssignment = singleInstanceAssignment;
            return this;
        }

        public Builder singleInstanceAssignment(Double singleInstanceAssignment) {
            return singleInstanceAssignment(Output.of(singleInstanceAssignment));
        }

        /**
         * @param target The target value of the metric that autoscaler should
         * maintain. This must be a positive value. A utilization
         * metric scales number of virtual machines handling requests
         * to increase or decrease proportionally to the metric.
         * For example, a good metric to use as a utilizationTarget is
         * www.googleapis.com/compute/instance/network/received_bytes_count.
         * The autoscaler will work to keep this value constant for each
         * of the instances.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<Double> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target The target value of the metric that autoscaler should
         * maintain. This must be a positive value. A utilization
         * metric scales number of virtual machines handling requests
         * to increase or decrease proportionally to the metric.
         * For example, a good metric to use as a utilizationTarget is
         * www.googleapis.com/compute/instance/network/received_bytes_count.
         * The autoscaler will work to keep this value constant for each
         * of the instances.
         * 
         * @return builder
         * 
         */
        public Builder target(Double target) {
            return target(Output.of(target));
        }

        /**
         * @param type Defines how target utilization value is expressed for a
         * Stackdriver Monitoring metric.
         * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Defines how target utilization value is expressed for a
         * Stackdriver Monitoring metric.
         * Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AutoscalarAutoscalingPolicyMetricArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
