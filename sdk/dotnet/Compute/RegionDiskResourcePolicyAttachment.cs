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
    /// Adds existing resource policies to a disk. You can only add one policy
    /// which will be applied to this disk for scheduling snapshot creation.
    /// 
    /// &gt; **Note:** This resource does not support zonal disks (`gcp.compute.Disk`). For zonal disks, please refer to the `gcp.compute.DiskResourcePolicyAttachment` resource.
    /// 
    /// ## Example Usage
    /// ### Region Disk Resource Policy Attachment Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var disk = new Gcp.Compute.Disk("disk", new()
    ///     {
    ///         Image = "debian-cloud/debian-11",
    ///         Size = 50,
    ///         Type = "pd-ssd",
    ///         Zone = "us-central1-a",
    ///     });
    /// 
    ///     var snapdisk = new Gcp.Compute.Snapshot("snapdisk", new()
    ///     {
    ///         SourceDisk = disk.Name,
    ///         Zone = "us-central1-a",
    ///     });
    /// 
    ///     var ssd = new Gcp.Compute.RegionDisk("ssd", new()
    ///     {
    ///         ReplicaZones = new[]
    ///         {
    ///             "us-central1-a",
    ///             "us-central1-f",
    ///         },
    ///         Snapshot = snapdisk.Id,
    ///         Size = 50,
    ///         Type = "pd-ssd",
    ///         Region = "us-central1",
    ///     });
    /// 
    ///     var attachment = new Gcp.Compute.RegionDiskResourcePolicyAttachment("attachment", new()
    ///     {
    ///         Disk = ssd.Name,
    ///         Region = "us-central1",
    ///     });
    /// 
    ///     var policy = new Gcp.Compute.ResourcePolicy("policy", new()
    ///     {
    ///         Region = "us-central1",
    ///         SnapshotSchedulePolicy = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyArgs
    ///         {
    ///             Schedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs
    ///             {
    ///                 DailySchedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs
    ///                 {
    ///                     DaysInCycle = 1,
    ///                     StartTime = "04:00",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var myImage = Gcp.Compute.GetImage.Invoke(new()
    ///     {
    ///         Family = "debian-11",
    ///         Project = "debian-cloud",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RegionDiskResourcePolicyAttachment can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default projects/{{project}}/regions/{{region}}/disks/{{disk}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{project}}/{{region}}/{{disk}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{region}}/{{disk}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment default {{disk}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment")]
    public partial class RegionDiskResourcePolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the regional disk in which the resource policies are attached to.
        /// </summary>
        [Output("disk")]
        public Output<string> Disk { get; private set; } = null!;

        /// <summary>
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
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
        /// A reference to the region where the disk resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a RegionDiskResourcePolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionDiskResourcePolicyAttachment(string name, RegionDiskResourcePolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, args ?? new RegionDiskResourcePolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionDiskResourcePolicyAttachment(string name, Input<string> id, RegionDiskResourcePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionDiskResourcePolicyAttachment:RegionDiskResourcePolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionDiskResourcePolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionDiskResourcePolicyAttachment Get(string name, Input<string> id, RegionDiskResourcePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionDiskResourcePolicyAttachment(name, id, state, options);
        }
    }

    public sealed class RegionDiskResourcePolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the regional disk in which the resource policies are attached to.
        /// </summary>
        [Input("disk", required: true)]
        public Input<string> Disk { get; set; } = null!;

        /// <summary>
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
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
        /// A reference to the region where the disk resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RegionDiskResourcePolicyAttachmentArgs()
        {
        }
        public static new RegionDiskResourcePolicyAttachmentArgs Empty => new RegionDiskResourcePolicyAttachmentArgs();
    }

    public sealed class RegionDiskResourcePolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the regional disk in which the resource policies are attached to.
        /// </summary>
        [Input("disk")]
        public Input<string>? Disk { get; set; }

        /// <summary>
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
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
        /// A reference to the region where the disk resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RegionDiskResourcePolicyAttachmentState()
        {
        }
        public static new RegionDiskResourcePolicyAttachmentState Empty => new RegionDiskResourcePolicyAttachmentState();
    }
}
