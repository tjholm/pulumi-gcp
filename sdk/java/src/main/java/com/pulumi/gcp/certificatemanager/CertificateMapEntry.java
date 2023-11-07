// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificatemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.certificatemanager.CertificateMapEntryArgs;
import com.pulumi.gcp.certificatemanager.inputs.CertificateMapEntryState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * CertificateMapEntry is a list of certificate configurations,
 * that have been issued for a particular hostname
 * 
 * ## Example Usage
 * ### Certificate Manager Certificate Map Entry Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.certificatemanager.CertificateMap;
 * import com.pulumi.gcp.certificatemanager.CertificateMapArgs;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorization;
 * import com.pulumi.gcp.certificatemanager.DnsAuthorizationArgs;
 * import com.pulumi.gcp.certificatemanager.Certificate;
 * import com.pulumi.gcp.certificatemanager.CertificateArgs;
 * import com.pulumi.gcp.certificatemanager.inputs.CertificateManagedArgs;
 * import com.pulumi.gcp.certificatemanager.CertificateMapEntry;
 * import com.pulumi.gcp.certificatemanager.CertificateMapEntryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var certificateMap = new CertificateMap(&#34;certificateMap&#34;, CertificateMapArgs.builder()        
 *             .description(&#34;My acceptance test certificate map&#34;)
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;terraform&#34;, true),
 *                 Map.entry(&#34;acc-test&#34;, true)
 *             ))
 *             .build());
 * 
 *         var instance = new DnsAuthorization(&#34;instance&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var instance2 = new DnsAuthorization(&#34;instance2&#34;, DnsAuthorizationArgs.builder()        
 *             .description(&#34;The default dnss&#34;)
 *             .domain(&#34;subdomain2.hashicorptest.com&#34;)
 *             .build());
 * 
 *         var certificate = new Certificate(&#34;certificate&#34;, CertificateArgs.builder()        
 *             .description(&#34;The default cert&#34;)
 *             .scope(&#34;DEFAULT&#34;)
 *             .managed(CertificateManagedArgs.builder()
 *                 .domains(                
 *                     instance.domain(),
 *                     instance2.domain())
 *                 .dnsAuthorizations(                
 *                     instance.id(),
 *                     instance2.id())
 *                 .build())
 *             .build());
 * 
 *         var default_ = new CertificateMapEntry(&#34;default&#34;, CertificateMapEntryArgs.builder()        
 *             .description(&#34;My acceptance test certificate map entry&#34;)
 *             .map(certificateMap.name())
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;terraform&#34;, true),
 *                 Map.entry(&#34;acc-test&#34;, true)
 *             ))
 *             .certificates(certificate.id())
 *             .matcher(&#34;PRIMARY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CertificateMapEntry can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default {{project}}/{{map}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:certificatemanager/certificateMapEntry:CertificateMapEntry default {{map}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:certificatemanager/certificateMapEntry:CertificateMapEntry")
public class CertificateMapEntry extends com.pulumi.resources.CustomResource {
    /**
     * A set of Certificates defines for the given hostname.
     * There can be defined up to fifteen certificates in each Certificate Map Entry.
     * Each certificate must match pattern projects/*{@literal /}locations/*{@literal /}certificates/*.
     * 
     */
    @Export(name="certificates", refs={List.class,String.class}, tree="[0,1]")
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
     * Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC &#34;Zulu&#34; format,
     * with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC &#34;Zulu&#34; format,
     * with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A human-readable description of the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
     * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
     * selecting a proper certificate.
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostname;

    /**
     * @return A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
     * for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
     * selecting a proper certificate.
     * 
     */
    public Output<Optional<String>> hostname() {
        return Codegen.optional(this.hostname);
    }
    /**
     * Set of labels associated with a Certificate Map Entry.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Set of labels associated with a Certificate Map Entry.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * A map entry that is inputted into the cetrificate map
     * 
     * ***
     * 
     */
    @Export(name="map", refs={String.class}, tree="[0]")
    private Output<String> map;

    /**
     * @return A map entry that is inputted into the cetrificate map
     * 
     * ***
     * 
     */
    public Output<String> map() {
        return this.map;
    }
    /**
     * A predefined matcher for particular cases, other than SNI selection
     * 
     */
    @Export(name="matcher", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> matcher;

    /**
     * @return A predefined matcher for particular cases, other than SNI selection
     * 
     */
    public Output<Optional<String>> matcher() {
        return Codegen.optional(this.matcher);
    }
    /**
     * A user-defined name of the Certificate Map Entry. Certificate Map Entry
     * names must be unique globally and match pattern
     * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A user-defined name of the Certificate Map Entry. Certificate Map Entry
     * names must be unique globally and match pattern
     * &#39;projects/*{@literal /}locations/*{@literal /}certificateMaps/*{@literal /}certificateMapEntries/*&#39;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
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
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * A serving state of this Certificate Map Entry.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return A serving state of this Certificate Map Entry.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC &#34;Zulu&#34; format,
     * with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC &#34;Zulu&#34; format,
     * with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertificateMapEntry(String name) {
        this(name, CertificateMapEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertificateMapEntry(String name, CertificateMapEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertificateMapEntry(String name, CertificateMapEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificateMapEntry:CertificateMapEntry", name, args == null ? CertificateMapEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertificateMapEntry(String name, Output<String> id, @Nullable CertificateMapEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:certificatemanager/certificateMapEntry:CertificateMapEntry", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
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
    public static CertificateMapEntry get(String name, Output<String> id, @Nullable CertificateMapEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertificateMapEntry(name, id, state, options);
    }
}
