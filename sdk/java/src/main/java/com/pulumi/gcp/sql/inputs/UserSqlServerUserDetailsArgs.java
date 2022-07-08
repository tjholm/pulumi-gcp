// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.sql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserSqlServerUserDetailsArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserSqlServerUserDetailsArgs Empty = new UserSqlServerUserDetailsArgs();

    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    @Import(name="serverRoles")
    private @Nullable Output<List<String>> serverRoles;

    public Optional<Output<List<String>>> serverRoles() {
        return Optional.ofNullable(this.serverRoles);
    }

    private UserSqlServerUserDetailsArgs() {}

    private UserSqlServerUserDetailsArgs(UserSqlServerUserDetailsArgs $) {
        this.disabled = $.disabled;
        this.serverRoles = $.serverRoles;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserSqlServerUserDetailsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserSqlServerUserDetailsArgs $;

        public Builder() {
            $ = new UserSqlServerUserDetailsArgs();
        }

        public Builder(UserSqlServerUserDetailsArgs defaults) {
            $ = new UserSqlServerUserDetailsArgs(Objects.requireNonNull(defaults));
        }

        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        public Builder serverRoles(@Nullable Output<List<String>> serverRoles) {
            $.serverRoles = serverRoles;
            return this;
        }

        public Builder serverRoles(List<String> serverRoles) {
            return serverRoles(Output.of(serverRoles));
        }

        public Builder serverRoles(String... serverRoles) {
            return serverRoles(List.of(serverRoles));
        }

        public UserSqlServerUserDetailsArgs build() {
            return $;
        }
    }

}
