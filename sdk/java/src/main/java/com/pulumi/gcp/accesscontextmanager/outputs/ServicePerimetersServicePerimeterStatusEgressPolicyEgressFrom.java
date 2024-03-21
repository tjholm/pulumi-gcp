// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.accesscontextmanager.outputs.ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom {
    /**
     * @return &#39;A list of identities that are allowed access through this `EgressPolicy`.
     * To specify an identity or identity group, use the IAM v1 format
     * specified [here](https://cloud.google.com/iam/docs/principal-identifiers.md#v1).
     * The following prefixes are supprted: user, group, serviceAccount, principal, and principalSet.&#39;
     * 
     */
    private @Nullable List<String> identities;
    /**
     * @return Specifies the type of identities that are allowed access to outside the
     * perimeter. If left unspecified, then members of `identities` field will
     * be allowed access.
     * Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.
     * 
     */
    private @Nullable String identityType;
    /**
     * @return Whether to enforce traffic restrictions based on `sources` field. If the `sources` field is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
     * Possible values are: `SOURCE_RESTRICTION_UNSPECIFIED`, `SOURCE_RESTRICTION_ENABLED`, `SOURCE_RESTRICTION_DISABLED`.
     * 
     */
    private @Nullable String sourceRestriction;
    /**
     * @return Sources that this EgressPolicy authorizes access from.
     * Structure is documented below.
     * 
     */
    private @Nullable List<ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource> sources;

    private ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom() {}
    /**
     * @return &#39;A list of identities that are allowed access through this `EgressPolicy`.
     * To specify an identity or identity group, use the IAM v1 format
     * specified [here](https://cloud.google.com/iam/docs/principal-identifiers.md#v1).
     * The following prefixes are supprted: user, group, serviceAccount, principal, and principalSet.&#39;
     * 
     */
    public List<String> identities() {
        return this.identities == null ? List.of() : this.identities;
    }
    /**
     * @return Specifies the type of identities that are allowed access to outside the
     * perimeter. If left unspecified, then members of `identities` field will
     * be allowed access.
     * Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.
     * 
     */
    public Optional<String> identityType() {
        return Optional.ofNullable(this.identityType);
    }
    /**
     * @return Whether to enforce traffic restrictions based on `sources` field. If the `sources` field is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
     * Possible values are: `SOURCE_RESTRICTION_UNSPECIFIED`, `SOURCE_RESTRICTION_ENABLED`, `SOURCE_RESTRICTION_DISABLED`.
     * 
     */
    public Optional<String> sourceRestriction() {
        return Optional.ofNullable(this.sourceRestriction);
    }
    /**
     * @return Sources that this EgressPolicy authorizes access from.
     * Structure is documented below.
     * 
     */
    public List<ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource> sources() {
        return this.sources == null ? List.of() : this.sources;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> identities;
        private @Nullable String identityType;
        private @Nullable String sourceRestriction;
        private @Nullable List<ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource> sources;
        public Builder() {}
        public Builder(ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.identities = defaults.identities;
    	      this.identityType = defaults.identityType;
    	      this.sourceRestriction = defaults.sourceRestriction;
    	      this.sources = defaults.sources;
        }

        @CustomType.Setter
        public Builder identities(@Nullable List<String> identities) {

            this.identities = identities;
            return this;
        }
        public Builder identities(String... identities) {
            return identities(List.of(identities));
        }
        @CustomType.Setter
        public Builder identityType(@Nullable String identityType) {

            this.identityType = identityType;
            return this;
        }
        @CustomType.Setter
        public Builder sourceRestriction(@Nullable String sourceRestriction) {

            this.sourceRestriction = sourceRestriction;
            return this;
        }
        @CustomType.Setter
        public Builder sources(@Nullable List<ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource> sources) {

            this.sources = sources;
            return this;
        }
        public Builder sources(ServicePerimetersServicePerimeterStatusEgressPolicyEgressFromSource... sources) {
            return sources(List.of(sources));
        }
        public ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom build() {
            final var _resultValue = new ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom();
            _resultValue.identities = identities;
            _resultValue.identityType = identityType;
            _resultValue.sourceRestriction = sourceRestriction;
            _resultValue.sources = sources;
            return _resultValue;
        }
    }
}
