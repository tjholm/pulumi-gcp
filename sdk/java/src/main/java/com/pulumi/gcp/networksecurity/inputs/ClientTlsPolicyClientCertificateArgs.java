// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networksecurity.inputs.ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs;
import com.pulumi.gcp.networksecurity.inputs.ClientTlsPolicyClientCertificateGrpcEndpointArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientTlsPolicyClientCertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientTlsPolicyClientCertificateArgs Empty = new ClientTlsPolicyClientCertificateArgs();

    /**
     * The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
     * Structure is documented below.
     * 
     */
    @Import(name="certificateProviderInstance")
    private @Nullable Output<ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs> certificateProviderInstance;

    /**
     * @return The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs>> certificateProviderInstance() {
        return Optional.ofNullable(this.certificateProviderInstance);
    }

    /**
     * gRPC specific configuration to access the gRPC server to obtain the cert and private key.
     * Structure is documented below.
     * 
     */
    @Import(name="grpcEndpoint")
    private @Nullable Output<ClientTlsPolicyClientCertificateGrpcEndpointArgs> grpcEndpoint;

    /**
     * @return gRPC specific configuration to access the gRPC server to obtain the cert and private key.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClientTlsPolicyClientCertificateGrpcEndpointArgs>> grpcEndpoint() {
        return Optional.ofNullable(this.grpcEndpoint);
    }

    private ClientTlsPolicyClientCertificateArgs() {}

    private ClientTlsPolicyClientCertificateArgs(ClientTlsPolicyClientCertificateArgs $) {
        this.certificateProviderInstance = $.certificateProviderInstance;
        this.grpcEndpoint = $.grpcEndpoint;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientTlsPolicyClientCertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientTlsPolicyClientCertificateArgs $;

        public Builder() {
            $ = new ClientTlsPolicyClientCertificateArgs();
        }

        public Builder(ClientTlsPolicyClientCertificateArgs defaults) {
            $ = new ClientTlsPolicyClientCertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateProviderInstance The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificateProviderInstance(@Nullable Output<ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs> certificateProviderInstance) {
            $.certificateProviderInstance = certificateProviderInstance;
            return this;
        }

        /**
         * @param certificateProviderInstance The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificateProviderInstance(ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs certificateProviderInstance) {
            return certificateProviderInstance(Output.of(certificateProviderInstance));
        }

        /**
         * @param grpcEndpoint gRPC specific configuration to access the gRPC server to obtain the cert and private key.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder grpcEndpoint(@Nullable Output<ClientTlsPolicyClientCertificateGrpcEndpointArgs> grpcEndpoint) {
            $.grpcEndpoint = grpcEndpoint;
            return this;
        }

        /**
         * @param grpcEndpoint gRPC specific configuration to access the gRPC server to obtain the cert and private key.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder grpcEndpoint(ClientTlsPolicyClientCertificateGrpcEndpointArgs grpcEndpoint) {
            return grpcEndpoint(Output.of(grpcEndpoint));
        }

        public ClientTlsPolicyClientCertificateArgs build() {
            return $;
        }
    }

}
