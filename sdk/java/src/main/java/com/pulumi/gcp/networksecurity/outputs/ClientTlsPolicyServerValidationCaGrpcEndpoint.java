// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ClientTlsPolicyServerValidationCaGrpcEndpoint {
    /**
     * @return The target URI of the gRPC endpoint. Only UDS path is supported, and should start with &#34;unix:&#34;.
     * 
     */
    private String targetUri;

    private ClientTlsPolicyServerValidationCaGrpcEndpoint() {}
    /**
     * @return The target URI of the gRPC endpoint. Only UDS path is supported, and should start with &#34;unix:&#34;.
     * 
     */
    public String targetUri() {
        return this.targetUri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientTlsPolicyServerValidationCaGrpcEndpoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String targetUri;
        public Builder() {}
        public Builder(ClientTlsPolicyServerValidationCaGrpcEndpoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.targetUri = defaults.targetUri;
        }

        @CustomType.Setter
        public Builder targetUri(String targetUri) {
            this.targetUri = Objects.requireNonNull(targetUri);
            return this;
        }
        public ClientTlsPolicyServerValidationCaGrpcEndpoint build() {
            final var o = new ClientTlsPolicyServerValidationCaGrpcEndpoint();
            o.targetUri = targetUri;
            return o;
        }
    }
}
