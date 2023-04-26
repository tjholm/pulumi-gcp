// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase
{
    /// <summary>
    /// ## Example Usage
    /// ### Firebasehosting Channel Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultHostingSite = new Gcp.Firebase.HostingSite("defaultHostingSite", new()
    ///     {
    ///         Project = "my-project-name",
    ///         SiteId = "site-with-channel",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultHostingChannel = new Gcp.Firebase.HostingChannel("defaultHostingChannel", new()
    ///     {
    ///         SiteId = defaultHostingSite.SiteId,
    ///         ChannelId = "channel-basic",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Firebasehosting Channel Full
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.Firebase.HostingSite("default", new()
    ///     {
    ///         Project = "my-project-name",
    ///         SiteId = "site-with-channel",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var full = new Gcp.Firebase.HostingChannel("full", new()
    ///     {
    ///         SiteId = @default.SiteId,
    ///         ChannelId = "channel-full",
    ///         Ttl = "86400s",
    ///         RetainedReleaseCount = 20,
    ///         Labels = 
    ///         {
    ///             { "some-key", "some-value" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Channel can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firebase/hostingChannel:HostingChannel default sites/{{site_id}}/channels/{{channel_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firebase/hostingChannel:HostingChannel default {{site_id}}/{{channel_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:firebase/hostingChannel:HostingChannel")]
    public partial class HostingChannel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. Immutable. A unique ID within the site that identifies the channel.
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel
        /// will not be automatically deleted. This field is present in the output whether it's
        /// set directly or via the `ttl` field.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Text labels used for extra metadata and/or filtering
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified resource name for the channel, in the format:
        /// sites/SITE_ID/channels/CHANNEL_ID
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other
        /// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        [Output("retainedReleaseCount")]
        public Output<int> RetainedReleaseCount { get; private set; } = null!;

        /// <summary>
        /// Required. The ID of the site in which to create this channel.
        /// </summary>
        [Output("siteId")]
        public Output<string> SiteId { get; private set; } = null!;

        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided
        /// duration past the time of the request. A duration in seconds with up to nine fractional
        /// digits, terminated by 's'. Example: "86400s" (one day).
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a HostingChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostingChannel(string name, HostingChannelArgs args, CustomResourceOptions? options = null)
            : base("gcp:firebase/hostingChannel:HostingChannel", name, args ?? new HostingChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostingChannel(string name, Input<string> id, HostingChannelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firebase/hostingChannel:HostingChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostingChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostingChannel Get(string name, Input<string> id, HostingChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new HostingChannel(name, id, state, options);
        }
    }

    public sealed class HostingChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Immutable. A unique ID within the site that identifies the channel.
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel
        /// will not be automatically deleted. This field is present in the output whether it's
        /// set directly or via the `ttl` field.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Text labels used for extra metadata and/or filtering
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other
        /// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        [Input("retainedReleaseCount")]
        public Input<int>? RetainedReleaseCount { get; set; }

        /// <summary>
        /// Required. The ID of the site in which to create this channel.
        /// </summary>
        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided
        /// duration past the time of the request. A duration in seconds with up to nine fractional
        /// digits, terminated by 's'. Example: "86400s" (one day).
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public HostingChannelArgs()
        {
        }
        public static new HostingChannelArgs Empty => new HostingChannelArgs();
    }

    public sealed class HostingChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Immutable. A unique ID within the site that identifies the channel.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// The time at which the channel will be automatically deleted. If null, the channel
        /// will not be automatically deleted. This field is present in the output whether it's
        /// set directly or via the `ttl` field.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Text labels used for extra metadata and/or filtering
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The fully-qualified resource name for the channel, in the format:
        /// sites/SITE_ID/channels/CHANNEL_ID
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of previous releases to retain on the channel for rollback or other
        /// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
        /// </summary>
        [Input("retainedReleaseCount")]
        public Input<int>? RetainedReleaseCount { get; set; }

        /// <summary>
        /// Required. The ID of the site in which to create this channel.
        /// </summary>
        [Input("siteId")]
        public Input<string>? SiteId { get; set; }

        /// <summary>
        /// Input only. A time-to-live for this channel. Sets `expire_time` to the provided
        /// duration past the time of the request. A duration in seconds with up to nine fractional
        /// digits, terminated by 's'. Example: "86400s" (one day).
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        public HostingChannelState()
        {
        }
        public static new HostingChannelState Empty => new HostingChannelState();
    }
}
