// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.organizations.inputs.IAMBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IAMBindingState extends com.pulumi.resources.ResourceArgs {

    public static final IAMBindingState Empty = new IAMBindingState();

    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Import(name="condition")
    private @Nullable Output<IAMBindingConditionArgs> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Optional<Output<IAMBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * (Computed) The etag of the organization&#39;s IAM policy.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return (Computed) The etag of the organization&#39;s IAM policy.
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
     * The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
     * `organizations/{{org_id}}/roles/{{role_id}}`.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private IAMBindingState() {}

    private IAMBindingState(IAMBindingState $) {
        this.condition = $.condition;
        this.etag = $.etag;
        this.members = $.members;
        this.orgId = $.orgId;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IAMBindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IAMBindingState $;

        public Builder() {
            $ = new IAMBindingState();
        }

        public Builder(IAMBindingState defaults) {
            $ = new IAMBindingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param condition An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder condition(@Nullable Output<IAMBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        /**
         * @param condition An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder condition(IAMBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param etag (Computed) The etag of the organization&#39;s IAM policy.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag (Computed) The etag of the organization&#39;s IAM policy.
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
         * @param orgId The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
         * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
         * will not be inferred from the provider.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
         * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
         * will not be inferred from the provider.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
         * `organizations/{{org_id}}/roles/{{role_id}}`.
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
         * `gcp.organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
         * `organizations/{{org_id}}/roles/{{role_id}}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public IAMBindingState build() {
            return $;
        }
    }

}
