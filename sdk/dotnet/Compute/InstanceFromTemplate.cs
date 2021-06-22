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
    /// Manages a VM instance resource within GCE. For more information see
    /// [the official documentation](https://cloud.google.com/compute/docs/instances)
    /// and
    /// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
    /// 
    /// This resource is specifically to create a compute instance from a given
    /// `source_instance_template`. To create an instance without a template, use the
    /// `gcp.compute.Instance` resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tplInstanceTemplate = new Gcp.Compute.InstanceTemplate("tplInstanceTemplate", new Gcp.Compute.InstanceTemplateArgs
    ///         {
    ///             MachineType = "e2-medium",
    ///             Disks = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///                 {
    ///                     SourceImage = "debian-cloud/debian-9",
    ///                     AutoDelete = true,
    ///                     DiskSizeGb = 100,
    ///                     Boot = true,
    ///                 },
    ///             },
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceTemplateNetworkInterfaceArgs
    ///                 {
    ///                     Network = "default",
    ///                 },
    ///             },
    ///             Metadata = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             CanIpForward = true,
    ///         });
    ///         var tplInstanceFromTemplate = new Gcp.Compute.InstanceFromTemplate("tplInstanceFromTemplate", new Gcp.Compute.InstanceFromTemplateArgs
    ///         {
    ///             Zone = "us-central1-a",
    ///             SourceInstanceTemplate = tplInstanceTemplate.Id,
    ///             CanIpForward = false,
    ///             Labels = 
    ///             {
    ///                 { "my_key", "my_value" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:compute/instanceFromTemplate:InstanceFromTemplate")]
    public partial class InstanceFromTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
        /// stopping the instance without setting this field, the update will fail.
        /// </summary>
        [Output("allowStoppingForUpdate")]
        public Output<bool> AllowStoppingForUpdate { get; private set; } = null!;

        /// <summary>
        /// List of disks attached to the instance
        /// </summary>
        [Output("attachedDisks")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateAttachedDisk>> AttachedDisks { get; private set; } = null!;

        /// <summary>
        /// The boot disk for the instance.
        /// </summary>
        [Output("bootDisk")]
        public Output<Outputs.InstanceFromTemplateBootDisk> BootDisk { get; private set; } = null!;

        /// <summary>
        /// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        /// </summary>
        [Output("canIpForward")]
        public Output<bool> CanIpForward { get; private set; } = null!;

        /// <summary>
        /// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
        /// to create.
        /// </summary>
        [Output("confidentialInstanceConfig")]
        public Output<Outputs.InstanceFromTemplateConfidentialInstanceConfig> ConfidentialInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        [Output("cpuPlatform")]
        public Output<string> CpuPlatform { get; private set; } = null!;

        /// <summary>
        /// Current status of the instance.
        /// </summary>
        [Output("currentStatus")]
        public Output<string> CurrentStatus { get; private set; } = null!;

        /// <summary>
        /// Whether deletion protection is enabled on this instance.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// A brief description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Desired status of the instance. Either "RUNNING" or "TERMINATED".
        /// </summary>
        [Output("desiredStatus")]
        public Output<string> DesiredStatus { get; private set; } = null!;

        /// <summary>
        /// Whether the instance has virtual displays enabled.
        /// </summary>
        [Output("enableDisplay")]
        public Output<bool> EnableDisplay { get; private set; } = null!;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance.
        /// </summary>
        [Output("guestAccelerators")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateGuestAccelerator>> GuestAccelerators { get; private set; } = null!;

        /// <summary>
        /// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
        /// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
        /// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The server-assigned unique identifier of this instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The unique fingerprint of the labels.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs assigned to the instance.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Metadata key/value pairs made available within the instance.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        [Output("metadataFingerprint")]
        public Output<string> MetadataFingerprint { get; private set; } = null!;

        /// <summary>
        /// Metadata startup scripts made available within the instance.
        /// </summary>
        [Output("metadataStartupScript")]
        public Output<string> MetadataStartupScript { get; private set; } = null!;

        /// <summary>
        /// The minimum CPU platform specified for the VM instance.
        /// </summary>
        [Output("minCpuPlatform")]
        public Output<string> MinCpuPlatform { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The networks attached to the instance.
        /// </summary>
        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        /// <summary>
        /// Configures network performance settings for the instance. If not specified, the instance will be created with its
        /// default network performance configuration.
        /// </summary>
        [Output("networkPerformanceConfig")]
        public Output<Outputs.InstanceFromTemplateNetworkPerformanceConfig> NetworkPerformanceConfig { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
        /// self_link nor project are provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        [Output("reservationAffinity")]
        public Output<Outputs.InstanceFromTemplateReservationAffinity> ReservationAffinity { get; private set; } = null!;

        /// <summary>
        /// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
        /// instance to recreate. Currently a max of 1 resource policy is supported.
        /// </summary>
        [Output("resourcePolicies")]
        public Output<string> ResourcePolicies { get; private set; } = null!;

        /// <summary>
        /// The scheduling strategy being used by the instance.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.InstanceFromTemplateScheduling> Scheduling { get; private set; } = null!;

        /// <summary>
        /// The scratch disks attached to the instance.
        /// </summary>
        [Output("scratchDisks")]
        public Output<ImmutableArray<Outputs.InstanceFromTemplateScratchDisk>> ScratchDisks { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The service account to attach to the instance.
        /// </summary>
        [Output("serviceAccount")]
        public Output<Outputs.InstanceFromTemplateServiceAccount> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// The shielded vm config being used by the instance.
        /// </summary>
        [Output("shieldedInstanceConfig")]
        public Output<Outputs.InstanceFromTemplateShieldedInstanceConfig> ShieldedInstanceConfig { get; private set; } = null!;

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Output("sourceInstanceTemplate")]
        public Output<string> SourceInstanceTemplate { get; private set; } = null!;

        /// <summary>
        /// The list of tags attached to the instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        [Output("tagsFingerprint")]
        public Output<string> TagsFingerprint { get; private set; } = null!;

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceFromTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceFromTemplate(string name, InstanceFromTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, args ?? new InstanceFromTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceFromTemplate(string name, Input<string> id, InstanceFromTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/instanceFromTemplate:InstanceFromTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceFromTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceFromTemplate Get(string name, Input<string> id, InstanceFromTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceFromTemplate(name, id, state, options);
        }
    }

    public sealed class InstanceFromTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
        /// stopping the instance without setting this field, the update will fail.
        /// </summary>
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        [Input("attachedDisks")]
        private InputList<Inputs.InstanceFromTemplateAttachedDiskArgs>? _attachedDisks;

        /// <summary>
        /// List of disks attached to the instance
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateAttachedDiskArgs> AttachedDisks
        {
            get => _attachedDisks ?? (_attachedDisks = new InputList<Inputs.InstanceFromTemplateAttachedDiskArgs>());
            set => _attachedDisks = value;
        }

        /// <summary>
        /// The boot disk for the instance.
        /// </summary>
        [Input("bootDisk")]
        public Input<Inputs.InstanceFromTemplateBootDiskArgs>? BootDisk { get; set; }

        /// <summary>
        /// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        /// <summary>
        /// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
        /// to create.
        /// </summary>
        [Input("confidentialInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateConfidentialInstanceConfigArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// Whether deletion protection is enabled on this instance.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// A brief description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired status of the instance. Either "RUNNING" or "TERMINATED".
        /// </summary>
        [Input("desiredStatus")]
        public Input<string>? DesiredStatus { get; set; }

        /// <summary>
        /// Whether the instance has virtual displays enabled.
        /// </summary>
        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs>? _guestAccelerators;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceFromTemplateGuestAcceleratorArgs>());
            set => _guestAccelerators = value;
        }

        /// <summary>
        /// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
        /// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
        /// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs assigned to the instance.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata key/value pairs made available within the instance.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Metadata startup scripts made available within the instance.
        /// </summary>
        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        /// <summary>
        /// The minimum CPU platform specified for the VM instance.
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// The networks attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceFromTemplateNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// Configures network performance settings for the instance. If not specified, the instance will be created with its
        /// default network performance configuration.
        /// </summary>
        [Input("networkPerformanceConfig")]
        public Input<Inputs.InstanceFromTemplateNetworkPerformanceConfigArgs>? NetworkPerformanceConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
        /// self_link nor project are provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        [Input("reservationAffinity")]
        public Input<Inputs.InstanceFromTemplateReservationAffinityArgs>? ReservationAffinity { get; set; }

        /// <summary>
        /// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
        /// instance to recreate. Currently a max of 1 resource policy is supported.
        /// </summary>
        [Input("resourcePolicies")]
        public Input<string>? ResourcePolicies { get; set; }

        /// <summary>
        /// The scheduling strategy being used by the instance.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.InstanceFromTemplateSchedulingArgs>? Scheduling { get; set; }

        [Input("scratchDisks")]
        private InputList<Inputs.InstanceFromTemplateScratchDiskArgs>? _scratchDisks;

        /// <summary>
        /// The scratch disks attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateScratchDiskArgs> ScratchDisks
        {
            get => _scratchDisks ?? (_scratchDisks = new InputList<Inputs.InstanceFromTemplateScratchDiskArgs>());
            set => _scratchDisks = value;
        }

        /// <summary>
        /// The service account to attach to the instance.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.InstanceFromTemplateServiceAccountArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// The shielded vm config being used by the instance.
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateShieldedInstanceConfigArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Input("sourceInstanceTemplate", required: true)]
        public Input<string> SourceInstanceTemplate { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of tags attached to the instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceFromTemplateArgs()
        {
        }
    }

    public sealed class InstanceFromTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
        /// stopping the instance without setting this field, the update will fail.
        /// </summary>
        [Input("allowStoppingForUpdate")]
        public Input<bool>? AllowStoppingForUpdate { get; set; }

        [Input("attachedDisks")]
        private InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs>? _attachedDisks;

        /// <summary>
        /// List of disks attached to the instance
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs> AttachedDisks
        {
            get => _attachedDisks ?? (_attachedDisks = new InputList<Inputs.InstanceFromTemplateAttachedDiskGetArgs>());
            set => _attachedDisks = value;
        }

        /// <summary>
        /// The boot disk for the instance.
        /// </summary>
        [Input("bootDisk")]
        public Input<Inputs.InstanceFromTemplateBootDiskGetArgs>? BootDisk { get; set; }

        /// <summary>
        /// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        /// <summary>
        /// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
        /// to create.
        /// </summary>
        [Input("confidentialInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateConfidentialInstanceConfigGetArgs>? ConfidentialInstanceConfig { get; set; }

        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        [Input("cpuPlatform")]
        public Input<string>? CpuPlatform { get; set; }

        /// <summary>
        /// Current status of the instance.
        /// </summary>
        [Input("currentStatus")]
        public Input<string>? CurrentStatus { get; set; }

        /// <summary>
        /// Whether deletion protection is enabled on this instance.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// A brief description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired status of the instance. Either "RUNNING" or "TERMINATED".
        /// </summary>
        [Input("desiredStatus")]
        public Input<string>? DesiredStatus { get; set; }

        /// <summary>
        /// Whether the instance has virtual displays enabled.
        /// </summary>
        [Input("enableDisplay")]
        public Input<bool>? EnableDisplay { get; set; }

        [Input("guestAccelerators")]
        private InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs>? _guestAccelerators;

        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs> GuestAccelerators
        {
            get => _guestAccelerators ?? (_guestAccelerators = new InputList<Inputs.InstanceFromTemplateGuestAcceleratorGetArgs>());
            set => _guestAccelerators = value;
        }

        /// <summary>
        /// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
        /// labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
        /// entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The server-assigned unique identifier of this instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The unique fingerprint of the labels.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs assigned to the instance.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The machine type to create.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata key/value pairs made available within the instance.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        [Input("metadataFingerprint")]
        public Input<string>? MetadataFingerprint { get; set; }

        /// <summary>
        /// Metadata startup scripts made available within the instance.
        /// </summary>
        [Input("metadataStartupScript")]
        public Input<string>? MetadataStartupScript { get; set; }

        /// <summary>
        /// The minimum CPU platform specified for the VM instance.
        /// </summary>
        [Input("minCpuPlatform")]
        public Input<string>? MinCpuPlatform { get; set; }

        /// <summary>
        /// A unique name for the resource, required by GCE.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs>? _networkInterfaces;

        /// <summary>
        /// The networks attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceFromTemplateNetworkInterfaceGetArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// Configures network performance settings for the instance. If not specified, the instance will be created with its
        /// default network performance configuration.
        /// </summary>
        [Input("networkPerformanceConfig")]
        public Input<Inputs.InstanceFromTemplateNetworkPerformanceConfigGetArgs>? NetworkPerformanceConfig { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
        /// self_link nor project are provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specifies the reservations that this instance can consume from.
        /// </summary>
        [Input("reservationAffinity")]
        public Input<Inputs.InstanceFromTemplateReservationAffinityGetArgs>? ReservationAffinity { get; set; }

        /// <summary>
        /// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
        /// instance to recreate. Currently a max of 1 resource policy is supported.
        /// </summary>
        [Input("resourcePolicies")]
        public Input<string>? ResourcePolicies { get; set; }

        /// <summary>
        /// The scheduling strategy being used by the instance.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.InstanceFromTemplateSchedulingGetArgs>? Scheduling { get; set; }

        [Input("scratchDisks")]
        private InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs>? _scratchDisks;

        /// <summary>
        /// The scratch disks attached to the instance.
        /// </summary>
        public InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs> ScratchDisks
        {
            get => _scratchDisks ?? (_scratchDisks = new InputList<Inputs.InstanceFromTemplateScratchDiskGetArgs>());
            set => _scratchDisks = value;
        }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The service account to attach to the instance.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.InstanceFromTemplateServiceAccountGetArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// The shielded vm config being used by the instance.
        /// </summary>
        [Input("shieldedInstanceConfig")]
        public Input<Inputs.InstanceFromTemplateShieldedInstanceConfigGetArgs>? ShieldedInstanceConfig { get; set; }

        /// <summary>
        /// Name or self link of an instance
        /// template to create the instance based on.
        /// </summary>
        [Input("sourceInstanceTemplate")]
        public Input<string>? SourceInstanceTemplate { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of tags attached to the instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        [Input("tagsFingerprint")]
        public Input<string>? TagsFingerprint { get; set; }

        /// <summary>
        /// The zone that the machine should be created in. If not
        /// set, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceFromTemplateState()
        {
        }
    }
}
