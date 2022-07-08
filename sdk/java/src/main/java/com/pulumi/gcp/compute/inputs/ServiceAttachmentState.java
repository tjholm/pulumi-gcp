// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.ServiceAttachmentConnectedEndpointArgs;
import com.pulumi.gcp.compute.inputs.ServiceAttachmentConsumerAcceptListArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAttachmentState Empty = new ServiceAttachmentState();

    /**
     * An array of the consumer forwarding rules connected to this service attachment.
     * 
     */
    @Import(name="connectedEndpoints")
    private @Nullable Output<List<ServiceAttachmentConnectedEndpointArgs>> connectedEndpoints;

    /**
     * @return An array of the consumer forwarding rules connected to this service attachment.
     * 
     */
    public Optional<Output<List<ServiceAttachmentConnectedEndpointArgs>>> connectedEndpoints() {
        return Optional.ofNullable(this.connectedEndpoints);
    }

    /**
     * The connection preference to use for this service attachment. Valid
     * values include &#34;ACCEPT_AUTOMATIC&#34;, &#34;ACCEPT_MANUAL&#34;.
     * 
     */
    @Import(name="connectionPreference")
    private @Nullable Output<String> connectionPreference;

    /**
     * @return The connection preference to use for this service attachment. Valid
     * values include &#34;ACCEPT_AUTOMATIC&#34;, &#34;ACCEPT_MANUAL&#34;.
     * 
     */
    public Optional<Output<String>> connectionPreference() {
        return Optional.ofNullable(this.connectionPreference);
    }

    /**
     * An array of projects that are allowed to connect to this service
     * attachment.
     * Structure is documented below.
     * 
     */
    @Import(name="consumerAcceptLists")
    private @Nullable Output<List<ServiceAttachmentConsumerAcceptListArgs>> consumerAcceptLists;

    /**
     * @return An array of projects that are allowed to connect to this service
     * attachment.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<ServiceAttachmentConsumerAcceptListArgs>>> consumerAcceptLists() {
        return Optional.ofNullable(this.consumerAcceptLists);
    }

    /**
     * An array of projects that are not allowed to connect to this service
     * attachment.
     * 
     */
    @Import(name="consumerRejectLists")
    private @Nullable Output<List<String>> consumerRejectLists;

    /**
     * @return An array of projects that are not allowed to connect to this service
     * attachment.
     * 
     */
    public Optional<Output<List<String>>> consumerRejectLists() {
        return Optional.ofNullable(this.consumerRejectLists);
    }

    /**
     * An optional description of this resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If specified, the domain name will be used during the integration between
     * the PSC connected endpoints and the Cloud DNS. For example, this is a
     * valid domain name: &#34;p.mycompany.com.&#34;. Current max number of domain names
     * supported is 1.
     * 
     */
    @Import(name="domainNames")
    private @Nullable Output<List<String>> domainNames;

    /**
     * @return If specified, the domain name will be used during the integration between
     * the PSC connected endpoints and the Cloud DNS. For example, this is a
     * valid domain name: &#34;p.mycompany.com.&#34;. Current max number of domain names
     * supported is 1.
     * 
     */
    public Optional<Output<List<String>>> domainNames() {
        return Optional.ofNullable(this.domainNames);
    }

    /**
     * If true, enable the proxy protocol which is for supplying client TCP/IP
     * address data in TCP connections that traverse proxies on their way to
     * destination servers.
     * 
     */
    @Import(name="enableProxyProtocol")
    private @Nullable Output<Boolean> enableProxyProtocol;

    /**
     * @return If true, enable the proxy protocol which is for supplying client TCP/IP
     * address data in TCP connections that traverse proxies on their way to
     * destination servers.
     * 
     */
    public Optional<Output<Boolean>> enableProxyProtocol() {
        return Optional.ofNullable(this.enableProxyProtocol);
    }

    /**
     * Fingerprint of this resource. This field is used internally during updates of this resource.
     * 
     */
    @Import(name="fingerprint")
    private @Nullable Output<String> fingerprint;

    /**
     * @return Fingerprint of this resource. This field is used internally during updates of this resource.
     * 
     */
    public Optional<Output<String>> fingerprint() {
        return Optional.ofNullable(this.fingerprint);
    }

    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * An array of subnets that is provided for NAT in this service attachment.
     * 
     */
    @Import(name="natSubnets")
    private @Nullable Output<List<String>> natSubnets;

    /**
     * @return An array of subnets that is provided for NAT in this service attachment.
     * 
     */
    public Optional<Output<List<String>>> natSubnets() {
        return Optional.ofNullable(this.natSubnets);
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
     * URL of the region where the resource resides.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return URL of the region where the resource resides.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The URI of the created resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * The URL of a forwarding rule that represents the service identified by
     * this service attachment.
     * 
     */
    @Import(name="targetService")
    private @Nullable Output<String> targetService;

    /**
     * @return The URL of a forwarding rule that represents the service identified by
     * this service attachment.
     * 
     */
    public Optional<Output<String>> targetService() {
        return Optional.ofNullable(this.targetService);
    }

    private ServiceAttachmentState() {}

    private ServiceAttachmentState(ServiceAttachmentState $) {
        this.connectedEndpoints = $.connectedEndpoints;
        this.connectionPreference = $.connectionPreference;
        this.consumerAcceptLists = $.consumerAcceptLists;
        this.consumerRejectLists = $.consumerRejectLists;
        this.description = $.description;
        this.domainNames = $.domainNames;
        this.enableProxyProtocol = $.enableProxyProtocol;
        this.fingerprint = $.fingerprint;
        this.name = $.name;
        this.natSubnets = $.natSubnets;
        this.project = $.project;
        this.region = $.region;
        this.selfLink = $.selfLink;
        this.targetService = $.targetService;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAttachmentState $;

        public Builder() {
            $ = new ServiceAttachmentState();
        }

        public Builder(ServiceAttachmentState defaults) {
            $ = new ServiceAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectedEndpoints An array of the consumer forwarding rules connected to this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder connectedEndpoints(@Nullable Output<List<ServiceAttachmentConnectedEndpointArgs>> connectedEndpoints) {
            $.connectedEndpoints = connectedEndpoints;
            return this;
        }

        /**
         * @param connectedEndpoints An array of the consumer forwarding rules connected to this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder connectedEndpoints(List<ServiceAttachmentConnectedEndpointArgs> connectedEndpoints) {
            return connectedEndpoints(Output.of(connectedEndpoints));
        }

        /**
         * @param connectedEndpoints An array of the consumer forwarding rules connected to this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder connectedEndpoints(ServiceAttachmentConnectedEndpointArgs... connectedEndpoints) {
            return connectedEndpoints(List.of(connectedEndpoints));
        }

        /**
         * @param connectionPreference The connection preference to use for this service attachment. Valid
         * values include &#34;ACCEPT_AUTOMATIC&#34;, &#34;ACCEPT_MANUAL&#34;.
         * 
         * @return builder
         * 
         */
        public Builder connectionPreference(@Nullable Output<String> connectionPreference) {
            $.connectionPreference = connectionPreference;
            return this;
        }

        /**
         * @param connectionPreference The connection preference to use for this service attachment. Valid
         * values include &#34;ACCEPT_AUTOMATIC&#34;, &#34;ACCEPT_MANUAL&#34;.
         * 
         * @return builder
         * 
         */
        public Builder connectionPreference(String connectionPreference) {
            return connectionPreference(Output.of(connectionPreference));
        }

        /**
         * @param consumerAcceptLists An array of projects that are allowed to connect to this service
         * attachment.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder consumerAcceptLists(@Nullable Output<List<ServiceAttachmentConsumerAcceptListArgs>> consumerAcceptLists) {
            $.consumerAcceptLists = consumerAcceptLists;
            return this;
        }

        /**
         * @param consumerAcceptLists An array of projects that are allowed to connect to this service
         * attachment.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder consumerAcceptLists(List<ServiceAttachmentConsumerAcceptListArgs> consumerAcceptLists) {
            return consumerAcceptLists(Output.of(consumerAcceptLists));
        }

        /**
         * @param consumerAcceptLists An array of projects that are allowed to connect to this service
         * attachment.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder consumerAcceptLists(ServiceAttachmentConsumerAcceptListArgs... consumerAcceptLists) {
            return consumerAcceptLists(List.of(consumerAcceptLists));
        }

        /**
         * @param consumerRejectLists An array of projects that are not allowed to connect to this service
         * attachment.
         * 
         * @return builder
         * 
         */
        public Builder consumerRejectLists(@Nullable Output<List<String>> consumerRejectLists) {
            $.consumerRejectLists = consumerRejectLists;
            return this;
        }

        /**
         * @param consumerRejectLists An array of projects that are not allowed to connect to this service
         * attachment.
         * 
         * @return builder
         * 
         */
        public Builder consumerRejectLists(List<String> consumerRejectLists) {
            return consumerRejectLists(Output.of(consumerRejectLists));
        }

        /**
         * @param consumerRejectLists An array of projects that are not allowed to connect to this service
         * attachment.
         * 
         * @return builder
         * 
         */
        public Builder consumerRejectLists(String... consumerRejectLists) {
            return consumerRejectLists(List.of(consumerRejectLists));
        }

        /**
         * @param description An optional description of this resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description of this resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param domainNames If specified, the domain name will be used during the integration between
         * the PSC connected endpoints and the Cloud DNS. For example, this is a
         * valid domain name: &#34;p.mycompany.com.&#34;. Current max number of domain names
         * supported is 1.
         * 
         * @return builder
         * 
         */
        public Builder domainNames(@Nullable Output<List<String>> domainNames) {
            $.domainNames = domainNames;
            return this;
        }

        /**
         * @param domainNames If specified, the domain name will be used during the integration between
         * the PSC connected endpoints and the Cloud DNS. For example, this is a
         * valid domain name: &#34;p.mycompany.com.&#34;. Current max number of domain names
         * supported is 1.
         * 
         * @return builder
         * 
         */
        public Builder domainNames(List<String> domainNames) {
            return domainNames(Output.of(domainNames));
        }

        /**
         * @param domainNames If specified, the domain name will be used during the integration between
         * the PSC connected endpoints and the Cloud DNS. For example, this is a
         * valid domain name: &#34;p.mycompany.com.&#34;. Current max number of domain names
         * supported is 1.
         * 
         * @return builder
         * 
         */
        public Builder domainNames(String... domainNames) {
            return domainNames(List.of(domainNames));
        }

        /**
         * @param enableProxyProtocol If true, enable the proxy protocol which is for supplying client TCP/IP
         * address data in TCP connections that traverse proxies on their way to
         * destination servers.
         * 
         * @return builder
         * 
         */
        public Builder enableProxyProtocol(@Nullable Output<Boolean> enableProxyProtocol) {
            $.enableProxyProtocol = enableProxyProtocol;
            return this;
        }

        /**
         * @param enableProxyProtocol If true, enable the proxy protocol which is for supplying client TCP/IP
         * address data in TCP connections that traverse proxies on their way to
         * destination servers.
         * 
         * @return builder
         * 
         */
        public Builder enableProxyProtocol(Boolean enableProxyProtocol) {
            return enableProxyProtocol(Output.of(enableProxyProtocol));
        }

        /**
         * @param fingerprint Fingerprint of this resource. This field is used internally during updates of this resource.
         * 
         * @return builder
         * 
         */
        public Builder fingerprint(@Nullable Output<String> fingerprint) {
            $.fingerprint = fingerprint;
            return this;
        }

        /**
         * @param fingerprint Fingerprint of this resource. This field is used internally during updates of this resource.
         * 
         * @return builder
         * 
         */
        public Builder fingerprint(String fingerprint) {
            return fingerprint(Output.of(fingerprint));
        }

        /**
         * @param name Name of the resource. The name must be 1-63 characters long, and
         * comply with RFC1035. Specifically, the name must be 1-63 characters
         * long and match the regular expression `a-z?`
         * which means the first character must be a lowercase letter, and all
         * following characters must be a dash, lowercase letter, or digit,
         * except the last character, which cannot be a dash.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource. The name must be 1-63 characters long, and
         * comply with RFC1035. Specifically, the name must be 1-63 characters
         * long and match the regular expression `a-z?`
         * which means the first character must be a lowercase letter, and all
         * following characters must be a dash, lowercase letter, or digit,
         * except the last character, which cannot be a dash.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param natSubnets An array of subnets that is provided for NAT in this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder natSubnets(@Nullable Output<List<String>> natSubnets) {
            $.natSubnets = natSubnets;
            return this;
        }

        /**
         * @param natSubnets An array of subnets that is provided for NAT in this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder natSubnets(List<String> natSubnets) {
            return natSubnets(Output.of(natSubnets));
        }

        /**
         * @param natSubnets An array of subnets that is provided for NAT in this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder natSubnets(String... natSubnets) {
            return natSubnets(List.of(natSubnets));
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
         * @param region URL of the region where the resource resides.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region URL of the region where the resource resides.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink The URI of the created resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param targetService The URL of a forwarding rule that represents the service identified by
         * this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder targetService(@Nullable Output<String> targetService) {
            $.targetService = targetService;
            return this;
        }

        /**
         * @param targetService The URL of a forwarding rule that represents the service identified by
         * this service attachment.
         * 
         * @return builder
         * 
         */
        public Builder targetService(String targetService) {
            return targetService(Output.of(targetService));
        }

        public ServiceAttachmentState build() {
            return $;
        }
    }

}
