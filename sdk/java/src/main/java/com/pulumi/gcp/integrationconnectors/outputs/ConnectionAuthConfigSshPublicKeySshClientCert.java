// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.integrationconnectors.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ConnectionAuthConfigSshPublicKeySshClientCert {
    /**
     * @return The resource name of the secret version in the format,
     * format as: projects/*&#47;secrets/*&#47;versions/*.
     * 
     */
    private String secretVersion;

    private ConnectionAuthConfigSshPublicKeySshClientCert() {}
    /**
     * @return The resource name of the secret version in the format,
     * format as: projects/*&#47;secrets/*&#47;versions/*.
     * 
     */
    public String secretVersion() {
        return this.secretVersion;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConnectionAuthConfigSshPublicKeySshClientCert defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String secretVersion;
        public Builder() {}
        public Builder(ConnectionAuthConfigSshPublicKeySshClientCert defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.secretVersion = defaults.secretVersion;
        }

        @CustomType.Setter
        public Builder secretVersion(String secretVersion) {
            if (secretVersion == null) {
              throw new MissingRequiredPropertyException("ConnectionAuthConfigSshPublicKeySshClientCert", "secretVersion");
            }
            this.secretVersion = secretVersion;
            return this;
        }
        public ConnectionAuthConfigSshPublicKeySshClientCert build() {
            final var _resultValue = new ConnectionAuthConfigSshPublicKeySshClientCert();
            _resultValue.secretVersion = secretVersion;
            return _resultValue;
        }
    }
}
