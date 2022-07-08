// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.PerInstanceConfigArgs;
import com.pulumi.gcp.compute.inputs.PerInstanceConfigState;
import com.pulumi.gcp.compute.outputs.PerInstanceConfigPreservedState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A config defined for a single managed instance that belongs to an instance group manager. It preserves the instance name
 * across instance group manager operations and can define stateful disks or metadata that are unique to the instance.
 * 
 * To get more information about PerInstanceConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/stateful-migs#per-instance_configs)
 * 
 * ## Example Usage
 * ### Stateful Igm
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var myImage = Output.of(ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build()));
 * 
 *         var igm_basic = new InstanceTemplate(&#34;igm-basic&#34;, InstanceTemplateArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .canIpForward(false)
 *             .tags(            
 *                 &#34;foo&#34;,
 *                 &#34;bar&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(myImage.apply(getImageResult -&gt; getImageResult.selfLink()))
 *                 .autoDelete(true)
 *                 .boot(true)
 *                 .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .serviceAccount(InstanceTemplateServiceAccountArgs.builder()
 *                 .scopes(                
 *                     &#34;userinfo-email&#34;,
 *                     &#34;compute-ro&#34;,
 *                     &#34;storage-ro&#34;)
 *                 .build())
 *             .build());
 * 
 *         var igm_no_tp = new InstanceGroupManager(&#34;igm-no-tp&#34;, InstanceGroupManagerArgs.builder()        
 *             .description(&#34;Test instance group manager&#34;)
 *             .versions(InstanceGroupManagerVersionArgs.builder()
 *                 .name(&#34;prod&#34;)
 *                 .instanceTemplate(igm_basic.selfLink())
 *                 .build())
 *             .baseInstanceName(&#34;igm-no-tp&#34;)
 *             .zone(&#34;us-central1-c&#34;)
 *             .targetSize(2)
 *             .build());
 * 
 *         var default_ = new Disk(&#34;default&#34;, DiskArgs.builder()        
 *             .type(&#34;pd-ssd&#34;)
 *             .zone(google_compute_instance_group_manager.igm().zone())
 *             .image(&#34;debian-8-jessie-v20170523&#34;)
 *             .physicalBlockSizeBytes(4096)
 *             .build());
 * 
 *         var withDisk = new PerInstanceConfig(&#34;withDisk&#34;, PerInstanceConfigArgs.builder()        
 *             .zone(google_compute_instance_group_manager.igm().zone())
 *             .instanceGroupManager(google_compute_instance_group_manager.igm().name())
 *             .preservedState(PerInstanceConfigPreservedStateArgs.builder()
 *                 .metadata(Map.ofEntries(
 *                     Map.entry(&#34;foo&#34;, &#34;bar&#34;),
 *                     Map.entry(&#34;instance_template&#34;, igm_basic.selfLink())
 *                 ))
 *                 .disks(PerInstanceConfigPreservedStateDiskArgs.builder()
 *                     .deviceName(&#34;my-stateful-disk&#34;)
 *                     .source(default_.id())
 *                     .mode(&#34;READ_ONLY&#34;)
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
 * PerInstanceConfig can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{project}}/{{zone}}/{{instance_group_manager}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{zone}}/{{instance_group_manager}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/perInstanceConfig:PerInstanceConfig default {{instance_group_manager}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/perInstanceConfig:PerInstanceConfig")
public class PerInstanceConfig extends com.pulumi.resources.CustomResource {
    /**
     * The instance group manager this instance config is part of.
     * 
     */
    @Export(name="instanceGroupManager", type=String.class, parameters={})
    private Output<String> instanceGroupManager;

    /**
     * @return The instance group manager this instance config is part of.
     * 
     */
    public Output<String> instanceGroupManager() {
        return this.instanceGroupManager;
    }
    /**
     * The minimal action to perform on the instance during an update.
     * Default is `NONE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     * 
     */
    @Export(name="minimalAction", type=String.class, parameters={})
    private Output</* @Nullable */ String> minimalAction;

    /**
     * @return The minimal action to perform on the instance during an update.
     * Default is `NONE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     * 
     */
    public Output<Optional<String>> minimalAction() {
        return Codegen.optional(this.minimalAction);
    }
    /**
     * The most disruptive action to perform on the instance during an update.
     * Default is `REPLACE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     * 
     */
    @Export(name="mostDisruptiveAllowedAction", type=String.class, parameters={})
    private Output</* @Nullable */ String> mostDisruptiveAllowedAction;

    /**
     * @return The most disruptive action to perform on the instance during an update.
     * Default is `REPLACE`. Possible values are:
     * * REPLACE
     * * RESTART
     * * REFRESH
     * * NONE
     * 
     */
    public Output<Optional<String>> mostDisruptiveAllowedAction() {
        return Codegen.optional(this.mostDisruptiveAllowedAction);
    }
    /**
     * The name for this per-instance config and its corresponding instance.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name for this per-instance config and its corresponding instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The preserved state for this instance.
     * Structure is documented below.
     * 
     */
    @Export(name="preservedState", type=PerInstanceConfigPreservedState.class, parameters={})
    private Output</* @Nullable */ PerInstanceConfigPreservedState> preservedState;

    /**
     * @return The preserved state for this instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PerInstanceConfigPreservedState>> preservedState() {
        return Codegen.optional(this.preservedState);
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
     * When true, deleting this config will immediately remove any specified state from the underlying instance.
     * When false, deleting this config will *not* immediately remove any state from the underlying instance.
     * State will be removed on the next instance recreation or update.
     * 
     */
    @Export(name="removeInstanceStateOnDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> removeInstanceStateOnDestroy;

    /**
     * @return When true, deleting this config will immediately remove any specified state from the underlying instance.
     * When false, deleting this config will *not* immediately remove any state from the underlying instance.
     * State will be removed on the next instance recreation or update.
     * 
     */
    public Output<Optional<Boolean>> removeInstanceStateOnDestroy() {
        return Codegen.optional(this.removeInstanceStateOnDestroy);
    }
    /**
     * Zone where the containing instance group manager is located
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output</* @Nullable */ String> zone;

    /**
     * @return Zone where the containing instance group manager is located
     * 
     */
    public Output<Optional<String>> zone() {
        return Codegen.optional(this.zone);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PerInstanceConfig(String name) {
        this(name, PerInstanceConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PerInstanceConfig(String name, PerInstanceConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PerInstanceConfig(String name, PerInstanceConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/perInstanceConfig:PerInstanceConfig", name, args == null ? PerInstanceConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PerInstanceConfig(String name, Output<String> id, @Nullable PerInstanceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/perInstanceConfig:PerInstanceConfig", name, state, makeResourceOptions(options, id));
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
    public static PerInstanceConfig get(String name, Output<String> id, @Nullable PerInstanceConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PerInstanceConfig(name, id, state, options);
    }
}
