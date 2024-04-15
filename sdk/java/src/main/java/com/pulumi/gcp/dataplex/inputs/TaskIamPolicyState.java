// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TaskIamPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final TaskIamPolicyState Empty = new TaskIamPolicyState();

    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The lake in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="lake")
    private @Nullable Output<String> lake;

    /**
     * @return The lake in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Optional<Output<String>> lake() {
        return Optional.ofNullable(this.lake);
    }

    /**
     * The location in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Import(name="policyData")
    private @Nullable Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Optional<Output<String>> policyData() {
        return Optional.ofNullable(this.policyData);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    @Import(name="taskId")
    private @Nullable Output<String> taskId;

    public Optional<Output<String>> taskId() {
        return Optional.ofNullable(this.taskId);
    }

    private TaskIamPolicyState() {}

    private TaskIamPolicyState(TaskIamPolicyState $) {
        this.etag = $.etag;
        this.lake = $.lake;
        this.location = $.location;
        this.policyData = $.policyData;
        this.project = $.project;
        this.taskId = $.taskId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TaskIamPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TaskIamPolicyState $;

        public Builder() {
            $ = new TaskIamPolicyState();
        }

        public Builder(TaskIamPolicyState defaults) {
            $ = new TaskIamPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param etag (Computed) The etag of the IAM policy.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag (Computed) The etag of the IAM policy.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param lake The lake in which the task will be created in.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder lake(@Nullable Output<String> lake) {
            $.lake = lake;
            return this;
        }

        /**
         * @param lake The lake in which the task will be created in.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder lake(String lake) {
            return lake(Output.of(lake));
        }

        /**
         * @param location The location in which the task will be created in.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location in which the task will be created in.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(@Nullable Output<String> policyData) {
            $.policyData = policyData;
            return this;
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(String policyData) {
            return policyData(Output.of(policyData));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public Builder taskId(@Nullable Output<String> taskId) {
            $.taskId = taskId;
            return this;
        }

        public Builder taskId(String taskId) {
            return taskId(Output.of(taskId));
        }

        public TaskIamPolicyState build() {
            return $;
        }
    }

}
