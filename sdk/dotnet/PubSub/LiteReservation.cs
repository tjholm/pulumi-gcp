// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub
{
    /// <summary>
    /// A named resource representing a shared pool of capacity.
    /// 
    /// To get more information about Reservation, see:
    /// 
    /// * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.reservations)
    /// * How-to Guides
    ///     * [Managing Reservations](https://cloud.google.com/pubsub/lite/docs/reservations)
    /// 
    /// ## Example Usage
    /// ### Pubsub Lite Reservation Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var example = new Gcp.PubSub.LiteReservation("example", new()
    ///     {
    ///         Project = project.Apply(getProjectResult =&gt; getProjectResult.Number),
    ///         ThroughputCapacity = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Reservation can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/reservations/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Reservation can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default projects/{{project}}/locations/{{region}}/reservations/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:pubsub/liteReservation:LiteReservation default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:pubsub/liteReservation:LiteReservation")]
    public partial class LiteReservation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the reservation.
        /// 
        /// 
        /// - - -
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
        /// The region of the pubsub lite reservation.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The reserved throughput capacity. Every unit of throughput capacity is
        /// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
        /// messages.
        /// </summary>
        [Output("throughputCapacity")]
        public Output<int> ThroughputCapacity { get; private set; } = null!;


        /// <summary>
        /// Create a LiteReservation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LiteReservation(string name, LiteReservationArgs args, CustomResourceOptions? options = null)
            : base("gcp:pubsub/liteReservation:LiteReservation", name, args ?? new LiteReservationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LiteReservation(string name, Input<string> id, LiteReservationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:pubsub/liteReservation:LiteReservation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LiteReservation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LiteReservation Get(string name, Input<string> id, LiteReservationState? state = null, CustomResourceOptions? options = null)
        {
            return new LiteReservation(name, id, state, options);
        }
    }

    public sealed class LiteReservationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the reservation.
        /// 
        /// 
        /// - - -
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
        /// The region of the pubsub lite reservation.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The reserved throughput capacity. Every unit of throughput capacity is
        /// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
        /// messages.
        /// </summary>
        [Input("throughputCapacity", required: true)]
        public Input<int> ThroughputCapacity { get; set; } = null!;

        public LiteReservationArgs()
        {
        }
        public static new LiteReservationArgs Empty => new LiteReservationArgs();
    }

    public sealed class LiteReservationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the reservation.
        /// 
        /// 
        /// - - -
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
        /// The region of the pubsub lite reservation.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The reserved throughput capacity. Every unit of throughput capacity is
        /// equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed
        /// messages.
        /// </summary>
        [Input("throughputCapacity")]
        public Input<int>? ThroughputCapacity { get; set; }

        public LiteReservationState()
        {
        }
        public static new LiteReservationState Empty => new LiteReservationState();
    }
}
