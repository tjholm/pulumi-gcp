// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs extends com.pulumi.resources.ResourceArgs {

    public static final AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs Empty = new AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs();

    /**
     * The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * (Optional, Beta, Deprecated)
     * Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
     * A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     * &gt; **Warning:** `monitoring_interval` is deprecated and will be removed in a future release.
     * 
     * @deprecated
     * `monitoring_interval` is deprecated and will be removed in a future release.
     * 
     */
    @Deprecated /* `monitoring_interval` is deprecated and will be removed in a future release. */
    @Import(name="monitoringInterval")
    private @Nullable Output<String> monitoringInterval;

    /**
     * @return (Optional, Beta, Deprecated)
     * Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
     * A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     * &gt; **Warning:** `monitoring_interval` is deprecated and will be removed in a future release.
     * 
     * @deprecated
     * `monitoring_interval` is deprecated and will be removed in a future release.
     * 
     */
    @Deprecated /* `monitoring_interval` is deprecated and will be removed in a future release. */
    public Optional<Output<String>> monitoringInterval() {
        return Optional.ofNullable(this.monitoringInterval);
    }

    /**
     * Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
     * If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.
     * 
     */
    @Import(name="monitoringIntervalDays")
    private @Nullable Output<Integer> monitoringIntervalDays;

    /**
     * @return Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
     * If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.
     * 
     */
    public Optional<Output<Integer>> monitoringIntervalDays() {
        return Optional.ofNullable(this.monitoringIntervalDays);
    }

    /**
     * Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.
     * 
     */
    @Import(name="stalenessDays")
    private @Nullable Output<Integer> stalenessDays;

    /**
     * @return Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.
     * 
     */
    public Optional<Output<Integer>> stalenessDays() {
        return Optional.ofNullable(this.stalenessDays);
    }

    private AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs() {}

    private AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs(AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs $) {
        this.disabled = $.disabled;
        this.monitoringInterval = $.monitoringInterval;
        this.monitoringIntervalDays = $.monitoringIntervalDays;
        this.stalenessDays = $.stalenessDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs $;

        public Builder() {
            $ = new AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs();
        }

        public Builder(AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs defaults) {
            $ = new AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disabled The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled The monitoring schedule for snapshot analysis. For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param monitoringInterval (Optional, Beta, Deprecated)
         * Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
         * A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * &gt; **Warning:** `monitoring_interval` is deprecated and will be removed in a future release.
         * 
         * @return builder
         * 
         * @deprecated
         * `monitoring_interval` is deprecated and will be removed in a future release.
         * 
         */
        @Deprecated /* `monitoring_interval` is deprecated and will be removed in a future release. */
        public Builder monitoringInterval(@Nullable Output<String> monitoringInterval) {
            $.monitoringInterval = monitoringInterval;
            return this;
        }

        /**
         * @param monitoringInterval (Optional, Beta, Deprecated)
         * Configuration of the snapshot analysis based monitoring pipeline running interval. The value is rolled up to full day.
         * A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * &gt; **Warning:** `monitoring_interval` is deprecated and will be removed in a future release.
         * 
         * @return builder
         * 
         * @deprecated
         * `monitoring_interval` is deprecated and will be removed in a future release.
         * 
         */
        @Deprecated /* `monitoring_interval` is deprecated and will be removed in a future release. */
        public Builder monitoringInterval(String monitoringInterval) {
            return monitoringInterval(Output.of(monitoringInterval));
        }

        /**
         * @param monitoringIntervalDays Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
         * If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.
         * 
         * @return builder
         * 
         */
        public Builder monitoringIntervalDays(@Nullable Output<Integer> monitoringIntervalDays) {
            $.monitoringIntervalDays = monitoringIntervalDays;
            return this;
        }

        /**
         * @param monitoringIntervalDays Configuration of the snapshot analysis based monitoring pipeline running interval. The value indicates number of days. The default value is 1.
         * If both FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days and [FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval][] are set when creating/updating EntityTypes/Features, FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days will be used.
         * 
         * @return builder
         * 
         */
        public Builder monitoringIntervalDays(Integer monitoringIntervalDays) {
            return monitoringIntervalDays(Output.of(monitoringIntervalDays));
        }

        /**
         * @param stalenessDays Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.
         * 
         * @return builder
         * 
         */
        public Builder stalenessDays(@Nullable Output<Integer> stalenessDays) {
            $.stalenessDays = stalenessDays;
            return this;
        }

        /**
         * @param stalenessDays Customized export features time window for snapshot analysis. Unit is one day. The default value is 21 days. Minimum value is 1 day. Maximum value is 4000 days.
         * 
         * @return builder
         * 
         */
        public Builder stalenessDays(Integer stalenessDays) {
            return stalenessDays(Output.of(stalenessDays));
        }

        public AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs build() {
            return $;
        }
    }

}
