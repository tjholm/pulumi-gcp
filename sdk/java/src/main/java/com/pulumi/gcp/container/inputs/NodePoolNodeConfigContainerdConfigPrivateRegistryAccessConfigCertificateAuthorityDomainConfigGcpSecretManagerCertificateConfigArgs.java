// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs Empty = new NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs();

    /**
     * URI for the secret that hosts a certificate. Must be in the format &#39;projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST&#39;.
     * 
     */
    @Import(name="secretUri", required=true)
    private Output<String> secretUri;

    /**
     * @return URI for the secret that hosts a certificate. Must be in the format &#39;projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST&#39;.
     * 
     */
    public Output<String> secretUri() {
        return this.secretUri;
    }

    private NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs() {}

    private NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs(NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs $) {
        this.secretUri = $.secretUri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs $;

        public Builder() {
            $ = new NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs();
        }

        public Builder(NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs defaults) {
            $ = new NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param secretUri URI for the secret that hosts a certificate. Must be in the format &#39;projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST&#39;.
         * 
         * @return builder
         * 
         */
        public Builder secretUri(Output<String> secretUri) {
            $.secretUri = secretUri;
            return this;
        }

        /**
         * @param secretUri URI for the secret that hosts a certificate. Must be in the format &#39;projects/PROJECT_NUM/secrets/SECRET_NAME/versions/VERSION_OR_LATEST&#39;.
         * 
         * @return builder
         * 
         */
        public Builder secretUri(String secretUri) {
            return secretUri(Output.of(secretUri));
        }

        public NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs build() {
            if ($.secretUri == null) {
                throw new MissingRequiredPropertyException("NodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfigArgs", "secretUri");
            }
            return $;
        }
    }

}
