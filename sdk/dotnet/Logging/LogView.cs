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
    /// Describes a view over log entries in a bucket.
    /// 
    /// To get more information about LogView, see:
    /// 
    /// * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.locations.buckets.views)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/logging/docs/apis)
    /// 
    /// ## Example Usage
    /// ### Logging Log View Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var loggingLogViewProjectBucketConfig = new Gcp.Logging.ProjectBucketConfig("loggingLogViewProjectBucketConfig", new()
    ///     {
    ///         Project = "my-project-name",
    ///         Location = "global",
    ///         RetentionDays = 30,
    ///         BucketId = "_Default",
    ///     });
    /// 
    ///     var loggingLogViewLogView = new Gcp.Logging.LogView("loggingLogViewLogView", new()
    ///     {
    ///         Bucket = loggingLogViewProjectBucketConfig.Id,
    ///         Description = "A logging view configured with Terraform",
    ///         Filter = "SOURCE(\"projects/myproject\") AND resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LogView can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/logView:LogView default {{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/logView:LogView")]
    public partial class LogView : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket of the resource
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Output only. The creation timestamp of the view.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Describes this view.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the view. For example: \`projects/my-project/locations/global/buckets/my-bucket/views/my-view\`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of the resource.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Output only. The last update timestamp of the view.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a LogView resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogView(string name, LogViewArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/logView:LogView", name, args ?? new LogViewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogView(string name, Input<string> id, LogViewState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/logView:LogView", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogView resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogView Get(string name, Input<string> id, LogViewState? state = null, CustomResourceOptions? options = null)
        {
            return new LogView(name, id, state, options);
        }
    }

    public sealed class LogViewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket of the resource
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Describes this view.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the view. For example: \`projects/my-project/locations/global/buckets/my-bucket/views/my-view\`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the resource.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        public LogViewArgs()
        {
        }
        public static new LogViewArgs Empty => new LogViewArgs();
    }

    public sealed class LogViewState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket of the resource
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Output only. The creation timestamp of the view.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Describes this view.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the view. For example: \`projects/my-project/locations/global/buckets/my-bucket/views/my-view\`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the resource.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Output only. The last update timestamp of the view.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public LogViewState()
        {
        }
        public static new LogViewState Empty => new LogViewState();
    }
}
