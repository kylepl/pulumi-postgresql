// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GrantArgs extends com.pulumi.resources.ResourceArgs {

    public static final GrantArgs Empty = new GrantArgs();

    /**
     * The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
     * 
     */
    @Import(name="columns")
    private @Nullable Output<List<String>> columns;

    /**
     * @return The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
     * 
     */
    public Optional<Output<List<String>>> columns() {
        return Optional.ofNullable(this.columns);
    }

    /**
     * The database to grant privileges on for this role.
     * 
     */
    @Import(name="database", required=true)
    private Output<String> database;

    /**
     * @return The database to grant privileges on for this role.
     * 
     */
    public Output<String> database() {
        return this.database;
    }

    /**
     * The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
     * 
     */
    @Import(name="objectType", required=true)
    private Output<String> objectType;

    /**
     * @return The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
     * 
     */
    public Output<String> objectType() {
        return this.objectType;
    }

    /**
     * The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
     * 
     */
    @Import(name="objects")
    private @Nullable Output<List<String>> objects;

    /**
     * @return The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
     * 
     */
    public Optional<Output<List<String>>> objects() {
        return Optional.ofNullable(this.objects);
    }

    /**
     * The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     * 
     */
    @Import(name="privileges", required=true)
    private Output<List<String>> privileges;

    /**
     * @return The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
     * 
     */
    public Output<List<String>> privileges() {
        return this.privileges;
    }

    /**
     * The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
     * 
     */
    @Import(name="schema")
    private @Nullable Output<String> schema;

    /**
     * @return The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
     * 
     */
    public Optional<Output<String>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     * 
     */
    @Import(name="withGrantOption")
    private @Nullable Output<Boolean> withGrantOption;

    /**
     * @return Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> withGrantOption() {
        return Optional.ofNullable(this.withGrantOption);
    }

    private GrantArgs() {}

    private GrantArgs(GrantArgs $) {
        this.columns = $.columns;
        this.database = $.database;
        this.objectType = $.objectType;
        this.objects = $.objects;
        this.privileges = $.privileges;
        this.role = $.role;
        this.schema = $.schema;
        this.withGrantOption = $.withGrantOption;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GrantArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GrantArgs $;

        public Builder() {
            $ = new GrantArgs();
        }

        public Builder(GrantArgs defaults) {
            $ = new GrantArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param columns The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
         * 
         * @return builder
         * 
         */
        public Builder columns(@Nullable Output<List<String>> columns) {
            $.columns = columns;
            return this;
        }

        /**
         * @param columns The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
         * 
         * @return builder
         * 
         */
        public Builder columns(List<String> columns) {
            return columns(Output.of(columns));
        }

        /**
         * @param columns The columns upon which to grant the privileges. Required when `object_type` is `column`. You cannot specify this option if the `object_type` is not `column`.
         * 
         * @return builder
         * 
         */
        public Builder columns(String... columns) {
            return columns(List.of(columns));
        }

        /**
         * @param database The database to grant privileges on for this role.
         * 
         * @return builder
         * 
         */
        public Builder database(Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The database to grant privileges on for this role.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param objectType The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
         * 
         * @return builder
         * 
         */
        public Builder objectType(Output<String> objectType) {
            $.objectType = objectType;
            return this;
        }

        /**
         * @param objectType The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server, column).
         * 
         * @return builder
         * 
         */
        public Builder objectType(String objectType) {
            return objectType(Output.of(objectType));
        }

        /**
         * @param objects The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
         * 
         * @return builder
         * 
         */
        public Builder objects(@Nullable Output<List<String>> objects) {
            $.objects = objects;
            return this;
        }

        /**
         * @param objects The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
         * 
         * @return builder
         * 
         */
        public Builder objects(List<String> objects) {
            return objects(Output.of(objects));
        }

        /**
         * @param objects The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `object_type` is `database` or `schema`. When `object_type` is `column`, only one value is allowed.
         * 
         * @return builder
         * 
         */
        public Builder objects(String... objects) {
            return objects(List.of(objects));
        }

        /**
         * @param privileges The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(Output<List<String>> privileges) {
            $.privileges = privileges;
            return this;
        }

        /**
         * @param privileges The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(List<String> privileges) {
            return privileges(Output.of(privileges));
        }

        /**
         * @param privileges The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
         * 
         * @return builder
         * 
         */
        public Builder privileges(String... privileges) {
            return privileges(List.of(privileges));
        }

        /**
         * @param role The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role to grant privileges on, Set it to &#34;public&#34; for all roles.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param schema The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<String> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema The database schema to grant privileges on for this role (Required except if object_type is &#34;database&#34;)
         * 
         * @return builder
         * 
         */
        public Builder schema(String schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param withGrantOption Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder withGrantOption(@Nullable Output<Boolean> withGrantOption) {
            $.withGrantOption = withGrantOption;
            return this;
        }

        /**
         * @param withGrantOption Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder withGrantOption(Boolean withGrantOption) {
            return withGrantOption(Output.of(withGrantOption));
        }

        public GrantArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            $.objectType = Objects.requireNonNull($.objectType, "expected parameter 'objectType' to be non-null");
            $.privileges = Objects.requireNonNull($.privileges, "expected parameter 'privileges' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
