// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.FirewallArgs;
import com.pulumi.gcp.compute.inputs.FirewallState;
import com.pulumi.gcp.compute.outputs.FirewallAllow;
import com.pulumi.gcp.compute.outputs.FirewallDeny;
import com.pulumi.gcp.compute.outputs.FirewallLogConfig;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Each network has its own firewall controlling access to and from the
 * instances.
 * 
 * All traffic to instances, even from other instances, is blocked by the
 * firewall unless firewall rules are created to allow it.
 * 
 * The default network has automatically created firewall rules that are
 * shown in default firewall rules. No manually created network has
 * automatically created firewall rules except for a default &#34;allow&#34; rule for
 * outgoing traffic and a default &#34;deny&#34; for incoming traffic. For all
 * networks except the default network, you must create any firewall rules
 * you need.
 * 
 * To get more information about Firewall, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/firewalls)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)
 * 
 * ## Example Usage
 * ### Firewall Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultFirewall = new Firewall(&#34;defaultFirewall&#34;, FirewallArgs.builder()        
 *             .network(defaultNetwork.name())
 *             .allows(            
 *                 FirewallAllowArgs.builder()
 *                     .protocol(&#34;icmp&#34;)
 *                     .build(),
 *                 FirewallAllowArgs.builder()
 *                     .protocol(&#34;tcp&#34;)
 *                     .ports(                    
 *                         &#34;80&#34;,
 *                         &#34;8080&#34;,
 *                         &#34;1000-2000&#34;)
 *                     .build())
 *             .sourceTags(&#34;web&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Firewall With Target Tags
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var rules = new Firewall(&#34;rules&#34;, FirewallArgs.builder()        
 *             .allows(FirewallAllowArgs.builder()
 *                 .ports(                
 *                     &#34;80&#34;,
 *                     &#34;8080&#34;,
 *                     &#34;1000-2000&#34;)
 *                 .protocol(&#34;tcp&#34;)
 *                 .build())
 *             .description(&#34;Creates firewall rule targeting tagged instances&#34;)
 *             .network(&#34;default&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .sourceTags(&#34;foo&#34;)
 *             .targetTags(&#34;web&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Firewall can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/firewall:Firewall default projects/{{project}}/global/firewalls/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/firewall:Firewall default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/firewall:Firewall default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/firewall:Firewall")
public class Firewall extends com.pulumi.resources.CustomResource {
    /**
     * The list of ALLOW rules specified by this firewall. Each rule
     * specifies a protocol and port-range tuple that describes a permitted
     * connection.
     * Structure is documented below.
     * 
     */
    @Export(name="allows", type=List.class, parameters={FirewallAllow.class})
    private Output</* @Nullable */ List<FirewallAllow>> allows;

    /**
     * @return The list of ALLOW rules specified by this firewall. Each rule
     * specifies a protocol and port-range tuple that describes a permitted
     * connection.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<FirewallAllow>>> allows() {
        return Codegen.optional(this.allows);
    }
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTimestamp", type=String.class, parameters={})
    private Output<String> creationTimestamp;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Output<String> creationTimestamp() {
        return this.creationTimestamp;
    }
    /**
     * The list of DENY rules specified by this firewall. Each rule specifies
     * a protocol and port-range tuple that describes a denied connection.
     * Structure is documented below.
     * 
     */
    @Export(name="denies", type=List.class, parameters={FirewallDeny.class})
    private Output</* @Nullable */ List<FirewallDeny>> denies;

    /**
     * @return The list of DENY rules specified by this firewall. Each rule specifies
     * a protocol and port-range tuple that describes a denied connection.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<FirewallDeny>>> denies() {
        return Codegen.optional(this.denies);
    }
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If destination ranges are specified, the firewall will apply only to
     * traffic that has destination IP address in these ranges. These ranges
     * must be expressed in CIDR format. Only IPv4 is supported.
     * 
     */
    @Export(name="destinationRanges", type=List.class, parameters={String.class})
    private Output<List<String>> destinationRanges;

    /**
     * @return If destination ranges are specified, the firewall will apply only to
     * traffic that has destination IP address in these ranges. These ranges
     * must be expressed in CIDR format. Only IPv4 is supported.
     * 
     */
    public Output<List<String>> destinationRanges() {
        return this.destinationRanges;
    }
    /**
     * Direction of traffic to which this firewall applies; default is
     * INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
     * destinationRanges; For EGRESS traffic, it is NOT supported to specify
     * `source_ranges` OR `source_tags`. For INGRESS traffic, one of `source_ranges`,
     * `source_tags` or `source_service_accounts` is required.
     * Possible values are `INGRESS` and `EGRESS`.
     * 
     */
    @Export(name="direction", type=String.class, parameters={})
    private Output<String> direction;

    /**
     * @return Direction of traffic to which this firewall applies; default is
     * INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
     * destinationRanges; For EGRESS traffic, it is NOT supported to specify
     * `source_ranges` OR `source_tags`. For INGRESS traffic, one of `source_ranges`,
     * `source_tags` or `source_service_accounts` is required.
     * Possible values are `INGRESS` and `EGRESS`.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * Denotes whether the firewall rule is disabled, i.e not applied to the
     * network it is associated with. When set to true, the firewall rule is
     * not enforced and the network behaves as if it did not exist. If this
     * is unspecified, the firewall rule will be enabled.
     * 
     */
    @Export(name="disabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Denotes whether the firewall rule is disabled, i.e not applied to the
     * network it is associated with. When set to true, the firewall rule is
     * not enforced and the network behaves as if it did not exist. If this
     * is unspecified, the firewall rule will be enabled.
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * This field denotes whether to enable logging for a particular firewall rule.
     * If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
     * 
     * @deprecated
     * Deprecated in favor of log_config
     * 
     */
    @Deprecated /* Deprecated in favor of log_config */
    @Export(name="enableLogging", type=Boolean.class, parameters={})
    private Output<Boolean> enableLogging;

    /**
     * @return This field denotes whether to enable logging for a particular firewall rule.
     * If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
     * 
     */
    public Output<Boolean> enableLogging() {
        return this.enableLogging;
    }
    /**
     * This field denotes the logging options for a particular firewall rule.
     * If defined, logging is enabled, and logs will be exported to Cloud Logging.
     * Structure is documented below.
     * 
     */
    @Export(name="logConfig", type=FirewallLogConfig.class, parameters={})
    private Output</* @Nullable */ FirewallLogConfig> logConfig;

    /**
     * @return This field denotes the logging options for a particular firewall rule.
     * If defined, logging is enabled, and logs will be exported to Cloud Logging.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FirewallLogConfig>> logConfig() {
        return Codegen.optional(this.logConfig);
    }
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name or self_link of the network to attach this firewall to.
     * 
     */
    @Export(name="network", type=String.class, parameters={})
    private Output<String> network;

    /**
     * @return The name or self_link of the network to attach this firewall to.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * Priority for this rule. This is an integer between 0 and 65535, both
     * inclusive. When not specified, the value assumed is 1000. Relative
     * priorities determine precedence of conflicting rules. Lower value of
     * priority implies higher precedence (eg, a rule with priority 0 has
     * higher precedence than a rule with priority 1). DENY rules take
     * precedence over ALLOW rules having equal priority.
     * 
     */
    @Export(name="priority", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return Priority for this rule. This is an integer between 0 and 65535, both
     * inclusive. When not specified, the value assumed is 1000. Relative
     * priorities determine precedence of conflicting rules. Lower value of
     * priority implies higher precedence (eg, a rule with priority 0 has
     * higher precedence than a rule with priority 1). DENY rules take
     * precedence over ALLOW rules having equal priority.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", type=String.class, parameters={})
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * If source ranges are specified, the firewall will apply only to
     * traffic that has source IP address in these ranges. These ranges must
     * be expressed in CIDR format. One or both of sourceRanges and
     * sourceTags may be set. If both properties are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP that belongs to a tag listed in the sourceTags property. The
     * connection does not need to match both properties for the firewall to
     * apply. Only IPv4 is supported. For INGRESS traffic, one of `source_ranges`,
     * `source_tags` or `source_service_accounts` is required.
     * 
     */
    @Export(name="sourceRanges", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> sourceRanges;

    /**
     * @return If source ranges are specified, the firewall will apply only to
     * traffic that has source IP address in these ranges. These ranges must
     * be expressed in CIDR format. One or both of sourceRanges and
     * sourceTags may be set. If both properties are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP that belongs to a tag listed in the sourceTags property. The
     * connection does not need to match both properties for the firewall to
     * apply. Only IPv4 is supported. For INGRESS traffic, one of `source_ranges`,
     * `source_tags` or `source_service_accounts` is required.
     * 
     */
    public Output<Optional<List<String>>> sourceRanges() {
        return Codegen.optional(this.sourceRanges);
    }
    /**
     * If source service accounts are specified, the firewall will apply only
     * to traffic originating from an instance with a service account in this
     * list. Source service accounts cannot be used to control traffic to an
     * instance&#39;s external IP address because service accounts are associated
     * with an instance, not an IP address. sourceRanges can be set at the
     * same time as sourceServiceAccounts. If both are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP belongs to an instance with service account listed in
     * sourceServiceAccount. The connection does not need to match both
     * properties for the firewall to apply. sourceServiceAccounts cannot be
     * used at the same time as sourceTags or targetTags. For INGRESS traffic,
     * one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
     * 
     */
    @Export(name="sourceServiceAccounts", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> sourceServiceAccounts;

    /**
     * @return If source service accounts are specified, the firewall will apply only
     * to traffic originating from an instance with a service account in this
     * list. Source service accounts cannot be used to control traffic to an
     * instance&#39;s external IP address because service accounts are associated
     * with an instance, not an IP address. sourceRanges can be set at the
     * same time as sourceServiceAccounts. If both are set, the firewall will
     * apply to traffic that has source IP address within sourceRanges OR the
     * source IP belongs to an instance with service account listed in
     * sourceServiceAccount. The connection does not need to match both
     * properties for the firewall to apply. sourceServiceAccounts cannot be
     * used at the same time as sourceTags or targetTags. For INGRESS traffic,
     * one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
     * 
     */
    public Output<Optional<List<String>>> sourceServiceAccounts() {
        return Codegen.optional(this.sourceServiceAccounts);
    }
    /**
     * If source tags are specified, the firewall will apply only to traffic
     * with source IP that belongs to a tag listed in source tags. Source
     * tags cannot be used to control traffic to an instance&#39;s external IP
     * address. Because tags are associated with an instance, not an IP
     * address. One or both of sourceRanges and sourceTags may be set. If
     * both properties are set, the firewall will apply to traffic that has
     * source IP address within sourceRanges OR the source IP that belongs to
     * a tag listed in the sourceTags property. The connection does not need
     * to match both properties for the firewall to apply. For INGRESS traffic,
     * one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
     * 
     */
    @Export(name="sourceTags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> sourceTags;

    /**
     * @return If source tags are specified, the firewall will apply only to traffic
     * with source IP that belongs to a tag listed in source tags. Source
     * tags cannot be used to control traffic to an instance&#39;s external IP
     * address. Because tags are associated with an instance, not an IP
     * address. One or both of sourceRanges and sourceTags may be set. If
     * both properties are set, the firewall will apply to traffic that has
     * source IP address within sourceRanges OR the source IP that belongs to
     * a tag listed in the sourceTags property. The connection does not need
     * to match both properties for the firewall to apply. For INGRESS traffic,
     * one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
     * 
     */
    public Output<Optional<List<String>>> sourceTags() {
        return Codegen.optional(this.sourceTags);
    }
    /**
     * A list of service accounts indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * targetServiceAccounts cannot be used at the same time as targetTags or
     * sourceTags. If neither targetServiceAccounts nor targetTags are
     * specified, the firewall rule applies to all instances on the specified
     * network.
     * 
     */
    @Export(name="targetServiceAccounts", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> targetServiceAccounts;

    /**
     * @return A list of service accounts indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * targetServiceAccounts cannot be used at the same time as targetTags or
     * sourceTags. If neither targetServiceAccounts nor targetTags are
     * specified, the firewall rule applies to all instances on the specified
     * network.
     * 
     */
    public Output<Optional<List<String>>> targetServiceAccounts() {
        return Codegen.optional(this.targetServiceAccounts);
    }
    /**
     * A list of instance tags indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * If no targetTags are specified, the firewall rule applies to all
     * instances on the specified network.
     * 
     */
    @Export(name="targetTags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> targetTags;

    /**
     * @return A list of instance tags indicating sets of instances located in the
     * network that may make network connections as specified in allowed[].
     * If no targetTags are specified, the firewall rule applies to all
     * instances on the specified network.
     * 
     */
    public Output<Optional<List<String>>> targetTags() {
        return Codegen.optional(this.targetTags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Firewall(String name) {
        this(name, FirewallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Firewall(String name, FirewallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Firewall(String name, FirewallArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewall:Firewall", name, args == null ? FirewallArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Firewall(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/firewall:Firewall", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Firewall get(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Firewall(name, id, state, options);
    }
}
