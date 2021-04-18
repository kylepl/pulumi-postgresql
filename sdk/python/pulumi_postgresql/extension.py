# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ExtensionArgs', 'Extension']

@pulumi.input_type
class ExtensionArgs:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Extension resource.
        :param pulumi.Input[str] database: Which database to create the extension on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
               objects
        :param pulumi.Input[str] name: The name of the extension.
        :param pulumi.Input[str] schema: Sets the schema of an extension.
        :param pulumi.Input[str] version: Sets the version number of the extension.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Which database to create the extension on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
        objects
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the extension.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the schema of an extension.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the version number of the extension.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _ExtensionState:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Extension resources.
        :param pulumi.Input[str] database: Which database to create the extension on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
               objects
        :param pulumi.Input[str] name: The name of the extension.
        :param pulumi.Input[str] schema: Sets the schema of an extension.
        :param pulumi.Input[str] version: Sets the version number of the extension.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if drop_cascade is not None:
            pulumi.set(__self__, "drop_cascade", drop_cascade)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Which database to create the extension on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
        objects
        """
        return pulumi.get(self, "drop_cascade")

    @drop_cascade.setter
    def drop_cascade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drop_cascade", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the extension.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the schema of an extension.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the version number of the extension.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Extension(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``Extension`` resource creates and manages an extension on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        my_extension = postgresql.Extension("myExtension")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Which database to create the extension on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
               objects
        :param pulumi.Input[str] name: The name of the extension.
        :param pulumi.Input[str] schema: Sets the schema of an extension.
        :param pulumi.Input[str] version: Sets the version number of the extension.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExtensionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``Extension`` resource creates and manages an extension on a PostgreSQL
        server.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        my_extension = postgresql.Extension("myExtension")
        ```

        :param str resource_name: The name of the resource.
        :param ExtensionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtensionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExtensionArgs.__new__(ExtensionArgs)

            __props__.__dict__["database"] = database
            __props__.__dict__["drop_cascade"] = drop_cascade
            __props__.__dict__["name"] = name
            __props__.__dict__["schema"] = schema
            __props__.__dict__["version"] = version
        super(Extension, __self__).__init__(
            'postgresql:index/extension:Extension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            drop_cascade: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Extension':
        """
        Get an existing Extension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Which database to create the extension on. Defaults to provider database.
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
               objects
        :param pulumi.Input[str] name: The name of the extension.
        :param pulumi.Input[str] schema: Sets the schema of an extension.
        :param pulumi.Input[str] version: Sets the version number of the extension.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExtensionState.__new__(_ExtensionState)

        __props__.__dict__["database"] = database
        __props__.__dict__["drop_cascade"] = drop_cascade
        __props__.__dict__["name"] = name
        __props__.__dict__["schema"] = schema
        __props__.__dict__["version"] = version
        return Extension(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Which database to create the extension on. Defaults to provider database.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> pulumi.Output[Optional[bool]]:
        """
        When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those
        objects
        """
        return pulumi.get(self, "drop_cascade")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the extension.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        """
        Sets the schema of an extension.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Sets the version number of the extension.
        """
        return pulumi.get(self, "version")

