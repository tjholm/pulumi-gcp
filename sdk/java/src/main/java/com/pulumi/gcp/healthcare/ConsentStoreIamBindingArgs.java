// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.healthcare.inputs.ConsentStoreIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConsentStoreIamBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConsentStoreIamBindingArgs Empty = new ConsentStoreIamBindingArgs();

    @Import(name="condition")
    private @Nullable Output<ConsentStoreIamBindingConditionArgs> condition;

    public Optional<Output<ConsentStoreIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="consentStoreId", required=true)
    private Output<String> consentStoreId;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> consentStoreId() {
        return this.consentStoreId;
    }

    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="dataset", required=true)
    private Output<String> dataset;

    /**
     * @return Identifies the dataset addressed by this request. Must be in the format
     * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> dataset() {
        return this.dataset;
    }

    @Import(name="members", required=true)
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }

    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private ConsentStoreIamBindingArgs() {}

    private ConsentStoreIamBindingArgs(ConsentStoreIamBindingArgs $) {
        this.condition = $.condition;
        this.consentStoreId = $.consentStoreId;
        this.dataset = $.dataset;
        this.members = $.members;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConsentStoreIamBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConsentStoreIamBindingArgs $;

        public Builder() {
            $ = new ConsentStoreIamBindingArgs();
        }

        public Builder(ConsentStoreIamBindingArgs defaults) {
            $ = new ConsentStoreIamBindingArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<ConsentStoreIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(ConsentStoreIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param consentStoreId Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder consentStoreId(Output<String> consentStoreId) {
            $.consentStoreId = consentStoreId;
            return this;
        }

        /**
         * @param consentStoreId Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder consentStoreId(String consentStoreId) {
            return consentStoreId(Output.of(consentStoreId));
        }

        /**
         * @param dataset Identifies the dataset addressed by this request. Must be in the format
         * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder dataset(Output<String> dataset) {
            $.dataset = dataset;
            return this;
        }

        /**
         * @param dataset Identifies the dataset addressed by this request. Must be in the format
         * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
         * Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder dataset(String dataset) {
            return dataset(Output.of(dataset));
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
         * @param role The role that should be applied. Only one
         * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
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
         * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public ConsentStoreIamBindingArgs build() {
            $.consentStoreId = Objects.requireNonNull($.consentStoreId, "expected parameter 'consentStoreId' to be non-null");
            $.dataset = Objects.requireNonNull($.dataset, "expected parameter 'dataset' to be non-null");
            $.members = Objects.requireNonNull($.members, "expected parameter 'members' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
