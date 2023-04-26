// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// An `Instance` is the runtime dataplane in Apigee.
    /// 
    /// To get more information about Instance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances/create)
    /// * How-to Guides
    ///     * [Creating a runtime instance](https://cloud.google.com/apigee/docs/api-platform/get-started/create-instance)
    /// 
    /// ## Example Usage
    /// ### Apigee Instance Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Gcp.Organizations.GetClientConfig.Invoke();
    /// 
    ///     var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork");
    /// 
    ///     var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 16,
    ///         Network = apigeeNetwork.Id,
    ///     });
    /// 
    ///     var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new()
    ///     {
    ///         Network = apigeeNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             apigeeRange.Name,
    ///         },
    ///     });
    /// 
    ///     var apigeeOrg = new Gcp.Apigee.Organization("apigeeOrg", new()
    ///     {
    ///         AnalyticsRegion = "us-central1",
    ///         ProjectId = current.Apply(getClientConfigResult =&gt; getClientConfigResult.Project),
    ///         AuthorizedNetwork = apigeeNetwork.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigeeVpcConnection,
    ///         },
    ///     });
    /// 
    ///     var apigeeInstance = new Gcp.Apigee.Instance("apigeeInstance", new()
    ///     {
    ///         Location = "us-central1",
    ///         OrgId = apigeeOrg.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Apigee Instance Cidr Range
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Gcp.Organizations.GetClientConfig.Invoke();
    /// 
    ///     var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork");
    /// 
    ///     var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 22,
    ///         Network = apigeeNetwork.Id,
    ///     });
    /// 
    ///     var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new()
    ///     {
    ///         Network = apigeeNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             apigeeRange.Name,
    ///         },
    ///     });
    /// 
    ///     var apigeeOrg = new Gcp.Apigee.Organization("apigeeOrg", new()
    ///     {
    ///         AnalyticsRegion = "us-central1",
    ///         ProjectId = current.Apply(getClientConfigResult =&gt; getClientConfigResult.Project),
    ///         AuthorizedNetwork = apigeeNetwork.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigeeVpcConnection,
    ///         },
    ///     });
    /// 
    ///     var apigeeInstance = new Gcp.Apigee.Instance("apigeeInstance", new()
    ///     {
    ///         Location = "us-central1",
    ///         OrgId = apigeeOrg.Id,
    ///         PeeringCidrRange = "SLASH_22",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Apigee Instance Ip Range
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Gcp.Organizations.GetClientConfig.Invoke();
    /// 
    ///     var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork");
    /// 
    ///     var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 22,
    ///         Network = apigeeNetwork.Id,
    ///     });
    /// 
    ///     var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new()
    ///     {
    ///         Network = apigeeNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             apigeeRange.Name,
    ///         },
    ///     });
    /// 
    ///     var apigeeOrg = new Gcp.Apigee.Organization("apigeeOrg", new()
    ///     {
    ///         AnalyticsRegion = "us-central1",
    ///         ProjectId = current.Apply(getClientConfigResult =&gt; getClientConfigResult.Project),
    ///         AuthorizedNetwork = apigeeNetwork.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigeeVpcConnection,
    ///         },
    ///     });
    /// 
    ///     var apigeeInstance = new Gcp.Apigee.Instance("apigeeInstance", new()
    ///     {
    ///         Location = "us-central1",
    ///         OrgId = apigeeOrg.Id,
    ///         IpRange = "10.87.8.0/22",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Apigee Instance Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var current = Gcp.Organizations.GetClientConfig.Invoke();
    /// 
    ///     var apigeeNetwork = new Gcp.Compute.Network("apigeeNetwork");
    /// 
    ///     var apigeeRange = new Gcp.Compute.GlobalAddress("apigeeRange", new()
    ///     {
    ///         Purpose = "VPC_PEERING",
    ///         AddressType = "INTERNAL",
    ///         PrefixLength = 16,
    ///         Network = apigeeNetwork.Id,
    ///     });
    /// 
    ///     var apigeeVpcConnection = new Gcp.ServiceNetworking.Connection("apigeeVpcConnection", new()
    ///     {
    ///         Network = apigeeNetwork.Id,
    ///         Service = "servicenetworking.googleapis.com",
    ///         ReservedPeeringRanges = new[]
    ///         {
    ///             apigeeRange.Name,
    ///         },
    ///     });
    /// 
    ///     var apigeeKeyring = new Gcp.Kms.KeyRing("apigeeKeyring", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var apigeeKey = new Gcp.Kms.CryptoKey("apigeeKey", new()
    ///     {
    ///         KeyRing = apigeeKeyring.Id,
    ///     });
    /// 
    ///     var apigeeSa = new Gcp.Projects.ServiceIdentity("apigeeSa", new()
    ///     {
    ///         Project = google_project.Project.Project_id,
    ///         Service = google_project_service.Apigee.Service,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var apigeeSaKeyuser = new Gcp.Kms.CryptoKeyIAMBinding("apigeeSaKeyuser", new()
    ///     {
    ///         CryptoKeyId = apigeeKey.Id,
    ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
    ///         Members = new[]
    ///         {
    ///             apigeeSa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///         },
    ///     });
    /// 
    ///     var apigeeOrg = new Gcp.Apigee.Organization("apigeeOrg", new()
    ///     {
    ///         AnalyticsRegion = "us-central1",
    ///         DisplayName = "apigee-org",
    ///         Description = "Auto-provisioned Apigee Org.",
    ///         ProjectId = current.Apply(getClientConfigResult =&gt; getClientConfigResult.Project),
    ///         AuthorizedNetwork = apigeeNetwork.Id,
    ///         RuntimeDatabaseEncryptionKeyName = apigeeKey.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             apigeeVpcConnection,
    ///             apigeeSaKeyuser,
    ///         },
    ///     });
    /// 
    ///     var apigeeInstance = new Gcp.Apigee.Instance("apigeeInstance", new()
    ///     {
    ///         Location = "us-central1",
    ///         Description = "Auto-managed Apigee Runtime Instance",
    ///         DisplayName = "my-instance-name",
    ///         OrgId = apigeeOrg.Id,
    ///         DiskEncryptionKeyName = apigeeKey.Id,
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
    ///  $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/instances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Customer accept list represents the list of projects (id/number) on customer
        /// side that can privately connect to the service attachment. It is an optional field
        /// which the customers can provide during the instance creation. By default, the customer
        /// project associated with the Apigee organization will be included to the list.
        /// </summary>
        [Output("consumerAcceptLists")]
        public Output<ImmutableArray<string>> ConsumerAcceptLists { get; private set; } = null!;

        /// <summary>
        /// Description of the instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        /// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        /// </summary>
        [Output("diskEncryptionKeyName")]
        public Output<string?> DiskEncryptionKeyName { get; private set; } = null!;

        /// <summary>
        /// Display name of the instance.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// IP range represents the customer-provided CIDR block of length 22 that will be used for
        /// the Apigee instance creation. This optional range, if provided, should be freely
        /// available as part of larger named range the customer has allocated to the Service
        /// Networking peering. If this is not provided, Apigee will automatically request for any
        /// available /22 CIDR block from Service Networking. The customer should use this CIDR block
        /// for configuring their firewall needs to allow traffic from Apigee.
        /// Input format: "a.b.c.d/22"
        /// </summary>
        [Output("ipRange")]
        public Output<string?> IpRange { get; private set; } = null!;

        /// <summary>
        /// Required. Compute Engine location where the instance resides.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource ID of the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Apigee Organization associated with the Apigee instance,
        /// in the format `organizations/{{org_name}}`.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// The size of the CIDR block range that will be reserved by the instance. For valid values,
        /// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
        /// </summary>
        [Output("peeringCidrRange")]
        public Output<string> PeeringCidrRange { get; private set; } = null!;

        /// <summary>
        /// Output only. Port number of the exposed Apigee endpoint.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Output only. Resource name of the service attachment created for the instance in
        /// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
        /// forward traffic to this service attachment using the PSC endpoints.
        /// </summary>
        [Output("serviceAttachment")]
        public Output<string> ServiceAttachment { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("consumerAcceptLists")]
        private InputList<string>? _consumerAcceptLists;

        /// <summary>
        /// Optional. Customer accept list represents the list of projects (id/number) on customer
        /// side that can privately connect to the service attachment. It is an optional field
        /// which the customers can provide during the instance creation. By default, the customer
        /// project associated with the Apigee organization will be included to the list.
        /// </summary>
        public InputList<string> ConsumerAcceptLists
        {
            get => _consumerAcceptLists ?? (_consumerAcceptLists = new InputList<string>());
            set => _consumerAcceptLists = value;
        }

        /// <summary>
        /// Description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        /// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        /// </summary>
        [Input("diskEncryptionKeyName")]
        public Input<string>? DiskEncryptionKeyName { get; set; }

        /// <summary>
        /// Display name of the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// IP range represents the customer-provided CIDR block of length 22 that will be used for
        /// the Apigee instance creation. This optional range, if provided, should be freely
        /// available as part of larger named range the customer has allocated to the Service
        /// Networking peering. If this is not provided, Apigee will automatically request for any
        /// available /22 CIDR block from Service Networking. The customer should use this CIDR block
        /// for configuring their firewall needs to allow traffic from Apigee.
        /// Input format: "a.b.c.d/22"
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// Required. Compute Engine location where the instance resides.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource ID of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Apigee Organization associated with the Apigee instance,
        /// in the format `organizations/{{org_name}}`.
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// The size of the CIDR block range that will be reserved by the instance. For valid values,
        /// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
        /// </summary>
        [Input("peeringCidrRange")]
        public Input<string>? PeeringCidrRange { get; set; }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("consumerAcceptLists")]
        private InputList<string>? _consumerAcceptLists;

        /// <summary>
        /// Optional. Customer accept list represents the list of projects (id/number) on customer
        /// side that can privately connect to the service attachment. It is an optional field
        /// which the customers can provide during the instance creation. By default, the customer
        /// project associated with the Apigee organization will be included to the list.
        /// </summary>
        public InputList<string> ConsumerAcceptLists
        {
            get => _consumerAcceptLists ?? (_consumerAcceptLists = new InputList<string>());
            set => _consumerAcceptLists = value;
        }

        /// <summary>
        /// Description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        /// Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        /// </summary>
        [Input("diskEncryptionKeyName")]
        public Input<string>? DiskEncryptionKeyName { get; set; }

        /// <summary>
        /// Display name of the instance.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// IP range represents the customer-provided CIDR block of length 22 that will be used for
        /// the Apigee instance creation. This optional range, if provided, should be freely
        /// available as part of larger named range the customer has allocated to the Service
        /// Networking peering. If this is not provided, Apigee will automatically request for any
        /// available /22 CIDR block from Service Networking. The customer should use this CIDR block
        /// for configuring their firewall needs to allow traffic from Apigee.
        /// Input format: "a.b.c.d/22"
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// Required. Compute Engine location where the instance resides.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource ID of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Apigee Organization associated with the Apigee instance,
        /// in the format `organizations/{{org_name}}`.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The size of the CIDR block range that will be reserved by the instance. For valid values,
        /// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
        /// </summary>
        [Input("peeringCidrRange")]
        public Input<string>? PeeringCidrRange { get; set; }

        /// <summary>
        /// Output only. Port number of the exposed Apigee endpoint.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Output only. Resource name of the service attachment created for the instance in
        /// the format: projects/*/regions/*/serviceAttachments/* Apigee customers can privately
        /// forward traffic to this service attachment using the PSC endpoints.
        /// </summary>
        [Input("serviceAttachment")]
        public Input<string>? ServiceAttachment { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
