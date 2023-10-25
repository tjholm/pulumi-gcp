// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetRegionNetworkEndpointGroup
    {
        /// <summary>
        /// Use this data source to access a Region Network Endpoint Group's attributes.
        /// 
        /// The RNEG may be found by providing either a `self_link`, or a `name` and a `region`.
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
        ///     var rneg1 = Gcp.Compute.GetRegionNetworkEndpointGroup.Invoke(new()
        ///     {
        ///         Name = "k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///         Region = "us-central1",
        ///     });
        /// 
        ///     var rneg2 = Gcp.Compute.GetRegionNetworkEndpointGroup.Invoke(new()
        ///     {
        ///         SelfLink = "https://www.googleapis.com/compute/v1/projects/myproject/regions/us-central1/networkEndpointGroups/k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegionNetworkEndpointGroupResult> InvokeAsync(GetRegionNetworkEndpointGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionNetworkEndpointGroupResult>("gcp:compute/getRegionNetworkEndpointGroup:getRegionNetworkEndpointGroup", args ?? new GetRegionNetworkEndpointGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access a Region Network Endpoint Group's attributes.
        /// 
        /// The RNEG may be found by providing either a `self_link`, or a `name` and a `region`.
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
        ///     var rneg1 = Gcp.Compute.GetRegionNetworkEndpointGroup.Invoke(new()
        ///     {
        ///         Name = "k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///         Region = "us-central1",
        ///     });
        /// 
        ///     var rneg2 = Gcp.Compute.GetRegionNetworkEndpointGroup.Invoke(new()
        ///     {
        ///         SelfLink = "https://www.googleapis.com/compute/v1/projects/myproject/regions/us-central1/networkEndpointGroups/k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegionNetworkEndpointGroupResult> Invoke(GetRegionNetworkEndpointGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionNetworkEndpointGroupResult>("gcp:compute/getRegionNetworkEndpointGroup:getRegionNetworkEndpointGroup", args ?? new GetRegionNetworkEndpointGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionNetworkEndpointGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Network Endpoint Group name. Provide either this or a `self_link`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project to list versions in. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// A reference to the region where the Serverless REGs Reside. Provide either this or a `self_link`.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The Network Endpoint Group self\_link.
        /// </summary>
        [Input("selfLink")]
        public string? SelfLink { get; set; }

        public GetRegionNetworkEndpointGroupArgs()
        {
        }
        public static new GetRegionNetworkEndpointGroupArgs Empty => new GetRegionNetworkEndpointGroupArgs();
    }

    public sealed class GetRegionNetworkEndpointGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Network Endpoint Group name. Provide either this or a `self_link`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project to list versions in. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region where the Serverless REGs Reside. Provide either this or a `self_link`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The Network Endpoint Group self\_link.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public GetRegionNetworkEndpointGroupInvokeArgs()
        {
        }
        public static new GetRegionNetworkEndpointGroupInvokeArgs Empty => new GetRegionNetworkEndpointGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionNetworkEndpointGroupResult
    {
        public readonly ImmutableArray<Outputs.GetRegionNetworkEndpointGroupAppEngineResult> AppEngines;
        public readonly ImmutableArray<Outputs.GetRegionNetworkEndpointGroupCloudFunctionResult> CloudFunctions;
        public readonly ImmutableArray<Outputs.GetRegionNetworkEndpointGroupCloudRunResult> CloudRuns;
        /// <summary>
        /// The RNEG description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The network to which all network endpoints in the RNEG belong.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Type of network endpoints in this network endpoint group.
        /// </summary>
        public readonly string NetworkEndpointType;
        public readonly string? Project;
        /// <summary>
        /// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment.
        /// </summary>
        public readonly string PscTargetService;
        public readonly string? Region;
        public readonly string? SelfLink;
        public readonly ImmutableArray<Outputs.GetRegionNetworkEndpointGroupServerlessDeploymentResult> ServerlessDeployments;
        /// <summary>
        /// subnetwork to which all network endpoints in the RNEG belong.
        /// </summary>
        public readonly string Subnetwork;

        [OutputConstructor]
        private GetRegionNetworkEndpointGroupResult(
            ImmutableArray<Outputs.GetRegionNetworkEndpointGroupAppEngineResult> appEngines,

            ImmutableArray<Outputs.GetRegionNetworkEndpointGroupCloudFunctionResult> cloudFunctions,

            ImmutableArray<Outputs.GetRegionNetworkEndpointGroupCloudRunResult> cloudRuns,

            string description,

            string id,

            string? name,

            string network,

            string networkEndpointType,

            string? project,

            string pscTargetService,

            string? region,

            string? selfLink,

            ImmutableArray<Outputs.GetRegionNetworkEndpointGroupServerlessDeploymentResult> serverlessDeployments,

            string subnetwork)
        {
            AppEngines = appEngines;
            CloudFunctions = cloudFunctions;
            CloudRuns = cloudRuns;
            Description = description;
            Id = id;
            Name = name;
            Network = network;
            NetworkEndpointType = networkEndpointType;
            Project = project;
            PscTargetService = pscTargetService;
            Region = region;
            SelfLink = selfLink;
            ServerlessDeployments = serverlessDeployments;
            Subnetwork = subnetwork;
        }
    }
}
