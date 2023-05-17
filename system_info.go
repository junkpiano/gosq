package gosq

type SystemInfo struct {
	ALMs struct {
		Github_Config string `json:"Github Config"`
		Gitlab_Config string `json:"Gitlab Config"`
	} `json:"ALMs"`
	Bundled struct {
		Config     string `json:"config"`
		Csharp     string `json:"csharp"`
		Flex       string `json:"flex"`
		Go         string `json:"go"`
		Iac        string `json:"iac"`
		Jacoco     string `json:"jacoco"`
		Java       string `json:"java"`
		Javascript string `json:"javascript"`
		Kotlin     string `json:"kotlin"`
		Php        string `json:"php"`
		Python     string `json:"python"`
		Ruby       string `json:"ruby"`
		Sonarscala string `json:"sonarscala"`
		Text       string `json:"text"`
		Vbnet      string `json:"vbnet"`
		Web        string `json:"web"`
		XML        string `json:"xml"`
	} `json:"Bundled"`
	Compute_Engine_Database_Connection struct {
		Pool_Active_Connections   int64 `json:"Pool Active Connections"`
		Pool_Idle_Connections     int64 `json:"Pool Idle Connections"`
		Pool_Max_Connections      int64 `json:"Pool Max Connections"`
		Pool_Max_Lifetime__ms     int64 `json:"Pool Max Lifetime (ms)"`
		Pool_Max_Wait__ms         int64 `json:"Pool Max Wait (ms)"`
		Pool_Min_Idle_Connections int64 `json:"Pool Min Idle Connections"`
		Pool_Total_Connections    int64 `json:"Pool Total Connections"`
	} `json:"Compute Engine Database Connection"`
	Compute_Engine_JVM_Properties struct {
		Awt_toolkit                   string `json:"awt.toolkit"`
		Com_redhat_fips               string `json:"com.redhat.fips"`
		Com_zaxxer_hikari_poolNumber  string `json:"com.zaxxer.hikari.pool_number"`
		File_encoding                 string `json:"file.encoding"`
		File_separator                string `json:"file.separator"`
		Ftp_nonProxyHosts             string `json:"ftp.nonProxyHosts"`
		GopherProxySet                string `json:"gopherProxySet"`
		HTTP_nonProxyHosts            string `json:"http.nonProxyHosts"`
		Java_awt_graphicsenv          string `json:"java.awt.graphicsenv"`
		Java_awt_headless             string `json:"java.awt.headless"`
		Java_awt_printerjob           string `json:"java.awt.printerjob"`
		Java_class_path               string `json:"java.class.path"`
		Java_class_version            string `json:"java.class.version"`
		Java_home                     string `json:"java.home"`
		Java_io_tmpdir                string `json:"java.io.tmpdir"`
		Java_library_path             string `json:"java.library.path"`
		Java_runtime_name             string `json:"java.runtime.name"`
		Java_runtime_version          string `json:"java.runtime.version"`
		Java_specification_name       string `json:"java.specification.name"`
		Java_specification_vendor     string `json:"java.specification.vendor"`
		Java_specification_version    string `json:"java.specification.version"`
		Java_vendor                   string `json:"java.vendor"`
		Java_vendor_url               string `json:"java.vendor.url"`
		Java_vendor_url_bug           string `json:"java.vendor.url.bug"`
		Java_vendor_version           string `json:"java.vendor.version"`
		Java_version                  string `json:"java.version"`
		Java_version_date             string `json:"java.version.date"`
		Java_vm_compressedOopsMode    string `json:"java.vm.compressedOopsMode"`
		Java_vm_info                  string `json:"java.vm.info"`
		Java_vm_name                  string `json:"java.vm.name"`
		Java_vm_specification_name    string `json:"java.vm.specification.name"`
		Java_vm_specification_vendor  string `json:"java.vm.specification.vendor"`
		Java_vm_specification_version string `json:"java.vm.specification.version"`
		Java_vm_vendor                string `json:"java.vm.vendor"`
		Java_vm_version               string `json:"java.vm.version"`
		Jdk_debug                     string `json:"jdk.debug"`
		Jdk_vendor_version            string `json:"jdk.vendor.version"`
		Line_separator                string `json:"line.separator"`
		Os_arch                       string `json:"os.arch"`
		Os_name                       string `json:"os.name"`
		Os_version                    string `json:"os.version"`
		Path_separator                string `json:"path.separator"`
		SocksNonProxyHosts            string `json:"socksNonProxyHosts"`
		Sun_arch_data_model           string `json:"sun.arch.data.model"`
		Sun_boot_library_path         string `json:"sun.boot.library.path"`
		Sun_cpu_endian                string `json:"sun.cpu.endian"`
		Sun_cpu_isalist               string `json:"sun.cpu.isalist"`
		Sun_io_unicode_encoding       string `json:"sun.io.unicode.encoding"`
		Sun_java_command              string `json:"sun.java.command"`
		Sun_java_launcher             string `json:"sun.java.launcher"`
		Sun_jnu_encoding              string `json:"sun.jnu.encoding"`
		Sun_management_compiler       string `json:"sun.management.compiler"`
		Sun_os_patch_level            string `json:"sun.os.patch.level"`
		User_country                  string `json:"user.country"`
		User_dir                      string `json:"user.dir"`
		User_home                     string `json:"user.home"`
		User_language                 string `json:"user.language"`
		User_name                     string `json:"user.name"`
		User_timezone                 string `json:"user.timezone"`
	} `json:"Compute Engine JVM Properties"`
	Compute_Engine_JVM_State struct {
		Free_Memory__MB        int64  `json:"Free Memory (MB)"`
		Heap_Committed__MB     int64  `json:"Heap Committed (MB)"`
		Heap_Init__MB          int64  `json:"Heap Init (MB)"`
		Heap_Max__MB           int64  `json:"Heap Max (MB)"`
		Heap_Used__MB          int64  `json:"Heap Used (MB)"`
		Max_Memory__MB         int64  `json:"Max Memory (MB)"`
		Non_Heap_Committed__MB int64  `json:"Non Heap Committed (MB)"`
		Non_Heap_Init__MB      int64  `json:"Non Heap Init (MB)"`
		Non_Heap_Used__MB      int64  `json:"Non Heap Used (MB)"`
		System_Load_Average    string `json:"System Load Average"`
		Threads                int64  `json:"Threads"`
	} `json:"Compute Engine JVM State"`
	Compute_Engine_Logging struct {
		Logs_Dir   string `json:"Logs Dir"`
		Logs_Level string `json:"Logs Level"`
	} `json:"Compute Engine Logging"`
	Compute_Engine_Tasks struct {
		In_Progress              int64 `json:"In Progress"`
		Longest_Time_Pending__ms int64 `json:"Longest Time Pending (ms)"`
		Max_Worker_Count         int64 `json:"Max Worker Count"`
		Pending                  int64 `json:"Pending"`
		Processed_With_Error     int64 `json:"Processed With Error"`
		Processed_With_Success   int64 `json:"Processed With Success"`
		Processing_Time__ms      int64 `json:"Processing Time (ms)"`
		Worker_Count             int64 `json:"Worker Count"`
		Workers_Paused           bool  `json:"Workers Paused"`
	} `json:"Compute Engine Tasks"`
	Database struct {
		Database                      string `json:"Database"`
		Database_Version              string `json:"Database Version"`
		Default_transaction_isolation string `json:"Default transaction isolation"`
		Driver                        string `json:"Driver"`
		Driver_Version                string `json:"Driver Version"`
		URL                           string `json:"URL"`
		Username                      string `json:"Username"`
	} `json:"Database"`
	Health         string        `json:"Health"`
	Health_Causes  []interface{} `json:"Health Causes"`
	Plugins        struct{}      `json:"Plugins"`
	Search_Indexes struct {
		Index_components___Docs            int64  `json:"Index components - Docs"`
		Index_components___Shards          int64  `json:"Index components - Shards"`
		Index_components___Store_Size      string `json:"Index components - Store Size"`
		Index_issues___Docs                int64  `json:"Index issues - Docs"`
		Index_issues___Shards              int64  `json:"Index issues - Shards"`
		Index_issues___Store_Size          string `json:"Index issues - Store Size"`
		Index_metadatas___Docs             int64  `json:"Index metadatas - Docs"`
		Index_metadatas___Shards           int64  `json:"Index metadatas - Shards"`
		Index_metadatas___Store_Size       string `json:"Index metadatas - Store Size"`
		Index_projectmeasures___Docs       int64  `json:"Index projectmeasures - Docs"`
		Index_projectmeasures___Shards     int64  `json:"Index projectmeasures - Shards"`
		Index_projectmeasures___Store_Size string `json:"Index projectmeasures - Store Size"`
		Index_rules___Docs                 int64  `json:"Index rules - Docs"`
		Index_rules___Shards               int64  `json:"Index rules - Shards"`
		Index_rules___Store_Size           string `json:"Index rules - Store Size"`
		Index_users___Docs                 int64  `json:"Index users - Docs"`
		Index_users___Shards               int64  `json:"Index users - Shards"`
		Index_users___Store_Size           string `json:"Index users - Store Size"`
		Index_views___Docs                 int64  `json:"Index views - Docs"`
		Index_views___Shards               int64  `json:"Index views - Shards"`
		Index_views___Store_Size           string `json:"Index views - Store Size"`
	} `json:"Search Indexes"`
	Search_State struct {
		CPU_Usage                             int64  `json:"CPU Usage (%)"`
		Disk_Available                        string `json:"Disk Available"`
		Field_Data_Circuit_Breaker_Estimation string `json:"Field Data Circuit Breaker Estimation"`
		Field_Data_Circuit_Breaker_Limit      string `json:"Field Data Circuit Breaker Limit"`
		Field_Data_Memory                     string `json:"Field Data Memory"`
		JVM_Heap_Max                          string `json:"JVM Heap Max"`
		JVM_Heap_Usage                        string `json:"JVM Heap Usage"`
		JVM_Heap_Used                         string `json:"JVM Heap Used"`
		JVM_Non_Heap_Used                     string `json:"JVM Non Heap Used"`
		JVM_Threads                           int64  `json:"JVM Threads"`
		Max_File_Descriptors                  int64  `json:"Max File Descriptors"`
		Open_File_Descriptors                 int64  `json:"Open File Descriptors"`
		Query_Cache_Memory                    string `json:"Query Cache Memory"`
		Request_Cache_Memory                  string `json:"Request Cache Memory"`
		Request_Circuit_Breaker_Estimation    string `json:"Request Circuit Breaker Estimation"`
		Request_Circuit_Breaker_Limit         string `json:"Request Circuit Breaker Limit"`
		State                                 string `json:"State"`
		Store_Size                            string `json:"Store Size"`
		Translog_Size                         string `json:"Translog Size"`
	} `json:"Search State"`
	Server_Push_Connections struct {
		SonarLint_Connected_Clients int64 `json:"SonarLint Connected Clients"`
	} `json:"Server Push Connections"`
	Settings struct {
		Default_New_Code_Definition                              string `json:"Default New Code Definition"`
		Devactivity_status                                       string `json:"devactivity.status"`
		HTTP_nonProxyHosts                                       string `json:"http.nonProxyHosts"`
		Process_gracefulStopTimeout                              string `json:"process.gracefulStopTimeout"`
		Process_index                                            string `json:"process.index"`
		Process_key                                              string `json:"process.key"`
		Process_sharedDir                                        string `json:"process.sharedDir"`
		Projects_default_visibility                              string `json:"projects.default.visibility"`
		Qualitygate_default                                      string `json:"qualitygate.default"`
		Sonar_auth_saml_certificate_secured                      string `json:"sonar.auth.saml.certificate.secured"`
		Sonar_auth_saml_enabled                                  string `json:"sonar.auth.saml.enabled"`
		Sonar_auth_saml_loginURL                                 string `json:"sonar.auth.saml.loginUrl"`
		Sonar_auth_saml_providerID                               string `json:"sonar.auth.saml.providerId"`
		Sonar_auth_saml_providerName                             string `json:"sonar.auth.saml.providerName"`
		Sonar_auth_saml_signature_enabled                        string `json:"sonar.auth.saml.signature.enabled"`
		Sonar_auth_saml_sp_certificate_secured                   string `json:"sonar.auth.saml.sp.certificate.secured"`
		Sonar_auth_saml_sp_privateKey_secured                    string `json:"sonar.auth.saml.sp.privateKey.secured"`
		Sonar_auth_saml_user_login                               string `json:"sonar.auth.saml.user.login"`
		Sonar_auth_saml_user_name                                string `json:"sonar.auth.saml.user.name"`
		Sonar_authenticator_ignoreStartupFailure                 string `json:"sonar.authenticator.ignoreStartupFailure"`
		Sonar_autoDatabaseUpgrade                                string `json:"sonar.autoDatabaseUpgrade"`
		Sonar_blueGreenEnabled                                   string `json:"sonar.blueGreenEnabled"`
		Sonar_buildbreaker_skip                                  string `json:"sonar.buildbreaker.skip"`
		Sonar_c_predefinedMacros                                 string `json:"sonar.c.predefinedMacros"`
		Sonar_ce_gracefulStopTimeOutInMs                         string `json:"sonar.ce.gracefulStopTimeOutInMs"`
		Sonar_ce_javaAdditionalOpts                              string `json:"sonar.ce.javaAdditionalOpts"`
		Sonar_ce_javaOpts                                        string `json:"sonar.ce.javaOpts"`
		Sonar_cluster_enabled                                    string `json:"sonar.cluster.enabled"`
		Sonar_cluster_kubernetes                                 string `json:"sonar.cluster.kubernetes"`
		Sonar_cluster_name                                       string `json:"sonar.cluster.name"`
		Sonar_cluster_node_name                                  string `json:"sonar.cluster.node.name"`
		Sonar_cluster_node_port                                  string `json:"sonar.cluster.node.port"`
		Sonar_cluster_web_startupLeader                          string `json:"sonar.cluster.web.startupLeader"`
		Sonar_core_id                                            string `json:"sonar.core.id"`
		Sonar_core_serverBaseURL                                 string `json:"sonar.core.serverBaseURL"`
		Sonar_core_startTime                                     string `json:"sonar.core.startTime"`
		Sonar_core_treemap_colormetric                           string `json:"sonar.core.treemap.colormetric"`
		Sonar_core_treemap_sizemetric                            string `json:"sonar.core.treemap.sizemetric"`
		Sonar_cpd_crossProject                                   string `json:"sonar.cpd.cross_project"`
		Sonar_dbcleaner_branchesToKeepWhenInactive               string `json:"sonar.dbcleaner.branchesToKeepWhenInactive"`
		Sonar_dbcleaner_daysBeforeDeletingInactiveBranchesAndPRs string `json:"sonar.dbcleaner.daysBeforeDeletingInactiveBranchesAndPRs"`
		Sonar_dbcleaner_monthsBeforeDeletingAllSnapshots         string `json:"sonar.dbcleaner.monthsBeforeDeletingAllSnapshots"`
		Sonar_dbcleaner_weeksBeforeDeletingAllSnapshots          string `json:"sonar.dbcleaner.weeksBeforeDeletingAllSnapshots"`
		Sonar_dryRun_cache_lastUpdate                            string `json:"sonar.dryRun.cache.lastUpdate"`
		Sonar_es_port                                            string `json:"sonar.es.port"`
		Sonar_forceAuthentication                                string `json:"sonar.forceAuthentication"`
		Sonar_genericcoverage_suffixes                           string `json:"sonar.genericcoverage.suffixes"`
		Sonar_governance_report_view_frequency                   string `json:"sonar.governance.report.view.frequency"`
		Sonar_java_coveragePlugin                                string `json:"sonar.java.coveragePlugin"`
		Sonar_javascript_jQueryObjectAliases                     string `json:"sonar.javascript.jQueryObjectAliases"`
		Sonar_jdbc_driverPath                                    string `json:"sonar.jdbc.driverPath"`
		Sonar_jdbc_maxActive                                     string `json:"sonar.jdbc.maxActive"`
		Sonar_jdbc_maxWait                                       string `json:"sonar.jdbc.maxWait"`
		Sonar_jdbc_minIdle                                       string `json:"sonar.jdbc.minIdle"`
		Sonar_jdbc_password                                      string `json:"sonar.jdbc.password"`
		Sonar_jdbc_url                                           string `json:"sonar.jdbc.url"`
		Sonar_jdbc_username                                      string `json:"sonar.jdbc.username"`
		Sonar_lf_enableGravatar                                  string `json:"sonar.lf.enableGravatar"`
		Sonar_lf_logoWidthPx                                     string `json:"sonar.lf.logoWidthPx"`
		Sonar_log_jsonOutput                                     string `json:"sonar.log.jsonOutput"`
		Sonar_organisation                                       string `json:"sonar.organisation"`
		Sonar_path_data                                          string `json:"sonar.path.data"`
		Sonar_path_home                                          string `json:"sonar.path.home"`
		Sonar_path_logs                                          string `json:"sonar.path.logs"`
		Sonar_path_temp                                          string `json:"sonar.path.temp"`
		Sonar_path_web                                           string `json:"sonar.path.web"`
		Sonar_plsql_file_suffixes                                string `json:"sonar.plsql.file.suffixes"`
		Sonar_plugins_risk_consent                               string `json:"sonar.plugins.risk.consent"`
		Sonar_preview_excludePlugins                             string `json:"sonar.preview.excludePlugins"`
		Sonar_report_dashboard_name                              string `json:"sonar.report.dashboard.name"`
		Sonar_report_frequency                                   string `json:"sonar.report.frequency"`
		Sonar_report_ignoreSslErrors                             string `json:"sonar.report.ignoreSslErrors"`
		Sonar_report_lastDate                                    string `json:"sonar.report.last_date"`
		Sonar_report_lastDate_devReport                          string `json:"sonar.report.last_date.dev_report"`
		Sonar_report_lastDate_managementReport                   string `json:"sonar.report.last_date.management_report"`
		Sonar_report_login                                       string `json:"sonar.report.login"`
		Sonar_report_subject                                     string `json:"sonar.report.subject"`
		Sonar_reports                                            string `json:"sonar.reports"`
		Sonar_scm_disabled                                       string `json:"sonar.scm.disabled"`
		Sonar_scm_enabled                                        string `json:"sonar.scm.enabled"`
		Sonar_search_host                                        string `json:"sonar.search.host"`
		Sonar_search_javaAdditionalOpts                          string `json:"sonar.search.javaAdditionalOpts"`
		Sonar_search_javaOpts                                    string `json:"sonar.search.javaOpts"`
		Sonar_search_port                                        string `json:"sonar.search.port"`
		Sonar_technicalDebt_ratingGrid                           string `json:"sonar.technicalDebt.ratingGrid"`
		Sonar_telemetry_compression                              string `json:"sonar.telemetry.compression"`
		Sonar_telemetry_enable                                   string `json:"sonar.telemetry.enable"`
		Sonar_telemetry_frequencyInSeconds                       string `json:"sonar.telemetry.frequencyInSeconds"`
		Sonar_telemetry_url                                      string `json:"sonar.telemetry.url"`
		Sonar_updatecenter_activate                              string `json:"sonar.updatecenter.activate"`
		Sonar_web_gracefulStopTimeOutInMs                        string `json:"sonar.web.gracefulStopTimeOutInMs"`
		Sonar_web_javaAdditionalOpts                             string `json:"sonar.web.javaAdditionalOpts"`
		Sonar_web_javaOpts                                       string `json:"sonar.web.javaOpts"`
		Sonar_web_sso_emailHeader                                string `json:"sonar.web.sso.emailHeader"`
		Sonar_web_sso_enable                                     string `json:"sonar.web.sso.enable"`
		Sonar_web_sso_groupsHeader                               string `json:"sonar.web.sso.groupsHeader"`
		Sonar_web_sso_loginHeader                                string `json:"sonar.web.sso.loginHeader"`
		Sonar_web_sso_nameHeader                                 string `json:"sonar.web.sso.nameHeader"`
		Sonar_web_sso_refreshIntervalInMinutes                   string `json:"sonar.web.sso.refreshIntervalInMinutes"`
	} `json:"Settings"`
	System struct {
		Accepted_external_identity_providers                                      string `json:"Accepted external identity providers"`
		Data_Dir                                                                  string `json:"Data Dir"`
		Docker                                                                    bool   `json:"Docker"`
		Edition                                                                   string `json:"Edition"`
		External_identity_providers_whose_users_are_allowed_to_sign_themselves_up string `json:"External identity providers whose users are allowed to sign themselves up"`
		Force_authentication                                                      bool   `json:"Force authentication"`
		High_Availability                                                         bool   `json:"High Availability"`
		Home_Dir                                                                  string `json:"Home Dir"`
		Official_Distribution                                                     bool   `json:"Official Distribution"`
		Processors                                                                int64  `json:"Processors"`
		Server_ID                                                                 string `json:"Server ID"`
		Temp_Dir                                                                  string `json:"Temp Dir"`
		Version                                                                   string `json:"Version"`
	} `json:"System"`
	Web_Database_Connection struct {
		Pool_Active_Connections   int64 `json:"Pool Active Connections"`
		Pool_Idle_Connections     int64 `json:"Pool Idle Connections"`
		Pool_Max_Connections      int64 `json:"Pool Max Connections"`
		Pool_Max_Lifetime__ms     int64 `json:"Pool Max Lifetime (ms)"`
		Pool_Max_Wait__ms         int64 `json:"Pool Max Wait (ms)"`
		Pool_Min_Idle_Connections int64 `json:"Pool Min Idle Connections"`
		Pool_Total_Connections    int64 `json:"Pool Total Connections"`
	} `json:"Web Database Connection"`
	Web_JVM_Properties struct {
		Awt_toolkit                                           string `json:"awt.toolkit"`
		Catalina_base                                         string `json:"catalina.base"`
		Catalina_home                                         string `json:"catalina.home"`
		Catalina_useNaming                                    string `json:"catalina.useNaming"`
		Com_redhat_fips                                       string `json:"com.redhat.fips"`
		Com_zaxxer_hikari_poolNumber                          string `json:"com.zaxxer.hikari.pool_number"`
		File_encoding                                         string `json:"file.encoding"`
		File_separator                                        string `json:"file.separator"`
		Ftp_nonProxyHosts                                     string `json:"ftp.nonProxyHosts"`
		GopherProxySet                                        string `json:"gopherProxySet"`
		HTTP_agent                                            string `json:"http.agent"`
		HTTP_nonProxyHosts                                    string `json:"http.nonProxyHosts"`
		Java_awt_graphicsenv                                  string `json:"java.awt.graphicsenv"`
		Java_awt_headless                                     string `json:"java.awt.headless"`
		Java_awt_printerjob                                   string `json:"java.awt.printerjob"`
		Java_class_path                                       string `json:"java.class.path"`
		Java_class_version                                    string `json:"java.class.version"`
		Java_home                                             string `json:"java.home"`
		Java_io_tmpdir                                        string `json:"java.io.tmpdir"`
		Java_library_path                                     string `json:"java.library.path"`
		Java_runtime_name                                     string `json:"java.runtime.name"`
		Java_runtime_version                                  string `json:"java.runtime.version"`
		Java_specification_name                               string `json:"java.specification.name"`
		Java_specification_vendor                             string `json:"java.specification.vendor"`
		Java_specification_version                            string `json:"java.specification.version"`
		Java_vendor                                           string `json:"java.vendor"`
		Java_vendor_url                                       string `json:"java.vendor.url"`
		Java_vendor_url_bug                                   string `json:"java.vendor.url.bug"`
		Java_vendor_version                                   string `json:"java.vendor.version"`
		Java_version                                          string `json:"java.version"`
		Java_version_date                                     string `json:"java.version.date"`
		Java_vm_compressedOopsMode                            string `json:"java.vm.compressedOopsMode"`
		Java_vm_info                                          string `json:"java.vm.info"`
		Java_vm_name                                          string `json:"java.vm.name"`
		Java_vm_specification_name                            string `json:"java.vm.specification.name"`
		Java_vm_specification_vendor                          string `json:"java.vm.specification.vendor"`
		Java_vm_specification_version                         string `json:"java.vm.specification.version"`
		Java_vm_vendor                                        string `json:"java.vm.vendor"`
		Java_vm_version                                       string `json:"java.vm.version"`
		Jdk_debug                                             string `json:"jdk.debug"`
		Jdk_vendor_version                                    string `json:"jdk.vendor.version"`
		Line_separator                                        string `json:"line.separator"`
		LogbackDisableServletContainerInitializer             string `json:"logbackDisableServletContainerInitializer"`
		Org_apache_catalina_startup_EXITONINITFAILURE         string `json:"org.apache.catalina.startup.EXIT_ON_INIT_FAILURE"`
		Org_apache_tomcat_util_buf_UDecoder_ALLOWENCODEDSLASH string `json:"org.apache.tomcat.util.buf.UDecoder.ALLOW_ENCODED_SLASH"`
		Os_arch                                               string `json:"os.arch"`
		Os_name                                               string `json:"os.name"`
		Os_version                                            string `json:"os.version"`
		Path_separator                                        string `json:"path.separator"`
		SocksNonProxyHosts                                    string `json:"socksNonProxyHosts"`
		Sun_arch_data_model                                   string `json:"sun.arch.data.model"`
		Sun_boot_library_path                                 string `json:"sun.boot.library.path"`
		Sun_cpu_endian                                        string `json:"sun.cpu.endian"`
		Sun_cpu_isalist                                       string `json:"sun.cpu.isalist"`
		Sun_font_fontmanager                                  string `json:"sun.font.fontmanager"`
		Sun_io_unicode_encoding                               string `json:"sun.io.unicode.encoding"`
		Sun_java_command                                      string `json:"sun.java.command"`
		Sun_java_launcher                                     string `json:"sun.java.launcher"`
		Sun_jnu_encoding                                      string `json:"sun.jnu.encoding"`
		Sun_management_compiler                               string `json:"sun.management.compiler"`
		Sun_os_patch_level                                    string `json:"sun.os.patch.level"`
		User_country                                          string `json:"user.country"`
		User_dir                                              string `json:"user.dir"`
		User_home                                             string `json:"user.home"`
		User_language                                         string `json:"user.language"`
		User_name                                             string `json:"user.name"`
		User_timezone                                         string `json:"user.timezone"`
	} `json:"Web JVM Properties"`
	Web_JVM_State struct {
		Free_Memory__MB        int64  `json:"Free Memory (MB)"`
		Heap_Committed__MB     int64  `json:"Heap Committed (MB)"`
		Heap_Init__MB          int64  `json:"Heap Init (MB)"`
		Heap_Max__MB           int64  `json:"Heap Max (MB)"`
		Heap_Used__MB          int64  `json:"Heap Used (MB)"`
		Max_Memory__MB         int64  `json:"Max Memory (MB)"`
		Non_Heap_Committed__MB int64  `json:"Non Heap Committed (MB)"`
		Non_Heap_Init__MB      int64  `json:"Non Heap Init (MB)"`
		Non_Heap_Used__MB      int64  `json:"Non Heap Used (MB)"`
		System_Load_Average    string `json:"System Load Average"`
		Threads                int64  `json:"Threads"`
	} `json:"Web JVM State"`
	Web_Logging struct {
		Logs_Dir   string `json:"Logs Dir"`
		Logs_Level string `json:"Logs Level"`
	} `json:"Web Logging"`
}
