// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase.Inputs
{

    public sealed class ExtensionsInstanceConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedEventTypes")]
        private InputList<string>? _allowedEventTypes;

        /// <summary>
        /// List of extension events selected by consumer that extension is allowed to
        /// emit, identified by their types.
        /// </summary>
        public InputList<string> AllowedEventTypes
        {
            get => _allowedEventTypes ?? (_allowedEventTypes = new InputList<string>());
            set => _allowedEventTypes = value;
        }

        /// <summary>
        /// (Output)
        /// The time at which the Extension Instance Config was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Fully qualified Eventarc resource name that consumers should use for event triggers.
        /// </summary>
        [Input("eventarcChannel")]
        public Input<string>? EventarcChannel { get; set; }

        /// <summary>
        /// The ref of the Extension from the Registry (e.g. publisher-id/awesome-extension)
        /// </summary>
        [Input("extensionRef", required: true)]
        public Input<string> ExtensionRef { get; set; } = null!;

        /// <summary>
        /// The version of the Extension from the Registry (e.g. 1.0.3). If left blank, latest is assumed.
        /// </summary>
        [Input("extensionVersion")]
        public Input<string>? ExtensionVersion { get; set; }

        /// <summary>
        /// (Output)
        /// The unique identifier for this configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("params", required: true)]
        private InputMap<string>? _params;

        /// <summary>
        /// Environment variables that may be configured for the Extension
        /// </summary>
        public InputMap<string> Params
        {
            get => _params ?? (_params = new InputMap<string>());
            set => _params = value;
        }

        /// <summary>
        /// (Output)
        /// Postinstall instructions to be shown for this Extension, with
        /// template strings representing function and parameter values substituted
        /// with actual values. These strings include: ${param:FOO},
        /// ${function:myFunc.url},
        /// ${function:myFunc.name}, and ${function:myFunc.location}
        /// 
        /// - - -
        /// </summary>
        [Input("populatedPostinstallContent")]
        public Input<string>? PopulatedPostinstallContent { get; set; }

        [Input("systemParams")]
        private InputMap<string>? _systemParams;

        /// <summary>
        /// Params whose values are only available at deployment time.
        /// Unlike other params, these will not be set as environment variables on
        /// functions. See a full list of system parameters at
        /// https://firebase.google.com/docs/extensions/publishers/parameters#system_parameters
        /// </summary>
        public InputMap<string> SystemParams
        {
            get => _systemParams ?? (_systemParams = new InputMap<string>());
            set => _systemParams = value;
        }

        public ExtensionsInstanceConfigGetArgs()
        {
        }
        public static new ExtensionsInstanceConfigGetArgs Empty => new ExtensionsInstanceConfigGetArgs();
    }
}
