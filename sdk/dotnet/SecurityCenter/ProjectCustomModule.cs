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
    /// Represents an instance of a Security Health Analytics custom module, including
    /// its full module name, display name, enablement state, and last updated time.
    /// You can create a custom module at the organization, folder, or project level.
    /// Custom modules that you create at the organization or folder level are inherited
    /// by the child folders and projects.
    /// 
    /// To get more information about ProjectCustomModule, see:
    /// 
    /// * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/projects.securityHealthAnalyticsSettings.customModules)
    /// * How-to Guides
    ///     * [Overview of custom modules for Security Health Analytics](https://cloud.google.com/security-command-center/docs/custom-modules-sha-overview)
    /// 
    /// ## Example Usage
    /// ### Scc Project Custom Module Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Gcp.SecurityCenter.ProjectCustomModule("example", new()
    ///     {
    ///         CustomConfig = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigArgs
    ///         {
    ///             Description = "The rotation period of the identified cryptokey resource exceeds 30 days.",
    ///             Predicate = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigPredicateArgs
    ///             {
    ///                 Expression = "resource.rotationPeriod &gt; duration(\"2592000s\")",
    ///             },
    ///             Recommendation = "Set the rotation period to at most 30 days.",
    ///             ResourceSelector = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigResourceSelectorArgs
    ///             {
    ///                 ResourceTypes = new[]
    ///                 {
    ///                     "cloudkms.googleapis.com/CryptoKey",
    ///                 },
    ///             },
    ///             Severity = "MEDIUM",
    ///         },
    ///         DisplayName = "basic_custom_module",
    ///         EnablementState = "ENABLED",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Scc Project Custom Module Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Gcp.SecurityCenter.ProjectCustomModule("example", new()
    ///     {
    ///         CustomConfig = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigArgs
    ///         {
    ///             CustomOutput = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigCustomOutputArgs
    ///             {
    ///                 Properties = new[]
    ///                 {
    ///                     new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigCustomOutputPropertyArgs
    ///                     {
    ///                         Name = "duration",
    ///                         ValueExpression = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigCustomOutputPropertyValueExpressionArgs
    ///                         {
    ///                             Description = "description of the expression",
    ///                             Expression = "resource.rotationPeriod",
    ///                             Location = "location of the expression",
    ///                             Title = "Purpose of the expression",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Description = "Description of the custom module",
    ///             Predicate = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigPredicateArgs
    ///             {
    ///                 Description = "description of the expression",
    ///                 Expression = "resource.rotationPeriod &gt; duration(\"2592000s\")",
    ///                 Location = "location of the expression",
    ///                 Title = "Purpose of the expression",
    ///             },
    ///             Recommendation = "Steps to resolve violation",
    ///             ResourceSelector = new Gcp.SecurityCenter.Inputs.ProjectCustomModuleCustomConfigResourceSelectorArgs
    ///             {
    ///                 ResourceTypes = new[]
    ///                 {
    ///                     "cloudkms.googleapis.com/CryptoKey",
    ///                 },
    ///             },
    ///             Severity = "LOW",
    ///         },
    ///         DisplayName = "full_custom_module",
    ///         EnablementState = "ENABLED",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ProjectCustomModule can be imported using any of these accepted formats* `projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, ProjectCustomModule can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:securitycenter/projectCustomModule:ProjectCustomModule")]
    public partial class ProjectCustomModule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If empty, indicates that the custom module was created in the organization,folder,
        /// or project in which you are viewing the custom module. Otherwise, ancestor_module
        /// specifies the organization or folder from which the custom module is inherited.
        /// </summary>
        [Output("ancestorModule")]
        public Output<string> AncestorModule { get; private set; } = null!;

        /// <summary>
        /// The user specified custom configuration for the module.
        /// Structure is documented below.
        /// </summary>
        [Output("customConfig")]
        public Output<Outputs.ProjectCustomModuleCustomConfig> CustomConfig { get; private set; } = null!;

        /// <summary>
        /// The display name of the Security Health Analytics custom module. This
        /// display name becomes the finding category for all findings that are
        /// returned by this custom module. The display name must be between 1 and
        /// 128 characters, start with a lowercase letter, and contain alphanumeric
        /// characters or underscores only.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The enablement state of the custom module.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Output("enablementState")]
        public Output<string> EnablementState { get; private set; } = null!;

        /// <summary>
        /// The editor that last updated the custom module.
        /// </summary>
        [Output("lastEditor")]
        public Output<string> LastEditor { get; private set; } = null!;

        /// <summary>
        /// Name of the property for the custom output.
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
        /// The time at which the custom module was last updated.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        /// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectCustomModule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectCustomModule(string name, ProjectCustomModuleArgs args, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/projectCustomModule:ProjectCustomModule", name, args ?? new ProjectCustomModuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectCustomModule(string name, Input<string> id, ProjectCustomModuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:securitycenter/projectCustomModule:ProjectCustomModule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectCustomModule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectCustomModule Get(string name, Input<string> id, ProjectCustomModuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectCustomModule(name, id, state, options);
        }
    }

    public sealed class ProjectCustomModuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user specified custom configuration for the module.
        /// Structure is documented below.
        /// </summary>
        [Input("customConfig", required: true)]
        public Input<Inputs.ProjectCustomModuleCustomConfigArgs> CustomConfig { get; set; } = null!;

        /// <summary>
        /// The display name of the Security Health Analytics custom module. This
        /// display name becomes the finding category for all findings that are
        /// returned by this custom module. The display name must be between 1 and
        /// 128 characters, start with a lowercase letter, and contain alphanumeric
        /// characters or underscores only.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The enablement state of the custom module.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("enablementState", required: true)]
        public Input<string> EnablementState { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectCustomModuleArgs()
        {
        }
        public static new ProjectCustomModuleArgs Empty => new ProjectCustomModuleArgs();
    }

    public sealed class ProjectCustomModuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If empty, indicates that the custom module was created in the organization,folder,
        /// or project in which you are viewing the custom module. Otherwise, ancestor_module
        /// specifies the organization or folder from which the custom module is inherited.
        /// </summary>
        [Input("ancestorModule")]
        public Input<string>? AncestorModule { get; set; }

        /// <summary>
        /// The user specified custom configuration for the module.
        /// Structure is documented below.
        /// </summary>
        [Input("customConfig")]
        public Input<Inputs.ProjectCustomModuleCustomConfigGetArgs>? CustomConfig { get; set; }

        /// <summary>
        /// The display name of the Security Health Analytics custom module. This
        /// display name becomes the finding category for all findings that are
        /// returned by this custom module. The display name must be between 1 and
        /// 128 characters, start with a lowercase letter, and contain alphanumeric
        /// characters or underscores only.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The enablement state of the custom module.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("enablementState")]
        public Input<string>? EnablementState { get; set; }

        /// <summary>
        /// The editor that last updated the custom module.
        /// </summary>
        [Input("lastEditor")]
        public Input<string>? LastEditor { get; set; }

        /// <summary>
        /// Name of the property for the custom output.
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
        /// The time at which the custom module was last updated.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        /// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ProjectCustomModuleState()
        {
        }
        public static new ProjectCustomModuleState Empty => new ProjectCustomModuleState();
    }
}
