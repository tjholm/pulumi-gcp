// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents an Image resource.
    /// 
    /// Google Compute Engine uses operating system images to create the root
    /// persistent disks for your instances. You specify an image when you create
    /// an instance. Images contain a boot loader, an operating system, and a
    /// root file system. Linux operating system images are also capable of
    /// running containers on Compute Engine.
    /// 
    /// Images can be either public or custom.
    /// 
    /// Public images are provided and maintained by Google, open-source
    /// communities, and third-party vendors. By default, all projects have
    /// access to these images and can use them to create instances.  Custom
    /// images are available only to your project. You can create a custom image
    /// from root persistent disks and other images. Then, use the custom image
    /// to create an instance.
    /// 
    /// To get more information about Image, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/images)
    /// 
    /// ## Example Usage
    /// ### Image Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Gcp.Compute.Image("example", new()
    ///     {
    ///         RawDisk = new Gcp.Compute.Inputs.ImageRawDiskArgs
    ///         {
    ///             Source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Image Guest Os
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Gcp.Compute.Image("example", new()
    ///     {
    ///         GuestOsFeatures = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.ImageGuestOsFeatureArgs
    ///             {
    ///                 Type = "SECURE_BOOT",
    ///             },
    ///             new Gcp.Compute.Inputs.ImageGuestOsFeatureArgs
    ///             {
    ///                 Type = "MULTI_IP_SUBNET",
    ///             },
    ///         },
    ///         RawDisk = new Gcp.Compute.Inputs.ImageRawDiskArgs
    ///         {
    ///             Source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Image Basic Storage Location
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Gcp.Compute.Image("example", new()
    ///     {
    ///         RawDisk = new Gcp.Compute.Inputs.ImageRawDiskArgs
    ///         {
    ///             Source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz",
    ///         },
    ///         StorageLocations = new[]
    ///         {
    ///             "us-central1",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Image can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/image:Image default projects/{{project}}/global/images/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/image:Image default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/image:Image default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/image:Image")]
    public partial class Image : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in
        /// bytes).
        /// </summary>
        [Output("archiveSizeBytes")]
        public Output<int> ArchiveSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Output("diskSizeGb")]
        public Output<int> DiskSizeGb { get; private set; } = null!;

        /// <summary>
        /// The name of the image family to which this image belongs. You can
        /// create disks by specifying an image family instead of a specific
        /// image name. The image family always returns its latest image that is
        /// not deprecated. The name of the image family must comply with
        /// RFC1035.
        /// </summary>
        [Output("family")]
        public Output<string?> Family { get; private set; } = null!;

        /// <summary>
        /// A list of features to enable on the guest operating system.
        /// Applicable only for bootable images.
        /// Structure is documented below.
        /// </summary>
        [Output("guestOsFeatures")]
        public Output<ImmutableArray<Outputs.ImageGuestOsFeature>> GuestOsFeatures { get; private set; } = null!;

        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key.
        /// After you encrypt an image with a customer-supplied key, you must
        /// provide the same key if you use the image later (e.g. to create a
        /// disk from the image)
        /// Structure is documented below.
        /// </summary>
        [Output("imageEncryptionKey")]
        public Output<Outputs.ImageImageEncryptionKey?> ImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used
        /// internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        [Output("licenses")]
        public Output<ImmutableArray<string>> Licenses { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The parameters of the raw disk image.
        /// Structure is documented below.
        /// </summary>
        [Output("rawDisk")]
        public Output<Outputs.ImageRawDisk?> RawDisk { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The source disk to create this image based on.
        /// You must provide either this property or the
        /// rawDisk.source property but not both to create an image.
        /// </summary>
        [Output("sourceDisk")]
        public Output<string?> SourceDisk { get; private set; } = null!;

        /// <summary>
        /// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
        /// URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Output("sourceImage")]
        public Output<string?> SourceImage { get; private set; } = null!;

        /// <summary>
        /// URL of the source snapshot used to create this image.
        /// In order to create an image, you must provide the full or partial URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The sourceImage URL
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Output("sourceSnapshot")]
        public Output<string?> SourceSnapshot { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage bucket storage location of the image
        /// (regional or multi-regional).
        /// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
        /// </summary>
        [Output("storageLocations")]
        public Output<ImmutableArray<string>> StorageLocations { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/image:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/image:Image", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
        {
            return new Image(name, id, state, options);
        }
    }

    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// The name of the image family to which this image belongs. You can
        /// create disks by specifying an image family instead of a specific
        /// image name. The image family always returns its latest image that is
        /// not deprecated. The name of the image family must comply with
        /// RFC1035.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.ImageGuestOsFeatureArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system.
        /// Applicable only for bootable images.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ImageGuestOsFeatureArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.ImageGuestOsFeatureArgs>());
            set => _guestOsFeatures = value;
        }

        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key.
        /// After you encrypt an image with a customer-supplied key, you must
        /// provide the same key if you use the image later (e.g. to create a
        /// disk from the image)
        /// Structure is documented below.
        /// </summary>
        [Input("imageEncryptionKey")]
        public Input<Inputs.ImageImageEncryptionKeyArgs>? ImageEncryptionKey { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The parameters of the raw disk image.
        /// Structure is documented below.
        /// </summary>
        [Input("rawDisk")]
        public Input<Inputs.ImageRawDiskArgs>? RawDisk { get; set; }

        /// <summary>
        /// The source disk to create this image based on.
        /// You must provide either this property or the
        /// rawDisk.source property but not both to create an image.
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        /// <summary>
        /// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
        /// URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Input("sourceImage")]
        public Input<string>? SourceImage { get; set; }

        /// <summary>
        /// URL of the source snapshot used to create this image.
        /// In order to create an image, you must provide the full or partial URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The sourceImage URL
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Input("sourceSnapshot")]
        public Input<string>? SourceSnapshot { get; set; }

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// Cloud Storage bucket storage location of the image
        /// (regional or multi-regional).
        /// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public ImageArgs()
        {
        }
        public static new ImageArgs Empty => new ImageArgs();
    }

    public sealed class ImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the image tar.gz archive stored in Google Cloud Storage (in
        /// bytes).
        /// </summary>
        [Input("archiveSizeBytes")]
        public Input<int>? ArchiveSizeBytes { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Size of the image when restored onto a persistent disk (in GB).
        /// </summary>
        [Input("diskSizeGb")]
        public Input<int>? DiskSizeGb { get; set; }

        /// <summary>
        /// The name of the image family to which this image belongs. You can
        /// create disks by specifying an image family instead of a specific
        /// image name. The image family always returns its latest image that is
        /// not deprecated. The name of the image family must comply with
        /// RFC1035.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        [Input("guestOsFeatures")]
        private InputList<Inputs.ImageGuestOsFeatureGetArgs>? _guestOsFeatures;

        /// <summary>
        /// A list of features to enable on the guest operating system.
        /// Applicable only for bootable images.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ImageGuestOsFeatureGetArgs> GuestOsFeatures
        {
            get => _guestOsFeatures ?? (_guestOsFeatures = new InputList<Inputs.ImageGuestOsFeatureGetArgs>());
            set => _guestOsFeatures = value;
        }

        /// <summary>
        /// Encrypts the image using a customer-supplied encryption key.
        /// After you encrypt an image with a customer-supplied key, you must
        /// provide the same key if you use the image later (e.g. to create a
        /// disk from the image)
        /// Structure is documented below.
        /// </summary>
        [Input("imageEncryptionKey")]
        public Input<Inputs.ImageImageEncryptionKeyGetArgs>? ImageEncryptionKey { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used
        /// internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this Image.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("licenses")]
        private InputList<string>? _licenses;

        /// <summary>
        /// Any applicable license URI.
        /// </summary>
        public InputList<string> Licenses
        {
            get => _licenses ?? (_licenses = new InputList<string>());
            set => _licenses = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The parameters of the raw disk image.
        /// Structure is documented below.
        /// </summary>
        [Input("rawDisk")]
        public Input<Inputs.ImageRawDiskGetArgs>? RawDisk { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The source disk to create this image based on.
        /// You must provide either this property or the
        /// rawDisk.source property but not both to create an image.
        /// </summary>
        [Input("sourceDisk")]
        public Input<string>? SourceDisk { get; set; }

        /// <summary>
        /// URL of the source image used to create this image. In order to create an image, you must provide the full or partial
        /// URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Input("sourceImage")]
        public Input<string>? SourceImage { get; set; }

        /// <summary>
        /// URL of the source snapshot used to create this image.
        /// In order to create an image, you must provide the full or partial URL of one of the following:
        /// * The selfLink URL
        /// * This property
        /// * The sourceImage URL
        /// * The rawDisk.source URL
        /// * The sourceDisk URL
        /// </summary>
        [Input("sourceSnapshot")]
        public Input<string>? SourceSnapshot { get; set; }

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// Cloud Storage bucket storage location of the image
        /// (regional or multi-regional).
        /// Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public ImageState()
        {
        }
        public static new ImageState Empty => new ImageState();
    }
}
