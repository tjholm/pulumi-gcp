// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.activedirectory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.activedirectory.DomainArgs;
import com.pulumi.gcp.activedirectory.inputs.DomainState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a Microsoft AD domain
 * 
 * To get more information about Domain, see:
 * 
 * * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains)
 * * How-to Guides
 *     * [Managed Microsoft Active Directory Quickstart](https://cloud.google.com/managed-microsoft-ad/docs/quickstarts)
 * 
 * ## Example Usage
 * ### Active Directory Domain Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.activedirectory.Domain;
 * import com.pulumi.gcp.activedirectory.DomainArgs;
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
 *         var ad_domain = new Domain(&#34;ad-domain&#34;, DomainArgs.builder()        
 *             .domainName(&#34;tfgen.org.com&#34;)
 *             .locations(&#34;us-central1&#34;)
 *             .reservedIpRange(&#34;192.168.255.0/24&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Domain can be imported using any of these accepted formats:
 * 
 * ```sh
 *  $ pulumi import gcp:activedirectory/domain:Domain default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:activedirectory/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The name of delegated administrator account used to perform Active Directory operations.
     * If not specified, setupadmin will be used.
     * 
     */
    @Export(name="admin", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> admin;

    /**
     * @return The name of delegated administrator account used to perform Active Directory operations.
     * If not specified, setupadmin will be used.
     * 
     */
    public Output<Optional<String>> admin() {
        return Codegen.optional(this.admin);
    }
    /**
     * The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
     * If CIDR subnets overlap between networks, domain creation will fail.
     * 
     */
    @Export(name="authorizedNetworks", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authorizedNetworks;

    /**
     * @return The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
     * If CIDR subnets overlap between networks, domain creation will fail.
     * 
     */
    public Output<Optional<List<String>>> authorizedNetworks() {
        return Codegen.optional(this.authorizedNetworks);
    }
    /**
     * The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     * 
     * ***
     * 
     */
    @Export(name="domainName", refs={String.class}, tree="[0]")
    private Output<String> domainName;

    /**
     * @return The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
     * https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
     * 
     * ***
     * 
     */
    public Output<String> domainName() {
        return this.domainName;
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
     * The fully-qualified domain name of the exposed domain used by clients to connect to the service.
     * Similar to what would be chosen for an Active Directory set up on an internal network.
     * 
     */
    @Export(name="fqdn", refs={String.class}, tree="[0]")
    private Output<String> fqdn;

    /**
     * @return The fully-qualified domain name of the exposed domain used by clients to connect to the service.
     * Similar to what would be chosen for an Active Directory set up on an internal network.
     * 
     */
    public Output<String> fqdn() {
        return this.fqdn;
    }
    /**
     * Resource labels that can contain user-provided metadata
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Resource labels that can contain user-provided metadata
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
     * e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     * 
     */
    @Export(name="locations", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> locations;

    /**
     * @return Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
     * e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
     * 
     */
    public Output<List<String>> locations() {
        return this.locations;
    }
    /**
     * The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
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
     * The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
     * Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
     * 
     */
    @Export(name="reservedIpRange", refs={String.class}, tree="[0]")
    private Output<String> reservedIpRange;

    /**
     * @return The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
     * Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
     * 
     */
    public Output<String> reservedIpRange() {
        return this.reservedIpRange;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:activedirectory/domain:Domain", name, args == null ? DomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Domain(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:activedirectory/domain:Domain", name, state, makeResourceOptions(options, id));
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
    public static Domain get(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
