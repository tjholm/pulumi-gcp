// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DatabaseMigrationService.Inputs
{

    public sealed class ConnectionProfileMysqlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
        /// </summary>
        [Input("cloudSqlId")]
        public Input<string>? CloudSqlId { get; set; }

        /// <summary>
        /// Required. The IP or hostname of the source MySQL database.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
        /// This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
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
        /// (Output)
        /// Output only. Indicates If this connection profile password is stored.
        /// </summary>
        [Input("passwordSet")]
        public Input<bool>? PasswordSet { get; set; }

        /// <summary>
        /// Required. The network port of the source MySQL database.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// SSL configuration for the destination to connect to the source database.
        /// Structure is documented below.
        /// </summary>
        [Input("ssl")]
        public Input<Inputs.ConnectionProfileMysqlSslArgs>? Ssl { get; set; }

        /// <summary>
        /// Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ConnectionProfileMysqlArgs()
        {
        }
        public static new ConnectionProfileMysqlArgs Empty => new ConnectionProfileMysqlArgs();
    }
}
