// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.serviceaccount.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountState extends com.pulumi.resources.ResourceArgs {

    public static final AccountState Empty = new AccountState();

    /**
     * The account id that is used to generate the service
     * account email address and a stable unique id. It is unique within a project,
     * must be 6-30 characters long, and match the regular expression `a-z`
     * to comply with RFC1035. Changing this forces a new service account to be created.
     * 
     */
    @Import(name="accountId")
    private @Nullable Output<String> accountId;

    /**
     * @return The account id that is used to generate the service
     * account email address and a stable unique id. It is unique within a project,
     * must be 6-30 characters long, and match the regular expression `a-z`
     * to comply with RFC1035. Changing this forces a new service account to be created.
     * 
     */
    public Optional<Output<String>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    /**
     * If set to true, skip service account creation if a service account with the same email already exists.
     * 
     */
    @Import(name="createIgnoreAlreadyExists")
    private @Nullable Output<Boolean> createIgnoreAlreadyExists;

    /**
     * @return If set to true, skip service account creation if a service account with the same email already exists.
     * 
     */
    public Optional<Output<Boolean>> createIgnoreAlreadyExists() {
        return Optional.ofNullable(this.createIgnoreAlreadyExists);
    }

    /**
     * A text description of the service account.
     * Must be less than or equal to 256 UTF-8 bytes.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A text description of the service account.
     * Must be less than or equal to 256 UTF-8 bytes.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
     * Must be set after creation to disable a service account.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
     * Must be set after creation to disable a service account.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * The display name for the service account.
     * Can be updated without creating a new resource.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name for the service account.
     * Can be updated without creating a new resource.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The e-mail address of the service account. This value
     * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
     * that would grant the service account privileges.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The e-mail address of the service account. This value
     * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
     * that would grant the service account privileges.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
     * 
     */
    @Import(name="member")
    private @Nullable Output<String> member;

    /**
     * @return The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
     * 
     */
    public Optional<Output<String>> member() {
        return Optional.ofNullable(this.member);
    }

    /**
     * The fully-qualified name of the service account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The fully-qualified name of the service account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project that the service account will be created in.
     * Defaults to the provider project configuration.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The unique id of the service account.
     * 
     */
    @Import(name="uniqueId")
    private @Nullable Output<String> uniqueId;

    /**
     * @return The unique id of the service account.
     * 
     */
    public Optional<Output<String>> uniqueId() {
        return Optional.ofNullable(this.uniqueId);
    }

    private AccountState() {}

    private AccountState(AccountState $) {
        this.accountId = $.accountId;
        this.createIgnoreAlreadyExists = $.createIgnoreAlreadyExists;
        this.description = $.description;
        this.disabled = $.disabled;
        this.displayName = $.displayName;
        this.email = $.email;
        this.member = $.member;
        this.name = $.name;
        this.project = $.project;
        this.uniqueId = $.uniqueId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountState $;

        public Builder() {
            $ = new AccountState();
        }

        public Builder(AccountState defaults) {
            $ = new AccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountId The account id that is used to generate the service
         * account email address and a stable unique id. It is unique within a project,
         * must be 6-30 characters long, and match the regular expression `a-z`
         * to comply with RFC1035. Changing this forces a new service account to be created.
         * 
         * @return builder
         * 
         */
        public Builder accountId(@Nullable Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId The account id that is used to generate the service
         * account email address and a stable unique id. It is unique within a project,
         * must be 6-30 characters long, and match the regular expression `a-z`
         * to comply with RFC1035. Changing this forces a new service account to be created.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param createIgnoreAlreadyExists If set to true, skip service account creation if a service account with the same email already exists.
         * 
         * @return builder
         * 
         */
        public Builder createIgnoreAlreadyExists(@Nullable Output<Boolean> createIgnoreAlreadyExists) {
            $.createIgnoreAlreadyExists = createIgnoreAlreadyExists;
            return this;
        }

        /**
         * @param createIgnoreAlreadyExists If set to true, skip service account creation if a service account with the same email already exists.
         * 
         * @return builder
         * 
         */
        public Builder createIgnoreAlreadyExists(Boolean createIgnoreAlreadyExists) {
            return createIgnoreAlreadyExists(Output.of(createIgnoreAlreadyExists));
        }

        /**
         * @param description A text description of the service account.
         * Must be less than or equal to 256 UTF-8 bytes.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A text description of the service account.
         * Must be less than or equal to 256 UTF-8 bytes.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disabled Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
         * Must be set after creation to disable a service account.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
         * Must be set after creation to disable a service account.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param displayName The display name for the service account.
         * Can be updated without creating a new resource.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name for the service account.
         * Can be updated without creating a new resource.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param email The e-mail address of the service account. This value
         * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
         * that would grant the service account privileges.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The e-mail address of the service account. This value
         * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
         * that would grant the service account privileges.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param member The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
         * 
         * @return builder
         * 
         */
        public Builder member(@Nullable Output<String> member) {
            $.member = member;
            return this;
        }

        /**
         * @param member The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
         * 
         * @return builder
         * 
         */
        public Builder member(String member) {
            return member(Output.of(member));
        }

        /**
         * @param name The fully-qualified name of the service account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The fully-qualified name of the service account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project that the service account will be created in.
         * Defaults to the provider project configuration.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project that the service account will be created in.
         * Defaults to the provider project configuration.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param uniqueId The unique id of the service account.
         * 
         * @return builder
         * 
         */
        public Builder uniqueId(@Nullable Output<String> uniqueId) {
            $.uniqueId = uniqueId;
            return this;
        }

        /**
         * @param uniqueId The unique id of the service account.
         * 
         * @return builder
         * 
         */
        public Builder uniqueId(String uniqueId) {
            return uniqueId(Output.of(uniqueId));
        }

        public AccountState build() {
            return $;
        }
    }

}
