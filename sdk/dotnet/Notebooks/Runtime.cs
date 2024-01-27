// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Notebooks
{
    /// <summary>
    /// A Cloud AI Platform Notebook runtime.
    /// 
    /// &gt; **Note:** Due to limitations of the Notebooks Runtime API, many fields
    /// in this resource do not properly detect drift. These fields will also not
    /// appear in state once imported.
    /// 
    /// To get more information about Runtime, see:
    /// 
    /// * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
    /// 
    /// ## Example Usage
    /// ### Notebook Runtime Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var runtime = new Gcp.Notebooks.Runtime("runtime", new()
    ///     {
    ///         AccessConfig = new Gcp.Notebooks.Inputs.RuntimeAccessConfigArgs
    ///         {
    ///             AccessType = "SINGLE_USER",
    ///             RuntimeOwner = "admin@hashicorptest.com",
    ///         },
    ///         Location = "us-central1",
    ///         VirtualMachine = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineArgs
    ///         {
    ///             VirtualMachineConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigArgs
    ///             {
    ///                 DataDisk = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs
    ///                 {
    ///                     InitializeParams = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs
    ///                     {
    ///                         DiskSizeGb = 100,
    ///                         DiskType = "PD_STANDARD",
    ///                     },
    ///                 },
    ///                 MachineType = "n1-standard-4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Notebook Runtime Basic Gpu
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var runtimeGpu = new Gcp.Notebooks.Runtime("runtimeGpu", new()
    ///     {
    ///         AccessConfig = new Gcp.Notebooks.Inputs.RuntimeAccessConfigArgs
    ///         {
    ///             AccessType = "SINGLE_USER",
    ///             RuntimeOwner = "admin@hashicorptest.com",
    ///         },
    ///         Location = "us-central1",
    ///         SoftwareConfig = new Gcp.Notebooks.Inputs.RuntimeSoftwareConfigArgs
    ///         {
    ///             InstallGpuDriver = true,
    ///         },
    ///         VirtualMachine = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineArgs
    ///         {
    ///             VirtualMachineConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigArgs
    ///             {
    ///                 AcceleratorConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigArgs
    ///                 {
    ///                     CoreCount = 1,
    ///                     Type = "NVIDIA_TESLA_V100",
    ///                 },
    ///                 DataDisk = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs
    ///                 {
    ///                     InitializeParams = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs
    ///                     {
    ///                         DiskSizeGb = 100,
    ///                         DiskType = "PD_STANDARD",
    ///                     },
    ///                 },
    ///                 MachineType = "n1-standard-4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Notebook Runtime Basic Container
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var runtimeContainer = new Gcp.Notebooks.Runtime("runtimeContainer", new()
    ///     {
    ///         AccessConfig = new Gcp.Notebooks.Inputs.RuntimeAccessConfigArgs
    ///         {
    ///             AccessType = "SINGLE_USER",
    ///             RuntimeOwner = "admin@hashicorptest.com",
    ///         },
    ///         Location = "us-central1",
    ///         VirtualMachine = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineArgs
    ///         {
    ///             VirtualMachineConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigArgs
    ///             {
    ///                 ContainerImages = new[]
    ///                 {
    ///                     new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigContainerImageArgs
    ///                     {
    ///                         Repository = "gcr.io/deeplearning-platform-release/base-cpu",
    ///                         Tag = "latest",
    ///                     },
    ///                     new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigContainerImageArgs
    ///                     {
    ///                         Repository = "gcr.io/deeplearning-platform-release/beam-notebooks",
    ///                         Tag = "latest",
    ///                     },
    ///                 },
    ///                 DataDisk = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs
    ///                 {
    ///                     InitializeParams = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs
    ///                     {
    ///                         DiskSizeGb = 100,
    ///                         DiskType = "PD_STANDARD",
    ///                     },
    ///                 },
    ///                 MachineType = "n1-standard-4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Notebook Runtime Kernels
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var runtimeContainer = new Gcp.Notebooks.Runtime("runtimeContainer", new()
    ///     {
    ///         AccessConfig = new Gcp.Notebooks.Inputs.RuntimeAccessConfigArgs
    ///         {
    ///             AccessType = "SINGLE_USER",
    ///             RuntimeOwner = "admin@hashicorptest.com",
    ///         },
    ///         Labels = 
    ///         {
    ///             { "k", "val" },
    ///         },
    ///         Location = "us-central1",
    ///         SoftwareConfig = new Gcp.Notebooks.Inputs.RuntimeSoftwareConfigArgs
    ///         {
    ///             Kernels = new[]
    ///             {
    ///                 new Gcp.Notebooks.Inputs.RuntimeSoftwareConfigKernelArgs
    ///                 {
    ///                     Repository = "gcr.io/deeplearning-platform-release/base-cpu",
    ///                     Tag = "latest",
    ///                 },
    ///             },
    ///         },
    ///         VirtualMachine = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineArgs
    ///         {
    ///             VirtualMachineConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigArgs
    ///             {
    ///                 DataDisk = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs
    ///                 {
    ///                     InitializeParams = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs
    ///                     {
    ///                         DiskSizeGb = 100,
    ///                         DiskType = "PD_STANDARD",
    ///                     },
    ///                 },
    ///                 MachineType = "n1-standard-4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Notebook Runtime Script
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var runtimeContainer = new Gcp.Notebooks.Runtime("runtimeContainer", new()
    ///     {
    ///         AccessConfig = new Gcp.Notebooks.Inputs.RuntimeAccessConfigArgs
    ///         {
    ///             AccessType = "SINGLE_USER",
    ///             RuntimeOwner = "admin@hashicorptest.com",
    ///         },
    ///         Labels = 
    ///         {
    ///             { "k", "val" },
    ///         },
    ///         Location = "us-central1",
    ///         SoftwareConfig = new Gcp.Notebooks.Inputs.RuntimeSoftwareConfigArgs
    ///         {
    ///             PostStartupScriptBehavior = "RUN_EVERY_START",
    ///         },
    ///         VirtualMachine = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineArgs
    ///         {
    ///             VirtualMachineConfig = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigArgs
    ///             {
    ///                 DataDisk = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskArgs
    ///                 {
    ///                     InitializeParams = new Gcp.Notebooks.Inputs.RuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsArgs
    ///                     {
    ///                         DiskSizeGb = 100,
    ///                         DiskType = "PD_STANDARD",
    ///                     },
    ///                 },
    ///                 MachineType = "n1-standard-4",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Runtime can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/runtimes/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Runtime can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:notebooks/runtime:Runtime default projects/{{project}}/locations/{{location}}/runtimes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:notebooks/runtime:Runtime default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:notebooks/runtime:Runtime default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:notebooks/runtime:Runtime")]
    public partial class Runtime : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// Structure is documented below.
        /// </summary>
        [Output("accessConfig")]
        public Output<Outputs.RuntimeAccessConfig?> AccessConfig { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// The health state of this runtime. For a list of possible output
        /// values, see `https://cloud.google.com/vertex-ai/docs/workbench/
        /// reference/rest/v1/projects.locations.runtimes#healthstate`.
        /// </summary>
        [Output("healthState")]
        public Output<string> HealthState { get; private set; } = null!;

        /// <summary>
        /// The labels to associate with this runtime. Label **keys** must
        /// contain 1 to 63 characters, and must conform to [RFC 1035]
        /// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
        /// empty, but, if present, must contain 1 to 63 characters, and must
        /// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
        /// more than 32 labels can be associated with a cluster.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// A reference to the zone where the machine resides.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Contains Runtime daemon metrics such as Service status and JupyterLab
        /// status
        /// Structure is documented below.
        /// </summary>
        [Output("metrics")]
        public Output<ImmutableArray<Outputs.RuntimeMetric>> Metrics { get; private set; } = null!;

        /// <summary>
        /// The name specified for the Notebook runtime.
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The config settings for software inside the runtime.
        /// Structure is documented below.
        /// </summary>
        [Output("softwareConfig")]
        public Output<Outputs.RuntimeSoftwareConfig> SoftwareConfig { get; private set; } = null!;

        /// <summary>
        /// The state of this runtime.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// Structure is documented below.
        /// </summary>
        [Output("virtualMachine")]
        public Output<Outputs.RuntimeVirtualMachine?> VirtualMachine { get; private set; } = null!;


        /// <summary>
        /// Create a Runtime resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Runtime(string name, RuntimeArgs args, CustomResourceOptions? options = null)
            : base("gcp:notebooks/runtime:Runtime", name, args ?? new RuntimeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Runtime(string name, Input<string> id, RuntimeState? state = null, CustomResourceOptions? options = null)
            : base("gcp:notebooks/runtime:Runtime", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Runtime resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Runtime Get(string name, Input<string> id, RuntimeState? state = null, CustomResourceOptions? options = null)
        {
            return new Runtime(name, id, state, options);
        }
    }

    public sealed class RuntimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// Structure is documented below.
        /// </summary>
        [Input("accessConfig")]
        public Input<Inputs.RuntimeAccessConfigArgs>? AccessConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels to associate with this runtime. Label **keys** must
        /// contain 1 to 63 characters, and must conform to [RFC 1035]
        /// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
        /// empty, but, if present, must contain 1 to 63 characters, and must
        /// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
        /// more than 32 labels can be associated with a cluster.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// A reference to the zone where the machine resides.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name specified for the Notebook runtime.
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
        /// The config settings for software inside the runtime.
        /// Structure is documented below.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.RuntimeSoftwareConfigArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// Structure is documented below.
        /// </summary>
        [Input("virtualMachine")]
        public Input<Inputs.RuntimeVirtualMachineArgs>? VirtualMachine { get; set; }

        public RuntimeArgs()
        {
        }
        public static new RuntimeArgs Empty => new RuntimeArgs();
    }

    public sealed class RuntimeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config settings for accessing runtime.
        /// Structure is documented below.
        /// </summary>
        [Input("accessConfig")]
        public Input<Inputs.RuntimeAccessConfigGetArgs>? AccessConfig { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The health state of this runtime. For a list of possible output
        /// values, see `https://cloud.google.com/vertex-ai/docs/workbench/
        /// reference/rest/v1/projects.locations.runtimes#healthstate`.
        /// </summary>
        [Input("healthState")]
        public Input<string>? HealthState { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels to associate with this runtime. Label **keys** must
        /// contain 1 to 63 characters, and must conform to [RFC 1035]
        /// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
        /// empty, but, if present, must contain 1 to 63 characters, and must
        /// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
        /// more than 32 labels can be associated with a cluster.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// A reference to the zone where the machine resides.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metrics")]
        private InputList<Inputs.RuntimeMetricGetArgs>? _metrics;

        /// <summary>
        /// Contains Runtime daemon metrics such as Service status and JupyterLab
        /// status
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.RuntimeMetricGetArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.RuntimeMetricGetArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The name specified for the Notebook runtime.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The config settings for software inside the runtime.
        /// Structure is documented below.
        /// </summary>
        [Input("softwareConfig")]
        public Input<Inputs.RuntimeSoftwareConfigGetArgs>? SoftwareConfig { get; set; }

        /// <summary>
        /// The state of this runtime.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Use a Compute Engine VM image to start the managed notebook instance.
        /// Structure is documented below.
        /// </summary>
        [Input("virtualMachine")]
        public Input<Inputs.RuntimeVirtualMachineGetArgs>? VirtualMachine { get; set; }

        public RuntimeState()
        {
        }
        public static new RuntimeState Empty => new RuntimeState();
    }
}
