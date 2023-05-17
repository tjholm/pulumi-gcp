// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for BigQuery Connection Connection. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigquery.ConnectionIamPolicy`: Authoritative. Sets the IAM policy for the connection and replaces any existing policy already attached.
    /// * `gcp.bigquery.ConnectionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the connection are preserved.
    /// * `gcp.bigquery.ConnectionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the connection are preserved.
    /// 
    /// &gt; **Note:** `gcp.bigquery.ConnectionIamPolicy` **cannot** be used in conjunction with `gcp.bigquery.ConnectionIamBinding` and `gcp.bigquery.ConnectionIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.bigquery.ConnectionIamBinding` resources **can be** used in conjunction with `gcp.bigquery.ConnectionIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_bigquery\_connection\_iam\_policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var admin = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.BigQuery.ConnectionIamPolicy("policy", new()
    ///     {
    ///         Project = google_bigquery_connection.Connection.Project,
    ///         Location = google_bigquery_connection.Connection.Location,
    ///         ConnectionId = google_bigquery_connection.Connection.Connection_id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigquery\_connection\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.BigQuery.ConnectionIamBinding("binding", new()
    ///     {
    ///         Project = google_bigquery_connection.Connection.Project,
    ///         Location = google_bigquery_connection.Connection.Location,
    ///         ConnectionId = google_bigquery_connection.Connection.Connection_id,
    ///         Role = "roles/viewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigquery\_connection\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.BigQuery.ConnectionIamMember("member", new()
    ///     {
    ///         Project = google_bigquery_connection.Connection.Project,
    ///         Location = google_bigquery_connection.Connection.Location,
    ///         ConnectionId = google_bigquery_connection.Connection.Connection_id,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/connections/{{connection_id}} * {{project}}/{{location}}/{{connection_id}} * {{location}}/{{connection_id}} * {{connection_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Connection connection IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy")]
    public partial class ConnectionIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional connection id that should be assigned to the created connection.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("connectionId")]
        public Output<string> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the connection should reside.
        /// Cloud SQL instance must be in the same location as the connection
        /// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        /// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        /// Spanner Connections same as spanner region
        /// AWS allowed regions are aws-us-east-1
        /// Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionIamPolicy(string name, ConnectionIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy", name, args ?? new ConnectionIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionIamPolicy(string name, Input<string> id, ConnectionIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionIamPolicy Get(string name, Input<string> id, ConnectionIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ConnectionIamPolicy(name, id, state, options);
        }
    }

    public sealed class ConnectionIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional connection id that should be assigned to the created connection.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        /// <summary>
        /// The geographic location where the connection should reside.
        /// Cloud SQL instance must be in the same location as the connection
        /// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        /// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        /// Spanner Connections same as spanner region
        /// AWS allowed regions are aws-us-east-1
        /// Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConnectionIamPolicyArgs()
        {
        }
        public static new ConnectionIamPolicyArgs Empty => new ConnectionIamPolicyArgs();
    }

    public sealed class ConnectionIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional connection id that should be assigned to the created connection.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The geographic location where the connection should reside.
        /// Cloud SQL instance must be in the same location as the connection
        /// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        /// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        /// Spanner Connections same as spanner region
        /// AWS allowed regions are aws-us-east-1
        /// Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConnectionIamPolicyState()
        {
        }
        public static new ConnectionIamPolicyState Empty => new ConnectionIamPolicyState();
    }
}
