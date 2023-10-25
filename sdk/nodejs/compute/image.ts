// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

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
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.compute.Image("example", {rawDisk: {
 *     source: "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
 * }});
 * ```
 * ### Image Guest Os
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.compute.Image("example", {
 *     guestOsFeatures: [
 *         {
 *             type: "SECURE_BOOT",
 *         },
 *         {
 *             type: "MULTI_IP_SUBNET",
 *         },
 *     ],
 *     rawDisk: {
 *         source: "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
 *     },
 * });
 * ```
 * ### Image Basic Storage Location
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.compute.Image("example", {
 *     rawDisk: {
 *         source: "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
 *     },
 *     storageLocations: ["us-central1"],
 * });
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
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in
     * bytes).
     */
    public /*out*/ readonly archiveSizeBytes!: pulumi.Output<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    public readonly diskSizeGb!: pulumi.Output<number>;
    /**
     * The name of the image family to which this image belongs. You can
     * create disks by specifying an image family instead of a specific
     * image name. The image family always returns its latest image that is
     * not deprecated. The name of the image family must comply with
     * RFC1035.
     */
    public readonly family!: pulumi.Output<string | undefined>;
    /**
     * A list of features to enable on the guest operating system.
     * Applicable only for bootable images.
     * Structure is documented below.
     */
    public readonly guestOsFeatures!: pulumi.Output<outputs.compute.ImageGuestOsFeature[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key.
     * After you encrypt an image with a customer-supplied key, you must
     * provide the same key if you use the image later (e.g. to create a
     * disk from the image)
     * Structure is documented below.
     */
    public readonly imageEncryptionKey!: pulumi.Output<outputs.compute.ImageImageEncryptionKey | undefined>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used
     * internally during updates.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this Image.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Any applicable license URI.
     */
    public readonly licenses!: pulumi.Output<string[]>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The parameters of the raw disk image.
     * Structure is documented below.
     */
    public readonly rawDisk!: pulumi.Output<outputs.compute.ImageRawDisk | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The source disk to create this image based on.
     * You must provide either this property or the
     * rawDisk.source property but not both to create an image.
     */
    public readonly sourceDisk!: pulumi.Output<string | undefined>;
    /**
     * URL of the source image used to create this image. In order to create an image, you must provide the full or partial
     * URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    public readonly sourceImage!: pulumi.Output<string | undefined>;
    /**
     * URL of the source snapshot used to create this image.
     * In order to create an image, you must provide the full or partial URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The sourceImage URL
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    public readonly sourceSnapshot!: pulumi.Output<string | undefined>;
    /**
     * Cloud Storage bucket storage location of the image
     * (regional or multi-regional).
     * Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
     */
    public readonly storageLocations!: pulumi.Output<string[]>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageState | undefined;
            resourceInputs["archiveSizeBytes"] = state ? state.archiveSizeBytes : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["diskSizeGb"] = state ? state.diskSizeGb : undefined;
            resourceInputs["family"] = state ? state.family : undefined;
            resourceInputs["guestOsFeatures"] = state ? state.guestOsFeatures : undefined;
            resourceInputs["imageEncryptionKey"] = state ? state.imageEncryptionKey : undefined;
            resourceInputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["licenses"] = state ? state.licenses : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["rawDisk"] = state ? state.rawDisk : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["sourceDisk"] = state ? state.sourceDisk : undefined;
            resourceInputs["sourceImage"] = state ? state.sourceImage : undefined;
            resourceInputs["sourceSnapshot"] = state ? state.sourceSnapshot : undefined;
            resourceInputs["storageLocations"] = state ? state.storageLocations : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["diskSizeGb"] = args ? args.diskSizeGb : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["guestOsFeatures"] = args ? args.guestOsFeatures : undefined;
            resourceInputs["imageEncryptionKey"] = args ? args.imageEncryptionKey : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["licenses"] = args ? args.licenses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rawDisk"] = args ? args.rawDisk : undefined;
            resourceInputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            resourceInputs["sourceImage"] = args ? args.sourceImage : undefined;
            resourceInputs["sourceSnapshot"] = args ? args.sourceSnapshot : undefined;
            resourceInputs["storageLocations"] = args ? args.storageLocations : undefined;
            resourceInputs["archiveSizeBytes"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * Size of the image tar.gz archive stored in Google Cloud Storage (in
     * bytes).
     */
    archiveSizeBytes?: pulumi.Input<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    diskSizeGb?: pulumi.Input<number>;
    /**
     * The name of the image family to which this image belongs. You can
     * create disks by specifying an image family instead of a specific
     * image name. The image family always returns its latest image that is
     * not deprecated. The name of the image family must comply with
     * RFC1035.
     */
    family?: pulumi.Input<string>;
    /**
     * A list of features to enable on the guest operating system.
     * Applicable only for bootable images.
     * Structure is documented below.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.ImageGuestOsFeature>[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key.
     * After you encrypt an image with a customer-supplied key, you must
     * provide the same key if you use the image later (e.g. to create a
     * disk from the image)
     * Structure is documented below.
     */
    imageEncryptionKey?: pulumi.Input<inputs.compute.ImageImageEncryptionKey>;
    /**
     * The fingerprint used for optimistic locking of this resource. Used
     * internally during updates.
     */
    labelFingerprint?: pulumi.Input<string>;
    /**
     * Labels to apply to this Image.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Any applicable license URI.
     */
    licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The parameters of the raw disk image.
     * Structure is documented below.
     */
    rawDisk?: pulumi.Input<inputs.compute.ImageRawDisk>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The source disk to create this image based on.
     * You must provide either this property or the
     * rawDisk.source property but not both to create an image.
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * URL of the source image used to create this image. In order to create an image, you must provide the full or partial
     * URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * URL of the source snapshot used to create this image.
     * In order to create an image, you must provide the full or partial URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The sourceImage URL
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * Cloud Storage bucket storage location of the image
     * (regional or multi-regional).
     * Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * An optional description of this resource. Provide this property when
     * you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Size of the image when restored onto a persistent disk (in GB).
     */
    diskSizeGb?: pulumi.Input<number>;
    /**
     * The name of the image family to which this image belongs. You can
     * create disks by specifying an image family instead of a specific
     * image name. The image family always returns its latest image that is
     * not deprecated. The name of the image family must comply with
     * RFC1035.
     */
    family?: pulumi.Input<string>;
    /**
     * A list of features to enable on the guest operating system.
     * Applicable only for bootable images.
     * Structure is documented below.
     */
    guestOsFeatures?: pulumi.Input<pulumi.Input<inputs.compute.ImageGuestOsFeature>[]>;
    /**
     * Encrypts the image using a customer-supplied encryption key.
     * After you encrypt an image with a customer-supplied key, you must
     * provide the same key if you use the image later (e.g. to create a
     * disk from the image)
     * Structure is documented below.
     */
    imageEncryptionKey?: pulumi.Input<inputs.compute.ImageImageEncryptionKey>;
    /**
     * Labels to apply to this Image.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Any applicable license URI.
     */
    licenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the resource; provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     *
     *
     * - - -
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The parameters of the raw disk image.
     * Structure is documented below.
     */
    rawDisk?: pulumi.Input<inputs.compute.ImageRawDisk>;
    /**
     * The source disk to create this image based on.
     * You must provide either this property or the
     * rawDisk.source property but not both to create an image.
     */
    sourceDisk?: pulumi.Input<string>;
    /**
     * URL of the source image used to create this image. In order to create an image, you must provide the full or partial
     * URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * URL of the source snapshot used to create this image.
     * In order to create an image, you must provide the full or partial URL of one of the following:
     * * The selfLink URL
     * * This property
     * * The sourceImage URL
     * * The rawDisk.source URL
     * * The sourceDisk URL
     */
    sourceSnapshot?: pulumi.Input<string>;
    /**
     * Cloud Storage bucket storage location of the image
     * (regional or multi-regional).
     * Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
     */
    storageLocations?: pulumi.Input<pulumi.Input<string>[]>;
}
