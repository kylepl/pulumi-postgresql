// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.postgresql.inputs.FunctionArgArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FunctionArgs extends com.pulumi.resources.ResourceArgs {

    public static final FunctionArgs Empty = new FunctionArgs();

    /**
     * List of arguments for the function.
     * 
     */
    @Import(name="args")
    private @Nullable Output<List<FunctionArgArgs>> args;

    /**
     * @return List of arguments for the function.
     * 
     */
    public Optional<Output<List<FunctionArgArgs>>> args() {
        return Optional.ofNullable(this.args);
    }

    /**
     * Function body.
     * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
     * 
     */
    @Import(name="body", required=true)
    private Output<String> body;

    /**
     * @return Function body.
     * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
     * 
     */
    public Output<String> body() {
        return this.body;
    }

    /**
     * The database where the function is located.
     * If not specified, the function is created in the current database.
     * 
     */
    @Import(name="database")
    private @Nullable Output<String> database;

    /**
     * @return The database where the function is located.
     * If not specified, the function is created in the current database.
     * 
     */
    public Optional<Output<String>> database() {
        return Optional.ofNullable(this.database);
    }

    /**
     * True to automatically drop objects that depend on the function (such as
     * operators or triggers), and in turn all objects that depend on those objects. Default is false.
     * 
     */
    @Import(name="dropCascade")
    private @Nullable Output<Boolean> dropCascade;

    /**
     * @return True to automatically drop objects that depend on the function (such as
     * operators or triggers), and in turn all objects that depend on those objects. Default is false.
     * 
     */
    public Optional<Output<Boolean>> dropCascade() {
        return Optional.ofNullable(this.dropCascade);
    }

    /**
     * The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
     * 
     */
    @Import(name="language")
    private @Nullable Output<String> language;

    /**
     * @return The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
     * 
     */
    public Optional<Output<String>> language() {
        return Optional.ofNullable(this.language);
    }

    /**
     * The name of the argument.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the argument.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
     * 
     */
    @Import(name="parallel")
    private @Nullable Output<String> parallel;

    /**
     * @return Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
     * 
     */
    public Optional<Output<String>> parallel() {
        return Optional.ofNullable(this.parallel);
    }

    /**
     * Type that the function returns. It can be computed from the OUT arguments. Default is void.
     * 
     */
    @Import(name="returns")
    private @Nullable Output<String> returns;

    /**
     * @return Type that the function returns. It can be computed from the OUT arguments. Default is void.
     * 
     */
    public Optional<Output<String>> returns() {
        return Optional.ofNullable(this.returns);
    }

    /**
     * The schema where the function is located.
     * If not specified, the function is created in the current schema.
     * 
     */
    @Import(name="schema")
    private @Nullable Output<String> schema;

    /**
     * @return The schema where the function is located.
     * If not specified, the function is created in the current schema.
     * 
     */
    public Optional<Output<String>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
     * 
     */
    @Import(name="securityDefiner")
    private @Nullable Output<Boolean> securityDefiner;

    /**
     * @return If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
     * 
     */
    public Optional<Output<Boolean>> securityDefiner() {
        return Optional.ofNullable(this.securityDefiner);
    }

    /**
     * If the function should always return NULL when any of the inputs is NULL. Default is false.
     * 
     */
    @Import(name="strict")
    private @Nullable Output<Boolean> strict;

    /**
     * @return If the function should always return NULL when any of the inputs is NULL. Default is false.
     * 
     */
    public Optional<Output<Boolean>> strict() {
        return Optional.ofNullable(this.strict);
    }

    /**
     * Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
     * 
     */
    @Import(name="volatility")
    private @Nullable Output<String> volatility;

    /**
     * @return Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
     * 
     */
    public Optional<Output<String>> volatility() {
        return Optional.ofNullable(this.volatility);
    }

    private FunctionArgs() {}

    private FunctionArgs(FunctionArgs $) {
        this.args = $.args;
        this.body = $.body;
        this.database = $.database;
        this.dropCascade = $.dropCascade;
        this.language = $.language;
        this.name = $.name;
        this.parallel = $.parallel;
        this.returns = $.returns;
        this.schema = $.schema;
        this.securityDefiner = $.securityDefiner;
        this.strict = $.strict;
        this.volatility = $.volatility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FunctionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FunctionArgs $;

        public Builder() {
            $ = new FunctionArgs();
        }

        public Builder(FunctionArgs defaults) {
            $ = new FunctionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param args List of arguments for the function.
         * 
         * @return builder
         * 
         */
        public Builder args(@Nullable Output<List<FunctionArgArgs>> args) {
            $.args = args;
            return this;
        }

        /**
         * @param args List of arguments for the function.
         * 
         * @return builder
         * 
         */
        public Builder args(List<FunctionArgArgs> args) {
            return args(Output.of(args));
        }

        /**
         * @param args List of arguments for the function.
         * 
         * @return builder
         * 
         */
        public Builder args(FunctionArgArgs... args) {
            return args(List.of(args));
        }

        /**
         * @param body Function body.
         * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
         * 
         * @return builder
         * 
         */
        public Builder body(Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body Function body.
         * This should be the body content within the `AS $$` and the final `$$`. It will also accept the `AS $$` and `$$` if added.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param database The database where the function is located.
         * If not specified, the function is created in the current database.
         * 
         * @return builder
         * 
         */
        public Builder database(@Nullable Output<String> database) {
            $.database = database;
            return this;
        }

        /**
         * @param database The database where the function is located.
         * If not specified, the function is created in the current database.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            return database(Output.of(database));
        }

        /**
         * @param dropCascade True to automatically drop objects that depend on the function (such as
         * operators or triggers), and in turn all objects that depend on those objects. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(@Nullable Output<Boolean> dropCascade) {
            $.dropCascade = dropCascade;
            return this;
        }

        /**
         * @param dropCascade True to automatically drop objects that depend on the function (such as
         * operators or triggers), and in turn all objects that depend on those objects. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder dropCascade(Boolean dropCascade) {
            return dropCascade(Output.of(dropCascade));
        }

        /**
         * @param language The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
         * 
         * @return builder
         * 
         */
        public Builder language(@Nullable Output<String> language) {
            $.language = language;
            return this;
        }

        /**
         * @param language The function programming language. Can be one of internal, sql, c, plpgsql. Default is plpgsql.
         * 
         * @return builder
         * 
         */
        public Builder language(String language) {
            return language(Output.of(language));
        }

        /**
         * @param name The name of the argument.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the argument.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parallel Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
         * 
         * @return builder
         * 
         */
        public Builder parallel(@Nullable Output<String> parallel) {
            $.parallel = parallel;
            return this;
        }

        /**
         * @param parallel Indicates if the function is parallel safe. Can be one of UNSAFE, RESTRICTED, or SAFE. Default is UNSAFE.
         * 
         * @return builder
         * 
         */
        public Builder parallel(String parallel) {
            return parallel(Output.of(parallel));
        }

        /**
         * @param returns Type that the function returns. It can be computed from the OUT arguments. Default is void.
         * 
         * @return builder
         * 
         */
        public Builder returns(@Nullable Output<String> returns) {
            $.returns = returns;
            return this;
        }

        /**
         * @param returns Type that the function returns. It can be computed from the OUT arguments. Default is void.
         * 
         * @return builder
         * 
         */
        public Builder returns(String returns) {
            return returns(Output.of(returns));
        }

        /**
         * @param schema The schema where the function is located.
         * If not specified, the function is created in the current schema.
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<String> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema The schema where the function is located.
         * If not specified, the function is created in the current schema.
         * 
         * @return builder
         * 
         */
        public Builder schema(String schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param securityDefiner If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder securityDefiner(@Nullable Output<Boolean> securityDefiner) {
            $.securityDefiner = securityDefiner;
            return this;
        }

        /**
         * @param securityDefiner If the function should execute with the permissions of the owner, rather than the permissions of the caller. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder securityDefiner(Boolean securityDefiner) {
            return securityDefiner(Output.of(securityDefiner));
        }

        /**
         * @param strict If the function should always return NULL when any of the inputs is NULL. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder strict(@Nullable Output<Boolean> strict) {
            $.strict = strict;
            return this;
        }

        /**
         * @param strict If the function should always return NULL when any of the inputs is NULL. Default is false.
         * 
         * @return builder
         * 
         */
        public Builder strict(Boolean strict) {
            return strict(Output.of(strict));
        }

        /**
         * @param volatility Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
         * 
         * @return builder
         * 
         */
        public Builder volatility(@Nullable Output<String> volatility) {
            $.volatility = volatility;
            return this;
        }

        /**
         * @param volatility Defines the volatility of the function. Can be one of VOLATILE, STABLE, or IMMUTABLE. Default is VOLATILE.
         * 
         * @return builder
         * 
         */
        public Builder volatility(String volatility) {
            return volatility(Output.of(volatility));
        }

        public FunctionArgs build() {
            $.body = Objects.requireNonNull($.body, "expected parameter 'body' to be non-null");
            return $;
        }
    }

}
