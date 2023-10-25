// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable
{
    /// <summary>
    /// ## +---
    /// 
    /// subcategory: "Cloud Bigtable"
    /// description: |-
    ///   Creates a Google Bigtable instance.
    /// ---
    /// 
    /// # gcp.bigtable.Instance
    /// 
    /// Creates a Google Bigtable instance. For more information see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/bigtable/docs)
    /// 
    /// ## Example Usage
    /// ### Simple Instance
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var production_instance = new Gcp.BigTable.Instance("production-instance", new()
    ///     {
    ///         Clusters = new[]
    ///         {
    ///             new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///             {
    ///                 ClusterId = "tf-instance-cluster",
    ///                 NumNodes = 1,
    ///                 StorageType = "HDD",
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "my-label", "prod-label" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Replicated Instance
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var production_instance = new Gcp.BigTable.Instance("production-instance", new()
    ///     {
    ///         Clusters = new[]
    ///         {
    ///             new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///             {
    ///                 ClusterId = "tf-instance-cluster1",
    ///                 NumNodes = 1,
    ///                 StorageType = "HDD",
    ///                 Zone = "us-central1-c",
    ///             },
    ///             new Gcp.BigTable.Inputs.InstanceClusterArgs
    ///             {
    ///                 AutoscalingConfig = new Gcp.BigTable.Inputs.InstanceClusterAutoscalingConfigArgs
    ///                 {
    ///                     CpuTarget = 50,
    ///                     MaxNodes = 3,
    ///                     MinNodes = 1,
    ///                 },
    ///                 ClusterId = "tf-instance-cluster2",
    ///                 StorageType = "HDD",
    ///                 Zone = "us-central1-b",
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "my-label", "prod-label" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Bigtable Instances can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instance:Instance default projects/{{project}}/instances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instance:Instance default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instance:Instance default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigtable/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A block of cluster configuration options. This can be specified at least once, and up 
        /// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
        /// to default to the backend value. See structure below.
        /// 
        /// -----
        /// </summary>
        [Output("clusters")]
        public Output<ImmutableArray<Outputs.InstanceCluster>> Clusters { get; private set; } = null!;

        /// <summary>
        /// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        /// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        /// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        /// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        /// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        /// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        /// 
        /// 
        /// -----
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance. Must be 6-33 characters and must only contain hyphens, lowercase letters and numbers.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusters")]
        private InputList<Inputs.InstanceClusterArgs>? _clusters;

        /// <summary>
        /// A block of cluster configuration options. This can be specified at least once, and up 
        /// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
        /// to default to the backend value. See structure below.
        /// 
        /// -----
        /// </summary>
        public InputList<Inputs.InstanceClusterArgs> Clusters
        {
            get => _clusters ?? (_clusters = new InputList<Inputs.InstanceClusterArgs>());
            set => _clusters = value;
        }

        /// <summary>
        /// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        /// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        /// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        /// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        /// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        /// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        /// 
        /// 
        /// -----
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance. Must be 6-33 characters and must only contain hyphens, lowercase letters and numbers.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("clusters")]
        private InputList<Inputs.InstanceClusterGetArgs>? _clusters;

        /// <summary>
        /// A block of cluster configuration options. This can be specified at least once, and up 
        /// to as many as possible within 8 cloud regions. Removing the field entirely from the config will cause the provider
        /// to default to the backend value. See structure below.
        /// 
        /// -----
        /// </summary>
        public InputList<Inputs.InstanceClusterGetArgs> Clusters
        {
            get => _clusters ?? (_clusters = new InputList<Inputs.InstanceClusterGetArgs>());
            set => _clusters = value;
        }

        /// <summary>
        /// Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        /// in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        /// It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        /// and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        /// `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        /// is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        /// 
        /// 
        /// -----
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance. Must be 6-33 characters and must only contain hyphens, lowercase letters and numbers.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
