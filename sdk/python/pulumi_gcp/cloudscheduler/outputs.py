# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'JobAppEngineHttpTarget',
    'JobAppEngineHttpTargetAppEngineRouting',
    'JobHttpTarget',
    'JobHttpTargetOauthToken',
    'JobHttpTargetOidcToken',
    'JobPubsubTarget',
    'JobRetryConfig',
]

@pulumi.output_type
class JobAppEngineHttpTarget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "relativeUri":
            suggest = "relative_uri"
        elif key == "appEngineRouting":
            suggest = "app_engine_routing"
        elif key == "httpMethod":
            suggest = "http_method"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobAppEngineHttpTarget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobAppEngineHttpTarget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobAppEngineHttpTarget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 relative_uri: str,
                 app_engine_routing: Optional['outputs.JobAppEngineHttpTargetAppEngineRouting'] = None,
                 body: Optional[str] = None,
                 headers: Optional[Mapping[str, str]] = None,
                 http_method: Optional[str] = None):
        """
        :param str relative_uri: The relative URI.
               The relative URL must begin with "/" and must be a valid HTTP relative URL.
               It can contain a path, query string arguments, and \\# fragments.
               If the relative URL is empty, then the root path "/" will be used.
               No spaces are allowed, and the maximum length allowed is 2083 characters
        :param 'JobAppEngineHttpTargetAppEngineRoutingArgs' app_engine_routing: App Engine Routing setting for the job.
               Structure is documented below.
        :param str body: HTTP request body.
               A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
               It is an error to set body on a job with an incompatible HttpMethod.
               A base64-encoded string.
        :param Mapping[str, str] headers: This map contains the header field names and values.
               Repeated headers are not supported, but a header value can contain commas.
        :param str http_method: Which HTTP method to use for the request.
        """
        pulumi.set(__self__, "relative_uri", relative_uri)
        if app_engine_routing is not None:
            pulumi.set(__self__, "app_engine_routing", app_engine_routing)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)

    @property
    @pulumi.getter(name="relativeUri")
    def relative_uri(self) -> str:
        """
        The relative URI.
        The relative URL must begin with "/" and must be a valid HTTP relative URL.
        It can contain a path, query string arguments, and \\# fragments.
        If the relative URL is empty, then the root path "/" will be used.
        No spaces are allowed, and the maximum length allowed is 2083 characters
        """
        return pulumi.get(self, "relative_uri")

    @property
    @pulumi.getter(name="appEngineRouting")
    def app_engine_routing(self) -> Optional['outputs.JobAppEngineHttpTargetAppEngineRouting']:
        """
        App Engine Routing setting for the job.
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine_routing")

    @property
    @pulumi.getter
    def body(self) -> Optional[str]:
        """
        HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
        A base64-encoded string.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def headers(self) -> Optional[Mapping[str, str]]:
        """
        This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[str]:
        """
        Which HTTP method to use for the request.
        """
        return pulumi.get(self, "http_method")


@pulumi.output_type
class JobAppEngineHttpTargetAppEngineRouting(dict):
    def __init__(__self__, *,
                 instance: Optional[str] = None,
                 service: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str instance: App instance.
               By default, the job is sent to an instance which is available when the job is attempted.
        :param str service: App service.
               By default, the job is sent to the service which is the default service when the job is attempted.
        :param str version: App version.
               By default, the job is sent to the version which is the default version when the job is attempted.
        """
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        """
        App instance.
        By default, the job is sent to an instance which is available when the job is attempted.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def service(self) -> Optional[str]:
        """
        App service.
        By default, the job is sent to the service which is the default service when the job is attempted.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        App version.
        By default, the job is sent to the version which is the default version when the job is attempted.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class JobHttpTarget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "httpMethod":
            suggest = "http_method"
        elif key == "oauthToken":
            suggest = "oauth_token"
        elif key == "oidcToken":
            suggest = "oidc_token"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobHttpTarget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobHttpTarget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobHttpTarget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 uri: str,
                 body: Optional[str] = None,
                 headers: Optional[Mapping[str, str]] = None,
                 http_method: Optional[str] = None,
                 oauth_token: Optional['outputs.JobHttpTargetOauthToken'] = None,
                 oidc_token: Optional['outputs.JobHttpTargetOidcToken'] = None):
        """
        :param str uri: The full URI path that the request will be sent to.
        :param str body: HTTP request body.
               A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
               It is an error to set body on a job with an incompatible HttpMethod.
               A base64-encoded string.
        :param Mapping[str, str] headers: This map contains the header field names and values.
               Repeated headers are not supported, but a header value can contain commas.
        :param str http_method: Which HTTP method to use for the request.
        :param 'JobHttpTargetOauthTokenArgs' oauth_token: Contains information needed for generating an OAuth token.
               This type of authorization should be used when sending requests to a GCP endpoint.
               Structure is documented below.
        :param 'JobHttpTargetOidcTokenArgs' oidc_token: Contains information needed for generating an OpenID Connect token.
               This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
               Structure is documented below.
        """
        pulumi.set(__self__, "uri", uri)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if oauth_token is not None:
            pulumi.set(__self__, "oauth_token", oauth_token)
        if oidc_token is not None:
            pulumi.set(__self__, "oidc_token", oidc_token)

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        The full URI path that the request will be sent to.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def body(self) -> Optional[str]:
        """
        HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
        A base64-encoded string.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def headers(self) -> Optional[Mapping[str, str]]:
        """
        This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[str]:
        """
        Which HTTP method to use for the request.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="oauthToken")
    def oauth_token(self) -> Optional['outputs.JobHttpTargetOauthToken']:
        """
        Contains information needed for generating an OAuth token.
        This type of authorization should be used when sending requests to a GCP endpoint.
        Structure is documented below.
        """
        return pulumi.get(self, "oauth_token")

    @property
    @pulumi.getter(name="oidcToken")
    def oidc_token(self) -> Optional['outputs.JobHttpTargetOidcToken']:
        """
        Contains information needed for generating an OpenID Connect token.
        This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
        Structure is documented below.
        """
        return pulumi.get(self, "oidc_token")


@pulumi.output_type
class JobHttpTargetOauthToken(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceAccountEmail":
            suggest = "service_account_email"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobHttpTargetOauthToken. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobHttpTargetOauthToken.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobHttpTargetOauthToken.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 service_account_email: str,
                 scope: Optional[str] = None):
        """
        :param str service_account_email: Service account email to be used for generating OAuth token.
               The service account must be within the same project as the job.
        :param str scope: OAuth scope to be used for generating OAuth access token. If not specified,
               "https://www.googleapis.com/auth/cloud-platform" will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        Service account email to be used for generating OAuth token.
        The service account must be within the same project as the job.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        OAuth scope to be used for generating OAuth access token. If not specified,
        "https://www.googleapis.com/auth/cloud-platform" will be used.
        """
        return pulumi.get(self, "scope")


@pulumi.output_type
class JobHttpTargetOidcToken(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceAccountEmail":
            suggest = "service_account_email"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobHttpTargetOidcToken. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobHttpTargetOidcToken.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobHttpTargetOidcToken.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 service_account_email: str,
                 audience: Optional[str] = None):
        """
        :param str service_account_email: Service account email to be used for generating OAuth token.
               The service account must be within the same project as the job.
        :param str audience: Audience to be used when generating OIDC token. If not specified,
               the URI specified in target will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> str:
        """
        Service account email to be used for generating OAuth token.
        The service account must be within the same project as the job.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def audience(self) -> Optional[str]:
        """
        Audience to be used when generating OIDC token. If not specified,
        the URI specified in target will be used.
        """
        return pulumi.get(self, "audience")


@pulumi.output_type
class JobPubsubTarget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "topicName":
            suggest = "topic_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobPubsubTarget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobPubsubTarget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobPubsubTarget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 topic_name: str,
                 attributes: Optional[Mapping[str, str]] = None,
                 data: Optional[str] = None):
        """
        :param str topic_name: The full resource name for the Cloud Pub/Sub topic to which
               messages will be published when a job is delivered. ~>**NOTE:**
               The topic name must be in the same format as required by PubSub's
               PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
        :param Mapping[str, str] attributes: Attributes for PubsubMessage.
               Pubsub message must contain either non-empty data, or at least one attribute.
        :param str data: The message payload for PubsubMessage.
               Pubsub message must contain either non-empty data, or at least one attribute.
               A base64-encoded string.
        """
        pulumi.set(__self__, "topic_name", topic_name)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> str:
        """
        The full resource name for the Cloud Pub/Sub topic to which
        messages will be published when a job is delivered. ~>**NOTE:**
        The topic name must be in the same format as required by PubSub's
        PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
        """
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Mapping[str, str]]:
        """
        Attributes for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def data(self) -> Optional[str]:
        """
        The message payload for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
        A base64-encoded string.
        """
        return pulumi.get(self, "data")


@pulumi.output_type
class JobRetryConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxBackoffDuration":
            suggest = "max_backoff_duration"
        elif key == "maxDoublings":
            suggest = "max_doublings"
        elif key == "maxRetryDuration":
            suggest = "max_retry_duration"
        elif key == "minBackoffDuration":
            suggest = "min_backoff_duration"
        elif key == "retryCount":
            suggest = "retry_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in JobRetryConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        JobRetryConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        JobRetryConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_backoff_duration: Optional[str] = None,
                 max_doublings: Optional[int] = None,
                 max_retry_duration: Optional[str] = None,
                 min_backoff_duration: Optional[str] = None,
                 retry_count: Optional[int] = None):
        """
        :param str max_backoff_duration: The maximum amount of time to wait before retrying a job after it fails.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param int max_doublings: The time between retries will double maxDoublings times.
               A job's retry interval starts at minBackoffDuration,
               then doubles maxDoublings times, then increases linearly,
               and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
        :param str max_retry_duration: The time limit for retrying a failed job, measured from time when an execution was first attempted.
               If specified with retryCount, the job will be retried until both limits are reached.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param str min_backoff_duration: The minimum amount of time to wait before retrying a job after it fails.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param int retry_count: The number of attempts that the system will make to run a
               job using the exponential backoff procedure described by maxDoublings.
               Values greater than 5 and negative values are not allowed.
        """
        if max_backoff_duration is not None:
            pulumi.set(__self__, "max_backoff_duration", max_backoff_duration)
        if max_doublings is not None:
            pulumi.set(__self__, "max_doublings", max_doublings)
        if max_retry_duration is not None:
            pulumi.set(__self__, "max_retry_duration", max_retry_duration)
        if min_backoff_duration is not None:
            pulumi.set(__self__, "min_backoff_duration", min_backoff_duration)
        if retry_count is not None:
            pulumi.set(__self__, "retry_count", retry_count)

    @property
    @pulumi.getter(name="maxBackoffDuration")
    def max_backoff_duration(self) -> Optional[str]:
        """
        The maximum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "max_backoff_duration")

    @property
    @pulumi.getter(name="maxDoublings")
    def max_doublings(self) -> Optional[int]:
        """
        The time between retries will double maxDoublings times.
        A job's retry interval starts at minBackoffDuration,
        then doubles maxDoublings times, then increases linearly,
        and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
        """
        return pulumi.get(self, "max_doublings")

    @property
    @pulumi.getter(name="maxRetryDuration")
    def max_retry_duration(self) -> Optional[str]:
        """
        The time limit for retrying a failed job, measured from time when an execution was first attempted.
        If specified with retryCount, the job will be retried until both limits are reached.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "max_retry_duration")

    @property
    @pulumi.getter(name="minBackoffDuration")
    def min_backoff_duration(self) -> Optional[str]:
        """
        The minimum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "min_backoff_duration")

    @property
    @pulumi.getter(name="retryCount")
    def retry_count(self) -> Optional[int]:
        """
        The number of attempts that the system will make to run a
        job using the exponential backoff procedure described by maxDoublings.
        Values greater than 5 and negative values are not allowed.
        """
        return pulumi.get(self, "retry_count")


