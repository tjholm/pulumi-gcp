// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Tpu
{
    public static class GetTensorflowVersions
    {
        /// <summary>
        /// Get TensorFlow versions available for a project. For more information see the [official documentation](https://cloud.google.com/tpu/docs/) and [API](https://cloud.google.com/tpu/docs/reference/rest/v1/projects.locations.tensorflowVersions).
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
        ///     var available = Gcp.Tpu.GetTensorflowVersions.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Configure Basic TPU Node With Available Version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var available = Gcp.Tpu.GetTensorflowVersions.Invoke();
        /// 
        ///     var tpu = new Gcp.Tpu.Node("tpu", new()
        ///     {
        ///         Zone = "us-central1-b",
        ///         AcceleratorType = "v3-8",
        ///         TensorflowVersion = available.Apply(getTensorflowVersionsResult =&gt; getTensorflowVersionsResult.Versions[0]),
        ///         CidrBlock = "10.2.0.0/29",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTensorflowVersionsResult> InvokeAsync(GetTensorflowVersionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTensorflowVersionsResult>("gcp:tpu/getTensorflowVersions:getTensorflowVersions", args ?? new GetTensorflowVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Get TensorFlow versions available for a project. For more information see the [official documentation](https://cloud.google.com/tpu/docs/) and [API](https://cloud.google.com/tpu/docs/reference/rest/v1/projects.locations.tensorflowVersions).
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
        ///     var available = Gcp.Tpu.GetTensorflowVersions.Invoke();
        /// 
        /// });
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Configure Basic TPU Node With Available Version
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var available = Gcp.Tpu.GetTensorflowVersions.Invoke();
        /// 
        ///     var tpu = new Gcp.Tpu.Node("tpu", new()
        ///     {
        ///         Zone = "us-central1-b",
        ///         AcceleratorType = "v3-8",
        ///         TensorflowVersion = available.Apply(getTensorflowVersionsResult =&gt; getTensorflowVersionsResult.Versions[0]),
        ///         CidrBlock = "10.2.0.0/29",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTensorflowVersionsResult> Invoke(GetTensorflowVersionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTensorflowVersionsResult>("gcp:tpu/getTensorflowVersions:getTensorflowVersions", args ?? new GetTensorflowVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTensorflowVersionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project to list versions for. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The zone to list versions for. If it
        /// is not provided, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetTensorflowVersionsArgs()
        {
        }
        public static new GetTensorflowVersionsArgs Empty => new GetTensorflowVersionsArgs();
    }

    public sealed class GetTensorflowVersionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project to list versions for. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The zone to list versions for. If it
        /// is not provided, the provider zone is used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetTensorflowVersionsInvokeArgs()
        {
        }
        public static new GetTensorflowVersionsInvokeArgs Empty => new GetTensorflowVersionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTensorflowVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Project;
        /// <summary>
        /// The list of TensorFlow versions available for the given project and zone.
        /// </summary>
        public readonly ImmutableArray<string> Versions;
        public readonly string Zone;

        [OutputConstructor]
        private GetTensorflowVersionsResult(
            string id,

            string project,

            ImmutableArray<string> versions,

            string zone)
        {
            Id = id;
            Project = project;
            Versions = versions;
            Zone = zone;
        }
    }
}
