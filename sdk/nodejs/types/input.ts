// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface FunctionArg {
    /**
     * An expression to be used as default value if the parameter is not specified.
     */
    default?: pulumi.Input<string>;
    /**
     * Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
     */
    mode?: pulumi.Input<string>;
    /**
     * The name of the argument.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the argument.
     */
    type: pulumi.Input<string>;
}

export interface ProviderClientcert {
    cert: pulumi.Input<string>;
    key: pulumi.Input<string>;
}

export interface SchemaPolicy {
    /**
     * Should the specified ROLE have CREATE privileges to the specified SCHEMA.
     */
    create?: pulumi.Input<boolean>;
    /**
     * Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
     */
    createWithGrant?: pulumi.Input<boolean>;
    /**
     * The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
     */
    role?: pulumi.Input<string>;
    /**
     * Should the specified ROLE have USAGE privileges to the specified SCHEMA.
     */
    usage?: pulumi.Input<boolean>;
    /**
     * Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
     */
    usageWithGrant?: pulumi.Input<boolean>;
}
export namespace config {
}
