// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    public static class GetUptimeCheckIPs
    {
        /// <summary>
        /// Returns the list of IP addresses that checkers run from. For more information see
        /// the [official documentation](https://cloud.google.com/monitoring/uptime-checks#get-ips).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ips = Gcp.Monitoring.GetUptimeCheckIPs.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ipList"] = ips.Apply(getUptimeCheckIPsResult =&gt; getUptimeCheckIPsResult.UptimeCheckIps),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUptimeCheckIPsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUptimeCheckIPsResult>("gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Returns the list of IP addresses that checkers run from. For more information see
        /// the [official documentation](https://cloud.google.com/monitoring/uptime-checks#get-ips).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ips = Gcp.Monitoring.GetUptimeCheckIPs.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["ipList"] = ips.Apply(getUptimeCheckIPsResult =&gt; getUptimeCheckIPsResult.UptimeCheckIps),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUptimeCheckIPsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUptimeCheckIPsResult>("gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetUptimeCheckIPsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of uptime check IPs used by Stackdriver Monitoring. Each `uptime_check_ip` contains:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUptimeCheckIPsUptimeCheckIpResult> UptimeCheckIps;

        [OutputConstructor]
        private GetUptimeCheckIPsResult(
            string id,

            ImmutableArray<Outputs.GetUptimeCheckIPsUptimeCheckIpResult> uptimeCheckIps)
        {
            Id = id;
            UptimeCheckIps = uptimeCheckIps;
        }
    }
}
