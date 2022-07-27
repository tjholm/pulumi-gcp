// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceTemplateArgs;
import com.pulumi.gcp.compute.inputs.InstanceTemplateState;
import com.pulumi.gcp.compute.outputs.InstanceTemplateAdvancedMachineFeatures;
import com.pulumi.gcp.compute.outputs.InstanceTemplateConfidentialInstanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceTemplateDisk;
import com.pulumi.gcp.compute.outputs.InstanceTemplateGuestAccelerator;
import com.pulumi.gcp.compute.outputs.InstanceTemplateNetworkInterface;
import com.pulumi.gcp.compute.outputs.InstanceTemplateNetworkPerformanceConfig;
import com.pulumi.gcp.compute.outputs.InstanceTemplateReservationAffinity;
import com.pulumi.gcp.compute.outputs.InstanceTemplateScheduling;
import com.pulumi.gcp.compute.outputs.InstanceTemplateServiceAccount;
import com.pulumi.gcp.compute.outputs.InstanceTemplateShieldedInstanceConfig;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a VM instance template resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceAccount.Account;
 * import com.pulumi.gcp.serviceAccount.AccountArgs;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.Disk;
 * import com.pulumi.gcp.compute.DiskArgs;
 * import com.pulumi.gcp.compute.ResourcePolicy;
 * import com.pulumi.gcp.compute.ResourcePolicyArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs;
 * import com.pulumi.gcp.compute.inputs.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateSchedulingArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateServiceAccountArgs;
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
 *             .accountId(&#34;service-account-id&#34;)
 *             .displayName(&#34;Service Account&#34;)
 *             .build());
 * 
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *         var foobar = new Disk(&#34;foobar&#34;, DiskArgs.builder()        
 *             .image(myImage.applyValue(getImageResult -&gt; getImageResult.selfLink()))
 *             .size(10)
 *             .type(&#34;pd-ssd&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .build());
 * 
 *         var dailyBackup = new ResourcePolicy(&#34;dailyBackup&#34;, ResourcePolicyArgs.builder()        
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
 *         var defaultInstanceTemplate = new InstanceTemplate(&#34;defaultInstanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .description(&#34;This template is used to create app server instances.&#34;)
 *             .tags(            
 *                 &#34;foo&#34;,
 *                 &#34;bar&#34;)
 *             .labels(Map.of(&#34;environment&#34;, &#34;dev&#34;))
 *             .instanceDescription(&#34;description assigned to instances&#34;)
 *             .machineType(&#34;e2-medium&#34;)
 *             .canIpForward(false)
 *             .scheduling(InstanceTemplateSchedulingArgs.builder()
 *                 .automaticRestart(true)
 *                 .onHostMaintenance(&#34;MIGRATE&#34;)
 *                 .build())
 *             .disks(            
 *                 InstanceTemplateDiskArgs.builder()
 *                     .sourceImage(&#34;debian-cloud/debian-9&#34;)
 *                     .autoDelete(true)
 *                     .boot(true)
 *                     .resourcePolicies(dailyBackup.id())
 *                     .build(),
 *                 InstanceTemplateDiskArgs.builder()
 *                     .source(foobar.name())
 *                     .autoDelete(false)
 *                     .boot(false)
 *                     .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .metadata(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .serviceAccount(InstanceTemplateServiceAccountArgs.builder()
 *                 .email(defaultAccount.email())
 *                 .scopes(&#34;cloud-platform&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Automatic Envoy Deployment
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateSchedulingArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateServiceAccountArgs;
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
 *         final var default = ComputeFunctions.getDefaultServiceAccount();
 * 
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *         var foobar = new InstanceTemplate(&#34;foobar&#34;, InstanceTemplateArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .canIpForward(false)
 *             .tags(            
 *                 &#34;foo&#34;,
 *                 &#34;bar&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(myImage.applyValue(getImageResult -&gt; getImageResult.selfLink()))
 *                 .autoDelete(true)
 *                 .boot(true)
 *                 .build())
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .scheduling(InstanceTemplateSchedulingArgs.builder()
 *                 .preemptible(false)
 *                 .automaticRestart(true)
 *                 .build())
 *             .metadata(Map.ofEntries(
 *                 Map.entry(&#34;gce-software-declaration&#34;, &#34;&#34;&#34;
 * {
 *   &#34;softwareRecipes&#34;: [{
 *     &#34;name&#34;: &#34;install-gce-service-proxy-agent&#34;,
 *     &#34;desired_state&#34;: &#34;INSTALLED&#34;,
 *     &#34;installSteps&#34;: [{
 *       &#34;scriptRun&#34;: {
 *         &#34;script&#34;: &#34;#! /bin/bash\nZONE=$(curl --silent http://metadata.google.internal/computeMetadata/v1/instance/zone -H Metadata-Flavor:Google | cut -d/ -f4 )\nexport SERVICE_PROXY_AGENT_DIRECTORY=$(mktemp -d)\nsudo gsutil cp   gs://gce-service-proxy-&#34;$ZONE&#34;/service-proxy-agent/releases/service-proxy-agent-0.2.tgz   &#34;$SERVICE_PROXY_AGENT_DIRECTORY&#34;   || sudo gsutil cp     gs://gce-service-proxy/service-proxy-agent/releases/service-proxy-agent-0.2.tgz     &#34;$SERVICE_PROXY_AGENT_DIRECTORY&#34;\nsudo tar -xzf &#34;$SERVICE_PROXY_AGENT_DIRECTORY&#34;/service-proxy-agent-0.2.tgz -C &#34;$SERVICE_PROXY_AGENT_DIRECTORY&#34;\n&#34;$SERVICE_PROXY_AGENT_DIRECTORY&#34;/service-proxy-agent/service-proxy-agent-bootstrap.sh&#34;
 *       }
 *     }]
 *   }]
 * }
 *                 &#34;&#34;&#34;),
 *                 Map.entry(&#34;gce-service-proxy&#34;, &#34;&#34;&#34;
 * {
 *   &#34;api-version&#34;: &#34;0.2&#34;,
 *   &#34;proxy-spec&#34;: {
 *     &#34;proxy-port&#34;: 15001,
 *     &#34;network&#34;: &#34;my-network&#34;,
 *     &#34;tracing&#34;: &#34;ON&#34;,
 *     &#34;access-log&#34;: &#34;/var/log/envoy/access.log&#34;
 *   }
 *   &#34;service&#34;: {
 *     &#34;serving-ports&#34;: [80, 81]
 *   },
 *  &#34;labels&#34;: {
 *    &#34;app_name&#34;: &#34;bookserver_app&#34;,
 *    &#34;app_version&#34;: &#34;STABLE&#34;
 *   }
 * }
 *                 &#34;&#34;&#34;),
 *                 Map.entry(&#34;enable-guest-attributes&#34;, &#34;true&#34;),
 *                 Map.entry(&#34;enable-osconfig&#34;, &#34;true&#34;)
 *             ))
 *             .serviceAccount(InstanceTemplateServiceAccountArgs.builder()
 *                 .email(default_.email())
 *                 .scopes(&#34;cloud-platform&#34;)
 *                 .build())
 *             .labels(Map.of(&#34;gce-service-proxy&#34;, &#34;on&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Using with Instance Group Manager
 * 
 * Instance Templates cannot be updated after creation with the Google
 * Cloud Platform API. In order to update an Instance Template, this provider will
 * create a replacement. In order to effectively
 * use an Instance Template resource with an [Instance Group Manager resource](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html).
 * Either omit the Instance Template `name` attribute, or specify a partial name
 * with `name_prefix`. Example:
 * 
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
 * import com.pulumi.gcp.compute.InstanceGroupManager;
 * import com.pulumi.gcp.compute.InstanceGroupManagerArgs;
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
 *         var instanceTemplate = new InstanceTemplate(&#34;instanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .namePrefix(&#34;instance-template-&#34;)
 *             .machineType(&#34;e2-medium&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .disks()
 *             .networkInterfaces()
 *             .build());
 * 
 *         var instanceGroupManager = new InstanceGroupManager(&#34;instanceGroupManager&#34;, InstanceGroupManagerArgs.builder()        
 *             .instanceTemplate(instanceTemplate.id())
 *             .baseInstanceName(&#34;instance-group-manager&#34;)
 *             .zone(&#34;us-central1-f&#34;)
 *             .targetSize(&#34;1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With this setup, this provider generates a unique name for your Instance
 * Template and can then update the Instance Group manager without conflict before
 * destroying the previous Instance Template.
 * 
 * ## Deploying the Latest Image
 * 
 * A common way to use instance templates and managed instance groups is to deploy the
 * latest image in a family, usually the latest build of your application. There are two
 * ways to do this in the provider, and they have their pros and cons. The difference ends
 * up being in how &#34;latest&#34; is interpreted. You can either deploy the latest image available
 * when the provider runs, or you can have each instance check what the latest image is when
 * it&#39;s being created, either as part of a scaling event or being rebuilt by the instance
 * group manager.
 * 
 * If you&#39;re not sure, we recommend deploying the latest image available when the provider runs,
 * because this means all the instances in your group will be based on the same image, always,
 * and means that no upgrades or changes to your instances happen outside of a `pulumi up`.
 * You can achieve this by using the `gcp.compute.Image`
 * data source, which will retrieve the latest image on every `pulumi apply`, and will update
 * the template to use that specific image:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
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
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *         var instanceTemplate = new InstanceTemplate(&#34;instanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .namePrefix(&#34;instance-template-&#34;)
 *             .machineType(&#34;e2-medium&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(google_compute_image.my_image().self_link())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * To have instances update to the latest on every scaling event or instance re-creation,
 * use the family as the image for the disk, and it will use GCP&#39;s default behavior, setting
 * the image for the template to the family:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
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
 *         var instanceTemplate = new InstanceTemplate(&#34;instanceTemplate&#34;, InstanceTemplateArgs.builder()        
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage(&#34;debian-cloud/debian-9&#34;)
 *                 .build())
 *             .machineType(&#34;e2-medium&#34;)
 *             .namePrefix(&#34;instance-template-&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Instance templates can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default projects/{{project}}/global/instanceTemplates/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/instanceTemplate:InstanceTemplate default {{name}}
 * ```
 * 
 *  [custom-vm-types]https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]https://cloud.google.com/network-tiers/docs/overview
 * 
 */
@ResourceType(type="gcp:compute/instanceTemplate:InstanceTemplate")
public class InstanceTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Configure Nested Virtualisation and Simultaneous Hyper Threading on this VM. Structure is documented below
     * 
     */
    @Export(name="advancedMachineFeatures", type=InstanceTemplateAdvancedMachineFeatures.class, parameters={})
    private Output</* @Nullable */ InstanceTemplateAdvancedMachineFeatures> advancedMachineFeatures;

    /**
     * @return Configure Nested Virtualisation and Simultaneous Hyper Threading on this VM. Structure is documented below
     * 
     */
    public Output<Optional<InstanceTemplateAdvancedMachineFeatures>> advancedMachineFeatures() {
        return Codegen.optional(this.advancedMachineFeatures);
    }
    /**
     * Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     * 
     */
    @Export(name="canIpForward", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> canIpForward;

    /**
     * @return Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     * 
     */
    public Output<Optional<Boolean>> canIpForward() {
        return Codegen.optional(this.canIpForward);
    }
    /**
     * Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    @Export(name="confidentialInstanceConfig", type=InstanceTemplateConfidentialInstanceConfig.class, parameters={})
    private Output<InstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfig;

    /**
     * @return Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    public Output<InstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfig() {
        return this.confidentialInstanceConfig;
    }
    /**
     * A brief description of this resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A brief description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     * 
     */
    @Export(name="disks", type=List.class, parameters={InstanceTemplateDisk.class})
    private Output<List<InstanceTemplateDisk>> disks;

    /**
     * @return Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     * 
     */
    public Output<List<InstanceTemplateDisk>> disks() {
        return this.disks;
    }
    /**
     * ) Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
     * 
     */
    @Export(name="enableDisplay", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableDisplay;

    /**
     * @return ) Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
     * 
     */
    public Output<Optional<Boolean>> enableDisplay() {
        return Codegen.optional(this.enableDisplay);
    }
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * 
     */
    @Export(name="guestAccelerators", type=List.class, parameters={InstanceTemplateGuestAccelerator.class})
    private Output</* @Nullable */ List<InstanceTemplateGuestAccelerator>> guestAccelerators;

    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * 
     */
    public Output<Optional<List<InstanceTemplateGuestAccelerator>>> guestAccelerators() {
        return Codegen.optional(this.guestAccelerators);
    }
    /**
     * A brief description to use for instances
     * created from this template.
     * 
     */
    @Export(name="instanceDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceDescription;

    /**
     * @return A brief description to use for instances
     * created from this template.
     * 
     */
    public Output<Optional<String>> instanceDescription() {
        return Codegen.optional(this.instanceDescription);
    }
    /**
     * A set of ket/value label pairs to assign to disk created from
     * this template
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of ket/value label pairs to assign to disk created from
     * this template
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
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
     * Metadata key/value pairs to make available from
     * within instances created from this template.
     * 
     */
    @Export(name="metadata", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> metadata;

    /**
     * @return Metadata key/value pairs to make available from
     * within instances created from this template.
     * 
     */
    public Output<Optional<Map<String,Object>>> metadata() {
        return Codegen.optional(this.metadata);
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
     * An alternative to using the
     * startup-script metadata key, mostly to match the compute_instance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     * 
     */
    @Export(name="metadataStartupScript", type=String.class, parameters={})
    private Output</* @Nullable */ String> metadataStartupScript;

    /**
     * @return An alternative to using the
     * startup-script metadata key, mostly to match the compute_instance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     * 
     */
    public Output<Optional<String>> metadataStartupScript() {
        return Codegen.optional(this.metadataStartupScript);
    }
    /**
     * Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * 
     */
    @Export(name="minCpuPlatform", type=String.class, parameters={})
    private Output</* @Nullable */ String> minCpuPlatform;

    /**
     * @return Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * 
     */
    public Output<Optional<String>> minCpuPlatform() {
        return Codegen.optional(this.minCpuPlatform);
    }
    /**
     * The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    @Export(name="namePrefix", type=String.class, parameters={})
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * Networks to attach to instances created from
     * this template. This can be specified multiple times for multiple networks.
     * Structure is documented below.
     * 
     */
    @Export(name="networkInterfaces", type=List.class, parameters={InstanceTemplateNetworkInterface.class})
    private Output</* @Nullable */ List<InstanceTemplateNetworkInterface>> networkInterfaces;

    /**
     * @return Networks to attach to instances created from
     * this template. This can be specified multiple times for multiple networks.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceTemplateNetworkInterface>>> networkInterfaces() {
        return Codegen.optional(this.networkInterfaces);
    }
    /**
     * Configures network performance settings for the instance created from the
     * template. Structure is documented below. **Note**: `machine_type`
     * must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
     * the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
     * in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
     * in order for this setting to take effect.
     * 
     */
    @Export(name="networkPerformanceConfig", type=InstanceTemplateNetworkPerformanceConfig.class, parameters={})
    private Output</* @Nullable */ InstanceTemplateNetworkPerformanceConfig> networkPerformanceConfig;

    /**
     * @return Configures network performance settings for the instance created from the
     * template. Structure is documented below. **Note**: `machine_type`
     * must be a [supported type](https://cloud.google.com/compute/docs/networking/configure-vm-with-high-bandwidth-configuration),
     * the `image` used must include the [`GVNIC`](https://cloud.google.com/compute/docs/networking/using-gvnic#create-instance-gvnic-image)
     * in `guest-os-features`, and `network_interface.0.nic-type` must be `GVNIC`
     * in order for this setting to take effect.
     * 
     */
    public Output<Optional<InstanceTemplateNetworkPerformanceConfig>> networkPerformanceConfig() {
        return Codegen.optional(this.networkPerformanceConfig);
    }
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Specifies the reservations that this instance can consume from.
     * Structure is documented below.
     * 
     */
    @Export(name="reservationAffinity", type=InstanceTemplateReservationAffinity.class, parameters={})
    private Output</* @Nullable */ InstanceTemplateReservationAffinity> reservationAffinity;

    /**
     * @return Specifies the reservations that this instance can consume from.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceTemplateReservationAffinity>> reservationAffinity() {
        return Codegen.optional(this.reservationAffinity);
    }
    /**
     * The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    @Export(name="scheduling", type=InstanceTemplateScheduling.class, parameters={})
    private Output<InstanceTemplateScheduling> scheduling;

    /**
     * @return The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    public Output<InstanceTemplateScheduling> scheduling() {
        return this.scheduling;
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
     * Service account to attach to the instance. Structure is documented below.
     * 
     */
    @Export(name="serviceAccount", type=InstanceTemplateServiceAccount.class, parameters={})
    private Output</* @Nullable */ InstanceTemplateServiceAccount> serviceAccount;

    /**
     * @return Service account to attach to the instance. Structure is documented below.
     * 
     */
    public Output<Optional<InstanceTemplateServiceAccount>> serviceAccount() {
        return Codegen.optional(this.serviceAccount);
    }
    /**
     * Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * 
     */
    @Export(name="shieldedInstanceConfig", type=InstanceTemplateShieldedInstanceConfig.class, parameters={})
    private Output<InstanceTemplateShieldedInstanceConfig> shieldedInstanceConfig;

    /**
     * @return Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * 
     */
    public Output<InstanceTemplateShieldedInstanceConfig> shieldedInstanceConfig() {
        return this.shieldedInstanceConfig;
    }
    /**
     * Tags to attach to the instance.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags to attach to the instance.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceTemplate(String name) {
        this(name, InstanceTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceTemplate(String name, InstanceTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceTemplate(String name, InstanceTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceTemplate:InstanceTemplate", name, args == null ? InstanceTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceTemplate(String name, Output<String> id, @Nullable InstanceTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceTemplate:InstanceTemplate", name, state, makeResourceOptions(options, id));
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
    public static InstanceTemplate get(String name, Output<String> id, @Nullable InstanceTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceTemplate(name, id, state, options);
    }
}
