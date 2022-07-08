// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.monitoring.outputs.UptimeCheckConfigContentMatcherJsonPathMatcher;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UptimeCheckConfigContentMatcher {
    /**
     * @return String or regex content to match (max 1024 bytes)
     * 
     */
    private final String content;
    /**
     * @return Information needed to perform a JSONPath content match. Used for `ContentMatcherOption::MATCHES_JSON_PATH` and `ContentMatcherOption::NOT_MATCHES_JSON_PATH`.
     * Structure is documented below.
     * 
     */
    private final @Nullable UptimeCheckConfigContentMatcherJsonPathMatcher jsonPathMatcher;
    /**
     * @return The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
     * Default value is `CONTAINS_STRING`.
     * Possible values are `CONTAINS_STRING`, `NOT_CONTAINS_STRING`, `MATCHES_REGEX`, `NOT_MATCHES_REGEX`, `MATCHES_JSON_PATH`, and `NOT_MATCHES_JSON_PATH`.
     * 
     */
    private final @Nullable String matcher;

    @CustomType.Constructor
    private UptimeCheckConfigContentMatcher(
        @CustomType.Parameter("content") String content,
        @CustomType.Parameter("jsonPathMatcher") @Nullable UptimeCheckConfigContentMatcherJsonPathMatcher jsonPathMatcher,
        @CustomType.Parameter("matcher") @Nullable String matcher) {
        this.content = content;
        this.jsonPathMatcher = jsonPathMatcher;
        this.matcher = matcher;
    }

    /**
     * @return String or regex content to match (max 1024 bytes)
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return Information needed to perform a JSONPath content match. Used for `ContentMatcherOption::MATCHES_JSON_PATH` and `ContentMatcherOption::NOT_MATCHES_JSON_PATH`.
     * Structure is documented below.
     * 
     */
    public Optional<UptimeCheckConfigContentMatcherJsonPathMatcher> jsonPathMatcher() {
        return Optional.ofNullable(this.jsonPathMatcher);
    }
    /**
     * @return The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
     * Default value is `CONTAINS_STRING`.
     * Possible values are `CONTAINS_STRING`, `NOT_CONTAINS_STRING`, `MATCHES_REGEX`, `NOT_MATCHES_REGEX`, `MATCHES_JSON_PATH`, and `NOT_MATCHES_JSON_PATH`.
     * 
     */
    public Optional<String> matcher() {
        return Optional.ofNullable(this.matcher);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UptimeCheckConfigContentMatcher defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String content;
        private @Nullable UptimeCheckConfigContentMatcherJsonPathMatcher jsonPathMatcher;
        private @Nullable String matcher;

        public Builder() {
    	      // Empty
        }

        public Builder(UptimeCheckConfigContentMatcher defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.jsonPathMatcher = defaults.jsonPathMatcher;
    	      this.matcher = defaults.matcher;
        }

        public Builder content(String content) {
            this.content = Objects.requireNonNull(content);
            return this;
        }
        public Builder jsonPathMatcher(@Nullable UptimeCheckConfigContentMatcherJsonPathMatcher jsonPathMatcher) {
            this.jsonPathMatcher = jsonPathMatcher;
            return this;
        }
        public Builder matcher(@Nullable String matcher) {
            this.matcher = matcher;
            return this;
        }        public UptimeCheckConfigContentMatcher build() {
            return new UptimeCheckConfigContentMatcher(content, jsonPathMatcher, matcher);
        }
    }
}
