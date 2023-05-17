// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Tags
{
    /// <summary>
    /// A TagBinding represents a connection between a TagValue and a Regional cloud resource (currently project, folder, or organization). Once a TagBinding is created, the TagValue is applied to all the descendants of the cloud resource.
    /// 
    /// To get more information about TagBinding, see:
    /// 
    /// * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v3/tagBindings)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/resource-manager/docs/tags/tags-creating-and-managing)
    /// 
    /// ## Example Usage
    /// 
    /// To bind a tag to a Cloud Run instance:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Gcp.Organizations.Project("project", new()
    ///     {
    ///         OrgId = "123456789",
    ///         ProjectId = "project_id",
    ///     });
    /// 
    ///     var key = new Gcp.Tags.TagKey("key", new()
    ///     {
    ///         Description = "For keyname resources.",
    ///         Parent = "organizations/123456789",
    ///         ShortName = "keyname",
    ///     });
    /// 
    ///     var @value = new Gcp.Tags.TagValue("value", new()
    ///     {
    ///         Description = "For valuename resources.",
    ///         Parent = key.Name.Apply(name =&gt; $"tagKeys/{name}"),
    ///         ShortName = "valuename",
    ///     });
    /// 
    ///     var binding = new Gcp.Tags.LocationTagBinding("binding", new()
    ///     {
    ///         Location = "us-central1",
    ///         Parent = project.Number.Apply(number =&gt; $"//run.googleapis.com/projects/{number}/locations/{google_cloud_run_service.Default.Location}/services/{google_cloud_run_service.Default.Name}"),
    ///         TagValue = @value.Name.Apply(name =&gt; $"tagValues/{name}"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// To bind a (firewall) tag to compute instance:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Gcp.Organizations.Project("project", new()
    ///     {
    ///         OrgId = "123456789",
    ///         ProjectId = "project_id",
    ///     });
    /// 
    ///     var key = new Gcp.Tags.TagKey("key", new()
    ///     {
    ///         Description = "For keyname resources.",
    ///         Parent = "organizations/123456789",
    ///         ShortName = "keyname",
    ///     });
    /// 
    ///     var @value = new Gcp.Tags.TagValue("value", new()
    ///     {
    ///         Description = "For valuename resources.",
    ///         Parent = key.Name.Apply(name =&gt; $"tagKeys/{name}"),
    ///         ShortName = "valuename",
    ///     });
    /// 
    ///     var binding = new Gcp.Tags.LocationTagBinding("binding", new()
    ///     {
    ///         Location = "us-central1",
    ///         Parent = project.Number.Apply(number =&gt; $"//compute.googleapis.com/projects/{number}/zones/us-central1-a/instances/{google_compute_instance.Instance.Instance_id}"),
    ///         TagValue = @value.Name.Apply(name =&gt; $"tagValues/{name}"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TagBinding can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:tags/locationTagBinding:LocationTagBinding default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:tags/locationTagBinding:LocationTagBinding")]
    public partial class LocationTagBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The generated id for the TagBinding. This is a string of the form: `tagBindings/{parent}/{tag-value-name}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The TagValue of the TagBinding. Must be of the form tagValues/456.
        /// </summary>
        [Output("tagValue")]
        public Output<string> TagValue { get; private set; } = null!;


        /// <summary>
        /// Create a LocationTagBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationTagBinding(string name, LocationTagBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:tags/locationTagBinding:LocationTagBinding", name, args ?? new LocationTagBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationTagBinding(string name, Input<string> id, LocationTagBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:tags/locationTagBinding:LocationTagBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LocationTagBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationTagBinding Get(string name, Input<string> id, LocationTagBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new LocationTagBinding(name, id, state, options);
        }
    }

    public sealed class LocationTagBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// The TagValue of the TagBinding. Must be of the form tagValues/456.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

        public LocationTagBindingArgs()
        {
        }
        public static new LocationTagBindingArgs Empty => new LocationTagBindingArgs();
    }

    public sealed class LocationTagBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The generated id for the TagBinding. This is a string of the form: `tagBindings/{parent}/{tag-value-name}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// The TagValue of the TagBinding. Must be of the form tagValues/456.
        /// </summary>
        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

        public LocationTagBindingState()
        {
        }
        public static new LocationTagBindingState Empty => new LocationTagBindingState();
    }
}
