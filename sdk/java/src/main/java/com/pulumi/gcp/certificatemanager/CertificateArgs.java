// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
import com.pulumi.gcp.certificatemanager.inputs.CertificateSelfManagedArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateArgs Empty = new CertificateArgs();

    /**
     * A human-readable description of the resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A human-readable description of the resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Set of label tags associated with the Certificate resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the Certificate resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The Certificate Manager location. If not specified, &#34;global&#34; is used.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The Certificate Manager location. If not specified, &#34;global&#34; is used.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it&#39;s authorized to do so.
     * Structure is documented below.
     * 
     */
    @Import(name="managed")
    private @Nullable Output<CertificateManagedArgs> managed;

    /**
     * @return Configuration and state of a Managed Certificate.
     * Certificate Manager provisions and renews Managed Certificates
     * automatically, for as long as it&#39;s authorized to do so.
     * Structure is documented below.
     * 
     */
    public Optional<Output<CertificateManagedArgs>> managed() {
        return Optional.ofNullable(this.managed);
    }

    /**
     * A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A user-defined name of the certificate. Certificate names must be unique
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
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
     * The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
     * served from non-core Google data centers.
     * Currently allowed only for managed certificates.
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return The scope of the certificate.
     * DEFAULT: Certificates with default scope are served from core Google data centers.
     * If unsure, choose this option.
     * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
     * served from non-core Google data centers.
     * Currently allowed only for managed certificates.
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user&#39;s responsibility.
     * Structure is documented below.
     * 
     */
    @Import(name="selfManaged")
    private @Nullable Output<CertificateSelfManagedArgs> selfManaged;

    /**
     * @return Certificate data for a SelfManaged Certificate.
     * SelfManaged Certificates are uploaded by the user. Updating such
     * certificates before they expire remains the user&#39;s responsibility.
     * Structure is documented below.
     * 
     */
    public Optional<Output<CertificateSelfManagedArgs>> selfManaged() {
        return Optional.ofNullable(this.selfManaged);
    }

    private CertificateArgs() {}

    private CertificateArgs(CertificateArgs $) {
        this.description = $.description;
        this.labels = $.labels;
        this.location = $.location;
        this.managed = $.managed;
        this.name = $.name;
        this.project = $.project;
        this.scope = $.scope;
        this.selfManaged = $.selfManaged;
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
         * @param description A human-readable description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A human-readable description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param labels Set of label tags associated with the Certificate resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Set of label tags associated with the Certificate resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The Certificate Manager location. If not specified, &#34;global&#34; is used.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The Certificate Manager location. If not specified, &#34;global&#34; is used.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param managed Configuration and state of a Managed Certificate.
         * Certificate Manager provisions and renews Managed Certificates
         * automatically, for as long as it&#39;s authorized to do so.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder managed(@Nullable Output<CertificateManagedArgs> managed) {
            $.managed = managed;
            return this;
        }

        /**
         * @param managed Configuration and state of a Managed Certificate.
         * Certificate Manager provisions and renews Managed Certificates
         * automatically, for as long as it&#39;s authorized to do so.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder managed(CertificateManagedArgs managed) {
            return managed(Output.of(managed));
        }

        /**
         * @param name A user-defined name of the certificate. Certificate names must be unique
         * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
         * and all following characters must be a dash, underscore, letter or digit.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A user-defined name of the certificate. Certificate names must be unique
         * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
         * and all following characters must be a dash, underscore, letter or digit.
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
         * @param scope The scope of the certificate.
         * DEFAULT: Certificates with default scope are served from core Google data centers.
         * If unsure, choose this option.
         * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
         * served from non-core Google data centers.
         * Currently allowed only for managed certificates.
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The scope of the certificate.
         * DEFAULT: Certificates with default scope are served from core Google data centers.
         * If unsure, choose this option.
         * EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
         * served from non-core Google data centers.
         * Currently allowed only for managed certificates.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param selfManaged Certificate data for a SelfManaged Certificate.
         * SelfManaged Certificates are uploaded by the user. Updating such
         * certificates before they expire remains the user&#39;s responsibility.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder selfManaged(@Nullable Output<CertificateSelfManagedArgs> selfManaged) {
            $.selfManaged = selfManaged;
            return this;
        }

        /**
         * @param selfManaged Certificate data for a SelfManaged Certificate.
         * SelfManaged Certificates are uploaded by the user. Updating such
         * certificates before they expire remains the user&#39;s responsibility.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder selfManaged(CertificateSelfManagedArgs selfManaged) {
            return selfManaged(Output.of(selfManaged));
        }

        public CertificateArgs build() {
            return $;
        }
    }

}
