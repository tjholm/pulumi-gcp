// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.alloydb.inputs.ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterAutomatedBackupPolicyWeeklyScheduleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterAutomatedBackupPolicyWeeklyScheduleArgs Empty = new ClusterAutomatedBackupPolicyWeeklyScheduleArgs();

    /**
     * The days of the week to perform a backup. At least one day of the week must be provided.
     * Each value may be one of `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, and `SUNDAY`.
     * 
     */
    @Import(name="daysOfWeeks")
    private @Nullable Output<List<String>> daysOfWeeks;

    /**
     * @return The days of the week to perform a backup. At least one day of the week must be provided.
     * Each value may be one of `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, and `SUNDAY`.
     * 
     */
    public Optional<Output<List<String>>> daysOfWeeks() {
        return Optional.ofNullable(this.daysOfWeeks);
    }

    /**
     * The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).
     * Structure is documented below.
     * 
     */
    @Import(name="startTimes", required=true)
    private Output<List<ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs>> startTimes;

    /**
     * @return The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).
     * Structure is documented below.
     * 
     */
    public Output<List<ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs>> startTimes() {
        return this.startTimes;
    }

    private ClusterAutomatedBackupPolicyWeeklyScheduleArgs() {}

    private ClusterAutomatedBackupPolicyWeeklyScheduleArgs(ClusterAutomatedBackupPolicyWeeklyScheduleArgs $) {
        this.daysOfWeeks = $.daysOfWeeks;
        this.startTimes = $.startTimes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterAutomatedBackupPolicyWeeklyScheduleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterAutomatedBackupPolicyWeeklyScheduleArgs $;

        public Builder() {
            $ = new ClusterAutomatedBackupPolicyWeeklyScheduleArgs();
        }

        public Builder(ClusterAutomatedBackupPolicyWeeklyScheduleArgs defaults) {
            $ = new ClusterAutomatedBackupPolicyWeeklyScheduleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param daysOfWeeks The days of the week to perform a backup. At least one day of the week must be provided.
         * Each value may be one of `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, and `SUNDAY`.
         * 
         * @return builder
         * 
         */
        public Builder daysOfWeeks(@Nullable Output<List<String>> daysOfWeeks) {
            $.daysOfWeeks = daysOfWeeks;
            return this;
        }

        /**
         * @param daysOfWeeks The days of the week to perform a backup. At least one day of the week must be provided.
         * Each value may be one of `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, and `SUNDAY`.
         * 
         * @return builder
         * 
         */
        public Builder daysOfWeeks(List<String> daysOfWeeks) {
            return daysOfWeeks(Output.of(daysOfWeeks));
        }

        /**
         * @param daysOfWeeks The days of the week to perform a backup. At least one day of the week must be provided.
         * Each value may be one of `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, and `SUNDAY`.
         * 
         * @return builder
         * 
         */
        public Builder daysOfWeeks(String... daysOfWeeks) {
            return daysOfWeeks(List.of(daysOfWeeks));
        }

        /**
         * @param startTimes The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder startTimes(Output<List<ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs>> startTimes) {
            $.startTimes = startTimes;
            return this;
        }

        /**
         * @param startTimes The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder startTimes(List<ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs> startTimes) {
            return startTimes(Output.of(startTimes));
        }

        /**
         * @param startTimes The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder startTimes(ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs... startTimes) {
            return startTimes(List.of(startTimes));
        }

        public ClusterAutomatedBackupPolicyWeeklyScheduleArgs build() {
            $.startTimes = Objects.requireNonNull($.startTimes, "expected parameter 'startTimes' to be non-null");
            return $;
        }
    }

}
