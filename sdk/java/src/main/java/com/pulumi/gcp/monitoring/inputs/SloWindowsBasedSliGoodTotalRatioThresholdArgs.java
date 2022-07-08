// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.monitoring.inputs.SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceArgs;
import com.pulumi.gcp.monitoring.inputs.SloWindowsBasedSliGoodTotalRatioThresholdPerformanceArgs;
import java.lang.Double;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SloWindowsBasedSliGoodTotalRatioThresholdArgs extends com.pulumi.resources.ResourceArgs {

    public static final SloWindowsBasedSliGoodTotalRatioThresholdArgs Empty = new SloWindowsBasedSliGoodTotalRatioThresholdArgs();

    /**
     * Basic SLI to evaluate to judge window quality.
     * Structure is documented below.
     * 
     */
    @Import(name="basicSliPerformance")
    private @Nullable Output<SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceArgs> basicSliPerformance;

    /**
     * @return Basic SLI to evaluate to judge window quality.
     * Structure is documented below.
     * 
     */
    public Optional<Output<SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceArgs>> basicSliPerformance() {
        return Optional.ofNullable(this.basicSliPerformance);
    }

    /**
     * Request-based SLI to evaluate to judge window quality.
     * Structure is documented below.
     * 
     */
    @Import(name="performance")
    private @Nullable Output<SloWindowsBasedSliGoodTotalRatioThresholdPerformanceArgs> performance;

    /**
     * @return Request-based SLI to evaluate to judge window quality.
     * Structure is documented below.
     * 
     */
    public Optional<Output<SloWindowsBasedSliGoodTotalRatioThresholdPerformanceArgs>> performance() {
        return Optional.ofNullable(this.performance);
    }

    /**
     * A duration string, e.g. 10s.
     * Good service is defined to be the count of requests made to
     * this service that return in no more than threshold.
     * 
     */
    @Import(name="threshold")
    private @Nullable Output<Double> threshold;

    /**
     * @return A duration string, e.g. 10s.
     * Good service is defined to be the count of requests made to
     * this service that return in no more than threshold.
     * 
     */
    public Optional<Output<Double>> threshold() {
        return Optional.ofNullable(this.threshold);
    }

    private SloWindowsBasedSliGoodTotalRatioThresholdArgs() {}

    private SloWindowsBasedSliGoodTotalRatioThresholdArgs(SloWindowsBasedSliGoodTotalRatioThresholdArgs $) {
        this.basicSliPerformance = $.basicSliPerformance;
        this.performance = $.performance;
        this.threshold = $.threshold;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SloWindowsBasedSliGoodTotalRatioThresholdArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SloWindowsBasedSliGoodTotalRatioThresholdArgs $;

        public Builder() {
            $ = new SloWindowsBasedSliGoodTotalRatioThresholdArgs();
        }

        public Builder(SloWindowsBasedSliGoodTotalRatioThresholdArgs defaults) {
            $ = new SloWindowsBasedSliGoodTotalRatioThresholdArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param basicSliPerformance Basic SLI to evaluate to judge window quality.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder basicSliPerformance(@Nullable Output<SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceArgs> basicSliPerformance) {
            $.basicSliPerformance = basicSliPerformance;
            return this;
        }

        /**
         * @param basicSliPerformance Basic SLI to evaluate to judge window quality.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder basicSliPerformance(SloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceArgs basicSliPerformance) {
            return basicSliPerformance(Output.of(basicSliPerformance));
        }

        /**
         * @param performance Request-based SLI to evaluate to judge window quality.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder performance(@Nullable Output<SloWindowsBasedSliGoodTotalRatioThresholdPerformanceArgs> performance) {
            $.performance = performance;
            return this;
        }

        /**
         * @param performance Request-based SLI to evaluate to judge window quality.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder performance(SloWindowsBasedSliGoodTotalRatioThresholdPerformanceArgs performance) {
            return performance(Output.of(performance));
        }

        /**
         * @param threshold A duration string, e.g. 10s.
         * Good service is defined to be the count of requests made to
         * this service that return in no more than threshold.
         * 
         * @return builder
         * 
         */
        public Builder threshold(@Nullable Output<Double> threshold) {
            $.threshold = threshold;
            return this;
        }

        /**
         * @param threshold A duration string, e.g. 10s.
         * Good service is defined to be the count of requests made to
         * this service that return in no more than threshold.
         * 
         * @return builder
         * 
         */
        public Builder threshold(Double threshold) {
            return threshold(Output.of(threshold));
        }

        public SloWindowsBasedSliGoodTotalRatioThresholdArgs build() {
            return $;
        }
    }

}
