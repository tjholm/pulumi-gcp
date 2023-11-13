// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.ImageArgs;
import com.pulumi.gcp.compute.inputs.ImageState;
import com.pulumi.gcp.compute.outputs.ImageGuestOsFeature;
import com.pulumi.gcp.compute.outputs.ImageImageEncryptionKey;
import com.pulumi.gcp.compute.outputs.ImageRawDisk;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents an Image resource.
 * 
 * Google Compute Engine uses operating system images to create the root
 * persistent disks for your instances. You specify an image when you create
 * an instance. Images contain a boot loader, an operating system, and a
 * root file system. Linux operating system images are also capable of
 * running containers on Compute Engine.
 * 
 * Images can be either public or custom.
 * 
 * Public images are provided and maintained by Google, open-source
 * communities, and third-party vendors. By default, all projects have
 * access to these images and can use them to create instances.  Custom
 * images are available only to your project. You can create a custom image
 * from root persistent disks and other images. Then, use the custom image
 * to create an instance.
 * 
 * To get more information about Image, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/images)
 * 
 * ## Example Usage
 * ### Image Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Image;
 * import com.pulumi.gcp.compute.ImageArgs;
 * import com.pulumi.gcp.compute.inputs.ImageRawDiskArgs;
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
 *         var example = new Image(&#34;example&#34;, ImageArgs.builder()        
 *             .rawDisk(ImageRawDiskArgs.builder()
 *                 .source(&#34;https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Image Guest Os
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Image;
 * import com.pulumi.gcp.compute.ImageArgs;
 * import com.pulumi.gcp.compute.inputs.ImageGuestOsFeatureArgs;
 * import com.pulumi.gcp.compute.inputs.ImageRawDiskArgs;
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
 *         var example = new Image(&#34;example&#34;, ImageArgs.builder()        
 *             .guestOsFeatures(            
 *                 ImageGuestOsFeatureArgs.builder()
 *                     .type(&#34;SECURE_BOOT&#34;)
 *                     .build(),
 *                 ImageGuestOsFeatureArgs.builder()
 *                     .type(&#34;MULTI_IP_SUBNET&#34;)
 *                     .build())
 *             .rawDisk(ImageRawDiskArgs.builder()
 *                 .source(&#34;https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Image Basic Storage Location
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Image;
 * import com.pulumi.gcp.compute.ImageArgs;
 * import com.pulumi.gcp.compute.inputs.ImageRawDiskArgs;
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
 *         var example = new Image(&#34;example&#34;, ImageArgs.builder()        
 *             .rawDisk(ImageRawDiskArgs.builder()
 *                 .source(&#34;https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz&#34;)
 *                 .build())
 *             .storageLocations(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Image can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/image:Image default projects/{{project}}/global/images/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/image:Image default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/image:Image default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in
     * bytes).
     * 
     */
    @Export(name="archiveSizeBytes", refs={Integer.class}, tree="[0]")
    private Output<Integer> archiveSizeBytes;

    /**
     * @return Size of the image tar.gz archive stored in Google Cloud Storage (in
     * bytes).
     * 
     */
    public Output<Integer> archiveSizeBytes() {
        return this.archiveSizeBytes;
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
     * An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource. Provide this property when
     * you create the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     * 
     */
    @Export(name="diskSizeGb", refs={Integer.class}, tree="[0]")
    private Output<Integer> diskSizeGb;

    /**
     * @return Size of the image when restored onto a persistent disk (in GB).
     * 
     */
    public Output<Integer> diskSizeGb() {
        return this.diskSizeGb;
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
     * The name of the image family to which this image belongs. You can
     * create disks by specifying an image family instead of a specific
     * image name. The image family always returns its latest image that is
     * not deprecated. The name of the image family must comply with
     * RFC1035.
     * 
     */
    @Export(name="family", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> family;

    /**
     * @return The name of the image family to which this image belongs. You can
     * create disks by specifying an image family instead of a specific
     * image name. The image family always returns its latest image that is
     * not deprecated. The name of the image family must comply with
     * RFC1035.
     * 
     */
    public Output<Optional<String>> family() {
        return Codegen.optional(this.family);
    }
    /**
     * A list of features to enable on the guest operating system.
     * Applicable only for bootable images.
     * Structure is documented below.
     * 
     */
    @Export(name="guestOsFeatures", refs={List.class,ImageGuestOsFeature.class}, tree="[0,1]")
    private Output<List<ImageGuestOsFeature>> guestOsFeatures;

    /**
     * @return A list of features to enable on the guest operating system.
     * Applicable only for bootable images.
     * Structure is documented below.
     * 
     */
    public Output<List<ImageGuestOsFeature>> guestOsFeatures() {
        return this.guestOsFeatures;
    }
    /**
     * Encrypts the image using a customer-supplied encryption key.
     * After you encrypt an image with a customer-supplied key, you must
     * provide the same key if you use the image later (e.g. to create a
     * disk from the image)
     * Structure is documented below.
     * 
     */
    @Export(name="imageEncryptionKey", refs={ImageImageEncryptionKey.class}, tree="[0]")
    private Output</* @Nullable */ ImageImageEncryptionKey> imageEncryptionKey;

    /**
     * @return Encrypts the image using a customer-supplied encryption key.
     * After you encrypt an image with a customer-supplied key, you must
     * provide the same key if you use the image later (e.g. to create a
     * disk from the image)
     * Structure is documented below.
     * 
     */
    public Output<Optional<ImageImageEncryptionKey>> imageEncryptionKey() {
        return Codegen.optional(this.imageEncryptionKey);
    }
    /**
     * The fingerprint used for optimistic locking of this resource. Used
     * internally during updates.
     * 
     */
    @Export(name="labelFingerprint", refs={String.class}, tree="[0]")
    private Output<String> labelFingerprint;

    /**
     * @return The fingerprint used for optimistic locking of this resource. Used
     * internally during updates.
     * 
     */
    public Output<String> labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * Labels to apply to this Image.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels to apply to this Image.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Any applicable license URI.
     * 
     */
    @Export(name="licenses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> licenses;

    /**
     * @return Any applicable license URI.
     * 
     */
    public Output<List<String>> licenses() {
        return this.licenses;
    }
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     * 
     * ***
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
     * The parameters of the raw disk image.
     * Structure is documented below.
     * 
     */
    @Export(name="rawDisk", refs={ImageRawDisk.class}, tree="[0]")
    private Output</* @Nullable */ ImageRawDisk> rawDisk;

    /**
     * @return The parameters of the raw disk image.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ImageRawDisk>> rawDisk() {
        return Codegen.optional(this.rawDisk);
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
     * The source disk to create this image based on.
     * You must provide either this property or the
     * rawDisk.source property but not both to create an image.
     * 
     */
    @Export(name="sourceDisk", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceDisk;

    /**
     * @return The source disk to create this image based on.
     * You must provide either this property or the
     * rawDisk.source property but not both to create an image.
     * 
     */
    public Output<Optional<String>> sourceDisk() {
        return Codegen.optional(this.sourceDisk);
    }
    /**
     * URL of the source image used to create this image. In order to create an image, you must provide the full or partial
     * URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The rawDisk.source URL
     * * The sourceDisk URL
     * 
     */
    @Export(name="sourceImage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceImage;

    /**
     * @return URL of the source image used to create this image. In order to create an image, you must provide the full or partial
     * URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The rawDisk.source URL
     * * The sourceDisk URL
     * 
     */
    public Output<Optional<String>> sourceImage() {
        return Codegen.optional(this.sourceImage);
    }
    /**
     * URL of the source snapshot used to create this image.
     * In order to create an image, you must provide the full or partial URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The sourceImage URL
     * * The rawDisk.source URL
     * * The sourceDisk URL
     * 
     */
    @Export(name="sourceSnapshot", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceSnapshot;

    /**
     * @return URL of the source snapshot used to create this image.
     * In order to create an image, you must provide the full or partial URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The sourceImage URL
     * * The rawDisk.source URL
     * * The sourceDisk URL
     * 
     */
    public Output<Optional<String>> sourceSnapshot() {
        return Codegen.optional(this.sourceSnapshot);
    }
    /**
     * Cloud Storage bucket storage location of the image
     * (regional or multi-regional).
     * Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
     * 
     */
    @Export(name="storageLocations", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> storageLocations;

    /**
     * @return Cloud Storage bucket storage location of the image
     * (regional or multi-regional).
     * Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
     * 
     */
    public Output<List<String>> storageLocations() {
        return this.storageLocations;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Image(String name) {
        this(name, ImageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Image(String name, @Nullable ImageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Image(String name, @Nullable ImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/image:Image", name, args == null ? ImageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Image(String name, Output<String> id, @Nullable ImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/image:Image", name, state, makeResourceOptions(options, id));
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
    public static Image get(String name, Output<String> id, @Nullable ImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Image(name, id, state, options);
    }
}
