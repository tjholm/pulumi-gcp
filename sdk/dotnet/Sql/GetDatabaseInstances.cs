// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql
{
    public static class GetDatabaseInstances
    {
        /// <summary>
        /// Use this data source to get information about a list of Cloud SQL instances in a project. You can also apply some filters over this list to get a more filtered list of Cloud SQL instances.
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
        ///     var qa = Gcp.Sql.GetDatabaseInstances.Invoke(new()
        ///     {
        ///         Project = "test-project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseInstancesResult> InvokeAsync(GetDatabaseInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseInstancesResult>("gcp:sql/getDatabaseInstances:getDatabaseInstances", args ?? new GetDatabaseInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a list of Cloud SQL instances in a project. You can also apply some filters over this list to get a more filtered list of Cloud SQL instances.
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
        ///     var qa = Gcp.Sql.GetDatabaseInstances.Invoke(new()
        ///     {
        ///         Project = "test-project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseInstancesResult> Invoke(GetDatabaseInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseInstancesResult>("gcp:sql/getDatabaseInstances:getDatabaseInstances", args ?? new GetDatabaseInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// To filter out the Cloud SQL instances which are of the specified database version.
        /// </summary>
        [Input("databaseVersion")]
        public string? DatabaseVersion { get; set; }

        /// <summary>
        /// The ID of the project in which the resources belong. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances which are located in the specified region.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances based on the current serving state of the database instance. Supported values include `SQL_INSTANCE_STATE_UNSPECIFIED`, `RUNNABLE`, `SUSPENDED`, `PENDING_DELETE`, `PENDING_CREATE`, `MAINTENANCE`, `FAILED`.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances based on the tier(or machine type) of the database instances.
        /// </summary>
        [Input("tier")]
        public string? Tier { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances which are located in the specified zone. This zone refers to the Compute Engine zone that the instance is currently serving from.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetDatabaseInstancesArgs()
        {
        }
        public static new GetDatabaseInstancesArgs Empty => new GetDatabaseInstancesArgs();
    }

    public sealed class GetDatabaseInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// To filter out the Cloud SQL instances which are of the specified database version.
        /// </summary>
        [Input("databaseVersion")]
        public Input<string>? DatabaseVersion { get; set; }

        /// <summary>
        /// The ID of the project in which the resources belong. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances which are located in the specified region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances based on the current serving state of the database instance. Supported values include `SQL_INSTANCE_STATE_UNSPECIFIED`, `RUNNABLE`, `SUSPENDED`, `PENDING_DELETE`, `PENDING_CREATE`, `MAINTENANCE`, `FAILED`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances based on the tier(or machine type) of the database instances.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        /// <summary>
        /// To filter out the Cloud SQL instances which are located in the specified zone. This zone refers to the Compute Engine zone that the instance is currently serving from.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetDatabaseInstancesInvokeArgs()
        {
        }
        public static new GetDatabaseInstancesInvokeArgs Empty => new GetDatabaseInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseInstancesResult
    {
        public readonly string? DatabaseVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDatabaseInstancesInstanceResult> Instances;
        public readonly string? Project;
        public readonly string? Region;
        public readonly string? State;
        public readonly string? Tier;
        public readonly string? Zone;

        [OutputConstructor]
        private GetDatabaseInstancesResult(
            string? databaseVersion,

            string id,

            ImmutableArray<Outputs.GetDatabaseInstancesInstanceResult> instances,

            string? project,

            string? region,

            string? state,

            string? tier,

            string? zone)
        {
            DatabaseVersion = databaseVersion;
            Id = id;
            Instances = instances;
            Project = project;
            Region = region;
            State = state;
            Tier = tier;
            Zone = zone;
        }
    }
}
