// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceGroupArgs;
import com.pulumi.gcp.compute.inputs.InstanceGroupState;
import com.pulumi.gcp.compute.outputs.InstanceGroupNamedPort;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a group of dissimilar Compute Engine virtual machine instances.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
 * 
 * ## Example Usage
 * 
 * ### Empty Instance Group
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceGroup;
 * import com.pulumi.gcp.compute.InstanceGroupArgs;
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
 *         var test = new InstanceGroup("test", InstanceGroupArgs.builder()
 *             .name("test")
 *             .description("Test instance group")
 *             .zone("us-central1-a")
 *             .network(default_.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example Usage - With instances and named ports
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceGroup;
 * import com.pulumi.gcp.compute.InstanceGroupArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupNamedPortArgs;
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
 *         var webservers = new InstanceGroup("webservers", InstanceGroupArgs.builder()
 *             .name("webservers")
 *             .description("Test instance group")
 *             .instances(            
 *                 test.id(),
 *                 test2.id())
 *             .namedPorts(            
 *                 InstanceGroupNamedPortArgs.builder()
 *                     .name("http")
 *                     .port("8080")
 *                     .build(),
 *                 InstanceGroupNamedPortArgs.builder()
 *                     .name("https")
 *                     .port("8443")
 *                     .build())
 *             .zone("us-central1-a")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example Usage - Recreating an instance group in use
 * Recreating an instance group that&#39;s in use by another resource will give a
 * `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
 * as shown in this example to avoid this type of error.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.InstanceGroup;
 * import com.pulumi.gcp.compute.InstanceGroupArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupNamedPortArgs;
 * import com.pulumi.gcp.compute.HttpsHealthCheck;
 * import com.pulumi.gcp.compute.HttpsHealthCheckArgs;
 * import com.pulumi.gcp.compute.BackendService;
 * import com.pulumi.gcp.compute.BackendServiceArgs;
 * import com.pulumi.gcp.compute.inputs.BackendServiceBackendArgs;
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
 *         final var debianImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family("debian-11")
 *             .project("debian-cloud")
 *             .build());
 * 
 *         var stagingVm = new Instance("stagingVm", InstanceArgs.builder()
 *             .name("staging-vm")
 *             .machineType("e2-medium")
 *             .zone("us-central1-c")
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(debianImage.applyValue(getImageResult -> getImageResult.selfLink()))
 *                     .build())
 *                 .build())
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .network("default")
 *                 .build())
 *             .build());
 * 
 *         var stagingGroup = new InstanceGroup("stagingGroup", InstanceGroupArgs.builder()
 *             .name("staging-instance-group")
 *             .zone("us-central1-c")
 *             .instances(stagingVm.id())
 *             .namedPorts(            
 *                 InstanceGroupNamedPortArgs.builder()
 *                     .name("http")
 *                     .port("8080")
 *                     .build(),
 *                 InstanceGroupNamedPortArgs.builder()
 *                     .name("https")
 *                     .port("8443")
 *                     .build())
 *             .build());
 * 
 *         var stagingHealth = new HttpsHealthCheck("stagingHealth", HttpsHealthCheckArgs.builder()
 *             .name("staging-health")
 *             .requestPath("/health_check")
 *             .build());
 * 
 *         var stagingService = new BackendService("stagingService", BackendServiceArgs.builder()
 *             .name("staging-service")
 *             .portName("https")
 *             .protocol("HTTPS")
 *             .backends(BackendServiceBackendArgs.builder()
 *                 .group(stagingGroup.id())
 *                 .build())
 *             .healthChecks(stagingHealth.id())
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
 * Instance groups can be imported using the `zone` and `name` with an optional `project`, e.g.
 * 
 * * `projects/{{project_id}}/zones/{{zone}}/instanceGroups/{{instance_group_id}}`
 * 
 * * `{{project_id}}/{{zone}}/{{instance_group_id}}`
 * 
 * * `{{zone}}/{{instance_group_id}}`
 * 
 * When using the `pulumi import` command, instance groups can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceGroup:InstanceGroup default {{zone}}/{{instance_group_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceGroup:InstanceGroup default {{project_id}}/{{zone}}/{{instance_group_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceGroup:InstanceGroup default projects/{{project_id}}/zones/{{zone}}/instanceGroups/{{instance_group_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/instanceGroup:InstanceGroup")
public class InstanceGroup extends com.pulumi.resources.CustomResource {
    /**
     * An optional textual description of the instance
     * group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional textual description of the instance
     * group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The list of instances in the group, in `self_link` format.
     * When adding instances they must all be in the same network and zone as the instance group.
     * 
     */
    @Export(name="instances", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> instances;

    /**
     * @return The list of instances in the group, in `self_link` format.
     * When adding instances they must all be in the same network and zone as the instance group.
     * 
     */
    public Output<List<String>> instances() {
        return this.instances;
    }
    /**
     * The name of the instance group. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the instance group. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The named port configuration. See the section below
     * for details on configuration. Structure is documented below.
     * 
     */
    @Export(name="namedPorts", refs={List.class,InstanceGroupNamedPort.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceGroupNamedPort>> namedPorts;

    /**
     * @return The named port configuration. See the section below
     * for details on configuration. Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceGroupNamedPort>>> namedPorts() {
        return Codegen.optional(this.namedPorts);
    }
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * The number of instances in the group.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The number of instances in the group.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The zone that this instance group should be created in.
     * 
     * ***
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The zone that this instance group should be created in.
     * 
     * ***
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceGroup(String name) {
        this(name, InstanceGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceGroup(String name, @Nullable InstanceGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceGroup(String name, @Nullable InstanceGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceGroup:InstanceGroup", name, args == null ? InstanceGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceGroup(String name, Output<String> id, @Nullable InstanceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceGroup:InstanceGroup", name, state, makeResourceOptions(options, id));
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
    public static InstanceGroup get(String name, Output<String> id, @Nullable InstanceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceGroup(name, id, state, options);
    }
}
