// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceFromTemplateArgs;
import com.pulumi.gcp.compute.inputs.InstanceFromTemplateState;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateAdvancedMachineFeatures;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateAttachedDisk;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateBootDisk;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateConfidentialInstanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateGuestAccelerator;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateNetworkInterface;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateNetworkPerformanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateReservationAffinity;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateScheduling;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateScratchDisk;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateServiceAccount;
import com.pulumi.gcp.compute.outputs.InstanceFromTemplateShieldedInstanceConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Manages a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 * 
 * This resource is specifically to create a compute instance from a given
 * `source_instance_template`. To create an instance without a template, use the
 * `gcp.compute.Instance` resource.
 * 
 * ## Example Usage
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
 * import com.pulumi.gcp.compute.InstanceFromTemplate;
 * import com.pulumi.gcp.compute.InstanceFromTemplateArgs;
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
 *         var tplInstanceTemplate = new InstanceTemplate(&#34;tplInstanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(&#34;debian-cloud/debian-11&#34;)
 *                 .autoDelete(true)
 *                 .diskSizeGb(100)
 *                 .boot(true)
 *                 .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .metadata(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .canIpForward(true)
 *             .build());
 * 
 *         var tplInstanceFromTemplate = new InstanceFromTemplate(&#34;tplInstanceFromTemplate&#34;, InstanceFromTemplateArgs.builder()        
 *             .zone(&#34;us-central1-a&#34;)
 *             .sourceInstanceTemplate(tplInstanceTemplate.selfLinkUnique())
 *             .canIpForward(false)
 *             .labels(Map.of(&#34;my_key&#34;, &#34;my_value&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:compute/instanceFromTemplate:InstanceFromTemplate")
public class InstanceFromTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Controls for advanced machine-related behavior features.
     * 
     */
    @Export(name="advancedMachineFeatures", type=InstanceFromTemplateAdvancedMachineFeatures.class, parameters={})
    private Output<InstanceFromTemplateAdvancedMachineFeatures> advancedMachineFeatures;

    /**
     * @return Controls for advanced machine-related behavior features.
     * 
     */
    public Output<InstanceFromTemplateAdvancedMachineFeatures> advancedMachineFeatures() {
        return this.advancedMachineFeatures;
    }
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     * 
     */
    @Export(name="allowStoppingForUpdate", type=Boolean.class, parameters={})
    private Output<Boolean> allowStoppingForUpdate;

    /**
     * @return If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     * 
     */
    public Output<Boolean> allowStoppingForUpdate() {
        return this.allowStoppingForUpdate;
    }
    /**
     * List of disks attached to the instance
     * 
     */
    @Export(name="attachedDisks", type=List.class, parameters={InstanceFromTemplateAttachedDisk.class})
    private Output<List<InstanceFromTemplateAttachedDisk>> attachedDisks;

    /**
     * @return List of disks attached to the instance
     * 
     */
    public Output<List<InstanceFromTemplateAttachedDisk>> attachedDisks() {
        return this.attachedDisks;
    }
    /**
     * The boot disk for the instance.
     * 
     */
    @Export(name="bootDisk", type=InstanceFromTemplateBootDisk.class, parameters={})
    private Output<InstanceFromTemplateBootDisk> bootDisk;

    /**
     * @return The boot disk for the instance.
     * 
     */
    public Output<InstanceFromTemplateBootDisk> bootDisk() {
        return this.bootDisk;
    }
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     * 
     */
    @Export(name="canIpForward", type=Boolean.class, parameters={})
    private Output<Boolean> canIpForward;

    /**
     * @return Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     * 
     */
    public Output<Boolean> canIpForward() {
        return this.canIpForward;
    }
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     * 
     */
    @Export(name="confidentialInstanceConfig", type=InstanceFromTemplateConfidentialInstanceConfig.class, parameters={})
    private Output<InstanceFromTemplateConfidentialInstanceConfig> confidentialInstanceConfig;

    /**
     * @return The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     * 
     */
    public Output<InstanceFromTemplateConfidentialInstanceConfig> confidentialInstanceConfig() {
        return this.confidentialInstanceConfig;
    }
    /**
     * The CPU platform used by this instance.
     * 
     */
    @Export(name="cpuPlatform", type=String.class, parameters={})
    private Output<String> cpuPlatform;

    /**
     * @return The CPU platform used by this instance.
     * 
     */
    public Output<String> cpuPlatform() {
        return this.cpuPlatform;
    }
    /**
     * Current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING,
     * SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance
     * life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).
     * 
     */
    @Export(name="currentStatus", type=String.class, parameters={})
    private Output<String> currentStatus;

    /**
     * @return Current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING,
     * SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance
     * life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).
     * 
     */
    public Output<String> currentStatus() {
        return this.currentStatus;
    }
    /**
     * Whether deletion protection is enabled on this instance.
     * 
     */
    @Export(name="deletionProtection", type=Boolean.class, parameters={})
    private Output<Boolean> deletionProtection;

    /**
     * @return Whether deletion protection is enabled on this instance.
     * 
     */
    public Output<Boolean> deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * A brief description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return A brief description of the resource.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Desired status of the instance. Either &#34;RUNNING&#34; or &#34;TERMINATED&#34;.
     * 
     */
    @Export(name="desiredStatus", type=String.class, parameters={})
    private Output<String> desiredStatus;

    /**
     * @return Desired status of the instance. Either &#34;RUNNING&#34; or &#34;TERMINATED&#34;.
     * 
     */
    public Output<String> desiredStatus() {
        return this.desiredStatus;
    }
    /**
     * Whether the instance has virtual displays enabled.
     * 
     */
    @Export(name="enableDisplay", type=Boolean.class, parameters={})
    private Output<Boolean> enableDisplay;

    /**
     * @return Whether the instance has virtual displays enabled.
     * 
     */
    public Output<Boolean> enableDisplay() {
        return this.enableDisplay;
    }
    /**
     * List of the type and count of accelerator cards attached to the instance.
     * 
     */
    @Export(name="guestAccelerators", type=List.class, parameters={InstanceFromTemplateGuestAccelerator.class})
    private Output<List<InstanceFromTemplateGuestAccelerator>> guestAccelerators;

    /**
     * @return List of the type and count of accelerator cards attached to the instance.
     * 
     */
    public Output<List<InstanceFromTemplateGuestAccelerator>> guestAccelerators() {
        return this.guestAccelerators;
    }
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="hostname", type=String.class, parameters={})
    private Output<String> hostname;

    /**
     * @return A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * The server-assigned unique identifier of this instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The server-assigned unique identifier of this instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The unique fingerprint of the labels.
     * 
     */
    @Export(name="labelFingerprint", type=String.class, parameters={})
    private Output<String> labelFingerprint;

    /**
     * @return The unique fingerprint of the labels.
     * 
     */
    public Output<String> labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * A set of key/value label pairs assigned to the instance.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs assigned to the instance.
     * 
     */
    public Output<Map<String,String>> labels() {
        return this.labels;
    }
    /**
     * The machine type to create.
     * 
     */
    @Export(name="machineType", type=String.class, parameters={})
    private Output<String> machineType;

    /**
     * @return The machine type to create.
     * 
     */
    public Output<String> machineType() {
        return this.machineType;
    }
    /**
     * Metadata key/value pairs made available within the instance.
     * 
     */
    @Export(name="metadata", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> metadata;

    /**
     * @return Metadata key/value pairs made available within the instance.
     * 
     */
    public Output<Map<String,String>> metadata() {
        return this.metadata;
    }
    /**
     * The unique fingerprint of the metadata.
     * 
     */
    @Export(name="metadataFingerprint", type=String.class, parameters={})
    private Output<String> metadataFingerprint;

    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    public Output<String> metadataFingerprint() {
        return this.metadataFingerprint;
    }
    /**
     * Metadata startup scripts made available within the instance.
     * 
     */
    @Export(name="metadataStartupScript", type=String.class, parameters={})
    private Output<String> metadataStartupScript;

    /**
     * @return Metadata startup scripts made available within the instance.
     * 
     */
    public Output<String> metadataStartupScript() {
        return this.metadataStartupScript;
    }
    /**
     * The minimum CPU platform specified for the VM instance.
     * 
     */
    @Export(name="minCpuPlatform", type=String.class, parameters={})
    private Output<String> minCpuPlatform;

    /**
     * @return The minimum CPU platform specified for the VM instance.
     * 
     */
    public Output<String> minCpuPlatform() {
        return this.minCpuPlatform;
    }
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The networks attached to the instance.
     * 
     */
    @Export(name="networkInterfaces", type=List.class, parameters={InstanceFromTemplateNetworkInterface.class})
    private Output<List<InstanceFromTemplateNetworkInterface>> networkInterfaces;

    /**
     * @return The networks attached to the instance.
     * 
     */
    public Output<List<InstanceFromTemplateNetworkInterface>> networkInterfaces() {
        return this.networkInterfaces;
    }
    /**
     * Configures network performance settings for the instance. If not specified, the instance will be created with its
     * default network performance configuration.
     * 
     */
    @Export(name="networkPerformanceConfig", type=InstanceFromTemplateNetworkPerformanceConfig.class, parameters={})
    private Output<InstanceFromTemplateNetworkPerformanceConfig> networkPerformanceConfig;

    /**
     * @return Configures network performance settings for the instance. If not specified, the instance will be created with its
     * default network performance configuration.
     * 
     */
    public Output<InstanceFromTemplateNetworkPerformanceConfig> networkPerformanceConfig() {
        return this.networkPerformanceConfig;
    }
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Specifies the reservations that this instance can consume from.
     * 
     */
    @Export(name="reservationAffinity", type=InstanceFromTemplateReservationAffinity.class, parameters={})
    private Output<InstanceFromTemplateReservationAffinity> reservationAffinity;

    /**
     * @return Specifies the reservations that this instance can consume from.
     * 
     */
    public Output<InstanceFromTemplateReservationAffinity> reservationAffinity() {
        return this.reservationAffinity;
    }
    /**
     * A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.
     * 
     */
    @Export(name="resourcePolicies", type=String.class, parameters={})
    private Output<String> resourcePolicies;

    /**
     * @return A list of self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.
     * 
     */
    public Output<String> resourcePolicies() {
        return this.resourcePolicies;
    }
    /**
     * The scheduling strategy being used by the instance.
     * 
     */
    @Export(name="scheduling", type=InstanceFromTemplateScheduling.class, parameters={})
    private Output<InstanceFromTemplateScheduling> scheduling;

    /**
     * @return The scheduling strategy being used by the instance.
     * 
     */
    public Output<InstanceFromTemplateScheduling> scheduling() {
        return this.scheduling;
    }
    /**
     * The scratch disks attached to the instance.
     * 
     */
    @Export(name="scratchDisks", type=List.class, parameters={InstanceFromTemplateScratchDisk.class})
    private Output<List<InstanceFromTemplateScratchDisk>> scratchDisks;

    /**
     * @return The scratch disks attached to the instance.
     * 
     */
    public Output<List<InstanceFromTemplateScratchDisk>> scratchDisks() {
        return this.scratchDisks;
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
     * The service account to attach to the instance.
     * 
     */
    @Export(name="serviceAccount", type=InstanceFromTemplateServiceAccount.class, parameters={})
    private Output<InstanceFromTemplateServiceAccount> serviceAccount;

    /**
     * @return The service account to attach to the instance.
     * 
     */
    public Output<InstanceFromTemplateServiceAccount> serviceAccount() {
        return this.serviceAccount;
    }
    /**
     * The shielded vm config being used by the instance.
     * 
     */
    @Export(name="shieldedInstanceConfig", type=InstanceFromTemplateShieldedInstanceConfig.class, parameters={})
    private Output<InstanceFromTemplateShieldedInstanceConfig> shieldedInstanceConfig;

    /**
     * @return The shielded vm config being used by the instance.
     * 
     */
    public Output<InstanceFromTemplateShieldedInstanceConfig> shieldedInstanceConfig() {
        return this.shieldedInstanceConfig;
    }
    /**
     * Name or self link of an instance
     * template to create the instance based on. It is recommended to reference
     * instance templates through their unique id (`self_link_unique` attribute).
     * 
     */
    @Export(name="sourceInstanceTemplate", type=String.class, parameters={})
    private Output<String> sourceInstanceTemplate;

    /**
     * @return Name or self link of an instance
     * template to create the instance based on. It is recommended to reference
     * instance templates through their unique id (`self_link_unique` attribute).
     * 
     */
    public Output<String> sourceInstanceTemplate() {
        return this.sourceInstanceTemplate;
    }
    /**
     * The list of tags attached to the instance.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output<List<String>> tags;

    /**
     * @return The list of tags attached to the instance.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * The unique fingerprint of the tags.
     * 
     */
    @Export(name="tagsFingerprint", type=String.class, parameters={})
    private Output<String> tagsFingerprint;

    /**
     * @return The unique fingerprint of the tags.
     * 
     */
    public Output<String> tagsFingerprint() {
        return this.tagsFingerprint;
    }
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return The zone that the machine should be created in. If not
     * set, the provider zone is used.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceFromTemplate(String name) {
        this(name, InstanceFromTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceFromTemplate(String name, InstanceFromTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceFromTemplate(String name, InstanceFromTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, args == null ? InstanceFromTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceFromTemplate(String name, Output<String> id, @Nullable InstanceFromTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, state, makeResourceOptions(options, id));
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
    public static InstanceFromTemplate get(String name, Output<String> id, @Nullable InstanceFromTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceFromTemplate(name, id, state, options);
    }
}
