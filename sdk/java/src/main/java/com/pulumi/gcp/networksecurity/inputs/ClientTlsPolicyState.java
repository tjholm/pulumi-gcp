// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.networksecurity.inputs.ClientTlsPolicyClientCertificateArgs;
import com.pulumi.gcp.networksecurity.inputs.ClientTlsPolicyServerValidationCaArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientTlsPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final ClientTlsPolicyState Empty = new ClientTlsPolicyState();

    /**
     * Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     * Structure is documented below.
     * 
     */
    @Import(name="clientCertificate")
    private @Nullable Output<ClientTlsPolicyClientCertificateArgs> clientCertificate;

    /**
     * @return Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ClientTlsPolicyClientCertificateArgs>> clientCertificate() {
        return Optional.ofNullable(this.clientCertificate);
    }

    /**
     * Time the ClientTlsPolicy was created in UTC.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Time the ClientTlsPolicy was created in UTC.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * A free-text description of the resource. Max length 1024 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A free-text description of the resource. Max length 1024 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Set of label tags associated with the ClientTlsPolicy resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the ClientTlsPolicy resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location of the client tls policy.
     * The default value is `global`.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the client tls policy.
     * The default value is `global`.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Name of the ClientTlsPolicy resource.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the ClientTlsPolicy resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     * Structure is documented below.
     * 
     */
    @Import(name="serverValidationCas")
    private @Nullable Output<List<ClientTlsPolicyServerValidationCaArgs>> serverValidationCas;

    /**
     * @return Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<ClientTlsPolicyServerValidationCaArgs>>> serverValidationCas() {
        return Optional.ofNullable(this.serverValidationCas);
    }

    /**
     * Server Name Indication string to present to the server during TLS handshake. E.g: &#34;secure.example.com&#34;.
     * 
     */
    @Import(name="sni")
    private @Nullable Output<String> sni;

    /**
     * @return Server Name Indication string to present to the server during TLS handshake. E.g: &#34;secure.example.com&#34;.
     * 
     */
    public Optional<Output<String>> sni() {
        return Optional.ofNullable(this.sni);
    }

    /**
     * Time the ClientTlsPolicy was updated in UTC.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Time the ClientTlsPolicy was updated in UTC.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private ClientTlsPolicyState() {}

    private ClientTlsPolicyState(ClientTlsPolicyState $) {
        this.clientCertificate = $.clientCertificate;
        this.createTime = $.createTime;
        this.description = $.description;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.serverValidationCas = $.serverValidationCas;
        this.sni = $.sni;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientTlsPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientTlsPolicyState $;

        public Builder() {
            $ = new ClientTlsPolicyState();
        }

        public Builder(ClientTlsPolicyState defaults) {
            $ = new ClientTlsPolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientCertificate Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(@Nullable Output<ClientTlsPolicyClientCertificateArgs> clientCertificate) {
            $.clientCertificate = clientCertificate;
            return this;
        }

        /**
         * @param clientCertificate Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder clientCertificate(ClientTlsPolicyClientCertificateArgs clientCertificate) {
            return clientCertificate(Output.of(clientCertificate));
        }

        /**
         * @param createTime Time the ClientTlsPolicy was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Time the ClientTlsPolicy was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param labels Set of label tags associated with the ClientTlsPolicy resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Set of label tags associated with the ClientTlsPolicy resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location of the client tls policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the client tls policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name Name of the ClientTlsPolicy resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the ClientTlsPolicy resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param serverValidationCas Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serverValidationCas(@Nullable Output<List<ClientTlsPolicyServerValidationCaArgs>> serverValidationCas) {
            $.serverValidationCas = serverValidationCas;
            return this;
        }

        /**
         * @param serverValidationCas Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serverValidationCas(List<ClientTlsPolicyServerValidationCaArgs> serverValidationCas) {
            return serverValidationCas(Output.of(serverValidationCas));
        }

        /**
         * @param serverValidationCas Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder serverValidationCas(ClientTlsPolicyServerValidationCaArgs... serverValidationCas) {
            return serverValidationCas(List.of(serverValidationCas));
        }

        /**
         * @param sni Server Name Indication string to present to the server during TLS handshake. E.g: &#34;secure.example.com&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sni(@Nullable Output<String> sni) {
            $.sni = sni;
            return this;
        }

        /**
         * @param sni Server Name Indication string to present to the server during TLS handshake. E.g: &#34;secure.example.com&#34;.
         * 
         * @return builder
         * 
         */
        public Builder sni(String sni) {
            return sni(Output.of(sni));
        }

        /**
         * @param updateTime Time the ClientTlsPolicy was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Time the ClientTlsPolicy was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public ClientTlsPolicyState build() {
            return $;
        }
    }

}
