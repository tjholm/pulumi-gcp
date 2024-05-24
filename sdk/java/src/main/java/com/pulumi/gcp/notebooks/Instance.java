// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.notebooks.InstanceArgs;
import com.pulumi.gcp.notebooks.inputs.InstanceState;
import com.pulumi.gcp.notebooks.outputs.InstanceAcceleratorConfig;
import com.pulumi.gcp.notebooks.outputs.InstanceContainerImage;
import com.pulumi.gcp.notebooks.outputs.InstanceReservationAffinity;
import com.pulumi.gcp.notebooks.outputs.InstanceShieldedInstanceConfig;
import com.pulumi.gcp.notebooks.outputs.InstanceVmImage;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **Warning:** `google_notebook_instance` is deprecated and will be removed in a future major release. Use `gcp.workbench.Instance` instead.
 * 
 * A Cloud AI Platform Notebook instance.
 * 
 * &gt; **Note:** Due to limitations of the Notebooks Instance API, many fields
 * in this resource do not properly detect drift. These fields will also not
 * appear in state once imported.
 * 
 * To get more information about Instance, see:
 * 
 * * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
 * 
 * ## Example Usage
 * 
 * ### Notebook Instance Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.notebooks.Instance;
 * import com.pulumi.gcp.notebooks.InstanceArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceVmImageArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("notebooks-instance")
 *             .location("us-west1-a")
 *             .machineType("e2-medium")
 *             .vmImage(InstanceVmImageArgs.builder()
 *                 .project("deeplearning-platform-release")
 *                 .imageFamily("tf-latest-cpu")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Notebook Instance Basic Stopped
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.notebooks.Instance;
 * import com.pulumi.gcp.notebooks.InstanceArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceVmImageArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("notebooks-instance")
 *             .location("us-west1-a")
 *             .machineType("e2-medium")
 *             .vmImage(InstanceVmImageArgs.builder()
 *                 .project("deeplearning-platform-release")
 *                 .imageFamily("tf-latest-cpu")
 *                 .build())
 *             .desiredState("STOPPED")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Notebook Instance Basic Container
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.notebooks.Instance;
 * import com.pulumi.gcp.notebooks.InstanceArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceContainerImageArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("notebooks-instance")
 *             .location("us-west1-a")
 *             .machineType("e2-medium")
 *             .metadata(Map.of("proxy-mode", "service_account"))
 *             .containerImage(InstanceContainerImageArgs.builder()
 *                 .repository("gcr.io/deeplearning-platform-release/base-cpu")
 *                 .tag("latest")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Notebook Instance Basic Gpu
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.notebooks.Instance;
 * import com.pulumi.gcp.notebooks.InstanceArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceAcceleratorConfigArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceVmImageArgs;
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
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("notebooks-instance")
 *             .location("us-west1-a")
 *             .machineType("n1-standard-1")
 *             .installGpuDriver(true)
 *             .acceleratorConfig(InstanceAcceleratorConfigArgs.builder()
 *                 .type("NVIDIA_TESLA_T4")
 *                 .coreCount(1)
 *                 .build())
 *             .vmImage(InstanceVmImageArgs.builder()
 *                 .project("deeplearning-platform-release")
 *                 .imageFamily("tf-latest-gpu")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Notebook Instance Full
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
 * import com.pulumi.gcp.compute.inputs.GetNetworkArgs;
 * import com.pulumi.gcp.compute.inputs.GetSubnetworkArgs;
 * import com.pulumi.gcp.notebooks.Instance;
 * import com.pulumi.gcp.notebooks.InstanceArgs;
 * import com.pulumi.gcp.notebooks.inputs.InstanceVmImageArgs;
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
 *         final var myNetwork = ComputeFunctions.getNetwork(GetNetworkArgs.builder()
 *             .name("default")
 *             .build());
 * 
 *         final var mySubnetwork = ComputeFunctions.getSubnetwork(GetSubnetworkArgs.builder()
 *             .name("default")
 *             .region("us-central1")
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .name("notebooks-instance")
 *             .location("us-central1-a")
 *             .machineType("e2-medium")
 *             .vmImage(InstanceVmImageArgs.builder()
 *                 .project("deeplearning-platform-release")
 *                 .imageFamily("tf-latest-cpu")
 *                 .build())
 *             .instanceOwners("my{@literal @}service-account.com")
 *             .serviceAccount("my{@literal @}service-account.com")
 *             .installGpuDriver(true)
 *             .bootDiskType("PD_SSD")
 *             .bootDiskSizeGb(110)
 *             .noPublicIp(true)
 *             .noProxyAccess(true)
 *             .network(myNetwork.applyValue(getNetworkResult -> getNetworkResult.id()))
 *             .subnet(mySubnetwork.applyValue(getSubnetworkResult -> getSubnetworkResult.id()))
 *             .labels(Map.of("k", "val"))
 *             .metadata(Map.of("terraform", "true"))
 *             .serviceAccountScopes(            
 *                 "https://www.googleapis.com/auth/bigquery",
 *                 "https://www.googleapis.com/auth/devstorage.read_write",
 *                 "https://www.googleapis.com/auth/cloud-platform",
 *                 "https://www.googleapis.com/auth/userinfo.email")
 *             .tags(            
 *                 "foo",
 *                 "bar")
 *             .diskEncryption("CMEK")
 *             .kmsKey("my-crypto-key")
 *             .desiredState("ACTIVE")
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
 * Instance can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/instances/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, Instance can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:notebooks/instance:Instance default projects/{{project}}/locations/{{location}}/instances/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:notebooks/instance:Instance default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:notebooks/instance:Instance default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:notebooks/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The hardware accelerator used on this instance. If you use accelerators,
     * make sure that your configuration has enough vCPUs and memory to support the
     * machineType you have selected.
     * Structure is documented below.
     * 
     */
    @Export(name="acceleratorConfig", refs={InstanceAcceleratorConfig.class}, tree="[0]")
    private Output</* @Nullable */ InstanceAcceleratorConfig> acceleratorConfig;

    /**
     * @return The hardware accelerator used on this instance. If you use accelerators,
     * make sure that your configuration has enough vCPUs and memory to support the
     * machineType you have selected.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceAcceleratorConfig>> acceleratorConfig() {
        return Codegen.optional(this.acceleratorConfig);
    }
    /**
     * The size of the boot disk in GB attached to this instance,
     * up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
     * If not specified, this defaults to 100.
     * 
     */
    @Export(name="bootDiskSizeGb", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> bootDiskSizeGb;

    /**
     * @return The size of the boot disk in GB attached to this instance,
     * up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
     * If not specified, this defaults to 100.
     * 
     */
    public Output<Optional<Integer>> bootDiskSizeGb() {
        return Codegen.optional(this.bootDiskSizeGb);
    }
    /**
     * Possible disk types for notebook instances.
     * Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
     * 
     */
    @Export(name="bootDiskType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bootDiskType;

    /**
     * @return Possible disk types for notebook instances.
     * Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
     * 
     */
    public Output<Optional<String>> bootDiskType() {
        return Codegen.optional(this.bootDiskType);
    }
    /**
     * Use a container image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    @Export(name="containerImage", refs={InstanceContainerImage.class}, tree="[0]")
    private Output</* @Nullable */ InstanceContainerImage> containerImage;

    /**
     * @return Use a container image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceContainerImage>> containerImage() {
        return Codegen.optional(this.containerImage);
    }
    /**
     * Instance creation time
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Instance creation time
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Specify a custom Cloud Storage path where the GPU driver is stored.
     * If not specified, we&#39;ll automatically choose from official GPU drivers.
     * 
     */
    @Export(name="customGpuDriverPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customGpuDriverPath;

    /**
     * @return Specify a custom Cloud Storage path where the GPU driver is stored.
     * If not specified, we&#39;ll automatically choose from official GPU drivers.
     * 
     */
    public Output<Optional<String>> customGpuDriverPath() {
        return Codegen.optional(this.customGpuDriverPath);
    }
    /**
     * The size of the data disk in GB attached to this instance,
     * up to a maximum of 64000 GB (64 TB).
     * You can choose the size of the data disk based on how big your notebooks and data are.
     * If not specified, this defaults to 100.
     * 
     */
    @Export(name="dataDiskSizeGb", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> dataDiskSizeGb;

    /**
     * @return The size of the data disk in GB attached to this instance,
     * up to a maximum of 64000 GB (64 TB).
     * You can choose the size of the data disk based on how big your notebooks and data are.
     * If not specified, this defaults to 100.
     * 
     */
    public Output<Optional<Integer>> dataDiskSizeGb() {
        return Codegen.optional(this.dataDiskSizeGb);
    }
    /**
     * Possible disk types for notebook instances.
     * Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
     * 
     */
    @Export(name="dataDiskType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dataDiskType;

    /**
     * @return Possible disk types for notebook instances.
     * Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
     * 
     */
    public Output<Optional<String>> dataDiskType() {
        return Codegen.optional(this.dataDiskType);
    }
    /**
     * Desired state of the Notebook Instance. Set this field to `ACTIVE` to start the Instance, and `STOPPED` to stop the Instance.
     * 
     */
    @Export(name="desiredState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desiredState;

    /**
     * @return Desired state of the Notebook Instance. Set this field to `ACTIVE` to start the Instance, and `STOPPED` to stop the Instance.
     * 
     */
    public Output<Optional<String>> desiredState() {
        return Codegen.optional(this.desiredState);
    }
    /**
     * Disk encryption method used on the boot and data disks, defaults to GMEK.
     * Possible values are: `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, `CMEK`.
     * 
     */
    @Export(name="diskEncryption", refs={String.class}, tree="[0]")
    private Output<String> diskEncryption;

    /**
     * @return Disk encryption method used on the boot and data disks, defaults to GMEK.
     * Possible values are: `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, `CMEK`.
     * 
     */
    public Output<String> diskEncryption() {
        return this.diskEncryption;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Whether the end user authorizes Google Cloud to install GPU driver
     * on this instance. If this field is empty or set to false, the GPU driver
     * won&#39;t be installed. Only applicable to instances with GPUs.
     * 
     */
    @Export(name="installGpuDriver", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> installGpuDriver;

    /**
     * @return Whether the end user authorizes Google Cloud to install GPU driver
     * on this instance. If this field is empty or set to false, the GPU driver
     * won&#39;t be installed. Only applicable to instances with GPUs.
     * 
     */
    public Output<Optional<Boolean>> installGpuDriver() {
        return Codegen.optional(this.installGpuDriver);
    }
    /**
     * The list of owners of this instance after creation.
     * Format: alias{@literal @}example.com.
     * Currently supports one owner only.
     * If not specified, all of the service account users of
     * your VM instance&#39;s service account can use the instance.
     * 
     */
    @Export(name="instanceOwners", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> instanceOwners;

    /**
     * @return The list of owners of this instance after creation.
     * Format: alias{@literal @}example.com.
     * Currently supports one owner only.
     * If not specified, all of the service account users of
     * your VM instance&#39;s service account can use the instance.
     * 
     */
    public Output<Optional<List<String>>> instanceOwners() {
        return Codegen.optional(this.instanceOwners);
    }
    /**
     * The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
     * Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
     * 
     */
    @Export(name="kmsKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKey;

    /**
     * @return The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
     * Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
     * 
     */
    public Output<Optional<String>> kmsKey() {
        return Codegen.optional(this.kmsKey);
    }
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels to apply to this instance. These can be later modified by the setLabels method.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * A reference to the zone where the machine resides.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return A reference to the zone where the machine resides.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * A reference to a machine type which defines VM kind.
     * 
     */
    @Export(name="machineType", refs={String.class}, tree="[0]")
    private Output<String> machineType;

    /**
     * @return A reference to a machine type which defines VM kind.
     * 
     */
    public Output<String> machineType() {
        return this.machineType;
    }
    /**
     * Custom metadata to apply to this instance.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> metadata;

    /**
     * @return Custom metadata to apply to this instance.
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    public Output<Optional<Map<String,String>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * The name specified for the Notebook instance.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name specified for the Notebook instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the VPC that this instance is in.
     * Format: projects/{project_id}/global/networks/{network_id}
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The name of the VPC that this instance is in.
     * Format: projects/{project_id}/global/networks/{network_id}
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The type of vNIC driver.
     * Possible values are: `UNSPECIFIED_NIC_TYPE`, `VIRTIO_NET`, `GVNIC`.
     * 
     */
    @Export(name="nicType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nicType;

    /**
     * @return The type of vNIC driver.
     * Possible values are: `UNSPECIFIED_NIC_TYPE`, `VIRTIO_NET`, `GVNIC`.
     * 
     */
    public Output<Optional<String>> nicType() {
        return Codegen.optional(this.nicType);
    }
    /**
     * The notebook instance will not register with the proxy..
     * 
     */
    @Export(name="noProxyAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noProxyAccess;

    /**
     * @return The notebook instance will not register with the proxy..
     * 
     */
    public Output<Optional<Boolean>> noProxyAccess() {
        return Codegen.optional(this.noProxyAccess);
    }
    /**
     * No public IP will be assigned to this instance.
     * 
     */
    @Export(name="noPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noPublicIp;

    /**
     * @return No public IP will be assigned to this instance.
     * 
     */
    public Output<Optional<Boolean>> noPublicIp() {
        return Codegen.optional(this.noPublicIp);
    }
    /**
     * If true, the data disk will not be auto deleted when deleting the instance.
     * 
     */
    @Export(name="noRemoveDataDisk", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noRemoveDataDisk;

    /**
     * @return If true, the data disk will not be auto deleted when deleting the instance.
     * 
     */
    public Output<Optional<Boolean>> noRemoveDataDisk() {
        return Codegen.optional(this.noRemoveDataDisk);
    }
    /**
     * Path to a Bash script that automatically runs after a
     * notebook instance fully boots up. The path must be a URL
     * or Cloud Storage path (gs://path-to-file/file-name).
     * 
     */
    @Export(name="postStartupScript", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postStartupScript;

    /**
     * @return Path to a Bash script that automatically runs after a
     * notebook instance fully boots up. The path must be a URL
     * or Cloud Storage path (gs://path-to-file/file-name).
     * 
     */
    public Output<Optional<String>> postStartupScript() {
        return Codegen.optional(this.postStartupScript);
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
     * The proxy endpoint that is used to access the Jupyter notebook.
     * Only returned when the resource is in a `PROVISIONED` state. If
     * needed you can utilize `pulumi up -refresh-only` to await
     * the population of this value.
     * 
     */
    @Export(name="proxyUri", refs={String.class}, tree="[0]")
    private Output<String> proxyUri;

    /**
     * @return The proxy endpoint that is used to access the Jupyter notebook.
     * Only returned when the resource is in a `PROVISIONED` state. If
     * needed you can utilize `pulumi up -refresh-only` to await
     * the population of this value.
     * 
     */
    public Output<String> proxyUri() {
        return this.proxyUri;
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Reservation Affinity for consuming Zonal reservation.
     * Structure is documented below.
     * 
     */
    @Export(name="reservationAffinity", refs={InstanceReservationAffinity.class}, tree="[0]")
    private Output</* @Nullable */ InstanceReservationAffinity> reservationAffinity;

    /**
     * @return Reservation Affinity for consuming Zonal reservation.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceReservationAffinity>> reservationAffinity() {
        return Codegen.optional(this.reservationAffinity);
    }
    /**
     * The service account on this instance, giving access to other
     * Google Cloud services. You can use any service account within
     * the same project, but you must have the service account user
     * permission to use the instance. If not specified,
     * the Compute Engine default service account is used.
     * 
     */
    @Export(name="serviceAccount", refs={String.class}, tree="[0]")
    private Output<String> serviceAccount;

    /**
     * @return The service account on this instance, giving access to other
     * Google Cloud services. You can use any service account within
     * the same project, but you must have the service account user
     * permission to use the instance. If not specified,
     * the Compute Engine default service account is used.
     * 
     */
    public Output<String> serviceAccount() {
        return this.serviceAccount;
    }
    /**
     * Optional. The URIs of service account scopes to be included in Compute Engine instances.
     * If not specified, the following scopes are defined:
     * - https://www.googleapis.com/auth/cloud-platform
     * - https://www.googleapis.com/auth/userinfo.email
     * 
     */
    @Export(name="serviceAccountScopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serviceAccountScopes;

    /**
     * @return Optional. The URIs of service account scopes to be included in Compute Engine instances.
     * If not specified, the following scopes are defined:
     * - https://www.googleapis.com/auth/cloud-platform
     * - https://www.googleapis.com/auth/userinfo.email
     * 
     */
    public Output<List<String>> serviceAccountScopes() {
        return this.serviceAccountScopes;
    }
    /**
     * A set of Shielded Instance options. Check [Images using supported Shielded VM features]
     * Not all combinations are valid
     * Structure is documented below.
     * 
     */
    @Export(name="shieldedInstanceConfig", refs={InstanceShieldedInstanceConfig.class}, tree="[0]")
    private Output<InstanceShieldedInstanceConfig> shieldedInstanceConfig;

    /**
     * @return A set of Shielded Instance options. Check [Images using supported Shielded VM features]
     * Not all combinations are valid
     * Structure is documented below.
     * 
     */
    public Output<InstanceShieldedInstanceConfig> shieldedInstanceConfig() {
        return this.shieldedInstanceConfig;
    }
    /**
     * The state of this instance.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The state of this instance.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The name of the subnet that this instance is in.
     * Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
     * 
     */
    @Export(name="subnet", refs={String.class}, tree="[0]")
    private Output<String> subnet;

    /**
     * @return The name of the subnet that this instance is in.
     * Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
     * 
     */
    public Output<String> subnet() {
        return this.subnet;
    }
    /**
     * The Compute Engine tags to add to instance.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return The Compute Engine tags to add to instance.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * Instance update time.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Instance update time.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Use a Compute Engine VM image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    @Export(name="vmImage", refs={InstanceVmImage.class}, tree="[0]")
    private Output</* @Nullable */ InstanceVmImage> vmImage;

    /**
     * @return Use a Compute Engine VM image to start the notebook instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<InstanceVmImage>> vmImage() {
        return Codegen.optional(this.vmImage);
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
        super("gcp:notebooks/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:notebooks/instance:Instance", name, state, makeResourceOptions(options, id));
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
