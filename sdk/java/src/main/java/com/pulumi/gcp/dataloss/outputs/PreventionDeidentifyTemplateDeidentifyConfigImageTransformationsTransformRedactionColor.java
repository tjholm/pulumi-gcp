// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor {
    /**
     * @return The amount of blue in the color as a value in the interval [0, 1].
     * 
     */
    private @Nullable Double blue;
    /**
     * @return The amount of green in the color as a value in the interval [0, 1].
     * 
     */
    private @Nullable Double green;
    /**
     * @return The amount of red in the color as a value in the interval [0, 1].
     * 
     */
    private @Nullable Double red;

    private PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor() {}
    /**
     * @return The amount of blue in the color as a value in the interval [0, 1].
     * 
     */
    public Optional<Double> blue() {
        return Optional.ofNullable(this.blue);
    }
    /**
     * @return The amount of green in the color as a value in the interval [0, 1].
     * 
     */
    public Optional<Double> green() {
        return Optional.ofNullable(this.green);
    }
    /**
     * @return The amount of red in the color as a value in the interval [0, 1].
     * 
     */
    public Optional<Double> red() {
        return Optional.ofNullable(this.red);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Double blue;
        private @Nullable Double green;
        private @Nullable Double red;
        public Builder() {}
        public Builder(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.blue = defaults.blue;
    	      this.green = defaults.green;
    	      this.red = defaults.red;
        }

        @CustomType.Setter
        public Builder blue(@Nullable Double blue) {
            this.blue = blue;
            return this;
        }
        @CustomType.Setter
        public Builder green(@Nullable Double green) {
            this.green = green;
            return this;
        }
        @CustomType.Setter
        public Builder red(@Nullable Double red) {
            this.red = red;
            return this;
        }
        public PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor build() {
            final var o = new PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor();
            o.blue = blue;
            o.green = green;
            o.red = red;
            return o;
        }
    }
}
