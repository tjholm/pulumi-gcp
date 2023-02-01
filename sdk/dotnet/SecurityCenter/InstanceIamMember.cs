// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecurityCenter
{
    /// <summary>
    /// Represents a Data Fusion instance.
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/data-fusion/docs/reference/rest/v1beta1/projects.locations.instances)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/data-fusion/docs/)
    /// 
    /// ## Example Usage
    /// ### Data Fusion Instance Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basicInstance = new Gcp.DataFusion.Instance("basicInstance", new()
    ///     {
    ///         Options = 
    ///         {
    ///             { "prober_test_run", "true" },
    ///         },
    ///         Region = "us-central1",
    ///         Type = "BASIC",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Fusion Instance Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Gcp.AppEngine.GetDefaultServiceAccount.Invoke();
    /// 
    ///     var network = new Gcp.Compute.Network("network");
    /// 
    ///     var privateIpAlloc = new Gcp.Compute.GlobalAddress("privateIpAlloc", new()
    ///     {
    ///         AddressType = "INTERNAL",
    ///         Purpose = "VPC_PEERING",
    ///         PrefixLength = 22,
    ///         Network = network.Id,
    ///     });
    /// 
    ///     var extendedInstance = new Gcp.DataFusion.Instance("extendedInstance", new()
    ///     {
    ///         Description = "My Data Fusion instance",
    ///         DisplayName = "My Data Fusion instance",
    ///         Region = "us-central1",
    ///         Type = "BASIC",
    ///         EnableStackdriverLogging = true,
    ///         EnableStackdriverMonitoring = true,
    ///         PrivateInstance = true,
    ///         Version = "6.6.0",
    ///         DataprocServiceAccount = @default.Apply(@default =&gt; @default.Apply(getDefaultServiceAccountResult =&gt; getDefaultServiceAccountResult.Email)),
    ///         Labels = 
    ///         {
    ///             { "example_key", "example_value" },
    ///         },
    ///         NetworkConfig = new Gcp.DataFusion.Inputs.InstanceNetworkConfigArgs
    ///         {
    ///             Network = "default",
    ///             IpAllocation = Output.Tuple(privateIpAlloc.Address, privateIpAlloc.PrefixLength).Apply(values =&gt;
    ///             {
    ///                 var address = values.Item1;
    ///                 var prefixLength = values.Item2;
    ///                 return $"{address}/{prefixLength}";
    ///             }),
    ///         },
    ///         Options = 
    ///         {
    ///             { "prober_test_run", "true" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Fusion Instance Cmek
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRing("keyRing", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var cryptoKey = new Gcp.Kms.CryptoKey("cryptoKey", new()
    ///     {
    ///         KeyRing = keyRing.Id,
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var cryptoKeyBinding = new Gcp.Kms.CryptoKeyIAMBinding("cryptoKeyBinding", new()
    ///     {
    ///         CryptoKeyId = cryptoKey.Id,
    ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
    ///         Members = new[]
    ///         {
    ///             $"serviceAccount:service-{project.Apply(getProjectResult =&gt; getProjectResult.Number)}@gcp-sa-datafusion.iam.gserviceaccount.com",
    ///         },
    ///     });
    /// 
    ///     var cmek = new Gcp.DataFusion.Instance("cmek", new()
    ///     {
    ///         Region = "us-central1",
    ///         Type = "BASIC",
    ///         CryptoKeyConfig = new Gcp.DataFusion.Inputs.InstanceCryptoKeyConfigArgs
    ///         {
    ///             KeyReference = cryptoKey.Id,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             cryptoKeyBinding,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Fusion Instance Enterprise
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var enterpriseInstance = new Gcp.DataFusion.Instance("enterpriseInstance", new()
    ///     {
    ///         EnableRbac = true,
    ///         Options = 
    ///         {
    ///             { "prober_test_run", "true" },
    ///         },
    ///         Region = "us-central1",
    ///         Type = "ENTERPRISE",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Fusion Instance Event
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var eventTopic = new Gcp.PubSub.Topic("eventTopic");
    /// 
    ///     var eventInstance = new Gcp.DataFusion.Instance("eventInstance", new()
    ///     {
    ///         Region = "us-central1",
    ///         Type = "BASIC",
    ///         Version = "6.7.0",
    ///         EventPublishConfig = new Gcp.DataFusion.Inputs.InstanceEventPublishConfigArgs
    ///         {
    ///             Enabled = true,
    ///             Topic = eventTopic.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Data Fusion Instance Zone
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zone = new Gcp.DataFusion.Instance("zone", new()
    ///     {
    ///         Region = "us-central1",
    ///         Type = "DEVELOPER",
    ///         Zone = "us-central1-a",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Instance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/instanceIamMember:InstanceIamMember default projects/{{project}}/locations/{{region}}/instances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/instanceIamMember:InstanceIamMember default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/instanceIamMember:InstanceIamMember default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/instanceIamMember:InstanceIamMember default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:securitycenter/instanceIamMember:InstanceIamMember")]
    public partial class InstanceIamMember : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.InstanceIamMemberCondition?> Condition { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceIamMember(string name, InstanceIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/instanceIamMember:InstanceIamMember", name, args ?? new InstanceIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceIamMember(string name, Input<string> id, InstanceIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/instanceIamMember:InstanceIamMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InstanceIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceIamMember Get(string name, Input<string> id, InstanceIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceIamMember(name, id, state, options);
        }
    }

    public sealed class InstanceIamMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.InstanceIamMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public InstanceIamMemberArgs()
        {
        }
        public static new InstanceIamMemberArgs Empty => new InstanceIamMemberArgs();
    }

    public sealed class InstanceIamMemberState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.InstanceIamMemberConditionGetArgs>? Condition { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The ID of the instance or a fully qualified identifier for the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the Data Fusion instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        public InstanceIamMemberState()
        {
        }
        public static new InstanceIamMemberState Empty => new InstanceIamMemberState();
    }
}
