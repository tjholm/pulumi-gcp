// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents an Autoscaler resource.
    /// 
    /// Autoscalers allow you to automatically scale virtual machine instances in
    /// managed instance groups according to an autoscaling policy that you
    /// define.
    /// 
    /// To get more information about RegionAutoscaler, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionAutoscalers)
    /// * How-to Guides
    ///     * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)
    /// 
    /// ## Example Usage
    /// ### Region Autoscaler Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foobarInstanceTemplate = new Gcp.Compute.InstanceTemplate("foobarInstanceTemplate", new()
    ///     {
    ///         MachineType = "e2-standard-4",
    ///         Disks = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.InstanceTemplateDiskArgs
    ///             {
    ///                 SourceImage = "debian-cloud/debian-11",
    ///                 DiskSizeGb = 250,
    ///             },
    ///         },
    ///         NetworkInterfaces = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.InstanceTemplateNetworkInterfaceArgs
    ///             {
    ///                 Network = "default",
    ///                 AccessConfigs = new[]
    ///                 {
    ///                     new Gcp.Compute.Inputs.InstanceTemplateNetworkInterfaceAccessConfigArgs
    ///                     {
    ///                         NetworkTier = "PREMIUM",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         ServiceAccount = new Gcp.Compute.Inputs.InstanceTemplateServiceAccountArgs
    ///         {
    ///             Scopes = new[]
    ///             {
    ///                 "https://www.googleapis.com/auth/devstorage.read_only",
    ///                 "https://www.googleapis.com/auth/logging.write",
    ///                 "https://www.googleapis.com/auth/monitoring.write",
    ///                 "https://www.googleapis.com/auth/pubsub",
    ///                 "https://www.googleapis.com/auth/service.management.readonly",
    ///                 "https://www.googleapis.com/auth/servicecontrol",
    ///                 "https://www.googleapis.com/auth/trace.append",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var foobarTargetPool = new Gcp.Compute.TargetPool("foobarTargetPool");
    /// 
    ///     var foobarRegionInstanceGroupManager = new Gcp.Compute.RegionInstanceGroupManager("foobarRegionInstanceGroupManager", new()
    ///     {
    ///         Region = "us-central1",
    ///         Versions = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.RegionInstanceGroupManagerVersionArgs
    ///             {
    ///                 InstanceTemplate = foobarInstanceTemplate.Id,
    ///                 Name = "primary",
    ///             },
    ///         },
    ///         TargetPools = new[]
    ///         {
    ///             foobarTargetPool.Id,
    ///         },
    ///         BaseInstanceName = "foobar",
    ///     });
    /// 
    ///     var foobarRegionAutoscaler = new Gcp.Compute.RegionAutoscaler("foobarRegionAutoscaler", new()
    ///     {
    ///         Region = "us-central1",
    ///         Target = foobarRegionInstanceGroupManager.Id,
    ///         AutoscalingPolicy = new Gcp.Compute.Inputs.RegionAutoscalerAutoscalingPolicyArgs
    ///         {
    ///             MaxReplicas = 5,
    ///             MinReplicas = 1,
    ///             CooldownPeriod = 60,
    ///             CpuUtilization = new Gcp.Compute.Inputs.RegionAutoscalerAutoscalingPolicyCpuUtilizationArgs
    ///             {
    ///                 Target = 0.5,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var debian9 = Gcp.Compute.GetImage.Invoke(new()
    ///     {
    ///         Family = "debian-11",
    ///         Project = "debian-cloud",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RegionAutoscaler can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default projects/{{project}}/regions/{{region}}/autoscalers/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionAutoscaler:RegionAutoscaler default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/regionAutoscaler:RegionAutoscaler")]
    public partial class RegionAutoscaler : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can
        /// define one or more of the policies for an autoscaler: cpuUtilization,
        /// customMetricUtilizations, and loadBalancingUtilization.
        /// If none of these are specified, the default will be to autoscale based
        /// on cpuUtilization to 0.6 or 60%.
        /// Structure is documented below.
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<Outputs.RegionAutoscalerAutoscalingPolicy> AutoscalingPolicy { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// A description of a scaling schedule.
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the instance group resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a RegionAutoscaler resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionAutoscaler(string name, RegionAutoscalerArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionAutoscaler:RegionAutoscaler", name, args ?? new RegionAutoscalerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionAutoscaler(string name, Input<string> id, RegionAutoscalerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionAutoscaler:RegionAutoscaler", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionAutoscaler resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionAutoscaler Get(string name, Input<string> id, RegionAutoscalerState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionAutoscaler(name, id, state, options);
        }
    }

    public sealed class RegionAutoscalerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can
        /// define one or more of the policies for an autoscaler: cpuUtilization,
        /// customMetricUtilizations, and loadBalancingUtilization.
        /// If none of these are specified, the default will be to autoscale based
        /// on cpuUtilization to 0.6 or 60%.
        /// Structure is documented below.
        /// </summary>
        [Input("autoscalingPolicy", required: true)]
        public Input<Inputs.RegionAutoscalerAutoscalingPolicyArgs> AutoscalingPolicy { get; set; } = null!;

        /// <summary>
        /// A description of a scaling schedule.
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// URL of the region where the instance group resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        public RegionAutoscalerArgs()
        {
        }
        public static new RegionAutoscalerArgs Empty => new RegionAutoscalerArgs();
    }

    public sealed class RegionAutoscalerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can
        /// define one or more of the policies for an autoscaler: cpuUtilization,
        /// customMetricUtilizations, and loadBalancingUtilization.
        /// If none of these are specified, the default will be to autoscale based
        /// on cpuUtilization to 0.6 or 60%.
        /// Structure is documented below.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.RegionAutoscalerAutoscalingPolicyGetArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// A description of a scaling schedule.
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// URL of the region where the instance group resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public RegionAutoscalerState()
        {
        }
        public static new RegionAutoscalerState Empty => new RegionAutoscalerState();
    }
}
