// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Inputs
{

    public sealed class WorkstationConfigContainerArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// Arguments passed to the entrypoint.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// If set, overrides the default ENTRYPOINT specified by the image.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("env")]
        private InputMap<string>? _env;

        /// <summary>
        /// Environment variables passed to the container.
        /// The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
        /// </summary>
        public InputMap<string> Env
        {
            get => _env ?? (_env = new InputMap<string>());
            set => _env = value;
        }

        /// <summary>
        /// Docker image defining the container. This image must be accessible by the config's service account.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// If set, overrides the USER specified in the image with the given uid.
        /// </summary>
        [Input("runAsUser")]
        public Input<int>? RunAsUser { get; set; }

        /// <summary>
        /// If set, overrides the default DIR specified by the image.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public WorkstationConfigContainerArgs()
        {
        }
        public static new WorkstationConfigContainerArgs Empty => new WorkstationConfigContainerArgs();
    }
}
