// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.datacatalog;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.datacatalog.inputs.PolicyTagIamMemberConditionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyTagIamMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyTagIamMemberArgs Empty = new PolicyTagIamMemberArgs();

    @Import(name="condition")
    private @Nullable Output<PolicyTagIamMemberConditionArgs> condition;

    public Optional<Output<PolicyTagIamMemberConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    @Import(name="member", required=true)
    private Output<String> member;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    public Output<String> member() {
        return this.member;
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="policyTag", required=true)
    private Output<String> policyTag;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> policyTag() {
        return this.policyTag;
    }

    /**
     * The role that should be applied. Only one
     * `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private PolicyTagIamMemberArgs() {}

    private PolicyTagIamMemberArgs(PolicyTagIamMemberArgs $) {
        this.condition = $.condition;
        this.member = $.member;
        this.policyTag = $.policyTag;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyTagIamMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyTagIamMemberArgs $;

        public Builder() {
            $ = new PolicyTagIamMemberArgs();
        }

        public Builder(PolicyTagIamMemberArgs defaults) {
            $ = new PolicyTagIamMemberArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<PolicyTagIamMemberConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(PolicyTagIamMemberConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param member Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
         * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
         * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
         * 
         * @return builder
         * 
         */
        public Builder member(Output<String> member) {
            $.member = member;
            return this;
        }

        /**
         * @param member Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
         * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
         * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
         * 
         * @return builder
         * 
         */
        public Builder member(String member) {
            return member(Output.of(member));
        }

        /**
         * @param policyTag Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder policyTag(Output<String> policyTag) {
            $.policyTag = policyTag;
            return this;
        }

        /**
         * @param policyTag Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder policyTag(String policyTag) {
            return policyTag(Output.of(policyTag));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
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
         * `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public PolicyTagIamMemberArgs build() {
            if ($.member == null) {
                throw new MissingRequiredPropertyException("PolicyTagIamMemberArgs", "member");
            }
            if ($.policyTag == null) {
                throw new MissingRequiredPropertyException("PolicyTagIamMemberArgs", "policyTag");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("PolicyTagIamMemberArgs", "role");
            }
            return $;
        }
    }

}
