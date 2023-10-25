// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Beyondcorp
{
    public static class GetAppConnector
    {
        /// <summary>
        /// Get information about a Google BeyondCorp App Connector.
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
        ///     var my_beyondcorp_app_connector = Gcp.Beyondcorp.GetAppConnector.Invoke(new()
        ///     {
        ///         Name = "my-beyondcorp-app-connector",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppConnectorResult> InvokeAsync(GetAppConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppConnectorResult>("gcp:beyondcorp/getAppConnector:getAppConnector", args ?? new GetAppConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Google BeyondCorp App Connector.
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
        ///     var my_beyondcorp_app_connector = Gcp.Beyondcorp.GetAppConnector.Invoke(new()
        ///     {
        ///         Name = "my-beyondcorp-app-connector",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppConnectorResult> Invoke(GetAppConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppConnectorResult>("gcp:beyondcorp/getAppConnector:getAppConnector", args ?? new GetAppConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App Connector.
        /// 
        /// - - -
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetAppConnectorArgs()
        {
        }
        public static new GetAppConnectorArgs Empty => new GetAppConnectorArgs();
    }

    public sealed class GetAppConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the App Connector.
        /// 
        /// - - -
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetAppConnectorInvokeArgs()
        {
        }
        public static new GetAppConnectorInvokeArgs Empty => new GetAppConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppConnectorResult
    {
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetAppConnectorPrincipalInfoResult> PrincipalInfos;
        public readonly string? Project;
        public readonly string? Region;
        public readonly string State;

        [OutputConstructor]
        private GetAppConnectorResult(
            string displayName,

            string id,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.GetAppConnectorPrincipalInfoResult> principalInfos,

            string? project,

            string? region,

            string state)
        {
            DisplayName = displayName;
            Id = id;
            Labels = labels;
            Name = name;
            PrincipalInfos = principalInfos;
            Project = project;
            Region = region;
            State = state;
        }
    }
}
