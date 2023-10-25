// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Spanner
{
    public static class GetDatabaseIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a Spanner database.
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Gcp.Spanner.GetDatabaseIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_spanner_database.Database.Project,
        ///         Database = google_spanner_database.Database.Name,
        ///         Instance = google_spanner_database.Database.Instance,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabaseIamPolicyResult> InvokeAsync(GetDatabaseIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseIamPolicyResult>("gcp:spanner/getDatabaseIamPolicy:getDatabaseIamPolicy", args ?? new GetDatabaseIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a Spanner database.
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Gcp.Spanner.GetDatabaseIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_spanner_database.Database.Project,
        ///         Database = google_spanner_database.Database.Name,
        ///         Instance = google_spanner_database.Database.Instance,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseIamPolicyResult> Invoke(GetDatabaseIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseIamPolicyResult>("gcp:spanner/getDatabaseIamPolicy:getDatabaseIamPolicy", args ?? new GetDatabaseIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
        /// </summary>
        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDatabaseIamPolicyArgs()
        {
        }
        public static new GetDatabaseIamPolicyArgs Empty => new GetDatabaseIamPolicyArgs();
    }

    public sealed class GetDatabaseIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatabaseIamPolicyInvokeArgs()
        {
        }
        public static new GetDatabaseIamPolicyInvokeArgs Empty => new GetDatabaseIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseIamPolicyResult
    {
        public readonly string Database;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Instance;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;

        [OutputConstructor]
        private GetDatabaseIamPolicyResult(
            string database,

            string etag,

            string id,

            string instance,

            string policyData,

            string project)
        {
            Database = database;
            Etag = etag;
            Id = id;
            Instance = instance;
            PolicyData = policyData;
            Project = project;
        }
    }
}
