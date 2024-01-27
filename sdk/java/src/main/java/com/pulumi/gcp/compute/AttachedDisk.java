// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.AttachedDiskArgs;
import com.pulumi.gcp.compute.inputs.AttachedDiskState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Persistent disks can be attached to a compute instance using the `attached_disk`
 * section within the compute instance configuration.
 * However there may be situations where managing the attached disks via the compute
 * instance config isn&#39;t preferable or possible, such as attaching dynamic
 * numbers of disks using the `count` variable.
 * 
 * To get more information about attaching disks, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
 * * How-to Guides
 *     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
 * 
 * **Note:** When using `gcp.compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = [&#34;attached_disk&#34;]` on the `gcp.compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.AttachedDisk;
 * import com.pulumi.gcp.compute.AttachedDiskArgs;
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
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .zone(&#34;us-west1-a&#34;)
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(&#34;debian-cloud/debian-11&#34;)
 *                     .build())
 *                 .build())
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .build());
 * 
 *         var defaultAttachedDisk = new AttachedDisk(&#34;defaultAttachedDisk&#34;, AttachedDiskArgs.builder()        
 *             .disk(google_compute_disk.default().id())
 *             .instance(defaultInstance.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Attached Disk can be imported the following ways* `projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}` * `{{project}}/{{zone}}/{{instance.name}}/{{disk.name}}` When using the `pulumi import` command, Attached Disk can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/attachedDisk:AttachedDisk default projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/attachedDisk:AttachedDisk default {{project}}/{{zone}}/{{instance.name}}/{{disk.name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/attachedDisk:AttachedDisk")
public class AttachedDisk extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a unique device name of your choice that is
     * reflected into the /dev/disk/by-id/google-* tree of a Linux operating
     * system running within the instance. This name can be used to
     * reference the device for mounting, resizing, and so on, from within
     * the instance.
     * 
     * If not specified, the server chooses a default device name to apply
     * to this disk, in the form persistent-disks-x, where x is a number
     * assigned by Google Compute Engine.
     * 
     */
    @Export(name="deviceName", refs={String.class}, tree="[0]")
    private Output<String> deviceName;

    /**
     * @return Specifies a unique device name of your choice that is
     * reflected into the /dev/disk/by-id/google-* tree of a Linux operating
     * system running within the instance. This name can be used to
     * reference the device for mounting, resizing, and so on, from within
     * the instance.
     * 
     * If not specified, the server chooses a default device name to apply
     * to this disk, in the form persistent-disks-x, where x is a number
     * assigned by Google Compute Engine.
     * 
     */
    public Output<String> deviceName() {
        return this.deviceName;
    }
    /**
     * `name` or `self_link` of the disk that will be attached.
     * 
     * ***
     * 
     */
    @Export(name="disk", refs={String.class}, tree="[0]")
    private Output<String> disk;

    /**
     * @return `name` or `self_link` of the disk that will be attached.
     * 
     * ***
     * 
     */
    public Output<String> disk() {
        return this.disk;
    }
    /**
     * `name` or `self_link` of the compute instance that the disk will be attached to.
     * If the `self_link` is provided then `zone` and `project` are extracted from the
     * self link. If only the name is used then `zone` and `project` must be defined
     * as properties on the resource or provider.
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output<String> instance;

    /**
     * @return `name` or `self_link` of the compute instance that the disk will be attached to.
     * If the `self_link` is provided then `zone` and `project` are extracted from the
     * self link. If only the name is used then `zone` and `project` must be defined
     * as properties on the resource or provider.
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * The mode in which to attach this disk, either READ_WRITE or
     * READ_ONLY. If not specified, the default is to attach the disk in
     * READ_WRITE mode.
     * 
     * Possible values:
     * &#34;READ_ONLY&#34;
     * &#34;READ_WRITE&#34;
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return The mode in which to attach this disk, either READ_WRITE or
     * READ_ONLY. If not specified, the default is to attach the disk in
     * READ_WRITE mode.
     * 
     * Possible values:
     * &#34;READ_ONLY&#34;
     * &#34;READ_WRITE&#34;
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * The project that the referenced compute instance is a part of. If `instance` is referenced by its
     * `self_link` the project defined in the link will take precedence.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project that the referenced compute instance is a part of. If `instance` is referenced by its
     * `self_link` the project defined in the link will take precedence.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The zone that the referenced compute instance is located within. If `instance` is referenced by its
     * `self_link` the zone defined in the link will take precedence.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The zone that the referenced compute instance is located within. If `instance` is referenced by its
     * `self_link` the zone defined in the link will take precedence.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AttachedDisk(String name) {
        this(name, AttachedDiskArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AttachedDisk(String name, AttachedDiskArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AttachedDisk(String name, AttachedDiskArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/attachedDisk:AttachedDisk", name, args == null ? AttachedDiskArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AttachedDisk(String name, Output<String> id, @Nullable AttachedDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/attachedDisk:AttachedDisk", name, state, makeResourceOptions(options, id));
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
    public static AttachedDisk get(String name, Output<String> id, @Nullable AttachedDiskState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AttachedDisk(name, id, state, options);
    }
}
