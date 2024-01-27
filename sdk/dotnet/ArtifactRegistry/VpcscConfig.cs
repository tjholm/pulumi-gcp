// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ArtifactRegistry
{
    /// <summary>
    /// ## Example Usage
    /// ### Artifact Registry Vpcsc Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_config = new Gcp.ArtifactRegistry.VpcscConfig("my-config", new()
    ///     {
    ///         Location = "us-central1",
    ///         VpcscPolicy = "ALLOW",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPCSCConfig can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, VPCSCConfig can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:artifactregistry/vpcscConfig:VpcscConfig")]
    public partial class VpcscConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the location this config is located in.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the project's VPC SC Config.
        /// Always of the form: projects/{project}/location/{location}/vpcscConfig
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
        /// The VPC SC policy for project and location.
        /// Possible values are: `DENY`, `ALLOW`.
        /// </summary>
        [Output("vpcscPolicy")]
        public Output<string?> VpcscPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a VpcscConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcscConfig(string name, VpcscConfigArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/vpcscConfig:VpcscConfig", name, args ?? new VpcscConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcscConfig(string name, Input<string> id, VpcscConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/vpcscConfig:VpcscConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcscConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcscConfig Get(string name, Input<string> id, VpcscConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcscConfig(name, id, state, options);
        }
    }

    public sealed class VpcscConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the location this config is located in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The VPC SC policy for project and location.
        /// Possible values are: `DENY`, `ALLOW`.
        /// </summary>
        [Input("vpcscPolicy")]
        public Input<string>? VpcscPolicy { get; set; }

        public VpcscConfigArgs()
        {
        }
        public static new VpcscConfigArgs Empty => new VpcscConfigArgs();
    }

    public sealed class VpcscConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the location this config is located in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the project's VPC SC Config.
        /// Always of the form: projects/{project}/location/{location}/vpcscConfig
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
        /// The VPC SC policy for project and location.
        /// Possible values are: `DENY`, `ALLOW`.
        /// </summary>
        [Input("vpcscPolicy")]
        public Input<string>? VpcscPolicy { get; set; }

        public VpcscConfigState()
        {
        }
        public static new VpcscConfigState Empty => new VpcscConfigState();
    }
}
