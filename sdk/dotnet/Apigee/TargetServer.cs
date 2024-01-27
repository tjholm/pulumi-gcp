// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// TargetServer configuration. TargetServers are used to decouple a proxy TargetEndpoint HTTPTargetConnections from concrete URLs for backend services.
    /// 
    /// To get more information about TargetServer, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.targetservers/create)
    /// * How-to Guides
    ///     * [Load balancing across backend servers](https://cloud.google.com/apigee/docs/api-platform/deploy/load-balancing-across-backend-servers)
    /// 
    /// ## Example Usage
    /// ### Apigee Target Server Test Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Gcp.Organizations.Project("project", new()
    ///     {
    ///         OrgId = "123456789",
    ///         BillingAccount = "000000-0000000-0000000-000000",
    ///     });
    /// 
    ///     var apigee = new Gcp.Projects.Service("apigee", new()
    ///     {
    ///         Project = project.ProjectId,
    ///         ServiceName = "apigee.googleapis.com",
    ///     });
    /// 
    ///     var servicenetworking = new Gcp.Projects.Service("servicenetworking", new()
    ///     {
    ///         Project = project.ProjectId,
    ///         ServiceName = "servicenetworking.googleapis.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigee,
    ///         },
    ///     });
    /// 
    ///     var compute = new Gcp.Projects.Service("compute", new()
    ///     {
    ///         Project = project.ProjectId,
    ///         ServiceName = "compute.googleapis.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             servicenetworking,
    ///         },
    ///     });
    /// 
    ///     var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork", new()
    ///     {
    ///         Project = project.ProjectId,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             compute,
    ///         },
    ///     });
    /// 
    ///     var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 16,
    ///         Network = apigeeNetwork.Id,
    ///         Project = project.ProjectId,
    ///     });
    /// 
    ///     var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new()
    ///     {
    ///         Network = apigeeNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             apigeeRange.Name,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             servicenetworking,
    ///         },
    ///     });
    /// 
    ///     var apigeeOrg = new Gcp.Apigee.Organization("apigeeOrg", new()
    ///     {
    ///         AnalyticsRegion = "us-central1",
    ///         ProjectId = project.ProjectId,
    ///         AuthorizedNetwork = apigeeNetwork.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigeeVpcConnection,
    ///             apigee,
    ///         },
    ///     });
    /// 
    ///     var apigeeEnvironment = new Gcp.Apigee.Environment("apigeeEnvironment", new()
    ///     {
    ///         OrgId = apigeeOrg.Id,
    ///         Description = "Apigee Environment",
    ///         DisplayName = "environment-1",
    ///     });
    /// 
    ///     var apigeeTargetServer = new Gcp.Apigee.TargetServer("apigeeTargetServer", new()
    ///     {
    ///         Description = "Apigee Target Server",
    ///         Protocol = "HTTP",
    ///         Host = "abc.foo.com",
    ///         Port = 8080,
    ///         EnvId = apigeeEnvironment.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TargetServer can be imported using any of these accepted formats* `{{env_id}}/targetservers/{{name}}` * `{{env_id}}/{{name}}` When using the `pulumi import` command, TargetServer can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/targetservers/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/targetServer:TargetServer")]
    public partial class TargetServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description of this TargetServer.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/environments/{{env_name}}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("envId")]
        public Output<string> EnvId { get; private set; } = null!;

        /// <summary>
        /// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool?> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Immutable. The protocol used by this TargetServer.
        /// Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("sSlInfo")]
        public Output<Outputs.TargetServerSSlInfo?> SSlInfo { get; private set; } = null!;


        /// <summary>
        /// Create a TargetServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetServer(string name, TargetServerArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/targetServer:TargetServer", name, args ?? new TargetServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetServer(string name, Input<string> id, TargetServerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/targetServer:TargetServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetServer Get(string name, Input<string> id, TargetServerState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetServer(name, id, state, options);
        }
    }

    public sealed class TargetServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of this TargetServer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/environments/{{env_name}}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("envId", required: true)]
        public Input<string> EnvId { get; set; } = null!;

        /// <summary>
        /// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Immutable. The protocol used by this TargetServer.
        /// Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("sSlInfo")]
        public Input<Inputs.TargetServerSSlInfoArgs>? SSlInfo { get; set; }

        public TargetServerArgs()
        {
        }
        public static new TargetServerArgs Empty => new TargetServerArgs();
    }

    public sealed class TargetServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of this TargetServer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Apigee environment group associated with the Apigee environment,
        /// in the format `organizations/{{org_name}}/environments/{{env_name}}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("envId")]
        public Input<string>? EnvId { get; set; }

        /// <summary>
        /// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Immutable. The protocol used by this TargetServer.
        /// Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("sSlInfo")]
        public Input<Inputs.TargetServerSSlInfoGetArgs>? SSlInfo { get; set; }

        public TargetServerState()
        {
        }
        public static new TargetServerState Empty => new TargetServerState();
    }
}
