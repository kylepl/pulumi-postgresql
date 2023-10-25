// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.
 *
 * > **Note:** This resource needs Postgresql version 9 or above.
 */
export class DefaultPrivileges extends pulumi.CustomResource {
    /**
     * Get an existing DefaultPrivileges resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultPrivilegesState, opts?: pulumi.CustomResourceOptions): DefaultPrivileges {
        return new DefaultPrivileges(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/defaultPrivileges:DefaultPrivileges';

    /**
     * Returns true if the given object is an instance of DefaultPrivileges.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultPrivileges {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultPrivileges.__pulumiType;
    }

    /**
     * The database to grant default privileges for this role.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    public readonly objectType!: pulumi.Output<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    public readonly privileges!: pulumi.Output<string[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    public readonly schema!: pulumi.Output<string | undefined>;
    /**
     * Permit the grant recipient to grant it to others
     */
    public readonly withGrantOption!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DefaultPrivileges resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultPrivilegesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultPrivilegesArgs | DefaultPrivilegesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultPrivilegesState | undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["objectType"] = state ? state.objectType : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["privileges"] = state ? state.privileges : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
            resourceInputs["withGrantOption"] = state ? state.withGrantOption : undefined;
        } else {
            const args = argsOrState as DefaultPrivilegesArgs | undefined;
            if ((!args || args.database === undefined) && !opts.urn) {
                throw new Error("Missing required property 'database'");
            }
            if ((!args || args.objectType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'objectType'");
            }
            if ((!args || args.owner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'owner'");
            }
            if ((!args || args.privileges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privileges'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["objectType"] = args ? args.objectType : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["privileges"] = args ? args.privileges : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["withGrantOption"] = args ? args.withGrantOption : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "postgresql:index/defaultPrivileg:DefaultPrivileg" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DefaultPrivileges.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultPrivileges resources.
 */
export interface DefaultPrivilegesState {
    /**
     * The database to grant default privileges for this role.
     */
    database?: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    objectType?: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    owner?: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    privileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    role?: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    schema?: pulumi.Input<string>;
    /**
     * Permit the grant recipient to grant it to others
     */
    withGrantOption?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DefaultPrivileges resource.
 */
export interface DefaultPrivilegesArgs {
    /**
     * The database to grant default privileges for this role.
     */
    database: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
     */
    objectType: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
     */
    owner: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
     */
    privileges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on.
     */
    role: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role.
     */
    schema?: pulumi.Input<string>;
    /**
     * Permit the grant recipient to grant it to others
     */
    withGrantOption?: pulumi.Input<boolean>;
}
