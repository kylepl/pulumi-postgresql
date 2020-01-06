# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('postgresql')

connect_timeout = __config__.get('connectTimeout') or (utilities.get_env_int('PGCONNECT_TIMEOUT') or 180)
"""
Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
"""

database = __config__.get('database') or (utilities.get_env('PGDATABASE') or 'postgres')
"""
The name of the database to connect to in order to conenct to (defaults to `postgres`).
"""

database_username = __config__.get('databaseUsername')
"""
Database username associated to the connected user (for user name maps)
"""

expected_version = __config__.get('expectedVersion')
"""
Specify the expected version of PostgreSQL.
"""

host = __config__.get('host') or utilities.get_env('PGHOST')
"""
Name of PostgreSQL server address to connect to
"""

max_connections = __config__.get('maxConnections')
"""
Maximum number of connections to establish to the database. Zero means unlimited.
"""

password = __config__.get('password') or utilities.get_env('PGPASSWORD')
"""
Password to be used if the PostgreSQL server demands password authentication
"""

port = __config__.get('port') or (utilities.get_env_int('PGPORT') or 5432)
"""
The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
"""

ssl_mode = __config__.get('sslMode')

sslmode = __config__.get('sslmode') or utilities.get_env('PGSSLMODE')
"""
This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
PostgreSQL server
"""

superuser = __config__.get('superuser')
"""
Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
Refreshing state password from Postgres)
"""

username = __config__.get('username') or (utilities.get_env('PGUSER') or 'postgres')
"""
PostgreSQL user name to connect as
"""

