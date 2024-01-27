// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam
{
    /// <summary>
    /// A configuration for an external identity provider.
    /// 
    /// To get more information about WorkforcePoolProvider, see:
    /// 
    /// * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools.providers)
    /// * How-to Guides
    ///     * [Configure a provider within the workforce pool](https://cloud.google.com/iam/docs/manage-workforce-identity-pools-providers#configure_a_provider_within_the_workforce_pool)
    /// 
    /// &gt; **Note:** Ask your Google Cloud account team to request access to workforce identity federation for your
    /// billing/quota project. The account team notifies you when the project is granted access.
    /// 
    /// ## Example Usage
    /// ### Iam Workforce Pool Provider Saml Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.Iam.WorkforcePool("pool", new()
    ///     {
    ///         WorkforcePoolId = "example-pool",
    ///         Parent = "organizations/123456789",
    ///         Location = "global",
    ///     });
    /// 
    ///     var example = new Gcp.Iam.WorkforcePoolProvider("example", new()
    ///     {
    ///         WorkforcePoolId = pool.WorkforcePoolId,
    ///         Location = pool.Location,
    ///         ProviderId = "example-prvdr",
    ///         AttributeMapping = 
    ///         {
    ///             { "google.subject", "assertion.sub" },
    ///         },
    ///         Saml = new Gcp.Iam.Inputs.WorkforcePoolProviderSamlArgs
    ///         {
    ///             IdpMetadataXml = "&lt;?xml version=\"1.0\"?&gt;&lt;md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://test.com\"&gt;&lt;md:IDPSSODescriptor protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"&gt; &lt;md:KeyDescriptor use=\"signing\"&gt;&lt;ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"&gt;&lt;ds:X509Data&gt;&lt;ds:X509Certificate&gt;MIIDpDCCAoygAwIBAgIGAX7/5qPhMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi00NTg0MjExHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjIwMjE2MDAxOTEyWhcNMzIwMjE2MDAyMDEyWjCBkjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNDU4NDIxMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxrBl7GKz52cRpxF9xCsirnRuMxnhFBaUrsHqAQrLqWmdlpNYZTVg+T9iQ+aq/iE68L+BRZcZniKIvW58wqqS0ltXVvIkXuDSvnvnkkI5yMIVErR20K8jSOKQm1FmK+fgAJ4koshFiu9oLiqu0Ejc0DuL3/XRsb4RuxjktKTb1khgBBtb+7idEk0sFR0RPefAweXImJkDHDm7SxjDwGJUubbqpdTxasPr0W+AHI1VUzsUsTiHAoyb0XDkYqHfDzhj/ZdIEl4zHQ3bEZvlD984ztAnmX2SuFLLKfXeAAGHei8MMixJvwxYkkPeYZ/5h8WgBZPP4heS2CPjwYExt29L8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQARjJFz++a9Z5IQGFzsZMrX2EDR5ML4xxUiQkbhld1S1PljOLcYFARDmUC2YYHOueU4ee8Jid9nPGEUebV/4Jok+b+oQh+dWMgiWjSLI7h5q4OYZ3VJtdlVwgMFt2iz+/4yBKMUZ50g3Qgg36vE34us+eKitg759JgCNsibxn0qtJgSPm0sgP2L6yTaLnoEUbXBRxCwynTSkp9ZijZqEzbhN0e2dWv7Rx/nfpohpDP6vEiFImKFHpDSv3M/5de1ytQzPFrZBYt9WlzlYwE1aD9FHCxdd+rWgYMVVoRaRmndpV/Rq3QUuDuFJtaoX11bC7ExkOpg9KstZzA63i3VcfYv&lt;/ds:X509Certificate&gt;&lt;/ds:X509Data&gt;&lt;/ds:KeyInfo&gt;&lt;/md:KeyDescriptor&gt;&lt;md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://test.com/sso\"/&gt;&lt;/md:IDPSSODescriptor&gt;&lt;/md:EntityDescriptor&gt;",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Iam Workforce Pool Provider Saml Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.Iam.WorkforcePool("pool", new()
    ///     {
    ///         WorkforcePoolId = "example-pool",
    ///         Parent = "organizations/123456789",
    ///         Location = "global",
    ///     });
    /// 
    ///     var example = new Gcp.Iam.WorkforcePoolProvider("example", new()
    ///     {
    ///         WorkforcePoolId = pool.WorkforcePoolId,
    ///         Location = pool.Location,
    ///         ProviderId = "example-prvdr",
    ///         AttributeMapping = 
    ///         {
    ///             { "google.subject", "assertion.sub" },
    ///         },
    ///         Saml = new Gcp.Iam.Inputs.WorkforcePoolProviderSamlArgs
    ///         {
    ///             IdpMetadataXml = "&lt;?xml version=\"1.0\"?&gt;&lt;md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://test.com\"&gt;&lt;md:IDPSSODescriptor protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"&gt; &lt;md:KeyDescriptor use=\"signing\"&gt;&lt;ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"&gt;&lt;ds:X509Data&gt;&lt;ds:X509Certificate&gt;MIIDpDCCAoygAwIBAgIGAX7/5qPhMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi00NTg0MjExHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjIwMjE2MDAxOTEyWhcNMzIwMjE2MDAyMDEyWjCBkjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNDU4NDIxMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxrBl7GKz52cRpxF9xCsirnRuMxnhFBaUrsHqAQrLqWmdlpNYZTVg+T9iQ+aq/iE68L+BRZcZniKIvW58wqqS0ltXVvIkXuDSvnvnkkI5yMIVErR20K8jSOKQm1FmK+fgAJ4koshFiu9oLiqu0Ejc0DuL3/XRsb4RuxjktKTb1khgBBtb+7idEk0sFR0RPefAweXImJkDHDm7SxjDwGJUubbqpdTxasPr0W+AHI1VUzsUsTiHAoyb0XDkYqHfDzhj/ZdIEl4zHQ3bEZvlD984ztAnmX2SuFLLKfXeAAGHei8MMixJvwxYkkPeYZ/5h8WgBZPP4heS2CPjwYExt29L8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQARjJFz++a9Z5IQGFzsZMrX2EDR5ML4xxUiQkbhld1S1PljOLcYFARDmUC2YYHOueU4ee8Jid9nPGEUebV/4Jok+b+oQh+dWMgiWjSLI7h5q4OYZ3VJtdlVwgMFt2iz+/4yBKMUZ50g3Qgg36vE34us+eKitg759JgCNsibxn0qtJgSPm0sgP2L6yTaLnoEUbXBRxCwynTSkp9ZijZqEzbhN0e2dWv7Rx/nfpohpDP6vEiFImKFHpDSv3M/5de1ytQzPFrZBYt9WlzlYwE1aD9FHCxdd+rWgYMVVoRaRmndpV/Rq3QUuDuFJtaoX11bC7ExkOpg9KstZzA63i3VcfYv&lt;/ds:X509Certificate&gt;&lt;/ds:X509Data&gt;&lt;/ds:KeyInfo&gt;&lt;/md:KeyDescriptor&gt;&lt;md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://test.com/sso\"/&gt;&lt;/md:IDPSSODescriptor&gt;&lt;/md:EntityDescriptor&gt;",
    ///         },
    ///         DisplayName = "Display name",
    ///         Description = "A sample SAML workforce pool provider.",
    ///         Disabled = false,
    ///         AttributeCondition = "true",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Iam Workforce Pool Provider Oidc Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.Iam.WorkforcePool("pool", new()
    ///     {
    ///         WorkforcePoolId = "example-pool",
    ///         Parent = "organizations/123456789",
    ///         Location = "global",
    ///     });
    /// 
    ///     var example = new Gcp.Iam.WorkforcePoolProvider("example", new()
    ///     {
    ///         WorkforcePoolId = pool.WorkforcePoolId,
    ///         Location = pool.Location,
    ///         ProviderId = "example-prvdr",
    ///         AttributeMapping = 
    ///         {
    ///             { "google.subject", "assertion.sub" },
    ///         },
    ///         Oidc = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcArgs
    ///         {
    ///             IssuerUri = "https://accounts.thirdparty.com",
    ///             ClientId = "client-id",
    ///             ClientSecret = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcClientSecretArgs
    ///             {
    ///                 Value = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcClientSecretValueArgs
    ///                 {
    ///                     PlainText = "client-secret",
    ///                 },
    ///             },
    ///             WebSsoConfig = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcWebSsoConfigArgs
    ///             {
    ///                 ResponseType = "CODE",
    ///                 AssertionClaimsBehavior = "MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Iam Workforce Pool Provider Oidc Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.Iam.WorkforcePool("pool", new()
    ///     {
    ///         WorkforcePoolId = "example-pool",
    ///         Parent = "organizations/123456789",
    ///         Location = "global",
    ///     });
    /// 
    ///     var example = new Gcp.Iam.WorkforcePoolProvider("example", new()
    ///     {
    ///         WorkforcePoolId = pool.WorkforcePoolId,
    ///         Location = pool.Location,
    ///         ProviderId = "example-prvdr",
    ///         AttributeMapping = 
    ///         {
    ///             { "google.subject", "assertion.sub" },
    ///         },
    ///         Oidc = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcArgs
    ///         {
    ///             IssuerUri = "https://accounts.thirdparty.com",
    ///             ClientId = "client-id",
    ///             ClientSecret = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcClientSecretArgs
    ///             {
    ///                 Value = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcClientSecretValueArgs
    ///                 {
    ///                     PlainText = "client-secret",
    ///                 },
    ///             },
    ///             WebSsoConfig = new Gcp.Iam.Inputs.WorkforcePoolProviderOidcWebSsoConfigArgs
    ///             {
    ///                 ResponseType = "CODE",
    ///                 AssertionClaimsBehavior = "MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS",
    ///                 AdditionalScopes = new[]
    ///                 {
    ///                     "groups",
    ///                     "roles",
    ///                 },
    ///             },
    ///         },
    ///         DisplayName = "Display name",
    ///         Description = "A sample OIDC workforce pool provider.",
    ///         Disabled = false,
    ///         AttributeCondition = "true",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WorkforcePoolProvider can be imported using any of these accepted formats* `locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}` * `{{location}}/{{workforce_pool_id}}/{{provider_id}}` When using the `pulumi import` command, WorkforcePoolProvider can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workforcePoolProvider:WorkforcePoolProvider default locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workforcePoolProvider:WorkforcePoolProvider default {{location}}/{{workforce_pool_id}}/{{provider_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:iam/workforcePoolProvider:WorkforcePoolProvider")]
    public partial class WorkforcePoolProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A [Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// </summary>
        [Output("attributeCondition")]
        public Output<string?> AttributeCondition { get; private set; } = null!;

        /// <summary>
        /// Maps attributes from the authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings.
        /// This is also the subject that appears in Cloud Logging logs. This is a required field and
        /// the mapped subject cannot exceed 127 bytes.
        /// * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to
        /// resources using an IAM `principalSet` binding; access applies to all members of the group.
        /// * `google.display_name`: The name of the authenticated user. This is an optional field and
        /// the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo.
        /// This is an optional field. When set, the image will be visible as the user's profile picture.
        /// If not set, a generic user icon will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute}
        /// is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
        /// The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
        /// to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute.
        /// For example, the following maps the sub claim of the incoming credential to the `subject` attribute
        /// on a Google token:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// An object containing a list of `"key": value` pairs.
        /// Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        [Output("attributeMapping")]
        public Output<ImmutableDictionary<string, string>?> AttributeMapping { get; private set; } = null!;

        /// <summary>
        /// A user-specified description of the provider. Cannot exceed 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// A user-specified display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The location for the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Output only. The resource name of the provider.
        /// Format: `locations/{location}/workforcePools/{workforcePoolId}/providers/{providerId}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Represents an OpenId Connect 1.0 identity provider.
        /// Structure is documented below.
        /// </summary>
        [Output("oidc")]
        public Output<Outputs.WorkforcePoolProviderOidc?> Oidc { get; private set; } = null!;

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name.
        /// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("providerId")]
        public Output<string> ProviderId { get; private set; } = null!;

        /// <summary>
        /// Represents a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Output("saml")]
        public Output<Outputs.WorkforcePoolProviderSaml?> Saml { get; private set; } = null!;

        /// <summary>
        /// The current state of the provider.
        /// * STATE_UNSPECIFIED: State unspecified.
        /// * ACTIVE: The provider is active and may be used to validate authentication credentials.
        /// * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
        /// deleted after approximately 30 days. You can restore a soft-deleted provider using
        /// [providers.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools.providers/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePoolProvider).
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name.
        /// The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
        /// It must start with a letter, and cannot have a trailing hyphen.
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Output("workforcePoolId")]
        public Output<string> WorkforcePoolId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkforcePoolProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkforcePoolProvider(string name, WorkforcePoolProviderArgs args, CustomResourceOptions? options = null)
            : base("gcp:iam/workforcePoolProvider:WorkforcePoolProvider", name, args ?? new WorkforcePoolProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkforcePoolProvider(string name, Input<string> id, WorkforcePoolProviderState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iam/workforcePoolProvider:WorkforcePoolProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkforcePoolProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkforcePoolProvider Get(string name, Input<string> id, WorkforcePoolProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkforcePoolProvider(name, id, state, options);
        }
    }

    public sealed class WorkforcePoolProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A [Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// </summary>
        [Input("attributeCondition")]
        public Input<string>? AttributeCondition { get; set; }

        [Input("attributeMapping")]
        private InputMap<string>? _attributeMapping;

        /// <summary>
        /// Maps attributes from the authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings.
        /// This is also the subject that appears in Cloud Logging logs. This is a required field and
        /// the mapped subject cannot exceed 127 bytes.
        /// * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to
        /// resources using an IAM `principalSet` binding; access applies to all members of the group.
        /// * `google.display_name`: The name of the authenticated user. This is an optional field and
        /// the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo.
        /// This is an optional field. When set, the image will be visible as the user's profile picture.
        /// If not set, a generic user icon will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute}
        /// is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
        /// The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
        /// to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute.
        /// For example, the following maps the sub claim of the incoming credential to the `subject` attribute
        /// on a Google token:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// An object containing a list of `"key": value` pairs.
        /// Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> AttributeMapping
        {
            get => _attributeMapping ?? (_attributeMapping = new InputMap<string>());
            set => _attributeMapping = value;
        }

        /// <summary>
        /// A user-specified description of the provider. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A user-specified display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The location for the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Represents an OpenId Connect 1.0 identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.WorkforcePoolProviderOidcArgs>? Oidc { get; set; }

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name.
        /// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        /// <summary>
        /// Represents a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("saml")]
        public Input<Inputs.WorkforcePoolProviderSamlArgs>? Saml { get; set; }

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name.
        /// The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
        /// It must start with a letter, and cannot have a trailing hyphen.
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workforcePoolId", required: true)]
        public Input<string> WorkforcePoolId { get; set; } = null!;

        public WorkforcePoolProviderArgs()
        {
        }
        public static new WorkforcePoolProviderArgs Empty => new WorkforcePoolProviderArgs();
    }

    public sealed class WorkforcePoolProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A [Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// </summary>
        [Input("attributeCondition")]
        public Input<string>? AttributeCondition { get; set; }

        [Input("attributeMapping")]
        private InputMap<string>? _attributeMapping;

        /// <summary>
        /// Maps attributes from the authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings.
        /// This is also the subject that appears in Cloud Logging logs. This is a required field and
        /// the mapped subject cannot exceed 127 bytes.
        /// * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to
        /// resources using an IAM `principalSet` binding; access applies to all members of the group.
        /// * `google.display_name`: The name of the authenticated user. This is an optional field and
        /// the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo.
        /// This is an optional field. When set, the image will be visible as the user's profile picture.
        /// If not set, a generic user icon will be displayed instead.
        /// This attribute cannot be referenced in IAM bindings.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute}
        /// is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes.
        /// The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a workforce pool
        /// to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/locations/{location}/workforcePools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/locations/{location}/workforcePools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For OIDC providers, you must supply a custom mapping that includes the `google.subject` attribute.
        /// For example, the following maps the sub claim of the incoming credential to the `subject` attribute
        /// on a Google token:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// An object containing a list of `"key": value` pairs.
        /// Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> AttributeMapping
        {
            get => _attributeMapping ?? (_attributeMapping = new InputMap<string>());
            set => _attributeMapping = value;
        }

        /// <summary>
        /// A user-specified description of the provider. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A user-specified display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The location for the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Output only. The resource name of the provider.
        /// Format: `locations/{location}/workforcePools/{workforcePoolId}/providers/{providerId}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Represents an OpenId Connect 1.0 identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.WorkforcePoolProviderOidcGetArgs>? Oidc { get; set; }

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name.
        /// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// Represents a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("saml")]
        public Input<Inputs.WorkforcePoolProviderSamlGetArgs>? Saml { get; set; }

        /// <summary>
        /// The current state of the provider.
        /// * STATE_UNSPECIFIED: State unspecified.
        /// * ACTIVE: The provider is active and may be used to validate authentication credentials.
        /// * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
        /// deleted after approximately 30 days. You can restore a soft-deleted provider using
        /// [providers.undelete](https://cloud.google.com/iam/docs/reference/rest/v1/locations.workforcePools.providers/undelete#google.iam.admin.v1.WorkforcePools.UndeleteWorkforcePoolProvider).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name.
        /// The IDs must be a globally unique string of 6 to 63 lowercase letters, digits, or hyphens.
        /// It must start with a letter, and cannot have a trailing hyphen.
        /// The prefix `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workforcePoolId")]
        public Input<string>? WorkforcePoolId { get; set; }

        public WorkforcePoolProviderState()
        {
        }
        public static new WorkforcePoolProviderState Empty => new WorkforcePoolProviderState();
    }
}
