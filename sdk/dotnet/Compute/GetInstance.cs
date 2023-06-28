// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetInstance
    {
        /// <summary>
        /// Get information about a VM instance resource within GCE. For more information see
        /// [the official documentation](https://cloud.google.com/compute/docs/instances)
        /// and
        /// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var appserver = Gcp.Compute.GetInstance.Invoke(new()
        ///     {
        ///         Name = "primary-application-server",
        ///         Zone = "us-central1-a",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("gcp:compute/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a VM instance resource within GCE. For more information see
        /// [the official documentation](https://cloud.google.com/compute/docs/instances)
        /// and
        /// [API](https://cloud.google.com/compute/docs/reference/latest/instances).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var appserver = Gcp.Compute.GetInstance.Invoke(new()
        ///     {
        ///         Name = "primary-application-server",
        ///         Zone = "us-central1-a",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("gcp:compute/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If `self_link` is provided, this value is ignored.  If neither `self_link`
        /// nor `project` are provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The self link of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("selfLink")]
        public string? SelfLink { get; set; }

        /// <summary>
        /// The zone of the instance. If `self_link` is provided, this
        /// value is ignored.  If neither `self_link` nor `zone` are provided, the
        /// provider zone is used.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If `self_link` is provided, this value is ignored.  If neither `self_link`
        /// nor `project` are provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The self link of the instance. One of `name` or `self_link` must be provided.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The zone of the instance. If `self_link` is provided, this
        /// value is ignored.  If neither `self_link` nor `zone` are provided, the
        /// provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceAdvancedMachineFeatureResult> AdvancedMachineFeatures;
        public readonly bool AllowStoppingForUpdate;
        /// <summary>
        /// List of disks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceAttachedDiskResult> AttachedDisks;
        /// <summary>
        /// The boot disk for the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceBootDiskResult> BootDisks;
        /// <summary>
        /// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
        /// </summary>
        public readonly bool CanIpForward;
        public readonly ImmutableArray<Outputs.GetInstanceConfidentialInstanceConfigResult> ConfidentialInstanceConfigs;
        /// <summary>
        /// The CPU platform used by this instance.
        /// </summary>
        public readonly string CpuPlatform;
        /// <summary>
        /// The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
        /// </summary>
        public readonly string CurrentStatus;
        /// <summary>
        /// Whether deletion protection is enabled on this instance.
        /// </summary>
        public readonly bool DeletionProtection;
        /// <summary>
        /// A brief description of the resource.
        /// </summary>
        public readonly string Description;
        public readonly string DesiredStatus;
        /// <summary>
        /// Whether the instance has virtual displays enabled.
        /// </summary>
        public readonly bool EnableDisplay;
        /// <summary>
        /// List of the type and count of accelerator cards attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceGuestAcceleratorResult> GuestAccelerators;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The server-assigned unique identifier of this instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The unique fingerprint of the labels.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// A set of key/value label pairs assigned to the disk.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The machine type to create.
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Metadata key/value pairs made available within the instance.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// The unique fingerprint of the metadata.
        /// </summary>
        public readonly string MetadataFingerprint;
        public readonly string MetadataStartupScript;
        /// <summary>
        /// The minimum CPU platform specified for the VM instance.
        /// </summary>
        public readonly string MinCpuPlatform;
        public readonly string? Name;
        /// <summary>
        /// The networks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// The network performance configuration setting for the instance, if set. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceNetworkPerformanceConfigResult> NetworkPerformanceConfigs;
        public readonly ImmutableArray<Outputs.GetInstanceParamResult> Params;
        public readonly string? Project;
        public readonly ImmutableArray<Outputs.GetInstanceReservationAffinityResult> ReservationAffinities;
        public readonly ImmutableArray<string> ResourcePolicies;
        /// <summary>
        /// The scheduling strategy being used by the instance. Structure is documented below
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSchedulingResult> Schedulings;
        /// <summary>
        /// The scratch disks attached to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceScratchDiskResult> ScratchDisks;
        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        public readonly string? SelfLink;
        /// <summary>
        /// The service account to attach to the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceServiceAccountResult> ServiceAccounts;
        /// <summary>
        /// The shielded vm config being used by the instance. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceShieldedInstanceConfigResult> ShieldedInstanceConfigs;
        /// <summary>
        /// The list of tags attached to the instance.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The unique fingerprint of the tags.
        /// </summary>
        public readonly string TagsFingerprint;
        public readonly string? Zone;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.GetInstanceAdvancedMachineFeatureResult> advancedMachineFeatures,

            bool allowStoppingForUpdate,

            ImmutableArray<Outputs.GetInstanceAttachedDiskResult> attachedDisks,

            ImmutableArray<Outputs.GetInstanceBootDiskResult> bootDisks,

            bool canIpForward,

            ImmutableArray<Outputs.GetInstanceConfidentialInstanceConfigResult> confidentialInstanceConfigs,

            string cpuPlatform,

            string currentStatus,

            bool deletionProtection,

            string description,

            string desiredStatus,

            bool enableDisplay,

            ImmutableArray<Outputs.GetInstanceGuestAcceleratorResult> guestAccelerators,

            string hostname,

            string id,

            string instanceId,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string machineType,

            ImmutableDictionary<string, string> metadata,

            string metadataFingerprint,

            string metadataStartupScript,

            string minCpuPlatform,

            string? name,

            ImmutableArray<Outputs.GetInstanceNetworkInterfaceResult> networkInterfaces,

            ImmutableArray<Outputs.GetInstanceNetworkPerformanceConfigResult> networkPerformanceConfigs,

            ImmutableArray<Outputs.GetInstanceParamResult> @params,

            string? project,

            ImmutableArray<Outputs.GetInstanceReservationAffinityResult> reservationAffinities,

            ImmutableArray<string> resourcePolicies,

            ImmutableArray<Outputs.GetInstanceSchedulingResult> schedulings,

            ImmutableArray<Outputs.GetInstanceScratchDiskResult> scratchDisks,

            string? selfLink,

            ImmutableArray<Outputs.GetInstanceServiceAccountResult> serviceAccounts,

            ImmutableArray<Outputs.GetInstanceShieldedInstanceConfigResult> shieldedInstanceConfigs,

            ImmutableArray<string> tags,

            string tagsFingerprint,

            string? zone)
        {
            AdvancedMachineFeatures = advancedMachineFeatures;
            AllowStoppingForUpdate = allowStoppingForUpdate;
            AttachedDisks = attachedDisks;
            BootDisks = bootDisks;
            CanIpForward = canIpForward;
            ConfidentialInstanceConfigs = confidentialInstanceConfigs;
            CpuPlatform = cpuPlatform;
            CurrentStatus = currentStatus;
            DeletionProtection = deletionProtection;
            Description = description;
            DesiredStatus = desiredStatus;
            EnableDisplay = enableDisplay;
            GuestAccelerators = guestAccelerators;
            Hostname = hostname;
            Id = id;
            InstanceId = instanceId;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            MachineType = machineType;
            Metadata = metadata;
            MetadataFingerprint = metadataFingerprint;
            MetadataStartupScript = metadataStartupScript;
            MinCpuPlatform = minCpuPlatform;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            NetworkPerformanceConfigs = networkPerformanceConfigs;
            Params = @params;
            Project = project;
            ReservationAffinities = reservationAffinities;
            ResourcePolicies = resourcePolicies;
            Schedulings = schedulings;
            ScratchDisks = scratchDisks;
            SelfLink = selfLink;
            ServiceAccounts = serviceAccounts;
            ShieldedInstanceConfigs = shieldedInstanceConfigs;
            Tags = tags;
            TagsFingerprint = tagsFingerprint;
            Zone = zone;
        }
    }
}
