// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
    /// 
    /// * `gcp.kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
    /// * `gcp.kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
    /// * `gcp.kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
    /// 
    /// &gt; **Note:** `gcp.kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `gcp.kms.KeyRingIAMBinding` and `gcp.kms.KeyRingIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.kms.KeyRingIAMBinding` resources **can be** used in conjunction with `gcp.kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## gcp.kms.KeyRingIAMPolicy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyring = new Gcp.Kms.KeyRing("keyring", new()
    ///     {
    ///         Name = "keyring-example",
    ///         Location = "global",
    ///     });
    /// 
    ///     var admin = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/editor",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var keyRing = new Gcp.Kms.KeyRingIAMPolicy("key_ring", new()
    ///     {
    ///         KeyRingId = keyring.Id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyring = new Gcp.Kms.KeyRing("keyring", new()
    ///     {
    ///         Name = "keyring-example",
    ///         Location = "global",
    ///     });
    /// 
    ///     var admin = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/editor",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///                 Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionInputArgs
    ///                 {
    ///                     Title = "expires_after_2019_12_31",
    ///                     Description = "Expiring at midnight of 2019-12-31",
    ///                     Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var keyRing = new Gcp.Kms.KeyRingIAMPolicy("key_ring", new()
    ///     {
    ///         KeyRingId = keyring.Id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.kms.KeyRingIAMBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMBinding("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMBinding("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Kms.Inputs.KeyRingIAMBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.kms.KeyRingIAMMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMMember("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMMember("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Kms.Inputs.KeyRingIAMMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.kms.KeyRingIAMBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMBinding("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMBinding("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Kms.Inputs.KeyRingIAMBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.kms.KeyRingIAMMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMMember("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRingIAMMember("key_ring", new()
    ///     {
    ///         KeyRingId = "your-key-ring-id",
    ///         Role = "roles/cloudkms.admin",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Kms.Inputs.KeyRingIAMMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies
    /// 
    /// IAM policy imports use the identifier of the Cloud KMS key ring only. For example:
    /// 
    /// * `{{project_id}}/{{location}}/{{key_ring_name}}`
    /// 
    /// An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
    /// 
    /// tf
    /// 
    /// import {
    /// 
    ///   id = "{{project_id}}/{{location}}/{{key_ring_name}}"
    /// 
    ///   to = google_kms_key_ring_iam_policy.default
    /// 
    /// }
    /// 
    /// The `pulumi import` command can also be used:
    /// 
    /// ```sh
    /// $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy default {{project_id}}/{{location}}/{{key_ring_name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy")]
    public partial class KeyRingIAMPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the key ring's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("keyRingId")]
        public Output<string> KeyRingId { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a KeyRingIAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRingIAMPolicy(string name, KeyRingIAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, args ?? new KeyRingIAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyRingIAMPolicy(string name, Input<string> id, KeyRingIAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRingIAMPolicy Get(string name, Input<string> id, KeyRingIAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyRingIAMPolicy(name, id, state, options);
        }
    }

    public sealed class KeyRingIAMPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public KeyRingIAMPolicyArgs()
        {
        }
        public static new KeyRingIAMPolicyArgs Empty => new KeyRingIAMPolicyArgs();
    }

    public sealed class KeyRingIAMPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the key ring's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId")]
        public Input<string>? KeyRingId { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public KeyRingIAMPolicyState()
        {
        }
        public static new KeyRingIAMPolicyState Empty => new KeyRingIAMPolicyState();
    }
}
