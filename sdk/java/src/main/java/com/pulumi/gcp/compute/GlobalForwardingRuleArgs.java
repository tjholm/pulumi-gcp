// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.GlobalForwardingRuleMetadataFilterArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GlobalForwardingRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final GlobalForwardingRuleArgs Empty = new GlobalForwardingRuleArgs();

    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The IP address that this forwarding rule serves. When a client sends
     * traffic to this IP address, the forwarding rule directs the traffic to
     * the target that you specify in the forwarding rule. The
     * loadBalancingScheme and the forwarding rule&#39;s target determine the
     * type of IP address that you can use. For detailed information, refer
     * to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
     * An address can be specified either by a literal IP address or a
     * reference to an existing Address resource. If you don&#39;t specify a
     * reserved IP address, an ephemeral IP address is assigned.
     * The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
     * that has validateForProxyless field set to true.
     * For Private Service Connect forwarding rules that forward traffic to
     * Google APIs, IP address must be provided.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The IP address that this forwarding rule serves. When a client sends
     * traffic to this IP address, the forwarding rule directs the traffic to
     * the target that you specify in the forwarding rule. The
     * loadBalancingScheme and the forwarding rule&#39;s target determine the
     * type of IP address that you can use. For detailed information, refer
     * to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
     * An address can be specified either by a literal IP address or a
     * reference to an existing Address resource. If you don&#39;t specify a
     * reserved IP address, an ephemeral IP address is assigned.
     * The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
     * that has validateForProxyless field set to true.
     * For Private Service Connect forwarding rules that forward traffic to
     * Google APIs, IP address must be provided.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * The IP protocol to which this rule applies. When the load balancing scheme is
     * INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
     * global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
     * and addressType of INTERNAL
     * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
     * 
     */
    @Import(name="ipProtocol")
    private @Nullable Output<String> ipProtocol;

    /**
     * @return The IP protocol to which this rule applies. When the load balancing scheme is
     * INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
     * global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
     * and addressType of INTERNAL
     * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
     * 
     */
    public Optional<Output<String>> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    /**
     * The IP Version that will be used by this global forwarding rule.
     * Possible values are `IPV4` and `IPV6`.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<String> ipVersion;

    /**
     * @return The IP Version that will be used by this global forwarding rule.
     * Possible values are `IPV4` and `IPV6`.
     * 
     */
    public Optional<Output<String>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * Labels to apply to this rule.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels to apply to this rule.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Specifies the forwarding rule type. * `EXTERNAL` is used for: * Classic Cloud VPN gateways * Protocol forwarding to VMs
     * from an external IP address * The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP *
     * `INTERNAL` is used for: * Protocol forwarding to VMs from an internal IP address * Internal TCP/UDP load balancers *
     * `INTERNAL_MANAGED` is used for: * Internal HTTP(S) load balancers * `INTERNAL_SELF_MANAGED` is used for: * Traffic
     * Director * `EXTERNAL_MANAGED` is used for: * Global external HTTP(S) load balancers For more information about
     * forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values:
     * INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL, EXTERNAL_MANAGED
     * 
     */
    @Import(name="loadBalancingScheme")
    private @Nullable Output<String> loadBalancingScheme;

    /**
     * @return Specifies the forwarding rule type. * `EXTERNAL` is used for: * Classic Cloud VPN gateways * Protocol forwarding to VMs
     * from an external IP address * The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP *
     * `INTERNAL` is used for: * Protocol forwarding to VMs from an internal IP address * Internal TCP/UDP load balancers *
     * `INTERNAL_MANAGED` is used for: * Internal HTTP(S) load balancers * `INTERNAL_SELF_MANAGED` is used for: * Traffic
     * Director * `EXTERNAL_MANAGED` is used for: * Global external HTTP(S) load balancers For more information about
     * forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values:
     * INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL, EXTERNAL_MANAGED
     * 
     */
    public Optional<Output<String>> loadBalancingScheme() {
        return Optional.ofNullable(this.loadBalancingScheme);
    }

    /**
     * Opaque filter criteria used by Loadbalancer to restrict routing
     * configuration to a limited set xDS compliant clients. In their xDS
     * requests to Loadbalancer, xDS clients present node metadata. If a
     * match takes place, the relevant routing configuration is made available
     * to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set
     * to MATCH_ANY, at least one of the filterLabels must match the
     * corresponding label provided in the metadata. If its filterMatchCriteria
     * is set to MATCH_ALL, then all of its filterLabels must match with
     * corresponding labels in the provided metadata.
     * metadataFilters specified here can be overridden by those specified in
     * the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     * 
     */
    @Import(name="metadataFilters")
    private @Nullable Output<List<GlobalForwardingRuleMetadataFilterArgs>> metadataFilters;

    /**
     * @return Opaque filter criteria used by Loadbalancer to restrict routing
     * configuration to a limited set xDS compliant clients. In their xDS
     * requests to Loadbalancer, xDS clients present node metadata. If a
     * match takes place, the relevant routing configuration is made available
     * to those proxies.
     * For each metadataFilter in this list, if its filterMatchCriteria is set
     * to MATCH_ANY, at least one of the filterLabels must match the
     * corresponding label provided in the metadata. If its filterMatchCriteria
     * is set to MATCH_ALL, then all of its filterLabels must match with
     * corresponding labels in the provided metadata.
     * metadataFilters specified here can be overridden by those specified in
     * the UrlMap that this ForwardingRule references.
     * metadataFilters only applies to Loadbalancers that have their
     * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<GlobalForwardingRuleMetadataFilterArgs>>> metadataFilters() {
        return Optional.ofNullable(this.metadataFilters);
    }

    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * This field is not used for external load balancing. For `INTERNAL` and `INTERNAL_SELF_MANAGED` load balancing, this
     * field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not
     * specified, the default network will be used.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return This field is not used for external load balancing. For `INTERNAL` and `INTERNAL_SELF_MANAGED` load balancing, this
     * field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not
     * specified, the default network will be used.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * This field is used along with the target field for TargetHttpProxy,
     * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
     * TargetPool, TargetInstance.
     * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
     * addressed to ports in the specified range will be forwarded to target.
     * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
     * disjoint port ranges.
     * Some types of forwarding target have constraints on the acceptable
     * ports:
     * * TargetHttpProxy: 80, 8080
     * * TargetHttpsProxy: 443
     * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     *   1883, 5222
     * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     *   1883, 5222
     * * TargetVpnGateway: 500, 4500
     * 
     */
    @Import(name="portRange")
    private @Nullable Output<String> portRange;

    /**
     * @return This field is used along with the target field for TargetHttpProxy,
     * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
     * TargetPool, TargetInstance.
     * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
     * addressed to ports in the specified range will be forwarded to target.
     * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
     * disjoint port ranges.
     * Some types of forwarding target have constraints on the acceptable
     * ports:
     * * TargetHttpProxy: 80, 8080
     * * TargetHttpsProxy: 443
     * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     *   1883, 5222
     * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
     *   1883, 5222
     * * TargetVpnGateway: 500, 4500
     * 
     */
    public Optional<Output<String>> portRange() {
        return Optional.ofNullable(this.portRange);
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
     * The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in
     * the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing
     * resource. The forwarded traffic must be of a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load
     * balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.
     * 
     */
    @Import(name="target", required=true)
    private Output<String> target;

    /**
     * @return The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in
     * the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing
     * resource. The forwarded traffic must be of a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load
     * balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    private GlobalForwardingRuleArgs() {}

    private GlobalForwardingRuleArgs(GlobalForwardingRuleArgs $) {
        this.description = $.description;
        this.ipAddress = $.ipAddress;
        this.ipProtocol = $.ipProtocol;
        this.ipVersion = $.ipVersion;
        this.labels = $.labels;
        this.loadBalancingScheme = $.loadBalancingScheme;
        this.metadataFilters = $.metadataFilters;
        this.name = $.name;
        this.network = $.network;
        this.portRange = $.portRange;
        this.project = $.project;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GlobalForwardingRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GlobalForwardingRuleArgs $;

        public Builder() {
            $ = new GlobalForwardingRuleArgs();
        }

        public Builder(GlobalForwardingRuleArgs defaults) {
            $ = new GlobalForwardingRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description An optional description of this resource. Provide this property when
         * you create the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description of this resource. Provide this property when
         * you create the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ipAddress The IP address that this forwarding rule serves. When a client sends
         * traffic to this IP address, the forwarding rule directs the traffic to
         * the target that you specify in the forwarding rule. The
         * loadBalancingScheme and the forwarding rule&#39;s target determine the
         * type of IP address that you can use. For detailed information, refer
         * to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
         * An address can be specified either by a literal IP address or a
         * reference to an existing Address resource. If you don&#39;t specify a
         * reserved IP address, an ephemeral IP address is assigned.
         * The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
         * that has validateForProxyless field set to true.
         * For Private Service Connect forwarding rules that forward traffic to
         * Google APIs, IP address must be provided.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The IP address that this forwarding rule serves. When a client sends
         * traffic to this IP address, the forwarding rule directs the traffic to
         * the target that you specify in the forwarding rule. The
         * loadBalancingScheme and the forwarding rule&#39;s target determine the
         * type of IP address that you can use. For detailed information, refer
         * to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
         * An address can be specified either by a literal IP address or a
         * reference to an existing Address resource. If you don&#39;t specify a
         * reserved IP address, an ephemeral IP address is assigned.
         * The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
         * that has validateForProxyless field set to true.
         * For Private Service Connect forwarding rules that forward traffic to
         * Google APIs, IP address must be provided.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param ipProtocol The IP protocol to which this rule applies. When the load balancing scheme is
         * INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
         * global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
         * and addressType of INTERNAL
         * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(@Nullable Output<String> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param ipProtocol The IP protocol to which this rule applies. When the load balancing scheme is
         * INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
         * global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
         * and addressType of INTERNAL
         * Possible values are `TCP`, `UDP`, `ESP`, `AH`, `SCTP`, and `ICMP`.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(String ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        /**
         * @param ipVersion The IP Version that will be used by this global forwarding rule.
         * Possible values are `IPV4` and `IPV6`.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<String> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion The IP Version that will be used by this global forwarding rule.
         * Possible values are `IPV4` and `IPV6`.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(String ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param labels Labels to apply to this rule.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels to apply to this rule.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param loadBalancingScheme Specifies the forwarding rule type. * `EXTERNAL` is used for: * Classic Cloud VPN gateways * Protocol forwarding to VMs
         * from an external IP address * The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP *
         * `INTERNAL` is used for: * Protocol forwarding to VMs from an internal IP address * Internal TCP/UDP load balancers *
         * `INTERNAL_MANAGED` is used for: * Internal HTTP(S) load balancers * `INTERNAL_SELF_MANAGED` is used for: * Traffic
         * Director * `EXTERNAL_MANAGED` is used for: * Global external HTTP(S) load balancers For more information about
         * forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values:
         * INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL, EXTERNAL_MANAGED
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingScheme(@Nullable Output<String> loadBalancingScheme) {
            $.loadBalancingScheme = loadBalancingScheme;
            return this;
        }

        /**
         * @param loadBalancingScheme Specifies the forwarding rule type. * `EXTERNAL` is used for: * Classic Cloud VPN gateways * Protocol forwarding to VMs
         * from an external IP address * The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP *
         * `INTERNAL` is used for: * Protocol forwarding to VMs from an internal IP address * Internal TCP/UDP load balancers *
         * `INTERNAL_MANAGED` is used for: * Internal HTTP(S) load balancers * `INTERNAL_SELF_MANAGED` is used for: * Traffic
         * Director * `EXTERNAL_MANAGED` is used for: * Global external HTTP(S) load balancers For more information about
         * forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values:
         * INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL, EXTERNAL_MANAGED
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingScheme(String loadBalancingScheme) {
            return loadBalancingScheme(Output.of(loadBalancingScheme));
        }

        /**
         * @param metadataFilters Opaque filter criteria used by Loadbalancer to restrict routing
         * configuration to a limited set xDS compliant clients. In their xDS
         * requests to Loadbalancer, xDS clients present node metadata. If a
         * match takes place, the relevant routing configuration is made available
         * to those proxies.
         * For each metadataFilter in this list, if its filterMatchCriteria is set
         * to MATCH_ANY, at least one of the filterLabels must match the
         * corresponding label provided in the metadata. If its filterMatchCriteria
         * is set to MATCH_ALL, then all of its filterLabels must match with
         * corresponding labels in the provided metadata.
         * metadataFilters specified here can be overridden by those specified in
         * the UrlMap that this ForwardingRule references.
         * metadataFilters only applies to Loadbalancers that have their
         * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metadataFilters(@Nullable Output<List<GlobalForwardingRuleMetadataFilterArgs>> metadataFilters) {
            $.metadataFilters = metadataFilters;
            return this;
        }

        /**
         * @param metadataFilters Opaque filter criteria used by Loadbalancer to restrict routing
         * configuration to a limited set xDS compliant clients. In their xDS
         * requests to Loadbalancer, xDS clients present node metadata. If a
         * match takes place, the relevant routing configuration is made available
         * to those proxies.
         * For each metadataFilter in this list, if its filterMatchCriteria is set
         * to MATCH_ANY, at least one of the filterLabels must match the
         * corresponding label provided in the metadata. If its filterMatchCriteria
         * is set to MATCH_ALL, then all of its filterLabels must match with
         * corresponding labels in the provided metadata.
         * metadataFilters specified here can be overridden by those specified in
         * the UrlMap that this ForwardingRule references.
         * metadataFilters only applies to Loadbalancers that have their
         * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metadataFilters(List<GlobalForwardingRuleMetadataFilterArgs> metadataFilters) {
            return metadataFilters(Output.of(metadataFilters));
        }

        /**
         * @param metadataFilters Opaque filter criteria used by Loadbalancer to restrict routing
         * configuration to a limited set xDS compliant clients. In their xDS
         * requests to Loadbalancer, xDS clients present node metadata. If a
         * match takes place, the relevant routing configuration is made available
         * to those proxies.
         * For each metadataFilter in this list, if its filterMatchCriteria is set
         * to MATCH_ANY, at least one of the filterLabels must match the
         * corresponding label provided in the metadata. If its filterMatchCriteria
         * is set to MATCH_ALL, then all of its filterLabels must match with
         * corresponding labels in the provided metadata.
         * metadataFilters specified here can be overridden by those specified in
         * the UrlMap that this ForwardingRule references.
         * metadataFilters only applies to Loadbalancers that have their
         * loadBalancingScheme set to INTERNAL_SELF_MANAGED.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder metadataFilters(GlobalForwardingRuleMetadataFilterArgs... metadataFilters) {
            return metadataFilters(List.of(metadataFilters));
        }

        /**
         * @param name Name of the resource; provided by the client when the resource is
         * created. The name must be 1-63 characters long, and comply with
         * RFC1035. Specifically, the name must be 1-63 characters long and match
         * the regular expression `a-z?` which means the
         * first character must be a lowercase letter, and all following
         * characters must be a dash, lowercase letter, or digit, except the last
         * character, which cannot be a dash.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource; provided by the client when the resource is
         * created. The name must be 1-63 characters long, and comply with
         * RFC1035. Specifically, the name must be 1-63 characters long and match
         * the regular expression `a-z?` which means the
         * first character must be a lowercase letter, and all following
         * characters must be a dash, lowercase letter, or digit, except the last
         * character, which cannot be a dash.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param network This field is not used for external load balancing. For `INTERNAL` and `INTERNAL_SELF_MANAGED` load balancing, this
         * field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not
         * specified, the default network will be used.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network This field is not used for external load balancing. For `INTERNAL` and `INTERNAL_SELF_MANAGED` load balancing, this
         * field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not
         * specified, the default network will be used.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param portRange This field is used along with the target field for TargetHttpProxy,
         * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
         * TargetPool, TargetInstance.
         * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
         * addressed to ports in the specified range will be forwarded to target.
         * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
         * disjoint port ranges.
         * Some types of forwarding target have constraints on the acceptable
         * ports:
         * * TargetHttpProxy: 80, 8080
         * * TargetHttpsProxy: 443
         * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
         *   1883, 5222
         * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
         *   1883, 5222
         * * TargetVpnGateway: 500, 4500
         * 
         * @return builder
         * 
         */
        public Builder portRange(@Nullable Output<String> portRange) {
            $.portRange = portRange;
            return this;
        }

        /**
         * @param portRange This field is used along with the target field for TargetHttpProxy,
         * TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
         * TargetPool, TargetInstance.
         * Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
         * addressed to ports in the specified range will be forwarded to target.
         * Forwarding rules with the same [IPAddress, IPProtocol] pair must have
         * disjoint port ranges.
         * Some types of forwarding target have constraints on the acceptable
         * ports:
         * * TargetHttpProxy: 80, 8080
         * * TargetHttpsProxy: 443
         * * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
         *   1883, 5222
         * * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
         *   1883, 5222
         * * TargetVpnGateway: 500, 4500
         * 
         * @return builder
         * 
         */
        public Builder portRange(String portRange) {
            return portRange(Output.of(portRange));
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
         * @param target The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in
         * the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing
         * resource. The forwarded traffic must be of a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load
         * balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.
         * 
         * @return builder
         * 
         */
        public Builder target(Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target The URL of the target resource to receive the matched traffic. For regional forwarding rules, this target must live in
         * the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing
         * resource. The forwarded traffic must be of a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load
         * balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        public GlobalForwardingRuleArgs build() {
            $.target = Objects.requireNonNull($.target, "expected parameter 'target' to be non-null");
            return $;
        }
    }

}
