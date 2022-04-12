// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class JobSparkConfigArgs : Pulumi.ResourceArgs
    {
        [Input("archiveUris")]
        private InputList<string>? _archiveUris;

        /// <summary>
        /// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
        /// </summary>
        public InputList<string> ArchiveUris
        {
            get => _archiveUris ?? (_archiveUris = new InputList<string>());
            set => _archiveUris = value;
        }

        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("fileUris")]
        private InputList<string>? _fileUris;

        /// <summary>
        /// HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
        /// </summary>
        public InputList<string> FileUris
        {
            get => _fileUris ?? (_fileUris = new InputList<string>());
            set => _fileUris = value;
        }

        [Input("jarFileUris")]
        private InputList<string>? _jarFileUris;

        /// <summary>
        /// HCFS URIs of jar files to be added to the Spark CLASSPATH.
        /// </summary>
        public InputList<string> JarFileUris
        {
            get => _jarFileUris ?? (_jarFileUris = new InputList<string>());
            set => _jarFileUris = value;
        }

        [Input("loggingConfig")]
        public Input<Inputs.JobSparkConfigLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`. Conflicts with `main_jar_file_uri`
        /// </summary>
        [Input("mainClass")]
        public Input<string>? MainClass { get; set; }

        /// <summary>
        /// The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with `main_class`
        /// </summary>
        [Input("mainJarFileUri")]
        public Input<string>? MainJarFileUri { get; set; }

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

        public JobSparkConfigArgs()
        {
        }
    }
}
