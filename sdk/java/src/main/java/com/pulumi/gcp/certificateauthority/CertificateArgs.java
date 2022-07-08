// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.certificateauthority.inputs.CertificateConfigArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateArgs Empty = new CertificateArgs();

    /**
     * Certificate Authority name.
     * 
     */
    @Import(name="certificateAuthority")
    private @Nullable Output<String> certificateAuthority;

    /**
     * @return Certificate Authority name.
     * 
     */
    public Optional<Output<String>> certificateAuthority() {
        return Optional.ofNullable(this.certificateAuthority);
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
     * Labels with user-defined metadata to apply to this resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata to apply to this resource.
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
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return Location of the Certificate. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * 
     */
    public Output<String> location() {
        return this.location;
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
    @Import(name="pool", required=true)
    private Output<String> pool;

    /**
     * @return The name of the CaPool this Certificate belongs to.
     * 
     */
    public Output<String> pool() {
        return this.pool;
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

    private CertificateArgs() {}

    private CertificateArgs(CertificateArgs $) {
        this.certificateAuthority = $.certificateAuthority;
        this.certificateTemplate = $.certificateTemplate;
        this.config = $.config;
        this.labels = $.labels;
        this.lifetime = $.lifetime;
        this.location = $.location;
        this.name = $.name;
        this.pemCsr = $.pemCsr;
        this.pool = $.pool;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateArgs $;

        public Builder() {
            $ = new CertificateArgs();
        }

        public Builder(CertificateArgs defaults) {
            $ = new CertificateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificateAuthority Certificate Authority name.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(@Nullable Output<String> certificateAuthority) {
            $.certificateAuthority = certificateAuthority;
            return this;
        }

        /**
         * @param certificateAuthority Certificate Authority name.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(String certificateAuthority) {
            return certificateAuthority(Output.of(certificateAuthority));
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
         * @param labels Labels with user-defined metadata to apply to this resource.
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
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Location of the Certificate. A full list of valid locations can be found by
         * running `gcloud privateca locations list`.
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
        public Builder pool(Output<String> pool) {
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

        public CertificateArgs build() {
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            $.pool = Objects.requireNonNull($.pool, "expected parameter 'pool' to be non-null");
            return $;
        }
    }

}
