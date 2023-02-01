// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase
{
    public static class GetWebAppConfig
    {
        public static Task<GetWebAppConfigResult> InvokeAsync(GetWebAppConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppConfigResult>("gcp:firebase/getWebAppConfig:getWebAppConfig", args ?? new GetWebAppConfigArgs(), options.WithDefaults());

        public static Output<GetWebAppConfigResult> Invoke(GetWebAppConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppConfigResult>("gcp:firebase/getWebAppConfig:getWebAppConfig", args ?? new GetWebAppConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// the id of the firebase web app
        /// </summary>
        [Input("webAppId", required: true)]
        public string WebAppId { get; set; } = null!;

        public GetWebAppConfigArgs()
        {
        }
        public static new GetWebAppConfigArgs Empty => new GetWebAppConfigArgs();
    }

    public sealed class GetWebAppConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// the id of the firebase web app
        /// </summary>
        [Input("webAppId", required: true)]
        public Input<string> WebAppId { get; set; } = null!;

        public GetWebAppConfigInvokeArgs()
        {
        }
        public static new GetWebAppConfigInvokeArgs Empty => new GetWebAppConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppConfigResult
    {
        /// <summary>
        /// The API key associated with the web App.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// The domain Firebase Auth configures for OAuth redirects, in the format:
        /// projectId.firebaseapp.com
        /// </summary>
        public readonly string AuthDomain;
        /// <summary>
        /// The default Firebase Realtime Database URL.
        /// </summary>
        public readonly string DatabaseUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the project's default GCP resource location. The location is one of the available GCP resource
        /// locations.
        /// This field is omitted if the default GCP resource location has not been finalized yet. To set your project's
        /// default GCP resource location, call defaultLocation.finalize after you add Firebase services to your project.
        /// </summary>
        public readonly string LocationId;
        /// <summary>
        /// The unique Google-assigned identifier of the Google Analytics web stream associated with the Firebase Web App.
        /// Firebase SDKs use this ID to interact with Google Analytics APIs.
        /// This field is only present if the App is linked to a web stream in a Google Analytics App + Web property.
        /// Learn more about this ID and Google Analytics web streams in the Analytics documentation.
        /// To generate a measurementId and link the Web App with a Google Analytics web stream,
        /// call projects.addGoogleAnalytics.
        /// </summary>
        public readonly string MeasurementId;
        /// <summary>
        /// The sender ID for use with Firebase Cloud Messaging.
        /// </summary>
        public readonly string MessagingSenderId;
        public readonly string? Project;
        /// <summary>
        /// The default Cloud Storage for Firebase storage bucket name.
        /// </summary>
        public readonly string StorageBucket;
        public readonly string WebAppId;

        [OutputConstructor]
        private GetWebAppConfigResult(
            string apiKey,

            string authDomain,

            string databaseUrl,

            string id,

            string locationId,

            string measurementId,

            string messagingSenderId,

            string? project,

            string storageBucket,

            string webAppId)
        {
            ApiKey = apiKey;
            AuthDomain = authDomain;
            DatabaseUrl = databaseUrl;
            Id = id;
            LocationId = locationId;
            MeasurementId = measurementId;
            MessagingSenderId = messagingSenderId;
            Project = project;
            StorageBucket = storageBucket;
            WebAppId = webAppId;
        }
    }
}
