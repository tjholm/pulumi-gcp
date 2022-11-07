// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiFeatureStoreIamBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final AiFeatureStoreIamBindingArgs Empty = new AiFeatureStoreIamBindingArgs();

    @Import(name="condition")
    private @Nullable Output<AiFeatureStoreIamBindingConditionArgs> condition;

    public Optional<Output<AiFeatureStoreIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    @Import(name="featurestore", required=true)
    private Output<String> featurestore;

    public Output<String> featurestore() {
        return this.featurestore;
    }

    @Import(name="members", required=true)
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The region of the dataset. eg us-central1
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region of the dataset. eg us-central1
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    @Import(name="role", required=true)
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }

    private AiFeatureStoreIamBindingArgs() {}

    private AiFeatureStoreIamBindingArgs(AiFeatureStoreIamBindingArgs $) {
        this.condition = $.condition;
        this.featurestore = $.featurestore;
        this.members = $.members;
        this.project = $.project;
        this.region = $.region;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiFeatureStoreIamBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiFeatureStoreIamBindingArgs $;

        public Builder() {
            $ = new AiFeatureStoreIamBindingArgs();
        }

        public Builder(AiFeatureStoreIamBindingArgs defaults) {
            $ = new AiFeatureStoreIamBindingArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<AiFeatureStoreIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(AiFeatureStoreIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        public Builder featurestore(Output<String> featurestore) {
            $.featurestore = featurestore;
            return this;
        }

        public Builder featurestore(String featurestore) {
            return featurestore(Output.of(featurestore));
        }

        public Builder members(Output<List<String>> members) {
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
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
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
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param region The region of the dataset. eg us-central1
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region of the dataset. eg us-central1
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public AiFeatureStoreIamBindingArgs build() {
            $.featurestore = Objects.requireNonNull($.featurestore, "expected parameter 'featurestore' to be non-null");
            $.members = Objects.requireNonNull($.members, "expected parameter 'members' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
