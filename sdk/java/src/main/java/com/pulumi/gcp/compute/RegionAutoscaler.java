// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RegionAutoscalerArgs;
import com.pulumi.gcp.compute.inputs.RegionAutoscalerState;
import com.pulumi.gcp.compute.outputs.RegionAutoscalerAutoscalingPolicy;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents an Autoscaler resource.
 * 
 * Autoscalers allow you to automatically scale virtual machine instances in
 * managed instance groups according to an autoscaling policy that you
 * define.
 * 
 * To get more information about RegionAutoscaler, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
 * * How-to Guides
 *     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
 * 
 * ## Example Usage
 * 
 * ### Region Autoscaler Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateServiceAccountArgs;
 * import com.pulumi.gcp.compute.TargetPool;
 * import com.pulumi.gcp.compute.TargetPoolArgs;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManager;
 * import com.pulumi.gcp.compute.RegionInstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.RegionInstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.compute.RegionAutoscaler;
 * import com.pulumi.gcp.compute.RegionAutoscalerArgs;
 * import com.pulumi.gcp.compute.inputs.RegionAutoscalerAutoscalingPolicyArgs;
 * import com.pulumi.gcp.compute.inputs.RegionAutoscalerAutoscalingPolicyCpuUtilizationArgs;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
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
 *         var foobarInstanceTemplate = new InstanceTemplate("foobarInstanceTemplate", InstanceTemplateArgs.builder()
 *             .name("my-instance-template")
 *             .machineType("e2-standard-4")
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage("debian-cloud/debian-11")
 *                 .diskSizeGb(250)
 *                 .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network("default")
 *                 .accessConfigs(InstanceTemplateNetworkInterfaceAccessConfigArgs.builder()
 *                     .networkTier("PREMIUM")
 *                     .build())
 *                 .build())
 *             .serviceAccount(InstanceTemplateServiceAccountArgs.builder()
 *                 .scopes(                
 *                     "https://www.googleapis.com/auth/devstorage.read_only",
 *                     "https://www.googleapis.com/auth/logging.write",
 *                     "https://www.googleapis.com/auth/monitoring.write",
 *                     "https://www.googleapis.com/auth/pubsub",
 *                     "https://www.googleapis.com/auth/service.management.readonly",
 *                     "https://www.googleapis.com/auth/servicecontrol",
 *                     "https://www.googleapis.com/auth/trace.append")
 *                 .build())
 *             .build());
 * 
 *         var foobarTargetPool = new TargetPool("foobarTargetPool", TargetPoolArgs.builder()
 *             .name("my-target-pool")
 *             .build());
 * 
 *         var foobarRegionInstanceGroupManager = new RegionInstanceGroupManager("foobarRegionInstanceGroupManager", RegionInstanceGroupManagerArgs.builder()
 *             .name("my-region-igm")
 *             .region("us-central1")
 *             .versions(RegionInstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(foobarInstanceTemplate.id())
 *                 .name("primary")
 *                 .build())
 *             .targetPools(foobarTargetPool.id())
 *             .baseInstanceName("foobar")
 *             .build());
 * 
 *         var foobar = new RegionAutoscaler("foobar", RegionAutoscalerArgs.builder()
 *             .name("my-region-autoscaler")
 *             .region("us-central1")
 *             .target(foobarRegionInstanceGroupManager.id())
 *             .autoscalingPolicy(RegionAutoscalerAutoscalingPolicyArgs.builder()
 *                 .maxReplicas(5)
 *                 .minReplicas(1)
 *                 .cooldownPeriod(60)
 *                 .cpuUtilization(RegionAutoscalerAutoscalingPolicyCpuUtilizationArgs.builder()
 *                     .target(0.5)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         final var debian9 = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family("debian-11")
 *             .project("debian-cloud")
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
 * RegionAutoscaler can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/regions/{{region}}/autoscalers/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, RegionAutoscaler can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/regionAutoscaler:RegionAutoscaler")
public class RegionAutoscaler extends com.pulumi.resources.CustomResource {
    /**
     * The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     * 
     */
    @Export(name="autoscalingPolicy", refs={RegionAutoscalerAutoscalingPolicy.class}, tree="[0]")
    private Output<RegionAutoscalerAutoscalingPolicy> autoscalingPolicy;

    /**
     * @return The configuration parameters for the autoscaling algorithm. You can
     * define one or more of the policies for an autoscaler: cpuUtilization,
     * customMetricUtilizations, and loadBalancingUtilization.
     * If none of these are specified, the default will be to autoscale based
     * on cpuUtilization to 0.6 or 60%.
     * Structure is documented below.
     * 
     */
    public Output<RegionAutoscalerAutoscalingPolicy> autoscalingPolicy() {
        return this.autoscalingPolicy;
    }
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTimestamp", refs={String.class}, tree="[0]")
    private Output<String> creationTimestamp;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Output<String> creationTimestamp() {
        return this.creationTimestamp;
    }
    /**
     * An optional description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Name of the resource. The name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. The name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * URL of the region where the instance group resides.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return URL of the region where the instance group resides.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * URL of the managed instance group that this autoscaler will scale.
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output<String> target;

    /**
     * @return URL of the managed instance group that this autoscaler will scale.
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionAutoscaler(java.lang.String name) {
        this(name, RegionAutoscalerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionAutoscaler(java.lang.String name, RegionAutoscalerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionAutoscaler(java.lang.String name, RegionAutoscalerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionAutoscaler:RegionAutoscaler", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionAutoscaler(java.lang.String name, Output<java.lang.String> id, @Nullable RegionAutoscalerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionAutoscaler:RegionAutoscaler", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionAutoscalerArgs makeArgs(RegionAutoscalerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionAutoscalerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static RegionAutoscaler get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionAutoscalerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionAutoscaler(name, id, state, options);
    }
}
