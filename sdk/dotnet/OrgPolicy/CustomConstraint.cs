// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OrgPolicy
{
    /// <summary>
    /// Custom constraints are created by administrators to provide more granular and customizable control over the specific fields that are restricted by your organization policies.
    /// 
    /// To get more information about CustomConstraint, see:
    /// 
    /// * [API documentation](https://cloud.google.com/resource-manager/docs/reference/orgpolicy/rest/v2/organizations.constraints)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints)
    ///     * [Supported Services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services)
    /// 
    /// ## Example Usage
    /// ### Org Policy Custom Constraint Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var constraint = new Gcp.OrgPolicy.CustomConstraint("constraint", new()
    ///     {
    ///         ActionType = "ALLOW",
    ///         Condition = "resource.management.autoUpgrade == false",
    ///         MethodTypes = new[]
    ///         {
    ///             "CREATE",
    ///             "UPDATE",
    ///         },
    ///         Parent = "organizations/123456789",
    ///         ResourceTypes = new[]
    ///         {
    ///             "container.googleapis.com/NodePool",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Org Policy Custom Constraint Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var constraint = new Gcp.OrgPolicy.CustomConstraint("constraint", new()
    ///     {
    ///         ActionType = "ALLOW",
    ///         Condition = "resource.management.autoUpgrade == false",
    ///         Description = "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.",
    ///         DisplayName = "Disable GKE auto upgrade",
    ///         MethodTypes = new[]
    ///         {
    ///             "CREATE",
    ///             "UPDATE",
    ///         },
    ///         Parent = "organizations/123456789",
    ///         ResourceTypes = new[]
    ///         {
    ///             "container.googleapis.com/NodePool",
    ///         },
    ///     });
    /// 
    ///     var @bool = new Gcp.OrgPolicy.Policy("bool", new()
    ///     {
    ///         Parent = "organizations/123456789",
    ///         Spec = new Gcp.OrgPolicy.Inputs.PolicySpecArgs
    ///         {
    ///             Rules = new[]
    ///             {
    ///                 new Gcp.OrgPolicy.Inputs.PolicySpecRuleArgs
    ///                 {
    ///                     Enforce = "TRUE",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CustomConstraint can be imported using any of these accepted formats* `{{parent}}/customConstraints/{{name}}` When using the `pulumi import` command, CustomConstraint can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:orgpolicy/customConstraint:CustomConstraint default {{parent}}/customConstraints/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:orgpolicy/customConstraint:CustomConstraint")]
    public partial class CustomConstraint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The action to take if the condition is met.
        /// Possible values are: `ALLOW`, `DENY`.
        /// </summary>
        [Output("actionType")]
        public Output<string> ActionType { get; private set; } = null!;

        /// <summary>
        /// A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
        /// </summary>
        [Output("condition")]
        public Output<string> Condition { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description of the constraint to display as an error message when the policy is violated.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name for the constraint.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
        /// </summary>
        [Output("methodTypes")]
        public Output<ImmutableArray<string>> MethodTypes { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of the custom constraint. This is unique within the organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
        /// </summary>
        [Output("resourceTypes")]
        public Output<ImmutableArray<string>> ResourceTypes { get; private set; } = null!;

        /// <summary>
        /// Output only. The timestamp representing when the constraint was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CustomConstraint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomConstraint(string name, CustomConstraintArgs args, CustomResourceOptions? options = null)
            : base("gcp:orgpolicy/customConstraint:CustomConstraint", name, args ?? new CustomConstraintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomConstraint(string name, Input<string> id, CustomConstraintState? state = null, CustomResourceOptions? options = null)
            : base("gcp:orgpolicy/customConstraint:CustomConstraint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomConstraint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomConstraint Get(string name, Input<string> id, CustomConstraintState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomConstraint(name, id, state, options);
        }
    }

    public sealed class CustomConstraintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take if the condition is met.
        /// Possible values are: `ALLOW`, `DENY`.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// A human-friendly description of the constraint to display as an error message when the policy is violated.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A human-friendly name for the constraint.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("methodTypes", required: true)]
        private InputList<string>? _methodTypes;

        /// <summary>
        /// A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
        /// </summary>
        public InputList<string> MethodTypes
        {
            get => _methodTypes ?? (_methodTypes = new InputList<string>());
            set => _methodTypes = value;
        }

        /// <summary>
        /// Immutable. The name of the custom constraint. This is unique within the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        [Input("resourceTypes", required: true)]
        private InputList<string>? _resourceTypes;

        /// <summary>
        /// Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
        /// </summary>
        public InputList<string> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<string>());
            set => _resourceTypes = value;
        }

        public CustomConstraintArgs()
        {
        }
        public static new CustomConstraintArgs Empty => new CustomConstraintArgs();
    }

    public sealed class CustomConstraintState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action to take if the condition is met.
        /// Possible values are: `ALLOW`, `DENY`.
        /// </summary>
        [Input("actionType")]
        public Input<string>? ActionType { get; set; }

        /// <summary>
        /// A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// A human-friendly description of the constraint to display as an error message when the policy is violated.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A human-friendly name for the constraint.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("methodTypes")]
        private InputList<string>? _methodTypes;

        /// <summary>
        /// A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
        /// </summary>
        public InputList<string> MethodTypes
        {
            get => _methodTypes ?? (_methodTypes = new InputList<string>());
            set => _methodTypes = value;
        }

        /// <summary>
        /// Immutable. The name of the custom constraint. This is unique within the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        [Input("resourceTypes")]
        private InputList<string>? _resourceTypes;

        /// <summary>
        /// Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
        /// </summary>
        public InputList<string> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<string>());
            set => _resourceTypes = value;
        }

        /// <summary>
        /// Output only. The timestamp representing when the constraint was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CustomConstraintState()
        {
        }
        public static new CustomConstraintState Empty => new CustomConstraintState();
    }
}
