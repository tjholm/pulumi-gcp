// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// Manages a organization-level logging sink. For more information see:
    /// * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/organizations.sinks)
    /// * How-to Guides
    ///     * [Exporting Logs](https://cloud.google.com/logging/docs/export)
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
    ///         var log_bucket = new Gcp.Storage.Bucket("log-bucket", new Gcp.Storage.BucketArgs
    ///         {
    ///             Location = "US",
    ///         });
    ///         var my_sink = new Gcp.Logging.OrganizationSink("my-sink", new Gcp.Logging.OrganizationSinkArgs
    ///         {
    ///             Description = "some explanation on what this is",
    ///             OrgId = "123456789",
    ///             Destination = log_bucket.Name.Apply(name =&gt; $"storage.googleapis.com/{name}"),
    ///             Filter = "resource.type = gce_instance AND severity &gt;= WARNING",
    ///         });
    ///         var log_writer = new Gcp.Projects.IAMMember("log-writer", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Project = "your-project-id",
    ///             Role = "roles/storage.objectCreator",
    ///             Member = my_sink.WriterIdentity,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Organization-level logging sinks can be imported using this format
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/organizationSink:OrganizationSink my_sink organizations/{{organization_id}}/sinks/{{sink_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/organizationSink:OrganizationSink")]
    public partial class OrganizationSink : Pulumi.CustomResource
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Output("bigqueryOptions")]
        public Output<Outputs.OrganizationSinkBigqueryOptions> BigqueryOptions { get; private set; } = null!;

        /// <summary>
        /// A description of this exclusion.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// If set to True, then this exclusion is disabled and it does not exclude any log entries.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        /// </summary>
        [Output("exclusions")]
        public Output<ImmutableArray<Outputs.OrganizationSinkExclusion>> Exclusions { get; private set; } = null!;

        /// <summary>
        /// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// Whether or not to include children organizations in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        /// </summary>
        [Output("includeChildren")]
        public Output<bool?> IncludeChildren { get; private set; } = null!;

        /// <summary>
        /// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the organization to be exported to the sink.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Output("writerIdentity")]
        public Output<string> WriterIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSink(string name, OrganizationSinkArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/organizationSink:OrganizationSink", name, args ?? new OrganizationSinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSink(string name, Input<string> id, OrganizationSinkState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/organizationSink:OrganizationSink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSink Get(string name, Input<string> id, OrganizationSinkState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationSink(name, id, state, options);
        }
    }

    public sealed class OrganizationSinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.OrganizationSinkBigqueryOptionsArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// A description of this exclusion.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// If set to True, then this exclusion is disabled and it does not exclude any log entries.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.OrganizationSinkExclusionArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        /// </summary>
        public InputList<Inputs.OrganizationSinkExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.OrganizationSinkExclusionArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Whether or not to include children organizations in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        /// </summary>
        [Input("includeChildren")]
        public Input<bool>? IncludeChildren { get; set; }

        /// <summary>
        /// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The numeric ID of the organization to be exported to the sink.
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public OrganizationSinkArgs()
        {
        }
    }

    public sealed class OrganizationSinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options that affect sinks exporting data to BigQuery. Structure documented below.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.OrganizationSinkBigqueryOptionsGetArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// A description of this exclusion.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination of the sink (or, in other words, where logs are written to). Can be a
        /// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// The writer associated with the sink must have access to write to the above resource.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// If set to True, then this exclusion is disabled and it does not exclude any log entries.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.OrganizationSinkExclusionGetArgs>? _exclusions;

        /// <summary>
        /// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
        /// </summary>
        public InputList<Inputs.OrganizationSinkExclusionGetArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.OrganizationSinkExclusionGetArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Whether or not to include children organizations in the sink export. If true, logs
        /// associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
        /// </summary>
        [Input("includeChildren")]
        public Input<bool>? IncludeChildren { get; set; }

        /// <summary>
        /// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The numeric ID of the organization to be exported to the sink.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The identity associated with this sink. This identity must be granted write access to the
        /// configured `destination`.
        /// </summary>
        [Input("writerIdentity")]
        public Input<string>? WriterIdentity { get; set; }

        public OrganizationSinkState()
        {
        }
    }
}
