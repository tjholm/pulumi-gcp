// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Datastream.Inputs
{

    public sealed class ConnectionProfileOracleProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectionAttributes")]
        private InputMap<string>? _connectionAttributes;

        /// <summary>
        /// Connection string attributes
        /// </summary>
        public InputMap<string> ConnectionAttributes
        {
            get => _connectionAttributes ?? (_connectionAttributes = new InputMap<string>());
            set => _connectionAttributes = value;
        }

        /// <summary>
        /// Database for the Oracle connection.
        /// </summary>
        [Input("databaseService", required: true)]
        public Input<string> DatabaseService { get; set; } = null!;

        /// <summary>
        /// Hostname for the SSH tunnel.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// SSH password.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port for the SSH tunnel.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Username for the SSH tunnel.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ConnectionProfileOracleProfileArgs()
        {
        }
        public static new ConnectionProfileOracleProfileArgs Empty => new ConnectionProfileOracleProfileArgs();
    }
}
