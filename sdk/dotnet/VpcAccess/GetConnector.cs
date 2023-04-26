// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.VpcAccess
{
    public static class GetConnector
    {
        /// <summary>
        /// Get a Serverless VPC Access connector.
        /// 
        /// To get more information about Connector, see:
        /// 
        /// * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
        /// * How-to Guides
        ///     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
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
        ///     var sample = Gcp.VpcAccess.GetConnector.Invoke(new()
        ///     {
        ///         Name = "vpc-con",
        ///     });
        /// 
        ///     var connector = new Gcp.VpcAccess.Connector("connector", new()
        ///     {
        ///         IpCidrRange = "10.8.0.0/28",
        ///         Network = "default",
        ///         Region = "us-central1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConnectorResult> InvokeAsync(GetConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectorResult>("gcp:vpcaccess/getConnector:getConnector", args ?? new GetConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Serverless VPC Access connector.
        /// 
        /// To get more information about Connector, see:
        /// 
        /// * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
        /// * How-to Guides
        ///     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
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
        ///     var sample = Gcp.VpcAccess.GetConnector.Invoke(new()
        ///     {
        ///         Name = "vpc-con",
        ///     });
        /// 
        ///     var connector = new Gcp.VpcAccess.Connector("connector", new()
        ///     {
        ///         IpCidrRange = "10.8.0.0/28",
        ///         Network = "default",
        ///         Region = "us-central1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConnectorResult> Invoke(GetConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectorResult>("gcp:vpcaccess/getConnector:getConnector", args ?? new GetConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
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

        public GetConnectorArgs()
        {
        }
        public static new GetConnectorArgs Empty => new GetConnectorArgs();
    }

    public sealed class GetConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
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

        public GetConnectorInvokeArgs()
        {
        }
        public static new GetConnectorInvokeArgs Empty => new GetConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectorResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpCidrRange;
        public readonly string MachineType;
        public readonly int MaxInstances;
        public readonly int MaxThroughput;
        public readonly int MinInstances;
        public readonly int MinThroughput;
        public readonly string Name;
        public readonly string Network;
        public readonly string? Project;
        public readonly string? Region;
        public readonly string SelfLink;
        public readonly string State;
        public readonly ImmutableArray<Outputs.GetConnectorSubnetResult> Subnets;

        [OutputConstructor]
        private GetConnectorResult(
            string id,

            string ipCidrRange,

            string machineType,

            int maxInstances,

            int maxThroughput,

            int minInstances,

            int minThroughput,

            string name,

            string network,

            string? project,

            string? region,

            string selfLink,

            string state,

            ImmutableArray<Outputs.GetConnectorSubnetResult> subnets)
        {
            Id = id;
            IpCidrRange = ipCidrRange;
            MachineType = machineType;
            MaxInstances = maxInstances;
            MaxThroughput = maxThroughput;
            MinInstances = minInstances;
            MinThroughput = minThroughput;
            Name = name;
            Network = network;
            Project = project;
            Region = region;
            SelfLink = selfLink;
            State = state;
            Subnets = subnets;
        }
    }
}
