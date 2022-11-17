// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.postgresql.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FunctionArg {
    /**
     * @return An expression to be used as default value if the parameter is not specified.
     * 
     */
    private @Nullable String default_;
    /**
     * @return Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
     * 
     */
    private @Nullable String mode;
    /**
     * @return The name of the argument.
     * 
     */
    private @Nullable String name;
    /**
     * @return The type of the argument.
     * 
     */
    private String type;

    private FunctionArg() {}
    /**
     * @return An expression to be used as default value if the parameter is not specified.
     * 
     */
    public Optional<String> default_() {
        return Optional.ofNullable(this.default_);
    }
    /**
     * @return Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
     * 
     */
    public Optional<String> mode() {
        return Optional.ofNullable(this.mode);
    }
    /**
     * @return The name of the argument.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The type of the argument.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FunctionArg defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String default_;
        private @Nullable String mode;
        private @Nullable String name;
        private String type;
        public Builder() {}
        public Builder(FunctionArg defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.default_ = defaults.default_;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter("default")
        public Builder default_(@Nullable String default_) {
            this.default_ = default_;
            return this;
        }
        @CustomType.Setter
        public Builder mode(@Nullable String mode) {
            this.mode = mode;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public FunctionArg build() {
            final var o = new FunctionArg();
            o.default_ = default_;
            o.mode = mode;
            o.name = name;
            o.type = type;
            return o;
        }
    }
}
