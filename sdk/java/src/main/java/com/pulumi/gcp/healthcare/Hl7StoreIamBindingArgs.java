// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.healthcare.inputs.Hl7StoreIamBindingConditionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Hl7StoreIamBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final Hl7StoreIamBindingArgs Empty = new Hl7StoreIamBindingArgs();

    @Import(name="condition")
    private @Nullable Output<Hl7StoreIamBindingConditionArgs> condition;

    public Optional<Output<Hl7StoreIamBindingConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Import(name="hl7V2StoreId", required=true)
    private Output<String> hl7V2StoreId;

    /**
     * @return The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> hl7V2StoreId() {
        return this.hl7V2StoreId;
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
     * 
     */
    @Import(name="members", required=true)
    private Output<List<String>> members;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }

    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private Hl7StoreIamBindingArgs() {}

    private Hl7StoreIamBindingArgs(Hl7StoreIamBindingArgs $) {
        this.condition = $.condition;
        this.hl7V2StoreId = $.hl7V2StoreId;
        this.members = $.members;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Hl7StoreIamBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Hl7StoreIamBindingArgs $;

        public Builder() {
            $ = new Hl7StoreIamBindingArgs();
        }

        public Builder(Hl7StoreIamBindingArgs defaults) {
            $ = new Hl7StoreIamBindingArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<Hl7StoreIamBindingConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(Hl7StoreIamBindingConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param hl7V2StoreId The HL7v2 store ID, in the form
         * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
         * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder hl7V2StoreId(Output<String> hl7V2StoreId) {
            $.hl7V2StoreId = hl7V2StoreId;
            return this;
        }

        /**
         * @param hl7V2StoreId The HL7v2 store ID, in the form
         * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
         * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder hl7V2StoreId(String hl7V2StoreId) {
            return hl7V2StoreId(Output.of(hl7V2StoreId));
        }

        /**
         * @param members Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * 
         * @return builder
         * 
         */
        public Builder members(Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
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
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
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
         * `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
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
         * `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Hl7StoreIamBindingArgs build() {
            if ($.hl7V2StoreId == null) {
                throw new MissingRequiredPropertyException("Hl7StoreIamBindingArgs", "hl7V2StoreId");
            }
            if ($.members == null) {
                throw new MissingRequiredPropertyException("Hl7StoreIamBindingArgs", "members");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("Hl7StoreIamBindingArgs", "role");
            }
            return $;
        }
    }

}
