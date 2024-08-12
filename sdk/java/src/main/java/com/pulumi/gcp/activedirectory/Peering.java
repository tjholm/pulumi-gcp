// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.activedirectory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.activedirectory.PeeringArgs;
import com.pulumi.gcp.activedirectory.inputs.PeeringState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Active Directory Peering Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.activedirectory.Domain;
 * import com.pulumi.gcp.activedirectory.DomainArgs;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.activedirectory.Peering;
 * import com.pulumi.gcp.activedirectory.PeeringArgs;
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
 *         var source_network = new Network("source-network", NetworkArgs.builder()
 *             .name("ad-network")
 *             .build());
 * 
 *         var ad_domain = new Domain("ad-domain", DomainArgs.builder()
 *             .domainName("ad.test.hashicorptest.com")
 *             .locations("us-central1")
 *             .reservedIpRange("192.168.255.0/24")
 *             .authorizedNetworks(source_network.id())
 *             .build());
 * 
 *         var peered_project = new Project("peered-project", ProjectArgs.builder()
 *             .name("my-peered-project")
 *             .projectId("my-peered-project")
 *             .orgId("123456789")
 *             .billingAccount("000000-0000000-0000000-000000")
 *             .build());
 * 
 *         var compute = new Service("compute", ServiceArgs.builder()
 *             .project(peered_project.projectId())
 *             .service("compute.googleapis.com")
 *             .build());
 * 
 *         var peered_network = new Network("peered-network", NetworkArgs.builder()
 *             .project(compute.project())
 *             .name("ad-peered-network")
 *             .build());
 * 
 *         var ad_domain_peering = new Peering("ad-domain-peering", PeeringArgs.builder()
 *             .domainResource(ad_domain.name())
 *             .peeringId("ad-domain-peering")
 *             .authorizedNetwork(peered_network.id())
 *             .labels(Map.of("foo", "bar"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:activedirectory/peering:Peering")
public class Peering extends com.pulumi.resources.CustomResource {
    /**
     * The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
     * 
     */
    @Export(name="authorizedNetwork", refs={String.class}, tree="[0]")
    private Output<String> authorizedNetwork;

    /**
     * @return The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
     * 
     */
    public Output<String> authorizedNetwork() {
        return this.authorizedNetwork;
    }
    /**
     * Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
     * 
     */
    @Export(name="domainResource", refs={String.class}, tree="[0]")
    private Output<String> domainResource;

    /**
     * @return Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
     * 
     */
    public Output<String> domainResource() {
        return this.domainResource;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
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
     * Unique name of the peering in this scope including projects and location using the form: projects/{projectId}/locations/global/peerings/{peeringId}.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name of the peering in this scope including projects and location using the form: projects/{projectId}/locations/global/peerings/{peeringId}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * ***
     * 
     */
    @Export(name="peeringId", refs={String.class}, tree="[0]")
    private Output<String> peeringId;

    /**
     * @return - - -
     * 
     */
    public Output<String> peeringId() {
        return this.peeringId;
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
     * The current state of this Peering.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return The current state of this Peering.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * Additional information about the current status of this peering, if available.
     * 
     */
    @Export(name="statusMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> statusMessage;

    /**
     * @return Additional information about the current status of this peering, if available.
     * 
     */
    public Output<Optional<String>> statusMessage() {
        return Codegen.optional(this.statusMessage);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Peering(java.lang.String name) {
        this(name, PeeringArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Peering(java.lang.String name, PeeringArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Peering(java.lang.String name, PeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:activedirectory/peering:Peering", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Peering(java.lang.String name, Output<java.lang.String> id, @Nullable PeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:activedirectory/peering:Peering", name, state, makeResourceOptions(options, id), false);
    }

    private static PeeringArgs makeArgs(PeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PeeringArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Peering get(java.lang.String name, Output<java.lang.String> id, @Nullable PeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Peering(name, id, state, options);
    }
}
