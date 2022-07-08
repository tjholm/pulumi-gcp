// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOrganizationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationPlainArgs Empty = new GetOrganizationPlainArgs();

    /**
     * The domain name of the Organization.
     * 
     */
    @Import(name="domain")
    private @Nullable String domain;

    /**
     * @return The domain name of the Organization.
     * 
     */
    public Optional<String> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * The Organization&#39;s numeric ID, including an optional `organizations/` prefix.
     * 
     */
    @Import(name="organization")
    private @Nullable String organization;

    /**
     * @return The Organization&#39;s numeric ID, including an optional `organizations/` prefix.
     * 
     */
    public Optional<String> organization() {
        return Optional.ofNullable(this.organization);
    }

    private GetOrganizationPlainArgs() {}

    private GetOrganizationPlainArgs(GetOrganizationPlainArgs $) {
        this.domain = $.domain;
        this.organization = $.organization;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationPlainArgs $;

        public Builder() {
            $ = new GetOrganizationPlainArgs();
        }

        public Builder(GetOrganizationPlainArgs defaults) {
            $ = new GetOrganizationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The domain name of the Organization.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable String domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param organization The Organization&#39;s numeric ID, including an optional `organizations/` prefix.
         * 
         * @return builder
         * 
         */
        public Builder organization(@Nullable String organization) {
            $.organization = organization;
            return this;
        }

        public GetOrganizationPlainArgs build() {
            return $;
        }
    }

}
