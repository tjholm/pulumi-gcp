// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static class GetBucketObject
    {
        /// <summary>
        /// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Example picture stored within a folder.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var picture = Output.Create(Gcp.Storage.GetBucketObject.InvokeAsync(new Gcp.Storage.GetBucketObjectArgs
        ///         {
        ///             Bucket = "image-store",
        ///             Name = "folder/butterfly01.jpg",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketObjectResult> InvokeAsync(GetBucketObjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectResult>("gcp:storage/getBucketObject:getBucketObject", args ?? new GetBucketObjectArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Example picture stored within a folder.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var picture = Output.Create(Gcp.Storage.GetBucketObject.InvokeAsync(new Gcp.Storage.GetBucketObjectArgs
        ///         {
        ///             Bucket = "image-store",
        ///             Name = "folder/butterfly01.jpg",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBucketObjectResult> Invoke(GetBucketObjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBucketObjectResult>("gcp:storage/getBucketObject:getBucketObject", args ?? new GetBucketObjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketObjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the containing bucket.
        /// </summary>
        [Input("bucket")]
        public string? Bucket { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetBucketObjectArgs()
        {
        }
    }

    public sealed class GetBucketObjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the containing bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetBucketObjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketObjectResult
    {
        public readonly string? Bucket;
        /// <summary>
        /// (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
        /// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
        /// </summary>
        public readonly string CacheControl;
        public readonly string Content;
        /// <summary>
        /// (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
        /// </summary>
        public readonly string ContentDisposition;
        /// <summary>
        /// (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
        /// </summary>
        public readonly string ContentEncoding;
        /// <summary>
        /// (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
        /// </summary>
        public readonly string ContentLanguage;
        /// <summary>
        /// (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// (Computed) Base 64 CRC32 hash of the uploaded data.
        /// </summary>
        public readonly string Crc32c;
        public readonly ImmutableArray<Outputs.GetBucketObjectCustomerEncryptionResult> CustomerEncryptions;
        public readonly string DetectMd5hash;
        public readonly bool EventBasedHold;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KmsKeyName;
        /// <summary>
        /// (Computed) Base 64 MD5 hash of the uploaded data.
        /// </summary>
        public readonly string Md5hash;
        /// <summary>
        /// (Computed) A url reference to download this object.
        /// </summary>
        public readonly string MediaLink;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly string? Name;
        public readonly string OutputName;
        /// <summary>
        /// (Computed) A url reference to this object.
        /// </summary>
        public readonly string SelfLink;
        public readonly string Source;
        /// <summary>
        /// (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
        /// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
        /// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
        /// </summary>
        public readonly string StorageClass;
        public readonly bool TemporaryHold;

        [OutputConstructor]
        private GetBucketObjectResult(
            string? bucket,

            string cacheControl,

            string content,

            string contentDisposition,

            string contentEncoding,

            string contentLanguage,

            string contentType,

            string crc32c,

            ImmutableArray<Outputs.GetBucketObjectCustomerEncryptionResult> customerEncryptions,

            string detectMd5hash,

            bool eventBasedHold,

            string id,

            string kmsKeyName,

            string md5hash,

            string mediaLink,

            ImmutableDictionary<string, string> metadata,

            string? name,

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
