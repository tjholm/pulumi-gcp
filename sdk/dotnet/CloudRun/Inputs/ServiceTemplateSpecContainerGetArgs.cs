// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class ServiceTemplateSpecContainerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// Arguments to the entrypoint.
        /// The docker image's CMD is used if this is not provided.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Entrypoint array. Not executed within a shell.
        /// The docker image's ENTRYPOINT is used if this is not provided.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("envFroms")]
        private InputList<Inputs.ServiceTemplateSpecContainerEnvFromGetArgs>? _envFroms;

        /// <summary>
        /// (Optional, Deprecated)
        /// List of sources to populate environment variables in the container.
        /// All invalid keys will be reported as an event when the container is starting.
        /// When a key exists in multiple sources, the value associated with the last source will
        /// take precedence. Values defined by an Env with a duplicate key will take
        /// precedence.
        /// Structure is documented below.
        /// </summary>
        [Obsolete(@"Not supported by Cloud Run fully managed")]
        public InputList<Inputs.ServiceTemplateSpecContainerEnvFromGetArgs> EnvFroms
        {
            get => _envFroms ?? (_envFroms = new InputList<Inputs.ServiceTemplateSpecContainerEnvFromGetArgs>());
            set => _envFroms = value;
        }

        [Input("envs")]
        private InputList<Inputs.ServiceTemplateSpecContainerEnvGetArgs>? _envs;

        /// <summary>
        /// List of environment variables to set in the container.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerEnvGetArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.ServiceTemplateSpecContainerEnvGetArgs>());
            set => _envs = value;
        }

        /// <summary>
        /// Docker image name. This is most often a reference to a container located
        /// in the container registry, such as gcr.io/cloudrun/hello
        /// </summary>
        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        /// <summary>
        /// Periodic probe of container liveness. Container will be restarted if the probe fails. More info:
        /// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
        /// Structure is documented below.
        /// </summary>
        [Input("livenessProbe")]
        public Input<Inputs.ServiceTemplateSpecContainerLivenessProbeGetArgs>? LivenessProbe { get; set; }

        [Input("ports")]
        private InputList<Inputs.ServiceTemplateSpecContainerPortGetArgs>? _ports;

        /// <summary>
        /// List of open ports in the container.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServiceTemplateSpecContainerPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Compute Resources required by this container. Used to set values such as max memory
        /// Structure is documented below.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.ServiceTemplateSpecContainerResourcesGetArgs>? Resources { get; set; }

        /// <summary>
        /// Startup probe of application within the container.
        /// All other probes are disabled if a startup probe is provided, until it
        /// succeeds. Container will not be added to service endpoints if the probe fails.
        /// Structure is documented below.
        /// </summary>
        [Input("startupProbe")]
        public Input<Inputs.ServiceTemplateSpecContainerStartupProbeGetArgs>? StartupProbe { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.ServiceTemplateSpecContainerVolumeMountGetArgs>? _volumeMounts;

        /// <summary>
        /// Volume to mount into the container's filesystem.
        /// Only supports SecretVolumeSources.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceTemplateSpecContainerVolumeMountGetArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.ServiceTemplateSpecContainerVolumeMountGetArgs>());
            set => _volumeMounts = value;
        }

        /// <summary>
        /// (Optional, Deprecated)
        /// Container's working directory.
        /// If not specified, the container runtime's default will be used, which
        /// might be configured in the container image.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public ServiceTemplateSpecContainerGetArgs()
        {
        }
        public static new ServiceTemplateSpecContainerGetArgs Empty => new ServiceTemplateSpecContainerGetArgs();
    }
}
