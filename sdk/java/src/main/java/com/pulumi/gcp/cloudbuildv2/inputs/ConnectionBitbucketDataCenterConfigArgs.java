// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuildv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionBitbucketDataCenterConfigAuthorizerCredentialArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionBitbucketDataCenterConfigServiceDirectoryConfigArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionBitbucketDataCenterConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionBitbucketDataCenterConfigArgs Empty = new ConnectionBitbucketDataCenterConfigArgs();

    /**
     * Required. A http access token with the `REPO_ADMIN` scope access.
     * Structure is documented below.
     * 
     */
    @Import(name="authorizerCredential", required=true)
    private Output<ConnectionBitbucketDataCenterConfigAuthorizerCredentialArgs> authorizerCredential;

    /**
     * @return Required. A http access token with the `REPO_ADMIN` scope access.
     * Structure is documented below.
     * 
     */
    public Output<ConnectionBitbucketDataCenterConfigAuthorizerCredentialArgs> authorizerCredential() {
        return this.authorizerCredential;
    }

    /**
     * The URI of the Bitbucket Data Center host this connection is for.
     * 
     */
    @Import(name="hostUri", required=true)
    private Output<String> hostUri;

    /**
     * @return The URI of the Bitbucket Data Center host this connection is for.
     * 
     */
    public Output<String> hostUri() {
        return this.hostUri;
    }

    /**
     * Required. A http access token with the `REPO_READ` access.
     * Structure is documented below.
     * 
     */
    @Import(name="readAuthorizerCredential", required=true)
    private Output<ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialArgs> readAuthorizerCredential;

    /**
     * @return Required. A http access token with the `REPO_READ` access.
     * Structure is documented below.
     * 
     */
    public Output<ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialArgs> readAuthorizerCredential() {
        return this.readAuthorizerCredential;
    }

    /**
     * (Output)
     * Output only. Version of the Bitbucket Data Center running on the `host_uri`.
     * 
     */
    @Import(name="serverVersion")
    private @Nullable Output<String> serverVersion;

    /**
     * @return (Output)
     * Output only. Version of the Bitbucket Data Center running on the `host_uri`.
     * 
     */
    public Optional<Output<String>> serverVersion() {
        return Optional.ofNullable(this.serverVersion);
    }

    /**
     * Configuration for using Service Directory to privately connect to a Bitbucket Data Center. This should only be set if the Bitbucket Data Center is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the Bitbucket Data Center will be made over the public internet.
     * Structure is documented below.
     * 
     */
    @Import(name="serviceDirectoryConfig")
    private @Nullable Output<ConnectionBitbucketDataCenterConfigServiceDirectoryConfigArgs> serviceDirectoryConfig;

    /**
     * @return Configuration for using Service Directory to privately connect to a Bitbucket Data Center. This should only be set if the Bitbucket Data Center is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the Bitbucket Data Center will be made over the public internet.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionBitbucketDataCenterConfigServiceDirectoryConfigArgs>> serviceDirectoryConfig() {
        return Optional.ofNullable(this.serviceDirectoryConfig);
    }

    /**
     * SSL certificate to use for requests to the Bitbucket Data Center.
     * 
     */
    @Import(name="sslCa")
    private @Nullable Output<String> sslCa;

    /**
     * @return SSL certificate to use for requests to the Bitbucket Data Center.
     * 
     */
    public Optional<Output<String>> sslCa() {
        return Optional.ofNullable(this.sslCa);
    }

    /**
     * Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as `projects/*&#47;secrets/*&#47;versions/*`.
     * 
     */
    @Import(name="webhookSecretSecretVersion", required=true)
    private Output<String> webhookSecretSecretVersion;

    /**
     * @return Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as `projects/*&#47;secrets/*&#47;versions/*`.
     * 
     */
    public Output<String> webhookSecretSecretVersion() {
        return this.webhookSecretSecretVersion;
    }

    private ConnectionBitbucketDataCenterConfigArgs() {}

    private ConnectionBitbucketDataCenterConfigArgs(ConnectionBitbucketDataCenterConfigArgs $) {
        this.authorizerCredential = $.authorizerCredential;
        this.hostUri = $.hostUri;
        this.readAuthorizerCredential = $.readAuthorizerCredential;
        this.serverVersion = $.serverVersion;
        this.serviceDirectoryConfig = $.serviceDirectoryConfig;
        this.sslCa = $.sslCa;
        this.webhookSecretSecretVersion = $.webhookSecretSecretVersion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionBitbucketDataCenterConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionBitbucketDataCenterConfigArgs $;

        public Builder() {
            $ = new ConnectionBitbucketDataCenterConfigArgs();
        }

        public Builder(ConnectionBitbucketDataCenterConfigArgs defaults) {
            $ = new ConnectionBitbucketDataCenterConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizerCredential Required. A http access token with the `REPO_ADMIN` scope access.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder authorizerCredential(Output<ConnectionBitbucketDataCenterConfigAuthorizerCredentialArgs> authorizerCredential) {
            $.authorizerCredential = authorizerCredential;
            return this;
        }

        /**
         * @param authorizerCredential Required. A http access token with the `REPO_ADMIN` scope access.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder authorizerCredential(ConnectionBitbucketDataCenterConfigAuthorizerCredentialArgs authorizerCredential) {
            return authorizerCredential(Output.of(authorizerCredential));
        }

        /**
         * @param hostUri The URI of the Bitbucket Data Center host this connection is for.
         * 
         * @return builder
         * 
         */
        public Builder hostUri(Output<String> hostUri) {
            $.hostUri = hostUri;
            return this;
        }

        /**
         * @param hostUri The URI of the Bitbucket Data Center host this connection is for.
         * 
         * @return builder
         * 
         */
        public Builder hostUri(String hostUri) {
            return hostUri(Output.of(hostUri));
        }

        /**
         * @param readAuthorizerCredential Required. A http access token with the `REPO_READ` access.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder readAuthorizerCredential(Output<ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialArgs> readAuthorizerCredential) {
            $.readAuthorizerCredential = readAuthorizerCredential;
            return this;
        }

        /**
         * @param readAuthorizerCredential Required. A http access token with the `REPO_READ` access.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder readAuthorizerCredential(ConnectionBitbucketDataCenterConfigReadAuthorizerCredentialArgs readAuthorizerCredential) {
            return readAuthorizerCredential(Output.of(readAuthorizerCredential));
        }

        /**
         * @param serverVersion (Output)
         * Output only. Version of the Bitbucket Data Center running on the `host_uri`.
         * 
         * @return builder
         * 
         */
        public Builder serverVersion(@Nullable Output<String> serverVersion) {
            $.serverVersion = serverVersion;
            return this;
        }

        /**
         * @param serverVersion (Output)
         * Output only. Version of the Bitbucket Data Center running on the `host_uri`.
         * 
         * @return builder
         * 
         */
        public Builder serverVersion(String serverVersion) {
            return serverVersion(Output.of(serverVersion));
        }

        /**
         * @param serviceDirectoryConfig Configuration for using Service Directory to privately connect to a Bitbucket Data Center. This should only be set if the Bitbucket Data Center is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the Bitbucket Data Center will be made over the public internet.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serviceDirectoryConfig(@Nullable Output<ConnectionBitbucketDataCenterConfigServiceDirectoryConfigArgs> serviceDirectoryConfig) {
            $.serviceDirectoryConfig = serviceDirectoryConfig;
            return this;
        }

        /**
         * @param serviceDirectoryConfig Configuration for using Service Directory to privately connect to a Bitbucket Data Center. This should only be set if the Bitbucket Data Center is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the Bitbucket Data Center will be made over the public internet.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serviceDirectoryConfig(ConnectionBitbucketDataCenterConfigServiceDirectoryConfigArgs serviceDirectoryConfig) {
            return serviceDirectoryConfig(Output.of(serviceDirectoryConfig));
        }

        /**
         * @param sslCa SSL certificate to use for requests to the Bitbucket Data Center.
         * 
         * @return builder
         * 
         */
        public Builder sslCa(@Nullable Output<String> sslCa) {
            $.sslCa = sslCa;
            return this;
        }

        /**
         * @param sslCa SSL certificate to use for requests to the Bitbucket Data Center.
         * 
         * @return builder
         * 
         */
        public Builder sslCa(String sslCa) {
            return sslCa(Output.of(sslCa));
        }

        /**
         * @param webhookSecretSecretVersion Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as `projects/*&#47;secrets/*&#47;versions/*`.
         * 
         * @return builder
         * 
         */
        public Builder webhookSecretSecretVersion(Output<String> webhookSecretSecretVersion) {
            $.webhookSecretSecretVersion = webhookSecretSecretVersion;
            return this;
        }

        /**
         * @param webhookSecretSecretVersion Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as `projects/*&#47;secrets/*&#47;versions/*`.
         * 
         * @return builder
         * 
         */
        public Builder webhookSecretSecretVersion(String webhookSecretSecretVersion) {
            return webhookSecretSecretVersion(Output.of(webhookSecretSecretVersion));
        }

        public ConnectionBitbucketDataCenterConfigArgs build() {
            if ($.authorizerCredential == null) {
                throw new MissingRequiredPropertyException("ConnectionBitbucketDataCenterConfigArgs", "authorizerCredential");
            }
            if ($.hostUri == null) {
                throw new MissingRequiredPropertyException("ConnectionBitbucketDataCenterConfigArgs", "hostUri");
            }
            if ($.readAuthorizerCredential == null) {
                throw new MissingRequiredPropertyException("ConnectionBitbucketDataCenterConfigArgs", "readAuthorizerCredential");
            }
            if ($.webhookSecretSecretVersion == null) {
                throw new MissingRequiredPropertyException("ConnectionBitbucketDataCenterConfigArgs", "webhookSecretSecretVersion");
            }
            return $;
        }
    }

}
