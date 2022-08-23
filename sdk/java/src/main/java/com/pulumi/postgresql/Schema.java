// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.postgresql.SchemaArgs;
import com.pulumi.postgresql.Utilities;
import com.pulumi.postgresql.inputs.SchemaState;
import com.pulumi.postgresql.outputs.SchemaPolicy;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="postgresql:index/schema:Schema")
public class Schema extends com.pulumi.resources.CustomResource {
    /**
     * The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
     * 
     */
    @Export(name="database", type=String.class, parameters={})
    private Output<String> database;

    /**
     * @return The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
     * 
     */
    public Output<String> database() {
        return this.database;
    }
    /**
     * When true, will also drop all the objects that are contained in the schema. (Default: false)
     * 
     */
    @Export(name="dropCascade", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dropCascade;

    /**
     * @return When true, will also drop all the objects that are contained in the schema. (Default: false)
     * 
     */
    public Output<Optional<Boolean>> dropCascade() {
        return Codegen.optional(this.dropCascade);
    }
    /**
     * When true, use the existing schema if it exists. (Default: true)
     * 
     */
    @Export(name="ifNotExists", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ifNotExists;

    /**
     * @return When true, use the existing schema if it exists. (Default: true)
     * 
     */
    public Output<Optional<Boolean>> ifNotExists() {
        return Codegen.optional(this.ifNotExists);
    }
    /**
     * The name of the schema. Must be unique in the PostgreSQL
     * database instance where it is configured.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the schema. Must be unique in the PostgreSQL
     * database instance where it is configured.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ROLE who owns the schema.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output<String> owner;

    /**
     * @return The ROLE who owns the schema.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * Can be specified multiple times for each policy.  Each
     * policy block supports fields documented below.
     * 
     * @deprecated
     * Use postgresql_grant resource instead (with object_type=&#34;schema&#34;)
     * 
     */
    @Deprecated /* Use postgresql_grant resource instead (with object_type=""schema"") */
    @Export(name="policies", type=List.class, parameters={SchemaPolicy.class})
    private Output<List<SchemaPolicy>> policies;

    /**
     * @return Can be specified multiple times for each policy.  Each
     * policy block supports fields documented below.
     * 
     */
    public Output<List<SchemaPolicy>> policies() {
        return this.policies;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Schema(String name) {
        this(name, SchemaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Schema(String name, @Nullable SchemaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Schema(String name, @Nullable SchemaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/schema:Schema", name, args == null ? SchemaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Schema(String name, Output<String> id, @Nullable SchemaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("postgresql:index/schema:Schema", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Schema get(String name, Output<String> id, @Nullable SchemaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Schema(name, id, state, options);
    }
}
