// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.integrationconnectors;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.integrationconnectors.ManagedZoneArgs;
import com.pulumi.gcp.integrationconnectors.inputs.ManagedZoneState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An Integration connectors Managed Zone.
 * 
 * To get more information about ManagedZone, see:
 * 
 * * [API documentation](https://cloud.google.com/integration-connectors/docs/reference/rest/v1/projects.locations.global.managedZones)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/integration-connectors/docs)
 * 
 * ## Example Usage
 * 
 * ### Integration Connectors Managed Zone
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.projects.IAMMember;
 * import com.pulumi.gcp.projects.IAMMemberArgs;
 * import com.pulumi.gcp.projects.Service;
 * import com.pulumi.gcp.projects.ServiceArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.dns.ManagedZone;
 * import com.pulumi.gcp.dns.ManagedZoneArgs;
 * import com.pulumi.gcp.dns.inputs.ManagedZonePrivateVisibilityConfigArgs;
 * import com.pulumi.gcp.integrationconnectors.ManagedZone;
 * import com.pulumi.gcp.integrationconnectors.ManagedZoneArgs;
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
 *         var targetProject = new Project("targetProject", ProjectArgs.builder()
 *             .projectId("tf-test_75092")
 *             .name("tf-test_2605")
 *             .orgId("123456789")
 *             .billingAccount("000000-0000000-0000000-000000")
 *             .build());
 * 
 *         final var testProject = OrganizationsFunctions.getProject();
 * 
 *         var dnsPeerBinding = new IAMMember("dnsPeerBinding", IAMMemberArgs.builder()
 *             .project(targetProject.projectId())
 *             .role("roles/dns.peer")
 *             .member(String.format("serviceAccount:service-%s{@literal @}gcp-sa-connectors.iam.gserviceaccount.com", testProject.applyValue(getProjectResult -> getProjectResult.number())))
 *             .build());
 * 
 *         var dns = new Service("dns", ServiceArgs.builder()
 *             .project(targetProject.projectId())
 *             .service("dns.googleapis.com")
 *             .build());
 * 
 *         var compute = new Service("compute", ServiceArgs.builder()
 *             .project(targetProject.projectId())
 *             .service("compute.googleapis.com")
 *             .build());
 * 
 *         var network = new Network("network", NetworkArgs.builder()
 *             .project(targetProject.projectId())
 *             .name("test")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var zone = new ManagedZone("zone", ManagedZoneArgs.builder()
 *             .name("tf-test-dns_34535")
 *             .dnsName("private_22375.example.com.")
 *             .visibility("private")
 *             .privateVisibilityConfig(ManagedZonePrivateVisibilityConfigArgs.builder()
 *                 .networks(ManagedZonePrivateVisibilityConfigNetworkArgs.builder()
 *                     .networkUrl(network.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var testmanagedzone = new ManagedZone("testmanagedzone", ManagedZoneArgs.builder()
 *             .name("test")
 *             .description("tf created description")
 *             .labels(Map.of("intent", "example"))
 *             .targetProject(targetProject.projectId())
 *             .targetVpc("test")
 *             .dns(zone.dnsName())
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
 * ManagedZone can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/global/managedZones/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, ManagedZone can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:integrationconnectors/managedZone:ManagedZone default projects/{{project}}/locations/global/managedZones/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:integrationconnectors/managedZone:ManagedZone default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:integrationconnectors/managedZone:ManagedZone default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:integrationconnectors/managedZone:ManagedZone")
public class ManagedZone extends com.pulumi.resources.CustomResource {
    /**
     * Time the Namespace was created in UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time the Namespace was created in UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * DNS Name of the resource.
     * 
     */
    @Export(name="dns", refs={String.class}, tree="[0]")
    private Output<String> dns;

    /**
     * @return DNS Name of the resource.
     * 
     */
    public Output<String> dns() {
        return this.dns;
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
     * Resource labels to represent user provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Resource labels to represent user provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Name of Managed Zone needs to be created.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of Managed Zone needs to be created.
     * 
     * ***
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
     * The name of the Target Project.
     * 
     */
    @Export(name="targetProject", refs={String.class}, tree="[0]")
    private Output<String> targetProject;

    /**
     * @return The name of the Target Project.
     * 
     */
    public Output<String> targetProject() {
        return this.targetProject;
    }
    /**
     * The name of the Target Project VPC Network.
     * 
     */
    @Export(name="targetVpc", refs={String.class}, tree="[0]")
    private Output<String> targetVpc;

    /**
     * @return The name of the Target Project VPC Network.
     * 
     */
    public Output<String> targetVpc() {
        return this.targetVpc;
    }
    /**
     * Time the Namespace was updated in UTC.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time the Namespace was updated in UTC.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ManagedZone(String name) {
        this(name, ManagedZoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ManagedZone(String name, ManagedZoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ManagedZone(String name, ManagedZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:integrationconnectors/managedZone:ManagedZone", name, args == null ? ManagedZoneArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ManagedZone(String name, Output<String> id, @Nullable ManagedZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:integrationconnectors/managedZone:ManagedZone", name, state, makeResourceOptions(options, id));
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
    public static ManagedZone get(String name, Output<String> id, @Nullable ManagedZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ManagedZone(name, id, state, options);
    }
}
