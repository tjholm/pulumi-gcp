// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.billing.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.billing.inputs.AccountIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountIamBindingState extends com.pulumi.resources.ResourceArgs {

    public static final AccountIamBindingState Empty = new AccountIamBindingState();

    /**
     * The billing account id.
     * 
     * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
     * 
     */
    @Import(name="billingAccountId")
    private @Nullable Output<String> billingAccountId;

    /**
     * @return The billing account id.
     * 
     * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
     * 
     */
    public Optional<Output<String>> billingAccountId() {
        return Optional.ofNullable(this.billingAccountId);
    }

    @Import(name="condition")
    private @Nullable Output<AccountIamBindingConditionArgs> condition;

    public Optional<Output<AccountIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * (Computed) The etag of the billing account&#39;s IAM policy.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return (Computed) The etag of the billing account&#39;s IAM policy.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The role that should be applied. Only one
     * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     * `gcp.billing.AccountIamPolicy` only:
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
     * 
     * `gcp.billing.AccountIamPolicy` only:
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private AccountIamBindingState() {}

    private AccountIamBindingState(AccountIamBindingState $) {
        this.billingAccountId = $.billingAccountId;
        this.condition = $.condition;
        this.etag = $.etag;
        this.members = $.members;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountIamBindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountIamBindingState $;

        public Builder() {
            $ = new AccountIamBindingState();
        }

        public Builder(AccountIamBindingState defaults) {
            $ = new AccountIamBindingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param billingAccountId The billing account id.
         * 
         * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
         * 
         * @return builder
         * 
         */
        public Builder billingAccountId(@Nullable Output<String> billingAccountId) {
            $.billingAccountId = billingAccountId;
            return this;
        }

        /**
         * @param billingAccountId The billing account id.
         * 
         * For `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`:
         * 
         * @return builder
         * 
         */
        public Builder billingAccountId(String billingAccountId) {
            return billingAccountId(Output.of(billingAccountId));
        }

        public Builder condition(@Nullable Output<AccountIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(AccountIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param etag (Computed) The etag of the billing account&#39;s IAM policy.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag (Computed) The etag of the billing account&#39;s IAM policy.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param members Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
         * 
         * `gcp.billing.AccountIamPolicy` only:
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
         * `gcp.billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
         * 
         * `gcp.billing.AccountIamPolicy` only:
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public AccountIamBindingState build() {
            return $;
        }
    }

}
