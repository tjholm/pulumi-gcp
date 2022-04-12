// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class JobPrestoConfigArgs : Pulumi.ResourceArgs
    {
        [Input("clientTags")]
        private InputList<string>? _clientTags;

        /// <summary>
        /// Presto client tags to attach to this query.
        /// </summary>
        public InputList<string> ClientTags
        {
            get => _clientTags ?? (_clientTags = new InputList<string>());
            set => _clientTags = value;
        }

        /// <summary>
        /// Whether to continue executing queries if a query fails. Setting to true can be useful when executing independent parallel queries. Defaults to false.
        /// </summary>
        [Input("continueOnFailure")]
        public Input<bool>? ContinueOnFailure { get; set; }

        [Input("loggingConfig")]
        public Input<Inputs.JobPrestoConfigLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The format in which query output will be displayed. See the Presto documentation for supported output formats.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// A mapping of property names to values. Used to set Presto session properties Equivalent to using the --session flag in the Presto CLI.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The HCFS URI of the script that contains SQL queries.
        /// Conflicts with `query_list`
        /// </summary>
        [Input("queryFileUri")]
        public Input<string>? QueryFileUri { get; set; }

        [Input("queryLists")]
        private InputList<string>? _queryLists;

        /// <summary>
        /// The list of SQL queries or statements to execute as part of the job.
        /// Conflicts with `query_file_uri`
        /// </summary>
        public InputList<string> QueryLists
        {
            get => _queryLists ?? (_queryLists = new InputList<string>());
            set => _queryLists = value;
        }

        public JobPrestoConfigArgs()
        {
        }
    }
}
