// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.diagflow.outputs.CxPageFormParameter;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class CxPageForm {
    /**
     * @return Parameters to collect from the user.
     * Structure is documented below.
     * 
     */
    private final @Nullable List<CxPageFormParameter> parameters;

    @CustomType.Constructor
    private CxPageForm(@CustomType.Parameter("parameters") @Nullable List<CxPageFormParameter> parameters) {
        this.parameters = parameters;
    }

    /**
     * @return Parameters to collect from the user.
     * Structure is documented below.
     * 
     */
    public List<CxPageFormParameter> parameters() {
        return this.parameters == null ? List.of() : this.parameters;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CxPageForm defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<CxPageFormParameter> parameters;

        public Builder() {
    	      // Empty
        }

        public Builder(CxPageForm defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.parameters = defaults.parameters;
        }

        public Builder parameters(@Nullable List<CxPageFormParameter> parameters) {
            this.parameters = parameters;
            return this;
        }
        public Builder parameters(CxPageFormParameter... parameters) {
            return parameters(List.of(parameters));
        }        public CxPageForm build() {
            return new CxPageForm(parameters);
        }
    }
}
