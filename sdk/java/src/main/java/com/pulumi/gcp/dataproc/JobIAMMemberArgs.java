// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dataproc.inputs.JobIAMMemberConditionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class JobIAMMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final JobIAMMemberArgs Empty = new JobIAMMemberArgs();

    @Import(name="condition")
    private @Nullable Output<JobIAMMemberConditionArgs> condition;

    public Optional<Output<JobIAMMemberConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    @Import(name="jobId", required=true)
    private Output<String> jobId;

    public Output<String> jobId() {
        return this.jobId;
    }

    @Import(name="member", required=true)
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }

    /**
     * The project in which the job belongs. If it
     * is not provided, the provider will use a default.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project in which the job belongs. If it
     * is not provided, the provider will use a default.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The region in which the job belongs. If it
     * is not provided, the provider will use a default.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which the job belongs. If it
     * is not provided, the provider will use a default.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The role that should be applied. Only one
     * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private JobIAMMemberArgs() {}

    private JobIAMMemberArgs(JobIAMMemberArgs $) {
        this.condition = $.condition;
        this.jobId = $.jobId;
        this.member = $.member;
        this.project = $.project;
        this.region = $.region;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JobIAMMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JobIAMMemberArgs $;

        public Builder() {
            $ = new JobIAMMemberArgs();
        }

        public Builder(JobIAMMemberArgs defaults) {
            $ = new JobIAMMemberArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<JobIAMMemberConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(JobIAMMemberConditionArgs condition) {
            return condition(Output.of(condition));
        }

        public Builder jobId(Output<String> jobId) {
            $.jobId = jobId;
            return this;
        }

        public Builder jobId(String jobId) {
            return jobId(Output.of(jobId));
        }

        public Builder member(Output<String> member) {
            $.member = member;
            return this;
        }

        public Builder member(String member) {
            return member(Output.of(member));
        }

        /**
         * @param project The project in which the job belongs. If it
         * is not provided, the provider will use a default.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project in which the job belongs. If it
         * is not provided, the provider will use a default.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param region The region in which the job belongs. If it
         * is not provided, the provider will use a default.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which the job belongs. If it
         * is not provided, the provider will use a default.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public JobIAMMemberArgs build() {
            $.jobId = Objects.requireNonNull($.jobId, "expected parameter 'jobId' to be non-null");
            $.member = Objects.requireNonNull($.member, "expected parameter 'member' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
