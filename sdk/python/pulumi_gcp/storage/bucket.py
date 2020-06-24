# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Bucket(pulumi.CustomResource):
    bucket_policy_only: pulumi.Output[bool]
    """
    Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
    """
    cors: pulumi.Output[list]
    """
    The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.

      * `maxAgeSeconds` (`float`) - The value, in seconds, to return in the [Access-Control-Max-Age header](https://www.w3.org/TR/cors/#access-control-max-age-response-header) used in preflight responses.
      * `methods` (`list`) - The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
      * `origins` (`list`) - The list of [Origins](https://tools.ietf.org/html/rfc6454) eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
      * `responseHeaders` (`list`) - The list of HTTP headers other than the [simple response headers](https://www.w3.org/TR/cors/#simple-response-header) to give permission for the user-agent to share across domains.
    """
    default_event_based_hold: pulumi.Output[bool]
    encryption: pulumi.Output[dict]
    """
    The bucket's encryption configuration.

      * `defaultKmsKeyName` (`str`)
    """
    force_destroy: pulumi.Output[bool]
    """
    When deleting a bucket, this
    boolean option will delete all contained objects. If you try to delete a
    bucket that contains objects, the provider will fail that run.
    """
    labels: pulumi.Output[dict]
    """
    A set of key/value label pairs to assign to the bucket.
    """
    lifecycle_rules: pulumi.Output[list]
    """
    The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.

      * `action` (`dict`) - The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
        * `storage_class` (`str`) - The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        * `type` (`str`) - The type of the action of this Lifecycle Rule. Supported values include: `Delete` and `SetStorageClass`.

      * `condition` (`dict`) - The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
        * `age` (`float`) - Minimum age of an object in days to satisfy this condition.
        * `createdBefore` (`str`) - Creation date of an object in RFC 3339 (e.g. `2017-06-13`) to satisfy this condition.
        * `matchesStorageClasses` (`list`) - [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects to satisfy this condition. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `STANDARD`, `DURABLE_REDUCED_AVAILABILITY`.
        * `numNewerVersions` (`float`) - Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
        * `withState` (`str`) - Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: `"LIVE"`, `"ARCHIVED"`, `"ANY"`.
    """
    location: pulumi.Output[str]
    """
    The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
    """
    logging: pulumi.Output[dict]
    """
    The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.

      * `logBucket` (`str`) - The bucket that will receive log objects.
      * `logObjectPrefix` (`str`) - The object prefix for log objects. If it's not provided,
        by default GCS sets this to this bucket's name.
    """
    name: pulumi.Output[str]
    """
    The name of the bucket.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    requester_pays: pulumi.Output[bool]
    """
    Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
    """
    retention_policy: pulumi.Output[dict]
    """
    Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.

      * `isLocked` (`bool`) - If set to `true`, the bucket will be [locked](https://cloud.google.com/storage/docs/using-bucket-lock#lock-bucket) and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
      * `retentionPeriod` (`float`) - The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    storage_class: pulumi.Output[str]
    """
    The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
    """
    url: pulumi.Output[str]
    """
    The base URL of the bucket, in the format `gs://<bucket-name>`.
    """
    versioning: pulumi.Output[dict]
    """
    The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.

      * `enabled` (`bool`) - While set to `true`, versioning is fully enabled for this bucket.
    """
    website: pulumi.Output[dict]
    """
    Configuration if the bucket acts as a website. Structure is documented below.

      * `mainPageSuffix` (`str`) - Behaves as the bucket's directory index where
        missing objects are treated as potential directories.
      * `notFoundPage` (`str`) - The custom object to return when a requested
        resource is not found.
    """
    def __init__(__self__, resource_name, opts=None, bucket_policy_only=None, cors=None, default_event_based_hold=None, encryption=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, requester_pays=None, retention_policy=None, storage_class=None, versioning=None, website=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a new bucket in Google cloud storage service (GCS).
        Once a bucket has been created, its location can't be changed.
        [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
        using the [`storage.BucketACL`](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html) resource.

        For more information see
        [the official documentation](https://cloud.google.com/storage/docs/overview)
        and
        [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).

        **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
        determined which will require enabling the compute api.


        ## Example Usage

        ### Creating A Private Bucket In Standard Storage, In The EU Region. Bucket Configured As Static Website And CORS Configurations

        ```python
        import pulumi
        import pulumi_gcp as gcp

        static_site = gcp.storage.Bucket("static-site",
            bucket_policy_only=True,
            cors=[{
                "maxAgeSeconds": 3600,
                "method": [
                    "GET",
                    "HEAD",
                    "PUT",
                    "POST",
                    "DELETE",
                ],
                "origin": ["http://image-store.com"],
                "responseHeader": ["*"],
            }],
            force_destroy=True,
            location="EU",
            website={
                "mainPageSuffix": "index.html",
                "notFoundPage": "404.html",
            })
        ```

        ### Life Cycle Settings For Storage Bucket Objects

        ```python
        import pulumi
        import pulumi_gcp as gcp

        auto_expire = gcp.storage.Bucket("auto-expire",
            force_destroy=True,
            lifecycle_rules=[{
                "action": {
                    "type": "Delete",
                },
                "condition": {
                    "age": "3",
                },
            }],
            location="US")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        :param pulumi.Input[list] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[dict] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, the provider will fail that run.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[list] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[dict] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[dict] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] storage_class: The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        :param pulumi.Input[dict] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[dict] website: Configuration if the bucket acts as a website. Structure is documented below.

        The **cors** object supports the following:

          * `maxAgeSeconds` (`pulumi.Input[float]`) - The value, in seconds, to return in the [Access-Control-Max-Age header](https://www.w3.org/TR/cors/#access-control-max-age-response-header) used in preflight responses.
          * `methods` (`pulumi.Input[list]`) - The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
          * `origins` (`pulumi.Input[list]`) - The list of [Origins](https://tools.ietf.org/html/rfc6454) eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
          * `responseHeaders` (`pulumi.Input[list]`) - The list of HTTP headers other than the [simple response headers](https://www.w3.org/TR/cors/#simple-response-header) to give permission for the user-agent to share across domains.

        The **encryption** object supports the following:

          * `defaultKmsKeyName` (`pulumi.Input[str]`)

        The **lifecycle_rules** object supports the following:

          * `action` (`pulumi.Input[dict]`) - The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
            * `storage_class` (`pulumi.Input[str]`) - The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
            * `type` (`pulumi.Input[str]`) - The type of the action of this Lifecycle Rule. Supported values include: `Delete` and `SetStorageClass`.

          * `condition` (`pulumi.Input[dict]`) - The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
            * `age` (`pulumi.Input[float]`) - Minimum age of an object in days to satisfy this condition.
            * `createdBefore` (`pulumi.Input[str]`) - Creation date of an object in RFC 3339 (e.g. `2017-06-13`) to satisfy this condition.
            * `matchesStorageClasses` (`pulumi.Input[list]`) - [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects to satisfy this condition. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `STANDARD`, `DURABLE_REDUCED_AVAILABILITY`.
            * `numNewerVersions` (`pulumi.Input[float]`) - Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
            * `withState` (`pulumi.Input[str]`) - Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: `"LIVE"`, `"ARCHIVED"`, `"ANY"`.

        The **logging** object supports the following:

          * `logBucket` (`pulumi.Input[str]`) - The bucket that will receive log objects.
          * `logObjectPrefix` (`pulumi.Input[str]`) - The object prefix for log objects. If it's not provided,
            by default GCS sets this to this bucket's name.

        The **retention_policy** object supports the following:

          * `isLocked` (`pulumi.Input[bool]`) - If set to `true`, the bucket will be [locked](https://cloud.google.com/storage/docs/using-bucket-lock#lock-bucket) and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
          * `retentionPeriod` (`pulumi.Input[float]`) - The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.

        The **versioning** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - While set to `true`, versioning is fully enabled for this bucket.

        The **website** object supports the following:

          * `mainPageSuffix` (`pulumi.Input[str]`) - Behaves as the bucket's directory index where
            missing objects are treated as potential directories.
          * `notFoundPage` (`pulumi.Input[str]`) - The custom object to return when a requested
            resource is not found.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['bucket_policy_only'] = bucket_policy_only
            __props__['cors'] = cors
            __props__['default_event_based_hold'] = default_event_based_hold
            __props__['encryption'] = encryption
            __props__['force_destroy'] = force_destroy
            __props__['labels'] = labels
            __props__['lifecycle_rules'] = lifecycle_rules
            __props__['location'] = location
            __props__['logging'] = logging
            __props__['name'] = name
            __props__['project'] = project
            __props__['requester_pays'] = requester_pays
            __props__['retention_policy'] = retention_policy
            __props__['storage_class'] = storage_class
            __props__['versioning'] = versioning
            __props__['website'] = website
            __props__['self_link'] = None
            __props__['url'] = None
        super(Bucket, __self__).__init__(
            'gcp:storage/bucket:Bucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bucket_policy_only=None, cors=None, default_event_based_hold=None, encryption=None, force_destroy=None, labels=None, lifecycle_rules=None, location=None, logging=None, name=None, project=None, requester_pays=None, retention_policy=None, self_link=None, storage_class=None, url=None, versioning=None, website=None):
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bucket_policy_only: Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        :param pulumi.Input[list] cors: The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[dict] encryption: The bucket's encryption configuration.
        :param pulumi.Input[bool] force_destroy: When deleting a bucket, this
               boolean option will delete all contained objects. If you try to delete a
               bucket that contains objects, the provider will fail that run.
        :param pulumi.Input[dict] labels: A set of key/value label pairs to assign to the bucket.
        :param pulumi.Input[list] lifecycle_rules: The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        :param pulumi.Input[str] location: The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        :param pulumi.Input[dict] logging: The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        :param pulumi.Input[str] name: The name of the bucket.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] requester_pays: Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        :param pulumi.Input[dict] retention_policy: Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] storage_class: The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        :param pulumi.Input[str] url: The base URL of the bucket, in the format `gs://<bucket-name>`.
        :param pulumi.Input[dict] versioning: The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        :param pulumi.Input[dict] website: Configuration if the bucket acts as a website. Structure is documented below.

        The **cors** object supports the following:

          * `maxAgeSeconds` (`pulumi.Input[float]`) - The value, in seconds, to return in the [Access-Control-Max-Age header](https://www.w3.org/TR/cors/#access-control-max-age-response-header) used in preflight responses.
          * `methods` (`pulumi.Input[list]`) - The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
          * `origins` (`pulumi.Input[list]`) - The list of [Origins](https://tools.ietf.org/html/rfc6454) eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
          * `responseHeaders` (`pulumi.Input[list]`) - The list of HTTP headers other than the [simple response headers](https://www.w3.org/TR/cors/#simple-response-header) to give permission for the user-agent to share across domains.

        The **encryption** object supports the following:

          * `defaultKmsKeyName` (`pulumi.Input[str]`)

        The **lifecycle_rules** object supports the following:

          * `action` (`pulumi.Input[dict]`) - The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
            * `storage_class` (`pulumi.Input[str]`) - The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
            * `type` (`pulumi.Input[str]`) - The type of the action of this Lifecycle Rule. Supported values include: `Delete` and `SetStorageClass`.

          * `condition` (`pulumi.Input[dict]`) - The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
            * `age` (`pulumi.Input[float]`) - Minimum age of an object in days to satisfy this condition.
            * `createdBefore` (`pulumi.Input[str]`) - Creation date of an object in RFC 3339 (e.g. `2017-06-13`) to satisfy this condition.
            * `matchesStorageClasses` (`pulumi.Input[list]`) - [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects to satisfy this condition. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `STANDARD`, `DURABLE_REDUCED_AVAILABILITY`.
            * `numNewerVersions` (`pulumi.Input[float]`) - Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
            * `withState` (`pulumi.Input[str]`) - Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: `"LIVE"`, `"ARCHIVED"`, `"ANY"`.

        The **logging** object supports the following:

          * `logBucket` (`pulumi.Input[str]`) - The bucket that will receive log objects.
          * `logObjectPrefix` (`pulumi.Input[str]`) - The object prefix for log objects. If it's not provided,
            by default GCS sets this to this bucket's name.

        The **retention_policy** object supports the following:

          * `isLocked` (`pulumi.Input[bool]`) - If set to `true`, the bucket will be [locked](https://cloud.google.com/storage/docs/using-bucket-lock#lock-bucket) and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
          * `retentionPeriod` (`pulumi.Input[float]`) - The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.

        The **versioning** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - While set to `true`, versioning is fully enabled for this bucket.

        The **website** object supports the following:

          * `mainPageSuffix` (`pulumi.Input[str]`) - Behaves as the bucket's directory index where
            missing objects are treated as potential directories.
          * `notFoundPage` (`pulumi.Input[str]`) - The custom object to return when a requested
            resource is not found.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket_policy_only"] = bucket_policy_only
        __props__["cors"] = cors
        __props__["default_event_based_hold"] = default_event_based_hold
        __props__["encryption"] = encryption
        __props__["force_destroy"] = force_destroy
        __props__["labels"] = labels
        __props__["lifecycle_rules"] = lifecycle_rules
        __props__["location"] = location
        __props__["logging"] = logging
        __props__["name"] = name
        __props__["project"] = project
        __props__["requester_pays"] = requester_pays
        __props__["retention_policy"] = retention_policy
        __props__["self_link"] = self_link
        __props__["storage_class"] = storage_class
        __props__["url"] = url
        __props__["versioning"] = versioning
        __props__["website"] = website
        return Bucket(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

