// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceArgs;
import com.pulumi.gcp.compute.inputs.InstanceState;
import com.pulumi.gcp.compute.outputs.InstanceAdvancedMachineFeatures;
import com.pulumi.gcp.compute.outputs.InstanceAttachedDisk;
import com.pulumi.gcp.compute.outputs.InstanceBootDisk;
import com.pulumi.gcp.compute.outputs.InstanceConfidentialInstanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceGuestAccelerator;
import com.pulumi.gcp.compute.outputs.InstanceNetworkInterface;
import com.pulumi.gcp.compute.outputs.InstanceNetworkPerformanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceParams;
import com.pulumi.gcp.compute.outputs.InstanceReservationAffinity;
import com.pulumi.gcp.compute.outputs.InstanceScheduling;
import com.pulumi.gcp.compute.outputs.InstanceScratchDisk;
import com.pulumi.gcp.compute.outputs.InstanceServiceAccount;
import com.pulumi.gcp.compute.outputs.InstanceShieldedInstanceConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceScratchDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceServiceAccountArgs;
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
 *         var defaultAccount = new Account(&#34;defaultAccount&#34;, AccountArgs.builder()        
 *             .accountId(&#34;service_account_id&#34;)
 *             .displayName(&#34;Service Account&#34;)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .tags(            
 *                 &#34;foo&#34;,
 *                 &#34;bar&#34;)
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(&#34;debian-cloud/debian-11&#34;)
 *                     .labels(Map.of(&#34;my_label&#34;, &#34;value&#34;))
 *                     .build())
 *                 .build())
 *             .scratchDisks(InstanceScratchDiskArgs.builder()
 *                 .interface_(&#34;SCSI&#34;)
 *                 .build())
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .accessConfigs()
 *                 .build())
 *             .metadata(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .metadataStartupScript(&#34;echo hi &gt; /test.txt&#34;)
 *             .serviceAccount(InstanceServiceAccountArgs.builder()
 *                 .email(defaultAccount.email())
 *                 .scopes(&#34;cloud-platform&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Instances can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instance:Instance default projects/{{project}}/zones/{{zone}}/instances/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instance:Instance default {{project}}/{{zone}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instance:Instance default {{name}}
 * ```
 * 
 *  [custom-vm-types]https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]https://cloud.google.com/network-tiers/docs/overview [extended-custom-vm-type]https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#extendedmemory
 * 
 */
@ResourceType(type="gcp:compute/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * Configure Nested Virtualisation and Simultaneous Hyper Threading  on this VM. Structure is documented below
     * 
     */
    @Export(name="advancedMachineFeatures", refs={InstanceAdvancedMachineFeatures.class}, tree="[0]")
    private Output</* @Nullable */ InstanceAdvancedMachineFeatures> advancedMachineFeatures;

    /**
     * @return Configure Nested Virtualisation and Simultaneous Hyper Threading  on this VM. Structure is documented below
     * 
     */
    public Output<Optional<InstanceAdvancedMachineFeatures>> advancedMachineFeatures() {
        return Codegen.optional(this.advancedMachineFeatures);
    }
    /**
     * If true, allows this prvider to stop the instance to update its properties.
     * If you try to update a property that requires stopping the instance without setting this field, the update will fail.
     * 
     */
    @Export(name="allowStoppingForUpdate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowStoppingForUpdate;

    /**
     * @return If true, allows this prvider to stop the instance to update its properties.
     * If you try to update a property that requires stopping the instance without setting this field, the update will fail.
     * 
     */
    public Output<Optional<Boolean>> allowStoppingForUpdate() {
        return Codegen.optional(this.allowStoppingForUpdate);
    }
    /**
     * Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
     * 
     */
    @Export(name="attachedDisks", refs={List.class,InstanceAttachedDisk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceAttachedDisk>> attachedDisks;

    /**
     * @return Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceAttachedDisk>>> attachedDisks() {
        return Codegen.optional(this.attachedDisks);
    }
    /**
     * The boot disk for the instance.
     * Structure is documented below.
     * 
     */
    @Export(name="bootDisk", refs={InstanceBootDisk.class}, tree="[0]")
    private Output<InstanceBootDisk> bootDisk;

    /**
     * @return The boot disk for the instance.
     * Structure is documented below.
     * 
     */
    public Output<InstanceBootDisk> bootDisk() {
        return this.bootDisk;
    }
    /**
     * Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs.
     * This defaults to false.
     * 
     */
    @Export(name="canIpForward", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> canIpForward;

    /**
     * @return Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs.
     * This defaults to false.
     * 
     */
    public Output<Optional<Boolean>> canIpForward() {
        return Codegen.optional(this.canIpForward);
    }
    /**
     * Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    @Export(name="confidentialInstanceConfig", refs={InstanceConfidentialInstanceConfig.class}, tree="[0]")
    private Output<InstanceConfidentialInstanceConfig> confidentialInstanceConfig;

    /**
     * @return Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    public Output<InstanceConfidentialInstanceConfig> confidentialInstanceConfig() {
        return this.confidentialInstanceConfig;
    }
    /**
     * The CPU platform used by this instance.
     * 
     */
    @Export(name="cpuPlatform", refs={String.class}, tree="[0]")
    private Output<String> cpuPlatform;

    /**
     * @return The CPU platform used by this instance.
     * 
     */
    public Output<String> cpuPlatform() {
        return this.cpuPlatform;
    }
    /**
     * The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
     * 
     */
    @Export(name="currentStatus", refs={String.class}, tree="[0]")
    private Output<String> currentStatus;

    /**
     * @return The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
     * 
     */
    public Output<String> currentStatus() {
        return this.currentStatus;
    }
    /**
     * Enable deletion protection on this instance. Defaults to false.
     * **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return Enable deletion protection on this instance. Defaults to false.
     * **Note:** you must disable deletion protection before removing the resource (e.g., via `pulumi destroy`), or the instance cannot be deleted and the provider run will not complete successfully.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * A brief description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Desired status of the instance. Either
     * `&#34;RUNNING&#34;` or `&#34;TERMINATED&#34;`.
     * 
     */
    @Export(name="desiredStatus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desiredStatus;

    /**
     * @return Desired status of the instance. Either
     * `&#34;RUNNING&#34;` or `&#34;TERMINATED&#34;`.
     * 
     */
    public Output<Optional<String>> desiredStatus() {
        return Codegen.optional(this.desiredStatus);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    @Export(name="enableDisplay", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDisplay;

    /**
     * @return Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    public Output<Optional<Boolean>> enableDisplay() {
        return Codegen.optional(this.enableDisplay);
    }
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * **Note:** GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
     * 
     */
    @Export(name="guestAccelerators", refs={List.class,InstanceGuestAccelerator.class}, tree="[0,1]")
    private Output<List<InstanceGuestAccelerator>> guestAccelerators;

    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * **Note:** GPU accelerators can only be used with `on_host_maintenance` option set to TERMINATE.
     * 
     */
    public Output<List<InstanceGuestAccelerator>> guestAccelerators() {
        return this.guestAccelerators;
    }
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
     * Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
     * The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostname;

    /**
     * @return A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid.
     * Valid format is a series of labels 1-63 characters long matching the regular expression `a-z`, concatenated with periods.
     * The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> hostname() {
        return Codegen.optional(this.hostname);
    }
    /**
     * The server-assigned unique identifier of this instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
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
    @Export(name="labelFingerprint", refs={String.class}, tree="[0]")
    private Output<String> labelFingerprint;

    /**
     * @return The unique fingerprint of the labels.
     * 
     */
    public Output<String> labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * A map of key/value label pairs to assign to the instance.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A map of key/value label pairs to assign to the instance.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The machine type to create.
     * 
     * **Note:** If you want to update this value (resize the VM) after initial creation, you must set `allow_stopping_for_update` to `true`.
     * 
     * [Custom machine types][custom-vm-types] can be formatted as `custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY_MB`, e.g. `custom-6-20480` for 6 vCPU and 20GB of RAM.
     * 
     * There is a limit of 6.5 GB per CPU unless you add [extended memory][extended-custom-vm-type]. You must do this explicitly by adding the suffix `-ext`, e.g. `custom-2-15360-ext` for 2 vCPU and 15 GB of memory.
     * 
     */
    @Export(name="machineType", refs={String.class}, tree="[0]")
    private Output<String> machineType;

    /**
     * @return The machine type to create.
     * 
     * **Note:** If you want to update this value (resize the VM) after initial creation, you must set `allow_stopping_for_update` to `true`.
     * 
     * [Custom machine types][custom-vm-types] can be formatted as `custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY_MB`, e.g. `custom-6-20480` for 6 vCPU and 20GB of RAM.
     * 
     * There is a limit of 6.5 GB per CPU unless you add [extended memory][extended-custom-vm-type]. You must do this explicitly by adding the suffix `-ext`, e.g. `custom-2-15360-ext` for 2 vCPU and 15 GB of memory.
     * 
     */
    public Output<String> machineType() {
        return this.machineType;
    }
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Ssh keys attached in the Cloud Console will be removed.
     * Add them to your config in order to keep them attached to your instance. A
     * list of default metadata values (e.g. ssh-keys) can be found [here](https://cloud.google.com/compute/docs/metadata/default-metadata-values)
     * 
     * &gt; Depending on the OS you choose for your instance, some metadata keys have
     * special functionality.  Most linux-based images will run the content of
     * `metadata.startup-script` in a shell on every boot.  At a minimum,
     * Debian, CentOS, RHEL, SLES, Container-Optimized OS, and Ubuntu images
     * support this key.  Windows instances require other keys depending on the format
     * of the script and the time you would like it to run - see [this table](https://cloud.google.com/compute/docs/startupscript#providing_a_startup_script_for_windows_instances).
     * For the convenience of the users of `metadata.startup-script`,
     * we provide a special attribute, `metadata_startup_script`, which is documented below.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> metadata;

    /**
     * @return Metadata key/value pairs to make available from
     * within the instance. Ssh keys attached in the Cloud Console will be removed.
     * Add them to your config in order to keep them attached to your instance. A
     * list of default metadata values (e.g. ssh-keys) can be found [here](https://cloud.google.com/compute/docs/metadata/default-metadata-values)
     * 
     * &gt; Depending on the OS you choose for your instance, some metadata keys have
     * special functionality.  Most linux-based images will run the content of
     * `metadata.startup-script` in a shell on every boot.  At a minimum,
     * Debian, CentOS, RHEL, SLES, Container-Optimized OS, and Ubuntu images
     * support this key.  Windows instances require other keys depending on the format
     * of the script and the time you would like it to run - see [this table](https://cloud.google.com/compute/docs/startupscript#providing_a_startup_script_for_windows_instances).
     * For the convenience of the users of `metadata.startup-script`,
     * we provide a special attribute, `metadata_startup_script`, which is documented below.
     * 
     */
    public Output<Optional<Map<String,String>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * The unique fingerprint of the metadata.
     * 
     */
    @Export(name="metadataFingerprint", refs={String.class}, tree="[0]")
    private Output<String> metadataFingerprint;

    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    public Output<String> metadataFingerprint() {
        return this.metadataFingerprint;
    }
    /**
     * An alternative to using the
     * startup-script metadata key, except this one forces the instance to be recreated
     * (thus re-running the script) if it is changed. This replaces the startup-script
     * metadata key on the created instance and thus the two mechanisms are not
     * allowed to be used simultaneously.  Users are free to use either mechanism - the
     * only distinction is that this separate attribute will cause a recreate on
     * modification.  On import, `metadata_startup_script` will not be set - if you
     * choose to specify it you will see a diff immediately after import causing a
     * destroy/recreate operation. If importing an instance and specifying this value
     * is desired, you will need to modify your state file.
     * 
     */
    @Export(name="metadataStartupScript", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> metadataStartupScript;

    /**
     * @return An alternative to using the
     * startup-script metadata key, except this one forces the instance to be recreated
     * (thus re-running the script) if it is changed. This replaces the startup-script
     * metadata key on the created instance and thus the two mechanisms are not
     * allowed to be used simultaneously.  Users are free to use either mechanism - the
     * only distinction is that this separate attribute will cause a recreate on
     * modification.  On import, `metadata_startup_script` will not be set - if you
     * choose to specify it you will see a diff immediately after import causing a
     * destroy/recreate operation. If importing an instance and specifying this value
     * is desired, you will need to modify your state file.
     * 
     */
    public Output<Optional<String>> metadataStartupScript() {
        return Codegen.optional(this.metadataStartupScript);
    }
    /**
     * Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    @Export(name="minCpuPlatform", refs={String.class}, tree="[0]")
    private Output<String> minCpuPlatform;

    /**
     * @return Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
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
    @Export(name="name", refs={String.class}, tree="[0]")
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
     * Networks to attach to the instance. This can
     * be specified multiple times. Structure is documented below.
     * 
     * ***
     * 
     */
    @Export(name="networkInterfaces", refs={List.class,InstanceNetworkInterface.class}, tree="[0,1]")
    private Output<List<InstanceNetworkInterface>> networkInterfaces;

    /**
     * @return Networks to attach to the instance. This can
     * be specified multiple times. Structure is documented below.
     * 
     * ***
     * 
     */
    public Output<List<InstanceNetworkInterface>> networkInterfaces() {
        return this.networkInterfaces;
    }
    /**
     * Configures network performance settings for the instance. Structure is
     * documented below. **Note**: `machine_type` must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
     * the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
     * in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
     * in order for this setting to take effect.
     * 
     */
    @Export(name="networkPerformanceConfig", refs={InstanceNetworkPerformanceConfig.class}, tree="[0]")
    private Output</* @Nullable */ InstanceNetworkPerformanceConfig> networkPerformanceConfig;

    /**
     * @return Configures network performance settings for the instance. Structure is
     * documented below. **Note**: `machine_type` must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
     * the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
     * in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
     * in order for this setting to take effect.
     * 
     */
    public Output<Optional<InstanceNetworkPerformanceConfig>> networkPerformanceConfig() {
        return Codegen.optional(this.networkPerformanceConfig);
    }
    /**
     * Additional instance parameters.
     * .
     * 
     */
    @Export(name="params", refs={InstanceParams.class}, tree="[0]")
    private Output</* @Nullable */ InstanceParams> params;

    /**
     * @return Additional instance parameters.
     * .
     * 
     */
    public Output<Optional<InstanceParams>> params() {
        return Codegen.optional(this.params);
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
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Specifies the reservations that this instance can consume from.
     * Structure is documented below.
     * 
     */
    @Export(name="reservationAffinity", refs={InstanceReservationAffinity.class}, tree="[0]")
    private Output<InstanceReservationAffinity> reservationAffinity;

    /**
     * @return Specifies the reservations that this instance can consume from.
     * Structure is documented below.
     * 
     */
    public Output<InstanceReservationAffinity> reservationAffinity() {
        return this.reservationAffinity;
    }
    /**
     * - A list of self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
     * 
     */
    @Export(name="resourcePolicies", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourcePolicies;

    /**
     * @return - A list of self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
     * 
     */
    public Output<Optional<String>> resourcePolicies() {
        return Codegen.optional(this.resourcePolicies);
    }
    /**
     * The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    @Export(name="scheduling", refs={InstanceScheduling.class}, tree="[0]")
    private Output<InstanceScheduling> scheduling;

    /**
     * @return The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    public Output<InstanceScheduling> scheduling() {
        return this.scheduling;
    }
    /**
     * Scratch disks to attach to the instance. This can be
     * specified multiple times for multiple scratch disks. Structure is documented below.
     * 
     */
    @Export(name="scratchDisks", refs={List.class,InstanceScratchDisk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceScratchDisk>> scratchDisks;

    /**
     * @return Scratch disks to attach to the instance. This can be
     * specified multiple times for multiple scratch disks. Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceScratchDisk>>> scratchDisks() {
        return Codegen.optional(this.scratchDisks);
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
     * Service account to attach to the instance.
     * Structure is documented below.
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    @Export(name="serviceAccount", refs={InstanceServiceAccount.class}, tree="[0]")
    private Output</* @Nullable */ InstanceServiceAccount> serviceAccount;

    /**
     * @return Service account to attach to the instance.
     * Structure is documented below.
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    public Output<Optional<InstanceServiceAccount>> serviceAccount() {
        return Codegen.optional(this.serviceAccount);
    }
    /**
     * Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    @Export(name="shieldedInstanceConfig", refs={InstanceShieldedInstanceConfig.class}, tree="[0]")
    private Output<InstanceShieldedInstanceConfig> shieldedInstanceConfig;

    /**
     * @return Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * **Note**: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
     * 
     */
    public Output<InstanceShieldedInstanceConfig> shieldedInstanceConfig() {
        return this.shieldedInstanceConfig;
    }
    /**
     * A list of network tags to attach to the instance.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of network tags to attach to the instance.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The unique fingerprint of the tags.
     * 
     */
    @Export(name="tagsFingerprint", refs={String.class}, tree="[0]")
    private Output<String> tagsFingerprint;

    /**
     * @return The unique fingerprint of the tags.
     * 
     */
    public Output<String> tagsFingerprint() {
        return this.tagsFingerprint;
    }
    /**
     * The zone that the machine should be created in. If it is not provided, the provider zone is used.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The zone that the machine should be created in. If it is not provided, the provider zone is used.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
