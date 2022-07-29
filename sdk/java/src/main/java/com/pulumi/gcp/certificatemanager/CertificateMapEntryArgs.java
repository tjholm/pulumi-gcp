// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertificateMapEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertificateMapEntryArgs Empty = new CertificateMapEntryArgs();

    /**
     * A set of Certificates defines for the given hostname.
     * There can be defined up to fifteen certificates in each Certificate Map Entry.
     * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
     * 
     */
    @Import(name="certificates", required=true)
    private Output<List<String>> certificates;

    /**
     * @return A set of Certificates defines for the given hostname.
     * There can be defined up to fifteen certificates in each Certificate Map Entry.
     * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
     * 
     */
    public Output<List<String>> certificates() {
        return this.certificates;
    }

    /**
     * CertificateMapEntry is a list of certificate configurations,
     * that have been issued for a particular hostname
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return CertificateMapEntry is a list of certificate configurations,
     * that have been issued for a particular hostname
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
     * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
     * selecting a proper certificate.
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
     * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
     * selecting a proper certificate.
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * Set of labels associated with a Certificate Map Entry.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Set of labels associated with a Certificate Map Entry.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * A map entry that is inputted into the cetrificate map
     * 
     */
    @Import(name="map", required=true)
    private Output<String> map;

    /**
     * @return A map entry that is inputted into the cetrificate map
     * 
     */
    public Output<String> map() {
        return this.map;
    }

    /**
     * A predefined matcher for particular cases, other than SNI selection
     * Possible values are `MATCHER_UNSPECIFIED` and `PRIMARY`.
     * 
     */
    @Import(name="matcher")
    private @Nullable Output<String> matcher;

    /**
     * @return A predefined matcher for particular cases, other than SNI selection
     * Possible values are `MATCHER_UNSPECIFIED` and `PRIMARY`.
     * 
     */
    public Optional<Output<String>> matcher() {
        return Optional.ofNullable(this.matcher);
    }

    /**
     * A user-defined name of the Certificate Map Entry. Certificate Map Entry
     * names must be unique globally and match pattern
     * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A user-defined name of the Certificate Map Entry. Certificate Map Entry
     * names must be unique globally and match pattern
     * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
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

    private CertificateMapEntryArgs() {}

    private CertificateMapEntryArgs(CertificateMapEntryArgs $) {
        this.certificates = $.certificates;
        this.description = $.description;
        this.hostname = $.hostname;
        this.labels = $.labels;
        this.map = $.map;
        this.matcher = $.matcher;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertificateMapEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertificateMapEntryArgs $;

        public Builder() {
            $ = new CertificateMapEntryArgs();
        }

        public Builder(CertificateMapEntryArgs defaults) {
            $ = new CertificateMapEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificates A set of Certificates defines for the given hostname.
         * There can be defined up to fifteen certificates in each Certificate Map Entry.
         * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
         * 
         * @return builder
         * 
         */
        public Builder certificates(Output<List<String>> certificates) {
            $.certificates = certificates;
            return this;
        }

        /**
         * @param certificates A set of Certificates defines for the given hostname.
         * There can be defined up to fifteen certificates in each Certificate Map Entry.
         * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
         * 
         * @return builder
         * 
         */
        public Builder certificates(List<String> certificates) {
            return certificates(Output.of(certificates));
        }

        /**
         * @param certificates A set of Certificates defines for the given hostname.
         * There can be defined up to fifteen certificates in each Certificate Map Entry.
         * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
         * 
         * @return builder
         * 
         */
        public Builder certificates(String... certificates) {
            return certificates(List.of(certificates));
        }

        /**
         * @param description CertificateMapEntry is a list of certificate configurations,
         * that have been issued for a particular hostname
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description CertificateMapEntry is a list of certificate configurations,
         * that have been issued for a particular hostname
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param hostname A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
         * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
         * selecting a proper certificate.
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
         * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
         * selecting a proper certificate.
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param labels Set of labels associated with a Certificate Map Entry.
         * An object containing a list of &#34;key&#34;: value pairs.
         * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Set of labels associated with a Certificate Map Entry.
         * An object containing a list of &#34;key&#34;: value pairs.
         * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param map A map entry that is inputted into the cetrificate map
         * 
         * @return builder
         * 
         */
        public Builder map(Output<String> map) {
            $.map = map;
            return this;
        }

        /**
         * @param map A map entry that is inputted into the cetrificate map
         * 
         * @return builder
         * 
         */
        public Builder map(String map) {
            return map(Output.of(map));
        }

        /**
         * @param matcher A predefined matcher for particular cases, other than SNI selection
         * Possible values are `MATCHER_UNSPECIFIED` and `PRIMARY`.
         * 
         * @return builder
         * 
         */
        public Builder matcher(@Nullable Output<String> matcher) {
            $.matcher = matcher;
            return this;
        }

        /**
         * @param matcher A predefined matcher for particular cases, other than SNI selection
         * Possible values are `MATCHER_UNSPECIFIED` and `PRIMARY`.
         * 
         * @return builder
         * 
         */
        public Builder matcher(String matcher) {
            return matcher(Output.of(matcher));
        }

        /**
         * @param name A user-defined name of the Certificate Map Entry. Certificate Map Entry
         * names must be unique globally and match pattern
         * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A user-defined name of the Certificate Map Entry. Certificate Map Entry
         * names must be unique globally and match pattern
         * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
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

        public CertificateMapEntryArgs build() {
            $.certificates = Objects.requireNonNull($.certificates, "expected parameter 'certificates' to be non-null");
            $.map = Objects.requireNonNull($.map, "expected parameter 'map' to be non-null");
            return $;
        }
    }

}
