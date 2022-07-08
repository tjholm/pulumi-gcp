// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.certificatemanager.inputs.DnsAuthorizationDnsResourceRecordArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DnsAuthorizationState extends com.pulumi.resources.ResourceArgs {

    public static final DnsAuthorizationState Empty = new DnsAuthorizationState();

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
     * The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
     * usable by certificate.
     * 
     */
    @Import(name="dnsResourceRecords")
    private @Nullable Output<List<DnsAuthorizationDnsResourceRecordArgs>> dnsResourceRecords;

    /**
     * @return The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
     * usable by certificate.
     * 
     */
    public Optional<Output<List<DnsAuthorizationDnsResourceRecordArgs>>> dnsResourceRecords() {
        return Optional.ofNullable(this.dnsResourceRecords);
    }

    /**
     * A domain which is being authorized. A DnsAuthorization resource covers a
     * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
     * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return A domain which is being authorized. A DnsAuthorization resource covers a
     * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
     * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * Set of label tags associated with the EdgeCache resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the EdgeCache resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource; provided by the client when the resource is created.
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

    private DnsAuthorizationState() {}

    private DnsAuthorizationState(DnsAuthorizationState $) {
        this.description = $.description;
        this.dnsResourceRecords = $.dnsResourceRecords;
        this.domain = $.domain;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsAuthorizationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsAuthorizationState $;

        public Builder() {
            $ = new DnsAuthorizationState();
        }

        public Builder(DnsAuthorizationState defaults) {
            $ = new DnsAuthorizationState(Objects.requireNonNull(defaults));
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
         * @param dnsResourceRecords The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
         * usable by certificate.
         * 
         * @return builder
         * 
         */
        public Builder dnsResourceRecords(@Nullable Output<List<DnsAuthorizationDnsResourceRecordArgs>> dnsResourceRecords) {
            $.dnsResourceRecords = dnsResourceRecords;
            return this;
        }

        /**
         * @param dnsResourceRecords The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
         * usable by certificate.
         * 
         * @return builder
         * 
         */
        public Builder dnsResourceRecords(List<DnsAuthorizationDnsResourceRecordArgs> dnsResourceRecords) {
            return dnsResourceRecords(Output.of(dnsResourceRecords));
        }

        /**
         * @param dnsResourceRecords The structure describing the DNS Resource Record that needs to be added to DNS configuration for the authorization to be
         * usable by certificate.
         * 
         * @return builder
         * 
         */
        public Builder dnsResourceRecords(DnsAuthorizationDnsResourceRecordArgs... dnsResourceRecords) {
            return dnsResourceRecords(List.of(dnsResourceRecords));
        }

        /**
         * @param domain A domain which is being authorized. A DnsAuthorization resource covers a
         * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
         * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain A domain which is being authorized. A DnsAuthorization resource covers a
         * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
         * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param labels Set of label tags associated with the EdgeCache resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Set of label tags associated with the EdgeCache resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param name Name of the resource; provided by the client when the resource is created.
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
         * @param name Name of the resource; provided by the client when the resource is created.
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

        public DnsAuthorizationState build() {
            return $;
        }
    }

}
