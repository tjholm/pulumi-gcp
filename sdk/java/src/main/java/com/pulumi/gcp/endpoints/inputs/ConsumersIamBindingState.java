// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.endpoints.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.endpoints.inputs.ConsumersIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConsumersIamBindingState extends com.pulumi.resources.ResourceArgs {

    public static final ConsumersIamBindingState Empty = new ConsumersIamBindingState();

    @Import(name="condition")
    private @Nullable Output<ConsumersIamBindingConditionArgs> condition;

    public Optional<Output<ConsumersIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    @Import(name="consumerProject")
    private @Nullable Output<String> consumerProject;

    public Optional<Output<String>> consumerProject() {
        return Optional.ofNullable(this.consumerProject);
    }

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

    @Import(name="members")
    private @Nullable Output<List<String>> members;

    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The role that should be applied. Only one
     * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private ConsumersIamBindingState() {}

    private ConsumersIamBindingState(ConsumersIamBindingState $) {
        this.condition = $.condition;
        this.consumerProject = $.consumerProject;
        this.etag = $.etag;
        this.members = $.members;
        this.role = $.role;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConsumersIamBindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConsumersIamBindingState $;

        public Builder() {
            $ = new ConsumersIamBindingState();
        }

        public Builder(ConsumersIamBindingState defaults) {
            $ = new ConsumersIamBindingState(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<ConsumersIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(ConsumersIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        public Builder consumerProject(@Nullable Output<String> consumerProject) {
            $.consumerProject = consumerProject;
            return this;
        }

        public Builder consumerProject(String consumerProject) {
            return consumerProject(Output.of(consumerProject));
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

        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public ConsumersIamBindingState build() {
            return $;
        }
    }

}
