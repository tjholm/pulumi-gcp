// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// ## Example Usage
    /// ### Bigquery Dataset Access Basic User
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var bqowner = new Gcp.ServiceAccount.Account("bqowner", new()
    ///     {
    ///         AccountId = "bqowner",
    ///     });
    /// 
    ///     var access = new Gcp.BigQuery.DatasetAccess("access", new()
    ///     {
    ///         DatasetId = dataset.DatasetId,
    ///         Role = "OWNER",
    ///         UserByEmail = bqowner.Email,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Access View
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @private = new Gcp.BigQuery.Dataset("private", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var publicDataset = new Gcp.BigQuery.Dataset("publicDataset", new()
    ///     {
    ///         DatasetId = "example_dataset2",
    ///     });
    /// 
    ///     var publicTable = new Gcp.BigQuery.Table("publicTable", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = publicDataset.DatasetId,
    ///         TableId = "example_table",
    ///         View = new Gcp.BigQuery.Inputs.TableViewArgs
    ///         {
    ///             Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///             UseLegacySql = false,
    ///         },
    ///     });
    /// 
    ///     var access = new Gcp.BigQuery.DatasetAccess("access", new()
    ///     {
    ///         DatasetId = @private.DatasetId,
    ///         View = new Gcp.BigQuery.Inputs.DatasetAccessViewArgs
    ///         {
    ///             ProjectId = publicTable.Project,
    ///             DatasetId = publicDataset.DatasetId,
    ///             TableId = publicTable.TableId,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Access Authorized Dataset
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @private = new Gcp.BigQuery.Dataset("private", new()
    ///     {
    ///         DatasetId = "private",
    ///     });
    /// 
    ///     var @public = new Gcp.BigQuery.Dataset("public", new()
    ///     {
    ///         DatasetId = "public",
    ///     });
    /// 
    ///     var access = new Gcp.BigQuery.DatasetAccess("access", new()
    ///     {
    ///         DatasetId = @private.DatasetId,
    ///         AuthorizedDataset = new Gcp.BigQuery.Inputs.DatasetAccessAuthorizedDatasetArgs
    ///         {
    ///             Dataset = new Gcp.BigQuery.Inputs.DatasetAccessAuthorizedDatasetDatasetArgs
    ///             {
    ///                 ProjectId = @public.Project,
    ///                 DatasetId = @public.DatasetId,
    ///             },
    ///             TargetTypes = new[]
    ///             {
    ///                 "VIEWS",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Access Authorized Routine
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var publicDataset = new Gcp.BigQuery.Dataset("publicDataset", new()
    ///     {
    ///         DatasetId = "public_dataset",
    ///         Description = "This dataset is public",
    ///     });
    /// 
    ///     var publicRoutine = new Gcp.BigQuery.Routine("publicRoutine", new()
    ///     {
    ///         DatasetId = publicDataset.DatasetId,
    ///         RoutineId = "public_routine",
    ///         RoutineType = "TABLE_VALUED_FUNCTION",
    ///         Language = "SQL",
    ///         DefinitionBody = @"SELECT 1 + value AS value
    /// ",
    ///         Arguments = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.RoutineArgumentArgs
    ///             {
    ///                 Name = "value",
    ///                 ArgumentKind = "FIXED_TYPE",
    ///                 DataType = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["typeKind"] = "INT64",
    ///                 }),
    ///             },
    ///         },
    ///         ReturnTableType = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["columns"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "value",
    ///                     ["type"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["typeKind"] = "INT64",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var @private = new Gcp.BigQuery.Dataset("private", new()
    ///     {
    ///         DatasetId = "private_dataset",
    ///         Description = "This dataset is private",
    ///     });
    /// 
    ///     var authorizedRoutine = new Gcp.BigQuery.DatasetAccess("authorizedRoutine", new()
    ///     {
    ///         DatasetId = @private.DatasetId,
    ///         Routine = new Gcp.BigQuery.Inputs.DatasetAccessRoutineArgs
    ///         {
    ///             ProjectId = publicRoutine.Project,
    ///             DatasetId = publicRoutine.DatasetId,
    ///             RoutineId = publicRoutine.RoutineId,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:bigquery/datasetAccess:DatasetAccess")]
    public partial class DatasetAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
        /// stored in state as a different member type
        /// </summary>
        [Output("apiUpdatedMember")]
        public Output<bool> ApiUpdatedMember { get; private set; } = null!;

        /// <summary>
        /// Grants all resources of particular types in a particular dataset read access to the current dataset.
        /// Structure is documented below.
        /// </summary>
        [Output("authorizedDataset")]
        public Output<Outputs.DatasetAccessAuthorizedDataset?> AuthorizedDataset { get; private set; } = null!;

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Output("groupByEmail")]
        public Output<string?> GroupByEmail { get; private set; } = null!;

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Output("iamMember")]
        public Output<string?> IamMember { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Basic, predefined, and custom roles are
        /// supported. Predefined roles that have equivalent basic roles are
        /// swapped by the API to their basic counterparts, and will show a diff
        /// post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// A routine from a different dataset to grant access to. Queries
        /// executed against that routine will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that routine is updated by any user, access to the routine
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Output("routine")]
        public Output<Outputs.DatasetAccessRoutine?> Routine { get; private set; } = null!;

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Output("specialGroup")]
        public Output<string?> SpecialGroup { get; private set; } = null!;

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Output("userByEmail")]
        public Output<string?> UserByEmail { get; private set; } = null!;

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Output("view")]
        public Output<Outputs.DatasetAccessView?> View { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetAccess(string name, DatasetAccessArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetAccess:DatasetAccess", name, args ?? new DatasetAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetAccess(string name, Input<string> id, DatasetAccessState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetAccess:DatasetAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetAccess Get(string name, Input<string> id, DatasetAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetAccess(name, id, state, options);
        }
    }

    public sealed class DatasetAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Grants all resources of particular types in a particular dataset read access to the current dataset.
        /// Structure is documented below.
        /// </summary>
        [Input("authorizedDataset")]
        public Input<Inputs.DatasetAccessAuthorizedDatasetArgs>? AuthorizedDataset { get; set; }

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Input("groupByEmail")]
        public Input<string>? GroupByEmail { get; set; }

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Input("iamMember")]
        public Input<string>? IamMember { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Basic, predefined, and custom roles are
        /// supported. Predefined roles that have equivalent basic roles are
        /// swapped by the API to their basic counterparts, and will show a diff
        /// post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// A routine from a different dataset to grant access to. Queries
        /// executed against that routine will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that routine is updated by any user, access to the routine
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("routine")]
        public Input<Inputs.DatasetAccessRoutineArgs>? Routine { get; set; }

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Input("specialGroup")]
        public Input<string>? SpecialGroup { get; set; }

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Input("userByEmail")]
        public Input<string>? UserByEmail { get; set; }

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("view")]
        public Input<Inputs.DatasetAccessViewArgs>? View { get; set; }

        public DatasetAccessArgs()
        {
        }
        public static new DatasetAccessArgs Empty => new DatasetAccessArgs();
    }

    public sealed class DatasetAccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
        /// stored in state as a different member type
        /// </summary>
        [Input("apiUpdatedMember")]
        public Input<bool>? ApiUpdatedMember { get; set; }

        /// <summary>
        /// Grants all resources of particular types in a particular dataset read access to the current dataset.
        /// Structure is documented below.
        /// </summary>
        [Input("authorizedDataset")]
        public Input<Inputs.DatasetAccessAuthorizedDatasetGetArgs>? AuthorizedDataset { get; set; }

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("datasetId")]
        public Input<string>? DatasetId { get; set; }

        /// <summary>
        /// A domain to grant access to. Any users signed in with the
        /// domain specified will be granted the specified access
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// An email address of a Google Group to grant access to.
        /// </summary>
        [Input("groupByEmail")]
        public Input<string>? GroupByEmail { get; set; }

        /// <summary>
        /// Some other type of member that appears in the IAM Policy but isn't a user,
        /// group, domain, or special group. For example: `allUsers`
        /// </summary>
        [Input("iamMember")]
        public Input<string>? IamMember { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes the rights granted to the user specified by the other
        /// member of the access object. Basic, predefined, and custom roles are
        /// supported. Predefined roles that have equivalent basic roles are
        /// swapped by the API to their basic counterparts, and will show a diff
        /// post-create. See
        /// [official docs](https://cloud.google.com/bigquery/docs/access-control).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// A routine from a different dataset to grant access to. Queries
        /// executed against that routine will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that routine is updated by any user, access to the routine
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("routine")]
        public Input<Inputs.DatasetAccessRoutineGetArgs>? Routine { get; set; }

        /// <summary>
        /// A special group to grant access to. Possible values include:
        /// </summary>
        [Input("specialGroup")]
        public Input<string>? SpecialGroup { get; set; }

        /// <summary>
        /// An email address of a user to grant access to. For example:
        /// fred@example.com
        /// </summary>
        [Input("userByEmail")]
        public Input<string>? UserByEmail { get; set; }

        /// <summary>
        /// A view from a different dataset to grant access to. Queries
        /// executed against that view will have read access to tables in
        /// this dataset. The role field is not required when this field is
        /// set. If that view is updated by any user, access to the view
        /// needs to be granted again via an update operation.
        /// Structure is documented below.
        /// </summary>
        [Input("view")]
        public Input<Inputs.DatasetAccessViewGetArgs>? View { get; set; }

        public DatasetAccessState()
        {
        }
        public static new DatasetAccessState Empty => new DatasetAccessState();
    }
}
