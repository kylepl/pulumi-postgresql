// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.ProviderArgs;
import com.pulumi.postgresql.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the postgresql package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:postgresql")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * AWS profile to use for IAM auth
     * 
     */
    @Export(name="awsRdsIamProfile", type=String.class, parameters={})
    private Output</* @Nullable */ String> awsRdsIamProfile;

    /**
     * @return AWS profile to use for IAM auth
     * 
     */
    public Output<Optional<String>> awsRdsIamProfile() {
        return Codegen.optional(this.awsRdsIamProfile);
    }
    /**
     * The name of the database to connect to in order to conenct to (defaults to `postgres`).
     * 
     */
    @Export(name="database", type=String.class, parameters={})
    private Output</* @Nullable */ String> database;

    /**
     * @return The name of the database to connect to in order to conenct to (defaults to `postgres`).
     * 
     */
    public Output<Optional<String>> database() {
        return Codegen.optional(this.database);
    }
    /**
     * Database username associated to the connected user (for user name maps)
     * 
     */
    @Export(name="databaseUsername", type=String.class, parameters={})
    private Output</* @Nullable */ String> databaseUsername;

    /**
     * @return Database username associated to the connected user (for user name maps)
     * 
     */
    public Output<Optional<String>> databaseUsername() {
        return Codegen.optional(this.databaseUsername);
    }
    /**
     * Specify the expected version of PostgreSQL.
     * 
     */
    @Export(name="expectedVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> expectedVersion;

    /**
     * @return Specify the expected version of PostgreSQL.
     * 
     */
    public Output<Optional<String>> expectedVersion() {
        return Codegen.optional(this.expectedVersion);
    }
    /**
     * Name of PostgreSQL server address to connect to
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output</* @Nullable */ String> host;

    /**
     * @return Name of PostgreSQL server address to connect to
     * 
     */
    public Output<Optional<String>> host() {
        return Codegen.optional(this.host);
    }
    /**
     * Password to be used if the PostgreSQL server demands password authentication
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Password to be used if the PostgreSQL server demands password authentication
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    @Export(name="scheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> scheme;

    public Output<Optional<String>> scheme() {
        return Codegen.optional(this.scheme);
    }
    /**
     * @deprecated
     * Rename PostgreSQL provider `ssl_mode` attribute to `sslmode`
     * 
     */
    @Deprecated /* Rename PostgreSQL provider `ssl_mode` attribute to `sslmode` */
    @Export(name="sslMode", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslMode;

    public Output<Optional<String>> sslMode() {
        return Codegen.optional(this.sslMode);
    }
    /**
     * This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     * 
     */
    @Export(name="sslmode", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslmode;

    /**
     * @return This option determines whether or with what priority a secure SSL TCP/IP connection will be negotiated with the
     * PostgreSQL server
     * 
     */
    public Output<Optional<String>> sslmode() {
        return Codegen.optional(this.sslmode);
    }
    /**
     * The SSL server root certificate file path. The file must contain PEM encoded data.
     * 
     */
    @Export(name="sslrootcert", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslrootcert;

    /**
     * @return The SSL server root certificate file path. The file must contain PEM encoded data.
     * 
     */
    public Output<Optional<String>> sslrootcert() {
        return Codegen.optional(this.sslrootcert);
    }
    /**
     * PostgreSQL user name to connect as
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output</* @Nullable */ String> username;

    /**
     * @return PostgreSQL user name to connect as
     * 
     */
    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
