// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.storage.inputs.TransferJobScheduleArgs;
import com.pulumi.gcp.storage.inputs.TransferJobTransferSpecArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransferJobState extends com.pulumi.resources.ResourceArgs {

    public static final TransferJobState Empty = new TransferJobState();

    /**
     * When the Transfer Job was created.
     * 
     */
    @Import(name="creationTime")
    private @Nullable Output<String> creationTime;

    /**
     * @return When the Transfer Job was created.
     * 
     */
    public Optional<Output<String>> creationTime() {
        return Optional.ofNullable(this.creationTime);
    }

    /**
     * When the Transfer Job was deleted.
     * 
     */
    @Import(name="deletionTime")
    private @Nullable Output<String> deletionTime;

    /**
     * @return When the Transfer Job was deleted.
     * 
     */
    public Optional<Output<String>> deletionTime() {
        return Optional.ofNullable(this.deletionTime);
    }

    /**
     * Unique description to identify the Transfer Job.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Unique description to identify the Transfer Job.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When the Transfer Job was last modified.
     * 
     */
    @Import(name="lastModificationTime")
    private @Nullable Output<String> lastModificationTime;

    /**
     * @return When the Transfer Job was last modified.
     * 
     */
    public Optional<Output<String>> lastModificationTime() {
        return Optional.ofNullable(this.lastModificationTime);
    }

    /**
     * The name of the Transfer Job.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Transfer Job.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<TransferJobScheduleArgs> schedule;

    /**
     * @return Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
     * 
     */
    public Optional<Output<TransferJobScheduleArgs>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Transfer specification. Structure documented below.
     * 
     */
    @Import(name="transferSpec")
    private @Nullable Output<TransferJobTransferSpecArgs> transferSpec;

    /**
     * @return Transfer specification. Structure documented below.
     * 
     */
    public Optional<Output<TransferJobTransferSpecArgs>> transferSpec() {
        return Optional.ofNullable(this.transferSpec);
    }

    private TransferJobState() {}

    private TransferJobState(TransferJobState $) {
        this.creationTime = $.creationTime;
        this.deletionTime = $.deletionTime;
        this.description = $.description;
        this.lastModificationTime = $.lastModificationTime;
        this.name = $.name;
        this.project = $.project;
        this.schedule = $.schedule;
        this.status = $.status;
        this.transferSpec = $.transferSpec;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransferJobState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransferJobState $;

        public Builder() {
            $ = new TransferJobState();
        }

        public Builder(TransferJobState defaults) {
            $ = new TransferJobState(Objects.requireNonNull(defaults));
        }

        /**
         * @param creationTime When the Transfer Job was created.
         * 
         * @return builder
         * 
         */
        public Builder creationTime(@Nullable Output<String> creationTime) {
            $.creationTime = creationTime;
            return this;
        }

        /**
         * @param creationTime When the Transfer Job was created.
         * 
         * @return builder
         * 
         */
        public Builder creationTime(String creationTime) {
            return creationTime(Output.of(creationTime));
        }

        /**
         * @param deletionTime When the Transfer Job was deleted.
         * 
         * @return builder
         * 
         */
        public Builder deletionTime(@Nullable Output<String> deletionTime) {
            $.deletionTime = deletionTime;
            return this;
        }

        /**
         * @param deletionTime When the Transfer Job was deleted.
         * 
         * @return builder
         * 
         */
        public Builder deletionTime(String deletionTime) {
            return deletionTime(Output.of(deletionTime));
        }

        /**
         * @param description Unique description to identify the Transfer Job.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Unique description to identify the Transfer Job.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param lastModificationTime When the Transfer Job was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModificationTime(@Nullable Output<String> lastModificationTime) {
            $.lastModificationTime = lastModificationTime;
            return this;
        }

        /**
         * @param lastModificationTime When the Transfer Job was last modified.
         * 
         * @return builder
         * 
         */
        public Builder lastModificationTime(String lastModificationTime) {
            return lastModificationTime(Output.of(lastModificationTime));
        }

        /**
         * @param name The name of the Transfer Job.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Transfer Job.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param schedule Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<TransferJobScheduleArgs> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder schedule(TransferJobScheduleArgs schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param status Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param transferSpec Transfer specification. Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder transferSpec(@Nullable Output<TransferJobTransferSpecArgs> transferSpec) {
            $.transferSpec = transferSpec;
            return this;
        }

        /**
         * @param transferSpec Transfer specification. Structure documented below.
         * 
         * @return builder
         * 
         */
        public Builder transferSpec(TransferJobTransferSpecArgs transferSpec) {
            return transferSpec(Output.of(transferSpec));
        }

        public TransferJobState build() {
            return $;
        }
    }

}
