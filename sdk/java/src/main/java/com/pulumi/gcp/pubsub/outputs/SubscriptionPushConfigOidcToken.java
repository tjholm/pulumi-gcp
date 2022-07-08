// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.pubsub.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SubscriptionPushConfigOidcToken {
    /**
     * @return Audience to be used when generating OIDC token. The audience claim
     * identifies the recipients that the JWT is intended for. The audience
     * value is a single case-sensitive string. Having multiple values (array)
     * for the audience field is not supported. More info about the OIDC JWT
     * token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
     * Note: if not specified, the Push endpoint URL will be used.
     * 
     */
    private final @Nullable String audience;
    /**
     * @return Service account email to be used for generating the OIDC token.
     * The caller (for subscriptions.create, subscriptions.patch, and
     * subscriptions.modifyPushConfig RPCs) must have the
     * iam.serviceAccounts.actAs permission for the service account.
     * 
     */
    private final String serviceAccountEmail;

    @CustomType.Constructor
    private SubscriptionPushConfigOidcToken(
        @CustomType.Parameter("audience") @Nullable String audience,
        @CustomType.Parameter("serviceAccountEmail") String serviceAccountEmail) {
        this.audience = audience;
        this.serviceAccountEmail = serviceAccountEmail;
    }

    /**
     * @return Audience to be used when generating OIDC token. The audience claim
     * identifies the recipients that the JWT is intended for. The audience
     * value is a single case-sensitive string. Having multiple values (array)
     * for the audience field is not supported. More info about the OIDC JWT
     * token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
     * Note: if not specified, the Push endpoint URL will be used.
     * 
     */
    public Optional<String> audience() {
        return Optional.ofNullable(this.audience);
    }
    /**
     * @return Service account email to be used for generating the OIDC token.
     * The caller (for subscriptions.create, subscriptions.patch, and
     * subscriptions.modifyPushConfig RPCs) must have the
     * iam.serviceAccounts.actAs permission for the service account.
     * 
     */
    public String serviceAccountEmail() {
        return this.serviceAccountEmail;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SubscriptionPushConfigOidcToken defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String audience;
        private String serviceAccountEmail;

        public Builder() {
    	      // Empty
        }

        public Builder(SubscriptionPushConfigOidcToken defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.audience = defaults.audience;
    	      this.serviceAccountEmail = defaults.serviceAccountEmail;
        }

        public Builder audience(@Nullable String audience) {
            this.audience = audience;
            return this;
        }
        public Builder serviceAccountEmail(String serviceAccountEmail) {
            this.serviceAccountEmail = Objects.requireNonNull(serviceAccountEmail);
            return this;
        }        public SubscriptionPushConfigOidcToken build() {
            return new SubscriptionPushConfigOidcToken(audience, serviceAccountEmail);
        }
    }
}
