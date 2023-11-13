// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.certificateauthority.inputs.CertificateCertificateDescriptionArgs;
import com.pulumi.gcp.certificateauthority.inputs.CertificateConfigArgs;
import com.pulumi.gcp.certificateauthority.inputs.CertificateRevocationDetailArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateState extends com.pulumi.resources.ResourceArgs {

    public static final CertificateState Empty = new CertificateState();

    /**
     * The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
     * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
     * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificate_authority`
     * should be set to `my-ca`.
     * 
     */
    @Import(name="certificateAuthority")
    private @Nullable Output<String> certificateAuthority;

    /**
     * @return The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
     * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
     * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificate_authority`
     * should be set to `my-ca`.
     * 
     */
    public Optional<Output<String>> certificateAuthority() {
        return Optional.ofNullable(this.certificateAuthority);
    }

    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
     * Structure is documented below.
     * 
     */
    @Import(name="certificateDescriptions")
    private @Nullable Output<List<CertificateCertificateDescriptionArgs>> certificateDescriptions;

    /**
     * @return Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<CertificateCertificateDescriptionArgs>>> certificateDescriptions() {
        return Optional.ofNullable(this.certificateDescriptions);
    }

    /**
     * The resource name for a CertificateTemplate used to issue this certificate,
     * in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`. If this is specified,
     * the caller must have the necessary permission to use this template. If this is
     * omitted, no template will be used. This template must be in the same location
     * as the Certificate.
     * 
     */
    @Import(name="certificateTemplate")
    private @Nullable Output<String> certificateTemplate;

    /**
     * @return The resource name for a CertificateTemplate used to issue this certificate,
     * in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`. If this is specified,
     * the caller must have the necessary permission to use this template. If this is
     * omitted, no template will be used. This template must be in the same location
     * as the Certificate.
     * 
     */
    public Optional<Output<String>> certificateTemplate() {
        return Optional.ofNullable(this.certificateTemplate);
    }

    /**
     * The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     * 
     */
    @Import(name="config")
    private @Nullable Output<CertificateConfigArgs> config;

    /**
     * @return The config used to create a self-signed X.509 certificate or CSR.
     * Structure is documented below.
     * 
     */
    public Optional<Output<CertificateConfigArgs>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The time that this resource was created on the server.
     * This is in RFC3339 text format.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time that this resource was created on the server.
     * This is in RFC3339 text format.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * The resource name of the issuing CertificateAuthority in the format `projects/*{@literal /}locations/*{@literal /}caPools/*{@literal /}certificateAuthorities/*`.
     * 
     */
    @Import(name="issuerCertificateAuthority")
    private @Nullable Output<String> issuerCertificateAuthority;

    /**
     * @return The resource name of the issuing CertificateAuthority in the format `projects/*{@literal /}locations/*{@literal /}caPools/*{@literal /}certificateAuthorities/*`.
     * 
     */
    public Optional<Output<String>> issuerCertificateAuthority() {
        return Optional.ofNullable(this.issuerCertificateAuthority);
    }

    /**
     * Labels with user-defined metadata to apply to this resource.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata to apply to this resource.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
     * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Import(name="lifetime")
    private @Nullable Output<String> lifetime;

    /**
     * @return The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
     * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
     * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<Output<String>> lifetime() {
        return Optional.ofNullable(this.lifetime);
    }

    /**
     * Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     * ***
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The name for this Certificate.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name for this Certificate.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Output only. The pem-encoded, signed X.509 certificate.
     * 
     */
    @Import(name="pemCertificate")
    private @Nullable Output<String> pemCertificate;

    /**
     * @return Output only. The pem-encoded, signed X.509 certificate.
     * 
     */
    public Optional<Output<String>> pemCertificate() {
        return Optional.ofNullable(this.pemCertificate);
    }

    /**
     * The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
     * 
     */
    @Import(name="pemCertificateChains")
    private @Nullable Output<List<String>> pemCertificateChains;

    /**
     * @return The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
     * 
     */
    public Optional<Output<List<String>>> pemCertificateChains() {
        return Optional.ofNullable(this.pemCertificateChains);
    }

    /**
     * Immutable. A pem-encoded X.509 certificate signing request (CSR).
     * 
     */
    @Import(name="pemCsr")
    private @Nullable Output<String> pemCsr;

    /**
     * @return Immutable. A pem-encoded X.509 certificate signing request (CSR).
     * 
     */
    public Optional<Output<String>> pemCsr() {
        return Optional.ofNullable(this.pemCsr);
    }

    /**
     * The name of the CaPool this Certificate belongs to.
     * 
     */
    @Import(name="pool")
    private @Nullable Output<String> pool;

    /**
     * @return The name of the CaPool this Certificate belongs to.
     * 
     */
    public Optional<Output<String>> pool() {
        return Optional.ofNullable(this.pool);
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
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * Output only. Details regarding the revocation of this Certificate. This Certificate is
     * considered revoked if and only if this field is present.
     * Structure is documented below.
     * 
     */
    @Import(name="revocationDetails")
    private @Nullable Output<List<CertificateRevocationDetailArgs>> revocationDetails;

    /**
     * @return Output only. Details regarding the revocation of this Certificate. This Certificate is
     * considered revoked if and only if this field is present.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<CertificateRevocationDetailArgs>>> revocationDetails() {
        return Optional.ofNullable(this.revocationDetails);
    }

    /**
     * Output only. The time at which this CertificateAuthority was updated.
     * This is in RFC3339 text format.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Output only. The time at which this CertificateAuthority was updated.
     * This is in RFC3339 text format.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private CertificateState() {}

    private CertificateState(CertificateState $) {
        this.certificateAuthority = $.certificateAuthority;
        this.certificateDescriptions = $.certificateDescriptions;
        this.certificateTemplate = $.certificateTemplate;
        this.config = $.config;
        this.createTime = $.createTime;
        this.effectiveLabels = $.effectiveLabels;
        this.issuerCertificateAuthority = $.issuerCertificateAuthority;
        this.labels = $.labels;
        this.lifetime = $.lifetime;
        this.location = $.location;
        this.name = $.name;
        this.pemCertificate = $.pemCertificate;
        this.pemCertificateChains = $.pemCertificateChains;
        this.pemCsr = $.pemCsr;
        this.pool = $.pool;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.revocationDetails = $.revocationDetails;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateState $;

        public Builder() {
            $ = new CertificateState();
        }

        public Builder(CertificateState defaults) {
            $ = new CertificateState(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateAuthority The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
         * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
         * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificate_authority`
         * should be set to `my-ca`.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(@Nullable Output<String> certificateAuthority) {
            $.certificateAuthority = certificateAuthority;
            return this;
        }

        /**
         * @param certificateAuthority The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
         * a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
         * argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificate_authority`
         * should be set to `my-ca`.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(String certificateAuthority) {
            return certificateAuthority(Output.of(certificateAuthority));
        }

        /**
         * @param certificateDescriptions Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificateDescriptions(@Nullable Output<List<CertificateCertificateDescriptionArgs>> certificateDescriptions) {
            $.certificateDescriptions = certificateDescriptions;
            return this;
        }

        /**
         * @param certificateDescriptions Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificateDescriptions(List<CertificateCertificateDescriptionArgs> certificateDescriptions) {
            return certificateDescriptions(Output.of(certificateDescriptions));
        }

        /**
         * @param certificateDescriptions Output only. Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder certificateDescriptions(CertificateCertificateDescriptionArgs... certificateDescriptions) {
            return certificateDescriptions(List.of(certificateDescriptions));
        }

        /**
         * @param certificateTemplate The resource name for a CertificateTemplate used to issue this certificate,
         * in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`. If this is specified,
         * the caller must have the necessary permission to use this template. If this is
         * omitted, no template will be used. This template must be in the same location
         * as the Certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateTemplate(@Nullable Output<String> certificateTemplate) {
            $.certificateTemplate = certificateTemplate;
            return this;
        }

        /**
         * @param certificateTemplate The resource name for a CertificateTemplate used to issue this certificate,
         * in the format `projects/*{@literal /}locations/*{@literal /}certificateTemplates/*`. If this is specified,
         * the caller must have the necessary permission to use this template. If this is
         * omitted, no template will be used. This template must be in the same location
         * as the Certificate.
         * 
         * @return builder
         * 
         */
        public Builder certificateTemplate(String certificateTemplate) {
            return certificateTemplate(Output.of(certificateTemplate));
        }

        /**
         * @param config The config used to create a self-signed X.509 certificate or CSR.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<CertificateConfigArgs> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The config used to create a self-signed X.509 certificate or CSR.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder config(CertificateConfigArgs config) {
            return config(Output.of(config));
        }

        /**
         * @param createTime The time that this resource was created on the server.
         * This is in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time that this resource was created on the server.
         * This is in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param issuerCertificateAuthority The resource name of the issuing CertificateAuthority in the format `projects/*{@literal /}locations/*{@literal /}caPools/*{@literal /}certificateAuthorities/*`.
         * 
         * @return builder
         * 
         */
        public Builder issuerCertificateAuthority(@Nullable Output<String> issuerCertificateAuthority) {
            $.issuerCertificateAuthority = issuerCertificateAuthority;
            return this;
        }

        /**
         * @param issuerCertificateAuthority The resource name of the issuing CertificateAuthority in the format `projects/*{@literal /}locations/*{@literal /}caPools/*{@literal /}certificateAuthorities/*`.
         * 
         * @return builder
         * 
         */
        public Builder issuerCertificateAuthority(String issuerCertificateAuthority) {
            return issuerCertificateAuthority(Output.of(issuerCertificateAuthority));
        }

        /**
         * @param labels Labels with user-defined metadata to apply to this resource.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels with user-defined metadata to apply to this resource.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param lifetime The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
         * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
         * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder lifetime(@Nullable Output<String> lifetime) {
            $.lifetime = lifetime;
            return this;
        }

        /**
         * @param lifetime The desired lifetime of the CA certificate. Used to create the &#34;notBeforeTime&#34; and
         * &#34;notAfterTime&#34; fields inside an X.509 certificate. A duration in seconds with up to nine
         * fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder lifetime(String lifetime) {
            return lifetime(Output.of(lifetime));
        }

        /**
         * @param location Location of the Certificate. A full list of valid locations can be found by
         * running `gcloud privateca locations list`.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Location of the Certificate. A full list of valid locations can be found by
         * running `gcloud privateca locations list`.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name The name for this Certificate.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name for this Certificate.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pemCertificate Output only. The pem-encoded, signed X.509 certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemCertificate(@Nullable Output<String> pemCertificate) {
            $.pemCertificate = pemCertificate;
            return this;
        }

        /**
         * @param pemCertificate Output only. The pem-encoded, signed X.509 certificate.
         * 
         * @return builder
         * 
         */
        public Builder pemCertificate(String pemCertificate) {
            return pemCertificate(Output.of(pemCertificate));
        }

        /**
         * @param pemCertificateChains The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
         * 
         * @return builder
         * 
         */
        public Builder pemCertificateChains(@Nullable Output<List<String>> pemCertificateChains) {
            $.pemCertificateChains = pemCertificateChains;
            return this;
        }

        /**
         * @param pemCertificateChains The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
         * 
         * @return builder
         * 
         */
        public Builder pemCertificateChains(List<String> pemCertificateChains) {
            return pemCertificateChains(Output.of(pemCertificateChains));
        }

        /**
         * @param pemCertificateChains The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
         * 
         * @return builder
         * 
         */
        public Builder pemCertificateChains(String... pemCertificateChains) {
            return pemCertificateChains(List.of(pemCertificateChains));
        }

        /**
         * @param pemCsr Immutable. A pem-encoded X.509 certificate signing request (CSR).
         * 
         * @return builder
         * 
         */
        public Builder pemCsr(@Nullable Output<String> pemCsr) {
            $.pemCsr = pemCsr;
            return this;
        }

        /**
         * @param pemCsr Immutable. A pem-encoded X.509 certificate signing request (CSR).
         * 
         * @return builder
         * 
         */
        public Builder pemCsr(String pemCsr) {
            return pemCsr(Output.of(pemCsr));
        }

        /**
         * @param pool The name of the CaPool this Certificate belongs to.
         * 
         * @return builder
         * 
         */
        public Builder pool(@Nullable Output<String> pool) {
            $.pool = pool;
            return this;
        }

        /**
         * @param pool The name of the CaPool this Certificate belongs to.
         * 
         * @return builder
         * 
         */
        public Builder pool(String pool) {
            return pool(Output.of(pool));
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
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param revocationDetails Output only. Details regarding the revocation of this Certificate. This Certificate is
         * considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder revocationDetails(@Nullable Output<List<CertificateRevocationDetailArgs>> revocationDetails) {
            $.revocationDetails = revocationDetails;
            return this;
        }

        /**
         * @param revocationDetails Output only. Details regarding the revocation of this Certificate. This Certificate is
         * considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder revocationDetails(List<CertificateRevocationDetailArgs> revocationDetails) {
            return revocationDetails(Output.of(revocationDetails));
        }

        /**
         * @param revocationDetails Output only. Details regarding the revocation of this Certificate. This Certificate is
         * considered revoked if and only if this field is present.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder revocationDetails(CertificateRevocationDetailArgs... revocationDetails) {
            return revocationDetails(List.of(revocationDetails));
        }

        /**
         * @param updateTime Output only. The time at which this CertificateAuthority was updated.
         * This is in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Output only. The time at which this CertificateAuthority was updated.
         * This is in RFC3339 text format.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public CertificateState build() {
            return $;
        }
    }

}
