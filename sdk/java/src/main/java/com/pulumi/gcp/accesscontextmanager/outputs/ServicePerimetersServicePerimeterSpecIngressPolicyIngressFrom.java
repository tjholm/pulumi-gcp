// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.accesscontextmanager.outputs.ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom {
    /**
     * @return A list of identities that are allowed access through this `EgressPolicy`.
     * Should be in the format of email address. The email address should
     * represent individual user or service account only.
     * 
     */
    private final @Nullable List<String> identities;
    /**
     * @return Specifies the type of identities that are allowed access to outside the
     * perimeter. If left unspecified, then members of `identities` field will
     * be allowed access.
     * Possible values are `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, and `ANY_SERVICE_ACCOUNT`.
     * 
     */
    private final @Nullable String identityType;
    /**
     * @return Sources that this `IngressPolicy` authorizes access from.
     * Structure is documented below.
     * 
     */
    private final @Nullable List<ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource> sources;

    @CustomType.Constructor
    private ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom(
        @CustomType.Parameter("identities") @Nullable List<String> identities,
        @CustomType.Parameter("identityType") @Nullable String identityType,
        @CustomType.Parameter("sources") @Nullable List<ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource> sources) {
        this.identities = identities;
        this.identityType = identityType;
        this.sources = sources;
    }

    /**
     * @return A list of identities that are allowed access through this `EgressPolicy`.
     * Should be in the format of email address. The email address should
     * represent individual user or service account only.
     * 
     */
    public List<String> identities() {
        return this.identities == null ? List.of() : this.identities;
    }
    /**
     * @return Specifies the type of identities that are allowed access to outside the
     * perimeter. If left unspecified, then members of `identities` field will
     * be allowed access.
     * Possible values are `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, and `ANY_SERVICE_ACCOUNT`.
     * 
     */
    public Optional<String> identityType() {
        return Optional.ofNullable(this.identityType);
    }
    /**
     * @return Sources that this `IngressPolicy` authorizes access from.
     * Structure is documented below.
     * 
     */
    public List<ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource> sources() {
        return this.sources == null ? List.of() : this.sources;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> identities;
        private @Nullable String identityType;
        private @Nullable List<ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource> sources;

        public Builder() {
    	      // Empty
        }

        public Builder(ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.identities = defaults.identities;
    	      this.identityType = defaults.identityType;
    	      this.sources = defaults.sources;
        }

        public Builder identities(@Nullable List<String> identities) {
            this.identities = identities;
            return this;
        }
        public Builder identities(String... identities) {
            return identities(List.of(identities));
        }
        public Builder identityType(@Nullable String identityType) {
            this.identityType = identityType;
            return this;
        }
        public Builder sources(@Nullable List<ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource> sources) {
            this.sources = sources;
            return this;
        }
        public Builder sources(ServicePerimetersServicePerimeterSpecIngressPolicyIngressFromSource... sources) {
            return sources(List.of(sources));
        }        public ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom build() {
            return new ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom(identities, identityType, sources);
        }
    }
}
