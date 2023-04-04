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
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 location_id: pulumi.Input[str],
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 feature_settings: Optional[pulumi.Input['ApplicationFeatureSettingsArgs']] = None,
                 iap: Optional[pulumi.Input['ApplicationIapArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input['ApplicationFeatureSettingsArgs'] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input['ApplicationIapArgs'] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        """
        pulumi.set(__self__, "location_id", location_id)
        if auth_domain is not None:
            pulumi.set(__self__, "auth_domain", auth_domain)
        if database_type is not None:
            pulumi.set(__self__, "database_type", database_type)
        if feature_settings is not None:
            pulumi.set(__self__, "feature_settings", feature_settings)
        if iap is not None:
            pulumi.set(__self__, "iap", iap)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if serving_status is not None:
            pulumi.set(__self__, "serving_status", serving_status)

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> pulumi.Input[str]:
        """
        The [location](https://cloud.google.com/appengine/docs/locations)
        to serve the app from.
        """
        return pulumi.get(self, "location_id")

    @location_id.setter
    def location_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "location_id", value)

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain to authenticate users with when using App Engine's User API.
        """
        return pulumi.get(self, "auth_domain")

    @auth_domain.setter
    def auth_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_domain", value)

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_type")

    @database_type.setter
    def database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_type", value)

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> Optional[pulumi.Input['ApplicationFeatureSettingsArgs']]:
        """
        A block of optional settings to configure specific App Engine features:
        """
        return pulumi.get(self, "feature_settings")

    @feature_settings.setter
    def feature_settings(self, value: Optional[pulumi.Input['ApplicationFeatureSettingsArgs']]):
        pulumi.set(self, "feature_settings", value)

    @property
    @pulumi.getter
    def iap(self) -> Optional[pulumi.Input['ApplicationIapArgs']]:
        """
        Settings for enabling Cloud Identity Aware Proxy
        """
        return pulumi.get(self, "iap")

    @iap.setter
    def iap(self, value: Optional[pulumi.Input['ApplicationIapArgs']]):
        pulumi.set(self, "iap", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID to create the application under.
        ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
        you may get a "Permission denied" error.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> Optional[pulumi.Input[str]]:
        """
        The serving status of the app.
        """
        return pulumi.get(self, "serving_status")

    @serving_status.setter
    def serving_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serving_status", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 code_bucket: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 default_bucket: Optional[pulumi.Input[str]] = None,
                 default_hostname: Optional[pulumi.Input[str]] = None,
                 feature_settings: Optional[pulumi.Input['ApplicationFeatureSettingsArgs']] = None,
                 gcr_domain: Optional[pulumi.Input[str]] = None,
                 iap: Optional[pulumi.Input['ApplicationIapArgs']] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 url_dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationUrlDispatchRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[str] app_id: Identifier of the app, usually `{PROJECT_ID}`
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[str] code_bucket: The GCS bucket code is being stored in for this app.
        :param pulumi.Input[str] default_bucket: The GCS bucket content is being stored in for this app.
        :param pulumi.Input[str] default_hostname: The default hostname for this app.
        :param pulumi.Input['ApplicationFeatureSettingsArgs'] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[str] gcr_domain: The GCR domain used for storing managed Docker images for this app.
        :param pulumi.Input['ApplicationIapArgs'] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] name: Unique name of the app, usually `apps/{PROJECT_ID}`
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationUrlDispatchRuleArgs']]] url_dispatch_rules: A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if auth_domain is not None:
            pulumi.set(__self__, "auth_domain", auth_domain)
        if code_bucket is not None:
            pulumi.set(__self__, "code_bucket", code_bucket)
        if database_type is not None:
            pulumi.set(__self__, "database_type", database_type)
        if default_bucket is not None:
            pulumi.set(__self__, "default_bucket", default_bucket)
        if default_hostname is not None:
            pulumi.set(__self__, "default_hostname", default_hostname)
        if feature_settings is not None:
            pulumi.set(__self__, "feature_settings", feature_settings)
        if gcr_domain is not None:
            pulumi.set(__self__, "gcr_domain", gcr_domain)
        if iap is not None:
            pulumi.set(__self__, "iap", iap)
        if location_id is not None:
            pulumi.set(__self__, "location_id", location_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if serving_status is not None:
            pulumi.set(__self__, "serving_status", serving_status)
        if url_dispatch_rules is not None:
            pulumi.set(__self__, "url_dispatch_rules", url_dispatch_rules)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the app, usually `{PROJECT_ID}`
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain to authenticate users with when using App Engine's User API.
        """
        return pulumi.get(self, "auth_domain")

    @auth_domain.setter
    def auth_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_domain", value)

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The GCS bucket code is being stored in for this app.
        """
        return pulumi.get(self, "code_bucket")

    @code_bucket.setter
    def code_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code_bucket", value)

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database_type")

    @database_type.setter
    def database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_type", value)

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The GCS bucket content is being stored in for this app.
        """
        return pulumi.get(self, "default_bucket")

    @default_bucket.setter
    def default_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_bucket", value)

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The default hostname for this app.
        """
        return pulumi.get(self, "default_hostname")

    @default_hostname.setter
    def default_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_hostname", value)

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> Optional[pulumi.Input['ApplicationFeatureSettingsArgs']]:
        """
        A block of optional settings to configure specific App Engine features:
        """
        return pulumi.get(self, "feature_settings")

    @feature_settings.setter
    def feature_settings(self, value: Optional[pulumi.Input['ApplicationFeatureSettingsArgs']]):
        pulumi.set(self, "feature_settings", value)

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The GCR domain used for storing managed Docker images for this app.
        """
        return pulumi.get(self, "gcr_domain")

    @gcr_domain.setter
    def gcr_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcr_domain", value)

    @property
    @pulumi.getter
    def iap(self) -> Optional[pulumi.Input['ApplicationIapArgs']]:
        """
        Settings for enabling Cloud Identity Aware Proxy
        """
        return pulumi.get(self, "iap")

    @iap.setter
    def iap(self, value: Optional[pulumi.Input['ApplicationIapArgs']]):
        pulumi.set(self, "iap", value)

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> Optional[pulumi.Input[str]]:
        """
        The [location](https://cloud.google.com/appengine/docs/locations)
        to serve the app from.
        """
        return pulumi.get(self, "location_id")

    @location_id.setter
    def location_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the app, usually `apps/{PROJECT_ID}`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID to create the application under.
        ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
        you may get a "Permission denied" error.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> Optional[pulumi.Input[str]]:
        """
        The serving status of the app.
        """
        return pulumi.get(self, "serving_status")

    @serving_status.setter
    def serving_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serving_status", value)

    @property
    @pulumi.getter(name="urlDispatchRules")
    def url_dispatch_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationUrlDispatchRuleArgs']]]]:
        """
        A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        return pulumi.get(self, "url_dispatch_rules")

    @url_dispatch_rules.setter
    def url_dispatch_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationUrlDispatchRuleArgs']]]]):
        pulumi.set(self, "url_dispatch_rules", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 feature_settings: Optional[pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']]] = None,
                 iap: Optional[pulumi.Input[pulumi.InputType['ApplicationIapArgs']]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="your-project-id",
            org_id="1234567")
        app = gcp.appengine.Application("app",
            project=my_project.project_id,
            location_id="us-central")
        ```

        ## Import

        Applications can be imported using the ID of the project the application belongs to, e.g.

        ```sh
         $ pulumi import gcp:appengine/application:Application app your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[pulumi.InputType['ApplicationIapArgs']] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="your-project-id",
            org_id="1234567")
        app = gcp.appengine.Application("app",
            project=my_project.project_id,
            location_id="us-central")
        ```

        ## Import

        Applications can be imported using the ID of the project the application belongs to, e.g.

        ```sh
         $ pulumi import gcp:appengine/application:Application app your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 feature_settings: Optional[pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']]] = None,
                 iap: Optional[pulumi.Input[pulumi.InputType['ApplicationIapArgs']]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["auth_domain"] = auth_domain
            __props__.__dict__["database_type"] = database_type
            __props__.__dict__["feature_settings"] = feature_settings
            __props__.__dict__["iap"] = iap
            if location_id is None and not opts.urn:
                raise TypeError("Missing required property 'location_id'")
            __props__.__dict__["location_id"] = location_id
            __props__.__dict__["project"] = project
            __props__.__dict__["serving_status"] = serving_status
            __props__.__dict__["app_id"] = None
            __props__.__dict__["code_bucket"] = None
            __props__.__dict__["default_bucket"] = None
            __props__.__dict__["default_hostname"] = None
            __props__.__dict__["gcr_domain"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["url_dispatch_rules"] = None
        super(Application, __self__).__init__(
            'gcp:appengine/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            auth_domain: Optional[pulumi.Input[str]] = None,
            code_bucket: Optional[pulumi.Input[str]] = None,
            database_type: Optional[pulumi.Input[str]] = None,
            default_bucket: Optional[pulumi.Input[str]] = None,
            default_hostname: Optional[pulumi.Input[str]] = None,
            feature_settings: Optional[pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']]] = None,
            gcr_domain: Optional[pulumi.Input[str]] = None,
            iap: Optional[pulumi.Input[pulumi.InputType['ApplicationIapArgs']]] = None,
            location_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            serving_status: Optional[pulumi.Input[str]] = None,
            url_dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationUrlDispatchRuleArgs']]]]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Identifier of the app, usually `{PROJECT_ID}`
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[str] code_bucket: The GCS bucket code is being stored in for this app.
        :param pulumi.Input[str] default_bucket: The GCS bucket content is being stored in for this app.
        :param pulumi.Input[str] default_hostname: The default hostname for this app.
        :param pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[str] gcr_domain: The GCR domain used for storing managed Docker images for this app.
        :param pulumi.Input[pulumi.InputType['ApplicationIapArgs']] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] name: Unique name of the app, usually `apps/{PROJECT_ID}`
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationUrlDispatchRuleArgs']]]] url_dispatch_rules: A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["auth_domain"] = auth_domain
        __props__.__dict__["code_bucket"] = code_bucket
        __props__.__dict__["database_type"] = database_type
        __props__.__dict__["default_bucket"] = default_bucket
        __props__.__dict__["default_hostname"] = default_hostname
        __props__.__dict__["feature_settings"] = feature_settings
        __props__.__dict__["gcr_domain"] = gcr_domain
        __props__.__dict__["iap"] = iap
        __props__.__dict__["location_id"] = location_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["serving_status"] = serving_status
        __props__.__dict__["url_dispatch_rules"] = url_dispatch_rules
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Identifier of the app, usually `{PROJECT_ID}`
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> pulumi.Output[str]:
        """
        The domain to authenticate users with when using App Engine's User API.
        """
        return pulumi.get(self, "auth_domain")

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> pulumi.Output[str]:
        """
        The GCS bucket code is being stored in for this app.
        """
        return pulumi.get(self, "code_bucket")

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_type")

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> pulumi.Output[str]:
        """
        The GCS bucket content is being stored in for this app.
        """
        return pulumi.get(self, "default_bucket")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> pulumi.Output[str]:
        """
        The default hostname for this app.
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> pulumi.Output['outputs.ApplicationFeatureSettings']:
        """
        A block of optional settings to configure specific App Engine features:
        """
        return pulumi.get(self, "feature_settings")

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> pulumi.Output[str]:
        """
        The GCR domain used for storing managed Docker images for this app.
        """
        return pulumi.get(self, "gcr_domain")

    @property
    @pulumi.getter
    def iap(self) -> pulumi.Output['outputs.ApplicationIap']:
        """
        Settings for enabling Cloud Identity Aware Proxy
        """
        return pulumi.get(self, "iap")

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> pulumi.Output[str]:
        """
        The [location](https://cloud.google.com/appengine/docs/locations)
        to serve the app from.
        """
        return pulumi.get(self, "location_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the app, usually `apps/{PROJECT_ID}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID to create the application under.
        ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
        you may get a "Permission denied" error.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> pulumi.Output[str]:
        """
        The serving status of the app.
        """
        return pulumi.get(self, "serving_status")

    @property
    @pulumi.getter(name="urlDispatchRules")
    def url_dispatch_rules(self) -> pulumi.Output[Sequence['outputs.ApplicationUrlDispatchRule']]:
        """
        A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        return pulumi.get(self, "url_dispatch_rules")

