// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DnsAuthorizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsAuthorizationArgs Empty = new DnsAuthorizationArgs();

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
     * A domain which is being authorized. A DnsAuthorization resource covers a
     * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
     * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return A domain which is being authorized. A DnsAuthorization resource covers a
     * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
     * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
     * 
     */
    public Output<String> domain() {
        return this.domain;
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

    private DnsAuthorizationArgs() {}

    private DnsAuthorizationArgs(DnsAuthorizationArgs $) {
        this.description = $.description;
        this.domain = $.domain;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsAuthorizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsAuthorizationArgs $;

        public Builder() {
            $ = new DnsAuthorizationArgs();
        }

        public Builder(DnsAuthorizationArgs defaults) {
            $ = new DnsAuthorizationArgs(Objects.requireNonNull(defaults));
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
         * @param domain A domain which is being authorized. A DnsAuthorization resource covers a
         * single domain and its wildcard, e.g. authorization for &#34;example.com&#34; can
         * be used to issue certificates for &#34;example.com&#34; and &#34;*.example.com&#34;.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
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

        public DnsAuthorizationArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            return $;
        }
    }

}
