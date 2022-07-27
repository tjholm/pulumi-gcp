// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.NodeTemplateArgs;
import com.pulumi.gcp.compute.inputs.NodeTemplateState;
import com.pulumi.gcp.compute.outputs.NodeTemplateNodeTypeFlexibility;
import com.pulumi.gcp.compute.outputs.NodeTemplateServerBinding;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a NodeTemplate resource. Node templates specify properties
 * for creating sole-tenant nodes, such as node type, vCPU and memory
 * requirements, node affinity labels, and region.
 * 
 * To get more information about NodeTemplate, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
 * * How-to Guides
 *     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
 * 
 * ## Example Usage
 * ### Node Template Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.NodeTemplate;
 * import com.pulumi.gcp.compute.NodeTemplateArgs;
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
 *         var template = new NodeTemplate(&#34;template&#34;, NodeTemplateArgs.builder()        
 *             .nodeType(&#34;n1-node-96-624&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Node Template Server Binding
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetNodeTypesArgs;
 * import com.pulumi.gcp.compute.NodeTemplate;
 * import com.pulumi.gcp.compute.NodeTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.NodeTemplateServerBindingArgs;
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
 *         final var central1a = ComputeFunctions.getNodeTypes(GetNodeTypesArgs.builder()
 *             .zone(&#34;us-central1-a&#34;)
 *             .build());
 * 
 *         var template = new NodeTemplate(&#34;template&#34;, NodeTemplateArgs.builder()        
 *             .nodeAffinityLabels(Map.of(&#34;foo&#34;, &#34;baz&#34;))
 *             .nodeType(&#34;n1-node-96-624&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .serverBinding(NodeTemplateServerBindingArgs.builder()
 *                 .type(&#34;RESTART_NODE_ON_MINIMAL_SERVERS&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * NodeTemplate can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/nodeTemplate:NodeTemplate")
public class NodeTemplate extends com.pulumi.resources.CustomResource {
    /**
     * CPU overcommit.
     * Default value is `NONE`.
     * Possible values are `ENABLED` and `NONE`.
     * 
     */
    @Export(name="cpuOvercommitType", type=String.class, parameters={})
    private Output</* @Nullable */ String> cpuOvercommitType;

    /**
     * @return CPU overcommit.
     * Default value is `NONE`.
     * Possible values are `ENABLED` and `NONE`.
     * 
     */
    public Output<Optional<String>> cpuOvercommitType() {
        return Codegen.optional(this.cpuOvercommitType);
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
     * Labels to use for node affinity, which will be used in
     * instance scheduling.
     * 
     */
    @Export(name="nodeAffinityLabels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> nodeAffinityLabels;

    /**
     * @return Labels to use for node affinity, which will be used in
     * instance scheduling.
     * 
     */
    public Output<Optional<Map<String,String>>> nodeAffinityLabels() {
        return Codegen.optional(this.nodeAffinityLabels);
    }
    /**
     * Node type to use for nodes group that are created from this template.
     * Only one of nodeTypeFlexibility and nodeType can be specified.
     * 
     */
    @Export(name="nodeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> nodeType;

    /**
     * @return Node type to use for nodes group that are created from this template.
     * Only one of nodeTypeFlexibility and nodeType can be specified.
     * 
     */
    public Output<Optional<String>> nodeType() {
        return Codegen.optional(this.nodeType);
    }
    /**
     * Flexible properties for the desired node type. Node groups that
     * use this node template will create nodes of a type that matches
     * these properties. Only one of nodeTypeFlexibility and nodeType can
     * be specified.
     * Structure is documented below.
     * 
     */
    @Export(name="nodeTypeFlexibility", type=NodeTemplateNodeTypeFlexibility.class, parameters={})
    private Output</* @Nullable */ NodeTemplateNodeTypeFlexibility> nodeTypeFlexibility;

    /**
     * @return Flexible properties for the desired node type. Node groups that
     * use this node template will create nodes of a type that matches
     * these properties. Only one of nodeTypeFlexibility and nodeType can
     * be specified.
     * Structure is documented below.
     * 
     */
    public Output<Optional<NodeTemplateNodeTypeFlexibility>> nodeTypeFlexibility() {
        return Codegen.optional(this.nodeTypeFlexibility);
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
     * Region where nodes using the node template will be created.
     * If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return Region where nodes using the node template will be created.
     * If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * The server binding policy for nodes using this template. Determines
     * where the nodes should restart following a maintenance event.
     * Structure is documented below.
     * 
     */
    @Export(name="serverBinding", type=NodeTemplateServerBinding.class, parameters={})
    private Output<NodeTemplateServerBinding> serverBinding;

    /**
     * @return The server binding policy for nodes using this template. Determines
     * where the nodes should restart following a maintenance event.
     * Structure is documented below.
     * 
     */
    public Output<NodeTemplateServerBinding> serverBinding() {
        return this.serverBinding;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NodeTemplate(String name) {
        this(name, NodeTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NodeTemplate(String name, @Nullable NodeTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NodeTemplate(String name, @Nullable NodeTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/nodeTemplate:NodeTemplate", name, args == null ? NodeTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NodeTemplate(String name, Output<String> id, @Nullable NodeTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/nodeTemplate:NodeTemplate", name, state, makeResourceOptions(options, id));
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
    public static NodeTemplate get(String name, Output<String> id, @Nullable NodeTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NodeTemplate(name, id, state, options);
    }
}
