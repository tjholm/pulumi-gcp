// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig
    {
        /// <summary>
        /// Required. Resource name of an existing Dataproc Metastore service. Example: * `projects/`
        /// </summary>
        public readonly string DataprocMetastoreService;

        [OutputConstructor]
        private WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(string dataprocMetastoreService)
        {
            DataprocMetastoreService = dataprocMetastoreService;
        }
    }
}
