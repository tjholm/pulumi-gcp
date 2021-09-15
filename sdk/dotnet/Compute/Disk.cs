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
    /// Persistent disks are durable storage devices that function similarly to
    /// the physical disks in a desktop or a server. Compute Engine manages the
    /// hardware behind these devices to ensure data redundancy and optimize
    /// performance for you. Persistent disks are available as either standard
    /// hard disk drives (HDD) or solid-state drives (SSD).
    /// 
    /// Persistent disks are located independently from your virtual machine
    /// instances, so you can detach or move persistent disks to keep your data
    /// even after you delete your instances. Persistent disk performance scales
    /// automatically with size, so you can resize your existing persistent disks
    /// or add more persistent disks to an instance to meet your performance and
    /// storage space requirements.
    /// 
    /// Add a persistent disk to your instance when you need reliable and
    /// affordable storage with consistent performance characteristics.
    /// 
    /// To get more information about Disk, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/disks)
    /// * How-to Guides
    ///     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
    /// 
    /// &gt; **Warning:** All arguments including `disk_encryption_key.raw_key` will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// 
    /// ## Example Usage
    /// ### Disk Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.Compute.Disk("default", new Gcp.Compute.DiskArgs
    ///         {
    ///             Image = "debian-8-jessie-v20170523",
    ///             Labels = 
    ///             {
    ///                 { "environment", "dev" },
    ///             },
    ///             PhysicalBlockSizeBytes = 4096,
    ///             Type = "pd-ssd",
    ///             Zone = "us-central1-a",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Disk can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/disk:Disk default projects/{{project}}/zones/{{zone}}/disks/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/disk:Disk default {{project}}/{{zone}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/disk:Disk default {{zone}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/disk:Disk default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/disk:Disk")]
    public partial class Disk : Pulumi.CustomResource
    {
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
        /// Encrypts the disk using a customer-supplied encryption key.
        /// After you encrypt a disk with a customer-supplied key, you must
        /// provide the same key if you use the disk later (e.g. to create a disk
        /// snapshot or an image, or to attach the disk to a virtual machine).
        /// Customer-supplied encryption keys do not protect access to metadata of
        /// the disk.
        /// If you do not provide an encryption key when creating the disk, then
        /// the disk will be encrypted using an automatically generated key and
        /// you do not need to provide a key to use the disk later.
        /// Structure is documented below.
        /// </summary>
        [Output("diskEncryptionKey")]
        public Output<Outputs.DiskDiskEncryptionKey?> DiskEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The image from which to initialize this disk. This can be
        /// one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
        /// `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
        /// `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
        /// `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
        /// images names must include the family name. If they don't, use the
        /// [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
        /// For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
        /// These images can be referred by family name here.
        /// </summary>
        [Output("image")]
        public Output<string?> Image { get; private set; } = null!;

        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        [Output("interface")]
        public Output<string?> Interface { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this disk.  A list of key-&gt;value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Last attach timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastAttachTimestamp")]
        public Output<string> LastAttachTimestamp { get; private set; } = null!;

        /// <summary>
        /// Last detach timestamp in RFC3339 text format.
        /// </summary>
        [Output("lastDetachTimestamp")]
        public Output<string> LastDetachTimestamp { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        [Output("multiWriter")]
        public Output<bool?> MultiWriter { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present
        /// in a request, a default value is used. Currently supported sizes
        /// are 4096 and 16384, other sizes may be added in the future.
        /// If an unsupported value is requested, the error message will list
        /// the supported values for the caller's project.
        /// </summary>
        [Output("physicalBlockSizeBytes")]
        public Output<int> PhysicalBlockSizeBytes { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Indicates how many IOPS must be provisioned for the disk.
        /// </summary>
        [Output("provisionedIops")]
        public Output<int?> ProvisionedIops { get; private set; } = null!;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// ~&gt;**NOTE** This value does not support updating the
        /// resource policy, as resource policies can not be updated more than
        /// one at a time. Use
        /// `gcp.compute.DiskResourcePolicyAttachment`
        /// to allow for updating the resource policy attached to the disk.
        /// </summary>
        [Output("resourcePolicies")]
        public Output<ImmutableArray<string>> ResourcePolicies { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Size of the persistent disk, specified in GB. You can specify this
        /// field when creating a persistent disk using the `image` or
        /// `snapshot` parameter, or specify it alone to create an empty
        /// persistent disk.
        /// If you specify this field along with `image` or `snapshot`,
        /// the value must not be less than the size of the image
        /// or the size of the snapshot.
        /// ~&gt;**NOTE** If you change the size, the provider updates the disk size
        /// if upsizing is detected but recreates the disk if downsizing is requested.
        /// You can add `lifecycle.prevent_destroy` in the config to prevent destroying
        /// and recreating.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as
        /// a partial or full URL to the resource. If the snapshot is in another
        /// project than this disk, you must supply a full URL. For example, the
        /// following are valid values:
        /// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
        /// * `projects/project/global/snapshots/snapshot`
        /// * `global/snapshots/snapshot`
        /// * `snapshot`
        /// </summary>
        [Output("snapshot")]
        public Output<string?> Snapshot { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if
        /// the source image is protected by a customer-supplied encryption key.
        /// Structure is documented below.
        /// </summary>
        [Output("sourceImageEncryptionKey")]
        public Output<Outputs.DiskSourceImageEncryptionKey?> SourceImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
        /// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
        /// under the same name, the source image ID would identify the exact version of the image that was used.
        /// </summary>
        [Output("sourceImageId")]
        public Output<string> SourceImageId { get; private set; } = null!;

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required
        /// if the source snapshot is protected by a customer-supplied encryption
        /// key.
        /// Structure is documented below.
        /// </summary>
        [Output("sourceSnapshotEncryptionKey")]
        public Output<Outputs.DiskSourceSnapshotEncryptionKey?> SourceSnapshotEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
        /// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
        /// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        /// </summary>
        [Output("sourceSnapshotId")]
        public Output<string> SourceSnapshotId { get; private set; } = null!;

        /// <summary>
        /// URL of the disk type resource describing which disk type to use to
        /// create the disk. Provide this when creating the disk.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;

        /// <summary>
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/disk:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/disk:Disk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, state, options);
        }
    }

    public sealed class DiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Encrypts the disk using a customer-supplied encryption key.
        /// After you encrypt a disk with a customer-supplied key, you must
        /// provide the same key if you use the disk later (e.g. to create a disk
        /// snapshot or an image, or to attach the disk to a virtual machine).
        /// Customer-supplied encryption keys do not protect access to metadata of
        /// the disk.
        /// If you do not provide an encryption key when creating the disk, then
        /// the disk will be encrypted using an automatically generated key and
        /// you do not need to provide a key to use the disk later.
        /// Structure is documented below.
        /// </summary>
        [Input("diskEncryptionKey")]
        public Input<Inputs.DiskDiskEncryptionKeyArgs>? DiskEncryptionKey { get; set; }

        /// <summary>
        /// The image from which to initialize this disk. This can be
        /// one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
        /// `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
        /// `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
        /// `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
        /// images names must include the family name. If they don't, use the
        /// [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
        /// For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
        /// These images can be referred by family name here.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this disk.  A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        [Input("multiWriter")]
        public Input<bool>? MultiWriter { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present
        /// in a request, a default value is used. Currently supported sizes
        /// are 4096 and 16384, other sizes may be added in the future.
        /// If an unsupported value is requested, the error message will list
        /// the supported values for the caller's project.
        /// </summary>
        [Input("physicalBlockSizeBytes")]
        public Input<int>? PhysicalBlockSizeBytes { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates how many IOPS must be provisioned for the disk.
        /// </summary>
        [Input("provisionedIops")]
        public Input<int>? ProvisionedIops { get; set; }

        [Input("resourcePolicies")]
        private InputList<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// ~&gt;**NOTE** This value does not support updating the
        /// resource policy, as resource policies can not be updated more than
        /// one at a time. Use
        /// `gcp.compute.DiskResourcePolicyAttachment`
        /// to allow for updating the resource policy attached to the disk.
        /// </summary>
        public InputList<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// Size of the persistent disk, specified in GB. You can specify this
        /// field when creating a persistent disk using the `image` or
        /// `snapshot` parameter, or specify it alone to create an empty
        /// persistent disk.
        /// If you specify this field along with `image` or `snapshot`,
        /// the value must not be less than the size of the image
        /// or the size of the snapshot.
        /// ~&gt;**NOTE** If you change the size, the provider updates the disk size
        /// if upsizing is detected but recreates the disk if downsizing is requested.
        /// You can add `lifecycle.prevent_destroy` in the config to prevent destroying
        /// and recreating.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as
        /// a partial or full URL to the resource. If the snapshot is in another
        /// project than this disk, you must supply a full URL. For example, the
        /// following are valid values:
        /// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
        /// * `projects/project/global/snapshots/snapshot`
        /// * `global/snapshots/snapshot`
        /// * `snapshot`
        /// </summary>
        [Input("snapshot")]
        public Input<string>? Snapshot { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if
        /// the source image is protected by a customer-supplied encryption key.
        /// Structure is documented below.
        /// </summary>
        [Input("sourceImageEncryptionKey")]
        public Input<Inputs.DiskSourceImageEncryptionKeyArgs>? SourceImageEncryptionKey { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required
        /// if the source snapshot is protected by a customer-supplied encryption
        /// key.
        /// Structure is documented below.
        /// </summary>
        [Input("sourceSnapshotEncryptionKey")]
        public Input<Inputs.DiskSourceSnapshotEncryptionKeyArgs>? SourceSnapshotEncryptionKey { get; set; }

        /// <summary>
        /// URL of the disk type resource describing which disk type to use to
        /// create the disk. Provide this when creating the disk.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DiskArgs()
        {
        }
    }

    public sealed class DiskState : Pulumi.ResourceArgs
    {
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
        /// Encrypts the disk using a customer-supplied encryption key.
        /// After you encrypt a disk with a customer-supplied key, you must
        /// provide the same key if you use the disk later (e.g. to create a disk
        /// snapshot or an image, or to attach the disk to a virtual machine).
        /// Customer-supplied encryption keys do not protect access to metadata of
        /// the disk.
        /// If you do not provide an encryption key when creating the disk, then
        /// the disk will be encrypted using an automatically generated key and
        /// you do not need to provide a key to use the disk later.
        /// Structure is documented below.
        /// </summary>
        [Input("diskEncryptionKey")]
        public Input<Inputs.DiskDiskEncryptionKeyGetArgs>? DiskEncryptionKey { get; set; }

        /// <summary>
        /// The image from which to initialize this disk. This can be
        /// one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
        /// `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
        /// `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
        /// `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
        /// images names must include the family name. If they don't, use the
        /// [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
        /// For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
        /// These images can be referred by family name here.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this disk.  A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Last attach timestamp in RFC3339 text format.
        /// </summary>
        [Input("lastAttachTimestamp")]
        public Input<string>? LastAttachTimestamp { get; set; }

        /// <summary>
        /// Last detach timestamp in RFC3339 text format.
        /// </summary>
        [Input("lastDetachTimestamp")]
        public Input<string>? LastDetachTimestamp { get; set; }

        /// <summary>
        /// Indicates whether or not the disk can be read/write attached to more than one instance.
        /// </summary>
        [Input("multiWriter")]
        public Input<bool>? MultiWriter { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Physical block size of the persistent disk, in bytes. If not present
        /// in a request, a default value is used. Currently supported sizes
        /// are 4096 and 16384, other sizes may be added in the future.
        /// If an unsupported value is requested, the error message will list
        /// the supported values for the caller's project.
        /// </summary>
        [Input("physicalBlockSizeBytes")]
        public Input<int>? PhysicalBlockSizeBytes { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Indicates how many IOPS must be provisioned for the disk.
        /// </summary>
        [Input("provisionedIops")]
        public Input<int>? ProvisionedIops { get; set; }

        [Input("resourcePolicies")]
        private InputList<string>? _resourcePolicies;

        /// <summary>
        /// Resource policies applied to this disk for automatic snapshot creations.
        /// ~&gt;**NOTE** This value does not support updating the
        /// resource policy, as resource policies can not be updated more than
        /// one at a time. Use
        /// `gcp.compute.DiskResourcePolicyAttachment`
        /// to allow for updating the resource policy attached to the disk.
        /// </summary>
        public InputList<string> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<string>());
            set => _resourcePolicies = value;
        }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Size of the persistent disk, specified in GB. You can specify this
        /// field when creating a persistent disk using the `image` or
        /// `snapshot` parameter, or specify it alone to create an empty
        /// persistent disk.
        /// If you specify this field along with `image` or `snapshot`,
        /// the value must not be less than the size of the image
        /// or the size of the snapshot.
        /// ~&gt;**NOTE** If you change the size, the provider updates the disk size
        /// if upsizing is detected but recreates the disk if downsizing is requested.
        /// You can add `lifecycle.prevent_destroy` in the config to prevent destroying
        /// and recreating.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The source snapshot used to create this disk. You can provide this as
        /// a partial or full URL to the resource. If the snapshot is in another
        /// project than this disk, you must supply a full URL. For example, the
        /// following are valid values:
        /// * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
        /// * `projects/project/global/snapshots/snapshot`
        /// * `global/snapshots/snapshot`
        /// * `snapshot`
        /// </summary>
        [Input("snapshot")]
        public Input<string>? Snapshot { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source image. Required if
        /// the source image is protected by a customer-supplied encryption key.
        /// Structure is documented below.
        /// </summary>
        [Input("sourceImageEncryptionKey")]
        public Input<Inputs.DiskSourceImageEncryptionKeyGetArgs>? SourceImageEncryptionKey { get; set; }

        /// <summary>
        /// The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
        /// persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
        /// under the same name, the source image ID would identify the exact version of the image that was used.
        /// </summary>
        [Input("sourceImageId")]
        public Input<string>? SourceImageId { get; set; }

        /// <summary>
        /// The customer-supplied encryption key of the source snapshot. Required
        /// if the source snapshot is protected by a customer-supplied encryption
        /// key.
        /// Structure is documented below.
        /// </summary>
        [Input("sourceSnapshotEncryptionKey")]
        public Input<Inputs.DiskSourceSnapshotEncryptionKeyGetArgs>? SourceSnapshotEncryptionKey { get; set; }

        /// <summary>
        /// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
        /// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
        /// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        /// </summary>
        [Input("sourceSnapshotId")]
        public Input<string>? SourceSnapshotId { get; set; }

        /// <summary>
        /// URL of the disk type resource describing which disk type to use to
        /// create the disk. Provide this when creating the disk.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        /// <summary>
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DiskState()
        {
        }
    }
}
