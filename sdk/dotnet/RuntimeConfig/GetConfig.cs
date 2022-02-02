// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.RuntimeConfig
{
    public static class GetConfig
    {
        public static Task<GetConfigResult> InvokeAsync(GetConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigResult>("gcp:runtimeconfig/getConfig:getConfig", args ?? new GetConfigArgs(), options.WithDefaults());

        public static Output<GetConfigResult> Invoke(GetConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigResult>("gcp:runtimeconfig/getConfig:getConfig", args ?? new GetConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Runtime Configurator configuration.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetConfigArgs()
        {
        }
    }

    public sealed class GetConfigInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Runtime Configurator configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Project;

        [OutputConstructor]
        private GetConfigResult(
            string description,

            string id,

            string name,

            string? project)
        {
            Description = description;
            Id = id;
            Name = name;
            Project = project;
        }
    }
}
