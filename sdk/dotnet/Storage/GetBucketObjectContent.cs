// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static class GetBucketObjectContent
    {
        /// <summary>
        /// Gets an existing object content inside an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        /// 
        /// &gt; **Warning:** The object content will be saved in the state, and visiable to everyone who has access to the state file.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Example file object  stored within a folder.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var key = Gcp.Storage.GetBucketObjectContent.Invoke(new()
        ///     {
        ///         Name = "encryptedkey",
        ///         Bucket = "keystore",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["encrypted"] = key.Apply(getBucketObjectContentResult =&gt; getBucketObjectContentResult.Content),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketObjectContentResult> InvokeAsync(GetBucketObjectContentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectContentResult>("gcp:storage/getBucketObjectContent:getBucketObjectContent", args ?? new GetBucketObjectContentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing object content inside an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        /// 
        /// &gt; **Warning:** The object content will be saved in the state, and visiable to everyone who has access to the state file.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Example file object  stored within a folder.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var key = Gcp.Storage.GetBucketObjectContent.Invoke(new()
        ///     {
        ///         Name = "encryptedkey",
        ///         Bucket = "keystore",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["encrypted"] = key.Apply(getBucketObjectContentResult =&gt; getBucketObjectContentResult.Content),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBucketObjectContentResult> Invoke(GetBucketObjectContentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketObjectContentResult>("gcp:storage/getBucketObjectContent:getBucketObjectContent", args ?? new GetBucketObjectContentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketObjectContentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the containing bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        /// <summary>
        /// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
        /// </summary>
        [Input("content")]
        public string? Content { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBucketObjectContentArgs()
        {
        }
        public static new GetBucketObjectContentArgs Empty => new GetBucketObjectContentArgs();
    }

    public sealed class GetBucketObjectContentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the containing bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetBucketObjectContentInvokeArgs()
        {
        }
        public static new GetBucketObjectContentInvokeArgs Empty => new GetBucketObjectContentInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketObjectContentResult
    {
        public readonly string Bucket;
        public readonly string CacheControl;
        /// <summary>
        /// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object content.
        /// </summary>
        public readonly string? Content;
        public readonly string ContentDisposition;
        public readonly string ContentEncoding;
        public readonly string ContentLanguage;
        public readonly string ContentType;
        public readonly string Crc32c;
        public readonly ImmutableArray<Outputs.GetBucketObjectContentCustomerEncryptionResult> CustomerEncryptions;
        public readonly string DetectMd5hash;
        public readonly bool EventBasedHold;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KmsKeyName;
        public readonly string Md5hash;
        public readonly string MediaLink;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string Name;
        public readonly string OutputName;
        public readonly string SelfLink;
        public readonly string Source;
        public readonly string StorageClass;
        public readonly bool TemporaryHold;

        [OutputConstructor]
        private GetBucketObjectContentResult(
            string bucket,

            string cacheControl,

            string? content,

            string contentDisposition,

            string contentEncoding,

            string contentLanguage,

            string contentType,

            string crc32c,

            ImmutableArray<Outputs.GetBucketObjectContentCustomerEncryptionResult> customerEncryptions,

            string detectMd5hash,

            bool eventBasedHold,

            string id,

            string kmsKeyName,

            string md5hash,

            string mediaLink,

            ImmutableDictionary<string, string> metadata,

            string name,

            string outputName,

            string selfLink,

            string source,

            string storageClass,

            bool temporaryHold)
        {
            Bucket = bucket;
            CacheControl = cacheControl;
            Content = content;
            ContentDisposition = contentDisposition;
            ContentEncoding = contentEncoding;
            ContentLanguage = contentLanguage;
            ContentType = contentType;
            Crc32c = crc32c;
            CustomerEncryptions = customerEncryptions;
            DetectMd5hash = detectMd5hash;
            EventBasedHold = eventBasedHold;
            Id = id;
            KmsKeyName = kmsKeyName;
            Md5hash = md5hash;
            MediaLink = mediaLink;
            Metadata = metadata;
            Name = name;
            OutputName = outputName;
            SelfLink = selfLink;
            Source = source;
            StorageClass = storageClass;
            TemporaryHold = temporaryHold;
        }
    }
}
