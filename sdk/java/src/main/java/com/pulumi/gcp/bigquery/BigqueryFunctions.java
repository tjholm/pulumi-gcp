// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.inputs.GetDefaultServiceAccountArgs;
import com.pulumi.gcp.bigquery.inputs.GetDefaultServiceAccountPlainArgs;
import com.pulumi.gcp.bigquery.outputs.GetDefaultServiceAccountResult;
import java.util.concurrent.CompletableFuture;

public final class BigqueryFunctions {
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDefaultServiceAccountResult> getDefaultServiceAccount() {
        return getDefaultServiceAccount(GetDefaultServiceAccountArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDefaultServiceAccountResult> getDefaultServiceAccountPlain() {
        return getDefaultServiceAccountPlain(GetDefaultServiceAccountPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDefaultServiceAccountResult> getDefaultServiceAccount(GetDefaultServiceAccountArgs args) {
        return getDefaultServiceAccount(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDefaultServiceAccountResult> getDefaultServiceAccountPlain(GetDefaultServiceAccountPlainArgs args) {
        return getDefaultServiceAccountPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDefaultServiceAccountResult> getDefaultServiceAccount(GetDefaultServiceAccountArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", TypeShape.of(GetDefaultServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of a project&#39;s unique BigQuery service account.
     * 
     * Each Google Cloud project has a unique service account used by BigQuery. When using
     * BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
     * this account needs to be granted the
     * `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
     * 
     * For more information see
     * [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.bigquery.BigqueryFunctions;
     * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
     * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var bqSa = BigqueryFunctions.getDefaultServiceAccount();
     * 
     *         var keySaUser = new CryptoKeyIAMMember(&#34;keySaUser&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.key().id())
     *             .role(&#34;roles/cloudkms.cryptoKeyEncrypterDecrypter&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, bqSa.applyValue(getDefaultServiceAccountResult -&gt; getDefaultServiceAccountResult.email())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDefaultServiceAccountResult> getDefaultServiceAccountPlain(GetDefaultServiceAccountPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", TypeShape.of(GetDefaultServiceAccountResult.class), args, Utilities.withVersion(options));
    }
}
