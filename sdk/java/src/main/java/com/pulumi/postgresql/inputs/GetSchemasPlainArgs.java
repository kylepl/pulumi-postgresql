// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSchemasPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSchemasPlainArgs Empty = new GetSchemasPlainArgs();

    /**
     * The PostgreSQL database which will be queried for schema names.
     * 
     */
    @Import(name="database", required=true)
    private String database;

    /**
     * @return The PostgreSQL database which will be queried for schema names.
     * 
     */
    public String database() {
        return this.database;
    }

    /**
     * Determines whether to include system schemas (pg_ prefix and information_schema). &#39;public&#39; will always be included. Defaults to ``false``.
     * 
     */
    @Import(name="includeSystemSchemas")
    private @Nullable Boolean includeSystemSchemas;

    /**
     * @return Determines whether to include system schemas (pg_ prefix and information_schema). &#39;public&#39; will always be included. Defaults to ``false``.
     * 
     */
    public Optional<Boolean> includeSystemSchemas() {
        return Optional.ofNullable(this.includeSystemSchemas);
    }

    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
     * 
     */
    @Import(name="likeAllPatterns")
    private @Nullable List<String> likeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
     * 
     */
    public Optional<List<String>> likeAllPatterns() {
        return Optional.ofNullable(this.likeAllPatterns);
    }

    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
     * 
     */
    @Import(name="likeAnyPatterns")
    private @Nullable List<String> likeAnyPatterns;

    /**
     * @return List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
     * 
     */
    public Optional<List<String>> likeAnyPatterns() {
        return Optional.ofNullable(this.likeAnyPatterns);
    }

    /**
     * List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     * 
     */
    @Import(name="notLikeAllPatterns")
    private @Nullable List<String> notLikeAllPatterns;

    /**
     * @return List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
     * 
     */
    public Optional<List<String>> notLikeAllPatterns() {
        return Optional.ofNullable(this.notLikeAllPatterns);
    }

    /**
     * Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
     * 
     */
    @Import(name="regexPattern")
    private @Nullable String regexPattern;

    /**
     * @return Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
     * 
     */
    public Optional<String> regexPattern() {
        return Optional.ofNullable(this.regexPattern);
    }

    private GetSchemasPlainArgs() {}

    private GetSchemasPlainArgs(GetSchemasPlainArgs $) {
        this.database = $.database;
        this.includeSystemSchemas = $.includeSystemSchemas;
        this.likeAllPatterns = $.likeAllPatterns;
        this.likeAnyPatterns = $.likeAnyPatterns;
        this.notLikeAllPatterns = $.notLikeAllPatterns;
        this.regexPattern = $.regexPattern;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSchemasPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSchemasPlainArgs $;

        public Builder() {
            $ = new GetSchemasPlainArgs();
        }

        public Builder(GetSchemasPlainArgs defaults) {
            $ = new GetSchemasPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param database The PostgreSQL database which will be queried for schema names.
         * 
         * @return builder
         * 
         */
        public Builder database(String database) {
            $.database = database;
            return this;
        }

        /**
         * @param includeSystemSchemas Determines whether to include system schemas (pg_ prefix and information_schema). &#39;public&#39; will always be included. Defaults to ``false``.
         * 
         * @return builder
         * 
         */
        public Builder includeSystemSchemas(@Nullable Boolean includeSystemSchemas) {
            $.includeSystemSchemas = includeSystemSchemas;
            return this;
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(@Nullable List<String> likeAllPatterns) {
            $.likeAllPatterns = likeAllPatterns;
            return this;
        }

        /**
         * @param likeAllPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAllPatterns(String... likeAllPatterns) {
            return likeAllPatterns(List.of(likeAllPatterns));
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(@Nullable List<String> likeAnyPatterns) {
            $.likeAnyPatterns = likeAnyPatterns;
            return this;
        }

        /**
         * @param likeAnyPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
         * 
         * @return builder
         * 
         */
        public Builder likeAnyPatterns(String... likeAnyPatterns) {
            return likeAnyPatterns(List.of(likeAnyPatterns));
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(@Nullable List<String> notLikeAllPatterns) {
            $.notLikeAllPatterns = notLikeAllPatterns;
            return this;
        }

        /**
         * @param notLikeAllPatterns List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
         * 
         * @return builder
         * 
         */
        public Builder notLikeAllPatterns(String... notLikeAllPatterns) {
            return notLikeAllPatterns(List.of(notLikeAllPatterns));
        }

        /**
         * @param regexPattern Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
         * 
         * @return builder
         * 
         */
        public Builder regexPattern(@Nullable String regexPattern) {
            $.regexPattern = regexPattern;
            return this;
        }

        public GetSchemasPlainArgs build() {
            $.database = Objects.requireNonNull($.database, "expected parameter 'database' to be non-null");
            return $;
        }
    }

}