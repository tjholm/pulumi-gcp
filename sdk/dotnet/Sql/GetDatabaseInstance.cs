// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    public static class GetDatabaseInstance
    {
        /// <summary>
        /// Use this data source to get information about a Cloud SQL instance.
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
        ///     var qa = Gcp.Sql.GetDatabaseInstance.Invoke(new()
        ///     {
        ///         Name = "test-sql-instance",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseInstanceResult> InvokeAsync(GetDatabaseInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseInstanceResult>("gcp:sql/getDatabaseInstance:getDatabaseInstance", args ?? new GetDatabaseInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a Cloud SQL instance.
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
        ///     var qa = Gcp.Sql.GetDatabaseInstance.Invoke(new()
        ///     {
        ///         Name = "test-sql-instance",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseInstanceResult> Invoke(GetDatabaseInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseInstanceResult>("gcp:sql/getDatabaseInstance:getDatabaseInstance", args ?? new GetDatabaseInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the instance.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDatabaseInstanceArgs()
        {
        }
        public static new GetDatabaseInstanceArgs Empty => new GetDatabaseInstanceArgs();
    }

    public sealed class GetDatabaseInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the instance.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatabaseInstanceInvokeArgs()
        {
        }
        public static new GetDatabaseInstanceInvokeArgs Empty => new GetDatabaseInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseInstanceResult
    {
        public readonly ImmutableArray<string> AvailableMaintenanceVersions;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceCloneResult> Clones;
        public readonly string ConnectionName;
        public readonly string DatabaseVersion;
        public readonly bool DeletionProtection;
        public readonly string EncryptionKeyName;
        public readonly string FirstIpAddress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceType;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceIpAddressResult> IpAddresses;
        public readonly string MaintenanceVersion;
        public readonly string MasterInstanceName;
        public readonly string Name;
        public readonly string PrivateIpAddress;
        public readonly string? Project;
        public readonly string PublicIpAddress;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceReplicaConfigurationResult> ReplicaConfigurations;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceRestoreBackupContextResult> RestoreBackupContexts;
        public readonly string RootPassword;
        public readonly string SelfLink;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceServerCaCertResult> ServerCaCerts;
        public readonly string ServiceAccountEmailAddress;
        public readonly ImmutableArray<Outputs.GetDatabaseInstanceSettingResult> Settings;

        [OutputConstructor]
        private GetDatabaseInstanceResult(
            ImmutableArray<string> availableMaintenanceVersions,

            ImmutableArray<Outputs.GetDatabaseInstanceCloneResult> clones,

            string connectionName,

            string databaseVersion,

            bool deletionProtection,

            string encryptionKeyName,

            string firstIpAddress,

            string id,

            string instanceType,

            ImmutableArray<Outputs.GetDatabaseInstanceIpAddressResult> ipAddresses,

            string maintenanceVersion,

            string masterInstanceName,

            string name,

            string privateIpAddress,

            string? project,

            string publicIpAddress,

            string region,

            ImmutableArray<Outputs.GetDatabaseInstanceReplicaConfigurationResult> replicaConfigurations,

            ImmutableArray<Outputs.GetDatabaseInstanceRestoreBackupContextResult> restoreBackupContexts,

            string rootPassword,

            string selfLink,

            ImmutableArray<Outputs.GetDatabaseInstanceServerCaCertResult> serverCaCerts,

            string serviceAccountEmailAddress,

            ImmutableArray<Outputs.GetDatabaseInstanceSettingResult> settings)
        {
            AvailableMaintenanceVersions = availableMaintenanceVersions;
            Clones = clones;
            ConnectionName = connectionName;
            DatabaseVersion = databaseVersion;
            DeletionProtection = deletionProtection;
            EncryptionKeyName = encryptionKeyName;
            FirstIpAddress = firstIpAddress;
            Id = id;
            InstanceType = instanceType;
            IpAddresses = ipAddresses;
            MaintenanceVersion = maintenanceVersion;
            MasterInstanceName = masterInstanceName;
            Name = name;
            PrivateIpAddress = privateIpAddress;
            Project = project;
            PublicIpAddress = publicIpAddress;
            Region = region;
            ReplicaConfigurations = replicaConfigurations;
            RestoreBackupContexts = restoreBackupContexts;
            RootPassword = rootPassword;
            SelfLink = selfLink;
            ServerCaCerts = serverCaCerts;
            ServiceAccountEmailAddress = serviceAccountEmailAddress;
            Settings = settings;
        }
    }
}
