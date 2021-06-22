// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplateJobHiveJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
        /// </summary>
        [Input("continueOnFailure")]
        public Input<bool>? ContinueOnFailure { get; set; }

        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see (https://cloud.google.com/dataproc/docs/concepts/cluster-properties).
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The HCFS URI of the script that contains SQL queries.
        /// </summary>
        [Input("queryFileUri")]
        public Input<string>? QueryFileUri { get; set; }

        /// <summary>
        /// A list of queries.
        /// </summary>
        [Input("queryList")]
        public Input<Inputs.WorkflowTemplateJobHiveJobQueryListArgs>? QueryList { get; set; }

        [Input("scriptVariables")]
        private InputMap<string>? _scriptVariables;

        /// <summary>
        /// Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET `name="value";`).
        /// </summary>
        public InputMap<string> ScriptVariables
        {
            get => _scriptVariables ?? (_scriptVariables = new InputMap<string>());
            set => _scriptVariables = value;
        }

        public WorkflowTemplateJobHiveJobArgs()
        {
        }
    }
}
