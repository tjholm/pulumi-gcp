// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.storage.outputs.TransferJobScheduleScheduleEndDate;
import com.pulumi.gcp.storage.outputs.TransferJobScheduleScheduleStartDate;
import com.pulumi.gcp.storage.outputs.TransferJobScheduleStartTimeOfDay;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TransferJobSchedule {
    /**
     * @return Interval between the start of each scheduled transfer. If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    private final @Nullable String repeatInterval;
    /**
     * @return The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure documented below.
     * 
     */
    private final @Nullable TransferJobScheduleScheduleEndDate scheduleEndDate;
    /**
     * @return The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure documented below.
     * 
     */
    private final TransferJobScheduleScheduleStartDate scheduleStartDate;
    /**
     * @return The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer&#39;s start time in a day is specified in your local timezone. Structure documented below.
     * 
     */
    private final @Nullable TransferJobScheduleStartTimeOfDay startTimeOfDay;

    @CustomType.Constructor
    private TransferJobSchedule(
        @CustomType.Parameter("repeatInterval") @Nullable String repeatInterval,
        @CustomType.Parameter("scheduleEndDate") @Nullable TransferJobScheduleScheduleEndDate scheduleEndDate,
        @CustomType.Parameter("scheduleStartDate") TransferJobScheduleScheduleStartDate scheduleStartDate,
        @CustomType.Parameter("startTimeOfDay") @Nullable TransferJobScheduleStartTimeOfDay startTimeOfDay) {
        this.repeatInterval = repeatInterval;
        this.scheduleEndDate = scheduleEndDate;
        this.scheduleStartDate = scheduleStartDate;
        this.startTimeOfDay = startTimeOfDay;
    }

    /**
     * @return Interval between the start of each scheduled transfer. If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<String> repeatInterval() {
        return Optional.ofNullable(this.repeatInterval);
    }
    /**
     * @return The last day the recurring transfer will be run. If `schedule_end_date` is the same as `schedule_start_date`, the transfer will be executed only once. Structure documented below.
     * 
     */
    public Optional<TransferJobScheduleScheduleEndDate> scheduleEndDate() {
        return Optional.ofNullable(this.scheduleEndDate);
    }
    /**
     * @return The first day the recurring transfer is scheduled to run. If `schedule_start_date` is in the past, the transfer will run for the first time on the following day. Structure documented below.
     * 
     */
    public TransferJobScheduleScheduleStartDate scheduleStartDate() {
        return this.scheduleStartDate;
    }
    /**
     * @return The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer&#39;s start time in a day is specified in your local timezone. Structure documented below.
     * 
     */
    public Optional<TransferJobScheduleStartTimeOfDay> startTimeOfDay() {
        return Optional.ofNullable(this.startTimeOfDay);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TransferJobSchedule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String repeatInterval;
        private @Nullable TransferJobScheduleScheduleEndDate scheduleEndDate;
        private TransferJobScheduleScheduleStartDate scheduleStartDate;
        private @Nullable TransferJobScheduleStartTimeOfDay startTimeOfDay;

        public Builder() {
    	      // Empty
        }

        public Builder(TransferJobSchedule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.repeatInterval = defaults.repeatInterval;
    	      this.scheduleEndDate = defaults.scheduleEndDate;
    	      this.scheduleStartDate = defaults.scheduleStartDate;
    	      this.startTimeOfDay = defaults.startTimeOfDay;
        }

        public Builder repeatInterval(@Nullable String repeatInterval) {
            this.repeatInterval = repeatInterval;
            return this;
        }
        public Builder scheduleEndDate(@Nullable TransferJobScheduleScheduleEndDate scheduleEndDate) {
            this.scheduleEndDate = scheduleEndDate;
            return this;
        }
        public Builder scheduleStartDate(TransferJobScheduleScheduleStartDate scheduleStartDate) {
            this.scheduleStartDate = Objects.requireNonNull(scheduleStartDate);
            return this;
        }
        public Builder startTimeOfDay(@Nullable TransferJobScheduleStartTimeOfDay startTimeOfDay) {
            this.startTimeOfDay = startTimeOfDay;
            return this;
        }        public TransferJobSchedule build() {
            return new TransferJobSchedule(repeatInterval, scheduleEndDate, scheduleStartDate, startTimeOfDay);
        }
    }
}
