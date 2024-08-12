// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuildv2.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConnectionBitbucketDataCenterConfigReadAuthorizerCredential {
    /**
     * @return Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: `projects/*&#47;secrets/*&#47;versions/*`.
     * 
     */
    private String userTokenSecretVersion;
    /**
     * @return (Output)
     * Output only. The username associated to this token.
     * 
     */
    private @Nullable String username;

    private ConnectionBitbucketDataCenterConfigReadAuthorizerCredential() {}
    /**
     * @return Required. A SecretManager resource containing the user token that authorizes the Cloud Build connection. Format: `projects/*&#47;secrets/*&#47;versions/*`.
     * 
     */
    public String userTokenSecretVersion() {
        return this.userTokenSecretVersion;
    }
    /**
     * @return (Output)
     * Output only. The username associated to this token.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectionBitbucketDataCenterConfigReadAuthorizerCredential defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String userTokenSecretVersion;
        private @Nullable String username;
        public Builder() {}
        public Builder(ConnectionBitbucketDataCenterConfigReadAuthorizerCredential defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.userTokenSecretVersion = defaults.userTokenSecretVersion;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder userTokenSecretVersion(String userTokenSecretVersion) {
            if (userTokenSecretVersion == null) {
              throw new MissingRequiredPropertyException("ConnectionBitbucketDataCenterConfigReadAuthorizerCredential", "userTokenSecretVersion");
            }
            this.userTokenSecretVersion = userTokenSecretVersion;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        public ConnectionBitbucketDataCenterConfigReadAuthorizerCredential build() {
            final var _resultValue = new ConnectionBitbucketDataCenterConfigReadAuthorizerCredential();
            _resultValue.userTokenSecretVersion = userTokenSecretVersion;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
