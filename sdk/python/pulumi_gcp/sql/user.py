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

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input['UserPasswordPolicyArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] deletion_policy: The deletion policy for the user.
               Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
               for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field, unless type is set to either CLOUD_IAM_USER
               or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
               and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] type: The user type. It determines the method to authenticate the
               user during login. The default is the database's built-in user type. Flags
               include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        pulumi.set(__self__, "instance", instance)
        if deletion_policy is not None:
            pulumi.set(__self__, "deletion_policy", deletion_policy)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_policy is not None:
            pulumi.set(__self__, "password_policy", password_policy)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The deletion policy for the user.
        Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
        for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        """
        return pulumi.get(self, "deletion_policy")

    @deletion_policy.setter
    def deletion_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deletion_policy", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The host the user can connect from. This is only supported
        for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
        Can be an IP address. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user. Changing this forces a new resource
        to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the user. Can be updated. For Postgres
        instances this is a Required field, unless type is set to either CLOUD_IAM_USER
        or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
        and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> Optional[pulumi.Input['UserPasswordPolicyArgs']]:
        return pulumi.get(self, "password_policy")

    @password_policy.setter
    def password_policy(self, value: Optional[pulumi.Input['UserPasswordPolicyArgs']]):
        pulumi.set(self, "password_policy", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The user type. It determines the method to authenticate the
        user during login. The default is the database's built-in user type. Flags
        include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input['UserPasswordPolicyArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 sql_server_user_details: Optional[pulumi.Input[Sequence[pulumi.Input['UserSqlServerUserDetailArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] deletion_policy: The deletion policy for the user.
               Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
               for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field, unless type is set to either CLOUD_IAM_USER
               or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
               and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] type: The user type. It determines the method to authenticate the
               user during login. The default is the database's built-in user type. Flags
               include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        if deletion_policy is not None:
            pulumi.set(__self__, "deletion_policy", deletion_policy)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_policy is not None:
            pulumi.set(__self__, "password_policy", password_policy)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if sql_server_user_details is not None:
            pulumi.set(__self__, "sql_server_user_details", sql_server_user_details)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The deletion policy for the user.
        Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
        for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        """
        return pulumi.get(self, "deletion_policy")

    @deletion_policy.setter
    def deletion_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deletion_policy", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The host the user can connect from. This is only supported
        for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
        Can be an IP address. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user. Changing this forces a new resource
        to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the user. Can be updated. For Postgres
        instances this is a Required field, unless type is set to either CLOUD_IAM_USER
        or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
        and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> Optional[pulumi.Input['UserPasswordPolicyArgs']]:
        return pulumi.get(self, "password_policy")

    @password_policy.setter
    def password_policy(self, value: Optional[pulumi.Input['UserPasswordPolicyArgs']]):
        pulumi.set(self, "password_policy", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sqlServerUserDetails")
    def sql_server_user_details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserSqlServerUserDetailArgs']]]]:
        return pulumi.get(self, "sql_server_user_details")

    @sql_server_user_details.setter
    def sql_server_user_details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserSqlServerUserDetailArgs']]]]):
        pulumi.set(self, "sql_server_user_details", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The user type. It determines the method to authenticate the
        user during login. The default is the database's built-in user type. Flags
        include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input[pulumi.InputType['UserPasswordPolicyArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        Example creating a SQL User.

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="MYSQL_5_7",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
            ))
        users = gcp.sql.User("users",
            instance=main.name,
            host="me.com",
            password="changeme")
        ```

        Example creating a Cloud IAM User. (For MySQL, specify `cloudsql_iam_authentication`)

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="POSTGRES_9_6",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
                database_flags=[gcp.sql.DatabaseInstanceSettingsDatabaseFlagArgs(
                    name="cloudsql.iam_authentication",
                    value="on",
                )],
            ))
        users = gcp.sql.User("users",
            instance=main.name,
            type="CLOUD_IAM_USER")
        ```

        ## Import

        SQL users for MySQL databases can be imported using the `project`, `instance`, `host` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/main-instance/my-domain.com/me
        ```

         SQL users for PostgreSQL databases can be imported using the `project`, `instance` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/main-instance/me
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deletion_policy: The deletion policy for the user.
               Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
               for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field, unless type is set to either CLOUD_IAM_USER
               or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
               and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] type: The user type. It determines the method to authenticate the
               user during login. The default is the database's built-in user type. Flags
               include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Example creating a SQL User.

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="MYSQL_5_7",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
            ))
        users = gcp.sql.User("users",
            instance=main.name,
            host="me.com",
            password="changeme")
        ```

        Example creating a Cloud IAM User. (For MySQL, specify `cloudsql_iam_authentication`)

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="POSTGRES_9_6",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
                database_flags=[gcp.sql.DatabaseInstanceSettingsDatabaseFlagArgs(
                    name="cloudsql.iam_authentication",
                    value="on",
                )],
            ))
        users = gcp.sql.User("users",
            instance=main.name,
            type="CLOUD_IAM_USER")
        ```

        ## Import

        SQL users for MySQL databases can be imported using the `project`, `instance`, `host` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/main-instance/my-domain.com/me
        ```

         SQL users for PostgreSQL databases can be imported using the `project`, `instance` and `name`, e.g.

        ```sh
         $ pulumi import gcp:sql/user:User users my-project/main-instance/me
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deletion_policy: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input[pulumi.InputType['UserPasswordPolicyArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["deletion_policy"] = deletion_policy
            __props__.__dict__["host"] = host
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__.__dict__["instance"] = instance
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["password_policy"] = password_policy
            __props__.__dict__["project"] = project
            __props__.__dict__["type"] = type
            __props__.__dict__["sql_server_user_details"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'gcp:sql/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deletion_policy: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            password_policy: Optional[pulumi.Input[pulumi.InputType['UserPasswordPolicyArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            sql_server_user_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserSqlServerUserDetailArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deletion_policy: The deletion policy for the user.
               Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
               for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        :param pulumi.Input[str] host: The host the user can connect from. This is only supported
               for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
               Can be an IP address. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the user. Changing this forces a new resource
               to be created.
        :param pulumi.Input[str] password: The password for the user. Can be updated. For Postgres
               instances this is a Required field, unless type is set to either CLOUD_IAM_USER
               or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
               and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] type: The user type. It determines the method to authenticate the
               user during login. The default is the database's built-in user type. Flags
               include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["deletion_policy"] = deletion_policy
        __props__.__dict__["host"] = host
        __props__.__dict__["instance"] = instance
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["password_policy"] = password_policy
        __props__.__dict__["project"] = project
        __props__.__dict__["sql_server_user_details"] = sql_server_user_details
        __props__.__dict__["type"] = type
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deletionPolicy")
    def deletion_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The deletion policy for the user.
        Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
        for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.
        """
        return pulumi.get(self, "deletion_policy")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The host the user can connect from. This is only supported
        for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
        Can be an IP address. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the user. Changing this forces a new resource
        to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for the user. Can be updated. For Postgres
        instances this is a Required field, unless type is set to either CLOUD_IAM_USER
        or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
        and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> pulumi.Output[Optional['outputs.UserPasswordPolicy']]:
        return pulumi.get(self, "password_policy")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="sqlServerUserDetails")
    def sql_server_user_details(self) -> pulumi.Output[Sequence['outputs.UserSqlServerUserDetail']]:
        return pulumi.get(self, "sql_server_user_details")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The user type. It determines the method to authenticate the
        user during login. The default is the database's built-in user type. Flags
        include "BUILT_IN", "CLOUD_IAM_USER", or "CLOUD_IAM_SERVICE_ACCOUNT".
        """
        return pulumi.get(self, "type")

