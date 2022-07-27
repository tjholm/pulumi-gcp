// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    /// <summary>
    /// Allows management of a single API service for a Google Cloud Platform project.
    /// 
    /// For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
    /// or run `gcloud services list --available`.
    /// 
    /// This resource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
    /// to use.
    /// 
    /// To get more information about `gcp.projects.Service`, see:
    /// 
    /// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
    /// * How-to Guides
    ///     * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)
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
    ///         var project = new Gcp.Projects.Service("project", new Gcp.Projects.ServiceArgs
    ///         {
    ///             DisableDependentServices = true,
    ///             Project = "your-project-id",
    ///             ServiceName = "iam.googleapis.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Project services can be imported using the `project_id` and `service`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/service:Service my_project your-project-id/iam.googleapis.com
    /// ```
    /// 
    ///  Note that unlike other resources that fail if they already exist, `terraform apply` can be successfully used to verify already enabled services. This means that when importing existing resources into Terraform, you can either import the `google_project_service` resources or treat them as new infrastructure and run `terraform apply` to add them to state.
    /// </summary>
    [GcpResourceType("gcp:projects/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// If `true`, services that are enabled
        /// and which depend on this service should also be disabled when this service is
        /// destroyed. If `false` or unset, an error will be generated if any enabled
        /// services depend on this service when destroying it.
        /// </summary>
        [Output("disableDependentServices")]
        public Output<bool?> DisableDependentServices { get; private set; } = null!;

        /// <summary>
        /// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        /// </summary>
        [Output("disableOnDestroy")]
        public Output<bool?> DisableOnDestroy { get; private set; } = null!;

        /// <summary>
        /// The project ID. If not provided, the provider project
        /// is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The service to enable.
        /// </summary>
        [Output("service")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `true`, services that are enabled
        /// and which depend on this service should also be disabled when this service is
        /// destroyed. If `false` or unset, an error will be generated if any enabled
        /// services depend on this service when destroying it.
        /// </summary>
        [Input("disableDependentServices")]
        public Input<bool>? DisableDependentServices { get; set; }

        /// <summary>
        /// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        /// </summary>
        [Input("disableOnDestroy")]
        public Input<bool>? DisableOnDestroy { get; set; }

        /// <summary>
        /// The project ID. If not provided, the provider project
        /// is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service to enable.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If `true`, services that are enabled
        /// and which depend on this service should also be disabled when this service is
        /// destroyed. If `false` or unset, an error will be generated if any enabled
        /// services depend on this service when destroying it.
        /// </summary>
        [Input("disableDependentServices")]
        public Input<bool>? DisableDependentServices { get; set; }

        /// <summary>
        /// If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        /// </summary>
        [Input("disableOnDestroy")]
        public Input<bool>? DisableOnDestroy { get; set; }

        /// <summary>
        /// The project ID. If not provided, the provider project
        /// is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service to enable.
        /// </summary>
        [Input("service")]
        public Input<string>? ServiceName { get; set; }

        public ServiceState()
        {
        }
    }
}
