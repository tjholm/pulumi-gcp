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
 * ### Region Autoscaler Basic
 * ```java
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
 *         var foobarInstanceTemplate = new InstanceTemplate(&#34;foobarInstanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .machineType(&#34;e2-standard-4&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(&#34;debian-cloud/debian-11&#34;)
 *                 .diskSizeGb(250)
 *                 .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .accessConfigs(InstanceTemplateNetworkInterfaceAccessConfigArgs.builder()
 *                     .networkTier(&#34;PREMIUM&#34;)
 *                     .build())
 *                 .build())
 *             .serviceAccount(InstanceTemplateServiceAccountArgs.builder()
 *                 .scopes(                
 *                     &#34;https://www.googleapis.com/auth/devstorage.read_only&#34;,
 *                     &#34;https://www.googleapis.com/auth/logging.write&#34;,
 *                     &#34;https://www.googleapis.com/auth/monitoring.write&#34;,
 *                     &#34;https://www.googleapis.com/auth/pubsub&#34;,
 *                     &#34;https://www.googleapis.com/auth/service.management.readonly&#34;,
 *                     &#34;https://www.googleapis.com/auth/servicecontrol&#34;,
 *                     &#34;https://www.googleapis.com/auth/trace.append&#34;)
 *                 .build())
 *             .build());
 * 
 *         var foobarTargetPool = new TargetPool(&#34;foobarTargetPool&#34;);
 * 
 *         var foobarRegionInstanceGroupManager = new RegionInstanceGroupManager(&#34;foobarRegionInstanceGroupManager&#34;, RegionInstanceGroupManagerArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .versions(RegionInstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(foobarInstanceTemplate.id())
 *                 .name(&#34;primary&#34;)
 *                 .build())
 *             .targetPools(foobarTargetPool.id())
 *             .baseInstanceName(&#34;foobar&#34;)
 *             .build());
 * 
 *         var foobarRegionAutoscaler = new RegionAutoscaler(&#34;foobarRegionAutoscaler&#34;, RegionAutoscalerArgs.builder()        
 *             .region(&#34;us-central1&#34;)
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
 *             .family(&#34;debian-11&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RegionAutoscaler can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/autoscalers/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, RegionAutoscaler can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{name}}
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
    public RegionAutoscaler(String name) {
        this(name, RegionAutoscalerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionAutoscaler(String name, RegionAutoscalerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionAutoscaler(String name, RegionAutoscalerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionAutoscaler:RegionAutoscaler", name, args == null ? RegionAutoscalerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RegionAutoscaler(String name, Output<String> id, @Nullable RegionAutoscalerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionAutoscaler:RegionAutoscaler", name, state, makeResourceOptions(options, id));
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
    public static RegionAutoscaler get(String name, Output<String> id, @Nullable RegionAutoscalerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionAutoscaler(name, id, state, options);
    }
}
