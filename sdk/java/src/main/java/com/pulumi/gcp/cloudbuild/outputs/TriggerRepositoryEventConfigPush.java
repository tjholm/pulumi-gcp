// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuild.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TriggerRepositoryEventConfigPush {
    /**
     * @return Regex of branches to match.  Specify only one of branch or tag.
     * 
     */
    private @Nullable String branch;
    /**
     * @return When true, only trigger a build if the revision regex does NOT match the git_ref regex.
     * 
     */
    private @Nullable Boolean invertRegex;
    /**
     * @return Regex of tags to match.  Specify only one of branch or tag.
     * 
     */
    private @Nullable String tag;

    private TriggerRepositoryEventConfigPush() {}
    /**
     * @return Regex of branches to match.  Specify only one of branch or tag.
     * 
     */
    public Optional<String> branch() {
        return Optional.ofNullable(this.branch);
    }
    /**
     * @return When true, only trigger a build if the revision regex does NOT match the git_ref regex.
     * 
     */
    public Optional<Boolean> invertRegex() {
        return Optional.ofNullable(this.invertRegex);
    }
    /**
     * @return Regex of tags to match.  Specify only one of branch or tag.
     * 
     */
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TriggerRepositoryEventConfigPush defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String branch;
        private @Nullable Boolean invertRegex;
        private @Nullable String tag;
        public Builder() {}
        public Builder(TriggerRepositoryEventConfigPush defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branch = defaults.branch;
    	      this.invertRegex = defaults.invertRegex;
    	      this.tag = defaults.tag;
        }

        @CustomType.Setter
        public Builder branch(@Nullable String branch) {
            this.branch = branch;
            return this;
        }
        @CustomType.Setter
        public Builder invertRegex(@Nullable Boolean invertRegex) {
            this.invertRegex = invertRegex;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {
            this.tag = tag;
            return this;
        }
        public TriggerRepositoryEventConfigPush build() {
            final var o = new TriggerRepositoryEventConfigPush();
            o.branch = branch;
            o.invertRegex = invertRegex;
            o.tag = tag;
            return o;
        }
    }
}
