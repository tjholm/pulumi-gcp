// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecurityPolicyAdvancedOptionsConfig {
    /**
     * @return Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
     * * DISABLED - Don&#39;t parse JSON payloads in POST bodies.
     * * STANDARD - Parse JSON payloads in POST bodies.
     * 
     */
    private final @Nullable String jsonParsing;
    /**
     * @return Log level to use. Defaults to `NORMAL`.
     * * NORMAL - Normal log level.
     * * VERBOSE - Verbose log level.
     * 
     */
    private final @Nullable String logLevel;

    @CustomType.Constructor
    private SecurityPolicyAdvancedOptionsConfig(
        @CustomType.Parameter("jsonParsing") @Nullable String jsonParsing,
        @CustomType.Parameter("logLevel") @Nullable String logLevel) {
        this.jsonParsing = jsonParsing;
        this.logLevel = logLevel;
    }

    /**
     * @return Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
     * * DISABLED - Don&#39;t parse JSON payloads in POST bodies.
     * * STANDARD - Parse JSON payloads in POST bodies.
     * 
     */
    public Optional<String> jsonParsing() {
        return Optional.ofNullable(this.jsonParsing);
    }
    /**
     * @return Log level to use. Defaults to `NORMAL`.
     * * NORMAL - Normal log level.
     * * VERBOSE - Verbose log level.
     * 
     */
    public Optional<String> logLevel() {
        return Optional.ofNullable(this.logLevel);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecurityPolicyAdvancedOptionsConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String jsonParsing;
        private @Nullable String logLevel;

        public Builder() {
    	      // Empty
        }

        public Builder(SecurityPolicyAdvancedOptionsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.jsonParsing = defaults.jsonParsing;
    	      this.logLevel = defaults.logLevel;
        }

        public Builder jsonParsing(@Nullable String jsonParsing) {
            this.jsonParsing = jsonParsing;
            return this;
        }
        public Builder logLevel(@Nullable String logLevel) {
            this.logLevel = logLevel;
            return this;
        }        public SecurityPolicyAdvancedOptionsConfig build() {
            return new SecurityPolicyAdvancedOptionsConfig(jsonParsing, logLevel);
        }
    }
}
