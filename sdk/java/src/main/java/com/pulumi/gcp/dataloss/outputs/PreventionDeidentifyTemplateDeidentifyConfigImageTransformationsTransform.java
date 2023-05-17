// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.dataloss.outputs.PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes;
import com.pulumi.gcp.dataloss.outputs.PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText;
import com.pulumi.gcp.dataloss.outputs.PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor;
import com.pulumi.gcp.dataloss.outputs.PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform {
    /**
     * @return Apply transformation to all findings not specified in other ImageTransformation&#39;s selectedInfoTypes.
     * 
     */
    private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes allInfoTypes;
    /**
     * @return Apply transformation to all text that doesn&#39;t match an infoType.
     * 
     */
    private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText allText;
    /**
     * @return The color to use when redacting content from an image. If not specified, the default is black.
     * Structure is documented below.
     * 
     */
    private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor redactionColor;
    /**
     * @return Apply transformation to the selected infoTypes.
     * Structure is documented below.
     * 
     */
    private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes selectedInfoTypes;

    private PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform() {}
    /**
     * @return Apply transformation to all findings not specified in other ImageTransformation&#39;s selectedInfoTypes.
     * 
     */
    public Optional<PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes> allInfoTypes() {
        return Optional.ofNullable(this.allInfoTypes);
    }
    /**
     * @return Apply transformation to all text that doesn&#39;t match an infoType.
     * 
     */
    public Optional<PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText> allText() {
        return Optional.ofNullable(this.allText);
    }
    /**
     * @return The color to use when redacting content from an image. If not specified, the default is black.
     * Structure is documented below.
     * 
     */
    public Optional<PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor> redactionColor() {
        return Optional.ofNullable(this.redactionColor);
    }
    /**
     * @return Apply transformation to the selected infoTypes.
     * Structure is documented below.
     * 
     */
    public Optional<PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes> selectedInfoTypes() {
        return Optional.ofNullable(this.selectedInfoTypes);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes allInfoTypes;
        private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText allText;
        private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor redactionColor;
        private @Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes selectedInfoTypes;
        public Builder() {}
        public Builder(PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allInfoTypes = defaults.allInfoTypes;
    	      this.allText = defaults.allText;
    	      this.redactionColor = defaults.redactionColor;
    	      this.selectedInfoTypes = defaults.selectedInfoTypes;
        }

        @CustomType.Setter
        public Builder allInfoTypes(@Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes allInfoTypes) {
            this.allInfoTypes = allInfoTypes;
            return this;
        }
        @CustomType.Setter
        public Builder allText(@Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText allText) {
            this.allText = allText;
            return this;
        }
        @CustomType.Setter
        public Builder redactionColor(@Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor redactionColor) {
            this.redactionColor = redactionColor;
            return this;
        }
        @CustomType.Setter
        public Builder selectedInfoTypes(@Nullable PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes selectedInfoTypes) {
            this.selectedInfoTypes = selectedInfoTypes;
            return this;
        }
        public PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform build() {
            final var o = new PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform();
            o.allInfoTypes = allInfoTypes;
            o.allText = allText;
            o.redactionColor = redactionColor;
            o.selectedInfoTypes = selectedInfoTypes;
            return o;
        }
    }
}
