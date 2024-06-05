// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTaskIamPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTaskIamPolicyArgs Empty = new GetTaskIamPolicyArgs();

    /**
     * The lake in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="lake", required=true)
    private Output<String> lake;

    /**
     * @return The lake in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> lake() {
        return this.lake;
    }

    /**
     * The location in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location in which the task will be created in.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
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

    @Import(name="taskId", required=true)
    private Output<String> taskId;

    public Output<String> taskId() {
        return this.taskId;
    }

    private GetTaskIamPolicyArgs() {}

    private GetTaskIamPolicyArgs(GetTaskIamPolicyArgs $) {
        this.lake = $.lake;
        this.location = $.location;
        this.project = $.project;
        this.taskId = $.taskId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTaskIamPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTaskIamPolicyArgs $;

        public Builder() {
            $ = new GetTaskIamPolicyArgs();
        }

        public Builder(GetTaskIamPolicyArgs defaults) {
            $ = new GetTaskIamPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param lake The lake in which the task will be created in.
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder lake(Output<String> lake) {
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
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
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
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
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

        public Builder taskId(Output<String> taskId) {
            $.taskId = taskId;
            return this;
        }

        public Builder taskId(String taskId) {
            return taskId(Output.of(taskId));
        }

        public GetTaskIamPolicyArgs build() {
            if ($.lake == null) {
                throw new MissingRequiredPropertyException("GetTaskIamPolicyArgs", "lake");
            }
            if ($.taskId == null) {
                throw new MissingRequiredPropertyException("GetTaskIamPolicyArgs", "taskId");
            }
            return $;
        }
    }

}
