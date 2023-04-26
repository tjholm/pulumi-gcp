// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetSnapshot
    {
        /// <summary>
        /// To get more information about Snapshot, see:
        /// 
        /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
        /// * How-to Guides
        ///     * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
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
        ///     var snapshot = Gcp.Compute.GetSnapshot.Invoke(new()
        ///     {
        ///         Name = "my-snapshot",
        ///     });
        /// 
        ///     var latest_snapshot = Gcp.Compute.GetSnapshot.Invoke(new()
        ///     {
        ///         Filter = "name != my-snapshot",
        ///         MostRecent = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("gcp:compute/getSnapshot:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// To get more information about Snapshot, see:
        /// 
        /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
        /// * How-to Guides
        ///     * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
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
        ///     var snapshot = Gcp.Compute.GetSnapshot.Invoke(new()
        ///     {
        ///         Name = "my-snapshot",
        ///     });
        /// 
        ///     var latest_snapshot = Gcp.Compute.GetSnapshot.Invoke(new()
        ///     {
        ///         Filter = "name != my-snapshot",
        ///         MostRecent = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("gcp:compute/getSnapshot:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to retrieve the compute snapshot.
        /// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
        /// If multiple compute snapshot match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// If `filter` is provided, ensures the most recent snapshot is returned when multiple compute snapshot match.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The name of the compute snapshot. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to retrieve the compute snapshot.
        /// See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
        /// If multiple compute snapshot match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// If `filter` is provided, ensures the most recent snapshot is returned when multiple compute snapshot match.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        /// <summary>
        /// The name of the compute snapshot. One of `name` or `filter` must be provided.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        public readonly string ChainName;
        public readonly string CreationTimestamp;
        public readonly string Description;
        public readonly int DiskSizeGb;
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LabelFingerprint;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<string> Licenses;
        public readonly bool? MostRecent;
        public readonly string? Name;
        public readonly string? Project;
        public readonly string SelfLink;
        public readonly ImmutableArray<Outputs.GetSnapshotSnapshotEncryptionKeyResult> SnapshotEncryptionKeys;
        public readonly int SnapshotId;
        public readonly string SourceDisk;
        public readonly ImmutableArray<Outputs.GetSnapshotSourceDiskEncryptionKeyResult> SourceDiskEncryptionKeys;
        public readonly int StorageBytes;
        public readonly ImmutableArray<string> StorageLocations;
        public readonly string Zone;

        [OutputConstructor]
        private GetSnapshotResult(
            string chainName,

            string creationTimestamp,

            string description,

            int diskSizeGb,

            string? filter,

            string id,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> licenses,

            bool? mostRecent,

            string? name,

            string? project,

            string selfLink,

            ImmutableArray<Outputs.GetSnapshotSnapshotEncryptionKeyResult> snapshotEncryptionKeys,

            int snapshotId,

            string sourceDisk,

            ImmutableArray<Outputs.GetSnapshotSourceDiskEncryptionKeyResult> sourceDiskEncryptionKeys,

            int storageBytes,

            ImmutableArray<string> storageLocations,

            string zone)
        {
            ChainName = chainName;
            CreationTimestamp = creationTimestamp;
            Description = description;
            DiskSizeGb = diskSizeGb;
            Filter = filter;
            Id = id;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Licenses = licenses;
            MostRecent = mostRecent;
            Name = name;
            Project = project;
            SelfLink = selfLink;
            SnapshotEncryptionKeys = snapshotEncryptionKeys;
            SnapshotId = snapshotId;
            SourceDisk = sourceDisk;
            SourceDiskEncryptionKeys = sourceDiskEncryptionKeys;
            StorageBytes = storageBytes;
            StorageLocations = storageLocations;
            Zone = zone;
        }
    }
}
