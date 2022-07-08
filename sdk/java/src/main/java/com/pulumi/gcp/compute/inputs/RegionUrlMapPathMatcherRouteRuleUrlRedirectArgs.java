// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs Empty = new RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs();

    /**
     * The host that will be used in the redirect response instead of the one that was
     * supplied in the request. The value must be between 1 and 255 characters.
     * 
     */
    @Import(name="hostRedirect")
    private @Nullable Output<String> hostRedirect;

    /**
     * @return The host that will be used in the redirect response instead of the one that was
     * supplied in the request. The value must be between 1 and 255 characters.
     * 
     */
    public Optional<Output<String>> hostRedirect() {
        return Optional.ofNullable(this.hostRedirect);
    }

    /**
     * If set to true, the URL scheme in the redirected request is set to https. If set to
     * false, the URL scheme of the redirected request will remain the same as that of the
     * request. This must only be set for UrlMaps used in TargetHttpProxys. Setting this
     * true for TargetHttpsProxy is not permitted. The default is set to false.
     * 
     */
    @Import(name="httpsRedirect")
    private @Nullable Output<Boolean> httpsRedirect;

    /**
     * @return If set to true, the URL scheme in the redirected request is set to https. If set to
     * false, the URL scheme of the redirected request will remain the same as that of the
     * request. This must only be set for UrlMaps used in TargetHttpProxys. Setting this
     * true for TargetHttpsProxy is not permitted. The default is set to false.
     * 
     */
    public Optional<Output<Boolean>> httpsRedirect() {
        return Optional.ofNullable(this.httpsRedirect);
    }

    /**
     * The path that will be used in the redirect response instead of the one that was
     * supplied in the request. pathRedirect cannot be supplied together with
     * prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the
     * original request will be used for the redirect. The value must be between 1 and 1024
     * characters.
     * 
     */
    @Import(name="pathRedirect")
    private @Nullable Output<String> pathRedirect;

    /**
     * @return The path that will be used in the redirect response instead of the one that was
     * supplied in the request. pathRedirect cannot be supplied together with
     * prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the
     * original request will be used for the redirect. The value must be between 1 and 1024
     * characters.
     * 
     */
    public Optional<Output<String>> pathRedirect() {
        return Optional.ofNullable(this.pathRedirect);
    }

    /**
     * The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,
     * retaining the remaining portion of the URL before redirecting the request.
     * prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or
     * neither. If neither is supplied, the path of the original request will be used for
     * the redirect. The value must be between 1 and 1024 characters.
     * 
     */
    @Import(name="prefixRedirect")
    private @Nullable Output<String> prefixRedirect;

    /**
     * @return The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,
     * retaining the remaining portion of the URL before redirecting the request.
     * prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or
     * neither. If neither is supplied, the path of the original request will be used for
     * the redirect. The value must be between 1 and 1024 characters.
     * 
     */
    public Optional<Output<String>> prefixRedirect() {
        return Optional.ofNullable(this.prefixRedirect);
    }

    /**
     * The HTTP Status code to use for this RedirectAction. Supported values are:
     * * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.
     * * FOUND, which corresponds to 302.
     * * SEE_OTHER which corresponds to 303.
     * * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method
     *   will be retained.
     * * PERMANENT_REDIRECT, which corresponds to 308. In this case,
     *   the request method will be retained.
     * 
     */
    @Import(name="redirectResponseCode")
    private @Nullable Output<String> redirectResponseCode;

    /**
     * @return The HTTP Status code to use for this RedirectAction. Supported values are:
     * * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.
     * * FOUND, which corresponds to 302.
     * * SEE_OTHER which corresponds to 303.
     * * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method
     *   will be retained.
     * * PERMANENT_REDIRECT, which corresponds to 308. In this case,
     *   the request method will be retained.
     * 
     */
    public Optional<Output<String>> redirectResponseCode() {
        return Optional.ofNullable(this.redirectResponseCode);
    }

    /**
     * If set to true, any accompanying query portion of the original URL is removed prior
     * to redirecting the request. If set to false, the query portion of the original URL is
     * retained.
     * This field is required to ensure an empty block is not set. The normal default value is false.
     * 
     */
    @Import(name="stripQuery")
    private @Nullable Output<Boolean> stripQuery;

    /**
     * @return If set to true, any accompanying query portion of the original URL is removed prior
     * to redirecting the request. If set to false, the query portion of the original URL is
     * retained.
     * This field is required to ensure an empty block is not set. The normal default value is false.
     * 
     */
    public Optional<Output<Boolean>> stripQuery() {
        return Optional.ofNullable(this.stripQuery);
    }

    private RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs() {}

    private RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs(RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs $) {
        this.hostRedirect = $.hostRedirect;
        this.httpsRedirect = $.httpsRedirect;
        this.pathRedirect = $.pathRedirect;
        this.prefixRedirect = $.prefixRedirect;
        this.redirectResponseCode = $.redirectResponseCode;
        this.stripQuery = $.stripQuery;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs $;

        public Builder() {
            $ = new RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs();
        }

        public Builder(RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs defaults) {
            $ = new RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostRedirect The host that will be used in the redirect response instead of the one that was
         * supplied in the request. The value must be between 1 and 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder hostRedirect(@Nullable Output<String> hostRedirect) {
            $.hostRedirect = hostRedirect;
            return this;
        }

        /**
         * @param hostRedirect The host that will be used in the redirect response instead of the one that was
         * supplied in the request. The value must be between 1 and 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder hostRedirect(String hostRedirect) {
            return hostRedirect(Output.of(hostRedirect));
        }

        /**
         * @param httpsRedirect If set to true, the URL scheme in the redirected request is set to https. If set to
         * false, the URL scheme of the redirected request will remain the same as that of the
         * request. This must only be set for UrlMaps used in TargetHttpProxys. Setting this
         * true for TargetHttpsProxy is not permitted. The default is set to false.
         * 
         * @return builder
         * 
         */
        public Builder httpsRedirect(@Nullable Output<Boolean> httpsRedirect) {
            $.httpsRedirect = httpsRedirect;
            return this;
        }

        /**
         * @param httpsRedirect If set to true, the URL scheme in the redirected request is set to https. If set to
         * false, the URL scheme of the redirected request will remain the same as that of the
         * request. This must only be set for UrlMaps used in TargetHttpProxys. Setting this
         * true for TargetHttpsProxy is not permitted. The default is set to false.
         * 
         * @return builder
         * 
         */
        public Builder httpsRedirect(Boolean httpsRedirect) {
            return httpsRedirect(Output.of(httpsRedirect));
        }

        /**
         * @param pathRedirect The path that will be used in the redirect response instead of the one that was
         * supplied in the request. pathRedirect cannot be supplied together with
         * prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the
         * original request will be used for the redirect. The value must be between 1 and 1024
         * characters.
         * 
         * @return builder
         * 
         */
        public Builder pathRedirect(@Nullable Output<String> pathRedirect) {
            $.pathRedirect = pathRedirect;
            return this;
        }

        /**
         * @param pathRedirect The path that will be used in the redirect response instead of the one that was
         * supplied in the request. pathRedirect cannot be supplied together with
         * prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the
         * original request will be used for the redirect. The value must be between 1 and 1024
         * characters.
         * 
         * @return builder
         * 
         */
        public Builder pathRedirect(String pathRedirect) {
            return pathRedirect(Output.of(pathRedirect));
        }

        /**
         * @param prefixRedirect The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,
         * retaining the remaining portion of the URL before redirecting the request.
         * prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or
         * neither. If neither is supplied, the path of the original request will be used for
         * the redirect. The value must be between 1 and 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder prefixRedirect(@Nullable Output<String> prefixRedirect) {
            $.prefixRedirect = prefixRedirect;
            return this;
        }

        /**
         * @param prefixRedirect The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,
         * retaining the remaining portion of the URL before redirecting the request.
         * prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or
         * neither. If neither is supplied, the path of the original request will be used for
         * the redirect. The value must be between 1 and 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder prefixRedirect(String prefixRedirect) {
            return prefixRedirect(Output.of(prefixRedirect));
        }

        /**
         * @param redirectResponseCode The HTTP Status code to use for this RedirectAction. Supported values are:
         * * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.
         * * FOUND, which corresponds to 302.
         * * SEE_OTHER which corresponds to 303.
         * * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method
         *   will be retained.
         * * PERMANENT_REDIRECT, which corresponds to 308. In this case,
         *   the request method will be retained.
         * 
         * @return builder
         * 
         */
        public Builder redirectResponseCode(@Nullable Output<String> redirectResponseCode) {
            $.redirectResponseCode = redirectResponseCode;
            return this;
        }

        /**
         * @param redirectResponseCode The HTTP Status code to use for this RedirectAction. Supported values are:
         * * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.
         * * FOUND, which corresponds to 302.
         * * SEE_OTHER which corresponds to 303.
         * * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method
         *   will be retained.
         * * PERMANENT_REDIRECT, which corresponds to 308. In this case,
         *   the request method will be retained.
         * 
         * @return builder
         * 
         */
        public Builder redirectResponseCode(String redirectResponseCode) {
            return redirectResponseCode(Output.of(redirectResponseCode));
        }

        /**
         * @param stripQuery If set to true, any accompanying query portion of the original URL is removed prior
         * to redirecting the request. If set to false, the query portion of the original URL is
         * retained.
         * This field is required to ensure an empty block is not set. The normal default value is false.
         * 
         * @return builder
         * 
         */
        public Builder stripQuery(@Nullable Output<Boolean> stripQuery) {
            $.stripQuery = stripQuery;
            return this;
        }

        /**
         * @param stripQuery If set to true, any accompanying query portion of the original URL is removed prior
         * to redirecting the request. If set to false, the query portion of the original URL is
         * retained.
         * This field is required to ensure an empty block is not set. The normal default value is false.
         * 
         * @return builder
         * 
         */
        public Builder stripQuery(Boolean stripQuery) {
            return stripQuery(Output.of(stripQuery));
        }

        public RegionUrlMapPathMatcherRouteRuleUrlRedirectArgs build() {
            return $;
        }
    }

}
