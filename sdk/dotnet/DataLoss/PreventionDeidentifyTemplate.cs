// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss
{
    /// <summary>
    /// Allows creation of templates to de-identify content.
    /// 
    /// To get more information about DeidentifyTemplate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// DeidentifyTemplate can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate")]
    public partial class PreventionDeidentifyTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration of the deidentify template
        /// Structure is documented below.
        /// </summary>
        [Output("deidentifyConfig")]
        public Output<Outputs.PreventionDeidentifyTemplateDeidentifyConfig> DeidentifyConfig { get; private set; } = null!;

        /// <summary>
        /// A description of the template.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User set display name of the template.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of the template in any of the following formats:
        /// * `projects/{{project}}`
        /// * `projects/{{project}}/locations/{{location}}`
        /// * `organizations/{{organization_id}}`
        /// * `organizations/{{organization_id}}/locations/{{location}}`
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;


        /// <summary>
        /// Create a PreventionDeidentifyTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PreventionDeidentifyTemplate(string name, PreventionDeidentifyTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate", name, args ?? new PreventionDeidentifyTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PreventionDeidentifyTemplate(string name, Input<string> id, PreventionDeidentifyTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PreventionDeidentifyTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PreventionDeidentifyTemplate Get(string name, Input<string> id, PreventionDeidentifyTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new PreventionDeidentifyTemplate(name, id, state, options);
        }
    }

    public sealed class PreventionDeidentifyTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the deidentify template
        /// Structure is documented below.
        /// </summary>
        [Input("deidentifyConfig", required: true)]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigArgs> DeidentifyConfig { get; set; } = null!;

        /// <summary>
        /// A description of the template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User set display name of the template.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The parent of the template in any of the following formats:
        /// * `projects/{{project}}`
        /// * `projects/{{project}}/locations/{{location}}`
        /// * `organizations/{{organization_id}}`
        /// * `organizations/{{organization_id}}/locations/{{location}}`
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public PreventionDeidentifyTemplateArgs()
        {
        }
    }

    public sealed class PreventionDeidentifyTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the deidentify template
        /// Structure is documented below.
        /// </summary>
        [Input("deidentifyConfig")]
        public Input<Inputs.PreventionDeidentifyTemplateDeidentifyConfigGetArgs>? DeidentifyConfig { get; set; }

        /// <summary>
        /// A description of the template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User set display name of the template.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at [https://cloud.google.com/dlp/docs/infotypes-reference](https://cloud.google.com/dlp/docs/infotypes-reference) when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the template in any of the following formats:
        /// * `projects/{{project}}`
        /// * `projects/{{project}}/locations/{{location}}`
        /// * `organizations/{{organization_id}}`
        /// * `organizations/{{organization_id}}/locations/{{location}}`
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        public PreventionDeidentifyTemplateState()
        {
        }
    }
}
