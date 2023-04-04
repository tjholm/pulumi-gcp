// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AssetIamMemberCondition {
    private @Nullable String description;
    private String expression;
    private String title;

    private AssetIamMemberCondition() {}
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public String expression() {
        return this.expression;
    }
    public String title() {
        return this.title;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AssetIamMemberCondition defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private String expression;
        private String title;
        public Builder() {}
        public Builder(AssetIamMemberCondition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.expression = defaults.expression;
    	      this.title = defaults.title;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder expression(String expression) {
            this.expression = Objects.requireNonNull(expression);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        public AssetIamMemberCondition build() {
            final var o = new AssetIamMemberCondition();
            o.description = description;
            o.expression = expression;
            o.title = title;
            return o;
        }
    }
}
