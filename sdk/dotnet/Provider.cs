// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    /// <summary>
    /// The provider type for the postgresql package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [PostgreSqlResourceType("pulumi:providers:postgresql")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// AWS profile to use for IAM auth
        /// </summary>
        [Output("awsRdsIamProfile")]
        public Output<string?> AwsRdsIamProfile { get; private set; } = null!;

        /// <summary>
        /// AWS region to use for IAM auth
        /// </summary>
        [Output("awsRdsIamRegion")]
        public Output<string?> AwsRdsIamRegion { get; private set; } = null!;

        /// <summary>
        /// The name of the database to connect to in order to conenct to (defaults to `postgres`).
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// Database username associated to the connected user (for user name maps)
        /// </summary>
        [Output("databaseUsername")]
        public Output<string?> DatabaseUsername { get; private set; } = null!;

        /// <summary>
        /// Specify the expected version of PostgreSQL.
        /// </summary>
        [Output("expectedVersion")]
        public Output<string?> ExpectedVersion { get; private set; } = null!;

        /// <summary>
        /// Name of PostgreSQL server address to connect to
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// Password to be used if the PostgreSQL server demands password authentication
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        [Output("sslMode")]
        public Output<string?> SslMode { get; private set; } = null!;

        /// <summary>
        /// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
        /// PostgreSQL server
        /// </summary>
        [Output("sslmode")]
        public Output<string?> Sslmode { get; private set; } = null!;

        /// <summary>
        /// The SSL server root certificate file path. The file must contain PEM encoded data.
        /// </summary>
        [Output("sslrootcert")]
        public Output<string?> Sslrootcert { get; private set; } = null!;

        /// <summary>
        /// PostgreSQL user name to connect as
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use rds_iam instead of password authentication (see:
        /// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
        /// </summary>
        [Input("awsRdsIamAuth", json: true)]
        public Input<bool>? AwsRdsIamAuth { get; set; }

        /// <summary>
        /// AWS profile to use for IAM auth
        /// </summary>
        [Input("awsRdsIamProfile")]
        public Input<string>? AwsRdsIamProfile { get; set; }

        /// <summary>
        /// AWS region to use for IAM auth
        /// </summary>
        [Input("awsRdsIamRegion")]
        public Input<string>? AwsRdsIamRegion { get; set; }

        /// <summary>
        /// SSL client certificate if required by the database.
        /// </summary>
        [Input("clientcert", json: true)]
        public Input<Inputs.ProviderClientcertArgs>? Clientcert { get; set; }

        /// <summary>
        /// Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
        /// </summary>
        [Input("connectTimeout", json: true)]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// The name of the database to connect to in order to conenct to (defaults to `postgres`).
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Database username associated to the connected user (for user name maps)
        /// </summary>
        [Input("databaseUsername")]
        public Input<string>? DatabaseUsername { get; set; }

        /// <summary>
        /// Specify the expected version of PostgreSQL.
        /// </summary>
        [Input("expectedVersion")]
        public Input<string>? ExpectedVersion { get; set; }

        /// <summary>
        /// Name of PostgreSQL server address to connect to
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Maximum number of connections to establish to the database. Zero means unlimited.
        /// </summary>
        [Input("maxConnections", json: true)]
        public Input<int>? MaxConnections { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password to be used if the PostgreSQL server demands password authentication
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The PostgreSQL port number to connect to at the server host, or socket file name extension for Unix-domain connections
        /// </summary>
        [Input("port", json: true)]
        public Input<int>? Port { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        /// <summary>
        /// This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
        /// PostgreSQL server
        /// </summary>
        [Input("sslmode")]
        public Input<string>? Sslmode { get; set; }

        /// <summary>
        /// The SSL server root certificate file path. The file must contain PEM encoded data.
        /// </summary>
        [Input("sslrootcert")]
        public Input<string>? Sslrootcert { get; set; }

        /// <summary>
        /// Specify if the user to connect as is a Postgres superuser or not.If not, some feature might be disabled (e.g.:
        /// Refreshing state password from Postgres)
        /// </summary>
        [Input("superuser", json: true)]
        public Input<bool>? Superuser { get; set; }

        /// <summary>
        /// PostgreSQL user name to connect as
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ProviderArgs()
        {
            ConnectTimeout = Utilities.GetEnvInt32("PGCONNECT_TIMEOUT") ?? 180;
            Sslmode = Utilities.GetEnv("PGSSLMODE");
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
