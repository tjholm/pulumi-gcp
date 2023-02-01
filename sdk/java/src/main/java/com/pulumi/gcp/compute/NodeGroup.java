// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.NodeGroupArgs;
import com.pulumi.gcp.compute.inputs.NodeGroupState;
import com.pulumi.gcp.compute.outputs.NodeGroupAutoscalingPolicy;
import com.pulumi.gcp.compute.outputs.NodeGroupMaintenanceWindow;
import com.pulumi.gcp.compute.outputs.NodeGroupShareSettings;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Node Group Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.NodeTemplate;
 * import com.pulumi.gcp.compute.NodeTemplateArgs;
 * import com.pulumi.gcp.compute.NodeGroup;
 * import com.pulumi.gcp.compute.NodeGroupArgs;
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
 *         var soletenant_tmpl = new NodeTemplate(&#34;soletenant-tmpl&#34;, NodeTemplateArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .nodeType(&#34;n1-node-96-624&#34;)
 *             .build());
 * 
 *         var nodes = new NodeGroup(&#34;nodes&#34;, NodeGroupArgs.builder()        
 *             .zone(&#34;us-central1-f&#34;)
 *             .description(&#34;example google_compute_node_group for Terraform Google Provider&#34;)
 *             .size(1)
 *             .nodeTemplate(soletenant_tmpl.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Node Group Autoscaling Policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.NodeTemplate;
 * import com.pulumi.gcp.compute.NodeTemplateArgs;
 * import com.pulumi.gcp.compute.NodeGroup;
 * import com.pulumi.gcp.compute.NodeGroupArgs;
 * import com.pulumi.gcp.compute.inputs.NodeGroupMaintenanceWindowArgs;
 * import com.pulumi.gcp.compute.inputs.NodeGroupAutoscalingPolicyArgs;
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
 *         var soletenant_tmpl = new NodeTemplate(&#34;soletenant-tmpl&#34;, NodeTemplateArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .nodeType(&#34;n1-node-96-624&#34;)
 *             .build());
 * 
 *         var nodes = new NodeGroup(&#34;nodes&#34;, NodeGroupArgs.builder()        
 *             .zone(&#34;us-central1-f&#34;)
 *             .description(&#34;example google_compute_node_group for Terraform Google Provider&#34;)
 *             .maintenancePolicy(&#34;RESTART_IN_PLACE&#34;)
 *             .maintenanceWindow(NodeGroupMaintenanceWindowArgs.builder()
 *                 .startTime(&#34;08:00&#34;)
 *                 .build())
 *             .initialSize(1)
 *             .nodeTemplate(soletenant_tmpl.id())
 *             .autoscalingPolicy(NodeGroupAutoscalingPolicyArgs.builder()
 *                 .mode(&#34;ONLY_SCALE_OUT&#34;)
 *                 .minNodes(1)
 *                 .maxNodes(10)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Node Group Share Settings
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.Project;
 * import com.pulumi.gcp.organizations.ProjectArgs;
 * import com.pulumi.gcp.compute.NodeTemplate;
 * import com.pulumi.gcp.compute.NodeTemplateArgs;
 * import com.pulumi.gcp.compute.NodeGroup;
 * import com.pulumi.gcp.compute.NodeGroupArgs;
 * import com.pulumi.gcp.compute.inputs.NodeGroupShareSettingsArgs;
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
 *         var guestProject = new Project(&#34;guestProject&#34;, ProjectArgs.builder()        
 *             .projectId(&#34;project-id&#34;)
 *             .orgId(&#34;123456789&#34;)
 *             .build());
 * 
 *         var soletenant_tmpl = new NodeTemplate(&#34;soletenant-tmpl&#34;, NodeTemplateArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .nodeType(&#34;n1-node-96-624&#34;)
 *             .build());
 * 
 *         var nodes = new NodeGroup(&#34;nodes&#34;, NodeGroupArgs.builder()        
 *             .zone(&#34;us-central1-f&#34;)
 *             .description(&#34;example google_compute_node_group for Terraform Google Provider&#34;)
 *             .size(1)
 *             .nodeTemplate(soletenant_tmpl.id())
 *             .shareSettings(NodeGroupShareSettingsArgs.builder()
 *                 .shareType(&#34;SPECIFIC_PROJECTS&#34;)
 *                 .projectMaps(NodeGroupShareSettingsProjectMapArgs.builder()
 *                     .id(guestProject.projectId())
 *                     .projectId(guestProject.projectId())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * NodeGroup can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{project}}/{{zone}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{zone}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/nodeGroup:NodeGroup")
public class NodeGroup extends com.pulumi.resources.CustomResource {
    /**
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.
     * Structure is documented below.
     * 
     */
    @Export(name="autoscalingPolicy", type=NodeGroupAutoscalingPolicy.class, parameters={})
    private Output<NodeGroupAutoscalingPolicy> autoscalingPolicy;

    /**
     * @return If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.
     * Structure is documented below.
     * 
     */
    public Output<NodeGroupAutoscalingPolicy> autoscalingPolicy() {
        return this.autoscalingPolicy;
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
     * An optional textual description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional textual description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
     * 
     */
    @Export(name="initialSize", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> initialSize;

    /**
     * @return The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
     * 
     */
    public Output<Optional<Integer>> initialSize() {
        return Codegen.optional(this.initialSize);
    }
    /**
     * Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
     * 
     */
    @Export(name="maintenancePolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> maintenancePolicy;

    /**
     * @return Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
     * 
     */
    public Output<Optional<String>> maintenancePolicy() {
        return Codegen.optional(this.maintenancePolicy);
    }
    /**
     * contains properties for the timeframe of maintenance
     * Structure is documented below.
     * 
     */
    @Export(name="maintenanceWindow", type=NodeGroupMaintenanceWindow.class, parameters={})
    private Output</* @Nullable */ NodeGroupMaintenanceWindow> maintenanceWindow;

    /**
     * @return contains properties for the timeframe of maintenance
     * Structure is documented below.
     * 
     */
    public Output<Optional<NodeGroupMaintenanceWindow>> maintenanceWindow() {
        return Codegen.optional(this.maintenanceWindow);
    }
    /**
     * Name of the resource.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The URL of the node template to which this node group belongs.
     * 
     */
    @Export(name="nodeTemplate", type=String.class, parameters={})
    private Output<String> nodeTemplate;

    /**
     * @return The URL of the node template to which this node group belongs.
     * 
     */
    public Output<String> nodeTemplate() {
        return this.nodeTemplate;
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
     * Share settings for the node group.
     * Structure is documented below.
     * 
     */
    @Export(name="shareSettings", type=NodeGroupShareSettings.class, parameters={})
    private Output<NodeGroupShareSettings> shareSettings;

    /**
     * @return Share settings for the node group.
     * Structure is documented below.
     * 
     */
    public Output<NodeGroupShareSettings> shareSettings() {
        return this.shareSettings;
    }
    /**
     * The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
     * 
     */
    @Export(name="size", type=Integer.class, parameters={})
    private Output<Integer> size;

    /**
     * @return The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * Zone where this node group is located
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return Zone where this node group is located
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NodeGroup(String name) {
        this(name, NodeGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NodeGroup(String name, NodeGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NodeGroup(String name, NodeGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/nodeGroup:NodeGroup", name, args == null ? NodeGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NodeGroup(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/nodeGroup:NodeGroup", name, state, makeResourceOptions(options, id));
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
    public static NodeGroup get(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NodeGroup(name, id, state, options);
    }
}
