// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReservationAssignmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReservationAssignmentArgs Empty = new ReservationAssignmentArgs();

    /**
     * The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
     * 
     */
    @Import(name="assignee", required=true)
    private Output<String> assignee;

    /**
     * @return The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
     * 
     */
    public Output<String> assignee() {
        return this.assignee;
    }

    /**
     * Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
     * 
     */
    @Import(name="jobType", required=true)
    private Output<String> jobType;

    /**
     * @return Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
     * 
     */
    public Output<String> jobType() {
        return this.jobType;
    }

    /**
     * The location for the resource
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The project for the resource
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The reservation for the resource
     * 
     */
    @Import(name="reservation", required=true)
    private Output<String> reservation;

    /**
     * @return The reservation for the resource
     * 
     */
    public Output<String> reservation() {
        return this.reservation;
    }

    private ReservationAssignmentArgs() {}

    private ReservationAssignmentArgs(ReservationAssignmentArgs $) {
        this.assignee = $.assignee;
        this.jobType = $.jobType;
        this.location = $.location;
        this.project = $.project;
        this.reservation = $.reservation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReservationAssignmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReservationAssignmentArgs $;

        public Builder() {
            $ = new ReservationAssignmentArgs();
        }

        public Builder(ReservationAssignmentArgs defaults) {
            $ = new ReservationAssignmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assignee The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
         * 
         * @return builder
         * 
         */
        public Builder assignee(Output<String> assignee) {
            $.assignee = assignee;
            return this;
        }

        /**
         * @param assignee The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
         * 
         * @return builder
         * 
         */
        public Builder assignee(String assignee) {
            return assignee(Output.of(assignee));
        }

        /**
         * @param jobType Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
         * 
         * @return builder
         * 
         */
        public Builder jobType(Output<String> jobType) {
            $.jobType = jobType;
            return this;
        }

        /**
         * @param jobType Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
         * 
         * @return builder
         * 
         */
        public Builder jobType(String jobType) {
            return jobType(Output.of(jobType));
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param reservation The reservation for the resource
         * 
         * @return builder
         * 
         */
        public Builder reservation(Output<String> reservation) {
            $.reservation = reservation;
            return this;
        }

        /**
         * @param reservation The reservation for the resource
         * 
         * @return builder
         * 
         */
        public Builder reservation(String reservation) {
            return reservation(Output.of(reservation));
        }

        public ReservationAssignmentArgs build() {
            $.assignee = Objects.requireNonNull($.assignee, "expected parameter 'assignee' to be non-null");
            $.jobType = Objects.requireNonNull($.jobType, "expected parameter 'jobType' to be non-null");
            $.reservation = Objects.requireNonNull($.reservation, "expected parameter 'reservation' to be non-null");
            return $;
        }
    }

}
