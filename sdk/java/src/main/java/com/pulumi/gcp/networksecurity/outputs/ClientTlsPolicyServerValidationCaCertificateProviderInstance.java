// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ClientTlsPolicyServerValidationCaCertificateProviderInstance {
    /**
     * @return Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to &#34;google_cloud_private_spiffe&#34; to use Certificate Authority Service certificate provider instance.
     * 
     */
    private String pluginInstance;

    private ClientTlsPolicyServerValidationCaCertificateProviderInstance() {}
    /**
     * @return Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to &#34;google_cloud_private_spiffe&#34; to use Certificate Authority Service certificate provider instance.
     * 
     */
    public String pluginInstance() {
        return this.pluginInstance;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientTlsPolicyServerValidationCaCertificateProviderInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String pluginInstance;
        public Builder() {}
        public Builder(ClientTlsPolicyServerValidationCaCertificateProviderInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.pluginInstance = defaults.pluginInstance;
        }

        @CustomType.Setter
        public Builder pluginInstance(String pluginInstance) {
            this.pluginInstance = Objects.requireNonNull(pluginInstance);
            return this;
        }
        public ClientTlsPolicyServerValidationCaCertificateProviderInstance build() {
            final var o = new ClientTlsPolicyServerValidationCaCertificateProviderInstance();
            o.pluginInstance = pluginInstance;
            return o;
        }
    }
}
