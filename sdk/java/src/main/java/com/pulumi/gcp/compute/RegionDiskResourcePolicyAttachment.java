// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RegionDiskResourcePolicyAttachmentArgs;
import com.pulumi.gcp.compute.inputs.RegionDiskResourcePolicyAttachmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Adds existing resource policies to a disk. You can only add one policy
 * which will be applied to this disk for scheduling snapshot creation.
 * 
 * &gt; **Note:** This resource does not support zonal disks (`gcp.compute.Disk`). For zonal disks, please refer to the `gcp.compute.DiskResourcePolicyAttachment` resource.
 * 
 * ## Example Usage
 * ### Region Disk Resource Policy Attachment Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Disk;
 * import com.pulumi.gcp.compute.DiskArgs;
 * import com.pulumi.gcp.compute.Snapshot;
 * import com.pulumi.gcp.compute.SnapshotArgs;
 * import com.pulumi.gcp.compute.RegionDisk;
 * import com.pulumi.gcp.compute.RegionDiskArgs;
 * import com.pulumi.gcp.compute.RegionDiskResourcePolicyAttachment;
 * import com.pulumi.gcp.compute.RegionDiskResourcePolicyAttachmentArgs;
 * import com.pulumi.gcp.compute.ResourcePolicy;
 * import com.pulumi.gcp.compute.ResourcePolicyArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs;
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
 *         var disk = new Disk(&#34;disk&#34;, DiskArgs.builder()        
 *             .image(&#34;debian-cloud/debian-9&#34;)
 *             .size(50)
 *             .type(&#34;pd-ssd&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .build());
 * 
 *         var snapdisk = new Snapshot(&#34;snapdisk&#34;, SnapshotArgs.builder()        
 *             .sourceDisk(disk.name())
 *             .zone(&#34;us-central1-a&#34;)
 *             .build());
 * 
 *         var ssd = new RegionDisk(&#34;ssd&#34;, RegionDiskArgs.builder()        
 *             .replicaZones(            
 *                 &#34;us-central1-a&#34;,
 *                 &#34;us-central1-f&#34;)
 *             .snapshot(snapdisk.id())
 *             .size(50)
 *             .type(&#34;pd-ssd&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var attachment = new RegionDiskResourcePolicyAttachment(&#34;attachment&#34;, RegionDiskResourcePolicyAttachmentArgs.builder()        
 *             .disk(ssd.name())
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var policy = new ResourcePolicy(&#34;policy&#34;, ResourcePolicyArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .snapshotSchedulePolicy(ResourcePolicySnapshotSchedulePolicyArgs.builder()
 *                 .schedule(ResourcePolicySnapshotSchedulePolicyScheduleArgs.builder()
 *                     .dailySchedule(ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs.builder()
 *                         .daysInCycle(1)
 *                         .startTime(&#34;04:00&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RegionDiskResourcePolicyAttachment can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{project}}/{{region}}/{{disk}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{region}}/{{disk}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{disk}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment")
public class RegionDiskResourcePolicyAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The name of the regional disk in which the resource policies are attached to.
     * 
     */
    @Export(name="disk", type=String.class, parameters={})
    private Output<String> disk;

    /**
     * @return The name of the regional disk in which the resource policies are attached to.
     * 
     */
    public Output<String> disk() {
        return this.disk;
    }
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
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
     * A reference to the region where the disk resides.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return A reference to the region where the disk resides.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionDiskResourcePolicyAttachment(String name) {
        this(name, RegionDiskResourcePolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionDiskResourcePolicyAttachment(String name, RegionDiskResourcePolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionDiskResourcePolicyAttachment(String name, RegionDiskResourcePolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, args == null ? RegionDiskResourcePolicyAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RegionDiskResourcePolicyAttachment(String name, Output<String> id, @Nullable RegionDiskResourcePolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, state, makeResourceOptions(options, id));
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
    public static RegionDiskResourcePolicyAttachment get(String name, Output<String> id, @Nullable RegionDiskResourcePolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionDiskResourcePolicyAttachment(name, id, state, options);
    }
}
