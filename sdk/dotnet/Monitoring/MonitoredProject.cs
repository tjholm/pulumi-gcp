// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// A [project being monitored](https://cloud.google.com/monitoring/settings/multiple-projects#create-multi) by a Metrics Scope.
    /// 
    /// To get more information about MonitoredProject, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/locations.global.metricsScopes.projects)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/settings/manage-api)
    /// 
    /// ## Example Usage
    /// ### Monitoring Monitored Project Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.Monitoring.MonitoredProject("primary", new()
    ///     {
    ///         MetricsScope = "my-project-name",
    ///     });
    /// 
    ///     var basic = new Gcp.Organizations.Project("basic", new()
    ///     {
    ///         OrgId = "123456789",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MonitoredProject can be imported using any of these accepted formats* `v1/locations/global/metricsScopes/{{name}}` * `{{name}}` When using the `pulumi import` command, MonitoredProject can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:monitoring/monitoredProject:MonitoredProject default v1/locations/global/metricsScopes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:monitoring/monitoredProject:MonitoredProject default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:monitoring/monitoredProject:MonitoredProject")]
    public partial class MonitoredProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. The time when this `MonitoredProject` was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("metricsScope")]
        public Output<string> MetricsScope { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MonitoredProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MonitoredProject(string name, MonitoredProjectArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/monitoredProject:MonitoredProject", name, args ?? new MonitoredProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MonitoredProject(string name, Input<string> id, MonitoredProjectState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/monitoredProject:MonitoredProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MonitoredProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MonitoredProject Get(string name, Input<string> id, MonitoredProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new MonitoredProject(name, id, state, options);
        }
    }

    public sealed class MonitoredProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("metricsScope", required: true)]
        public Input<string> MetricsScope { get; set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MonitoredProjectArgs()
        {
        }
        public static new MonitoredProjectArgs Empty => new MonitoredProjectArgs();
    }

    public sealed class MonitoredProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The time when this `MonitoredProject` was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("metricsScope")]
        public Input<string>? MetricsScope { get; set; }

        /// <summary>
        /// Immutable. The resource name of the `MonitoredProject`. On input, the resource name includes the scoping project ID and monitored project ID. On output, it contains the equivalent project numbers. Example: `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MonitoredProjectState()
        {
        }
        public static new MonitoredProjectState Empty => new MonitoredProjectState();
    }
}
