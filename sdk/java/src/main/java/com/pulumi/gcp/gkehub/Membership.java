// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkehub.MembershipArgs;
import com.pulumi.gcp.gkehub.inputs.MembershipState;
import com.pulumi.gcp.gkehub.outputs.MembershipAuthority;
import com.pulumi.gcp.gkehub.outputs.MembershipEndpoint;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Membership contains information about a member cluster.
 * 
 * To get more information about Membership, see:
 * 
 * * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
 * * How-to Guides
 *     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
 * 
 * ## Example Usage
 * 
 * ### Gkehub Membership Regional
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.gkehub.Membership;
 * import com.pulumi.gcp.gkehub.MembershipArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointGkeClusterArgs;
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
 *         var primary = new Cluster("primary", ClusterArgs.builder()
 *             .name("basic-cluster")
 *             .location("us-central1-a")
 *             .initialNodeCount(1)
 *             .deletionProtection(false)
 *             .network("default")
 *             .subnetwork("default")
 *             .build());
 * 
 *         var membership = new Membership("membership", MembershipArgs.builder()
 *             .membershipId("basic")
 *             .location("us-west1")
 *             .endpoint(MembershipEndpointArgs.builder()
 *                 .gkeCluster(MembershipEndpointGkeClusterArgs.builder()
 *                     .resourceLink(primary.id().applyValue(id -> String.format("//container.googleapis.com/%s", id)))
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Gkehub Membership Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.gkehub.Membership;
 * import com.pulumi.gcp.gkehub.MembershipArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointGkeClusterArgs;
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
 *         var primary = new Cluster("primary", ClusterArgs.builder()
 *             .name("basic-cluster")
 *             .location("us-central1-a")
 *             .initialNodeCount(1)
 *             .deletionProtection("true")
 *             .network("default")
 *             .subnetwork("default")
 *             .build());
 * 
 *         var membership = new Membership("membership", MembershipArgs.builder()
 *             .membershipId("basic")
 *             .endpoint(MembershipEndpointArgs.builder()
 *                 .gkeCluster(MembershipEndpointGkeClusterArgs.builder()
 *                     .resourceLink(primary.id().applyValue(id -> String.format("//container.googleapis.com/%s", id)))
 *                     .build())
 *                 .build())
 *             .labels(Map.of("env", "test"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Gkehub Membership Issuer
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterWorkloadIdentityConfigArgs;
 * import com.pulumi.gcp.gkehub.Membership;
 * import com.pulumi.gcp.gkehub.MembershipArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipEndpointGkeClusterArgs;
 * import com.pulumi.gcp.gkehub.inputs.MembershipAuthorityArgs;
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
 *         var primary = new Cluster("primary", ClusterArgs.builder()
 *             .name("basic-cluster")
 *             .location("us-central1-a")
 *             .initialNodeCount(1)
 *             .workloadIdentityConfig(ClusterWorkloadIdentityConfigArgs.builder()
 *                 .workloadPool("my-project-name.svc.id.goog")
 *                 .build())
 *             .deletionProtection("true")
 *             .network("default")
 *             .subnetwork("default")
 *             .build());
 * 
 *         var membership = new Membership("membership", MembershipArgs.builder()
 *             .membershipId("basic")
 *             .endpoint(MembershipEndpointArgs.builder()
 *                 .gkeCluster(MembershipEndpointGkeClusterArgs.builder()
 *                     .resourceLink(primary.id())
 *                     .build())
 *                 .build())
 *             .authority(MembershipAuthorityArgs.builder()
 *                 .issuer(primary.id().applyValue(id -> String.format("https://container.googleapis.com/v1/%s", id)))
 *                 .build())
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
 * Membership can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}`
 * 
 * * `{{project}}/{{location}}/{{membership_id}}`
 * 
 * * `{{location}}/{{membership_id}}`
 * 
 * When using the `pulumi import` command, Membership can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membership:Membership default projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membership:Membership default {{project}}/{{location}}/{{membership_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:gkehub/membership:Membership default {{location}}/{{membership_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkehub/membership:Membership")
public class Membership extends com.pulumi.resources.CustomResource {
    /**
     * Authority encodes how Google will recognize identities from this Membership.
     * See the workload identity documentation for more details:
     * https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     * Structure is documented below.
     * 
     */
    @Export(name="authority", refs={MembershipAuthority.class}, tree="[0]")
    private Output</* @Nullable */ MembershipAuthority> authority;

    /**
     * @return Authority encodes how Google will recognize identities from this Membership.
     * See the workload identity documentation for more details:
     * https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     * Structure is documented below.
     * 
     */
    public Output<Optional<MembershipAuthority>> authority() {
        return Codegen.optional(this.authority);
    }
    /**
     * The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
     * 
     * &gt; **Warning:** `description` is deprecated and will be removed in a future major release.
     * 
     * @deprecated
     * `description` is deprecated and will be removed in a future major release.
     * 
     */
    @Deprecated /* `description` is deprecated and will be removed in a future major release. */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
     * 
     * &gt; **Warning:** `description` is deprecated and will be removed in a future major release.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
     * Structure is documented below.
     * 
     */
    @Export(name="endpoint", refs={MembershipEndpoint.class}, tree="[0]")
    private Output</* @Nullable */ MembershipEndpoint> endpoint;

    /**
     * @return If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
     * Structure is documented below.
     * 
     */
    public Output<Optional<MembershipEndpoint>> endpoint() {
        return Codegen.optional(this.endpoint);
    }
    /**
     * Labels to apply to this membership.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels to apply to this membership.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Location of the membership.
     * The default value is `global`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> location;

    /**
     * @return Location of the membership.
     * The default value is `global`.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * The client-provided identifier of the membership.
     * 
     * ***
     * 
     */
    @Export(name="membershipId", refs={String.class}, tree="[0]")
    private Output<String> membershipId;

    /**
     * @return The client-provided identifier of the membership.
     * 
     * ***
     * 
     */
    public Output<String> membershipId() {
        return this.membershipId;
    }
    /**
     * The unique identifier of the membership.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique identifier of the membership.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Membership(String name) {
        this(name, MembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Membership(String name, MembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Membership(String name, MembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/membership:Membership", name, args == null ? MembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Membership(String name, Output<String> id, @Nullable MembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/membership:Membership", name, state, makeResourceOptions(options, id));
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
    public static Membership get(String name, Output<String> id, @Nullable MembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Membership(name, id, state, options);
    }
}
