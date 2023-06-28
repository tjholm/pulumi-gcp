// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkehub.FeatureArgs;
import com.pulumi.gcp.gkehub.outputs.FeatureResourceState;
import com.pulumi.gcp.gkehub.outputs.FeatureSpec;
import com.pulumi.gcp.gkehub.outputs.FeatureState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Feature represents the settings and status of any Hub Feature.
 * 
 * To get more information about Feature, see:
 * 
 * * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.features)
 * * How-to Guides
 *     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
 * 
 * ## Example Usage
 * ### Gkehub Feature Multi Cluster Ingress
 * ```java
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
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecMulticlusteringressArgs;
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
 *         var cluster = new Cluster(&#34;cluster&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1-a&#34;)
 *             .initialNodeCount(1)
 *             .build());
 * 
 *         var membership = new Membership(&#34;membership&#34;, MembershipArgs.builder()        
 *             .membershipId(&#34;my-membership&#34;)
 *             .endpoint(MembershipEndpointArgs.builder()
 *                 .gkeCluster(MembershipEndpointGkeClusterArgs.builder()
 *                     .resourceLink(cluster.id().applyValue(id -&gt; String.format(&#34;//container.googleapis.com/%s&#34;, id)))
 *                     .build())
 *                 .build())
 *             .description(&#34;Membership&#34;)
 *             .build());
 * 
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .spec(FeatureSpecArgs.builder()
 *                 .multiclusteringress(FeatureSpecMulticlusteringressArgs.builder()
 *                     .configMembership(membership.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkehub Feature Multi Cluster Service Discovery
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
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
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .labels(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .location(&#34;global&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkehub Feature Anthos Service Mesh
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
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
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Enable Fleet Observability For Default Logs With Copy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .spec(FeatureSpecArgs.builder()
 *                 .fleetobservability(FeatureSpecFleetobservabilityArgs.builder()
 *                     .loggingConfig(FeatureSpecFleetobservabilityLoggingConfigArgs.builder()
 *                         .defaultConfig(FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs.builder()
 *                             .mode(&#34;COPY&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Enable Fleet Observability For Scope Logs With Move
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .spec(FeatureSpecArgs.builder()
 *                 .fleetobservability(FeatureSpecFleetobservabilityArgs.builder()
 *                     .loggingConfig(FeatureSpecFleetobservabilityLoggingConfigArgs.builder()
 *                         .fleetScopeLogsConfig(FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs.builder()
 *                             .mode(&#34;MOVE&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Enable Fleet Observability For Both Default And Scope Logs
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkehub.Feature;
 * import com.pulumi.gcp.gkehub.FeatureArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs;
 * import com.pulumi.gcp.gkehub.inputs.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var feature = new Feature(&#34;feature&#34;, FeatureArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .spec(FeatureSpecArgs.builder()
 *                 .fleetobservability(FeatureSpecFleetobservabilityArgs.builder()
 *                     .loggingConfig(FeatureSpecFleetobservabilityLoggingConfigArgs.builder()
 *                         .defaultConfig(FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs.builder()
 *                             .mode(&#34;COPY&#34;)
 *                             .build())
 *                         .fleetScopeLogsConfig(FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs.builder()
 *                             .mode(&#34;MOVE&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Feature can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/feature:Feature default projects/{{project}}/locations/{{location}}/features/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/feature:Feature default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/feature:Feature default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkehub/feature:Feature")
public class Feature extends com.pulumi.resources.CustomResource {
    /**
     * Output only. When the Feature resource was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Output only. When the Feature resource was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Output only. When the Feature resource was deleted.
     * 
     */
    @Export(name="deleteTime", type=String.class, parameters={})
    private Output<String> deleteTime;

    /**
     * @return Output only. When the Feature resource was deleted.
     * 
     */
    public Output<String> deleteTime() {
        return this.deleteTime;
    }
    /**
     * GCP labels for this Feature.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return GCP labels for this Feature.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location for the resource
     * 
     * ***
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The full, unique name of this Feature resource
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The full, unique name of this Feature resource
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
     * State of the Feature resource itself.
     * Structure is documented below.
     * 
     */
    @Export(name="resourceStates", type=List.class, parameters={FeatureResourceState.class})
    private Output<List<FeatureResourceState>> resourceStates;

    /**
     * @return State of the Feature resource itself.
     * Structure is documented below.
     * 
     */
    public Output<List<FeatureResourceState>> resourceStates() {
        return this.resourceStates;
    }
    /**
     * Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
     * Structure is documented below.
     * 
     */
    @Export(name="spec", type=FeatureSpec.class, parameters={})
    private Output</* @Nullable */ FeatureSpec> spec;

    /**
     * @return Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FeatureSpec>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * (Output)
     * Output only. The &#34;running state&#34; of the Feature in this Hub.
     * Structure is documented below.
     * 
     */
    @Export(name="states", type=List.class, parameters={FeatureState.class})
    private Output<List<FeatureState>> states;

    /**
     * @return (Output)
     * Output only. The &#34;running state&#34; of the Feature in this Hub.
     * Structure is documented below.
     * 
     */
    public Output<List<FeatureState>> states() {
        return this.states;
    }
    /**
     * (Output)
     * The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return (Output)
     * The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Feature(String name) {
        this(name, FeatureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Feature(String name, FeatureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Feature(String name, FeatureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/feature:Feature", name, args == null ? FeatureArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Feature(String name, Output<String> id, @Nullable com.pulumi.gcp.gkehub.inputs.FeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/feature:Feature", name, state, makeResourceOptions(options, id));
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
    public static Feature get(String name, Output<String> id, @Nullable com.pulumi.gcp.gkehub.inputs.FeatureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Feature(name, id, state, options);
    }
}
