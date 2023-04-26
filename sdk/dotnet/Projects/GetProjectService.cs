// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    public static class GetProjectService
    {
        /// <summary>
        /// Verify the API service for the Google Cloud Platform project to see if it is enabled or not.
        /// 
        /// For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
        /// or run `gcloud services list --available`.
        /// 
        /// This datasource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
        /// to use.
        /// 
        /// 
        /// To get more information about `gcp.projects.Service`, see:
        /// 
        /// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
        /// * How-to Guides
        ///     * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)
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
        ///     var my_project_service = Gcp.Projects.GetProjectService.Invoke(new()
        ///     {
        ///         Service = "my-project-service",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectServiceResult> InvokeAsync(GetProjectServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectServiceResult>("gcp:projects/getProjectService:getProjectService", args ?? new GetProjectServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Verify the API service for the Google Cloud Platform project to see if it is enabled or not.
        /// 
        /// For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
        /// or run `gcloud services list --available`.
        /// 
        /// This datasource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
        /// to use.
        /// 
        /// 
        /// To get more information about `gcp.projects.Service`, see:
        /// 
        /// * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
        /// * How-to Guides
        ///     * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)
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
        ///     var my_project_service = Gcp.Projects.GetProjectService.Invoke(new()
        ///     {
        ///         Service = "my-project-service",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectServiceResult> Invoke(GetProjectServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectServiceResult>("gcp:projects/getProjectService:getProjectService", args ?? new GetProjectServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The name of the Google Platform project service.
        /// </summary>
        [Input("service", required: true)]
        public string Service { get; set; } = null!;

        public GetProjectServiceArgs()
        {
        }
        public static new GetProjectServiceArgs Empty => new GetProjectServiceArgs();
    }

    public sealed class GetProjectServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the Google Platform project service.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public GetProjectServiceInvokeArgs()
        {
        }
        public static new GetProjectServiceInvokeArgs Empty => new GetProjectServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectServiceResult
    {
        public readonly bool DisableDependentServices;
        public readonly bool DisableOnDestroy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Project;
        public readonly string Service;

        [OutputConstructor]
        private GetProjectServiceResult(
            bool disableDependentServices,

            bool disableOnDestroy,

            string id,

            string? project,

            string service)
        {
            DisableDependentServices = disableDependentServices;
            DisableOnDestroy = disableOnDestroy;
            Id = id;
            Project = project;
            Service = service;
        }
    }
}
